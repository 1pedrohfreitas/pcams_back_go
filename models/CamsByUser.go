package models

import "time"

type CamsByUser struct {
	ID                  uint      `json:"id"`
	Alias               string    `json:"alias"`
	UrlCamStream        string    `json:"urlcamstream"`
	Icon                T_Images  `json:"icon"`
	StartPermissionDate time.Time `json:"startpermissiondate"`
	StopPermissionDate  time.Time `json:"stoppermissiondate"`
}
