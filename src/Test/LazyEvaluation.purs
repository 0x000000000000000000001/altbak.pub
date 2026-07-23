module Test.LazyEvaluation where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

-- PureScript is strict by default. We simulate Lazy Evaluation (Thunks)
-- by wrapping computations in functions taking `Unit`.
-- Forcing these thunks deeply tests memory allocation and function application overhead.
newtype Lazy a = Lazy (Unit -> a)

force :: forall a. Lazy a -> a
force (Lazy f) = f unit

defer :: forall a. (Unit -> a) -> Lazy a
defer f = Lazy f

-- Builds a chain of n thunks.
-- When forced, it will consume n frames on the call stack.
-- Kept at 4000 to prevent V8 stack overflow (limit ~10400).
buildThunks :: Int -> Lazy Int -> Lazy Int
buildThunks 0 acc = acc
buildThunks n acc = buildThunks (n - 1) (defer \_ -> force acc + 1)

-- Evaluate the 1000-deep thunk chain 1000 times.
-- This allocates and evaluates 1 million thunks dynamically.
runManyTimes :: Int -> Int -> Int
runManyTimes 0 acc = acc
runManyTimes n acc = runManyTimes (n - 1) (acc + force (buildThunks 1000 (defer \_ -> 0)))

describe :: Effect Unit
describe = log "Lazy Evaluation (1M Thunks Forced, 1k Depth):"

-- The result should be 1000 * 1000 = 1000000
act :: Effect Unit
act = do
  dummy <- Bench.opaque 1000
  logShow $ runManyTimes dummy 0
