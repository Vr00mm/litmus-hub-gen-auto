package workflow

import (
	"hub-gen-auto/pkg/types"
)

func generate() {
	var workflows []types.Workflow

	for _, composant := range composants.Objects {

		var workflow types.WorkflowArgow
		var workflowChartVersion types.WorkflowChartVersion

		workflowChartVersion.APIVersion = "litmuchaos.io/v1alpha1"
		workflowChartVersion.Kind = "ChartServiceVersion"
		workflowChartVersion.Metadata.CreatedAt = ""
		workflowChartVersion.Metadata.Name = composant.Namespace + "-" + composant.Name
		workflowChartVersion.Metadata.Version = "0.1.0"
		workflowChartVersion.Metadata.Annotations.Categories = composant.Name
		workflowChartVersion.Metadata.Annotations.ChartDescription = "Injects chaos on " + composant.Name + " component."
		workflowChartVersion.Spec.DisplayName = composant.Namespace + "-" + composant.Name
		workflowChartVersion.Spec.CategoryDescription = ""
		workflowChartVersion.Spec.Provider.Name = "auto-hub-gen"
		workflowChartVersion.Spec.Links = append(workflowChartVersion.Spec.Links)
		workflowChartVersion.Spec.Icon = append(workflowChartVersion.Spec.Icon)

		workflow.APIVersion = "argoproj.io/v1alpha1"
		workflow.Kind = "Workflow"
		workflow.Metadata.GenerateName = composant.Namespace + "-" + composant.Name + "-"
		workflow.Metadata.Namespace = composant.Namespace
		workflow.Metadata.Labels.Subject = composant.Namespace + "-" + composant.Name
		workflow.Spec.Entrypoint = "argowf-chaos"
		workflow.Spec.ServiceAccountName = "argo-chaos"
		workflow.Spec.SecurityContext.RunAsNonRoot = true
		workflow.Spec.SecurityContext.RunAsUser = 1000
		workflow.Spec.Arguments.Parameters = append(workflow.Spec.Arguments.Parameters, struct {
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

		for _, experiment := range composant.Experiments {
			var stepExperiment []struct {
				Name     string `yaml:"name"`
				Template string `yaml:"template"`
			}
			stepExperiment = append(stepExperiment, struct {
				Name     string "yaml:\"name\""
				Template string "yaml:\"template\""
			}{experiment.Name, experiment.Name})
			workflowSteps = append(workflowSteps, stepExperiment)

			workflowChartVersion.Spec.Experiments = append(workflowChartVersion.Spec.Experiments, experiment.ChaosExperiment.Name)
			for _, keyword := range experiment.ChartVersion.Spec.Keywords {
				workflowChartVersion.Spec.Keywords = append(workflowChartVersion.Spec.Keywords, keyword)
			}
			for _, platform := range experiment.ChartVersion.Spec.Platforms {
				workflowChartVersion.Spec.Platforms = append(workflowChartVersion.Spec.Platforms, platform)
			}
			for name, email := range experiment.ChartVersion.Spec.Maintainers {
				workflowChartVersion.Spec.Maintainers = append(workflowChartVersion.Spec.Maintainers, map[string]string{"name": name, "email": email})
			}
		}
		workflowChartVersion.Spec.Keywords = utils.removeDuplicateStr(workflowChartVersion.Spec.Keywords)
		workflowChartVersion.Spec.Platforms = utils.removeDuplicateStr(workflowChartVersion.Spec.Platforms)
		workflowChartVersion.Spec.Maintainers = utils.removeDuplicateMap(workflowChartVersion.Spec.Maintainers)
	}
}
