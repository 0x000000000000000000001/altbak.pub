<?php
$exports['runAstTreeFFI'] = function($limit) {
    $build = function($n) use (&$build) {
        if ($n === 0) return (object)['tag' => 'Val', 'value' => 1];
        return (object)['tag' => 'Add',
            'left' => (object)['tag' => 'Mul', 'left' => (object)['tag' => 'Val', 'value' => $n], 'right' => $build($n - 1)],
            'right' => (object)['tag' => 'Sub', 'left' => $build($n - 1), 'right' => (object)['tag' => 'Val', 'value' => 1]]];
    };
    $eval = function($tree) use (&$eval) {
        switch ($tree->tag) {
            case 'Val': return $tree->value;
            case 'Add': return $eval($tree->left) + $eval($tree->right);
            case 'Mul': return $eval($tree->left) * $eval($tree->right);
            case 'Sub': return $eval($tree->left) - $eval($tree->right);
        }
        throw new \LogicException('Unknown AST tag');
    };
    return $eval($build((int)$limit));
};
return $exports;
