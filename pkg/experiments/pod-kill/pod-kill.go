package podKill

import (
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
	"strings"
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
	containers, _ := composant.GetContainers()
	exp := generateExperiment(composant)
	exp.Type = composant.Type
	exp.TemplateName = "pod-delete"
	exp.Composant = composant.GetName()
	exp.Container = strings.Join(containers, ", ")
	experiments = append(experiments, exp)
	return experiments
}
