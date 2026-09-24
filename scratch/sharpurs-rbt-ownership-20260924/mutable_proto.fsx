// Mutable red-black tree prototype: same Okasaki algorithm, in-place updates.
// Measures the ceiling of an ownership-based mutation transform.

[<AllowNullLiteral>]
type Node(c: int, l: Node, k: int, r: Node) =
    member val Color = c with get, set   // 0 = R, 1 = B
    member val Left = l with get, set
    member val Key = k with get, set
    member val Right = r with get, set

let rec balance (node: Node) =
    let y = node.Left
    if y <> null && y.Color = 0 then
        let x = y.Left
        if x <> null && x.Color = 0 then
            // Left-left: y becomes the red root, x and node become black.
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
                // Left-right: w becomes the red root.
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
                // Right-right.
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
                    // Right-left.
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
    if node = null then 0
    else 1 + max (depth node.Left) (depth node.Right)

// ---- reference persistent implementation, to compare depths ----
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

let main () =
    // Equivalence of resulting depth for many sizes.
    let mutable mismatches = 0
    for n in [ 0; 1; 2; 3; 5; 10; 50; 100; 1000; 4096; 20000 ] do
        let mutable m: Node = null
        for k in n .. -1 .. 1 do m <- insert k m
        let pm = pbuildTree n E
        if depth m <> pdepth pm then
            mismatches <- mismatches + 1
            printfn "depth mismatch n=%d mutable=%d persistent=%d" n (depth m) (pdepth pm)
    printfn "depth equivalence mismatches: %d" mismatches

    let now () = float (System.Diagnostics.Stopwatch.GetTimestamp()) / float System.Diagnostics.Stopwatch.Frequency * 1000000.0
    let best n f =
        let mutable b = infinity
        for _ in 1 .. n do
            let t0 = now ()
            let _ = f ()
            let t1 = now ()
            if t1 - t0 < b then b <- t1 - t0
        b

    // warm
    for _ in 1 .. 3 do
        depth (buildTree 100000 null) |> ignore

    let mutableUs = best 15 (fun () -> depth (buildTree 100000 null))
    let persistentUs = best 15 (fun () -> pdepth (pbuildTree 100000 E))
    let check = depth (buildTree 100000 null)
    printfn "mutable:    depth=%d  %.3f ms" check (mutableUs / 1000.0)
    printfn "persistent: depth=%d  %.3f ms" (pdepth (pbuildTree 100000 E)) (persistentUs / 1000.0)

    // Allocation floor: build 2.58M nodes without any algorithm.
    let d0 = System.GC.GetTotalAllocatedBytes(true)
    let t0 = now ()
    let rec alloc n acc = if n = 0 then acc else alloc (n - 1) (Node(0, acc, n, acc))
    let big = alloc 2_580_000 null
    let t1 = now ()
    let d1 = System.GC.GetTotalAllocatedBytes(true)
    printfn "alloc 2.58M nodes: %.3f ms, %.1f MB (sink %b)" ((t1 - t0) / 1000.0) ((float (d1 - d0)) / 1e6) (big <> null)
    0

main () |> ignore
