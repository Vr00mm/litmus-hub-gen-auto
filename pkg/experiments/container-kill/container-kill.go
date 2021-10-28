package containerKill

import (
	"fmt"
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/types"
)

var experimentRequirements = []string{"isRoot"}

var availableLibs = []string{"litmus", "pumba"}

var config = map[string]string{
	"Name": "container-kill",
	"Description": ``,

}

var parameters = map[string]string{
			"TARGET_CONTAINER": "",
			"RAMP_TIME": "",
			"LIB": "litmus",
			"TARGET_PODS": "",
			"CHAOS_INTERVAL": "10",
			"SIGNAL": "SIGKILL",
			"SOCKET_PATH": "/var/run/docker.sock",
			"CONTAINER_RUNTIME": "docker",
			"TOTAL_CHAOS_DURATION": "60",
			"PODS_AFFECTED_PERC": "",
			"LIB_IMAGE": "litmuschaos/go-runner:latest",
			"SEQUENCE": "parallel",
}

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
	fmt.Printf("Le composant: %s a comme containers %v\n", composant.GetName(), containers)
	for container := range containers {
		exp := generateContainerKill(composant, containers[container])
		experiments = append(experiments, exp)
	}
	return experiments
}
