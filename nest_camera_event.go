package main

import (
	"time"
)

// NestCameraEvent struct to persist camera events
type NestCameraEvent struct {
	ID               uint       `gorm:"primary_key" json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	HasSound         bool       `json:"has_sound"`
	HasMotion        bool       `json:"has_motion"`
	HasPerson        bool       `json:"has_person"`
	StartTime        time.Time  `sql:"unique_index" json:"start_time"`
	EndTime          *time.Time `json:"end_time"`
	UrlsExpireTime   time.Time  `json:"urls_expire_time"`
	WebURL           string     `sql:"type:text" json:"web_url"`
	AppURL           string     `sql:"type:text" json:"app_url"`
	ImageURL         string     `sql:"type:text" json:"image_url"`
	AnimatedImageURL string     `sql:"type:text" json:"animated_image_url"`
}
