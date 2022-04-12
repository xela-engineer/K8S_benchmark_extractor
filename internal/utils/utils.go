package utils

import (
	"fmt"
	"os/exec"
)

func TestKubectl() bool {
	cmd := "kubectl get pods --all-namespaces "
	_, err := exec.Command("timeout", "10s", "bash", "-c", cmd).Output() // Get k8s obj json
	if err != nil {
		fmt.Printf("command kubectl timeout or cannot get all namespaces: %v\n", err)
		return false
	}

	return true
}

func NewService() {

}
