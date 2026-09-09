let objMap = Map.empty<string, obj>
let unbox<'a> (x: obj) : 'a = unbox x
let (|Unbox|) (x: obj) = unbox x

let undefined = Unchecked.defaultof<obj>
let Prim_undefined = undefined
let intMod a b = unbox<int> a % unbox<int> b
let semiringInt = 0

// PureScript's Euclidean Int modulo, including zero and Int32.MinValue / -1.
let sharpurs_int_mod (left: int) (right: int) : int =
    if right = 0 || right = -1 then 0
    else
        let remainder = left % right
        if remainder < 0 then
            if right > 0 then remainder + right else remainder - right
        else remainder

let sharpurs_apply (func: obj) (arg: obj) : obj =
    if isNull func then failwith "sharpurs_apply: func is null!"
    match func with
    | :? (obj -> obj) as invoke ->
        try invoke arg
        // Keep the exception boundary of MethodInfo.Invoke for callers and FFI.
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    | _ ->
        let method = func.GetType().GetMethods() |> Array.find (fun m -> m.Name = "Invoke" && m.GetParameters().Length = 1)
        method.Invoke(func, [| arg |])



let mutable fallbackCalls = 0
let events = ResizeArray<int>()
let failure = System.InvalidOperationException("int-arithmetic fixture failure")
let mutable currentValue = 0
let mutable reads = 0
let IntArithmetic_track : obj = box (fun (label: obj) -> box (fun (value: obj) -> events.Add(unbox<int> label); value))
let IntArithmetic_explode : obj = box (fun (_: obj) -> events.Add(99); raise failure : obj)
let IntArithmetic_readCurrent : obj = box (fun (_: obj) -> reads <- reads + 1; box currentValue)
let ArithmeticFallback_track = IntArithmetic_track
let ArithmeticFallback_explode = IntArithmetic_explode
let ArithmeticFallback_readCurrent = IntArithmetic_readCurrent
let Data_Unit_unit = box ()
let binary operation : obj = box (fun (x: obj) -> box (fun (y: obj) ->
    fallbackCalls <- fallbackCalls + 1
    operation x y))
let intAdd = binary (fun x y -> box (unbox<int> x + unbox<int> y))
let intSub = binary (fun x y -> box (unbox<int> x - unbox<int> y))
let numberAdd = binary (fun x y -> box (unbox<float> x + unbox<float> y))
let numberSub = binary (fun x y -> box (unbox<float> x - unbox<float> y))
let Data_Semiring_semiringInt : obj = box (Map.ofList ["add", intAdd])
let Data_Ring_ringInt : obj = box (Map.ofList ["sub", intSub; "Semiring0", box (fun (_: obj) -> Data_Semiring_semiringInt)])
let Data_Semiring_semiringNumber : obj = box (Map.ofList ["add", numberAdd])
let Data_Ring_ringNumber : obj = box (Map.ofList ["sub", numberSub; "Semiring0", box (fun (_: obj) -> Data_Semiring_semiringNumber)])
let Data_Semiring_add : obj = box (fun (dict: obj) -> Map.find "add" (unbox<Map<string,obj>> dict))
let Data_Ring_sub : obj = box (fun (dict: obj) -> Map.find "sub" (unbox<Map<string,obj>> dict))
// Deliberately noncanonical dictionaries exercise object-ABI dispatch rather
// than deriving semantics from the types of their Int operands.
let customSemiring : obj = box (Map.ofList ["add", binary (fun x y -> box (unbox<int> y - unbox<int> x))])
let customRing : obj = box (Map.ofList ["sub", intAdd])


let IntArithmetic_sub  = (sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt))))

let IntArithmetic_add  = (sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt))))

let IntArithmetic_visibleSub_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((box x)))))) (box ((box y))))

let IntArithmetic_visibleSub_direct_apply (x: obj) (y: obj) : obj =
    try IntArithmetic_visibleSub_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_visibleSub = (box (fun (x: obj) -> (box (fun (y: obj) -> (IntArithmetic_visibleSub_direct x y)))))

