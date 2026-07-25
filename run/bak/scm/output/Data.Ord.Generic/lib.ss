#!r6rs
#!chezscheme
(library
  (Data.Ord.Generic lib)
  (export
    genericCompare
    genericCompare$p
    genericOrdArgument
    genericOrdConstructor
    genericOrdNoArguments
    genericOrdNoConstructors
    genericOrdProduct
    genericOrdSum)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define genericOrdNoConstructors
    (scm:list (scm:cons (scm:string->symbol "genericCompare'") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Ordering.EQ)))))

  (scm:define genericOrdNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericCompare'") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Ordering.EQ)))))

  (scm:define genericOrdArgument
    (scm:lambda (dictOrd0)
      (scm:list (scm:cons (scm:string->symbol "genericCompare'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) v1) v12)))))))

  (scm:define genericCompare$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericCompare'"))))

  (scm:define genericOrdConstructor
    (scm:lambda (dictGenericOrd0)
      (scm:list (scm:cons (scm:string->symbol "genericCompare'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericOrd0 (scm:string->symbol "genericCompare'")) v1) v12)))))))

  (scm:define genericOrdProduct
    (scm:lambda (dictGenericOrd0)
      (scm:lambda (dictGenericOrd11)
        (scm:list (scm:cons (scm:string->symbol "genericCompare'") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:let ([v24 (((rt:record-ref dictGenericOrd0 (scm:string->symbol "genericCompare'")) (Data.Generic.Rep.Product-value0 v2)) (Data.Generic.Rep.Product-value0 v13))])
              (scm:cond
                [(Data.Ordering.EQ? v24) (((rt:record-ref dictGenericOrd11 (scm:string->symbol "genericCompare'")) (Data.Generic.Rep.Product-value1 v2)) (Data.Generic.Rep.Product-value1 v13))]
                [scm:else v24])))))))))

  (scm:define genericOrdSum
    (scm:lambda (dictGenericOrd0)
      (scm:lambda (dictGenericOrd11)
        (scm:list (scm:cons (scm:string->symbol "genericCompare'") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Generic.Rep.Inl? v2) (scm:cond
                [(Data.Generic.Rep.Inl? v13) (((rt:record-ref dictGenericOrd0 (scm:string->symbol "genericCompare'")) (Data.Generic.Rep.Inl-value0 v2)) (Data.Generic.Rep.Inl-value0 v13))]
                [(Data.Generic.Rep.Inr? v13) Data.Ordering.LT]
                [scm:else (rt:fail)])]
              [(Data.Generic.Rep.Inr? v2) (scm:cond
                [(Data.Generic.Rep.Inr? v13) (((rt:record-ref dictGenericOrd11 (scm:string->symbol "genericCompare'")) (Data.Generic.Rep.Inr-value0 v2)) (Data.Generic.Rep.Inr-value0 v13))]
                [(Data.Generic.Rep.Inl? v13) Data.Ordering.GT]
                [scm:else (rt:fail)])]
              [scm:else (rt:fail)]))))))))

  (scm:define genericCompare
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericOrd1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            (((rt:record-ref dictGenericOrd1 (scm:string->symbol "genericCompare'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3))))))))
