#!r6rs
#!chezscheme
(library
  (Data.Array.ST.Partial lib)
  (export
    peek
    peekImpl
    poke
    pokeImpl)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Monad.ST.Uncurried lib) Control.Monad.ST.Uncurried.)
    (Data.Array.ST.Partial foreign))

  (scm:define poke
    (scm:lambda (_)
      (Control.Monad.ST.Uncurried.runSTFn3 pokeImpl)))

  (scm:define peek
    (scm:lambda (_)
      (Control.Monad.ST.Uncurried.runSTFn2 peekImpl))))
