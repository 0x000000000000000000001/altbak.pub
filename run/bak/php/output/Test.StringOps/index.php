<?php

namespace Test\StringOps;

// ALL IMPORTS: Data.Array, Data.Either, Data.Function, Data.Maybe, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.String, Data.String.CodePoints, Data.String.Common, Data.String.Pattern, Data.String.Regex, Data.String.Regex.Flags, Effect, Effect.Console, Partial.Unsafe, Prelude, Prim, Test.StringOps
// TO REQUIRE: Data.Array, Data.Either, Data.Function, Data.Maybe, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.String, Data.String.CodePoints, Data.String.Common, Data.String.Pattern, Data.String.Regex, Data.String.Regex.Flags, Effect, Effect.Console, Partial.Unsafe, Prelude, Test.StringOps
require_once __DIR__ . '/../Data.Array/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.String/index.php';
require_once __DIR__ . '/../Data.String.CodePoints/index.php';
require_once __DIR__ . '/../Data.String.Common/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
require_once __DIR__ . '/../Data.String.Regex/index.php';
require_once __DIR__ . '/../Data.String.Regex.Flags/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.StringOps/index.php';

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
\PhpursThunks::$thunks['Test_StringOps_regexPattern'] = function() { $v = (($GLOBALS['Partial_Unsafe__unsafePartial'] ?? \PhpursThunks::eval('Partial_Unsafe__unsafePartial')))(function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_1_0 = ((((($GLOBALS['Data_String_Regex_regexImpl'] ?? \PhpursThunks::eval('Data_String_Regex_regexImpl')))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left'))))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))("(hello|world)[0-9]+"))((($GLOBALS['Data_String_Regex_renderFlags'] ?? \PhpursThunks::eval('Data_String_Regex_renderFlags')))(($GLOBALS['Data_String_Regex_Flags_noFlags'] ?? \PhpursThunks::eval('Data_String_Regex_Flags_noFlags'))));
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Right"))) {
$__t1 = ($__local_var_1_0)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Test_StringOps_runStringOps'] = function() { $v = function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_StringOps_runStringOps"), recVars=[];
  $loop_1_0 = null;
  $loop_1_0 = (function() use (&$loop_1_0) {
  $__fn = function($v_2, $v1_3 = null, $v2_4 = null) use (&$loop_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "loop_1_0"), recVars=["loop_1_0"];
  while (true) {
switch ($v_2) {
case 0:
$__t5 = $v2_4;
break;
default:
$concatted_5_1 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))($v1_3))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))($v_2)))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("world"))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v_2))(1)))));
$__tco_2 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_2))(1);
$__tco_3 = ((($GLOBALS['Data_String_CodePoints_take'] ?? \PhpursThunks::eval('Data_String_CodePoints_take')))(10))($concatted_5_1);
$__tco_4 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v2_4))((($GLOBALS['Data_Array_length'] ?? \PhpursThunks::eval('Data_Array_length')))(((($GLOBALS['Data_String_Common_split'] ?? \PhpursThunks::eval('Data_String_Common_split')))("e"))((((($GLOBALS['Data_String_Regex_replace'] ?? \PhpursThunks::eval('Data_String_Regex_replace')))(($GLOBALS['Test_StringOps_regexPattern'] ?? \PhpursThunks::eval('Test_StringOps_regexPattern'))))("matched"))($concatted_5_1))));
$v_2 = $__tco_2;
$v1_3 = $__tco_3;
$v2_4 = $__tco_4;
continue 2;
$__t5 = null;
break;
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($loop_1_0)($n_0))("hello"))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_StringOps_describe'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))("String Operations (1k Regex/Split):"); return $v; };
\PhpursThunks::$thunks['Test_StringOps_act'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))((($GLOBALS['Test_StringOps_runStringOps'] ?? \PhpursThunks::eval('Test_StringOps_runStringOps')))(1000))); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






