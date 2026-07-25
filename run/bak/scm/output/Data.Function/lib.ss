#!r6rs
#!chezscheme
(library
  (Data.Function lib)
  (export
    apply
    applyFlipped
    applyN
    const
    flip
    on)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define on
    (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((f0 (g1 x2)) (g1 y3)))))))

  (scm:define flip
    (scm:lambda (f0)
      (scm:lambda (b1)
        (scm:lambda (a2)
          ((f0 a2) b1)))))

  (scm:define const
    (scm:lambda (a0)
      (scm:lambda (v1)
        a0)))

  (scm:define applyN
    (scm:lambda (f0)
      (scm:letrec ([go1 (scm:lambda (n2)
        (scm:lambda (acc3)
          (scm:cond
            [(scm:fx<=? n2 0) acc3]
            [scm:else ((go1 (scm:fx- n2 1)) (f0 acc3))])))])
        go1)))

  (scm:define applyFlipped
    (scm:lambda (x0)
      (scm:lambda (f1)
        (f1 x0))))

  (scm:define apply
    (scm:lambda (f0)
      (scm:lambda (x1)
        (f0 x1)))))
