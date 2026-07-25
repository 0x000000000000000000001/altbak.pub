#!r6rs
#!chezscheme
(library
  (Data.Maybe lib)
  (export
    Just
    Just-value0
    Just?
    Nothing
    Nothing?
    altMaybe
    alternativeMaybe
    applicativeMaybe
    applyMaybe
    bindMaybe
    boundedMaybe
    eq1Maybe
    eqMaybe
    extendMaybe
    fromJust
    fromMaybe
    fromMaybe$p
    functorMaybe
    genericMaybe
    identity
    invariantMaybe
    isJust
    isNothing
    maybe
    maybe$p
    monadMaybe
    monoidMaybe
    optional
    ord1Maybe
    ordMaybe
    plusMaybe
    semigroupMaybe
    semiringMaybe
    showMaybe)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define Nothing
    (scm:quote Nothing))

  (scm:define Nothing?
    (scm:lambda (v)
      (scm:eq? (scm:quote Nothing) v)))

  (scm:define-record-type (Just$ Just Just?)
    (scm:fields (scm:immutable value0 Just-value0)))

  (scm:define showMaybe
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (scm:cond
          [(Just? v1) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Just ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Just-value0 v1))) (rt:string->pstring ")"))]
          [(Nothing? v1) (rt:string->pstring "Nothing")]
          [scm:else (rt:fail)]))))))

  (scm:define semigroupMaybe
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [(Nothing? v1) v12]
            [(Nothing? v12) v1]
            [(scm:and (Just? v1) (Just? v12)) (Just (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (Just-value0 v1)) (Just-value0 v12)))]
            [scm:else (rt:fail)])))))))

  (scm:define optional
    (scm:lambda (dictAlt0)
      (scm:lambda (dictApplicative1)
        (scm:lambda (a2)
          (((rt:record-ref dictAlt0 (scm:string->symbol "alt")) (((rt:record-ref ((rt:record-ref dictAlt0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Just) a2)) ((rt:record-ref dictApplicative1 (scm:string->symbol "pure")) Nothing))))))

  (scm:define monoidMaybe
    (scm:lambda (dictSemigroup0)
      (scm:let ([semigroupMaybe11 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [(Nothing? v1) v12]
            [(Nothing? v12) v1]
            [(scm:and (Just? v1) (Just? v12)) (Just (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (Just-value0 v1)) (Just-value0 v12)))]
            [scm:else (rt:fail)])))))])
        (scm:list (scm:cons (scm:string->symbol "mempty") Nothing) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
          semigroupMaybe11))))))

  (scm:define maybe$p
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Nothing? v22) (v0 Data.Unit.unit)]
            [(Just? v22) (v11 (Just-value0 v22))]
            [scm:else (rt:fail)])))))

  (scm:define maybe
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Nothing? v22) v0]
            [(Just? v22) (v11 (Just-value0 v22))]
            [scm:else (rt:fail)])))))

  (scm:define isNothing
    (scm:lambda (v20)
      (scm:cond
        [(Nothing? v20) #t]
        [(Just? v20) #f]
        [scm:else (rt:fail)])))

  (scm:define isJust
    (scm:lambda (v20)
      (scm:cond
        [(Nothing? v20) #f]
        [(Just? v20) #t]
        [scm:else (rt:fail)])))

  (scm:define genericMaybe
    (scm:list (scm:cons (scm:string->symbol "to") (scm:lambda (x0)
      (scm:cond
        [(Data.Generic.Rep.Inl? x0) Nothing]
        [(Data.Generic.Rep.Inr? x0) (Just (Data.Generic.Rep.Inr-value0 x0))]
        [scm:else (rt:fail)]))) (scm:cons (scm:string->symbol "from") (scm:lambda (x0)
      (scm:cond
        [(Nothing? x0) (Data.Generic.Rep.Inl Data.Generic.Rep.NoArguments)]
        [(Just? x0) (Data.Generic.Rep.Inr (Just-value0 x0))]
        [scm:else (rt:fail)])))))

  (scm:define functorMaybe
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Just? v11) (Just (v0 (Just-value0 v11)))]
          [scm:else Nothing]))))))

  (scm:define invariantMaybe
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [(Just? v12) (Just (f0 (Just-value0 v12)))]
            [scm:else Nothing])))))))

  (scm:define fromMaybe$p
    (scm:lambda (a0)
      ((maybe$p a0) identity)))

  (scm:define fromMaybe
    (scm:lambda (a0)
      (scm:lambda (v21)
        (scm:cond
          [(Nothing? v21) a0]
          [(Just? v21) (Just-value0 v21)]
          [scm:else (rt:fail)]))))

  (scm:define fromJust
    (scm:lambda (_)
      (scm:lambda (v1)
        (scm:cond
          [(Just? v1) (Just-value0 v1)]
          [scm:else (rt:fail)]))))

  (scm:define extendMaybe
    (scm:list (scm:cons (scm:string->symbol "extend") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Nothing? v11) Nothing]
          [scm:else (Just (v0 v11))])))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorMaybe))))

  (scm:define eqMaybe
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:cond
            [(Nothing? x1) (Nothing? y2)]
            [scm:else (scm:and (Just? x1) (scm:and (Just? y2) (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Just-value0 x1)) (Just-value0 y2))))])))))))

  (scm:define ordMaybe
    (scm:lambda (dictOrd0)
      (scm:let*
        ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))]
         [eqMaybe12 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:cond
              [(Nothing? x2) (Nothing? y3)]
              [scm:else (scm:and (Just? x2) (scm:and (Just? y3) (((rt:record-ref _1 (scm:string->symbol "eq")) (Just-value0 x2)) (Just-value0 y3))))])))))])
          (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x3)
            (scm:lambda (y4)
              (scm:cond
                [(Nothing? x3) (scm:cond
                  [(Nothing? y4) Data.Ordering.EQ]
                  [scm:else Data.Ordering.LT])]
                [(Nothing? y4) Data.Ordering.GT]
                [(scm:and (Just? x3) (Just? y4)) (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Just-value0 x3)) (Just-value0 y4))]
                [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
            eqMaybe12))))))

  (scm:define eq1Maybe
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:cond
            [(Nothing? x1) (Nothing? y2)]
            [scm:else (scm:and (Just? x1) (scm:and (Just? y2) (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Just-value0 x1)) (Just-value0 y2))))])))))))

  (scm:define ord1Maybe
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:cond
            [(Nothing? x1) (scm:cond
              [(Nothing? y2) Data.Ordering.EQ]
              [scm:else Data.Ordering.LT])]
            [(Nothing? y2) Data.Ordering.GT]
            [(scm:and (Just? x1) (Just? y2)) (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Just-value0 x1)) (Just-value0 y2))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1Maybe))))

  (scm:define boundedMaybe
    (scm:lambda (dictBounded0)
      (scm:let*
        ([_1 ((rt:record-ref dictBounded0 (scm:string->symbol "Ord0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Eq0")) (scm:quote undefined))]
         [ordMaybe13 (scm:let ([eqMaybe13 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x3)
          (scm:lambda (y4)
            (scm:cond
              [(Nothing? x3) (Nothing? y4)]
              [scm:else (scm:and (Just? x3) (scm:and (Just? y4) (((rt:record-ref _2 (scm:string->symbol "eq")) (Just-value0 x3)) (Just-value0 y4))))])))))])
          (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x4)
            (scm:lambda (y5)
              (scm:cond
                [(Nothing? x4) (scm:cond
                  [(Nothing? y5) Data.Ordering.EQ]
                  [scm:else Data.Ordering.LT])]
                [(Nothing? y5) Data.Ordering.GT]
                [(scm:and (Just? x4) (Just? y5)) (((rt:record-ref _1 (scm:string->symbol "compare")) (Just-value0 x4)) (Just-value0 y5))]
                [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
            eqMaybe13))))])
          (scm:list (scm:cons (scm:string->symbol "top") (Just (rt:record-ref dictBounded0 (scm:string->symbol "top")))) (scm:cons (scm:string->symbol "bottom") Nothing) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
            ordMaybe13))))))

  (scm:define applyMaybe
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Just? v0) (scm:cond
            [(Just? v11) (Just ((Just-value0 v0) (Just-value0 v11)))]
            [scm:else Nothing])]
          [(Nothing? v0) Nothing]
          [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorMaybe))))

  (scm:define bindMaybe
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Just? v0) (v11 (Just-value0 v0))]
          [(Nothing? v0) Nothing]
          [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyMaybe))))

  (scm:define semiringMaybe
    (scm:lambda (dictSemiring0)
      (scm:list (scm:cons (scm:string->symbol "zero") Nothing) (scm:cons (scm:string->symbol "one") (Just (rt:record-ref dictSemiring0 (scm:string->symbol "one")))) (scm:cons (scm:string->symbol "add") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [(Nothing? v1) v12]
            [(Nothing? v12) v1]
            [(scm:and (Just? v1) (Just? v12)) (Just (((rt:record-ref dictSemiring0 (scm:string->symbol "add")) (Just-value0 v1)) (Just-value0 v12)))]
            [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "mul") (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Just? x1) (Just? y2)) (Just (((rt:record-ref dictSemiring0 (scm:string->symbol "mul")) (Just-value0 x1)) (Just-value0 y2)))]
            [scm:else Nothing])))))))

  (scm:define applicativeMaybe
    (scm:list (scm:cons (scm:string->symbol "pure") Just) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyMaybe))))

  (scm:define monadMaybe
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeMaybe)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindMaybe))))

  (scm:define altMaybe
    (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Nothing? v0) v11]
          [scm:else v0])))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorMaybe))))

  (scm:define plusMaybe
    (scm:list (scm:cons (scm:string->symbol "empty") Nothing) (scm:cons (scm:string->symbol "Alt0") (scm:lambda (_)
      altMaybe))))

  (scm:define alternativeMaybe
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeMaybe)) (scm:cons (scm:string->symbol "Plus1") (scm:lambda (_)
      plusMaybe)))))
