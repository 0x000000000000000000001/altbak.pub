#!r6rs
#!chezscheme
(library
  (Data.Void lib)
  (export
    absurd)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define absurd
    (scm:lambda (a0)
      (scm:letrec ([spin1 (scm:lambda (v2)
        (spin1 v2))])
        (spin1 a0)))))
