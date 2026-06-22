#!r6rs
#!chezscheme
(library
  (Test.Church lib)
  (export
    act
    addC
    c10
    c100
    c100k
    c10k
    describe
    fromInt
    mulC
    succC
    toInt
    zeroC)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define zeroC
    (scm:lambda (v0)
      (scm:lambda (x1)
        x1)))

  (scm:define toInt
    (scm:lambda (n0)
      ((n0 (scm:lambda (x1)
        (scm:fx+ x1 1))) 0)))

  (scm:define succC
    (scm:lambda (n0)
      (scm:lambda (f1)
        (scm:lambda (x2)
          (f1 ((n0 f1) x2))))))

  (scm:define mulC
    (scm:lambda (m0)
      (scm:lambda (n1)
        (scm:lambda (f2)
          (scm:lambda (x3)
            ((m0 (n1 f2)) x3))))))

  (scm:define fromInt
    (scm:lambda (v0)
      (scm:cond
        [(scm:fx=? v0 0) zeroC]
        [scm:else (scm:let ([_1 (fromInt (scm:fx- v0 1))])
          (scm:lambda (f2)
            (scm:lambda (x3)
              (f2 ((_1 f2) x3)))))])))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Church Numerals (100k Closure Applications):")))

  (scm:define c10
    (fromInt 10))

  (scm:define c100
    (scm:lambda (f0)
      (scm:lambda (x1)
        ((c10 (c10 f0)) x1))))

  (scm:define c10k
    (scm:lambda (f0)
      (scm:lambda (x1)
        ((c10 (c10 (c100 f0))) x1))))

  (scm:define c100k
    (scm:lambda (f0)
      (scm:lambda (x1)
        ((c10 (c10 (c100 (c10 f0)))) x1))))

  (scm:define addC
    (scm:lambda (m0)
      (scm:lambda (n1)
        (scm:lambda (f2)
          (scm:lambda (x3)
            ((m0 f2) ((n1 f2) x3)))))))

  (scm:define act
    (Effect.Console.log (Data.Show.showIntImpl ((c10 (c10 (c100 (c10 (scm:lambda (x0)
      (scm:fx+ x0 1)))))) 0)))))
