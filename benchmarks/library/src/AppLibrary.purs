module AppLibrary where

import Prelude

import Bench (runBench)
import Data.Foldable (for_)
import Data.Traversable (traverse)
import Effect (Effect)
import Effect.Console (log)
import Test.Library.Arrays as Arrays
import Test.Library.Records as Records
import Test.Library.Refs as Refs
import Test.Library.Strings as Strings
import Test.Library.Traversal as Traversal

type Case = { label :: String, prepare :: Effect (Effect Int) }

-- Keep this order and these labels aligned with ../cases.json, whose expected
-- values are independent oracles rather than results learned from a warm-up.
cases :: Array Case
cases =
  [ { label: "array/map-10000", prepare: Arrays.prepareMap }
  , { label: "array/filter-10000", prepare: Arrays.prepareFilter }
  , { label: "array/foldl-sum-10000", prepare: Arrays.prepareFoldl }
  , { label: "array/sort-descending-10000", prepare: Arrays.prepareSort }
  , { label: "array/eq-neq-last-10000", prepare: Arrays.prepareEqNeq }
  , { label: "array/append-1000", prepare: Arrays.prepareAppend }
  , { label: "foldable/traverse-maybe-1000", prepare: Traversal.prepareSuccess }
  , { label: "foldable/traverse-maybe-fail-first-1000", prepare: Traversal.prepareFailFirst }
  , { label: "foldable/traverse-maybe-fail-last-1000", prepare: Traversal.prepareFailLast }
  , { label: "string/cu-length-unicode-1000", prepare: Strings.prepareCodeUnitsLength }
  , { label: "string/cp-length-unicode-1000", prepare: Strings.prepareCodePointsLength }
  , { label: "string/indexOf-hit-10000", prepare: Strings.prepareIndexHit }
  , { label: "string/indexOf-miss-10000", prepare: Strings.prepareIndexMiss }
  , { label: "string/cp-take-unicode-1500", prepare: Strings.prepareTakeCodePoints }
  , { label: "string/split-1000", prepare: Strings.prepareSplit }
  , { label: "string/concat-1000", prepare: Strings.prepareConcat }
  , { label: "record/set-small-1000", prepare: Records.prepareSetSmall }
  , { label: "record/set-medium-1000", prepare: Records.prepareSetMedium }
  , { label: "record/set-large-1000", prepare: Records.prepareSetLarge }
  , { label: "st/ref-write-10000", prepare: Refs.prepareSTWrite }
  , { label: "st/ref-modify-10000", prepare: Refs.prepareSTModify }
  , { label: "effect/ref-write-10000", prepare: Refs.prepareEffectWrite }
  , { label: "effect/ref-modify-10000", prepare: Refs.prepareEffectModify }
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
