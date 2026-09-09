[<AutoOpen>]
module PureScript_Control_Alt

open System
open System.Collections.Generic

let Control_Alt_Altusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Alt_altArray  = (sharpurs_apply (box ((box Control_Alt_Altusd_Dict))) (box ((box ((Map.add "alt" (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupArray)))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Functor_functorArray))))) Map.empty)))))))

let Control_Alt_alt  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "alt" (unbox<Map<string, obj>> ((box v))))))))
