package internal

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/google/uuid"
	"github.com/yosssi/gohtml"
	"html/template"
)

type Attrs struct {
	Props    map[string]any
	Children map[string]any
	Payload  map[string]any
}

func Render(t string, data *Attrs) (string, error) {
	attrs := make(map[string]any)

	attrs["props"] = data.Props
	attrs["payload"] = data.Payload

	for k, v := range data.Children {
		v := v

		attrs[k] = v
	}

	temp := template.Must(template.New(fmt.Sprintf("temp-%s", uuid.New().String())).Parse(t))

	buf := new(bytes.Buffer)

	err := temp.Execute(gohtml.NewWriter(buf), attrs)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
