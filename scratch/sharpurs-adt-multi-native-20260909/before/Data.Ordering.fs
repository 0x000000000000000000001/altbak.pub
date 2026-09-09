[<AutoOpen>]
module PureScript_Data_Ordering

open System
open System.Collections.Generic

type Data_Ordering_Ordering =
  | Data_Ordering_LTusd_Ctor
  | Data_Ordering_GTusd_Ctor
  | Data_Ordering_EQusd_Ctor

let Data_Ordering_LT  = (box Data_Ordering_LTusd_Ctor)

let Data_Ordering_GT  = (box Data_Ordering_GTusd_Ctor)

let Data_Ordering_EQ  = (box Data_Ordering_EQusd_Ctor)

let Data_Ordering_showOrdering  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Ordering_LTusd_Ctor -> ((box "LT")) | Data_Ordering_GTusd_Ctor -> ((box "GT")) | Data_Ordering_EQusd_Ctor -> ((box "EQ"))))))) Map.empty))))))

let Data_Ordering_semigroupOrdering  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Ordering_LTusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (Data_Ordering_GTusd_Ctor, _) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Ordering_EQusd_Ctor, y) -> ((box y))))))))) Map.empty))))))

let Data_Ordering_invert  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Ordering_GTusd_Ctor -> ((box Data_Ordering_LTusd_Ctor)) | Data_Ordering_EQusd_Ctor -> ((box Data_Ordering_EQusd_Ctor)) | Data_Ordering_LTusd_Ctor -> ((box Data_Ordering_GTusd_Ctor)))))

let Data_Ordering_eqOrdering  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Ordering_LTusd_Ctor, Data_Ordering_LTusd_Ctor) -> ((box true)) | (Data_Ordering_GTusd_Ctor, Data_Ordering_GTusd_Ctor) -> ((box true)) | (Data_Ordering_EQusd_Ctor, Data_Ordering_EQusd_Ctor) -> ((box true)) | (_, _) -> ((box false))))))))) Map.empty))))))
