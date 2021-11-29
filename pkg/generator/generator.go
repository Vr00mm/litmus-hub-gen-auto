package generator

import (
	"fmt"
	//        "encoding/json"
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
	"hub-gen-auto/pkg/utils"
	"hub-gen-auto/pkg/workflow"

	"gopkg.in/yaml.v3"

	//	"strings"
	containerKill "hub-gen-auto/pkg/experiments/container-kill"
	podCPUHog "hub-gen-auto/pkg/experiments/pod-cpu-hog"
	podKill "hub-gen-auto/pkg/experiments/pod-delete"
	podIOStress "hub-gen-auto/pkg/experiments/pod-io-stress"
	podMemHog "hub-gen-auto/pkg/experiments/pod-memory-hog"
	podNetCorruption "hub-gen-auto/pkg/experiments/pod-network-corruption"
	podNetDuplication "hub-gen-auto/pkg/experiments/pod-network-duplication"
	podNetLatency "hub-gen-auto/pkg/experiments/pod-network-latency"
	podNetLoss "hub-gen-auto/pkg/experiments/pod-network-loss"
)

var experimentsList = []string{"container-kill", "pod-delete", "pod-cpu-hog-exec", "pod-memory-hog-exec", "pod-io-stress", "pod-network-corruption", "pod-network-duplication", "pod-network-loss", "pod-network-latency"}
var experimentsManifests map[string]types.ChaosChart
var hubChartData []byte

func init() {
	experimentsManifests = utils.GetExperimentsManifests(experimentsList)
	hubChartData = utils.MakeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/generic.chartserviceversion.yaml")
	containerKill.ExperimentsManifests = experimentsManifests
	podCPUHog.ExperimentsManifests = experimentsManifests
	podKill.ExperimentsManifests = experimentsManifests
	podIOStress.ExperimentsManifests = experimentsManifests
	podMemHog.ExperimentsManifests = experimentsManifests
	podNetCorruption.ExperimentsManifests = experimentsManifests
	podNetDuplication.ExperimentsManifests = experimentsManifests
	podNetLatency.ExperimentsManifests = experimentsManifests
	podNetLoss.ExperimentsManifests = experimentsManifests
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
	case "pod-delete":
		exps := podKill.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-cpu-hog-exec":
		exps := podCPUHog.Generate(composant)
		for exp := range exps {
			experiments = append(experiments, exps[exp])
		}
		return experiments
	case "pod-memory-hog-exec":
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
	var hubPackage types.HubPackage
	var hubChart types.ChaosChartVersion

	if err := yaml.Unmarshal(hubChartData, &hubChart); err != nil {
		panic(err)
	}

	hubChart.Metadata.Name = clusterName + "-" + namespace.Namespace
	hubChart.Spec.DisplayName = clusterName + "-" + namespace.Namespace

	hubChart.Spec.Experiments = []string{}
	hubChart.Spec.Keywords = []string{}

	hubPackage.PackageName = clusterName + "-" + namespace.Namespace
	for _, experiment := range experiments {
		hubPackage.Experiments = append(hubPackage.Experiments, struct {
			Name string "yaml:\"name\""
			CSV  string "yaml:\"CSV\""
			Desc string "yaml:\"desc\""
		}{experiment.ChaosExperiment.Metadata.Name, experiment.ChaosExperiment.Metadata.Name + ".chartserviceversion.yaml", experiment.ChaosExperiment.Metadata.Name})
		hubChart.Spec.Experiments = append(hubChart.Spec.Experiments, experiment.ChaosExperiment.Metadata.Name)
	}
	project.Name = clusterName + "-" + namespace.Namespace
	project.Namespace = namespace.Namespace
	project.Description = "Experiments and workflow for namespace " + namespace.Namespace + " on cluster " + clusterName
	project.Platform = clusterName
	project.Experiments = experiments
	project.Workflows = workflows
	project.Package = hubPackage
	project.Chart = hubChart
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
			fmt.Printf("Cannot find determinist labels in namespace %v in %v \n", namespace.Namespace, clusterName)
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
