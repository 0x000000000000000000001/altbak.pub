#!r6rs
#!chezscheme
(library
  (Control.Biapply lib)
  (export
    biapply
    biapplyFirst
    biapplySecond
    biapplyTuple
    bilift2
    bilift3
    identity)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Bifunctor lib) Data.Bifunctor.)
    (prefix (Data.Function lib) Data.Function.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define biapplyTuple
    (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v0)
      (scm:lambda (v11)
        (Data.Tuple.Tuple* ((Data.Tuple.Tuple-value0 v0) (Data.Tuple.Tuple-value0 v11)) ((Data.Tuple.Tuple-value1 v0) (Data.Tuple.Tuple-value1 v11)))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
      Data.Bifunctor.bifunctorTuple))))

  (scm:define biapply
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "biapply"))))

  (scm:define biapplyFirst
    (scm:lambda (dictBiapply0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (((rt:record-ref dictBiapply0 (scm:string->symbol "biapply")) ((((rt:record-ref ((rt:record-ref dictBiapply0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined)) (scm:string->symbol "bimap")) (scm:lambda (v3)
            identity)) (scm:lambda (v3)
            identity)) a1)) b2)))))

  (scm:define biapplySecond
    (scm:lambda (dictBiapply0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (((rt:record-ref dictBiapply0 (scm:string->symbol "biapply")) ((((rt:record-ref ((rt:record-ref dictBiapply0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined)) (scm:string->symbol "bimap")) Data.Function.const) Data.Function.const) a1)) b2)))))

  (scm:define bilift2
    (scm:lambda (dictBiapply0)
      (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (a3)
            (scm:lambda (b4)
              (((rt:record-ref dictBiapply0 (scm:string->symbol "biapply")) ((((rt:record-ref ((rt:record-ref dictBiapply0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined)) (scm:string->symbol "bimap")) f1) g2) a3)) b4)))))))

  (scm:define bilift3
    (scm:lambda (dictBiapply0)
      (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (a3)
            (scm:lambda (b4)
              (scm:lambda (c5)
                (((rt:record-ref dictBiapply0 (scm:string->symbol "biapply")) (((rt:record-ref dictBiapply0 (scm:string->symbol "biapply")) ((((rt:record-ref ((rt:record-ref dictBiapply0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined)) (scm:string->symbol "bimap")) f1) g2) a3)) b4)) c5)))))))))
