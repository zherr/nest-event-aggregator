package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
)

func logNestCamEvent(nestCamResponse NestCameraResponse) {
	dbEndpoint, present := os.LookupEnv("NEST_DB_ENDPOINT")
	if !present {
		log.Fatalln("NEST_DB_ENDPOINT not set. Please see README for more details.")
	}

	db, err := gorm.Open("postgres", dbEndpoint)
	if err != nil {
		panic("failed to connect database")
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
