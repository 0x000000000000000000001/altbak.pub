package main

import (
 "encoding/json"
 "flag"
 "fmt"
 "os"
 "runtime"
 "runtime/debug"
 "runtime/pprof"
 "time"
 rt "gopurs/output/gopurs_runtime"
 ps "gopurs/output/purescript"
)

var sink string
type testCase struct { Name string; Get func() rt.Value }
var cases = []testCase{
 {"AstTree", ps.Get_Test_AstTree_act}, {"Fib", ps.Get_Test_Fib_act},
 {"ListOps", ps.Get_Test_ListOps_act}, {"TCO", ps.Get_Test_TCO_act},
 {"Records", ps.Get_Test_Records_act}, {"Ackermann", ps.Get_Test_Ackermann_act},
 {"Church", ps.Get_Test_Church_act}, {"Primes", ps.Get_Test_Primes_act},
 {"RBTree", ps.Get_Test_RBTree_act}, {"Polymorphism", ps.Get_Test_Polymorphism_act},
 {"StateMonad", ps.Get_Test_StateMonad_act}, {"LazyEvaluation", ps.Get_Test_LazyEvaluation_act},
 {"ArrayOps", ps.Get_Test_ArrayOps_act}, {"RowToList", ps.Get_Test_RowToList_act},
}
type sample struct {
 Test string `json:"test"`
 Result string `json:"result"`
 Repetitions int `json:"repetitions"`
 Sample int `json:"sample"`
 NsPerOp float64 `json:"ns_per_op"`
 BytesPerOp float64 `json:"bytes_per_op"`
 AllocsPerOp float64 `json:"allocs_per_op"`
 GCs uint32 `json:"gcs"`
}
func main() {
 name := flag.String("test", "", "one test or all")
 profile := flag.String("profile", "", "write exact allocation profile for one test")
 count := flag.Int("count", 1, "profile executions")
 flag.Parse()
 if *profile != "" { runtime.MemProfileRate = 1 }
 debug.SetGCPercent(800)
 enc := json.NewEncoder(os.Stdout)
 for _, tc := range cases {
  if *name != "" && *name != tc.Name { continue }
  act := tc.Get()
  run := func() string { return rt.Apply(act, rt.Value{}).StrVal() }
  expected := run()
  run(); run()
  if *profile != "" {
   runtime.GC(); runtime.GC()
   before,err:=os.Create(*profile+".before"); if err!=nil {panic(err)}
   if err=pprof.Lookup("allocs").WriteTo(before,0); err!=nil {panic(err)}
   if err=before.Close(); err!=nil {panic(err)}
   for i:=0; i<*count; i++ { sink=run(); if sink!=expected { panic("unstable result") } }
   runtime.GC(); runtime.GC()
   f,err:=os.Create(*profile); if err!=nil {panic(err)}
   if err=pprof.Lookup("allocs").WriteTo(f,0); err!=nil {panic(err)}
   if err=f.Close(); err!=nil {panic(err)}
   fmt.Fprintf(os.Stderr,"profile %s: %d executions, result %s\n",tc.Name,*count,expected)
   return
  }
  start:=time.Now(); sink=run(); elapsed:=time.Since(start)
  reps:=int((80*time.Millisecond)/elapsed); if reps<1 {reps=1}; if reps>20000 {reps=20000}
  for s:=0;s<3;s++ {
   runtime.GC()
   var before,after runtime.MemStats
   runtime.ReadMemStats(&before)
   begin:=time.Now()
   for i:=0;i<reps;i++ {sink=run(); if sink!=expected {panic("unstable result")}}
   duration:=time.Since(begin)
   runtime.ReadMemStats(&after)
   err:=enc.Encode(sample{tc.Name,expected,reps,s+1,float64(duration.Nanoseconds())/float64(reps),float64(after.TotalAlloc-before.TotalAlloc)/float64(reps),float64(after.Mallocs-before.Mallocs)/float64(reps),after.NumGC-before.NumGC})
   if err!=nil {panic(err)}
  }
 }
}
