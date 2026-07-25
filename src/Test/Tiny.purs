module Test.Tiny where

import Prelude

data Shape = Circle Int | Rect Int Int

area :: Shape -> Int
area (Circle r) = r * r
area (Rect w h) = w * h
