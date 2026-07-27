<?php

namespace Data\Int;

// ALL IMPORTS: Control.Category, Control.Semigroupoid, Data.Boolean, Data.Bounded, Data.CommutativeRing, Data.DivisionRing, Data.Eq, Data.EuclideanRing, Data.HeytingAlgebra, Data.Int, Data.Int.Bits, Data.Maybe, Data.Number, Data.Ord, Data.Ordering, Data.Ring, Data.Semiring, Data.Show, Prelude, Prim
// TO REQUIRE: Control.Category, Control.Semigroupoid, Data.Boolean, Data.Bounded, Data.CommutativeRing, Data.DivisionRing, Data.Eq, Data.EuclideanRing, Data.HeytingAlgebra, Data.Int, Data.Int.Bits, Data.Maybe, Data.Number, Data.Ord, Data.Ordering, Data.Ring, Data.Semiring, Data.Show, Prelude
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.CommutativeRing/index.php';
require_once __DIR__ . '/../Data.DivisionRing/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Int/index.php';
require_once __DIR__ . '/../Data.Int.Bits/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Number/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Prelude/index.php';

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
$ffi_Data_Int = \call_user_func(function() {
  $exports = [];
$fromNumberImpl = function($just, $nothing = null, $n = null) use (&$fromNumberImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$fromNumberImpl) {

            return $fromNumberImpl(...\array_merge($__args, $more));
        };
    }
    // JS bitwise OR 0 limits to 32-bit signed integer. 
    // PHP integers are typically 64-bit, but to match JS semantics we can just cast to int.
    return (intval($n) == $n) ? $just(intval($n)) : $nothing;
};

$toNumber = function($n) {
    return floatval($n);
};

$fromStringAsImpl = function($just, $nothing = null, $radix = null, $s = null) use (&$fromStringAsImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$fromStringAsImpl) {

            return $fromStringAsImpl(...\array_merge($__args, $more));
        };
    }
    
    if ($radix < 11) {
        $digits = "[0-" . ($radix - 1) . "]";
    } else if ($radix === 11) {
        $digits = "[0-9a]";
    } else {
        $digits = "[0-9a-" . chr(86 + $radix) . "]";
    }
    $pattern = "/^[\+\-]?" . $digits . "+$/i";

    if (preg_match($pattern, $s)) {
        $i = intval(base_convert($s, $radix, 10));
        return $just($i);
    } else {
        return $nothing;
    }
};

$toStringAs = function($radix, $i = null) use (&$toStringAs) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$toStringAs) {

            return $toStringAs(...\array_merge($__args, $more));
        };
    }
    return base_convert($i, 10, $radix);
};

$quot = function($x, $y = null) use (&$quot) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$quot) {

            return $quot(...\array_merge($__args, $more));
        };
    }
    return intdiv($x, $y);
};

$rem = function($x, $y = null) use (&$rem) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$rem) {

            return $rem(...\array_merge($__args, $more));
        };
    }
    return $x % $y;
};

$pow = function($x, $y = null) use (&$pow) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$pow) {

            return $pow(...\array_merge($__args, $more));
        };
    }
    return intval(pow($x, $y));
};

