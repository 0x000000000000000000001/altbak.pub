<?php

namespace Data\String\Regex\Flags;

// ALL IMPORTS: Control.Alternative, Control.MonadPlus, Data.Eq, Data.Functor, Data.HeytingAlgebra, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Show, Data.String, Data.String.Common, Data.String.Regex.Flags, Data.Symbol, Prelude, Prim
// TO REQUIRE: Control.Alternative, Control.MonadPlus, Data.Eq, Data.Functor, Data.HeytingAlgebra, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Show, Data.String, Data.String.Common, Data.String.Regex.Flags, Data.Symbol, Prelude
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.String/index.php';
require_once __DIR__ . '/../Data.String.Common/index.php';
require_once __DIR__ . '/../Data.String.Regex.Flags/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_String_Regex_Flags_eqArray
$GLOBALS['Data_String_Regex_Flags_eqArray'] = (object)["eq" => ($GLOBALS['Data_Eq_eqArrayImpl'])($GLOBALS['Data_Eq_eqStringImpl'])];

// Data_String_Regex_Flags_RegexFlags
function majData_majString_majRegex_majFlags_majRegexmajFlags($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majRegex_majFlags_majRegexmajFlags';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_Regex_Flags_RegexFlags'] = __NAMESPACE__ . '\\majData_majString_majRegex_majFlags_majRegexmajFlags';

