//go:build js && wasm

package console

import "syscall/js"

var console = js.Global().Get("console")

func Assert(assertion func(...any) bool, astArgs []any, args ...any) {
	a := make([]js.Value, len(args))

	for i := range args {
		a[i] = js.ValueOf(args[i])
	}

	fn := js.FuncOf(func(this Value, args []Value) any {
		return assertion(astArgs...)
	})

	console.Call("assert", a...)
}

func Clear() {
	console.Call("clear")
}

func Count(label string) {
	switch label == "" {
	case true:
		console.Call("count")
	case false:
		console.Call("count", label)
	}
}

func CountReset(label string) {
	switch label == "" {
	case true:
		console.Call("countReset", label)
	case false:
		console.Call()
	}
}

func DebugObject(objs ...any) {
	o := make([]js.Value, len(objs))

	for i := range objs {
		o[i] = js.ValueOf(objs[i])
	}

	console.Call("debug", o...)
}

func DebugMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	console.Call("debug", msg, s...)
}

func Dir(obj any) {
	console.Call("dir", js.ValueOf(obj))
}

func DirXML(obj any) {
	console.Call("dirxml", js.ValueOf(obj))
}

func ErrObject(objs ...any) {
	o := make([]js.Value, len(objs))

	for i := range objs {
		o[i] = js.ValueOf(objs[i])
	}

	console.Call("error", o...)
}

func ErrMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	console.Call("error", msg, s...)
}

func Group(label string) {
	switch label == "" {
	case true:
		console.Call("group")
	case false:
		console.Call("group", label)
	}
}

func GroupCollapsed(label string) {
	switch label == "" {
	case true:
		console.Call("groupCollapsed")
	case false:
		console.Call("groupCollapsed", label)
	}
}

func GroupEnd() {
	console.Call("groupEnd")
}

func InfoObject(objs ...any) {
	o := make([]js.Value, len(objs))

	for i := range objs {
		o[i] = js.ValueOf(objs[i])
	}

	console.Call("info", o...)
}

func InfoMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	console.Call("info", msg, s...)
}

func LogObject(objs ...any) {
	o := make([]js.Value, len(objs))

	for i := range objs {
		o[i] = js.ValueOf(objs[i])
	}

	console.Call("log", o...)
}

func LogMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	console.Call("log", msg, s...)
}

func Table(obj any, restrictions ...any) {
	switch obj.(type) {
	case [][]any:
		ob := obj.([][]any)

		o := make([]js.Value, len(ob))

		for i := range ob {
			nOb := make([]js.Value, len(ob[i]))
			for j := range ob[i] {
				nOb[j] = js.ValueOf(ob[i][j])
			}

			o[i] = js.ValueOf(nOb)
		}

		console.Call("table", o)
		// TODO: continue implementing
	}
}

func Time(label string) {
	switch label == "" {
	case true:
		console.Call("time")
	case false:
		console.Call("time", label)
	}
}

func TimeEnd(label string) {
	switch label == "" {
	case true:
		console.Call("timeEnd")
	case false:
		console.Call("timeEnd", label)
	}
}

func TimeLog(label string, vals ...any) {
	v := make([]js.Value, len(vals))

	for i := range vals {
		v[i] = js.ValueOf(vals[i])
	}

	switch label == "" {
	case true:
		console.Call("timeLog", v...)
	case false:
		console.Call("timeLog", label, v...)
	}
}

func Trace(objs ...any) {
	o := make([]js.Value, len(objs))

	for i := range objs {
		o[i] = js.ValueOf(objs[i])
	}

	console.Call("trace", o...)
}

func WarnObject(objs ...any) {
	o := make([]js.Value, len(objs))

	for i := range objs {
		o[i] = js.ValueOf(objs[i])
	}

	console.Call("warn", o...)
}

func WarnMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	console.Call("warn", msg, s...)
}
