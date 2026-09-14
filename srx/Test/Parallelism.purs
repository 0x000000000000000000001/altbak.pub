module Test.Parallelism where

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Aff (Aff, delay, Milliseconds(..), forkAff, joinFiber)
import Data.Traversable (traverse)
import Data.Array (replicate)
import Data.Foldable (sum)

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
  -- The CPU heavy pure computation is returned to prevent DCE
  pure (fib n)

act :: Aff String
act = do
  -- We launch heavy tasks in parallel using forkAff
  fibers <- traverse (\_ -> forkAff (heavyTask 42)) (replicate 10 unit)

  -- Wait for all of them to complete and collect results
  results <- traverse joinFiber fibers

  -- Use the result so the compiler doesn't optimize it away
  pure $ "Sum of results: " <> show (sum results)
