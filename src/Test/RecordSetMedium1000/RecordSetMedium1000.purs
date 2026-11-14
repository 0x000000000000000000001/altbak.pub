module Test.RecordSetMedium1000 (prepare) where

import Prelude

import Bench as Bench
import Effect (Effect)

type Medium =
  { a :: Int, b :: Int, c :: Int, d :: Int
  , e :: Int, f :: Int, g :: Int, h :: Int
  }

sumMedium :: Medium -> Int
sumMedium r = r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h

updateMedium :: Int -> Medium -> Medium
updateMedium 0 r = r
updateMedium n r = updateMedium (n - 1) (r { h = r.h + 1 })

-- Consume both the original and final records; each invocation starts
-- from the same immutable fixture. Scalar replacement remains allowed.
prepare :: Effect (Effect Int)
prepare = do
  fixture <- Bench.opaque { a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8 }
  pure do
    input <- Bench.opaque fixture
    count <- Bench.opaque 1000
    let result = updateMedium count input
    pure (sumMedium result + sumMedium input)
