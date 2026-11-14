module Test.Library.Refs
  ( prepareSTWrite
  , prepareSTModify
  , prepareEffectWrite
  , prepareEffectModify
  ) where

import Prelude

import Bench as Bench
import Control.Monad.ST (ST)
import Control.Monad.ST as ST
import Control.Monad.ST.Ref as STRef
import Effect (Effect)
import Effect.Ref as Ref

-- Every invocation allocates a fresh reference. The timed work includes its
-- allocation, all 10,000 updates and the final read; only the bound is prepared.
prepareSTWrite :: Effect (Effect Int)
prepareSTWrite = do
  count <- Bench.opaque 10000
  pure do
    n <- Bench.opaque count
    pure $ ST.run do
      ref <- STRef.new 0
      writeST 1 n ref
      STRef.read ref

prepareSTModify :: Effect (Effect Int)
prepareSTModify = do
  count <- Bench.opaque 10000
  pure do
    n <- Bench.opaque count
    pure $ ST.run do
      ref <- STRef.new 0
      modifyST n ref
      STRef.read ref

prepareEffectWrite :: Effect (Effect Int)
prepareEffectWrite = do
  count <- Bench.opaque 10000
  pure do
    n <- Bench.opaque count
    ref <- Ref.new 0
    writeEffect 1 n ref
    Ref.read ref

prepareEffectModify :: Effect (Effect Int)
prepareEffectModify = do
  count <- Bench.opaque 10000
  pure do
    n <- Bench.opaque count
    ref <- Ref.new 0
    modifyEffect n ref
    Ref.read ref

writeST :: forall r. Int -> Int -> STRef.STRef r Int -> ST r Unit
writeST i n ref
  | i > n = pure unit
  | otherwise = do
      _ <- STRef.write i ref
      writeST (i + 1) n ref

modifyST :: forall r. Int -> STRef.STRef r Int -> ST r Unit
modifyST remaining ref
  | remaining <= 0 = pure unit
  | otherwise = do
      _ <- STRef.modify (_ + 1) ref
      modifyST (remaining - 1) ref

writeEffect :: Int -> Int -> Ref.Ref Int -> Effect Unit
writeEffect i n ref
  | i > n = pure unit
  | otherwise = do
      Ref.write i ref
      writeEffect (i + 1) n ref

modifyEffect :: Int -> Ref.Ref Int -> Effect Unit
modifyEffect remaining ref
  | remaining <= 0 = pure unit
  | otherwise = do
      _ <- Ref.modify (_ + 1) ref
      modifyEffect (remaining - 1) ref
