module Test.Church where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

-- Church Numerals represents integers entirely as functions.
-- This heavily benchmarks currying, closure allocation,
-- and higher-order function applications.
type Church a = (a -> a) -> a -> a

zeroC :: forall a. Church a
zeroC _ x = x

succC :: forall a. Church a -> Church a
succC n f x = f (n f x)

addC :: forall a. Church a -> Church a -> Church a
addC m n f x = m f (n f x)

mulC :: forall a. Church a -> Church a -> Church a
mulC m n f x = m (n f) x

fromInt :: Int -> Church Int
fromInt 0 = zeroC
fromInt n = succC (fromInt (n - 1))

toInt :: Church Int -> Int
toInt n = n (\x -> x + 1) 0

c10 :: Int -> Church Int
c10 n = fromInt n

c100 :: Int -> Church Int
c100 n = mulC (c10 n) (c10 n)

c10k :: Int -> Church Int
c10k n = mulC (c100 n) (c100 n)

c100k :: Int -> Church Int
c100k n = mulC (c10k n) (c10 n)

describe :: Effect Unit
describe = log "Church Numerals (100k Closure Applications):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 10
  logShow $ toInt (c100k dummy)
