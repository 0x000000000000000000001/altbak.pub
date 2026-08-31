const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Data.Map.Internal/corefn.json'));
const types = data.typeTable;
const printType = (id) => {
  if (id === null || id === undefined) return "Any";
  const t = types[id];
  if (!t) return "Unknown";
  if (t.type === "Func") return `Func([${t.args.map(printType).join(", ")}], ${printType(t.ret)})`;
  if (t.type === "ConstrainedType") return `ConstrainedType([${t.constraints.map(c => c.fqn.join(".")).join(", ")}], ${printType(t.body)})`;
  if (t.type === "ForAll") return `ForAll([${t.vars.join(", ")}], ${printType(t.body)})`;
  if (t.type === "TypeApp") return `TypeApp(${printType(t.constructor)}, ${t.args.map(printType).join(", ")})`;
  if (t.type === "TypeVar") return `TypeVar(${t.name})`;
  if (t.type === "Adt") return `Adt(${t.fqn.join(".")})`;
  return JSON.stringify(t);
};
const lookupDecl = data.decls.find(d => d.identifier === "lookup");
console.log(printType(lookupDecl.expression.annotation.type));
