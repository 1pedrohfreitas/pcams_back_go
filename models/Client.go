package models

import "time"

type Clients struct {
	ID         uint      `json:"id"`
	FullName   string    `json:"fullname"`
	Alias      string    `json:"alias"`
	Status     int       `json:"status"`
	Created_at time.Time `json:"password"`
	Updated_at time.Time `json:"updated_at"`
	CodClient  string    `json:"codclient"`
}
