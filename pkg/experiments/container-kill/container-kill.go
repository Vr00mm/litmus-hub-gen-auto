package containerKill

import (
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
)

var experimentRequirements = []string{"isRoot"}

func checkRequirements(composant resources.Object) bool {
	return requirements.CheckRequirements(experimentRequirements, composant)
}

func generateContainerKill(composant resources.Object, container string) types.Experiment {
	var experiment types.Experiment
	experiment.Name = "container-kill-" + composant.Type + "-" + composant.GetName() + "-" + container
	experiment.Template = "container-kill"
	experiment.Label = composant.GetUniqueLabel()
	experiment.Kind = composant.Type
	experiment.Args = map[string]string{
		"TARGET_CONTAINER": container,
	}
	return experiment
}

func Generate(composant resources.Object) []types.Experiment {
	var experiments []types.Experiment
	containers, _ := composant.GetContainers()
	for _, container := range containers {
		exp := generateContainerKill(composant, container)
		experiments = append(experiments, exp)
	}
	return experiments
}
