module Test.STRefModify10000 (prepare) where

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
      modifyST n ref
      STRef.read ref

modifyST :: forall r. Int -> STRef.STRef r Int -> ST r Unit
modifyST remaining ref
  | remaining <= 0 = pure unit
  | otherwise = do
      _ <- STRef.modify (_ + 1) ref
      modifyST (remaining - 1) ref
