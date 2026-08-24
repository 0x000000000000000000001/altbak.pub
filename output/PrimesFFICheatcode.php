<?php
$exports['runPrimesFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    if ($n < 2) return 0;
    $sieve = array_fill(0, $n + 1, 1);
    
    for ($p = 2; $p * $p <= $n; $p++) {
        if ($sieve[$p]) {
            for ($i = $p * $p; $i <= $n; $i += $p) {
                $sieve[$i] = 0;
            }
        }
    }
    
    $sum = 0;
    for ($p = 2; $p <= $n; $p++) {
        if ($sieve[$p]) {
            $sum += $p;
        }
    }
    return $sum;
};
return $exports;
