//go:build js && wasm

package dom

import "syscall/js"

type Win struct {
	window js.Value
}

var Window = &Win{
	window: js.Global().Get("window"),
}

func (w *Win) GetElementById(id string) Element {
	return &element{
		elem: w.window.Call("getElementById", id),
	}
}

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

func (w *Win) On(name OnEvent, fn func(args ...js.Value) interface{}) {
	w.window.Call("addEventListener", string(name), js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fn(args...)
	}))
}
