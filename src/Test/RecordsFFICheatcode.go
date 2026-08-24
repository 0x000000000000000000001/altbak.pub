package Test_RecordsFFICheatcode

type dictD_cc struct { e int; f int }
type dictB_cc struct { c int; d dictD_cc }
type dictR_cc struct { a int; b dictB_cc }

func RunRecordsFFICheatcode(limit int) int {
	n := int(limit)
	r := &dictR_cc{a: 0, b: dictB_cc{c: 0, d: dictD_cc{e: 0, f: 0}}}
	for n > 0 {
		r.a += 1
		r.b.c += 2
		r.b.d.e += 3
		r.b.d.f += (n % 5)
		n--
	}
	return r.b.d.f
}
