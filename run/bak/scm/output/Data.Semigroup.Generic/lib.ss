#!r6rs
#!chezscheme
(library
  (Data.Semigroup.Generic lib)
  (export
    genericAppend
    genericAppend$p
    genericSemigroupArgument
    genericSemigroupConstructor
    genericSemigroupNoArguments
    genericSemigroupNoConstructors
    genericSemigroupProduct)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.))

  (scm:define genericSemigroupNoConstructors
    (scm:list (scm:cons (scm:string->symbol "genericAppend'") (scm:lambda (a0)
      (scm:lambda (v1)
        a0)))))

  (scm:define genericSemigroupNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericAppend'") (scm:lambda (a0)
      (scm:lambda (v1)
        a0)))))

  (scm:define genericSemigroupArgument
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "genericAppend'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) v1) v12)))))))

  (scm:define genericAppend$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericAppend'"))))

  (scm:define genericSemigroupConstructor
    (scm:lambda (dictGenericSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "genericAppend'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericSemigroup0 (scm:string->symbol "genericAppend'")) v1) v12)))))))

  (scm:define genericSemigroupProduct
    (scm:lambda (dictGenericSemigroup0)
      (scm:lambda (dictGenericSemigroup11)
        (scm:list (scm:cons (scm:string->symbol "genericAppend'") (scm:lambda (v2)
          (scm:lambda (v13)
            (Data.Generic.Rep.Product* (((rt:record-ref dictGenericSemigroup0 (scm:string->symbol "genericAppend'")) (Data.Generic.Rep.Product-value0 v2)) (Data.Generic.Rep.Product-value0 v13)) (((rt:record-ref dictGenericSemigroup11 (scm:string->symbol "genericAppend'")) (Data.Generic.Rep.Product-value1 v2)) (Data.Generic.Rep.Product-value1 v13))))))))))

  (scm:define genericAppend
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericSemigroup1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (((rt:record-ref dictGenericSemigroup1 (scm:string->symbol "genericAppend'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3)))))))))
