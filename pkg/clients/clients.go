package clients

import (
	"fmt"
        "os"
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/tools/clientcmd"

       )
func InitKubeClient(kubeconfig string) (*kubernetes.Clientset, error) {
        // use the current context in kubeconfig
        var err error
        var clientset *kubernetes.Clientset

        if _, err := os.Stat(kubeconfig); os.IsNotExist(err) {
                fmt.Fprintf(os.Stderr, "[WARNING] Kubeconfig not found %q: %v\n", kubeconfig, err)
                return clientset, err
        }


        config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
        if err != nil {
                fmt.Fprintf(os.Stderr, "[WARNING] Failed to build config from %q: %v\n", kubeconfig, err)
                return clientset, err

        }

        // create the clientset
        clientset, err = kubernetes.NewForConfig(config)
        if err != nil {
                fmt.Fprintf(os.Stderr, "[WARNING] Failed to create client from %q: %v\n", kubeconfig, err)
                return clientset, err
        }

        // test connectivity for k8s cluster
//        _, err = clientset.CoreV1().Namespaces().Get(context.TODO(), defaultNamespace, metav1.GetOptions{})
//        if err != nil {
//                fmt.Fprintf(os.Stderr, "Failed to get namespace %q: %v\n", defaultNamespace, err)
//                os.Exit(1)
//        }
        return clientset, err
}
