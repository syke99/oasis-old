//go:build js && wasm

package dom

import "syscall/js"

// Win is a wrapper type to provide
// methods to manipulate the global
// HTML "window" element
type Win struct {
	window js.Value
}

// Window is the global HTML "window" element
var Window = &Win{
	window: js.Global().Get("window"),
}

// GetElementById returns the child Element with
// the matching id
func (w *Win) GetElementById(id string) Element {
	return &element{
		elem: w.window.Call("getElementById", id),
	}
}

// GetElementsByClassName returns the child Elements with
// the matching class
func (w *Win) GetElementsByClassName(class string) []Element {
	e := w.window.Call("getElementsByClassName", class)

	l := e.Length()

	elems := make([]Element, l)

	for i := 0; i < l; i++ {
		elems[i] = &element{
			elem: e.Index(i),
		}
	}

	return elems
}

// On adds an event listener to a specific "on" event
func (w *Win) On(name OnEvent, fn func(args ...js.Value) interface{}) {
	w.window.Call("addEventListener", string(name), js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fn(args...)
	}))
}
