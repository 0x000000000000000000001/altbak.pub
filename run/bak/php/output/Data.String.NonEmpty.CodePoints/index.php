<?php

namespace Data\String\NonEmpty\CodePoints;

// ALL IMPORTS: Control.Semigroupoid, Data.Array.NonEmpty, Data.Maybe, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.String.CodePoints, Data.String.NonEmpty.CodePoints, Data.String.NonEmpty.Internal, Data.String.Pattern, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Array.NonEmpty, Data.Maybe, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.String.CodePoints, Data.String.NonEmpty.CodePoints, Data.String.NonEmpty.Internal, Data.String.Pattern, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.String.CodePoints/index.php';
require_once __DIR__ . '/../Data.String.NonEmpty.CodePoints/index.php';
require_once __DIR__ . '/../Data.String.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
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


// Data_String_NonEmpty_CodePoints_fromJust
$GLOBALS['Data_String_NonEmpty_CodePoints_fromJust'] = function($v_0 = null) {
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
};

// Data_String_NonEmpty_CodePoints_snoc
$GLOBALS['Data_String_NonEmpty_CodePoints_snoc'] = (function() {
  $__fn = function($c_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])($s_1))(($GLOBALS['Data_String_CodePoints_singleton'])($c_0));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_NonEmpty_CodePoints_singleton
