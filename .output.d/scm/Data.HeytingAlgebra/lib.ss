#!r6rs
#!chezscheme
(library
  (Data.HeytingAlgebra lib)
  (export
    boolConj
    boolDisj
    boolNot
    conj
    conjRecord
    disj
    disjRecord
    ff
    ffRecord
    heytingAlgebraBoolean
    heytingAlgebraFunction
    heytingAlgebraProxy
    heytingAlgebraRecord
    heytingAlgebraRecordCons
    heytingAlgebraRecordNil
    heytingAlgebraUnit
    implies
    impliesRecord
    not
    notRecord
    tt
    ttRecord)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.HeytingAlgebra foreign))

  (scm:define ttRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "ttRecord"))))

  (scm:define tt
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "tt"))))

  (scm:define notRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "notRecord"))))

  (scm:define not
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "not"))))

  (scm:define impliesRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "impliesRecord"))))

  (scm:define implies
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "implies"))))

  (scm:define heytingAlgebraUnit
    (scm:list (scm:cons (scm:string->symbol "ff") Data.Unit.unit) (scm:cons (scm:string->symbol "tt") Data.Unit.unit) (scm:cons (scm:string->symbol "implies") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Unit.unit))) (scm:cons (scm:string->symbol "conj") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Unit.unit))) (scm:cons (scm:string->symbol "disj") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Unit.unit))) (scm:cons (scm:string->symbol "not") (scm:lambda (v0)
      Data.Unit.unit))))

  (scm:define heytingAlgebraRecordNil
    (scm:list (scm:cons (scm:string->symbol "conjRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:list))))) (scm:cons (scm:string->symbol "disjRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:list))))) (scm:cons (scm:string->symbol "ffRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:list)))) (scm:cons (scm:string->symbol "impliesRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:list))))) (scm:cons (scm:string->symbol "notRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:list)))) (scm:cons (scm:string->symbol "ttRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:list))))))

  (scm:define heytingAlgebraProxy
    (scm:list (scm:cons (scm:string->symbol "conj") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy))) (scm:cons (scm:string->symbol "disj") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy))) (scm:cons (scm:string->symbol "implies") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy))) (scm:cons (scm:string->symbol "ff") Type.Proxy.Proxy) (scm:cons (scm:string->symbol "not") (scm:lambda (v0)
      Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "tt") Type.Proxy.Proxy)))

  (scm:define ffRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "ffRecord"))))

  (scm:define ff
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "ff"))))

  (scm:define disjRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "disjRecord"))))

  (scm:define disj
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "disj"))))

  (rt:define-lazy $lazy-heytingAlgebraBoolean "heytingAlgebraBoolean" "Data.HeytingAlgebra"
    (scm:list (scm:cons (scm:string->symbol "ff") #f) (scm:cons (scm:string->symbol "tt") #t) (scm:cons (scm:string->symbol "implies") (scm:lambda (a0)
      (scm:lambda (b1)
        (((rt:record-ref ($lazy-heytingAlgebraBoolean) (scm:string->symbol "disj")) ((rt:record-ref ($lazy-heytingAlgebraBoolean) (scm:string->symbol "not")) a0)) b1)))) (scm:cons (scm:string->symbol "conj") boolConj) (scm:cons (scm:string->symbol "disj") boolDisj) (scm:cons (scm:string->symbol "not") boolNot)))

  (scm:define heytingAlgebraBoolean
    ($lazy-heytingAlgebraBoolean))

  (scm:define conjRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "conjRecord"))))

  (scm:define heytingAlgebraRecord
    (scm:lambda (_)
      (scm:lambda (dictHeytingAlgebraRecord1)
        (scm:list (scm:cons (scm:string->symbol "ff") (((rt:record-ref dictHeytingAlgebraRecord1 (scm:string->symbol "ffRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "tt") (((rt:record-ref dictHeytingAlgebraRecord1 (scm:string->symbol "ttRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "conj") ((rt:record-ref dictHeytingAlgebraRecord1 (scm:string->symbol "conjRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "disj") ((rt:record-ref dictHeytingAlgebraRecord1 (scm:string->symbol "disjRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "implies") ((rt:record-ref dictHeytingAlgebraRecord1 (scm:string->symbol "impliesRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "not") ((rt:record-ref dictHeytingAlgebraRecord1 (scm:string->symbol "notRecord")) Type.Proxy.Proxy))))))

  (scm:define conj
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "conj"))))

  (scm:define heytingAlgebraFunction
    (scm:lambda (dictHeytingAlgebra0)
      (scm:let*
        ([ff11 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "ff"))]
         [tt12 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "tt"))])
          (scm:list (scm:cons (scm:string->symbol "ff") (scm:lambda (v3)
            ff11)) (scm:cons (scm:string->symbol "tt") (scm:lambda (v3)
            tt12)) (scm:cons (scm:string->symbol "implies") (scm:lambda (f3)
            (scm:lambda (g4)
              (scm:lambda (a5)
                (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "implies")) (f3 a5)) (g4 a5)))))) (scm:cons (scm:string->symbol "conj") (scm:lambda (f3)
            (scm:lambda (g4)
              (scm:lambda (a5)
                (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "conj")) (f3 a5)) (g4 a5)))))) (scm:cons (scm:string->symbol "disj") (scm:lambda (f3)
            (scm:lambda (g4)
              (scm:lambda (a5)
                (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "disj")) (f3 a5)) (g4 a5)))))) (scm:cons (scm:string->symbol "not") (scm:lambda (f3)
            (scm:lambda (a4)
              ((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "not")) (f3 a4)))))))))

  (scm:define heytingAlgebraRecordCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (_)
        (scm:lambda (dictHeytingAlgebraRecord2)
          (scm:lambda (dictHeytingAlgebra3)
            (scm:let*
              ([ff14 (rt:record-ref dictHeytingAlgebra3 (scm:string->symbol "ff"))]
               [tt15 (rt:record-ref dictHeytingAlgebra3 (scm:string->symbol "tt"))])
                (scm:list (scm:cons (scm:string->symbol "conjRecord") (scm:lambda (v6)
                  (scm:lambda (ra7)
                    (scm:lambda (rb8)
                      (scm:let*
                        ([key9 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                         [get10 (Record.Unsafe.unsafeGet key9)])
                          (((Record.Unsafe.unsafeSet key9) (((rt:record-ref dictHeytingAlgebra3 (scm:string->symbol "conj")) (get10 ra7)) (get10 rb8))) ((((rt:record-ref dictHeytingAlgebraRecord2 (scm:string->symbol "conjRecord")) Type.Proxy.Proxy) ra7) rb8))))))) (scm:cons (scm:string->symbol "disjRecord") (scm:lambda (v6)
                  (scm:lambda (ra7)
                    (scm:lambda (rb8)
                      (scm:let*
                        ([key9 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                         [get10 (Record.Unsafe.unsafeGet key9)])
                          (((Record.Unsafe.unsafeSet key9) (((rt:record-ref dictHeytingAlgebra3 (scm:string->symbol "disj")) (get10 ra7)) (get10 rb8))) ((((rt:record-ref dictHeytingAlgebraRecord2 (scm:string->symbol "disjRecord")) Type.Proxy.Proxy) ra7) rb8))))))) (scm:cons (scm:string->symbol "impliesRecord") (scm:lambda (v6)
                  (scm:lambda (ra7)
                    (scm:lambda (rb8)
                      (scm:let*
                        ([key9 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                         [get10 (Record.Unsafe.unsafeGet key9)])
                          (((Record.Unsafe.unsafeSet key9) (((rt:record-ref dictHeytingAlgebra3 (scm:string->symbol "implies")) (get10 ra7)) (get10 rb8))) ((((rt:record-ref dictHeytingAlgebraRecord2 (scm:string->symbol "impliesRecord")) Type.Proxy.Proxy) ra7) rb8))))))) (scm:cons (scm:string->symbol "ffRecord") (scm:lambda (v6)
                  (scm:lambda (row7)
                    (((Record.Unsafe.unsafeSet ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)) ff14) (((rt:record-ref dictHeytingAlgebraRecord2 (scm:string->symbol "ffRecord")) Type.Proxy.Proxy) row7))))) (scm:cons (scm:string->symbol "notRecord") (scm:lambda (v6)
                  (scm:lambda (row7)
                    (scm:let ([key8 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)])
                      (((Record.Unsafe.unsafeSet key8) ((rt:record-ref dictHeytingAlgebra3 (scm:string->symbol "not")) ((Record.Unsafe.unsafeGet key8) row7))) (((rt:record-ref dictHeytingAlgebraRecord2 (scm:string->symbol "notRecord")) Type.Proxy.Proxy) row7)))))) (scm:cons (scm:string->symbol "ttRecord") (scm:lambda (v6)
                  (scm:lambda (row7)
                    (((Record.Unsafe.unsafeSet ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)) tt15) (((rt:record-ref dictHeytingAlgebraRecord2 (scm:string->symbol "ttRecord")) Type.Proxy.Proxy) row7)))))))))))))
