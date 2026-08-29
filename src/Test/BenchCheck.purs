module Test.BenchCheck where
import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench (benchNow)

act :: Effect String
act = do
  t1 <- benchNow
  t2 <- benchNow
  pure ( "Delta: " <> show (t2 - t1))
