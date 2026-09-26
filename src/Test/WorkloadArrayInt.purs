module Test.WorkloadArrayInt where

-- Purust-only workload, not part of the published kernel table.
-- Replay: --test WorkloadArrayInt --expected 35406173408
-- (about 7485 us on the current compiler).
--
-- Paired with WorkloadArrayNumber: same algorithm, other element type.

import Prelude

import Bench as Bench
import Data.Array (range, unsafeIndex)
import Data.Array as Array
import Effect (Effect)
import Effect.Console (log)
import Partial.Unsafe (unsafePartial)

-- Purust-only workload (not part of the published kernel table). Paired with
--
-- Purust-only workload: not part of the published kernel table.
-- Replay with: --test WorkloadArrayInt --expected 35406173408
-- (measured about 7380 us on the current compiler).-- `WorkloadArrayNumber`, which runs the same algorithm on `Number` elements:
-- the ratio measures what unboxed arrays buy on a real code path.

build :: Int -> Array Int
build n = map (\i -> i * 3 + 1) (range 1 n)

sumValues :: Array Int -> Int
sumValues xs = Array.foldl (+) 0 xs

sumFiltered :: Array Int -> Int
sumFiltered xs = Array.foldl (+) 0 (Array.filter (\x -> x `mod` 2 == 0) xs)

indexedReads :: Array Int -> Int -> Int -> Int -> Int
indexedReads source size count start = go count start 0
  where
  go remaining index checksum
    | remaining == 0 = checksum
    | otherwise =
        let value = unsafePartial (unsafeIndex source index)
            next = if index + 1 == size then 0 else index + 1
        in go (remaining - 1) next (checksum + value)

describe :: Effect Unit
describe = log "Workload: Array Int (map, fold, filter+fold, indexed reads):"

act :: Effect Int
act = do
  n <- Bench.opaque 100000
  reads <- Bench.opaque 8388608
  let xs = build n
  pure (sumValues xs + sumFiltered xs + indexedReads xs 1024 reads 5)
