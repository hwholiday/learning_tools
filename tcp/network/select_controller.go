package network



func SwitchController(data []byte, client *TcpClient) {
	if len(data) <= 4 {
		client.Write([]byte("类型错误"))
	}
}
