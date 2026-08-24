<?php
$exports['runLazyEvaluationFFI'] = function($limit) {
    $n = (int)$limit;
    
    $defer = function($f) { return $f; };
    $force = function($l) { return $l(); };
    
    $buildThunks = function($depth, $acc) use (&$buildThunks, $defer, $force) {
        if ($depth === 0) return $acc;
        return $buildThunks($depth - 1, $defer(function() use ($force, $acc) {
            return $force($acc) + 1;
        }));
    };
    
    $runManyTimes = function($times, $acc) use (&$runManyTimes, $buildThunks, $defer, $force) {
        if ($times === 0) return $acc;
        $t = $buildThunks(1000, $defer(function() { return 0; }));
        return $runManyTimes($times - 1, $acc + $force($t));
    };
    
    return $runManyTimes($n, 0);
};
return $exports;
