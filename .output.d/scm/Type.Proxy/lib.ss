#!r6rs
#!chezscheme
(library
  (Type.Proxy lib)
  (export
    Proxy
    Proxy?)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Proxy
    (scm:quote Proxy))

  (scm:define Proxy?
    (scm:lambda (v)
      (scm:eq? (scm:quote Proxy) v))))
