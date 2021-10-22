package generator

import (
	"fmt"
//        "encoding/json"

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

func Generate(clusterName string, res []*resources.Resources) {
        for _, namespace := range res {
		fmt.Printf("Le nom est : %s\n", namespace.Namespace)
		for _, composant := range namespace.Deploys.Items {
	                fmt.Printf("Composant Trouvé : %s\n", composant.Name)
			ready := requirements.CheckHaveLabels(composant)
			if ready = true {
				fmt.Printf("Le composant a bien les labels necessaire !!")
                        }
                }
        }

//	file, _ := json.MarshalIndent(manifest, "", " ")
//	fmt.Println(string(file))
}
