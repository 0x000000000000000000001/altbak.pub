#!r6rs
#!chezscheme
(library
  (Test.TCO lib)
  (export
    act
    deepTailRec
    describe)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.EuclideanRing lib) Data.EuclideanRing.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Tail Call Optimization (10M calls):")))

  (scm:define deepTailRec
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx=? v0 0) v11]
          [scm:else ((deepTailRec (scm:fx- v0 1)) (scm:fx+ v11 ((Data.EuclideanRing.intMod v0) 3)))]))))

  (scm:define act
    (Effect.Console.log (Data.Show.showIntImpl ((deepTailRec 10000000) 0)))))
