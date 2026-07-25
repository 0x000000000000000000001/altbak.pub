#!r6rs
#!chezscheme
(library
  (Data.Monoid.Endo lib)
  (export
    Endo
    boundedEndo
    eqEndo
    monoidEndo
    ordEndo
    semigroupEndo
    showEndo)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Endo
    (scm:lambda (x0)
      x0))

  (scm:define showEndo
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Endo ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupEndo
    (scm:lambda (dictSemigroupoid0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemigroupoid0 (scm:string->symbol "compose")) v1) v12)))))))

  (scm:define ordEndo
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define monoidEndo
    (scm:lambda (dictCategory0)
      (scm:let*
        ([_1 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))]
         [semigroupEndo12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref _1 (scm:string->symbol "compose")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictCategory0 (scm:string->symbol "identity"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEndo12))))))

  (scm:define eqEndo
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define boundedEndo
    (scm:lambda (dictBounded0)
      dictBounded0)))
