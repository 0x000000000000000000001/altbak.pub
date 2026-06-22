module Test.TCO where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

-- A deeply recursive function (Tail Recursive)
-- In PureScript to JS, it will be transformed into a 'while' loop
-- In Erlang/Scheme, it relies on native VM optimization
deepTailRec :: Int -> Int -> Int
deepTailRec 0 acc = acc
deepTailRec n acc = deepTailRec (n - 1) (acc + (n `mod` 3))

describe :: Effect Unit
describe = log "Tail Call Optimization (10M calls):"

act :: Effect Unit
act = logShow $ deepTailRec 10000000 0
