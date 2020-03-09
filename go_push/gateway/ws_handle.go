package gateway

func (w *WsConnection) WsHandle() {
	var (
		err error
		msg *WSMessage
	)
	for {
		if msg, err = w.ReadMsg(); err != nil {
            w.close()
		}
	}

}
