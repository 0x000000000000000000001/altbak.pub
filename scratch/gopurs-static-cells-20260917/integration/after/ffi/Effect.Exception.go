package Effect_Exception

import (
	"errors"
	"fmt"
)
func Error(msg string) error { return errors.New(msg) }
func Message(e error) string { return e.Error() }
func Name(e error) string { return "Error" }
func ShowErrorImpl(e error) string { return e.Error() }
func StackImpl(just func(string) interface{}, nothing interface{}, e error) interface{} { return nothing }

func ThrowException(e error, _ interface{}) interface{} {
	panic(e)
}

func CatchException(c func(error) func(interface{}) interface{}, t func(interface{}) interface{}, _ interface{}) (res interface{}) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				res = c(err)(nil)
			} else {
				res = c(fmt.Errorf("%v", r))(nil)
			}
		}
	}()
	return t(nil)
}

func ErrorWithCause(msg string, cause error) error { return errors.New(msg) }
func ErrorWithName(name string, msg string) error { return errors.New(msg) }
