<?php
$exports['runStateMonadFFI'] = function($input) {
    $pure = function($value) {
        return function($state) use ($value) { return (object)['value' => $value, 'state' => $state]; };
    };
    $bind = function($action, $next) {
        return function($state) use ($action, $next) {
            $first = $action($state);
            return $next($first->value)($first->state);
        };
    };
    $get = function($state) { return (object)['value' => $state, 'state' => $state]; };
    $put = function($state) {
        return function($_) use ($state) { return (object)['value' => null, 'state' => $state]; };
    };
    $modify = function($f) use ($bind, $get, $put) {
        return $bind($get, function($state) use ($f, $put) { return $put($f($state)); });
    };
    $chain = function($depth) use (&$chain, $pure, $bind, $modify) {
        if ($depth === 0) return $pure(null);
        return $bind($modify(function($value) { return $value + 1; }),
            function($_) use (&$chain, $depth) { return $chain($depth - 1); });
    };
    $total = 0;
    for ($i = 0; $i < 20; $i++) $total += $chain((int)$input)(0)->state;
    return $total;
};
return $exports;
