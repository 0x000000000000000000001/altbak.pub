<?php
$exports['runAstTreeFFICheatcode'] = function($limit) {
    $buildTree = function($depth) use (&$buildTree) {
        if ($depth === 0) {
            return (object)["type" => "Literal", "value" => 1];
        }
        return (object)["type" => "Add", "left" => $buildTree($depth - 1), "right" => $buildTree($depth - 1)];
    };
    $evalTree = function($node) use (&$evalTree) {
        if ($node->type === "Literal") {
            return $node->value;
        }
        if ($node->type === "Add") {
            return $evalTree($node->left) + $evalTree($node->right);
        }
        return 0;
    };
    $tree = $buildTree((int)$limit);
    return $evalTree($tree);
};
return $exports;
