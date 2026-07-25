#!r6rs
#!chezscheme
(library
  (Effect.Class.Console lib)
  (export
    clear
    debug
    debugShow
    error
    errorShow
    info
    infoShow
    log
    logShow
    time
    timeEnd
    timeLog
    warn
    warnShow)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define warnShow
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (dictShow1)
        (scm:lambda (x2)
          ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.warn ((rt:record-ref dictShow1 (scm:string->symbol "show")) x2)))))))

  (scm:define warn
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (x1)
        ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.warn x1)))))

  (scm:define timeLog
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (x1)
        ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.timeLog x1)))))

  (scm:define timeEnd
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (x1)
        ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.timeEnd x1)))))

  (scm:define time
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (x1)
        ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.time x1)))))

  (scm:define logShow
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (dictShow1)
        (scm:lambda (x2)
          ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.log ((rt:record-ref dictShow1 (scm:string->symbol "show")) x2)))))))

  (scm:define log
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (x1)
        ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.log x1)))))

  (scm:define infoShow
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (dictShow1)
        (scm:lambda (x2)
          ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.info ((rt:record-ref dictShow1 (scm:string->symbol "show")) x2)))))))

  (scm:define info
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (x1)
        ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.info x1)))))

  (scm:define errorShow
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (dictShow1)
        (scm:lambda (x2)
          ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.error ((rt:record-ref dictShow1 (scm:string->symbol "show")) x2)))))))

  (scm:define error
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (x1)
        ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.error x1)))))

  (scm:define debugShow
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (dictShow1)
        (scm:lambda (x2)
          ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.debug ((rt:record-ref dictShow1 (scm:string->symbol "show")) x2)))))))

  (scm:define debug
    (scm:lambda (dictMonadEffect0)
      (scm:lambda (x1)
        ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) (Effect.Console.debug x1)))))

  (scm:define clear
    (scm:lambda (dictMonadEffect0)
      ((rt:record-ref dictMonadEffect0 (scm:string->symbol "liftEffect")) Effect.Console.clear))))
