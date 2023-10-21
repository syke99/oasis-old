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
	// AfterElements inserts a set of
	// Elements from texts inside the
	// innerHTML of the Element that is
	// the parent of this Element, right
	// after this Element
	AfterElements(elements ...Element)
	// AfterText inserts a set of HTML
	// Text nodes from texts inside the
	// innerHTML of the Element that is
	// the parent of this Element, right
	// after this Element
	AfterText(texts ...string)
	// AddEventListener allows you to add a custom event listener
	// to an Element with the given name
	AddEventListener(name string, fn func(args ...js.Value) interface{})
	// Append allows you to append a
	// child Element to a parent Element's
	// children Elements
	Append(elm Element)
	// BeforeElements inserts a set of
	// Elements from texts inside the
	// innerHTML of the Element that is
	// the parent of this Element, right
	// before this Element
	BeforeElements(elements ...Element)
	// BeforeText inserts a set of HTML
	// Text nodes from texts inside the
	// innerHTML of the Element that is
	// the parent of this Element, right
	// before this Element
	BeforeText(texts ...string)
	// Closest traverses the element and
	// its parents (heading toward the
	// document root) until it finds an
	// Element that matches the specified
	// set of CSS selectors
	Closest(cssSelectors ...string) Element
	// DispatchEvent dispatches the provided Event to the Element
	// it is called on
	DispatchEvent(event *Event)
	// GetAttribute returns the given value
	// for the specified attr of the Element
	// it was called on
	GetAttribute(attr string) js.Value
	// GetAssignedSlot returns a <slot>
	// Element that the Element this
	// method was called on is assigned
	// to
	GetAssignedSlot() Element
	// GetChildElementCount returns the
	// number of child Elements of this
	// Element
	GetChildElementCount() int
	// GetChildren returns a slice of
	// all child Elements of this Element
	GetChildren() []Element
	// GetClasses returns a slice of all
	// classes for a given Element
	GetClasses() []string
	// GetClassName returns this Element's
	// className property, or an empty string
	GetClassName() string
	// GetClientHeight returns this
	// Element's clientHeight property
	GetClientHeight() int
	// GetClientLeft returns this
	// Element's clientLeft property
	GetClientLeft() int
	// GetClientTop returns this
	// Element's clientTop property
	GetClientTop() int
	// GetClientWidth returns this
	// Element's clientWidth property
	GetClientWidth() int
	// GetElementById returns the child Element with
	// the matching id
	GetElementById(id string) Element
	// GetElementsByClassName returns the child Elements with
	// the matching class
	GetElementsByClassName(class string) []Element
	// GetFirstChildElement returns an element's
	// first child Element, or null if there are
	// no child elements. It does not include
	// children such as Text Nodes
	GetFirstChildElement() Element
	// GetFirstChildNode returns an element's
	// first child Node as an Element,
	// including all children, such as Text
	// Nodes
	GetFirstChildNode() Element
	// GetId returns the ID of the Element
	// it was called on as a string
	GetId() string
	// GetInnerHTML returns the inner HTML content of
	// the Element it was called on
	GetInnerHTML() string
	// GetLastChildElement returns an element's
	// last child Element, or null if there are
	// no child elements. It does not include
	// children such as Text Nodes
	GetLastChildElement() Element
	// GetLastChildNode returns an element's
	// last child Node as an Element,
	// including all children, such as Text
	// Nodes
	GetLastChildNode() Element
	// GetLocalName returns an Element's
	// localName property
	GetLocalName() string
	// GetNamespaceURI returns a pointer to
	// a string representing an Element's
	// namespaceURI property. If the returned
	// pointer is null, this Element does not
	// have a namespaceURI
	GetNamespaceURI() string
	// GetNextElementSibling returns the Element
	// immediately following the specified one
	// in its parent's children list,
	// or null if the specified Element is the
	// last one in the list
	GetNextElementSibling() Element
	// GetPrefix returns this Element's
	// prefix if it has one, or an empty
	// string
	GetPrefix() string
	// GetPreviousElementSibling returns the Element
	// immediately before the specified one
	// in its parent's children list,
	// or null if the specified Element is the
	// last one in the list
	GetPreviousElementSibling() Element
	// GetOuterHTML returns the outer HTML content of
	// the Element it was called on
	GetOuterHTML() string
	// GetScrollHeight returns this Element's
	// scrollHeight property
	GetScrollHeight() int
	// GetScrollLeft returns this Element's
	// scrollLeft property
	GetScrollLeft() int
	// GetScrollTop returns this Element's
	// scrollTop property
	GetScrollTop() int
	// GetScrollWidth returns this Element's
	// scrollWidth property
	GetScrollWidth() int
	// GetSlot returns this Element's
	// slot attribute, or an empty string
	GetSlot() string
	// GetTagName returns the Element's
	// HTML tag name
	GetTagName() string
	// HasAttribute returns a boolean
	// value indicating whether the
	// current Element has the
	// specified attribute or not
	HasAttribute(attribute string) bool
	// HasAttributes returns a boolean
	// value indicating whether the
	// current Element has any attributes
	// or not
	HasAttributes() bool
	// HasPointerCapture checks whether
	// the Element on which it is called
	// has pointer capture for the pointer
	// identified by the given pointer ID
	HasPointerCapture(pointerId js.Value) bool
	// InsertAdjacentElement inserts a given
	// Element at a given position relative
	// to the Element it is called from
	InsertAdjacentElement(position InsertPosition, elm Element)
	// InsertAdjacentHTML parses the specified
	// text as HTML or XML and inserts the
	// resulting nodes into the DOM tree at a
	// specified position relative to the
	// Element it is called from
	InsertAdjacentHTML(position InsertPosition, html string)
	// InsertAdjacentText , given a relative
	// position and a string (data), inserts
	// a new text node at the given position
	// relative to the Element it is called from
	InsertAdjacentText(position InsertPosition, data string)
	// Matches takes in a slice of CSS
	// selectors to match against and
	// compares the Element this method
	// was called on to the Element(s)
	// with matching CSS Selectors and
	// returns a bool describing whether
	// they match
	Matches(cssSelectors ...string) bool
	// On adds an event listener to a specific "on" event
	On(name OnEvent, fn func(args ...js.Value) interface{})
	// Prepend allows you to prepend a
	// child Element to a parent Element's
	// children Elements
	Prepend(elm Element)
	// QuerySelector returns the first
	// child Element of the parent
	// Element this method was called on
	// with a matching group of CSS
	// selectors
	QuerySelector(cssSelectors ...string) Element
	// QuerySelectorAll works like
	// QuerySelector, except for
	// it returns ALL child Elements
	// with matching sets of CSS
	// selectors instead of just the
	// first child Element with a
	// set of matching CSS selectors
	QuerySelectorAll(cssSelectors ...string) []Element
	// Remove removes the Element
	// from the DOM
	Remove()
	// RemoveAttribute removes the given
	// attr from the Element it was called
	// on
	RemoveAttribute(attr string)
	// RemoveEventListener removes the event listener from an Element
	// with the given name. fn must match the same event listener that
	// was set for the given name whenever AddEventListener was called
	RemoveEventListener(name string, fn func(args ...js.Value) interface{}, isCapture bool)
	// RequestPointerLock lets you asynchronously
	// ask for the pointer to be locked on the
	// given Element
	RequestPointerLock(unadjustedMovement bool)
	// ReplaceChildren replaces the children
	// Elements of the parent Element it was
	// called on with children
	ReplaceChildren(children ...Element)
	// RequestFullScreen takes in a
	// *FullScreenOpts to control
	// the transition to Full Screen
	// and requests with the given
	// opts if not already in fullscreen
	// mode
	RequestFullScreen(opts *FullScreenOpts)

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
	// SetAttribute sets the given value
	// for the specified attr of the Element
	// it was called on
	SetAttribute(attr string, val string)
	// SetInnerHTML sets the inner HTML content of
	// the Element it was called on
	SetInnerHTML(inner string)
	// SetOuterHTML sets the outer HTML content of
	// the Element it was called on
	SetOuterHTML(outer string)
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

