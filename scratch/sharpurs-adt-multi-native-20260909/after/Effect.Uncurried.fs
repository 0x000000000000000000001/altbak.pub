[<AutoOpen>]
module PureScript_Effect_Uncurried

open System
open System.Collections.Generic

module Effect_Uncurried_FFI =
    let mkEffectFn1 = box (fun (f: obj) -> box (fun (a: obj) ->
        let res = sharpurs_apply f a
        sharpurs_apply res null
    ))
    
    let mkEffectFn2 = box (fun (f: obj) -> box (fun (a: obj) -> box (fun (b: obj) ->
        let fa = sharpurs_apply f a
        let fab = sharpurs_apply fa b
        sharpurs_apply fab null
    )))
    
    let mkEffectFn3 = box (fun (f: obj) -> box (fun (a: obj) -> box (fun (b: obj) -> box (fun (c: obj) ->
        let fa = sharpurs_apply f a
        let fab = sharpurs_apply fa b
        let fabc = sharpurs_apply fab c
        sharpurs_apply fabc null
    ))))
    
    let mkEffectFn4 = box (fun (f: obj) -> box (fun (a: obj) -> box (fun (b: obj) -> box (fun (c: obj) -> box (fun (d: obj) ->
        let fa = sharpurs_apply f a
        let fab = sharpurs_apply fa b
        let fabc = sharpurs_apply fab c
        let fabcd = sharpurs_apply fabc d
        sharpurs_apply fabcd null
    )))))
    let mkEffectFn5 _ = undefined
    let mkEffectFn6 _ = undefined
    let mkEffectFn7 _ = undefined
    let mkEffectFn8 _ = undefined
    let mkEffectFn9 _ = undefined
    let mkEffectFn10 _ = undefined
    
    let runEffectFn1 = box (fun (eff: obj) -> box (fun (a: obj) -> box (fun _ ->
        let effT = eff :?> (obj -> obj)
        effT a
    )))
    
    let runEffectFn2 = box (fun (eff: obj) -> box (fun (a: obj) -> box (fun (b: obj) -> box (fun _ ->
        try
            let effT = eff :?> (obj -> (obj -> obj))
            effT a b
        with
        | :? System.InvalidCastException ->
            // The eff is a .NET function that takes 2 arguments or we need to apply manually.
            // But since it's an EffectFn2 from our FFI, it's box (fun a -> box (fun b -> ...))
            // So we CAN cast it to (obj -> obj) to get the inner function!
            let eff1 = eff :?> (obj -> obj)
            let eff2 = eff1 a :?> (obj -> obj)
            eff2 b
    ))))
    
    let runEffectFn3 = box (fun (eff: obj) -> box (fun (a: obj) -> box (fun (b: obj) -> box (fun (c: obj) -> box (fun _ ->
        let eff1 = eff :?> (obj -> obj)
        let eff2 = eff1 a :?> (obj -> obj)
        let eff3 = eff2 b :?> (obj -> obj)
        eff3 c
    )))))
    
    let runEffectFn4 = box (fun (eff: obj) -> box (fun (a: obj) -> box (fun (b: obj) -> box (fun (c: obj) -> box (fun (d: obj) -> box (fun _ ->
        let eff1 = eff :?> (obj -> obj)
        let eff2 = eff1 a :?> (obj -> obj)
        let eff3 = eff2 b :?> (obj -> obj)
        let eff4 = eff3 c :?> (obj -> obj)
        eff4 d
    ))))))
    
    let runEffectFn5 _ = undefined
    let runEffectFn6 _ = undefined
    let runEffectFn7 _ = undefined
    let runEffectFn8 _ = undefined
    let runEffectFn9 _ = undefined
    let runEffectFn10 _ = undefined
    

let Effect_Uncurried_mkEffectFn1 = box (Effect_Uncurried_FFI.``mkEffectFn1``)
let Effect_Uncurried_mkEffectFn10 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``mkEffectFn10`` (unbox arg0)))
let Effect_Uncurried_mkEffectFn2 = box (Effect_Uncurried_FFI.``mkEffectFn2``)
let Effect_Uncurried_mkEffectFn3 = box (Effect_Uncurried_FFI.``mkEffectFn3``)
let Effect_Uncurried_mkEffectFn4 = box (Effect_Uncurried_FFI.``mkEffectFn4``)
let Effect_Uncurried_mkEffectFn5 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``mkEffectFn5`` (unbox arg0)))
let Effect_Uncurried_mkEffectFn6 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``mkEffectFn6`` (unbox arg0)))
let Effect_Uncurried_mkEffectFn7 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``mkEffectFn7`` (unbox arg0)))
let Effect_Uncurried_mkEffectFn8 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``mkEffectFn8`` (unbox arg0)))
let Effect_Uncurried_mkEffectFn9 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``mkEffectFn9`` (unbox arg0)))
let Effect_Uncurried_runEffectFn1 = box (Effect_Uncurried_FFI.``runEffectFn1``)
let Effect_Uncurried_runEffectFn10 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``runEffectFn10`` (unbox arg0)))
let Effect_Uncurried_runEffectFn2 = box (Effect_Uncurried_FFI.``runEffectFn2``)
let Effect_Uncurried_runEffectFn3 = box (Effect_Uncurried_FFI.``runEffectFn3``)
let Effect_Uncurried_runEffectFn4 = box (Effect_Uncurried_FFI.``runEffectFn4``)
let Effect_Uncurried_runEffectFn5 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``runEffectFn5`` (unbox arg0)))
let Effect_Uncurried_runEffectFn6 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``runEffectFn6`` (unbox arg0)))
let Effect_Uncurried_runEffectFn7 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``runEffectFn7`` (unbox arg0)))
let Effect_Uncurried_runEffectFn8 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``runEffectFn8`` (unbox arg0)))
let Effect_Uncurried_runEffectFn9 = box (fun (arg0: obj) -> box (Effect_Uncurried_FFI.``runEffectFn9`` (unbox arg0)))


