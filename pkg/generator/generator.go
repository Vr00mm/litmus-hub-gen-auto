package generator

import (
	"fmt"
	//        "encoding/json"
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"strings"
	//        "hub-gen-auto/pkg/experiments"
)

var experimentList = []string{"container-kill", "test-123"}

type Workflow struct {
	Name        string   `yaml:"name"`
	Experiments []string `yaml:"experiments"`
}

type Experiment struct {
	Name      string      `yaml:"name"`
	Template  string      `yaml:"template"`
	Label     string      `yaml:"label"`
	Kind      string      `yaml:"kind"`
	Args      interface{} `yaml:"args"`
}

type Hub struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Namespace   string       `yaml:"namespace"`
	Platform    string       `yaml:"platform"`
	GitURL      string       `yaml:"gitUrl"`
	Workflows   []Workflow   `yaml:"workflows"`
	Experiments []Experiment `yaml:"experiments"`
}

type Projects struct {
	ClusterName    	string `yaml:"cluster_name"`
	Namespaces 	[]Hub  `yaml:"namespaces"`
}

func generateExperiment(experimentName string, composant resources.Object) (Experiment){
	var experiment Experiment
        experiment.Name = "test"
        experiment.Template = experimentName
	experiment.Label = "test"
	experiment.Kind = composant.Type
	experiment.Args = "test"
	return experiment
}

func generateExperiments(composants *resources.Resources, experimentsList []string){
	var experiments []Experiment
	for _, experimentName := range experimentsList {
		for _, composant := range composants.Objects {
			ready := requirements.CheckRequirements(experimentName, composant)
			if !ready {
				continue
			}
			experiment := generateExperiment(experimentName, composant)
			experiments = append(experiments, experiment)
		}
	}
}

func Generate(clusterName string, res []*resources.Resources) {
	var projects Projects
	projects.ClusterName = clusterName

	for _, namespace := range res {

                ready, err := requirements.CheckHaveLabels(namespace, []string{"composant"})
                if !ready {
	                fmt.Printf("Erreur lors de la vérification des labels: %v\n", err)
                        continue
                }


		targetLabels, err := requirements.FindUniqueLabels(namespace)
		if err != nil {
		}

		if len(namespace.Objects) == len(targetLabels) && len(targetLabels) > 0 {
			skip := true
			for _, labels := range newVar {
				for _, label := range labels {
					if strings.HasPrefix(label, "composant") {
						skip = false
					}
				}
			}
			if skip {
				fmt.Printf("Le namespace %s n'est pret pour le chaos: %s\n", namespace.Namespace)
				continue
			}
		} else {
                                fmt.Printf("Le namespace %s n'est pret pour le chaos: %s\n", namespace.Namespace)
                                continue
		}
                fmt.Printf("Le namespace %s est pret pour le chaos: %s\n", namespace.Namespace)

                fmt.Printf("Lancement de la génération des experiments.\n")

		var experiments []Experiment
		experiments = generateExperiments(namespace, experimentsList)
                fmt.Printf("Génération des experiments terminée.\n")

                fmt.Printf("Lancement de la génération des workflows.\n")
		var workflows []Workflow
		workflows = generateWorkflows(namespace, workflowsList)
                fmt.Printf("Génération des workflows terminée.\n")

                fmt.Printf("Packaging du hub.\n")
                var hub Hub
                hub = generateHub(namespace, experiments, workflows)
                fmt.Printf("Packaging terminé.\n")


	}

	//	file, _ := json.MarshalIndent(manifest, "", " ")
	//	fmt.Println(string(file))
}
