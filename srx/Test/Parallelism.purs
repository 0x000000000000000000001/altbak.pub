module Test.Parallelism where

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Aff (Aff, launchAff_, delay, Milliseconds(..), forkAff, joinFiber)
import Data.Traversable (traverse)
import Data.Array (replicate)

describe :: Effect Unit
describe = log "Parallelism (40 x Fib 35):"

fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib n = fib (n - 1) + fib (n - 2)

heavyTask :: Int -> Aff Unit
heavyTask n = do
  -- We yield the fiber to ensure the computation happens inside the spawned goroutine
  _ <- delay (Milliseconds 0.0)
  -- The CPU heavy pure computation
  let _res = fib n
  pure unit

act :: Effect Unit
act = launchAff_ do
  -- We launch heavy tasks in parallel using forkAff
  fibers <- traverse (\_ -> forkAff (heavyTask 35)) (replicate 200 unit)

  -- Wait for all of them to complete
  _ <- traverse joinFiber fibers

  pure unit
