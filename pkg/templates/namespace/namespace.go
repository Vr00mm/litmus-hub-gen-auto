package namespace

import (
	"bytes"
	"hub-gen-auto/pkg/types"
	gotemplate "text/template"
)

var Template string

func init() {
	Template = `# Chaos Engineering

## Exercices de chaos pour le cluster {{.ClusterName}} dans le namespace {{.Namespace}}

 {{range $composant, $experiments := .Composants -}}
 ### Exercices pour le composant {{$composant}}

| Exercice | Container | Lien | Commentaire
| ------ | ------ | ------ | ------ |
{{range _, $experiment := $experiments -}}
| {{$experiment.TemplateName}} |  | [documentation/experiments/clusters/{{$.ClusterName}}/${{.Namespace}}/{{$experiment.ChaosExperiment.Metadata.Name}}.md](PlDb) | |
{{end}}`
}

func sortByComponent(namespace types.Manifest) map[string][]types.ChaosChart{
	result := make(map[string][]types.ChaosChart)
	for _, experiment := range namespace.Experiments {
		composantName := experiment.Type + "/" + experiment.Composant
		result[composantName] = append(result[composantName], experiment)
	}
	return result
}
func Generate(data struct {
	ClusterName string
	Namespace   types.Manifest
}) string {

	t, err := gotemplate.New(data.ClusterName + "-" + data.Namespace.Namespace).Parse(Template)
	if err != nil {
		panic(err)
	}

	sortedData := sortByComponent(data.Namespace)
	var tpl bytes.Buffer
	err = t.Execute(&tpl, struct {
		ClusterName string
		Namespace   string
		Composants map[string][]types.ChaosChart
	}{data.ClusterName, data.Namespace.Namespace, sortedData})
	if err != nil {
		panic(err)
	}
	return tpl.String()
}
