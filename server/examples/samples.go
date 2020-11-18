package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	kamel "github.com/apache/camel-k/pkg/client/camel/clientset/versioned"
)

func main() {
	// Bootstrap k8s configuration from local 	Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Println("Using kubeconfig file: ", kubeconfig)
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	exampleClient, err := kamel.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building example clientset: %v", err)
	}
	ctx := context.Background()

	clientset, _ := kubernetes.NewForConfig(cfg)
	clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})

	kamelets, err := exampleClient.CamelV1alpha1().Kamelets("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing all kamelets: %v", err)
	}

	fmt.Println("List of Kamelets")
	for _, kamelet := range kamelets.Items {
		fmt.Println(kamelet.Name, " - ", kamelet.Labels["camel.apache.org/kamelet.type"])
	}

	kameletbindings, err := exampleClient.CamelV1alpha1().KameletBindings("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing all kamelets: %v", err)
	}

	fmt.Println("List of KameletBindings")
	for _, kameletbinding := range kameletbindings.Items {
		fmt.Println(kameletbinding.Name, " ", kameletbinding.Spec.Source.Ref.Name, " ", kameletbinding.Spec.Sink.Ref.Name)
	}
}
