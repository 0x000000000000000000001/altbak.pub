#!r6rs
#!chezscheme
(library
  (Spago.Generated.BuildInfo lib)
  (export
    packages
    pursVersion
    spagoVersion)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define spagoVersion
    (rt:string->pstring "1.0.4"))

  (scm:define pursVersion
    (rt:string->pstring "0.15.16"))

  (scm:define packages
    (scm:list (scm:cons (scm:string->symbol "ps-scheme-test") (rt:string->pstring "0.0.0")))))
