package domain

var GameDatabase []Game

type Game struct {
	Id    int     `son:"id"`
	Name  string  `json:"game_title"`
	Price float64 `json:"price"`
}
