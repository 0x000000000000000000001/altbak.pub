[<AutoOpen>]
module PureScript_Data_Time_Component_Gen

open System
open System.Collections.Generic

let Data_Time_Component_Gen_genSecond  = (box (fun (dictMonadGen: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_Gen_genBoundedEnum))) (box ((box dictMonadGen)))))) (box ((box Data_Time_Component_boundedEnumSecond))))))

let Data_Time_Component_Gen_genMinute  = (box (fun (dictMonadGen: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_Gen_genBoundedEnum))) (box ((box dictMonadGen)))))) (box ((box Data_Time_Component_boundedEnumMinute))))))

let Data_Time_Component_Gen_genMillisecond  = (box (fun (dictMonadGen: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_Gen_genBoundedEnum))) (box ((box dictMonadGen)))))) (box ((box Data_Time_Component_boundedEnumMillisecond))))))

let Data_Time_Component_Gen_genHour  = (box (fun (dictMonadGen: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_Gen_genBoundedEnum))) (box ((box dictMonadGen)))))) (box ((box Data_Time_Component_boundedEnumHour))))))
