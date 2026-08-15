const fs = require('fs');
let code = fs.readFileSync('output/purescript/Test_RBTree.go', 'utf8');

code = code.replace(
  /\/\/ TAST \(Let\): __local_var_2_0 -> gopurs_runtime\.Value\nvar __local_var_2_0 gopurs_runtime\.Value = gopurs_runtime\.Value\{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe\.Pointer\(Call_Test_RBTree_ins\(x_0, s_1\)\)\}\nvar __t1 \*Constructor_Test_RBTree_T\n\{\nif \(__local_var_2_0\.Type == 9 && __local_var_2_0\.IntVal == 3983586014 && __local_var_2_0\.UnsafePtr != nil\) \{\n__t1 = \(&Constructor_Test_RBTree_T\{1, 1583507464, \(\*Constructor_Test_RBTree_T\)\(__local_var_2_0\.UnsafePtr\)\.V1, \(\*Constructor_Test_RBTree_T\)\(__local_var_2_0\.UnsafePtr\)\.V2, \(\*Constructor_Test_RBTree_T\)\(__local_var_2_0\.UnsafePtr\)\.V3\}\)\ngoto end_branch_1\n\} else \{\n\n\}\n\}\n\{\nif \(__local_var_2_0\.Type == 9 && __local_var_2_0\.IntVal == 3983586014 && __local_var_2_0\.UnsafePtr == nil\) \{\n__t1 = \(\*Constructor_Test_RBTree_T\)\(nil\)\ngoto end_branch_1\n\} else \{\n\n\}\n\}/,
  `__local_var_2_0 := Call_Test_RBTree_ins(x_0, s_1)
var __t1 *Constructor_Test_RBTree_T
{
if (__local_var_2_0 != nil) {
__t1 = (&Constructor_Test_RBTree_T{1, 1583507464, __local_var_2_0.V1, __local_var_2_0.V2, __local_var_2_0.V3})
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0 == nil) {
__t1 = (*Constructor_Test_RBTree_T)(nil)
goto end_branch_1
} else {

}
}`
);

fs.writeFileSync('output/purescript/Test_RBTree.go', code);
