//go:build js && wasm

package dom

import "syscall/js"

// Doc is a wrapper type to provide
// methods to manipulate the global
// HTML "document" element
type Doc struct {
	doc js.Value
}

// Document is the global HTML "document" element
var Document = &Doc{
	doc: js.Global().Get("document"),
}

// GetElementById returns the child Element with
// the matching id
func (d *Doc) GetElementById(id string) Element {
	return &element{
		elem: d.doc.Call("getElementById", id),
	}
}

// GetElementsByClassName returns the child Elements with
// the matching class
func (d *Doc) GetElementsByClassName(class string) []Element {
	e := d.doc.Call("getElementsByClassName", class)

	l := e.Length()

	elems := make([]Element, l)

	for i := 0; i < l; i++ {
		elems[i] = &element{
			elem: e.Index(i),
		}
	}

	return elems
}
