[<AutoOpen>]
module PureScript_Test_BenchCheck

open System
open System.Collections.Generic

let Test_BenchCheck_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((box Bench_benchNow)))))) (box ((box (fun (t1: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((box Bench_benchNow)))))) (box ((box (fun (t2: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "Delta: ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showNumber)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringNumber)))))) (box ((box t2)))))) (box ((box t1)))))))))))))))))))))))
