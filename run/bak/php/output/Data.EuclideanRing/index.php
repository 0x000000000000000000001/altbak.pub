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
\PhpursThunks::$thunks['Data_EuclideanRing_mod'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_EuclideanRing_mod"), recVars=[];
  $__res = ($dict_0)->mod;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_EuclideanRing_gcd'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $dictEuclideanRing_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_EuclideanRing_gcd"), recVars=["Data_EuclideanRing_gcd"];
  while (true) {
$zero_2_0 = ((((((($dictEuclideanRing_1)->CommutativeRing0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Ring0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Semiring0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->zero;
$__res = (function() use ($dictEq_0, $dictEuclideanRing_1, $zero_2_0) {
  $__fn = function($a_3, $b_4 = null) use ($dictEq_0, $dictEuclideanRing_1, $zero_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_EuclideanRing_gcd"];
  if (((($dictEq_0)->eq)($b_4))($zero_2_0)) {
$__t1 = $a_3;
} else {
$__t1 = ((((($GLOBALS['Data_EuclideanRing_gcd'] ?? \PhpursThunks::eval('Data_EuclideanRing_gcd')))($dictEq_0))($dictEuclideanRing_1))($b_4))(((($dictEuclideanRing_1)->mod)($a_3))($b_4));
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_EuclideanRing_euclideanRingNumber'] = function() { $v = (object)["degree" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = 1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "div" => ($GLOBALS['Data_EuclideanRing_numDiv'] ?? \PhpursThunks::eval('Data_EuclideanRing_numDiv')), "mod" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = 0.0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "CommutativeRing0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_CommutativeRing_commutativeRingNumber'] ?? \PhpursThunks::eval('Data_CommutativeRing_commutativeRingNumber'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_EuclideanRing_euclideanRingInt'] = function() { $v = (object)["degree" => ($GLOBALS['Data_EuclideanRing_intDegree'] ?? \PhpursThunks::eval('Data_EuclideanRing_intDegree')), "div" => ($GLOBALS['Data_EuclideanRing_intDiv'] ?? \PhpursThunks::eval('Data_EuclideanRing_intDiv')), "mod" => ($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')), "CommutativeRing0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_CommutativeRing_commutativeRingInt'] ?? \PhpursThunks::eval('Data_CommutativeRing_commutativeRingInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_EuclideanRing_div'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_EuclideanRing_div"), recVars=[];
  $__res = ($dict_0)->div;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_EuclideanRing_lcm'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $dictEuclideanRing_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_EuclideanRing_lcm"), recVars=[];
  $Semiring0_2_0 = (((((($dictEuclideanRing_1)->CommutativeRing0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Ring0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Semiring0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $zero_3_1 = ($Semiring0_2_0)->zero;
  $gcd2_4_2 = ((($GLOBALS['Data_EuclideanRing_gcd'] ?? \PhpursThunks::eval('Data_EuclideanRing_gcd')))($dictEq_0))($dictEuclideanRing_1);
  $__res = (function() use ($Semiring0_2_0, $dictEq_0, $dictEuclideanRing_1, $gcd2_4_2, $zero_3_1) {
  $__fn = function($a_5, $b_6 = null) use ($Semiring0_2_0, $dictEq_0, $dictEuclideanRing_1, $gcd2_4_2, $zero_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (((($GLOBALS['Data_HeytingAlgebra_boolDisj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolDisj')))(((($dictEq_0)->eq)($a_5))($zero_3_1)))(((($dictEq_0)->eq)($b_6))($zero_3_1))) {
$__t3 = $zero_3_1;
} else {
$__t3 = ((($dictEuclideanRing_1)->div)(((($Semiring0_2_0)->mul)($a_5))($b_6)))((($gcd2_4_2)($a_5))($b_6));
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_EuclideanRing_degree'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_EuclideanRing_degree"), recVars=[];
  $__res = ($dict_0)->degree;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
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
\PhpursThunks::$thunks['Data_EuclideanRing_intDegree'] = function() use (&$ffi_Data_EuclideanRing) { return $ffi_Data_EuclideanRing['intDegree']; };
\PhpursThunks::$thunks['Data_EuclideanRing_intDiv'] = function() use (&$ffi_Data_EuclideanRing) { return $ffi_Data_EuclideanRing['intDiv']; };
\PhpursThunks::$thunks['Data_EuclideanRing_intMod'] = function() use (&$ffi_Data_EuclideanRing) { return $ffi_Data_EuclideanRing['intMod']; };
\PhpursThunks::$thunks['Data_EuclideanRing_numDiv'] = function() use (&$ffi_Data_EuclideanRing) { return $ffi_Data_EuclideanRing['numDiv']; };









