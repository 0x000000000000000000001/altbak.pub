package kernels

func Call_Test_Fib_fib(v_0_loop int64) int64 {
fib:
for {
if false { continue fib }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0) == (int64(0)) {
__t0 = int64(0)
goto end_branch_0
} else {

}
}
{
if (v_0) == (int64(1)) {
__t0 = int64(1)
goto end_branch_0
} else {

}
}
{
__t0 = (Call_Test_Fib_fib((v_0) - (int64(1)))) + (Call_Test_Fib_fib((v_0) - (int64(2))))
}
end_branch_0:
return __t0
}
}

func Call_Test_Ackermann_ackermann(v_0_loop int64, v1_1_loop int64) int64 {
ackermann:
for {
if false { continue ackermann }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t0 int64
{
if (v_0) == (int64(0)) {
__t0 = (v1_1) + (int64(1))
goto end_branch_0
} else {

}
}
{
if (v1_1) == (int64(0)) {
v_0_loop = (v_0) - (int64(1))
v1_1_loop = int64(1)
continue ackermann
__t0 = func() int64 { panic("unreachable") }()
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (int64(1))
v1_1_loop = Call_Test_Ackermann_ackermann(v_0, (v1_1) - (int64(1)))
continue ackermann
__t0 = func() int64 { panic("unreachable") }()
}
end_branch_0:
return __t0
}
}

func Call_Test_TCO_deepTailRec(v_0_loop int64, v1_1_loop int64) int64 {
deepTailRec:
for {
if false { continue deepTailRec }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t0 int64
{
if (v_0) == (int64(0)) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (int64(1))
v1_1_loop = (v1_1) + ((v_0) % (int64(3)))
continue deepTailRec
__t0 = func() int64 { panic("unreachable") }()
}
end_branch_0:
return __t0
}
}
