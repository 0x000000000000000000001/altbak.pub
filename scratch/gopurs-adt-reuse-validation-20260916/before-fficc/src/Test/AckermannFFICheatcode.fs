module Test.AckermannFFICheatcode

let rec ack m n =
    if m = 0 then n + 1
    elif n = 0 then ack (m - 1) 1
    else ack (m - 1) (ack m (n - 1))
let runAckermannFFICheatcode (m: obj) = ack (unbox<int> m) 4 :> obj
