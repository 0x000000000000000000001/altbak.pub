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
    
    if mode == "verify" {
        root := kernels.Call_Test_RBTree_buildTree(n,nil)
        v := kernels.Validate(root,n)
        if v.Depth != kernels.Call_Test_RBTree_depth(root) { panic("depth mismatch") }
        fmt.Println(v.Count,v.Sum,v.Depth,v.BlackHeight,-1)
        return
    }
    if mode != "bench" { panic("unknown mode") }
    var sum int64
    for i := int64(0); i < count; i++ {
        
        root := kernels.Call_Test_RBTree_buildTree(n,nil)
        sum += kernels.Call_Test_RBTree_depth(root)
    }
    fmt.Println(sum)
}
