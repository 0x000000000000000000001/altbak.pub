package main
import("encoding/json";"fmt";"os";"runtime";"time";r "gopurs/output/gopurs_runtime";p "gopurs/output/purescript")
var sink int64
var next = r.Func(func(x r.Value)r.Value{return r.Int(x.IntVal+1)})
func run(){sink=r.Apply2(p.Call_Test_Church_c100k(10),next,r.Int(0)).IntVal;if sink!=100000{panic(sink)}}
func main(){for i:=0;i<10;i++{run()}; samples:=[]map[string]any{};for sample:=0;sample<7;sample++{runtime.GC();var a,b runtime.MemStats;runtime.ReadMemStats(&a);start:=time.Now();for i:=0;i<100;i++{run()};ns:=time.Since(start).Nanoseconds();runtime.ReadMemStats(&b);samples=append(samples,map[string]any{"ns_per_call":float64(ns)/100,"allocs_per_call":float64(b.Mallocs-a.Mallocs)/100,"bytes_per_call":float64(b.TotalAlloc-a.TotalAlloc)/100,"result":sink})};if err:=json.NewEncoder(os.Stdout).Encode(samples);err!=nil{panic(err)};_ = fmt.Sprint(sink)}
