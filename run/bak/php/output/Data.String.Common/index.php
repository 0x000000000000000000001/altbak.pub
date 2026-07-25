<?php

namespace Data\String\Common;

// ALL IMPORTS: Data.Eq, Data.Ordering, Data.String.Common, Data.String.Pattern, Prelude, Prim
// TO REQUIRE: Data.Eq, Data.Ordering, Data.String.Common, Data.String.Pattern, Prelude
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.String.Common/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
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
\PhpursThunks::$thunks['Data_String_Common_null'] = function() { $v = function($s_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Eq_eqStringImpl'] ?? \PhpursThunks::eval('Data_Eq_eqStringImpl')))($s_0))("");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_String_Common_localeCompare'] = function() { $v = (((($GLOBALS['Data_String_Common__localeCompare'] ?? \PhpursThunks::eval('Data_String_Common__localeCompare')))(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT")); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_String_Common = \call_user_func(function() {
  $exports = [];
$_localeCompare = function($lt, $eq = null, $gt = null, $s1 = null, $s2 = null) use (&$_localeCompare) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_localeCompare) {

            return $_localeCompare(...\array_merge($__args, $more));
        };
    }
    $result = strcmp($s1, $s2);
    return $result < 0 ? $lt : ($result > 0 ? $gt : $eq);
};

$replace = function($s1, $s2 = null, $s3 = null) use (&$replace) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$replace) {

            return $replace(...\array_merge($__args, $more));
        };
    }
    $pos = strpos($s3, $s1);
    if ($pos !== false) {
        return substr_replace($s3, $s2, $pos, strlen($s1));
    }
    return $s3;
};

$replaceAll = function($s1, $s2 = null, $s3 = null) use (&$replaceAll) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$replaceAll) {

            return $replaceAll(...\array_merge($__args, $more));
        };
    }
    return str_replace($s1, $s2, $s3);
};

$split = function($sep, $s = null) use (&$split) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$split) {

            return $split(...\array_merge($__args, $more));
        };
    }
    if ($sep === "") {
        return str_split($s);
    }
    return explode($sep, $s);
};

$toLower = function($s) use (&$toLower) {
    return mb_strtolower($s, 'UTF-8');
};

$toUpper = function($s) use (&$toUpper) {
    return mb_strtoupper($s, 'UTF-8');
};

$trim = function($s) use (&$trim) {
    return trim($s);
};

$joinWith = function($s, $xs = null) use (&$joinWith) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$joinWith) {

            return $joinWith(...\array_merge($__args, $more));
        };
    }
    return implode($s, $xs);
};

$exports['_localeCompare'] = $_localeCompare;
$exports['replace'] = $replace;
$exports['replaceAll'] = $replaceAll;
$exports['split'] = $split;
$exports['toLower'] = $toLower;
$exports['toUpper'] = $toUpper;
$exports['trim'] = $trim;
$exports['joinWith'] = $joinWith;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_String_Common__localeCompare'] = function() use (&$ffi_Data_String_Common) { return $ffi_Data_String_Common['_localeCompare']; };
\PhpursThunks::$thunks['Data_String_Common_joinWith'] = function() use (&$ffi_Data_String_Common) { return $ffi_Data_String_Common['joinWith']; };
\PhpursThunks::$thunks['Data_String_Common_replace'] = function() use (&$ffi_Data_String_Common) { return $ffi_Data_String_Common['replace']; };
\PhpursThunks::$thunks['Data_String_Common_replaceAll'] = function() use (&$ffi_Data_String_Common) { return $ffi_Data_String_Common['replaceAll']; };
\PhpursThunks::$thunks['Data_String_Common_split'] = function() use (&$ffi_Data_String_Common) { return $ffi_Data_String_Common['split']; };
\PhpursThunks::$thunks['Data_String_Common_toLower'] = function() use (&$ffi_Data_String_Common) { return $ffi_Data_String_Common['toLower']; };
\PhpursThunks::$thunks['Data_String_Common_toUpper'] = function() use (&$ffi_Data_String_Common) { return $ffi_Data_String_Common['toUpper']; };
\PhpursThunks::$thunks['Data_String_Common_trim'] = function() use (&$ffi_Data_String_Common) { return $ffi_Data_String_Common['trim']; };




