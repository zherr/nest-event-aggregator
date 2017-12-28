package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"time"
)

func logNestCamEvent(nestCamResponse NestCameraResponse) {
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

	db, err := gorm.Open("postgres", dbEndpoint)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&NestCameraEvent{})

	var existingEvent NestCameraEvent
	timeFormat := "2006-01-02 15:04:05"
	startTimeStr := nestCamResponse.NestCameraEvent.StartTime.Format(timeFormat)
	startTime, err := time.Parse(timeFormat, startTimeStr)
	notFound := db.Where(&NestCameraEvent{StartTime: startTime}).First(&existingEvent).RecordNotFound()
	if notFound {
		nestCamResponse.NestCameraEvent.StartTime = startTime
		log.Printf("Event logged | Start: %s", startTimeStr)
		db.Create(&nestCamResponse.NestCameraEvent)
	}
}
