package main

import (
	"log"
	"os"

	"github.com/alexshinningsun/K8S_benchmark_extractor/internal/k8sExtractor"
	"github.com/alexshinningsun/K8S_benchmark_extractor/internal/kubebench"
	"github.com/alexshinningsun/K8S_benchmark_extractor/internal/kubehunter"
	"github.com/alexshinningsun/K8S_benchmark_extractor/internal/utils"
	"github.com/kelseyhightower/envconfig"
)

type environmentVariables struct {
	Path     string `envconfig:"K8SRESULT_PATH"`
	App_dir  string
	Platform string `envconfig:"K8S_CLOUD_PLATFORM"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var errPtrHunter *error
	var errPtrKubeBench *error
	var errPtrExtractor *error
	errK8S := utils.TestKubectl() // Test kubectl command
	if !errK8S {
		panic(errK8S)
	}

	path, err := os.Getwd() // Get the current working directory
	if err != nil {
		log.Println(err)
	}

	var env = environmentVariables{Path: path, App_dir: path} // Get environment variable
	if err := envconfig.Process("APP", &env); err != nil {    // Extract env. variables fomr ./.env
		panic(err)
	}
	if env.Path == "" {
		env.Path = path
	}
	if _, err := os.Stat(env.Path + "/fetcher_result/"); os.IsNotExist(err) {
		err = os.Mkdir(env.Path+"/fetcher_result/", 0755) // Create a directory for storing resutl
	}
	check(err)
	//wg = new(sync.WaitGroup)
	errPtrHunter = new(error)
	errPtrKubeBench = new(error)
	errPtrExtractor = new(error)

	sK8sManifestExtractor := &k8sExtractor.Service{ // Part C
		Path:    env.Path,
		App_dir: env.App_dir,
		Err:     errPtrExtractor,
		Wg:      nil,
	}
	sK8sManifestExtractor.ExtractK8sObjects()

	sK8sHunter := &kubehunter.Service{ // Part A
		Path:    env.Path,
		App_dir: env.App_dir,
		Err:     errPtrHunter,
		Wg:      nil,
	}
	sK8sHunter.Execkubehunter()

	sK8sbench := &kubebench.Service{ // Part B
		Path:     env.Path,
		App_dir:  env.App_dir,
		Platform: env.Platform,
		Err:      errPtrKubeBench,
		Wg:       nil,
	}
	sK8sbench.Execkubebench()
	/*db, err := database.NewDatabase(env.MongoURI, env.DatabaseName, env.DatabaseUsername, env.DatabasePassword)
	if err != nil {
		panic(err)
	}*/

}
