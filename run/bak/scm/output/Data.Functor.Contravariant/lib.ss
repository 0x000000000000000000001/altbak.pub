#!r6rs
#!chezscheme
(library
  (Data.Functor.Contravariant lib)
  (export
    cmap
    cmapFlipped
    coerce
    contravariantConst
    imapC)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Void lib) Data.Void.))

  (scm:define contravariantConst
    (scm:list (scm:cons (scm:string->symbol "cmap") (scm:lambda (v0)
      (scm:lambda (v11)
        v11)))))

  (scm:define cmap
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "cmap"))))

  (scm:define cmapFlipped
    (scm:lambda (dictContravariant0)
      (scm:lambda (x1)
        (scm:lambda (f2)
          (((rt:record-ref dictContravariant0 (scm:string->symbol "cmap")) f2) x1)))))

  (scm:define coerce
    (scm:lambda (dictContravariant0)
      (scm:lambda (dictFunctor1)
        (scm:lambda (a2)
          (((rt:record-ref dictFunctor1 (scm:string->symbol "map")) Data.Void.absurd) (((rt:record-ref dictContravariant0 (scm:string->symbol "cmap")) Data.Void.absurd) a2))))))

  (scm:define imapC
    (scm:lambda (dictContravariant0)
      (scm:lambda (v1)
        (scm:lambda (f2)
          ((rt:record-ref dictContravariant0 (scm:string->symbol "cmap")) f2))))))
