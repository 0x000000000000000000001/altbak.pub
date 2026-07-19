<?php

$iterations = 10000000;

// ============================================================================
// Scenario 1: Current phpurs_curry_fallback (Array merge + splat)
// ============================================================================
function current_curry_fallback($fn, $args, $expected) {
    // Simplified for expected=2, given=1
    return function($a) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num > 1) {
            $merged = \array_merge($args, \func_get_args());
            $res = $fn(...\array_slice($merged, 0, $expected));
            return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a;
        return $fn(...$args);
    };
}

$intSub_current = function($x = null, $y = null) use (&$intSub_current) {
    if (\func_num_args() < 2) return current_curry_fallback($intSub_current, \func_get_args(), 2);
    return $x - $y;
};

$mappend_current = function($x = null, $y = null) use (&$mappend_current) {
    if (\func_num_args() < 2) return current_curry_fallback($mappend_current, \func_get_args(), 2);
    return $x + $y;
};

$start = microtime(true);
$v_5 = $iterations;
$v1_6 = 0;
while (true) {
    if ($v_5 === 0) break;
    $tco_2 = ($intSub_current($v_5))(1);
    $tco_3 = ($mappend_current($v1_6))(1);
    $v_5 = $tco_2;
    $v1_6 = $tco_3;
}
$end = microtime(true);
echo "1. Current phpurs_curry_fallback: " . ($end - $start) . " seconds\n";


// ============================================================================
// Scenario 2: Optimized phpurs_curry_fallback (Direct call, no splat/merge)
// ============================================================================
function optimized_curry_fallback($fn, $args, $expected) {
    // Specifically optimized for 2-arity functions where 1 arg is given
    return function($a) use ($fn, $args) {
        // Omitting func_num_args() check for this benchmark focus
        return $fn($args[0], $a);
    };
}

$intSub_opt = function($x = null, $y = null) use (&$intSub_opt) {
    if (\func_num_args() < 2) return optimized_curry_fallback($intSub_opt, \func_get_args(), 2);
    return $x - $y;
};

$mappend_opt = function($x = null, $y = null) use (&$mappend_opt) {
    if (\func_num_args() < 2) return optimized_curry_fallback($mappend_opt, \func_get_args(), 2);
    return $x + $y;
};

$start = microtime(true);
$v_5 = $iterations;
$v1_6 = 0;
while (true) {
    if ($v_5 === 0) break;
    $tco_2 = ($intSub_opt($v_5))(1);
    $tco_3 = ($mappend_opt($v1_6))(1);
    $v_5 = $tco_2;
    $v1_6 = $tco_3;
}
$end = microtime(true);
echo "2. Optimized phpurs_curry_fallback: " . ($end - $start) . " seconds\n";


// ============================================================================
// Scenario 3: Uncurried (Backend-Optimizer emits PhpCall(fn, [arg1, arg2]))
// ============================================================================
$intSub_uncurried = function($x, $y) {
    return $x - $y;
};

$mappend_uncurried = function($x, $y) {
    return $x + $y;
};

$start = microtime(true);
$v_5 = $iterations;
$v1_6 = 0;
while (true) {
    if ($v_5 === 0) break;
    $tco_2 = $intSub_uncurried($v_5, 1);
    $tco_3 = $mappend_uncurried($v1_6, 1);
    $v_5 = $tco_2;
    $v1_6 = $tco_3;
}
$end = microtime(true);
echo "3. UncurriedApp (direct 2-arity call): " . ($end - $start) . " seconds\n";


// ============================================================================
// Scenario 4: Old Nested Closures (Pre backend-optimizer phpurs generation)
// ============================================================================
$intSub_nested = function($x) {
    return function($y) use ($x) {
        return $x - $y;
    };
};

$mappend_nested = function($x) {
    return function($y) use ($x) {
        return $x + $y;
    };
};

$start = microtime(true);
$v_5 = $iterations;
$v1_6 = 0;
while (true) {
    if ($v_5 === 0) break;
    $tco_2 = ($intSub_nested($v_5))(1);
    $tco_3 = ($mappend_nested($v1_6))(1);
    $v_5 = $tco_2;
    $v1_6 = $tco_3;
}
$end = microtime(true);
echo "4. Old Nested Closures (no fallback): " . ($end - $start) . " seconds\n";
