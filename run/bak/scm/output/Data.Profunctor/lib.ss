#!r6rs
#!chezscheme
(library
  (Data.Profunctor lib)
  (export
    arr
    dimap
    identity
    lcmap
    profunctorFn
    rmap
    unwrapIso
    wrapIso)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define profunctorFn
    (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (a2b0)
      (scm:lambda (c2d1)
        (scm:lambda (b2c2)
          (scm:lambda (x3)
            (c2d1 (b2c2 (a2b0 x3))))))))))

  (scm:define dimap
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "dimap"))))

  (scm:define lcmap
    (scm:lambda (dictProfunctor0)
      (scm:lambda (a2b1)
        (((rt:record-ref dictProfunctor0 (scm:string->symbol "dimap")) a2b1) identity))))

  (scm:define rmap
    (scm:lambda (dictProfunctor0)
      (scm:lambda (b2c1)
        (((rt:record-ref dictProfunctor0 (scm:string->symbol "dimap")) identity) b2c1))))

  (scm:define unwrapIso
    (scm:lambda (dictProfunctor0)
      (scm:lambda (_)
        (((rt:record-ref dictProfunctor0 (scm:string->symbol "dimap")) Unsafe.Coerce.unsafeCoerce) Unsafe.Coerce.unsafeCoerce))))

  (scm:define wrapIso
    (scm:lambda (dictProfunctor0)
      (scm:lambda (_)
        (scm:lambda (v2)
          (((rt:record-ref dictProfunctor0 (scm:string->symbol "dimap")) Unsafe.Coerce.unsafeCoerce) Unsafe.Coerce.unsafeCoerce)))))

  (scm:define arr
    (scm:lambda (dictCategory0)
      (scm:let ([identity11 (rt:record-ref dictCategory0 (scm:string->symbol "identity"))])
        (scm:lambda (dictProfunctor2)
          (scm:lambda (f3)
            ((((rt:record-ref dictProfunctor2 (scm:string->symbol "dimap")) identity) f3) identity11)))))))
