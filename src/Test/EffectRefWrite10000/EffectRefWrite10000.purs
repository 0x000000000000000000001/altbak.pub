module Test.EffectRefWrite10000 (prepare) where

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
    writeEffect 1 n ref
    Ref.read ref

writeEffect :: Int -> Int -> Ref.Ref Int -> Effect Unit
writeEffect i n ref
  | i > n = pure unit
  | otherwise = do
      Ref.write i ref
      writeEffect (i + 1) n ref
