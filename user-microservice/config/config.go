package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	GetEnv() string
	GetURI() string
	GetDBName() string
}

type config struct {
	env    string
	uri    string
	dbName string
}

//Load configuration with -env default : develop
func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	localErr := godotenv.Load(".env.local")
	if localErr != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	env := flag.String("env", "develop", "To switch configurations.")
	flag.Parse()
	var uri string
	if *env == "develop" {
		uri = os.Getenv("MONGO_URI")
	} else {
		uri = os.Getenv("DOCKER_MONGO_URI")
	}
	dbName := os.Getenv("DB_NAME")
	return &config{env: *env, uri: uri, dbName: dbName}

}

func (t *config) GetEnv() string {
	return t.env
}

func (t *config) GetURI() string {
	return t.uri
}
func (t *config) GetDBName() string {
	return t.dbName
}