let IntArithmetic_visibleAdd_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box x)))))) (box ((box y))))

let IntArithmetic_visibleAdd_direct_apply (x: obj) (y: obj) : obj =
    try IntArithmetic_visibleAdd_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_visibleAdd = (box (fun (x: obj) -> (box (fun (y: obj) -> (IntArithmetic_visibleAdd_direct x y)))))

let IntArithmetic_subInt_direct (x: obj) (y: obj) : obj = (box ((unbox<int> (box ((box x)))) - (unbox<int> (box ((box y))))))

let IntArithmetic_subInt_direct_apply (x: obj) (y: obj) : obj =
    try IntArithmetic_subInt_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_subInt = (box (fun (x: obj) -> (box (fun (y: obj) -> (IntArithmetic_subInt_direct x y)))))

let IntArithmetic_secondFailureSub  = (box (fun (x: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 1)))))) (box ((box x))))))) - (unbox<int> (box ((sharpurs_apply (box ((box IntArithmetic_explode))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 2)))))) (box ((box x))))))))))))))

let IntArithmetic_secondFailureAdd  = (box (fun (x: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 1)))))) (box ((box x))))))) + (unbox<int> (box ((sharpurs_apply (box ((box IntArithmetic_explode))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 2)))))) (box ((box x))))))))))))))

let IntArithmetic_partialSub  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((box 5))))

let IntArithmetic_partialAdd  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box 5))))

let IntArithmetic_orderedSub_direct (x: obj) (y: obj) : obj = (box ((unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 1)))))) (box ((box x))))))) - (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 2)))))) (box ((box y)))))))))

let IntArithmetic_orderedSub_direct_apply (x: obj) (y: obj) : obj =
    try IntArithmetic_orderedSub_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_orderedSub = (box (fun (x: obj) -> (box (fun (y: obj) -> (IntArithmetic_orderedSub_direct x y)))))

let IntArithmetic_orderedAdd_direct (x: obj) (y: obj) : obj = (box ((unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 1)))))) (box ((box x))))))) + (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 2)))))) (box ((box y)))))))))

let IntArithmetic_orderedAdd_direct_apply (x: obj) (y: obj) : obj =
    try IntArithmetic_orderedAdd_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_orderedAdd = (box (fun (x: obj) -> (box (fun (y: obj) -> (IntArithmetic_orderedAdd_direct x y)))))

let IntArithmetic_numberSub  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringNumber)))))) (box ((box x)))))) (box ((box y))))))))

let IntArithmetic_numberAdd  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box x)))))) (box ((box y))))))))

let IntArithmetic_genericSub  = (box (fun (dictRing: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box dictRing)))))) (box ((box x)))))) (box ((box y))))))))))

let IntArithmetic_genericAdd  = (box (fun (dictSemiring: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box dictSemiring)))))) (box ((box x)))))) (box ((box y))))))))))

let IntArithmetic_firstFailureSub  = (box (fun (x: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((box IntArithmetic_explode))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 1)))))) (box ((box x)))))))))) - (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 2)))))) (box ((box x)))))))))))

let IntArithmetic_firstFailureAdd  = (box (fun (x: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((box IntArithmetic_explode))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 1)))))) (box ((box x)))))))))) + (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_track))) (box ((box 2)))))) (box ((box x)))))))))))

let IntArithmetic_delayedSub  = (box (fun (offset: obj) -> (box (fun (v: obj) -> (box ((unbox<int> (box ((box offset)))) - (unbox<int> (box ((sharpurs_apply (box ((box IntArithmetic_readCurrent))) (box ((box Data_Unit_unit)))))))))))))

let IntArithmetic_delayedAdd  = (box (fun (offset: obj) -> (box (fun (v: obj) -> (box ((unbox<int> (box ((box offset)))) + (unbox<int> (box ((sharpurs_apply (box ((box IntArithmetic_readCurrent))) (box ((box Data_Unit_unit)))))))))))))

