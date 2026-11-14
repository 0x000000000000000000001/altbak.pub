module Test.Library.Records where

import Prelude

import Bench as Bench
import Effect (Effect)

type Small = { a :: Int, b :: Int }

type Medium =
  { a :: Int, b :: Int, c :: Int, d :: Int
  , e :: Int, f :: Int, g :: Int, h :: Int
  }

type Large =
  { a :: Int, b :: Int, c :: Int, d :: Int
  , e :: Int, f :: Int, g :: Int, h :: Int
  , i :: Int, j :: Int, k :: Int, l :: Int
  , m :: Int, n :: Int, o :: Int, p :: Int
  }

sumSmall :: Small -> Int
sumSmall r = r.a + r.b

sumMedium :: Medium -> Int
sumMedium r = r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h

sumLarge :: Large -> Int
sumLarge r =
  r.a + r.b + r.c + r.d + r.e + r.f + r.g + r.h
    + r.i + r.j + r.k + r.l + r.m + r.n + r.o + r.p

updateSmall :: Int -> Small -> Small
updateSmall 0 r = r
updateSmall n r = updateSmall (n - 1) (r { b = r.b + 1 })

updateMedium :: Int -> Medium -> Medium
updateMedium 0 r = r
updateMedium n r = updateMedium (n - 1) (r { h = r.h + 1 })

updateLarge :: Int -> Large -> Large
updateLarge 0 r = r
updateLarge n r = updateLarge (n - 1) (r { p = r.p + 1 })

-- Include the original record in the checksum as well as the final version.
-- Each invocation reuses the immutable fixture and starts from the same value.
-- Scalar replacement and other semantics-preserving optimizations are allowed.
prepareSetSmall :: Effect (Effect Int)
prepareSetSmall = do
  fixture <- Bench.opaque { a: 1, b: 2 }
  pure do
    input <- Bench.opaque fixture
    count <- Bench.opaque 1000
    let result = updateSmall count input
    pure (sumSmall result + sumSmall input)

prepareSetMedium :: Effect (Effect Int)
prepareSetMedium = do
  fixture <- Bench.opaque { a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8 }
  pure do
    input <- Bench.opaque fixture
    count <- Bench.opaque 1000
    let result = updateMedium count input
    pure (sumMedium result + sumMedium input)

prepareSetLarge :: Effect (Effect Int)
prepareSetLarge = do
  fixture <- Bench.opaque
    { a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8
    , i: 9, j: 10, k: 11, l: 12, m: 13, n: 14, o: 15, p: 16
    }
  pure do
    input <- Bench.opaque fixture
    count <- Bench.opaque 1000
    let result = updateLarge count input
    pure (sumLarge result + sumLarge input)
