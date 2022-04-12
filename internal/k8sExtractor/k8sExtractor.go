package k8sExtractor

import (
	"fmt"
	"os/exec"
	"sync"
)

const (
	runScript     = "/internal/k8sExtractor/run.sh"
	fileDirPrefix = "offline-kube-objects"
)

type Service struct {
	Path    string
	App_dir string
	Err     *error
	Wg      *sync.WaitGroup
}

// This is a K8s Manifest Extractor
func (s *Service) ExtractK8sObjects() {
	var err error
	fileDir := s.Path + "/fetcher_result/" + fileDirPrefix
	fmt.Printf("[ExtractK8sObjects] Start extracting K8S manifests......\n")
	_, err = exec.Command(s.App_dir+runScript, fileDir).Output()
	if err != nil {
		fmt.Printf("[ExtractK8sObjects] Error in extracting K8S manifests: %v\n", err)
		return
	}
	fmt.Printf("[ExtractK8sObjects] Extracting K8S manifests completed \n")
	return
}
