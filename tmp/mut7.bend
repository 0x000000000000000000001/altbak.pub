import Base

def f(x: Nat) -> Nat:
  m = (x + 1n : Nat)
  match m:
    case 0n: 0n
    case 1n+p: p

def main() -> Nat:
  f(10n)
