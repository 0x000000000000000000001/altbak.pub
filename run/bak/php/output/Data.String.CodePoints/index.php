<?php

namespace Data\String\CodePoints;

// ALL IMPORTS: Control.Semigroupoid, Data.Array, Data.Boolean, Data.Bounded, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Ord, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.String.CodePoints, Data.String.CodeUnits, Data.String.Common, Data.String.Pattern, Data.String.Unsafe, Data.Tuple, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Array, Data.Boolean, Data.Bounded, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Ord, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.String.CodePoints, Data.String.CodeUnits, Data.String.Common, Data.String.Pattern, Data.String.Unsafe, Data.Tuple, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Int/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.String.CodePoints/index.php';
require_once __DIR__ . '/../Data.String.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.Common/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
require_once __DIR__ . '/../Data.String.Unsafe/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
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
\PhpursThunks::$thunks['Data_String_CodePoints_CodePoint'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_CodePoint"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_unsurrogate'] = function() { $v = (function() {
  $__fn = function($lead_0, $trail_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_unsurrogate"), recVars=[];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(((($GLOBALS['Data_Semiring_intMul'] ?? \PhpursThunks::eval('Data_Semiring_intMul')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($lead_0))(55296)))(1024)))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($trail_1))(56320))))(65536);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_showCodePoint'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(CodePoint 0x"))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_String_Common_toUpper'] ?? \PhpursThunks::eval('Data_String_Common_toUpper')))(((($GLOBALS['Data_Int_toStringAs'] ?? \PhpursThunks::eval('Data_Int_toStringAs')))(16))($v_0))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_isTrail'] = function() { $v = function($cu_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_isTrail"), recVars=[];
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(56320))($cu_0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(56320))($cu_0))->tag === "GT")))))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($cu_0))(57343)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($cu_0))(57343))->tag === "GT"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_isLead'] = function() { $v = function($cu_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_isLead"), recVars=[];
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(55296))($cu_0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(55296))($cu_0))->tag === "GT")))))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($cu_0))(56319)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($cu_0))(56319))->tag === "GT"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_uncons'] = function() { $v = function($s_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_uncons"), recVars=[];
  $v_1_0 = (($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($s_0);
  switch ($v_1_0) {
case 0:
$__t4 = new Phpurs_Data0("Nothing");
break;
case 1:
$__t4 = new Phpurs_Data1("Just", (object)["head" => (($GLOBALS['Data_Enum_toCharCode'] ?? \PhpursThunks::eval('Data_Enum_toCharCode')))(((($GLOBALS['Data_String_Unsafe_charAt'] ?? \PhpursThunks::eval('Data_String_Unsafe_charAt')))(0))($s_0)), "tail" => ""]);
break;
default:
$cu1_2_1 = (($GLOBALS['Data_Enum_toCharCode'] ?? \PhpursThunks::eval('Data_Enum_toCharCode')))(((($GLOBALS['Data_String_Unsafe_charAt'] ?? \PhpursThunks::eval('Data_String_Unsafe_charAt')))(1))($s_0));
$cu0_3_2 = (($GLOBALS['Data_Enum_toCharCode'] ?? \PhpursThunks::eval('Data_Enum_toCharCode')))(((($GLOBALS['Data_String_Unsafe_charAt'] ?? \PhpursThunks::eval('Data_String_Unsafe_charAt')))(0))($s_0));
if (((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))((($GLOBALS['Data_String_CodePoints_isLead'] ?? \PhpursThunks::eval('Data_String_CodePoints_isLead')))($cu0_3_2)))((($GLOBALS['Data_String_CodePoints_isTrail'] ?? \PhpursThunks::eval('Data_String_CodePoints_isTrail')))($cu1_2_1))) {
$__t3 = new Phpurs_Data1("Just", (object)["head" => ((($GLOBALS['Data_String_CodePoints_unsurrogate'] ?? \PhpursThunks::eval('Data_String_CodePoints_unsurrogate')))($cu0_3_2))($cu1_2_1), "tail" => ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))(2))($s_0)]);
} else {
$__t3 = new Phpurs_Data1("Just", (object)["head" => $cu0_3_2, "tail" => ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))(1))($s_0)]);
};
$__t4 = $__t3;
break;
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_unconsButWithTuple'] = function() { $v = function($s_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_unconsButWithTuple"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_String_CodePoints_uncons'] ?? \PhpursThunks::eval('Data_String_CodePoints_uncons')))($s_0);
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", (($__local_var_1_0)->value0)->head, (($__local_var_1_0)->value0)->tail));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_toCodePointArrayFallback'] = function() { $v = function($s_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_toCodePointArrayFallback"), recVars=[];
  $__res = (((($GLOBALS['Data_Unfoldable_unfoldableArray'] ?? \PhpursThunks::eval('Data_Unfoldable_unfoldableArray')))->unfoldr)(($GLOBALS['Data_String_CodePoints_unconsButWithTuple'] ?? \PhpursThunks::eval('Data_String_CodePoints_unconsButWithTuple'))))($s_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_unsafeCodePointAt0Fallback'] = function() { $v = function($s_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_unsafeCodePointAt0Fallback"), recVars=[];
  $cu0_1_0 = (($GLOBALS['Data_Enum_toCharCode'] ?? \PhpursThunks::eval('Data_Enum_toCharCode')))(((($GLOBALS['Data_String_Unsafe_charAt'] ?? \PhpursThunks::eval('Data_String_Unsafe_charAt')))(0))($s_0));
  if (((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))((($GLOBALS['Data_String_CodePoints_isLead'] ?? \PhpursThunks::eval('Data_String_CodePoints_isLead')))($cu0_1_0)))((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($s_0)))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($s_0)))(1))->tag === "GT")))) {
$cu1_2_2 = (($GLOBALS['Data_Enum_toCharCode'] ?? \PhpursThunks::eval('Data_Enum_toCharCode')))(((($GLOBALS['Data_String_Unsafe_charAt'] ?? \PhpursThunks::eval('Data_String_Unsafe_charAt')))(1))($s_0));
if ((($GLOBALS['Data_String_CodePoints_isTrail'] ?? \PhpursThunks::eval('Data_String_CodePoints_isTrail')))($cu1_2_2)) {
$__t3 = ((($GLOBALS['Data_String_CodePoints_unsurrogate'] ?? \PhpursThunks::eval('Data_String_CodePoints_unsurrogate')))($cu0_1_0))($cu1_2_2);
} else {
$__t3 = $cu0_1_0;
};
$__t1 = $__t3;
} else {
$__t1 = $cu0_1_0;
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_unsafeCodePointAt0'] = function() { $v = (($GLOBALS['Data_String_CodePoints__unsafeCodePointAt0'] ?? \PhpursThunks::eval('Data_String_CodePoints__unsafeCodePointAt0')))(($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0Fallback'] ?? \PhpursThunks::eval('Data_String_CodePoints_unsafeCodePointAt0Fallback'))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_toCodePointArray'] = function() { $v = ((($GLOBALS['Data_String_CodePoints__toCodePointArray'] ?? \PhpursThunks::eval('Data_String_CodePoints__toCodePointArray')))(($GLOBALS['Data_String_CodePoints_toCodePointArrayFallback'] ?? \PhpursThunks::eval('Data_String_CodePoints_toCodePointArrayFallback'))))(($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0'] ?? \PhpursThunks::eval('Data_String_CodePoints_unsafeCodePointAt0'))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_length'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length'))))(($GLOBALS['Data_String_CodePoints_toCodePointArray'] ?? \PhpursThunks::eval('Data_String_CodePoints_toCodePointArray'))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_lastIndexOf'] = function() { $v = (function() {
  $__fn = function($p_0, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_lastIndexOf"), recVars=[];
  $__local_var_2_0 = ((($GLOBALS['Data_String_CodeUnits_lastIndexOf'] ?? \PhpursThunks::eval('Data_String_CodeUnits_lastIndexOf')))($p_0))($s_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($GLOBALS['Data_String_CodePoints_length'] ?? \PhpursThunks::eval('Data_String_CodePoints_length')))(((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))(($__local_var_2_0)->value0))($s_1)));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_indexOf'] = function() { $v = (function() {
  $__fn = function($p_0, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_indexOf"), recVars=[];
  $__local_var_2_0 = ((($GLOBALS['Data_String_CodeUnits_indexOf'] ?? \PhpursThunks::eval('Data_String_CodeUnits_indexOf')))($p_0))($s_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($GLOBALS['Data_String_CodePoints_length'] ?? \PhpursThunks::eval('Data_String_CodePoints_length')))(((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))(($__local_var_2_0)->value0))($s_1)));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_fromCharCode'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_CodeUnits_singleton'] ?? \PhpursThunks::eval('Data_String_CodeUnits_singleton'))))((((($GLOBALS['Data_Enum_toEnumWithDefaults'] ?? \PhpursThunks::eval('Data_Enum_toEnumWithDefaults')))(($GLOBALS['Data_Enum_boundedEnumChar'] ?? \PhpursThunks::eval('Data_Enum_boundedEnumChar'))))(($GLOBALS['Data_Bounded_bottomChar'] ?? \PhpursThunks::eval('Data_Bounded_bottomChar'))))(($GLOBALS['Data_Bounded_topChar'] ?? \PhpursThunks::eval('Data_Bounded_topChar')))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_singletonFallback'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_singletonFallback"), recVars=[];
  if (( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(65535)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(65535))->tag === "GT")))) {
$__t0 = (($GLOBALS['Data_String_CodePoints_fromCharCode'] ?? \PhpursThunks::eval('Data_String_CodePoints_fromCharCode')))($v_0);
} else {
$__t0 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_String_CodePoints_fromCharCode'] ?? \PhpursThunks::eval('Data_String_CodePoints_fromCharCode')))(((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(((($GLOBALS['Data_EuclideanRing_intDiv'] ?? \PhpursThunks::eval('Data_EuclideanRing_intDiv')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(65536)))(1024)))(55296))))((($GLOBALS['Data_String_CodePoints_fromCharCode'] ?? \PhpursThunks::eval('Data_String_CodePoints_fromCharCode')))(((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(((($GLOBALS['Data_EuclideanRing_intMod'] ?? \PhpursThunks::eval('Data_EuclideanRing_intMod')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(65536)))(1024)))(56320)));
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_fromCodePointArray'] = function() { $v = (($GLOBALS['Data_String_CodePoints__fromCodePointArray'] ?? \PhpursThunks::eval('Data_String_CodePoints__fromCodePointArray')))(($GLOBALS['Data_String_CodePoints_singletonFallback'] ?? \PhpursThunks::eval('Data_String_CodePoints_singletonFallback'))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_singleton'] = function() { $v = (($GLOBALS['Data_String_CodePoints__singleton'] ?? \PhpursThunks::eval('Data_String_CodePoints__singleton')))(($GLOBALS['Data_String_CodePoints_singletonFallback'] ?? \PhpursThunks::eval('Data_String_CodePoints_singletonFallback'))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_takeFallback'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_takeFallback"), recVars=["Data_String_CodePoints_takeFallback"];
  while (true) {
if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(1))->tag === "LT"))) {
$__t2 = "";
} else {
$v2_2_0 = (($GLOBALS['Data_String_CodePoints_uncons'] ?? \PhpursThunks::eval('Data_String_CodePoints_uncons')))($v1_1);
if ((is_object($v2_2_0) && (($v2_2_0)->tag === "Just"))) {
$__t1 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_String_CodePoints_singleton'] ?? \PhpursThunks::eval('Data_String_CodePoints_singleton')))((($v2_2_0)->value0)->head)))(((($GLOBALS['Data_String_CodePoints_takeFallback'] ?? \PhpursThunks::eval('Data_String_CodePoints_takeFallback')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1)))((($v2_2_0)->value0)->tail));
} else {
$__t1 = $v1_1;
};
$__t2 = $__t1;
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_take'] = function() { $v = (($GLOBALS['Data_String_CodePoints__take'] ?? \PhpursThunks::eval('Data_String_CodePoints__take')))(($GLOBALS['Data_String_CodePoints_takeFallback'] ?? \PhpursThunks::eval('Data_String_CodePoints_takeFallback'))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_lastIndexOf__prime__'] = function() { $v = (function() {
  $__fn = function($p_0, $i_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_lastIndexOf'"), recVars=[];
  $__local_var_3_0 = (((($GLOBALS['Data_String_CodeUnits_lastIndexOf__prime__'] ?? \PhpursThunks::eval('Data_String_CodeUnits_lastIndexOf__prime__')))($p_0))((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))(((($GLOBALS['Data_String_CodePoints_take'] ?? \PhpursThunks::eval('Data_String_CodePoints_take')))($i_1))($s_2))))($s_2);
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($GLOBALS['Data_String_CodePoints_length'] ?? \PhpursThunks::eval('Data_String_CodePoints_length')))(((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))(($__local_var_3_0)->value0))($s_2)));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_splitAt'] = function() { $v = (function() {
  $__fn = function($i_0, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_splitAt"), recVars=[];
  $before_2_0 = ((($GLOBALS['Data_String_CodePoints_take'] ?? \PhpursThunks::eval('Data_String_CodePoints_take')))($i_0))($s_1);
  $__res = (object)["before" => $before_2_0, "after" => ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($before_2_0)))($s_1)];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_eqCodePoint'] = function() { $v = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_ordCodePoint'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_String_CodePoints_eqCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_eqCodePoint'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_drop'] = function() { $v = (function() {
  $__fn = function($n_0, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_drop"), recVars=[];
  $__res = ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))(((($GLOBALS['Data_String_CodePoints_take'] ?? \PhpursThunks::eval('Data_String_CodePoints_take')))($n_0))($s_1))))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_indexOf__prime__'] = function() { $v = (function() {
  $__fn = function($p_0, $i_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_indexOf'"), recVars=[];
  $s__prime___3_0 = ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))(((($GLOBALS['Data_String_CodePoints_take'] ?? \PhpursThunks::eval('Data_String_CodePoints_take')))($i_1))($s_2))))($s_2);
  $__local_var_4_1 = ((($GLOBALS['Data_String_CodeUnits_indexOf'] ?? \PhpursThunks::eval('Data_String_CodeUnits_indexOf')))($p_0))($s__prime___3_0);
  if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($i_1))((($GLOBALS['Data_String_CodePoints_length'] ?? \PhpursThunks::eval('Data_String_CodePoints_length')))(((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))(($__local_var_4_1)->value0))($s__prime___3_0))));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_countTail'] = function() { $v = (function() {
  $__fn = function($p_0, $s_1 = null, $accum_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_countTail"), recVars=["Data_String_CodePoints_countTail"];
  while (true) {
$v_3_0 = (($GLOBALS['Data_String_CodePoints_uncons'] ?? \PhpursThunks::eval('Data_String_CodePoints_uncons')))($s_1);
if (((is_object($v_3_0) && (($v_3_0)->tag === "Just")) && ($p_0)((($v_3_0)->value0)->head))) {
$__tco_2 = $p_0;
$__tco_3 = (($v_3_0)->value0)->tail;
$__tco_4 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($accum_2))(1);
$p_0 = $__tco_2;
$s_1 = $__tco_3;
$accum_2 = $__tco_4;
continue ;
$__t1 = null;
} else {
$__t1 = $accum_2;
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_countFallback'] = function() { $v = (function() {
  $__fn = function($p_0, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_countFallback"), recVars=[];
  $__res = (((($GLOBALS['Data_String_CodePoints_countTail'] ?? \PhpursThunks::eval('Data_String_CodePoints_countTail')))($p_0))($s_1))(0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_countPrefix'] = function() { $v = ((($GLOBALS['Data_String_CodePoints__countPrefix'] ?? \PhpursThunks::eval('Data_String_CodePoints__countPrefix')))(($GLOBALS['Data_String_CodePoints_countFallback'] ?? \PhpursThunks::eval('Data_String_CodePoints_countFallback'))))(($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0'] ?? \PhpursThunks::eval('Data_String_CodePoints_unsafeCodePointAt0'))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_dropWhile'] = function() { $v = (function() {
  $__fn = function($p_0, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_dropWhile"), recVars=[];
  $__res = ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))(((($GLOBALS['Data_String_CodePoints_take'] ?? \PhpursThunks::eval('Data_String_CodePoints_take')))(((($GLOBALS['Data_String_CodePoints_countPrefix'] ?? \PhpursThunks::eval('Data_String_CodePoints_countPrefix')))($p_0))($s_1)))($s_1))))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_takeWhile'] = function() { $v = (function() {
  $__fn = function($p_0, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_takeWhile"), recVars=[];
  $__res = ((($GLOBALS['Data_String_CodePoints_take'] ?? \PhpursThunks::eval('Data_String_CodePoints_take')))(((($GLOBALS['Data_String_CodePoints_countPrefix'] ?? \PhpursThunks::eval('Data_String_CodePoints_countPrefix')))($p_0))($s_1)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_codePointFromChar'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_CodePoints_CodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_CodePoint'))))(($GLOBALS['Data_Enum_toCharCode'] ?? \PhpursThunks::eval('Data_Enum_toCharCode'))); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_codePointAtFallback'] = function() { $v = (function() {
  $__fn = function($n_0, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_codePointAtFallback"), recVars=["Data_String_CodePoints_codePointAtFallback"];
  while (true) {
$v_2_0 = (($GLOBALS['Data_String_CodePoints_uncons'] ?? \PhpursThunks::eval('Data_String_CodePoints_uncons')))($s_1);
if ((is_object($v_2_0) && (($v_2_0)->tag === "Just"))) {
if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))($n_0))(0)) {
$__t4 = new Phpurs_Data1("Just", (($v_2_0)->value0)->head);
} else {
$__tco_2 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($n_0))(1);
$__tco_3 = (($v_2_0)->value0)->tail;
$n_0 = $__tco_2;
$s_1 = $__tco_3;
continue ;
$__t4 = null;
};
$__t1 = $__t4;
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_codePointAt'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_CodePoints_codePointAt"), recVars=[];
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(0))->tag === "LT"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
switch ($v_0) {
case 0:
switch ($v1_1) {
case "":
$__t1 = new Phpurs_Data0("Nothing");
break;
default:
$__t1 = new Phpurs_Data1("Just", (($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0'] ?? \PhpursThunks::eval('Data_String_CodePoints_unsafeCodePointAt0')))($v1_1));
break;
};
$__t0 = $__t1;
break;
default:
$__t0 = ((((((($GLOBALS['Data_String_CodePoints__codePointAt'] ?? \PhpursThunks::eval('Data_String_CodePoints__codePointAt')))(($GLOBALS['Data_String_CodePoints_codePointAtFallback'] ?? \PhpursThunks::eval('Data_String_CodePoints_codePointAtFallback'))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")))(($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0'] ?? \PhpursThunks::eval('Data_String_CodePoints_unsafeCodePointAt0'))))($v_0))($v1_1);
break;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_boundedCodePoint'] = function() { $v = (object)["bottom" => 0, "top" => 1114111, "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_String_CodePoints_ordCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_ordCodePoint'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_boundedEnumCodePoint'] = function() { $v = (object)["cardinality" => ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))(1114111))(1), "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_String_CodePoints_boundedEnumCodePoint","Data_String_CodePoints_enumCodePoint"];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "toEnum" => function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_String_CodePoints_boundedEnumCodePoint","Data_String_CodePoints_enumCodePoint"];
  if (((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(0))->tag === "LT")))))(( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(1114111)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(1114111))->tag === "GT"))))) {
$__t0 = new Phpurs_Data1("Just", $n_0);
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_String_CodePoints_boundedEnumCodePoint","Data_String_CodePoints_enumCodePoint"];
  $__res = ($GLOBALS['Data_String_CodePoints_boundedCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_boundedCodePoint'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_String_CodePoints_boundedEnumCodePoint","Data_String_CodePoints_enumCodePoint"];
  $__res = ($GLOBALS['Data_String_CodePoints_enumCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_enumCodePoint'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_CodePoints_enumCodePoint'] = function() { $v = (object)["succ" => ((($GLOBALS['Data_Enum_defaultSucc'] ?? \PhpursThunks::eval('Data_Enum_defaultSucc')))((($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_boundedEnumCodePoint')))->toEnum))((($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_boundedEnumCodePoint')))->fromEnum), "pred" => ((($GLOBALS['Data_Enum_defaultPred'] ?? \PhpursThunks::eval('Data_Enum_defaultPred')))((($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_boundedEnumCodePoint')))->toEnum))((($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_boundedEnumCodePoint')))->fromEnum), "Ord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_String_CodePoints_boundedEnumCodePoint","Data_String_CodePoints_enumCodePoint"];
  $__res = ($GLOBALS['Data_String_CodePoints_ordCodePoint'] ?? \PhpursThunks::eval('Data_String_CodePoints_ordCodePoint'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_String_CodePoints = \call_user_func(function() {
  $exports = [];
if (!\function_exists('Data_String_CodePoints_utf8_ord')) {
    function Data_String_CodePoints_utf8_ord($char) {
        if ($char === '') return 0;
        $c0 = ord($char[0]);
        if ($c0 < 0x80) {
            return $c0;
        } elseif ($c0 < 0xE0) {
            return (($c0 & 0x1F) << 6) | (ord($char[1]) & 0x3F);
        } elseif ($c0 < 0xF0) {
            return (($c0 & 0x0F) << 12) | ((ord($char[1]) & 0x3F) << 6) | (ord($char[2]) & 0x3F);
        } else {
            return (($c0 & 0x07) << 18) | ((ord($char[1]) & 0x3F) << 12) | ((ord($char[2]) & 0x3F) << 6) | (ord($char[3]) & 0x3F);
        }
    }
}

if (!\function_exists('Data_String_CodePoints_utf8_chr')) {
    function Data_String_CodePoints_utf8_chr($code) {
        if ($code < 0x80) {
            return chr($code);
        } elseif ($code < 0x800) {
            return chr(0xC0 | ($code >> 6)) . chr(0x80 | ($code & 0x3F));
        } elseif ($code < 0x10000) {
            return chr(0xE0 | ($code >> 12)) . chr(0x80 | (($code >> 6) & 0x3F)) . chr(0x80 | ($code & 0x3F));
        } else {
            return chr(0xF0 | ($code >> 18)) . chr(0x80 | (($code >> 12) & 0x3F)) . chr(0x80 | (($code >> 6) & 0x3F)) . chr(0x80 | ($code & 0x3F));
        }
    }
}

$_unsafeCodePointAt0 = function($fallback, $str = null) use (&$_unsafeCodePointAt0) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_unsafeCodePointAt0) {

            return $_unsafeCodePointAt0(...\array_merge($__args, $more));
        };
    }
    return Data_String_CodePoints_utf8_ord(iconv_substr($str, 0, 1, 'UTF-8'));
};

$_codePointAt = function($fallback, $just = null, $nothing = null, $unsafeCodePointAt0 = null, $index = null, $str = null) use (&$_codePointAt) {
    if (\func_num_args() < 6) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_codePointAt) {

            return $_codePointAt(...\array_merge($__args, $more));
        };
    }
    $len = iconv_strlen($str, 'UTF-8');
    if ($index < 0 || $index >= $len) return $nothing;
    return $just($unsafeCodePointAt0(iconv_substr($str, $index, 1, 'UTF-8')));
};

$_countPrefix = function($fallback, $unsafeCodePointAt0 = null, $pred = null, $str = null) use (&$_countPrefix) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_countPrefix) {

            return $_countPrefix(...\array_merge($__args, $more));
        };
    }
    $len = iconv_strlen($str, 'UTF-8');
    for ($i = 0; $i < $len; $i++) {
        $char = iconv_substr($str, $i, 1, 'UTF-8');
        $cp = $unsafeCodePointAt0($char);
        if (!$pred($cp)) return $i;
    }
    return $len;
};

