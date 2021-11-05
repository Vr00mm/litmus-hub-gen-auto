package md

import (
	mdCluster "hub-gen-auto/pkg/templates/cluster"
	mdNamespace "hub-gen-auto/pkg/templates/namespace"
	"hub-gen-auto/pkg/types"
	"hub-gen-auto/pkg/utils"
)

func Generate(clusterName string, composant []types.Manifest) {
	clusterTemplate := mdCluster.Generate(struct {
		ClusterName string
		Namespaces  []types.Manifest
	}{clusterName, composant})
	utils.WriteToFile("documentation/clusters/"+clusterName+".md", clusterTemplate)
	for _, namespace := range composant {
		namespaceTemplate := mdNamespace.Generate(struct {
			ClusterName string
			Namespace   types.Manifest
		}{clusterName, namespace})
		utils.WriteToFile("documentation/clusters/"+clusterName+"/"+namespace.Namespace+".md", namespaceTemplate)
	}
}
