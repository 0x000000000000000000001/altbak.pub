<?php

namespace Data\Function\Uncurried;

// ALL IMPORTS: Data.Unit, Prim
// TO REQUIRE: Data.Unit
require_once __DIR__ . '/../Data.Unit/index.php';

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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Function_Uncurried = \call_user_func(function() {
  $exports = [];
$mkFn0 = function($fn) use (&$mkFn0) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn0) {
            return $mkFn0(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn) {
        return $fn(null);
    };
};
$mkFn1 = function($fn) use (&$mkFn1) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn1) {
            return $mkFn1(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null) use ($fn) {
        return $fn($a0);
    };
};
$mkFn2 = function($fn) use (&$mkFn2) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn2) {
            return $mkFn2(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null) use ($fn) {
        return $fn($a0)($a1);
    };
};
$mkFn3 = function($fn) use (&$mkFn3) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn3) {
            return $mkFn3(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null, $a2 = null) use ($fn) {
        return $fn($a0)($a1)($a2);
    };
};
$mkFn4 = function($fn) use (&$mkFn4) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn4) {
            return $mkFn4(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null, $a2 = null, $a3 = null) use ($fn) {
        return $fn($a0)($a1)($a2)($a3);
    };
};
$mkFn5 = function($fn) use (&$mkFn5) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn5) {
            return $mkFn5(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null) use ($fn) {
        return $fn($a0)($a1)($a2)($a3)($a4);
    };
};
$mkFn6 = function($fn) use (&$mkFn6) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn6) {
            return $mkFn6(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null) use ($fn) {
        return $fn($a0)($a1)($a2)($a3)($a4)($a5);
    };
};
$mkFn7 = function($fn) use (&$mkFn7) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn7) {
            return $mkFn7(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null, $a6 = null) use ($fn) {
        return $fn($a0)($a1)($a2)($a3)($a4)($a5)($a6);
    };
};
$mkFn8 = function($fn) use (&$mkFn8) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn8) {
            return $mkFn8(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null, $a6 = null, $a7 = null) use ($fn) {
        return $fn($a0)($a1)($a2)($a3)($a4)($a5)($a6)($a7);
    };
};
$mkFn9 = function($fn) use (&$mkFn9) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn9) {
            return $mkFn9(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null, $a6 = null, $a7 = null, $a8 = null) use ($fn) {
        return $fn($a0)($a1)($a2)($a3)($a4)($a5)($a6)($a7)($a8);
    };
};
$mkFn10 = function($fn) use (&$mkFn10) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mkFn10) {
            return $mkFn10(...\array_merge($__args, $more));
        };
    }
    return function($a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null, $a6 = null, $a7 = null, $a8 = null, $a9 = null) use ($fn) {
        return $fn($a0)($a1)($a2)($a3)($a4)($a5)($a6)($a7)($a8)($a9);
    };
};
$runFn0 = function($fn = null) use (&$runFn0) {
    if (\func_num_args() < 1) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn0) {
            return $runFn0(...\array_merge($__args, $more));
        };
    }
    return $fn();
};
$runFn1 = function($fn = null, $a0 = null) use (&$runFn1) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn1) {
            return $runFn1(...\array_merge($__args, $more));
        };
    }
    return $fn($a0);
};
$runFn2 = function($fn = null, $a0 = null, $a1 = null) use (&$runFn2) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn2) {
            return $runFn2(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1);
};
$runFn3 = function($fn = null, $a0 = null, $a1 = null, $a2 = null) use (&$runFn3) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn3) {
            return $runFn3(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1, $a2);
};
$runFn4 = function($fn = null, $a0 = null, $a1 = null, $a2 = null, $a3 = null) use (&$runFn4) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn4) {
            return $runFn4(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1, $a2, $a3);
};
$runFn5 = function($fn = null, $a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null) use (&$runFn5) {
    if (\func_num_args() < 6) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn5) {
            return $runFn5(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1, $a2, $a3, $a4);
};
$runFn6 = function($fn = null, $a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null) use (&$runFn6) {
    if (\func_num_args() < 7) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn6) {
            return $runFn6(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1, $a2, $a3, $a4, $a5);
};
$runFn7 = function($fn = null, $a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null, $a6 = null) use (&$runFn7) {
    if (\func_num_args() < 8) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn7) {
            return $runFn7(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1, $a2, $a3, $a4, $a5, $a6);
};
$runFn8 = function($fn = null, $a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null, $a6 = null, $a7 = null) use (&$runFn8) {
    if (\func_num_args() < 9) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn8) {
            return $runFn8(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1, $a2, $a3, $a4, $a5, $a6, $a7);
};
$runFn9 = function($fn = null, $a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null, $a6 = null, $a7 = null, $a8 = null) use (&$runFn9) {
    if (\func_num_args() < 10) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn9) {
            return $runFn9(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1, $a2, $a3, $a4, $a5, $a6, $a7, $a8);
};
$runFn10 = function($fn = null, $a0 = null, $a1 = null, $a2 = null, $a3 = null, $a4 = null, $a5 = null, $a6 = null, $a7 = null, $a8 = null, $a9 = null) use (&$runFn10) {
    if (\func_num_args() < 11) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runFn10) {
            return $runFn10(...\array_merge($__args, $more));
        };
    }
    return $fn($a0, $a1, $a2, $a3, $a4, $a5, $a6, $a7, $a8, $a9);
};
$exports['mkFn0'] = $mkFn0;
$exports['runFn0'] = $runFn0;
$exports['mkFn1'] = $mkFn1;
$exports['runFn1'] = $runFn1;
$exports['mkFn2'] = $mkFn2;
$exports['runFn2'] = $runFn2;
$exports['mkFn3'] = $mkFn3;
$exports['runFn3'] = $runFn3;
$exports['mkFn4'] = $mkFn4;
$exports['runFn4'] = $runFn4;
$exports['mkFn5'] = $mkFn5;
$exports['runFn5'] = $runFn5;
$exports['mkFn6'] = $mkFn6;
$exports['runFn6'] = $runFn6;
$exports['mkFn7'] = $mkFn7;
$exports['runFn7'] = $runFn7;
$exports['mkFn8'] = $mkFn8;
$exports['runFn8'] = $runFn8;
$exports['mkFn9'] = $mkFn9;
$exports['runFn9'] = $runFn9;
$exports['mkFn10'] = $mkFn10;
$exports['runFn10'] = $runFn10;
return $exports;
  return $exports;
});
function majData_majFunction_majUncurried_mkmajFn0($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn0';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn0'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn0'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn0';

function majData_majFunction_majUncurried_mkmajFn10($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn10';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn10'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn10'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn10';

function majData_majFunction_majUncurried_mkmajFn2($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn2'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn2'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn2';

function majData_majFunction_majUncurried_mkmajFn3($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn3';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn3'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn3'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn3';

function majData_majFunction_majUncurried_mkmajFn4($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn4';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn4'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn4'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn4';

function majData_majFunction_majUncurried_mkmajFn5($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn5';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn5'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn5'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn5';

function majData_majFunction_majUncurried_mkmajFn6($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn6';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn6'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn6'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn6';

function majData_majFunction_majUncurried_mkmajFn7($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn7';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn7'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn7'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn7';

function majData_majFunction_majUncurried_mkmajFn8($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn8';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn8'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn8'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn8';

function majData_majFunction_majUncurried_mkmajFn9($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn9';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['mkFn9'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_mkFn9'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn9';

function majData_majFunction_majUncurried_runmajFn0($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn0';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn0'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Function_Uncurried_runFn0'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn0';

function majData_majFunction_majUncurried_runmajFn10($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null, $v7 = null, $v8 = null, $v9 = null, $v10 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn10';
  if ($__num < 11) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 11);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn10'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6, $v7, $v8, $v9, $v10);
}
$GLOBALS['Data_Function_Uncurried_runFn10'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn10';

function majData_majFunction_majUncurried_runmajFn2($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn2';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn2'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_Function_Uncurried_runFn2'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn2';

function majData_majFunction_majUncurried_runmajFn3($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn3';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn3'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_Function_Uncurried_runFn3'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn3';

function majData_majFunction_majUncurried_runmajFn4($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn4';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn4'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_Function_Uncurried_runFn4'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn4';

function majData_majFunction_majUncurried_runmajFn5($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn5';
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn5'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5);
}
$GLOBALS['Data_Function_Uncurried_runFn5'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn5';

function majData_majFunction_majUncurried_runmajFn6($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn6';
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn6'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6);
}
$GLOBALS['Data_Function_Uncurried_runFn6'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn6';

function majData_majFunction_majUncurried_runmajFn7($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null, $v7 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn7';
  if ($__num < 8) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 8);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn7'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6, $v7);
}
$GLOBALS['Data_Function_Uncurried_runFn7'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn7';

function majData_majFunction_majUncurried_runmajFn8($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null, $v7 = null, $v8 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn8';
  if ($__num < 9) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 9);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn8'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6, $v7, $v8);
}
$GLOBALS['Data_Function_Uncurried_runFn8'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn8';

function majData_majFunction_majUncurried_runmajFn9($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null, $v7 = null, $v8 = null, $v9 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn9';
  if ($__num < 10) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 10);
  }
  global $ffi_Data_Function_Uncurried;
  $f = ($ffi_Data_Function_Uncurried['runFn9'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6, $v7, $v8, $v9);
}
$GLOBALS['Data_Function_Uncurried_runFn9'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn9';





// Data_Function_Uncurried_runFn1
function majData_majFunction_majUncurried_runmajFn1($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunction_majUncurried_runmajFn1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $f_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Function_Uncurried_runFn1'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_runmajFn1';

// Data_Function_Uncurried_mkFn1
function majData_majFunction_majUncurried_mkmajFn1($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunction_majUncurried_mkmajFn1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $f_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Function_Uncurried_mkFn1'] = __NAMESPACE__ . '\\majData_majFunction_majUncurried_mkmajFn1';

