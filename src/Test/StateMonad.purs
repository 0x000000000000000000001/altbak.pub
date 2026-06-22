module Test.StateMonad where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

-- A pure, custom implementation of the State Monad.
-- This relies purely on function closures rather than data structures.
newtype State s a = State (s -> { val :: a, state :: s })

runState :: forall s a. State s a -> s -> { val :: a, state :: s }
runState (State f) s = f s

-- The Bind operation creates a deeply nested closure
bindState :: forall s a b. State s a -> (a -> State s b) -> State s b
bindState (State f) g = State \s ->
  let r1 = f s
      State g' = g r1.val
  in g' r1.state

pureState :: forall s a. a -> State s a
pureState a = State \s -> { val: a, state: s }

get :: forall s. State s s
get = State \s -> { val: s, state: s }

put :: forall s. s -> State s Unit
put s = State \_ -> { val: unit, state: s }

modify :: forall s. (s -> s) -> State s Unit
modify f = bindState get \s -> put (f s)

-- Builds a chained sequence of State modifications.
-- Resolving this dives deep into the Call Stack (depth = n).
-- We cap n at 6000 to avoid 'Maximum call stack size exceeded' in V8.
chainModifications :: Int -> State Int Unit
chainModifications 0 = pureState unit
chainModifications n = bindState (modify (\x -> x + 1)) \_ -> chainModifications (n - 1)

-- We run the 6000-depth State monad 2000 times to accumulate work.
-- Total: 12 million closures allocated and deeply evaluated.
runManyTimes :: Int -> Int -> Int
runManyTimes 0 acc = acc
runManyTimes n acc = runManyTimes (n - 1) (acc + (runState (chainModifications 6000) 0).state)

describe :: Effect Unit
describe = log "State Monad (12M Binds, 6k Stack Depth):"

-- The result should be 2000 * 6000 = 12000000
act :: Effect Unit
act = logShow $ runManyTimes 2000 0
