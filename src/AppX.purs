module AppX where

import Prelude

import App as App
import Effect (Effect)
import Test.AffOperations as AffOperations
import Test.FileOps as FileOps
import Test.STArray as STArray
import Test.StringOps as StringOps

main :: Effect Unit
main = do
  App.main

  FileOps.describe
  FileOps.act

  STArray.describe
  STArray.act

  StringOps.describe
  StringOps.act

  AffOperations.describe
  AffOperations.act
