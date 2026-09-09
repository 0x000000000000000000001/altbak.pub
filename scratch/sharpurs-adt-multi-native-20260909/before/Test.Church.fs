[<AutoOpen>]
module PureScript_Test_Church

open System
open System.Collections.Generic

let Test_Church_zeroC  = (box (fun (v: obj) -> (box (fun (x: obj) -> (box x)))))

let Test_Church_toInt  = (box (fun (n: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box n))) (box ((box (fun (x: obj) -> (box ((unbox<int> (box ((box x)))) + (unbox<int> (box ((box 1))))))))))))) (box ((box 0))))))

let Test_Church_succC  = (box (fun (n: obj) -> (box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box f))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box n))) (box ((box f)))))) (box ((box x)))))))))))))

let Test_Church_mulC  = (box (fun (m: obj) -> (box (fun (n: obj) -> (box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box m))) (box ((sharpurs_apply (box ((box n))) (box ((box f))))))))) (box ((box x))))))))))))

let rec Test_Church_fromInt_tco (v: obj) : obj = ((match ((unbox ((box v)))) with | LitInt 0 () -> ((box Test_Church_zeroC)) | n -> ((sharpurs_apply (box ((box Test_Church_succC))) (box ((Test_Church_fromInt_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))))))))))
and Test_Church_fromInt = box (fun (v: obj) -> Test_Church_fromInt_tco v)


let Test_Church_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Church Numerals (100k Closure Applications):"))))

let Test_Church_c10  = (box (fun (n: obj) -> (sharpurs_apply (box ((box Test_Church_fromInt))) (box ((box n))))))

let Test_Church_c100  = (box (fun (n: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Test_Church_mulC))) (box ((sharpurs_apply (box ((box Test_Church_c10))) (box ((box n))))))))) (box ((sharpurs_apply (box ((box Test_Church_c10))) (box ((box n)))))))))

let Test_Church_c10k  = (box (fun (n: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Test_Church_mulC))) (box ((sharpurs_apply (box ((box Test_Church_c100))) (box ((box n))))))))) (box ((sharpurs_apply (box ((box Test_Church_c100))) (box ((box n)))))))))

let Test_Church_c100k  = (box (fun (n: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Test_Church_mulC))) (box ((sharpurs_apply (box ((box Test_Church_c10k))) (box ((box n))))))))) (box ((sharpurs_apply (box ((box Test_Church_c10))) (box ((box n)))))))))

let Test_Church_addC  = (box (fun (m: obj) -> (box (fun (n: obj) -> (box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box m))) (box ((box f)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box n))) (box ((box f)))))) (box ((box x)))))))))))))))

let Test_Church_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_Church_toInt))) (box ((sharpurs_apply (box ((box Test_Church_c100k))) (box ((box dummy))))))))))))))))))
