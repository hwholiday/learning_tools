package elo

import (
	"fmt"
	"math"
	"strconv"
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
	elo.Ea=Decimal(elo.Ea,"%.2f")
	elo.Eb = 1-elo.Ea
	var Sb float64
	if elo.Sa==0{
		Sb=1
	}
	elo.Ra = elo.A + uint32(Decimal(float64(K)*(elo.Sa-elo.Ea),"%.0f"))
	elo.Rb = elo.B + uint32(Decimal(float64(K)*(Sb-elo.Eb),"%.0f"))
	fmt.Println("预期A选手的胜负值", elo.Ea)
	fmt.Println("预期B选手的胜负值", elo.Eb)
	fmt.Println("A玩家进行了一场比赛之后的Rating", elo.Ra)
	fmt.Println("B玩家进行了一场比赛之后的Rating", elo.Rb)
}
//f 保留2 位小数 %.2f
func Decimal(value float64,f string) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf(f, value), 64)
	return value
}