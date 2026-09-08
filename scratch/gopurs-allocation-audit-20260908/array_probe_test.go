package main

import (
 "testing"
 ps "gopurs/output/purescript"
)

var intSink int64
var benchN int64 = 900

func TestArrayCounterfactual(t *testing.T) {
 for n:=int64(-32);n<=2048;n++ {
  want:=ps.Call_Test_ArrayOps_sumEvens(n)
  if got:=copiedOriginalSumEvens(n);got!=want {t.Fatalf("copy n=%d got=%d want=%d",n,got,want)}
  if got:=withoutArrayRoundtrip(n);got!=want {t.Fatalf("no roundtrip n=%d got=%d want=%d",n,got,want)}
 }
}

func BenchmarkArray(b *testing.B) {
 for _,tc:=range []struct{name string; fn func(int64)int64}{
  {"compiled_original",ps.Call_Test_ArrayOps_sumEvens},
  {"copied_original",copiedOriginalSumEvens},
  {"without_roundtrip",withoutArrayRoundtrip},
 } {
  b.Run(tc.name,func(b *testing.B){b.ReportAllocs();for i:=0;i<b.N;i++{intSink=tc.fn(benchN)}})
 }
}
