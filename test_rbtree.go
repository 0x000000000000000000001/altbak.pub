package main
import "fmt"
import pkg_Test_RBTree "altbak/output/Test.RBTree"
import "altbak/output/gopurs_runtime"
func main() {
    fmt.Println("Running RBTree...")
    gopurs_runtime.Apply(pkg_Test_RBTree.Get_act(), gopurs_runtime.Value{})
    fmt.Println("Done")
}
