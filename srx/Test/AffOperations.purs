module Test.AffOperations where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Effect.Aff (Aff, delay, Milliseconds(..))
import Effect.Class (liftEffect)
import Data.Int (toNumber)
import Bench as Bench

describe :: Effect Unit
describe = log "Aff Operations (Asynchronous Delays)"

-- This measures a 10 ms timer delay and its scheduling overhead, not the
-- throughput of arbitrary asynchronous computations.
act :: Aff String
act = do
  milliseconds <- liftEffect $ Bench.opaque 10
  _ <- delay (Milliseconds (toNumber milliseconds))
  pure (show milliseconds)