$_fromCodePointArray = function($singleton, $cps = null) use (&$_fromCodePointArray) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_fromCodePointArray) {

            return $_fromCodePointArray(...\array_merge($__args, $more));
        };
    }
    $result = "";
    foreach ($cps as $cp) {
        $result .= Data_String_CodePoints_utf8_chr($cp);
    }
    return $result;
};

$_singleton = function($fallback, $cp = null) use (&$_singleton) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_singleton) {

            return $_singleton(...\array_merge($__args, $more));
        };
    }
    return Data_String_CodePoints_utf8_chr($cp);
};

$_take = function($fallback, $n = null, $str = null) use (&$_take) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_take) {

            return $_take(...\array_merge($__args, $more));
        };
    }
    return iconv_substr($str, 0, $n, 'UTF-8');
};

$_toCodePointArray = function($fallback, $unsafeCodePointAt0 = null, $str = null) use (&$_toCodePointArray) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_toCodePointArray) {

            return $_toCodePointArray(...\array_merge($__args, $more));
        };
    }
    $len = iconv_strlen($str, 'UTF-8');
    $arr = [];
    for ($i = 0; $i < $len; $i++) {
        $arr[] = $unsafeCodePointAt0(iconv_substr($str, $i, 1, 'UTF-8'));
    }
    return $arr;
};

