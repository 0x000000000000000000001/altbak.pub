module Bench where

import Prelude
import Effect (Effect)
import Effect.Console (log)

foreign import benchNow :: Effect Number
foreign import opaque :: forall a. a -> Effect a
foreign import formatNumber :: Number -> String

runBench :: Effect Unit -> Effect String -> Effect Number
runBench describe act = do
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

  -- Note: We intentionally avoid using `Data.Foldable.foldr` or array literals here.
  -- This core benchmark suite runs against highly experimental and WIP AOT compilers 
  -- (like `purust`) that might crash on unboxed array literals or fail 
  -- to parse the massive CoreFn dependency graph brought in by Foldable/Array.
  -- We stick to primitive nested `min` calls to maintain maximum compiler compatibility.
  --
  -- Why take the `min` instead of average or median?
  -- In benchmarking, OS jitter, context switches, and background tasks only ever *add*
  -- noise and execution time. The minimum time recorded across multiple warm iterations
  -- represents the most accurate measurement of the algorithm's true baseline cost 
  -- with the least amount of external OS interference.
  let best = min (min (min (min d1 d2) (min d3 d4)) (min (min d5 d6) (min d7 d8))) (min d9 d10)

  log ("\n(Execution time - best of 10)\n\n" <> formatNumber best <> " μs\n")
  pure best


