package actions


import (
    "context"
    "log"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func List(resource *string, namespace *string) {

    if *resource == "deploy" {
        listDeployments(namespace)
    } else {
        listPods(namespace)
    }
}


func listDeployments(namespace *string) {
    deploycl := ClientSet.AppsV1().Deployments(*namespace)
    deployments, err := deploycl.List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        log.Println("me do panic")
        panic(err)
    }
    for _, item := range deployments.Items {
        log.Println("Deploy Name: ", item.GetName())
        log.Println("NameSpace: ", item.GetNamespace())
        log.Println("UID: ", item.GetUID())
        log.Println("-----------------")
    }
}


func listPods(namespace *string) {
    pods, err := ClientSet.CoreV1().Pods(*namespace).List(context.Background(), metav1.ListOptions{})
    if err != nil {
        panic(err)
    }
    for _, item := range pods.Items {
        log.Println("Pod Name: ", item.GetName())
        log.Println("NameSpace: ", item.GetNamespace())
        for _, status := range item.Status.ContainerStatuses {
            log.Println("Status Ready: ", status.Ready)
        }
        for _, container := range item.Spec.Containers {
            log.Println("Container Image: ", container.Image)
        }
        log.Println("-----------------")
    }
}
