<?php

namespace Data\String\NonEmpty\Internal;

// ALL IMPORTS: Control.Bind, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Maybe, Data.Monoid, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.Show, Data.String, Data.String.CodeUnits, Data.String.Common, Data.String.NonEmpty.Internal, Data.String.Pattern, Data.Symbol, Prelude, Prim, Prim.TypeError, Type.Proxy, Unsafe.Coerce
// TO REQUIRE: Control.Bind, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Maybe, Data.Monoid, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.Show, Data.String, Data.String.CodeUnits, Data.String.Common, Data.String.NonEmpty.Internal, Data.String.Pattern, Data.Symbol, Prelude, Type.Proxy, Unsafe.Coerce
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.String/index.php';
require_once __DIR__ . '/../Data.String.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.Common/index.php';
require_once __DIR__ . '/../Data.String.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Type.Proxy/index.php';
require_once __DIR__ . '/../Unsafe.Coerce/index.php';

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




// Data_String_NonEmpty_Internal_fromJust
function majData_majString_majNonmajEmpty_majInternal_frommajJust($v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_frommajJust';
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
$GLOBALS['Data_String_NonEmpty_Internal_fromJust'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_frommajJust';

// Data_String_NonEmpty_Internal_NonEmptyString
function majData_majString_majNonmajEmpty_majInternal_majNonmajEmptymajString($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_majNonmajEmptymajString';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_NonEmptyString'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_majNonmajEmptymajString';

// Data_String_NonEmpty_Internal_NonEmptyReplacement
function majData_majString_majNonmajEmpty_majInternal_majNonmajEmptymajReplacement($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_majNonmajEmptymajReplacement';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_NonEmptyReplacement'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_majNonmajEmptymajReplacement';

// Data_String_NonEmpty_Internal_toUpper
function majData_majString_majNonmajEmpty_majInternal_tomajUpper(string $v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_tomajUpper';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Data\String\Common\majData_majString_majCommon_tomajUpper($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_toUpper'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_tomajUpper';

// Data_String_NonEmpty_Internal_toString
function majData_majString_majNonmajEmpty_majInternal_tomajString(string $v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_tomajString';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_toString'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_tomajString';

// Data_String_NonEmpty_Internal_toLower
function majData_majString_majNonmajEmpty_majInternal_tomajLower(string $v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_tomajLower';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Data\String\Common\majData_majString_majCommon_tomajLower($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_toLower'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_tomajLower';

// Data_String_NonEmpty_Internal_showNonEmptyString
$GLOBALS['Data_String_NonEmpty_Internal_showNonEmptyString'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(NonEmptyString.unsafeFromString "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Show_showString'])->{'show'})($v_0)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_NonEmpty_Internal_showNonEmptyReplacement
$GLOBALS['Data_String_NonEmpty_Internal_showNonEmptyReplacement'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(NonEmptyReplacement "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_String_NonEmpty_Internal_showNonEmptyString'])->{'show'})($v_0)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_NonEmpty_Internal_semigroupNonEmptyString
$GLOBALS['Data_String_NonEmpty_Internal_semigroupNonEmptyString'] = $GLOBALS['Data_Semigroup_semigroupString'];

// Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement
$GLOBALS['Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement'] = $GLOBALS['Data_Semigroup_semigroupString'];

// Data_String_NonEmpty_Internal_replaceAll
function majData_majString_majNonmajEmpty_majInternal_replacemajAll(string $pat_0, $v_1 = null, $v1_2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_replacemajAll';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Data\String\Common\majData_majString_majCommon_replacemajAll($pat_0, $v_1, $v1_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_replaceAll'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_replacemajAll';

// Data_String_NonEmpty_Internal_replace
function majData_majString_majNonmajEmpty_majInternal_replace(string $pat_0, $v_1 = null, $v1_2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_replace';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Data\String\Common\majData_majString_majCommon_replace($pat_0, $v_1, $v1_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_replace'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_replace';

// Data_String_NonEmpty_Internal_prependString
function majData_majString_majNonmajEmpty_majInternal_prependmajString(string $s1_0, $v_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_prependmajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($s1_0))($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_prependString'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_prependmajString';

// Data_String_NonEmpty_Internal_ordNonEmptyString
$GLOBALS['Data_String_NonEmpty_Internal_ordNonEmptyString'] = $GLOBALS['Data_Ord_ordString'];

// Data_String_NonEmpty_Internal_ordNonEmptyReplacement
$GLOBALS['Data_String_NonEmpty_Internal_ordNonEmptyReplacement'] = $GLOBALS['Data_Ord_ordString'];

// Data_String_NonEmpty_Internal_nonEmptyNonEmpty
function majData_majString_majNonmajEmpty_majInternal_nonmajEmptymajNonmajEmpty($dictIsSymbol_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_nonmajEmptymajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["nes" => function($p_1) use ($dictIsSymbol_0) {
  $__num = \func_num_args();
  $__res = (($dictIsSymbol_0)->{'reflectSymbol'})($p_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_nonEmptyNonEmpty'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_nonmajEmptymajNonmajEmpty';

// Data_String_NonEmpty_Internal_nes
function majData_majString_majNonmajEmpty_majInternal_nes($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_nes';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'nes'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_nes'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_nes';

// Data_String_NonEmpty_Internal_makeNonEmptyBad
function majData_majString_majNonmajEmpty_majInternal_makemajNonmajEmptymajBad($_dollar__unused_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_makemajNonmajEmptymajBad';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["nes" => function($v_1) {
  $__num = \func_num_args();
  $__res = "";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_makeNonEmptyBad'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_makemajNonmajEmptymajBad';

// Data_String_NonEmpty_Internal_localeCompare
function majData_majString_majNonmajEmpty_majInternal_localemajCompare(string $v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_localemajCompare';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Common\majData_majString_majCommon_localemajCompare($v_0, $v1_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_localeCompare'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_localemajCompare';

// Data_String_NonEmpty_Internal_liftS
function majData_majString_majNonmajEmpty_majInternal_liftmajS($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_liftmajS';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_liftS'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_liftmajS';

// Data_String_NonEmpty_Internal_startsWith_closure
$GLOBALS['Data_String_NonEmpty_Internal_startsWith_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_liftS']))($GLOBALS['Data_String_CodeUnits_startsWith']);

// Data_String_NonEmpty_Internal_startsWith
function majData_majString_majNonmajEmpty_majInternal_startsmajWith(string $v_0, $v_1 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_startsmajWith';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_Internal_startsWith_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_startsWith'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_startsmajWith';

// Data_String_NonEmpty_Internal_joinWith1
function majData_majString_majNonmajEmpty_majInternal_joinmajWith1($dictFoldable1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_joinmajWith1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFoldable1_0)->{'Foldable0'})(null);
  $__local_var_2_1 = (($GLOBALS['Data_Monoid_monoidString'])->{'Semigroup0'})(null);
  $__res = function($v_3) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_NonEmptyString']))(function($xs_4) use ($__local_var_1_0, $__local_var_2_1, $v_3) {
  $__num = \func_num_args();
  $__res = ((((($__local_var_1_0)->{'foldl'})((function() use ($__local_var_2_1, $v_3) {
  $__fn = function($v_5, $v1_6 = null) use ($__local_var_2_1, $v_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if (($v_5)->{'init'}) {
$__t2 = (object)["init" => false, "acc" => $v1_6];
goto end_branch_2;;
};
  $__t2 = (object)["init" => false, "acc" => ((($__local_var_2_1)->{'append'})(($v_5)->{'acc'}))(((($__local_var_2_1)->{'append'})($v_3))($v1_6))];
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["init" => true, "acc" => ($GLOBALS['Data_Monoid_monoidString'])->{'mempty'}]))($xs_4))->{'acc'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_joinWith1'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_joinmajWith1';

// Data_String_NonEmpty_Internal_joinWith
function majData_majString_majNonmajEmpty_majInternal_joinmajWith($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_joinmajWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($GLOBALS['Data_Monoid_monoidString'])->{'Semigroup0'})(null);
  $__res = function($splice_2) use ($__local_var_1_0, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_3) use ($__local_var_1_0, $dictFoldable_0, $splice_2) {
  $__num = \func_num_args();
  $__res = ((((($dictFoldable_0)->{'foldl'})((function() use ($__local_var_1_0, $splice_2) {
  $__fn = function($v_4, $v1_5 = null) use ($__local_var_1_0, $splice_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if (($v_4)->{'init'}) {
$__t1 = (object)["init" => false, "acc" => $v1_5];
goto end_branch_1;;
};
  $__t1 = (object)["init" => false, "acc" => ((($__local_var_1_0)->{'append'})(($v_4)->{'acc'}))(((($__local_var_1_0)->{'append'})($splice_2))($v1_5))];
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["init" => true, "acc" => ($GLOBALS['Data_Monoid_monoidString'])->{'mempty'}]))($xs_3))->{'acc'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Unsafe_Coerce_unsafeCoerce']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_joinWith'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_joinmajWith';

// Data_String_NonEmpty_Internal_join1With
function majData_majString_majNonmajEmpty_majInternal_join1majWith($dictFoldable1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_join1majWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $joinWith2_1_0 = ($GLOBALS['Data_String_NonEmpty_Internal_joinWith'])((($dictFoldable1_0)->{'Foldable0'})(null));
  $__res = function($splice_2) use ($joinWith2_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_NonEmptyString']))(($joinWith2_1_0)($splice_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_join1With'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_join1majWith';

// Data_String_NonEmpty_Internal_fromString
function majData_majString_majNonmajEmpty_majInternal_frommajString(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_frommajString';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = match ($v_0) { "" => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just($v_0) };
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_fromString'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_frommajString';

// Data_String_NonEmpty_Internal_stripPrefix
function majData_majString_majNonmajEmpty_majInternal_stripmajPrefix(string $pat_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_stripmajPrefix';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_bindMaybe'])->{'bind'})(\Data\String\CodeUnits\majData_majString_majCodemajUnits_stripmajPrefix($pat_0, $a_1)))($GLOBALS['Data_String_NonEmpty_Internal_fromString']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_stripPrefix'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_stripmajPrefix';

// Data_String_NonEmpty_Internal_stripSuffix
function majData_majString_majNonmajEmpty_majInternal_stripmajSuffix(string $pat_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_stripmajSuffix';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_bindMaybe'])->{'bind'})(\Data\String\CodeUnits\majData_majString_majCodemajUnits_stripmajSuffix($pat_0, $a_1)))($GLOBALS['Data_String_NonEmpty_Internal_fromString']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_stripSuffix'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_stripmajSuffix';

// Data_String_NonEmpty_Internal_trim
function majData_majString_majNonmajEmpty_majInternal_trim(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_trim';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = \Data\String\Common\majData_majString_majCommon_trim($v_0);
  $__res = match ($__local_var_1_0) { "" => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just($__local_var_1_0) };
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_trim'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_trim';

// Data_String_NonEmpty_Internal_unsafeFromString
function majData_majString_majNonmajEmpty_majInternal_unsafemajFrommajString($_dollar__unused_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_unsafemajFrommajString';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_fromJust']))($GLOBALS['Data_String_NonEmpty_Internal_fromString']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_unsafeFromString'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_unsafemajFrommajString';

// Data_String_NonEmpty_Internal_eqNonEmptyString
$GLOBALS['Data_String_NonEmpty_Internal_eqNonEmptyString'] = $GLOBALS['Data_Eq_eqString'];

// Data_String_NonEmpty_Internal_eqNonEmptyReplacement
$GLOBALS['Data_String_NonEmpty_Internal_eqNonEmptyReplacement'] = $GLOBALS['Data_Eq_eqString'];

// Data_String_NonEmpty_Internal_endsWith_closure
$GLOBALS['Data_String_NonEmpty_Internal_endsWith_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_liftS']))($GLOBALS['Data_String_CodeUnits_endsWith']);

// Data_String_NonEmpty_Internal_endsWith
function majData_majString_majNonmajEmpty_majInternal_endsmajWith(string $v_0, $v_1 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_endsmajWith';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_Internal_endsWith_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_endsWith'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_endsmajWith';

// Data_String_NonEmpty_Internal_contains_closure
$GLOBALS['Data_String_NonEmpty_Internal_contains_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_NonEmpty_Internal_liftS']))($GLOBALS['Data_String_CodeUnits_contains']);

// Data_String_NonEmpty_Internal_contains
function majData_majString_majNonmajEmpty_majInternal_contains(string $v_0, $v_1 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_contains';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_NonEmpty_Internal_contains_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_contains'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_contains';

// Data_String_NonEmpty_Internal_appendString
function majData_majString_majNonmajEmpty_majInternal_appendmajString(string $v_0, $s2_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majNonmajEmpty_majInternal_appendmajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($v_0))($s2_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_NonEmpty_Internal_appendString'] = __NAMESPACE__ . '\\majData_majString_majNonmajEmpty_majInternal_appendmajString';

