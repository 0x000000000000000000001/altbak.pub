<?php

namespace Data\Enum\Generic;

// ALL IMPORTS: Control.Apply, Control.Semigroupoid, Data.Boolean, Data.Bounded.Generic, Data.Enum, Data.Enum.Generic, Data.Eq, Data.EuclideanRing, Data.Function, Data.Functor, Data.Generic.Rep, Data.HeytingAlgebra, Data.Maybe, Data.Newtype, Data.Ord, Data.Ring, Data.Semiring, Prelude, Prim
// TO REQUIRE: Control.Apply, Control.Semigroupoid, Data.Boolean, Data.Bounded.Generic, Data.Enum, Data.Enum.Generic, Data.Eq, Data.EuclideanRing, Data.Function, Data.Functor, Data.Generic.Rep, Data.HeytingAlgebra, Data.Maybe, Data.Newtype, Data.Ord, Data.Ring, Data.Semiring, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Bounded.Generic/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Enum.Generic/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
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


// Data_Enum_Generic_greaterThanOrEq
$GLOBALS['Data_Enum_Generic_greaterThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT"));
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1 = null, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (is_object((($__local_var_0_0)($a1_1))($a2_2)) && (((($__local_var_0_0)($a1_1))($a2_2))->{'tag'} === "LT")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Enum_Generic_lessThan
$GLOBALS['Data_Enum_Generic_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT"));
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1 = null, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (is_object((($__local_var_0_0)($a1_1))($a2_2)) && (((($__local_var_0_0)($a1_1))($a2_2))->{'tag'} === "LT"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Enum_Generic_genericToEnum'
$GLOBALS['Data_Enum_Generic_genericToEnum__prime__'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['genericToEnum__prime__'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericToEnum
$GLOBALS['Data_Enum_Generic_genericToEnum'] = function($dictGeneric_0 = null) {
  $__num = \func_num_args();
  $to_1_0 = ($dictGeneric_0)['to'];
  $__res = function($dictGenericBoundedEnum_2 = null) use ($to_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($to_1_0)))(($dictGenericBoundedEnum_2)['genericToEnum__prime__']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericSucc'
$GLOBALS['Data_Enum_Generic_genericSucc__prime__'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['genericSucc__prime__'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericSucc
$GLOBALS['Data_Enum_Generic_genericSucc'] = function($dictGeneric_0 = null) {
  $__num = \func_num_args();
  $to_1_0 = ($dictGeneric_0)['to'];
  $from_2_1 = ($dictGeneric_0)['from'];
  $__res = function($dictGenericEnum_3 = null) use ($from_2_1, $to_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($to_1_0)))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictGenericEnum_3)['genericSucc__prime__']))($from_2_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericPred'
$GLOBALS['Data_Enum_Generic_genericPred__prime__'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['genericPred__prime__'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericPred
$GLOBALS['Data_Enum_Generic_genericPred'] = function($dictGeneric_0 = null) {
  $__num = \func_num_args();
  $to_1_0 = ($dictGeneric_0)['to'];
  $from_2_1 = ($dictGeneric_0)['from'];
  $__res = function($dictGenericEnum_3 = null) use ($from_2_1, $to_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($to_1_0)))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictGenericEnum_3)['genericPred__prime__']))($from_2_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericFromEnum'
$GLOBALS['Data_Enum_Generic_genericFromEnum__prime__'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['genericFromEnum__prime__'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericFromEnum
$GLOBALS['Data_Enum_Generic_genericFromEnum'] = function($dictGeneric_0 = null) {
  $__num = \func_num_args();
  $from_1_0 = ($dictGeneric_0)['from'];
  $__res = function($dictGenericBoundedEnum_2 = null) use ($from_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictGenericBoundedEnum_2)['genericFromEnum__prime__']))($from_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericEnumSum
$GLOBALS['Data_Enum_Generic_genericEnumSum'] = (function() {
  $__fn = function($dictGenericEnum_0 = null, $dictGenericTop_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $genericTop_prime_2_0 = ($dictGenericTop_1)['genericTop__prime__'];
  $__res = (function() use ($dictGenericEnum_0, $genericTop_prime_2_0) {
  $__fn = function($dictGenericEnum1_3 = null, $dictGenericBottom_4 = null) use ($dictGenericEnum_0, $genericTop_prime_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $genericBottom_prime_5_1 = ($dictGenericBottom_4)['genericBottom__prime__'];
  $__res = ["genericPred__prime__" => function($v_6 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop_prime_2_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Inl"))) {
$__t2 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Inl']))((($dictGenericEnum_0)['genericPred__prime__'])(($v_6)->{'value0'}));
goto end_branch_2;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Inr"))) {
$v1_7_3 = (($dictGenericEnum1_3)['genericPred__prime__'])(($v_6)->{'value0'});
$__t4 = null;;
if ((is_object($v1_7_3) && (($v1_7_3)->{'tag'} === "Nothing"))) {
$__t4 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", $genericTop_prime_2_0));
goto end_branch_4;;
};
if ((is_object($v1_7_3) && (($v1_7_3)->{'tag'} === "Just"))) {
$__t4 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", ($v1_7_3)->{'value0'}));
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t2 = $__t4;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_6 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericBottom_prime_5_1) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Inl"))) {
$v1_7_6 = (($dictGenericEnum_0)['genericSucc__prime__'])(($v_6)->{'value0'});
$__t7 = null;;
if ((is_object($v1_7_6) && (($v1_7_6)->{'tag'} === "Nothing"))) {
$__t7 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", $genericBottom_prime_5_1));
goto end_branch_7;;
};
if ((is_object($v1_7_6) && (($v1_7_6)->{'tag'} === "Just"))) {
$__t7 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", ($v1_7_6)->{'value0'}));
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t5 = $__t7;
goto end_branch_5;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Inr"))) {
$__t5 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Inr']))((($dictGenericEnum1_3)['genericSucc__prime__'])(($v_6)->{'value0'}));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
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
})();

// Data_Enum_Generic_genericEnumProduct
$GLOBALS['Data_Enum_Generic_genericEnumProduct'] = (function() {
  $__fn = function($dictGenericEnum_0 = null, $dictGenericTop_1 = null, $dictGenericBottom_2 = null, $dictGenericEnum1_3 = null, $dictGenericTop1_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $genericTop_prime_5_0 = ($dictGenericTop1_4)['genericTop__prime__'];
  $__res = function($dictGenericBottom1_6 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop_prime_5_0) {
  $__num = \func_num_args();
  $genericBottom_prime_7_1 = ($dictGenericBottom1_6)['genericBottom__prime__'];
  $__res = ["genericPred__prime__" => function($v_8 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop_prime_5_0) {
  $__num = \func_num_args();
  $v1_9_2 = (($dictGenericEnum1_3)['genericPred__prime__'])(($v_8)->{'value1'});
  $__t3 = null;;
  if ((is_object($v1_9_2) && (($v1_9_2)->{'tag'} === "Just"))) {
$__t3 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($v_8)->{'value0'}, ($v1_9_2)->{'value0'}));
goto end_branch_3;;
};
  if ((is_object($v1_9_2) && (($v1_9_2)->{'tag'} === "Nothing"))) {
$__t3 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($a_10 = null) use ($genericTop_prime_5_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Product", $a_10, $genericTop_prime_5_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictGenericEnum_0)['genericPred__prime__'])(($v_8)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_8 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericBottom_prime_7_1) {
  $__num = \func_num_args();
  $v1_9_4 = (($dictGenericEnum1_3)['genericSucc__prime__'])(($v_8)->{'value1'});
  $__t5 = null;;
  if ((is_object($v1_9_4) && (($v1_9_4)->{'tag'} === "Just"))) {
$__t5 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($v_8)->{'value0'}, ($v1_9_4)->{'value0'}));
goto end_branch_5;;
};
  if ((is_object($v1_9_4) && (($v1_9_4)->{'tag'} === "Nothing"))) {
$__t5 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($a_10 = null) use ($genericBottom_prime_7_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Product", $a_10, $genericBottom_prime_7_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictGenericEnum_0)['genericSucc__prime__'])(($v_8)->{'value0'}));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 5 ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
  };
  return $__fn;
})();

// Data_Enum_Generic_genericEnumNoArguments
$GLOBALS['Data_Enum_Generic_genericEnumNoArguments'] = ["genericPred__prime__" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_Generic_genericEnumConstructor
$GLOBALS['Data_Enum_Generic_genericEnumConstructor'] = function($dictGenericEnum_0 = null) {
  $__num = \func_num_args();
  $__res = ["genericPred__prime__" => function($v_1 = null) use ($dictGenericEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Constructor']))((($dictGenericEnum_0)['genericPred__prime__'])($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_1 = null) use ($dictGenericEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Constructor']))((($dictGenericEnum_0)['genericSucc__prime__'])($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericEnumArgument
$GLOBALS['Data_Enum_Generic_genericEnumArgument'] = function($dictEnum_0 = null) {
  $__num = \func_num_args();
  $__res = ["genericPred__prime__" => function($v_1 = null) use ($dictEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Argument']))((($dictEnum_0)['pred'])($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_1 = null) use ($dictEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Argument']))((($dictEnum_0)['succ'])($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericCardinality'
$GLOBALS['Data_Enum_Generic_genericCardinality__prime__'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['genericCardinality__prime__'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericCardinality
$GLOBALS['Data_Enum_Generic_genericCardinality'] = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericBoundedEnum_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($dictGenericBoundedEnum_1)['genericCardinality__prime__'];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Enum_Generic_genericBoundedEnumSum
$GLOBALS['Data_Enum_Generic_genericBoundedEnumSum'] = function($dictGenericBoundedEnum_0 = null) {
  $__num = \func_num_args();
  $genericCardinality_prime1_1_0 = ($dictGenericBoundedEnum_0)['genericCardinality__prime__'];
  $__res = function($dictGenericBoundedEnum1_2 = null) use ($dictGenericBoundedEnum_0, $genericCardinality_prime1_1_0) {
  $__num = \func_num_args();
  $__res = ["genericCardinality__prime__" => ($genericCardinality_prime1_1_0 + ($dictGenericBoundedEnum1_2)['genericCardinality__prime__']), "genericToEnum__prime__" => function($n_3 = null) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality_prime1_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($GLOBALS['Data_Enum_Generic_greaterThanOrEq'])($n_3))(0)))((($GLOBALS['Data_Enum_Generic_lessThan'])($n_3))($genericCardinality_prime1_1_0))) {
$__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Inl']))((($dictGenericBoundedEnum_0)['genericToEnum__prime__'])($n_3));
goto end_branch_1;;
};
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Inr']))((($dictGenericBoundedEnum1_2)['genericToEnum__prime__'])(($n_3 - $genericCardinality_prime1_1_0)));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_3 = null) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality_prime1_1_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Inl"))) {
$__t2 = (($dictGenericBoundedEnum_0)['genericFromEnum__prime__'])(($v_3)->{'value0'});
goto end_branch_2;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Inr"))) {
$__t2 = ((($dictGenericBoundedEnum1_2)['genericFromEnum__prime__'])(($v_3)->{'value0'}) + $genericCardinality_prime1_1_0);
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericBoundedEnumProduct
$GLOBALS['Data_Enum_Generic_genericBoundedEnumProduct'] = function($dictGenericBoundedEnum_0 = null) {
  $__num = \func_num_args();
  $genericCardinality_prime1_1_0 = ($dictGenericBoundedEnum_0)['genericCardinality__prime__'];
  $__res = function($dictGenericBoundedEnum1_2 = null) use ($dictGenericBoundedEnum_0, $genericCardinality_prime1_1_0) {
  $__num = \func_num_args();
  $genericCardinality_prime2_3_1 = ($dictGenericBoundedEnum1_2)['genericCardinality__prime__'];
  $__res = ["genericCardinality__prime__" => ($genericCardinality_prime1_1_0 * $genericCardinality_prime2_3_1), "genericToEnum__prime__" => function($n_4 = null) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality_prime2_3_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_applyMaybe'])['apply'])(((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Product']))((($dictGenericBoundedEnum_0)['genericToEnum__prime__'])(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])['div'])($n_4))($genericCardinality_prime2_3_1)))))((($dictGenericBoundedEnum1_2)['genericToEnum__prime__'])(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])['mod'])($n_4))($genericCardinality_prime2_3_1)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v1_4 = null) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality_prime2_3_1) {
  $__num = \func_num_args();
  $__res = (((($dictGenericBoundedEnum_0)['genericFromEnum__prime__'])(($v1_4)->{'value0'}) * $genericCardinality_prime2_3_1) + (($dictGenericBoundedEnum1_2)['genericFromEnum__prime__'])(($v1_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericBoundedEnumNoArguments
$GLOBALS['Data_Enum_Generic_genericBoundedEnumNoArguments'] = ["genericCardinality__prime__" => 1, "genericToEnum__prime__" => function($i_0 = null) {
  $__num = \func_num_args();
  $__res = match ($i_0) { 0 => new Phpurs_Data1("Just", new Phpurs_Data0("NoArguments")), default => new Phpurs_Data0("Nothing") };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = 0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_Generic_genericBoundedEnumConstructor
$GLOBALS['Data_Enum_Generic_genericBoundedEnumConstructor'] = function($dictGenericBoundedEnum_0 = null) {
  $__num = \func_num_args();
  $__res = ["genericCardinality__prime__" => ($dictGenericBoundedEnum_0)['genericCardinality__prime__'], "genericToEnum__prime__" => function($i_1 = null) use ($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Constructor']))((($dictGenericBoundedEnum_0)['genericToEnum__prime__'])($i_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_1 = null) use ($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = (($dictGenericBoundedEnum_0)['genericFromEnum__prime__'])($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Enum_Generic_genericBoundedEnumArgument
$GLOBALS['Data_Enum_Generic_genericBoundedEnumArgument'] = function($dictBoundedEnum_0 = null) {
  $__num = \func_num_args();
  $__res = ["genericCardinality__prime__" => ($dictBoundedEnum_0)['cardinality'], "genericToEnum__prime__" => function($i_1 = null) use ($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])($GLOBALS['Data_Generic_Rep_Argument']))((($dictBoundedEnum_0)['toEnum'])($i_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_1 = null) use ($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = (($dictBoundedEnum_0)['fromEnum'])($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

