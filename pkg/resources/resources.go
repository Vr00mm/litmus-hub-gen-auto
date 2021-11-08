package resources

import (
	"context"
	"fmt"

	//	appsv1 "k8s.io/api/apps/v1"
	"encoding/json"
	"os"

	"github.com/buger/jsonparser"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	//	"byte"
)

type Object struct {
	Type                 string
	JsonData             string
	UniqueLabels         []string
	GeneratedExperiments []string
}

// Resources represents the k8s resources
type Resources struct {
	clientset kubernetes.Interface
	Namespace string
	Objects   []Object
}

func (obj *Object) GetUniqueLabel() string {
	return obj.UniqueLabels[0]
}

func (obj *Object) AddGeneratedExperiment(experimentName string) {
	obj.GeneratedExperiments = append(obj.GeneratedExperiments, experimentName)
}

func (obj *Object) AddUniqueLabel(label string) {
	obj.UniqueLabels = append(obj.UniqueLabels, label)
}

func (obj *Object) GetLabels() map[string]string {
	var results map[string]string
	//if data, _, _, err := jsonparser.Get([]byte(obj.JsonData), "metadata", "labels"); err == nil {
	if data, _, _, err := jsonparser.Get([]byte(obj.JsonData), "spec", "template", "metadata", "labels"); err == nil {
		json.Unmarshal(data, &results)
		return results
	}
	return results
}

func (obj *Object) GetName() string {
	var data []byte
	if data, _, _, err := jsonparser.Get([]byte(obj.JsonData), "metadata", "name"); err == nil {
		return string(data)
	}
	return string(data)
}

func (obj *Object) GetLabel(label string) (string, bool) {
	if data, _, _, err := jsonparser.Get([]byte(obj.JsonData), "metadata", "labels", label); err == nil {
		return string(data), true
	}
	return "", false
}

func (obj *Object) GetContainers() ([]string, bool) {
	var results []string
	jsonparser.ArrayEach([]byte(obj.JsonData), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		data, _, _, _ := jsonparser.Get(value, "name")
		results = append(results, string(data))
	}, "spec", "template", "spec", "containers")
	return results, true
}

// NewResources resturns Resources for the namespace
func GetResources(clientset kubernetes.Interface) ([]*Resources, error) {
	var err error
	var result []*Resources

	// Get namespaces list
	namespaces, err := GetNamespaces(clientset)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get k8s namespaces: %v\n", err)
		return result, err
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
