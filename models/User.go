package models

import (
	"time"
)

type User struct {
	ID         uint      `json:"id"`
	FullName   string    `json:"fullname"`
	Alias      string    `json:"alias"`
	UserName   string    `json:"username"`
	UserType   int       `json:"usertype"`
	Status     int       `json:"status"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
