#!r6rs
#!chezscheme
(library
  (Data.Unfoldable lib)
  (export
    fromJust
    fromMaybe
    none
    replicate
    replicateA
    unfoldableArray
    unfoldableMaybe
    unfoldr
    unfoldrArrayImpl)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unfoldable1 lib) Data.Unfoldable1.)
    (prefix (Data.Unit lib) Data.Unit.)
    (Data.Unfoldable foreign))

  (scm:define fromJust
    (scm:lambda (v0)
      (scm:cond
        [(Data.Maybe.Just? v0) (Data.Maybe.Just-value0 v0)]
        [scm:else (rt:fail)])))

  (scm:define unfoldr
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "unfoldr"))))

  (scm:define unfoldableMaybe
    (scm:list (scm:cons (scm:string->symbol "unfoldr") (scm:lambda (f0)
      (scm:lambda (b1)
        (scm:let ([_2 (f0 b1)])
          (scm:cond
            [(Data.Maybe.Just? _2) (Data.Maybe.Just (Data.Tuple.Tuple-value0 (Data.Maybe.Just-value0 _2)))]
            [scm:else Data.Maybe.Nothing]))))) (scm:cons (scm:string->symbol "Unfoldable10") (scm:lambda (_)
      Data.Unfoldable1.unfoldable1Maybe))))

  (scm:define unfoldableArray
    (scm:list (scm:cons (scm:string->symbol "unfoldr") ((((unfoldrArrayImpl Data.Maybe.isNothing) fromJust) Data.Tuple.fst) Data.Tuple.snd)) (scm:cons (scm:string->symbol "Unfoldable10") (scm:lambda (_)
      Data.Unfoldable1.unfoldable1Array))))

  (scm:define replicate
    (scm:lambda (dictUnfoldable0)
      (scm:lambda (n1)
        (scm:lambda (v2)
          (((rt:record-ref dictUnfoldable0 (scm:string->symbol "unfoldr")) (scm:lambda (i3)
            (scm:cond
              [(scm:fx<=? i3 0) Data.Maybe.Nothing]
              [scm:else (Data.Maybe.Just (Data.Tuple.Tuple* v2 (scm:fx- i3 1)))]))) n1)))))

  (scm:define replicateA
    (scm:lambda (dictApplicative0)
      (scm:lambda (dictUnfoldable1)
        (scm:lambda (dictTraversable2)
          (scm:let ([sequence3 ((rt:record-ref dictTraversable2 (scm:string->symbol "sequence")) dictApplicative0)])
            (scm:lambda (n4)
              (scm:lambda (m5)
                (sequence3 (((rt:record-ref dictUnfoldable1 (scm:string->symbol "unfoldr")) (scm:lambda (i6)
                  (scm:cond
                    [(scm:fx<=? i6 0) Data.Maybe.Nothing]
                    [scm:else (Data.Maybe.Just (Data.Tuple.Tuple* m5 (scm:fx- i6 1)))]))) n4)))))))))

  (scm:define none
    (scm:lambda (dictUnfoldable0)
      (((rt:record-ref dictUnfoldable0 (scm:string->symbol "unfoldr")) (scm:lambda (v1)
        Data.Maybe.Nothing)) Data.Unit.unit)))

  (scm:define fromMaybe
    (scm:lambda (dictUnfoldable0)
      ((rt:record-ref dictUnfoldable0 (scm:string->symbol "unfoldr")) (scm:lambda (b1)
        (scm:cond
          [(Data.Maybe.Just? b1) (Data.Maybe.Just (Data.Tuple.Tuple* (Data.Maybe.Just-value0 b1) Data.Maybe.Nothing))]
          [scm:else Data.Maybe.Nothing]))))))
