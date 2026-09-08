"""Create the isolated direct-call experiment from the current generated F#."""

from pathlib import Path
import difflib
import shutil


audit = Path(__file__).resolve().parent
source = audit.parent.parent / "output" / "Main"
target = audit / "direct-call"
old = '''    let method = func.GetType().GetMethods() |> Array.find (fun m -> m.Name = "Invoke" && m.GetParameters().Length = 1)
    method.Invoke(func, [| arg |])'''
new = '''    match func with
    | :? (obj -> obj) as invoke -> invoke arg
    | _ ->
        let method = func.GetType().GetMethods() |> Array.find (fun m -> m.Name = "Invoke" && m.GetParameters().Length = 1)
        method.Invoke(func, [| arg |])'''
original = (source / "Sharpurs_Prelude.fs").read_text()
if original.count(old) != 1:
    raise SystemExit("Original reflection implementation differs; inspect before reproducing.")

target.mkdir(exist_ok=True)
for file in source.iterdir():
    if file.is_file() and file.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"}:
        shutil.copy2(file, target / file.name)

modified = original.replace(old, new)
(target / "Sharpurs_Prelude.fs").write_text(modified)
(audit / "direct-call.patch").write_text("".join(difflib.unified_diff(
    original.splitlines(keepends=True),
    modified.splitlines(keepends=True),
    fromfile="original/Sharpurs_Prelude.fs",
    tofile="direct-call/Sharpurs_Prelude.fs",
)))
print(f"Prepared {target}; original generated program unchanged.")
