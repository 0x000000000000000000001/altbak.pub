#!r6rs
#!chezscheme
(library
  (Data.Functor.Invariant lib)
  (export
    imap
    imapF
    invariantAdditive
    invariantAlternate
    invariantArray
    invariantConj
    invariantDisj
    invariantDual
    invariantEndo
    invariantFn
    invariantMultiplicative)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Semigroupoid lib) Control.Semigroupoid.)
    (prefix (Data.Functor lib) Data.Functor.))

  (scm:define invariantMultiplicative
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (f0 v12)))))))

  (scm:define invariantEndo
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (ab0)
      (scm:lambda (ba1)
        (scm:lambda (v2)
          (scm:lambda (x3)
            (ab0 (v2 (ba1 x3))))))))))

  (scm:define invariantDual
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (f0 v12)))))))

  (scm:define invariantDisj
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (f0 v12)))))))

  (scm:define invariantConj
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (f0 v12)))))))

  (scm:define invariantAdditive
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (f0 v12)))))))

  (scm:define imapF
    (scm:lambda (dictFunctor0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1)))))

  (scm:define invariantArray
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (Data.Functor.arrayMap f0))))))

  (scm:define invariantFn
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        ((rt:record-ref Control.Semigroupoid.semigroupoidFn (scm:string->symbol "compose")) f0))))))

  (scm:define imap
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "imap"))))

  (scm:define invariantAlternate
    (scm:lambda (dictInvariant0)
      (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            ((((rt:record-ref dictInvariant0 (scm:string->symbol "imap")) f1) g2) v3)))))))))
