//go:build js && wasm

package dom

import "syscall/js"

type Doc struct {
	doc js.Value
}

func Document() *Doc {
	return &Doc{
		doc: js.Global().Get("document"),
	}
}

func (d *Doc) GetElementById(id string) Element {
	return *element{
		elem: e.elem.Call("getElementById", id),
	}
}

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
