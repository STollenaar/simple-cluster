package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	clientSet *kubernetes.Clientset
)

// We initialize our clientSet to be able to utilize the kubernetes API
func init() {
	config := ctrl.GetConfigOrDie()
	clientSet = kubernetes.NewForConfigOrDie(config)
}

// Listing pods in given namespace and returing the pods
func ListPods(namespace string) []coreV1.Pod {
	pods, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	return pods.Items
}

// Going through the pods, getting the logs and printing any errors
func GetErrorLogs(pods []coreV1.Pod) {
	for _, pod := range pods {
		req := clientSet.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &coreV1.PodLogOptions{})
		podLogs, err := req.Stream(context.TODO())
		if err != nil {
			fmt.Printf("Error opening stream for %s in %s\n", pod.Name, pod.Namespace)
			continue
		}
		defer podLogs.Close()

		buffer := new(bytes.Buffer)
		_, err = io.Copy(buffer, podLogs)
		if err != nil {
			fmt.Printf("Error in copy information from podLogs to buffer for %s in %s\n", pod.Name, pod.Namespace)
		}
		str := buffer.String()
		lines := strings.Split(str, "\n")

		for _, line := range lines {
			if strings.Contains(strings.ToLower(line), "error") || strings.Contains(strings.ToLower(line), "exception") {
				fmt.Println(line)
			}
		}
	}
}
