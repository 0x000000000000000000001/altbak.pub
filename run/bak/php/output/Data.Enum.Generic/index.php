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

// Data_Enum_Generic_lessThan
$GLOBALS['Data_Enum_Generic_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Enum_Generic_genericToEnum'
function majData_majEnum_majGeneric_genericmajTomajEnum__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajTomajEnum__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericToEnum__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericToEnum__prime__'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajTomajEnum__prime__';

// Data_Enum_Generic_genericToEnum
function majData_majEnum_majGeneric_genericmajTomajEnum($dictGeneric_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajTomajEnum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $to_1_0 = ($dictGeneric_0)->{'to'};
  $__res = function($dictGenericBoundedEnum_2) use ($to_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($to_1_0)))(($dictGenericBoundedEnum_2)->{'genericToEnum__prime__'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericToEnum'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajTomajEnum';

// Data_Enum_Generic_genericSucc'
function majData_majEnum_majGeneric_genericmajSucc__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajSucc__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericSucc__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericSucc__prime__'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajSucc__prime__';

// Data_Enum_Generic_genericSucc
function majData_majEnum_majGeneric_genericmajSucc($dictGeneric_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajSucc';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $to_1_0 = ($dictGeneric_0)->{'to'};
  $from_2_1 = ($dictGeneric_0)->{'from'};
  $__res = function($dictGenericEnum_3) use ($from_2_1, $to_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($to_1_0)))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictGenericEnum_3)->{'genericSucc__prime__'}))($from_2_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericSucc'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajSucc';

// Data_Enum_Generic_genericPred'
function majData_majEnum_majGeneric_genericmajPred__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajPred__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericPred__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericPred__prime__'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajPred__prime__';

// Data_Enum_Generic_genericPred
function majData_majEnum_majGeneric_genericmajPred($dictGeneric_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajPred';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $to_1_0 = ($dictGeneric_0)->{'to'};
  $from_2_1 = ($dictGeneric_0)->{'from'};
  $__res = function($dictGenericEnum_3) use ($from_2_1, $to_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($to_1_0)))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictGenericEnum_3)->{'genericPred__prime__'}))($from_2_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericPred'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajPred';

// Data_Enum_Generic_genericFromEnum'
function majData_majEnum_majGeneric_genericmajFrommajEnum__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajFrommajEnum__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericFromEnum__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericFromEnum__prime__'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajFrommajEnum__prime__';

// Data_Enum_Generic_genericFromEnum
function majData_majEnum_majGeneric_genericmajFrommajEnum($dictGeneric_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajFrommajEnum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $from_1_0 = ($dictGeneric_0)->{'from'};
  $__res = function($dictGenericBoundedEnum_2) use ($from_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictGenericBoundedEnum_2)->{'genericFromEnum__prime__'}))($from_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericFromEnum'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajFrommajEnum';

// Data_Enum_Generic_genericEnumSum
function majData_majEnum_majGeneric_genericmajEnummajSum($dictGenericEnum_0, $dictGenericTop_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajEnummajSum';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $genericTop_prime_2_0 = ($dictGenericTop_1)->{'genericTop__prime__'};
  $__res = (function() use ($dictGenericEnum_0, $genericTop_prime_2_0) {
  $__fn = function($dictGenericEnum1_3, $dictGenericBottom_4 = null) use ($dictGenericEnum_0, $genericTop_prime_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $genericBottom_prime_5_1 = ($dictGenericBottom_4)->{'genericBottom__prime__'};
  $__res = (object)["genericPred__prime__" => function($v_6) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop_prime_2_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_6 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inl) {
$__t2 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Inl']))((($dictGenericEnum_0)->{'genericPred__prime__'})(($v_6)->{'value0'}));
goto end_branch_2;;
};
  if ($v_6 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inr) {
$v1_7_3 = (($dictGenericEnum1_3)->{'genericPred__prime__'})(($v_6)->{'value0'});
$__t4 = null;;
if ($v1_7_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = new \Data\Maybe\Data_Maybe_Just(new \Data\Generic\Rep\Data_Generic_Rep_Inl($genericTop_prime_2_0));
goto end_branch_4;;
};
if ($v1_7_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = new \Data\Maybe\Data_Maybe_Just(new \Data\Generic\Rep\Data_Generic_Rep_Inr(($v1_7_3)->{'value0'}));
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
}, "genericSucc__prime__" => function($v_6) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericBottom_prime_5_1) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v_6 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inl) {
$v1_7_6 = (($dictGenericEnum_0)->{'genericSucc__prime__'})(($v_6)->{'value0'});
$__t7 = null;;
if ($v1_7_6 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = new \Data\Maybe\Data_Maybe_Just(new \Data\Generic\Rep\Data_Generic_Rep_Inr($genericBottom_prime_5_1));
goto end_branch_7;;
};
if ($v1_7_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = new \Data\Maybe\Data_Maybe_Just(new \Data\Generic\Rep\Data_Generic_Rep_Inl(($v1_7_6)->{'value0'}));
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t5 = $__t7;
goto end_branch_5;;
};
  if ($v_6 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inr) {
$__t5 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Inr']))((($dictGenericEnum1_3)->{'genericSucc__prime__'})(($v_6)->{'value0'}));
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericEnumSum'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajEnummajSum';

// Data_Enum_Generic_genericEnumProduct
function majData_majEnum_majGeneric_genericmajEnummajProduct($dictGenericEnum_0, $dictGenericTop_1 = null, $dictGenericBottom_2 = null, $dictGenericEnum1_3 = null, $dictGenericTop1_4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajEnummajProduct';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $genericTop_prime_5_0 = ($dictGenericTop1_4)->{'genericTop__prime__'};
  $__res = function($dictGenericBottom1_6) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop_prime_5_0) {
  $__num = \func_num_args();
  $genericBottom_prime_7_1 = ($dictGenericBottom1_6)->{'genericBottom__prime__'};
  $__res = (object)["genericPred__prime__" => function($v_8) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop_prime_5_0) {
  $__num = \func_num_args();
  $v1_9_2 = (($dictGenericEnum1_3)->{'genericPred__prime__'})(($v_8)->{'value1'});
  $__t3 = null;;
  if ($v1_9_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(new \Data\Generic\Rep\Data_Generic_Rep_Product(($v_8)->{'value0'}, ($v1_9_2)->{'value0'}));
goto end_branch_3;;
};
  if ($v1_9_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($a_10) use ($genericTop_prime_5_0) {
  $__num = \func_num_args();
  $__res = new \Data\Generic\Rep\Data_Generic_Rep_Product($a_10, $genericTop_prime_5_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictGenericEnum_0)->{'genericPred__prime__'})(($v_8)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_8) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericBottom_prime_7_1) {
  $__num = \func_num_args();
  $v1_9_4 = (($dictGenericEnum1_3)->{'genericSucc__prime__'})(($v_8)->{'value1'});
  $__t5 = null;;
  if ($v1_9_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = new \Data\Maybe\Data_Maybe_Just(new \Data\Generic\Rep\Data_Generic_Rep_Product(($v_8)->{'value0'}, ($v1_9_4)->{'value0'}));
goto end_branch_5;;
};
  if ($v1_9_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($a_10) use ($genericBottom_prime_7_1) {
  $__num = \func_num_args();
  $__res = new \Data\Generic\Rep\Data_Generic_Rep_Product($a_10, $genericBottom_prime_7_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictGenericEnum_0)->{'genericSucc__prime__'})(($v_8)->{'value0'}));
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
  return 5 < $__num ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericEnumProduct'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajEnummajProduct';

// Data_Enum_Generic_genericEnumNoArguments
$GLOBALS['Data_Enum_Generic_genericEnumNoArguments'] = (object)["genericPred__prime__" => function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_Generic_genericEnumConstructor
function majData_majEnum_majGeneric_genericmajEnummajConstructor($dictGenericEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajEnummajConstructor';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericPred__prime__" => function($v_1) use ($dictGenericEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Constructor']))((($dictGenericEnum_0)->{'genericPred__prime__'})($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_1) use ($dictGenericEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Constructor']))((($dictGenericEnum_0)->{'genericSucc__prime__'})($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericEnumConstructor'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajEnummajConstructor';

// Data_Enum_Generic_genericEnumArgument
function majData_majEnum_majGeneric_genericmajEnummajArgument($dictEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajEnummajArgument';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericPred__prime__" => function($v_1) use ($dictEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Argument']))((($dictEnum_0)->{'pred'})($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_1) use ($dictEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Argument']))((($dictEnum_0)->{'succ'})($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericEnumArgument'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajEnummajArgument';

// Data_Enum_Generic_genericCardinality'
function majData_majEnum_majGeneric_genericmajCardinality__prime__($dict_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajCardinality__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericCardinality__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericCardinality__prime__'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajCardinality__prime__';

// Data_Enum_Generic_genericCardinality
function majData_majEnum_majGeneric_genericmajCardinality($dictGeneric_0, $dictGenericBoundedEnum_1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajCardinality';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($dictGenericBoundedEnum_1)->{'genericCardinality__prime__'};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericCardinality'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajCardinality';

// Data_Enum_Generic_genericBoundedEnumSum
function majData_majEnum_majGeneric_genericmajBoundedmajEnummajSum($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajBoundedmajEnummajSum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $genericCardinality_prime1_1_0 = ($dictGenericBoundedEnum_0)->{'genericCardinality__prime__'};
  $__res = function($dictGenericBoundedEnum1_2) use ($dictGenericBoundedEnum_0, $genericCardinality_prime1_1_0) {
  $__num = \func_num_args();
  $__res = (object)["genericCardinality__prime__" => ($genericCardinality_prime1_1_0 + ($dictGenericBoundedEnum1_2)->{'genericCardinality__prime__'}), "genericToEnum__prime__" => function($n_3) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality_prime1_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Enum_Generic_greaterThanOrEq'])($n_3))(0)))((($GLOBALS['Data_Enum_Generic_lessThan'])($n_3))($genericCardinality_prime1_1_0))) {
$__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Inl']))((($dictGenericBoundedEnum_0)->{'genericToEnum__prime__'})($n_3));
goto end_branch_1;;
};
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Inr']))((($dictGenericBoundedEnum1_2)->{'genericToEnum__prime__'})(($n_3 - $genericCardinality_prime1_1_0)));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_3) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality_prime1_1_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_3 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inl) {
$__t2 = (($dictGenericBoundedEnum_0)->{'genericFromEnum__prime__'})(($v_3)->{'value0'});
goto end_branch_2;;
};
  if ($v_3 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inr) {
$__t2 = ((($dictGenericBoundedEnum1_2)->{'genericFromEnum__prime__'})(($v_3)->{'value0'}) + $genericCardinality_prime1_1_0);
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericBoundedEnumSum'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajBoundedmajEnummajSum';

// Data_Enum_Generic_genericBoundedEnumProduct
function majData_majEnum_majGeneric_genericmajBoundedmajEnummajProduct($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajBoundedmajEnummajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $genericCardinality_prime1_1_0 = ($dictGenericBoundedEnum_0)->{'genericCardinality__prime__'};
  $__res = function($dictGenericBoundedEnum1_2) use ($dictGenericBoundedEnum_0, $genericCardinality_prime1_1_0) {
  $__num = \func_num_args();
  $genericCardinality_prime2_3_1 = ($dictGenericBoundedEnum1_2)->{'genericCardinality__prime__'};
  $__res = (object)["genericCardinality__prime__" => ($genericCardinality_prime1_1_0 * $genericCardinality_prime2_3_1), "genericToEnum__prime__" => function($n_4) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality_prime2_3_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Product']))((($dictGenericBoundedEnum_0)->{'genericToEnum__prime__'})(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])->{'div'})($n_4))($genericCardinality_prime2_3_1)))))((($dictGenericBoundedEnum1_2)->{'genericToEnum__prime__'})(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])->{'mod'})($n_4))($genericCardinality_prime2_3_1)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v1_4) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality_prime2_3_1) {
  $__num = \func_num_args();
  $__res = (((($dictGenericBoundedEnum_0)->{'genericFromEnum__prime__'})(($v1_4)->{'value0'}) * $genericCardinality_prime2_3_1) + (($dictGenericBoundedEnum1_2)->{'genericFromEnum__prime__'})(($v1_4)->{'value1'}));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericBoundedEnumProduct'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajBoundedmajEnummajProduct';

// Data_Enum_Generic_genericBoundedEnumNoArguments
$GLOBALS['Data_Enum_Generic_genericBoundedEnumNoArguments'] = (object)["genericCardinality__prime__" => 1, "genericToEnum__prime__" => function($i_0) {
  $__num = \func_num_args();
  $__res = match ($i_0) { 0 => new \Data\Maybe\Data_Maybe_Just(new \Data\Generic\Rep\Data_Generic_Rep_NoArguments()), default => new \Data\Maybe\Data_Maybe_Nothing() };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_0) {
  $__num = \func_num_args();
  $__res = 0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Enum_Generic_genericBoundedEnumConstructor
function majData_majEnum_majGeneric_genericmajBoundedmajEnummajConstructor($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajBoundedmajEnummajConstructor';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericCardinality__prime__" => ($dictGenericBoundedEnum_0)->{'genericCardinality__prime__'}, "genericToEnum__prime__" => function($i_1) use ($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Constructor']))((($dictGenericBoundedEnum_0)->{'genericToEnum__prime__'})($i_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_1) use ($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = (($dictGenericBoundedEnum_0)->{'genericFromEnum__prime__'})($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericBoundedEnumConstructor'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajBoundedmajEnummajConstructor';

// Data_Enum_Generic_genericBoundedEnumArgument
function majData_majEnum_majGeneric_genericmajBoundedmajEnummajArgument($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEnum_majGeneric_genericmajBoundedmajEnummajArgument';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericCardinality__prime__" => ($dictBoundedEnum_0)->{'cardinality'}, "genericToEnum__prime__" => function($i_1) use ($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($GLOBALS['Data_Generic_Rep_Argument']))((($dictBoundedEnum_0)->{'toEnum'})($i_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_1) use ($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = (($dictBoundedEnum_0)->{'fromEnum'})($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Enum_Generic_genericBoundedEnumArgument'] = __NAMESPACE__ . '\\majData_majEnum_majGeneric_genericmajBoundedmajEnummajArgument';

