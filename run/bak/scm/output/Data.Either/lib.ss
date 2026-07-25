#!r6rs
#!chezscheme
(library
  (Data.Either lib)
  (export
    Left
    Left-value0
    Left?
    Right
    Right-value0
    Right?
    altEither
    applicativeEither
    applyEither
    bindEither
    blush
    boundedEither
    choose
    either
    eq1Either
    eqEither
    extendEither
    fromLeft
    fromLeft$p
    fromRight
    fromRight$p
    functorEither
    genericEither
    hush
    invariantEither
    isLeft
    isRight
    monadEither
    note
    note$p
    ord1Either
    ordEither
    semigroupEither
    showEither)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define-record-type (Left$ Left Left?)
    (scm:fields (scm:immutable value0 Left-value0)))

  (scm:define-record-type (Right$ Right Right?)
    (scm:fields (scm:immutable value0 Right-value0)))

  (scm:define showEither
    (scm:lambda (dictShow0)
      (scm:lambda (dictShow11)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (scm:cond
            [(Left? v2) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Left ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Left-value0 v2))) (rt:string->pstring ")"))]
            [(Right? v2) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Right ") ((rt:record-ref dictShow11 (scm:string->symbol "show")) (Right-value0 v2))) (rt:string->pstring ")"))]
            [scm:else (rt:fail)])))))))

  (scm:define note$p
    (scm:lambda (f0)
      (scm:lambda (v21)
        (scm:cond
          [(Data.Maybe.Nothing? v21) (Left (f0 Data.Unit.unit))]
          [(Data.Maybe.Just? v21) (Right (Data.Maybe.Just-value0 v21))]
          [scm:else (rt:fail)]))))

  (scm:define note
    (scm:lambda (a0)
      (scm:lambda (v21)
        (scm:cond
          [(Data.Maybe.Nothing? v21) (Left a0)]
          [(Data.Maybe.Just? v21) (Right (Data.Maybe.Just-value0 v21))]
          [scm:else (rt:fail)]))))

  (scm:define genericEither
    (scm:list (scm:cons (scm:string->symbol "to") (scm:lambda (x0)
      (scm:cond
        [(Data.Generic.Rep.Inl? x0) (Left (Data.Generic.Rep.Inl-value0 x0))]
        [(Data.Generic.Rep.Inr? x0) (Right (Data.Generic.Rep.Inr-value0 x0))]
        [scm:else (rt:fail)]))) (scm:cons (scm:string->symbol "from") (scm:lambda (x0)
      (scm:cond
        [(Left? x0) (Data.Generic.Rep.Inl (Left-value0 x0))]
        [(Right? x0) (Data.Generic.Rep.Inr (Right-value0 x0))]
        [scm:else (rt:fail)])))))

  (scm:define functorEither
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (scm:cond
          [(Left? m1) (Left (Left-value0 m1))]
          [(Right? m1) (Right (f0 (Right-value0 m1)))]
          [scm:else (rt:fail)]))))))

  (scm:define invariantEither
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (m2)
          (scm:cond
            [(Left? m2) (Left (Left-value0 m2))]
            [(Right? m2) (Right (f0 (Right-value0 m2)))]
            [scm:else (rt:fail)])))))))

  (scm:define fromRight$p
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Right? v11) (Right-value0 v11)]
          [scm:else (v0 Data.Unit.unit)]))))

  (scm:define fromRight
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Right? v11) (Right-value0 v11)]
          [scm:else v0]))))

  (scm:define fromLeft$p
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Left? v11) (Left-value0 v11)]
          [scm:else (v0 Data.Unit.unit)]))))

  (scm:define fromLeft
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Left? v11) (Left-value0 v11)]
          [scm:else v0]))))

  (scm:define extendEither
    (scm:list (scm:cons (scm:string->symbol "extend") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Left? v11) (Left (Left-value0 v11))]
          [scm:else (Right (v0 v11))])))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorEither))))

  (scm:define eqEither
    (scm:lambda (dictEq0)
      (scm:lambda (dictEq11)
        (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:cond
              [(Left? x2) (scm:and (Left? y3) (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Left-value0 x2)) (Left-value0 y3)))]
              [scm:else (scm:and (Right? x2) (scm:and (Right? y3) (((rt:record-ref dictEq11 (scm:string->symbol "eq")) (Right-value0 x2)) (Right-value0 y3))))]))))))))

  (scm:define ordEither
    (scm:lambda (dictOrd0)
      (scm:let ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))])
        (scm:lambda (dictOrd12)
          (scm:let*
            ([_3 ((rt:record-ref dictOrd12 (scm:string->symbol "Eq0")) (scm:quote undefined))]
             [eqEither24 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x4)
              (scm:lambda (y5)
                (scm:cond
                  [(Left? x4) (scm:and (Left? y5) (((rt:record-ref _1 (scm:string->symbol "eq")) (Left-value0 x4)) (Left-value0 y5)))]
                  [scm:else (scm:and (Right? x4) (scm:and (Right? y5) (((rt:record-ref _3 (scm:string->symbol "eq")) (Right-value0 x4)) (Right-value0 y5))))])))))])
              (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x5)
                (scm:lambda (y6)
                  (scm:cond
                    [(Left? x5) (scm:cond
                      [(Left? y6) (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Left-value0 x5)) (Left-value0 y6))]
                      [scm:else Data.Ordering.LT])]
                    [(Left? y6) Data.Ordering.GT]
                    [(scm:and (Right? x5) (Right? y6)) (((rt:record-ref dictOrd12 (scm:string->symbol "compare")) (Right-value0 x5)) (Right-value0 y6))]
                    [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                eqEither24))))))))

  (scm:define eq1Either
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq11)
        (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:cond
              [(Left? x2) (scm:and (Left? y3) (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Left-value0 x2)) (Left-value0 y3)))]
              [scm:else (scm:and (Right? x2) (scm:and (Right? y3) (((rt:record-ref dictEq11 (scm:string->symbol "eq")) (Right-value0 x2)) (Right-value0 y3))))]))))))))

  (scm:define ord1Either
    (scm:lambda (dictOrd0)
      (scm:let*
        ([ordEither11 (ordEither dictOrd0)]
         [_2 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))]
         [eq1Either13 (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq13)
          (scm:lambda (x4)
            (scm:lambda (y5)
              (scm:cond
                [(Left? x4) (scm:and (Left? y5) (((rt:record-ref _2 (scm:string->symbol "eq")) (Left-value0 x4)) (Left-value0 y5)))]
                [scm:else (scm:and (Right? x4) (scm:and (Right? y5) (((rt:record-ref dictEq13 (scm:string->symbol "eq")) (Right-value0 x4)) (Right-value0 y5))))]))))))])
          (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd14)
            (rt:record-ref (ordEither11 dictOrd14) (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
            eq1Either13))))))

  (scm:define either
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Left? v22) (v0 (Left-value0 v22))]
            [(Right? v22) (v11 (Right-value0 v22))]
            [scm:else (rt:fail)])))))

  (scm:define hush
    (scm:lambda (v20)
      (scm:cond
        [(Left? v20) Data.Maybe.Nothing]
        [(Right? v20) (Data.Maybe.Just (Right-value0 v20))]
        [scm:else (rt:fail)])))

  (scm:define isLeft
    (scm:lambda (v20)
      (scm:cond
        [(Left? v20) #t]
        [(Right? v20) #f]
        [scm:else (rt:fail)])))

  (scm:define isRight
    (scm:lambda (v20)
      (scm:cond
        [(Left? v20) #f]
        [(Right? v20) #t]
        [scm:else (rt:fail)])))

  (scm:define choose
    (scm:lambda (dictAlt0)
      (scm:let ([_1 ((rt:record-ref dictAlt0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (a2)
          (scm:lambda (b3)
            (((rt:record-ref dictAlt0 (scm:string->symbol "alt")) (((rt:record-ref _1 (scm:string->symbol "map")) Left) a2)) (((rt:record-ref _1 (scm:string->symbol "map")) Right) b3)))))))

  (scm:define boundedEither
    (scm:lambda (dictBounded0)
      (scm:let*
        ([bottom1 (rt:record-ref dictBounded0 (scm:string->symbol "bottom"))]
         [ordEither12 (ordEither ((rt:record-ref dictBounded0 (scm:string->symbol "Ord0")) (scm:quote undefined)))])
          (scm:lambda (dictBounded13)
            (scm:let ([ordEither24 (ordEither12 ((rt:record-ref dictBounded13 (scm:string->symbol "Ord0")) (scm:quote undefined)))])
              (scm:list (scm:cons (scm:string->symbol "top") (Right (rt:record-ref dictBounded13 (scm:string->symbol "top")))) (scm:cons (scm:string->symbol "bottom") (Left bottom1)) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
                ordEither24))))))))

  (scm:define blush
    (scm:lambda (v20)
      (scm:cond
        [(Left? v20) (Data.Maybe.Just (Left-value0 v20))]
        [(Right? v20) Data.Maybe.Nothing]
        [scm:else (rt:fail)])))

  (scm:define applyEither
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Left? v0) (Left (Left-value0 v0))]
          [(Right? v0) (scm:cond
            [(Left? v11) (Left (Left-value0 v11))]
            [(Right? v11) (Right ((Right-value0 v0) (Right-value0 v11)))]
            [scm:else (rt:fail)])]
          [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorEither))))

  (scm:define bindEither
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v20)
      (scm:cond
        [(Left? v20) (scm:let ([_1 (Left-value0 v20)])
          (scm:lambda (v2)
            (Left _1)))]
        [(Right? v20) (scm:let ([_1 (Right-value0 v20)])
          (scm:lambda (f2)
            (f2 _1)))]
        [scm:else (rt:fail)]))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyEither))))

  (scm:define semigroupEither
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:cond
            [(Left? x1) (Left (Left-value0 x1))]
            [(Right? x1) (scm:cond
              [(Left? y2) (Left (Left-value0 y2))]
              [(Right? y2) (Right (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (Right-value0 x1)) (Right-value0 y2)))]
              [scm:else (rt:fail)])]
            [scm:else (rt:fail)])))))))

  (scm:define applicativeEither
    (scm:list (scm:cons (scm:string->symbol "pure") Right) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyEither))))

  (scm:define monadEither
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeEither)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindEither))))

  (scm:define altEither
    (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Left? v0) v11]
          [scm:else v0])))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorEither)))))
