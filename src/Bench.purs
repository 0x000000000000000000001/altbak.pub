module Bench where
import Prelude
import Effect (Effect)
import Effect.Console (log)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)

foreign import benchNow :: Effect Number
foreign import opaque :: forall a. a -> Effect a
foreign import formatNumber :: Number -> String


runBench :: Effect Unit -> Effect Unit -> Effect Number
runBench describe act = do
  log "--------------------------------------------------\n\n(Test)\n"
  describe
  log "\n(Output)\n"
  t1 <- benchNow
  act
  t2 <- benchNow
  let dt = t2 - t1
  log ("\n(Execution time)\n\n" <> formatNumber dt <> " μs\n")
  pure dt

runBenchAff :: Effect Unit -> Aff Unit -> Aff Number
runBenchAff describe act = do
  liftEffect $ log "--------------------------------------------------\n\n(Test)\n"
  liftEffect describe
  liftEffect $ log "\n(Output)\n"
  t1 <- liftEffect benchNow
  act
  t2 <- liftEffect benchNow
  let dt = t2 - t1
  liftEffect $ log ("\n(Execution time)\n\n" <> formatNumber dt <> " μs\n")
  pure dt
