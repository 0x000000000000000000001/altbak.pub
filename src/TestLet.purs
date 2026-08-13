module TestLet where

import Prelude

foreign import opaque :: Int -> Int

addOne :: Int -> Int
addOne x = 
  let 
    y = opaque x
    z = y + y
  in z + z
