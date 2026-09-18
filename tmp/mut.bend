import Base
def f(x: Nat) -> Nat:
  match x:
    case 0n: 0n
    case 1n+p: g(p)

def g(x: Nat) -> Nat:
  match x:
    case 0n: 0n
    case 1n+p: f(p)

def main() -> Nat:
  f(10n)
