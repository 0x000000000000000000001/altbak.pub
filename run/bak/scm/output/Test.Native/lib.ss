#!r6rs
#!chezscheme
(library
  (Test.Native lib)
  (export
    loopNative)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define loopNative
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx=? v0 0) v11]
          [scm:else ((loopNative (scm:fx- v0 1)) (scm:fx+ v11 1))])))))
