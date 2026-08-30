<?php

$writeFileSync = function($path, $content = null) use (&$writeFileSync) {
    if (func_num_args() < 2) {
        $__args = func_get_args();
        return function(...$more) use ($__args, &$writeFileSync) {
            return $writeFileSync(...array_merge($__args, $more));
        };
    }
    return function() use ($path, $content) {
        file_put_contents($path, $content);
    };
};

$readFileSync = function($path) {
    return function() use ($path) {
        return file_get_contents($path);
    };
};

$loopE = function($n, $action = null) use (&$loopE) {
    if (func_num_args() < 2) {
        $__args = func_get_args();
        return function(...$more) use ($__args, &$loopE) {
            return $loopE(...array_merge($__args, $more));
        };
    }
    return function() use ($n, $action) {
        for ($i = 0; $i < $n; $i++) {
            $action();
        }
    };
};

$exports['writeFileSync'] = $writeFileSync;
$exports['readFileSync'] = $readFileSync;
$exports['loopE'] = $loopE;

return $exports;
