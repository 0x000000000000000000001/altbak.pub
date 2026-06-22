#!r6rs
#!chezscheme
(library
  (Control.Bind lib)
  (export
    arrayBind
    bind
    bindArray
    bindFlipped
    bindFn
    bindProxy
    composeKleisli
    composeKleisliFlipped
    discard
    discardProxy
    discardUnit
    identity
    ifM
    join)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Apply lib) Control.Apply.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Control.Bind foreign))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define discard
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "discard"))))

  (scm:define bindProxy
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      Control.Apply.applyProxy))))

  (scm:define bindFn
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (m0)
      (scm:lambda (f1)
        (scm:lambda (x2)
          ((f1 (m0 x2)) x2))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      Control.Apply.applyFn))))

  (scm:define bindArray
    (scm:list (scm:cons (scm:string->symbol "bind") arrayBind) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      Control.Apply.applyArray))))

  (scm:define bind
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bind"))))

  (scm:define bindFlipped
    (scm:lambda (dictBind0)
      (scm:lambda (b1)
        (scm:lambda (a2)
          (((rt:record-ref dictBind0 (scm:string->symbol "bind")) a2) b1)))))

  (scm:define composeKleisliFlipped
    (scm:lambda (dictBind0)
      (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (a3)
            (((rt:record-ref dictBind0 (scm:string->symbol "bind")) (g2 a3)) f1))))))

  (scm:define composeKleisli
    (scm:lambda (dictBind0)
      (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (a3)
            (((rt:record-ref dictBind0 (scm:string->symbol "bind")) (f1 a3)) g2))))))

  (scm:define discardProxy
    (scm:list (scm:cons (scm:string->symbol "discard") (scm:lambda (dictBind0)
      (rt:record-ref dictBind0 (scm:string->symbol "bind"))))))

  (scm:define discardUnit
    (scm:list (scm:cons (scm:string->symbol "discard") (scm:lambda (dictBind0)
      (rt:record-ref dictBind0 (scm:string->symbol "bind"))))))

  (scm:define ifM
    (scm:lambda (dictBind0)
      (scm:lambda (cond1)
        (scm:lambda (t2)
          (scm:lambda (f3)
            (((rt:record-ref dictBind0 (scm:string->symbol "bind")) cond1) (scm:lambda (cond$p4)
              (scm:cond
                [cond$p4 t2]
                [scm:else f3]))))))))

  (scm:define join
    (scm:lambda (dictBind0)
      (scm:lambda (m1)
        (((rt:record-ref dictBind0 (scm:string->symbol "bind")) m1) identity)))))
