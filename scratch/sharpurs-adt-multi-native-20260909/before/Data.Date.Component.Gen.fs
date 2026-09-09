[<AutoOpen>]
module PureScript_Data_Date_Component_Gen

open System
open System.Collections.Generic

let Data_Date_Component_Gen_toEnum  = (sharpurs_apply (box ((box Data_Enum_toEnum))) (box ((box Data_Date_Component_boundedEnumYear))))

let Data_Date_Component_Gen_genYear  = (box (fun (dictMonadGen: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Apply0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Bind1" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Monad0" (unbox<Map<string, obj>> ((box dictMonadGen)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box Data_Maybe_fromJust))) (box ((box Prim_undefined)))))))))))))) (box ((box Data_Date_Component_Gen_toEnum))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_Gen_Class_chooseInt))) (box ((box dictMonadGen)))))) (box ((box 1900)))))) (box ((box 2100)))))))))

let Data_Date_Component_Gen_genWeekday  = (box (fun (dictMonadGen: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_Gen_genBoundedEnum))) (box ((box dictMonadGen)))))) (box ((box Data_Date_Component_boundedEnumWeekday))))))

let Data_Date_Component_Gen_genMonth  = (box (fun (dictMonadGen: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_Gen_genBoundedEnum))) (box ((box dictMonadGen)))))) (box ((box Data_Date_Component_boundedEnumMonth))))))

let Data_Date_Component_Gen_genDay  = (box (fun (dictMonadGen: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_Gen_genBoundedEnum))) (box ((box dictMonadGen)))))) (box ((box Data_Date_Component_boundedEnumDay))))))
