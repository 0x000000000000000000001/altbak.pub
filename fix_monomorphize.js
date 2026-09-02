const fs = require('fs');
const file = '/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Monomorphize.purs';
let content = fs.readFileSync(file, 'utf8');

const target = `                in
                  if Array.length specs.newBinds > 0 then
                    { binds: Array.snoc acc.binds (Rec (specs.newBinds <> bs)), polyMap: specs.polyMap }
                  else
                    { binds: Array.snoc acc.binds (Rec bs), polyMap: acc.polyMap }`;

const replacement = `                in
                  if Array.length specs.newBinds > 0 then
                    if Array.length bs == 1 then
                      { binds: acc.binds <> map (\\b -> Rec [b]) (specs.newBinds <> bs), polyMap: specs.polyMap }
                    else
                      { binds: Array.snoc acc.binds (Rec (specs.newBinds <> bs)), polyMap: specs.polyMap }
                  else
                    { binds: Array.snoc acc.binds (Rec bs), polyMap: acc.polyMap }`;

if (content.includes(target)) {
    content = content.replace(target, replacement);
    fs.writeFileSync(file, content);
    console.log("Patched Monomorphize.purs");
} else {
    console.log("Target not found!");
}
