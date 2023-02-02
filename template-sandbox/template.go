package template

import (
	_ "embed"
	"text/template"
)

type htmlTemplateContent string

//go:embed tmpls/components.html.gotmpl
var componentsTemplateContent htmlTemplateContent
var components = template.Must(template.New("components").Parse(string(componentsTemplateContent)))

var NameAge = template.Must(components.Clone())
var _ = template.Must(NameAge.Parse(`{{ template "name" .Name }}
{{ template "age" .Age }}
`))

var AgeName = template.Must(components.Clone())
var _ = template.Must(AgeName.Parse(`{{ template "age" .Age }}
{{ template "name" .Name}}
`))
