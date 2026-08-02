module Bench where
import Prelude
import Effect (Effect)
import Effect.Console (log)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)

foreign import benchNow :: Effect Number
foreign import opaque :: forall a. a -> Effect a
foreign import formatNumber :: Number -> String
foreign import keepAlive :: Effect Unit

runBench :: Effect Unit -> Effect Unit -> Effect Number
runBench describe act = do
  describe
  t1 <- benchNow
  act
  t2 <- benchNow
  let dt = t2 - t1
  log ("\nExecution time: " <> formatNumber dt <> " μs\n")
  pure dt

runBenchAff :: Effect Unit -> Aff Unit -> Aff Number
runBenchAff describe act = do
  liftEffect describe
  t1 <- liftEffect benchNow
  act
  t2 <- liftEffect benchNow
  let dt = t2 - t1
  liftEffect $ log ("\nExecution time: " <> formatNumber dt <> " μs\n")
  pure dt
