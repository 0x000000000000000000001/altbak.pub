module AppFFI where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench (runBench, formatNumber)
import Test.AstTreeFFI as AstTreeFFI
import Test.RBTreeFFI as RBTreeFFI

main :: Effect Unit
main = do
  t1 <- runBench AstTreeFFI.describe AstTreeFFI.act
  t2 <- runBench RBTreeFFI.describe RBTreeFFI.act
  let totalMs = (t1 / 1000.0) + (t2 / 1000.0)
  log $ "\n==================================================\n\nTotal exec time: " <> formatNumber totalMs <> " ms\n"
