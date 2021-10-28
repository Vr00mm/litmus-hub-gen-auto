package types

type LitmusChartVersion struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string `yaml:"name"`
		Version     string `yaml:"version"`
		Annotations struct {
			Categories string `yaml:"categories"`
			Vendor     string `yaml:"vendor"`
			Support    string `yaml:"support"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec struct {
		DisplayName         string   `yaml:"displayName"`
		CategoryDescription string   `yaml:"categoryDescription"`
		Keywords            []string `yaml:"keywords"`
		Platforms           []string `yaml:"platforms"`
		Maturity            string   `yaml:"maturity"`
		ChaosType           string   `yaml:"chaosType"`
		Maintainers         []struct {
			Name  string `yaml:"name"`
			Email string `yaml:"email"`
		} `yaml:"maintainers"`
		MinKubeVersion string `yaml:"minKubeVersion"`
		Provider       struct {
			Name string `yaml:"name"`
		} `yaml:"provider"`
		Labels struct {
			AppKubernetesIoComponent string `yaml:"app.kubernetes.io/component"`
			AppKubernetesIoVersion   string `yaml:"app.kubernetes.io/version"`
		} `yaml:"labels"`
		Links []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		} `yaml:"links"`
		Icon []struct {
			URL       interface{} `yaml:"url"`
			Mediatype string      `yaml:"mediatype"`
		} `yaml:"icon"`
		Chaosexpcrdlink string `yaml:"chaosexpcrdlink"`
	} `yaml:"spec"`
}


type LitmusExperiment struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			Name                     string `yaml:"name"`
			AppKubernetesIoPartOf    string `yaml:"app.kubernetes.io/part-of"`
			AppKubernetesIoComponent string `yaml:"app.kubernetes.io/component"`
			AppKubernetesIoVersion   string `yaml:"app.kubernetes.io/version"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Definition struct {
			Scope       string `yaml:"scope"`
			Permissions []struct {
				APIGroups []string `yaml:"apiGroups"`
				Resources []string `yaml:"resources"`
				Verbs     []string `yaml:"verbs"`
			} `yaml:"permissions"`
			Image           string   `yaml:"image"`
			ImagePullPolicy string   `yaml:"imagePullPolicy"`
			Args            []string `yaml:"args"`
			Command         []string `yaml:"command"`
			Env             []struct {
				Name  string `yaml:"name"`
				Value string `yaml:"value"`
			} `yaml:"env"`
			Labels struct {
				Name                     string `yaml:"name"`
				AppKubernetesIoPartOf    string `yaml:"app.kubernetes.io/part-of"`
				AppKubernetesIoComponent string `yaml:"app.kubernetes.io/component"`
				AppKubernetesIoVersion   string `yaml:"app.kubernetes.io/version"`
			} `yaml:"labels"`
		} `yaml:"definition"`
	} `yaml:"spec"`
}

type Workflow struct {
	Name        string   `yaml:"name"`
	Experiments []string `yaml:"experiments"`
}

type Experiment struct {
	Name        string            `yaml:"name"`
	Template    string            `yaml:"template"`
	Permissions string            `yaml:"permissions"`
	Label       string            `yaml:"label"`
	Kind        string            `yaml:"kind"`
	Args        map[string]string `yaml:"args"`
}

type Manifest struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Namespace   string       `yaml:"namespace"`
	Platform    string       `yaml:"platform"`
	GitURL      string       `yaml:"gitUrl"`
	Workflows   []Workflow   `yaml:"workflows"`
	Experiments []Experiment `yaml:"experiments"`
}
