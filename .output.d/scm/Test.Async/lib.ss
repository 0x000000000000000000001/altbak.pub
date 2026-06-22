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
    (scm:let ([_0 (runAsyncTest 1000)])
      (scm:lambda ()
        (scm:let ([_ (_0)])
          ((Effect.Console.log (rt:string->pstring "Done"))))))))
