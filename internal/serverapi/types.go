package serverapi

type RespLogin struct {
	ID       string `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

type RespError struct {
	Error string `json:"error"`
}
