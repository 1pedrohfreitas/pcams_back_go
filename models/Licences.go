package models

import "time"

type Licences struct {
	ID                  uint       `json:"id"`
	UnicCode            string     `json:"uniccode"`
	ClientId            Clients    `json:"clientid"`
	InstallerId         Installers `json:"installerid"`
	LicenceType         string     `json:"licencetype"`
	StartDateValidate   time.Time  `json:"startvalidate"`
	EndDateValidate     time.Time  `json:"endvalidate"`
	PrevDateEndValidate time.Time  `json:"prevendvalidate"`
	Created_at          time.Time  `json:"created_at"`
	Updated_at          time.Time  `json:"updated_at"`
}
