open System
open System.Diagnostics

[<EntryPoint>]
let main argv =
    let sw = Stopwatch.StartNew()
    printfn "Hello world"
    sw.Stop()
    printfn "%d ms" sw.ElapsedMilliseconds
    0
