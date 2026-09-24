// Compare method IL between two assemblies without loading their dependencies.
open System
open System.IO
open System.Reflection.Metadata
open System.Reflection.PortableExecutable

let dump (path: string) =
    use fs = File.OpenRead(path)
    use pe = new PEReader(fs)
    let md = pe.GetMetadataReader()
    let hash (bytes: byte array) =
        use sha = System.Security.Cryptography.SHA256.Create()
        Convert.ToHexString(sha.ComputeHash(bytes))
    let typeName (handle: TypeDefinitionHandle) =
        if handle.IsNil then ""
        else
            let t = md.GetTypeDefinition(handle)
            let ns = md.GetString(t.Namespace)
            let n = md.GetString(t.Name)
            if ns = "" then n else ns + "." + n
    [ for handle in md.MethodDefinitions do
        let m = md.GetMethodDefinition(handle)
        let il =
            if m.RelativeVirtualAddress = 0 then [||]
            else (pe.GetMethodBody(m.RelativeVirtualAddress)).GetILBytes()
        let nested =
            if m.GetDeclaringType().IsNil then ""
            else
                let t = md.GetTypeDefinition(m.GetDeclaringType())
                if t.IsNested then "nested/" else ""
        yield
            (sprintf "%s%s.%s" nested (typeName (m.GetDeclaringType())) (md.GetString(m.Name))),
            (il.Length, hash il) ]
    |> Map.ofList

let main (argv: string array) : int =
    let a = dump argv.[0]
    let b = dump argv.[1]
    let names = Map.toArray a |> Array.map fst
    let namesB = Map.toArray b |> Array.map fst |> Set.ofArray
    let onlyA = names |> Array.filter (fun n -> not (Set.contains n namesB))
    let common = names |> Array.filter (fun n -> Set.contains n namesB)
    let different =
        common
        |> Array.filter (fun n ->
            match Map.tryFind n a, Map.tryFind n b with
            | Some (la, ha), Some (lb, hb) -> la <> lb || ha <> hb
            | _ -> true)
    Console.WriteLine($"methods A={a.Count} B={b.Count}")
    Console.WriteLine($"only in A: {onlyA.Length}")
    onlyA |> Array.truncate 20 |> Array.iter (fun n -> Console.WriteLine("  A " + n))
    Console.WriteLine($"IL differs in {different.Length} common methods")
    different |> Array.truncate 60 |> Array.iter (fun n ->
        match Map.tryFind n a, Map.tryFind n b with
        | Some (la, ha), Some (lb, hb) -> Console.WriteLine($"  {n}: A[{la}] {ha.Substring(0, 8)} B[{lb}] {hb.Substring(0, 8)}")
        | _ -> ())
    0

main (fsi.CommandLineArgs |> Array.skip 1) |> ignore
