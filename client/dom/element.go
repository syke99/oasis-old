//go:build js && wasm

package dom

import (
	"strings"
	"syscall/js"
)

// Element is an interface for interacting
// with specific HTML elements
type Element interface {
	// AddEventListener allows you to add a custom event listener
	// to an Element with the given name
	AddEventListener(name string, fn func(args ...js.Value) interface{})
	// On adds an event listener to a specific "on" event
	On(name OnEvent, fn func(args ...js.Value) interface{})
	// RemoveEventListener removes the event listener from an Element
	// with the given name. fn must match the same event listener that
	// was set for the given name whenever AddEventListener was called
	RemoveEventListener(name string, fn func(args ...js.Value) interface{}, isCapture bool)
	// GetElementById returns the child Element with
	// the matching id
	GetElementById(id string) Element
	// GetElementsByClassName returns the child Elements with
	// the matching class
	GetElementsByClassName(class string) []Element
	// GetInnerHTML returns the inner HTML content of
	// the Element it was called on
	GetInnerHTML() string
	// SetInnerHTML sets the inner HTML content of
	// the Element it was called on
	SetInnerHTML(inner string)
	// GetOuterHTML returns the outer HTML content of
	// the Element it was called on
	GetOuterHTML() string
	// SetOuterHTML sets the outer HTML content of
	// the Element it was called on
	SetOuterHTML(outer string)
	// GetAttribute returns the given value
	// for the specified attr of the Element
	// it was called on
	GetAttribute(attr string) js.Value
	// SetAttribute sets the given value
	// for the specified attr of the Element
	// it was called on
	SetAttribute(attr string, val string)
	// GetId returns the ID of the Element
	// it was called on as a string
	GetId() string
	// GetClasses returns a slice of all
	// classes for a given Element
	GetClasses() []string
}

type element struct {
	elem js.Value
}

// AddEventListener allows you to add a custom event listener
// to an Element with the given name
func (e *element) AddEventListener(name string, fn func(args ...js.Value) interface{}) {
	e.elem.Call("addEventListener", name, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fn(args...)
	}))
}

// On adds an event listener to a specific "on" event
func (e *element) On(name OnEvent, fn func(args ...js.Value) interface{}) {
	e.elem.Call("addEventListener", string(name), js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fn(args...)
	}))
}

// RemoveEventListener removes the event listener from an Element
// with the given name. fn must match the same event listener that
// was set for the given name whenever AddEventListener was called
func (e *element) RemoveEventListener(name string, fn func(args ...js.Value) interface{}, isCapture bool) {
	e.elem.Call("removeEventListener", name, js.FuncOf(func(this js.Value, args []js.Value) any {
		return fn(args...)
	}), isCapture)
}

// GetElementById returns the child Element with
// the matching id
func (e *element) GetElementById(id string) Element {
	return &element{
		elem: e.elem.Call("getElementById", id),
	}
}

// GetElementsByClassName returns the child Elements with
// the matching class
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

// GetInnerHTML returns the inner HTML content of
// the Element it was called on
func (e *element) GetInnerHTML() string {
	return e.elem.Get("innerHTML").String()
}

// SetInnerHTML sets the inner HTML content of
// the Element it was called on
func (e *element) SetInnerHTML(inner string) {
	e.elem.Set("innerHTML", inner)
}

// GetOuterHTML returns the outer HTML content of
// the Element it was called on
func (e *element) GetOuterHTML() string {
	return e.elem.Get("outerHTML").String()
}

// SetOuterHTML sets the outer HTML content of
// the Element it was called on
func (e *element) SetOuterHTML(outer string) {
	e.elem.Set("outerHTML", outer)
}

// GetAttribute returns the given value
// for the specified attr of the Element
// it was called on
func (e *element) GetAttribute(attr string) js.Value {
	return e.elem.Get(attr)
}

// SetAttribute sets the given value
// for the specified attr of the Element
// it was called on
func (e *element) SetAttribute(attr string, val string) {
	e.elem.Set(attr, val)
}

// GetId returns the ID of the Element
// it was called on as a string
func (e *element) GetId() string {
	return e.elem.Get("id").String()
}

// GetClasses returns a slice of all
// classes for a given Element
func (e *element) GetClasses() []string {
	classes := e.elem.Get("class").String()

	return strings.Split(classes, "")
}