let IntArithmetic_capturedSub_direct (offset: obj) (value: obj) : obj = (box ((unbox<int> (box ((box offset)))) - (unbox<int> (box ((box value))))))

let IntArithmetic_capturedSub_direct_apply (offset: obj) (value: obj) : obj =
    try IntArithmetic_capturedSub_direct offset value
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_capturedSub = (box (fun (offset: obj) -> (box (fun (value: obj) -> (IntArithmetic_capturedSub_direct offset value)))))

let IntArithmetic_capturedAdd_direct (offset: obj) (value: obj) : obj = (box ((unbox<int> (box ((box offset)))) + (unbox<int> (box ((box value))))))

let IntArithmetic_capturedAdd_direct_apply (offset: obj) (value: obj) : obj =
    try IntArithmetic_capturedAdd_direct offset value
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_capturedAdd = (box (fun (offset: obj) -> (box (fun (value: obj) -> (IntArithmetic_capturedAdd_direct offset value)))))

let IntArithmetic_annotatedSub_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_sub))) (box ((box x)))))) (box ((box y))))

let IntArithmetic_annotatedSub_direct_apply (x: obj) (y: obj) : obj =
    try IntArithmetic_annotatedSub_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_annotatedSub = (box (fun (x: obj) -> (box (fun (y: obj) -> (IntArithmetic_annotatedSub_direct x y)))))

let IntArithmetic_annotatedAdd_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((box IntArithmetic_add))) (box ((box x)))))) (box ((box y))))

let IntArithmetic_annotatedAdd_direct_apply (x: obj) (y: obj) : obj =
    try IntArithmetic_annotatedAdd_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_annotatedAdd = (box (fun (x: obj) -> (box (fun (y: obj) -> (IntArithmetic_annotatedAdd_direct x y)))))

let IntArithmetic_addInt_direct (x: obj) (y: obj) : obj = (box ((unbox<int> (box ((box x)))) + (unbox<int> (box ((box y))))))

let IntArithmetic_addInt_direct_apply (x: obj) (y: obj) : obj =
    try IntArithmetic_addInt_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let IntArithmetic_addInt = (box (fun (x: obj) -> (box (fun (y: obj) -> (IntArithmetic_addInt_direct x y)))))

let ArithmeticFallback_sub  = (sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt))))

let ArithmeticFallback_add  = (sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt))))

let ArithmeticFallback_visibleSub_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((box x)))))) (box ((box y))))

let ArithmeticFallback_visibleSub_direct_apply (x: obj) (y: obj) : obj =
    try ArithmeticFallback_visibleSub_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_visibleSub = (box (fun (x: obj) -> (box (fun (y: obj) -> (ArithmeticFallback_visibleSub_direct x y)))))

let ArithmeticFallback_visibleAdd_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box x)))))) (box ((box y))))

let ArithmeticFallback_visibleAdd_direct_apply (x: obj) (y: obj) : obj =
    try ArithmeticFallback_visibleAdd_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_visibleAdd = (box (fun (x: obj) -> (box (fun (y: obj) -> (ArithmeticFallback_visibleAdd_direct x y)))))

let ArithmeticFallback_subInt_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((box x)))))) (box ((box y))))

let ArithmeticFallback_subInt_direct_apply (x: obj) (y: obj) : obj =
    try ArithmeticFallback_subInt_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_subInt = (box (fun (x: obj) -> (box (fun (y: obj) -> (ArithmeticFallback_subInt_direct x y)))))

let ArithmeticFallback_secondFailureSub  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 1)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box ArithmeticFallback_explode))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 2)))))) (box ((box x))))))))))))

let ArithmeticFallback_secondFailureAdd  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 1)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box ArithmeticFallback_explode))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 2)))))) (box ((box x))))))))))))

let ArithmeticFallback_partialSub  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((box 5))))

let ArithmeticFallback_partialAdd  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box 5))))

