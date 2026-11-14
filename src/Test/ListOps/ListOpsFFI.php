<?php
$exports['runListOpsFFI'] = function($limit) {
    $n = (int)$limit;
    
    $range = function($start) {
        return function($end) use ($start) {
            $go = function($curr, $acc) use (&$go, $start) {
                if ($curr < $start) return $acc;
                return $go($curr - 1, (object)["type" => "Cons", "value0" => $curr, "value1" => $acc]);
            };
            return $go($end, (object)["type" => "Nil"]);
        };
    };

    $filter = function($p) {
        return function($lst) use ($p) {
            $go = function($list, $acc) use (&$go, $p) {
                if ($list->type === "Nil") {
                    return $acc;
                }
                $x = $list->value0;
                $xs = $list->value1;
                if ($p($x)) {
                    return $go($xs, (object)["type" => "Cons", "value0" => $x, "value1" => $acc]);
                } else {
                    return $go($xs, $acc);
                }
            };
            return $go($lst, (object)["type" => "Nil"]);
        };
    };

    $foldl = function($f, $initial, $lst) {
        $go = function($list, $acc) use (&$go, $f) {
            if ($list->type === "Nil") return $acc;
            return $go($list->value1, $f($acc)($list->value0));
        };
        return $go($lst, $initial);
    };

    $rng = $range(1)($n);
    $filtered = $filter(function($x) { return $x % 2 === 0; })($rng);
    return $foldl(function($acc) { return function($x) use ($acc) { return $acc + $x; }; }, 0, $filtered);
};
return $exports;
