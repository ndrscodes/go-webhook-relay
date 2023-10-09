package rendering

import (
	"bytes"
	"log"
	"text/template"
	"webhook-proxy/rendering/resolvers"
)

type GoTemplateRenderer struct {
}

var tplCache map[string]*template.Template = map[string]*template.Template{}

func findTemplate(path string) (tpl *template.Template, err error) {
	tpl, ok := tplCache[path]
	if !ok {
		tpl, err = template.ParseFiles(path)
		if err != nil {
			return nil, err
		}
		tplCache[path] = tpl
		log.Default().Printf("cached template %v at %s", tpl, path)
	}
	return tpl, nil
}

func execute(tpl *template.Template, v any) (out *string, err error) {
	buffer := bytes.Buffer{}
	err = tpl.Execute(&buffer, v)
	if err != nil {
		return nil, err
	}
	res := buffer.String()
	return &res, nil
}

func (GoTemplateRenderer) Render(payload *map[string]any, resolver resolvers.Resolver) (text *string, err error) {
	path := resolver.Resolve(payload)

	tpl, err := findTemplate(path)
	if err != nil {
		return nil, err
	}

	return execute(tpl, payload)
}