let ArithmeticFallback_orderedSub_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 1)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 2)))))) (box ((box y)))))))

let ArithmeticFallback_orderedSub_direct_apply (x: obj) (y: obj) : obj =
    try ArithmeticFallback_orderedSub_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_orderedSub = (box (fun (x: obj) -> (box (fun (y: obj) -> (ArithmeticFallback_orderedSub_direct x y)))))

let ArithmeticFallback_orderedAdd_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 1)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 2)))))) (box ((box y)))))))

let ArithmeticFallback_orderedAdd_direct_apply (x: obj) (y: obj) : obj =
    try ArithmeticFallback_orderedAdd_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_orderedAdd = (box (fun (x: obj) -> (box (fun (y: obj) -> (ArithmeticFallback_orderedAdd_direct x y)))))

let ArithmeticFallback_numberSub  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringNumber)))))) (box ((box x)))))) (box ((box y))))))))

let ArithmeticFallback_numberAdd  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box x)))))) (box ((box y))))))))

let ArithmeticFallback_genericSub  = (box (fun (dictRing: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box dictRing)))))) (box ((box x)))))) (box ((box y))))))))))

let ArithmeticFallback_genericAdd  = (box (fun (dictSemiring: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box dictSemiring)))))) (box ((box x)))))) (box ((box y))))))))))

let ArithmeticFallback_firstFailureSub  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((sharpurs_apply (box ((box ArithmeticFallback_explode))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 1)))))) (box ((box x)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 2)))))) (box ((box x)))))))))

let ArithmeticFallback_firstFailureAdd  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((sharpurs_apply (box ((box ArithmeticFallback_explode))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 1)))))) (box ((box x)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_track))) (box ((box 2)))))) (box ((box x)))))))))

let ArithmeticFallback_delayedSub  = (box (fun (offset: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((box offset)))))) (box ((sharpurs_apply (box ((box ArithmeticFallback_readCurrent))) (box ((box Data_Unit_unit)))))))))))

let ArithmeticFallback_delayedAdd  = (box (fun (offset: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box offset)))))) (box ((sharpurs_apply (box ((box ArithmeticFallback_readCurrent))) (box ((box Data_Unit_unit)))))))))))

let ArithmeticFallback_capturedSub_direct (offset: obj) (value: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((box offset)))))) (box ((box value))))

let ArithmeticFallback_capturedSub_direct_apply (offset: obj) (value: obj) : obj =
    try ArithmeticFallback_capturedSub_direct offset value
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_capturedSub = (box (fun (offset: obj) -> (box (fun (value: obj) -> (ArithmeticFallback_capturedSub_direct offset value)))))

let ArithmeticFallback_capturedAdd_direct (offset: obj) (value: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box offset)))))) (box ((box value))))

let ArithmeticFallback_capturedAdd_direct_apply (offset: obj) (value: obj) : obj =
    try ArithmeticFallback_capturedAdd_direct offset value
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_capturedAdd = (box (fun (offset: obj) -> (box (fun (value: obj) -> (ArithmeticFallback_capturedAdd_direct offset value)))))

let ArithmeticFallback_annotatedSub_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_sub))) (box ((box x)))))) (box ((box y))))

let ArithmeticFallback_annotatedSub_direct_apply (x: obj) (y: obj) : obj =
    try ArithmeticFallback_annotatedSub_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_annotatedSub = (box (fun (x: obj) -> (box (fun (y: obj) -> (ArithmeticFallback_annotatedSub_direct x y)))))

let ArithmeticFallback_annotatedAdd_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((box ArithmeticFallback_add))) (box ((box x)))))) (box ((box y))))

let ArithmeticFallback_annotatedAdd_direct_apply (x: obj) (y: obj) : obj =
    try ArithmeticFallback_annotatedAdd_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_annotatedAdd = (box (fun (x: obj) -> (box (fun (y: obj) -> (ArithmeticFallback_annotatedAdd_direct x y)))))

