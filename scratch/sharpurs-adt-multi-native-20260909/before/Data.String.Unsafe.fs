[<AutoOpen>]
module PureScript_Data_String_Unsafe

open System
open System.Collections.Generic

let Data_String_Unsafe_char = box (fun (arg0: obj) -> box (Data.String.Unsafe.FFI.Char(unbox arg0)))
let Data_String_Unsafe_charAt = box (fun (arg0: obj) -> box (Data.String.Unsafe.FFI.CharAt(unbox arg0)))



