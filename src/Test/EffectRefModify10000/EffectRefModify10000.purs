module Test.EffectRefModify10000 (prepare) where

import Prelude

import Bench as Bench
import Effect.Ref as Ref
import Effect (Effect)

-- Allocate a fresh reference on every invocation. Allocation, all 10,000
-- updates and the final read are timed; only the bound is prepared.
prepare :: Effect (Effect Int)
prepare = do
  count <- Bench.opaque 10000
  pure do
    n <- Bench.opaque count
    ref <- Ref.new 0
    modifyEffect n ref
    Ref.read ref

modifyEffect :: Int -> Ref.Ref Int -> Effect Unit
modifyEffect remaining ref
  | remaining <= 0 = pure unit
  | otherwise = do
      _ <- Ref.modify (_ + 1) ref
      modifyEffect (remaining - 1) ref
