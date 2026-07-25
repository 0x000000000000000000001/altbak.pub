#!r6rs
#!chezscheme
(library
  (Data.Functor lib)
  (export
    arrayMap
    flap
    functorArray
    functorFn
    functorProxy
    map
    mapFlipped
    void
    voidLeft
    voidRight)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Semigroupoid lib) Control.Semigroupoid.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Functor foreign))

  (scm:define map
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "map"))))

  (scm:define mapFlipped
    (scm:lambda (dictFunctor0)
      (scm:lambda (fa1)
        (scm:lambda (f2)
          (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f2) fa1)))))

  (scm:define void
    (scm:lambda (dictFunctor0)
      ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (v1)
        Data.Unit.unit))))

  (scm:define voidLeft
    (scm:lambda (dictFunctor0)
      (scm:lambda (f1)
        (scm:lambda (x2)
          (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (v3)
            x2)) f1)))))

  (scm:define voidRight
    (scm:lambda (dictFunctor0)
      (scm:lambda (x1)
        ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (v2)
          x1)))))

  (scm:define functorProxy
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy)))))

  (scm:define functorFn
    (scm:list (scm:cons (scm:string->symbol "map") (rt:record-ref Control.Semigroupoid.semigroupoidFn (scm:string->symbol "compose")))))

  (scm:define functorArray
    (scm:list (scm:cons (scm:string->symbol "map") arrayMap)))

  (scm:define flap
    (scm:lambda (dictFunctor0)
      (scm:lambda (ff1)
        (scm:lambda (x2)
          (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (f3)
            (f3 x2))) ff1))))))
