module Bench.Extended (runBenchSync, runBenchAff) where

import Prelude

import Bench (benchNow, formatNumber)
import Effect (Effect)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)
import Effect.Console (log)

foreign import consumeResult :: String -> String -> Effect Unit

startReport :: Effect Unit -> Effect Unit
startReport describe = do
  log "--------------------------------------------------\n\n(Test)\n"
  describe
  log "\n(Output & Warm-up)\n"

finishReport :: Number -> Effect Number
finishReport best = do
  log ("\n(Execution time - best of 10)\n\n" <> formatNumber best <> " μs\n")
  pure best

measureSync :: String -> Effect String -> Effect Number
measureSync expected act = do
  start <- benchNow
  result <- act
  consumeResult expected result
  finish <- benchNow
  pure (finish - start)

measureAff :: String -> Aff String -> Aff Number
measureAff expected act = do
  start <- liftEffect benchNow
  result <- act
  liftEffect $ consumeResult expected result
  finish <- liftEffect benchNow
  pure (finish - start)

runBenchSync :: Effect Unit -> Effect String -> Effect Number
runBenchSync describe act = do
  startReport describe
  out <- act
  log out
  consumeResult out out
  act >>= consumeResult out
  act >>= consumeResult out
  d1 <- measureSync out act
  d2 <- measureSync out act
  d3 <- measureSync out act
  d4 <- measureSync out act
  d5 <- measureSync out act
  d6 <- measureSync out act
  d7 <- measureSync out act
  d8 <- measureSync out act
  d9 <- measureSync out act
  d10 <- measureSync out act
  finishReport $ min (min (min (min d1 d2) (min d3 d4)) (min (min d5 d6) (min d7 d8))) (min d9 d10)

runBenchAff :: Effect Unit -> Aff String -> Aff Number
runBenchAff describe act = do
  liftEffect $ startReport describe
  out <- act
  liftEffect $ log out
  liftEffect $ consumeResult out out
  act >>= (liftEffect <<< consumeResult out)
  act >>= (liftEffect <<< consumeResult out)
  d1 <- measureAff out act
  d2 <- measureAff out act
  d3 <- measureAff out act
  d4 <- measureAff out act
  d5 <- measureAff out act
  d6 <- measureAff out act
  d7 <- measureAff out act
  d8 <- measureAff out act
  d9 <- measureAff out act
  d10 <- measureAff out act
  liftEffect $ finishReport $ min (min (min (min d1 d2) (min d3 d4)) (min (min d5 d6) (min d7 d8))) (min d9 d10)
