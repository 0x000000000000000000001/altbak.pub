#!r6rs
#!chezscheme
(library
  (Data.Semiring.Generic lib)
  (export
    genericAdd
    genericAdd$p
    genericMul
    genericMul$p
    genericOne
    genericOne$p
    genericSemiringArgument
    genericSemiringConstructor
    genericSemiringNoArguments
    genericSemiringProduct
    genericZero
    genericZero$p)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.))

  (scm:define genericZero$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericZero'"))))

  (scm:define genericZero
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericSemiring1)
        ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (rt:record-ref dictGenericSemiring1 (scm:string->symbol "genericZero'"))))))

  (scm:define genericSemiringNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericAdd'") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Generic.Rep.NoArguments))) (scm:cons (scm:string->symbol "genericZero'") Data.Generic.Rep.NoArguments) (scm:cons (scm:string->symbol "genericMul'") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Generic.Rep.NoArguments))) (scm:cons (scm:string->symbol "genericOne'") Data.Generic.Rep.NoArguments)))

  (scm:define genericSemiringArgument
    (scm:lambda (dictSemiring0)
      (scm:list (scm:cons (scm:string->symbol "genericAdd'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemiring0 (scm:string->symbol "add")) v1) v12)))) (scm:cons (scm:string->symbol "genericZero'") (rt:record-ref dictSemiring0 (scm:string->symbol "zero"))) (scm:cons (scm:string->symbol "genericMul'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemiring0 (scm:string->symbol "mul")) v1) v12)))) (scm:cons (scm:string->symbol "genericOne'") (rt:record-ref dictSemiring0 (scm:string->symbol "one"))))))

  (scm:define genericOne$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericOne'"))))

  (scm:define genericOne
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericSemiring1)
        ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (rt:record-ref dictGenericSemiring1 (scm:string->symbol "genericOne'"))))))

  (scm:define genericMul$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericMul'"))))

  (scm:define genericMul
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericSemiring1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (((rt:record-ref dictGenericSemiring1 (scm:string->symbol "genericMul'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3))))))))

  (scm:define genericAdd$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericAdd'"))))

  (scm:define genericSemiringConstructor
    (scm:lambda (dictGenericSemiring0)
      (scm:list (scm:cons (scm:string->symbol "genericAdd'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericSemiring0 (scm:string->symbol "genericAdd'")) v1) v12)))) (scm:cons (scm:string->symbol "genericZero'") (rt:record-ref dictGenericSemiring0 (scm:string->symbol "genericZero'"))) (scm:cons (scm:string->symbol "genericMul'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericSemiring0 (scm:string->symbol "genericMul'")) v1) v12)))) (scm:cons (scm:string->symbol "genericOne'") (rt:record-ref dictGenericSemiring0 (scm:string->symbol "genericOne'"))))))

  (scm:define genericSemiringProduct
    (scm:lambda (dictGenericSemiring0)
      (scm:let*
        ([genericZero$p11 (rt:record-ref dictGenericSemiring0 (scm:string->symbol "genericZero'"))]
         [genericOne$p12 (rt:record-ref dictGenericSemiring0 (scm:string->symbol "genericOne'"))])
          (scm:lambda (dictGenericSemiring13)
            (scm:list (scm:cons (scm:string->symbol "genericAdd'") (scm:lambda (v4)
              (scm:lambda (v15)
                (Data.Generic.Rep.Product* (((rt:record-ref dictGenericSemiring0 (scm:string->symbol "genericAdd'")) (Data.Generic.Rep.Product-value0 v4)) (Data.Generic.Rep.Product-value0 v15)) (((rt:record-ref dictGenericSemiring13 (scm:string->symbol "genericAdd'")) (Data.Generic.Rep.Product-value1 v4)) (Data.Generic.Rep.Product-value1 v15)))))) (scm:cons (scm:string->symbol "genericZero'") (Data.Generic.Rep.Product* genericZero$p11 (rt:record-ref dictGenericSemiring13 (scm:string->symbol "genericZero'")))) (scm:cons (scm:string->symbol "genericMul'") (scm:lambda (v4)
              (scm:lambda (v15)
                (Data.Generic.Rep.Product* (((rt:record-ref dictGenericSemiring0 (scm:string->symbol "genericMul'")) (Data.Generic.Rep.Product-value0 v4)) (Data.Generic.Rep.Product-value0 v15)) (((rt:record-ref dictGenericSemiring13 (scm:string->symbol "genericMul'")) (Data.Generic.Rep.Product-value1 v4)) (Data.Generic.Rep.Product-value1 v15)))))) (scm:cons (scm:string->symbol "genericOne'") (Data.Generic.Rep.Product* genericOne$p12 (rt:record-ref dictGenericSemiring13 (scm:string->symbol "genericOne'")))))))))

  (scm:define genericAdd
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericSemiring1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (((rt:record-ref dictGenericSemiring1 (scm:string->symbol "genericAdd'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3)))))))))
