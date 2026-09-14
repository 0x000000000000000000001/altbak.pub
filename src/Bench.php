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
    return number_format($n, 2, '.', '');
};

$exports['benchNow'] = $benchNow;
$exports['opaque'] = $opaque;
$exports['formatNumber'] = $formatNumber;

return $exports;
