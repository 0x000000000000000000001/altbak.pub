package Test_AstTreeFFICheatcode

type exprType_cheatcode int

const (
	exprVal_cheatcode exprType_cheatcode = iota
	exprAdd_cheatcode
	exprMul_cheatcode
	exprSub_cheatcode
)

type expr_cheatcode struct {
	typ   exprType_cheatcode
	value int
	left  *expr_cheatcode
	right *expr_cheatcode
}

func evalAst_cheatcode(e *expr_cheatcode) int {
	switch e.typ {
	case exprVal_cheatcode:
		return e.value
	case exprAdd_cheatcode:
		return evalAst_cheatcode(e.left) + evalAst_cheatcode(e.right)
	case exprMul_cheatcode:
		return evalAst_cheatcode(e.left) * evalAst_cheatcode(e.right)
	case exprSub_cheatcode:
		return evalAst_cheatcode(e.left) - evalAst_cheatcode(e.right)
	}
	return 0
}

func buildTreeAst_cheatcode(n int) *expr_cheatcode {
	if n == 0 {
		return &expr_cheatcode{typ: exprVal_cheatcode, value: 1}
	}
	
	valN := &expr_cheatcode{typ: exprVal_cheatcode, value: n}
	val1 := &expr_cheatcode{typ: exprVal_cheatcode, value: 1}
	leftTree := buildTreeAst_cheatcode(n - 1)
	rightTree := buildTreeAst_cheatcode(n - 1)
	
	mulNode := &expr_cheatcode{typ: exprMul_cheatcode, left: valN, right: leftTree}
	subNode := &expr_cheatcode{typ: exprSub_cheatcode, left: rightTree, right: val1}
	
	return &expr_cheatcode{typ: exprAdd_cheatcode, left: mulNode, right: subNode}
}

// gopurs unboxes Int to float64 for FFICheatcode methods
func RunAstTreeFFICheatcode(limit float64) float64 {
	t := buildTreeAst_cheatcode(int(limit))
	return float64(evalAst_cheatcode(t))
}
