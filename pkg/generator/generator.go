package generator

import (
	"fmt"
//        "encoding/json"
	"strings"
        "hub-gen-auto/pkg/resources"
        "hub-gen-auto/pkg/requirements"

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
	Composant string      `yaml:"composant"`
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

type Project	struct {
        Name    string  `yaml:"name"`
	Project	[]Hub	`yaml:"project"`
}

type Projects struct {
	Name	string	`yaml:"name"`
	Projects	[]Project	`yaml:"projects"`
}


func generateFromResources() {

}

type Labels map[string]string

func Generate(clusterName string, res []*resources.Resources) {
	for _, namespace := range res {
                targetLabels := make(map[string]map[string]string, len(namespace.Deploys.Items)) //string //map[string]map[string]strin
		for _, composant := range namespace.Deploys.Items {
			ready := requirements.CheckHaveLabels(composant.ObjectMeta.Labels)
			if ready == true {
		                targetLabels[composant.Name] = make(map[string]string, len(composant.ObjectMeta.Labels)) //string //map[string]map[string]string
                                targetLabels[composant.Name] = composant.ObjectMeta.Labels
                        }
                }
                for _, composant := range namespace.Stss.Items {
                        ready := requirements.CheckHaveLabels(composant.ObjectMeta.Labels)
                        if ready == true {
                                targetLabels[composant.Name] = make(map[string]string, len(composant.ObjectMeta.Labels)) //string //map[string]map[string]string
                                targetLabels[composant.Name] = composant.ObjectMeta.Labels
                        }
                }

		newVar := requirements.FindUniqueLabelsBb(targetLabels)
		if len(namespace.Deploys.Items) + len (namespace.Stss.Items) == len(newVar) && len(newVar) > 0 {
			skip := false
			for _, labels := range newVar {
				for _, label := range labels {
					if !strings.HasPrefix(label, "composant") {
						skip = true
					}
				}
			}
			if !skip {
			 fmt.Printf("Le namespace %s est pret pour le chaos: %s\n", namespace.Namespace, newVar)
			}
		}
        }

//	file, _ := json.MarshalIndent(manifest, "", " ")
//	fmt.Println(string(file))
}
