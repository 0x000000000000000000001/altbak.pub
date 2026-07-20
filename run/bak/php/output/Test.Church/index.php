<?php

namespace Test\Church;

// ALL IMPORTS: Data.Function, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Prim, Test.Church
// TO REQUIRE: Data.Function, Data.Ring, Data.Semiring, Data.Show, Effect, Effect.Console, Prelude, Test.Church
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.Church/index.php';

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
\PhpursThunks::$thunks['Test_Church_zeroC'] = function() { $v = (function() {
  $__fn = function($v_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_zeroC"), recVars=[];
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_Church_toInt'] = function() { $v = function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_toInt"), recVars=[];
  $__res = (($n_0)(function($x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($x_1))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_Church_succC'] = function() { $v = (function() {
  $__fn = function($n_0, $f_1 = null, $x_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_succC"), recVars=[];
  $__res = ($f_1)((($n_0)($f_1))($x_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_Church_mulC'] = function() { $v = (function() {
  $__fn = function($m_0, $n_1 = null, $f_2 = null, $x_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_mulC"), recVars=[];
  $__res = (($m_0)(($n_1)($f_2)))($x_3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_Church_fromInt'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_fromInt"), recVars=["Test_Church_fromInt"];
  while (true) {
switch ($v_0) {
case 0:
$__t1 = ($GLOBALS['Test_Church_zeroC'] ?? \PhpursThunks::eval('Test_Church_zeroC'));
break;
default:
$__local_var_1_0 = (($GLOBALS['Test_Church_fromInt'] ?? \PhpursThunks::eval('Test_Church_fromInt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1));
$__t1 = (function() use ($__local_var_1_0) {
  $__fn = function($f_2, $x_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Test_Church_fromInt"];
  $__res = ($f_2)((($__local_var_1_0)($f_2))($x_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
break;
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_Church_describe'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))("Church Numerals (100k Closure Applications):"); return $v; };
\PhpursThunks::$thunks['Test_Church_c10'] = function() { $v = (($GLOBALS['Test_Church_fromInt'] ?? \PhpursThunks::eval('Test_Church_fromInt')))(10); return $v; };
\PhpursThunks::$thunks['Test_Church_c100'] = function() { $v = (function() {
  $__fn = function($f_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_c100"), recVars=[];
  $__res = ((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))($f_0)))($x_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_Church_c10k'] = function() { $v = (function() {
  $__fn = function($f_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_c10k"), recVars=[];
  $__res = ((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))((($GLOBALS['Test_Church_c100'] ?? \PhpursThunks::eval('Test_Church_c100')))($f_0))))($x_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_Church_c100k'] = function() { $v = (function() {
  $__fn = function($f_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_c100k"), recVars=[];
  $__res = ((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))((($GLOBALS['Test_Church_c100'] ?? \PhpursThunks::eval('Test_Church_c100')))((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))($f_0)))))($x_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_Church_addC'] = function() { $v = (function() {
  $__fn = function($m_0, $n_1 = null, $f_2 = null, $x_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_Church_addC"), recVars=[];
  $__res = (($m_0)($f_2))((($n_1)($f_2))($x_3));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Test_Church_act'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))((($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))((($GLOBALS['Test_Church_c100'] ?? \PhpursThunks::eval('Test_Church_c100')))((($GLOBALS['Test_Church_c10'] ?? \PhpursThunks::eval('Test_Church_c10')))(function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($x_0))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))))(0))); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };














