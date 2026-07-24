#!/bin/bash
tests=(
  "Test.Ackermann"
  "Test.AstTree"
  "Test.Church"
  "Test.Fib"
  "Test.LazyEvaluation"
  "Test.ListOps"
  "Test.Polymorphism"
  "Test.Primes"
  "Test.RBTree"
  "Test.Records"
  "Test.StateMonad"
  "Test.TCO"
)

for t in "${tests[@]}"; do
  echo "Profiling $t..."
  if ! PPROF=1 ./bin/go/run --test "$t" > "out_$t.log" 2>&1; then
      echo "Failed $t! See out_$t.log"
      exit 1
  fi
  if [ -f "output/cpu.prof" ]; then
      mv output/cpu.prof "cpu_$t.prof"
      mv output/mem.prof "mem_$t.prof" 2>/dev/null || true
  else
      if [ -f "cpu.prof" ]; then
         mv cpu.prof "cpu_$t.prof"
         mv mem.prof "mem_$t.prof" 2>/dev/null || true
      fi
  fi
done
echo "Done!"
