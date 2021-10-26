package containerKill

import (
	"hub-gen-auto/pkg/requirements"
)

var experimentRequirements = []string{"isRoot"}

func CheckRequirements() bool {
	return requirements.CheckRequirements(experimentRequirements)
}

func GenerateContainerKill(composant) {
	var experiment Experiment
        experiment.Name := composant.Name
        experiment.Template := experimentName
        experiment.Label := composant.Label
        experiment.Kind := composant.Kind
        experiment.Args := composant.Args
        return experiment
}

func Generate(composant) []Experiment {
	var experiments []Experiment
	containers := resources.getContainers(composant)
	for _, container := range containers {
		experiments = append(experiments, generateContainerKill(container))
	}
	return experiments
}
