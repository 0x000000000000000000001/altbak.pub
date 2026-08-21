const fs = require('fs');
const path = require('path');

const outputDir = 'run/bak/rust/output/purust_output';
const dirs = fs.readdirSync(outputDir).filter(d => d.startsWith('Purs_'));

for (const dir of dirs) {
    const toml = fs.readFileSync(path.join(outputDir, dir, 'Cargo.toml'), 'utf8');
    const match = toml.match(/Purs_Main =/);
    if (match) {
        console.log(dir + " depends on Purs_Main");
    }
}
