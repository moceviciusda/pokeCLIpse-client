package serverapi

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

type Pokemon struct {
	Name           string `json:"name"`
	Level          int    `json:"level"`
	HP             int    `json:"hp"`
	Attack         int    `json:"attack"`
	Defense        int    `json:"defense"`
	SpecialAttack  int    `json:"special_attack"`
	SpecialDefense int    `json:"special_defense"`
	Speed          int    `json:"speed"`
}
