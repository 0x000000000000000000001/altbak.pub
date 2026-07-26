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
$GLOBALS['Data_String_CodeUnits__charAt'] = $ffi_Data_String_CodeUnits['_charAt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits__indexOf'] = $ffi_Data_String_CodeUnits['_indexOf'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits__indexOfStartingAt'] = $ffi_Data_String_CodeUnits['_indexOfStartingAt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits__lastIndexOf'] = $ffi_Data_String_CodeUnits['_lastIndexOf'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits__lastIndexOfStartingAt'] = $ffi_Data_String_CodeUnits['_lastIndexOfStartingAt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits__toChar'] = $ffi_Data_String_CodeUnits['_toChar'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_countPrefix'] = $ffi_Data_String_CodeUnits['countPrefix'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_drop'] = $ffi_Data_String_CodeUnits['drop'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_fromCharArray'] = $ffi_Data_String_CodeUnits['fromCharArray'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_length'] = $ffi_Data_String_CodeUnits['length'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_singleton'] = $ffi_Data_String_CodeUnits['singleton'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_slice'] = $ffi_Data_String_CodeUnits['slice'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_splitAt'] = $ffi_Data_String_CodeUnits['splitAt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_take'] = $ffi_Data_String_CodeUnits['take'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodeUnits_toCharArray'] = $ffi_Data_String_CodeUnits['toCharArray'] ?? new class { public function __invoke(...$args) { return $this; } };


// Data_String_CodeUnits_uncons
$GLOBALS['Data_String_CodeUnits_uncons'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = match ($v_0) { "" => new Phpurs_Data0("Nothing"), default => new Phpurs_Data1("Just", ["head" => (($GLOBALS['Data_String_Unsafe_charAt'])(0))($v_0), "tail" => (($GLOBALS['Data_String_CodeUnits_drop'])(1))($v_0)]) };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodeUnits_toChar
$GLOBALS['Data_String_CodeUnits_toChar'] = (($GLOBALS['Data_String_CodeUnits__toChar'])($GLOBALS['Data_Maybe_Just']))(new Phpurs_Data0("Nothing"));

// Data_String_CodeUnits_takeWhile
$GLOBALS['Data_String_CodeUnits_takeWhile'] = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_String_CodeUnits_take'])((($GLOBALS['Data_String_CodeUnits_countPrefix'])($p_0))($s_1)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodeUnits_takeRight
$GLOBALS['Data_String_CodeUnits_takeRight'] = (function() {
  $__fn = function($i_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_String_CodeUnits_drop'])((($GLOBALS['Data_String_CodeUnits_length'])($s_1) - $i_0)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodeUnits_stripSuffix
$GLOBALS['Data_String_CodeUnits_stripSuffix'] = (function() {
  $__fn = function($v_0 = null, $str_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = (($GLOBALS['Data_String_CodeUnits_splitAt'])((($GLOBALS['Data_String_CodeUnits_length'])($str_1) - ($GLOBALS['Data_String_CodeUnits_length'])($v_0))))($str_1);
  $__t1 = null;;
  if (((($GLOBALS['Data_Eq_eqString'])['eq'])(($v1_2_0)['after']))($v_0)) {
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
})();

// Data_String_CodeUnits_stripPrefix
$GLOBALS['Data_String_CodeUnits_stripPrefix'] = (function() {
  $__fn = function($v_0 = null, $str_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = (($GLOBALS['Data_String_CodeUnits_splitAt'])(($GLOBALS['Data_String_CodeUnits_length'])($v_0)))($str_1);
  $__t1 = null;;
  if (((($GLOBALS['Data_Eq_eqString'])['eq'])(($v1_2_0)['before']))($v_0)) {
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
})();

// Data_String_CodeUnits_startsWith
$GLOBALS['Data_String_CodeUnits_startsWith'] = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_isJust']))(($GLOBALS['Data_String_CodeUnits_stripPrefix'])($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodeUnits_lastIndexOf'
$GLOBALS['Data_String_CodeUnits_lastIndexOf__prime__'] = (($GLOBALS['Data_String_CodeUnits__lastIndexOfStartingAt'])($GLOBALS['Data_Maybe_Just']))(new Phpurs_Data0("Nothing"));

// Data_String_CodeUnits_lastIndexOf
$GLOBALS['Data_String_CodeUnits_lastIndexOf'] = (($GLOBALS['Data_String_CodeUnits__lastIndexOf'])($GLOBALS['Data_Maybe_Just']))(new Phpurs_Data0("Nothing"));

// Data_String_CodeUnits_indexOf'
$GLOBALS['Data_String_CodeUnits_indexOf__prime__'] = (($GLOBALS['Data_String_CodeUnits__indexOfStartingAt'])($GLOBALS['Data_Maybe_Just']))(new Phpurs_Data0("Nothing"));

// Data_String_CodeUnits_indexOf
$GLOBALS['Data_String_CodeUnits_indexOf'] = (($GLOBALS['Data_String_CodeUnits__indexOf'])($GLOBALS['Data_Maybe_Just']))(new Phpurs_Data0("Nothing"));

// Data_String_CodeUnits_endsWith
$GLOBALS['Data_String_CodeUnits_endsWith'] = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_isJust']))(($GLOBALS['Data_String_CodeUnits_stripSuffix'])($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodeUnits_dropWhile
$GLOBALS['Data_String_CodeUnits_dropWhile'] = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_String_CodeUnits_drop'])((($GLOBALS['Data_String_CodeUnits_countPrefix'])($p_0))($s_1)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodeUnits_dropRight
$GLOBALS['Data_String_CodeUnits_dropRight'] = (function() {
  $__fn = function($i_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_String_CodeUnits_take'])((($GLOBALS['Data_String_CodeUnits_length'])($s_1) - $i_0)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodeUnits_contains
$GLOBALS['Data_String_CodeUnits_contains'] = function($pat_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_isJust']))(($GLOBALS['Data_String_CodeUnits_indexOf'])($pat_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodeUnits_charAt
$GLOBALS['Data_String_CodeUnits_charAt'] = (($GLOBALS['Data_String_CodeUnits__charAt'])($GLOBALS['Data_Maybe_Just']))(new Phpurs_Data0("Nothing"));

