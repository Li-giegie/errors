package errors

import (
	e "errors"
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	fmt.Println(NewError("i'm error"))
}

func TestNewErrors(t *testing.T) {
	var err = e.New("err 1!")
	fmt.Println(NewErrors("i'm error",err))
}

func TestNewErrorf(t *testing.T) {
	fmt.Println(NewErrorf("%v %v %v","abc",1,true))
}

func TestAppend(t *testing.T) {
	var err = e.New("err 2!")
	fmt.Println(Append(err," i'm error!"))
}

func TestAppends(t *testing.T) {
	var err = e.New("err 2!")
	fmt.Println(Appends(err," i'm error!",12,false))
}
