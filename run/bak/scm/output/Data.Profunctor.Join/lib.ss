#!r6rs
#!chezscheme
(library
  (Data.Profunctor.Join lib)
  (export
    Join
    eqJoin
    invariantJoin
    monoidJoin
    newtypeJoin
    ordJoin
    semigroupJoin
    showJoin)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Join
    (scm:lambda (x0)
      x0))

  (scm:define showJoin
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Join ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupJoin
    (scm:lambda (dictSemigroupoid0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemigroupoid0 (scm:string->symbol "compose")) v1) v12)))))))

  (scm:define ordJoin
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define newtypeJoin
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidJoin
    (scm:lambda (dictCategory0)
      (scm:let*
        ([_1 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))]
         [semigroupJoin12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref _1 (scm:string->symbol "compose")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictCategory0 (scm:string->symbol "identity"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupJoin12))))))

  (scm:define invariantJoin
    (scm:lambda (dictProfunctor0)
      (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            ((((rt:record-ref dictProfunctor0 (scm:string->symbol "dimap")) g2) f1) v3))))))))

  (scm:define eqJoin
    (scm:lambda (dictEq0)
      dictEq0)))
