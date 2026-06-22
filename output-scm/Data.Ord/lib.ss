#!r6rs
#!chezscheme
(library
  (Data.Ord lib)
  (export
    abs
    between
    clamp
    compare
    compare1
    compareRecord
    comparing
    greaterThan
    greaterThanOrEq
    lessThan
    lessThanOrEq
    max
    min
    ord1Array
    ordArray
    ordArrayImpl
    ordBoolean
    ordBooleanImpl
    ordChar
    ordCharImpl
    ordInt
    ordIntImpl
    ordNumber
    ordNumberImpl
    ordOrdering
    ordProxy
    ordRecord
    ordRecordCons
    ordRecordNil
    ordString
    ordStringImpl
    ordUnit
    ordVoid
    signum)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Eq lib) Data.Eq.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Ord foreign))

  (scm:define ordVoid
    (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Ordering.EQ))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Eq.eqVoid))))

  (scm:define ordUnit
    (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Ordering.EQ))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Eq.eqUnit))))

  (scm:define ordString
    (scm:list (scm:cons (scm:string->symbol "compare") (((ordStringImpl Data.Ordering.LT) Data.Ordering.EQ) Data.Ordering.GT)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Eq.eqString))))

  (scm:define ordRecordNil
    (scm:list (scm:cons (scm:string->symbol "compareRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          Data.Ordering.EQ)))) (scm:cons (scm:string->symbol "EqRecord0") (scm:lambda (_)
      Data.Eq.eqRowNil))))

  (scm:define ordProxy
    (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Ordering.EQ))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Eq.eqProxy))))

  (scm:define ordOrdering
    (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Data.Ordering.LT? v0) (scm:cond
            [(Data.Ordering.LT? v11) Data.Ordering.EQ]
            [scm:else Data.Ordering.LT])]
          [(Data.Ordering.EQ? v0) (scm:cond
            [(Data.Ordering.EQ? v11) Data.Ordering.EQ]
            [(Data.Ordering.LT? v11) Data.Ordering.GT]
            [(Data.Ordering.GT? v11) Data.Ordering.LT]
            [scm:else (rt:fail)])]
          [(Data.Ordering.GT? v0) (scm:cond
            [(Data.Ordering.GT? v11) Data.Ordering.EQ]
            [scm:else Data.Ordering.GT])]
          [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Ordering.eqOrdering))))

  (scm:define ordNumber
    (scm:list (scm:cons (scm:string->symbol "compare") (((ordNumberImpl Data.Ordering.LT) Data.Ordering.EQ) Data.Ordering.GT)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Eq.eqNumber))))

  (scm:define ordInt
    (scm:list (scm:cons (scm:string->symbol "compare") (((ordIntImpl Data.Ordering.LT) Data.Ordering.EQ) Data.Ordering.GT)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Eq.eqInt))))

  (scm:define ordChar
    (scm:list (scm:cons (scm:string->symbol "compare") (((ordCharImpl Data.Ordering.LT) Data.Ordering.EQ) Data.Ordering.GT)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Eq.eqChar))))

  (scm:define ordBoolean
    (scm:list (scm:cons (scm:string->symbol "compare") (((ordBooleanImpl Data.Ordering.LT) Data.Ordering.EQ) Data.Ordering.GT)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
      Data.Eq.eqBoolean))))

  (scm:define compareRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "compareRecord"))))

  (scm:define ordRecord
    (scm:lambda (_)
      (scm:lambda (dictOrdRecord1)
        (scm:let ([eqRec12 (scm:list (scm:cons (scm:string->symbol "eq") ((rt:record-ref ((rt:record-ref dictOrdRecord1 (scm:string->symbol "EqRecord0")) (scm:quote undefined)) (scm:string->symbol "eqRecord")) Type.Proxy.Proxy)))])
          (scm:list (scm:cons (scm:string->symbol "compare") ((rt:record-ref dictOrdRecord1 (scm:string->symbol "compareRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
            eqRec12)))))))

  (scm:define compare1
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "compare1"))))

  (scm:define compare
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "compare"))))

  (scm:define comparing
    (scm:lambda (dictOrd0)
      (scm:lambda (f1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (f1 x2)) (f1 y3)))))))

  (scm:define greaterThan
    (scm:lambda (dictOrd0)
      (scm:lambda (a11)
        (scm:lambda (a22)
          (Data.Ordering.GT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) a11) a22))))))

  (scm:define greaterThanOrEq
    (scm:lambda (dictOrd0)
      (scm:lambda (a11)
        (scm:lambda (a22)
          (scm:not (Data.Ordering.LT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) a11) a22)))))))

  (scm:define lessThan
    (scm:lambda (dictOrd0)
      (scm:lambda (a11)
        (scm:lambda (a22)
          (Data.Ordering.LT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) a11) a22))))))

  (scm:define signum
    (scm:lambda (dictOrd0)
      (scm:lambda (dictRing1)
        (scm:let*
          ([Semiring02 ((rt:record-ref dictRing1 (scm:string->symbol "Semiring0")) (scm:quote undefined))]
           [zero3 (rt:record-ref Semiring02 (scm:string->symbol "zero"))]
           [zero4 (rt:record-ref ((rt:record-ref dictRing1 (scm:string->symbol "Semiring0")) (scm:quote undefined)) (scm:string->symbol "zero"))]
           [one5 (rt:record-ref Semiring02 (scm:string->symbol "one"))])
            (scm:lambda (x6)
              (scm:cond
                [(Data.Ordering.LT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) x6) zero3)) (((rt:record-ref dictRing1 (scm:string->symbol "sub")) zero4) one5)]
                [(Data.Ordering.GT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) x6) zero3)) one5]
                [scm:else x6]))))))

  (scm:define lessThanOrEq
    (scm:lambda (dictOrd0)
      (scm:lambda (a11)
        (scm:lambda (a22)
          (scm:not (Data.Ordering.GT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) a11) a22)))))))

  (scm:define max
    (scm:lambda (dictOrd0)
      (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:let ([v3 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) x1) y2)])
            (scm:cond
              [(Data.Ordering.LT? v3) y2]
              [(Data.Ordering.EQ? v3) x1]
              [(Data.Ordering.GT? v3) x1]
              [scm:else (rt:fail)]))))))

  (scm:define min
    (scm:lambda (dictOrd0)
      (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:let ([v3 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) x1) y2)])
            (scm:cond
              [(Data.Ordering.LT? v3) x1]
              [(Data.Ordering.EQ? v3) x1]
              [(Data.Ordering.GT? v3) y2]
              [scm:else (rt:fail)]))))))

  (scm:define ordArray
    (scm:lambda (dictOrd0)
      (scm:let ([eqArray1 (scm:list (scm:cons (scm:string->symbol "eq") (Data.Eq.eqArrayImpl (rt:record-ref ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined)) (scm:string->symbol "eq")))))])
        (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (xs2)
          (scm:lambda (ys3)
            (((rt:record-ref ordInt (scm:string->symbol "compare")) 0) (((ordArrayImpl (scm:lambda (x4)
              (scm:lambda (y5)
                (scm:let ([v6 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) x4) y5)])
                  (scm:cond
                    [(Data.Ordering.EQ? v6) 0]
                    [(Data.Ordering.LT? v6) 1]
                    [(Data.Ordering.GT? v6) -1]
                    [scm:else (rt:fail)]))))) xs2) ys3))))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
          eqArray1))))))

  (scm:define ord1Array
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref (ordArray dictOrd0) (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      Data.Eq.eq1Array))))

  (scm:define ordRecordCons
    (scm:lambda (dictOrdRecord0)
      (scm:let ([eqRowCons1 ((Data.Eq.eqRowCons ((rt:record-ref dictOrdRecord0 (scm:string->symbol "EqRecord0")) (scm:quote undefined))) (scm:quote undefined))])
        (scm:lambda (_)
          (scm:lambda (dictIsSymbol3)
            (scm:let ([eqRowCons14 (eqRowCons1 dictIsSymbol3)])
              (scm:lambda (dictOrd5)
                (scm:let ([eqRowCons26 (eqRowCons14 ((rt:record-ref dictOrd5 (scm:string->symbol "Eq0")) (scm:quote undefined)))])
                  (scm:list (scm:cons (scm:string->symbol "compareRecord") (scm:lambda (v7)
                    (scm:lambda (ra8)
                      (scm:lambda (rb9)
                        (scm:let*
                          ([key10 ((rt:record-ref dictIsSymbol3 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                           [left11 (((rt:record-ref dictOrd5 (scm:string->symbol "compare")) ((Record.Unsafe.unsafeGet key10) ra8)) ((Record.Unsafe.unsafeGet key10) rb9))])
                            (scm:cond
                              [(scm:or (Data.Ordering.LT? left11) (scm:or (Data.Ordering.GT? left11) (scm:not (Data.Ordering.EQ? left11)))) left11]
                              [scm:else ((((rt:record-ref dictOrdRecord0 (scm:string->symbol "compareRecord")) Type.Proxy.Proxy) ra8) rb9)])))))) (scm:cons (scm:string->symbol "EqRecord0") (scm:lambda (_)
                    eqRowCons26)))))))))))

  (scm:define clamp
    (scm:lambda (dictOrd0)
      (scm:lambda (low1)
        (scm:lambda (hi2)
          (scm:lambda (x3)
            (scm:let*
              ([v4 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) low1) x3)]
               [_5 (scm:cond
                [(Data.Ordering.LT? v4) x3]
                [(Data.Ordering.EQ? v4) low1]
                [(Data.Ordering.GT? v4) low1]
                [scm:else (rt:fail)])]
               [v6 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) hi2) _5)])
                (scm:cond
                  [(Data.Ordering.LT? v6) hi2]
                  [(Data.Ordering.EQ? v6) hi2]
                  [(Data.Ordering.GT? v6) _5]
                  [scm:else (rt:fail)])))))))

  (scm:define between
    (scm:lambda (dictOrd0)
      (scm:lambda (low1)
        (scm:lambda (hi2)
          (scm:lambda (x3)
            (scm:cond
              [(Data.Ordering.LT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) x3) low1)) #f]
              [scm:else (scm:not (Data.Ordering.GT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) x3) hi2)))]))))))

  (scm:define abs
    (scm:lambda (dictOrd0)
      (scm:lambda (dictRing1)
        (scm:let*
          ([zero2 (rt:record-ref ((rt:record-ref dictRing1 (scm:string->symbol "Semiring0")) (scm:quote undefined)) (scm:string->symbol "zero"))]
           [zero3 (rt:record-ref ((rt:record-ref dictRing1 (scm:string->symbol "Semiring0")) (scm:quote undefined)) (scm:string->symbol "zero"))])
            (scm:lambda (x4)
              (scm:cond
                [(scm:not (Data.Ordering.LT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) x4) zero2))) x4]
                [scm:else (((rt:record-ref dictRing1 (scm:string->symbol "sub")) zero3) x4)])))))))
