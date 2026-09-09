[<AutoOpen>]
module PureScript_Data_CommutativeRing

open System
open System.Collections.Generic

let Data_CommutativeRing_ringRecord  = (sharpurs_apply (box ((box Data_Ring_ringRecord))) (box ((box Prim_undefined))))

let Data_CommutativeRing_CommutativeRingRecordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_CommutativeRing_CommutativeRingusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_CommutativeRing_commutativeRingUnit  = (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingusd_Dict))) (box ((box ((Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box Data_Ring_ringUnit))))) Map.empty))))))

let Data_CommutativeRing_commutativeRingRecordNil  = (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingRecordusd_Dict))) (box ((box ((Map.add "RingRecord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ring_ringRecordNil))))) Map.empty))))))

let Data_CommutativeRing_commutativeRingRecordCons  = (box (fun (dictIsSymbol: obj) -> (let ringRecordCons = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_ringRecordCons))) (box ((box dictIsSymbol)))))) (box ((box Prim_undefined)))) in (box (fun (usd__unused: obj) -> (box (fun (dictCommutativeRingRecord: obj) -> (let ringRecordCons1 = (sharpurs_apply (box ((box ringRecordCons))) (box ((sharpurs_apply (box ((Map.find "RingRecord0" (unbox<Map<string, obj>> ((box dictCommutativeRingRecord)))))) (box ((box Prim_undefined))))))) in (box (fun (dictCommutativeRing: obj) -> (let ringRecordCons2 = (sharpurs_apply (box ((box ringRecordCons1))) (box ((sharpurs_apply (box ((Map.find "Ring0" (unbox<Map<string, obj>> ((box dictCommutativeRing)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingRecordusd_Dict))) (box ((box ((Map.add "RingRecord0" (box ((box (fun (usd__unused: obj) -> (box ringRecordCons2))))) Map.empty)))))))))))))))))

let Data_CommutativeRing_commutativeRingRecord  = (box (fun (usd__unused: obj) -> (box (fun (dictCommutativeRingRecord: obj) -> (let ringRecord1 = (sharpurs_apply (box ((box Data_CommutativeRing_ringRecord))) (box ((sharpurs_apply (box ((Map.find "RingRecord0" (unbox<Map<string, obj>> ((box dictCommutativeRingRecord)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingusd_Dict))) (box ((box ((Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box ringRecord1))))) Map.empty)))))))))))

let Data_CommutativeRing_commutativeRingProxy  = (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingusd_Dict))) (box ((box ((Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box Data_Ring_ringProxy))))) Map.empty))))))

let Data_CommutativeRing_commutativeRingNumber  = (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingusd_Dict))) (box ((box ((Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box Data_Ring_ringNumber))))) Map.empty))))))

let Data_CommutativeRing_commutativeRingInt  = (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingusd_Dict))) (box ((box ((Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box Data_Ring_ringInt))))) Map.empty))))))

let Data_CommutativeRing_commutativeRingFn  = (box (fun (dictCommutativeRing: obj) -> (let ringFn = (sharpurs_apply (box ((box Data_Ring_ringFn))) (box ((sharpurs_apply (box ((Map.find "Ring0" (unbox<Map<string, obj>> ((box dictCommutativeRing)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingusd_Dict))) (box ((box ((Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box ringFn))))) Map.empty)))))))))
