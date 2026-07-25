#!r6rs
#!chezscheme
(library
  (Effect.Uncurried lib)
  (export
    mkEffectFn1
    mkEffectFn10
    mkEffectFn2
    mkEffectFn3
    mkEffectFn4
    mkEffectFn5
    mkEffectFn6
    mkEffectFn7
    mkEffectFn8
    mkEffectFn9
    monoidEffectFn1
    monoidEffectFn10
    monoidEffectFn2
    monoidEffectFn3
    monoidEffectFn4
    monoidEffectFn5
    monoidEffectFn6
    monoidEffectFn7
    monoidEffectFn8
    monoidEffectFn9
    runEffectFn1
    runEffectFn10
    runEffectFn2
    runEffectFn3
    runEffectFn4
    runEffectFn5
    runEffectFn6
    runEffectFn7
    runEffectFn8
    runEffectFn9
    semigroupEffectFn1
    semigroupEffectFn10
    semigroupEffectFn2
    semigroupEffectFn3
    semigroupEffectFn4
    semigroupEffectFn5
    semigroupEffectFn6
    semigroupEffectFn7
    semigroupEffectFn8
    semigroupEffectFn9)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Effect.Uncurried foreign))

  (scm:define semigroupEffectFn9
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4 c5 d6 e7 f8 g9 h10 i11)
            (scm:let*
              ([a$p12 (f11 a3 b4 c5 d6 e7 f8 g9 h10 i11)]
               [a$p13 (f22 a3 b4 c5 d6 e7 f8 g9 h10 i11)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p12) a$p13)))))))))

  (scm:define semigroupEffectFn8
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4 c5 d6 e7 f8 g9 h10)
            (scm:let*
              ([a$p11 (f11 a3 b4 c5 d6 e7 f8 g9 h10)]
               [a$p12 (f22 a3 b4 c5 d6 e7 f8 g9 h10)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p11) a$p12)))))))))

  (scm:define semigroupEffectFn7
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4 c5 d6 e7 f8 g9)
            (scm:let*
              ([a$p10 (f11 a3 b4 c5 d6 e7 f8 g9)]
               [a$p11 (f22 a3 b4 c5 d6 e7 f8 g9)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p10) a$p11)))))))))

  (scm:define semigroupEffectFn6
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4 c5 d6 e7 f8)
            (scm:let*
              ([a$p9 (f11 a3 b4 c5 d6 e7 f8)]
               [a$p10 (f22 a3 b4 c5 d6 e7 f8)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p9) a$p10)))))))))

  (scm:define semigroupEffectFn5
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4 c5 d6 e7)
            (scm:let*
              ([a$p8 (f11 a3 b4 c5 d6 e7)]
               [a$p9 (f22 a3 b4 c5 d6 e7)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p8) a$p9)))))))))

  (scm:define semigroupEffectFn4
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4 c5 d6)
            (scm:let*
              ([a$p7 (f11 a3 b4 c5 d6)]
               [a$p8 (f22 a3 b4 c5 d6)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p7) a$p8)))))))))

  (scm:define semigroupEffectFn3
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4 c5)
            (scm:let*
              ([a$p6 (f11 a3 b4 c5)]
               [a$p7 (f22 a3 b4 c5)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p6) a$p7)))))))))

  (scm:define semigroupEffectFn2
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4)
            (scm:let*
              ([a$p5 (f11 a3 b4)]
               [a$p6 (f22 a3 b4)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p5) a$p6)))))))))

  (scm:define semigroupEffectFn10
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3 b4 c5 d6 e7 f8 g9 h10 i11 j12)
            (scm:let*
              ([a$p13 (f11 a3 b4 c5 d6 e7 f8 g9 h10 i11 j12)]
               [a$p14 (f22 a3 b4 c5 d6 e7 f8 g9 h10 i11 j12)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p13) a$p14)))))))))

  (scm:define semigroupEffectFn1
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f11)
        (scm:lambda (f22)
          (scm:lambda (a3)
            (scm:let*
              ([a$p4 (f11 a3)]
               [a$p5 (f22 a3)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) a$p4) a$p5)))))))))

  (scm:define monoidEffectFn9
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn913 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6 c7 d8 e9 f10 g11 h12 i13)
              (scm:let*
                ([a$p14 (f13 a5 b6 c7 d8 e9 f10 g11 h12 i13)]
                 [a$p15 (f24 a5 b6 c7 d8 e9 f10 g11 h12 i13)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p14) a$p15)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15 v26 v37 v48 v59 v610 v711 v812)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn913))))))

  (scm:define monoidEffectFn8
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn813 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6 c7 d8 e9 f10 g11 h12)
              (scm:let*
                ([a$p13 (f13 a5 b6 c7 d8 e9 f10 g11 h12)]
                 [a$p14 (f24 a5 b6 c7 d8 e9 f10 g11 h12)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p13) a$p14)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15 v26 v37 v48 v59 v610 v711)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn813))))))

  (scm:define monoidEffectFn7
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn713 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6 c7 d8 e9 f10 g11)
              (scm:let*
                ([a$p12 (f13 a5 b6 c7 d8 e9 f10 g11)]
                 [a$p13 (f24 a5 b6 c7 d8 e9 f10 g11)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p12) a$p13)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15 v26 v37 v48 v59 v610)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn713))))))

  (scm:define monoidEffectFn6
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn613 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6 c7 d8 e9 f10)
              (scm:let*
                ([a$p11 (f13 a5 b6 c7 d8 e9 f10)]
                 [a$p12 (f24 a5 b6 c7 d8 e9 f10)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p11) a$p12)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15 v26 v37 v48 v59)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn613))))))

  (scm:define monoidEffectFn5
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn513 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6 c7 d8 e9)
              (scm:let*
                ([a$p10 (f13 a5 b6 c7 d8 e9)]
                 [a$p11 (f24 a5 b6 c7 d8 e9)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p10) a$p11)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15 v26 v37 v48)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn513))))))

  (scm:define monoidEffectFn4
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn413 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6 c7 d8)
              (scm:let*
                ([a$p9 (f13 a5 b6 c7 d8)]
                 [a$p10 (f24 a5 b6 c7 d8)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p9) a$p10)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15 v26 v37)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn413))))))

  (scm:define monoidEffectFn3
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn313 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6 c7)
              (scm:let*
                ([a$p8 (f13 a5 b6 c7)]
                 [a$p9 (f24 a5 b6 c7)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p8) a$p9)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15 v26)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn313))))))

  (scm:define monoidEffectFn2
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn213 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6)
              (scm:let*
                ([a$p7 (f13 a5 b6)]
                 [a$p8 (f24 a5 b6)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p7) a$p8)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn213))))))

  (scm:define monoidEffectFn10
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn1013 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5 b6 c7 d8 e9 f10 g11 h12 i13 j14)
              (scm:let*
                ([a$p15 (f13 a5 b6 c7 d8 e9 f10 g11 h12 i13 j14)]
                 [a$p16 (f24 a5 b6 c7 d8 e9 f10 g11 h12 i13 j14)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p15) a$p16)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4 v15 v26 v37 v48 v59 v610 v711 v812 v913)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn1013))))))

  (scm:define monoidEffectFn1
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupEffectFn113 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f13)
          (scm:lambda (f24)
            (scm:lambda (a5)
              (scm:let*
                ([a$p6 (f13 a5)]
                 [a$p7 (f24 a5)])
                  (((rt:record-ref _2 (scm:string->symbol "append")) a$p6) a$p7)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4)
            _1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupEffectFn113)))))))
