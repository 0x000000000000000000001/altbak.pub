<?php

namespace Data\String\CodeUnits;

// ALL IMPORTS: Control.Semigroupoid, Data.Eq, Data.Maybe, Data.Ring, Data.Semiring, Data.String.CodeUnits, Data.String.Pattern, Data.String.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Eq, Data.Maybe, Data.Ring, Data.Semiring, Data.String.CodeUnits, Data.String.Pattern, Data.String.Unsafe, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.String.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
require_once __DIR__ . '/../Data.String.Unsafe/index.php';
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
\PhpursThunks::$thunks['Data_String_CodeUnits_uncons'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  switch ($v_0) {
case "":
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
break;
default:
;
break;
};
  $__t0 = new Phpurs_Data1("Just", ["head" => ((($GLOBALS['Data_String_Unsafe_charAt'] ?? \PhpursThunks::eval('Data_String_Unsafe_charAt')))(0))($v_0), "tail" => ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))(1))($v_0)]);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_toChar'] = function() { $v = ((($GLOBALS['Data_String_CodeUnits__toChar'] ?? \PhpursThunks::eval('Data_String_CodeUnits__toChar')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_takeWhile'] = function() { $v = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))(((($GLOBALS['Data_String_CodeUnits_countPrefix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_countPrefix')))($p_0))($s_1)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_takeRight'] = function() { $v = (function() {
  $__fn = function($i_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))(((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($s_1) - $i_0)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_stripSuffix'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $str_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = ((($GLOBALS['Data_String_CodeUnits_splitAt'] ?? \PhpursThunks::eval('Data_String_CodeUnits_splitAt')))(((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($str_1) - (($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($v_0))))($str_1);
  $__t1 = null;;
  if ((($v1_2_0)['after'] === $v_0)) {
$__t1 = new Phpurs_Data1("Just", ($v1_2_0)['before']);
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_stripPrefix'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $str_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = ((($GLOBALS['Data_String_CodeUnits_splitAt'] ?? \PhpursThunks::eval('Data_String_CodeUnits_splitAt')))((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($v_0)))($str_1);
  $__t1 = null;;
  if ((($v1_2_0)['before'] === $v_0)) {
$__t1 = new Phpurs_Data1("Just", ($v1_2_0)['after']);
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_startsWith'] = function() { $v = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Maybe_isJust'] ?? \PhpursThunks::eval('Data_Maybe_isJust'))))((($GLOBALS['Data_String_CodeUnits_stripPrefix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_stripPrefix')))($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_lastIndexOf__prime__'] = function() { $v = ((($GLOBALS['Data_String_CodeUnits__lastIndexOfStartingAt'] ?? \PhpursThunks::eval('Data_String_CodeUnits__lastIndexOfStartingAt')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_lastIndexOf'] = function() { $v = ((($GLOBALS['Data_String_CodeUnits__lastIndexOf'] ?? \PhpursThunks::eval('Data_String_CodeUnits__lastIndexOf')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_indexOf__prime__'] = function() { $v = ((($GLOBALS['Data_String_CodeUnits__indexOfStartingAt'] ?? \PhpursThunks::eval('Data_String_CodeUnits__indexOfStartingAt')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_indexOf'] = function() { $v = ((($GLOBALS['Data_String_CodeUnits__indexOf'] ?? \PhpursThunks::eval('Data_String_CodeUnits__indexOf')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_endsWith'] = function() { $v = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Maybe_isJust'] ?? \PhpursThunks::eval('Data_Maybe_isJust'))))((($GLOBALS['Data_String_CodeUnits_stripSuffix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_stripSuffix')))($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_dropWhile'] = function() { $v = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_String_CodeUnits_drop'] ?? \PhpursThunks::eval('Data_String_CodeUnits_drop')))(((($GLOBALS['Data_String_CodeUnits_countPrefix'] ?? \PhpursThunks::eval('Data_String_CodeUnits_countPrefix')))($p_0))($s_1)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_dropRight'] = function() { $v = (function() {
  $__fn = function($i_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_String_CodeUnits_take'] ?? \PhpursThunks::eval('Data_String_CodeUnits_take')))(((($GLOBALS['Data_String_CodeUnits_length'] ?? \PhpursThunks::eval('Data_String_CodeUnits_length')))($s_1) - $i_0)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_contains'] = function() { $v = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Maybe_isJust'] ?? \PhpursThunks::eval('Data_Maybe_isJust'))))((($GLOBALS['Data_String_CodeUnits_indexOf'] ?? \PhpursThunks::eval('Data_String_CodeUnits_indexOf')))($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_CodeUnits_charAt'] = function() { $v = ((($GLOBALS['Data_String_CodeUnits__charAt'] ?? \PhpursThunks::eval('Data_String_CodeUnits__charAt')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_String_CodeUnits = \call_user_func(function() {
  $exports = [];
$fromCharArray = function($a) use (&$fromCharArray) {
    return implode("", $a);
};

$toCharArray = function($s) use (&$toCharArray) {
    if ($s === "") return [];
    return str_split($s);
};

$singleton = function($c) use (&$singleton) {
    return $c;
};

$_charAt = function($just, $nothing = null, $i = null, $s = null) use (&$_charAt) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_charAt) {

            return $_charAt(...\array_merge($__args, $more));
        };
    }
    return ($i >= 0 && $i < strlen($s)) ? $just($s[$i]) : $nothing;
};

$_toChar = function($just, $nothing = null, $s = null) use (&$_toChar) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_toChar) {

            return $_toChar(...\array_merge($__args, $more));
        };
    }
    return strlen($s) === 1 ? $just($s) : $nothing;
};

$length = function($s) use (&$length) {
    return strlen($s);
};

$countPrefix = function($p, $s = null) use (&$countPrefix) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$countPrefix) {

            return $countPrefix(...\array_merge($__args, $more));
        };
    }
    $i = 0;
    $len = strlen($s);
    while ($i < $len && $p($s[$i])) {
        $i++;
    }
    return $i;
};

$_indexOf = function($just, $nothing = null, $x = null, $s = null) use (&$_indexOf) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_indexOf) {

            return $_indexOf(...\array_merge($__args, $more));
        };
    }
    $i = strpos($s, $x);
    return $i === false ? $nothing : $just($i);
};

