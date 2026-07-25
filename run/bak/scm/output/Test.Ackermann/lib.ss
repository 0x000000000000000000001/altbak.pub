#!r6rs
#!chezscheme
(library
  (Test.Ackermann lib)
  (export
    ackermann
    act
    describe)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Ackermann (3, 4):")))

  (scm:define ackermann
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx=? v0 0) (scm:fx+ v11 1)]
          [(scm:fx=? v11 0) ((ackermann (scm:fx- v0 1)) 1)]
          [scm:else ((ackermann (scm:fx- v0 1)) ((ackermann v0) (scm:fx- v11 1)))]))))

  (scm:define act
    (scm:let ([_0 (Bench.opaque 3)])
      (scm:lambda ()
        (scm:let ([dummy1 (_0)])
          ((Effect.Console.log (Data.Show.showIntImpl ((ackermann dummy1) 4)))))))))
