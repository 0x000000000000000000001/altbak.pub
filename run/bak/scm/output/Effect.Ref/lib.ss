#!r6rs
#!chezscheme
(library
  (Effect.Ref lib)
  (export
    _new
    modify
    modify$p
    modifyImpl
    modify_
    new
    newWithSelf
    read
    write)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Unit lib) Data.Unit.)
    (Effect.Ref foreign))

  (scm:define new
    _new)

  (scm:define modify$p
    modifyImpl)

  (scm:define modify
    (scm:lambda (f0)
      (modifyImpl (scm:lambda (s1)
        (scm:let ([s$p2 (f0 s1)])
          (scm:list (scm:cons (scm:string->symbol "state") s$p2) (scm:cons (scm:string->symbol "value") s$p2)))))))

  (scm:define modify_
    (scm:lambda (f0)
      (scm:lambda (s1)
        (scm:let ([_2 ((modifyImpl (scm:lambda (s2)
          (scm:let ([s$p3 (f0 s2)])
            (scm:list (scm:cons (scm:string->symbol "state") s$p3) (scm:cons (scm:string->symbol "value") s$p3))))) s1)])
          (scm:lambda ()
            (scm:let ([a$p3 (_2)])
              Data.Unit.unit)))))))
