#!r6rs
#!chezscheme
(library
  (Data.Field lib)
  (export
    field)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define field
    (scm:lambda (dictEuclideanRing0)
      (scm:lambda (dictDivisionRing1)
        (scm:list (scm:cons (scm:string->symbol "EuclideanRing0") (scm:lambda (_)
          dictEuclideanRing0)) (scm:cons (scm:string->symbol "DivisionRing1") (scm:lambda (_)
          dictDivisionRing1)))))))
