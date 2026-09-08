package main
import (
    "encoding/json"
    "fmt"
    "os"
    "runtime"
    "strconv"
    "rbmemory/kernels"
)
type TreeSummary struct {
    Count int64 `json:"count"`
    Sum int64 `json:"sum"`
    Depth int64 `json:"depth"`
    BlackHeight int64 `json:"black_height"`
    Allocations int64 `json:"allocations"`
    PoolBytes int64 `json:"pool_bytes"`
    FallbackAllocations int64 `json:"fallback_allocations"`
}
func summary(root *kernels.Constructor_Test_RBTree_T, n int64) TreeSummary {
    v := kernels.Validate(root,n)
    if v.Depth != kernels.Call_Test_RBTree_depth(root) { panic("depth mismatch") }
    allocations, poolBytes, fallback := kernels.AllocationCounts()
    return TreeSummary{v.Count,v.Sum,v.Depth,v.BlackHeight,allocations,poolBytes,fallback}
}
func verifyOne(n int64) TreeSummary {
    kernels.ResetPool()
    root := kernels.Call_Test_RBTree_buildTree(n,nil)
    runtime.GC() // Untimed: all pooled and fallback pointers must remain valid.
    result := summary(root,n)
    runtime.KeepAlive(root)
    return result
}
func runBatch(count, n int64) (int64, int64, int64, int64) {
    var checksum int64
    for i:=int64(0); i<count; i++ {
        kernels.ResetPool()
        root := kernels.Call_Test_RBTree_buildTree(n,nil)
        checksum += kernels.Call_Test_RBTree_depth(root)
    }
    allocations,poolBytes,fallback := kernels.AllocationCounts()
    kernels.ResetPool() // Include cleanup of the last completed tree too.
    return checksum,allocations,poolBytes,fallback
}
type Memory struct {
    HeapAlloc uint64 `json:"heap_alloc_bytes"`
    HeapInuse uint64 `json:"heap_inuse_bytes"`
    HeapSys uint64 `json:"heap_sys_bytes"`
    TotalAlloc uint64 `json:"total_alloc_bytes"`
    Mallocs uint64 `json:"mallocs"`
    NumGC uint32 `json:"num_gc"`
    NextGC uint64 `json:"next_gc_bytes"`
    GCCPUFraction float64 `json:"gc_cpu_fraction"`
}
func memory() Memory {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return Memory{m.HeapAlloc,m.HeapInuse,m.HeapSys,m.TotalAlloc,m.Mallocs,m.NumGC,m.NextGC,m.GCCPUFraction}
}
func emit(value any) { if err:=json.NewEncoder(os.Stdout).Encode(value); err!=nil {panic(err)} }
func main() {
    if len(os.Args)!=5 {panic("usage: runner bench|verify|stats|sequence count n capMiB")}
    mode:=os.Args[1]
    count,err:=strconv.ParseInt(os.Args[2],10,64); if err!=nil {panic(err)}
    n,err:=strconv.ParseInt(os.Args[3],10,64); if err!=nil {panic(err)}
    capMiB,err:=strconv.Atoi(os.Args[4]); if err!=nil {panic(err)}
    if count<1 || count>1000 || n<0 || n>1000000 || capMiB<0 || capMiB>128 {panic("POC input out of range")}
    kernels.ConfigurePool(capMiB)
    switch mode {
    case "verify":
        emit(verifyOne(n))
    case "bench":
        checksum,_,_,_:=runBatch(count,n)
        fmt.Println(checksum)
    case "stats":
        before:=memory()
        checksum,allocations,poolBytes,fallback:=runBatch(count,n)
        after:=memory()
        runtime.GC()
        afterGC:=memory()
        emit(map[string]any{"checksum":checksum,"allocations":allocations,"pool_bytes":poolBytes,
            "fallback_allocations":fallback,"before":before,"after":after,"after_gc":afterGC,
            "collections_during_batch":after.NumGC-before.NumGC,
            "allocated_bytes_during_batch":after.TotalAlloc-before.TotalAlloc})
    case "sequence":
        values:=[]TreeSummary{}
        for _,size:=range []int64{n,1000,0,n} {values=append(values,verifyOne(size))}
        kernels.ResetPool()
        runtime.GC()
        emit(map[string]any{"trees":values,"after_gc":memory()})
    default: panic("unknown mode")
    }
}
