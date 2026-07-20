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
\PhpursThunks::$thunks['Data_Int_Even'] = function() { $v = ($GLOBALS['__phpurs_data0_Even'] ??= new Phpurs_Data0("Even")); return $v; };
\PhpursThunks::$thunks['Data_Int_Odd'] = function() { $v = ($GLOBALS['__phpurs_data0_Odd'] ??= new Phpurs_Data0("Odd")); return $v; };
\PhpursThunks::$thunks['Data_Int_showParity'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Even"))) {
$__t0 = "Even";
} else {
if ((is_object($v_0) && (($v_0)->tag === "Odd"))) {
$__t0 = "Odd";
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Int_radix'] = function() { $v = function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Int_radix"), recVars=[];
  if (((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(2)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(2))->tag === "LT")))))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(36)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(36))->tag === "GT"))))) {
$__t0 = new Phpurs_Data1("Just", $n_0);
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Int_odd'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Int_odd"), recVars=[];
  $__res = ((($GLOBALS['Data_Eq_eqBooleanImpl'] ?? \PhpursThunks::eval('Data_Eq_eqBooleanImpl')))(((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(((($GLOBALS['Data_Int_Bits_and'] ?? \PhpursThunks::eval('Data_Int_Bits_and')))($x_0))(1)))(0)))(false);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Int_octal'] = function() { $v = 8; return $v; };
\PhpursThunks::$thunks['Data_Int_hexadecimal'] = function() { $v = 16; return $v; };
\PhpursThunks::$thunks['Data_Int_fromStringAs'] = function() { $v = ((($GLOBALS['Data_Int_fromStringAsImpl'] ?? \PhpursThunks::eval('Data_Int_fromStringAsImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Int_fromString'] = function() { $v = (($GLOBALS['Data_Int_fromStringAs'] ?? \PhpursThunks::eval('Data_Int_fromStringAs')))(10); return $v; };
\PhpursThunks::$thunks['Data_Int_fromNumber'] = function() { $v = ((($GLOBALS['Data_Int_fromNumberImpl'] ?? \PhpursThunks::eval('Data_Int_fromNumberImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_Int_unsafeClamp'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Int_unsafeClamp"), recVars=[];
  if ((($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot')))((($GLOBALS['Data_Number_isFinite'] ?? \PhpursThunks::eval('Data_Number_isFinite')))($x_0))) {
$__t2 = 0;
} else {
if (( ! (is_object((((($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')))->compare)($x_0))((($GLOBALS['Data_Int_toNumber'] ?? \PhpursThunks::eval('Data_Int_toNumber')))(($GLOBALS['Data_Bounded_topInt'] ?? \PhpursThunks::eval('Data_Bounded_topInt'))))) && (((((($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')))->compare)($x_0))((($GLOBALS['Data_Int_toNumber'] ?? \PhpursThunks::eval('Data_Int_toNumber')))(($GLOBALS['Data_Bounded_topInt'] ?? \PhpursThunks::eval('Data_Bounded_topInt')))))->tag === "LT")))) {
$__t2 = ($GLOBALS['Data_Bounded_topInt'] ?? \PhpursThunks::eval('Data_Bounded_topInt'));
} else {
if (( ! (is_object((((($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')))->compare)($x_0))((($GLOBALS['Data_Int_toNumber'] ?? \PhpursThunks::eval('Data_Int_toNumber')))(($GLOBALS['Data_Bounded_bottomInt'] ?? \PhpursThunks::eval('Data_Bounded_bottomInt'))))) && (((((($GLOBALS['Data_Ord_ordNumber'] ?? \PhpursThunks::eval('Data_Ord_ordNumber')))->compare)($x_0))((($GLOBALS['Data_Int_toNumber'] ?? \PhpursThunks::eval('Data_Int_toNumber')))(($GLOBALS['Data_Bounded_bottomInt'] ?? \PhpursThunks::eval('Data_Bounded_bottomInt')))))->tag === "GT")))) {
$__t2 = ($GLOBALS['Data_Bounded_bottomInt'] ?? \PhpursThunks::eval('Data_Bounded_bottomInt'));
} else {
$__local_var_1_0 = (($GLOBALS['Data_Int_fromNumber'] ?? \PhpursThunks::eval('Data_Int_fromNumber')))($x_0);
if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Nothing"))) {
$__t1 = 0;
} else {
if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__t1 = ($__local_var_1_0)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__t2 = $__t1;
};
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Int_round'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Int_unsafeClamp'] ?? \PhpursThunks::eval('Data_Int_unsafeClamp'))))(($GLOBALS['Data_Number_round'] ?? \PhpursThunks::eval('Data_Number_round'))); return $v; };
\PhpursThunks::$thunks['Data_Int_trunc'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Int_unsafeClamp'] ?? \PhpursThunks::eval('Data_Int_unsafeClamp'))))(($GLOBALS['Data_Number_trunc'] ?? \PhpursThunks::eval('Data_Number_trunc'))); return $v; };
\PhpursThunks::$thunks['Data_Int_floor'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Int_unsafeClamp'] ?? \PhpursThunks::eval('Data_Int_unsafeClamp'))))(($GLOBALS['Data_Number_floor'] ?? \PhpursThunks::eval('Data_Number_floor'))); return $v; };
\PhpursThunks::$thunks['Data_Int_even'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Int_even"), recVars=[];
  $__res = ((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(((($GLOBALS['Data_Int_Bits_and'] ?? \PhpursThunks::eval('Data_Int_Bits_and')))($x_0))(1)))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Int_parity'] = function() { $v = function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Int_parity"), recVars=[];
  if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(((($GLOBALS['Data_Int_Bits_and'] ?? \PhpursThunks::eval('Data_Int_Bits_and')))($n_0))(1)))(0)) {
$__t0 = new Phpurs_Data0("Even");
} else {
$__t0 = new Phpurs_Data0("Odd");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Int_eqParity'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "Even"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "Even"));
} else {
$__t0 = ((is_object($x_0) && (($x_0)->tag === "Odd")) && (is_object($y_1) && (($y_1)->tag === "Odd")));
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Int_ordParity'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "Even"))) {
if ((is_object($y_1) && (($y_1)->tag === "Even"))) {
$__t1 = new Phpurs_Data0("EQ");
} else {
$__t1 = new Phpurs_Data0("LT");
};
$__t0 = $__t1;
} else {
if ((is_object($y_1) && (($y_1)->tag === "Even"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if (((is_object($x_0) && (($x_0)->tag === "Odd")) && (is_object($y_1) && (($y_1)->tag === "Odd")))) {
$__t0 = new Phpurs_Data0("EQ");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Int_eqParity'] ?? \PhpursThunks::eval('Data_Int_eqParity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Int_semiringParity'] = function() { $v = (object)["zero" => new Phpurs_Data0("Even"), "add" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((function() use ($x_0, $y_1, &$__fn) {
if ((is_object($x_0) && (($x_0)->tag === "Even"))) {
$__t1 = (is_object($y_1) && (($y_1)->tag === "Even"));
} else {
$__t1 = ((is_object($x_0) && (($x_0)->tag === "Odd")) && (is_object($y_1) && (($y_1)->tag === "Odd")));
};
return $__t1;
})()) {
$__t0 = new Phpurs_Data0("Even");
} else {
$__t0 = new Phpurs_Data0("Odd");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => new Phpurs_Data0("Odd"), "mul" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (((is_object($v_0) && (($v_0)->tag === "Odd")) && (is_object($v1_1) && (($v1_1)->tag === "Odd")))) {
$__t2 = new Phpurs_Data0("Odd");
} else {
$__t2 = new Phpurs_Data0("Even");
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Int_ringParity'] = function() { $v = (object)["sub" => (($GLOBALS['Data_Int_semiringParity'] ?? \PhpursThunks::eval('Data_Int_semiringParity')))->add, "Semiring0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Int_semiringParity'] ?? \PhpursThunks::eval('Data_Int_semiringParity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Int_divisionRingParity'] = function() { $v = (object)["recip" => (($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity, "Ring0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Int_ringParity'] ?? \PhpursThunks::eval('Data_Int_ringParity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Int_decimal'] = function() { $v = 10; return $v; };
\PhpursThunks::$thunks['Data_Int_commutativeRingParity'] = function() { $v = (object)["Ring0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Int_ringParity'] ?? \PhpursThunks::eval('Data_Int_ringParity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Int_euclideanRingParity'] = function() { $v = (object)["degree" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Even"))) {
$__t0 = 0;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Odd"))) {
$__t0 = 1;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
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
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
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
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("Even");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "CommutativeRing0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Int_commutativeRingParity'] ?? \PhpursThunks::eval('Data_Int_commutativeRingParity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Int_ceil'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Int_unsafeClamp'] ?? \PhpursThunks::eval('Data_Int_unsafeClamp'))))(($GLOBALS['Data_Number_ceil'] ?? \PhpursThunks::eval('Data_Number_ceil'))); return $v; };
\PhpursThunks::$thunks['Data_Int_boundedParity'] = function() { $v = (object)["bottom" => new Phpurs_Data0("Even"), "top" => new Phpurs_Data0("Odd"), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Int_ordParity'] ?? \PhpursThunks::eval('Data_Int_ordParity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Int_binary'] = function() { $v = 2; return $v; };
\PhpursThunks::$thunks['Data_Int_base36'] = function() { $v = 36; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






























