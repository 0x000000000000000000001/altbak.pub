#!r6rs
#!chezscheme
(library
  (TestEffect lib)
  (export
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define main
    (scm:let ([_0 (Effect.Console.log (rt:string->pstring "Hello"))])
      (scm:lambda ()
        (scm:let ([_ (_0)])
          ((Effect.Console.log (rt:string->pstring "World"))))))))
