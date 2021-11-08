package podIOStress

import (
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
)

var ExperimentsManifests map[string]types.ChaosChart

var experimentRequirements = []string{"isRoot"}

var availableLibs = []string{"litmus", "pumba"}

func generateExperiment(composant resources.Object, container string) types.ChaosChart {
	exp := ExperimentsManifests["pod-io-stress"]
	exp.ChaosExperiment.Metadata.Name = "pod-io-stress-" + composant.GetName() + "-" + container
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
		exp.Type = composant.Type
		exp.TemplateName = "pod-io-stress"
		exp.Composant = composant.GetName()
		exp.Container = containers[container]
		experiments = append(experiments, exp)
	}
	return experiments
}