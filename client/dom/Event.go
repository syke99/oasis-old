//go:build js && wasm

package dom

import "syscall/js"

// EventOptions is used to configure options
// for a new Event
type EventOptions struct {
	opts map[string]js.Value
}

func NewEventOptions(bubbles bool, cancelable bool, composed bool) *EventOptions {
	return &EventOptions{
		opts: map[string]js.Value{
			"bubbles":    js.ValueOf(bubbles),
			"cancelable": js.ValueOf(cancelable),
			"composed":   js.ValueOf(composed),
		},
	}
}

// Event is a new JavaScript Event
type Event struct {
	event js.Value
}

// NewEvent creates a new *Event with the
// provided options
func NewEvent(name string, opts *EventOptions) *Event {
	eventConstructor := js.Global().Get("Event")
	return &Event{
		event: eventConstructor.New(name, opts.opts),
	}
}
