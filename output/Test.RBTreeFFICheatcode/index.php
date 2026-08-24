<?php

namespace Test\RBTreeFFICheatcode;

// ALL IMPORTS: Bench, Control.Bind, Data.Function, Data.Show, Effect, Effect.Console, Prelude, Prim, Test.RBTreeFFICheatcode
// TO REQUIRE: Bench, Control.Bind, Data.Function, Data.Show, Effect, Effect.Console, Prelude, Test.RBTreeFFICheatcode
require_once __DIR__ . '/../Bench/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.RBTreeFFICheatcode/index.php';

if (!class_exists(__NAMESPACE__ . '\\Phpurs_Data0')) {
  class Phpurs_Data0 { public $tag; public function __construct($t) { $this->tag = $t; } }
  class Phpurs_Data1 { public $tag; public $value0; public function __construct($t, $value0) { $this->tag = $t; $this->value0 = $value0; } }
  class Phpurs_Data2 { public $tag; public $value0, $value1; public function __construct($t, $value0, $value1) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; } }
  class Phpurs_Data3 { public $tag; public $value0, $value1, $value2; public function __construct($t, $value0, $value1, $value2) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; } }
  class Phpurs_Data4 { public $tag; public $value0, $value1, $value2, $value3; public function __construct($t, $value0, $value1, $value2, $value3) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; } }
  class Phpurs_Data5 { public $tag; public $value0, $value1, $value2, $value3, $value4; public function __construct($t, $value0, $value1, $value2, $value3, $value4) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; } }
  class Phpurs_Data6 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; } }
  class Phpurs_Data7 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; } }
  class Phpurs_Data8 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; } }
  class Phpurs_Data9 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; } }
  class Phpurs_Data10 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; } }
  class Phpurs_Data11 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; $this->value10 = $value10; } }
  class Phpurs_Data12 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10, $value11; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10, $value11) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; $this->value10 = $value10; $this->value11 = $value11; } }
}
if (!\function_exists(__NAMESPACE__ . '\\phpurs_curry_fallback')) {
  function phpurs_curry_fallback($fn, $args, $expected) {
    $missing = $expected - \count($args);
    if ($missing === 1) {
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
    if ($missing === 2) {
      return function($a, $b = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 2) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b;
        return $fn(...$args);
      };
    }
    if ($missing === 3) {
      return function($a, $b = null, $c = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 2) { $args[] = $a; $args[] = $b; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 3) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b; $args[] = $c;
        return $fn(...$args);
      };
    }
    if ($missing === 4) {
      return function($a, $b = null, $c = null, $d = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 2) { $args[] = $a; $args[] = $b; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 3) { $args[] = $a; $args[] = $b; $args[] = $c; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 4) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b; $args[] = $c; $args[] = $d;
        return $fn(...$args);
      };
    }
    return function(...$more) use ($fn, $args, $expected) {
      $merged = \array_merge($args, $more);
      if (\count($merged) >= $expected) {
        $res = $fn(...\array_slice($merged, 0, $expected));
        if (\count($merged) > $expected) {
          return $res(...\array_slice($merged, $expected));
        }
        return $res;
      }
      return phpurs_curry_fallback($fn, $merged, $expected);
    };
  }
}
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Test_RBTreeFFICheatcode = \call_user_func(function() {
  $exports = [];
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
  return $exports;
});
function majTest_majRmajBmajTreemajFmajFmajImajCheatcode_runmajRmajBmajTreemajFmajFmajImajCheatcode(int $v0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majTest_majRmajBmajTreemajFmajFmajImajCheatcode_runmajRmajBmajTreemajFmajFmajImajCheatcode';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Test_RBTreeFFICheatcode;
  $f = (\array_key_exists('runRBTreeFFICheatcode', $ffi_Test_RBTreeFFICheatcode) ? $ffi_Test_RBTreeFFICheatcode['runRBTreeFFICheatcode'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Test_RBTreeFFICheatcode_runRBTreeFFICheatcode'] = __NAMESPACE__ . '\\majTest_majRmajBmajTreemajFmajFmajImajCheatcode_runmajRmajBmajTreemajFmajFmajImajCheatcode';





// Test_RBTreeFFICheatcode_describe
$GLOBALS['Test_RBTreeFFICheatcode_describe'] = \Effect\Console\majEffect_majConsole_log("Red-Black Tree FFICheatcode (100k Worst-Case Insertions):");

// Test_RBTreeFFICheatcode_act
$GLOBALS['Test_RBTreeFFICheatcode_act'] = (function() use (&$__fn) {
$__local_var_0_0 = \Bench\majBench_opaque(100000);
return function() use ($__local_var_0_0, &$__fn) {
$dummy_1_1 = phpurs_execute_effect($__local_var_0_0);
return phpurs_execute_effect(\Effect\Console\majEffect_majConsole_log(\Data\Show\majData_majShow_showmajIntmajImpl(\Test\RBTreeFFICheatcode\majTest_majRmajBmajTreemajFmajFmajImajCheatcode_runmajRmajBmajTreemajFmajFmajImajCheatcode($dummy_1_1))));
};
})();

