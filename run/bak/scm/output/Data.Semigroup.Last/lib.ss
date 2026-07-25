#!r6rs
#!chezscheme
(library
  (Data.Semigroup.Last lib)
  (export
    Last
    applicativeLast
    applyLast
    bindLast
    boundedLast
    eq1Last
    eqLast
    functorLast
    monadLast
    ord1Last
    ordLast
    semigroupLast
    showLast)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Last
    (scm:lambda (x0)
      x0))

  (scm:define showLast
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Last ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupLast
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (x1)
        x1)))))

  (scm:define ordLast
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define functorLast
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (f0 m1))))))

  (scm:define eqLast
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1Last
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (rt:record-ref dictEq0 (scm:string->symbol "eq"))))))

  (scm:define ord1Last
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref dictOrd0 (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1Last))))

  (scm:define boundedLast
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define applyLast
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (v0 v11)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorLast))))

  (scm:define bindLast
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (f1)
        (f1 v0)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyLast))))

  (scm:define applicativeLast
    (scm:list (scm:cons (scm:string->symbol "pure") Last) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyLast))))

  (scm:define monadLast
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeLast)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindLast)))))
