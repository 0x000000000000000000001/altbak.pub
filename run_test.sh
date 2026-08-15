#!/bin/bash
cd /Users/0x1/Documents/htdocs/gopurs/gopurs
npm run build --silent
cd /Users/0x1/Documents/htdocs/altbak.pub
mkdir -p src/AppTest
cat << 'EOF2' > src/AppTest.purs
module AppTest where
import Prelude
import Effect (Effect)
import Bench (opaque)
import Test.DerivingTraversable as Test.DerivingTraversable

main :: Effect Unit
main = do
  let keepAlive = {
        _dummy: unit
      }
  void $ opaque keepAlive
EOF2
spago build -q
node ../gopurs/gopurs/bin/gopurs.js --main AppTest
gofmt -w output/purescript/Test_DerivingTraversable.go
cd output
go build -o go_app ./main
