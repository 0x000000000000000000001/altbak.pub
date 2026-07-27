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
function majData_majString_majCodemajUnits__charmajAt($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__charmajAt';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['_charAt'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_String_CodeUnits__charAt'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__charmajAt';

function majData_majString_majCodemajUnits__indexmajOf($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__indexmajOf';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['_indexOf'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_String_CodeUnits__indexOf'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__indexmajOf';

function majData_majString_majCodemajUnits__indexmajOfmajStartingmajAt($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__indexmajOfmajStartingmajAt';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['_indexOfStartingAt'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_String_CodeUnits__indexOfStartingAt'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__indexmajOfmajStartingmajAt';

function majData_majString_majCodemajUnits__lastmajIndexmajOf($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__lastmajIndexmajOf';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['_lastIndexOf'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_String_CodeUnits__lastIndexOf'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__lastmajIndexmajOf';

function majData_majString_majCodemajUnits__lastmajIndexmajOfmajStartingmajAt($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__lastmajIndexmajOfmajStartingmajAt';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['_lastIndexOfStartingAt'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_String_CodeUnits__lastIndexOfStartingAt'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__lastmajIndexmajOfmajStartingmajAt';

function majData_majString_majCodemajUnits__tomajChar($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__tomajChar';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['_toChar'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_String_CodeUnits__toChar'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits__tomajChar';

function majData_majString_majCodemajUnits_countmajPrefix($v0, $v1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_countmajPrefix';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['countPrefix'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_CodeUnits_countPrefix'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_countmajPrefix';

function majData_majString_majCodemajUnits_drop(int $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_drop';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['drop'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_CodeUnits_drop'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_drop';

function majData_majString_majCodemajUnits_frommajCharmajArray($v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_frommajCharmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['fromCharArray'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_CodeUnits_fromCharArray'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_frommajCharmajArray';

function majData_majString_majCodemajUnits_length(string $v0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['length'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_CodeUnits_length'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_length';

function majData_majString_majCodemajUnits_singleton($v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['singleton'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_CodeUnits_singleton'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_singleton';

function majData_majString_majCodemajUnits_slice(int $v0, $v1 = null, $v2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_slice';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['slice'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_String_CodeUnits_slice'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_slice';

function majData_majString_majCodemajUnits_splitmajAt(int $v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_splitmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['splitAt'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_CodeUnits_splitAt'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_splitmajAt';

function majData_majString_majCodemajUnits_take(int $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_take';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['take'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_CodeUnits_take'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_take';

function majData_majString_majCodemajUnits_tomajCharmajArray(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_tomajCharmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_CodeUnits;
  $f = ($ffi_Data_String_CodeUnits['toCharArray'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_CodeUnits_toCharArray'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_tomajCharmajArray';





// Data_String_CodeUnits_uncons
function majData_majString_majCodemajUnits_uncons(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = match ($v_0) { "" => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just((object)["head" => \Data\String\Unsafe\majData_majString_majUnsafe_charmajAt(0, $v_0), "tail" => \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(1, $v_0)]) };
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_uncons'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_uncons';

// Data_String_CodeUnits_toChar_closure
$GLOBALS['Data_String_CodeUnits_toChar_closure'] = (($GLOBALS['Data_String_CodeUnits__toChar'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_CodeUnits_toChar
function majData_majString_majCodemajUnits_tomajChar(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_tomajChar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_CodeUnits_toChar_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_toChar'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_tomajChar';

// Data_String_CodeUnits_takeWhile
function majData_majString_majCodemajUnits_takemajWhile($p_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_takemajWhile';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\CodeUnits\majData_majString_majCodemajUnits_take(\Data\String\CodeUnits\majData_majString_majCodemajUnits_countmajPrefix($p_0, $s_1), $s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_takeWhile'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_takemajWhile';

// Data_String_CodeUnits_takeRight
function majData_majString_majCodemajUnits_takemajRight(int $i_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_takemajRight';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop((\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($s_1) - $i_0), $s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_takeRight'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_takemajRight';

// Data_String_CodeUnits_stripSuffix
function majData_majString_majCodemajUnits_stripmajSuffix(string $v_0, $str_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_stripmajSuffix';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = \Data\String\CodeUnits\majData_majString_majCodemajUnits_splitmajAt((\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($str_1) - \Data\String\CodeUnits\majData_majString_majCodemajUnits_length($v_0)), $str_1);
  $__t1 = null;;
  if (((($GLOBALS['Data_Eq_eqString'])->{'eq'})(($v1_2_0)->{'after'}))($v_0)) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($v1_2_0)->{'before'});
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_stripSuffix'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_stripmajSuffix';

// Data_String_CodeUnits_stripPrefix
function majData_majString_majCodemajUnits_stripmajPrefix(string $v_0, $str_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_stripmajPrefix';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = \Data\String\CodeUnits\majData_majString_majCodemajUnits_splitmajAt(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($v_0), $str_1);
  $__t1 = null;;
  if (((($GLOBALS['Data_Eq_eqString'])->{'eq'})(($v1_2_0)->{'before'}))($v_0)) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($v1_2_0)->{'after'});
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_stripPrefix'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_stripmajPrefix';

// Data_String_CodeUnits_startsWith
function majData_majString_majCodemajUnits_startsmajWith(string $pat_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_startsmajWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_isJust']))(($GLOBALS['Data_String_CodeUnits_stripPrefix'])($pat_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_startsWith'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_startsmajWith';

// Data_String_CodeUnits_lastIndexOf'_closure
$GLOBALS['Data_String_CodeUnits_lastIndexOf__prime___closure'] = (($GLOBALS['Data_String_CodeUnits__lastIndexOfStartingAt'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_CodeUnits_lastIndexOf'
function majData_majString_majCodemajUnits_lastmajIndexmajOf__prime__(string $v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_lastmajIndexmajOf__prime__';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_String_CodeUnits_lastIndexOf__prime___closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_lastIndexOf__prime__'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_lastmajIndexmajOf__prime__';

// Data_String_CodeUnits_lastIndexOf_closure
$GLOBALS['Data_String_CodeUnits_lastIndexOf_closure'] = (($GLOBALS['Data_String_CodeUnits__lastIndexOf'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_CodeUnits_lastIndexOf
function majData_majString_majCodemajUnits_lastmajIndexmajOf(string $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_lastmajIndexmajOf';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_CodeUnits_lastIndexOf_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_lastIndexOf'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_lastmajIndexmajOf';

// Data_String_CodeUnits_indexOf'_closure
$GLOBALS['Data_String_CodeUnits_indexOf__prime___closure'] = (($GLOBALS['Data_String_CodeUnits__indexOfStartingAt'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_CodeUnits_indexOf'
function majData_majString_majCodemajUnits_indexmajOf__prime__(string $v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_indexmajOf__prime__';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_String_CodeUnits_indexOf__prime___closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_indexOf__prime__'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_indexmajOf__prime__';

// Data_String_CodeUnits_indexOf_closure
$GLOBALS['Data_String_CodeUnits_indexOf_closure'] = (($GLOBALS['Data_String_CodeUnits__indexOf'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_CodeUnits_indexOf
function majData_majString_majCodemajUnits_indexmajOf(string $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_indexmajOf';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_CodeUnits_indexOf_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_indexOf'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_indexmajOf';

// Data_String_CodeUnits_endsWith
function majData_majString_majCodemajUnits_endsmajWith(string $pat_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_endsmajWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_isJust']))(($GLOBALS['Data_String_CodeUnits_stripSuffix'])($pat_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_endsWith'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_endsmajWith';

// Data_String_CodeUnits_dropWhile
function majData_majString_majCodemajUnits_dropmajWhile($p_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_dropmajWhile';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(\Data\String\CodeUnits\majData_majString_majCodemajUnits_countmajPrefix($p_0, $s_1), $s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_dropWhile'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_dropmajWhile';

// Data_String_CodeUnits_dropRight
function majData_majString_majCodemajUnits_dropmajRight(int $i_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_dropmajRight';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\CodeUnits\majData_majString_majCodemajUnits_take((\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($s_1) - $i_0), $s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_dropRight'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_dropmajRight';

// Data_String_CodeUnits_contains
function majData_majString_majCodemajUnits_contains(string $pat_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_contains';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_isJust']))(($GLOBALS['Data_String_CodeUnits_indexOf'])($pat_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_contains'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_contains';

// Data_String_CodeUnits_charAt_closure
$GLOBALS['Data_String_CodeUnits_charAt_closure'] = (($GLOBALS['Data_String_CodeUnits__charAt'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_CodeUnits_charAt
function majData_majString_majCodemajUnits_charmajAt(int $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajUnits_charmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_CodeUnits_charAt_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodeUnits_charAt'] = __NAMESPACE__ . '\\majData_majString_majCodemajUnits_charmajAt';

