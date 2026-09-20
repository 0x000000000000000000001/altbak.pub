package Test_ChurchFFI

type Church[A any] func(func(A) A) func(A) A

func zeroC[A any]() Church[A] {
	return func(f func(A) A) func(A) A {
		return func(x A) A {
			return x
		}
	}
}

func succC[A any](n Church[A]) Church[A] {
	return func(f func(A) A) func(A) A {
		return func(x A) A {
			return f(n(f)(x))
		}
	}
}

func addC[A any](m Church[A], n Church[A]) Church[A] {
	return func(f func(A) A) func(A) A {
		return func(x A) A {
			return m(f)(n(f)(x))
		}
	}
}

func mulC[A any](m Church[A], n Church[A]) Church[A] {
	return func(f func(A) A) func(A) A {
		return func(x A) A {
			return m(n(f))(x)
		}
	}
}

func fromInt(n int) Church[int] {
	if n == 0 {
		return zeroC[int]()
	}
	return succC(fromInt(n - 1))
}

func toInt(n Church[int]) int {
	return n(func(x int) int { return x + 1 })(0)
}

func c10(n int) Church[int] {
	return fromInt(n)
}

func c100(n int) Church[int] {
	return mulC(c10(n), c10(n))
}

func c10k(n int) Church[int] {
	return mulC(c100(n), c100(n))
}

func c100k(n int) Church[int] {
	return mulC(c10k(n), c10(n))
}

func RunChurchFFI(limit int) int {
	dummy := limit
	return (toInt(c100k(dummy)))
}
