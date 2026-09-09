type DirectCall_Color =
  | DirectCall_Warmusd_Ctor
  | DirectCall_Coolusd_Ctor

type DirectCall_Tree =
  | DirectCall_Tipusd_Ctor
  | DirectCall_Forkusd_Ctor of obj * obj * obj * obj

let DirectCall_Warm  = (box DirectCall_Warmusd_Ctor)

let DirectCall_Cool  = (box DirectCall_Coolusd_Ctor)

let DirectCall_Tip  = (box DirectCall_Tipusd_Ctor)

let DirectCall_Fork  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (fun (usd__arg3: obj) -> (fun (usd__arg4: obj) -> (box (DirectCall_Forkusd_Ctor(usd__arg1, usd__arg2, usd__arg3, usd__arg4)))))))))

let DirectCall_unary  = (box (fun (x: obj) -> (box x)))

let DirectCall_stringPair  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectCall_shadowed  = (box (fun (ordinal1: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ordinal1))) (box ((box 4)))))) (box ((box 3)))))) (box ((box 2)))))) (box ((box 1))))))

let DirectCall_seed  = (box 17)

let DirectCall_returned  = (box (fun (flag: obj) -> (box (fun (value: obj) -> (let saved = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 90)))))) (box ((box value)))) in (box (fun (other: obj) -> (match ((unbox ((box flag)))) with | LitBool true () -> ((box saved)) | _ -> ((box other))))))))))

let DirectCall_return_direct  = (box 7)

let DirectCall_return  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectCall_ordinal_direct (a: obj) (b: obj) (c: obj) (d: obj) : obj = (match ((unbox ((box ((unbox<int> (box ((box a)))) < (unbox<int> (box ((box b))))))))) with | LitBool true () -> ((box c)) | _ -> ((box d)))

let DirectCall_ordinal_direct_apply (a: obj) (b: obj) (c: obj) (d: obj) : obj =
    try DirectCall_ordinal_direct a b c d
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_ordinal = (box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (DirectCall_ordinal_direct a b c d)))))))))

let DirectCall_partial1  = (sharpurs_apply (box ((box DirectCall_ordinal))) (box ((box 1))))

let DirectCall_partial2  = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_ordinal))) (box ((box 1)))))) (box ((box 2))))

let DirectCall_partial3  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_ordinal))) (box ((box 1)))))) (box ((box 2)))))) (box ((box 3))))

let DirectCall_saturated  = (box (fun (x: obj) -> (DirectCall_ordinal_direct_apply ((box ((box x)))) ((box ((box 10)))) ((box ((box 20)))) ((box ((box 30)))))))

let DirectCall_orderedPartial  = (box (fun (x: obj) -> (sharpurs_apply (box ((box DirectCall_ordinal))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 1)))))) (box ((box x)))))))))

let DirectCall_ordered  = (box (fun (x: obj) -> (DirectCall_ordinal_direct_apply ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 1)))))) (box ((box x))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 2)))))) (box ((box 10))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 3)))))) (box ((box 20))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 4)))))) (box ((box 30))))))))))

let DirectCall_numberPair  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectCall_localShadow  = (box (fun (x: obj) -> (let ordinal1 = (box (fun (a: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box a))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ordinal1))) (box ((box x)))))) (box ((box 2)))))) (box ((box 3)))))) (box ((box 4)))))))

let DirectCall_imported_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectRemote_remote))) (box ((box x)))))) (box ((box y))))

let DirectCall_imported_direct_apply (x: obj) (y: obj) : obj =
    try DirectCall_imported_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_imported = (box (fun (x: obj) -> (box (fun (y: obj) -> (DirectCall_imported_direct x y)))))

let DirectCall_higher  = (box (fun (fn: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box fn))) (box ((box x))))))))

let DirectCall_guarded_direct (first: obj) (x: obj) (second: obj) (y: obj) : obj = (match ((unbox ((box first)))) with | LitBool true () -> ((box x)) | _ -> ((match ((unbox ((box second)))) with | LitBool true () -> ((box y)) | _ -> ((box DirectCall_seed)))))

let DirectCall_guarded_direct_apply (first: obj) (x: obj) (second: obj) (y: obj) : obj =
    try DirectCall_guarded_direct first x second y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_guarded = (box (fun (first: obj) -> (box (fun (x: obj) -> (box (fun (second: obj) -> (box (fun (y: obj) -> (DirectCall_guarded_direct first x second y)))))))))

let DirectCall_generic  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectCall_failBody_direct (x: obj) (v: obj) : obj = (sharpurs_apply (box ((box DirectCall_explode))) (box ((box x))))

let DirectCall_failBody_direct_apply (x: obj) (v: obj) : obj =
    try DirectCall_failBody_direct x v
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_failBody = (box (fun (x: obj) -> (box (fun (v: obj) -> (DirectCall_failBody_direct x v)))))

let DirectCall_captured_direct (x: obj) (y: obj) : obj = (DirectCall_ordinal_direct_apply ((box ((box x)))) ((box ((box y)))) ((box ((box DirectCall_seed)))) ((box ((box 29)))))

let DirectCall_captured_direct_apply (x: obj) (y: obj) : obj =
    try DirectCall_captured_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_captured = (box (fun (x: obj) -> (box (fun (y: obj) -> (DirectCall_captured_direct x y)))))

let DirectCall_bodyCall  = (box (fun (x: obj) -> (DirectCall_failBody_direct_apply ((box ((box x)))) ((box ((box 0)))))))

let DirectCall_asValue  = (box DirectCall_ordinal)

let DirectCall_arrange  = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (box (DirectCall_Forkusd_Ctor((box color), (box left), (box value), (box right))))))))))))

let DirectCall_argumentCall  = (box (fun (x: obj) -> (DirectCall_ordinal_direct_apply ((box ((sharpurs_apply (box ((box DirectCall_explode))) (box ((box x))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 2)))))) (box ((box 10))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 3)))))) (box ((box 20))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 4)))))) (box ((box 30))))))))))