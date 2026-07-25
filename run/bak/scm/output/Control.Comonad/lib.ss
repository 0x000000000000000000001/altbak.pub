#!r6rs
#!chezscheme
(library
  (Control.Comonad lib)
  (export
    extract)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define extract
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "extract")))))
