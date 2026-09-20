<?php
$exports['runArrayOpsFFI'] = function($limit) {
    $n = (int)$limit;
    
    $range = function($start) {
        return function($end) use ($start) {
            $arr = [];
            $step = $end >= $start ? 1 : -1;
            for ($i = $start; $step > 0 ? $i <= $end : $i >= $end; $i += $step) {
                $arr[] = $i;
            }
            return $arr;
        };
    };

    $filter = function($p) {
        return function($arr) use ($p) {
            $res = [];
            foreach ($arr as $x) {
                if ($p($x)) {
                    $res[] = $x;
                }
            }
            return $res;
        };
    };

    $foldl = function($f, $initial, $arr) {
        $acc = $initial;
        foreach ($arr as $x) {
            $acc = $f($acc)($x);
        }
        return $acc;
    };

    $rng = $range(1)($n);
    $filtered = $filter(function($x) { return $x % 2 === 0; })($rng);
    return $foldl(function($acc) { return function($x) use ($acc) { return $acc + $x; }; }, 0, $filtered);
};
return $exports;