$_indexOfStartingAt = function($just, $nothing = null, $x = null, $startAt = null, $s = null) use (&$_indexOfStartingAt) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_indexOfStartingAt) {

            return $_indexOfStartingAt(...\array_merge($__args, $more));
        };
    }
    if ($startAt < 0 || $startAt > strlen($s)) return $nothing;
    $i = strpos($s, $x, $startAt);
    return $i === false ? $nothing : $just($i);
};

$_lastIndexOf = function($just, $nothing = null, $x = null, $s = null) use (&$_lastIndexOf) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_lastIndexOf) {

            return $_lastIndexOf(...\array_merge($__args, $more));
        };
    }
    if ($x === "") {
        return $just(strlen($s));
    }
    $i = strrpos($s, $x);
    return $i === false ? $nothing : $just($i);
};

$_lastIndexOfStartingAt = function($just, $nothing = null, $x = null, $startAt = null, $s = null) use (&$_lastIndexOfStartingAt) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_lastIndexOfStartingAt) {

            return $_lastIndexOfStartingAt(...\array_merge($__args, $more));
        };
    }
    if ($x === "") return $just(\min($startAt, strlen($s)));
    if ($startAt < 0) return $nothing;
    if ($startAt > strlen($s)) $startAt = strlen($s);
    $i = strrpos(substr($s, 0, $startAt + strlen($x)), $x);
    // JS lastIndexOf searches backwards from startAt. PHP strrpos searches the whole string up to offset, or with negative offset.
    // Equivalent logic:
    $sub = substr($s, 0, $startAt + strlen($x));
    $pos = strrpos($sub, $x);
    if ($pos !== false && $pos <= $startAt) {
        return $just($pos);
    }
    return $nothing;
};

$take = function($n, $s = null) use (&$take) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$take) {

            return $take(...\array_merge($__args, $more));
        };
    }
    return substr($s, 0, $n);
};

$drop = function($n, $s = null) use (&$drop) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$drop) {

            return $drop(...\array_merge($__args, $more));
        };
    }
    return substr($s, $n);
};

$slice = function($b, $e = null, $s = null) use (&$slice) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$slice) {

            return $slice(...\array_merge($__args, $more));
        };
    }
    // JS slice with negative indices
    $len = strlen($s);
    if ($b < 0) $b = \max($len + $b, 0);
    else $b = \min($b, $len);
    if ($e < 0) $e = \max($len + $e, 0);
    else $e = \min($e, $len);
    if ($e <= $b) return "";
    return substr($s, $b, $e - $b);
};

$splitAt = function($i, $s = null) use (&$splitAt) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$splitAt) {

            return $splitAt(...\array_merge($__args, $more));
        };
    }
    return (object)[
        "before" => substr($s, 0, $i),
        "after" => substr($s, $i)
    ];
};

$exports['fromCharArray'] = $fromCharArray;
$exports['toCharArray'] = $toCharArray;
$exports['singleton'] = $singleton;
$exports['_charAt'] = $_charAt;
$exports['_toChar'] = $_toChar;
$exports['length'] = $length;
$exports['countPrefix'] = $countPrefix;
$exports['_indexOf'] = $_indexOf;
$exports['_indexOfStartingAt'] = $_indexOfStartingAt;
$exports['_lastIndexOf'] = $_lastIndexOf;
$exports['_lastIndexOfStartingAt'] = $_lastIndexOfStartingAt;
$exports['take'] = $take;
$exports['drop'] = $drop;
$exports['slice'] = $slice;
$exports['splitAt'] = $splitAt;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_String_CodeUnits__charAt'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['_charAt']; };
\PhpursThunks::$thunks['Data_String_CodeUnits__indexOf'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['_indexOf']; };
\PhpursThunks::$thunks['Data_String_CodeUnits__indexOfStartingAt'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['_indexOfStartingAt']; };
\PhpursThunks::$thunks['Data_String_CodeUnits__lastIndexOf'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['_lastIndexOf']; };
\PhpursThunks::$thunks['Data_String_CodeUnits__lastIndexOfStartingAt'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['_lastIndexOfStartingAt']; };
\PhpursThunks::$thunks['Data_String_CodeUnits__toChar'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['_toChar']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_countPrefix'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['countPrefix']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_drop'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['drop']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_fromCharArray'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['fromCharArray']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_length'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['length']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_singleton'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['singleton']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_slice'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['slice']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_splitAt'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['splitAt']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_take'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['take']; };
\PhpursThunks::$thunks['Data_String_CodeUnits_toCharArray'] = function() use (&$ffi_Data_String_CodeUnits) { return $ffi_Data_String_CodeUnits['toCharArray']; };


















