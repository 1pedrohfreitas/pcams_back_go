package models

import "time"

type CamAccessPermission struct {
	ID                  uint      `json:"id"`
	CamId               int       `json:"camid"`
	Alias               string    `json:"alias"`
	StartPermissionDate time.Time `json:"startpermissiondate"`
	StopPermissionDate  time.Time `json:"stoppermissiondate"`
	DurationPermitions  int       `json:"durationpermitions"`
}
