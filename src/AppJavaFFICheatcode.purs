module AppJavaFFICheatcode where

import Prelude
import Effect (Effect)
import AppFFICheatcode as AppFFICheatcode
import AppJavaFFI (warmupOptimized)

main :: Effect Unit
main = do
  warmupOptimized
  warmupOptimized
  warmupOptimized
  AppFFICheatcode.main
