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
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_fromJust'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Just"))) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_snoc'] = function() { $v = (function() {
  $__fn = function($c_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($s_1 . (($GLOBALS['Data_String_CodeUnits_singleton'] ?? \PhpursThunks::eval('Data_String_CodeUnits_singleton')))($c_0));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_singleton'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_NonEmptyString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_NonEmptyString'))))(($GLOBALS['Data_String_CodeUnits_singleton'] ?? \PhpursThunks::eval('Data_String_CodeUnits_singleton'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_liftS'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_takeWhile'] = function() { $v = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_fromString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_fromString'))))(function($v_1 = null) use ($f_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))(((($GLOBALS['Data_String_CodeUnits_countPrefix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_countPrefix')))($f_0))($v_1)))($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_lastIndexOf__prime__'] = function() { $v = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_liftS'))))((($GLOBALS['Data_String_CodeUnits_lastIndexOf__prime__'] ?? \PhpursThunks::eval('Data_String_CodeUnits_lastIndexOf__prime__')))($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_lastIndexOf'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_liftS'))))(($GLOBALS['Data_String_CodeUnits_lastIndexOf'] ?? \PhpursThunks::eval('Data_String_CodeUnits_lastIndexOf'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_indexOf__prime__'] = function() { $v = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_liftS'))))((($GLOBALS['Data_String_CodeUnits_indexOf__prime__'] ?? \PhpursThunks::eval('Data_String_CodeUnits_indexOf__prime__')))($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_indexOf'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_liftS'))))(($GLOBALS['Data_String_CodeUnits_indexOf'] ?? \PhpursThunks::eval('Data_String_CodeUnits_indexOf'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_fromNonEmptyString'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_length'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length'))))(($GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_fromNonEmptyString'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_splitAt'] = function() { $v = (function() {
  $__fn = function($i_0 = null, $nes_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = ((($GLOBALS['Data_String_CodeUnits_splitAt'] ?? \PhpursThunks::eval('Data_String_CodeUnits_splitAt')))($i_0))($nes_1);
  $__t1 = null;;
  switch (($v_2_0)['before']) {
case "":
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
break;
default:
;
break;
};
  $__t1 = new Phpurs_Data1("Just", ($v_2_0)['before']);
  end_branch_1:;
  $__t2 = null;;
  switch (($v_2_0)['after']) {
case "":
$__t2 = new Phpurs_Data0("Nothing");
goto end_branch_2;;
break;
default:
;
break;
};
  $__t2 = new Phpurs_Data1("Just", ($v_2_0)['after']);
  end_branch_2:;
  $__res = ["before" => $__t1, "after" => $__t2];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_take'] = function() { $v = (function() {
  $__fn = function($i_0 = null, $nes_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($i_0 < 1)) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data1("Just", ((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))($i_0))($nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_takeRight'] = function() { $v = (function() {
  $__fn = function($i_0 = null, $nes_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($i_0 < 1)) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data1("Just", ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))(((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($nes_1) - $i_0)))($nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_toChar'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_CodeUnits_toChar'] ?? \PhpursThunks::eval('Data_String_CodeUnits_toChar'))))(($GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_fromNonEmptyString'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_toCharArray'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_CodeUnits_toCharArray'] ?? \PhpursThunks::eval('Data_String_CodeUnits_toCharArray'))))(($GLOBALS['Data_String_NonEmpty_CodeUnits_fromNonEmptyString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_fromNonEmptyString'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_fromJust'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_fromJust'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Array_NonEmpty_fromArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_fromArray'))))(($GLOBALS['Data_String_NonEmpty_CodeUnits_toCharArray'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_toCharArray')))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_uncons'] = function() { $v = function($nes_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))(1))($nes_0);
  $__t1 = null;;
  switch ($__local_var_1_0) {
case "":
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
break;
default:
;
break;
};
  $__t1 = new Phpurs_Data1("Just", $__local_var_1_0);
  end_branch_1:;
  $__res = ["head" => ((($GLOBALS['Data_String_Unsafe_charAt'] ?? \PhpursThunks::eval('Data_String_Unsafe_charAt')))(0))($nes_0), "tail" => $__t1];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_fromFoldable1'] = function() { $v = function($dictFoldable1_0 = null) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable1_0)['foldMap1'])(($GLOBALS['Data_Semigroup_semigroupString'] ?? \PhpursThunks::eval('Data_Semigroup_semigroupString'))))(($GLOBALS['Data_String_NonEmpty_CodeUnits_singleton'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_singleton')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_fromCharArray'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  switch (count($v_0)) {
case 0:
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
break;
default:
;
break;
};
  $__t0 = new Phpurs_Data1("Just", (($GLOBALS['Data_String_CodeUnits_fromCharArray'] ?? \PhpursThunks::eval('Data_String_CodeUnits_fromCharArray')))($v_0));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_fromJust'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_fromJust'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_fromCharArray'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_fromCharArray'))))(($GLOBALS['Data_Array_NonEmpty_toArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_toArray')))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_dropWhile'] = function() { $v = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_Internal_fromString'] ?? \PhpursThunks::eval('Data_String_NonEmpty_Internal_fromString'))))(function($v_1 = null) use ($f_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))(((($GLOBALS['Data_String_CodeUnits_countPrefix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_countPrefix')))($f_0))($v_1)))($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_dropRight'] = function() { $v = (function() {
  $__fn = function($i_0 = null, $nes_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($i_0 >= (($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($nes_1))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data1("Just", ((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))(((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($nes_1) - $i_0)))($nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_drop'] = function() { $v = (function() {
  $__fn = function($i_0 = null, $nes_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($i_0 >= (($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($nes_1))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data1("Just", ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))($i_0))($nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_countPrefix'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_liftS'))))(($GLOBALS['Data_String_CodeUnits_countPrefix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_countPrefix'))); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_cons'] = function() { $v = (function() {
  $__fn = function($c_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_String_CodeUnits_singleton'] ?? \PhpursThunks::eval('Data_String_CodeUnits_singleton')))($c_0) . $s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_NonEmpty_CodeUnits_charAt'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_String_NonEmpty_CodeUnits_liftS'] ?? \PhpursThunks::eval('Data_String_NonEmpty_CodeUnits_liftS'))))(($GLOBALS['Data_String_CodeUnits_charAt'] ?? \PhpursThunks::eval('Data_String_CodeUnits_charAt'))); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };





























