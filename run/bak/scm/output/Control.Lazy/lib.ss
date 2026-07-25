#!r6rs
#!chezscheme
(library
  (Control.Lazy lib)
  (export
    defer
    fix
    lazyFn
    lazyUnit)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define lazyUnit
    (scm:list (scm:cons (scm:string->symbol "defer") (scm:lambda (v0)
      Data.Unit.unit))))

  (scm:define lazyFn
    (scm:list (scm:cons (scm:string->symbol "defer") (scm:lambda (f0)
      (scm:lambda (x1)
        ((f0 Data.Unit.unit) x1))))))

  (scm:define defer
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "defer"))))

  (scm:define fix
    (scm:lambda (dictLazy0)
      (scm:lambda (f1)
        (scm:letrec ([go2 ((rt:record-ref dictLazy0 (scm:string->symbol "defer")) (scm:lambda (v3)
          (f1 go2)))])
          go2)))))
