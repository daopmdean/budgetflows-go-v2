package model

import "time"

type HealthStatus struct {
	Message     string    `json:"message"`
	StartedTime time.Time `json:"timeStarted"`
	TimeNow     time.Time `json:"timeNow"`
	AppName     string    `json:"appName"`
	Flag        string    `json:"flag"`
}
