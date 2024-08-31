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
