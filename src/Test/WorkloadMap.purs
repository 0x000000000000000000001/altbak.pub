module Test.WorkloadMap where

-- Purust-only workload, not part of the published kernel table.
-- Replay: --test WorkloadMap --expected 400040000
-- (about 59856 us on the current compiler).
--
-- The shipped Data.Map. Measured: ~2.7 us per insert into a 20k map, with
-- ~161 allocations (8 for an empty map), while lookups stay ~170 ns. The
-- library stores sizes, so its algorithm is O(log n); the cost comes from
-- erasure boxing around its polymorphic helpers (`unsafeBalancedNode`,
-- `insertWith`): every node boxes subtrees into Value::Class and applies
-- eta closures. Removing that needs specialization (the monomorphisation
-- pass), not codegen cleanup: clone/dispatch tidy-ups measured flat.

import Prelude

import Bench as Bench
import Data.Array as Array
import Data.Foldable (sum)
import Data.Map as Map
import Data.Maybe (fromMaybe)
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Effect.Console (log)

-- Purust-only workload: `Data.Map` is a shipped library built on the same
--
-- Purust-only: not part of the published kernel table. Run with
-- +--test WorkloadMap --expected 400040000-- (measured ~59856 us on the current compiler;-- the oracle is exact, so a mismatch means the workload changed).
-- balanced-tree shape as the Red-Black Tree kernel, so insertion, lookup,
-- filtering and traversal all matter for real code.

describe :: Effect Unit
describe = log "Workload: Data.Map (build, lookups, filter, size):"

act :: Effect Int
act = do
  n <- Bench.opaque 20000
  let keys = Array.range 1 n
      table = Map.fromFoldable (map (\k -> Tuple k (k * 2)) keys) :: Map.Map Int Int
      lookups = map (\k -> fromMaybe 0 (Map.lookup k table)) keys
      filtered = Map.filter (\v -> v `mod` 2 == 0) table
  pure (Map.size filtered + sum lookups)
