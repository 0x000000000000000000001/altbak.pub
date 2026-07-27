<?php

namespace Data\String\NonEmpty\CodeUnits;

// ALL IMPORTS: Control.Semigroupoid, Data.Array.NonEmpty, Data.Maybe, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.String.CodeUnits, Data.String.NonEmpty.CodeUnits, Data.String.NonEmpty.Internal, Data.String.Pattern, Data.String.Unsafe, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Array.NonEmpty, Data.Maybe, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.String.CodeUnits, Data.String.NonEmpty.CodeUnits, Data.String.NonEmpty.Internal, Data.String.Pattern, Data.String.Unsafe, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.String.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.NonEmpty.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
require_once __DIR__ . '/../Data.String.Unsafe/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
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




// Data_String_NonEmpty_CodeUnits_lessThan
$GLOBALS['Data_String_NonEmpty_CodeUnits_lessThan'] = (function() use (&$__fn) {
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

// Data_String_NonEmpty_CodeUnits_fromJust
function majData_majString_majNonmajEmpty_majCodemajUnits_frommajJust($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_frommajJust';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_fromJust'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_frommajJust';

// Data_String_NonEmpty_CodeUnits_greaterThanOrEq
$GLOBALS['Data_String_NonEmpty_CodeUnits_greaterThanOrEq'] = (function() use (&$__fn) {
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

// Data_String_NonEmpty_CodeUnits_snoc
function majData_majString_majNonmajEmpty_majCodemajUnits_snoc($c_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_snoc';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($s_1))(\Data\String\CodeUnits\majData_majString_majCodemajUnits_singleton($c_0));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_snoc'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_snoc';

// Data_String_NonEmpty_CodeUnits_singleton_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_singleton_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_NonEmptyString']))($GLOBALS['Data_String_CodeUnits_singleton']);

// Data_String_NonEmpty_CodeUnits_singleton
function majData_majString_majNonmajEmpty_majCodemajUnits_singleton($v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_singleton_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_singleton'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_singleton';

// Data_String_NonEmpty_CodeUnits_liftS
function majData_majString_majNonmajEmpty_majCodemajUnits_liftmajS($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_liftmajS';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_liftS'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_liftmajS';

// Data_String_NonEmpty_CodeUnits_takeWhile
function majData_majString_majNonmajEmpty_majCodemajUnits_takemajWhile($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_takemajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_fromString']))(function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\String\CodeUnits\majData_majString_majCodemajUnits_take(\Data\String\CodeUnits\majData_majString_majCodemajUnits_countmajPrefix($f_0, $v_1), $v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_takeWhile'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_takemajWhile';

// Data_String_NonEmpty_CodeUnits_lastIndexOf'
function majData_majString_majNonmajEmpty_majCodemajUnits_lastmajIndexmajOf__prime__(string $pat_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_lastmajIndexmajOf__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS']))(($GLOBALS['Data_String_CodeUnits_lastIndexOf__prime__'])($pat_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_lastIndexOf__prime__'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_lastmajIndexmajOf__prime__';

// Data_String_NonEmpty_CodeUnits_lastIndexOf_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_lastIndexOf_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS']))($GLOBALS['Data_String_CodeUnits_lastIndexOf']);

// Data_String_NonEmpty_CodeUnits_lastIndexOf
function majData_majString_majNonmajEmpty_majCodemajUnits_lastmajIndexmajOf(string $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_lastmajIndexmajOf';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_lastIndexOf_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_lastIndexOf'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_lastmajIndexmajOf';

// Data_String_NonEmpty_CodeUnits_indexOf'
function majData_majString_majNonmajEmpty_majCodemajUnits_indexmajOf__prime__(string $pat_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_indexmajOf__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS']))(($GLOBALS['Data_String_CodeUnits_indexOf__prime__'])($pat_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_indexOf__prime__'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_indexmajOf__prime__';

// Data_String_NonEmpty_CodeUnits_indexOf_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_indexOf_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS']))($GLOBALS['Data_String_CodeUnits_indexOf']);

// Data_String_NonEmpty_CodeUnits_indexOf
function majData_majString_majNonmajEmpty_majCodemajUnits_indexmajOf(string $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_indexmajOf';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_indexOf_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_indexOf'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_indexmajOf';

// Data_String_NonEmpty_CodeUnits_fromNonEmptyString
function majData_majString_majNonmajEmpty_majCodemajUnits_frommajNonmajEmptymajString(string $v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_frommajNonmajEmptymajString';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyString'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_frommajNonmajEmptymajString';

// Data_String_NonEmpty_CodeUnits_length_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_length_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_CodeUnits_length']))($GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyString']);

// Data_String_NonEmpty_CodeUnits_length
function majData_majString_majNonmajEmpty_majCodemajUnits_length(string $v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_length_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_length'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_length';

// Data_String_NonEmpty_CodeUnits_splitAt
function majData_majString_majNonmajEmpty_majCodemajUnits_splitmajAt(int $i_0, $nes_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_splitmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = \Data\String\CodeUnits\majData_majString_majCodemajUnits_splitmajAt($i_0, $nes_1);
  $__res = (object)["before" => match (($v_2_0)->{'before'}) { "" => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just(($v_2_0)->{'before'}) }, "after" => match (($v_2_0)->{'after'}) { "" => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just(($v_2_0)->{'after'}) }];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_splitAt'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_splitmajAt';

// Data_String_NonEmpty_CodeUnits_take
function majData_majString_majNonmajEmpty_majCodemajUnits_take(int $i_0, $nes_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_take';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((($GLOBALS['Data_String_NonEmpty_CodeUnits_lessThan'])($i_0))(1)) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Just(\Data\String\CodeUnits\majData_majString_majCodemajUnits_take($i_0, $nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_take'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_take';

// Data_String_NonEmpty_CodeUnits_takeRight
function majData_majString_majNonmajEmpty_majCodemajUnits_takemajRight(int $i_0, $nes_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_takemajRight';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((($GLOBALS['Data_String_NonEmpty_CodeUnits_lessThan'])($i_0))(1)) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Just(\Data\String\CodeUnits\majData_majString_majCodemajUnits_drop((\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($nes_1) - $i_0), $nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_takeRight'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_takemajRight';

// Data_String_NonEmpty_CodeUnits_toChar_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_toChar_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_CodeUnits_toChar']))($GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyString']);

// Data_String_NonEmpty_CodeUnits_toChar
function majData_majString_majNonmajEmpty_majCodemajUnits_tomajChar(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_tomajChar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_toChar_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_toChar'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_tomajChar';

// Data_String_NonEmpty_CodeUnits_toCharArray_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_toCharArray_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_CodeUnits_toCharArray']))($GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyString']);

// Data_String_NonEmpty_CodeUnits_toCharArray
function majData_majString_majNonmajEmpty_majCodemajUnits_tomajCharmajArray(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_tomajCharmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_toCharArray_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_toCharArray'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_tomajCharmajArray';

// Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_fromJust']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromArray']))($GLOBALS['Data_String_NonEmpty_CodeUnits_toCharArray']));

// Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray
function majData_majString_majNonmajEmpty_majCodemajUnits_tomajNonmajEmptymajCharmajArray(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_tomajNonmajEmptymajCharmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_tomajNonmajEmptymajCharmajArray';

// Data_String_NonEmpty_CodeUnits_uncons
function majData_majString_majNonmajEmpty_majCodemajUnits_uncons(string $nes_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(1, $nes_0);
  $__res = (object)["head" => \Data\String\Unsafe\majData_majString_majUnsafe_charmajAt(0, $nes_0), "tail" => match ($__local_var_1_0) { "" => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just($__local_var_1_0) }];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_uncons'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_uncons';

// Data_String_NonEmpty_CodeUnits_fromFoldable1
function majData_majString_majNonmajEmpty_majCodemajUnits_frommajFoldable1($dictFoldable1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_frommajFoldable1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($dictFoldable1_0)->{'foldMap1'})($GLOBALS['Data_Semigroup_semigroupString']))($GLOBALS['Data_String_NonEmpty_CodeUnits_singleton']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_fromFoldable1'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_frommajFoldable1';

// Data_String_NonEmpty_CodeUnits_fromCharArray
function majData_majString_majNonmajEmpty_majCodemajUnits_frommajCharmajArray($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_frommajCharmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = match (count($v_0)) { 0 => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just(\Data\String\CodeUnits\majData_majString_majCodemajUnits_frommajCharmajArray($v_0)) };
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_fromCharArray'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_frommajCharmajArray';

// Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_fromJust']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_fromCharArray']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray
function majData_majString_majNonmajEmpty_majCodemajUnits_frommajNonmajEmptymajCharmajArray($v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_frommajNonmajEmptymajCharmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_frommajNonmajEmptymajCharmajArray';

// Data_String_NonEmpty_CodeUnits_dropWhile
function majData_majString_majNonmajEmpty_majCodemajUnits_dropmajWhile($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_dropmajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_fromString']))(function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(\Data\String\CodeUnits\majData_majString_majCodemajUnits_countmajPrefix($f_0, $v_1), $v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_dropWhile'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_dropmajWhile';

// Data_String_NonEmpty_CodeUnits_dropRight
function majData_majString_majNonmajEmpty_majCodemajUnits_dropmajRight(int $i_0, $nes_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_dropmajRight';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((($GLOBALS['Data_String_NonEmpty_CodeUnits_greaterThanOrEq'])($i_0))(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($nes_1))) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Just(\Data\String\CodeUnits\majData_majString_majCodemajUnits_take((\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($nes_1) - $i_0), $nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_dropRight'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_dropmajRight';

// Data_String_NonEmpty_CodeUnits_drop
function majData_majString_majNonmajEmpty_majCodemajUnits_drop(int $i_0, $nes_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_drop';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((($GLOBALS['Data_String_NonEmpty_CodeUnits_greaterThanOrEq'])($i_0))(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($nes_1))) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Just(\Data\String\CodeUnits\majData_majString_majCodemajUnits_drop($i_0, $nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_drop'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_drop';

// Data_String_NonEmpty_CodeUnits_countPrefix_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_countPrefix_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS']))($GLOBALS['Data_String_CodeUnits_countPrefix']);

// Data_String_NonEmpty_CodeUnits_countPrefix
function majData_majString_majNonmajEmpty_majCodemajUnits_countmajPrefix($v_0, $v_1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_countmajPrefix';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_countPrefix_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_countPrefix'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_countmajPrefix';

// Data_String_NonEmpty_CodeUnits_cons
function majData_majString_majNonmajEmpty_majCodemajUnits_cons($c_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_cons';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(\Data\String\CodeUnits\majData_majString_majCodemajUnits_singleton($c_0)))($s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_cons'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_cons';

// Data_String_NonEmpty_CodeUnits_charAt_closure
$GLOBALS['Data_String_NonEmpty_CodeUnits_charAt_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS']))($GLOBALS['Data_String_CodeUnits_charAt']);

// Data_String_NonEmpty_CodeUnits_charAt
function majData_majString_majNonmajEmpty_majCodemajUnits_charmajAt(int $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majCodemajUnits_charmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_CodeUnits_charAt_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_CodeUnits_charAt'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majCodemajUnits_charmajAt';

