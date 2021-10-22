package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

        "hub-gen-auto/pkg/generator"
        "hub-gen-auto/pkg/clients"
	"hub-gen-auto/pkg/resources"
)

const (
	defaultNamespace   = "default"
	defaultOutFile     = "data.csv"
	descNamespaceOpt   = "namespace to parse"
	descOutFileOpt     = "output filename"
	descShortOptSuffix = " (shorthand)"
//        availableTests = ["container-kill", "pod-cpu-hog-exec", "pod-delete", "pod-dns-error", "pod-dns-spoof", "pod-io-stress", "pod-memory-hog-exec", "pod-network-corruption", "pod-network-duplication", "pod-network-latency", "pod-network-loss", "volume-fill"]
)

var (
	dir       string
        kubeconfigs []string
	outFile   string
)

func init() {
	var kubeconfig string
	if kubeconfigEnv := os.Getenv("KUBECONFIG"); kubeconfigEnv != "" {
                kubeconfigs = strings.Split(kubeconfigEnv, ":")
	} else if home := os.Getenv("HOME"); home != "" {
		kubeconfigs = append(kubeconfigs, filepath.Join(home, ".kube", "config"))
        } else {
		flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.StringVar(&outFile, "outfile", defaultOutFile, descOutFileOpt)
	flag.StringVar(&outFile, "o", defaultOutFile, descOutFileOpt+descShortOptSuffix)
	flag.Parse()
}

func main() {
        for _, kubeconfig := range kubeconfigs {

            // Get ClusterName from filename
            clusterName := kubeconfig[strings.LastIndex(kubeconfig, "_")+1:]

            fmt.Printf("Connect to cluster: %s\n", clusterName)
            // Create Kubernetes Client
            clientset, err := clients.InitKubeClient(kubeconfig)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Failed to init k8s config: %v\n", err)
                continue
            }

            results, err := resources.GetResources(clientset)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Cannot get resources: %v\n", err)
                continue
            }

            generator.Generate(clusterName, results)
        }
        os.Exit(0)


}