let ArithmeticFallback_addInt_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box x)))))) (box ((box y))))

let ArithmeticFallback_addInt_direct_apply (x: obj) (y: obj) : obj =
    try ArithmeticFallback_addInt_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let ArithmeticFallback_addInt = (box (fun (x: obj) -> (box (fun (y: obj) -> (ArithmeticFallback_addInt_direct x y)))))


let mutable checks = 0
let check label condition = if not condition then failwith label else checks <- checks + 1
let apply = sharpurs_apply
let call fn x y = apply (apply fn (box x)) (box y) |> unbox<int>
let cases = [
    (-2147483648, -2147483648, 0, 0);
    (-2147483648, -2147483647, 1, -1);
    (-2147483648, -65536, 2147418112, -2147418112);
    (-2147483648, -1, 2147483647, -2147483647);
    (-2147483648, 0, -2147483648, -2147483648);
    (-2147483648, 1, -2147483647, 2147483647);
    (-2147483648, 65536, -2147418112, 2147418112);
    (-2147483648, 2147483646, -2, 2);
    (-2147483648, 2147483647, -1, 1);
    (-2147483647, -2147483648, 1, 1);
    (-2147483647, -2147483647, 2, 0);
    (-2147483647, -65536, 2147418113, -2147418111);
    (-2147483647, -1, -2147483648, -2147483646);
    (-2147483647, 0, -2147483647, -2147483647);
    (-2147483647, 1, -2147483646, -2147483648);
    (-2147483647, 65536, -2147418111, 2147418113);
    (-2147483647, 2147483646, -1, 3);
    (-2147483647, 2147483647, 0, 2);
    (-65536, -2147483648, 2147418112, 2147418112);
    (-65536, -2147483647, 2147418113, 2147418111);
    (-65536, -65536, -131072, 0);
    (-65536, -1, -65537, -65535);
    (-65536, 0, -65536, -65536);
    (-65536, 1, -65535, -65537);
    (-65536, 65536, 0, -131072);
    (-65536, 2147483646, 2147418110, 2147418114);
    (-65536, 2147483647, 2147418111, 2147418113);
    (-1, -2147483648, 2147483647, 2147483647);
    (-1, -2147483647, -2147483648, 2147483646);
    (-1, -65536, -65537, 65535);
    (-1, -1, -2, 0);
    (-1, 0, -1, -1);
    (-1, 1, 0, -2);
    (-1, 65536, 65535, -65537);
    (-1, 2147483646, 2147483645, -2147483647);
    (-1, 2147483647, 2147483646, -2147483648);
    (0, -2147483648, -2147483648, -2147483648);
    (0, -2147483647, -2147483647, 2147483647);
    (0, -65536, -65536, 65536);
    (0, -1, -1, 1);
    (0, 0, 0, 0);
    (0, 1, 1, -1);
    (0, 65536, 65536, -65536);
    (0, 2147483646, 2147483646, -2147483646);
    (0, 2147483647, 2147483647, -2147483647);
    (1, -2147483648, -2147483647, -2147483647);
    (1, -2147483647, -2147483646, -2147483648);
    (1, -65536, -65535, 65537);
    (1, -1, 0, 2);
    (1, 0, 1, 1);
    (1, 1, 2, 0);
    (1, 65536, 65537, -65535);
    (1, 2147483646, 2147483647, -2147483645);
    (1, 2147483647, -2147483648, -2147483646);
    (65536, -2147483648, -2147418112, -2147418112);
    (65536, -2147483647, -2147418111, -2147418113);
    (65536, -65536, 0, 131072);
    (65536, -1, 65535, 65537);
    (65536, 0, 65536, 65536);
    (65536, 1, 65537, 65535);
    (65536, 65536, 131072, 0);
    (65536, 2147483646, -2147418114, -2147418110);
    (65536, 2147483647, -2147418113, -2147418111);
    (2147483646, -2147483648, -2, -2);
    (2147483646, -2147483647, -1, -3);
    (2147483646, -65536, 2147418110, -2147418114);
    (2147483646, -1, 2147483645, 2147483647);
    (2147483646, 0, 2147483646, 2147483646);
    (2147483646, 1, 2147483647, 2147483645);
    (2147483646, 65536, -2147418114, 2147418110);
    (2147483646, 2147483646, -4, 0);
    (2147483646, 2147483647, -3, -1);
    (2147483647, -2147483648, -1, -1);
    (2147483647, -2147483647, 0, -2);
    (2147483647, -65536, 2147418111, -2147418113);
    (2147483647, -1, 2147483646, -2147483648);
    (2147483647, 0, 2147483647, 2147483647);
    (2147483647, 1, -2147483648, 2147483646);
    (2147483647, 65536, -2147418113, 2147418111);
    (2147483647, 2147483646, -3, 1);
    (2147483647, 2147483647, -2, 0);
    (286798192, -1840507377, -1553709185, 2127305569);
    (579669026, 1606425881, -2108872389, -1026756855);
    (2084135844, 605051571, -1605779881, 1479084273);
    (1893851254, -488874659, 1404976595, -1912241383);
    (-1394098408, 61814679, -1332283729, -1455913087);
    (-1439043574, 1744707553, 305663979, 1111216169);
    (-503105588, 1908412603, 1405307015, 1883449105);
    (740118238, -620336987, 119781251, 1360455225);
    (563151296, -482440673, 80710623, 1045591969);
    (782979314, -989561943, -206582629, 1772541257);
    (443519220, 799967171, 1243486391, -356447951);
    (1953401414, -1309025043, 644376371, -1032540839);
    (1732942184, 796929447, -1765095665, 936012737);
    (762367706, -1327607695, -565239989, 2089975401);
    (-1496379620, -1363943477, 1434644199, -132436143);
    (-1584450898, 1910757941, 326307043, 799758457);
    (488580624, -675397073, -186816449, 1163977697);
    (1980631490, -956645831, 1023985659, -1357689975);
    (-762949052, 2058107091, 1295158039, 1473911153);
    (-1569921002, -1969913731, 755132563, 399992729);
    (1668106168, -729356361, 938749807, -1897504767);
    (1752867242, 816637185, -1725462869, 936230057);
    (-175244692, -1158171941, -1333416633, 982927249);
    (-475444610, -2046562363, 1772960323, 1571117753);
    (107454048, 1345076799, 1452530847, -1237622751);
    (60962450, 1638655177, 1699617627, -1577692727);
    (116565908, -1560646173, -1444080265, 1677212081);
    (102365670, 1288194061, 1390559731, -1185828391);
    (-124466680, -549165625, -673632305, 424698945);
    (-1008441222, 1687346577, 678905355, 1599179497);
    (2136275388, -224622101, 1911653287, -1934069807);
    (1699255886, -90199723, 1609056163, 1789455609)
]
for x, y, sum, difference in cases do
    let before = fallbackCalls
    check "native sum matches JS" (call IntArithmetic_addInt x y = sum)
    check "native difference matches JS" (call IntArithmetic_subInt x y = difference)
    check "native Int operations bypass dictionaries" (fallbackCalls = before)
    check "generic oracle sum" (call ArithmeticFallback_addInt x y = sum)
    check "generic oracle difference" (call ArithmeticFallback_subInt x y = difference)
    check "oracle uses actual dictionary calls" (fallbackCalls = before + 2)
    check "visible TypeApp add fallback" (call IntArithmetic_visibleAdd x y = sum)
    check "visible TypeApp sub fallback" (call IntArithmetic_visibleSub x y = difference)
    check "generic Int add" (call (apply IntArithmetic_genericAdd Data_Semiring_semiringInt) x y = sum)
    check "generic Int sub" (call (apply IntArithmetic_genericSub Data_Ring_ringInt) x y = difference)
    check "annotated add alias" (call IntArithmetic_annotatedAdd x y = sum)
    check "annotated sub alias" (call IntArithmetic_annotatedSub x y = difference)
    check "custom Int semiring retains supplied operation" (call (apply IntArithmetic_genericAdd customSemiring) x y = (y - x))
    check "custom Int ring retains supplied operation" (call (apply IntArithmetic_genericSub customRing) x y = sum)
