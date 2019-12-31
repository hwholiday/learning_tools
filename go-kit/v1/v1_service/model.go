package v1_service

type Add struct {
	A int `json:"a"`
	B int `json:"b"`
}

type AddAck struct {
	Res int `json:"res"`
}
