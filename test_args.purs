module TestArgs where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Data.Array as Array
import Data.String (Pattern(..))
import Data.String as String
import Data.Maybe (Maybe(..))
import Node.Process as Process

main :: Effect Unit
main = do
  argsRaw <- Process.argv
  let args = Array.concatMap (\s -> String.split (Pattern " ") s) argsRaw
  let getArg key = case Array.elemIndex key args of
        Just i -> Array.index args (i + 1)
        Nothing -> Nothing
  log (show argsRaw)
  log (show args)
  log (show (getArg "--main"))
