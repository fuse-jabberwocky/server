package stimzi

import "encoding/json"

// KafkaTopic represents a Strimzi Kafka Topic
type KafkaTopic struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Labels struct {
			StrimziIoCluster string `json:"strimzi.io/cluster"`
		} `json:"labels"`
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Config struct {
		} `json:"config"`
		Partitions int    `json:"partitions"`
		Replicas   int    `json:"replicas"`
		TopicName  string `json:"topicName"`
	} `json:"spec"`
}

// FromUnstructuredObject returns a KafkaTopic from a Kubernetes UnstructuredObject
func FromUnstructuredObject(obj interface{}) (result *KafkaTopic, err error) {
	marhaled, _ := json.Marshal(obj)
	err = json.Unmarshal(marhaled, &result)

	return
}
