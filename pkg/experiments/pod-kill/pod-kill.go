package podKill

import (
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
)

var ExperimentsManifests map[string]types.ChaosChart

var experimentRequirements = []string{"isRoot"}

var availableLibs = []string{"litmus", "pumba"}

func generateExperiment(composant resources.Object) types.ChaosChart {
	exp := ExperimentsManifests["pod-kill"]
	exp.ChaosExperiment.Metadata.Name = "pod-kill-" + composant.GetName()
	return exp
}

func checkRequirements(composant resources.Object) bool {
	return requirements.CheckRequirements(experimentRequirements, composant)
}

func Generate(composant resources.Object) []types.ChaosChart {
	var experiments []types.ChaosChart
	exp := generateExperiment(composant)
	experiments = append(experiments, exp)
	return experiments
}
