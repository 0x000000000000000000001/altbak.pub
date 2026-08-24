<?php
$exports['runTCOFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    $i = $n;
    while ($i > 0) {
        $acc += $i;
        $i--;
    }
    return $acc;
};
return $exports;
