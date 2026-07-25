module Test.Native where

import Prelude

loopNative :: Int -> Int -> Int
loopNative 0 acc = acc
loopNative n acc = loopNative (n - 1) (acc + 1)
