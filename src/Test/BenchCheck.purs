module Test.BenchCheck where
import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench (benchNow)

act :: Effect Unit
act = do
  t1 <- benchNow
  t2 <- benchNow
  log $ "Delta: " <> show (t2 - t1)
