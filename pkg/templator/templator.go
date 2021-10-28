package templator

import (
	"fmt"
	"bytes"
	"hub-gen-auto/pkg/types"
	"text/template"
	"github.com/Masterminds/sprig"

	containerKill "hub-gen-auto/pkg/experiments/container-kill"
	podKill "hub-gen-auto/pkg/experiments/pod-kill"
)

type ExperimentTemplates struct {
	Path     string
	Experiment types.LitmusExperiment
	Engine		types.LitmusEngine
	ChartVersion types.LitmusChaosChart
}

func generateChartVersion(experiment types.Experiment) types.LitmusChartVersion {
	var tmpl types.LitmusChartVersion
	tmpl.APIVersion = "litmuchaos.io/v1alpha1"
	tmpl.Kind = "ChartServiceVersion"
	tmpl.Metadata.Name = experiment.Name
	tmpl.Metadata.Version = "0.1.0"
	tmpl.Metadata.Annotations.Categories = experiment.Name
	tmpl.Metadata.Annotations.Vendor = "chaosexperiment"
	tmpl.Metadata.Annotations.Support = "litmus"


	tmpl.Spec.DisplayName = "experiment.Template"
	tmpl.Spec.CategoryDescription = ""
	tmpl.Spec.Keywords = []string{"PKS", experiment.Platform, experiment.Name, experiment.Template}]
	tmpl.Spec.Platforms = []string{"PKS", experiment.Platform}] 
	tmpl.Spec.Maturity = ""
	tmpl.Spec.ChaosType = ""
	tmpl.Spec.Maintainers = append(tmpl.Spec.Maintainers, )
	tmpl.Spec.MinKubeVersion = ""
	tmpl.Spec.Provider.Name = ""
	tmpl.Spec.Labels.AppKubernetesIoComponent = ""
	tmpl.Spec.Labels.AppKubernetesIoVersion = ""
	tmpl.Spec.Icon = ""
	tmpl.Spec.Chaosexpcrdlink= ""
	return tmpl
}

func generateExperimentTemplate(experiment types.Experiment) types.LitmusExperiment {
	var tmpl types.LitmusExperiment
	tmpl.APIVersion = "litmuchaos.io/v1alpha1"
	tmpl.Kind = "ChaosExperiment"
	tmpl.Metadata.Name = experiment.Name
	tmpl.Metadata.Labels.Name = experiment.Name
	tmpl.Metadata.Labels.AppKubernetesIoComponent = "chaosexperiment"
	tmpl.Metadata.Labels.AppKubernetesIoPartOf = "litmus"
	tmpl.Metadata.Labels.AppKubernetesIoVersion = "latest"
	tmpl.Spec.Definition.Scope = "Namespaced"
	tmpl.Spec.Definition.Permissions = experiment.Permissions
	tmpl.Spec.Definition.Image = "litmuschaos/go-runner:latest"
	tmpl.Spec.Definition.ImagePullPolicy = "Always"
	tmpl.Spec.Definition.Command = []string{"/bin/bash"}
	tmpl.Spec.Definition.Args = []string{"-c", "./experiments -name "+experiment.Template}
	tmpl.Spec.Definition.Labels.AppKubernetesIoComponent = "experiment-job"
	tmpl.Spec.Definition.Labels.AppKubernetesIoPartOf = "litmus"
	tmpl.Spec.Definition.Labels.AppKubernetesIoVersion = "latest"
	tmpl.Spec.Definition.Labels.Name = experiment.Name

	for key, value := range experiment.Args{
		tmpl.Spec.Definition.Env = append(tmpl.Spec.Definition.Env, struct{Name string "yaml:\"name\""; Value string "yaml:\"value\""}{key, value})
	}
	return tmpl
}


func (t *template.Template) process(vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}

func templateExperiment(experiment types.Experiment) ExperimentTemplates {
	var template ExperimentTemplates
	template.Experiment = generateExperimentTemplate(experiment)
	template.ChartVersion = generateChartVersion(experiment)
	return template
}

func Template(hubs []types.Manifest) []ExperimentTemplate {
	var templates []ExperimentTemplate
	for hub := range hubs {
		for experiment := range hubs[hub].Experiments {
			tpl := templateExperiment(hubs[hub].Experiments[experiment])
			templates = append(templates, tpl)
		}
	}
	return templates
}
