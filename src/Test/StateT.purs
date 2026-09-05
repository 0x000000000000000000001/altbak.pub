module Test.StateT where

import Prelude
import Control.Monad.State.Trans (StateT, runStateT, get, put)
import Data.Identity (Identity)

testState :: Int -> Int -> Int -> StateT Int Identity Unit
testState n acc limit = do
  s <- get
  if s < limit
    then do
      put (s + n)
      testState n acc limit
    else pure unit

runTest :: Int -> Int
runTest limit = case runStateT (testState 1 0 limit) 0 of
  _ -> 42
