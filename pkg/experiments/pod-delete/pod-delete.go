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
	exp := ExperimentsManifests["pod-delete"]
	exp.ChaosExperiment.Metadata.Name = "pod-delete-" + composant.GetName()
	exp.ChaosEngine.Spec.Appinfo.Appkind = composant.Type
	exp.ChaosEngine.Spec.Appinfo.Applabel = composant.GetUniqueLabel()
	exp.ChaosEngine.Spec.Appinfo.Appns = composant.GetNamespace()

    var arrayExpParameters []struct{Name string "yaml:\"name\""; Spec struct{Components struct{Env []struct{Name string "yaml:\"name\""; Value string "yaml:\"value\""} "yaml:\"env\""} "yaml:\"components\""} "yaml:\"spec\""}
	var expParameters struct{Name string "yaml:\"name\""; Spec struct{Components struct{Env []struct{Name string "yaml:\"name\""; Value string "yaml:\"value\""} "yaml:\"env\""} "yaml:\"components\""} "yaml:\"spec\""}
    
	expParameters.Name = "pod-delete-" + composant.GetName()
	arrayExpParameters = append(arrayExpParameters, expParameters)
	exp.ChaosEngine.Spec.Experiments = arrayExpParameters
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
