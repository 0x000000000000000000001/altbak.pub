#!r6rs
#!chezscheme
(library
  (Test.LazyEvaluation lib)
  (export
    Lazy
    act
    buildThunks
    defer
    describe
    force
    runManyTimes)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define Lazy
    (scm:lambda (x0)
      x0))

  (scm:define force
    (scm:lambda (v0)
      (v0 Data.Unit.unit)))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Lazy Evaluation (16M Thunks Forced, 4k Depth):")))

  (scm:define defer
    (scm:lambda (f0)
      f0))

  (scm:define buildThunks
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx=? v0 0) v11]
          [scm:else ((buildThunks (scm:fx- v0 1)) (scm:lambda (v22)
            (scm:fx+ (v11 Data.Unit.unit) 1)))]))))

  (scm:define runManyTimes
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx=? v0 0) v11]
          [scm:else ((runManyTimes (scm:fx- v0 1)) (scm:fx+ v11 (((buildThunks 4000) (scm:lambda (v22)
            0)) Data.Unit.unit)))]))))

  (scm:define act
    (Effect.Console.log (Data.Show.showIntImpl ((runManyTimes 4000) 0)))))
