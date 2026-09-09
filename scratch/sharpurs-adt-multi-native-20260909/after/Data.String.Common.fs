[<AutoOpen>]
module PureScript_Data_String_Common

open System
open System.Collections.Generic

let Data_String_Common__localeCompare = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (fun (arg3: obj) -> box (fun (arg4: obj) -> box (Data.String.Common.FFI._LocaleCompare(unbox arg0, unbox arg1, unbox arg2, unbox arg3, unbox arg4)))))))
let Data_String_Common_joinWith = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data.String.Common.FFI.JoinWith(unbox arg0, unbox arg1))))
let Data_String_Common_replace = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (Data.String.Common.FFI.Replace(unbox arg0, unbox arg1, unbox arg2)))))
let Data_String_Common_replaceAll = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (Data.String.Common.FFI.ReplaceAll(unbox arg0, unbox arg1, unbox arg2)))))
let Data_String_Common_split = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data.String.Common.FFI.Split(unbox arg0, unbox arg1))))
let Data_String_Common_toLower = box (fun (arg0: obj) -> box (Data.String.Common.FFI.ToLower(unbox arg0)))
let Data_String_Common_toUpper = box (fun (arg0: obj) -> box (Data.String.Common.FFI.ToUpper(unbox arg0)))
let Data_String_Common_trim = box (fun (arg0: obj) -> box (Data.String.Common.FFI.Trim(unbox arg0)))


let Data_String_Common_null  = (box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqString)))))) (box ((box s)))))) (box ((box ""))))))

let Data_String_Common_localeCompare  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_String_Common__localeCompare))) (box ((box Data_Ordering_LTusd_Ctor)))))) (box ((box Data_Ordering_EQusd_Ctor)))))) (box ((box Data_Ordering_GTusd_Ctor))))
