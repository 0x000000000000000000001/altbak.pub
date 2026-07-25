#!r6rs
#!chezscheme
(library
  (Record.Unsafe lib)
  (export
    unsafeDelete
    unsafeGet
    unsafeHas
    unsafeSet)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Record.Unsafe foreign))
  )
