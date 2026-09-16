package Test_ChurchFFI


type Church func(func(int) int) func(int) int

func zeroC() Church {
	return func(f func(int) int) func(int) int {
		return func(x int) int {
			return x
		}
	}
}

func succC(n Church) Church {
	return func(f func(int) int) func(int) int {
		return func(x int) int {
			return f(n(f)(x))
		}
	}
}

func addC(m Church, n Church) Church {
	return func(f func(int) int) func(int) int {
		return func(x int) int {
			return m(f)(n(f)(x))
		}
	}
}

func mulC(m Church, n Church) Church {
	return func(f func(int) int) func(int) int {
		return func(x int) int {
			return m(n(f))(x)
		}
	}
}

func fromInt(n int) Church {
	if n == 0 {
		return zeroC()
	}
	return succC(fromInt(n - 1))
}

func toInt(n Church) int {
	return n(func(x int) int { return x + 1 })(0)
}

func c10(n int) Church {
	return fromInt(n)
}

func c100(n int) Church {
	return mulC(c10(n), c10(n))
}

func c10k(n int) Church {
	return mulC(c100(n), c100(n))
}

func c100k(n int) Church {
	return mulC(c10k(n), c10(n))
}

func RunChurchFFI(limit int) int {
	dummy := limit
	return (toInt(c100k(dummy)))
}
