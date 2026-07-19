<?php
require_once __DIR__ . '/../run/bak/php/output/App/main.mod.php';

$tests = [
    'Test_AstTree' => $GLOBALS['Test_AstTree_act'] ?? \PhpursThunks::eval('Test_AstTree_act'),
    'Test_Fib' => $GLOBALS['Test_Fib_act'] ?? \PhpursThunks::eval('Test_Fib_act'),
    'Test_ListOps' => $GLOBALS['Test_ListOps_act'] ?? \PhpursThunks::eval('Test_ListOps_act'),
    'Test_TCO' => $GLOBALS['Test_TCO_act'] ?? \PhpursThunks::eval('Test_TCO_act'),
    'Test_Records' => $GLOBALS['Test_Records_act'] ?? \PhpursThunks::eval('Test_Records_act'),
    'Test_Ackermann' => $GLOBALS['Test_Ackermann_act'] ?? \PhpursThunks::eval('Test_Ackermann_act'),
    'Test_Church' => $GLOBALS['Test_Church_act'] ?? \PhpursThunks::eval('Test_Church_act'),
    'Test_Primes' => $GLOBALS['Test_Primes_act'] ?? \PhpursThunks::eval('Test_Primes_act'),
    'Test_RBTree' => $GLOBALS['Test_RBTree_act'] ?? \PhpursThunks::eval('Test_RBTree_act'),
    'Test_Polymorphism' => $GLOBALS['Test_Polymorphism_act'] ?? \PhpursThunks::eval('Test_Polymorphism_act'),
    'Test_StateMonad' => $GLOBALS['Test_StateMonad_act'] ?? \PhpursThunks::eval('Test_StateMonad_act'),
    'Test_LazyEvaluation' => $GLOBALS['Test_LazyEvaluation_act'] ?? \PhpursThunks::eval('Test_LazyEvaluation_act'),
];

foreach ($tests as $name => $fn) {
    echo "Running $name...\n";
    $start = microtime(true);
    $fn();
    $end = microtime(true);
    echo "$name took: " . ($end - $start) . " seconds\n";
}
