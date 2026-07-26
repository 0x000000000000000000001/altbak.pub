<?php
class Data_Maybe_Just { public $value0; public function __construct($v) { $this->value0 = $v; } }
class Data_Maybe_Nothing { public function __construct() {} }

$start = microtime(true);
$sum = 0;
for ($i = 0; $i < 1000000; $i++) {
    $x = new Data_Maybe_Just($i);
    if ($x instanceof Data_Maybe_Just) {
        $sum += $x->value0;
    }
}
$timeObj = (microtime(true) - $start) * 1000;
echo "Objects: $timeObj ms\n";

$start = microtime(true);
$sum = 0;
for ($i = 0; $i < 1000000; $i++) {
    $x = ['Data_Maybe_Just', $i];
    if ($x[0] === 'Data_Maybe_Just') {
        $sum += $x[1];
    }
}
$timeArr = (microtime(true) - $start) * 1000;
echo "Arrays: $timeArr ms\n";
