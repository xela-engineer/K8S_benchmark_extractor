package kubehunter

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/thedevsaddam/gojsonq"
)

type Service struct {
	Path    string
	App_dir string
	Err     *error
	Wg      *sync.WaitGroup
}

func (s *Service) Execkubehunter() {
	var out []byte
	var jobUid string
	var getPodCMD string
	var podName string
	var getLogsCMD string
	createJobCMD := "kubectl create -f " + s.App_dir + "/assets/yamls/kube-hunter-job.yaml"
	removeJobCMD := "kubectl delete -f " + s.App_dir + "/assets/yamls/kube-hunter-job.yaml"
	getJobIdCMD := "kubectl get job kube-hunter -ojson "                   //| jq '.metadata.labels.\"controller-uid\" '
	out, *(s.Err) = exec.Command("/bin/bash", "-c", removeJobCMD).Output() // Delete kubehunter job on target k8s
	time.Sleep(2 * time.Second)
	out, *(s.Err) = exec.Command("/bin/bash", "-c", createJobCMD).Output() // Create kubehunter job on target k8s

	status_cmd := "kubectl get pods --all-namespaces | grep kube-hunter |awk  '{ print $4}'" // get status
	fmt.Printf("[kubehunter] Waiting pod to Completed status ......\n")
	// fmt.Printf("Waiting for kubehunter pod to Completed status: %v\n", *(s.Err))
	for i := 1; i < 50; { // the kube-bench pod need times to run.
		out, *(s.Err) = exec.Command("/bin/bash", "-c", status_cmd).Output() // Get job status
		if strings.TrimSuffix(string(out), "\n") == "Completed" {            // Wait until Completed status
			fmt.Printf("[kubehunter] Pod Completed\n")
			break
		}
	}
	out, *(s.Err) = exec.Command("/bin/bash", "-c", getJobIdCMD).Output() // Get the kube-hunter job ID from k8s
	uid := gojsonq.New().JSONString(string(out)).Find("metadata.labels.controller-uid")
	jobUid = uid.(string)
	//fmt.Printf("Result: %v\n", uid)
	getPodCMD = "kubectl get pods -l controller-uid=" + jobUid + " -o name | awk -F \"/\" '{print $2}'"
	out, *(s.Err) = exec.Command("/bin/bash", "-c", getPodCMD).Output() // Get the kube-hunter pod name by label
	if *(s.Err) != nil {
		fmt.Printf("Error in : %v\n", *(s.Err))
		return
	}
	saveFilePath := s.Path + "/fetcher_result/kube-hunter-result.json"
	podName = strings.TrimSuffix(string(out), "\n")
	getLogsCMD = "kubectl logs " + podName + " > " + saveFilePath
	out, *(s.Err) = exec.Command("/bin/bash", "-c", getLogsCMD).Output() // Get the kube-hunter pod's log and send to /tmp/...

	if *(s.Err) != nil {
		fmt.Printf("Error in running kubehunter in k8s (Get the kube-hunter pod's log): %v\n", *(s.Err))

	}

	exec.Command("/bin/bash", "-c", removeJobCMD).Output() // Delete kubehunter job on target k8s
	if s.Wg != nil {
		s.Wg.Done()
	}
	return
}
