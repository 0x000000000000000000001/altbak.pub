package Data_Int_Bits

func And(n1 int, n2 int) int { return n1 & n2 }
func Or(n1 int, n2 int) int { return n1 | n2 }
func Xor(n1 int, n2 int) int { return n1 ^ n2 }
func Shl(n1 int, n2 int) int { return n1 << n2 }
func Shr(n1 int, n2 int) int { return n1 >> n2 }
func Zshr(n1 int, n2 int) int { return int(uint(n1) >> uint(n2)) }
func Complement(n int) int { return ^n }
