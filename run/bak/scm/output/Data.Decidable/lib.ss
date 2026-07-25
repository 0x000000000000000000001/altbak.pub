#!r6rs
#!chezscheme
(library
  (Data.Decidable lib)
  (export
    decidableComparison
    decidableEquivalence
    decidableOp
    decidablePredicate
    identity
    lose
    lost)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Decide lib) Data.Decide.)
    (prefix (Data.Divisible lib) Data.Divisible.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define lose
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "lose"))))

  (scm:define lost
    (scm:lambda (dictDecidable0)
      ((rt:record-ref dictDecidable0 (scm:string->symbol "lose")) identity)))

  (scm:define decidablePredicate
    (scm:list (scm:cons (scm:string->symbol "lose") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:letrec ([spin2 (scm:lambda (v3)
          (spin2 v3))])
          (spin2 (f0 a1)))))) (scm:cons (scm:string->symbol "Decide0") (scm:lambda (_)
      Data.Decide.choosePredicate)) (scm:cons (scm:string->symbol "Divisible1") (scm:lambda (_)
      Data.Divisible.divisiblePredicate))))

  (scm:define decidableOp
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([chooseOp1 (Data.Decide.chooseOp ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)))]
         [divisibleOp2 (Data.Divisible.divisibleOp dictMonoid0)])
          (scm:list (scm:cons (scm:string->symbol "lose") (scm:lambda (f3)
            (scm:lambda (a4)
              (scm:letrec ([spin5 (scm:lambda (v6)
                (spin5 v6))])
                (spin5 (f3 a4)))))) (scm:cons (scm:string->symbol "Decide0") (scm:lambda (_)
            chooseOp1)) (scm:cons (scm:string->symbol "Divisible1") (scm:lambda (_)
            divisibleOp2))))))

  (scm:define decidableEquivalence
    (scm:list (scm:cons (scm:string->symbol "lose") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:letrec ([spin2 (scm:lambda (v3)
          (spin2 v3))])
          (spin2 (f0 a1)))))) (scm:cons (scm:string->symbol "Decide0") (scm:lambda (_)
      Data.Decide.chooseEquivalence)) (scm:cons (scm:string->symbol "Divisible1") (scm:lambda (_)
      Data.Divisible.divisibleEquivalence))))

  (scm:define decidableComparison
    (scm:list (scm:cons (scm:string->symbol "lose") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:lambda (v2)
          (scm:letrec ([spin3 (scm:lambda (v4)
            (spin3 v4))])
            (spin3 (f0 a1))))))) (scm:cons (scm:string->symbol "Decide0") (scm:lambda (_)
      Data.Decide.chooseComparison)) (scm:cons (scm:string->symbol "Divisible1") (scm:lambda (_)
      Data.Divisible.divisibleComparison)))))
