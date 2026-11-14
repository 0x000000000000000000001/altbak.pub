module Test.JsonDecoding where

import Prelude

import Data.Argonaut.Core (Json, stringify)
import Data.Argonaut.Decode (class DecodeJson, decodeJson, (.:), (.:?))
import Data.Argonaut.Decode.Error (JsonDecodeError(..), printJsonDecodeError)
import Data.Argonaut.Decode.Parser (decodeJsonStringWith)
import Data.Argonaut.Encode (class EncodeJson, encodeJson)
import Data.Argonaut.Parser (jsonParser)
import Data.Either (Either(..))
import Data.Maybe (Maybe)
import Effect (Effect)
import Partial.Unsafe (unsafeCrashWith)

type Profile =
  { city :: String
  , note :: Maybe String
  , scores :: Array Number
  }

type User =
  { id :: Int
  , name :: String
  , active :: Boolean
  , profile :: Maybe Profile
  , tags :: Array String
  }

type Item = { sku :: String, quantity :: Int, price :: Number }

data Event
  = View String (Maybe Int)
  | Purchase Int (Array Item)

instance decodeEvent :: DecodeJson Event where
  decodeJson json = do
    obj <- decodeJson json
    tag <- obj .: "tag"
    case tag of
      "view" -> View <$> obj .: "path" <*> obj .:? "duration"
      "purchase" -> Purchase <$> obj .: "orderId" <*> obj .: "items"
      _ -> Left (TypeMismatch "Event tag")

instance encodeEvent :: EncodeJson Event where
  encodeJson = case _ of
    View path duration -> encodeJson { tag: "view", path, duration }
    Purchase orderId items -> encodeJson { tag: "purchase", orderId, items }

type Payload =
  { version :: Int
  , next :: Maybe String
  , users :: Array User
  , events :: Array Event
  }

-- The FFI only supplies inputs, timing, allocation counts and output checks.
-- Parsing and decoding use the same ordinary Argonaut code in both backends.
foreign import drive :: forall a b. (String -> a) -> (a -> b) -> (String -> b) -> (a -> String) -> (b -> String) -> Effect Unit

parse :: String -> Json
parse input = case jsonParser input of
  Left err -> unsafeCrashWith err
  Right value -> value

decode :: Json -> Either JsonDecodeError Payload
decode = decodeJson

decodeText :: String -> Either JsonDecodeError Payload
decodeText = decodeJsonStringWith decode

fingerprint :: Either JsonDecodeError Payload -> String
fingerprint = case _ of
  Left err -> stringify $ encodeJson { error: printJsonDecodeError err }
  Right value -> stringify $ encodeJson { value }

main :: Effect Unit
main = drive parse decode decodeText stringify fingerprint
