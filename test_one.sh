#!/usr/bin/env bash
export PATH="$PWD/run/bak/js/node_modules/.bin:$PATH"
mod="Test.Church"
f="src/Test/Church.purs"
desc_line="    desc: $mod.describe,"
act_line="    act: $mod.act,"

cat <<EOF2 > /tmp/AppTest_$$.purs
module AppTest where
import Prelude
import Effect (Effect)
import Bench (opaque)
import $mod as $mod

main :: Effect Unit
main = void \$ opaque keepAlive
  where
  keepAlive = {
$desc_line
$act_line
    _dummy: unit
  }
EOF2

rm -f src/AppTest.purs
mv -f /tmp/AppTest_$$.purs src/AppTest.purs
rm -rf output/AppTest

spago build -q || echo "spago build failed"
node ../gopurs/gopurs/bin/gopurs.js --main AppTest || echo "gopurs failed"

cd output
go mod tidy
go build -o go_app ./main || echo "go build failed"
