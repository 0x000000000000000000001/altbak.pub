[<AutoOpen>]
module PureScript_Test_Implicit

open System
open System.Collections.Generic

let Test_Implicit_foo  = (box (fun (x: obj) -> (box x)))

let Test_Implicit_bar  = (sharpurs_apply (box ((box Test_Implicit_foo))) (box ((box 42))))
