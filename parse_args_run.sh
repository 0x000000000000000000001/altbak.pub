#!/usr/bin/env bash
for script in bin/js/run bin/es/run bin/php/run bin/go/run bin/scm/run bin/erl/run; do
  # Determine specific output cleanup
  CLEANUP=""
  if [[ "$script" == *"es/run"* ]]; then
    CLEANUP="rm -rf run/bak/es/output-purs/* run/bak/js/spago/*"
  elif [[ "$script" == *"go/run"* ]]; then
    CLEANUP="rm -rf run/bak/go/output/* run/bak/go/spago/*"
  elif [[ "$script" == *"js/run"* ]]; then
    CLEANUP="rm -rf run/bak/js/output/* run/bak/js/spago/*"
  elif [[ "$script" == *"php/run"* ]]; then
    CLEANUP="rm -rf run/bak/php/output/* run/bak/php/spago/*"
  elif [[ "$script" == *"scm/run"* ]]; then
    CLEANUP="rm -rf run/bak/scm/output/* run/bak/scm/spago/*"
  elif [[ "$script" == *"erl/run"* ]]; then
    CLEANUP="rm -rf run/bak/erl/output/* run/bak/erl/spago/*"
  fi

  cat << INNER_EOF > "$script.new"
#!/usr/bin/env bash
set -e
cd "\$(dirname "\$0")/../.."

CLEAN=0
MODE="pure"
TEST_NAME=""

while [[ \$# -gt 0 ]]
do
  case "\$1" in
    --clean)
    CLEAN=1
    shift
    ;;
    --test)
    MODE="test"
    TEST_NAME="\$2"
    shift 2
    ;;
    --x)
    MODE="x"
    shift
    ;;
    *)
    shift
    ;;
  esac
done

if [ "\$CLEAN" == "1" ]; then
  $CLEANUP
fi
INNER_EOF

  # Append everything after the original if [ "$1" == "--clean" ] ...
  awk '
    BEGIN { skip = 1 }
    /mkdir -p/ { skip = 0 }
    !skip {
      if ($0 ~ /if \[ "\$1" == "--test" \]; then/) {
        print "if [ \"$MODE\" == \"test\" ]; then"
      } else if ($0 ~ /elif \[ "\$1" == "--x" \]; then/) {
        print "elif [ \"$MODE\" == \"x\" ]; then"
      } else if ($0 ~ /TEST_NAME="\$2"/) {
        # remove it as we already parsed it
      } else {
        print $0
      }
    }
  ' "$script" >> "$script.new"

  chmod +x "$script.new"
  mv "$script.new" "$script"
done
