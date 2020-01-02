package v3_service

type Add struct {
	A int `json:"a"`
	B int `json:"b"`
}

type AddAck struct {
	Res int `json:"res"`
}

type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginAck struct {
	Token string `json:"token"`
}
