module Test.AffOperations where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Effect.Aff (launchAff_, delay, Milliseconds(..))
import Effect.Class (liftEffect)

describe :: Effect Unit
describe = log "Aff Operations (Asynchronous Delays):"

-- | A simple test that launches an Aff block with a delay,
-- | proving that the asynchronous runtime (or its synchronous fallback in PHP) is functioning.
act :: Effect Unit
act = launchAff_ do
  _ <- delay (Milliseconds 10.0)
  liftEffect $ log "10"
