//go:build js && wasm

package dom

import "syscall/js"

func Document() Element {
	return &element{
		elem: js.Global().Get("document"),
	}
}

type Element interface {
	GetElementById(id string) Element
	GetInnerHTML() string
	SetInnerHTML(inner string)
	GetOuterHTML() string
	SetOuterHTML(outer string)
}

type element struct {
	elem js.Value
}

func (e *element) GetElementById(id string) Element {
	return *element{
		elem: e.elem.Call("getElementById", id),
	}
}

func (e *element) GetInnerHTML() string {
	return e.elem.Get("innerHTML").String()
}

func (e *element) SetInnerHTML(inner string) {
	e.elem.Set("innerHTML", inner)
}

func (e *element) GetOuterHTML() string {
	return e.elem.Get("outerHTML").String()
}

func (e *element) SetOuterHTML(outer string) {
	e.elem.Set("outerHTML", inner)
}
