package generator

import (
	"fmt"
	//        "encoding/json"
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
	"hub-gen-auto/pkg/utils"

	//	"strings"
	containerKill "hub-gen-auto/pkg/experiments/container-kill"
	//	podKill "hub-gen-auto/pkg/experiments/pod-kill"
)

var experimentsList = []string{"container-kill"}

func init() {
	experimentsManifests := utils.GetExperimentsManifests(experimentsList)
	containerKill.ExperimentsManifests = experimentsManifests
}

func generateWorkflows(composants *resources.Resources) []types.Workflow {
	var workflows
	var workflowChartVersion types.WorkflowChartVersion
	var workflowArgow types.WorkflowChartVersion

	var workflows []types.Workflow
	var experiments []string
	for _, composant := range composants.Objects {
		var workflow types.Workflow
		for experiment := range composant.GeneratedExperiments {
			experiments = append(experiments, composant.GeneratedExperiments[experiment])
		}
		workflow.Experiments = experiments
		workflow.Name = composants.Namespace + "-" + composant.GetName()
		workflows = append(workflows, workflow)
	}
	return workflows

}

func generateExperiment(experimentName string, composant resources.Object) []types.ChaosChart {
	var experiments []types.ChaosChart
	switch experimentName {
	case "container-kill":
		exps := containerKill.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments

	default:
		fmt.Printf("Unsupported experiment %v, please provide the correct value of experiment\n", experimentName)
		return experiments
	}
}

func generateExperiments(composants *resources.Resources, experimentsList []string) []types.ChaosChart {
	var experiments []types.ChaosChart
	for _, experimentName := range experimentsList {
		for composant := range composants.Objects {
			exps := generateExperiment(experimentName, composants.Objects[composant])
			for _, experiment := range exps {
				experiments = append(experiments, experiment)
				composants.Objects[composant].AddGeneratedExperiment(experiment.ChartVersion.Metadata.Name)
			}
		}
	}
	return experiments
}

func generateHub(clusterName string, namespace *resources.Resources, experiments []types.ChaosChart, workflows []types.Workflow) types.Manifest {
	var project types.Manifest
	project.Name = clusterName + "-" + namespace.Namespace
	project.Description = "Experiments and workflow for namespace " + namespace.Namespace + " on cluster " + clusterName
	project.Platform = clusterName
	project.Experiments = experiments
	project.Workflows = workflows
	return project
}

func Generate(clusterName string, res []*resources.Resources) ([]types.Manifest, error) {
	var projects []types.Manifest
	var err error
	for _, namespace := range res {
		//ready, err := requirements.CheckHaveLabels(namespace, []string{"composant"})
		//if !ready {
		//        fmt.Printf("Cannot find labels in namespace :\n %v \n", namespace.Namespace)
		//        continue
		//}

		compliant := requirements.FindUniqueLabels(namespace)
		if !compliant {
			fmt.Printf("Cannot find determinist labels in namespace %v \n", namespace.Namespace)
			continue
		}

		fmt.Printf("Lancement de la génération des experiments pour le namespace %s.\n", namespace.Namespace)
		experiments := generateExperiments(namespace, experimentsList)
		if len(experiments) < 1 {
			fmt.Printf("Cannot generate experiments for: %v\n", namespace.Namespace)
			continue
		}

		fmt.Printf("Génération des experiments terminée.\n")
		fmt.Printf("Lancement de la génération des workflows.\n")

		workflows := generateWorkflows(namespace)
		if len(workflows) < 1 {
			fmt.Printf("Cannot generate workflows for: %v\n", namespace.Namespace)
			continue
		}
		fmt.Printf("Génération des workflows terminée.\n")

		fmt.Printf("Packaging du hub.\n")

		project := generateHub(clusterName, namespace, experiments, workflows)
		projects = append(projects, project)
		fmt.Printf("Packaging terminé.\n")

	}
	return projects, err

	//	file, _ := json.MarshalIndent(manifest, "", " ")
	//	fmt.Println(string(file))
}
