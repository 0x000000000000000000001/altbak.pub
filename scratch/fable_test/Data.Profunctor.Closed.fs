[<AutoOpen>]
module PureScript_Data_Profunctor_Closed

open System
open System.Collections.Generic

let Data_Profunctor_Closed_Closedusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Profunctor_Closed_closedFunction  = (sharpurs_apply (box ((box Data_Profunctor_Closed_Closedusd_Dict))) (box ((box ((Map.add "closed" (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (Map.add "Profunctor0" (box ((box (fun (usd__unused: obj) -> (box Data_Profunctor_profunctorFn))))) Map.empty)))))))

let Data_Profunctor_Closed_closed  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "closed" (unbox<Map<string, obj>> ((box v))))))))
