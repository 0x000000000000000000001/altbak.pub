module Test.Native where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import foldlArray :: forall a b. (b -> a -> b) -> b -> Array a -> b
foreign import arrayMap :: forall a b. (a -> b) -> Array a -> Array b
foreign import filter :: forall a. (a -> Boolean) -> Array a -> Array a

sumNative :: Array Int -> Int
sumNative arr = foldlArray (+) 0 arr

describe :: Effect Unit
describe = log "Native Intrinsics Test (900 elements):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 900
  let 
    arr1 = arrayMap (\x -> x) [1,2,3]
    arr2 = arrayMap (\x -> x * 2) arr1
    arr3 = filter (\x -> x > 0) arr2
  logShow $ sumNative arr3
