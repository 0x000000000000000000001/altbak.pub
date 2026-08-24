<?php
$exports['runStateMonadFFI'] = function($limit) {
    $n = (int)$limit;
    
    $pure = function($a) {
        return function($s) use ($a) {
            return (object)["value0" => $a, "value1" => $s];
        };
    };
    
    $bind = function($m, $f) {
        return function($s) use ($m, $f) {
            $res = $m($s);
            return $f($res->value0)($res->value1);
        };
    };
    
    $get = function($s) { return (object)["value0" => $s, "value1" => $s]; };
    $put = function($s) { return function($_) use ($s) { return (object)["value0" => null, "value1" => $s]; }; };
    
    $modify = function($f) use ($bind, $get, $put) {
        return $bind($get, function($s) use ($f, $put) {
            return $put($f($s));
        });
    };
    
    $go = function($i, $m) use (&$go, $bind, $modify) {
        if ($i <= 0) return $m;
        return $go($i - 1, $bind($m, function($_) use ($modify) {
            return $modify(function($s) { return $s + 1; });
        }));
    };
    
    $res = $go(60, $pure(null))(0);
    return $res->value1;
};
return $exports;
