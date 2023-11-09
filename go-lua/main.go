package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

var script = `function costPrice()
    if env.a=="aa" and env.b> 10  then
        return 2
    elseif env.a=="bb" and env.bindTime>1  and env.bindTime<9  then 
        return 3
    end
    return 4
end`

func main() {
	L := lua.NewState()
	defer L.Close()
	env := L.NewTable()
	env.RawSetString("a", lua.LString("aa"))
	env.RawSetString("b", lua.LNumber(20))
	L.SetGlobal("env", env)
	err := L.DoString(script)
	if err != nil {
		return
	}
	err = L.CallByParam(lua.P{
		Fn:      L.GetGlobal("costPrice"),
		NRet:    1,
		Protect: true,
	})
	if err != nil {
		return
	}
	ret := L.Get(-1)
	L.Pop(1)
	fmt.Println("ret.String()", ret.String())

}
