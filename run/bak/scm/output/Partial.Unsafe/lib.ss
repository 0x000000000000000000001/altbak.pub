#!r6rs
#!chezscheme
(library
  (Partial.Unsafe lib)
  (export
    _unsafePartial
    unsafeCrashWith
    unsafePartial)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Partial lib) Partial.)
    (Partial.Unsafe foreign))

  (scm:define unsafePartial
    _unsafePartial)

  (scm:define unsafeCrashWith
    (scm:lambda (msg0)
      (Partial._crashWith msg0))))
