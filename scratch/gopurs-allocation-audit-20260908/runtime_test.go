package main

import (
 "runtime"
 "sync"
 "testing"
 rt "gopurs/output/gopurs_runtime"
)

func TestConcurrentFunctionWrapping(t *testing.T) {
 var wg sync.WaitGroup
 for g:=0;g<2;g++ {
  wg.Add(1)
  go func(offset int) {
   defer wg.Done()
   for j:=0;j<1000;j++ {
    n:=int64(offset+j)
    f:=rt.Func(func(x rt.Value) rt.Value { return rt.Int(x.IntVal+n) })
    if rt.Apply(f,rt.Int(3)).IntVal!=n+3 { t.Error("wrong closure result") }
   }
  }(g*1000)
 }
 wg.Wait()
}

var retainedFunc rt.Value
var valueSink rt.Value

func TestSaturatedCallAllocations(t *testing.T) {
 f5:=rt.Func5(func(a,b,c,d,e rt.Value)rt.Value{return rt.Int(a.IntVal+b.IntVal+c.IntVal+d.IntVal+e.IntVal)})
 f6:=rt.Func6(func(a,b,c,d,e,f rt.Value)rt.Value{return rt.Int(a.IntVal+b.IntVal+c.IntVal+d.IntVal+e.IntVal+f.IntVal)})
 one:=rt.Int(1)
 a5:=testing.AllocsPerRun(1000,func(){valueSink=rt.Apply5(f5,one,one,one,one,one)})
 a6:=testing.AllocsPerRun(1000,func(){valueSink=rt.Apply6(f6,one,one,one,one,one,one)})
 t.Logf("Apply5 allocations/call=%g; Apply6 allocations/call=%g",a5,a6)
 if valueSink.IntVal!=6 {t.Error("wrong arithmetic result")}
}

func TestValueToAnyFloatCompatibility(t *testing.T) {
 // This audit probe records the current helper's failure without changing it.
 defer func(){if p:=recover();p!=nil {t.Logf("ValueToAny(Float(1.5)) panics: %v",p)} else {t.Error("expected static mismatch was not reproduced")}}()
 t.Logf("converted=%v",rt.ValueToAny(rt.Float(1.5)))
}

func TestLastClosureRetention(t *testing.T) {
 // Reach a steady baseline with a scalar closure in EscapeSink.
 rt.Func(func(x rt.Value)rt.Value{return x})
 runtime.GC()
 var baseline,retained,released runtime.MemStats
 runtime.ReadMemStats(&baseline)
 func(){
  data:=make([]byte,16<<20)
  data[0]=7
  retainedFunc=rt.Func(func(x rt.Value)rt.Value{return rt.Int(int64(data[0]))})
  valueSink=rt.Apply(retainedFunc,rt.Value{})
  retainedFunc=rt.Value{}
 }()
 runtime.GC();runtime.ReadMemStats(&retained)
 rt.Func(func(x rt.Value)rt.Value{return x})
 runtime.GC();runtime.ReadMemStats(&released)
 t.Logf("heap baseline=%d retained=%d released=%d; retained_delta=%d released_delta=%d",baseline.HeapAlloc,retained.HeapAlloc,released.HeapAlloc,int64(retained.HeapAlloc)-int64(baseline.HeapAlloc),int64(released.HeapAlloc)-int64(baseline.HeapAlloc))
 if int64(retained.HeapAlloc)-int64(released.HeapAlloc)<15<<20 { t.Error("expected last closure environment retention was not observed") }
}
