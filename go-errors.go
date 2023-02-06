package errors

import (
	"bytes"
	e "errors"
	"fmt"
)

func NewError(arg interface{}) error {
	return e.New(fmt.Sprint(arg))
}

func NewErrorf(format string,arg ...interface{}) error {
	return e.New(fmt.Sprintf(format,arg...))
}

func NewErrors(arg ...interface{}) error {
	var buf = new(bytes.Buffer)
	for _, i := range arg {
		buf.WriteString(fmt.Sprint(i))
	}

	return e.New(buf.String())
}

func Append(err error,arg interface{}) error {
	return e.New(err.Error() + fmt.Sprint(arg))
}

func Appends(err error,arg ...interface{}) error {
	var buf = new(bytes.Buffer)
	buf.WriteString(err.Error())
	for _, i := range arg {
		buf.WriteString(fmt.Sprint(i))
	}

	return e.New(buf.String())
}