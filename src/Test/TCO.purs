module Test.TCO where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

-- Une fonction profondément récursive (Tail Recursive)
-- En PureScript vers JS, elle sera transformée en boucle 'while'
-- En Erlang/Scheme, elle reposera sur l'optimisation native de la VM
deepTailRec :: Int -> Int -> Int
deepTailRec 0 acc = acc
deepTailRec n acc = deepTailRec (n - 1) (acc + (n `mod` 3))

describe :: Effect Unit
describe = log "Tail Call Optimization (10M calls):"

act :: Effect Unit
act = logShow $ deepTailRec 10000000 0
