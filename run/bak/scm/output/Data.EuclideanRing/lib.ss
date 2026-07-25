#!r6rs
#!chezscheme
(library
  (Data.EuclideanRing lib)
  (export
    degree
    div
    euclideanRingInt
    euclideanRingNumber
    gcd
    intDegree
    intDiv
    intMod
    lcm
    mod
    numDiv)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.CommutativeRing lib) Data.CommutativeRing.)
    (Data.EuclideanRing foreign))

  (scm:define mod
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "mod"))))

  (scm:define gcd
    (scm:lambda (dictEq0)
      (scm:lambda (dictEuclideanRing1)
        (scm:let ([zero2 (rt:record-ref ((rt:record-ref ((rt:record-ref ((rt:record-ref dictEuclideanRing1 (scm:string->symbol "CommutativeRing0")) (scm:quote undefined)) (scm:string->symbol "Ring0")) (scm:quote undefined)) (scm:string->symbol "Semiring0")) (scm:quote undefined)) (scm:string->symbol "zero"))])
          (scm:lambda (a3)
            (scm:lambda (b4)
              (scm:cond
                [(((rt:record-ref dictEq0 (scm:string->symbol "eq")) b4) zero2) a3]
                [scm:else ((((gcd dictEq0) dictEuclideanRing1) b4) (((rt:record-ref dictEuclideanRing1 (scm:string->symbol "mod")) a3) b4))])))))))

  (scm:define euclideanRingNumber
    (scm:list (scm:cons (scm:string->symbol "degree") (scm:lambda (v0)
      1)) (scm:cons (scm:string->symbol "div") numDiv) (scm:cons (scm:string->symbol "mod") (scm:lambda (v0)
      (scm:lambda (v11)
        0.0))) (scm:cons (scm:string->symbol "CommutativeRing0") (scm:lambda (_)
      Data.CommutativeRing.commutativeRingNumber))))

  (scm:define euclideanRingInt
    (scm:list (scm:cons (scm:string->symbol "degree") intDegree) (scm:cons (scm:string->symbol "div") intDiv) (scm:cons (scm:string->symbol "mod") intMod) (scm:cons (scm:string->symbol "CommutativeRing0") (scm:lambda (_)
      Data.CommutativeRing.commutativeRingInt))))

  (scm:define div
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "div"))))

  (scm:define lcm
    (scm:lambda (dictEq0)
      (scm:lambda (dictEuclideanRing1)
        (scm:let*
          ([Semiring02 ((rt:record-ref ((rt:record-ref ((rt:record-ref dictEuclideanRing1 (scm:string->symbol "CommutativeRing0")) (scm:quote undefined)) (scm:string->symbol "Ring0")) (scm:quote undefined)) (scm:string->symbol "Semiring0")) (scm:quote undefined))]
           [zero3 (rt:record-ref Semiring02 (scm:string->symbol "zero"))]
           [gcd24 ((gcd dictEq0) dictEuclideanRing1)])
            (scm:lambda (a5)
              (scm:lambda (b6)
                (scm:cond
                  [(scm:or (((rt:record-ref dictEq0 (scm:string->symbol "eq")) a5) zero3) (((rt:record-ref dictEq0 (scm:string->symbol "eq")) b6) zero3)) zero3]
                  [scm:else (((rt:record-ref dictEuclideanRing1 (scm:string->symbol "div")) (((rt:record-ref Semiring02 (scm:string->symbol "mul")) a5) b6)) ((gcd24 a5) b6))])))))))

  (scm:define degree
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "degree")))))
