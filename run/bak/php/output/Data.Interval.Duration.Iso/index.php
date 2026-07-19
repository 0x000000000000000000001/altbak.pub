<?php

namespace Data\Interval\Duration\Iso;

// ALL IMPORTS: Control.Applicative, Control.Plus, Control.Semigroupoid, Data.Either, Data.Eq, Data.Foldable, Data.Function, Data.HeytingAlgebra, Data.Interval.Duration, Data.Interval.Duration.Iso, Data.List, Data.List.NonEmpty, Data.List.Types, Data.Map, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Monoid.Additive, Data.Newtype, Data.Number, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Plus, Control.Semigroupoid, Data.Either, Data.Eq, Data.Foldable, Data.Function, Data.HeytingAlgebra, Data.Interval.Duration, Data.Interval.Duration.Iso, Data.List, Data.List.NonEmpty, Data.List.Types, Data.Map, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Monoid.Additive, Data.Newtype, Data.Number, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Interval.Duration/index.php';
require_once __DIR__ . '/../Data.Interval.Duration.Iso/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.NonEmpty/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Map/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Number/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_lookup'] = function() { $v = function($k_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_lookup"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = function($v_2) use (&$go_1_0, $k_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_2) && (($v_2)->tag === "Node"))) {
if ((is_object($k_0) && (($k_0)->tag === "Second"))) {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Second"))) {
$__t4 = new Phpurs_Data1("Just", ($v_2)->value3);
} else {
$__tco_3 = ($v_2)->value4;
$v_2 = $__tco_3;
continue ;
$__t4 = null;
};
$__t2 = $__t4;
} else {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Second"))) {
$__tco_5 = ($v_2)->value5;
$v_2 = $__tco_5;
continue ;
$__t2 = null;
} else {
if ((is_object($k_0) && (($k_0)->tag === "Minute"))) {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Minute"))) {
$__t7 = new Phpurs_Data1("Just", ($v_2)->value3);
} else {
$__tco_6 = ($v_2)->value4;
$v_2 = $__tco_6;
continue ;
$__t7 = null;
};
$__t2 = $__t7;
} else {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Minute"))) {
$__tco_8 = ($v_2)->value5;
$v_2 = $__tco_8;
continue ;
$__t2 = null;
} else {
if ((is_object($k_0) && (($k_0)->tag === "Hour"))) {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Hour"))) {
$__t10 = new Phpurs_Data1("Just", ($v_2)->value3);
} else {
$__tco_9 = ($v_2)->value4;
$v_2 = $__tco_9;
continue ;
$__t10 = null;
};
$__t2 = $__t10;
} else {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Hour"))) {
$__tco_11 = ($v_2)->value5;
$v_2 = $__tco_11;
continue ;
$__t2 = null;
} else {
if ((is_object($k_0) && (($k_0)->tag === "Day"))) {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Day"))) {
$__t13 = new Phpurs_Data1("Just", ($v_2)->value3);
} else {
$__tco_12 = ($v_2)->value4;
$v_2 = $__tco_12;
continue ;
$__t13 = null;
};
$__t2 = $__t13;
} else {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Day"))) {
$__tco_14 = ($v_2)->value5;
$v_2 = $__tco_14;
continue ;
$__t2 = null;
} else {
if ((is_object($k_0) && (($k_0)->tag === "Week"))) {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Week"))) {
$__t16 = new Phpurs_Data1("Just", ($v_2)->value3);
} else {
$__tco_15 = ($v_2)->value4;
$v_2 = $__tco_15;
continue ;
$__t16 = null;
};
$__t2 = $__t16;
} else {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Week"))) {
$__tco_17 = ($v_2)->value5;
$v_2 = $__tco_17;
continue ;
$__t2 = null;
} else {
if ((is_object($k_0) && (($k_0)->tag === "Month"))) {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Month"))) {
$__t19 = new Phpurs_Data1("Just", ($v_2)->value3);
} else {
$__tco_18 = ($v_2)->value4;
$v_2 = $__tco_18;
continue ;
$__t19 = null;
};
$__t2 = $__t19;
} else {
if ((is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Month"))) {
$__tco_20 = ($v_2)->value5;
$v_2 = $__tco_20;
continue ;
$__t2 = null;
} else {
if (((is_object($k_0) && (($k_0)->tag === "Year")) && (is_object(($v_2)->value2) && ((($v_2)->value2)->tag === "Year")))) {
$__t2 = new Phpurs_Data1("Just", ($v_2)->value3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
};
};
};
};
};
};
};
};
};
};
$__t1 = $__t2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_foldMap1'] = function() { $v = ((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldMap)(($GLOBALS['Data_List_Types_monoidList'] ?? \PhpursThunks::eval('Data_List_Types_monoidList'))); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_foldMap2'] = function() { $v = (function() use (&$__fn) {
$semigroupAdditive1_0_0 = (object)["append" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($v_0 + $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return ((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldMap)((object)["mempty" => 0.0, "Semigroup0" => function($dollar__unused_1) use ($semigroupAdditive1_0_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $semigroupAdditive1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_fold'] = function() { $v = (((($GLOBALS['Data_Foldable_foldMapDefaultR'] ?? \PhpursThunks::eval('Data_Foldable_foldMapDefaultR')))(($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'))))((($GLOBALS['Data_Monoid_monoidFn'] ?? \PhpursThunks::eval('Data_Monoid_monoidFn')))(($GLOBALS['Data_List_Types_monoidList'] ?? \PhpursThunks::eval('Data_List_Types_monoidList')))))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_toUnfoldable'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($b_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($source_2, $memo_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
$v_4_1 = (($GLOBALS['Data_Map_Internal_stepUnfoldr'] ?? \PhpursThunks::eval('Data_Map_Internal_stepUnfoldr')))($source_2);
if ((is_object($v_4_1) && (($v_4_1)->tag === "Nothing"))) {
$go_5_3 = null;
$go_5_3 = (function() use (&$go_5_3) {
  $__fn = function($b_6, $v_7 = null) use (&$go_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_3"), recVars=["go_1_0","go_5_3"];
  while (true) {
if ((is_object($v_7) && (($v_7)->tag === "Nil"))) {
$__t4 = $b_6;
} else {
if ((is_object($v_7) && (($v_7)->tag === "Cons"))) {
$__tco_5 = new Phpurs_Data2("Cons", ($v_7)->value0, $b_6);
$__tco_6 = ($v_7)->value1;
$b_6 = $__tco_5;
$v_7 = $__tco_6;
continue ;
$__t4 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__res = $__t4;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t2 = (($go_5_3)(new Phpurs_Data0("Nil")))($memo_3);
} else {
if ((is_object($v_4_1) && (($v_4_1)->tag === "Just"))) {
$__tco_7 = (($v_4_1)->value0)->value1;
$__tco_8 = new Phpurs_Data2("Cons", (($v_4_1)->value0)->value0, $memo_3);
$source_2 = $__tco_7;
$memo_3 = $__tco_8;
continue ;
$__t2 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_1_0)($b_0))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Map_Internal_toMapIter'] ?? \PhpursThunks::eval('Data_Map_Internal_toMapIter'))); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_IsEmpty'] = function() { $v = ($GLOBALS['__phpurs_data0_IsEmpty'] ??= new Phpurs_Data0("IsEmpty")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_InvalidWeekComponentUsage'] = function() { $v = ($GLOBALS['__phpurs_data0_InvalidWeekComponentUsage'] ??= new Phpurs_Data0("InvalidWeekComponentUsage")); return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_ContainsNegativeValue'] = function() { $v = function($value0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("ContainsNegativeValue", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_InvalidFractionalUse'] = function() { $v = function($value0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("InvalidFractionalUse", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_unIsoDuration'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_unIsoDuration"), recVars=[];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_showIsoDuration'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (("(IsoDuration (Duration " . (($GLOBALS['Data_Interval_Duration_show'] ?? \PhpursThunks::eval('Data_Interval_Duration_show')))($v_0)) . "))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_showError'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "IsEmpty"))) {
$__t0 = "(IsEmpty)";
} else {
if ((is_object($v_0) && (($v_0)->tag === "InvalidWeekComponentUsage"))) {
$__t0 = "(InvalidWeekComponentUsage)";
} else {
if ((is_object($v_0) && (($v_0)->tag === "ContainsNegativeValue"))) {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Minute"))) {
$__t1 = "(ContainsNegativeValue Minute)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Second"))) {
$__t1 = "(ContainsNegativeValue Second)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Hour"))) {
$__t1 = "(ContainsNegativeValue Hour)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Day"))) {
$__t1 = "(ContainsNegativeValue Day)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Week"))) {
$__t1 = "(ContainsNegativeValue Week)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Month"))) {
$__t1 = "(ContainsNegativeValue Month)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Year"))) {
$__t1 = "(ContainsNegativeValue Year)";
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
};
};
};
};
$__t0 = $__t1;
} else {
if ((is_object($v_0) && (($v_0)->tag === "InvalidFractionalUse"))) {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Minute"))) {
$__t2 = "(InvalidFractionalUse Minute)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Second"))) {
$__t2 = "(InvalidFractionalUse Second)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Hour"))) {
$__t2 = "(InvalidFractionalUse Hour)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Day"))) {
$__t2 = "(InvalidFractionalUse Day)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Week"))) {
$__t2 = "(InvalidFractionalUse Week)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Month"))) {
$__t2 = "(InvalidFractionalUse Month)";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Year"))) {
$__t2 = "(InvalidFractionalUse Year)";
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
};
};
};
};
$__t0 = $__t2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_prettyError'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_prettyError"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "IsEmpty"))) {
$__t0 = "Duration is empty (has no components)";
} else {
if ((is_object($v_0) && (($v_0)->tag === "InvalidWeekComponentUsage"))) {
$__t0 = "Week component of Duration is used with other components";
} else {
if ((is_object($v_0) && (($v_0)->tag === "ContainsNegativeValue"))) {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Minute"))) {
$__t1 = "Component `Minute` contains negative value";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Second"))) {
$__t1 = "Component `Second` contains negative value";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Hour"))) {
$__t1 = "Component `Hour` contains negative value";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Day"))) {
$__t1 = "Component `Day` contains negative value";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Week"))) {
$__t1 = "Component `Week` contains negative value";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Month"))) {
$__t1 = "Component `Month` contains negative value";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Year"))) {
$__t1 = "Component `Year` contains negative value";
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
};
};
};
};
$__t0 = $__t1;
} else {
if ((is_object($v_0) && (($v_0)->tag === "InvalidFractionalUse"))) {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Minute"))) {
$__t2 = "Invalid usage of Fractional value at component `Minute`";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Second"))) {
$__t2 = "Invalid usage of Fractional value at component `Second`";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Hour"))) {
$__t2 = "Invalid usage of Fractional value at component `Hour`";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Day"))) {
$__t2 = "Invalid usage of Fractional value at component `Day`";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Week"))) {
$__t2 = "Invalid usage of Fractional value at component `Week`";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Month"))) {
$__t2 = "Invalid usage of Fractional value at component `Month`";
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Year"))) {
$__t2 = "Invalid usage of Fractional value at component `Year`";
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
};
};
};
};
$__t0 = $__t2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_eqIsoDuration'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Interval_Duration_eq'] ?? \PhpursThunks::eval('Data_Interval_Duration_eq')))($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_ordIsoDuration'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Interval_Duration_compare'] ?? \PhpursThunks::eval('Data_Interval_Duration_compare')))($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Interval_Duration_Iso_eqIsoDuration'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_eqIsoDuration'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_eqError'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "IsEmpty"))) {
$__t1 = (is_object($y_1) && (($y_1)->tag === "IsEmpty"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "InvalidWeekComponentUsage"))) {
$__t1 = (is_object($y_1) && (($y_1)->tag === "InvalidWeekComponentUsage"));
} else {
if ((is_object($x_0) && (($x_0)->tag === "ContainsNegativeValue"))) {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Second"))) {
$__t2 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Second"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Minute"))) {
$__t2 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Minute"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Hour"))) {
$__t2 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Hour"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Day"))) {
$__t2 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Day"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Week"))) {
$__t2 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Week"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Month"))) {
$__t2 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Month"));
} else {
$__t2 = ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Year")) && (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Year")));
};
};
};
};
};
};
$__t1 = ((is_object($y_1) && (($y_1)->tag === "ContainsNegativeValue")) && $__t2);
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Second"))) {
$__t0 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Second"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Minute"))) {
$__t0 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Minute"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Hour"))) {
$__t0 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Hour"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Day"))) {
$__t0 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Day"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Week"))) {
$__t0 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Week"));
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Month"))) {
$__t0 = (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Month"));
} else {
$__t0 = ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Year")) && (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Year")));
};
};
};
};
};
};
$__t1 = ((is_object($x_0) && (($x_0)->tag === "InvalidFractionalUse")) && ((is_object($y_1) && (($y_1)->tag === "InvalidFractionalUse")) && $__t0));
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_ordError'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "IsEmpty"))) {
if ((is_object($y_1) && (($y_1)->tag === "IsEmpty"))) {
$__t1 = new Phpurs_Data0("EQ");
} else {
$__t1 = new Phpurs_Data0("LT");
};
$__t0 = $__t1;
} else {
if ((is_object($y_1) && (($y_1)->tag === "IsEmpty"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "InvalidWeekComponentUsage"))) {
if ((is_object($y_1) && (($y_1)->tag === "InvalidWeekComponentUsage"))) {
$__t2 = new Phpurs_Data0("EQ");
} else {
$__t2 = new Phpurs_Data0("LT");
};
$__t0 = $__t2;
} else {
if ((is_object($y_1) && (($y_1)->tag === "InvalidWeekComponentUsage"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_0) && (($x_0)->tag === "ContainsNegativeValue"))) {
if ((is_object($y_1) && (($y_1)->tag === "ContainsNegativeValue"))) {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Second"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Second"))) {
$__t5 = new Phpurs_Data0("EQ");
} else {
$__t5 = new Phpurs_Data0("LT");
};
$__t4 = $__t5;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Second"))) {
$__t4 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Minute"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Minute"))) {
$__t6 = new Phpurs_Data0("EQ");
} else {
$__t6 = new Phpurs_Data0("LT");
};
$__t4 = $__t6;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Minute"))) {
$__t4 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Hour"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Hour"))) {
$__t7 = new Phpurs_Data0("EQ");
} else {
$__t7 = new Phpurs_Data0("LT");
};
$__t4 = $__t7;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Hour"))) {
$__t4 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Day"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Day"))) {
$__t8 = new Phpurs_Data0("EQ");
} else {
$__t8 = new Phpurs_Data0("LT");
};
$__t4 = $__t8;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Day"))) {
$__t4 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Week"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Week"))) {
$__t9 = new Phpurs_Data0("EQ");
} else {
$__t9 = new Phpurs_Data0("LT");
};
$__t4 = $__t9;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Week"))) {
$__t4 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Month"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Month"))) {
$__t10 = new Phpurs_Data0("EQ");
} else {
$__t10 = new Phpurs_Data0("LT");
};
$__t4 = $__t10;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Month"))) {
$__t4 = new Phpurs_Data0("GT");
} else {
if (((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Year")) && (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Year")))) {
$__t4 = new Phpurs_Data0("EQ");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
};
};
};
};
};
};
};
};
};
};
};
$__t3 = $__t4;
} else {
$__t3 = new Phpurs_Data0("LT");
};
$__t0 = $__t3;
} else {
if ((is_object($y_1) && (($y_1)->tag === "ContainsNegativeValue"))) {
$__t0 = new Phpurs_Data0("GT");
} else {
if (((is_object($x_0) && (($x_0)->tag === "InvalidFractionalUse")) && (is_object($y_1) && (($y_1)->tag === "InvalidFractionalUse")))) {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Second"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Second"))) {
$__t12 = new Phpurs_Data0("EQ");
} else {
$__t12 = new Phpurs_Data0("LT");
};
$__t11 = $__t12;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Second"))) {
$__t11 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Minute"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Minute"))) {
$__t13 = new Phpurs_Data0("EQ");
} else {
$__t13 = new Phpurs_Data0("LT");
};
$__t11 = $__t13;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Minute"))) {
$__t11 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Hour"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Hour"))) {
$__t14 = new Phpurs_Data0("EQ");
} else {
$__t14 = new Phpurs_Data0("LT");
};
$__t11 = $__t14;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Hour"))) {
$__t11 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Day"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Day"))) {
$__t15 = new Phpurs_Data0("EQ");
} else {
$__t15 = new Phpurs_Data0("LT");
};
$__t11 = $__t15;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Day"))) {
$__t11 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Week"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Week"))) {
$__t16 = new Phpurs_Data0("EQ");
} else {
$__t16 = new Phpurs_Data0("LT");
};
$__t11 = $__t16;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Week"))) {
$__t11 = new Phpurs_Data0("GT");
} else {
if ((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Month"))) {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Month"))) {
$__t17 = new Phpurs_Data0("EQ");
} else {
$__t17 = new Phpurs_Data0("LT");
};
$__t11 = $__t17;
} else {
if ((is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Month"))) {
$__t11 = new Phpurs_Data0("GT");
} else {
if (((is_object(($x_0)->value0) && ((($x_0)->value0)->tag === "Year")) && (is_object(($y_1)->value0) && ((($y_1)->value0)->tag === "Year")))) {
$__t11 = new Phpurs_Data0("EQ");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
};
};
};
};
};
};
};
};
};
};
};
};
};
$__t0 = $__t11;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
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
  $__res = ($GLOBALS['Data_Interval_Duration_Iso_eqError'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_eqError'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_checkWeekUsage'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_checkWeekUsage"), recVars=[];
  if ((function() use ($v_0, &$__fn) {
$__local_var_1_1 = ((($GLOBALS['Data_Interval_Duration_Iso_lookup'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_lookup')))(new Phpurs_Data0("Week")))(($v_0)->asMap);
if ((is_object($__local_var_1_1) && (($__local_var_1_1)->tag === "Nothing"))) {
$__t2 = false;
} else {
if ((is_object($__local_var_1_1) && (($__local_var_1_1)->tag === "Just"))) {
$__t2 = true;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
if ((is_object(($v_0)->asMap) && ((($v_0)->asMap)->tag === "Leaf"))) {
$__t3 = false;
} else {
if ((is_object(($v_0)->asMap) && ((($v_0)->asMap)->tag === "Node"))) {
$__t3 = ((($v_0)->asMap)->value1 > 1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
return ($__t2 && $__t3);
})()) {
$__t0 = new Phpurs_Data2("Cons", new Phpurs_Data0("InvalidWeekComponentUsage"), new Phpurs_Data0("Nil"));
} else {
$__t0 = new Phpurs_Data0("Nil");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_checkNegativeValues'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_checkNegativeValues"), recVars=[];
  $__res = ((($GLOBALS['Data_Interval_Duration_Iso_foldMap1'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_foldMap1')))(function($v1_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((($v1_1)->value1 >= 0.0)) {
$__t0 = new Phpurs_Data0("Nil");
} else {
$__t0 = new Phpurs_Data2("Cons", new Phpurs_Data1("ContainsNegativeValue", ($v1_1)->value0), new Phpurs_Data0("Nil"));
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_0)->asList);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_checkFractionalUse'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_checkFractionalUse"), recVars=[];
  $__local_var_1_0 = ((($GLOBALS['Data_List_span'] ?? \PhpursThunks::eval('Data_List_span')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($a_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Number_floor'] ?? \PhpursThunks::eval('Data_Number_floor')))($a_1) === $a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Tuple_snd'] ?? \PhpursThunks::eval('Data_Tuple_snd')))))(($v_0)->asList);
  if (((is_object(($__local_var_1_0)->rest) && ((($__local_var_1_0)->rest)->tag === "Cons")) && (((($GLOBALS['Data_Interval_Duration_Iso_foldMap2'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_foldMap2')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Monoid_Additive_Additive'] ?? \PhpursThunks::eval('Data_Monoid_Additive_Additive'))))(($GLOBALS['Data_Number_abs'] ?? \PhpursThunks::eval('Data_Number_abs')))))(($GLOBALS['Data_Tuple_snd'] ?? \PhpursThunks::eval('Data_Tuple_snd')))))((($__local_var_1_0)->rest)->value1) > 0.0))) {
$__t1 = new Phpurs_Data2("Cons", new Phpurs_Data1("InvalidFractionalUse", ((($__local_var_1_0)->rest)->value0)->value0), new Phpurs_Data0("Nil"));
} else {
$__t1 = new Phpurs_Data0("Nil");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_checkEmptiness'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_checkEmptiness"), recVars=[];
  if ((is_object(($v_0)->asList) && ((($v_0)->asList)->tag === "Nil"))) {
$__t0 = new Phpurs_Data2("Cons", new Phpurs_Data0("IsEmpty"), new Phpurs_Data0("Nil"));
} else {
$__t0 = new Phpurs_Data0("Nil");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_checkValidIsoDuration'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_checkValidIsoDuration"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$__t1 = $v_2;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_3)->value0, $v_2);
$__tco_3 = ($v1_3)->value1;
$v_2 = $__tco_2;
$v1_3 = $__tco_3;
continue ;
$__t1 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Data_Interval_Duration_Iso_fold'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_fold')))([($GLOBALS['Data_Interval_Duration_Iso_checkWeekUsage'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_checkWeekUsage')), ($GLOBALS['Data_Interval_Duration_Iso_checkEmptiness'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_checkEmptiness')), ($GLOBALS['Data_Interval_Duration_Iso_checkFractionalUse'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_checkFractionalUse')), ($GLOBALS['Data_Interval_Duration_Iso_checkNegativeValues'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_checkNegativeValues'))]))((object)["asList" => (($go_1_0)(new Phpurs_Data0("Nil")))((($GLOBALS['Data_Interval_Duration_Iso_toUnfoldable'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_toUnfoldable')))($v_0)), "asMap" => $v_0]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_Duration_Iso_mkIsoDuration'] = function() { $v = function($d_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_Duration_Iso_mkIsoDuration"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_Interval_Duration_Iso_checkValidIsoDuration'] ?? \PhpursThunks::eval('Data_Interval_Duration_Iso_checkValidIsoDuration')))($d_0);
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Nil"))) {
$__t1 = new Phpurs_Data1("Right", $d_0);
} else {
if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Cons"))) {
$__t1 = new Phpurs_Data1("Left", new Phpurs_Data2("NonEmpty", ($__local_var_1_0)->value0, ($__local_var_1_0)->value1));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };

























