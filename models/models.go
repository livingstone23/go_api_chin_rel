package models

type Team struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`
}

func (b Team) TableName() string {
	return "teams"
}

type Player struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	TeamID int `db:"team_id" json:"teamId"`
	Teams Team `db:"team"json:"team"`
}

func (b Player) TableName() string {
	return "players"
}

type PlayerPicture struct {
	Id int `db:"id" json:"id"`
	Picture string `db:"picture" json:"picture"`
	PlayerID int `db:"player_id" json:"playerId"`
	Player Player `db:"player" json:"player"`
}

func (b PlayerPicture) TableName() string {
	return "player_pictures"
}
