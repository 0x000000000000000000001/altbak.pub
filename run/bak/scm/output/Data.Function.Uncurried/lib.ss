#!r6rs
#!chezscheme
(library
  (Data.Function.Uncurried lib)
  (export
    mkFn0
    mkFn1
    mkFn10
    mkFn2
    mkFn3
    mkFn4
    mkFn5
    mkFn6
    mkFn7
    mkFn8
    mkFn9
    runFn0
    runFn1
    runFn10
    runFn2
    runFn3
    runFn4
    runFn5
    runFn6
    runFn7
    runFn8
    runFn9)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Data.Function.Uncurried foreign))

  (scm:define runFn1
    (scm:lambda (f0)
      f0))

  (scm:define mkFn1
    (scm:lambda (f0)
      f0)))