$exports['fromNumberImpl'] = $fromNumberImpl;
$exports['toNumber'] = $toNumber;
$exports['fromStringAsImpl'] = $fromStringAsImpl;
$exports['toStringAs'] = $toStringAs;
$exports['quot'] = $quot;
$exports['rem'] = $rem;
$exports['pow'] = $pow;
return $exports;
  return $exports;
});
function majData_majInt_frommajNumbermajImpl($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majInt_frommajNumbermajImpl';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_Int;
  $f = ($ffi_Data_Int['fromNumberImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_Int_fromNumberImpl'] = __NAMESPACE__ . '\\majData_majInt_frommajNumbermajImpl';

function majData_majInt_frommajStringmajAsmajImpl($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majInt_frommajStringmajAsmajImpl';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_Int;
  $f = ($ffi_Data_Int['fromStringAsImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_Int_fromStringAsImpl'] = __NAMESPACE__ . '\\majData_majInt_frommajStringmajAsmajImpl';

function majData_majInt_pow(int $v0, $v1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majInt_pow';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Int;
  $f = ($ffi_Data_Int['pow'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Int_pow'] = __NAMESPACE__ . '\\majData_majInt_pow';

function majData_majInt_quot(int $v0, $v1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majInt_quot';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Int;
  $f = ($ffi_Data_Int['quot'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Int_quot'] = __NAMESPACE__ . '\\majData_majInt_quot';

function majData_majInt_rem(int $v0, $v1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majInt_rem';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Int;
  $f = ($ffi_Data_Int['rem'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Int_rem'] = __NAMESPACE__ . '\\majData_majInt_rem';

function majData_majInt_tomajNumber(int $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majInt_tomajNumber';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Int;
  $f = ($ffi_Data_Int['toNumber'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Int_toNumber'] = __NAMESPACE__ . '\\majData_majInt_tomajNumber';

function majData_majInt_tomajStringmajAs(int $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majInt_tomajStringmajAs';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_Int;
  $f = ($ffi_Data_Int['toStringAs'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_Int_toStringAs'] = __NAMESPACE__ . '\\majData_majInt_tomajStringmajAs';



final class Data_Int_Even { public $tag = 'Even'; public function __construct() {} }
final class Data_Int_Odd { public $tag = 'Odd'; public function __construct() {} }

// Data_Int_greaterThanOrEq
$GLOBALS['Data_Int_greaterThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Int_lessThanOrEq
$GLOBALS['Data_Int_lessThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Int_Even
$GLOBALS['Data_Int_Even'] = ($GLOBALS['__phpurs_data0_Even'] ??= new \Data\Int\Data_Int_Even());

// Data_Int_Odd
$GLOBALS['Data_Int_Odd'] = ($GLOBALS['__phpurs_data0_Odd'] ??= new \Data\Int\Data_Int_Odd());

// Data_Int_showParity
$GLOBALS['Data_Int_showParity'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Int\Data_Int_Even) {
$__t0 = "Even";
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Int\Data_Int_Odd) {
$__t0 = "Odd";
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Int_radix
function majData_majInt_radix(int $n_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_radix';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Int_greaterThanOrEq'])($n_0))(2)))((($GLOBALS['Data_Int_lessThanOrEq'])($n_0))(36))) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($n_0);
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_radix'] = __NAMESPACE__ . '\\majData_majInt_radix';

// Data_Int_odd
function majData_majInt_odd(int $x_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_odd';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})((($x_0 & 1) === 0)))(false);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_odd'] = __NAMESPACE__ . '\\majData_majInt_odd';

// Data_Int_octal
$GLOBALS['Data_Int_octal'] = 8;

// Data_Int_hexadecimal
$GLOBALS['Data_Int_hexadecimal'] = 16;

// Data_Int_fromStringAs_closure
$GLOBALS['Data_Int_fromStringAs_closure'] = (($GLOBALS['Data_Int_fromStringAsImpl'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_Int_fromStringAs
function majData_majInt_frommajStringmajAs(int $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_frommajStringmajAs';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Int_fromStringAs_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Int_fromStringAs'] = __NAMESPACE__ . '\\majData_majInt_frommajStringmajAs';

// Data_Int_fromString_closure
$GLOBALS['Data_Int_fromString_closure'] = ($GLOBALS['Data_Int_fromStringAs'])(10);

// Data_Int_fromString
function majData_majInt_frommajString(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_frommajString';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Int_fromString_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_fromString'] = __NAMESPACE__ . '\\majData_majInt_frommajString';

// Data_Int_fromNumber_closure
$GLOBALS['Data_Int_fromNumber_closure'] = (($GLOBALS['Data_Int_fromNumberImpl'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_Int_fromNumber
function majData_majInt_frommajNumber(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_frommajNumber';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Int_fromNumber_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_fromNumber'] = __NAMESPACE__ . '\\majData_majInt_frommajNumber';

// Data_Int_unsafeClamp
function majData_majInt_unsafemajClamp(float $x_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_unsafemajClamp';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t2 = null;;
  if ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'not'})(\Data\Number\majData_majNumber_ismajFinite($x_0))) {
$__t2 = 0;
goto end_branch_2;;
};
  if (($x_0 >= \Data\Int\majData_majInt_tomajNumber(($GLOBALS['Data_Bounded_boundedInt'])->{'top'}))) {
$__t2 = ($GLOBALS['Data_Bounded_boundedInt'])->{'top'};
goto end_branch_2;;
};
  if (($x_0 <= \Data\Int\majData_majInt_tomajNumber(($GLOBALS['Data_Bounded_boundedInt'])->{'bottom'}))) {
$__t2 = ($GLOBALS['Data_Bounded_boundedInt'])->{'bottom'};
goto end_branch_2;;
};
  $__local_var_1_0 = \Data\Int\majData_majInt_frommajNumber($x_0);
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($__local_var_1_0)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__t2 = $__t1;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_unsafeClamp'] = __NAMESPACE__ . '\\majData_majInt_unsafemajClamp';

// Data_Int_round_closure
$GLOBALS['Data_Int_round_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Int_unsafeClamp']))($GLOBALS['Data_Number_round']);

// Data_Int_round
function majData_majInt_round(float $v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_round';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Int_round_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_round'] = __NAMESPACE__ . '\\majData_majInt_round';

// Data_Int_trunc_closure
$GLOBALS['Data_Int_trunc_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Int_unsafeClamp']))($GLOBALS['Data_Number_trunc']);

// Data_Int_trunc
function majData_majInt_trunc(float $v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_trunc';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Int_trunc_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_trunc'] = __NAMESPACE__ . '\\majData_majInt_trunc';

// Data_Int_floor_closure
$GLOBALS['Data_Int_floor_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Int_unsafeClamp']))($GLOBALS['Data_Number_floor']);

// Data_Int_floor
function majData_majInt_floor(float $v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_floor';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Int_floor_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_floor'] = __NAMESPACE__ . '\\majData_majInt_floor';

// Data_Int_even
function majData_majInt_even(int $x_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_even';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($x_0 & 1) === 0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_even'] = __NAMESPACE__ . '\\majData_majInt_even';

// Data_Int_parity
function majData_majInt_parity(int $n_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_parity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = match (($n_0 & 1)) { 0 => new \Data\Int\Data_Int_Even(), default => new \Data\Int\Data_Int_Odd() };
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_parity'] = __NAMESPACE__ . '\\majData_majInt_parity';

// Data_Int_eqParity
$GLOBALS['Data_Int_eqParity'] = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_0 instanceof \Data\Int\Data_Int_Even) {
$__t0 = $y_1 instanceof \Data\Int\Data_Int_Even;
goto end_branch_0;;
};
  $__t0 = ($x_0 instanceof \Data\Int\Data_Int_Odd && $y_1 instanceof \Data\Int\Data_Int_Odd);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Int_ordParity
$GLOBALS['Data_Int_ordParity'] = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_0 instanceof \Data\Int\Data_Int_Even) {
$__t1 = null;;
if ($y_1 instanceof \Data\Int\Data_Int_Even) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($y_1 instanceof \Data\Int\Data_Int_Even) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if (($x_0 instanceof \Data\Int\Data_Int_Odd && $y_1 instanceof \Data\Int\Data_Int_Odd)) {
$__t0 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Int_eqParity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Int_semiringParity
$GLOBALS['Data_Int_semiringParity'] = (object)["zero" => new \Data\Int\Data_Int_Even(), "add" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (((($GLOBALS['Data_Int_eqParity'])->{'eq'})($x_0))($y_1)) {
$__t0 = new \Data\Int\Data_Int_Even();
goto end_branch_0;;
};
  $__t0 = new \Data\Int\Data_Int_Odd();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => new \Data\Int\Data_Int_Odd(), "mul" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if (($v_0 instanceof \Data\Int\Data_Int_Odd && $v1_1 instanceof \Data\Int\Data_Int_Odd)) {
$__t1 = new \Data\Int\Data_Int_Odd();
goto end_branch_1;;
};
  $__t1 = new \Data\Int\Data_Int_Even();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Int_ringParity
$GLOBALS['Data_Int_ringParity'] = (object)["sub" => ($GLOBALS['Data_Int_semiringParity'])->{'add'}, "Semiring0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Int_semiringParity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Int_divisionRingParity
$GLOBALS['Data_Int_divisionRingParity'] = (object)["recip" => function($x_0) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ring0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Int_ringParity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Int_decimal
$GLOBALS['Data_Int_decimal'] = 10;

// Data_Int_commutativeRingParity
$GLOBALS['Data_Int_commutativeRingParity'] = (object)["Ring0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Int_ringParity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Int_euclideanRingParity
$GLOBALS['Data_Int_euclideanRingParity'] = (object)["degree" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Int\Data_Int_Even) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Int\Data_Int_Odd) {
$__t0 = 1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "div" => (function() {
  $__fn = function($x_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "mod" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Int\Data_Int_Even();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "CommutativeRing0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Int_commutativeRingParity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Int_ceil_closure
$GLOBALS['Data_Int_ceil_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Int_unsafeClamp']))($GLOBALS['Data_Number_ceil']);

// Data_Int_ceil
function majData_majInt_ceil(float $v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInt_ceil';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Int_ceil_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Int_ceil'] = __NAMESPACE__ . '\\majData_majInt_ceil';

// Data_Int_boundedParity
$GLOBALS['Data_Int_boundedParity'] = (object)["bottom" => new \Data\Int\Data_Int_Even(), "top" => new \Data\Int\Data_Int_Odd(), "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Int_ordParity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Int_binary
$GLOBALS['Data_Int_binary'] = 2;

// Data_Int_base36
$GLOBALS['Data_Int_base36'] = 36;

