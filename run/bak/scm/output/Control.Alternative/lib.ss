#!r6rs
#!chezscheme
(library
  (Control.Alternative lib)
  (export
    alternativeArray
    guard)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Applicative lib) Control.Applicative.)
    (prefix (Control.Plus lib) Control.Plus.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define guard
    (scm:lambda (dictAlternative0)
      (scm:let ([empty1 (rt:record-ref ((rt:record-ref dictAlternative0 (scm:string->symbol "Plus1")) (scm:quote undefined)) (scm:string->symbol "empty"))])
        (scm:lambda (v2)
          (scm:cond
            [v2 ((rt:record-ref ((rt:record-ref dictAlternative0 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure")) Data.Unit.unit)]
            [scm:else empty1])))))

  (scm:define alternativeArray
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      Control.Applicative.applicativeArray)) (scm:cons (scm:string->symbol "Plus1") (scm:lambda (_)
      Control.Plus.plusArray)))))
