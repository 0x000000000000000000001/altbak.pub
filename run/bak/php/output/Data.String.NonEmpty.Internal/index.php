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
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_fromJust'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_fromJust"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Just"))) {
$__t0 = ($v_0)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_NonEmptyString'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_NonEmptyString"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_NonEmptyReplacement'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_NonEmptyReplacement"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_toUpper'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_toUpper"), recVars=[];
  $__res = (($GLOBALS['Data_String_Common_toUpper'] ?? \PhpursThunks::eval('Data_String_Common_toUpper')))($v_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_toString'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_toString"), recVars=[];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_toLower'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_toLower"), recVars=[];
  $__res = (($GLOBALS['Data_String_Common_toLower'] ?? \PhpursThunks::eval('Data_String_Common_toLower')))($v_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_showNonEmptyString'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (("(NonEmptyString.unsafeFromString " . (($GLOBALS['Data_Show_showStringImpl'] ?? \PhpursThunks::eval('Data_Show_showStringImpl')))($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_showNonEmptyReplacement'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (("(NonEmptyReplacement (NonEmptyString.unsafeFromString " . (($GLOBALS['Data_Show_showStringImpl'] ?? \PhpursThunks::eval('Data_Show_showStringImpl')))($v_0)) . "))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_semigroupNonEmptyString'] = function() { $v = ($GLOBALS['Data_Semigroup_semigroupString'] ?? \PhpursThunks::eval('Data_Semigroup_semigroupString')); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement'] = function() { $v = ($GLOBALS['Data_Semigroup_semigroupString'] ?? \PhpursThunks::eval('Data_Semigroup_semigroupString')); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_replaceAll'] = function() { $v = (function() {
  $__fn = function($pat_0, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_replaceAll"), recVars=[];
  $__res = (((($GLOBALS['Data_String_Common_replaceAll'] ?? \PhpursThunks::eval('Data_String_Common_replaceAll')))($pat_0))($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_replace'] = function() { $v = (function() {
  $__fn = function($pat_0, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_replace"), recVars=[];
  $__res = (((($GLOBALS['Data_String_Common_replace'] ?? \PhpursThunks::eval('Data_String_Common_replace')))($pat_0))($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_prependString'] = function() { $v = (function() {
  $__fn = function($s1_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_prependString"), recVars=[];
  $__res = ($s1_0 . $v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_ordNonEmptyString'] = function() { $v = ($GLOBALS['Data_Ord_ordString'] ?? \PhpursThunks::eval('Data_Ord_ordString')); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_ordNonEmptyReplacement'] = function() { $v = ($GLOBALS['Data_Ord_ordString'] ?? \PhpursThunks::eval('Data_Ord_ordString')); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_nonEmptyNonEmpty'] = function() { $v = function($dictIsSymbol_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_nonEmptyNonEmpty"), recVars=[];
  $__res = (object)["nes" => function($p_1) use ($dictIsSymbol_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictIsSymbol_0)->reflectSymbol)($p_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_nes'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_nes"), recVars=[];
  $__res = ($dict_0)->nes;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_makeNonEmptyBad'] = function() { $v = function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_makeNonEmptyBad"), recVars=[];
  $__res = (object)["nes" => function($v_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = "";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_localeCompare'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_localeCompare"), recVars=[];
  $__res = ((($GLOBALS['Data_String_Common_localeCompare'] ?? \PhpursThunks::eval('Data_String_Common_localeCompare')))($v_0))($v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_liftS'] = function() { $v = (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_liftS"), recVars=[];
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_startsWith'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_liftS'))))(($GLOBALS['Data_String_CodeUnits_startsWith'] ?? \PhpursThunks::eval('Data_String_CodeUnits_startsWith'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_joinWith1'] = function() { $v = function($dictFoldable1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_joinWith1"), recVars=[];
  $intercalate_1_0 = ((($GLOBALS['Data_Foldable_intercalate'] ?? \PhpursThunks::eval('Data_Foldable_intercalate')))((($dictFoldable1_0)->Foldable0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))))(($GLOBALS['Data_Monoid_monoidString'] ?? \PhpursThunks::eval('Data_Monoid_monoidString')));
  $__res = function($v_2) use ($intercalate_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_NonEmptyString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_NonEmptyString'))))(($intercalate_1_0)($v_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_joinWith'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_joinWith"), recVars=[];
  $intercalate_1_0 = ((($GLOBALS['Data_Foldable_intercalate'] ?? \PhpursThunks::eval('Data_Foldable_intercalate')))($dictFoldable_0))(($GLOBALS['Data_Monoid_monoidString'] ?? \PhpursThunks::eval('Data_Monoid_monoidString')));
  $__res = function($splice_2) use ($intercalate_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($intercalate_1_0)($splice_2)))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_join1With'] = function() { $v = function($dictFoldable1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_join1With"), recVars=[];
  $joinWith2_1_0 = (($GLOBALS['Data_String_NonEmpty_Internal_joinWith'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_joinWith')))((($dictFoldable1_0)->Foldable0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($splice_2) use ($joinWith2_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_NonEmptyString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_NonEmptyString'))))(($joinWith2_1_0)($splice_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_fromString'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_fromString"), recVars=[];
  switch ($v_0) {
case "":
$__t0 = new Phpurs_Data0("Nothing");
break;
default:
$__t0 = new Phpurs_Data1("Just", $v_0);
break;
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_stripPrefix'] = function() { $v = (function() {
  $__fn = function($pat_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_stripPrefix"), recVars=[];
  $__local_var_2_0 = ((($GLOBALS['Data_String_CodeUnits_stripPrefix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_stripPrefix')))($pat_0))($a_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
switch (($__local_var_2_0)->value0) {
case "":
$__t2 = new Phpurs_Data0("Nothing");
break;
default:
$__t2 = new Phpurs_Data1("Just", ($__local_var_2_0)->value0);
break;
};
$__t1 = $__t2;
} else {
if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_stripSuffix'] = function() { $v = (function() {
  $__fn = function($pat_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_stripSuffix"), recVars=[];
  $__local_var_2_0 = ((($GLOBALS['Data_String_CodeUnits_stripSuffix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_stripSuffix')))($pat_0))($a_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
switch (($__local_var_2_0)->value0) {
case "":
$__t2 = new Phpurs_Data0("Nothing");
break;
default:
$__t2 = new Phpurs_Data1("Just", ($__local_var_2_0)->value0);
break;
};
$__t1 = $__t2;
} else {
if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_trim'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_trim"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_String_Common_trim'] ?? \PhpursThunks::eval('Data_String_Common_trim')))($v_0);
  switch ($__local_var_1_0) {
case "":
$__t1 = new Phpurs_Data0("Nothing");
break;
default:
$__t1 = new Phpurs_Data1("Just", $__local_var_1_0);
break;
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_unsafeFromString'] = function() { $v = function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_unsafeFromString"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_fromJust'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_fromJust'))))(($GLOBALS['Data_String_NonEmpty_Internal_fromString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_fromString')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_eqNonEmptyString'] = function() { $v = ($GLOBALS['Data_Eq_eqString'] ?? \PhpursThunks::eval('Data_Eq_eqString')); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_eqNonEmptyReplacement'] = function() { $v = ($GLOBALS['Data_Eq_eqString'] ?? \PhpursThunks::eval('Data_Eq_eqString')); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_endsWith'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_liftS'))))(($GLOBALS['Data_String_CodeUnits_endsWith'] ?? \PhpursThunks::eval('Data_String_CodeUnits_endsWith'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_contains'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_liftS'))))(($GLOBALS['Data_String_CodeUnits_contains'] ?? \PhpursThunks::eval('Data_String_CodeUnits_contains'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_Internal_appendString'] = function() { $v = (function() {
  $__fn = function($v_0, $s2_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_String_NonEmpty_Internal_appendString"), recVars=[];
  $__res = ($v_0 . $s2_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




































