#!r6rs
#!chezscheme
(library
  (Effect.Console lib)
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
    (Effect.Console foreign))

  (scm:define warnShow
    (scm:lambda (dictShow0)
      (scm:lambda (a1)
        (warn ((rt:record-ref dictShow0 (scm:string->symbol "show")) a1)))))

  (scm:define logShow
    (scm:lambda (dictShow0)
      (scm:lambda (a1)
        (log ((rt:record-ref dictShow0 (scm:string->symbol "show")) a1)))))

  (scm:define infoShow
    (scm:lambda (dictShow0)
      (scm:lambda (a1)
        (info ((rt:record-ref dictShow0 (scm:string->symbol "show")) a1)))))

  (scm:define errorShow
    (scm:lambda (dictShow0)
      (scm:lambda (a1)
        (error ((rt:record-ref dictShow0 (scm:string->symbol "show")) a1)))))

  (scm:define debugShow
    (scm:lambda (dictShow0)
      (scm:lambda (a1)
        (debug ((rt:record-ref dictShow0 (scm:string->symbol "show")) a1))))))
