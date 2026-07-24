#!r6rs
#!chezscheme
(library
  (Test.Polymorphism lib)
  (export
    act
    describe
    intMonoidish
    mappend_
    mempty_
    polyLoop)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define mempty_
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "mempty_"))))

  (scm:define mappend_
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "mappend_"))))

  (scm:define polyLoop
    (scm:lambda (dictMonoidish0)
      (scm:let ([mempty_11 (rt:record-ref dictMonoidish0 (scm:string->symbol "mempty_"))])
        (scm:lambda (n_init2)
          (scm:lambda (acc_init3)
            (scm:letrec ([go4 (scm:lambda (v5)
              (scm:lambda (v16)
                (scm:cond
                  [(scm:fx=? v5 0) v16]
                  [scm:else ((go4 (scm:fx- v5 1)) (((rt:record-ref dictMonoidish0 (scm:string->symbol "mappend_")) v16) mempty_11))])))])
              ((go4 n_init2) acc_init3)))))))

  (scm:define intMonoidish
    (scm:list (scm:cons (scm:string->symbol "mempty_") 1) (scm:cons (scm:string->symbol "mappend_") (scm:lambda (x0)
      (scm:lambda (y1)
        (scm:fx+ x0 y1))))))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Polymorphism (10M Type Class Dict Lookups):")))

  (scm:define act
    (scm:let ([_0 (Bench.opaque 10000000)])
      (scm:lambda ()
        (scm:let ([dummy1 (_0)])
          ((scm:letrec ([go2 (scm:lambda (v3)
            (scm:lambda (v14)
              (scm:cond
                [(scm:fx=? v3 0) v14]
                [scm:else ((go2 (scm:fx- v3 1)) (scm:fx+ v14 1))])))])
            (Effect.Console.log (Data.Show.showIntImpl ((go2 dummy1) 0))))))))))
