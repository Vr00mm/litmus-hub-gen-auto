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
	var exp types.ChaosChart
	exp = ExperimentsManifests["container-kill"]
	exp.ChaosExperiment.Metadata.Name = "container-kill-" + composant.GetName() + "-" + container
	exp.ChaosEngine.Spec.Appinfo.Appkind = composant.Type
	exp.ChaosEngine.Spec.Appinfo.Applabel = composant.GetUniqueLabel()
	exp.ChaosEngine.Spec.Appinfo.Appns = composant.GetNamespace()
	return exp
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
		exp.Type = composant.Type
		exp.TemplateName = "container-kill"
		exp.Composant = composant.GetName()
		exp.Container = containers[container]
		experiments = append(experiments, exp)
	}
	return experiments
}
