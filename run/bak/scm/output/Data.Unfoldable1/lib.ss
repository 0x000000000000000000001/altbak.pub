#!r6rs
#!chezscheme
(library
  (Data.Unfoldable1 lib)
  (export
    fromJust
    iterateN
    range
    replicate1
    replicate1A
    singleton
    unfoldable1Array
    unfoldable1Maybe
    unfoldr1
    unfoldr1ArrayImpl)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (Data.Unfoldable1 foreign))

  (scm:define fromJust
    (scm:lambda (v0)
      (scm:cond
        [(Data.Maybe.Just? v0) (Data.Maybe.Just-value0 v0)]
        [scm:else (rt:fail)])))

  (scm:define unfoldr1
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "unfoldr1"))))

  (scm:define unfoldable1Maybe
    (scm:list (scm:cons (scm:string->symbol "unfoldr1") (scm:lambda (f0)
      (scm:lambda (b1)
        (Data.Maybe.Just (Data.Tuple.Tuple-value0 (f0 b1))))))))

  (scm:define unfoldable1Array
    (scm:list (scm:cons (scm:string->symbol "unfoldr1") ((((unfoldr1ArrayImpl Data.Maybe.isNothing) fromJust) Data.Tuple.fst) Data.Tuple.snd))))

  (scm:define replicate1
    (scm:lambda (dictUnfoldable10)
      (scm:lambda (n1)
        (scm:lambda (v2)
          (((rt:record-ref dictUnfoldable10 (scm:string->symbol "unfoldr1")) (scm:lambda (i3)
            (scm:cond
              [(scm:fx<=? i3 0) (Data.Tuple.Tuple* v2 Data.Maybe.Nothing)]
              [scm:else (Data.Tuple.Tuple* v2 (Data.Maybe.Just (scm:fx- i3 1)))]))) (scm:fx- n1 1))))))

  (scm:define replicate1A
    (scm:lambda (dictApply0)
      (scm:lambda (dictUnfoldable11)
        (scm:lambda (dictTraversable12)
          (scm:let ([sequence13 ((rt:record-ref dictTraversable12 (scm:string->symbol "sequence1")) dictApply0)])
            (scm:lambda (n4)
              (scm:lambda (m5)
                (sequence13 (((rt:record-ref dictUnfoldable11 (scm:string->symbol "unfoldr1")) (scm:lambda (i6)
                  (scm:cond
                    [(scm:fx<=? i6 0) (Data.Tuple.Tuple* m5 Data.Maybe.Nothing)]
                    [scm:else (Data.Tuple.Tuple* m5 (Data.Maybe.Just (scm:fx- i6 1)))]))) (scm:fx- n4 1))))))))))

  (scm:define singleton
    (scm:lambda (dictUnfoldable10)
      (scm:lambda (v1)
        (((rt:record-ref dictUnfoldable10 (scm:string->symbol "unfoldr1")) (scm:lambda (i2)
          (scm:cond
            [(scm:fx<=? i2 0) (Data.Tuple.Tuple* v1 Data.Maybe.Nothing)]
            [scm:else (Data.Tuple.Tuple* v1 (Data.Maybe.Just (scm:fx- i2 1)))]))) 0))))

  (scm:define range
    (scm:lambda (dictUnfoldable10)
      (scm:lambda (start1)
        (scm:lambda (end2)
          (((rt:record-ref dictUnfoldable10 (scm:string->symbol "unfoldr1")) (scm:let ([_3 (scm:cond
            [(scm:fx>=? end2 start1) 1]
            [scm:else -1])])
            (scm:lambda (i4)
              (scm:let ([i$p5 (scm:fx+ i4 _3)])
                (Data.Tuple.Tuple* i4 (scm:cond
                  [(scm:fx=? i4 end2) Data.Maybe.Nothing]
                  [scm:else (Data.Maybe.Just i$p5)])))))) start1)))))

  (scm:define iterateN
    (scm:lambda (dictUnfoldable10)
      (scm:lambda (n1)
        (scm:lambda (f2)
          (scm:lambda (s3)
            (((rt:record-ref dictUnfoldable10 (scm:string->symbol "unfoldr1")) (scm:lambda (v4)
              (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v4) (scm:cond
                [(scm:fx>? (Data.Tuple.Tuple-value1 v4) 0) (Data.Maybe.Just (Data.Tuple.Tuple* (f2 (Data.Tuple.Tuple-value0 v4)) (scm:fx- (Data.Tuple.Tuple-value1 v4) 1)))]
                [scm:else Data.Maybe.Nothing])))) (Data.Tuple.Tuple* s3 (scm:fx- n1 1)))))))))
