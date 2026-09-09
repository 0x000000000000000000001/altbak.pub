let IntCompare_lessThan  = (sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box Data_Ord_ordInt))))

let IntCompare_Reverse  = (box (fun (x: obj) -> (box x)))

let IntCompare_stringLess  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box Data_Ord_ordString)))))) (box ((box x)))))) (box ((box y))))))))

let IntCompare_partial  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box Data_Ord_ordInt)))))) (box ((box 5))))

let IntCompare_ordered  = (box (fun (x: obj) -> (box (fun (y: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntCompare_track))) (box ((box 1)))))) (box ((box x))))))) < (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntCompare_track))) (box ((box 2)))))) (box ((box y)))))))))))))

let IntCompare_numberLess  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box Data_Ord_ordNumber)))))) (box ((box x)))))) (box ((box y))))))))

let IntCompare_less  = (box (fun (x: obj) -> (box (fun (y: obj) -> (box ((unbox<int> (box ((box x)))) < (unbox<int> (box ((box y))))))))))

let IntCompare_greater  = (box (fun (x: obj) -> (box (fun (y: obj) -> (box ((unbox<int> (box ((box x)))) > (unbox<int> (box ((box y))))))))))

let IntCompare_genericLess  = (box (fun (dictOrd: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y))))))))))

let IntCompare_eqReverse  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqInt)))))) (box ((box l)))))) (box ((box r)))))))))))) Map.empty))))))

let IntCompare_ordReverse  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box Data_Ord_ordInt)))))) (box ((box y)))))) (box ((box x)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box IntCompare_eqReverse))))) Map.empty)))))))

let IntCompare_custom  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box IntCompare_ordReverse)))))) (box ((sharpurs_apply (box ((box IntCompare_Reverse))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box IntCompare_Reverse))) (box ((box y)))))))))))

let IntCompare_annotated  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box IntCompare_lessThan))) (box ((box x)))))) (box ((box y))))))))