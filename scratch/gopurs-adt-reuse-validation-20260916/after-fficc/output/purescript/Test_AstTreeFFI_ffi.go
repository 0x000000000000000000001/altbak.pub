package purescript

import "gopurs/output/gopurs_runtime"


type exprType int

const (
	exprVal exprType = iota
	exprAdd
	exprMul
	exprSub
)

type expr struct {
	typ   exprType
	value int
	left  *expr
	right *expr
}

func evalAst(e *expr) int {
	switch e.typ {
	case exprVal:
		return e.value
	case exprAdd:
		return evalAst(e.left) + evalAst(e.right)
	case exprMul:
		return evalAst(e.left) * evalAst(e.right)
	case exprSub:
		return evalAst(e.left) - evalAst(e.right)
	}
	return 0
}

func buildTreeAst(n int) *expr {
	if n == 0 {
		return &expr{typ: exprVal, value: 1}
	}
	
	valN := &expr{typ: exprVal, value: n}
	val1 := &expr{typ: exprVal, value: 1}
	leftTree := buildTreeAst(n - 1)
	rightTree := buildTreeAst(n - 1)
	
	mulNode := &expr{typ: exprMul, left: valN, right: leftTree}
	subNode := &expr{typ: exprSub, left: rightTree, right: val1}
	
	return &expr{typ: exprAdd, left: mulNode, right: subNode}
}

// gopurs unboxes Int to float64 for FFI methods
func Test_AstTreeFFI_RunAstTreeFFI(limit int) int {
	t := buildTreeAst(int(limit))
	return (evalAst(t))
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_AstTreeFFI_RunAstTreeFFI = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_AstTreeFFI_RunAstTreeFFI(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})