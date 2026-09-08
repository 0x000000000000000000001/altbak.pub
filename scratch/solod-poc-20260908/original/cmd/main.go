package main
import (
    "fmt"
    "os"
    "strconv"
    "poc/kernels"
)
func main() {
    if len(os.Args) != 4 { panic("usage: runner kind iterations n") }
    count, err := strconv.ParseInt(os.Args[2],10,64); if err != nil { panic(err) }
    n, err := strconv.ParseInt(os.Args[3],10,64); if err != nil { panic(err) }
    var sum int64
    switch os.Args[1] {
    case "fib":
        for i := int64(0); i < count; i++ { sum += kernels.Call_Test_Fib_fib(n+(i&1)) }
    case "ackermann":
        for i := int64(0); i < count; i++ { sum += kernels.Call_Test_Ackermann_ackermann(3,n+(i&1)) }
    case "tco":
        for i := int64(0); i < count; i++ { sum += kernels.Call_Test_TCO_deepTailRec(n+(i&1),0) }
    default: panic("unknown kind")
    }
    fmt.Println(sum)
}
