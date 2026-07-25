#!r6rs
#!chezscheme
(library
  (Data.Divisible lib)
  (export
    conquer
    divisibleComparison
    divisibleEquivalence
    divisibleOp
    divisiblePredicate)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Divide lib) Data.Divide.)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define divisiblePredicate
    (scm:list (scm:cons (scm:string->symbol "conquer") (scm:lambda (v0)
      #t)) (scm:cons (scm:string->symbol "Divide0") (scm:lambda (_)
      Data.Divide.dividePredicate))))

  (scm:define divisibleOp
    (scm:lambda (dictMonoid0)
      (scm:let ([divideOp1 (Data.Divide.divideOp ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)))])
        (scm:list (scm:cons (scm:string->symbol "conquer") (scm:let ([_2 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
          (scm:lambda (v3)
            _2))) (scm:cons (scm:string->symbol "Divide0") (scm:lambda (_)
          divideOp1))))))

  (scm:define divisibleEquivalence
    (scm:list (scm:cons (scm:string->symbol "conquer") (scm:lambda (v0)
      (scm:lambda (v11)
        #t))) (scm:cons (scm:string->symbol "Divide0") (scm:lambda (_)
      Data.Divide.divideEquivalence))))

  (scm:define divisibleComparison
    (scm:list (scm:cons (scm:string->symbol "conquer") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Ordering.EQ))) (scm:cons (scm:string->symbol "Divide0") (scm:lambda (_)
      Data.Divide.divideComparison))))

  (scm:define conquer
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "conquer")))))
