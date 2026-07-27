package Effect_Console



import (
	"fmt"
)

func Log(s string) func() interface{} {
	return func() interface{} {
		fmt.Println(s)
		return nil
	}
}

func Warn(s string) func() interface{} {
	return func() interface{} {
		fmt.Println("[WARN]", s)
		return nil
	}
}

func Error(s string) func() interface{} {
	return func() interface{} {
		fmt.Println("[ERROR]", s)
		return nil
	}
}

func Info(s string) func() interface{} {
	return func() interface{} {
		fmt.Println("[INFO]", s)
		return nil
	}
}

func Debug(s string) func() interface{} {
	return Log(s)
}

func Time(s string) func() interface{} {
	return Log(s)
}

func TimeLog(s string) func() interface{} {
	return Log(s)
}

func TimeEnd(s string) func() interface{} {
	return Log(s)
}

func Clear() {
}

func Group(s string) func() interface{} {
	return Log(s)
}

func GroupCollapsed(s string) func() interface{} {
	return Log(s)
}

func GroupEnd() {
	Clear()
}
