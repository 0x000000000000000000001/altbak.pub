package Test_FileOps

import (
	"os"
)

func WriteFileSync(path string, content string, _ interface{}) interface{} {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

func ReadFileSync(path string, _ interface{}) interface{} {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func LoopE(n int, action func(interface{}) interface{}, _ interface{}) interface{} {
	for i := 0; i < n; i++ {
		action(nil)
	}
	return nil
}
