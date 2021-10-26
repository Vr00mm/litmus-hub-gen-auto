package resources

import (
	"context"
	"fmt"
//	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"encoding/json"
	"github.com/buger/jsonparser"
//	"byte"
)

var (
	// ResourceTypes represents the set of resource types.
	// Resouces are grouped by the same level of abstraction.
	ResourceTypes   = []string{"hpa cronjob", "deploy job", "sts ds rs", "pod", "pvc", "svc", "ing"}
	normalizedNames = map[string]string{
		"ns":      "namespace",
		"svc":     "service",
		"pvc":     "persistentvolumeclaim",
		"pod":     "po",
		"sts":     "statefulset",
		"ds":      "daemonset",
		"rs":      "replicaset",
		"deploy":  "deployment",
		"job":     "job",
		"cronjob": "cj",
		"ing":     "ingress",
		"hpa":     "horizontalpodautoscaler"}
)

type Object struct {
	Type		string
	JsonData	string
}

// Resources represents the k8s resources
type Resources struct {
	clientset kubernetes.Interface
	Namespace string
	Objects   []Object
}

func (obj *Object) GetLabel(label string) (string, bool) {
	if data, _, _, err := jsonparser.Get([]byte(obj.JsonData), "person", "name", label); err == nil {
 		return string(data), true
	}
	return "", false
}

// NewResources resturns Resources for the namespace
func GetResources(clientset kubernetes.Interface) ([]*Resources, error) {
	var err error
	var result []*Resources

	// Get namespaces list
	namespaces, err := GetNamespaces(clientset)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get k8s namespaces: %v\n", err)
	}

	for _, namespace := range namespaces.Items {
		res := &Resources{clientset: clientset, Namespace: namespace.Name}

		// statefulset
		statefulsets, err := clientset.AppsV1().StatefulSets(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to get statefulsets in namespace %q: %v", namespace.Name, err)
		}

		// deployment
		deployments, err := clientset.AppsV1().Deployments(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to get deployments in namespace %q: %v", namespace.Name, err)
		}

		for _, statefulset := range statefulsets.Items {
			var object Object
			jsonData, err := json.Marshal(statefulset)
			if err != nil {
			    fmt.Println(err)
			    continue
			}
                        object.Type = "statefulset"
			object.JsonData = string(jsonData)
			res.Objects = append(res.Objects, object)
		}

                for _, deployment := range deployments.Items {
                        var object Object
                        jsonData, err := json.Marshal(deployment)
                        if err != nil {
                            fmt.Println(err)
                            continue
                        }
                        object.Type = "deployment"
			object.JsonData = string(jsonData)

                        res.Objects = append(res.Objects, object)
                }
		result = append(result, res)

	}

	return result, nil
}

// NewResources resturns Resources for the namespace
func GetNamespaces(clientset kubernetes.Interface) (*corev1.NamespaceList, error) {
	var err error
	var namespaces *corev1.NamespaceList

	// namespaces
	namespaces, err = clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get namespaces : %v", err)
	}

	return namespaces, nil
}
