module Test.RecordSetSmall1000 (prepare) where

import Prelude

import Bench as Bench
import Effect (Effect)

type Small = { a :: Int, b :: Int }

sumSmall :: Small -> Int
sumSmall r = r.a + r.b

updateSmall :: Int -> Small -> Small
updateSmall 0 r = r
updateSmall n r = updateSmall (n - 1) (r { b = r.b + 1 })

-- Consume both the original and final records; each invocation starts
-- from the same immutable fixture. Scalar replacement remains allowed.
prepare :: Effect (Effect Int)
prepare = do
  fixture <- Bench.opaque { a: 1, b: 2 }
  pure do
    input <- Bench.opaque fixture
    count <- Bench.opaque 1000
    let result = updateSmall count input
    pure (sumSmall result + sumSmall input)
