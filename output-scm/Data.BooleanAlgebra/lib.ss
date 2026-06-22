#!r6rs
#!chezscheme
(library
  (Data.BooleanAlgebra lib)
  (export
    booleanAlgebraBoolean
    booleanAlgebraFn
    booleanAlgebraProxy
    booleanAlgebraRecord
    booleanAlgebraRecordCons
    booleanAlgebraRecordNil
    booleanAlgebraUnit)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.HeytingAlgebra lib) Data.HeytingAlgebra.)
    (prefix (Type.Proxy lib) Type.Proxy.))

  (scm:define booleanAlgebraUnit
    (scm:list (scm:cons (scm:string->symbol "HeytingAlgebra0") (scm:lambda (_)
      Data.HeytingAlgebra.heytingAlgebraUnit))))

  (scm:define booleanAlgebraRecordNil
    (scm:list (scm:cons (scm:string->symbol "HeytingAlgebraRecord0") (scm:lambda (_)
      Data.HeytingAlgebra.heytingAlgebraRecordNil))))

  (scm:define booleanAlgebraRecordCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (_)
        (scm:lambda (dictBooleanAlgebraRecord2)
          (scm:let ([heytingAlgebraRecordCons13 (((Data.HeytingAlgebra.heytingAlgebraRecordCons dictIsSymbol0) (scm:quote undefined)) ((rt:record-ref dictBooleanAlgebraRecord2 (scm:string->symbol "HeytingAlgebraRecord0")) (scm:quote undefined)))])
            (scm:lambda (dictBooleanAlgebra4)
              (scm:let ([heytingAlgebraRecordCons25 (heytingAlgebraRecordCons13 ((rt:record-ref dictBooleanAlgebra4 (scm:string->symbol "HeytingAlgebra0")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "HeytingAlgebraRecord0") (scm:lambda (_)
                  heytingAlgebraRecordCons25))))))))))

  (scm:define booleanAlgebraRecord
    (scm:lambda (_)
      (scm:lambda (dictBooleanAlgebraRecord1)
        (scm:let*
          ([_2 ((rt:record-ref dictBooleanAlgebraRecord1 (scm:string->symbol "HeytingAlgebraRecord0")) (scm:quote undefined))]
           [heytingAlgebraRecord13 (scm:list (scm:cons (scm:string->symbol "ff") (((rt:record-ref _2 (scm:string->symbol "ffRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "tt") (((rt:record-ref _2 (scm:string->symbol "ttRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "conj") ((rt:record-ref _2 (scm:string->symbol "conjRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "disj") ((rt:record-ref _2 (scm:string->symbol "disjRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "implies") ((rt:record-ref _2 (scm:string->symbol "impliesRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "not") ((rt:record-ref _2 (scm:string->symbol "notRecord")) Type.Proxy.Proxy)))])
            (scm:list (scm:cons (scm:string->symbol "HeytingAlgebra0") (scm:lambda (_)
              heytingAlgebraRecord13)))))))

  (scm:define booleanAlgebraProxy
    (scm:list (scm:cons (scm:string->symbol "HeytingAlgebra0") (scm:lambda (_)
      Data.HeytingAlgebra.heytingAlgebraProxy))))

  (scm:define booleanAlgebraFn
    (scm:lambda (dictBooleanAlgebra0)
      (scm:let*
        ([_1 ((rt:record-ref dictBooleanAlgebra0 (scm:string->symbol "HeytingAlgebra0")) (scm:quote undefined))]
         [ff12 (rt:record-ref _1 (scm:string->symbol "ff"))]
         [heytingAlgebraFunction3 (scm:let ([tt13 (rt:record-ref _1 (scm:string->symbol "tt"))])
          (scm:list (scm:cons (scm:string->symbol "ff") (scm:lambda (v4)
            ff12)) (scm:cons (scm:string->symbol "tt") (scm:lambda (v4)
            tt13)) (scm:cons (scm:string->symbol "implies") (scm:lambda (f4)
            (scm:lambda (g5)
              (scm:lambda (a6)
                (((rt:record-ref _1 (scm:string->symbol "implies")) (f4 a6)) (g5 a6)))))) (scm:cons (scm:string->symbol "conj") (scm:lambda (f4)
            (scm:lambda (g5)
              (scm:lambda (a6)
                (((rt:record-ref _1 (scm:string->symbol "conj")) (f4 a6)) (g5 a6)))))) (scm:cons (scm:string->symbol "disj") (scm:lambda (f4)
            (scm:lambda (g5)
              (scm:lambda (a6)
                (((rt:record-ref _1 (scm:string->symbol "disj")) (f4 a6)) (g5 a6)))))) (scm:cons (scm:string->symbol "not") (scm:lambda (f4)
            (scm:lambda (a5)
              ((rt:record-ref _1 (scm:string->symbol "not")) (f4 a5)))))))])
          (scm:list (scm:cons (scm:string->symbol "HeytingAlgebra0") (scm:lambda (_)
            heytingAlgebraFunction3))))))

  (scm:define booleanAlgebraBoolean
    (scm:list (scm:cons (scm:string->symbol "HeytingAlgebra0") (scm:lambda (_)
      Data.HeytingAlgebra.heytingAlgebraBoolean)))))
