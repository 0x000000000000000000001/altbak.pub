package main
import (
    "fmt"
    "os"
    "strconv"
    "rbpoc/kernels"
)
func main() {
    if len(os.Args) != 4 { panic("usage: runner bench|verify count n") }
    mode := os.Args[1]
    count, err := strconv.ParseInt(os.Args[2],10,64); if err != nil { panic(err) }
    n, err := strconv.ParseInt(os.Args[3],10,64); if err != nil { panic(err) }
    if n < 0 || n > 100001 || count < 1 { panic("input out of POC range") }
    
    bits := int64(0)
    for k := n; k > 0; k >>= 1 { bits++ }
    capacity := (n+1)*(6*bits+4)
    pool := make([]kernels.Constructor_Test_RBTree_T, int(capacity))
    kernels.BeginArena(pool)

    if mode == "verify" {
        root := kernels.Call_Test_RBTree_buildTree(n,nil)
        v := kernels.Validate(root,n)
        if v.Depth != kernels.Call_Test_RBTree_depth(root) { panic("depth mismatch") }
        fmt.Println(v.Count,v.Sum,v.Depth,v.BlackHeight,kernels.NodesUsed())
        return
    }
    if mode != "bench" { panic("unknown mode") }
    var sum int64
    for i := int64(0); i < count; i++ {
        kernels.ResetArena()
        root := kernels.Call_Test_RBTree_buildTree(n,nil)
        sum += kernels.Call_Test_RBTree_depth(root)
    }
    fmt.Println(sum)
}
