package models

type T_Images struct {
	ID        uint   `json:"id,omitempty"`
	SourceImg string `json:"sourceimg"`
	Alias     string `json:"alias"`
	TypeImg   string `json:"typeimg,omitempty"`
}
