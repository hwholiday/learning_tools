package main

import (
	"fmt"
	"net/http"

	"github.com/lucas-clemente/quic-go"

	"github.com/lucas-clemente/quic-go/http3"

	"github.com/google/uuid"
)

const (
	addr = ":8888"
)

func main() {
	setUpQuic()
}

func setUpQuic() {
	mux := http.NewServeMux()
	mux.HandleFunc("/quic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request %+v \n", r.Header)
		_, _ = w.Write([]byte(uuid.New().String()))
	})

	server := http3.Server{
		Server:     &http.Server{Handler: mux, Addr: addr},
		QuicConfig: &quic.Config{},
	}
	fmt.Println("start http3 quic addr " + addr)
	if err := server.ListenAndServeTLS("./cert.pem", "./priv.key"); err != nil {
		panic(err)
	}
}
