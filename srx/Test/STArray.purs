module Test.STArray where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Data.Array.ST as STArray
import Control.Monad.ST (run)
import Data.Maybe (Maybe(..))
import Bench as Bench

sumArray :: Int -> Int
sumArray lastValue = run do
  arr <- STArray.new
  _ <- STArray.pushAll [1, 2, 3, 4, 5, 6, 7, 8, 9, lastValue] arr
  x <- STArray.pop arr
  case x of
    Just v -> pure v
    Nothing -> pure 0

describe :: Effect Unit
describe = log "STArray Operations:"

act :: Effect String
act = do
  lastValue <- Bench.opaque 10
  pure (show $ sumArray lastValue)
