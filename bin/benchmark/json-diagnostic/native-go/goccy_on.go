//go:build goccy

package main

import goccyjson "github.com/goccy/go-json"

func goccyUnmarshal(data []byte, target any) error {
	return goccyjson.Unmarshal(data, target)
}
