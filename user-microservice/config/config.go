package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Test interface {
	GetEnv() string
	GetURI() string
	GetDBName() string
}

type test struct {
	env    string
	uri    string
	dbName string
}

//Load configuration with -env default : develop
func Load() Test {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	localerr := godotenv.Load(".env.local")
	if localerr != nil {
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
	return &test{env: *env, uri: uri, dbName: dbName}

}

func (t *test) GetEnv() string {
	return t.env
}

func (t *test) GetURI() string {
	return t.uri
}
func (t *test) GetDBName() string {
	return t.dbName
}