// AfterElements inserts a set of
// Elements from texts inside the
// innerHTML of the Element that is
// the parent of this Element, right
// after this Element
func (e *element) AfterElements(elements ...Element) {
	el := make([]any, len(elements))

	for i := range elements {
		switch elements[i].(type) {
		case *element:
			el[i] = elements[i].(*element).elem
		default:
			console.ErrMessage(fmt.Sprintf("cannot insert element %s before element %s", elements[i].(*element).GetTagName(), e.GetTagName()), nil)
			return
		}
	}

	e.elem.Call("after", el...)
}

// AfterText inserts a set of HTML
// Text nodes from texts inside the
// innerHTML of the Element that is
// the parent of this Element, right
// after this Element
func (e *element) AfterText(texts ...string) {
	t := make([]any, len(texts))

	for i := range texts {
		t[i] = texts[i]
	}

	e.elem.Call("after", t...)
}

// AddEventListener allows you to add a custom event listener
// to an Element with the given name
func (e *element) AddEventListener(name string, fn func(args ...js.Value) interface{}) {
	e.elem.Call("addEventListener", name, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fn(args...)
	}))
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

// BeforeElements inserts a set of
// Elements from texts inside the
// innerHTML of the Element that is
// the parent of this Element, right
// before this Element
func (e *element) BeforeElements(elements ...Element) {
	el := make([]any, len(elements))

	for i := range elements {
		switch elements[i].(type) {
		case *element:
			el[i] = elements[i].(*element).elem
		default:
			console.ErrMessage(fmt.Sprintf("cannot insert element %s before element %s", elements[i].(*element).GetTagName(), e.GetTagName()), nil)
			return
		}
	}

	e.elem.Call("before", el...)
}

