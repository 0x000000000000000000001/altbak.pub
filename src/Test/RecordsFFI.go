package Test_RecordsFFI


type dictD struct {
	e int
	f int
}

type dictB struct {
	c int
	d dictD
}

type dictR struct {
	a int
	b dictB
}

func updateRec(n int, r dictR) dictR {
	if n == 0 {
		return r
	}
	
	newD := dictD{
		e: r.b.d.e + 3,
		f: r.b.d.f + (n % 5),
	}
	newB := dictB{
		c: r.b.c + 2,
		d: newD,
	}
	newR := dictR{
		a: r.a + 1,
		b: newB,
	}
	
	return updateRec(n - 1, newR)
}

func RunRecordsFFI(limit int) int {
	dummy := limit
	initial := dictR{a: 0, b: dictB{c: 0, d: dictD{e: 0, f: 0}}}
	res := updateRec(dummy, initial)
	return (res.b.d.f)
}
