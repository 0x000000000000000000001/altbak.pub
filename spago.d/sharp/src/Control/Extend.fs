module Control.Extend_FFI

let arrayExtend (f: obj) (arr: obj) =
    let f' = f :?> (obj -> obj)
    let arr' = unbox<obj[]> arr
    let res = Array.zeroCreate arr'.Length
    for i = 0 to arr'.Length - 1 do
        let subArr = arr'.[i..]
        res.[i] <- f' (box subArr)
    box res
