//go:build js && wasm

package console

import "syscall/js"

func Log(log any) {
	js.Global().Call("console", js.ValueOf(log))
}
