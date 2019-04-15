package main
import (
	"net"
	"os"
	"testing"
	"github.com/kataras/go-errors"
	"fmt"
)

func Test(t *testing.T) {
	checkErr := func(err error) {
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	checkErr(Get())
	checkErr(GetNil())
}

func Get() error {
	return errors.New("errr")
}

func GetNil() error {
	return nil
}
func Test1(t *testing.T) {

}
func Test_GetIp(t *testing.T)  {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(addrs)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}