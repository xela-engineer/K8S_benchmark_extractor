package kubebench

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Path     string
	App_dir  string
	Platform string
	Err      *error
	Wg       *sync.WaitGroup
}

func (s *Service) Execkubebench() {
	var out []byte
	var podName string
	var getLogsCMD string
	yamlLocation := s.App_dir + "/assets/yamls/"

	switch s.Platform { // Select yaml in kubebench git
	case "GCP":
		yamlLocation += "job-gke.yaml"
	case "AWS":
		yamlLocation += "job-eks.yaml"
	case "Azure":
		yamlLocation += "job-aks.yaml"
	case "other":
		yamlLocation += "job-aks.yaml"
	}
	createJobCMD := "kubectl create -f " + yamlLocation
	removeJobCMD := "kubectl delete -f " + yamlLocation
	getJobCMD := "kubectl get pods --all-namespaces | grep kube-bench- |awk  '{ print $2}' "
	out, *(s.Err) = exec.Command("/bin/bash", "-c", removeJobCMD).Output() // Delete kubehunter job on target k8s
	time.Sleep(2 * time.Second)
	out, *(s.Err) = exec.Command("/bin/bash", "-c", createJobCMD).Output() // Create kubehunter job on target k8s

	status_cmd := "kubectl get pods --all-namespaces | grep kube-bench |awk  '{ print $4}'" // get status
	fmt.Printf("[kubebench] Waiting pod to Completed status ......\n")
	// fmt.Printf("Waiting for kubehunter pod to Completed status: %v\n", *(s.Err))
	for i := 1; i < 50; { // the kube-bench pod need times to run.
		out, *(s.Err) = exec.Command("/bin/bash", "-c", status_cmd).Output() // Get job status
		if strings.TrimSuffix(string(out), "\n") == "Completed" {            // Wait until Completed status
			fmt.Printf("[kubebench] Pod Completed\n")
			break
		}
	}
	out, *(s.Err) = exec.Command("/bin/bash", "-c", getJobCMD).Output() // Get the kube-bench pod name from k8s
	saveFilePath := s.Path + "/fetcher_result/kube-bench-result.json"
	podName = strings.TrimSuffix(string(out), "\n")
	getLogsCMD = "kubectl logs " + podName + " > " + saveFilePath
	out, *(s.Err) = exec.Command("/bin/bash", "-c", getLogsCMD).Output() // Get the kube-hunter pod's log and send to /tmp/...

	if *(s.Err) != nil {
		fmt.Printf("Error in running kubebench in k8s: %v\n", *(s.Err))

	}

	exec.Command("/bin/bash", "-c", removeJobCMD).Output() // Delete kubehunter job on target k8s
	if s.Wg != nil {
		s.Wg.Done()
	}
	return
}
