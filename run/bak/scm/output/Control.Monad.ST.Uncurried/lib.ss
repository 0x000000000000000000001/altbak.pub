#!r6rs
#!chezscheme
(library
  (Control.Monad.ST.Uncurried lib)
  (export
    mkSTFn1
    mkSTFn10
    mkSTFn2
    mkSTFn3
    mkSTFn4
    mkSTFn5
    mkSTFn6
    mkSTFn7
    mkSTFn8
    mkSTFn9
    runSTFn1
    runSTFn10
    runSTFn2
    runSTFn3
    runSTFn4
    runSTFn5
    runSTFn6
    runSTFn7
    runSTFn8
    runSTFn9)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Control.Monad.ST.Uncurried foreign))
  )
