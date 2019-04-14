package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"io"
	"os"
)

type PseronB struct {
	IdentityPri  [32]byte //身份密钥对//IPK
	IdentityPub  [32]byte
	SignedPri    [32]byte //已签名的预共享密钥//SPK
	SignedPub    [32]byte
	OneTimePri   [32]byte //一次性预共享密钥//OPK
	OneTimePub   [32]byte
	EphemeralPri [32]byte //一个临时密钥对//EPK
	EphemeralPub [32]byte
	DH1          [32]byte
	DH2          [32]byte
	DH3          [32]byte
	DH4          [32]byte
}
type PseronA struct {
	IdentityPri  [32]byte //身份密钥对//IPK
	IdentityPub  [32]byte
	SignedPri    [32]byte //已签名的预共享密钥//SPK
	SignedPub    [32]byte
	OneTimePri   [32]byte //一次性预共享密钥//OpK
	OneTimePub   [32]byte
	EphemeralPri [32]byte //一个临时密钥对//EPK
	EphemeralPub [32]byte
	DH1          [32]byte
	DH2          [32]byte
	DH3          [32]byte
	DH4          [32]byte
	//DH1 = DH(IPK-A私钥, SPK-B公钥)
	//DH2 = DH(EPK-A私钥, IPK-B公钥)
	//DH3= DH(EPK-A私钥, SPK-B公钥)
	//DH4 = DH(IPK-A私钥, EPK--B公钥)
}

func main() {
	var a PseronA
	a.IdentityPri, a.IdentityPub = GetCurve25519KeypPair()
	a.SignedPri, a.SignedPub = GetCurve25519KeypPair()
	a.OneTimePri, a.OneTimePub = GetCurve25519KeypPair()
	a.EphemeralPri, a.EphemeralPub = GetCurve25519KeypPair()

	var b PseronB
	b.IdentityPri, b.IdentityPub = GetCurve25519KeypPair()
	b.SignedPri, b.SignedPub = GetCurve25519KeypPair()
	b.OneTimePri, b.OneTimePub = GetCurve25519KeypPair()
	b.EphemeralPri, b.EphemeralPub = GetCurve25519KeypPair()
	//DH1 = DH(IPK-A私钥, SPK-B公钥)
	//DH2 = DH(EPK-A私钥, IPK-B公钥)
	//DH3= DH(EPK-A私钥, SPK-B公钥)
	//DH4 = DH(IPK-A私钥, OPK--B公钥)
	a.DH1 = GetCurve25519Key(a.IdentityPri, b.SignedPub)
	a.DH2 = GetCurve25519Key(a.EphemeralPri, b.IdentityPub)
	a.DH3 = GetCurve25519Key(a.EphemeralPri, b.SignedPub)
	a.DH4 = GetCurve25519Key(a.EphemeralPri, b.OneTimePub)
	var aKey = bytes.Join([][]byte{a.DH1[:], a.DH2[:], a.DH3[:], a.DH4[:]}, []byte{})
	fmt.Println("aKey", aKey)
	fmt.Println("aKey", base64.StdEncoding.EncodeToString(aKey))
	fmt.Println("aKey", len(aKey))

	b.DH1 = GetCurve25519Key(b.SignedPri, a.IdentityPub)
	b.DH2 = GetCurve25519Key(b.IdentityPri, a.EphemeralPub)
	b.DH3 = GetCurve25519Key(b.SignedPri, a.EphemeralPub)
	b.DH4 = GetCurve25519Key(b.OneTimePub, b.EphemeralPub)
	var bKey = bytes.Join([][]byte{b.DH1[:], b.DH2[:], b.DH3[:], b.DH4[:]}, []byte{})
	fmt.Println("bKey", bKey)
	fmt.Println("bKey", base64.StdEncoding.EncodeToString(bKey))
	fmt.Println("bKey", len(bKey))

}

func GetCurve25519KeypPair() (Aprivate, Apublic [32]byte) {
	//产生随机数
	if _, err := io.ReadFull(rand.Reader, Aprivate[:]); err != nil {
		os.Exit(0)
	}
	curve25519.ScalarBaseMult(&Apublic, &Aprivate)
	return
}

func GetCurve25519Key(private, public [32]byte) (Key [32]byte) {
	//产生随机数
	curve25519.ScalarMult(&Key, &private, &public)
	return
}
