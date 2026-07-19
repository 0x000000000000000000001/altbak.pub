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
\PhpursThunks::$thunks['Data_Enum_Generic_genericToEnum__prime__'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericToEnum'"), recVars=[];
  $__res = ($dict_0)->genericToEnum__prime__;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericToEnum'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0, $dictGenericBoundedEnum_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericToEnum"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v1_2) use ($dictGeneric_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_2) && (($v1_2)->tag === "Just"))) {
$__t0 = new Phpurs_Data1("Just", (($dictGeneric_0)->to)(($v1_2)->value0));
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictGenericBoundedEnum_1)->genericToEnum__prime__);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericSucc__prime__'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericSucc'"), recVars=[];
  $__res = ($dict_0)->genericSucc__prime__;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericSucc'] = function() { $v = function($dictGeneric_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericSucc"), recVars=[];
  $from_1_0 = ($dictGeneric_0)->from;
  $__res = function($dictGenericEnum_2) use ($dictGeneric_0, $from_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v1_3) use ($dictGeneric_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($dictGeneric_0)->to)(($v1_3)->value0));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($dictGenericEnum_2)->genericSucc__prime__))($from_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericPred__prime__'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericPred'"), recVars=[];
  $__res = ($dict_0)->genericPred__prime__;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericPred'] = function() { $v = function($dictGeneric_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericPred"), recVars=[];
  $from_1_0 = ($dictGeneric_0)->from;
  $__res = function($dictGenericEnum_2) use ($dictGeneric_0, $from_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v1_3) use ($dictGeneric_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($dictGeneric_0)->to)(($v1_3)->value0));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($dictGenericEnum_2)->genericPred__prime__))($from_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericFromEnum__prime__'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericFromEnum'"), recVars=[];
  $__res = ($dict_0)->genericFromEnum__prime__;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericFromEnum'] = function() { $v = function($dictGeneric_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericFromEnum"), recVars=[];
  $from_1_0 = ($dictGeneric_0)->from;
  $__res = function($dictGenericBoundedEnum_2) use ($from_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($dictGenericBoundedEnum_2)->genericFromEnum__prime__))($from_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumSum'] = function() { $v = (function() {
  $__fn = function($dictGenericEnum_0, $dictGenericTop_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericEnumSum"), recVars=[];
  $genericTop__prime___2_0 = ($dictGenericTop_1)->genericTop__prime__;
  $__res = (function() use ($dictGenericEnum_0, $genericTop__prime___2_0) {
  $__fn = function($dictGenericEnum1_3, $dictGenericBottom_4 = null) use ($dictGenericEnum_0, $genericTop__prime___2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $genericBottom__prime___5_1 = ($dictGenericBottom_4)->genericBottom__prime__;
  $__res = (object)["genericPred__prime__" => function($v_6) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop__prime___2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_6) && (($v_6)->tag === "Inl"))) {
$__local_var_7_3 = (($dictGenericEnum_0)->genericPred__prime__)(($v_6)->value0);
if ((is_object($__local_var_7_3) && (($__local_var_7_3)->tag === "Just"))) {
$__t4 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", ($__local_var_7_3)->value0));
} else {
$__t4 = new Phpurs_Data0("Nothing");
};
$__t2 = $__t4;
} else {
if ((is_object($v_6) && (($v_6)->tag === "Inr"))) {
$v1_7_5 = (($dictGenericEnum1_3)->genericPred__prime__)(($v_6)->value0);
if ((is_object($v1_7_5) && (($v1_7_5)->tag === "Nothing"))) {
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", $genericTop__prime___2_0));
} else {
if ((is_object($v1_7_5) && (($v1_7_5)->tag === "Just"))) {
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", ($v1_7_5)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__t2 = $__t6;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_6) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericBottom__prime___5_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_6) && (($v_6)->tag === "Inl"))) {
$v1_7_8 = (($dictGenericEnum_0)->genericSucc__prime__)(($v_6)->value0);
if ((is_object($v1_7_8) && (($v1_7_8)->tag === "Nothing"))) {
$__t9 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", $genericBottom__prime___5_1));
} else {
if ((is_object($v1_7_8) && (($v1_7_8)->tag === "Just"))) {
$__t9 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", ($v1_7_8)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
};
};
$__t7 = $__t9;
} else {
if ((is_object($v_6) && (($v_6)->tag === "Inr"))) {
$__local_var_7_10 = (($dictGenericEnum1_3)->genericSucc__prime__)(($v_6)->value0);
if ((is_object($__local_var_7_10) && (($__local_var_7_10)->tag === "Just"))) {
$__t11 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", ($__local_var_7_10)->value0));
} else {
$__t11 = new Phpurs_Data0("Nothing");
};
$__t7 = $__t11;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
};
};
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
  $__fn = function($dictGenericEnum_0, $dictGenericTop_1 = null, $dictGenericBottom_2 = null, $dictGenericEnum1_3 = null, $dictGenericTop1_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericEnumProduct"), recVars=[];
  $genericTop__prime___5_0 = ($dictGenericTop1_4)->genericTop__prime__;
  $__res = function($dictGenericBottom1_6) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop__prime___5_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $genericBottom__prime___7_1 = ($dictGenericBottom1_6)->genericBottom__prime__;
  $__res = (object)["genericPred__prime__" => function($v_8) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericTop__prime___5_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v1_9_2 = (($dictGenericEnum1_3)->genericPred__prime__)(($v_8)->value1);
  if ((is_object($v1_9_2) && (($v1_9_2)->tag === "Just"))) {
$__t3 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($v_8)->value0, ($v1_9_2)->value0));
} else {
if ((is_object($v1_9_2) && (($v1_9_2)->tag === "Nothing"))) {
$__local_var_10_4 = (($dictGenericEnum_0)->genericPred__prime__)(($v_8)->value0);
if ((is_object($__local_var_10_4) && (($__local_var_10_4)->tag === "Just"))) {
$__t5 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($__local_var_10_4)->value0, $genericTop__prime___5_0));
} else {
$__t5 = new Phpurs_Data0("Nothing");
};
$__t3 = $__t5;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_8) use ($dictGenericEnum1_3, $dictGenericEnum_0, $genericBottom__prime___7_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v1_9_6 = (($dictGenericEnum1_3)->genericSucc__prime__)(($v_8)->value1);
  if ((is_object($v1_9_6) && (($v1_9_6)->tag === "Just"))) {
$__t7 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($v_8)->value0, ($v1_9_6)->value0));
} else {
if ((is_object($v1_9_6) && (($v1_9_6)->tag === "Nothing"))) {
$__local_var_10_8 = (($dictGenericEnum_0)->genericSucc__prime__)(($v_8)->value0);
if ((is_object($__local_var_10_8) && (($__local_var_10_8)->tag === "Just"))) {
$__t9 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($__local_var_10_8)->value0, $genericBottom__prime___7_1));
} else {
$__t9 = new Phpurs_Data0("Nothing");
};
$__t7 = $__t9;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
};
};
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
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumNoArguments'] = function() { $v = (object)["genericPred__prime__" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumConstructor'] = function() { $v = function($dictGenericEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericEnumConstructor"), recVars=[];
  $__res = (object)["genericPred__prime__" => function($v_1) use ($dictGenericEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = (($dictGenericEnum_0)->genericPred__prime__)($v_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_2_0)->value0);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_1) use ($dictGenericEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_2 = (($dictGenericEnum_0)->genericSucc__prime__)($v_1);
  if ((is_object($__local_var_2_2) && (($__local_var_2_2)->tag === "Just"))) {
$__t3 = new Phpurs_Data1("Just", ($__local_var_2_2)->value0);
} else {
$__t3 = new Phpurs_Data0("Nothing");
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericEnumArgument'] = function() { $v = function($dictEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericEnumArgument"), recVars=[];
  $__res = (object)["genericPred__prime__" => function($v_1) use ($dictEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = (($dictEnum_0)->pred)($v_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_2_0)->value0);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericSucc__prime__" => function($v_1) use ($dictEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_2 = (($dictEnum_0)->succ)($v_1);
  if ((is_object($__local_var_2_2) && (($__local_var_2_2)->tag === "Just"))) {
$__t3 = new Phpurs_Data1("Just", ($__local_var_2_2)->value0);
} else {
$__t3 = new Phpurs_Data0("Nothing");
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericCardinality__prime__'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericCardinality'"), recVars=[];
  $__res = ($dict_0)->genericCardinality__prime__;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericCardinality'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0, $dictGenericBoundedEnum_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericCardinality"), recVars=[];
  $__res = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(($dictGenericBoundedEnum_1)->genericCardinality__prime__);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumSum'] = function() { $v = function($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericBoundedEnumSum"), recVars=[];
  $genericCardinality__prime__1_1_0 = ($dictGenericBoundedEnum_0)->genericCardinality__prime__;
  $__res = function($dictGenericBoundedEnum1_2) use ($dictGenericBoundedEnum_0, $genericCardinality__prime__1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["genericCardinality__prime__" => ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($genericCardinality__prime__1_1_0)))((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(($dictGenericBoundedEnum1_2)->genericCardinality__prime__)), "genericToEnum__prime__" => function($n_3) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality__prime__1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_3))(0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_3))(0))->tag === "LT")))))((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_3))($genericCardinality__prime__1_1_0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_3))($genericCardinality__prime__1_1_0))->tag === "LT")))) {
$__local_var_4_4 = (($dictGenericBoundedEnum_0)->genericToEnum__prime__)($n_3);
if ((is_object($__local_var_4_4) && (($__local_var_4_4)->tag === "Just"))) {
$__t5 = new Phpurs_Data1("Just", new Phpurs_Data1("Inl", ($__local_var_4_4)->value0));
} else {
$__t5 = new Phpurs_Data0("Nothing");
};
$__t3 = $__t5;
} else {
$__local_var_4_1 = (($dictGenericBoundedEnum1_2)->genericToEnum__prime__)(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($n_3))($genericCardinality__prime__1_1_0));
if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data1("Inr", ($__local_var_4_1)->value0));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
$__t3 = $__t2;
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_3) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality__prime__1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_3) && (($v_3)->tag === "Inl"))) {
$__t6 = (($dictGenericBoundedEnum_0)->genericFromEnum__prime__)(($v_3)->value0);
} else {
if ((is_object($v_3) && (($v_3)->tag === "Inr"))) {
$__t6 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))((($dictGenericBoundedEnum1_2)->genericFromEnum__prime__)(($v_3)->value0)))((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($genericCardinality__prime__1_1_0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
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
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumProduct'] = function() { $v = function($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericBoundedEnumProduct"), recVars=[];
  $genericCardinality__prime__1_1_0 = ($dictGenericBoundedEnum_0)->genericCardinality__prime__;
  $__res = function($dictGenericBoundedEnum1_2) use ($dictGenericBoundedEnum_0, $genericCardinality__prime__1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $genericCardinality__prime__2_3_1 = ($dictGenericBoundedEnum1_2)->genericCardinality__prime__;
  $__res = (object)["genericCardinality__prime__" => ((($GLOBALS['Data_Semiring_intMul'] ?? \PhpursThunks::eval('Data_Semiring_intMul')))((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($genericCardinality__prime__1_1_0)))((($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($genericCardinality__prime__2_3_1)), "genericToEnum__prime__" => function($n_4) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality__prime__2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_5_2 = (($dictGenericBoundedEnum_0)->genericToEnum__prime__)(((($GLOBALS['Data_EuclideanRing_intDiv'] ?? \PhpursThunks::eval('Data_EuclideanRing_intDiv')))($n_4))($genericCardinality__prime__2_3_1));
  if ((is_object($__local_var_5_2) && (($__local_var_5_2)->tag === "Just"))) {
$__local_var_6_4 = (($dictGenericBoundedEnum1_2)->genericToEnum__prime__)(((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))($n_4))($genericCardinality__prime__2_3_1));
if ((is_object($__local_var_6_4) && (($__local_var_6_4)->tag === "Just"))) {
$__t5 = new Phpurs_Data1("Just", new Phpurs_Data2("Product", ($__local_var_5_2)->value0, ($__local_var_6_4)->value0));
} else {
$__t5 = new Phpurs_Data0("Nothing");
};
$__t3 = $__t5;
} else {
$__t3 = new Phpurs_Data0("Nothing");
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v1_4) use ($dictGenericBoundedEnum1_2, $dictGenericBoundedEnum_0, $genericCardinality__prime__2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(((($GLOBALS['Data_Semiring_intMul'] ?? \PhpursThunks::eval('Data_Semiring_intMul')))((($dictGenericBoundedEnum_0)->genericFromEnum__prime__)(($v1_4)->value0)))($genericCardinality__prime__2_3_1)))((($dictGenericBoundedEnum1_2)->genericFromEnum__prime__)(($v1_4)->value1));
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
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumNoArguments'] = function() { $v = (object)["genericCardinality__prime__" => 1, "genericToEnum__prime__" => function($i_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))($i_0))(0)) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data0("NoArguments"));
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = 0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumConstructor'] = function() { $v = function($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericBoundedEnumConstructor"), recVars=[];
  $__res = (object)["genericCardinality__prime__" => (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(($dictGenericBoundedEnum_0)->genericCardinality__prime__), "genericToEnum__prime__" => function($i_1) use ($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = (($dictGenericBoundedEnum_0)->genericToEnum__prime__)($i_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_2_0)->value0);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_1) use ($dictGenericBoundedEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictGenericBoundedEnum_0)->genericFromEnum__prime__)($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Enum_Generic_genericBoundedEnumArgument'] = function() { $v = function($dictBoundedEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Enum_Generic_genericBoundedEnumArgument"), recVars=[];
  $__res = (object)["genericCardinality__prime__" => (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))(($dictBoundedEnum_0)->cardinality), "genericToEnum__prime__" => function($i_1) use ($dictBoundedEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = (($dictBoundedEnum_0)->toEnum)($i_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_2_0)->value0);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "genericFromEnum__prime__" => function($v_1) use ($dictBoundedEnum_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictBoundedEnum_0)->fromEnum)($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






















