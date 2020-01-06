package v5_service


type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginAck struct {
	Token string `json:"token"`
}
