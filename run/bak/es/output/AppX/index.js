import * as App from "../App/index.js";
import * as Test$dAffOperations from "../Test.AffOperations/index.js";
import * as Test$dFileOps from "../Test.FileOps/index.js";
import * as Test$dSTArray from "../Test.STArray/index.js";
import * as Test$dStringOps from "../Test.StringOps/index.js";
const main = () => {
  App.main();
  Test$dFileOps.describe();
  Test$dFileOps.act();
  Test$dSTArray.describe();
  Test$dSTArray.act();
  Test$dStringOps.describe();
  Test$dStringOps.act();
  Test$dAffOperations.describe();
  return Test$dAffOperations.act();
};
export {main};
