package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func getDbConnection() (*gorm.DB, error) {
	dbHost, present := os.LookupEnv("NEST_DB_HOST")
	if !present {
		log.Fatalln("NEST_DB_ENDPOINT not set. Please see README for more details.")
	}
	dbUser, present := os.LookupEnv("NEST_DB_USER")
	if !present {
		log.Fatalln("NEST_DB_USER not set. Please see README for more details.")
	}
	dbName, present := os.LookupEnv("NEST_DB_NAME")
	if !present {
		log.Fatalln("NEST_DB_NAME not set. Please see README for more details.")
	}
	dbPassword, present := os.LookupEnv("NEST_DB_PASSWORD")
	if !present {
		log.Fatalln("NEST_DB_PASSWORD not set. Please see README for more details.")
	}

	dbEndpoint := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPassword)

	return gorm.Open("postgres", dbEndpoint)
}
