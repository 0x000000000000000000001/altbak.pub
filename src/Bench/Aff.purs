module Bench.Aff (runBenchAff) where

import Prelude

import Bench (benchNow, formatNumber)
import Effect (Effect)
import Effect.Aff (Aff)
import Effect.Class (liftEffect)
import Effect.Console (log)

runBenchAff :: Effect Unit -> Aff String -> Aff Number
runBenchAff describe act = do
  liftEffect $ log "--------------------------------------------------\n\n(Test)\n"
  liftEffect $ describe
  liftEffect $ log "\n(Output & Warm-up)\n"
  
  -- Warm-up (3 runs)
  out <- act
  liftEffect $ log out
  void act
  void act
  
  -- 10 measured runs
  t1 <- liftEffect benchNow
  void act
  t2 <- liftEffect benchNow
  let d1 = t2 - t1

  t3 <- liftEffect benchNow
  void act
  t4 <- liftEffect benchNow
  let d2 = t4 - t3

  t5 <- liftEffect benchNow
  void act
  t6 <- liftEffect benchNow
  let d3 = t6 - t5

  t7 <- liftEffect benchNow
  void act
  t8 <- liftEffect benchNow
  let d4 = t8 - t7

  t9 <- liftEffect benchNow
  void act
  t10 <- liftEffect benchNow
  let d5 = t10 - t9

  t11 <- liftEffect benchNow
  void act
  t12 <- liftEffect benchNow
  let d6 = t12 - t11

  t13 <- liftEffect benchNow
  void act
  t14 <- liftEffect benchNow
  let d7 = t14 - t13

  t15 <- liftEffect benchNow
  void act
  t16 <- liftEffect benchNow
  let d8 = t16 - t15

  t17 <- liftEffect benchNow
  void act
  t18 <- liftEffect benchNow
  let d9 = t18 - t17

  t19 <- liftEffect benchNow
  void act
  t20 <- liftEffect benchNow
  let d10 = t20 - t19

  let best = min (min (min (min d1 d2) (min d3 d4)) (min (min d5 d6) (min d7 d8))) (min d9 d10)

  liftEffect $ log ("\n(Execution time - best of 10)\n\n" <> formatNumber best <> " μs\n")
  pure best
