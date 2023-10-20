package islands

import (
	"encoding/json"
	"github.com/syke99/oasis/internal"
)

type props struct {
	props map[string]any
}

// Island is an interface for manipulating
// templates to be returned for a request
type Island interface {
	// GetName returns the name of an Island
	GetName() string
	// GetTemplate returns the template of an Island
	GetTemplate() string
	// GetProps returns the props for a given
	// Island as a map[string]any
	GetProps() map[string]any
	AddChild(child Island)
	// AddProp adds  prop to an Island. It will then be
	// available in an Island's template via {{ .props.name }}
	AddProp(name string, prop any) Island
	// AddProps adds  props to an Island at once
	AddProps(props map[string]any) Island
	// Hydrate takes in a payload that has been unmarshalled
	// from JSON into a map[string]any so that it can be
	// accessed in an Island's template via {{ .payload.* }}
	Hydrate(payload map[string]any) Island
	// HydrateBytes is just like Hydrate, except it accepts
	// a []byte representation of a JSON object
	HydrateBytes(payload []byte) (Island, error)
	// Render renders an Island's template and returns
	// the rendered template string or an error if one
	// occurs
	Render() (string, error)
}

// MustRender attempts to render an Island
// and returns the rendered template. It
// panics if an error is encountered
func MustRender(elem Island) string {
	str, err := elem.Render()
	if err != nil {
		panic(err)
	}
	return str
}

type node struct {
	name     string
	template string
	props    *props
	children map[string]Island
	payload  map[string]any
}

func NewIsland(name string, template string) Island {
	n := &node{
		name:     name,
		template: template,
		props:    &props{props: make(map[string]any, 0)},
		children: make(map[string]Island, 0),
	}

	return n
}

// GetName returns the name of an Island
func (n *node) GetName() string {
	return n.name
}

// GetTemplate returns the template of an Island
func (n *node) GetTemplate() string {
	return n.template
}

// GetProps returns the props for a given
// Island as a map[string]any
func (n *node) GetProps() map[string]any {
	return n.props.props
}

// AddProp adds  prop to an Island. It will then be
// available in an Island's template via {{ .props.name }}
func (n *node) AddProp(name string, prop any) Island {
	n.props.props[name] = prop
	return n
}

// AddProps adds  props to an Island at once
func (n *node) AddProps(props map[string]any) Island {
	if len(n.props.props) == 0 {
		n.props.props = props
	} else {
		for k, v := range props {
			n.props.props[k] = v
		}
	}

	return n
}

func (n *node) AddChild(child Island) {
	n.children[child.GetName()] = child
}

// Hydrate takes in a payload that has been unmarshalled
// from JSON into a map[string]any so that it can be
// accessed in an Island's template via {{ .payload.* }}
func (n *node) Hydrate(payload map[string]any) Island {
	n.payload = payload
	return n
}

// HydrateBytes is just like Hydrate, except it accepts
// a []byte representation of a JSON object
func (n *node) HydrateBytes(payload []byte) (Island, error) {
	p := make(map[string]any)

	err := json.Unmarshal(payload, &p)
	if err != nil {
		return nil, err
	}

	n.payload = p
	return n, nil
}

// Render renders an Island's template and returns
// the rendered template string or an error if one
// occurs
func (n *node) Render() (string, error) {
	childMap := make(map[string]string)

	for name, child := range n.children {
		renderedChild, err := child.Render()
		if err != nil {
			return "", err
		}

		childMap[name] = renderedChild
	}

	p := &internal.Attrs{
		Props:    n.props.props,
		Children: childMap,
		Payload:  n.payload,
	}
	return internal.Render(n.template, p)
}
