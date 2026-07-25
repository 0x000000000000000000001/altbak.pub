#!r6rs
#!chezscheme
(library
  (Data.Array.Partial lib)
  (export
    head
    init
    last
    tail)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Array lib) Data.Array.))

  (scm:define tail
    (scm:lambda (_)
      (scm:lambda (xs1)
        (Data.Array.sliceImpl 1 (rt:array-length xs1) xs1))))

  (scm:define last
    (scm:lambda (_)
      (scm:lambda (xs1)
        (rt:array-ref xs1 (scm:fx- (rt:array-length xs1) 1)))))

  (scm:define init
    (scm:lambda (_)
      (scm:lambda (xs1)
        (Data.Array.sliceImpl 0 (scm:fx- (rt:array-length xs1) 1) xs1))))

  (scm:define head
    (scm:lambda (_)
      (scm:lambda (xs1)
        (rt:array-ref xs1 0)))))
