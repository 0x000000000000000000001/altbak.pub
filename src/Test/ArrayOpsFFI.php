<?php
$exports['runArrayOpsFFI'] = function($limit) {
    $n = (int)$limit;
    
    $range = function($start) {
        return function($end) use ($start) {
            $arr = [];
            for ($i = $start; $i <= $end; $i++) {
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

    $sumArray = function($arr) {
        $sum = 0;
        foreach ($arr as $x) {
            $sum += $x;
        }
        return $sum;
    };

    $rng = $range(1)($n);
    $filtered = $filter(function($x) { return $x % 2 === 0; })($rng);
    return $sumArray($filtered);
};
return $exports;
