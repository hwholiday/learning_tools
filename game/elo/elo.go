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

}

func EloRating(elo Elo) (a uint32, b uint32) {
	var (
		Ea float64 //预期A选手的胜负值
		Eb float64 //预期B选手的胜负值
		Ra uint32  //A玩家进行了一场比赛之后的Rating
		Rb uint32  //B玩家进行了一场比赛之后的Rating
	)
	Ea = 1 / (1 + math.Pow(10, float64(elo.B-elo.A)/float64(D)))
	Ea = Decimal(Ea, "%.2f")
	Eb = 1 - Ea
	var Sb float64
	if elo.Sa == 0 {
		Sb = 1
	}
	Ra = elo.A + uint32(Decimal(float64(K)*(elo.Sa-Ea), "%.0f"))
	Rb = elo.B + uint32(Decimal(float64(K)*(Sb-Eb), "%.0f"))
	return Ra, Rb
}

//f 保留2 位小数 %.2f
func Decimal(value float64, f string) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf(f, value), 64)
	return value
}
