package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type environmentVariables struct {
	Path string `envconfig:"K8SRESULT_PATH"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	var env = environmentVariables{Path: path}
	if err := envconfig.Process("APP", &env); err != nil { // Extract env. variables fomr ./.env
		panic(err)
	}
	if env.Path == "" {
		env.Path = path
	}

	err = os.Mkdir(env.Path+"/fetcher_result/", 0755)
	check(err)

	/*db, err := database.NewDatabase(env.MongoURI, env.DatabaseName, env.DatabaseUsername, env.DatabasePassword)
	if err != nil {
		panic(err)
	}*/

}
