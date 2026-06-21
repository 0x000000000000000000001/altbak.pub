#!r6rs
#!chezscheme
(library
  (Control.Semigroupoid lib)
  (export
    compose
    composeFlipped
    semigroupoidFn)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define semigroupoidFn
    (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (x2)
          (f0 (g1 x2))))))))

  (scm:define compose
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "compose"))))

  (scm:define composeFlipped
    (scm:lambda (dictSemigroupoid0)
      (scm:lambda (f1)
        (scm:lambda (g2)
          (((rt:record-ref dictSemigroupoid0 (scm:string->symbol "compose")) g2) f1))))))
