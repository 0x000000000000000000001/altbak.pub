-- | Do not delete. Useful to avoid warnings from `spago build`.
-- | Since the `test/` directory exists (to hold `test/snapshots/`), 
-- | `purs compile` warns "No files found using pattern: test/**/*.purs"
-- | if there isn't at least one `.purs` file in it.
module Test.Main where

import Prelude
import Effect (Effect)

main :: Effect Unit
main = pure unit
