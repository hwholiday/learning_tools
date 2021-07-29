package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/hkdf"
	"io"
	"testing"
)

func TestKDF(t *testing.T) {
	hash := sha256.New

	secret := []byte("123456")

	salt := make([]byte, hash().Size())
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	info := []byte("user pwd")
	hkdf := hkdf.New(hash, secret, salt, info)
	key := make([]byte, 16)
	if _, err := io.ReadFull(hkdf, key); err != nil {
		panic(err)
	}
	fmt.Println("key", key)
	fmt.Println("key", base64.StdEncoding.EncodeToString(key[:]))
}
