//go:build js && wasm

package client

import (
	"github.com/syke99/oasis/client/console"
	"syscall/js"
)

// FuncMap is a map of functions that you
// want available for being called inside
// of HTML
type FuncMap map[string]func(args ...js.Value) interface{}

// Oasis is used for configuring the client's
// available WASM functionality
type Oasis struct {
	funcs FuncMap
}

func NewOasis() *Oasis {
	return &Oasis{}
}

// AddToFuncMap adds a function with the given name
// to o. If o doesn't have any functions added yet,
// it will create an underlying map to hold the
// newly added function
func (o *Oasis) AddToFuncMap(name string, fn func(args ...js.Value) interface{}) {
	if o.funcs == nil {
		o.funcs = make(FuncMap)
	}

	o.funcs[name] = fn
}

// AddFuncMap allows you to add an entire FuncMap
// to o so that multiple functions can be added
// at once
func (o *Oasis) AddFuncMap(fmap FuncMap) {
	if o.funcs == nil {
		o.funcs = fmap
		return
	}
	for k, v := range fmap {
		o.funcs[k] = v
	}
}

// Run should be called after your Oasis
// has been configured. This will register
// all funcs in the FuncMap held by the
// Oasis and make them available to be used
// as functions for values of attributes
// of HTML elements. It will log an error
// if no functions were added to the FuncMap
// before being called and then exit
func (o *Oasis) Run() {
	if o.funcs != nil {
		for k, v := range o.funcs {
			js.Global().Set(k, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				return v(args...)
			}))
		}
	} else {
		console.ErrMessage("attempted to run oasis app without funcmap; shutting down", nil)
	}
	<-make(chan struct{})
}
