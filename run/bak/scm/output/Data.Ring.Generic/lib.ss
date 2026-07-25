#!r6rs
#!chezscheme
(library
  (Data.Ring.Generic lib)
  (export
    genericRingArgument
    genericRingConstructor
    genericRingNoArguments
    genericRingProduct
    genericSub
    genericSub$p)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.))

  (scm:define genericSub$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericSub'"))))

  (scm:define genericSub
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericRing1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (((rt:record-ref dictGenericRing1 (scm:string->symbol "genericSub'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3))))))))

  (scm:define genericRingProduct
    (scm:lambda (dictGenericRing0)
      (scm:lambda (dictGenericRing11)
        (scm:list (scm:cons (scm:string->symbol "genericSub'") (scm:lambda (v2)
          (scm:lambda (v13)
            (Data.Generic.Rep.Product* (((rt:record-ref dictGenericRing0 (scm:string->symbol "genericSub'")) (Data.Generic.Rep.Product-value0 v2)) (Data.Generic.Rep.Product-value0 v13)) (((rt:record-ref dictGenericRing11 (scm:string->symbol "genericSub'")) (Data.Generic.Rep.Product-value1 v2)) (Data.Generic.Rep.Product-value1 v13))))))))))

  (scm:define genericRingNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericSub'") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Generic.Rep.NoArguments)))))

  (scm:define genericRingConstructor
    (scm:lambda (dictGenericRing0)
      (scm:list (scm:cons (scm:string->symbol "genericSub'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericRing0 (scm:string->symbol "genericSub'")) v1) v12)))))))

  (scm:define genericRingArgument
    (scm:lambda (dictRing0)
      (scm:list (scm:cons (scm:string->symbol "genericSub'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictRing0 (scm:string->symbol "sub")) v1) v12))))))))
