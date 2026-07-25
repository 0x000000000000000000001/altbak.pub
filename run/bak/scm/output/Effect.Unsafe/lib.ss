#!r6rs
#!chezscheme
(library
  (Effect.Unsafe lib)
  (export
    unsafePerformEffect)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Effect.Unsafe foreign))
  )
