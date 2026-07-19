module Test.StringOps where

import Effect as Effect
import Effect.Console as Effect.Console

import Prelude
import Data.String as S
import Data.String.Regex as R
import Data.String.Regex.Flags (noFlags)
import Data.Either (Either(..))
import Partial.Unsafe (unsafePartial)
import Data.Array (length)

regexPattern :: R.Regex
regexPattern = unsafePartial (case R.regex "(hello|world)[0-9]+" noFlags of Right r -> r)

runStringOps :: Int -> Int
runStringOps n = 
  let
    loop :: Int -> String -> Int -> Int
    loop 0 _ acc = acc
    loop i s acc =
      let 
        concatted = s <> show i <> "world" <> show (i+1)
        replaced = R.replace regexPattern "matched" concatted
        splitted = S.split (S.Pattern "e") replaced
      in loop (i - 1) (S.take 10 concatted) (acc + length splitted)
  in loop n "hello" 0


describe :: Effect.Effect Unit
describe = Effect.Console.log "String Operations (1k Regex/Split):"

act :: Effect.Effect Unit
act = Effect.Console.logShow $ runStringOps 1000
