package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/hkdf"
	"hash"
	"io"
	"os"
	"testing"
)

var PseronA_KDF_Prefix []byte
var PseronB_KDF_Prefix []byte

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

//x3dh
func Test_x3curve25519(t *testing.T) {
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

	b.DH1 = GetCurve25519Key(b.SignedPri, a.IdentityPub)
	b.DH2 = GetCurve25519Key(b.IdentityPri, a.EphemeralPub)
	b.DH3 = GetCurve25519Key(b.SignedPri, a.EphemeralPub)
	b.DH4 = GetCurve25519Key(b.OneTimePri, a.EphemeralPub)

	var aKey = bytes.Join([][]byte{a.DH1[:], a.DH2[:], a.DH3[:], a.DH4[:]}, []byte{})

	var bKey = bytes.Join([][]byte{b.DH1[:], b.DH2[:], b.DH3[:], b.DH4[:]}, []byte{})

	fmt.Println("ADH1", base64.StdEncoding.EncodeToString(a.DH1[:]))
	fmt.Println("ADH2", base64.StdEncoding.EncodeToString(a.DH2[:]))
	fmt.Println("ADH3", base64.StdEncoding.EncodeToString(a.DH3[:]))
	fmt.Println("ADH4", base64.StdEncoding.EncodeToString(a.DH4[:]))

	fmt.Println("BDH1", base64.StdEncoding.EncodeToString(b.DH1[:]))
	fmt.Println("BDH2", base64.StdEncoding.EncodeToString(b.DH2[:]))
	fmt.Println("BDH3", base64.StdEncoding.EncodeToString(b.DH3[:]))
	fmt.Println("BDH4", base64.StdEncoding.EncodeToString(b.DH4[:]))

	fmt.Println("aKey", base64.StdEncoding.EncodeToString(aKey))
	fmt.Println("aKey", base64.StdEncoding.EncodeToString(kdf(aKey)))

	fmt.Println("bKey", base64.StdEncoding.EncodeToString(bKey))
	fmt.Println("bKey", base64.StdEncoding.EncodeToString(kdf(bKey)))
	fmt.Println("x3DH结束")
	fmt.Println("开始计算Signal protocol(双棘轮)")

	for i := 1; i <= 3; i++ {
		aSalt := GetCurve25519Key(a.EphemeralPri, b.EphemeralPub)
		fmt.Println("计算Ａ的salt第 ", i, " 次", base64.StdEncoding.EncodeToString(aSalt[:]))
		fmt.Println("计算A的KEY第 ", i, " 次", base64.StdEncoding.EncodeToString(Signalkdf(aKey, aSalt, "A")))

		bSalt := GetCurve25519Key(b.EphemeralPri, a.EphemeralPub)
		fmt.Println("计算B的salt第 ", i, " 次", base64.StdEncoding.EncodeToString(bSalt[:]))
		fmt.Println("计算B的KEY第 ", i, " 次", base64.StdEncoding.EncodeToString(Signalkdf(bKey, bSalt, "B")))
	}

}

func kdf(data []byte) []byte {
	// create reader
	r := hkdf.New(
		func() hash.Hash {
			return sha256.New()
		},
		data,
		make([]byte, 32), []byte("1"),
	)
	var secret [32]byte
	_, err := r.Read(secret[:])
	if err != nil {
		panic(err)
	}
	return secret[:]
}

func Signalkdf(data []byte, salt [32]byte, tag string) []byte {
	// create reader
	if tag == "A" {
		r := hkdf.New(
			func() hash.Hash {
				return sha256.New()
			},
			append(PseronA_KDF_Prefix[:], data...),
			salt[:], nil,
		)
		var secret [64]byte
		_, err := r.Read(secret[:])
		if err != nil {
			panic(err)
		}
		head := secret[:32]
		PseronA_KDF_Prefix = head
		tail := secret[32:]
		return tail
	} else {
		r := hkdf.New(
			func() hash.Hash {
				return sha256.New()
			},
			append(PseronB_KDF_Prefix[:], data...),
			salt[:], nil,
		)
		var secret [64]byte
		_, err := r.Read(secret[:])
		if err != nil {
			panic(err)
		}
		head := secret[:32]
		PseronB_KDF_Prefix = head
		tail := secret[32:]
		return tail
	}
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
