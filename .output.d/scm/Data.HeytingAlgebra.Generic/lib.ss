#!r6rs
#!chezscheme
(library
  (Data.HeytingAlgebra.Generic lib)
  (export
    genericConj
    genericConj$p
    genericDisj
    genericDisj$p
    genericFF
    genericFF$p
    genericHeytingAlgebraArgument
    genericHeytingAlgebraConstructor
    genericHeytingAlgebraNoArguments
    genericHeytingAlgebraProduct
    genericImplies
    genericImplies$p
    genericNot
    genericNot$p
    genericTT
    genericTT$p)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.))

  (scm:define genericTT$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericTT'"))))

  (scm:define genericTT
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericHeytingAlgebra1)
        ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (rt:record-ref dictGenericHeytingAlgebra1 (scm:string->symbol "genericTT'"))))))

  (scm:define genericNot$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericNot'"))))

  (scm:define genericNot
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericHeytingAlgebra1)
        (scm:lambda (x2)
          ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) ((rt:record-ref dictGenericHeytingAlgebra1 (scm:string->symbol "genericNot'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)))))))

  (scm:define genericImplies$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericImplies'"))))

  (scm:define genericImplies
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericHeytingAlgebra1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (((rt:record-ref dictGenericHeytingAlgebra1 (scm:string->symbol "genericImplies'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3))))))))

  (scm:define genericHeytingAlgebraNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericFF'") Data.Generic.Rep.NoArguments) (scm:cons (scm:string->symbol "genericTT'") Data.Generic.Rep.NoArguments) (scm:cons (scm:string->symbol "genericImplies'") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Generic.Rep.NoArguments))) (scm:cons (scm:string->symbol "genericConj'") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Generic.Rep.NoArguments))) (scm:cons (scm:string->symbol "genericDisj'") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Generic.Rep.NoArguments))) (scm:cons (scm:string->symbol "genericNot'") (scm:lambda (v0)
      Data.Generic.Rep.NoArguments))))

  (scm:define genericHeytingAlgebraArgument
    (scm:lambda (dictHeytingAlgebra0)
      (scm:list (scm:cons (scm:string->symbol "genericFF'") (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "ff"))) (scm:cons (scm:string->symbol "genericTT'") (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "tt"))) (scm:cons (scm:string->symbol "genericImplies'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "implies")) v1) v12)))) (scm:cons (scm:string->symbol "genericConj'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "conj")) v1) v12)))) (scm:cons (scm:string->symbol "genericDisj'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "disj")) v1) v12)))) (scm:cons (scm:string->symbol "genericNot'") (scm:lambda (v1)
        ((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "not")) v1))))))

  (scm:define genericFF$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericFF'"))))

  (scm:define genericFF
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericHeytingAlgebra1)
        ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (rt:record-ref dictGenericHeytingAlgebra1 (scm:string->symbol "genericFF'"))))))

  (scm:define genericDisj$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericDisj'"))))

  (scm:define genericDisj
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericHeytingAlgebra1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (((rt:record-ref dictGenericHeytingAlgebra1 (scm:string->symbol "genericDisj'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3))))))))

  (scm:define genericConj$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericConj'"))))

  (scm:define genericHeytingAlgebraConstructor
    (scm:lambda (dictGenericHeytingAlgebra0)
      (scm:list (scm:cons (scm:string->symbol "genericFF'") (rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericFF'"))) (scm:cons (scm:string->symbol "genericTT'") (rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericTT'"))) (scm:cons (scm:string->symbol "genericImplies'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericImplies'")) v1) v12)))) (scm:cons (scm:string->symbol "genericConj'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericConj'")) v1) v12)))) (scm:cons (scm:string->symbol "genericDisj'") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericDisj'")) v1) v12)))) (scm:cons (scm:string->symbol "genericNot'") (scm:lambda (v1)
        ((rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericNot'")) v1))))))

  (scm:define genericHeytingAlgebraProduct
    (scm:lambda (dictGenericHeytingAlgebra0)
      (scm:let*
        ([genericFF$p11 (rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericFF'"))]
         [genericTT$p12 (rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericTT'"))])
          (scm:lambda (dictGenericHeytingAlgebra13)
            (scm:list (scm:cons (scm:string->symbol "genericFF'") (Data.Generic.Rep.Product* genericFF$p11 (rt:record-ref dictGenericHeytingAlgebra13 (scm:string->symbol "genericFF'")))) (scm:cons (scm:string->symbol "genericTT'") (Data.Generic.Rep.Product* genericTT$p12 (rt:record-ref dictGenericHeytingAlgebra13 (scm:string->symbol "genericTT'")))) (scm:cons (scm:string->symbol "genericImplies'") (scm:lambda (v4)
              (scm:lambda (v15)
                (Data.Generic.Rep.Product* (((rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericImplies'")) (Data.Generic.Rep.Product-value0 v4)) (Data.Generic.Rep.Product-value0 v15)) (((rt:record-ref dictGenericHeytingAlgebra13 (scm:string->symbol "genericImplies'")) (Data.Generic.Rep.Product-value1 v4)) (Data.Generic.Rep.Product-value1 v15)))))) (scm:cons (scm:string->symbol "genericConj'") (scm:lambda (v4)
              (scm:lambda (v15)
                (Data.Generic.Rep.Product* (((rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericConj'")) (Data.Generic.Rep.Product-value0 v4)) (Data.Generic.Rep.Product-value0 v15)) (((rt:record-ref dictGenericHeytingAlgebra13 (scm:string->symbol "genericConj'")) (Data.Generic.Rep.Product-value1 v4)) (Data.Generic.Rep.Product-value1 v15)))))) (scm:cons (scm:string->symbol "genericDisj'") (scm:lambda (v4)
              (scm:lambda (v15)
                (Data.Generic.Rep.Product* (((rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericDisj'")) (Data.Generic.Rep.Product-value0 v4)) (Data.Generic.Rep.Product-value0 v15)) (((rt:record-ref dictGenericHeytingAlgebra13 (scm:string->symbol "genericDisj'")) (Data.Generic.Rep.Product-value1 v4)) (Data.Generic.Rep.Product-value1 v15)))))) (scm:cons (scm:string->symbol "genericNot'") (scm:lambda (v4)
              (Data.Generic.Rep.Product* ((rt:record-ref dictGenericHeytingAlgebra0 (scm:string->symbol "genericNot'")) (Data.Generic.Rep.Product-value0 v4)) ((rt:record-ref dictGenericHeytingAlgebra13 (scm:string->symbol "genericNot'")) (Data.Generic.Rep.Product-value1 v4))))))))))

  (scm:define genericConj
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericHeytingAlgebra1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (((rt:record-ref dictGenericHeytingAlgebra1 (scm:string->symbol "genericConj'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) y3)))))))))
