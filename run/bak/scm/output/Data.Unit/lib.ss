#!r6rs
#!chezscheme
(library
  (Data.Unit lib)
  (export
    unit)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Data.Unit foreign))
  )
