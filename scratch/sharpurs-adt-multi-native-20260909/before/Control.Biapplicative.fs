[<AutoOpen>]
module PureScript_Control_Biapplicative

open System
open System.Collections.Generic

let Control_Biapplicative_Biapplicativeusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Biapplicative_bipure  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "bipure" (unbox<Map<string, obj>> ((box v))))))))

let Control_Biapplicative_biapplicativeTuple  = (sharpurs_apply (box ((box Control_Biapplicative_Biapplicativeusd_Dict))) (box ((box ((Map.add "bipure" (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (Map.add "Biapply0" (box ((box (fun (usd__unused: obj) -> (box Control_Biapply_biapplyTuple))))) Map.empty)))))))
