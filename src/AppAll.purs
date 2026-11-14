module AppAll where

import Prelude

import App as Core
import AppLibrary as WIP
import Effect (Effect)
import Effect.Console (log)

main :: Effect Unit
main = do
  Core.main
  log "=== Excluded WIP benchmarks ==="
  WIP.main
