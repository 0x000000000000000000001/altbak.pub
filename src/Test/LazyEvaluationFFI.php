<?php
$exports['runLazyEvaluationFFI'] = function($limit) {
    $n = (int)$limit;
    
    $defer = function($thunk) {
        $value = null;
        $evaluated = false;
        return function() use (&$value, &$evaluated, $thunk) {
            if (!$evaluated) {
                $value = $thunk();
                $evaluated = true;
            }
            return $value;
        };
    };
    
    $force = function($lzy) { return $lzy(); };
    
    $go = function($i, $acc) use (&$go, $defer, $force) {
        if ($i >= 1000) return $acc;
        $lazyVal = $defer(function() { return 1; });
        return $go($i + 1, $acc + $force($lazyVal));
    };
    
    $acc = 0;
    for ($i = 0; $i < $n; $i++) {
        $acc += $go(0, 0);
    }
    return $acc;
};
return $exports;