$exports['_unsafeCodePointAt0'] = $_unsafeCodePointAt0;
$exports['_codePointAt'] = $_codePointAt;
$exports['_countPrefix'] = $_countPrefix;
$exports['_fromCodePointArray'] = $_fromCodePointArray;
$exports['_singleton'] = $_singleton;
$exports['_take'] = $_take;
$exports['_toCodePointArray'] = $_toCodePointArray;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_String_CodePoints__codePointAt'] = function() use (&$ffi_Data_String_CodePoints) { return $ffi_Data_String_CodePoints['_codePointAt']; };
\PhpursThunks::$thunks['Data_String_CodePoints__countPrefix'] = function() use (&$ffi_Data_String_CodePoints) { return $ffi_Data_String_CodePoints['_countPrefix']; };
\PhpursThunks::$thunks['Data_String_CodePoints__fromCodePointArray'] = function() use (&$ffi_Data_String_CodePoints) { return $ffi_Data_String_CodePoints['_fromCodePointArray']; };
\PhpursThunks::$thunks['Data_String_CodePoints__singleton'] = function() use (&$ffi_Data_String_CodePoints) { return $ffi_Data_String_CodePoints['_singleton']; };
\PhpursThunks::$thunks['Data_String_CodePoints__take'] = function() use (&$ffi_Data_String_CodePoints) { return $ffi_Data_String_CodePoints['_take']; };
\PhpursThunks::$thunks['Data_String_CodePoints__toCodePointArray'] = function() use (&$ffi_Data_String_CodePoints) { return $ffi_Data_String_CodePoints['_toCodePointArray']; };
\PhpursThunks::$thunks['Data_String_CodePoints__unsafeCodePointAt0'] = function() use (&$ffi_Data_String_CodePoints) { return $ffi_Data_String_CodePoints['_unsafeCodePointAt0']; };







































