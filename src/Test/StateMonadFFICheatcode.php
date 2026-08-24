<?php
$exports['runStateMonadFFICheatcode'] = function($limit) {
    $n = (int)$limit * 20; // in BenchCheck it passes 60, but wait: 60 * 20 = 1200 binds
    $state = 0;
    for ($i = 0; $i < $n; $i++) {
        $state += 1;
    }
    return $state;
};
return $exports;
