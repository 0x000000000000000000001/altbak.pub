// File I/O and loop adapters for the extended Sharpurs benchmark.
// The signatures mirror the PureScript foreign imports in FileOps.purs.

let writeFileSync = box (fun (pathObj: obj) -> box (fun (contentObj: obj) -> box (fun (_: obj) ->
    let path = unbox<string> pathObj
    let parent = System.IO.Path.GetDirectoryName path
    if not (System.String.IsNullOrEmpty parent) then
        System.IO.Directory.CreateDirectory parent |> ignore
    System.IO.File.WriteAllText(path, unbox<string> contentObj)
    box ())))

let readFileSync = box (fun (pathObj: obj) -> box (fun (_: obj) ->
    box (System.IO.File.ReadAllText(unbox<string> pathObj))))

let loopE = box (fun (countObj: obj) -> box (fun (actionObj: obj) -> box (fun (_: obj) ->
    let count = unbox<int> countObj
    let action = unbox<obj -> obj> actionObj
    for _ in 1 .. count do
        action (box null) |> ignore
    box ())))
