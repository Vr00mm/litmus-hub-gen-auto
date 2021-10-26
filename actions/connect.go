package actions

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

var ClientSet *kubernetes.Clientset

func init() {
	log.Println("Start Kubernetes Connection")
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = filepath.Join(
			os.Getenv("HOME"), ".kube", "config",
		)
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Println("fatal build")
		log.Fatal(err)
	}
	ClientSet = kubernetes.NewForConfigOrDie(config)

}
