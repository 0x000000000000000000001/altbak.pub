module Test.Church where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

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

-- Compute 10^5 (100,000) using function composition
c10 :: Church Int
c10 = fromInt 10

c100 :: Church Int
c100 = mulC c10 c10

c10k :: Church Int
c10k = mulC c100 c100

c100k :: Church Int
c100k = mulC c10k c10

describe :: Effect Unit
describe = log "Church Numerals (100k Closure Applications):"

act :: Effect Unit
act = logShow $ toInt c100k
