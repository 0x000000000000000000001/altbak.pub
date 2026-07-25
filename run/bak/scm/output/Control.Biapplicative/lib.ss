#!r6rs
#!chezscheme
(library
  (Control.Biapplicative lib)
  (export
    biapplicativeTuple
    bipure)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Biapply lib) Control.Biapply.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define bipure
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bipure"))))

  (scm:define biapplicativeTuple
    (scm:list (scm:cons (scm:string->symbol "bipure") Data.Tuple.Tuple) (scm:cons (scm:string->symbol "Biapply0") (scm:lambda (_)
      Control.Biapply.biapplyTuple)))))
