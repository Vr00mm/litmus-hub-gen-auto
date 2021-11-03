package workflow

import (
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
	"hub-gen-auto/pkg/utils"
)

func Generate(composants *resources.Resources, experimentsManifests map[string]types.ChaosChart) []types.WorkflowChart {

	var workflows []types.WorkflowChart

	for _, composant := range composants.Objects {

		var workflowArgo types.WorkflowArgow
		var workflowChartVersion types.WorkflowChartVersion

		workflowChartVersion.APIVersion = "litmuchaos.io/v1alpha1"
		workflowChartVersion.Kind = "ChartServiceVersion"
		//workflowChartVersion.Metadata.CreatedAt = ""
		workflowChartVersion.Metadata.Name = composants.Namespace + "-" + composant.GetName()
		workflowChartVersion.Metadata.Version = "0.1.0"
		workflowChartVersion.Metadata.Annotations.Categories = composant.GetName()
		workflowChartVersion.Metadata.Annotations.ChartDescription = "Injects chaos on " + composant.GetName() + " component."
		workflowChartVersion.Spec.DisplayName = composants.Namespace + "-" + composant.GetName()
		workflowChartVersion.Spec.CategoryDescription = ""
		workflowChartVersion.Spec.Provider.Name = "auto-hub-gen"
		workflowChartVersion.Spec.Links = append(workflowChartVersion.Spec.Links)
		workflowChartVersion.Spec.Icon = append(workflowChartVersion.Spec.Icon)

		workflowArgo.APIVersion = "argoproj.io/v1alpha1"
		workflowArgo.Kind = "Workflow"
		workflowArgo.Metadata.GenerateName = composants.Namespace + "-" + composant.GetName() + "-"
		workflowArgo.Metadata.Namespace = composants.Namespace
		workflowArgo.Metadata.Labels.Subject = composants.Namespace + "-" + composant.GetName()
		workflowArgo.Spec.Entrypoint = "argowf-chaos"
		workflowArgo.Spec.ServiceAccountName = "argo-chaos"
		workflowArgo.Spec.SecurityContext.RunAsNonRoot = true
		workflowArgo.Spec.SecurityContext.RunAsUser = 1000
		workflowArgo.Spec.Arguments.Parameters = append(workflowArgo.Spec.Arguments.Parameters, struct {
			Name  string "yaml:\"name\""
			Value string "yaml:\"value\""
		}{"adminModeNamespace", "litmus"})

		var workflowSteps [][]struct {
			Name     string `yaml:"name"`
			Template string `yaml:"template"`
		}
		var stepInstall []struct {
			Name     string `yaml:"name"`
			Template string `yaml:"template"`
		}

		stepInstall = append(stepInstall, struct {
			Name     string "yaml:\"name\""
			Template string "yaml:\"template\""
		}{"install-chaos-experiments", "install-chaos-experiments"})
		workflowSteps = append(workflowSteps, stepInstall)

		for _, experiment := range composant.GeneratedExperiments {

			var stepExperiment []struct {
				Name     string `yaml:"name"`
				Template string `yaml:"template"`
			}
			stepExperiment = append(stepExperiment, struct {
				Name     string "yaml:\"name\""
				Template string "yaml:\"template\""
			}{experiment, experiment})

			workflowSteps = append(workflowSteps, stepExperiment)

			workflowChartVersion.Spec.Experiments = append(workflowChartVersion.Spec.Experiments, experiment)
			for _, keyword := range experimentsManifests[experiment].ChartVersion.Spec.Keywords {
				workflowChartVersion.Spec.Keywords = append(workflowChartVersion.Spec.Keywords, keyword)
			}
			for _, platform := range experimentsManifests[experiment].ChartVersion.Spec.Platforms {
				workflowChartVersion.Spec.Platforms = append(workflowChartVersion.Spec.Platforms, platform)
			}
			for _, link := range experimentsManifests[experiment].ChartVersion.Spec.Links {
				workflowChartVersion.Spec.Links = append(workflowChartVersion.Spec.Links, link)
			}
			for _, maintener := range experimentsManifests[experiment].ChartVersion.Spec.Maintainers {
				workflowChartVersion.Spec.Maintainers = append(workflowChartVersion.Spec.Maintainers, struct {
					Name  string "yaml:\"name\""
					Email string "yaml:\"email\""
				}{Name: maintener.Name, Email: maintener.Email})
			}
		}
		relou := &types.WorkflowTemplate{Name: "argowf-chaos", Steps: workflowSteps}
		workflowArgo.Spec.Templates = append(workflowArgo.Spec.Templates, relou)

		workflowChartVersion.Spec.Keywords = utils.RemoveDuplicateStr(workflowChartVersion.Spec.Keywords)
		workflowChartVersion.Spec.Platforms = utils.RemoveDuplicateStr(workflowChartVersion.Spec.Platforms)
		workflowChartVersion.Spec.Maintainers = utils.RemoveDuplicateMap(workflowChartVersion.Spec.Maintainers)

		var workflow types.WorkflowChart
		workflow.WorkflowArgow = workflowArgo
		workflow.WorkflowChartVersion = workflowChartVersion
		workflows = append(workflows, workflow)

	}
	return workflows
}
