package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_executeQuery_eventById(t *testing.T) {
	db, err := getDbConnection()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	defer db.Exec("TRUNCATE TABLE nest_camera_events;")

	expectedWebURL := "www.google.com"
	eventOne := NestCameraEvent{WebURL: expectedWebURL}
	db.Create(&eventOne)

	query := fmt.Sprintf("{eventById(id:%d){web_url}}", eventOne.ID)
	result := executeQuery(query, schema)
	resultJSON, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}

	data := struct {
		Data struct {
			NestCameraEvent NestCameraEvent `json:"eventById"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(resultJSON, &data)
	if err != nil {
		t.Error(err)
	}

	if data.Data.NestCameraEvent.WebURL != expectedWebURL {
		t.Errorf("Expected event WebURL to equal: %s, got: %s", expectedWebURL, data.Data.NestCameraEvent.WebURL)
	}
}

func Test_executeQuery_allEvents(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	startTime, err := time.Parse(timeFormat, "2017-12-28 10:00:00")
	if err != nil {
		t.Error(err)
	}

	eventOne := NestCameraEvent{WebURL: "www.google.com", StartTime: startTime}
	eventTwo := NestCameraEvent{WebURL: "www.nest.com", StartTime: startTime.Add(time.Hour)}
	db, err := getDbConnection()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	defer db.Exec("TRUNCATE TABLE nest_camera_events;")
	db.Create(&eventOne)
	db.Create(&eventTwo)

	query := "{allEvents{id}}"
	result := executeQuery(query, schema)
	resultJSON, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}

	data := struct {
		Data struct {
			NestCameraEvents []NestCameraEvent `json:"allEvents"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(resultJSON, &data)
	if err != nil {
		t.Error(err)
	}

	if len(data.Data.NestCameraEvents) != 2 {
		t.Errorf("Expected two events!")
	}
}

func Test_executeQuery_eventsBetween(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	startTime, err := time.Parse(timeFormat, "2017-12-25 10:00:00")
	if err != nil {
		t.Error(err)
	}

	eventOne := NestCameraEvent{WebURL: "www.google.com", StartTime: startTime}
	eventTwo := NestCameraEvent{WebURL: "www.nest.com", StartTime: startTime.Add(time.Hour)}
	eventThree := NestCameraEvent{WebURL: "www.what.com", StartTime: startTime.Add(time.Hour * 48)}
	db, err := getDbConnection()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	defer db.Exec("TRUNCATE TABLE nest_camera_events;")
	db.Create(&eventOne)
	db.Create(&eventTwo)
	db.Create(&eventThree)

	query := `{eventsBetween(start:"2017-12-25",end:"2017-12-26"){id}}`
	result := executeQuery(query, schema)
	resultJSON, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}

	data := struct {
		Data struct {
			NestCameraEvents []NestCameraEvent `json:"eventsBetween"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(resultJSON, &data)
	if err != nil {
		t.Error(err)
	}

	if len(data.Data.NestCameraEvents) != 2 {
		t.Errorf("Expected two events, got %d!", len(data.Data.NestCameraEvents))
	}

	if data.Data.NestCameraEvents[0].ID != eventOne.ID {
		t.Errorf("First event found is not expected id!")
	}

	if data.Data.NestCameraEvents[1].ID != eventTwo.ID {
		t.Errorf("Second event found is not expected id!")
	}
}
