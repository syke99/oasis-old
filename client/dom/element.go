//go:build js && wasm

package dom

import (
	"strings"
	"syscall/js"
)

type Element interface {
	AddEventListener(name string, fn func(args ...js.Value) interface{})
	On(name OnEvent, fn func(args ...js.Value) interface{})
	RemoveEventListener(name string, fn func(args ...js.Value) interface{}, isCapture bool)
	GetElementById(id string) Element
	GetElementsByClassName(class string) []Element
	GetInnerHTML() string
	SetInnerHTML(inner string)
	GetOuterHTML() string
	SetOuterHTML(outer string)
	SetAttribute(attr string, val string)
	GetAttribute(attr string) js.Value
	GetId() string
	GetClasses() []string
}

type element struct {
	elem js.Value
}

func (e *element) AddEventListener(name string, fn func(args ...js.Value) interface{}) {
	e.elem.Call("addEventListener", name, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fn(args...)
	}))
}

func (e *element) On(name OnEvent, fn func(args ...js.Value) interface{}) {
	e.elem.Call("addEventListener", string(name), js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fn(args...)
	}))
}

func (e *element) RemoveEventListener(name string, fn func(args ...js.Value) interface{}, isCapture bool) {
	e.elem.Call("removeEventListener", name, js.FuncOf(func(this js.Value, args []js.Value) any {
		return fn(args...)
	}), isCapture)
}

func (e *element) GetElementById(id string) Element {
	return &element{
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
	e.elem.Set("outerHTML", outer)
}

func (e *element) GetAttribute(attr string) js.Value {
	return e.elem.Get(attr)
}

func (e *element) SetAttribute(attr string, val string) {
	e.elem.Set(attr, val)
}

func (e *element) GetId() string {
	return e.elem.Get("id").String()
}

func (e *element) GetClasses() []string {
	classes := e.elem.Get("class").String()

	return strings.Split(classes, "")
}
