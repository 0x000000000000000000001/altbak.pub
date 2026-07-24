(library (Bench foreign)
  (export benchNow opaque formatNumber)
  (import (except (chezscheme) opaque)
          (prefix (purescm runtime) rt:))

  (define benchNow
    (lambda ()
      (let ([t (current-time)])
        (+ (* (time-second t) 1e6)
           (/ (time-nanosecond t) 1e3)))))

  (define (opaque a)
    (lambda () a))

  (define (formatNumber n)
    (rt:string->pstring (number->string n))))
