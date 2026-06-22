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
    (Effect.Console.log (Data.Show.showIntImpl (scm:letrec ([go0 (scm:lambda (v1)
      (scm:lambda (v12)
        (scm:cond
          [(scm:fx=? v1 0) v12]
          [scm:else ((go0 (scm:fx- v1 1)) (scm:fx+ v12 1))])))])
      ((go0 10000000) 0))))))
