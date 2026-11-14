module Test.EffectRefWrite10000 (prepare) where

import Prelude

import Bench as Bench
import Control.Monad.Rec.Class (Step(..), tailRecM)
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
writeEffect start n ref = tailRecM step start
  where
  step i
    | i > n = pure (Done unit)
    | otherwise = do
        Ref.write i ref
        pure (Loop (i + 1))