// BeforeText inserts a set of HTML
// Text nodes from texts inside the
// innerHTML of the Element that is
// the parent of this Element, right
// before this Element
func (e *element) BeforeText(texts ...string) {
	t := make([]any, len(texts))

	for i := range texts {
		t[i] = texts[i]
	}

	e.elem.Call("before", t...)
}

// Closest traverses the element and
// its parents (heading toward the
// document root) until it finds an
// Element that matches the specified
// set of CSS selectors
func (e *element) Closest(cssSelectors ...string) Element {
	selectors := make([]any, len(cssSelectors))

	for i := range cssSelectors {
		selectors[i] = cssSelectors[i]
	}

	return &element{
		elem: e.elem.Call("closest", selectors...),
	}
}

// DispatchEvent dispatches the provided Event to the Element
// it is called on
func (e *element) DispatchEvent(event *Event) {
	e.elem.Call("dispatchEvent", event.event)
}

// GetAttribute returns the given value
// for the specified attr of the Element
// it was called on
func (e *element) GetAttribute(attr string) js.Value {
	return e.elem.Get(attr)
}

func (e *element) GetAriaSort() string {
	return e.elem.Get("ariaSort").String()
}

// GetAriaValueMax returns this Element's
// ariaValueMax property, or an empty string
func (e *element) GetAriaValueMax() string {
	return e.elem.Get("ariaValueMax").String()
}

// GetAriaValueMin returns this Element's
// ariaValueMin property, or an empty string
func (e *element) GetAriaValueMin() string {
	return e.elem.Get("ariaValueMin").String()
}

// GetAriaValueNow returns this Element's
// ariaValueNow property, or an empty string
func (e *element) GetAriaValueNow() string {
	return e.elem.Get("ariaValueNow").String()
}

// GetAriaValueText returns this Element's
// ariaValueText property, or an empty string
func (e *element) GetAriaValueText() string {
	return e.elem.Get("ariaValueText").String()
}

// GetAssignedSlot returns a <slot>
// Element that the Element this
// method was called on is assigned
// to
func (e *element) GetAssignedSlot() Element {
	return &element{
		elem: e.elem.Get("assignedSlot"),
	}
}

// GetChildElementCount returns the
// number of child Elements of this
// Element
func (e *element) GetChildElementCount() int {
	return e.elem.Get("childElementCount").Int()
}

// GetChildren returns a slice of
// all child Elements of this Element
func (e *element) GetChildren() []Element {
	ch := e.elem.Get("children")

	children := make([]Element, ch.Length())

	for i := 0; i < ch.Length(); i++ {
		children[i] = &element{
			elem: ch.Index(i),
		}
	}
	return children
}

// GetClasses returns a slice of all
// classes for a given Element
func (e *element) GetClasses() []string {
	classes := e.elem.Get("class").String()

	return strings.Split(classes, "")
}

// GetClassName returns this Element's
// className property, or an empty string
func (e *element) GetClassName() string {
	return e.elem.Get("className").String()
}

// GetClientHeight returns this
// Element's clientHeight property
func (e *element) GetClientHeight() int {
	return e.elem.Get("clientHeight").Int()
}

// GetClientLeft returns this
// Element's clientLeft property
func (e *element) GetClientLeft() int {
	return e.elem.Get("clientLeft").Int()
}

// GetClientTop returns this
// Element's clientTop property
func (e *element) GetClientTop() int {
	return e.elem.Get("clientTop").Int()
}

