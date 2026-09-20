module Bench.Extended (runBenchSync, runBenchAff) where

import Prelude

import Bench (benchNow, formatNumber)
import Effect (Effect)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)
import Effect.Console (log)

foreign import consumeResult :: String -> Effect Unit

startReport :: Effect Unit -> Effect Unit
startReport describe = do
  log "--------------------------------------------------\n\n(Test)\n"
  describe
  log "\n(Output & Warm-up)\n"

finishReport :: Number -> Effect Number
finishReport best = do
  log ("\n(Execution time - best of 10)\n\n" <> formatNumber best <> " μs\n")
  pure best

measureSync :: Effect String -> Effect Number
measureSync act = do
  start <- benchNow
  result <- act
  consumeResult result
  finish <- benchNow
  pure (finish - start)

measureAff :: Aff String -> Aff Number
measureAff act = do
  start <- liftEffect benchNow
  result <- act
  liftEffect $ consumeResult result
  finish <- liftEffect benchNow
  pure (finish - start)

runBenchSync :: Effect Unit -> Effect String -> Effect Number
runBenchSync describe act = do
  startReport describe
  out <- act
  log out
  consumeResult out
  act >>= consumeResult
  act >>= consumeResult
  d1 <- measureSync act
  d2 <- measureSync act
  d3 <- measureSync act
  d4 <- measureSync act
  d5 <- measureSync act
  d6 <- measureSync act
  d7 <- measureSync act
  d8 <- measureSync act
  d9 <- measureSync act
  d10 <- measureSync act
  finishReport $ min (min (min (min d1 d2) (min d3 d4)) (min (min d5 d6) (min d7 d8))) (min d9 d10)

runBenchAff :: Effect Unit -> Aff String -> Aff Number
runBenchAff describe act = do
  liftEffect $ startReport describe
  out <- act
  liftEffect $ log out
  liftEffect $ consumeResult out
  act >>= (liftEffect <<< consumeResult)
  act >>= (liftEffect <<< consumeResult)
  d1 <- measureAff act
  d2 <- measureAff act
  d3 <- measureAff act
  d4 <- measureAff act
  d5 <- measureAff act
  d6 <- measureAff act
  d7 <- measureAff act
  d8 <- measureAff act
  d9 <- measureAff act
  d10 <- measureAff act
  liftEffect $ finishReport $ min (min (min (min d1 d2) (min d3 d4)) (min (min d5 d6) (min d7 d8))) (min d9 d10)
