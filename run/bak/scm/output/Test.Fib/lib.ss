#!r6rs
#!chezscheme
(library
  (Test.Fib lib)
  (export
    act
    describe
    fib)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define fib
    (scm:lambda (v0)
      (scm:cond
        [(scm:fx=? v0 0) 0]
        [(scm:fx=? v0 1) 1]
        [scm:else (scm:fx+ (fib (scm:fx- v0 1)) (fib (scm:fx- v0 2)))])))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Fibonacci:")))

  (scm:define act
    (scm:let ([_0 (Bench.opaque 10)])
      (scm:lambda ()
        (scm:let ([dummy1 (_0)])
          ((Effect.Console.log (Data.Show.showIntImpl (fib dummy1)))))))))
