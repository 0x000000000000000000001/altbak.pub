[<AutoOpen>]
module PureScript_Data_Functor_Contravariant

open System
open System.Collections.Generic

let Data_Functor_Contravariant_Contravariantusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Functor_Contravariant_contravariantConst  = (sharpurs_apply (box ((box Data_Functor_Contravariant_Contravariantusd_Dict))) (box ((box ((Map.add "cmap" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (_, x) -> ((sharpurs_apply (box ((box Data_Const_Const))) (box ((box x)))))))))))) Map.empty))))))

let Data_Functor_Contravariant_cmap  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "cmap" (unbox<Map<string, obj>> ((box v))))))))

let Data_Functor_Contravariant_cmapFlipped  = (box (fun (dictContravariant: obj) -> (box (fun (x: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_Contravariant_cmap))) (box ((box dictContravariant)))))) (box ((box f)))))) (box ((box x))))))))))

let Data_Functor_Contravariant_coerce  = (box (fun (dictContravariant: obj) -> (box (fun (dictFunctor: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box Data_Void_absurd)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_Contravariant_cmap))) (box ((box dictContravariant)))))) (box ((box Data_Void_absurd)))))) (box ((box a)))))))))))))

let Data_Functor_Contravariant_imapC  = (box (fun (dictContravariant: obj) -> (box (fun (v: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_Contravariant_cmap))) (box ((box dictContravariant)))))) (box ((box f))))))))))
