#!r6rs
#!chezscheme
(library
  (Control.Extend lib)
  (export
    arrayExtend
    composeCoKleisli
    composeCoKleisliFlipped
    duplicate
    extend
    extendArray
    extendFlipped
    extendFn
    identity)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Functor lib) Data.Functor.)
    (Control.Extend foreign))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define extendFn
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "extend") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (w3)
            (f1 (scm:lambda (w$p4)
              (g2 (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) w3) w$p4)))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
        Data.Functor.functorFn)))))

  (scm:define extendArray
    (scm:list (scm:cons (scm:string->symbol "extend") arrayExtend) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Functor.functorArray))))

  (scm:define extend
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "extend"))))

  (scm:define extendFlipped
    (scm:lambda (dictExtend0)
      (scm:lambda (w1)
        (scm:lambda (f2)
          (((rt:record-ref dictExtend0 (scm:string->symbol "extend")) f2) w1)))))

  (scm:define duplicate
    (scm:lambda (dictExtend0)
      ((rt:record-ref dictExtend0 (scm:string->symbol "extend")) identity)))

  (scm:define composeCoKleisliFlipped
    (scm:lambda (dictExtend0)
      (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (w3)
            (f1 (((rt:record-ref dictExtend0 (scm:string->symbol "extend")) g2) w3)))))))

  (scm:define composeCoKleisli
    (scm:lambda (dictExtend0)
      (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (w3)
            (g2 (((rt:record-ref dictExtend0 (scm:string->symbol "extend")) f1) w3))))))))
