package main

import(
    "fmt"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)


func main() {
    // Setup the client for Configmap and deployment.
    clientConfig, err := rest.InCLusterConfig()
    if err != nil {
        panic("Unable to get our client configuration")
    }

    clientset, err := kubernetes.NewForConfig(clientConfig)
    if err != nil {
        panic("Unable to create the clientset")
    }

    // apply the updated deployment
    prompt()
    fmt.Println("Updating deployment...")
    retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
        result, getErr := deploymentsClient.Get(context.TODO(), "http-server", metav1.GetOptions{})
        if getErr != nil {
            panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
        }
        result.Spec.Template.Spec.Containers[0].Env = //Contains the data from the ConfigMap
        _, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
        return updateErr
    })
    if retryErr != nil {
        panic(fmt.Errorf("Update failed: %v", retryErr))
    }
    fmt.Println("Updated deployment...")
}

// Set a watch for changes in the configmap
func WatchforConfigMapChanges () {
    for {
        watcher, err := clientset.CoreV1().ConfigMaps().Watch(context.TODO(),
            metav1.SingleObject(metav1.ObjectMeta{
            }))
    }
}
// Take the configmap data and update the deployment manifast with the data from the Configmap data


