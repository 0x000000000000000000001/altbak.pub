import fs from 'fs';
let code = fs.readFileSync('src/Bench.rs', 'utf8');

code = code.replace(
    'pub fn Bench_formatNumber(mut nObj: UnknownType) -> UnknownType {\n    let n = nObj.init_number.unwrap();\n    mk_string(&format!("{:.2}", n))\n}',
    'pub fn Bench_formatNumber(mut n: f64) -> String {\n    format!("{:.2}", n)\n}'
);

fs.writeFileSync('src/Bench.rs', code);
