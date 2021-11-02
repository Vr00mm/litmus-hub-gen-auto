package containerKill

import (
	"fmt"
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
)

var ExperimentsManifests map[string]types.ChaosChart

var experimentRequirements = []string{"isRoot"}

var availableLibs = []string{"litmus", "pumba"}

func generateExperiment(composant resources.Object, container string) types.ChaosChart {
	testVar := ExperimentsManifests["container-kill"]
	testVar.ChaosExperiment.Metadata.Name = composant.GetName()
	return testVar
}

func checkRequirements(composant resources.Object) bool {
	return requirements.CheckRequirements(experimentRequirements, composant)
}

func Generate(composant resources.Object) []types.ChaosChart {
	var experiments []types.ChaosChart
	containers, _ := composant.GetContainers()
	fmt.Printf("Le composant: %s a comme containers %v\n", composant.GetName(), containers)
	for container := range containers {
		exp := generateExperiment(composant, containers[container])
		experiments = append(experiments, exp)
	}
	return experiments
}