for y in [System.Int32.MinValue; -1; 0; 1; System.Int32.MaxValue] do
    check "partial add reusable" (unbox<int> (apply IntArithmetic_partialAdd (box y)) = 5 + y)
    check "partial sub reusable" (unbox<int> (apply IntArithmetic_partialSub (box y)) = 5 - y)
let capturedAdd = apply IntArithmetic_capturedAdd (box System.Int32.MaxValue)
let capturedSub = apply IntArithmetic_capturedSub (box System.Int32.MinValue)
for y in [-1; 0; 1; 7] do
    check "captured addition reused" (unbox<int> (apply capturedAdd (box y)) = System.Int32.MaxValue + y)
    check "captured subtraction reused" (unbox<int> (apply capturedSub (box y)) = System.Int32.MinValue - y)
let floatCall fn x y = apply (apply fn (box x)) (box y) |> unbox<float>
check "Number addition fallback" (floatCall IntArithmetic_numberAdd 1.25 1.5 = 2.75)
check "Number subtraction fallback" (floatCall IntArithmetic_numberSub 1.25 1.5 = -0.25)
for fn, expected in [IntArithmetic_orderedAdd, 10; ArithmeticFallback_orderedAdd, 10;
                     IntArithmetic_orderedSub, -4; ArithmeticFallback_orderedSub, -4] do
    events.Clear()
    check "ordered arithmetic result" (call fn 3 7 = expected)
    check "operands evaluate exactly once left to right" (List.ofSeq events = [1; 2])
