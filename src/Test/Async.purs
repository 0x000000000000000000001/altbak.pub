module Test.Async where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

foreign import runAsyncTest :: Int -> Effect Unit

describe :: Effect Unit
describe = log "Asynchronous Concurrency (1000 forks):"

act :: Effect Unit
act = runAsyncTest 1000
