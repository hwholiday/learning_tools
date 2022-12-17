package gocmp

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")
	code := m.Run()
	fmt.Println("TestMain end")
	os.Exit(code)
}

func TestAdd(t *testing.T) {
	c := add(1, 2)
	if !cmp.Equal(c, 3) {
		t.Errorf("add(%d,%d) = %d want %d", 1, 2, c, 3)
	}
}

func add(a, b int) int {
	return a + b + 1
}

func TestErr(t *testing.T) {
	got := testErr()
	if !cmp.Equal(got, nil) {
		t.Errorf("testErr() = %q want nil", got)
	}
}

func testErr() error {
	return errors.New("my is err")
}

type Handler struct {
	Version int
	token   string
}

type LoginDto struct {
	handler *Handler
	Name    string
	Pwd     string
}

func TestStruct(t *testing.T) {
	want := LoginDto{
		handler: &Handler{
			Version: 1,
			token:   "2",
		},
		Name: "name",
		Pwd:  "pwd",
	}
	got := LoginDto{
		handler: &Handler{
			Version: 1,
			token:   "222222",
		},
		Name: "name",
		Pwd:  "pwd",
	}

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(LoginDto{})) {
		t.Errorf("got %+v want %+v", got, want)
	}

	if !cmp.Equal(got, want, cmp.AllowUnexported(LoginDto{}, Handler{})) {
		t.Errorf("got %+v want %+v diff %+v", got, want, cmp.Diff(got, want, cmp.AllowUnexported(LoginDto{}, Handler{})))
	}

	if !cmp.Equal(got.handler, want.handler, cmp.AllowUnexported(Handler{})) {
		t.Errorf("got %+v want %+v diff %+v", got, want, cmp.Diff(got.handler, want.handler, cmp.AllowUnexported(Handler{})))
	}
}

func TestSlice(t *testing.T) {
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}
	if !cmp.Equal(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.SortSlices(func(x, y interface{}) bool { return x.(int) > y.(int) })) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
