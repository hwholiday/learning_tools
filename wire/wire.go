//+build wireinject

package main

import "github.com/google/wire"

func NewApp() Msg3 {
	wire.Build(NewMsg3, NewMsg2, NewMsg)
	return Msg3{}
}
