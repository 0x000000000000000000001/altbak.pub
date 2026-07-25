#!r6rs
#!chezscheme
(library
  (Type.Equality lib)
  (export
    from
    proof
    refl
    to)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define refl
    (scm:list (scm:cons (scm:string->symbol "proof") (scm:lambda (a0)
      a0)) (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define proof
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "proof"))))

  (scm:define to
    (scm:lambda (dictTypeEquals0)
      ((rt:record-ref dictTypeEquals0 (scm:string->symbol "proof")) (scm:lambda (a1)
        a1))))

  (scm:define from
    (scm:lambda (dictTypeEquals0)
      ((rt:record-ref dictTypeEquals0 (scm:string->symbol "proof")) (scm:lambda (a1)
        a1)))))
