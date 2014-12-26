package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	yaml "gopkg.in/yaml.v2"
)

type dbConfig struct {
	Database dbase
}

type dbase struct {
	User     string
	Password string
	Name     string
}

// DBInstance returns an instance of the db to use in the application
func DBInstance() *sqlx.DB {
	filePath := os.Getenv("CONFIG")
	if filePath == "" {
		filePath = "./config.yml"
	}
	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Fatal: %s\n", err)
	}

	config := dbConfig{}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalf("Fatal: %s\n", err)
	}

	cdb := config.Database
	configString := fmt.Sprintf("user=%s password=%s dbname=%s port=5432 sslmode=disable", cdb.User, cdb.Password, cdb.Name)
	db, err := sqlx.Connect("postgres", configString)
	if err != nil {
		fmt.Printf("DB connection error: %s\n\n", err)
	}
	return db
}
