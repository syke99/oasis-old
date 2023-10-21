//go:build js && wasm

package dom

import (
	"fmt"
	"github.com/syke99/oasis/client/console"
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
	// DispatchEvent dispatches the provided Event to the Element
	// it is called on
	DispatchEvent(event *Event)
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
	// GetTagName returns the Element's
	// HTML tag name
	GetTagName() string
	// Append allows you to append a
	// child Element to a parent Element's
	// children Elements
	Append(elm Element)
	// Prepend allows you to prepend a
	// child Element to a parent Element's
	// children Elements
	Prepend(elm Element)
	// Remove removes the Element
	// from the DOM
	Remove()
	// RemoveAttribute removes the given
	// attr from the Element it was called
	// on
	RemoveAttribute(attr string)
	// ReplaceChildren replaces the children
	// Elements of the parent Element it was
	// called on with children
	ReplaceChildren(children ...Element)
	// ReplaceWith replaces the
	// Element it was called on
	// with the given elements
	ReplaceWith(elems ...Element)
	// ScrollCoordinates scrolls to the given
	// coordinates inside the element it
	// was called on
	ScrollCoordinates(x int, y int)
	// ScrollOpts functions like
	// ScrollCoordinates, except it
	// takes a *ScrollOpts to allow
	// you to specify the behavior of
	// how scrolling happens instead
	// of just taking in coordinates
	ScrollOpts(opts *ScrollOpts)
	// ScrollByCoordinates allows you
	// to scroll the screen by the given
	// number of pixels along the x and y
	// axes
	ScrollByCoordinates(x int, y int)
	// ScrollToCoordinates scrolls to the
	// pixel coordinates of x and y
	ScrollToCoordinates(x int, y int)
	// ScrollToOptions functions like
	// ScrollToCoordinates, except it
	// takes a *ScrollOpts to allow
	// you to specify the behavior of
	// how scrolling happens instead
	// of just taking in coordinates
	ScrollToOptions(opts *ScrollOpts)
}

type element struct {
	elem js.Value
}

func NewElement(name string, initFunc js.Func, onMount js.Func, onDismount js.Func) Element {
	js.Global().Call("makeComponent", name, initFunc, onMount, onDismount)

	return &element{
		elem: any(Document.GetElementsByTagName(name)).(*element).elem.Index(0),
	}
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

// DispatchEvent dispatches the provided Event to the Element
// it is called on
func (e *element) DispatchEvent(event *Event) {
	e.elem.Call("dispatchEvent", event.event)
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

// GetTagName returns the Element's
// HTML tag name
func (e *element) GetTagName() string {
	return e.elem.Get("tagName").String()
}

// Append allows you to append a
// child Element to a parent Element's
// children Elements
func (e *element) Append(elm Element) {
	switch elm.(type) {
	case *element:
		e.elem.Call("append", elm.(*element).elem)
	default:
		console.ErrMessage(fmt.Sprintf("cannot append element %+v to parent element %s", elm, e.GetTagName()), nil)
	}
}

// Prepend allows you to prepend a
// child Element to a parent Element's
// children Elements
func (e *element) Prepend(elm Element) {
	switch elm.(type) {
	case *element:
		e.elem.Call("prepend", elm.(*element).elem)
	default:
		console.ErrMessage(fmt.Sprintf("cannot prepend element %+v to parent element %s", elm, e.GetTagName()), nil)
	}
}

// Remove removes the Element
// from the DOM
func (e *element) Remove() {
	e.elem.Call("remove")
}

// RemoveAttribute removes the given
// attr from the Element it was called
// on
func (e *element) RemoveAttribute(attr string) {
	e.elem.Call("removeAttribute", attr)
}

// ReplaceChildren replaces the children
// Elements of the parent Element it was
// called on with children
func (e *element) ReplaceChildren(children ...Element) {
	for i := range children {
		child := children[i]

		switch child.(type) {
		case *element:
			e.elem.Call("replaceChildren", child.(*element).elem)
		default:
			console.ErrMessage(fmt.Sprintf("cannot replace element %s children with child element %+v", e.GetTagName(), child), nil)
		}
	}
}

// ReplaceWith replaces the
// Element it was called on
// with the given elements
func (e *element) ReplaceWith(elems ...Element) {
	for i := range elems {
		elem := elems[i]

		switch elem.(type) {
		case *element:
			e.elem.Call("replaceWith", elem.(*element).elem)
		default:
			console.ErrMessage(fmt.Sprintf("cannot replace element %s with element %+v", e.GetTagName(), elem), nil)
		}
	}
}

type ScrollOpts struct {
	opts map[string]any
}

func NewScrollOpts() *ScrollOpts {
	return &ScrollOpts{opts: make(map[string]any)}
}

func (s *ScrollOpts) Top(top int) *ScrollOpts {
	s.opts["top"] = top
	return s
}

func (s *ScrollOpts) Left(left int) *ScrollOpts {
	s.opts["left"] = left
	return s
}

func (s *ScrollOpts) BehaviorSmooth() *ScrollOpts {
	s.opts["behavior"] = "smooth"
	return s
}

func (s *ScrollOpts) BehaviorInstant() *ScrollOpts {
	s.opts["behavior"] = "instant"
	return s
}

func (s *ScrollOpts) BehaviorAuto() *ScrollOpts {
	s.opts["behavior"] = "auto"
	return s
}

// ScrollCoordinates scrolls to the given
// coordinates inside the element it
// was called on
func (e *element) ScrollCoordinates(x int, y int) {
	e.elem.Call("scroll", x, y)
}

// ScrollOpts functions like
// ScrollCoordinates, except it
// takes a *ScrollOpts to allow
// you to specify the behavior of
// how scrolling happens instead
// of just taking in coordinates
func (e *element) ScrollOpts(opts *ScrollOpts) {
	e.elem.Call("scroll", opts.opts)
}

// ScrollByCoordinates allows you
// to scroll the screen by the given
// number of pixels along the x and y
// axes
func (e *element) ScrollByCoordinates(x int, y int) {
	e.elem.Call("scrollBy", x, y)
}

// ScrollByOpts functions like
// ScrollByCoordinates, except it
// takes a *ScrollOpts to allow
// you to specify the behavior of
// how scrolling happens instead
// of just taking in coordinates
func (e *element) ScrollByOpts(opts *ScrollOpts) {
	e.elem.Call("scrollBy", opts.opts)
}

// ScrollToCoordinates scrolls to the
// pixel coordinates of x and y
func (e *element) ScrollToCoordinates(x int, y int) {
	e.elem.Call("scrollTo", x, y)
}

// ScrollToOptions functions like
// ScrollToCoordinates, except it
// takes a *ScrollOpts to allow
// you to specify the behavior of
// how scrolling happens instead
// of just taking in coordinates
func (e *element) ScrollToOptions(opts *ScrollOpts) {
	e.elem.Call("scrollTo", opts.opts)
}