// GetClientWidth returns this
// Element's clientWidth property
func (e *element) GetClientWidth() int {
	return e.elem.Get("clientWidth").Int()
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

// GetFirstChildElement returns an element's
// first child Element, or null if there are
// no child elements. It does not include
// children such as Text Nodes
func (e *element) GetFirstChildElement() Element {
	return &element{
		elem: e.elem.Get("firstElementChild"),
	}
}

// GetFirstChildNode returns an element's
// first child Node as an Element,
// including all children, such as Text
// Nodes
func (e *element) GetFirstChildNode() Element {
	return &element{
		elem: e.elem.Get("firstChild"),
	}
}

// GetId returns the ID of the Element
// it was called on as a string
func (e *element) GetId() string {
	return e.elem.Get("id").String()
}

// GetInnerHTML returns the inner HTML content of
// the Element it was called on
func (e *element) GetInnerHTML() string {
	return e.elem.Get("innerHTML").String()
}

// GetLastChildElement returns an element's
// last child Element, or null if there are
// no child elements. It does not include
// children such as Text Nodes
func (e *element) GetLastChildElement() Element {
	return &element{
		elem: e.elem.Get("lastElementChild"),
	}
}

// GetLastChildNode returns an element's
// last child Node as an Element,
// including all children, such as Text
// Nodes
func (e *element) GetLastChildNode() Element {
	return &element{
		elem: e.elem.Get("lastChild"),
	}
}

// GetLocalName returns an Element's
// localName property
func (e *element) GetLocalName() string {
	return e.elem.Get("localName").String()
}

// GetNamespaceURI returns a string
// representing an Element's
// namespaceURI property. If the returned
// string is empty, this Element does not
// have a namespaceURI
func (e *element) GetNamespaceURI() string {
	return e.elem.Get("namespaceURI").String()
}

// GetNextElementSibling returns the Element
// immediately following the specified one
// in its parent's children list,
// or null if the specified Element is the
// last one in the list
func (e *element) GetNextElementSibling() Element {
	el := e.elem.Get("nextElementSibling")

	if el.IsNull() {
		return nil
	}

	return &element{
		elem: el,
	}
}

// GetPrefix returns this Element's
// prefix if it has one, or an empty
// string
func (e *element) GetPrefix() string {
	return e.elem.Get("prefix").String()
}

// GetPreviousElementSibling returns the Element
// immediately before the specified one
// in its parent's children list,
// or null if the specified Element is the
// last one in the list
func (e *element) GetPreviousElementSibling() Element {
	el := e.elem.Get("previousElementSibling")

	if el.IsNull() {
		return nil
	}

	return &element{
		elem: el,
	}
}

// GetOuterHTML returns the outer HTML content of
// the Element it was called on
func (e *element) GetOuterHTML() string {
	return e.elem.Get("outerHTML").String()
}

// GetScrollHeight returns this Element's
// scrollHeight property
func (e *element) GetScrollHeight() int {
	return e.elem.Get("scrollHeight").Int()
}

// GetScrollLeft returns this Element's
// scrollLeft property
func (e *element) GetScrollLeft() int {
	return e.elem.Get("scrollLeft").Int()
}

// GetScrollTop returns this Element's
// scrollTop property
func (e *element) GetScrollTop() int {
	return e.elem.Get("scrollTop").Int()
}

// GetScrollWidth returns this Element's
// scrollWidth property
func (e *element) GetScrollWidth() int {
	return e.elem.Get("scrollWidth").Int()
}

// GetSlot returns this Element's
// slot attribute, or an empty string
func (e *element) GetSlot() string {
	return e.elem.Get("slot").String()
}

// GetTagName returns the Element's
// HTML tag name
func (e *element) GetTagName() string {
	return e.elem.Get("tagName").String()
}

// HasAttribute returns a boolean
// value indicating whether the
// current Element has the
// specified attribute or not
func (e *element) HasAttribute(attribute string) bool {
	return e.elem.Call("hasAttribute", attribute).Bool()
}

// HasAttributes returns a boolean
// value indicating whether the
// current Element has any attributes
// or not
func (e *element) HasAttributes() bool {
	return e.elem.Call("hasAttributes").Bool()
}

// HasPointerCapture checks whether
// the Element on which it is called
// has pointer capture for the pointer
// identified by the given pointer ID
func (e *element) HasPointerCapture(pointerId js.Value) bool {
	return e.elem.Call("hasPointerCapture", pointerId).Bool()
}

type InsertPosition string

var BeforeBegin InsertPosition = "beforebegin"
var AfterBegin InsertPosition = "afterbegin"
var BeforeEnd InsertPosition = "beforeend"
var AfterEnd InsertPosition = "afterend"

// InsertAdjacentElement inserts a given
// Element at a given position relative
// to the Element it is called from
func (e *element) InsertAdjacentElement(position InsertPosition, elm Element) {
	switch elm.(type) {
	case *element:
		e.elem.Call("insertAdjacentElement", position, elm.(*element).elem)
	default:
		console.ErrMessage(fmt.Sprintf("cannot insert element %s relative to parent element %s at position %s", elm.GetTagName(), e.GetTagName(), string(position)), nil)
	}
}

// InsertAdjacentHTML parses the specified
// text as HTML or XML and inserts the
// resulting nodes into the DOM tree at a
// specified position relative to the
// Element it is called from
func (e *element) InsertAdjacentHTML(position InsertPosition, html string) {
	e.elem.Call("insertAdjacentText", position, html)
}

// InsertAdjacentText , given a relative
// position and a string (data), inserts
// a new text node at the given position
// relative to the Element it is called from
func (e *element) InsertAdjacentText(position InsertPosition, data string) {
	e.elem.Call("insertAdjacentText", position, data)
}

// Matches takes in a slice of CSS
// selectors to match against and
// compares the Element this method
// was called on to the Element(s)
// with matching CSS Selectors and
// returns a bool describing whether
// they match
func (e *element) Matches(cssSelectors ...string) bool {
	selectors := make([]any, len(cssSelectors))

	for i := range cssSelectors {
		selectors[i] = cssSelectors[i]
	}

	return e.elem.Call("matches", selectors...).Bool()
}

// On adds an event listener to a specific "on" event
func (e *element) On(name OnEvent, fn func(args ...js.Value) interface{}) {
	e.elem.Call("addEventListener", string(name), js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fn(args...)
	}))
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

