<?php

$benchNow = function() {
    return hrtime(true) / 1000.0;
};

$opaque = function($a) {
    return function() use($a) {
        return $a;
    };
};

$formatNumber = function($n) {
    return number_format($n, 6, '.', '');
};

$measureBatch = function($iterations, $expected, $act) {
    // Phpurs groups foreign-function arguments; the returned closure is Effect.
    return function() use ($iterations, $expected, $act) {
        $result = 0;
        $start = hrtime(true);
        for ($i = 0; $i < $iterations; $i++) {
            $result = $act();
            $GLOBALS['altbak_benchmark_result'] = $result;
        }
        $elapsed = (hrtime(true) - $start) / 1000.0;
        if ($result !== $expected) throw new \RuntimeException('Unstable benchmark result');
        return $elapsed / $iterations;
    };
};

$exports['benchNow'] = $benchNow;
$exports['opaque'] = $opaque;
$exports['formatNumber'] = $formatNumber;
$exports['measureBatch'] = $measureBatch;

return $exports;
