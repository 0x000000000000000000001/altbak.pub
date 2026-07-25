#!r6rs
#!chezscheme
(library
  (Unsafe.Coerce lib)
  (export
    unsafeCoerce)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Unsafe.Coerce foreign))
  )
