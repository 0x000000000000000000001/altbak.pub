<?php
$exports['runPrimesFFI'] = function($limit) {
    $n = (int)$limit;
    
    $range = function($start) {
        return function($end) use ($start) {
            $go = function($curr, $acc) use (&$go, $start) {
                if ($curr < $start) return $acc;
                return (object)["type" => "Cons", "value0" => $curr, "value1" => $acc];
            };
            return $go($end, (object)["type" => "Nil"]);
        };
    };

    $filter = function($p) {
        return function($lst) use ($p) {
            $go = function($list, $acc) use (&$go, $p) {
                if ($list->type === "Nil") {
                    $rev = function($l, $a) use (&$rev) {
                        if ($l->type === "Nil") return $a;
                        return $rev($l->value1, (object)["type" => "Cons", "value0" => $l->value0, "value1" => $a]);
                    };
                    return $rev($acc, (object)["type" => "Nil"]);
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

    $sieve = function($lst) use (&$sieve, $filter) {
        if ($lst->type === "Nil") return (object)["type" => "Nil"];
        $p = $lst->value0;
        $xs = $lst->value1;
        $filtered = $filter(function($x) use ($p) { return $x % $p !== 0; })($xs);
        return (object)[
            "type" => "Cons",
            "value0" => $p,
            "value1" => $sieve($filtered)
        ];
    };

    $sumList = function($lst) {
        $go = function($list, $acc) use (&$go) {
            if ($list->type === "Nil") return $acc;
            return $go($list->value1, $acc + $list->value0);
        };
        return $go($lst, 0);
    };

    $rng = $range(2)($n);
    $s = $sieve($rng);
    return $sumList($s);
};
return $exports;
