package dto

type PageResultDTO struct {
	Data          []any  `json:"data"`
	TotalElements int    `json:"count"`
	Page          int    `json:"page"`
	StartIndex    int    `json:"startindex"`
	EndIndex      int    `json:"endindex"`
	Sort          string `json:"sort"`
	Offset        int    `json:"offset"`
}
