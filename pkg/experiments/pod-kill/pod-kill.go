package podKill

import (
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
)

var experimentRequirements = []string{}

func checkRequirements(composant resources.Object) bool {
	return requirements.CheckRequirements(experimentRequirements, composant)
}

func generatePodKill(composant resources.Object) types.Experiment {
	var experiment types.Experiment
	experiment.Name = "pod-kill-" + composant.Type + "-" + composant.GetName()
	experiment.Template = "pod-kill"
	experiment.Label = composant.GetUniqueLabel()
	experiment.Kind = composant.Type
	return experiment
}

func Generate(composant resources.Object) []types.Experiment {
	var experiments []types.Experiment
	exp := generatePodKill(composant)
	experiments = append(experiments, exp)
	return experiments
}
