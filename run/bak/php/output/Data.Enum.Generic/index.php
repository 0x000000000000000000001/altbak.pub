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
\PhpursThunks::$thunks['Data_Enum_Generic_genericToEnum__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericToEnum__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericToEnum'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericBoundedEnum_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v1_2 = null) use ($dictGeneric_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Just"))) {
$__t0 = new Phpurs_Data1("Just", (($dictGeneric_0)->{'to'})(($v1_2)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictGenericBoundedEnum_1)->{'genericToEnum__prime__'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericSucc__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericSucc__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericSucc'] = function() { $v = function($dictGeneric_0 = null) {
  $__num = \func_num_args();
  $from_1_0 = ($dictGeneric_0)->{'from'};
  $__res = function($dictGenericEnum_2 = null) use ($dictGeneric_0, $from_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v1_3 = null) use ($dictGeneric_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($dictGeneric_0)->{'to'})(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($dictGenericEnum_2)->{'genericSucc__prime__'}))($from_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericPred__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericPred__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericPred'] = function() { $v = function($dictGeneric_0 = null) {
  $__num = \func_num_args();
  $from_1_0 = ($dictGeneric_0)->{'from'};
  $__res = function($dictGenericEnum_2 = null) use ($dictGeneric_0, $from_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v1_3 = null) use ($dictGeneric_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($dictGeneric_0)->{'to'})(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($dictGenericEnum_2)->{'genericPred__prime__'}))($from_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericFromEnum__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericFromEnum__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericFromEnum'] = function() { $v = function($dictGeneric_0 = null) {
  $__num = \func_num_args();
  $from_1_0 = ($dictGeneric_0)->{'from'};
  $__res = function($dictGenericBoundedEnum_2 = null) use ($from_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($dictGenericBoundedEnum_2)->{'genericFromEnum__prime__'}))($from_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumSum'] = function() { $v = (function() {
  $__fn = function($dictGenericEnum_0 = null, $dictGenericTop_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $genericTop__prime___2_0 = ($dictGenericTop_1)->{'genericTop__prime__'};
  $__res = (function() use ($dictGenericEnum_0, $genericTop__prime___2_0) {
  $__fn = function($dictGenericEnum1_3 = null, $dictGenericBottom_4 = null) use ($dictGenericEnum_0, $genericTop__prime___2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $genericBottom__prime___5_1 = ($dictGenericBottom_4)->{'genericBottom__prime__'};
  $__res = (object)["genericPred__prime__" => function($v_6 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop__prime___2_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Inl"))) {
$__local_var_7_3 = (($dictGenericEnum_0)->{'genericPred__prime__'})(($v_6)->{'value0'});
$__t4 = null;;
if ((is_object($__local_var_7_3) && (($__local_var_7_3)->{'tag'} === "Just"))) {
$__t4 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", ($__local_var_7_3)->{'value0'}));
goto end_branch_4;;
};
$__t4 = new Phpurs_Data0("Nothing");
end_branch_4:;
$__t2 = $__t4;
goto end_branch_2;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Inr"))) {
$v1_7_5 = (($dictGenericEnum1_3)->{'genericPred__prime__'})(($v_6)->{'value0'});
$__t6 = null;;
if ((is_object($v1_7_5) && (($v1_7_5)->{'tag'} === "Nothing"))) {
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", $genericTop__prime___2_0));
goto end_branch_6;;
};
if ((is_object($v1_7_5) && (($v1_7_5)->{'tag'} === "Just"))) {
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", ($v1_7_5)->{'value0'}));
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t2 = $__t6;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_6 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericBottom__prime___5_1) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Inl"))) {
$v1_7_8 = (($dictGenericEnum_0)->{'genericSucc__prime__'})(($v_6)->{'value0'});
$__t9 = null;;
if ((is_object($v1_7_8) && (($v1_7_8)->{'tag'} === "Nothing"))) {
$__t9 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", $genericBottom__prime___5_1));
goto end_branch_9;;
};
if ((is_object($v1_7_8) && (($v1_7_8)->{'tag'} === "Just"))) {
$__t9 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", ($v1_7_8)->{'value0'}));
goto end_branch_9;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
end_branch_9:;
$__t7 = $__t9;
goto end_branch_7;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Inr"))) {
$__local_var_7_10 = (($dictGenericEnum1_3)->{'genericSucc__prime__'})(($v_6)->{'value0'});
$__t11 = null;;
if ((is_object($__local_var_7_10) && (($__local_var_7_10)->{'tag'} === "Just"))) {
$__t11 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", ($__local_var_7_10)->{'value0'}));
goto end_branch_11;;
};
$__t11 = new Phpurs_Data0("Nothing");
end_branch_11:;
$__t7 = $__t11;
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumProduct'] = function() { $v = (function() {
  $__fn = function($dictGenericEnum_0 = null, $dictGenericTop_1 = null, $dictGenericBottom_2 = null, $dictGenericEnum1_3 = null, $dictGenericTop1_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $genericTop__prime___5_0 = ($dictGenericTop1_4)->{'genericTop__prime__'};
  $__res = function($dictGenericBottom1_6 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop__prime___5_0) {
  $__num = \func_num_args();
  $genericBottom__prime___7_1 = ($dictGenericBottom1_6)->{'genericBottom__prime__'};
  $__res = (object)["genericPred__prime__" => function($v_8 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop__prime___5_0) {
  $__num = \func_num_args();
  $v1_9_2 = (($dictGenericEnum1_3)->{'genericPred__prime__'})(($v_8)->{'value1'});
  $__t3 = null;;
  if ((is_object($v1_9_2) && (($v1_9_2)->{'tag'} === "Just"))) {
$__t3 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($v_8)->{'value0'}, ($v1_9_2)->{'value0'}));
goto end_branch_3;;
};
  if ((is_object($v1_9_2) && (($v1_9_2)->{'tag'} === "Nothing"))) {
$__local_var_10_4 = (($dictGenericEnum_0)->{'genericPred__prime__'})(($v_8)->{'value0'});
$__t5 = null;;
if ((is_object($__local_var_10_4) && (($__local_var_10_4)->{'tag'} === "Just"))) {
$__t5 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($__local_var_10_4)->{'value0'}, $genericTop__prime___5_0));
goto end_branch_5;;
};
$__t5 = new Phpurs_Data0("Nothing");
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_8 = null) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericBottom__prime___7_1) {
  $__num = \func_num_args();
  $v1_9_6 = (($dictGenericEnum1_3)->{'genericSucc__prime__'})(($v_8)->{'value1'});
  $__t7 = null;;
  if ((is_object($v1_9_6) && (($v1_9_6)->{'tag'} === "Just"))) {
$__t7 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($v_8)->{'value0'}, ($v1_9_6)->{'value0'}));
goto end_branch_7;;
};
  if ((is_object($v1_9_6) && (($v1_9_6)->{'tag'} === "Nothing"))) {
$__local_var_10_8 = (($dictGenericEnum_0)->{'genericSucc__prime__'})(($v_8)->{'value0'});
$__t9 = null;;
if ((is_object($__local_var_10_8) && (($__local_var_10_8)->{'tag'} === "Just"))) {
$__t9 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($__local_var_10_8)->{'value0'}, $genericBottom__prime___7_1));
goto end_branch_9;;
};
$__t9 = new Phpurs_Data0("Nothing");
end_branch_9:;
$__t7 = $__t9;
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumNoArguments'] = function() { $v = (object)["genericPred__prime__" => function($v_0 = null) {
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
}]; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumConstructor'] = function() { $v = function($dictGenericEnum_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["genericPred__prime__" => function($v_1 = null) use ($dictGenericEnum_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (($dictGenericEnum_0)->{'genericPred__prime__'})($v_1);
  $__t1 = null;;
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_2_0)->{'value0'});
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_1 = null) use ($dictGenericEnum_0) {
  $__num = \func_num_args();
  $__local_var_2_2 = (($dictGenericEnum_0)->{'genericSucc__prime__'})($v_1);
  $__t3 = null;;
  if ((is_object($__local_var_2_2) && (($__local_var_2_2)->{'tag'} === "Just"))) {
$__t3 = new Phpurs_Data1("Just", ($__local_var_2_2)->{'value0'});
goto end_branch_3;;
};
  $__t3 = new Phpurs_Data0("Nothing");
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumArgument'] = function() { $v = function($dictEnum_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["genericPred__prime__" => function($v_1 = null) use ($dictEnum_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (($dictEnum_0)->{'pred'})($v_1);
  $__t1 = null;;
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_2_0)->{'value0'});
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_1 = null) use ($dictEnum_0) {
  $__num = \func_num_args();
  $__local_var_2_2 = (($dictEnum_0)->{'succ'})($v_1);
  $__t3 = null;;
  if ((is_object($__local_var_2_2) && (($__local_var_2_2)->{'tag'} === "Just"))) {
$__t3 = new Phpurs_Data1("Just", ($__local_var_2_2)->{'value0'});
goto end_branch_3;;
};
  $__t3 = new Phpurs_Data0("Nothing");
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericCardinality__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericCardinality__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericCardinality'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericBoundedEnum_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($dictGenericBoundedEnum_1)->{'genericCardinality__prime__'};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumSum'] = function() { $v = function($dictGenericBoundedEnum_0 = null) {
  $__num = \func_num_args();
  $genericCardinality__prime__1_1_0 = ($dictGenericBoundedEnum_0)->{'genericCardinality__prime__'};
  $__res = function($dictGenericBoundedEnum1_2 = null) use ($dictGenericBoundedEnum_0, $genericCardinality__prime__1_1_0) {
  $__num = \func_num_args();
  $__res = (object)["genericCardinality__prime__" => ($genericCardinality__prime__1_1_0 + ($dictGenericBoundedEnum1_2)->{'genericCardinality__prime__'}), "genericToEnum__prime__" => function($n_3 = null) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality__prime__1_1_0) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((($n_3 >= 0) && ($n_3 < $genericCardinality__prime__1_1_0))) {
$__local_var_4_4 = (($dictGenericBoundedEnum_0)->{'genericToEnum__prime__'})($n_3);
$__t5 = null;;
if ((is_object($__local_var_4_4) && (($__local_var_4_4)->{'tag'} === "Just"))) {
$__t5 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", ($__local_var_4_4)->{'value0'}));
goto end_branch_5;;
};
$__t5 = new Phpurs_Data0("Nothing");
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
  $__local_var_4_1 = (($dictGenericBoundedEnum1_2)->{'genericToEnum__prime__'})(($n_3 - $genericCardinality__prime__1_1_0));
  $__t2 = null;;
  if ((is_object($__local_var_4_1) && (($__local_var_4_1)->{'tag'} === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", ($__local_var_4_1)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = new Phpurs_Data0("Nothing");
  end_branch_2:;
  $__t3 = $__t2;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_3 = null) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality__prime__1_1_0) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Inl"))) {
$__t6 = (($dictGenericBoundedEnum_0)->{'genericFromEnum__prime__'})(($v_3)->{'value0'});
goto end_branch_6;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Inr"))) {
$__t6 = ((($dictGenericBoundedEnum1_2)->{'genericFromEnum__prime__'})(($v_3)->{'value0'}) + $genericCardinality__prime__1_1_0);
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
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
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumProduct'] = function() { $v = function($dictGenericBoundedEnum_0 = null) {
  $__num = \func_num_args();
  $genericCardinality__prime__1_1_0 = ($dictGenericBoundedEnum_0)->{'genericCardinality__prime__'};
  $__res = function($dictGenericBoundedEnum1_2 = null) use ($dictGenericBoundedEnum_0, $genericCardinality__prime__1_1_0) {
  $__num = \func_num_args();
  $genericCardinality__prime__2_3_1 = ($dictGenericBoundedEnum1_2)->{'genericCardinality__prime__'};
  $__res = (object)["genericCardinality__prime__" => ($genericCardinality__prime__1_1_0 * $genericCardinality__prime__2_3_1), "genericToEnum__prime__" => function($n_4 = null) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality__prime__2_3_1) {
  $__num = \func_num_args();
  $__local_var_5_2 = (($dictGenericBoundedEnum_0)->{'genericToEnum__prime__'})(($n_4 / $genericCardinality__prime__2_3_1));
  $__t3 = null;;
  if ((is_object($__local_var_5_2) && (($__local_var_5_2)->{'tag'} === "Just"))) {
$__local_var_6_4 = (($dictGenericBoundedEnum1_2)->{'genericToEnum__prime__'})(((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))($n_4))($genericCardinality__prime__2_3_1));
$__t5 = null;;
if ((is_object($__local_var_6_4) && (($__local_var_6_4)->{'tag'} === "Just"))) {
$__t5 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($__local_var_5_2)->{'value0'}, ($__local_var_6_4)->{'value0'}));
goto end_branch_5;;
};
$__t5 = new Phpurs_Data0("Nothing");
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
  $__t3 = new Phpurs_Data0("Nothing");
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v1_4 = null) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality__prime__2_3_1) {
  $__num = \func_num_args();
  $__res = (((($dictGenericBoundedEnum_0)->{'genericFromEnum__prime__'})(($v1_4)->{'value0'}) * $genericCardinality__prime__2_3_1) + (($dictGenericBoundedEnum1_2)->{'genericFromEnum__prime__'})(($v1_4)->{'value1'}));
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
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumNoArguments'] = function() { $v = (object)["genericCardinality__prime__" => 1, "genericToEnum__prime__" => function($i_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  switch ($i_0) {
case 0:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("NoArguments"));
goto end_branch_0;;
break;
default:
;
break;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = 0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumConstructor'] = function() { $v = function($dictGenericBoundedEnum_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["genericCardinality__prime__" => ($dictGenericBoundedEnum_0)->{'genericCardinality__prime__'}, "genericToEnum__prime__" => function($i_1 = null) use ($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (($dictGenericBoundedEnum_0)->{'genericToEnum__prime__'})($i_1);
  $__t1 = null;;
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_2_0)->{'value0'});
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_1 = null) use ($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = (($dictGenericBoundedEnum_0)->{'genericFromEnum__prime__'})($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumArgument'] = function() { $v = function($dictBoundedEnum_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["genericCardinality__prime__" => ($dictBoundedEnum_0)->{'cardinality'}, "genericToEnum__prime__" => function($i_1 = null) use ($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (($dictBoundedEnum_0)->{'toEnum'})($i_1);
  $__t1 = null;;
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_2_0)->{'value0'});
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_1 = null) use ($dictBoundedEnum_0) {
  $__num = \func_num_args();
  $__res = (($dictBoundedEnum_0)->{'fromEnum'})($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






















