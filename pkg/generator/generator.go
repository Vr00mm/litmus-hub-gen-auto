package generator

import (
	"fmt"
	//        "encoding/json"
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
	"hub-gen-auto/pkg/utils"
	"hub-gen-auto/pkg/workflow"

	//	"strings"
	containerKill "hub-gen-auto/pkg/experiments/container-kill"
	podCPUHog "hub-gen-auto/pkg/experiments/pod-cpu-hog"
	podIOStress "hub-gen-auto/pkg/experiments/pod-io-stress"
	podKill "hub-gen-auto/pkg/experiments/pod-kill"
	podMemHog "hub-gen-auto/pkg/experiments/pod-memory-hog"
	podNetCorruption "hub-gen-auto/pkg/experiments/pod-network-corruption"
	podNetDuplication "hub-gen-auto/pkg/experiments/pod-network-duplication"
	podNetLatency "hub-gen-auto/pkg/experiments/pod-network-latency"
	podNetLoss "hub-gen-auto/pkg/experiments/pod-network-loss"
)

var experimentsList = []string{"container-kill", "pod-kill", "pod-cpu-hog", "pod-mem-hog", "pod-io-stress"}
var experimentsManifests map[string]types.ChaosChart

func init() {
	experimentsManifests = utils.GetExperimentsManifests(experimentsList)
	containerKill.ExperimentsManifests = experimentsManifests
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
	case "pod-kill":
		exps := podKill.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-cpu-hog":
		exps := podCPUHog.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-mem-hog":
		exps := podMemHog.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-io-stress":
		exps := podIOStress.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-network-corruption":
		exps := podNetCorruption.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-network-duplication":
		exps := podNetDuplication.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-network-latency":
		exps := podNetLatency.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-network-loss":
		exps := podNetLoss.Generate(composant)
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

func generateHub(clusterName string, namespace *resources.Resources, experiments []types.ChaosChart, workflows []types.WorkflowChart) types.Manifest {
	var project types.Manifest
	project.Name = clusterName + "-" + namespace.Namespace
	project.Namespace = namespace.Namespace
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

		workflows := workflow.Generate(namespace, experimentsManifests)
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
