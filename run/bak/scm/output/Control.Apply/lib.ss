#!r6rs
#!chezscheme
(library
  (Control.Apply lib)
  (export
    apply
    applyArray
    applyFirst
    applyFn
    applyProxy
    applySecond
    arrayApply
    identity
    lift2
    lift3
    lift4
    lift5)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Function lib) Data.Function.)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Control.Apply foreign))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define applyProxy
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Functor.functorProxy))))

  (scm:define applyFn
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (x2)
          ((f0 x2) (g1 x2)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Functor.functorFn))))

  (scm:define applyArray
    (scm:list (scm:cons (scm:string->symbol "apply") arrayApply) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Functor.functorArray))))

  (scm:define apply
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "apply"))))

  (scm:define applyFirst
    (scm:lambda (dictApply0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Function.const) a1)) b2)))))

  (scm:define applySecond
    (scm:lambda (dictApply0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (v3)
            identity)) a1)) b2)))))

  (scm:define lift2
    (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) f1) a2)) b3))))))

  (scm:define lift3
    (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) f1) a2)) b3)) c4)))))))

  (scm:define lift4
    (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) f1) a2)) b3)) c4)) d5))))))))

  (scm:define lift5
    (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (scm:lambda (e6)
                  (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) f1) a2)) b3)) c4)) d5)) e6))))))))))
