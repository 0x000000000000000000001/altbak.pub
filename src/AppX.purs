module AppX where

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Class (liftEffect)
import Effect.Aff (launchAff_)
import Bench (formatNumber, runBenchAff, runBenchSync)
import Test.AffOperations as AffOperations
import Test.FileOps as FileOps
import Test.STArray as STArray
import Test.StringOps as StringOps
import Test.Parallelism as Parallelism

main :: Effect Unit
main = do
  -- extended tests
  ---- sync
  t15 <- runBenchSync FileOps.describe FileOps.act
  t16 <- runBenchSync STArray.describe STArray.act
  t17 <- runBenchSync StringOps.describe StringOps.act

  ---- async
  launchAff_ do
    t18 <- runBenchAff AffOperations.describe AffOperations.act
    t19 <- runBenchAff Parallelism.describe Parallelism.act

    let totalMs = (t15 / 1000.0) + (t16 / 1000.0) + (t17 / 1000.0) + (t18 / 1000.0) + (t19 / 1000.0)
    liftEffect $ log $ "\n==================================================\n\nTotal exec time: " <> formatNumber totalMs <> " ms\n"
