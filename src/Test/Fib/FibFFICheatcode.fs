module Test.FibFFICheatcode
let rec fib n = if n < 2 then n else fib (n - 1) + fib (n - 2)
let runFibFFICheatcode (n: obj) = fib (unbox<int> n) :> obj
