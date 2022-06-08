package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/lucas-clemente/quic-go/qlog"

	"github.com/lucas-clemente/quic-go/logging"

	"github.com/lucas-clemente/quic-go"

	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	pool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs:            pool,
			InsecureSkipVerify: true,
			KeyLogWriter:       os.Stdout,
		},
		QuicConfig: &quic.Config{
			Tracer: qlog.NewTracer(func(p logging.Perspective, connectionID []byte) io.WriteCloser {

				return os.Stdout
			}),
		},
	}
	defer roundTripper.Close()
	client := &http.Client{
		Transport: roundTripper,
		Timeout:   2 * time.Second,
	}
	var i = 0
	for i < 3 {
		i++
		res, err := client.Get("https://127.0.0.1:8888/quic")
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("收到返回值", string(data))
	}

}
