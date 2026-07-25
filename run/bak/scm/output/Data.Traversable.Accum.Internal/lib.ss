#!r6rs
#!chezscheme
(library
  (Data.Traversable.Accum.Internal lib)
  (export
    StateL
    StateR
    applicativeStateL
    applicativeStateR
    applyStateL
    applyStateR
    functorStateL
    functorStateR
    stateL
    stateR)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define StateR
    (scm:lambda (x0)
      x0))

  (scm:define StateL
    (scm:lambda (x0)
      x0))

  (scm:define stateR
    (scm:lambda (v0)
      v0))

  (scm:define stateL
    (scm:lambda (v0)
      v0))

  (scm:define functorStateR
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (k1)
        (scm:lambda (s2)
          (scm:let ([v3 (k1 s2)])
            (scm:list (scm:cons (scm:string->symbol "accum") (rt:record-ref v3 (scm:string->symbol "accum"))) (scm:cons (scm:string->symbol "value") (f0 (rt:record-ref v3 (scm:string->symbol "value"))))))))))))

  (scm:define functorStateL
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (k1)
        (scm:lambda (s2)
          (scm:let ([v3 (k1 s2)])
            (scm:list (scm:cons (scm:string->symbol "accum") (rt:record-ref v3 (scm:string->symbol "accum"))) (scm:cons (scm:string->symbol "value") (f0 (rt:record-ref v3 (scm:string->symbol "value"))))))))))))

  (scm:define applyStateR
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (f0)
      (scm:lambda (x1)
        (scm:lambda (s2)
          (scm:let*
            ([v3 (x1 s2)]
             [v14 (f0 (rt:record-ref v3 (scm:string->symbol "accum")))])
              (scm:list (scm:cons (scm:string->symbol "accum") (rt:record-ref v14 (scm:string->symbol "accum"))) (scm:cons (scm:string->symbol "value") ((rt:record-ref v14 (scm:string->symbol "value")) (rt:record-ref v3 (scm:string->symbol "value")))))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorStateR))))

  (scm:define applyStateL
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (f0)
      (scm:lambda (x1)
        (scm:lambda (s2)
          (scm:let*
            ([v3 (f0 s2)]
             [v14 (x1 (rt:record-ref v3 (scm:string->symbol "accum")))])
              (scm:list (scm:cons (scm:string->symbol "accum") (rt:record-ref v14 (scm:string->symbol "accum"))) (scm:cons (scm:string->symbol "value") ((rt:record-ref v3 (scm:string->symbol "value")) (rt:record-ref v14 (scm:string->symbol "value")))))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorStateL))))

  (scm:define applicativeStateR
    (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a0)
      (scm:lambda (s1)
        (scm:list (scm:cons (scm:string->symbol "accum") s1) (scm:cons (scm:string->symbol "value") a0))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyStateR))))

  (scm:define applicativeStateL
    (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a0)
      (scm:lambda (s1)
        (scm:list (scm:cons (scm:string->symbol "accum") s1) (scm:cons (scm:string->symbol "value") a0))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyStateL)))))
