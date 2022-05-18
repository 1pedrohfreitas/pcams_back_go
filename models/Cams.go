package models

import "time"

type Cams struct {
	ID           uint      `json:"id"`
	DeviceName   string    `json:"devicename"`
	Alias        string    `json:"alias"`
	StreamType   string    `json:"-"`
	UrlCamStream string    `json:"urlcamstream"`
	Status       int       `json:"status"`
	Icon         T_Images  `json:"icon"`
	Created_at   time.Time `json:"-"`
	Updated_at   time.Time `json:"-"`
}
