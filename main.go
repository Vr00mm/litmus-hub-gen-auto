package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"hub-gen-auto/pkg/clients"
	"hub-gen-auto/pkg/generator"
	"hub-gen-auto/pkg/templator"
	"hub-gen-auto/pkg/builder"
	"hub-gen-auto/pkg/resources"
	"hub-gen-auto/pkg/utils"
)

const (
	defaultNamespace = "default"
	defaultOutFile   = "data.csv"

	defaultContainerRuntime = "docker"
	defaultSocketPath       = "/var/run/docker.sock"
	defaultContainerPath    = "/var/lib/docker/containers"
	defaultExperimentLib    = "pumba"

	descContainerRuntimeOpt = "It supports docker, containerd, and crio runtimes. The default value is docker"
	descSocketPathOpt       = "It contains path of docker socket file by default(/var/run/docker.sock). For other runtimes provide the appropriate path."
	descContainerPathOpt    = "It contains path of docker volumes default(/var/lib/docker/containers). For other runtimes provide the appropriate path."
	descExperimentLibOpt    = "It supports litmus and pumba default(litmus). If pumba is not available for the experiments, it will fallback to litmus"

	descOutFileOpt = "output filename"

	descShortOptSuffix = " (shorthand)"

//        availableTests = ["container-kill", "pod-cpu-hog-exec", "pod-delete", "pod-dns-error", "pod-dns-spoof", "pod-io-stress", "pod-memory-hog-exec", "pod-network-corruption", "pod-network-duplication", "pod-network-latency", "pod-network-loss", "volume-fill"]
)

var (
	dir              string
	kubeconfigs      []string
	outFile          string
	containerRuntime string
	sockerPath       string
	containerPath    string
	experimentLib    string
)

func init() {

	utils.WriteArrayToFile([]string{"ligne1", "ligne2"}, "data.csv")
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
	flag.StringVar(&containerRuntime, "containerRuntime", defaultContainerRuntime, descContainerRuntimeOpt)
	flag.StringVar(&containerRuntime, "r", defaultContainerRuntime, descContainerRuntimeOpt+descShortOptSuffix)
	flag.StringVar(&sockerPath, "socketPath", defaultSocketPath, descSocketPathOpt)
	flag.StringVar(&sockerPath, "s", defaultSocketPath, descSocketPathOpt+descShortOptSuffix)
	flag.StringVar(&containerPath, "containerPath", defaultContainerPath, descContainerPathOpt)
	flag.StringVar(&containerPath, "c", defaultContainerPath, descContainerPathOpt+descShortOptSuffix)
	flag.StringVar(&experimentLib, "experimentLib", defaultExperimentLib, descExperimentLibOpt)
	flag.StringVar(&experimentLib, "l", defaultExperimentLib, descExperimentLibOpt+descShortOptSuffix)
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

		manifest, err := generator.Generate(clusterName, results)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot generate experiments parameters: %v\n", err)
			continue
		}

		templates, err := templator.Template(manifest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot templates experiments: %v\n", err)
			continue
		}

		err := builder.Build(templates)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot write experiments to disk: %v\n", err)
			continue
		}
	}
	os.Exit(0)

}
