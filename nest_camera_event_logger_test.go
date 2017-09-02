package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"testing"
	"time"
)

const dbEndpoint = "root:root@tcp(localhost:3306)/nest_test?parseTime=true"

func TestMain(m *testing.M) {
	err := os.Setenv("NEST_DB_ENDPOINT", dbEndpoint)
	if err != nil {
		panic("Unable to set NEST_DB_ENDPOINT for test")
	}
	os.Exit(m.Run())
}

func Test_logNestCamEvent(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05.000Z"
	timeStr := "2017-08-30T19:43:37.000Z"
	eventTime, err := time.Parse(timeLayout, timeStr)

	if err != nil {
		t.Error(err)
	}

	exampleNestCameraResponse := NestCameraResponse{
		Name:                  "Family Room",
		SoftwareVersion:       "205-600055",
		WhereID:               "12345",
		DeviceID:              "12345",
		StructureID:           "12345",
		IsOnline:              true,
		IsStreaming:           true,
		IsAudioInputEnabled:   true,
		LastIsOnlineChange:    eventTime,
		IsVideoHistoryEnabled: true,
		IsPublicShareEnabled:  false,
		NestCameraEvent: NestCameraEvent{
			HasSound:         false,
			HasMotion:        true,
			HasPerson:        false,
			StartTime:        eventTime,
			EndTime:          &eventTime,
			UrlsExpireTime:   eventTime,
			WebURL:           "www.nest.com",
			AppURL:           "www.nest.com",
			ImageURL:         "www.nest.com",
			AnimatedImageURL: "www.nest.com",
		},
		WhereName:   "12345",
		NameLong:    "1234555",
		WebURL:      "www.nest.com",
		AppURL:      "www.nest.com",
		SnapshotURL: "www.nest.com",
	}

	logNestCamEvent(exampleNestCameraResponse)

	dbEndpoint, _ := os.LookupEnv("NEST_DB_ENDPOINT")
	db, err := gorm.Open("mysql", dbEndpoint)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	defer db.Exec("TRUNCATE TABLE nest_camera_events;")

	var firstNestCamEvent NestCameraEvent
	db.First(&firstNestCamEvent)

	var count int
	db.Model(&firstNestCamEvent).Count(&count)
	if 1 != count {
		t.Errorf("Expected 1 camera event, got %v", count)
	}

	if firstNestCamEvent.HasSound != false {
		t.Errorf("Expected %v got %v", false, firstNestCamEvent.HasSound)
	}
	if firstNestCamEvent.HasMotion != true {
		t.Errorf("Expected %v got %v", true, firstNestCamEvent.HasMotion)
	}
	if firstNestCamEvent.HasPerson != false {
		t.Errorf("Expected %v got %v", false, firstNestCamEvent.HasPerson)
	}
	if !firstNestCamEvent.StartTime.Equal(eventTime) {
		t.Errorf("Expected %v got %v", eventTime, firstNestCamEvent.StartTime)
	}
	if firstNestCamEvent.WebURL != "www.nest.com" {
		t.Errorf("Expected %v got %v", "www.nest.com", firstNestCamEvent.WebURL)
	}

	// It should not log duplicate events
	logNestCamEvent(exampleNestCameraResponse)
	db.Model(&firstNestCamEvent).Count(&count)
	if 1 != count {
		t.Errorf("Expected 1 camera event, got %v", count)
	}

	// Unique by StartDate
	exampleNestCameraResponse.NestCameraEvent.StartTime = time.Now()

	logNestCamEvent(exampleNestCameraResponse)
	var secondNestCamEvent NestCameraEvent
	db.Model(&secondNestCamEvent).Count(&count)
	if 2 != count {
		t.Errorf("Expected 2 camera events, got %v", count)
	}
}
