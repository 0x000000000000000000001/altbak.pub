module Test.WorkloadList where

-- Purust-only workload, not part of the published kernel table.
-- Replay: --test WorkloadList --expected 2500149997
-- (about 6949 us on the current compiler).
--
-- The shipped Data.List, which the kernel table never exercises (its list case declares its own type).

import Prelude

import Bench as Bench
import Data.List as List
import Effect (Effect)
import Effect.Console (log)

-- Purust-only workload: the library `Data.List`, which the kernel table never
--
-- Purust-only: not part of the published kernel table. Run with
-- +--test WorkloadList --expected 2500149997-- (measured ~6949 us on the current compiler;-- the oracle is exact, so a mismatch means the workload changed).
-- exercises (its list case declares its own list type). Build, filter, fold and
-- a filtered length all run through the shipped library module.

describe :: Effect Unit
describe = log "Workload: Data.List (range, filter, foldl, filtered length):"

act :: Effect Int
act = do
  n <- Bench.opaque 100000
  let xs = List.range 1 n
      evens = List.filter (\x -> x `mod` 2 == 0) xs
      counted = List.length (List.filter (\x -> x > 3) xs)
  pure (List.foldl (+) 0 evens + counted)
