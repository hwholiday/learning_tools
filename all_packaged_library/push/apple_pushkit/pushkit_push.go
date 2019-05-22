package apple_pushkit

import (
	"crypto/tls"
	"fmt"
	"net"
	"encoding/hex"
	"encoding/binary"
	"io/ioutil"
	"golang.org/x/crypto/pkcs12"
	"time"
	"bytes"
	"crypto/x509"
	"errors"
	"strings"
)

var (
	ErrExpired = errors.New("certificate has expired or is not yet valid")
)

const (
	Development = "gateway.sandbox.push.apple.com:2195"
	Production  = "gateway.push.apple.com:2195"
)

type PushKit struct {
	FileName   string
	Pwd        string
	serverName string
	cert       tls.Certificate
	hostUrl    string
	isDebug    bool
}

func InitPushKit(filename, pwd string, isDebug bool) (*PushKit, error) {
	var (
		err     error
		cert    tls.Certificate
		pushkit = &PushKit{}
	)
	if cert, err = Load(filename, pwd); err != nil {
		return nil, err
	}
	pushkit.cert = cert
	pushkit.FileName = filename
	pushkit.isDebug = isDebug
	if isDebug {
		pushkit.hostUrl = Development
		pushkit.serverName = strings.Split(Development, ":")[0]
	} else {
		pushkit.serverName = strings.Split(Production, ":")[0]
		pushkit.hostUrl = Production
	}
	return pushkit, nil
}

func (p *PushKit) Push(token string, data []byte) error {
	conf := &tls.Config{
		Certificates: []tls.Certificate{p.cert},
		ServerName:   p.serverName,
	}
	conn, err := net.Dial("tcp", p.hostUrl)
	if err != nil {
		return errors.New(fmt.Sprintf("error Dial: %s", err.Error()))
	}
	tlsconn := tls.Client(conn, conf)
	// 强制握手，以验证身份握手被处理，否则会在第一次读写的时候进行尝试
	err = tlsconn.Handshake()
	if err != nil {
		return errors.New(fmt.Sprintf("error Handshake: %s", err.Error()))
	}
	/*state := tlsconn.ConnectionState()
	fmt.Printf("conn state %v\n", state)*/
	btoken, err := hex.DecodeString(token)
	if err != nil {
		return errors.New(fmt.Sprintf("tls error DecodeString: %s", err.Error()))
	}
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, uint8(1))
	binary.Write(buffer, binary.BigEndian, uint32(1))
	binary.Write(buffer, binary.BigEndian, uint32(time.Now().Unix()+60*60*24))
	binary.Write(buffer, binary.BigEndian, uint16(len(btoken)))
	binary.Write(buffer, binary.BigEndian, btoken)
	binary.Write(buffer, binary.BigEndian, uint16(len(data)))
	binary.Write(buffer, binary.BigEndian, data)
	pdu := buffer.Bytes()
	_, err = tlsconn.Write(pdu)
	if err != nil {
		return errors.New(fmt.Sprintf("tls error Write: %s", err.Error()))
	}
	tlsconn.Close()
	return nil
}

// Load a .p12 certificate from disk.
func Load(filename, password string) (tls.Certificate, error) {
	p12, err := ioutil.ReadFile(filename)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("Unable to load %s: %v", filename, err)
	}
	return Decode(p12, password)
}

// Decode and verify an in memory .p12 certificate (DER binary format).
func Decode(p12 []byte, password string) (tls.Certificate, error) {
	// decode an x509.Certificate to verify
	privateKey, cert, err := pkcs12.Decode(p12, password)
	if err != nil {
		return tls.Certificate{}, err
	}
	if err := verify(cert); err != nil {
		return tls.Certificate{}, err
	}
	// wraps x509 certificate as a tls.Certificate:
	return tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  privateKey,
		Leaf:        cert,
	}, nil
}

// verify checks if a certificate has expired
func verify(cert *x509.Certificate) error {
	_, err := cert.Verify(x509.VerifyOptions{})
	if err == nil {
		return nil
	}
	switch e := err.(type) {
	case x509.CertificateInvalidError:
		switch e.Reason {
		case x509.Expired:
			return ErrExpired
		default:
			return err
		}
	case x509.UnknownAuthorityError:
		return nil
	default:
		return err
	}
}
