#!r6rs
#!chezscheme
(library
  (Data.Newtype lib)
  (export
    ala
    alaF
    collect
    modify
    newtypeAdditive
    newtypeConj
    newtypeDisj
    newtypeDual
    newtypeEndo
    newtypeFirst
    newtypeLast
    newtypeMultiplicative
    over
    over2
    overF
    overF2
    traverse
    un
    under
    under2
    underF
    underF2
    unwrap
    wrap)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define wrap
    (scm:lambda (_)
      Unsafe.Coerce.unsafeCoerce))

  (scm:define unwrap
    (scm:lambda (_)
      Unsafe.Coerce.unsafeCoerce))

  (scm:define underF2
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (_)
          (scm:lambda (_)
            (scm:lambda (v4)
              Unsafe.Coerce.unsafeCoerce))))))

  (scm:define underF
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (_)
          (scm:lambda (_)
            (scm:lambda (v4)
              Unsafe.Coerce.unsafeCoerce))))))

  (scm:define under2
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (v2)
          Unsafe.Coerce.unsafeCoerce))))

  (scm:define under
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (v2)
          Unsafe.Coerce.unsafeCoerce))))

  (scm:define un
    (scm:lambda (_)
      (scm:lambda (v1)
        Unsafe.Coerce.unsafeCoerce)))

  (scm:define traverse
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (v2)
          Unsafe.Coerce.unsafeCoerce))))

  (scm:define overF2
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (_)
          (scm:lambda (_)
            (scm:lambda (v4)
              Unsafe.Coerce.unsafeCoerce))))))

  (scm:define overF
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (_)
          (scm:lambda (_)
            (scm:lambda (v4)
              Unsafe.Coerce.unsafeCoerce))))))

  (scm:define over2
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (v2)
          Unsafe.Coerce.unsafeCoerce))))

  (scm:define over
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (v2)
          Unsafe.Coerce.unsafeCoerce))))

  (scm:define newtypeMultiplicative
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define newtypeLast
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define newtypeFirst
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define newtypeEndo
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define newtypeDual
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define newtypeDisj
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define newtypeConj
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define newtypeAdditive
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define modify
    (scm:lambda (_)
      (scm:lambda (fn1)
        (scm:lambda (t2)
          (fn1 t2)))))

  (scm:define collect
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (v2)
          Unsafe.Coerce.unsafeCoerce))))

  (scm:define alaF
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (_)
          (scm:lambda (_)
            (scm:lambda (v4)
              Unsafe.Coerce.unsafeCoerce))))))

  (scm:define ala
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (_)
          (scm:lambda (v3)
            (scm:lambda (f4)
              (f4 Unsafe.Coerce.unsafeCoerce))))))))
