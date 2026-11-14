<?php
$exports['runRBTreeFFI'] = function($limit) {
    $empty = (object)['tag' => 'E'];
    $node = function($color, $left, $value, $right) {
        return (object)['tag' => 'T', 'color' => $color, 'left' => $left, 'value' => $value, 'right' => $right];
    };
    $red = function($tree) { return $tree->tag === 'T' && $tree->color === 'R'; };
    $balance = function($color, $a, $x, $b) use ($node, $red) {
        if ($color === 'B') {
            if ($red($a) && $red($a->left)) {
                $child = $a->left;
                return $node('R', $node('B', $child->left, $child->value, $child->right), $a->value, $node('B', $a->right, $x, $b));
            }
            if ($red($a) && $red($a->right)) {
                $child = $a->right;
                return $node('R', $node('B', $a->left, $a->value, $child->left), $child->value, $node('B', $child->right, $x, $b));
            }
            if ($red($b) && $red($b->left)) {
                $child = $b->left;
                return $node('R', $node('B', $a, $x, $child->left), $child->value, $node('B', $child->right, $b->value, $b->right));
            }
            if ($red($b) && $red($b->right)) {
                $child = $b->right;
                return $node('R', $node('B', $a, $x, $b->left), $b->value, $node('B', $child->left, $child->value, $child->right));
            }
        }
        return $node($color, $a, $x, $b);
    };
    $ins = function($value, $tree) use (&$ins, $balance, $node, $empty) {
        if ($tree->tag === 'E') return $node('R', $empty, $value, $empty);
        if ($value < $tree->value) return $balance($tree->color, $ins($value, $tree->left), $tree->value, $tree->right);
        if ($value > $tree->value) return $balance($tree->color, $tree->left, $tree->value, $ins($value, $tree->right));
        return $node($tree->color, $tree->left, $tree->value, $tree->right);
    };
    $depth = function($tree) use (&$depth) {
        return $tree->tag === 'E' ? 0 : 1 + max($depth($tree->left), $depth($tree->right));
    };
    $tree = $empty;
    for ($value = (int)$limit; $value > 0; $value--) {
        $redRoot = $ins($value, $tree);
        $tree = $node('B', $redRoot->left, $redRoot->value, $redRoot->right);
    }
    return $depth($tree);
};
return $exports;
