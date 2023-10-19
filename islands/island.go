package islands

import (
	"encoding/json"
	"github.com/syke99/oasis/internal"
)

type props struct {
	props map[string]any
}

type Island interface {
	AddChild(child Island)
	AddProp(name string, prop any) Island
	AddProps(props map[string]any) Island
	Hydrate(payload map[string]any) Island
	HydrateBytes(payload []byte) (Island, error)
	Render() (string, error)
	GetName() string
	GetTemplate() string
}

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
	children map[string]any
	payload  map[string]any
}

func NewIsland(name string, template string) Island {
	n := &node{
		name:     name,
		template: template,
		props:    &props{props: make(map[string]any, 0)},
		children: make(map[string]any, 0),
	}

	return n
}

func (n *node) GetName() string {
	return n.name
}

func (n *node) GetTemplate() string {
	return n.template
}

func (n *node) AddProp(name string, prop any) Island {
	n.props.props[name] = prop
	return n
}

func (n *node) AddProps(props map[string]any) Island {
	if len(n.props.props) == 0 {
		n.props.props = props
	} else {
		for k, v := range props {
			v := v

			n.props.props[k] = v
		}
	}

	return n
}

func (n *node) AddChild(child Island) {
	n.children[child.GetName()] = child
}

func (n *node) Hydrate(payload map[string]any) Island {
	n.payload = payload
	return n
}

func (n *node) HydrateBytes(payload []byte) (Island, error) {
	p := make(map[string]any)

	err := json.Unmarshal(payload, &p)
	if err != nil {
		return nil, err
	}

	n.payload = p
	return n, nil
}

func (n *node) Render() (string, error) {
	p := &internal.Attrs{
		Props:    n.props.props,
		Children: n.children,
		Payload:  n.payload,
	}
	return internal.Render(n.template, p)
}
