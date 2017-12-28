package main

import (
	"log"
	"time"
)

func logNestCamEvent(nestCamResponse NestCameraResponse) {
	db, err := getDbConnection()
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
