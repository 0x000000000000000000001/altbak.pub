const fs = require('fs');
const path = '/Users/0x1/Documents/htdocs/purust/purust/src/Main.purs';
let content = fs.readFileSync(path, 'utf8');

const injection = `
    allModules <- Ref.read modulesRef
    
    -- Transitive closure of imports
    tcRef <- Ref.new (Map.empty :: Map.Map String (Set.Set String))
    let initTc = Map.toUnfoldable allModules :: Array (Tuple String { code :: String, imports :: Array String })
    _ <- foldl (\\eff (Tuple k v) -> eff *> Ref.modify_ (Map.insert k (Set.fromFoldable v.imports)) tcRef) (pure unit) initTc
    
    let loop = do
          changed <- Ref.new false
          currMap <- Ref.read tcRef
          let arr = Map.toUnfoldable currMap :: Array (Tuple String (Set.Set String))
          _ <- foldl (\\eff (Tuple k imps) -> eff *> do
            let newImps = foldl (\\acc i -> 
                  case Map.lookup i currMap of
                    Just trans -> Set.union acc trans
                    Nothing -> acc
                ) imps (Set.toUnfoldable imps :: Array String)
            if Set.size newImps > Set.size imps then do
               Ref.write true changed
               Ref.modify_ (Map.insert k newImps) tcRef
            else pure unit
          ) (pure unit) arr
          isChanged <- Ref.read changed
          if isChanged then loop else pure unit
    loop
    finalTcMap <- Ref.read tcRef
`;

content = content.replace(
  "allModules <- Ref.read modulesRef",
  injection
);

content = content.replace(
  "let modDeps = \"purust_core = { path = \\\"../purust_core\\\" }\\nperceus_ptr = { path = \\\"/Users/0x1/Documents/htdocs/purust/purust/tests/runtime/perceus_ptr\\\" }\\nfancy-regex = \\\"0.13\\\"\\n\" <> String.joinWith \"\\n\" (map (\\i -> \"Purs_\" <> i <> \" = { path = \\\"../Purs_\" <> i <> \"\\\" }\") imp)",
  `let modDeps = "purust_core = { path = \\\"../purust_core\\\" }\\nperceus_ptr = { path = \\\"/Users/0x1/Documents/htdocs/purust/purust/tests/runtime/perceus_ptr\\\" }\\nfancy-regex = \\\"0.13\\\"\\n\" <> String.joinWith \"\\n\" (map (\\i -> \"Purs_\" <> i <> \" = { path = \\\"../Purs_\" <> i <> \"\\\" }\") (fromMaybe [] (map (\\s -> Set.toUnfoldable s :: Array String) (Map.lookup k finalTcMap))))`
);

content = content.replace(
  `let rustCode = "#![allow(warnings)]\\nuse perceus_ptr::PerceusPtr;\\nuse purust_core::*;\\n\" <> importsRust <> \"\\n\\n\" <> rsFile <> \"\\n\\n\" <> ffiContent <> \"\\n\\n\"`,
  `let rustCode = "#![allow(warnings)]\\nuse perceus_ptr::PerceusPtr;\\nuse purust_core::*;\\n\" <> importsRust <> \"\\n\\n\" <> rsFile <> \"\\n\\n\" <> ffiContent <> \"\\n\\n\"`
);

fs.writeFileSync(path, content);
