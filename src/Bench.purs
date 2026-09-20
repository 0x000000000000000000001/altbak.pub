module Bench where

import Prelude
import Effect (Effect)
import Effect.Console (log)

foreign import benchNow :: Effect Number
foreign import opaque :: forall a. a -> Effect a
foreign import formatNumber :: Number -> String

-- Returns microseconds per invocation. The native driver consumes every numeric result
-- and checks the final result against the separately validated warm-up value.
foreign import measureBatch :: Int -> Int -> Effect Int -> Effect Number

calibrate :: Int -> Effect Int -> Int -> Number -> Effect Int
calibrate expected act iterations target = do
  elapsed <- measureBatch iterations expected act
  if elapsed >= target || iterations >= 16777216
    then pure iterations
    else calibrate expected act (iterations * 2) (target / 2.0)

runBench :: Effect Unit -> Effect Int -> Effect Number
runBench describe act = do
  log "--------------------------------------------------\n\n(Test)\n"
  describe
  log "\n(Output & Warm-up)\n"
  out <- act
  log (show out)
  void act
  void act
  iterations <- calibrate out act 1 10000.0
  d1 <- measureBatch iterations out act
  d2 <- measureBatch iterations out act
  d3 <- measureBatch iterations out act
  d4 <- measureBatch iterations out act
  d5 <- measureBatch iterations out act
  d6 <- measureBatch iterations out act
  d7 <- measureBatch iterations out act
  d8 <- measureBatch iterations out act
  d9 <- measureBatch iterations out act
  d10 <- measureBatch iterations out act
  let best = min (min (min (min d1 d2) (min d3 d4)) (min (min d5 d6) (min d7 d8))) (min d9 d10)
  log ("\n(Execution time - best of 10)\n\n" <> formatNumber best <> " μs\n")
  log ("Batch iterations: " <> show iterations)
  pure best

runBenchSync :: Effect Unit -> Effect String -> Effect Number
runBenchSync describe act = do
  log "--------------------------------------------------\n\n(Test)\n"
  describe
  log "\n(Output & Warm-up)\n"
  
  -- Warm-up (3 runs)
  out <- act
  log out
  void act
  void act
  
  -- 10 measured runs
  t1 <- benchNow
  void act
  t2 <- benchNow
  let d1 = t2 - t1

  t3 <- benchNow
  void act
  t4 <- benchNow
  let d2 = t4 - t3

  t5 <- benchNow
  void act
  t6 <- benchNow
  let d3 = t6 - t5

  t7 <- benchNow
  void act
  t8 <- benchNow
  let d4 = t8 - t7

  t9 <- benchNow
  void act
  t10 <- benchNow
  let d5 = t10 - t9

  t11 <- benchNow
  void act
  t12 <- benchNow
  let d6 = t12 - t11

  t13 <- benchNow
  void act
  t14 <- benchNow
  let d7 = t14 - t13

  t15 <- benchNow
  void act
  t16 <- benchNow
  let d8 = t16 - t15

  t17 <- benchNow
  void act
  t18 <- benchNow
  let d9 = t18 - t17

  t19 <- benchNow
  void act
  t20 <- benchNow
  let d10 = t20 - t19

  let best = min (min (min (min d1 d2) (min d3 d4)) (min (min d5 d6) (min d7 d8))) (min d9 d10)

  log ("\n(Execution time - best of 10)\n\n" <> formatNumber best <> " μs\n")
  pure best
