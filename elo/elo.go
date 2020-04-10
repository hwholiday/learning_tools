package elo

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

const (
	// K is the default K-Factor
	K = 32
	// D is the default deviation
	D = 400
)

type Elo struct {
	//输入值
	A  uint32  //A玩家当前的Rating
	B  uint32  //B玩家当前的Rating
	Sa float64 //实际胜负值，胜=1，平=0.5，负=0   传入值默认A的胜负 / 1 A胜利 B失败 / 0 B胜利 A失败

	//返回值
	Ea float64 //预期A选手的胜负值
	Eb float64 //预期B选手的胜负值
	Ra uint32  //A玩家进行了一场比赛之后的Rating
	Rb uint32  //B玩家进行了一场比赛之后的Rating
}

func EloRating(elo Elo) {
	fmt.Println("A玩家当前的Rating", elo.A)
	fmt.Println("B玩家当前的Rating", elo.B)
	fmt.Println("胜负", elo.Sa)
	elo.Ea = 1 / (1 + math.Pow(10, float64(elo.B-elo.A)/float64(D)))
	elo.Ea, _ = decimal.NewFromFloatWithExponent(elo.Ea, -2).Float64()
	elo.Eb = 1-elo.Ea
	var Sb float64
	if elo.Sa==0{
		Sb=1
	}
	ra,_:=decimal.NewFromFloatWithExponent(float64(K)*(elo.Sa-elo.Ea), -0).Float64()
	rb,_:=decimal.NewFromFloatWithExponent(float64(K)*(Sb-elo.Eb), -0).Float64()
	elo.Ra = elo.A + uint32(ra)
	elo.Rb = elo.B + uint32(rb)
	fmt.Println("预期A选手的胜负值", elo.Ea)
	fmt.Println("预期B选手的胜负值", elo.Eb)
	fmt.Println("A玩家进行了一场比赛之后的Rating", elo.Ra)
	fmt.Println("B玩家进行了一场比赛之后的Rating", elo.Rb)
}
