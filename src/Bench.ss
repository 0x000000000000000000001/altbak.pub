(library (Bench foreign)
  (export benchNow opaque formatNumber)
  (import (except (chezscheme) opaque)
          (prefix (purescm runtime) rt:))

  (define origin (current-time 'time-monotonic))
  (define benchNow
    (lambda ()
      (let ([elapsed (time-difference (current-time 'time-monotonic) origin)])
        (+ (* (time-second elapsed) 1e6)
           (/ (time-nanosecond elapsed) 1e3)))))

  (define (opaque a)
    (lambda () a))

  (define (formatNumber n)
    (rt:string->pstring (number->string n))))
