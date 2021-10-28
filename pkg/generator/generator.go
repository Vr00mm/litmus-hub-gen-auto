package generator

import (
	"fmt"
	//        "encoding/json"
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"

	//	"strings"
	containerKill "hub-gen-auto/pkg/experiments/container-kill"
	podKill "hub-gen-auto/pkg/experiments/pod-kill"
)

var experimentsList = []string{"container-kill", "pod-kill"}

func generateWorkflows(composants *resources.Resources) []types.Workflow {
	var workflows []types.Workflow
	var experiments []string
	for _, composant := range composants.Objects {
		var workflow types.Workflow
		for _, experiment := range composant.GeneratedExperiments {
			experiments = append(experiments, experiment)
		}
		workflow.Experiments = experiments
		workflow.Name = composants.Namespace + "-" + composant.GetName()
		workflows = append(workflows, workflow)
	}
	return workflows

}

func generateExperiment(experimentName string, composant resources.Object) []types.Experiment {
	var experiments []types.Experiment
	switch experimentName {
	case "container-kill":
		exps := containerKill.Generate(composant)
		for _, exp := range exps {
			experiments = append(experiments, exp)
		}
		return experiments
	case "pod-kill":
		exps := podKill.Generate(composant)
		for _, exp := range exps {
			experiments = append(experiments, exp)
		}
		return experiments
	default:
		fmt.Printf("Unsupported experiment %v, please provide the correct value of experiment\n", experimentName)
		return experiments
	}

}

func generateExperiments(composants *resources.Resources, experimentsList []string) (*resources.Resources, []types.Experiment) {
	var experiments []types.Experiment
	var objs []resources.Object
	for _, experimentName := range experimentsList {
		for _, composant := range composants.Objects {
			exps := generateExperiment(experimentName, composant)
			for _, experiment := range exps {
				experiments = append(experiments, experiment)
				composant.AddGeneratedExperiment(experiment.Name)
				objs = append(objs, composant)
			}
		}
	}
	composants.Objects = objs
	return composants, experiments
}

func Generate(clusterName string, res []*resources.Resources) []types.Manifest {
	var projects []types.Manifest

	for _, namespace := range res {
		//ready, err := requirements.CheckHaveLabels(namespace, []string{"composant"})
		//if !ready {
		//        fmt.Printf("Cannot find labels in namespace :\n %v \n", namespace.Namespace)
		//        continue
		//}

		fmt.Printf("before the shet: %v \n\n", namespace)
		namespace, compliant := requirements.FindUniqueLabels(namespace)
		if !compliant {
			fmt.Printf("Cannot find determinist labels in namespace %v : %v\n", namespace.Namespace)
			continue

		}
		fmt.Printf("here is the shet: %v\n\n\n", namespace)

		fmt.Printf("Lancement de la génération des experiments.\n")
		var experiments []types.Experiment
		namespace, experiments = generateExperiments(namespace, experimentsList)
		if len(experiments) < 1 {
			fmt.Printf("Cannot generate experiments for: %v\n", namespace.Namespace)
			continue
		}

		fmt.Printf("Génération des experiments terminée.\n")

		fmt.Printf("Lancement de la génération des workflows.\n")
		var workflows []types.Workflow
		workflows = generateWorkflows(namespace)
		fmt.Printf("Génération des workflows terminée.\n")

		fmt.Printf("Packaging du hub.\n")

		var project types.Manifest
		project.Name = clusterName + "-" + namespace.Namespace
		project.Description = "Experiments and workflow for namespace " + namespace.Namespace + " on cluster " + clusterName
		project.Platform = clusterName
		project.Experiments = experiments
		project.Workflows = workflows
		fmt.Printf("Packaging terminé.\n")

		projects = append(projects, project)
	}
	return projects

	//	file, _ := json.MarshalIndent(manifest, "", " ")
	//	fmt.Println(string(file))
}
