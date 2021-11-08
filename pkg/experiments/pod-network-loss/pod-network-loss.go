package podNetLoss

import (
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
)

var ExperimentsManifests map[string]types.ChaosChart

var experimentRequirements = []string{"isRoot"}

var availableLibs = []string{"litmus", "pumba"}

func generateExperiment(composant resources.Object, container string) types.ChaosChart {
	exp := ExperimentsManifests["pod-network-loss"]
	exp.ChaosExperiment.Metadata.Name = "pod-network-loss-" + composant.GetName() + "-" + container
	return exp
}

func checkRequirements(composant resources.Object) bool {
	return requirements.CheckRequirements(experimentRequirements, composant)
}

func Generate(composant resources.Object) []types.ChaosChart {
	var experiments []types.ChaosChart
	containers, _ := composant.GetContainers()
	for container := range containers {
		exp := generateExperiment(composant, containers[container])
		exp.TemplateName = "pod-network-loss"
		exp.Container = containers[container]
		exp.Composant = composant.GetName()
		exp.Type = composant.Type
		experiments = append(experiments, exp)
	}
	return experiments
}
