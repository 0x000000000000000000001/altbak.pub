module Test.Ackermann where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

-- The Ackermann function is deeply recursive but NOT tail-recursive.
-- It is a classic benchmark for testing function call overhead,
-- stack management, and context switching in VMs.
ackermann :: Int -> Int -> Int
ackermann 0 n = n + 1
ackermann m 0 = ackermann (m - 1) 1
ackermann m n = ackermann (m - 1) (ackermann m (n - 1))

describe :: Effect Unit
describe = log "Ackermann (3, 8):"

act :: Effect Unit
act = logShow $ ackermann 3 8
