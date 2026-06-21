#!r6rs
#!chezscheme
(library
  (App lib)
  (export
    fib
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define fib
    (scm:lambda (v0)
      (scm:cond
        [(scm:fx=? v0 0) 0]
        [(scm:fx=? v0 1) 1]
        [scm:else (scm:fx+ (fib (scm:fx- v0 1)) (fib (scm:fx- v0 2)))])))

  (scm:define main
    (Effect.Console.log (Data.Show.showIntImpl (fib 40)))))
