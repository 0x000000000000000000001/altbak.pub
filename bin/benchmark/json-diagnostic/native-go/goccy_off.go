//go:build !goccy

package main

import "errors"

func goccyUnmarshal(data []byte, target any) error {
	return errors.New("built without the goccy tag")
}
