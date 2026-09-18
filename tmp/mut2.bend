import Base
def f(x: Nat) -> Nat:
  match x:
    case 0n: 0n
    case 1n+p: g(f, p)

def g(self: _, x: Nat) -> Nat:
  match x:
    case 0n: 0n
    case 1n+p: self(p)

def main() -> Nat:
  f(10n)
