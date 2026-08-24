<?php
$exports['runTCOFFI'] = function($limit) {
    $go = function($n, $acc) use (&$go) {
        if ($n <= 0) return $acc;
        return $go($n - 1, $acc + $n);
    };
    return $go((int)$limit, 0);
};
return $exports;
