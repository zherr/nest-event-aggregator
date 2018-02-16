package main

import (
	"testing"
)

func Test_getDbConnection(t *testing.T) {
	db, err := getDbConnection()
	db.AutoMigrate(&NestCameraEvent{})

	if err != nil {
		t.Error(err)
	}

	if db == nil {
		t.Error("No database connection opened")
	}
}
