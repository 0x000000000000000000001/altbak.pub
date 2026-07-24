#!r6rs
#!chezscheme
(library
  (Test.StateMonad lib)
  (export
    State
    act
    bindState
    chainModifications
    describe
    get
    modify
    pureState
    put
    runManyTimes
    runState)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define State
    (scm:lambda (x0)
      x0))

  (scm:define runState
    (scm:lambda (v0)
      (scm:lambda (s1)
        (v0 s1))))

  (scm:define put
    (scm:lambda (s0)
      (scm:lambda (v1)
        (scm:list (scm:cons (scm:string->symbol "val") Data.Unit.unit) (scm:cons (scm:string->symbol "state") s0)))))

  (scm:define pureState
    (scm:lambda (a0)
      (scm:lambda (s1)
        (scm:list (scm:cons (scm:string->symbol "val") a0) (scm:cons (scm:string->symbol "state") s1)))))

  (scm:define get
    (scm:lambda (s0)
      (scm:list (scm:cons (scm:string->symbol "val") s0) (scm:cons (scm:string->symbol "state") s0))))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "State Monad (1.2k Binds, 60 Stack Depth):")))

  (scm:define bindState
    (scm:lambda (v0)
      (scm:lambda (g1)
        (scm:lambda (s2)
          (scm:let ([r13 (v0 s2)])
            ((g1 (rt:record-ref r13 (scm:string->symbol "val"))) (rt:record-ref r13 (scm:string->symbol "state"))))))))

  (scm:define modify
    (scm:lambda (f0)
      (scm:lambda (s1)
        (scm:list (scm:cons (scm:string->symbol "val") Data.Unit.unit) (scm:cons (scm:string->symbol "state") (f0 s1))))))

  (scm:define chainModifications
    (scm:lambda (v0)
      (scm:cond
        [(scm:fx=? v0 0) (scm:lambda (s1)
          (scm:list (scm:cons (scm:string->symbol "val") Data.Unit.unit) (scm:cons (scm:string->symbol "state") s1)))]
        [scm:else (scm:lambda (s1)
          ((chainModifications (scm:fx- v0 1)) (scm:fx+ s1 1)))])))

  (scm:define runManyTimes
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx=? v0 0) v11]
          [scm:else ((runManyTimes (scm:fx- v0 1)) (scm:fx+ v11 (rt:record-ref ((chainModifications 60) 0) (scm:string->symbol "state"))))]))))

  (scm:define act
    (scm:let ([_0 (Bench.opaque 20)])
      (scm:lambda ()
        (scm:let ([dummy1 (_0)])
          ((Effect.Console.log (Data.Show.showIntImpl ((runManyTimes dummy1) 0)))))))))
