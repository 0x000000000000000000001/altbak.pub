<?php
$exports['runTCOFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    while ($n > 0) {
        $acc += $n;
        $n -= 1;
    }
    return $acc;
};
return $exports;
