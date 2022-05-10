package models

import "time"

type Cams struct {
	ID           uint      `json:"id"`
	DeviceName   string    `json:"devicename"`
	Alias        string    `json:"alias"`
	StreamType   string    `json:"streamtype"`
	UrlCamStream string    `json:"urlcamstream"`
	Status       int       `json:"status"`
	Created_at   time.Time `json:"password"`
	Updated_at   time.Time `json:"updated_at"`
}
