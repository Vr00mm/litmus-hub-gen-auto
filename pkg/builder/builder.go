package builder

import (
	"fmt"
	"hub-gen-auto/pkg/types"
	"os"

	"gopkg.in/yaml.v3"
)

func writeToFile(filename string, data string) {
	destination, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create:", err)
		return
	}
	defer destination.Close()
	fmt.Fprintf(destination, "%s", data)
}

func Build(clusterName string, manifests []types.Manifest) error {
	var err error
	os.MkdirAll("build/"+clusterName, os.ModePerm)

	for _, manifest := range manifests {
		os.MkdirAll("build/"+clusterName+"/"+manifest.Namespace, os.ModePerm)
		for _, experiment := range manifest.Experiments {
			os.MkdirAll("build/"+clusterName+"/"+manifest.Namespace, os.ModePerm)
			exp, _ := yaml.Marshal(&experiment.ChaosExperiment)
			engine, _ := yaml.Marshal(&experiment.ChaosEngine)
			chart, _ := yaml.Marshal(&experiment.ChaosExperiment)
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/experiment.yaml", string(exp))
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/engine.yaml", string(engine))
			writeToFile("build/"+clusterName+"/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/"+experiment.ChaosExperiment.Metadata.Name+".chartserviceversion.yaml", string(chart))
		}

	}
	return err
}
