module Test.LazyEvaluation where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

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

-- Evaluate the 4000-deep thunk chain 4000 times.
-- This allocates and evaluates 16 million thunks dynamically.
runManyTimes :: Int -> Int -> Int
runManyTimes 0 acc = acc
runManyTimes n acc = runManyTimes (n - 1) (acc + force (buildThunks 4000 (defer \_ -> 0)))

describe :: Effect Unit
describe = log "Lazy Evaluation (16M Thunks Forced, 4k Depth):"

-- The result should be 4000 * 4000 = 16000000
act :: Effect Unit
act = logShow $ runManyTimes 4000 0
