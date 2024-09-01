package serverapi

import "fmt"

type RespLogin struct {
	ID       string `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

type RespError struct {
	Error string `json:"error"`
}

type RespLocationInfo struct {
	Name     string `json:"name"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type Stats struct {
	Hp             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"special_attack"`
	SpecialDefense int `json:"special_defense"`
	Speed          int `json:"speed"`
}

type Pokemon struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
	Shiny bool   `json:"shiny"`
	Stats Stats  `json:"stats"`
}

func (s Stats) String() string {
	return fmt.Sprintf(
		`	HP: %d			Speed: %d
	Attack: %d		Special Attack: %d
	Defense: %d		Special Defense: %d`, s.Hp, s.Speed, s.Attack, s.SpecialAttack, s.Defense, s.SpecialDefense)
}
