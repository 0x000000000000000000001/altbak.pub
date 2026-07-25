#!/usr/bin/env bash

# Setup the environment like bin/go/run does
mkdir -p run/bak/go/output run/bak/go/spago
rm -rf output .spago spago.yaml spago.lock
ln -s run/bak/go/output output
ln -s run/bak/go/spago .spago
ln -s run/bak/go/spago.go.yaml spago.yaml
touch run/bak/go/spago.lock && ln -s run/bak/go/spago.lock spago.lock

# Build PureScript
spago build || echo "spago build failed, but continuing..."

set -e

node ../gopurs/bin/gopurs.js --main AppX

cd output
rm -f go.mod go.sum
go mod init gopurs/output > /dev/null 2>&1

# Get btree
go get github.com/google/btree

# Inject FFI for GoogleBTree
mkdir -p Test.GoogleBTree
cat << 'FFI' > Test.GoogleBTree/ffi.go
package Test_GoogleBTree

import (
    "gopurs/output/gopurs_runtime"
    "github.com/google/btree"
    "unsafe"
)

type IntItem int64
func (a IntItem) Less(b btree.Item) bool {
    return a < b.(IntItem)
}

func init() {
    Get_empty = func() gopurs_runtime.Value {
        tree := btree.New(32) // Degree 32 B-Tree
        return gopurs_runtime.Value{Type: 9, IntVal: 1, UnsafePtr: unsafe.Pointer(tree)}
    }
    
    Get_insert = func() gopurs_runtime.Value {
        return gopurs_runtime.Func2(func(v_key gopurs_runtime.Value, v_tree gopurs_runtime.Value) gopurs_runtime.Value {
            tree := (*btree.BTree)(v_tree.UnsafePtr)
            // Clone makes a copy-on-write clone (O(1) allocation)
            newTree := tree.Clone()
            newTree.ReplaceOrInsert(IntItem(v_key.IntVal))
            return gopurs_runtime.Value{Type: 9, IntVal: 1, UnsafePtr: unsafe.Pointer(newTree)}
        })
    }
    
    Get_size = func() gopurs_runtime.Value {
        return gopurs_runtime.Func(func(v_tree gopurs_runtime.Value) gopurs_runtime.Value {
            tree := (*btree.BTree)(v_tree.UnsafePtr)
            return gopurs_runtime.Int(int64(tree.Len()))
        })
    }
}
FFI

go mod tidy > /dev/null 2>&1
go build -o go_app .

echo "=== Running Google BTree Benchmark ==="
GOGC=1000 ./go_app

