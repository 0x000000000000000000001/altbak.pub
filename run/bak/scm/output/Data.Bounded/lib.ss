#!r6rs
#!chezscheme
(library
  (Data.Bounded lib)
  (export
    bottom
    bottomChar
    bottomInt
    bottomNumber
    bottomRecord
    boundedBoolean
    boundedChar
    boundedInt
    boundedNumber
    boundedOrdering
    boundedProxy
    boundedRecord
    boundedRecordCons
    boundedRecordNil
    boundedUnit
    top
    topChar
    topInt
    topNumber
    topRecord)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ord lib) Data.Ord.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Bounded foreign))

  (scm:define topRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "topRecord"))))

  (scm:define top
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "top"))))

  (scm:define boundedUnit
    (scm:list (scm:cons (scm:string->symbol "top") Data.Unit.unit) (scm:cons (scm:string->symbol "bottom") Data.Unit.unit) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
      Data.Ord.ordUnit))))

  (scm:define boundedRecordNil
    (scm:list (scm:cons (scm:string->symbol "topRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:list)))) (scm:cons (scm:string->symbol "bottomRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:list)))) (scm:cons (scm:string->symbol "OrdRecord0") (scm:lambda (_)
      Data.Ord.ordRecordNil))))

  (scm:define boundedProxy
    (scm:list (scm:cons (scm:string->symbol "bottom") Type.Proxy.Proxy) (scm:cons (scm:string->symbol "top") Type.Proxy.Proxy) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
      Data.Ord.ordProxy))))

  (scm:define boundedOrdering
    (scm:list (scm:cons (scm:string->symbol "top") Data.Ordering.GT) (scm:cons (scm:string->symbol "bottom") Data.Ordering.LT) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
      Data.Ord.ordOrdering))))

  (scm:define boundedNumber
    (scm:list (scm:cons (scm:string->symbol "top") topNumber) (scm:cons (scm:string->symbol "bottom") bottomNumber) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
      Data.Ord.ordNumber))))

  (scm:define boundedInt
    (scm:list (scm:cons (scm:string->symbol "top") topInt) (scm:cons (scm:string->symbol "bottom") bottomInt) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
      Data.Ord.ordInt))))

  (scm:define boundedChar
    (scm:list (scm:cons (scm:string->symbol "top") topChar) (scm:cons (scm:string->symbol "bottom") bottomChar) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
      Data.Ord.ordChar))))

  (scm:define boundedBoolean
    (scm:list (scm:cons (scm:string->symbol "top") #t) (scm:cons (scm:string->symbol "bottom") #f) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
      Data.Ord.ordBoolean))))

  (scm:define bottomRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bottomRecord"))))

  (scm:define boundedRecord
    (scm:lambda (_)
      (scm:lambda (dictBoundedRecord1)
        (scm:let*
          ([_2 ((rt:record-ref dictBoundedRecord1 (scm:string->symbol "OrdRecord0")) (scm:quote undefined))]
           [eqRec13 (scm:list (scm:cons (scm:string->symbol "eq") ((rt:record-ref ((rt:record-ref _2 (scm:string->symbol "EqRecord0")) (scm:quote undefined)) (scm:string->symbol "eqRecord")) Type.Proxy.Proxy)))]
           [ordRecord14 (scm:list (scm:cons (scm:string->symbol "compare") ((rt:record-ref _2 (scm:string->symbol "compareRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
            eqRec13)))])
            (scm:list (scm:cons (scm:string->symbol "top") (((rt:record-ref dictBoundedRecord1 (scm:string->symbol "topRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "bottom") (((rt:record-ref dictBoundedRecord1 (scm:string->symbol "bottomRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
              ordRecord14)))))))

  (scm:define bottom
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bottom"))))

  (scm:define boundedRecordCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (dictBounded1)
        (scm:let*
          ([top12 (rt:record-ref dictBounded1 (scm:string->symbol "top"))]
           [bottom13 (rt:record-ref dictBounded1 (scm:string->symbol "bottom"))]
           [Ord04 ((rt:record-ref dictBounded1 (scm:string->symbol "Ord0")) (scm:quote undefined))])
            (scm:lambda (_)
              (scm:lambda (_)
                (scm:lambda (dictBoundedRecord7)
                  (scm:let*
                    ([_8 ((rt:record-ref dictBoundedRecord7 (scm:string->symbol "OrdRecord0")) (scm:quote undefined))]
                     [_9 ((rt:record-ref _8 (scm:string->symbol "EqRecord0")) (scm:quote undefined))]
                     [_10 ((rt:record-ref Ord04 (scm:string->symbol "Eq0")) (scm:quote undefined))]
                     [ordRecordCons11 (scm:let ([eqRowCons211 (scm:list (scm:cons (scm:string->symbol "eqRecord") (scm:lambda (v11)
                      (scm:lambda (ra12)
                        (scm:lambda (rb13)
                          (scm:let ([get14 (Record.Unsafe.unsafeGet ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy))])
                            (scm:and (((rt:record-ref _10 (scm:string->symbol "eq")) (get14 ra12)) (get14 rb13)) ((((rt:record-ref _9 (scm:string->symbol "eqRecord")) Type.Proxy.Proxy) ra12) rb13))))))))])
                      (scm:list (scm:cons (scm:string->symbol "compareRecord") (scm:lambda (v12)
                        (scm:lambda (ra13)
                          (scm:lambda (rb14)
                            (scm:let*
                              ([key15 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                               [left16 (((rt:record-ref Ord04 (scm:string->symbol "compare")) ((Record.Unsafe.unsafeGet key15) ra13)) ((Record.Unsafe.unsafeGet key15) rb14))])
                                (scm:cond
                                  [(scm:or (Data.Ordering.LT? left16) (scm:or (Data.Ordering.GT? left16) (scm:not (Data.Ordering.EQ? left16)))) left16]
                                  [scm:else ((((rt:record-ref _8 (scm:string->symbol "compareRecord")) Type.Proxy.Proxy) ra13) rb14)])))))) (scm:cons (scm:string->symbol "EqRecord0") (scm:lambda (_)
                        eqRowCons211))))])
                      (scm:list (scm:cons (scm:string->symbol "topRecord") (scm:lambda (v12)
                        (scm:lambda (rowProxy13)
                          (((Record.Unsafe.unsafeSet ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)) top12) (((rt:record-ref dictBoundedRecord7 (scm:string->symbol "topRecord")) Type.Proxy.Proxy) rowProxy13))))) (scm:cons (scm:string->symbol "bottomRecord") (scm:lambda (v12)
                        (scm:lambda (rowProxy13)
                          (((Record.Unsafe.unsafeSet ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)) bottom13) (((rt:record-ref dictBoundedRecord7 (scm:string->symbol "bottomRecord")) Type.Proxy.Proxy) rowProxy13))))) (scm:cons (scm:string->symbol "OrdRecord0") (scm:lambda (_)
                        ordRecordCons11))))))))))))
