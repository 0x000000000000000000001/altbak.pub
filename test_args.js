const argsRaw = process.argv;
const args = argsRaw.flatMap(s => s.split(" "));
console.log(argsRaw);
console.log(args);
const i = args.indexOf("--main");
if (i !== -1) {
  console.log(args[i + 1]);
}
