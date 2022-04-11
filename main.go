package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type environmentVariables struct {
	Path string

	DatabaseUsername string `envconfig:"DATABASE_USERNAME"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD"`
	JWTRealm         string `required:"true"`
	JWTSecretKey     string `required:"true"`
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

	/*db, err := database.NewDatabase(env.MongoURI, env.DatabaseName, env.DatabaseUsername, env.DatabasePassword)
	if err != nil {
		panic(err)
	}*/

}
