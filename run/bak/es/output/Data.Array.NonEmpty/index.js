import * as $runtime from "../runtime.js";
import * as Control$dBind from "../Control.Bind/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dArray$dNonEmpty$dInternal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dFunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dNonEmpty from "../Data.NonEmpty/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dSemigroup$dFoldable from "../Data.Semigroup.Foldable/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const intercalate1 = dictSemigroup => {
  const semigroupJoinWith1 = {append: v => v1 => j => dictSemigroup.append(v(j))(dictSemigroup.append(j)(v1(j)))};
  return a => foldable => Data$dArray$dNonEmpty$dInternal.foldable1NonEmptyArray.foldMap1(semigroupJoinWith1)(x => v => x)(foldable)(a);
};
const transpose = x => Data$dArray.transpose(x);
const toArray = v => v;
const unionBy$p = eq => xs => x => Data$dArray.unionBy(eq)(xs)(x);
const union$p = dictEq => unionBy$p(dictEq.eq);
const unionBy = eq => xs => x => Data$dArray.unionBy(eq)(xs)(x);
const union = dictEq => unionBy(dictEq.eq);
const unzip = x => {
  const $0 = Data$dArray.unzip(x);
  return Data$dTuple.$Tuple($0._1, $0._2);
};
const updateAt = i => x => x$1 => Data$dArray._updateAt(Data$dMaybe.Just, Data$dMaybe.Nothing, i, x, x$1);
const zip = xs => ys => Data$dArray.zipWithImpl(Data$dTuple.Tuple, xs, ys);
const zipWith = f => xs => ys => Data$dArray.zipWithImpl(f, xs, ys);
const zipWithA = dictApplicative => f => xs => ys => Data$dTraversable.traversableArray.traverse(dictApplicative)(Data$dTraversable.identity)(Data$dArray.zipWithImpl(f, xs, ys));
const splitAt = i => xs => Data$dArray.splitAt(i)(xs);
const some = dictAlternative => Data$dArray.some(dictAlternative);
const snoc$p = xs => x => Data$dArray.snoc(xs)(x);
const snoc = xs => x => Data$dArray.snoc(xs)(x);
const singleton = x => [x];
const replicate = i => x => {
  const v = Data$dOrd.ordInt.compare(1)(i);
  return Data$dArray.replicateImpl(
    (() => {
      if (v === "LT") { return i; }
      if (v === "EQ") { return 1; }
      if (v === "GT") { return 1; }
      $runtime.fail();
    })(),
    x
  );
};
const range = x => y => Data$dArray.rangeImpl(x, y);
const prependArray = xs => ys => [...xs, ...ys];
const modifyAt = i => f => x => Data$dArray.modifyAt(i)(f)(x);
const intersectBy$p = eq => xs => Data$dArray.intersectBy(eq)(xs);
const intersectBy = eq => xs => x => Data$dArray.intersectBy(eq)(xs)(x);
const intersect$p = dictEq => intersectBy$p(dictEq.eq);
const intersect = dictEq => intersectBy(dictEq.eq);
const intercalate = dictSemigroup => intercalate1(dictSemigroup);
const insertAt = i => x => x$1 => Data$dArray._insertAt(Data$dMaybe.Just, Data$dMaybe.Nothing, i, x, x$1);
const fromFoldable1 = dictFoldable1 => {
  const $0 = dictFoldable1.Foldable0().foldr;
  return x => Data$dArray.fromFoldableImpl($0, x);
};
const fromArray = xs => {
  if (xs.length > 0) { return Data$dMaybe.$Maybe("Just", xs); }
  return Data$dMaybe.Nothing;
};
const fromFoldable = dictFoldable => {
  const $0 = dictFoldable.foldr;
  return x => {
    const $1 = Data$dArray.fromFoldableImpl($0, x);
    if ($1.length > 0) { return Data$dMaybe.$Maybe("Just", $1); }
    return Data$dMaybe.Nothing;
  };
};
const transpose$p = x => {
  const $0 = Data$dArray.transpose(x);
  if ($0.length > 0) { return Data$dMaybe.$Maybe("Just", $0); }
  return Data$dMaybe.Nothing;
};
const foldr1 = /* #__PURE__ */ (() => Data$dArray$dNonEmpty$dInternal.foldable1NonEmptyArray.foldr1)();
const foldl1 = /* #__PURE__ */ (() => Data$dArray$dNonEmpty$dInternal.foldable1NonEmptyArray.foldl1)();
const foldMap1 = dictSemigroup => Data$dArray$dNonEmpty$dInternal.foldable1NonEmptyArray.foldMap1(dictSemigroup);
const fold1 = dictSemigroup => Data$dArray$dNonEmpty$dInternal.foldable1NonEmptyArray.foldMap1(dictSemigroup)(Data$dSemigroup$dFoldable.identity);
const difference$p = dictEq => Data$dFoldable.foldrArray(Data$dArray.delete(dictEq));
const cons$p = x => xs => [x, ...xs];
const fromNonEmpty = v => [v._1, ...v._2];
const concatMap = b => a => Control$dBind.arrayBind(a)(b);
const concat = /* #__PURE__ */ (() => {
  const $0 = Data$dFunctor.arrayMap(toArray);
  return x => Data$dArray.concat($0(x));
})();
const appendArray = xs => ys => [...xs, ...ys];
const alterAt = i => f => x => Data$dArray.alterAt(i)(f)(x);
const head = x => {
  if (0 < x.length) { return x[0]; }
  $runtime.fail();
};
const init = x => {
  if (x.length === 0) { $runtime.fail(); }
  return Data$dArray.sliceImpl(0, x.length - 1 | 0, x);
};
const last = x => {
  const $0 = x.length - 1 | 0;
  if ($0 >= 0 && $0 < x.length) { return x[$0]; }
  $runtime.fail();
};
const tail = x => {
  const $0 = Data$dArray.unconsImpl(v => Data$dMaybe.Nothing, v => xs => Data$dMaybe.$Maybe("Just", xs), x);
  if ($0.tag === "Just") { return $0._1; }
  $runtime.fail();
};
const uncons = x => {
  const $0 = Data$dArray.unconsImpl(v => Data$dMaybe.Nothing, x$1 => xs => Data$dMaybe.$Maybe("Just", {head: x$1, tail: xs}), x);
  if ($0.tag === "Just") { return $0._1; }
  $runtime.fail();
};
const toNonEmpty = x => {
  const $0 = uncons(x);
  return Data$dNonEmpty.$NonEmpty($0.head, $0.tail);
};
const unsnoc = x => {
  const $0 = Data$dArray.unsnoc(x);
  if ($0.tag === "Just") { return $0._1; }
  $runtime.fail();
};
const all = p => x => Data$dArray.allImpl(p, x);
const any = p => x => Data$dArray.anyImpl(p, x);
const catMaybes = x => Data$dArray.mapMaybe(x$1 => x$1)(x);
const $$delete = dictEq => x => x$1 => Data$dArray.deleteBy(dictEq.eq)(x)(x$1);
const deleteAt = i => x => Data$dArray._deleteAt(Data$dMaybe.Just, Data$dMaybe.Nothing, i, x);
const deleteBy = f => x => x$1 => Data$dArray.deleteBy(f)(x)(x$1);
const difference = dictEq => xs => Data$dFoldable.foldrArray(Data$dArray.delete(dictEq))(xs);
const drop = i => x => {
  if (i < 1) { return x; }
  return Data$dArray.sliceImpl(i, x.length, x);
};
const dropEnd = i => x => {
  const $0 = x.length - i | 0;
  if ($0 < 1) { return []; }
  return Data$dArray.sliceImpl(0, $0, x);
};
const dropWhile = f => x => Data$dArray.span(f)(x).rest;
const elem = dictEq => x => x$1 => Data$dArray.elem(dictEq)(x)(x$1);
const elemIndex = dictEq => x => x$1 => Data$dArray.findIndexImpl(Data$dMaybe.Just, Data$dMaybe.Nothing, v => dictEq.eq(v)(x), x$1);
const elemLastIndex = dictEq => x => x$1 => Data$dArray.findLastIndexImpl(Data$dMaybe.Just, Data$dMaybe.Nothing, v => dictEq.eq(v)(x), x$1);
const filter = f => x => Data$dArray.filterImpl(f, x);
const filterA = dictApplicative => f => Data$dArray.filterA(dictApplicative)(f);
const find = p => x => Data$dArray.find(p)(x);
const findIndex = p => x => Data$dArray.findIndexImpl(Data$dMaybe.Just, Data$dMaybe.Nothing, p, x);
const findLastIndex = x => x$1 => Data$dArray.findLastIndexImpl(Data$dMaybe.Just, Data$dMaybe.Nothing, x, x$1);
const findMap = p => x => Data$dArray.findMapImpl(Data$dMaybe.Nothing, Data$dMaybe.isJust, p, x);
const foldM = dictMonad => f => acc => Data$dArray.foldM(dictMonad)(f)(acc);
const foldRecM = dictMonadRec => f => acc => Data$dArray.foldRecM(dictMonadRec)(f)(acc);
const index = x => $0 => {
  if ($0 >= 0 && $0 < x.length) { return Data$dMaybe.$Maybe("Just", x[$0]); }
  return Data$dMaybe.Nothing;
};
const length = x => x.length;
const mapMaybe = f => x => Data$dArray.mapMaybe(f)(x);
const notElem = dictEq => x => x$1 => Data$dArray.notElem(dictEq)(x)(x$1);
const partition = f => x => Data$dArray.partitionImpl(f, x);
const slice = start => end => x => Data$dArray.sliceImpl(start, end, x);
const span = f => x => Data$dArray.span(f)(x);
const take = i => x => {
  if (i < 1) { return []; }
  return Data$dArray.sliceImpl(0, i, x);
};
const takeEnd = i => x => Data$dArray.takeEnd(i)(x);
const takeWhile = f => x => Data$dArray.span(f)(x).init;
const toUnfoldable = dictUnfoldable => x => {
  const len = x.length;
  return dictUnfoldable.unfoldr(i => {
    if (i < len) { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(x[i], i + 1 | 0)); }
    return Data$dMaybe.Nothing;
  })(0);
};
const cons = x => x$1 => [x, ...x$1];
const group = dictEq => {
  const eq = dictEq.eq;
  return x => Data$dArray.groupBy(eq)(x);
};
const groupAllBy = op => Data$dArray.groupAllBy(op);
const groupAll = dictOrd => Data$dArray.groupAllBy(dictOrd.compare);
const groupBy = op => x => Data$dArray.groupBy(op)(x);
const insert = dictOrd => x => x$1 => Data$dArray.insertBy(dictOrd.compare)(x)(x$1);
const insertBy = f => x => x$1 => Data$dArray.insertBy(f)(x)(x$1);
const intersperse = x => x$1 => Data$dArray.intersperse(x)(x$1);
const mapWithIndex = f => Data$dFunctorWithIndex.mapWithIndexArray(f);
const modifyAtIndices = dictFoldable => is => f => x => Data$dArray.modifyAtIndices(dictFoldable)(is)(f)(x);
const nub = dictOrd => x => Data$dArray.nubBy(dictOrd.compare)(x);
const nubBy = f => x => Data$dArray.nubBy(f)(x);
const nubByEq = f => x => Data$dArray.nubByEq(f)(x);
const nubEq = dictEq => x => Data$dArray.nubByEq(dictEq.eq)(x);
const reverse = x => Data$dArray.reverse(x);
const scanl = f => x => x$1 => Data$dArray.scanlImpl(f, x, x$1);
const scanr = f => x => x$1 => Data$dArray.scanrImpl(f, x, x$1);
const sort = dictOrd => {
  const compare = dictOrd.compare;
  return x => Data$dArray.sortBy(compare)(x);
};
const sortBy = f => x => Data$dArray.sortBy(f)(x);
const sortWith = dictOrd => f => x => Data$dArray.sortWith(dictOrd)(f)(x);
const updateAtIndices = dictFoldable => pairs => x => Data$dArray.updateAtIndices(dictFoldable)(pairs)(x);
const unsafeIndex = () => x => $0 => x[$0];
const toUnfoldable1 = dictUnfoldable1 => xs => {
  const len = xs.length;
  return dictUnfoldable1.unfoldr1(i => Data$dTuple.$Tuple(xs[i], i < (len - 1 | 0) ? Data$dMaybe.$Maybe("Just", i + 1 | 0) : Data$dMaybe.Nothing))(0);
};
export {
  all,
  alterAt,
  any,
  appendArray,
  catMaybes,
  concat,
  concatMap,
  cons,
  cons$p,
  $$delete as delete,
  deleteAt,
  deleteBy,
  difference,
  difference$p,
  drop,
  dropEnd,
  dropWhile,
  elem,
  elemIndex,
  elemLastIndex,
  filter,
  filterA,
  find,
  findIndex,
  findLastIndex,
  findMap,
  fold1,
  foldM,
  foldMap1,
  foldRecM,
  foldl1,
  foldr1,
  fromArray,
  fromFoldable,
  fromFoldable1,
  fromNonEmpty,
  group,
  groupAll,
  groupAllBy,
  groupBy,
  head,
  index,
  init,
  insert,
  insertAt,
  insertBy,
  intercalate,
  intercalate1,
  intersect,
  intersect$p,
  intersectBy,
  intersectBy$p,
  intersperse,
  last,
  length,
  mapMaybe,
  mapWithIndex,
  modifyAt,
  modifyAtIndices,
  notElem,
  nub,
  nubBy,
  nubByEq,
  nubEq,
  partition,
  prependArray,
  range,
  replicate,
  reverse,
  scanl,
  scanr,
  singleton,
  slice,
  snoc,
  snoc$p,
  some,
  sort,
  sortBy,
  sortWith,
  span,
  splitAt,
  tail,
  take,
  takeEnd,
  takeWhile,
  toArray,
  toNonEmpty,
  toUnfoldable,
  toUnfoldable1,
  transpose,
  transpose$p,
  uncons,
  union,
  union$p,
  unionBy,
  unionBy$p,
  unsafeIndex,
  unsnoc,
  unzip,
  updateAt,
  updateAtIndices,
  zip,
  zipWith,
  zipWithA
};
