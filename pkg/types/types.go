package types

type ChaosChart struct {
	ChartVersion    ChaosChartVersion
	ChaosEngine     ChaosEngine
	ChaosExperiment ChaosExperiment
	Icon            []byte
}

type WorkflowArgow struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		GenerateName string `yaml:"generateName"`
		Namespace    string `yaml:"namespace"`
		Labels       struct {
			Subject string `yaml:"subject"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Entrypoint         string `yaml:"entrypoint"`
		ServiceAccountName string `yaml:"serviceAccountName"`
		SecurityContext    struct {
			RunAsUser    int  `yaml:"runAsUser"`
			RunAsNonRoot bool `yaml:"runAsNonRoot"`
		} `yaml:"securityContext"`
		Arguments struct {
			Parameters []struct {
				Name  string `yaml:"name"`
				Value string `yaml:"value"`
			} `yaml:"parameters"`
		} `yaml:"arguments"`
		Templates []struct {
			Name  string `yaml:"name"`
			Steps [][]struct {
				Name     string `yaml:"name"`
				Template string `yaml:"template"`
			} `yaml:"steps,omitempty"`
			Inputs struct {
				Artifacts []struct {
					Name string `yaml:"name"`
					Path string `yaml:"path"`
					Raw  struct {
						Data string `yaml:"data"`
					} `yaml:"raw"`
				} `yaml:"artifacts"`
			} `yaml:"inputs,omitempty"`
			Container struct {
				Image   string   `yaml:"image"`
				Command []string `yaml:"command"`
				Args    []string `yaml:"args"`
			} `yaml:"container,omitempty"`
		} `yaml:"templates"`
	} `yaml:"spec"`
}

type WorkflowChart struct {
	WorkflowChartVersion WorkflowChartVersion
	WorkflowArgow        WorkflowArgow
}

type WorkflowChartVersion struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string `yaml:"name"`
		Version     string `yaml:"version"`
		Annotations struct {
			Categories       string `yaml:"categories"`
			ChartDescription string `yaml:"chartDescription"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec struct {
		DisplayName         string   `yaml:"displayName"`
		CategoryDescription string   `yaml:"categoryDescription"`
		Experiments         []string `yaml:"experiments"`
		Keywords            []string `yaml:"keywords"`
		Platforms           []string `yaml:"platforms"`
		Maintainers         []struct {
			Name  string `yaml:"name"`
			Email string `yaml:"email"`
		} `yaml:"maintainers"`
		Provider struct {
			Name string `yaml:"name"`
		} `yaml:"provider"`
		Links []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		} `yaml:"links"`
		Icon []struct {
			URL       string `yaml:"url"`
			Mediatype string `yaml:"mediatype"`
		} `yaml:"icon"`
	} `yaml:"spec"`
}

type ChaosEngine struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		EngineState string `yaml:"engineState"`
		Appinfo     struct {
			Appns    string `yaml:"appns"`
			Applabel string `yaml:"applabel"`
			Appkind  string `yaml:"appkind"`
		} `yaml:"appinfo"`
		ChaosServiceAccount string `yaml:"chaosServiceAccount"`
		Experiments         []struct {
			Name string `yaml:"name"`
			Spec struct {
				Components struct {
					Env []struct {
						Name  string `yaml:"name"`
						Value string `yaml:"value"`
					} `yaml:"env"`
				} `yaml:"components"`
			} `yaml:"spec"`
		} `yaml:"experiments"`
	} `yaml:"spec"`
}

type ChaosChartVersion struct {
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

type ChaosExperiment struct {
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
	Experiments []ChaosChart `yaml:"experiments"`
}
