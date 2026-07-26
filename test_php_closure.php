<?php
function foo(): array|\Closure {
  return function() {};
}
foo();
echo "OK\n";
