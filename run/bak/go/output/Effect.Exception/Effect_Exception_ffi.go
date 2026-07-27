package Effect_Exception

import "errors"
func Error(msg string) error { return errors.New(msg) }
func Message(e error) string { return e.Error() }
func Name(e error) string { return "Error" }
func ShowErrorImpl(e error) string { return e.Error() }
func StackImpl(just func(string) interface{}, nothing interface{}, e error) interface{} { return nothing }
func ThrowException(e error) func() interface{} { return func() interface{} { panic(e) } }
func CatchException(c func(error) func() interface{}, t func() interface{}) func() interface{} {
	return func() (res interface{}) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					res = c(err)()
				} else {
					// Fallback if panic is not an error type
					res = c(errors.New("panic"))()
				}
			}
		}()
		return t()
	}
}
func ErrorWithCause(msg string, cause error) error { return errors.New(msg) }
func ErrorWithName(name string, msg string) error { return errors.New(msg) }
