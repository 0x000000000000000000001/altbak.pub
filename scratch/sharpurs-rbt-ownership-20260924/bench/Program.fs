module Bench

open System

// ---------------- persistent (Okasaki) ----------------
type T = E | T of int * T * int * T
let rec pins x s =
    match s with
    | E -> T(0, E, x, E)
    | T(color, a, y, b) ->
        if x < y then pbalance color (pins x a) y b
        elif x > y then pbalance color a y (pins x b)
        else T(color, a, y, b)
and pbalance color a x b =
    match color, a, x, b with
    | 1, T(0, T(0, a1, x1, b1), y1, c1), z1, d1 -> T(0, T(1, a1, x1, b1), y1, T(1, c1, z1, d1))
    | 1, T(0, a1, x1, T(0, b1, y1, c1)), z1, d1 -> T(0, T(1, a1, x1, b1), y1, T(1, c1, z1, d1))
    | 1, a1, x1, T(0, T(0, b1, y1, c1), z1, d1) -> T(0, T(1, a1, x1, b1), y1, T(1, c1, z1, d1))
    | 1, a1, x1, T(0, b1, y1, T(0, c1, z1, d1)) -> T(0, T(1, a1, x1, b1), y1, T(1, c1, z1, d1))
    | c, a1, x1, b1 -> T(c, a1, x1, b1)
let pinsert x s = match pins x s with | T(_, a, y, b) -> T(1, a, y, b) | E -> E
let rec pbuildTree n acc = if n = 0 then acc else pbuildTree (n - 1) (pinsert n acc)
let rec pdepth s = match s with E -> 0 | T(_, a, _, b) -> 1 + max (pdepth a) (pdepth b)

// ---------------- mutable (same algorithm, in place) ----------------
[<AllowNullLiteral>]
type Node(c: int, l: Node, k: int, r: Node) =
    member val Color = c with get, set
    member val Left = l with get, set
    member val Key = k with get, set
    member val Right = r with get, set

let balance (node: Node) =
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

let rec ins x (node: Node) =
    if node = null then Node(0, null, x, null)
    else
        if x < node.Key then
            node.Left <- ins x node.Left
            balance node
        elif x > node.Key then
            node.Right <- ins x node.Right
            balance node
        else node

let insert x (t: Node) =
    let result = ins x t
    result.Color <- 1
    result

let rec buildTree n (acc: Node) =
    if n = 0 then acc else buildTree (n - 1) (insert n acc)

let rec depth (node: Node) =
    if node = null then 0 else 1 + max (depth node.Left) (depth node.Right)

let now () = float (System.Diagnostics.Stopwatch.GetTimestamp()) / float System.Diagnostics.Stopwatch.Frequency * 1000000.0

[<EntryPoint>]
let main _ =
    let n = 100000
    let mutable sink = 0
    for _ in 1 .. 3 do
        sink <- sink + pdepth (pbuildTree n E)
        sink <- sink + depth (buildTree n null)
    let best f =
        let mutable b = infinity
        for _ in 1 .. 10 do
            let t0 = now ()
            let v = f ()
            let t1 = now ()
            sink <- sink + v
            if t1 - t0 < b then b <- t1 - t0
        b
    let p = best (fun () -> pdepth (pbuildTree n E))
    let m = best (fun () -> depth (buildTree n null))
    printfn "persistent: %.3f ms | mutable: %.3f ms | sink=%d" (p / 1000.0) (m / 1000.0) sink
    0