for fn, reference in [IntArithmetic_delayedAdd, ArithmeticFallback_delayedAdd;
                      IntArithmetic_delayedSub, ArithmeticFallback_delayedSub] do
    reads <- 0
    currentValue <- 10
    let delayed = apply fn (box 7)
    let oracle = apply reference (box 7)
    check "constructing closures does not force reads" (reads = 0)
    for value in [11; -1; System.Int32.MinValue; System.Int32.MaxValue] do
        currentValue <- value
        let before = reads
        let actual = apply delayed Data_Unit_unit
        check "each force reads once" (reads = before + 1)
        check "reused closure sees latest state" (actual = apply oracle Data_Unit_unit)
        check "oracle force reads once" (reads = before + 2)
let captured action = try action() |> ignore; failwith "expected failure" with ex -> ex
let rec chain (ex: System.Exception) =
    match ex with
    | :? System.Reflection.TargetInvocationException as wrapper -> 1 + chain wrapper.InnerException
    | cause when System.Object.ReferenceEquals(cause, failure) -> 0
    | cause -> failwithf "Unexpected exception: %A" cause
for fn, reference, expected in [IntArithmetic_firstFailureAdd, ArithmeticFallback_firstFailureAdd, [1; 99];
                                IntArithmetic_secondFailureAdd, ArithmeticFallback_secondFailureAdd, [1; 2; 99];
                                IntArithmetic_firstFailureSub, ArithmeticFallback_firstFailureSub, [1; 99];
                                IntArithmetic_secondFailureSub, ArithmeticFallback_secondFailureSub, [1; 2; 99]] do
    events.Clear()
    let actualChain = chain (captured (fun () -> apply fn (box 3)))
    check "argument exception stops further evaluation" (List.ofSeq events = expected)
    events.Clear()
    let oracleChain = chain (captured (fun () -> apply reference (box 3)))
    check "argument exception wrapper chain matches generic oracle" (actualChain = oracleChain && actualChain = 2)
    check "generic oracle has same exception order" (List.ofSeq events = expected)
printfn "int-arithmetic runtime: %d checks passed" checks