// QuerySelector returns the first
// child Element of the parent
// Element this method was called on
// with a matching group of CSS
// selectors
func (e *element) QuerySelector(cssSelectors ...string) Element {
	selectors := make([]any, len(cssSelectors))

	for i := range cssSelectors {
		selectors[i] = cssSelectors[i]
	}

	return &element{
		elem: e.elem.Call("querySelector", selectors...),
	}
}

// QuerySelectorAll works like
// QuerySelector, except for
// it returns ALL child Elements
// with matching sets of CSS
// selectors instead of just the
// first child Element with a
// set of matching CSS selectors
func (e *element) QuerySelectorAll(cssSelectors ...string) []Element {
	selectors := make([]any, len(cssSelectors))

	for i := range cssSelectors {
		selectors[i] = cssSelectors[i]
	}

	ch := e.elem.Call("querySelectorAll", selectors...)

	children := make([]Element, ch.Length())

	for i := 0; i < ch.Length(); i++ {
		children[i] = &element{
			elem: ch.Index(i),
		}
	}

	return children
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

// RemoveEventListener removes the event listener from an Element
// with the given name. fn must match the same event listener that
// was set for the given name whenever AddEventListener was called
func (e *element) RemoveEventListener(name string, fn func(args ...js.Value) interface{}, isCapture bool) {
	e.elem.Call("removeEventListener", name, js.FuncOf(func(this js.Value, args []js.Value) any {
		return fn(args...)
	}), isCapture)
}

type FullScreenOpts struct {
	opts map[string]any
}

func NewFullScreenOpts() *FullScreenOpts {
	return &FullScreenOpts{
		opts: make(map[string]any),
	}
}

func (f *FullScreenOpts) Show() {
	f.opts["navigationUI"] = "show"
}

func (f *FullScreenOpts) Hide() {
	f.opts["navigationUI"] = "hide"
}

func (f *FullScreenOpts) Auto() {
	f.opts["navigationUI"] = "auto"
}

// RequestFullScreen takes in a
// *FullScreenOpts to control
// the transition to Full Screen
// and requests with the given
// opts if not already in fullscreen
// mode
func (e *element) RequestFullScreen(opts *FullScreenOpts) {
	if !js.Global().Get("document").Get("fullscreenElement").Bool() {
		e.elem.Call("requestFullscreen", opts.opts).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) any {
			console.ErrObject(args[0])
			return nil
		}))
	}
}

// RequestPointerLock lets you asynchronously
// ask for the pointer to be locked on the
// given Element
func (e *element) RequestPointerLock(unadjustedMovement bool) {
	e.elem.Call("requestPointerLock", map[string]bool{
		"unadjustedMovement": unadjustedMovement,
	}).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) any {
		console.ErrObject(args[0])
		return nil
	}))
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

// SetAttribute sets the given value
// for the specified attr of the Element
// it was called on
func (e *element) SetAttribute(attr string, val string) {
	e.elem.Set(attr, val)
}

// SetInnerHTML sets the inner HTML content of
// the Element it was called on
func (e *element) SetInnerHTML(inner string) {
	e.elem.Set("innerHTML", inner)
}

// SetOuterHTML sets the outer HTML content of
// the Element it was called on
func (e *element) SetOuterHTML(outer string) {
	e.elem.Set("outerHTML", outer)
}
