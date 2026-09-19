[<AutoOpen>]
module PureScript_Control_Monad

open System
open System.Collections.Generic

let Control_Monad_Monadusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Monad_whenM  = (box (fun (dictMonad: obj) -> (let Bind1 = (sharpurs_apply (box ((Map.find "Bind1" (unbox<Map<string, obj>> ((box dictMonad)))))) (box ((box Prim_undefined)))) in let Applicative0 = (sharpurs_apply (box ((Map.find "Applicative0" (unbox<Map<string, obj>> ((box dictMonad)))))) (box ((box Prim_undefined)))) in (box (fun (mb: obj) -> (box (fun (m: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Bind1)))))) (box ((box mb)))))) (box ((box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_when))) (box ((box Applicative0)))))) (box ((box b)))))) (box ((box m))))))))))))))))

let Control_Monad_unlessM  = (box (fun (dictMonad: obj) -> (let Bind1 = (sharpurs_apply (box ((Map.find "Bind1" (unbox<Map<string, obj>> ((box dictMonad)))))) (box ((box Prim_undefined)))) in let Applicative0 = (sharpurs_apply (box ((Map.find "Applicative0" (unbox<Map<string, obj>> ((box dictMonad)))))) (box ((box Prim_undefined)))) in (box (fun (mb: obj) -> (box (fun (m: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Bind1)))))) (box ((box mb)))))) (box ((box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_unless))) (box ((box Applicative0)))))) (box ((box b)))))) (box ((box m))))))))))))))))

let Control_Monad_monadProxy  = (sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Control_Applicative_applicativeProxy))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Control_Bind_bindProxy))))) Map.empty)))))))

let Control_Monad_monadFn  = (sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Control_Applicative_applicativeFn))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Control_Bind_bindFn))))) Map.empty)))))))

let Control_Monad_monadArray  = (sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Control_Applicative_applicativeArray))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Control_Bind_bindArray))))) Map.empty)))))))

let Control_Monad_liftM1  = (box (fun (dictMonad: obj) -> (let Bind1 = (sharpurs_apply (box ((Map.find "Bind1" (unbox<Map<string, obj>> ((box dictMonad)))))) (box ((box Prim_undefined)))) in let Applicative0 = (sharpurs_apply (box ((Map.find "Applicative0" (unbox<Map<string, obj>> ((box dictMonad)))))) (box ((box Prim_undefined)))) in (box (fun (f: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Bind1)))))) (box ((box a)))))) (box ((box (fun (a_prime: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Applicative0)))))) (box ((sharpurs_apply (box ((box f))) (box ((box a_prime)))))))))))))))))))

let Control_Monad_ap  = (box (fun (dictMonad: obj) -> (let Bind1 = (sharpurs_apply (box ((Map.find "Bind1" (unbox<Map<string, obj>> ((box dictMonad)))))) (box ((box Prim_undefined)))) in let Applicative0 = (sharpurs_apply (box ((Map.find "Applicative0" (unbox<Map<string, obj>> ((box dictMonad)))))) (box ((box Prim_undefined)))) in (box (fun (f: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Bind1)))))) (box ((box f)))))) (box ((box (fun (f_prime: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Bind1)))))) (box ((box a)))))) (box ((box (fun (a_prime: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Applicative0)))))) (box ((sharpurs_apply (box ((box f_prime))) (box ((box a_prime))))))))))))))))))))))))
