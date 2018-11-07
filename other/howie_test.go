package main
import (
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