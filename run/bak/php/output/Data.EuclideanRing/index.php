<?php

namespace Data\EuclideanRing;

// ALL IMPORTS: Data.BooleanAlgebra, Data.CommutativeRing, Data.Eq, Data.EuclideanRing, Data.HeytingAlgebra, Data.Ring, Data.Semiring, Prim
// TO REQUIRE: Data.BooleanAlgebra, Data.CommutativeRing, Data.Eq, Data.EuclideanRing, Data.HeytingAlgebra, Data.Ring, Data.Semiring
require_once __DIR__ . '/../Data.BooleanAlgebra/index.php';
require_once __DIR__ . '/../Data.CommutativeRing/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';

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
$ffi_Data_EuclideanRing = \call_user_func(function() {
  $exports = [];
$intDegree = function($x) use (&$intDegree) { return abs($x); };
$intDiv = function($x, $y = null) use (&$intDiv) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$intDiv) {

            return $intDiv(...\array_merge($__args, $more));
        };
    }
    return (int)($x / $y);
};
$intMod = function($x, $y = null) use (&$intMod) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$intMod) {

            return $intMod(...\array_merge($__args, $more));
        };
    }
    return $x % $y;
};
$numDiv = function($x, $y = null) use (&$numDiv) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$numDiv) {

            return $numDiv(...\array_merge($__args, $more));
        };
    }
    return $x / $y;
};

$exports['intDegree'] = $intDegree;
$exports['intDiv'] = $intDiv;
$exports['intMod'] = $intMod;
$exports['numDiv'] = $numDiv;
return $exports;
  return $exports;
});
function majData_majEuclideanmajRing_intmajDegree(int $v0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majEuclideanmajRing_intmajDegree';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_EuclideanRing;
  $f = ($ffi_Data_EuclideanRing['intDegree'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_EuclideanRing_intDegree'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_intmajDegree';

function majData_majEuclideanmajRing_intmajDiv(int $v0, $v1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majEuclideanmajRing_intmajDiv';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_EuclideanRing;
  $f = ($ffi_Data_EuclideanRing['intDiv'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_EuclideanRing_intDiv'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_intmajDiv';

function majData_majEuclideanmajRing_intmajMod(int $v0, $v1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majEuclideanmajRing_intmajMod';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_EuclideanRing;
  $f = ($ffi_Data_EuclideanRing['intMod'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_EuclideanRing_intMod'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_intmajMod';

function majData_majEuclideanmajRing_nummajDiv(float $v0, $v1 = null): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majEuclideanmajRing_nummajDiv';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_EuclideanRing;
  $f = ($ffi_Data_EuclideanRing['numDiv'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_EuclideanRing_numDiv'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_nummajDiv';





// Data_EuclideanRing_mod
function majData_majEuclideanmajRing_mod($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEuclideanmajRing_mod';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'mod'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_EuclideanRing_mod'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_mod';

// Data_EuclideanRing_gcd
function majData_majEuclideanmajRing_gcd($dictEq_0, $dictEuclideanRing_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEuclideanmajRing_gcd';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_EuclideanRing_gcd_dictEq_0 = $dictEq_0;
  $__tco_var_Data_EuclideanRing_gcd_dictEuclideanRing_1 = $dictEuclideanRing_1;
  tco_loop_Data_EuclideanRing_gcd:;
  $dictEq_0 = $__tco_var_Data_EuclideanRing_gcd_dictEq_0;
  $dictEuclideanRing_1 = $__tco_var_Data_EuclideanRing_gcd_dictEuclideanRing_1;
  $zero_2_0 = ((((((($dictEuclideanRing_1)->{'CommutativeRing0'})(null))->{'Ring0'})(null))->{'Semiring0'})(null))->{'zero'};
  $__res = (function() use ($dictEq_0, $dictEuclideanRing_1, $zero_2_0) {
  $__fn = function($a_3, $b_4 = null) use ($dictEq_0, $dictEuclideanRing_1, $zero_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if (((($dictEq_0)->{'eq'})($b_4))($zero_2_0)) {
$__t1 = $a_3;
goto end_branch_1;;
};
  $__t1 = \Data\EuclideanRing\majData_majEuclideanmajRing_gcd($dictEq_0, $dictEuclideanRing_1, $b_4, ((($dictEuclideanRing_1)->{'mod'})($a_3))($b_4));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_EuclideanRing_gcd'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_gcd';

// Data_EuclideanRing_euclideanRingNumber
$GLOBALS['Data_EuclideanRing_euclideanRingNumber'] = (object)["degree" => function($v_0) {
  $__num = \func_num_args();
  $__res = 1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "div" => $GLOBALS['Data_EuclideanRing_numDiv'], "mod" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = 0.0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "CommutativeRing0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_CommutativeRing_commutativeRingNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_EuclideanRing_euclideanRingInt
$GLOBALS['Data_EuclideanRing_euclideanRingInt'] = (object)["degree" => $GLOBALS['Data_EuclideanRing_intDegree'], "div" => $GLOBALS['Data_EuclideanRing_intDiv'], "mod" => $GLOBALS['Data_EuclideanRing_intMod'], "CommutativeRing0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_CommutativeRing_commutativeRingInt'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_EuclideanRing_div
function majData_majEuclideanmajRing_div($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEuclideanmajRing_div';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'div'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_EuclideanRing_div'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_div';

// Data_EuclideanRing_lcm
function majData_majEuclideanmajRing_lcm($dictEq_0, $dictEuclideanRing_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEuclideanmajRing_lcm';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Semiring0_2_0 = (((((($dictEuclideanRing_1)->{'CommutativeRing0'})(null))->{'Ring0'})(null))->{'Semiring0'})(null);
  $zero_3_1 = ($Semiring0_2_0)->{'zero'};
  $gcd2_4_2 = (($GLOBALS['Data_EuclideanRing_gcd'])($dictEq_0))($dictEuclideanRing_1);
  $__res = (function() use ($Semiring0_2_0, $dictEq_0, $dictEuclideanRing_1, $gcd2_4_2, $zero_3_1) {
  $__fn = function($a_5, $b_6 = null) use ($Semiring0_2_0, $dictEq_0, $dictEuclideanRing_1, $gcd2_4_2, $zero_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(((($dictEq_0)->{'eq'})($a_5))($zero_3_1)))(((($dictEq_0)->{'eq'})($b_6))($zero_3_1))) {
$__t3 = $zero_3_1;
goto end_branch_3;;
};
  $__t3 = ((($dictEuclideanRing_1)->{'div'})(((($Semiring0_2_0)->{'mul'})($a_5))($b_6)))((($gcd2_4_2)($a_5))($b_6));
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_EuclideanRing_lcm'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_lcm';

// Data_EuclideanRing_degree
function majData_majEuclideanmajRing_degree($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEuclideanmajRing_degree';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'degree'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_EuclideanRing_degree'] = __NAMESPACE__ . '\\majData_majEuclideanmajRing_degree';

