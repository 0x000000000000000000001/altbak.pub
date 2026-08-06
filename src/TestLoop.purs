module TestLoop where
import Prelude
loop :: Int -> Int
loop 0 = 0
loop n = loop (n - 1)
