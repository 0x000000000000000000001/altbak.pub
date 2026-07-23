module Bench where
import Prelude
import Effect (Effect)
import Effect.Console (log)

foreign import benchNow :: Effect Number
foreign import opaque :: forall a. a -> Effect a
foreign import formatNumber :: Number -> String

runBench :: Effect Unit -> Effect Unit -> Effect Number
runBench describe act = do
  describe
  t1 <- benchNow
  act
  t2 <- benchNow
  let dt = t2 - t1
  log ("\nExecution time: " <> formatNumber dt <> " μs\n")
  pure dt
