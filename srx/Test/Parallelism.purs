module Test.Parallelism where

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Aff (Aff, delay, Milliseconds(..), forkAff, joinFiber)
import Data.Traversable (traverse)
import Data.Array (replicate)
import Data.Foldable (foldl)
import Effect.Class (liftEffect)
import Bench as Bench

describe :: Effect Unit
describe = log "Parallelism (10 x Fib 42)"

fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib n = fib (n - 1) + fib (n - 2)

heavyTask :: Int -> Aff Int
heavyTask n = do
  -- We yield the fiber to ensure the computation happens inside the spawned goroutine
  _ <- delay (Milliseconds 0.0)
  -- Each spawned task obtains its own opaque input after yielding. The result
  -- cannot be shared from the warm-up or computed before the task starts.
  input <- liftEffect $ Bench.opaque n
  pure (fib input)

-- Each addition is at most 1,000,000,006 + fib(42), within signed 32-bit Int.
-- The checksum therefore agrees across JS and native integer representations.
checksum :: Array Int -> Int
checksum = foldl (\acc value -> mod (acc + value) 1000000007) 0

act :: Aff String
act = do
  -- We launch heavy tasks in parallel using forkAff
  fibers <- traverse (\_ -> forkAff (heavyTask 42)) (replicate 10 unit)

  -- Wait for all of them to complete and collect results
  results <- traverse joinFiber fibers

  -- Use the result so the compiler doesn't optimize it away
  pure $ "Checksum: " <> show (checksum results)