// Data_String_Regex_Flags_unicode
$GLOBALS['Data_String_Regex_Flags_unicode'] = (object)["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => true];

// Data_String_Regex_Flags_sticky
$GLOBALS['Data_String_Regex_Flags_sticky'] = (object)["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => true, "unicode" => false];

// Data_String_Regex_Flags_showRegexFlags
$GLOBALS['Data_String_Regex_Flags_showRegexFlags'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($v_0)->{'global'}) {
$__t0 = [$GLOBALS['Data_Unit_unit']];
goto end_branch_0;;
};
  $__t0 = [];
  end_branch_0:;
  $__t1 = null;;
  if (($v_0)->{'ignoreCase'}) {
$__t1 = [$GLOBALS['Data_Unit_unit']];
goto end_branch_1;;
};
  $__t1 = [];
  end_branch_1:;
  $__t2 = null;;
  if (($v_0)->{'multiline'}) {
$__t2 = [$GLOBALS['Data_Unit_unit']];
goto end_branch_2;;
};
  $__t2 = [];
  end_branch_2:;
  $__t3 = null;;
  if (($v_0)->{'dotAll'}) {
$__t3 = [$GLOBALS['Data_Unit_unit']];
goto end_branch_3;;
};
  $__t3 = [];
  end_branch_3:;
  $__t4 = null;;
  if (($v_0)->{'sticky'}) {
$__t4 = [$GLOBALS['Data_Unit_unit']];
goto end_branch_4;;
};
  $__t4 = [];
  end_branch_4:;
  $__t5 = null;;
  if (($v_0)->{'unicode'}) {
$__t5 = [$GLOBALS['Data_Unit_unit']];
goto end_branch_5;;
};
  $__t5 = [];
  end_branch_5:;
  $usedFlags_1_0 = \Data\Semigroup\majData_majSemigroup_concatmajArray(\Data\Semigroup\majData_majSemigroup_concatmajArray(\Data\Semigroup\majData_majSemigroup_concatmajArray(\Data\Semigroup\majData_majSemigroup_concatmajArray(\Data\Semigroup\majData_majSemigroup_concatmajArray(\Data\Semigroup\majData_majSemigroup_concatmajArray([], \Data\Functor\majData_majFunctor_arraymajMap(function($v_1) {
  $__num = \func_num_args();
  $__res = "global";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__t0)), \Data\Functor\majData_majFunctor_arraymajMap(function($v_1) {
  $__num = \func_num_args();
  $__res = "ignoreCase";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__t1)), \Data\Functor\majData_majFunctor_arraymajMap(function($v_1) {
  $__num = \func_num_args();
  $__res = "multiline";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__t2)), \Data\Functor\majData_majFunctor_arraymajMap(function($v_1) {
  $__num = \func_num_args();
  $__res = "dotAll";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__t3)), \Data\Functor\majData_majFunctor_arraymajMap(function($v_1) {
  $__num = \func_num_args();
  $__res = "sticky";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__t4)), \Data\Functor\majData_majFunctor_arraymajMap(function($v_1) {
  $__num = \func_num_args();
  $__res = "unicode";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__t5));
  $__t7 = null;;
  if (((($GLOBALS['Data_String_Regex_Flags_eqArray'])->{'eq'})($usedFlags_1_0))([])) {
$__t7 = "noFlags";
goto end_branch_7;;
};
  $__t7 = (("(" . \Data\String\Common\majData_majString_majCommon_joinmajWith(" <> ", $usedFlags_1_0)) . ")");
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_Regex_Flags_semigroupRegexFlags
$GLOBALS['Data_String_Regex_Flags_semigroupRegexFlags'] = (object)["append" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = (object)["global" => (($v_0)->{'global'} || ($v1_1)->{'global'}), "ignoreCase" => (($v_0)->{'ignoreCase'} || ($v1_1)->{'ignoreCase'}), "multiline" => (($v_0)->{'multiline'} || ($v1_1)->{'multiline'}), "dotAll" => (($v_0)->{'dotAll'} || ($v1_1)->{'dotAll'}), "sticky" => (($v_0)->{'sticky'} || ($v1_1)->{'sticky'}), "unicode" => (($v_0)->{'unicode'} || ($v1_1)->{'unicode'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_Regex_Flags_noFlags
$GLOBALS['Data_String_Regex_Flags_noFlags'] = (object)["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => false];

// Data_String_Regex_Flags_newtypeRegexFlags
$GLOBALS['Data_String_Regex_Flags_newtypeRegexFlags'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_Regex_Flags_multiline
$GLOBALS['Data_String_Regex_Flags_multiline'] = (object)["global" => false, "ignoreCase" => false, "multiline" => true, "dotAll" => false, "sticky" => false, "unicode" => false];

// Data_String_Regex_Flags_monoidRegexFlags
$GLOBALS['Data_String_Regex_Flags_monoidRegexFlags'] = (object)["mempty" => $GLOBALS['Data_String_Regex_Flags_noFlags'], "Semigroup0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_Regex_Flags_semigroupRegexFlags'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_Regex_Flags_ignoreCase
$GLOBALS['Data_String_Regex_Flags_ignoreCase'] = (object)["global" => false, "ignoreCase" => true, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => false];

// Data_String_Regex_Flags_global
$GLOBALS['Data_String_Regex_Flags_global'] = (object)["global" => true, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => false];

// Data_String_Regex_Flags_eqRegexFlags
$GLOBALS['Data_String_Regex_Flags_eqRegexFlags'] = (function() use (&$__fn) {
$__local_var_0_0 = (object)["eqRecord" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($ra_1) {
  $__num = \func_num_args();
  $__res = function($rb_2) use ($ra_1) {
  $__num = \func_num_args();
  $__res = (($ra_1)->{'unicode'} === ($rb_2)->{'unicode'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__local_var_1_1 = (object)["eqRecord" => function($v_1) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__res = function($ra_2) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__res = function($rb_3) use ($__local_var_0_0, $ra_2) {
  $__num = \func_num_args();
  $__res = ((($ra_2)->{'sticky'} === ($rb_3)->{'sticky'}) && (((($__local_var_0_0)->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_2))($rb_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__local_var_2_2 = (object)["eqRecord" => function($v_2) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = function($ra_3) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = function($rb_4) use ($__local_var_1_1, $ra_3) {
  $__num = \func_num_args();
  $__res = ((($ra_3)->{'multiline'} === ($rb_4)->{'multiline'}) && (((($__local_var_1_1)->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_3))($rb_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__local_var_3_3 = (object)["eqRecord" => function($v_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = function($ra_4) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = function($rb_5) use ($__local_var_2_2, $ra_4) {
  $__num = \func_num_args();
  $__res = ((($ra_4)->{'ignoreCase'} === ($rb_5)->{'ignoreCase'}) && (((($__local_var_2_2)->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_4))($rb_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__local_var_4_4 = (object)["eqRecord" => function($v_4) use ($__local_var_3_3) {
  $__num = \func_num_args();
  $__res = function($ra_5) use ($__local_var_3_3) {
  $__num = \func_num_args();
  $__res = function($rb_6) use ($__local_var_3_3, $ra_5) {
  $__num = \func_num_args();
  $__res = ((($ra_5)->{'global'} === ($rb_6)->{'global'}) && (((($__local_var_3_3)->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_5))($rb_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
return (object)["eq" => function($ra_5) use ($__local_var_4_4) {
  $__num = \func_num_args();
  $__res = function($rb_6) use ($__local_var_4_4, $ra_5) {
  $__num = \func_num_args();
  $__res = ((($ra_5)->{'dotAll'} === ($rb_6)->{'dotAll'}) && (((($__local_var_4_4)->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_5))($rb_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_String_Regex_Flags_dotAll
$GLOBALS['Data_String_Regex_Flags_dotAll'] = (object)["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => true, "sticky" => false, "unicode" => false];

