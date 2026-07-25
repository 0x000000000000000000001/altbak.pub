#!r6rs
#!chezscheme
(library
  (Data.Functor.Coproduct.Inject lib)
  (export
    inj
    injectLeft
    injectReflexive
    injectRight
    prj)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Maybe lib) Data.Maybe.))

  (scm:define prj
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "prj"))))

  (scm:define injectReflexive
    (scm:list (scm:cons (scm:string->symbol "inj") (scm:lambda (x0)
      x0)) (scm:cons (scm:string->symbol "prj") Data.Maybe.Just)))

  (scm:define injectLeft
    (scm:list (scm:cons (scm:string->symbol "inj") (scm:lambda (x0)
      (Data.Either.Left x0))) (scm:cons (scm:string->symbol "prj") (scm:lambda (v20)
      (scm:cond
        [(Data.Either.Left? v20) (Data.Maybe.Just (Data.Either.Left-value0 v20))]
        [(Data.Either.Right? v20) Data.Maybe.Nothing]
        [scm:else (rt:fail)])))))

  (scm:define inj
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "inj"))))

  (scm:define injectRight
    (scm:lambda (dictInject0)
      (scm:list (scm:cons (scm:string->symbol "inj") (scm:lambda (x1)
        (Data.Either.Right ((rt:record-ref dictInject0 (scm:string->symbol "inj")) x1)))) (scm:cons (scm:string->symbol "prj") (scm:lambda (v21)
        (scm:cond
          [(Data.Either.Left? v21) Data.Maybe.Nothing]
          [(Data.Either.Right? v21) ((rt:record-ref dictInject0 (scm:string->symbol "prj")) (Data.Either.Right-value0 v21))]
          [scm:else (rt:fail)])))))))
