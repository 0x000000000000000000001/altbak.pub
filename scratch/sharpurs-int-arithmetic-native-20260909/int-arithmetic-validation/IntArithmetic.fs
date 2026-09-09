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