$GLOBALS['Data_String_NonEmpty_CodePoints_singleton'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_Internal_NonEmptyString']))($GLOBALS['Data_String_CodePoints_singleton']);

// Data_String_NonEmpty_CodePoints_liftS
$GLOBALS['Data_String_NonEmpty_CodePoints_liftS'] = (function() {
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
})();

// Data_String_NonEmpty_CodePoints_takeWhile
$GLOBALS['Data_String_NonEmpty_CodePoints_takeWhile'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_Internal_fromString']))(function($v_1 = null) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_String_CodePoints_take'])((($GLOBALS['Data_String_CodePoints_countPrefix'])($f_0))($v_1)))($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_NonEmpty_CodePoints_lastIndexOf'
$GLOBALS['Data_String_NonEmpty_CodePoints_lastIndexOf__prime__'] = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_liftS']))(($GLOBALS['Data_String_CodePoints_lastIndexOf__prime__'])($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_NonEmpty_CodePoints_lastIndexOf
$GLOBALS['Data_String_NonEmpty_CodePoints_lastIndexOf'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_liftS']))($GLOBALS['Data_String_CodePoints_lastIndexOf']);

// Data_String_NonEmpty_CodePoints_indexOf'
$GLOBALS['Data_String_NonEmpty_CodePoints_indexOf__prime__'] = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_liftS']))(($GLOBALS['Data_String_CodePoints_indexOf__prime__'])($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_NonEmpty_CodePoints_indexOf
$GLOBALS['Data_String_NonEmpty_CodePoints_indexOf'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_liftS']))($GLOBALS['Data_String_CodePoints_indexOf']);

// Data_String_NonEmpty_CodePoints_fromNonEmptyString
$GLOBALS['Data_String_NonEmpty_CodePoints_fromNonEmptyString'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_NonEmpty_CodePoints_length
$GLOBALS['Data_String_NonEmpty_CodePoints_length'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_CodePoints_length']))($GLOBALS['Data_String_NonEmpty_CodePoints_fromNonEmptyString']);

// Data_String_NonEmpty_CodePoints_splitAt
$GLOBALS['Data_String_NonEmpty_CodePoints_splitAt'] = (function() {
  $__fn = function($i_0 = null, $nes_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = (($GLOBALS['Data_String_CodePoints_splitAt'])($i_0))($nes_1);
  $__res = ["before" => match (($v_2_0)['before']) { "" => new Phpurs_Data0("Nothing"), default => new Phpurs_Data1("Just", ($v_2_0)['before']) }, "after" => match (($v_2_0)['after']) { "" => new Phpurs_Data0("Nothing"), default => new Phpurs_Data1("Just", ($v_2_0)['after']) }];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_NonEmpty_CodePoints_take
$GLOBALS['Data_String_NonEmpty_CodePoints_take'] = (function() {
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
  $__t0 = new Phpurs_Data1("Just", (($GLOBALS['Data_String_CodePoints_take'])($i_0))($nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_NonEmpty_CodePoints_toCodePointArray
$GLOBALS['Data_String_NonEmpty_CodePoints_toCodePointArray'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_CodePoints_toCodePointArray']))($GLOBALS['Data_String_NonEmpty_CodePoints_fromNonEmptyString']);

// Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray
$GLOBALS['Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_fromJust']))(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_Array_NonEmpty_fromArray']))($GLOBALS['Data_String_NonEmpty_CodePoints_toCodePointArray']));

// Data_String_NonEmpty_CodePoints_uncons
$GLOBALS['Data_String_NonEmpty_CodePoints_uncons'] = function($nes_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Data_String_CodePoints_codePointAt'])(0))($nes_0);
  $__t1 = null;;
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Just"))) {
$__t1 = ($__local_var_1_0)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__local_var_1_2 = (($GLOBALS['Data_String_CodeUnits_drop'])(($GLOBALS['Data_String_CodeUnits_length'])((($GLOBALS['Data_String_CodePoints_take'])(1))($nes_0))))($nes_0);
  $__res = ["head" => $__t1, "tail" => match ($__local_var_1_2) { "" => new Phpurs_Data0("Nothing"), default => new Phpurs_Data1("Just", $__local_var_1_2) }];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_NonEmpty_CodePoints_fromFoldable1
$GLOBALS['Data_String_NonEmpty_CodePoints_fromFoldable1'] = function($dictFoldable1_0 = null) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable1_0)['foldMap1'])($GLOBALS['Data_Semigroup_semigroupString']))($GLOBALS['Data_String_NonEmpty_CodePoints_singleton']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_NonEmpty_CodePoints_fromCodePointArray
$GLOBALS['Data_String_NonEmpty_CodePoints_fromCodePointArray'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = match (count($v_0)) { 0 => new Phpurs_Data0("Nothing"), default => new Phpurs_Data1("Just", ($GLOBALS['Data_String_CodePoints_fromCodePointArray'])($v_0)) };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray
$GLOBALS['Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_fromJust']))(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_fromCodePointArray']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_String_NonEmpty_CodePoints_dropWhile
$GLOBALS['Data_String_NonEmpty_CodePoints_dropWhile'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_Internal_fromString']))(function($v_1 = null) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_String_CodePoints_dropWhile'])($f_0))($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_NonEmpty_CodePoints_drop
$GLOBALS['Data_String_NonEmpty_CodePoints_drop'] = (function() {
  $__fn = function($i_0 = null, $nes_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($i_0 >= ($GLOBALS['Data_String_CodePoints_length'])($nes_1))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data1("Just", (($GLOBALS['Data_String_CodeUnits_drop'])(($GLOBALS['Data_String_CodeUnits_length'])((($GLOBALS['Data_String_CodePoints_take'])($i_0))($nes_1))))($nes_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_NonEmpty_CodePoints_countPrefix
$GLOBALS['Data_String_NonEmpty_CodePoints_countPrefix'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_liftS']))($GLOBALS['Data_String_CodePoints_countPrefix']);

// Data_String_NonEmpty_CodePoints_cons
$GLOBALS['Data_String_NonEmpty_CodePoints_cons'] = (function() {
  $__fn = function($c_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(($GLOBALS['Data_String_CodePoints_singleton'])($c_0)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_NonEmpty_CodePoints_codePointAt
$GLOBALS['Data_String_NonEmpty_CodePoints_codePointAt'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_NonEmpty_CodePoints_liftS']))($GLOBALS['Data_String_CodePoints_codePointAt']);

