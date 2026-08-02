<?php
$benchNow = function() {
    return microtime(true) * 1000000.0;
};
$opaque = function($a) {
    return function() use ($a) {
        return $a;
    };
};
$formatNumber = function($n) {
    return number_format($n, 2, '.', '');
};
$exports['benchNow'] = $benchNow;
$exports['opaque'] = $opaque;
$exports['formatNumber'] = $formatNumber;
$exports["keepAlive"] = function() {
    return function() {};
};
return $exports;
