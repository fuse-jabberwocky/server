package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// "github.com/apache/camel-k/pkg/client"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Println("Using kubeconfig file: ", kubeconfig)
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, _ := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := client.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
		LabelSelector: "camel.apache.org/integration=event-timer-sink"})
	if err != nil {
		log.Fatal(err)
	}
	firstPod := pods.Items[0]
	req := client.CoreV1().Pods("default").GetLogs(firstPod.GetName(), &corev1.PodLogOptions{})

	readCloser, err := req.Stream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer readCloser.Close()
	_, err = io.Copy(os.Stdout, readCloser)
	log.Fatal(err)

}
