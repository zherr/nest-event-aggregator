package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const dbEndpoint = "host=localhost user=postgres dbname=postgres sslmode=disable password="

func TestMain(m *testing.M) {
	err := os.Setenv("NEST_DB_HOST", "localhost")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("NEST_DB_NAME", "postgres")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("NEST_DB_USER", "postgres")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("NEST_DB_PASSWORD", "password")
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func Test_logNestCamEvent(t *testing.T) {
	jsonExample, err := ioutil.ReadFile("fixtures/nest_camera_response.json")
	if err != nil {
		panic(err)
	}
	var exampleNestCameraResponse NestCameraResponse
	json.Unmarshal(jsonExample, &exampleNestCameraResponse)

	logNestCamEvent(exampleNestCameraResponse)

	db, err := gorm.Open("postgres", dbEndpoint)
	if err != nil {
		panic(err)
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

	timeLayout := "2006-01-02T15:04:05.000Z"
	timeStr := "2017-08-30T19:43:37.000Z"
	expectedEventTime, err := time.Parse(timeLayout, timeStr)
	if err != nil {
		t.Error(err)
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
	if !firstNestCamEvent.StartTime.Equal(expectedEventTime) {
		t.Errorf("Expected %v got %v", expectedEventTime, firstNestCamEvent.StartTime)
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
