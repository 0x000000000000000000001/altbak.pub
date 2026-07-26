<?php
function Call_foo($a) { return $a; }
$GLOBALS['foo'] = function($a) { return Call_foo($a); };

$start = microtime(true);
for ($i=0; $i<10000000; $i++) {
    $x = Call_foo($i);
}
$t1 = microtime(true) - $start;

$start = microtime(true);
for ($i=0; $i<10000000; $i++) {
    $x = (function_exists('Call_foo') ? Call_foo($i) : ($GLOBALS['foo'])($i));
}
$t2 = microtime(true) - $start;

echo "Direct: $t1\n";
echo "Function_exists: $t2\n";
