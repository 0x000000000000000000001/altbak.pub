module Test.TCO where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

-- A deeply recursive function (Tail Recursive)
-- In PureScript to JS, it will be transformed into a 'while' loop
-- In Erlang/Scheme, it relies on native VM optimization
deepTailRec :: Int -> Int -> Int
deepTailRec 0 acc = acc
deepTailRec n acc = deepTailRec (n - 1) (acc + (n `mod` 3))

describe :: Effect Unit
describe = log "Tail Call Optimization (100k calls):"

act :: Effect String
act = do
  dummy <- Bench.opaque 100000
  pure (show ( deepTailRec dummy 0))
