package main

import (
	_ "github.com/go-sql-driver/mysql"
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

	db, err := gorm.Open("mysql", dbEndpoint)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&NestCameraEvent{})

	var existingEvent NestCameraEvent
	timeFormat := "2006-01-02 15:04:05"
	startTimeStr := nestCamResponse.NestCameraEvent.StartTime.Format(timeFormat)
	endTimeStr := nestCamResponse.NestCameraEvent.EndTime.Format(timeFormat)
	startTime, err := time.Parse(timeFormat, startTimeStr)
	endTime, err := time.Parse(timeFormat, endTimeStr)
	notFound := db.Where(&NestCameraEvent{StartTime: startTime, EndTime: endTime}).First(&existingEvent).RecordNotFound()
	if notFound {
		nestCamResponse.NestCameraEvent.StartTime = startTime
		nestCamResponse.NestCameraEvent.EndTime = endTime
		log.Printf("Event logged | Start: %s | End: %s\n", startTimeStr, endTimeStr)
		db.Create(&nestCamResponse.NestCameraEvent)
	}
}
