import assert from "node:assert/strict";
import * as P from "./output/AdtPilot/index.js";

let count = 0;
const equal = (actual, expected) => { assert.equal(actual, expected); count++; };
equal(P.depth(P.empty), 0);
equal(P.depth(P.singleton), 1);
equal(P.depth(P.asymmetric), 3);
equal(P.rootValue(P.empty), 0);
equal(P.rootValue(P.singleton), 2147483647);
equal(P.rootValue(P.leftChild(P.asymmetric)), -2147483648);
equal(P.isRed(P.rootColor(P.empty)), false);
equal(P.isRed(P.rootColor(P.singleton)), false);
equal(P.isRed(P.rootColor(P.leftChild(P.asymmetric))), true);

// A partially applied public constructor function can be reused independently.
const makeRed = P.singletonWith(P.R.value);
const redA = makeRed(7);
const redB = makeRed(-9);
equal(P.rootValue(redA), 7);
equal(P.rootValue(redB), -9);
equal(P.isRed(P.rootColor(redA)), true);
equal(P.depth(redA), 1);

// Repeated values, distinct paths and persistent subtree sharing remain valid.
const shared = P.T.create(P.B.value)(redA)(7)(redA);
equal(P.depth(shared), 2);
equal(P.leftChild(shared), redA);
equal(P.rootValue(redA), 7);

// Exercise recursion separately from balanced-tree insertion.
let skewed = P.E.value;
for (let i = 0; i < 1000; i++) skewed = new P.T(P.R.value, skewed, i, P.E.value);
equal(P.depth(skewed), 1000);
equal(P.rootValue(skewed), 999);

console.log(`${count} JavaScript reference assertions passed`);
