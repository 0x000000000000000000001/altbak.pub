(library (Bench foreign)
  (export benchNow opaque formatNumber)
  (import (chezscheme))

  (define (benchNow)
    (lambda () 0.0))

  (define (opaque a)
    (lambda () a))

  (define (formatNumber n)
    (number->string n)))
