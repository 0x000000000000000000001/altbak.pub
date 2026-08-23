package main

import "fmt"

func main() {
    orig := func() *struct{} {
        orig := fmt.Sprintf("hi")
        _ = orig
        clone := struct{}{}
        return &clone
    }()
    _ = orig
}
