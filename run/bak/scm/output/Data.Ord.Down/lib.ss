#!r6rs
#!chezscheme
(library
  (Data.Ord.Down lib)
  (export
    Down
    boundedDown
    eqDown
    newtypeDown
    ordDown
    showDown)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define Down
    (scm:lambda (x0)
      x0))

  (scm:define showDown
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Down ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define newtypeDown
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define eqDown
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define ordDown
    (scm:lambda (dictOrd0)
      (scm:let ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))])
        (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:let ([_4 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) v2) v13)])
              (scm:cond
                [(Data.Ordering.GT? _4) Data.Ordering.LT]
                [(Data.Ordering.EQ? _4) Data.Ordering.EQ]
                [(Data.Ordering.LT? _4) Data.Ordering.GT]
                [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
          _1))))))

  (scm:define boundedDown
    (scm:lambda (dictBounded0)
      (scm:let*
        ([_1 ((rt:record-ref dictBounded0 (scm:string->symbol "Ord0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Eq0")) (scm:quote undefined))]
         [ordDown13 (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (v3)
          (scm:lambda (v14)
            (scm:let ([_5 (((rt:record-ref _1 (scm:string->symbol "compare")) v3) v14)])
              (scm:cond
                [(Data.Ordering.GT? _5) Data.Ordering.LT]
                [(Data.Ordering.EQ? _5) Data.Ordering.EQ]
                [(Data.Ordering.LT? _5) Data.Ordering.GT]
                [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
          _2)))])
          (scm:list (scm:cons (scm:string->symbol "top") (rt:record-ref dictBounded0 (scm:string->symbol "bottom"))) (scm:cons (scm:string->symbol "bottom") (rt:record-ref dictBounded0 (scm:string->symbol "top"))) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
            ordDown13)))))))
