package ffi_loader

import . "github.com/purescript-native/go-runtime"

// The upstream go-ffi length counts code points. PureScript CodeUnits.length
// counts UTF-16 units, including both surrogates of supplementary characters.
func init() {
    Foreign("Data.String.CodeUnits")["length"] = func(value Any) Any {
        units := 0
        for _, char := range value.(string) {
            units++
            if char > 0xffff { units++ }
        }
        return units
    }
}
