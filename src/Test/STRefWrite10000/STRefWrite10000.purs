module Test.STRefWrite10000 (prepare) where

import Prelude

import Bench as Bench
import Control.Monad.ST (ST)
import Control.Monad.ST as ST
import Control.Monad.ST.Ref as STRef
import Effect (Effect)

-- Allocate a fresh reference on every invocation. Allocation, all 10,000
-- updates and the final read are timed; only the bound is prepared.
prepare :: Effect (Effect Int)
prepare = do
  count <- Bench.opaque 10000
  pure do
    n <- Bench.opaque count
    pure $ ST.run do
      ref <- STRef.new 0
      writeST 1 n ref
      STRef.read ref

writeST :: forall r. Int -> Int -> STRef.STRef r Int -> ST r Unit
writeST i n ref
  | i > n = pure unit
  | otherwise = do
      _ <- STRef.write i ref
      writeST (i + 1) n ref
