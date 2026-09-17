package purescript
import("reflect";"testing";r "gopurs/output/gopurs_runtime")
func TestCountedFunctions(t *testing.T){
for n:=int64(0);n<=12;n++{fn:=Call_Test_Church_fromInt(n);for repeat:=0;repeat<3;repeat++{trace:=[]int64{};f:=r.Func(func(x r.Value)r.Value{trace=append(trace,x.IntVal);return r.Int(x.IntVal*3-7)});want:=int64(11);expected:=[]int64{};for i:=int64(0);i<n;i++{expected=append(expected,want);want=want*3-7};var got r.Value;if repeat==1{got=r.Apply(r.Apply(fn,f),r.Int(11))}else{got=r.Apply2(fn,f,r.Int(11))};if got.IntVal!=want||!reflect.DeepEqual(trace,expected){t.Fatalf("n=%d repeat=%d got=%d want=%d trace=%v expected=%v",n,repeat,got.IntVal,want,trace,expected)}}}
}
func TestChurchCompositions(t *testing.T){f:=r.Func(func(x r.Value)r.Value{return r.Int(x.IntVal+1)});for n:=int64(0);n<=10;n++{got:=r.Apply2(Call_Test_Church_c100k(n),f,r.Int(0)).IntVal;want:=n*n*n*n*n;if got!=want{t.Fatalf("n=%d got=%d want=%d",n,got,want)}}}
func TestCountedPanicOrder(t *testing.T){calls:=0;f:=r.Func(func(x r.Value)r.Value{calls++;if calls==3{panic("callback")};return r.Int(x.IntVal+1)});defer func(){if recover()!="callback"||calls!=3{t.Fatalf("calls=%d",calls)}}();r.Apply2(Call_Test_Church_fromInt(5),f,r.Int(0));t.Fatal("missing panic")}
