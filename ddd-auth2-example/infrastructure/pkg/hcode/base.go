package hcode

const (
	EN = "en"
)

const LanguageLen = 0

var (
	OK                        = addCode(200)
	ServerErr                 = addCode(300)
	TranErr                   = addCode(301)
	ParameterErr              = addCode(304)
	RedisExecErr              = addCode(305)
	ResourcesNotFindErr       = addCode(306)
	SysParameterErr           = addCode(307)
	MgoExecErr                = addCode(308)
	TokenValidErr             = addCode(309)
	ResourcesAlreadyExistsErr = addCode(310)
)

func init() {

}
