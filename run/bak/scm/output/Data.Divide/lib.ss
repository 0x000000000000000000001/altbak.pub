#!r6rs
#!chezscheme
(library
  (Data.Divide lib)
  (export
    divide
    divideComparison
    divideEquivalence
    divideOp
    dividePredicate
    divided
    identity)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Comparison lib) Data.Comparison.)
    (prefix (Data.Equivalence lib) Data.Equivalence.)
    (prefix (Data.Op lib) Data.Op.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Predicate lib) Data.Predicate.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define dividePredicate
    (scm:list (scm:cons (scm:string->symbol "divide") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (a3)
            (scm:let ([v24 (f0 a3)])
              (scm:and (v1 (Data.Tuple.Tuple-value0 v24)) (v12 (Data.Tuple.Tuple-value1 v24))))))))) (scm:cons (scm:string->symbol "Contravariant0") (scm:lambda (_)
      Data.Predicate.contravariantPredicate))))

  (scm:define divideOp
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "divide") (scm:lambda (f1)
        (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:lambda (a4)
              (scm:let ([v25 (f1 a4)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (v2 (Data.Tuple.Tuple-value0 v25))) (v13 (Data.Tuple.Tuple-value1 v25))))))))) (scm:cons (scm:string->symbol "Contravariant0") (scm:lambda (_)
        Data.Op.contravariantOp)))))

  (scm:define divideEquivalence
    (scm:list (scm:cons (scm:string->symbol "divide") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (a3)
            (scm:lambda (b4)
              (scm:let*
                ([v25 (f0 a3)]
                 [v36 (f0 b4)])
                  (scm:and ((v1 (Data.Tuple.Tuple-value0 v25)) (Data.Tuple.Tuple-value0 v36)) ((v12 (Data.Tuple.Tuple-value1 v25)) (Data.Tuple.Tuple-value1 v36)))))))))) (scm:cons (scm:string->symbol "Contravariant0") (scm:lambda (_)
      Data.Equivalence.contravariantEquivalence))))

  (scm:define divideComparison
    (scm:list (scm:cons (scm:string->symbol "divide") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (a3)
            (scm:lambda (b4)
              (scm:let*
                ([v25 (f0 a3)]
                 [v36 (f0 b4)]
                 [_7 ((v1 (Data.Tuple.Tuple-value0 v25)) (Data.Tuple.Tuple-value0 v36))]
                 [_8 ((v12 (Data.Tuple.Tuple-value1 v25)) (Data.Tuple.Tuple-value1 v36))])
                  (scm:cond
                    [(Data.Ordering.LT? _7) Data.Ordering.LT]
                    [(Data.Ordering.GT? _7) Data.Ordering.GT]
                    [(Data.Ordering.EQ? _7) _8]
                    [scm:else (rt:fail)])))))))) (scm:cons (scm:string->symbol "Contravariant0") (scm:lambda (_)
      Data.Comparison.contravariantComparison))))

  (scm:define divide
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "divide"))))

  (scm:define divided
    (scm:lambda (dictDivide0)
      ((rt:record-ref dictDivide0 (scm:string->symbol "divide")) identity))))
