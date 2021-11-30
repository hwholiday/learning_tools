package hcode

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"sync"
)

//map[uint32]map[string]string
var code sync.Map

func addCode(c int) Code {
	if _, ok := code.Load(strconv.Itoa(c)); ok {
		panic(fmt.Sprintf("code: %d already exist", c))
	}
	code.Store(strconv.Itoa(c), nil)
	return Code(c)
}

func addDescription(c Code, explanation map[string]string) {
	if _, ok := code.Load(strconv.Itoa(c.Code())); !ok {
		panic(fmt.Sprintf("code: %d not exist", c))
	}
	code.Store(strconv.Itoa(c.Code()), explanation)
}

func Click() {
	code.Range(func(key, value interface{}) bool {
		if value == nil {
			panic(fmt.Sprintf("code: %s not add description", fmt.Sprint(key)))
			return false
		}
		if val, ok := value.(map[string]string); ok {
			if len(val) <= 0 {
				panic(fmt.Sprintf("code: %s not add description", fmt.Sprint(key)))
				return false
			} else {
				if len(val) != LanguageLen {
					panic(fmt.Sprintf("code: %s not add enough description cur %d need %d", fmt.Sprint(key), len(val), LanguageLen))
					return false
				} else {
					return true
				}
			}
		} else {
			panic(fmt.Sprintf("code: %s not add description", fmt.Sprint(key)))
			return false
		}
	})
}

type Code int

type Codes interface {
	Error() string
	Code() int
	Message(lang ...string) string
}

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}
func (e Code) Code() int { return int(e) }

func (e Code) Message(lang ...string) string {
	if cm, ok := code.Load(strconv.Itoa(e.Code())); ok {
		if cm == nil {
			return e.Error()
		}
		if val, ok := cm.(map[string]string); ok {
			var l string
			if len(lang) <= 0 {
				l = EN
			} else {
				l = lang[0]
				if len(l) <= 0 {
					l = EN
				}
			}
			l = strings.ToLower(l)
			if msg, ok := val[l]; ok {
				if len(lang) == 2 {
					return fmt.Sprintf(msg, lang[1])
				}
				return msg
			}
		}
	}
	return e.Error()
}

func EqualError(code Codes, err error) bool {
	return Cause(err).Code() == code.Code()
}

func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}

func String(e string) Code {
	if e == "" {
		return OK
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}
	return Code(i)
}
