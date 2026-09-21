module Test.JsonTypedAst where

import Prelude
import Data.Argonaut.Core (Json, stringify)
import Data.Argonaut.Parser (jsonParser)
import Data.Argonaut.Decode.Error (printJsonDecodeError)
import Data.Either (Either(..))
import Effect (Effect)
import Partial.Unsafe (unsafeCrashWith)
import PureScript.Backend.Optimizer.CoreFn (Ann, Module)
import PureScript.Backend.Optimizer.CoreFn.Json (decodeModule)
import Test.JsonTypedAst.Fingerprint (fingerprint)

foreign import drive :: forall a b. (String -> a) -> (a -> b) -> (a -> String) -> (b -> String) -> Effect Unit

parse :: String -> Json
parse input = case jsonParser input of
  Left err -> unsafeCrashWith err
  Right value -> value

decode :: Json -> Module Ann
decode input = case decodeModule input of
  Left err -> unsafeCrashWith (printJsonDecodeError err)
  Right value -> value

main :: Effect Unit
main = drive parse decode stringify fingerprint
