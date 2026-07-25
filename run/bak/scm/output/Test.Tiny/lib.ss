#!r6rs
#!chezscheme
(library
  (Test.Tiny lib)
  (export
    Circle
    Circle-value0
    Circle?
    Rect
    Rect*
    Rect-value0
    Rect-value1
    Rect?
    area)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define-record-type (Circle$ Circle Circle?)
    (scm:fields (scm:immutable value0 Circle-value0)))

  (scm:define-record-type (Rect$ Rect* Rect?)
    (scm:fields (scm:immutable value0 Rect-value0) (scm:immutable value1 Rect-value1)))

  (scm:define Rect
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Rect* value0 value1))))

  (scm:define area
    (scm:lambda (v0)
      (scm:cond
        [(Circle? v0) (scm:fx* (Circle-value0 v0) (Circle-value0 v0))]
        [(Rect? v0) (scm:fx* (Rect-value0 v0) (Rect-value1 v0))]
        [scm:else (rt:fail)]))))
