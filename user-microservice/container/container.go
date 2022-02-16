package container

import (
	"log"

	"t/db"
)

type Container interface {
	GetEnv() string
	GetURI() string
	GetDB() *db.MongoInstance
}

type container struct {
	env string
	uri string
	db  *db.MongoInstance
}

func NewContainer(env string, uri string, dbName string) Container {
	instance, err := db.Connect(uri, dbName)
	if err != nil {
		log.Fatal(err)
	}
	return &container{env: env, uri: uri, db: instance}
}

func (c *container) GetEnv() string {
	return c.env
}
func (c *container) GetURI() string {
	return c.uri
}
func (c *container) GetDB() *db.MongoInstance {
	return c.db
}
