module BenchFields

open System

// Same mutable algorithm, but with direct public fields instead of auto-properties.
[<AllowNullLiteral>]
type FNode =
    val mutable Color: int   // 0 = R, 1 = B
    val mutable Left: FNode
    val mutable Key: int
    val mutable Right: FNode
    new(c, l, k, r) = { Color = c; Left = l; Key = k; Right = r }

let balance (node: FNode) =
    let y = node.Left
    if y <> null && y.Color = 0 then
        let x = y.Left
        if x <> null && x.Color = 0 then
            let c = y.Right
            y.Right <- node
            node.Left <- c
            node.Color <- 1
            x.Color <- 1
            y.Color <- 0
            y
        else
            let w = y.Right
            if w <> null && w.Color = 0 then
                let b = w.Left
                let c = w.Right
                y.Right <- b
                node.Left <- c
                y.Color <- 1
                node.Color <- 1
                w.Left <- y
                w.Right <- node
                w.Color <- 0
                w
            else node
    else
        let y = node.Right
        if y <> null && y.Color = 0 then
            let x = y.Right
            if x <> null && x.Color = 0 then
                let b = y.Left
                y.Left <- node
                node.Right <- b
                node.Color <- 1
                x.Color <- 1
                y.Color <- 0
                y
            else
                let w = y.Left
                if w <> null && w.Color = 0 then
                    let b = w.Left
                    let c = w.Right
                    y.Left <- c
                    node.Right <- b
                    y.Color <- 1
                    node.Color <- 1
                    w.Left <- node
                    w.Right <- y
                    w.Color <- 0
                    w
                else node
        else node

let rec ins x (node: FNode) =
    if node = null then FNode(0, null, x, null)
    else
        if x < node.Key then
            node.Left <- ins x node.Left
            balance node
        elif x > node.Key then
            node.Right <- ins x node.Right
            balance node
        else node

let insert x (t: FNode) =
    let result = ins x t
    result.Color <- 1
    result

let rec buildTree n (acc: FNode) =
    if n = 0 then acc else buildTree (n - 1) (insert n acc)

let rec depth (node: FNode) =
    if node = null then 0 else 1 + max (depth node.Left) (depth node.Right)

let now () = float (System.Diagnostics.Stopwatch.GetTimestamp()) / float System.Diagnostics.Stopwatch.Frequency * 1000000.0

[<EntryPoint>]
let main _ =
    let n = 100000
    let mutable sink = 0
    for _ in 1 .. 3 do sink <- sink + depth (buildTree n null)
    let mutable bestMs = infinity
    for _ in 1 .. 10 do
        let t0 = now ()
        let v = depth (buildTree n null)
        let t1 = now ()
        sink <- sink + v
        if t1 - t0 < bestMs then bestMs <- t1 - t0
    printfn "mutable-fields: %.3f ms | sink=%d" (bestMs / 1000.0) sink
    0
