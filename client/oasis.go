//go:build js && wasm

package client

import "syscall/js"

type Oasis struct {
	funcs map[string]func(args ...js.Value)
}

func NewOasis() *Oasis {
	return &Oasis{}
}

func (o *Oasis) AddFunc(name string, fn func(args ...js.Value)) {
	if o.funcs == nil {
		o.funcs = make(map[string]func(args ...js.Value))
	}

	o.funcs[name] = fn
}

func (o *Oasis) Run() {
	if o.funcs == nil {
		return
	}

	for k, v := range o.funcs {
		js.Global().Set(k, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			v(args...)
			return nil
		}))
	}
	<-make(chan struct{})
}
