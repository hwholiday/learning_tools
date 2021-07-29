package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestDh(t *testing.T) {
	aP := big.NewInt(509)               //素数
	aG := big.NewInt(5)                 //底数g，任选
	aS := big.NewInt(123)               //随机数
	AA := big.NewInt(0).Exp(aG, aS, aP) //// aG^aS mod aP
	fmt.Println(AA)
	//发给B aG，AA,aP
	bS := big.NewInt(456) //随机数
	BB := big.NewInt(0).Exp(aG, bS, aP)
	fmt.Println(BB)
	BK := big.NewInt(0).Exp(AA, bS, aP)
	fmt.Println("密钥", BK)
	//发给A BB
	AK := big.NewInt(0).Exp(BB, aS, aP)
	fmt.Println("密钥", AK)

}
