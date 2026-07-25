#!r6rs
#!chezscheme
(library
  (Data.Functor.Costar lib)
  (export
    Costar
    applicativeCostar
    applyCostar
    bifunctorCostar
    bindCostar
    categoryCostar
    closedCostar
    distributiveCostar
    functorCostar
    hoistCostar
    invariantCostar
    monadCostar
    newtypeCostar
    profunctorCostar
    semigroupoidCostar
    strongCostar)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define Costar
    (scm:lambda (x0)
      x0))

  (scm:define semigroupoidCostar
    (scm:lambda (dictExtend0)
      (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (w3)
            (v1 (((rt:record-ref dictExtend0 (scm:string->symbol "extend")) v12) w3)))))))))

  (scm:define profunctorCostar
    (scm:lambda (dictFunctor0)
      (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1)])
              (scm:lambda (x5)
                (g2 (v3 (_4 x5))))))))))))

  (scm:define strongCostar
    (scm:lambda (dictComonad0)
      (scm:let*
        ([Functor01 ((rt:record-ref ((rt:record-ref dictComonad0 (scm:string->symbol "Extend0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [profunctorCostar12 (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (v4)
              (scm:let ([_5 ((rt:record-ref Functor01 (scm:string->symbol "map")) f2)])
                (scm:lambda (x6)
                  (g3 (v4 (_5 x6))))))))))])
          (scm:list (scm:cons (scm:string->symbol "first") (scm:lambda (v3)
            (scm:lambda (x4)
              (Data.Tuple.Tuple* (v3 (((rt:record-ref Functor01 (scm:string->symbol "map")) Data.Tuple.fst) x4)) (Data.Tuple.Tuple-value1 ((rt:record-ref dictComonad0 (scm:string->symbol "extract")) x4)))))) (scm:cons (scm:string->symbol "second") (scm:lambda (v3)
            (scm:lambda (x4)
              (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 ((rt:record-ref dictComonad0 (scm:string->symbol "extract")) x4)) (v3 (((rt:record-ref Functor01 (scm:string->symbol "map")) Data.Tuple.snd) x4)))))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
            profunctorCostar12))))))

  (scm:define newtypeCostar
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define hoistCostar
    (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (x2)
          (v1 (f0 x2))))))

  (scm:define functorCostar
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (x2)
          (f0 (v1 x2))))))))

  (scm:define invariantCostar
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v2)
          (scm:lambda (x3)
            (f0 (v2 x3)))))))))

  (rt:define-lazy $lazy-distributiveCostar "distributiveCostar" "Data.Functor.Costar"
    (scm:list (scm:cons (scm:string->symbol "distribute") (scm:lambda (dictFunctor0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (v3)
            (v3 a2))) f1))))) (scm:cons (scm:string->symbol "collect") (scm:lambda (dictFunctor0)
      (scm:lambda (f1)
        (scm:let*
          ([_2 ((rt:record-ref ($lazy-distributiveCostar) (scm:string->symbol "distribute")) dictFunctor0)]
           [_3 ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1)])
            (scm:lambda (x4)
              (_2 (_3 x4))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorCostar))))

  (scm:define distributiveCostar
    ($lazy-distributiveCostar))

  (scm:define closedCostar
    (scm:lambda (dictFunctor0)
      (scm:let ([profunctorCostar11 (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1)])
              (scm:lambda (x5)
                (g2 (v3 (_4 x5))))))))))])
        (scm:list (scm:cons (scm:string->symbol "closed") (scm:lambda (v2)
          (scm:lambda (g3)
            (scm:lambda (x4)
              (v2 (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (v15)
                (v15 x4))) g3)))))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
          profunctorCostar11))))))

  (scm:define categoryCostar
    (scm:lambda (dictComonad0)
      (scm:let*
        ([_1 ((rt:record-ref dictComonad0 (scm:string->symbol "Extend0")) (scm:quote undefined))]
         [semigroupoidCostar12 (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:lambda (w4)
              (v2 (((rt:record-ref _1 (scm:string->symbol "extend")) v13) w4)))))))])
          (scm:list (scm:cons (scm:string->symbol "identity") (rt:record-ref dictComonad0 (scm:string->symbol "extract"))) (scm:cons (scm:string->symbol "Semigroupoid0") (scm:lambda (_)
            semigroupoidCostar12))))))

  (scm:define bifunctorCostar
    (scm:lambda (dictContravariant0)
      (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref dictContravariant0 (scm:string->symbol "cmap")) f1)])
              (scm:lambda (x5)
                (g2 (v3 (_4 x5))))))))))))

  (scm:define applyCostar
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (a2)
          ((v0 a2) (v11 a2)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorCostar))))

  (scm:define bindCostar
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (f1)
        (scm:lambda (x2)
          ((f1 (v0 x2)) x2))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyCostar))))

  (scm:define applicativeCostar
    (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a0)
      (scm:lambda (v1)
        a0))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyCostar))))

  (scm:define monadCostar
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeCostar)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindCostar)))))
