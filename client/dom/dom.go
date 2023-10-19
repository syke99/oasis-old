//go:build js && wasm

package dom

import "syscall/js"

type Element interface {
	GetElementById(id string) Element
	GetElementsByClassName(class string) []Element
	GetInnerHTML() string
	SetInnerHTML(inner string)
	GetOuterHTML() string
	SetOuterHTML(outer string)
	SetAttribute(attr string, val string)
	GetId() string
}

type element struct {
	elem js.Value
}

func (e *element) GetElementById(id string) Element {
	return *element{
		elem: e.elem.Call("getElementById", id),
	}
}

func (e *element) GetElementsByClassName(class string) []Element {
	el := e.elem.Call("getElementsByClassName", class)

	l := el.Length()

	elems := make([]Element, l)

	for i := 0; i < l; i++ {
		elems[i] = &element{
			elem: el.Index(i),
		}
	}

	return elems
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

func (e *element) SetAttribute(attr string, val string) {
	e.elem.Set(attr, val)
}

func (e *element) GetId() string {
	return e.elem.Get("id").String()
}
