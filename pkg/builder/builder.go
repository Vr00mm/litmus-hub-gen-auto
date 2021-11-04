package builder

import (
	"fmt"
	"hub-gen-auto/pkg/types"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func writeToFile(filename string, data interface{}) {
	file, err := yaml.Marshal(data)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}
}

func Build(clusterName string, manifests []types.Manifest) error {
	var err error
	for _, manifest := range manifests {
		os.MkdirAll("build/"+clusterName+"/"+manifest.Namespace+"/icons/", os.ModePerm)
		writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+clusterName+"-"+manifest.Namespace+".chartserviceversion.yaml", manifest.Chart)
		writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+clusterName+"-"+manifest.Namespace+".package.yaml", manifest.Package)
		for _, experiment := range manifest.Experiments {
			os.MkdirAll("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/", os.ModePerm)
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/experiment.yaml", experiment.ChaosExperiment)
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/engine.yaml", experiment.ChaosEngine)
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/rbac.yaml", experiment.RBAC)
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/rbac-psp.yaml", experiment.PSP)
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/"+experiment.ChaosExperiment.Metadata.Name+".chartserviceversion.yaml", experiment.ChaosExperiment)
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/icons/"+experiment.ChaosExperiment.Metadata.Name+".png", experiment.Icon)
		}

	}
	return err
}