let Effect_Uncurried_semigroupEffectFn9  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn9))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (box (fun (e: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (h: obj) -> (box (fun (i: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn9))) (box ((box f1)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))) (box ((box g)))))) (box ((box h)))))) (box ((box i))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn9))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))) (box ((box g)))))) (box ((box h)))))) (box ((box i)))))))))))))))))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn8  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn8))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (box (fun (e: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (h: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn8))) (box ((box f1)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))) (box ((box g)))))) (box ((box h))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn8))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))) (box ((box g)))))) (box ((box h)))))))))))))))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn7  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn7))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (box (fun (e: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn7))) (box ((box f1)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))) (box ((box g))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn7))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))) (box ((box g)))))))))))))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn6  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn6))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (box (fun (e: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn6))) (box ((box f1)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn6))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))))))))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn5  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn5))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (box (fun (e: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn5))) (box ((box f1)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn5))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))))))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn4  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn4))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn4))) (box ((box f1)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn4))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn3  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn3))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn3))) (box ((box f1)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn3))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn2  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn2))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn2))) (box ((box f1)))))) (box ((box a)))))) (box ((box b))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn2))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn10  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn10))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (box (fun (e: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (h: obj) -> (box (fun (i: obj) -> (box (fun (j: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn10))) (box ((box f1)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))) (box ((box g)))))) (box ((box h)))))) (box ((box i)))))) (box ((box j))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn10))) (box ((box f2)))))) (box ((box a)))))) (box ((box b)))))) (box ((box c)))))) (box ((box d)))))) (box ((box e)))))) (box ((box f)))))) (box ((box g)))))) (box ((box h)))))) (box ((box i)))))) (box ((box j)))))))))))))))))))))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_semigroupEffectFn1  = (box (fun (dictSemigroup: obj) -> (let semigroupEffect = (sharpurs_apply (box ((box Effect_semigroupEffect))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f1: obj) -> (box (fun (f2: obj) -> (sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn1))) (box ((box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box semigroupEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn1))) (box ((box f1)))))) (box ((box a))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Uncurried_runEffectFn1))) (box ((box f2)))))) (box ((box a)))))))))))))))))) Map.empty)))))))))

let Effect_Uncurried_monoidEffectFn9  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn91 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn9))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn9))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (box (fun (v4: obj) -> (box (fun (v5: obj) -> (box (fun (v6: obj) -> (box (fun (v7: obj) -> (box (fun (v8: obj) -> (box mempty)))))))))))))))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn91))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn8  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn81 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn8))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn8))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (box (fun (v4: obj) -> (box (fun (v5: obj) -> (box (fun (v6: obj) -> (box (fun (v7: obj) -> (box mempty)))))))))))))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn81))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn7  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn71 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn7))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn7))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (box (fun (v4: obj) -> (box (fun (v5: obj) -> (box (fun (v6: obj) -> (box mempty)))))))))))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn71))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn6  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn61 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn6))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn6))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (box (fun (v4: obj) -> (box (fun (v5: obj) -> (box mempty)))))))))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn61))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn5  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn51 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn5))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn5))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (box (fun (v4: obj) -> (box mempty)))))))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn51))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn4  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn41 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn4))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn4))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (box mempty)))))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn41))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn3  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn31 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn3))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn3))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box mempty)))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn31))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn2  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn21 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn2))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn2))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box mempty)))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn21))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn10  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn101 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn10))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn10))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (box (fun (v4: obj) -> (box (fun (v5: obj) -> (box (fun (v6: obj) -> (box (fun (v7: obj) -> (box (fun (v8: obj) -> (box (fun (v9: obj) -> (box mempty)))))))))))))))))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn101))))) Map.empty))))))))))

let Effect_Uncurried_monoidEffectFn1  = (box (fun (dictMonoid: obj) -> (let mempty = (sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((sharpurs_apply (box ((box Effect_monoidEffect))) (box ((box dictMonoid))))))) in let semigroupEffectFn11 = (sharpurs_apply (box ((box Effect_Uncurried_semigroupEffectFn1))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Effect_Uncurried_mkEffectFn1))) (box ((box (fun (v: obj) -> (box mempty)))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEffectFn11))))) Map.empty))))))))))
