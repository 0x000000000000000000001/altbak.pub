#!r6rs
#!chezscheme
(library
  (Control.Monad lib)
  (export
    ap
    liftM1
    monadArray
    monadFn
    monadProxy
    unlessM
    whenM)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Applicative lib) Control.Applicative.)
    (prefix (Control.Bind lib) Control.Bind.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define whenM
    (scm:lambda (dictMonad0)
      (scm:let ([_1 ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined))])
        (scm:lambda (mb2)
          (scm:lambda (m3)
            (((rt:record-ref ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "bind")) mb2) (scm:lambda (b4)
              (scm:cond
                [b4 m3]
                [scm:else ((rt:record-ref _1 (scm:string->symbol "pure")) Data.Unit.unit)]))))))))

  (scm:define unlessM
    (scm:lambda (dictMonad0)
      (scm:let ([_1 ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined))])
        (scm:lambda (mb2)
          (scm:lambda (m3)
            (((rt:record-ref ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "bind")) mb2) (scm:lambda (b4)
              (scm:cond
                [(scm:not b4) m3]
                [b4 ((rt:record-ref _1 (scm:string->symbol "pure")) Data.Unit.unit)]
                [scm:else (rt:fail)]))))))))

  (scm:define monadProxy
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      Control.Applicative.applicativeProxy)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      Control.Bind.bindProxy))))

  (scm:define monadFn
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      Control.Applicative.applicativeFn)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      Control.Bind.bindFn))))

  (scm:define monadArray
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      Control.Applicative.applicativeArray)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      Control.Bind.bindArray))))

  (scm:define liftM1
    (scm:lambda (dictMonad0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (((rt:record-ref ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "bind")) a2) (scm:lambda (a$p3)
            ((rt:record-ref ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure")) (f1 a$p3))))))))

  (scm:define ap
    (scm:lambda (dictMonad0)
      (scm:let ([_1 ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined))])
        (scm:lambda (f2)
          (scm:lambda (a3)
            (((rt:record-ref _1 (scm:string->symbol "bind")) f2) (scm:lambda (f$p4)
              (((rt:record-ref _1 (scm:string->symbol "bind")) a3) (scm:lambda (a$p5)
                ((rt:record-ref ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure")) (f$p4 a$p5))))))))))))
