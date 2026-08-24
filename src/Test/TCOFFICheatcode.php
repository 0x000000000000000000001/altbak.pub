<?php
$exports['runTCOFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    while ($n > 0) {
        $acc += ($n % 3);
        $n--;
    }
    return $acc;
};
return $exports;
