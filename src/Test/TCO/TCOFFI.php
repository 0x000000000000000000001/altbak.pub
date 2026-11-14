<?php
$exports['runTCOFFI'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    while ($n > 0) {
        $acc += ($n % 3);
        $n -= 1;
    }
    return $acc;
};
return $exports;
