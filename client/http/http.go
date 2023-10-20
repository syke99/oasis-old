//go:build js && wasm

package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"syscall/js"
)

// Do preforms the given *http.Request
// and handles all asynchronicity to prevent
// an Oasis from crashing due to blocking calls,
// then returns the *http.Response or error; this
// provides an idiomatic way of preforming HTTP
// requests
func Do(req *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error

	// wrap HTTP request in try catch so is non-blocking
	httpRequestAsHandler(req).
		Invoke().
		Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			res = goHttpResponse(args[0])
			return nil
		})).
		Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			err = errors.New(args[0].Call("toString").String())
			return nil
		}))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func goHttpResponse(response js.Value) *http.Response {
	return &http.Response{
		Status:     response.Get("statusText").String(),
		StatusCode: response.Get("status").Int(),
		Header:     httpResponseHeaders(response),
		Body:       io.NopCloser(bytes.NewBuffer(httpResponseBody(response))),
	}
}

func httpResponseHeaders(response js.Value) http.Header {
	headers := make(http.Header)

	response.
		Get("headers").
		Call("forEach", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			key := args[0].String()
			val := []string{args[1].String()}

			headers[key] = val
			return nil
		}))

	return headers
}

func httpResponseBody(response js.Value) []byte {
	var body []byte

	response.
		Call("arrayBuffer").
		Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			l := this.Length()

			body = make([]byte, l)

			js.CopyBytesToGo(body, this)
			return nil
		}))

	return body
}

func httpRequestAsHandler(req *http.Request) js.Value {
	// Handler for the Promise
	// We need to return a Promise because HTTP requests are blocking in Go
	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		// Run this code asynchronously
		go func() {
			// Make the HTTP request
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				// Handle errors: reject the Promise if we have an error
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New(err.Error())
				reject.Invoke(errorObject)
				return
			}
			defer res.Body.Close()

			// Read the response body
			data, err := io.ReadAll(res.Body)
			if err != nil {
				// Handle errors here too
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New(err.Error())
				reject.Invoke(errorObject)
				return
			}

			// "data" is a byte slice, so we need to convert it to a JS Uint8Array object
			arrayConstructor := js.Global().Get("Uint8Array")
			dataJS := arrayConstructor.New(len(data))
			js.CopyBytesToJS(dataJS, data)

			// Create a Response object and pass the data
			responseConstructor := js.Global().Get("Response")
			response := responseConstructor.New(dataJS)

			// Resolve the Promise
			resolve.Invoke(response)
		}()

		// The handler of a Promise doesn't return any value
		return nil
	})

	return js.ValueOf(handler)
}
