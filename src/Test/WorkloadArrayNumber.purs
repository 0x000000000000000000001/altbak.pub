module Test.WorkloadArrayNumber where

-- Purust-only workload, not part of the published kernel table.
-- Replay: --test WorkloadArrayNumber --expected 2147483647
-- (about 17998 us on the current compiler).
--
-- Same algorithm as WorkloadArrayInt. There is no unboxed representation for Array Number, so the pair quantifies the gap a NumberArray variant would close.

import Prelude

import Bench as Bench
import Data.Array (range, unsafeIndex)
import Data.Array as Array
import Data.Int as Int
import Effect (Effect)
import Effect.Console (log)
import Partial.Unsafe (unsafePartial)

-- Same algorithm as `WorkloadArrayInt`, on `Number` elements. There is no
-- unboxed representation for `Array Number`, so the pair quantifies the gap
-- that a `NumberArray` variant would have to close.

build :: Int -> Array Number
build n = map (\i -> Int.toNumber i * 3.0 + 1.0) (range 1 n)

sumValues :: Array Number -> Number
sumValues xs = Array.foldl (+) 0.0 xs

sumFiltered :: Array Number -> Number
sumFiltered xs = Array.foldl (+) 0.0 (Array.filter (\x -> x `mod` 2.0 == 0.0) xs)

indexedReads :: Array Number -> Int -> Int -> Int -> Number
indexedReads source size count start = go count start 0.0
  where
  go remaining index checksum
    | remaining == 0 = checksum
    | otherwise =
        let value = unsafePartial (unsafeIndex source index)
            next = if index + 1 == size then 0 else index + 1
        in go (remaining - 1) next (checksum + value)

describe :: Effect Unit
describe = log "Workload: Array Number (map, fold, filter+fold, indexed reads):"

act :: Effect Int
act = do
  n <- Bench.opaque 100000
  reads <- Bench.opaque 8388608
  let xs = build n
  pure (Int.floor (sumValues xs + sumFiltered xs + indexedReads xs 1024 reads 5))
