package Effect_Console

import (
	"fmt"
	"strings"
	"sync/atomic"
)

var indentation int32 = 0

func getIndent() string {
	ind := atomic.LoadInt32(&indentation)
	if ind < 0 {
		ind = 0
	}
	return strings.Repeat("  ", int(ind))
}

func Log(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Warn(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Error(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Info(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Debug(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Time(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func TimeLog(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func TimeEnd(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Clear(_ interface{}) interface{} {
	return nil
}

func Group(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	atomic.AddInt32(&indentation, 1)
	return nil
}

func GroupCollapsed(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	atomic.AddInt32(&indentation, 1)
	return nil
}

func GroupEnd(_ interface{}) interface{} {
	ind := atomic.LoadInt32(&indentation)
	if ind > 0 {
		atomic.AddInt32(&indentation, -1)
	}
	return nil
}
