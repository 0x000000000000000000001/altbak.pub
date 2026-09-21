module Test.ArrayIndexing (BoxedArray, nativeReads, boxedReads) where

import Prelude

import Data.Array (unsafeIndex)
import Partial.Unsafe (unsafePartial)
import Unsafe.Coerce (unsafeCoerce)

-- The drivers supply the same integer values, outside the measured interval.
-- This opaque input models a decoded/FFI array kept in the runtime's boxed
-- representation. Only the selected element should be converted to Int.
foreign import data BoxedArray :: Type

nativeReads :: Array Int -> Int -> Int -> Int -> Int
nativeReads source size count start = go count start 0
  where
  go remaining index checksum
    | remaining == 0 = checksum
    | otherwise =
        let value = unsafePartial (unsafeIndex source index)
            next = if index + 1 == size then 0 else index + 1
        in go (remaining - 1) next (checksum + value)

boxedReads :: BoxedArray -> Int -> Int -> Int -> Int
boxedReads source size count start = go count start 0
  where
  go remaining index checksum
    | remaining == 0 = checksum
    | otherwise =
        let value = unsafePartial (unsafeIndex (unsafeCoerce source :: Array Int) index)
            next = if index + 1 == size then 0 else index + 1
        in go (remaining - 1) next (checksum + value)
