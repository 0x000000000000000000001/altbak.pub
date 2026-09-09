type DirectFallback_Color =
  | DirectFallback_Warmusd_Ctor
  | DirectFallback_Coolusd_Ctor

type DirectFallback_Tree =
  | DirectFallback_Tipusd_Ctor
  | DirectFallback_Forkusd_Ctor of obj * obj * obj * obj

let DirectFallback_Warm  = (box DirectFallback_Warmusd_Ctor)

let DirectFallback_Cool  = (box DirectFallback_Coolusd_Ctor)

let DirectFallback_Tip  = (box DirectFallback_Tipusd_Ctor)

let DirectFallback_Fork  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (fun (usd__arg3: obj) -> (fun (usd__arg4: obj) -> (box (DirectFallback_Forkusd_Ctor(usd__arg1, usd__arg2, usd__arg3, usd__arg4)))))))))

let DirectFallback_unary  = (box (fun (x: obj) -> (box x)))

let DirectFallback_stringPair  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectFallback_shadowed  = (box (fun (ordinal1: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ordinal1))) (box ((box 4)))))) (box ((box 3)))))) (box ((box 2)))))) (box ((box 1))))))

let DirectFallback_seed  = (box 17)

let DirectFallback_returned  = (box (fun (flag: obj) -> (box (fun (value: obj) -> (let saved = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 90)))))) (box ((box value)))) in (box (fun (other: obj) -> (match ((unbox ((box flag)))) with | LitBool true () -> ((box saved)) | _ -> ((box other))))))))))

let DirectFallback_return_direct  = (box 7)

let DirectFallback_return  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectFallback_ordinal  = (box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (match ((unbox ((box ((unbox<int> (box ((box a)))) < (unbox<int> (box ((box b))))))))) with | LitBool true () -> ((box c)) | _ -> ((box d)))))))))))

let DirectFallback_partial1  = (sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box 1))))

let DirectFallback_partial2  = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box 1)))))) (box ((box 2))))

let DirectFallback_partial3  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box 1)))))) (box ((box 2)))))) (box ((box 3))))

let DirectFallback_saturated  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box x)))))) (box ((box 10)))))) (box ((box 20)))))) (box ((box 30))))))

let DirectFallback_orderedPartial  = (box (fun (x: obj) -> (sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 1)))))) (box ((box x)))))))))

let DirectFallback_ordered  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 1)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 2)))))) (box ((box 10))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 3)))))) (box ((box 20))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 4)))))) (box ((box 30)))))))))

let DirectFallback_numberPair  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectFallback_localShadow  = (box (fun (x: obj) -> (let ordinal1 = (box (fun (a: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box a))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ordinal1))) (box ((box x)))))) (box ((box 2)))))) (box ((box 3)))))) (box ((box 4)))))))

let DirectFallback_imported  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box DirectRemote_remote))) (box ((box x)))))) (box ((box y))))))))

let DirectFallback_higher  = (box (fun (fn: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box fn))) (box ((box x))))))))

let DirectFallback_guarded  = (box (fun (first: obj) -> (box (fun (x: obj) -> (box (fun (second: obj) -> (box (fun (y: obj) -> (match ((unbox ((box first)))) with | LitBool true () -> ((box x)) | _ -> ((match ((unbox ((box second)))) with | LitBool true () -> ((box y)) | _ -> ((box DirectFallback_seed)))))))))))))

let DirectFallback_generic  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectFallback_failBody  = (box (fun (x: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((box DirectFallback_explode))) (box ((box x))))))))

let DirectFallback_captured  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box x)))))) (box ((box y)))))) (box ((box DirectFallback_seed)))))) (box ((box 29))))))))

let DirectFallback_bodyCall  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_failBody))) (box ((box x)))))) (box ((box 0))))))

let DirectFallback_asValue  = (box DirectFallback_ordinal)

let DirectFallback_arrange  = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (box (DirectFallback_Forkusd_Ctor((box color), (box left), (box value), (box right))))))))))))

let DirectFallback_argumentCall  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((sharpurs_apply (box ((box DirectFallback_explode))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 2)))))) (box ((box 10))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 3)))))) (box ((box 20))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 4)))))) (box ((box 30)))))))))