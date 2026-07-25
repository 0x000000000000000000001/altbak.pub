#!r6rs
#!chezscheme
(library
  (Data.DivisionRing lib)
  (export
    divisionringNumber
    leftDiv
    recip
    rightDiv)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ring lib) Data.Ring.))

  (scm:define recip
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "recip"))))

  (scm:define rightDiv
    (scm:lambda (dictDivisionRing0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictDivisionRing0 (scm:string->symbol "Ring0")) (scm:quote undefined)) (scm:string->symbol "Semiring0")) (scm:quote undefined)) (scm:string->symbol "mul")) a1) ((rt:record-ref dictDivisionRing0 (scm:string->symbol "recip")) b2))))))

  (scm:define leftDiv
    (scm:lambda (dictDivisionRing0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictDivisionRing0 (scm:string->symbol "Ring0")) (scm:quote undefined)) (scm:string->symbol "Semiring0")) (scm:quote undefined)) (scm:string->symbol "mul")) ((rt:record-ref dictDivisionRing0 (scm:string->symbol "recip")) b2)) a1)))))

  (scm:define divisionringNumber
    (scm:list (scm:cons (scm:string->symbol "recip") (scm:lambda (x0)
      (scm:fl/ 1.0 x0))) (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
      Data.Ring.ringNumber)))))
