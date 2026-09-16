<?php
$exports['runChurchFFI'] = function($limit) {
    $zeroC = function($f) { return function($x) { return $x; }; };
    $succC = function($n) {
        return function($f) use ($n) {
            return function($x) use ($n, $f) {
                return $f($n($f)($x));
            };
        };
    };
    $mulC = function($m) {
        return function($n) use ($m) {
            return function($f) use ($m, $n) {
                return function($x) use ($m, $n, $f) {
                    return $m($n($f))($x);
                };
            };
        };
    };
    $fromInt = function($n) use (&$fromInt, $zeroC, $succC) {
        if ($n === 0) return $zeroC;
        return $succC($fromInt($n - 1));
    };
    
    $c10 = function($n) use ($fromInt) { return $fromInt($n); };
    $c100 = function($n) use ($c10, $mulC) { return $mulC($c10($n))($c10($n)); };
    $c10k = function($n) use ($c100, $mulC) { return $mulC($c100($n))($c100($n)); };
    $c100k = function($n) use ($c10k, $c10, $mulC) { return $mulC($c10k($n))($c10($n)); };
    
    $toInt = function($n) {
        return $n(function($x) { return $x + 1; })(0);
    };
    
    return $toInt($c100k((int)$limit));
};
return $exports;
