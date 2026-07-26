import fs from 'fs';
import path from 'path';
console.log(process.cwd());
const p = "../phpurs-enums/src/Data/Enum.php";
console.log(path.resolve(p));
console.log(fs.existsSync(p));
