<?php
$exports['runRBTreeFFI'] = function($limit) {
    $n = (int)$limit;
    $empty = (object)["type" => "Leaf"];
    
    $insert = function($k, $v, $tree) use (&$insert) {
        $balance = function($color, $left, $key, $value, $right) {
            if ($color === 'B') {
                if ($left->type === "Node" && $left->color === 'R') {
                    if ($left->left->type === "Node" && $left->left->color === 'R') {
                        return (object)[
                            "type" => "Node",
                            "color" => 'R',
                            "left" => (object)["type" => "Node", "color" => 'B', "left" => $left->left->left, "key" => $left->left->key, "value" => $left->left->value, "right" => $left->left->right],
                            "key" => $left->key,
                            "value" => $left->value,
                            "right" => (object)["type" => "Node", "color" => 'B', "left" => $left->right, "key" => $key, "value" => $value, "right" => $right]
                        ];
                    }
                    if ($left->right->type === "Node" && $left->right->color === 'R') {
                        return (object)[
                            "type" => "Node",
                            "color" => 'R',
                            "left" => (object)["type" => "Node", "color" => 'B', "left" => $left->left, "key" => $left->key, "value" => $left->value, "right" => $left->right->left],
                            "key" => $left->right->key,
                            "value" => $left->right->value,
                            "right" => (object)["type" => "Node", "color" => 'B', "left" => $left->right->right, "key" => $key, "value" => $value, "right" => $right]
                        ];
                    }
                }
                if ($right->type === "Node" && $right->color === 'R') {
                    if ($right->left->type === "Node" && $right->left->color === 'R') {
                        return (object)[
                            "type" => "Node",
                            "color" => 'R',
                            "left" => (object)["type" => "Node", "color" => 'B', "left" => $left, "key" => $key, "value" => $value, "right" => $right->left->left],
                            "key" => $right->left->key,
                            "value" => $right->left->value,
                            "right" => (object)["type" => "Node", "color" => 'B', "left" => $right->left->right, "key" => $right->key, "value" => $right->value, "right" => $right->right]
                        ];
                    }
                    if ($right->right->type === "Node" && $right->right->color === 'R') {
                        return (object)[
                            "type" => "Node",
                            "color" => 'R',
                            "left" => (object)["type" => "Node", "color" => 'B', "left" => $left, "key" => $key, "value" => $value, "right" => $right->left],
                            "key" => $right->key,
                            "value" => $right->value,
                            "right" => (object)["type" => "Node", "color" => 'B', "left" => $right->right->left, "key" => $right->right->key, "value" => $right->right->value, "right" => $right->right->right]
                        ];
                    }
                }
            }
            return (object)["type" => "Node", "color" => $color, "left" => $left, "key" => $key, "value" => $value, "right" => $right];
        };
        
        $ins = function($node) use (&$ins, $k, $v, $balance) {
            if ($node->type === "Leaf") {
                return (object)["type" => "Node", "color" => 'R', "left" => (object)["type" => "Leaf"], "key" => $k, "value" => $v, "right" => (object)["type" => "Leaf"]];
            }
            if ($k < $node->key) {
                return $balance($node->color, $ins($node->left), $node->key, $node->value, $node->right);
            } else if ($k > $node->key) {
                return $balance($node->color, $node->left, $node->key, $node->value, $ins($node->right));
            } else {
                return (object)["type" => "Node", "color" => $node->color, "left" => $node->left, "key" => $k, "value" => $v, "right" => $node->right];
            }
        };
        
        $res = $ins($tree);
        $res->color = 'B';
        return $res;
    };

    $lookup = function($k, $node) use (&$lookup) {
        if ($node->type === "Leaf") {
            return (object)["type" => "Nothing"];
        }
        if ($k < $node->key) {
            return $lookup($k, $node->left);
        } else if ($k > $node->key) {
            return $lookup($k, $node->right);
        } else {
            return (object)["type" => "Just", "value0" => $node->value];
        }
    };
    
    $tree = $empty;
    for ($i = 0; $i < $n; $i++) {
        $tree = $insert($i, $i, $tree);
    }
    
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $res = $lookup($i, $tree);
        if ($res->type === "Just") {
            $sum += $res->value0;
        }
    }
    return $sum;
};
return $exports;
