import os

files = {
    'Test.FibFFI.php': """<?php
$exports['runFibFFI'] = function($limit) {
    $fib = function($n) use (&$fib) {
        if ($n === 0) return 0;
        if ($n === 1) return 1;
        return $fib($n - 1) + $fib($n - 2);
    };
    return $fib((int)$limit);
};
return $exports;
""",
    'Test.FibFFICheatcode.php': """<?php
$exports['runFibFFICheatcode'] = function($n) use (&$exports) {
    if ($n === 0) return 0;
    if ($n === 1) return 1;
    $f = $exports['runFibFFICheatcode'];
    return $f($n - 1) + $f($n - 2);
};
return $exports;
""",
    'Test.AstTreeFFI.php': """<?php
$exports['runAstTreeFFI'] = function($limit) {
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
""",
    'Test.AstTreeFFICheatcode.php': """<?php
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
""",
    'Test.AckermannFFI.php': """<?php
$exports['runAckermannFFI'] = function($ignore) {
    $ack = function($m, $n) use (&$ack) {
        if ($m === 0) return $n + 1;
        if ($m > 0 && $n === 0) return $ack($m - 1, 1);
        return $ack($m - 1, $ack($m, $n - 1));
    };
    return $ack(3, 4);
};
return $exports;
""",
    'Test.AckermannFFICheatcode.php': """<?php
$exports['runAckermannFFICheatcode'] = function($ignore) {
    $ack = function($m, $n) use (&$ack) {
        if ($m === 0) return $n + 1;
        if ($m > 0 && $n === 0) return $ack($m - 1, 1);
        return $ack($m - 1, $ack($m, $n - 1));
    };
    return $ack(3, 4);
};
return $exports;
""",
    'Test.PrimesFFI.php': """<?php
$exports['runPrimesFFI'] = function($limit) {
    $n = (int)$limit;
    
    $range = function($start) {
        return function($end) use ($start) {
            $go = function($curr, $acc) use (&$go, $start) {
                if ($curr < $start) return $acc;
                return (object)["type" => "Cons", "value0" => $curr, "value1" => $acc];
            };
            return $go($end, (object)["type" => "Nil"]);
        };
    };

    $filter = function($p) {
        return function($lst) use ($p) {
            $go = function($list, $acc) use (&$go, $p) {
                if ($list->type === "Nil") {
                    $rev = function($l, $a) use (&$rev) {
                        if ($l->type === "Nil") return $a;
                        return $rev($l->value1, (object)["type" => "Cons", "value0" => $l->value0, "value1" => $a]);
                    };
                    return $rev($acc, (object)["type" => "Nil"]);
                }
                $x = $list->value0;
                $xs = $list->value1;
                if ($p($x)) {
                    return $go($xs, (object)["type" => "Cons", "value0" => $x, "value1" => $acc]);
                } else {
                    return $go($xs, $acc);
                }
            };
            return $go($lst, (object)["type" => "Nil"]);
        };
    };

    $sieve = function($lst) use (&$sieve, $filter) {
        if ($lst->type === "Nil") return (object)["type" => "Nil"];
        $p = $lst->value0;
        $xs = $lst->value1;
        $filtered = $filter(function($x) use ($p) { return $x % $p !== 0; })($xs);
        return (object)[
            "type" => "Cons",
            "value0" => $p,
            "value1" => $sieve($filtered)
        ];
    };

    $sumList = function($lst) {
        $go = function($list, $acc) use (&$go) {
            if ($list->type === "Nil") return $acc;
            return $go($list->value1, $acc + $list->value0);
        };
        return $go($lst, 0);
    };

    $rng = $range(2)($n);
    $s = $sieve($rng);
    return $sumList($s);
};
return $exports;
""",
    'Test.PrimesFFICheatcode.php': """<?php
$exports['runPrimesFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    if ($n < 2) return 0;
    $sieve = array_fill(0, $n + 1, 1);
    
    for ($p = 2; $p * $p <= $n; $p++) {
        if ($sieve[$p]) {
            for ($i = $p * $p; $i <= $n; $i += $p) {
                $sieve[$i] = 0;
            }
        }
    }
    
    $sum = 0;
    for ($p = 2; $p <= $n; $p++) {
        if ($sieve[$p]) {
            $sum += $p;
        }
    }
    return $sum;
};
return $exports;
""",
    'Test.ListOpsFFI.php': """<?php
$exports['runListOpsFFI'] = function($limit) {
    $n = (int)$limit;
    
    $range = function($start) {
        return function($end) use ($start) {
            $go = function($curr, $acc) use (&$go, $start) {
                if ($curr < $start) return $acc;
                return (object)["type" => "Cons", "value0" => $curr, "value1" => $acc];
            };
            return $go($end, (object)["type" => "Nil"]);
        };
    };

    $filter = function($p) {
        return function($lst) use ($p) {
            $go = function($list, $acc) use (&$go, $p) {
                if ($list->type === "Nil") {
                    $rev = function($l, $a) use (&$rev) {
                        if ($l->type === "Nil") return $a;
                        return $rev($l->value1, (object)["type" => "Cons", "value0" => $l->value0, "value1" => $a]);
                    };
                    return $rev($acc, (object)["type" => "Nil"]);
                }
                $x = $list->value0;
                $xs = $list->value1;
                if ($p($x)) {
                    return $go($xs, (object)["type" => "Cons", "value0" => $x, "value1" => $acc]);
                } else {
                    return $go($xs, $acc);
                }
            };
            return $go($lst, (object)["type" => "Nil"]);
        };
    };

    $sumList = function($lst) {
        $go = function($list, $acc) use (&$go) {
            if ($list->type === "Nil") return $acc;
            return $go($list->value1, $acc + $list->value0);
        };
        return $go($lst, 0);
    };

    $rng = $range(1)($n);
    $filtered = $filter(function($x) { return $x % 2 === 0; })($rng);
    return $sumList($filtered);
};
return $exports;
""",
    'Test.ListOpsFFICheatcode.php': """<?php
$exports['runListOpsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 1; $i <= $n; $i++) {
        if ($i % 2 === 0) {
            $sum += $i;
        }
    }
    return $sum;
};
return $exports;
""",
    'Test.TCOFFI.php': """<?php
$exports['runTCOFFI'] = function($limit) {
    $go = function($n, $acc) use (&$go) {
        if ($n <= 0) return $acc;
        return $go($n - 1, $acc + $n);
    };
    return $go((int)$limit, 0);
};
return $exports;
""",
    'Test.TCOFFICheatcode.php': """<?php
$exports['runTCOFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    while ($n > 0) {
        $acc += $n;
        $n -= 1;
    }
    return $acc;
};
return $exports;
""",
    'Test.RecordsFFI.php': """<?php
$exports['runRecordsFFI'] = function($limit) {
    $n = (int)$limit;
    $rec = (object)["a" => (object)["b" => (object)["c" => (object)["d" => (object)["e" => 0]]]]];
    $go = function($i, $r) use (&$go, $n) {
        if ($i >= $n) return $r;
        $newR = (object)[
            "a" => (object)[
                "b" => (object)[
                    "c" => (object)[
                        "d" => (object)[
                            "e" => $r->a->b->c->d->e + 1
                        ]
                    ]
                ]
            ]
        ];
        return $go($i + 1, $newR);
    };
    $finalRec = $go(0, $rec);
    return $finalRec->a->b->c->d->e;
};
return $exports;
""",
    'Test.RecordsFFICheatcode.php': """<?php
$exports['runRecordsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $rec = (object)["a" => (object)["b" => (object)["c" => (object)["d" => (object)["e" => 0]]]]];
    for ($i = 0; $i < $n; $i++) {
        $rec = (object)[
            "a" => (object)[
                "b" => (object)[
                    "c" => (object)[
                        "d" => (object)[
                            "e" => $rec->a->b->c->d->e + 1
                        ]
                    ]
                ]
            ]
        ];
    }
    return $rec->a->b->c->d->e;
};
return $exports;
""",
    'Test.ChurchFFI.php': """<?php
$exports['runChurchFFI'] = function($limit) {
    $n = (int)$limit * 10000;
    $zero = function($f) { return function($x) { return $x; }; };
    $succ = function($n_church) {
        return function($f) use ($n_church) {
            return function($x) use ($n_church, $f) {
                return $f($n_church($f)($x));
            };
        };
    };
    
    $go = function($i, $acc) use (&$go, $succ, $n) {
        if ($i >= $n) return $acc;
        return $go($i + 1, $succ($acc));
    };
    
    $churchN = $go(0, $zero);
    return $churchN(function($x) { return $x + 1; })(0);
};
return $exports;
""",
    'Test.ChurchFFICheatcode.php': """<?php
$exports['runChurchFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    for ($i = 0; $i < $n * 10000; $i++) {
        $acc++;
    }
    return $acc;
};
return $exports;
""",
    'Test.RBTreeFFI.php': """<?php
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
""",
    'Test.RBTreeFFICheatcode.php': """<?php
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
""",
    'Test.PolymorphismFFI.php': """<?php
$exports['runPolymorphismFFI'] = function($limit) {
    $n = (int)$limit;
    
    $lengthString = function($s) { return strlen($s); };
    $lengthArray = function($arr) { return count($arr); };
    
    $computeLength = function($dict) {
        return function($x) use ($dict) {
            return $dict->length($x);
        };
    };
    
    $dictString = (object)["length" => $lengthString];
    $dictArray = (object)["length" => $lengthArray];
    
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $sum += $computeLength($dictString)("hello") + $computeLength($dictArray)([1, 2, 3]);
    }
    return $sum;
};
return $exports;
""",
    'Test.PolymorphismFFICheatcode.php': """<?php
$exports['runPolymorphismFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    
    $computeLength = function($x) {
        if (is_string($x)) return strlen($x);
        if (is_array($x)) return count($x);
        return 0;
    };
    
    for ($i = 0; $i < $n; $i++) {
        $sum += $computeLength("hello") + $computeLength([1, 2, 3]);
    }
    return $sum;
};
return $exports;
""",
    'Test.StateMonadFFI.php': """<?php
$exports['runStateMonadFFI'] = function($limit) {
    $n = (int)$limit;
    
    $pure = function($a) {
        return function($s) use ($a) {
            return (object)["value0" => $a, "value1" => $s];
        };
    };
    
    $bind = function($m, $f) {
        return function($s) use ($m, $f) {
            $res = $m($s);
            return $f($res->value0)($res->value1);
        };
    };
    
    $get = function($s) { return (object)["value0" => $s, "value1" => $s]; };
    $put = function($s) { return function($_) use ($s) { return (object)["value0" => null, "value1" => $s]; }; };
    
    $modify = function($f) use ($bind, $get, $put) {
        return $bind($get, function($s) use ($f, $put) {
            return $put($f($s));
        });
    };
    
    $go = function($i, $m) use (&$go, $bind, $modify) {
        if ($i <= 0) return $m;
        return $go($i - 1, $bind($m, function($_) use ($modify) {
            return $modify(function($s) { return $s + 1; });
        }));
    };
    
    $res = $go(60, $pure(null))(0);
    return $res->value1;
};
return $exports;
""",
    'Test.StateMonadFFICheatcode.php': """<?php
$exports['runStateMonadFFICheatcode'] = function($limit) {
    $state = 0;
    for ($i = 0; $i < 60; $i++) {
        for ($j = 0; $j < 20; $j++) {
            $state += 1;
        }
    }
    return $state;
};
return $exports;
""",
    'Test.LazyEvaluationFFI.php': """<?php
$exports['runLazyEvaluationFFI'] = function($limit) {
    $n = (int)$limit;
    
    $defer = function($thunk) {
        $value = null;
        $evaluated = false;
        return function() use (&$value, &$evaluated, $thunk) {
            if (!$evaluated) {
                $value = $thunk();
                $evaluated = true;
            }
            return $value;
        };
    };
    
    $force = function($lzy) { return $lzy(); };
    
    $go = function($i, $acc) use (&$go, $defer, $force) {
        if ($i >= 1000) return $acc;
        $lazyVal = $defer(function() { return 1; });
        return $go($i + 1, $acc + $force($lazyVal));
    };
    
    $acc = 0;
    for ($i = 0; $i < $n; $i++) {
        $acc += $go(0, 0);
    }
    return $acc;
};
return $exports;
""",
    'Test.LazyEvaluationFFICheatcode.php': """<?php
$exports['runLazyEvaluationFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    for ($i = 0; $i < $n; $i++) {
        $acc += 1000;
    }
    return $acc;
};
return $exports;
""",
    'Test.ArrayOpsFFI.php': """<?php
$exports['runArrayOpsFFI'] = function($limit) {
    $n = (int)$limit;
    
    $range = function($start) {
        return function($end) use ($start) {
            $arr = [];
            for ($i = $start; $i <= $end; $i++) {
                $arr[] = $i;
            }
            return $arr;
        };
    };

    $filter = function($p) {
        return function($arr) use ($p) {
            $res = [];
            foreach ($arr as $x) {
                if ($p($x)) {
                    $res[] = $x;
                }
            }
            return $res;
        };
    };

    $sumArray = function($arr) {
        $sum = 0;
        foreach ($arr as $x) {
            $sum += $x;
        }
        return $sum;
    };

    $rng = $range(1)($n);
    $filtered = $filter(function($x) { return $x % 2 === 0; })($rng);
    return $sumArray($filtered);
};
return $exports;
""",
    'Test.ArrayOpsFFICheatcode.php': """<?php
$exports['runArrayOpsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 1; $i <= $n; $i++) {
        if ($i % 2 === 0) {
            $sum += $i;
        }
    }
    return $sum;
};
return $exports;
""",
    'Test.RowToListFFI.php': """<?php
$exports['runRowToListFFI'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $rec = (object)["a" => 1, "b" => "hello", "c" => true];
        $countKeys = function($r) {
            return count(get_object_vars($r));
        };
        $sum += $countKeys($rec) + ($rec->a * 2);
    }
    return $sum;
};
return $exports;
""",
    'Test.RowToListFFICheatcode.php': """<?php
$exports['runRowToListFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $rec = (object)["a" => 1, "b" => "hello", "c" => true];
        $sum += $rec->a;
    }
    return $sum;
};
return $exports;
"""
}

for filename, content in files.items():
    with open(f"src/Test/{filename}", "w") as f:
        f.write(content)
