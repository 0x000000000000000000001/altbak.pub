import * as $runtime from "../runtime.js";
const $Shape = (tag, _1, _2) => ({tag, _1, _2});
const Circle = value0 => $Shape("Circle", value0);
const Rect = value0 => value1 => $Shape("Rect", value0, value1);
const area = v => {
  if (v.tag === "Circle") { return v._1 * v._1 | 0; }
  if (v.tag === "Rect") { return v._1 * v._2 | 0; }
  $runtime.fail();
};
export {$Shape, Circle, Rect, area};
