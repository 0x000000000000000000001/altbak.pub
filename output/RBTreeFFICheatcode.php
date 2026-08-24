<?php
$exports['runRBTreeFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $root = null;
    
    $insert = function($root, $key, $value) use (&$insert) {
        if ($root === null) {
            return (object)["color" => 'B', "left" => null, "key" => $key, "value" => $value, "right" => null];
        }
        
        $balance = function($color, $left, $key, $value, $right) {
            if ($color === 'B') {
                if ($left !== null && $left->color === 'R') {
                    if ($left->left !== null && $left->left->color === 'R') {
                        return (object)[
                            "color" => 'R',
                            "left" => (object)["color" => 'B', "left" => $left->left->left, "key" => $left->left->key, "value" => $left->left->value, "right" => $left->left->right],
                            "key" => $left->key,
                            "value" => $left->value,
                            "right" => (object)["color" => 'B', "left" => $left->right, "key" => $key, "value" => $value, "right" => $right]
                        ];
                    }
                    if ($left->right !== null && $left->right->color === 'R') {
                        return (object)[
                            "color" => 'R',
                            "left" => (object)["color" => 'B', "left" => $left->left, "key" => $left->key, "value" => $left->value, "right" => $left->right->left],
                            "key" => $left->right->key,
                            "value" => $left->right->value,
                            "right" => (object)["color" => 'B', "left" => $left->right->right, "key" => $key, "value" => $value, "right" => $right]
                        ];
                    }
                }
                if ($right !== null && $right->color === 'R') {
                    if ($right->left !== null && $right->left->color === 'R') {
                        return (object)[
                            "color" => 'R',
                            "left" => (object)["color" => 'B', "left" => $left, "key" => $key, "value" => $value, "right" => $right->left->left],
                            "key" => $right->left->key,
                            "value" => $right->left->value,
                            "right" => (object)["color" => 'B', "left" => $right->left->right, "key" => $right->key, "value" => $right->value, "right" => $right->right]
                        ];
                    }
                    if ($right->right !== null && $right->right->color === 'R') {
                        return (object)[
                            "color" => 'R',
                            "left" => (object)["color" => 'B', "left" => $left, "key" => $key, "value" => $value, "right" => $right->left],
                            "key" => $right->key,
                            "value" => $right->value,
                            "right" => (object)["color" => 'B', "left" => $right->right->left, "key" => $right->right->key, "value" => $right->right->value, "right" => $right->right->right]
                        ];
                    }
                }
            }
            return (object)["color" => $color, "left" => $left, "key" => $key, "value" => $value, "right" => $right];
        };

        $ins = function($node) use (&$ins, $key, $value, $balance) {
            if ($node === null) {
                return (object)["color" => 'R', "left" => null, "key" => $key, "value" => $value, "right" => null];
            }
            if ($key < $node->key) {
                return $balance($node->color, $ins($node->left), $node->key, $node->value, $node->right);
            } else if ($key > $node->key) {
                return $balance($node->color, $node->left, $node->key, $node->value, $ins($node->right));
            } else {
                return (object)["color" => $node->color, "left" => $node->left, "key" => $key, "value" => $value, "right" => $node->right];
            }
        };
        
        $res = $ins($root);
        $res->color = 'B';
        return $res;
    };
    
    for ($i = 0; $i < $n; $i++) {
        $root = $insert($root, $i, $i);
    }
    
    $lookup = function($node, $key) {
        while ($node !== null) {
            if ($key < $node->key) {
                $node = $node->left;
            } else if ($key > $node->key) {
                $node = $node->right;
            } else {
                return $node;
            }
        }
        return null;
    };

    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $node = $lookup($root, $i);
        if ($node !== null) {
            $sum += $node->value;
        }
    }
    return $sum;
};
return $exports;
