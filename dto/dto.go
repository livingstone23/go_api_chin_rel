package dto

type TeamDto struct {
	Name string `json:"name"`
}

type PlayerDto struct {
	Name string `json:"name"`
	Description string `json:"description"`
	TeamId int `json:"teamId"`
}

