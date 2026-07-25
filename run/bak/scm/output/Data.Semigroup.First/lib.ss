#!r6rs
#!chezscheme
(library
  (Data.Semigroup.First lib)
  (export
    First
    applicativeFirst
    applyFirst
    bindFirst
    boundedFirst
    eq1First
    eqFirst
    functorFirst
    monadFirst
    ord1First
    ordFirst
    semigroupFirst
    showFirst)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define First
    (scm:lambda (x0)
      x0))

  (scm:define showFirst
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(First ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupFirst
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (x0)
      (scm:lambda (v1)
        x0)))))

  (scm:define ordFirst
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define functorFirst
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (f0 m1))))))

  (scm:define eqFirst
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1First
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (rt:record-ref dictEq0 (scm:string->symbol "eq"))))))

  (scm:define ord1First
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref dictOrd0 (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1First))))

  (scm:define boundedFirst
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define applyFirst
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (v0 v11)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorFirst))))

  (scm:define bindFirst
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (f1)
        (f1 v0)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyFirst))))

  (scm:define applicativeFirst
    (scm:list (scm:cons (scm:string->symbol "pure") First) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyFirst))))

  (scm:define monadFirst
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeFirst)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindFirst)))))
