[<AutoOpen>]
module PureScript_Data_Char

open System
open System.Collections.Generic

let Data_Char_toCharCode  = (sharpurs_apply (box ((box Data_Enum_fromEnum))) (box ((box Data_Enum_boundedEnumChar))))

let Data_Char_fromCharCode  = (sharpurs_apply (box ((box Data_Enum_toEnum))) (box ((box Data_Enum_boundedEnumChar))))
