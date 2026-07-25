#!r6rs
#!chezscheme
(library
  (Data.Profunctor.Split lib)
  (export
    SplitF
    SplitF*
    SplitF-value0
    SplitF-value1
    SplitF-value2
    SplitF?
    functorSplit
    hoistSplit
    identity
    liftSplit
    lowerSplit
    profunctorSplit
    split
    unSplit)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define-record-type (SplitF$ SplitF* SplitF?)
    (scm:fields (scm:immutable value0 SplitF-value0) (scm:immutable value1 SplitF-value1) (scm:immutable value2 SplitF-value2)))

  (scm:define SplitF
    (scm:lambda (value0)
      (scm:lambda (value1)
        (scm:lambda (value2)
          (SplitF* value0 value1 value2)))))

  (scm:define unSplit
    (scm:lambda (f0)
      (scm:lambda (v1)
        (((f0 (SplitF-value0 v1)) (SplitF-value1 v1)) (SplitF-value2 v1)))))

  (scm:define split
    (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (fx2)
          (SplitF* f0 g1 fx2)))))

  (scm:define profunctorSplit
    (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (v2)
          (SplitF* (scm:lambda (x3)
            ((SplitF-value0 v2) (f0 x3))) (scm:lambda (x3)
            (g1 ((SplitF-value1 v2) x3))) (SplitF-value2 v2))))))))

  (scm:define lowerSplit
    (scm:lambda (dictInvariant0)
      (scm:lambda (v1)
        ((((rt:record-ref dictInvariant0 (scm:string->symbol "imap")) (SplitF-value1 v1)) (SplitF-value0 v1)) (SplitF-value2 v1)))))

  (scm:define liftSplit
    (scm:lambda (fx0)
      (SplitF* identity identity fx0)))

  (scm:define hoistSplit
    (scm:lambda (nat0)
      (scm:lambda (v1)
        (SplitF* (SplitF-value0 v1) (SplitF-value1 v1) (nat0 (SplitF-value2 v1))))))

  (scm:define functorSplit
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (v1)
        (SplitF* (SplitF-value0 v1) (scm:lambda (x2)
          (f0 ((SplitF-value1 v1) x2))) (SplitF-value2 v1))))))))
