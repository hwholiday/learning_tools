package goconvey

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAddV1(t *testing.T) {
	Convey("将两数相加", t, func(c C) {
		c.So(AddV1(1, 2), ShouldEqual, 3)
	})
}
//go test -v