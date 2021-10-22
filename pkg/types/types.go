package types

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

type Manifests []struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Namespace   string       `yaml:"namespace"`
	Platform    string       `yaml:"platform"`
	GitURL      string       `yaml:"gitUrl"`
	Workflows   []Workflow   `yaml:"workflows"`
	Experiments []Experiment `yaml:"experiments"`
}
