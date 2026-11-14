package ffi_loader

import (
    "testing"
    . "github.com/purescript-native/go-runtime"
)

func TestCodeUnitsLength(t *testing.T) {
    length := Foreign("Data.String.CodeUnits")["length"].(func(Any) Any)
    for _, item := range []struct{ text string; units int }{
        {"", 0}, {"abc", 3}, {"é", 1}, {"😀", 2}, {"a😀é", 4}, {"\uFFFD", 1},
    } {
        if got := length(item.text).(int); got != item.units {
            t.Fatalf("length(%q) = %d, expected %d UTF-16 units", item.text, got, item.units)
        }
    }
}
