[<AutoOpen>]
module PureScript_Test_RowToList

open System
open System.Collections.Generic

let Test_RowToList_RecordKeysusd_Dict  = (box (fun (x: obj) -> (box x)))

let Test_RowToList_keysNil  = (sharpurs_apply (box ((box Test_RowToList_RecordKeysusd_Dict))) (box ((box ((Map.add "keysImpl" (box ((box (fun (v: obj) -> (box 0))))) Map.empty))))))

let Test_RowToList_keysImpl  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "keysImpl" (unbox<Map<string, obj>> ((box v))))))))

let Test_RowToList_keysCons  = (box (fun (dictRecordKeys: obj) -> (sharpurs_apply (box ((box Test_RowToList_RecordKeysusd_Dict))) (box ((box ((Map.add "keysImpl" (box ((box (fun (v: obj) -> (box ((unbox<int> (box ((box 1)))) + (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_RowToList_keysImpl))) (box ((box dictRecordKeys)))))) (box ((box Type_Proxy_Proxyusd_Ctor))))))))))))) Map.empty))))))))

let Test_RowToList_keysCons1  = (sharpurs_apply (box ((box Test_RowToList_keysCons))) (box ((sharpurs_apply (box ((box Test_RowToList_keysCons))) (box ((sharpurs_apply (box ((box Test_RowToList_keysCons))) (box ((sharpurs_apply (box ((box Test_RowToList_keysCons))) (box ((sharpurs_apply (box ((box Test_RowToList_keysCons))) (box ((box Test_RowToList_keysNil))))))))))))))))

let Test_RowToList_keys  = (box (fun (usd__unused: obj) -> (box (fun (dictRecordKeys: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Test_RowToList_keysImpl))) (box ((box dictRecordKeys)))))) (box ((box Type_Proxy_Proxyusd_Ctor))))))))))

let Test_RowToList_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "RowToList (Keys Count):"))))

let Test_RowToList_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000))))))))) (box ((box (fun (usd__unused: obj) -> (let rec_var = (box ((Map.add "a" (box ((box 1))) (Map.add "b" (box ((box "two"))) (Map.add "c" (box ((box true))) (Map.add "d" (box ((box 4.0))) (Map.add "e" (box ((box "five"))) Map.empty))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_RowToList_keys))) (box ((box Prim_undefined)))))) (box ((box Test_RowToList_keysCons1)))))) (box ((box rec_var))))))))))))))))
