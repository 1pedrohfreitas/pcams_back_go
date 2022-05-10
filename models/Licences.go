package models

import "time"

type Licences struct {
	ID          uint      `json:"id"`
	UnicCode    string    `json:"uniccode"`
	ClientId    int       `json:"clientid"`
	InstallerId int       `json:"installerid"`
	LicenceType string    `json:"licencetype"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
