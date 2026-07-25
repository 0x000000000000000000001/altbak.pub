#!r6rs
#!chezscheme
(library
  (Control.Applicative lib)
  (export
    applicativeArray
    applicativeFn
    applicativeProxy
    liftA1
    pure
    unless
    when)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Apply lib) Control.Apply.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Type.Proxy lib) Type.Proxy.))

  (scm:define pure
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "pure"))))

  (scm:define unless
    (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [(scm:not v1) v12]
            [v1 ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) Data.Unit.unit)]
            [scm:else (rt:fail)])))))

  (scm:define when
    (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [v1 v12]
            [scm:else ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) Data.Unit.unit)])))))

  (scm:define liftA1
    (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "apply")) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) f1)) a2)))))

  (scm:define applicativeProxy
    (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (v0)
      Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      Control.Apply.applyProxy))))

  (scm:define applicativeFn
    (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (x0)
      (scm:lambda (v1)
        x0))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      Control.Apply.applyFn))))

  (scm:define applicativeArray
    (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (x0)
      (rt:make-array x0))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      Control.Apply.applyArray)))))
