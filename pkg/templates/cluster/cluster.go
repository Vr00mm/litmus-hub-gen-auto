package cluster

import (
	"bytes"
	"hub-gen-auto/pkg/types"
	gotemplate "text/template"
)

var Template string

func init() {
	Template = `# Chaos Engineering

## Exercices de chaos pour le cluster {{.ClusterName}}

### Exercices par namespace

| Namespace | Lien |
| ------ | ------ |
{{range $namespace := .Namespaces -}}
| {{$namespace.Namespace}} | [documentation/clusters/{{$.ClusterName}}/{{$namespace.Namespace}}.md](documentation/clusters/{{$.ClusterName}}/{{$namespace.Namespace}}.md) |
{{end}}`
}

func Generate(data struct {
	ClusterName string
	Namespaces  []types.Manifest
}) string {
	t, err := gotemplate.New(data.ClusterName).Parse(Template)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		panic(err)
	}
	return tpl.String()
}
