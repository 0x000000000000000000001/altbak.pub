-- @inline PerceusProbe.inspect never
module PerceusProbe where

data Shade = Red | Black
data Tree = Empty | Branch Shade Tree Int Tree

inspect :: Tree -> Int
inspect (Branch Black (Branch Red _ _ _) _ _) = 11
inspect (Branch _ _ _ (Branch Black _ _ _)) = 22
inspect _ = 33
