package builder

import (
	"hub-gen-auto/pkg/types"
	"hub-gen-auto/pkg/utils"
)

func Build(clusterName string, manifests []types.Manifest) error {
	var err error
	for _, manifest := range manifests {
		utils.WriteToFileAsYaml("build/"+clusterName+"/charts/"+manifest.Namespace+"/"+manifest.Namespace+".chartserviceversion.yaml", manifest.Chart)
		utils.WriteToFileAsYaml("build/"+clusterName+"/charts/"+manifest.Namespace+"/"+manifest.Namespace+".package.yaml", manifest.Package)
		for _, experiment := range manifest.Experiments {
			utils.WriteToFileAsYaml("build/"+clusterName+"/charts/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/experiment.yaml", experiment.ChaosExperiment)
			utils.WriteToFileAsYaml("build/"+clusterName+"/charts/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/engine.yaml", experiment.ChaosEngine)
			utils.WriteToFileAsYaml("build/"+clusterName+"/charts/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/rbac.yaml", experiment.RBAC)
			utils.WriteToFileAsYaml("build/"+clusterName+"/charts/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/rbac-psp.yaml", experiment.PSP)
			utils.WriteToFileAsYaml("build/"+clusterName+"/charts/"+manifest.Namespace+"/"+experiment.ChaosExperiment.Metadata.Name+"/"+experiment.ChaosExperiment.Metadata.Name+".chartserviceversion.yaml", experiment.ChaosExperiment)
			utils.WriteToFile("build/"+clusterName+"/charts/"+manifest.Namespace+"/icons/"+experiment.ChaosExperiment.Metadata.Name+".png", string(experiment.Icon))
		}

	}
	return err
}
