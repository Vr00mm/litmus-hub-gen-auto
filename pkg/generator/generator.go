package generator

import (
	"fmt"
	//        "encoding/json"
	"hub-gen-auto/pkg/requirements"
	"hub-gen-auto/pkg/resources"
	"strings"
	//        "hub-gen-auto/pkg/experiments"
)

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

func generateExperiment(experimentName, composant){
	var experiment Experiment
        experiment.Name := composant.Name
        experiment.Template := experimentName
	experiment.Label := composant.Label
	experiment.Kind := composant.Kind
	experiment.Args := composant.Args
	return experiment
}

func generateExperiments(composants []*resources.Resources, experiments []string){
	var experiments []Experiment
	for _, experimentName := range experiments {
		for _, composant := range composants {
			ready := requirements.checkRequirements(experimentName, composant)
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
		targetLabels := make(map[string]map[string]string, len(namespace.Deploys.Items))
		for _, composant := range namespace.Deploys.Items {
			ready := requirements.CheckHaveLabels(composant.ObjectMeta.Labels)
			if ready == true {
				targetLabels[composant.Name] = make(map[string]string, len(composant.ObjectMeta.Labels))
				targetLabels[composant.Name] = composant.ObjectMeta.Labels
			}
		}
		for _, composant := range namespace.Stss.Items {
			ready := requirements.CheckHaveLabels(composant.ObjectMeta.Labels)
			if ready == true {
				targetLabels[composant.Name] = make(map[string]string, len(composant.ObjectMeta.Labels))
				targetLabels[composant.Name] = composant.ObjectMeta.Labels
			}
		}

		newVar := requirements.FindUniqueLabels(targetLabels)

		if len(namespace.Deploys.Items)+len(namespace.Stss.Items) == len(newVar) && len(newVar) > 0 {
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
		experiments = generateExperiments(res, newVar)
                fmt.Printf("Génération des experiments terminée.\n")

                fmt.Printf("Lancement de la génération des workflows.\n")
		var workflows []Workflow
		workflows = generateWorkflows(res, newVar)
                fmt.Printf("Génération des workflows terminée.\n")

                fmt.Printf("Packaging du hub.\n")
                var hub Hub
                hub = generateHub(res, experiments, workflows)
                fmt.Printf("Packaging terminé.\n")


	}

	//	file, _ := json.MarshalIndent(manifest, "", " ")
	//	fmt.Println(string(file))
}
