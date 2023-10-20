//go:build js && wasm

package client

import "syscall/js"

type FuncMap map[string]func(args ...js.Value)

type Oasis struct {
	funcs FuncMap
}

func NewOasis() *Oasis {
	return &Oasis{}
}

func (o *Oasis) AddToFuncMap(name string, fn func(args ...js.Value)) {
	if o.funcs == nil {
		o.funcs = make(FuncMap)
	}

	o.funcs[name] = fn
}

func (o *Oasis) AddFuncMap(fmap FuncMap) {
	if o.funcs == nil {
		o.funcs = fmap
		return
	}
	for k, v := range fmap {
		o.funcs[k] = v
	}
}

func (o *Oasis) Run() {
	if o.funcs != nil {
		for k, v := range o.funcs {
			js.Global().Set(k, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				v(args...)
				return nil
			}))
		}
	}
	<-make(chan struct{})
}
