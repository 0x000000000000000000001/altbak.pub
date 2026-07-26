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
\PhpursThunks::$thunks['Data_String_Regex_Flags_guard'] = function() { $v = (($GLOBALS['Control_Alternative_guard'] ?? \PhpursThunks::eval('Control_Alternative_guard')))(($GLOBALS['Control_Alternative_alternativeArray'] ?? \PhpursThunks::eval('Control_Alternative_alternativeArray'))); return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_eq'] = function() { $v = (($GLOBALS['Data_Eq_eqArrayImpl'] ?? \PhpursThunks::eval('Data_Eq_eqArrayImpl')))(($GLOBALS['Data_Eq_eqStringImpl'] ?? \PhpursThunks::eval('Data_Eq_eqStringImpl'))); return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_RegexFlags'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_unicode'] = function() { $v = ["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => true]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_sticky'] = function() { $v = ["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => true, "unicode" => false]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_showRegexFlags'] = function() { $v = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $usedFlags_1_0 = ((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))(((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))(((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))(((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))(((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))(((($GLOBALS['Data_Semigroup_concatArray'] ?? \PhpursThunks::eval('Data_Semigroup_concatArray')))([]))(((($GLOBALS['Data_Functor_arrayMap'] ?? \PhpursThunks::eval('Data_Functor_arrayMap')))(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = "global";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_Regex_Flags_guard'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_guard')))(($v_0)['global'])))))(((($GLOBALS['Data_Functor_arrayMap'] ?? \PhpursThunks::eval('Data_Functor_arrayMap')))(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = "ignoreCase";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_Regex_Flags_guard'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_guard')))(($v_0)['ignoreCase'])))))(((($GLOBALS['Data_Functor_arrayMap'] ?? \PhpursThunks::eval('Data_Functor_arrayMap')))(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = "multiline";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_Regex_Flags_guard'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_guard')))(($v_0)['multiline'])))))(((($GLOBALS['Data_Functor_arrayMap'] ?? \PhpursThunks::eval('Data_Functor_arrayMap')))(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = "dotAll";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_Regex_Flags_guard'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_guard')))(($v_0)['dotAll'])))))(((($GLOBALS['Data_Functor_arrayMap'] ?? \PhpursThunks::eval('Data_Functor_arrayMap')))(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = "sticky";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_Regex_Flags_guard'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_guard')))(($v_0)['sticky'])))))(((($GLOBALS['Data_Functor_arrayMap'] ?? \PhpursThunks::eval('Data_Functor_arrayMap')))(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = "unicode";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_Regex_Flags_guard'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_guard')))(($v_0)['unicode'])));
  $__t1 = null;;
  if (((($GLOBALS['Data_String_Regex_Flags_eq'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_eq')))($usedFlags_1_0))([])) {
$__t1 = "noFlags";
goto end_branch_1;;
};
  $__t1 = (("(" . ((($GLOBALS['Data_String_Common_joinWith'] ?? \PhpursThunks::eval('Data_String_Common_joinWith')))(" <> "))($usedFlags_1_0)) . ")");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_semigroupRegexFlags'] = function() { $v = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["global" => (($v_0)['global'] || ($v1_1)['global']), "ignoreCase" => (($v_0)['ignoreCase'] || ($v1_1)['ignoreCase']), "multiline" => (($v_0)['multiline'] || ($v1_1)['multiline']), "dotAll" => (($v_0)['dotAll'] || ($v1_1)['dotAll']), "sticky" => (($v_0)['sticky'] || ($v1_1)['sticky']), "unicode" => (($v_0)['unicode'] || ($v1_1)['unicode'])];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_noFlags'] = function() { $v = ["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => false]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_newtypeRegexFlags'] = function() { $v = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_multiline'] = function() { $v = ["global" => false, "ignoreCase" => false, "multiline" => true, "dotAll" => false, "sticky" => false, "unicode" => false]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_monoidRegexFlags'] = function() { $v = ["mempty" => ($GLOBALS['Data_String_Regex_Flags_noFlags'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_noFlags')), "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_String_Regex_Flags_semigroupRegexFlags'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_semigroupRegexFlags'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_ignoreCase'] = function() { $v = ["global" => false, "ignoreCase" => true, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => false]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_global'] = function() { $v = ["global" => true, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => false]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_eqRegexFlags'] = function() { $v = ["eq" => ((((((($GLOBALS['Data_Eq_eqRowCons'] ?? \PhpursThunks::eval('Data_Eq_eqRowCons')))(((((($GLOBALS['Data_Eq_eqRowCons'] ?? \PhpursThunks::eval('Data_Eq_eqRowCons')))(((((($GLOBALS['Data_Eq_eqRowCons'] ?? \PhpursThunks::eval('Data_Eq_eqRowCons')))(((((($GLOBALS['Data_Eq_eqRowCons'] ?? \PhpursThunks::eval('Data_Eq_eqRowCons')))(((((($GLOBALS['Data_Eq_eqRowCons'] ?? \PhpursThunks::eval('Data_Eq_eqRowCons')))(((((($GLOBALS['Data_Eq_eqRowCons'] ?? \PhpursThunks::eval('Data_Eq_eqRowCons')))(($GLOBALS['Data_Eq_eqRowNil'] ?? \PhpursThunks::eval('Data_Eq_eqRowNil'))))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))(["reflectSymbol" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = "unicode";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(($GLOBALS['Data_Eq_eqBoolean'] ?? \PhpursThunks::eval('Data_Eq_eqBoolean')))))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))(["reflectSymbol" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = "sticky";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(($GLOBALS['Data_Eq_eqBoolean'] ?? \PhpursThunks::eval('Data_Eq_eqBoolean')))))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))(["reflectSymbol" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = "multiline";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(($GLOBALS['Data_Eq_eqBoolean'] ?? \PhpursThunks::eval('Data_Eq_eqBoolean')))))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))(["reflectSymbol" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = "ignoreCase";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(($GLOBALS['Data_Eq_eqBoolean'] ?? \PhpursThunks::eval('Data_Eq_eqBoolean')))))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))(["reflectSymbol" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = "global";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(($GLOBALS['Data_Eq_eqBoolean'] ?? \PhpursThunks::eval('Data_Eq_eqBoolean')))))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))(["reflectSymbol" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = "dotAll";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(($GLOBALS['Data_Eq_eqBoolean'] ?? \PhpursThunks::eval('Data_Eq_eqBoolean'))))['eqRecord'])(new Phpurs_Data0("Proxy"))]; return $v; };
\PhpursThunks::$thunks['Data_String_Regex_Flags_dotAll'] = function() { $v = ["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => true, "sticky" => false, "unicode" => false]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };

















