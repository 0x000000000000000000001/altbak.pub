import Base

# 1. AstTree
type Expr is Data:
  Val { v: Nat }
  Add { l: Expr, r: Expr }
  Mul { l: Expr, r: Expr }
  Sub { l: Expr, r: Expr }

def evalAst(e: Expr) -> Nat:
  match e:
    case Val{+v}:
      v
    case Add{+l, +r}:
      (evalAst(l) + evalAst(r) : Nat)
    case Mul{+l, +r}:
      (evalAst(l) * evalAst(r) : Nat)
    case Sub{+l, +r}:
      (evalAst(l) - evalAst(r) : Nat)

def buildTreeAst(+n: Nat) -> Expr:
  match n:
    case 0n:
      Val{1n}
    case 1n+p:
      Add{Mul{Val{n}, buildTreeAst(p)}, Sub{buildTreeAst(p), Val{1n}}}

def runAstTree(+limit: Nat) -> Nat:
  evalAst(buildTreeAst(limit))

# 2. Fib
def fib(+n: Nat) -> Nat:
  match n:
    case 0n: 0n
    case 1n+p:
      match p:
        case 0n: 1n
        case 1n+q: (fib(1n+q) + fib(q) : Nat)

def runFib(+limit: Nat) -> Nat:
  fib(limit)

def main() -> IO(Unit):
  do IO<Unit>:
    IO.print("--- Bend Benchmark ---")
    IO.print(Nat.show(runAstTree(3n)))
    IO.print(Nat.show(runFib(10n)))
    IO.print("Done")
