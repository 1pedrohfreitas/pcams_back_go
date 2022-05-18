package models

import "time"

type Login struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password,omitempty"`
	Status       int       `json:"status"`
	Type         string    `json:"type"`
	Token        string    `json:"token"`
	LastLogin_at time.Time `json:"lastLlgin_at"`
}
