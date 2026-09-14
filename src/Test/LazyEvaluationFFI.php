<?php
$exports['runLazyEvaluationFFI'] = function($count) {
    // Tail-recursive construction/repetition become loops; each force still
    // evaluates the full chain of 1000 closures, as in Test.LazyEvaluation.
    $result = 0;
    for ($repetition = 0; $repetition < (int)$count; $repetition++) {
        $thunk = function() { return 0; };
        for ($depth = 0; $depth < 1000; $depth++) {
            $previous = $thunk;
            $thunk = function() use ($previous) { return $previous() + 1; };
        }
        $result += $thunk();
    }
    return $result;
};
return $exports;
