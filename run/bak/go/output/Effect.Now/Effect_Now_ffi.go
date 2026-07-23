package Effect_Now

import (
	"gopurs/output/gopurs_runtime"
	"time"
)

var Now = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	// Retourne le temps en millisecondes
	return gopurs_runtime.Float(float64(time.Now().UnixNano()) / 1e6)
})

var GetTimezoneOffset = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	_, offset := time.Now().Zone()
	// offset est en secondes à l'est d'UTC. getTimezoneOffset demande les minutes à l'ouest.
	return gopurs_runtime.Float(float64(-offset / 60))
})
