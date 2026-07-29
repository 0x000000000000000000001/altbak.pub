package main

type Value struct {
	x int
}

func Apply(a, b Value) Value { return Value{} }

func Call_comparing(a Value, b func(Value) Value, c, d Value) Value { return Value{} }

func Get_comparing() Value {
	var dictOrd_0_box Value
	var f_1_box Value
	var x_2_box Value
	var y_3_box Value
	
	// This simulates line 292
	return Call_comparing(dictOrd_0_box, func(a0 Value) Value {
		return Apply(f_1_box, a0)
	}, x_2_box, y_3_box)
}

func main() {}
