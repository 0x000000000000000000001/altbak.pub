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
function majData_majString_majCommon__localemajCompare($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCommon__localemajCompare';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_String_Common;
  $f = ($ffi_Data_String_Common['_localeCompare'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_String_Common__localeCompare'] = __NAMESPACE__ . '\\majData_majString_majCommon__localemajCompare';

function majData_majString_majCommon_joinmajWith(string $v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCommon_joinmajWith';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_Common;
  $f = ($ffi_Data_String_Common['joinWith'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_Common_joinWith'] = __NAMESPACE__ . '\\majData_majString_majCommon_joinmajWith';

function majData_majString_majCommon_replace(string $v0, $v1 = null, $v2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCommon_replace';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_String_Common;
  $f = ($ffi_Data_String_Common['replace'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_String_Common_replace'] = __NAMESPACE__ . '\\majData_majString_majCommon_replace';

function majData_majString_majCommon_replacemajAll(string $v0, $v1 = null, $v2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCommon_replacemajAll';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_String_Common;
  $f = ($ffi_Data_String_Common['replaceAll'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_String_Common_replaceAll'] = __NAMESPACE__ . '\\majData_majString_majCommon_replacemajAll';

function majData_majString_majCommon_split(string $v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCommon_split';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_Common;
  $f = ($ffi_Data_String_Common['split'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_Common_split'] = __NAMESPACE__ . '\\majData_majString_majCommon_split';

function majData_majString_majCommon_tomajLower(string $v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCommon_tomajLower';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_Common;
  $f = ($ffi_Data_String_Common['toLower'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_Common_toLower'] = __NAMESPACE__ . '\\majData_majString_majCommon_tomajLower';

function majData_majString_majCommon_tomajUpper(string $v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCommon_tomajUpper';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_Common;
  $f = ($ffi_Data_String_Common['toUpper'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_Common_toUpper'] = __NAMESPACE__ . '\\majData_majString_majCommon_tomajUpper';

function majData_majString_majCommon_trim(string $v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCommon_trim';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_Common;
  $f = ($ffi_Data_String_Common['trim'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_Common_trim'] = __NAMESPACE__ . '\\majData_majString_majCommon_trim';





// Data_String_Common_null
function majData_majString_majCommon_null(string $s_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCommon_null';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Eq_eqString'])->{'eq'})($s_0))("");
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_Common_null'] = __NAMESPACE__ . '\\majData_majString_majCommon_null';

// Data_String_Common_localeCompare_closure
$GLOBALS['Data_String_Common_localeCompare_closure'] = ((($GLOBALS['Data_String_Common__localeCompare'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());

// Data_String_Common_localeCompare
function majData_majString_majCommon_localemajCompare(string $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCommon_localemajCompare';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_Common_localeCompare_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Common_localeCompare'] = __NAMESPACE__ . '\\majData_majString_majCommon_localemajCompare';

