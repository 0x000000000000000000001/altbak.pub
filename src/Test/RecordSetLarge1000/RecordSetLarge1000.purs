module Test.RecordSetLarge1000 (prepare) where

import Prelude

import Bench as Bench
import Effect (Effect)

type Large =
  { a :: Int, b :: Int, c :: Int, d :: Int
  , e :: Int, f :: Int, g :: Int, h :: Int
  , i :: Int, j :: Int, k :: Int, l :: Int
  , m :: Int, n :: Int, o :: Int, p :: Int
  }

sumLarge :: Large -> Int
sumLarge r =
  r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h
    + r.i + r.j + r.k + r.l + r.m + r.n + r.o + r.p

updateLarge :: Int -> Large -> Large
updateLarge 0 r = r
updateLarge n r = updateLarge (n - 1) (r { p = r.p + 1 })

-- Consume both the original and final records; each invocation starts
-- from the same immutable fixture. Scalar replacement remains allowed.
prepare :: Effect (Effect Int)
prepare = do
  fixture <- Bench.opaque
    { a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8
    , i: 9, j: 10, k: 11, l: 12, m: 13, n: 14, o: 15, p: 16
    }
  pure do
    input <- Bench.opaque fixture
    count <- Bench.opaque 1000
    let result = updateLarge count input
    pure (sumLarge result + sumLarge input)
