/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"

	"sigs.k8s.io/controller-runtime/pkg/client"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

	strimzi "./go/strimzi"
)

func main() {
	// Bootstrap k8s configuration from local 	Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Println("Using kubeconfig file: ", kubeconfig)
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	c, _ := client.New(cfg, client.Options{})
	// Kafkas
	kafkas := &unstructured.UnstructuredList{}
	kafkas.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "kafka.strimzi.io",
		Kind:    "Kafka",
		Version: "v1beta1",
	})
	_ = c.List(context.Background(), kafkas)

	data, err := yaml.Marshal(kafkas)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Kafkas found:")
	fmt.Printf(string(data))

	// Kafka Topics
	kafkaTopics := &unstructured.UnstructuredList{}
	kafkaTopics.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "kafka.strimzi.io",
		Kind:    "KafkaTopic",
		Version: "v1beta1",
	})
	_ = c.List(context.Background(), kafkaTopics)

	fmt.Println("Kafka Topics found:")
	for _, kafkaTopic := range kafkaTopics.Items {
		parsedTopic, _ := strimzi.FromUnstructuredObject(kafkaTopic.Object)
		fmt.Println(parsedTopic.Metadata.Name)
	}
}
