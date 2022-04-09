package main

import (
	"github.com/kelseyhightower/envconfig"
)

type environmentVariables struct {
	Port             int
	MongoURI         string `required:"true"`
	DatabaseName     string `required:"true"`
	DatabaseUsername string `envconfig:"DATABASE_USERNAME"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD"`
	JWTRealm         string `required:"true"`
	JWTSecretKey     string `required:"true"`
}

func main() {
	var env = environmentVariables{Port: 3001}
	if err := envconfig.Process("APP", &env); err != nil { // Extract env. variables fomr ./.env
		panic(err)
	}

	/*db, err := database.NewDatabase(env.MongoURI, env.DatabaseName, env.DatabaseUsername, env.DatabasePassword)
	if err != nil {
		panic(err)
	}*/

}
