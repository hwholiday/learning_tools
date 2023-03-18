package dove

import (
	"github.com/hwholiday/ghost/utils"
	"github.com/rs/zerolog/log"
)

const (
	DefaultWsPort = ":8081"
)
const (
	DefaultConnAcceptCrcId uint64 = 1
	DefaultConnCloseCrcId  uint64 = 2
)

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
)

var doveMode = ReleaseMode

var DefaultConnMax int64 = 10000

func SetConnMax(value int64) {
	DefaultConnMax = value
}
func SetMode(value string) {
	switch value {
	case DebugMode:
		doveMode = DebugMode
	case ReleaseMode:
		doveMode = ReleaseMode
	default:
		doveMode = ReleaseMode
	}
}

func ModeName() string {
	return doveMode
}

func setup() {
	utils.SetUpGlobalZeroLogConf(doveMode == DebugMode)
	log.Info().Str("dove run mode :", ModeName()).Send()
}
