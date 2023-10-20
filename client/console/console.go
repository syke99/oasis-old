//go:build js && wasm

package console

import (
	"syscall/js"
)

var console = js.Global().Get("console")

// Assert calls console.assert() with the given assertion, any args
// to be passed to the assertion, and then any args to be passed to
// console.assert
func Assert(assertion func(...any) bool, astArgs []any, args ...any) {
	a := make([]js.Value, len(args))

	for i := range args {
		a[i] = js.ValueOf(args[i])
	}

	fn := js.FuncOf(func(this js.Value, args []js.Value) any {
		return assertion(astArgs...)
	})

	all := make([]any, len(a)+1)

	all[0] = fn

	for i := range a {
		all[i+1] = a[i]
	}

	console.Call("assert", all...)
}

// Clear calls console.clear()
func Clear() {
	console.Call("clear")
}

// Count calls console.count();
// if label is not empty, it calls
// console.count(label)
func Count(label string) {
	switch label == "" {
	case true:
		console.Call("count")
	case false:
		console.Call("count", label)
	}
}

// CountReset calls console.countReset();
// if label is not empty, it calls
// console.countReset(label)
func CountReset(label string) {
	switch label == "" {
	case true:
		console.Call("countReset", label)
	case false:
		console.Call("countReset")
	}
}

// DebugObject calls console.debug(objs...)
func DebugObject(objs ...any) {
	console.Call("debug", objs...)
}

// DebugMessage calls console.debug(msg, subStr...)
func DebugMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	all := make([]any, len(s)+1)

	all[0] = msg

	for i := range s {
		all[i+1] = s[i]
	}

	console.Call("debug", all...)
}

// Dir calls console.dir(obj)
func Dir(obj any) {
	console.Call("dir", js.ValueOf(obj))
}

// DirXML calls console.dirxml(obj)
func DirXML(obj any) {
	console.Call("dirxml", js.ValueOf(obj))
}

// ErrObject calls console.error(objs...)
func ErrObject(objs ...any) {
	console.Call("error", objs...)
}

// ErrMessage calls console.error(msg, subStr...)
func ErrMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	all := make([]any, len(s)+1)

	all[0] = msg

	for i := range s {
		all[i+1] = s[i]
	}

	console.Call("error", all...)
}

// Group calls console.group();
// if label is not empty, it calls
// console.group(label)
func Group(label string) {
	switch label == "" {
	case true:
		console.Call("group")
	case false:
		console.Call("group", label)
	}
}

// GroupCollapsed calls console.groupCollapsed();
// if label is not empty, it calls
// console.groupCollapsed(label)
func GroupCollapsed(label string) {
	switch label == "" {
	case true:
		console.Call("groupCollapsed")
	case false:
		console.Call("groupCollapsed", label)
	}
}

// GroupEnd calls console.groupEnd()
func GroupEnd() {
	console.Call("groupEnd")
}

// InfoObject calls console.info(objs...)
func InfoObject(objs ...any) {
	console.Call("info", objs...)
}

// InfoMessage calls console.info(msg, subStr...)
func InfoMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	all := make([]any, len(s)+1)

	all[0] = msg

	for i := range s {
		all[i+1] = s[i]
	}

	console.Call("info", all...)
}

// LogObject calls console.log(objs...)
func LogObject(objs ...any) {
	console.Call("log", objs...)
}

// LogMessage calls console.log(msg, subStr...)
func LogMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	all := make([]any, len(s)+1)

	all[0] = msg

	for i := range s {
		all[i+1] = s[i]
	}

	console.Call("log", all...)
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

// Time calls console.time();
// if label is not empty, it calls
// console.time(label)
func Time(label string) {
	switch label == "" {
	case true:
		console.Call("time")
	case false:
		console.Call("time", label)
	}
}

// TimeEnd calls console.timeEnd();
// if label is not empty, it calls
// console.timeEnd(label)
func TimeEnd(label string) {
	switch label == "" {
	case true:
		console.Call("timeEnd")
	case false:
		console.Call("timeEnd", label)
	}
}

// TimeLog calls console.timeLog(vals...);
// if label is not empty, it calls
// console.timeLog(label, vals...)
func TimeLog(label string, vals ...any) {

	switch label == "" {
	case true:
		console.Call("timeLog", vals...)
	case false:
		v := make([]any, len(vals)+1)

		v[0] = label

		for i := range vals {
			v[i+1] = vals[i+1]
		}
		console.Call("timeLog", v...)
	}
}

// Trace calls console.trace(objs...)
func Trace(objs ...any) {
	console.Call("trace", objs...)
}

// WarnObject calls console.warn(objs...)
func WarnObject(objs ...any) {
	console.Call("warn", objs...)
}

// WarnMessage calls console.warn(msg, subStr...)
func WarnMessage(msg string, subStr []string) {
	s := make([]js.Value, len(subStr))

	for i := range subStr {
		s[i] = js.ValueOf(subStr[i])
	}

	all := make([]any, len(s)+1)

	all[0] = msg

	for i := range s {
		all[i+1] = s[i]
	}

	console.Call("warn", all...)
}
