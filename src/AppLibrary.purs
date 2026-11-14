module AppLibrary where

import Prelude

import Bench (runBench)
import Data.Foldable (for_)
import Data.Traversable (traverse)
import Effect (Effect)
import Effect.Console (log)
import Test.ArrayMap10000 as ArrayMap10000
import Test.ArrayFilter10000 as ArrayFilter10000
import Test.ArrayFoldlSum10000 as ArrayFoldlSum10000
import Test.ArraySortDescending10000 as ArraySortDescending10000
import Test.ArrayEqNeqLast10000 as ArrayEqNeqLast10000
import Test.ArrayAppend1000 as ArrayAppend1000
import Test.TraverseMaybe1000 as TraverseMaybe1000
import Test.TraverseMaybeFailFirst1000 as TraverseMaybeFailFirst1000
import Test.TraverseMaybeFailLast1000 as TraverseMaybeFailLast1000
import Test.StringCodeUnitsLengthUnicode1000 as StringCodeUnitsLengthUnicode1000
import Test.StringCodePointsLengthUnicode1000 as StringCodePointsLengthUnicode1000
import Test.StringIndexOfHit10000 as StringIndexOfHit10000
import Test.StringIndexOfMiss10000 as StringIndexOfMiss10000
import Test.StringTakeCodePointsUnicode1500 as StringTakeCodePointsUnicode1500
import Test.StringSplit1000 as StringSplit1000
import Test.StringConcat1000 as StringConcat1000
import Test.RecordSetSmall1000 as RecordSetSmall1000
import Test.RecordSetMedium1000 as RecordSetMedium1000
import Test.RecordSetLarge1000 as RecordSetLarge1000
import Test.STRefWrite10000 as STRefWrite10000
import Test.STRefModify10000 as STRefModify10000
import Test.EffectRefWrite10000 as EffectRefWrite10000
import Test.EffectRefModify10000 as EffectRefModify10000

type Case = { label :: String, prepare :: Effect (Effect Int) }

-- Keep this order and these labels aligned with ../bin/benchmark/library-cases.json.
-- Expected values are independent oracles rather than results learned from a warm-up.
cases :: Array Case
cases =
  [ { label: "array/map-10000", prepare: ArrayMap10000.prepare }
  , { label: "array/filter-10000", prepare: ArrayFilter10000.prepare }
  , { label: "array/foldl-sum-10000", prepare: ArrayFoldlSum10000.prepare }
  , { label: "array/sort-descending-10000", prepare: ArraySortDescending10000.prepare }
  , { label: "array/eq-neq-last-10000", prepare: ArrayEqNeqLast10000.prepare }
  , { label: "array/append-1000", prepare: ArrayAppend1000.prepare }
  , { label: "foldable/traverse-maybe-1000", prepare: TraverseMaybe1000.prepare }
  , { label: "foldable/traverse-maybe-fail-first-1000", prepare: TraverseMaybeFailFirst1000.prepare }
  , { label: "foldable/traverse-maybe-fail-last-1000", prepare: TraverseMaybeFailLast1000.prepare }
  , { label: "string/cu-length-unicode-1000", prepare: StringCodeUnitsLengthUnicode1000.prepare }
  , { label: "string/cp-length-unicode-1000", prepare: StringCodePointsLengthUnicode1000.prepare }
  , { label: "string/indexOf-hit-10000", prepare: StringIndexOfHit10000.prepare }
  , { label: "string/indexOf-miss-10000", prepare: StringIndexOfMiss10000.prepare }
  , { label: "string/cp-take-unicode-1500", prepare: StringTakeCodePointsUnicode1500.prepare }
  , { label: "string/split-1000", prepare: StringSplit1000.prepare }
  , { label: "string/concat-1000", prepare: StringConcat1000.prepare }
  , { label: "record/set-small-1000", prepare: RecordSetSmall1000.prepare }
  , { label: "record/set-medium-1000", prepare: RecordSetMedium1000.prepare }
  , { label: "record/set-large-1000", prepare: RecordSetLarge1000.prepare }
  , { label: "st/ref-write-10000", prepare: STRefWrite10000.prepare }
  , { label: "st/ref-modify-10000", prepare: STRefModify10000.prepare }
  , { label: "effect/ref-write-10000", prepare: EffectRefWrite10000.prepare }
  , { label: "effect/ref-modify-10000", prepare: EffectRefModify10000.prepare }
  ]

main :: Effect Unit
main = do
  log "Preparing library fixtures outside the measured actions..."
  prepared <- traverse (\item -> do
    act <- item.prepare
    pure { label: item.label, act }) cases
  log "Global warm-up in progress..."
  for_ [1, 2, 3] \_ ->
    for_ prepared \item -> void item.act
  log "Global warm-up complete. Starting library benchmarks..."
  for_ prepared \item -> void $ runBench (log item.label) item.act
