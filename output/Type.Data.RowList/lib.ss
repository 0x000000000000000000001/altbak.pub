#!r6rs
#!chezscheme
(library
  (Type.Data.RowList lib)
  (export
    RLProxy
    RLProxy?)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define RLProxy
    (scm:quote RLProxy))

  (scm:define RLProxy?
    (scm:lambda (v)
      (scm:eq? (scm:quote RLProxy) v))))
