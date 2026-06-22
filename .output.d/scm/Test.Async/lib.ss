#!r6rs
#!chezscheme
(library
  (Test.Async lib)
  (export
    act
    describe
    runAsyncTest)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Effect.Console lib) Effect.Console.)
    (Test.Async foreign))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Asynchronous Concurrency (1000 forks):")))

  (scm:define act
    (runAsyncTest 1000)))
