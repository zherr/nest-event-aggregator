package main

import "time"

// NestCameraResponse struct to unmarshal a Nest camera response
type NestCameraResponse struct {
	Name                  string    `json:"name"`
	SoftwareVersion       string    `json:"software_version"`
	WhereID               string    `json:"where_id"`
	DeviceID              string    `json:"device_id"`
	StructureID           string    `json:"structure_id"`
	IsOnline              bool      `json:"is_online"`
	IsStreaming           bool      `json:"is_streaming"`
	IsAudioInputEnabled   bool      `json:"is_audio_input_enabled"`
	LastIsOnlineChange    time.Time `json:"last_is_online_change"`
	IsVideoHistoryEnabled bool      `json:"is_video_history_enabled"`
	IsPublicShareEnabled  bool      `json:"is_public_share_enabled"`
	NestCameraEvent       `json:"last_event"`
	WhereName             string `json:"where_name"`
	NameLong              string `json:"name_long"`
	WebURL                string `json:"web_url"`
	AppURL                string `json:"app_url"`
	SnapshotURL           string `json:"snapshot_url"`
}
