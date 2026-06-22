#!r6rs
#!chezscheme
(library
  (Data.Ordering lib)
  (export
    EQ
    EQ?
    GT
    GT?
    LT
    LT?
    eqOrdering
    invert
    semigroupOrdering
    showOrdering)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define LT
    (scm:quote LT))

  (scm:define LT?
    (scm:lambda (v)
      (scm:eq? (scm:quote LT) v)))

  (scm:define GT
    (scm:quote GT))

  (scm:define GT?
    (scm:lambda (v)
      (scm:eq? (scm:quote GT) v)))

  (scm:define EQ
    (scm:quote EQ))

  (scm:define EQ?
    (scm:lambda (v)
      (scm:eq? (scm:quote EQ) v)))

  (scm:define showOrdering
    (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v0)
      (scm:cond
        [(LT? v0) (rt:string->pstring "LT")]
        [(GT? v0) (rt:string->pstring "GT")]
        [(EQ? v0) (rt:string->pstring "EQ")]
        [scm:else (rt:fail)])))))

  (scm:define semigroupOrdering
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(LT? v0) LT]
          [(GT? v0) GT]
          [(EQ? v0) v11]
          [scm:else (rt:fail)]))))))

  (scm:define invert
    (scm:lambda (v0)
      (scm:cond
        [(GT? v0) LT]
        [(EQ? v0) EQ]
        [(LT? v0) GT]
        [scm:else (rt:fail)])))

  (scm:define eqOrdering
    (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(LT? v0) (LT? v11)]
          [(GT? v0) (GT? v11)]
          [scm:else (scm:and (EQ? v0) (EQ? v11))])))))))
