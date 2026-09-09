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