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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_String_Regex_Flags_guard
$GLOBALS['Data_String_Regex_Flags_guard'] = ($GLOBALS['Control_Alternative_guard'])($GLOBALS['Control_Alternative_alternativeArray']);

// Data_String_Regex_Flags_eq
$GLOBALS['Data_String_Regex_Flags_eq'] = ($GLOBALS['Data_Eq_eqArrayImpl'])(($GLOBALS['Data_Eq_eqString'])->{'eq'});

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
  $usedFlags_1_0 = ((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})([]))(((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})(((($GLOBALS['Data_Functor_functorArray'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = "global";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_String_Regex_Flags_guard'])(($v_0)->{'global'}))))(((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})(((($GLOBALS['Data_Functor_functorArray'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = "ignoreCase";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_String_Regex_Flags_guard'])(($v_0)->{'ignoreCase'}))))(((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})(((($GLOBALS['Data_Functor_functorArray'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = "multiline";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_String_Regex_Flags_guard'])(($v_0)->{'multiline'}))))(((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})(((($GLOBALS['Data_Functor_functorArray'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = "dotAll";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_String_Regex_Flags_guard'])(($v_0)->{'dotAll'}))))(((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})(((($GLOBALS['Data_Functor_functorArray'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = "sticky";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_String_Regex_Flags_guard'])(($v_0)->{'sticky'}))))(((($GLOBALS['Data_Functor_functorArray'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = "unicode";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_String_Regex_Flags_guard'])(($v_0)->{'unicode'}))))))));
  $__t1 = null;;
  if ((($GLOBALS['Data_String_Regex_Flags_eq'])($usedFlags_1_0))([])) {
$__t1 = "noFlags";
goto end_branch_1;;
};
  $__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("("))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(\Data\String\Common\majData_majString_majCommon_joinmajWith(" <> ", $usedFlags_1_0)))(")"));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_Regex_Flags_semigroupRegexFlags
$GLOBALS['Data_String_Regex_Flags_semigroupRegexFlags'] = (object)["append" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["global" => ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($v_0)->{'global'}))(($v1_1)->{'global'}), "ignoreCase" => ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($v_0)->{'ignoreCase'}))(($v1_1)->{'ignoreCase'}), "multiline" => ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($v_0)->{'multiline'}))(($v1_1)->{'multiline'}), "dotAll" => ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($v_0)->{'dotAll'}))(($v1_1)->{'dotAll'}), "sticky" => ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($v_0)->{'sticky'}))(($v1_1)->{'sticky'}), "unicode" => ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($v_0)->{'unicode'}))(($v1_1)->{'unicode'})];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_String_Regex_Flags_noFlags
$GLOBALS['Data_String_Regex_Flags_noFlags'] = (object)["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => false, "sticky" => false, "unicode" => false];

// Data_String_Regex_Flags_newtypeRegexFlags
$GLOBALS['Data_String_Regex_Flags_newtypeRegexFlags'] = (object)["Coercible0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_Regex_Flags_multiline
$GLOBALS['Data_String_Regex_Flags_multiline'] = (object)["global" => false, "ignoreCase" => false, "multiline" => true, "dotAll" => false, "sticky" => false, "unicode" => false];

// Data_String_Regex_Flags_monoidRegexFlags
$GLOBALS['Data_String_Regex_Flags_monoidRegexFlags'] = (object)["mempty" => $GLOBALS['Data_String_Regex_Flags_noFlags'], "Semigroup0" => function($_dollar__unused_0) {
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
$GLOBALS['Data_String_Regex_Flags_eqRegexFlags'] = (object)["eq" => (function() {
  $__fn = function($ra_0, $rb_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(($ra_0)->{'dotAll'}))(($rb_1)->{'dotAll'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(($ra_0)->{'global'}))(($rb_1)->{'global'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(($ra_0)->{'ignoreCase'}))(($rb_1)->{'ignoreCase'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(($ra_0)->{'multiline'}))(($rb_1)->{'multiline'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(($ra_0)->{'sticky'}))(($rb_1)->{'sticky'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(($ra_0)->{'unicode'}))(($rb_1)->{'unicode'})))((((($GLOBALS['Data_Eq_eqRowNil'])->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_0))($rb_1)))))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_String_Regex_Flags_dotAll
$GLOBALS['Data_String_Regex_Flags_dotAll'] = (object)["global" => false, "ignoreCase" => false, "multiline" => false, "dotAll" => true, "sticky" => false, "unicode" => false];

