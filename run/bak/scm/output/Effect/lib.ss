#!r6rs
#!chezscheme
(library
  (Effect lib)
  (export
    applicativeEffect
    applyEffect
    bindE
    bindEffect
    forE
    foreachE
    functorEffect
    monadEffect
    monoidEffect
    pureE
    semigroupEffect
    untilE
    whileE)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Effect foreign))

  (rt:define-lazy $lazy-monadEffect "monadEffect" "Effect"
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      ($lazy-applicativeEffect))) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      ($lazy-bindEffect)))))

  (rt:define-lazy $lazy-bindEffect "bindEffect" "Effect"
    (scm:list (scm:cons (scm:string->symbol "bind") bindE) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      ($lazy-applyEffect)))))

  (rt:define-lazy $lazy-applyEffect "applyEffect" "Effect"
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:lambda ()
          (scm:let*
            ([f$p2 (f0)]
             [a$p3 (a1)])
              (((rt:record-ref ($lazy-applicativeEffect) (scm:string->symbol "pure")) (f$p2 a$p3)))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      ($lazy-functorEffect)))))

  (rt:define-lazy $lazy-applicativeEffect "applicativeEffect" "Effect"
    (scm:list (scm:cons (scm:string->symbol "pure") pureE) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      ($lazy-applyEffect)))))

  (rt:define-lazy $lazy-functorEffect "functorEffect" "Effect"
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:lambda ()
          (scm:let ([a$p2 (a1)])
            (f0 a$p2))))))))

  (scm:define monadEffect
    ($lazy-monadEffect))

  (scm:define bindEffect
    ($lazy-bindEffect))

  (scm:define applyEffect
    ($lazy-applyEffect))

  (scm:define applicativeEffect
    ($lazy-applicativeEffect))

  (scm:define functorEffect
    ($lazy-functorEffect))

  (scm:define semigroupEffect
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (a1)
        (scm:lambda (b2)
          (scm:lambda ()
            (scm:let*
              ([a$p3 (a1)]
               [a$p4 (b2)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p3) a$p4)))))))))

  (scm:define monoidEffect
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffect12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda ()
              (scm:let*
                ([a$p4 (a2)]
                 [a$p5 (b3)])
                  (((rt:record-ref _1 (scm:string->symbol "append")) a$p4) a$p5)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:let ([_3 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
            (scm:lambda ()
              _3))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffect12)))))))
