package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

const aesSalt = "dA1^&6("

func AesECBEncrypt(key string, plaintext []byte) (ciphertextStr string, err error) {
	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	if len(plaintext)%aes.BlockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}
	commonAeskey := []byte(key + aesSalt)
	var ciphertext []byte
	////128 192  256位的其中一个 长度 对应分别是 16 24  32字节长度
	if len(commonAeskey) != 16 && len(commonAeskey) != 24 && len(commonAeskey) != 32 {
		return "", fmt.Errorf("key size is not 16 or 24 or 32, but %d", len(commonAeskey))
	}
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return "", err
	}
	ciphertext = make([]byte, len(plaintext))
	NewECBEncrypter(block).CryptBlocks(ciphertext, plaintext)
	return hex.EncodeToString(ciphertext), nil
}

func AesECBDecrypt(key string, ciphertextStr string) (plaintext []byte, err error) {
	var ciphertext []byte
	ciphertext, err = hex.DecodeString(ciphertextStr)
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	// ECB mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	commonAeskey := []byte(key + aesSalt)
	////128 192  256位的其中一个 长度 对应分别是 16 24  32字节长度
	if len(commonAeskey) != 16 && len(commonAeskey) != 24 && len(commonAeskey) != 32 {
		return nil, fmt.Errorf("key size is not 16 or 24 or 32, but %d", len(commonAeskey))
	}
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	NewECBDecrypter(block).CryptBlocks(ciphertext, ciphertext)
	plaintext = PKCS5UnPadding(ciphertext)
	return plaintext, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
