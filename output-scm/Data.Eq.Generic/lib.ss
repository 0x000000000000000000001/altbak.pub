#!r6rs
#!chezscheme
(library
  (Data.Eq.Generic lib)
  (export
    genericEq
    genericEq$p
    genericEqArgument
    genericEqConstructor
    genericEqNoArguments
    genericEqNoConstructors
    genericEqProduct
    genericEqSum)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.))

  (scm:define genericEqNoConstructors
    (scm:list (scm:cons (scm:string->symbol "genericEq'") (scm:lambda (v0)
      (scm:lambda (v11)
        #t)))))

  (scm:define genericEqNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericEq'") (scm:lambda (v0)
      (scm:lambda (v11)
        #t)))))

  (scm:define genericEqArgument
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "genericEq'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictEq0 (scm:string->symbol "eq")) v1) v12)))))))

  (scm:define genericEq$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericEq'"))))

  (scm:define genericEqConstructor
    (scm:lambda (dictGenericEq0)
      (scm:list (scm:cons (scm:string->symbol "genericEq'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericEq0 (scm:string->symbol "genericEq'")) v1) v12)))))))

  (scm:define genericEqProduct
    (scm:lambda (dictGenericEq0)
      (scm:lambda (dictGenericEq11)
        (scm:list (scm:cons (scm:string->symbol "genericEq'") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:and (((rt:record-ref dictGenericEq0 (scm:string->symbol "genericEq'")) (Data.Generic.Rep.Product-value0 v2)) (Data.Generic.Rep.Product-value0 v13)) (((rt:record-ref dictGenericEq11 (scm:string->symbol "genericEq'")) (Data.Generic.Rep.Product-value1 v2)) (Data.Generic.Rep.Product-value1 v13))))))))))

  (scm:define genericEqSum
    (scm:lambda (dictGenericEq0)
      (scm:lambda (dictGenericEq11)
        (scm:list (scm:cons (scm:string->symbol "genericEq'") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Generic.Rep.Inl? v2) (scm:and (Data.Generic.Rep.Inl? v13) (((rt:record-ref dictGenericEq0 (scm:string->symbol "genericEq'")) (Data.Generic.Rep.Inl-value0 v2)) (Data.Generic.Rep.Inl-value0 v13)))]
              [scm:else (scm:and (Data.Generic.Rep.Inr? v2) (scm:and (Data.Generic.Rep.Inr? v13) (((rt:record-ref dictGenericEq11 (scm:string->symbol "genericEq'")) (Data.Generic.Rep.Inr-value0 v2)) (Data.Generic.Rep.Inr-value0 v13))))]))))))))

  (scm:define genericEq
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericEq1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            (((rt:record-ref dictGenericEq1 (scm:string->symbol "genericEq'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3))))))))
