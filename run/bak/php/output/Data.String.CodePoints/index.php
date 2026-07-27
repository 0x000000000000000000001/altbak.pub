<?php

namespace Data\String\CodePoints;

// ALL IMPORTS: Control.Semigroupoid, Data.Array, Data.Boolean, Data.Bounded, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Ord, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.String.CodePoints, Data.String.CodeUnits, Data.String.Common, Data.String.Pattern, Data.String.Unsafe, Data.Tuple, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Array, Data.Boolean, Data.Bounded, Data.Enum, Data.Eq, Data.EuclideanRing, Data.Functor, Data.HeytingAlgebra, Data.Int, Data.Maybe, Data.Ord, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.String.CodePoints, Data.String.CodeUnits, Data.String.Common, Data.String.Pattern, Data.String.Unsafe, Data.Tuple, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Int/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.String.CodePoints/index.php';
require_once __DIR__ . '/../Data.String.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.Common/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
require_once __DIR__ . '/../Data.String.Unsafe/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
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
$ffi_Data_String_CodePoints = \call_user_func(function() {
  $exports = [];
if (!\function_exists('Data_String_CodePoints_utf8_ord')) {
    function Data_String_CodePoints_utf8_ord($char) {
        if ($char === '') return 0;
        $c0 = ord($char[0]);
        if ($c0 < 0x80) {
            return $c0;
        } elseif ($c0 < 0xE0) {
            return (($c0 & 0x1F) << 6) | (ord($char[1]) & 0x3F);
        } elseif ($c0 < 0xF0) {
            return (($c0 & 0x0F) << 12) | ((ord($char[1]) & 0x3F) << 6) | (ord($char[2]) & 0x3F);
        } else {
            return (($c0 & 0x07) << 18) | ((ord($char[1]) & 0x3F) << 12) | ((ord($char[2]) & 0x3F) << 6) | (ord($char[3]) & 0x3F);
        }
    }
}

if (!\function_exists('Data_String_CodePoints_utf8_chr')) {
    function Data_String_CodePoints_utf8_chr($code) {
        if ($code < 0x80) {
            return chr($code);
        } elseif ($code < 0x800) {
            return chr(0xC0 | ($code >> 6)) . chr(0x80 | ($code & 0x3F));
        } elseif ($code < 0x10000) {
            return chr(0xE0 | ($code >> 12)) . chr(0x80 | (($code >> 6) & 0x3F)) . chr(0x80 | ($code & 0x3F));
        } else {
            return chr(0xF0 | ($code >> 18)) . chr(0x80 | (($code >> 12) & 0x3F)) . chr(0x80 | (($code >> 6) & 0x3F)) . chr(0x80 | ($code & 0x3F));
        }
    }
}

$_unsafeCodePointAt0 = function($fallback, $str = null) use (&$_unsafeCodePointAt0) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_unsafeCodePointAt0) {

            return $_unsafeCodePointAt0(...\array_merge($__args, $more));
        };
    }
    return Data_String_CodePoints_utf8_ord(iconv_substr($str, 0, 1, 'UTF-8'));
};

$_codePointAt = function($fallback, $just = null, $nothing = null, $unsafeCodePointAt0 = null, $index = null, $str = null) use (&$_codePointAt) {
    if (\func_num_args() < 6) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_codePointAt) {

            return $_codePointAt(...\array_merge($__args, $more));
        };
    }
    $len = iconv_strlen($str, 'UTF-8');
    if ($index < 0 || $index >= $len) return $nothing;
    return $just($unsafeCodePointAt0(iconv_substr($str, $index, 1, 'UTF-8')));
};

$_countPrefix = function($fallback, $unsafeCodePointAt0 = null, $pred = null, $str = null) use (&$_countPrefix) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_countPrefix) {

            return $_countPrefix(...\array_merge($__args, $more));
        };
    }
    $len = iconv_strlen($str, 'UTF-8');
    for ($i = 0; $i < $len; $i++) {
        $char = iconv_substr($str, $i, 1, 'UTF-8');
        $cp = $unsafeCodePointAt0($char);
        if (!$pred($cp)) return $i;
    }
    return $len;
};

$_fromCodePointArray = function($singleton, $cps = null) use (&$_fromCodePointArray) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_fromCodePointArray) {

            return $_fromCodePointArray(...\array_merge($__args, $more));
        };
    }
    $result = "";
    foreach ($cps as $cp) {
        $result .= Data_String_CodePoints_utf8_chr($cp);
    }
    return $result;
};

$_singleton = function($fallback, $cp = null) use (&$_singleton) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_singleton) {

            return $_singleton(...\array_merge($__args, $more));
        };
    }
    return Data_String_CodePoints_utf8_chr($cp);
};

$_take = function($fallback, $n = null, $str = null) use (&$_take) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_take) {

            return $_take(...\array_merge($__args, $more));
        };
    }
    return iconv_substr($str, 0, $n, 'UTF-8');
};

$_toCodePointArray = function($fallback, $unsafeCodePointAt0 = null, $str = null) use (&$_toCodePointArray) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_toCodePointArray) {

            return $_toCodePointArray(...\array_merge($__args, $more));
        };
    }
    $len = iconv_strlen($str, 'UTF-8');
    $arr = [];
    for ($i = 0; $i < $len; $i++) {
        $arr[] = $unsafeCodePointAt0(iconv_substr($str, $i, 1, 'UTF-8'));
    }
    return $arr;
};

$exports['_unsafeCodePointAt0'] = $_unsafeCodePointAt0;
$exports['_codePointAt'] = $_codePointAt;
$exports['_countPrefix'] = $_countPrefix;
$exports['_fromCodePointArray'] = $_fromCodePointArray;
$exports['_singleton'] = $_singleton;
$exports['_take'] = $_take;
$exports['_toCodePointArray'] = $_toCodePointArray;
return $exports;
  return $exports;
});
function majData_majString_majCodemajPoints__codemajPointmajAt($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__codemajPointmajAt';
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  global $ffi_Data_String_CodePoints;
  $f = ($ffi_Data_String_CodePoints['_codePointAt'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5);
}
$GLOBALS['Data_String_CodePoints__codePointAt'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__codemajPointmajAt';

function majData_majString_majCodemajPoints__countmajPrefix($v0, $v1 = null, $v2 = null, $v3 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__countmajPrefix';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_String_CodePoints;
  $f = ($ffi_Data_String_CodePoints['_countPrefix'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_String_CodePoints__countPrefix'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__countmajPrefix';

function majData_majString_majCodemajPoints__frommajCodemajPointmajArray($v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__frommajCodemajPointmajArray';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_CodePoints;
  $f = ($ffi_Data_String_CodePoints['_fromCodePointArray'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_CodePoints__fromCodePointArray'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__frommajCodemajPointmajArray';

function majData_majString_majCodemajPoints__singleton($v0, $v1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__singleton';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_CodePoints;
  $f = ($ffi_Data_String_CodePoints['_singleton'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_CodePoints__singleton'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__singleton';

function majData_majString_majCodemajPoints__take($v0, $v1 = null, $v2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__take';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_String_CodePoints;
  $f = ($ffi_Data_String_CodePoints['_take'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_String_CodePoints__take'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__take';

function majData_majString_majCodemajPoints__tomajCodemajPointmajArray($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__tomajCodemajPointmajArray';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_String_CodePoints;
  $f = ($ffi_Data_String_CodePoints['_toCodePointArray'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_String_CodePoints__toCodePointArray'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__tomajCodemajPointmajArray';

function majData_majString_majCodemajPoints__unsafemajCodemajPointmajAt0($v0, $v1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__unsafemajCodemajPointmajAt0';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_CodePoints;
  $f = ($ffi_Data_String_CodePoints['_unsafeCodePointAt0'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_CodePoints__unsafeCodePointAt0'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints__unsafemajCodemajPointmajAt0';





// Data_String_CodePoints_lessThanOrEq
$GLOBALS['Data_String_CodePoints_lessThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_String_CodePoints_greaterThan
$GLOBALS['Data_String_CodePoints_greaterThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_String_CodePoints_lessThan
$GLOBALS['Data_String_CodePoints_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_String_CodePoints_compare
$GLOBALS['Data_String_CodePoints_compare'] = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());

// Data_String_CodePoints_greaterThanOrEq
$GLOBALS['Data_String_CodePoints_greaterThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_String_CodePoints_CodePoint
function majData_majString_majCodemajPoints_majCodemajPoint($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_majCodemajPoint';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_CodePoint'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_majCodemajPoint';

// Data_String_CodePoints_showCodePoint
$GLOBALS['Data_String_CodePoints_showCodePoint'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(CodePoint 0x"))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(\Data\String\Common\majData_majString_majCommon_tomajUpper(\Data\Int\majData_majInt_tomajStringmajAs(16, $v_0))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_CodePoints_uncons
function majData_majString_majCodemajPoints_uncons(string $s_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $v_1_0 = \Data\String\CodeUnits\majData_majString_majCodemajUnits_length($s_0);
  $__t4 = null;;
  switch ($v_1_0) {
case 0:
$__t4 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_4;;
break;
default:
;
break;
};
  switch ($v_1_0) {
case 1:
$__t4 = new \Data\Maybe\Data_Maybe_Just((object)["head" => (($GLOBALS['Data_Enum_boundedEnumChar'])->{'fromEnum'})(\Data\String\Unsafe\majData_majString_majUnsafe_charmajAt(0, $s_0)), "tail" => ""]);
goto end_branch_4;;
break;
default:
;
break;
};
  $cu1_2_1 = (($GLOBALS['Data_Enum_boundedEnumChar'])->{'fromEnum'})(\Data\String\Unsafe\majData_majString_majUnsafe_charmajAt(1, $s_0));
  $cu0_3_2 = (($GLOBALS['Data_Enum_boundedEnumChar'])->{'fromEnum'})(\Data\String\Unsafe\majData_majString_majUnsafe_charmajAt(0, $s_0));
  $__t3 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])(55296))($cu0_3_2)))((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])($cu0_3_2))(56319))))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])(56320))($cu1_2_1)))((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])($cu1_2_1))(57343)))) {
$__t3 = new \Data\Maybe\Data_Maybe_Just((object)["head" => (((($cu0_3_2 - 55296) * 1024) + ($cu1_2_1 - 56320)) + 65536), "tail" => \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(2, $s_0)]);
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Just((object)["head" => $cu0_3_2, "tail" => \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(1, $s_0)]);
  end_branch_3:;
  $__t4 = $__t3;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_uncons'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_uncons';

// Data_String_CodePoints_unconsButWithTuple
function majData_majString_majCodemajPoints_unconsmajButmajWithmajTuple(string $s_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_unconsmajButmajWithmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'head'}, ($v_1)->{'tail'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\String\CodePoints\majData_majString_majCodemajPoints_uncons($s_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_unconsButWithTuple'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_unconsmajButmajWithmajTuple';

// Data_String_CodePoints_toCodePointArrayFallback
function majData_majString_majCodemajPoints_tomajCodemajPointmajArraymajFallback(string $s_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_tomajCodemajPointmajArraymajFallback';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Unfoldable_unfoldableArray'])->{'unfoldr'})($GLOBALS['Data_String_CodePoints_unconsButWithTuple']))($s_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_toCodePointArrayFallback'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_tomajCodemajPointmajArraymajFallback';

// Data_String_CodePoints_unsafeCodePointAt0Fallback
function majData_majString_majCodemajPoints_unsafemajCodemajPointmajAt0majFallback(string $s_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_unsafemajCodemajPointmajAt0majFallback';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $cu0_1_0 = (($GLOBALS['Data_Enum_boundedEnumChar'])->{'fromEnum'})(\Data\String\Unsafe\majData_majString_majUnsafe_charmajAt(0, $s_0));
  $__t1 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])(55296))($cu0_1_0)))((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])($cu0_1_0))(56319))))((($GLOBALS['Data_String_CodePoints_greaterThan'])(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($s_0)))(1))) {
$cu1_2_2 = (($GLOBALS['Data_Enum_boundedEnumChar'])->{'fromEnum'})(\Data\String\Unsafe\majData_majString_majUnsafe_charmajAt(1, $s_0));
$__t3 = null;;
if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])(56320))($cu1_2_2)))((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])($cu1_2_2))(57343))) {
$__t3 = (((($cu0_1_0 - 55296) * 1024) + ($cu1_2_2 - 56320)) + 65536);
goto end_branch_3;;
};
$__t3 = $cu0_1_0;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  $__t1 = $cu0_1_0;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_unsafeCodePointAt0Fallback'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_unsafemajCodemajPointmajAt0majFallback';

// Data_String_CodePoints_unsafeCodePointAt0_closure
$GLOBALS['Data_String_CodePoints_unsafeCodePointAt0_closure'] = ($GLOBALS['Data_String_CodePoints__unsafeCodePointAt0'])($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0Fallback']);

// Data_String_CodePoints_unsafeCodePointAt0
function majData_majString_majCodemajPoints_unsafemajCodemajPointmajAt0(string $v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_unsafemajCodemajPointmajAt0';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_unsafeCodePointAt0'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_unsafemajCodemajPointmajAt0';

// Data_String_CodePoints_toCodePointArray_closure
$GLOBALS['Data_String_CodePoints_toCodePointArray_closure'] = (($GLOBALS['Data_String_CodePoints__toCodePointArray'])($GLOBALS['Data_String_CodePoints_toCodePointArrayFallback']))($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0']);

// Data_String_CodePoints_toCodePointArray
function majData_majString_majCodemajPoints_tomajCodemajPointmajArray(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_tomajCodemajPointmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_toCodePointArray_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_toCodePointArray'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_tomajCodemajPointmajArray';

// Data_String_CodePoints_length_closure
$GLOBALS['Data_String_CodePoints_length_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_length']))($GLOBALS['Data_String_CodePoints_toCodePointArray']);

// Data_String_CodePoints_length
function majData_majString_majCodemajPoints_length(string $v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_length_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_length'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_length';

// Data_String_CodePoints_lastIndexOf
function majData_majString_majCodemajPoints_lastmajIndexmajOf(string $p_0, $s_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_lastmajIndexmajOf';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($i_2) use ($s_1) {
  $__num = \func_num_args();
  $__res = \Data\String\CodePoints\majData_majString_majCodemajPoints_length(\Data\String\CodeUnits\majData_majString_majCodemajUnits_take($i_2, $s_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\String\CodeUnits\majData_majString_majCodemajUnits_lastmajIndexmajOf($p_0, $s_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_lastIndexOf'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_lastmajIndexmajOf';

// Data_String_CodePoints_indexOf
function majData_majString_majCodemajPoints_indexmajOf(string $p_0, $s_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_indexmajOf';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($i_2) use ($s_1) {
  $__num = \func_num_args();
  $__res = \Data\String\CodePoints\majData_majString_majCodemajPoints_length(\Data\String\CodeUnits\majData_majString_majCodemajUnits_take($i_2, $s_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\String\CodeUnits\majData_majString_majCodemajUnits_indexmajOf($p_0, $s_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_indexOf'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_indexmajOf';

// Data_String_CodePoints_fromCharCode_closure
$GLOBALS['Data_String_CodePoints_fromCharCode_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_CodeUnits_singleton']))(((($GLOBALS['Data_Enum_toEnumWithDefaults'])($GLOBALS['Data_Enum_boundedEnumChar']))(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'}))(($GLOBALS['Data_Bounded_boundedChar'])->{'top'}));

// Data_String_CodePoints_fromCharCode
function majData_majString_majCodemajPoints_frommajCharmajCode(int $v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_frommajCharmajCode';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_fromCharCode_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_fromCharCode'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_frommajCharmajCode';

// Data_String_CodePoints_singletonFallback
function majData_majString_majCodemajPoints_singletonmajFallback(int $v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_singletonmajFallback';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])($v_0))(65535)) {
$__t0 = \Data\String\CodePoints\majData_majString_majCodemajPoints_frommajCharmajCode($v_0);
goto end_branch_0;;
};
  $__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(\Data\String\CodePoints\majData_majString_majCodemajPoints_frommajCharmajCode((((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])->{'div'})(($v_0 - 65536)))(1024) + 55296))))(\Data\String\CodePoints\majData_majString_majCodemajPoints_frommajCharmajCode((((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])->{'mod'})(($v_0 - 65536)))(1024) + 56320)));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_singletonFallback'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_singletonmajFallback';

// Data_String_CodePoints_fromCodePointArray_closure
$GLOBALS['Data_String_CodePoints_fromCodePointArray_closure'] = ($GLOBALS['Data_String_CodePoints__fromCodePointArray'])($GLOBALS['Data_String_CodePoints_singletonFallback']);

// Data_String_CodePoints_fromCodePointArray
function majData_majString_majCodemajPoints_frommajCodemajPointmajArray($v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_frommajCodemajPointmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_fromCodePointArray_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_fromCodePointArray'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_frommajCodemajPointmajArray';

// Data_String_CodePoints_singleton_closure
$GLOBALS['Data_String_CodePoints_singleton_closure'] = ($GLOBALS['Data_String_CodePoints__singleton'])($GLOBALS['Data_String_CodePoints_singletonFallback']);

// Data_String_CodePoints_singleton
function majData_majString_majCodemajPoints_singleton(int $v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_singleton_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_singleton'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_singleton';

// Data_String_CodePoints_takeFallback
function majData_majString_majCodemajPoints_takemajFallback(int $v_0, $v1_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_takemajFallback';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_String_CodePoints_takeFallback_v_0 = $v_0;
  $__tco_var_Data_String_CodePoints_takeFallback_v1_1 = $v1_1;
  tco_loop_Data_String_CodePoints_takeFallback:;
  $v_0 = $__tco_var_Data_String_CodePoints_takeFallback_v_0;
  $v1_1 = $__tco_var_Data_String_CodePoints_takeFallback_v1_1;
  $__t2 = null;;
  if ((($GLOBALS['Data_String_CodePoints_lessThan'])($v_0))(1)) {
$__t2 = "";
goto end_branch_2;;
};
  $v2_2_0 = \Data\String\CodePoints\majData_majString_majCodemajPoints_uncons($v1_1);
  $__t1 = null;;
  if ($v2_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(\Data\String\CodePoints\majData_majString_majCodemajPoints_singleton((($v2_2_0)->{'value0'})->{'head'})))(\Data\String\CodePoints\majData_majString_majCodemajPoints_takemajFallback(($v_0 - 1), (($v2_2_0)->{'value0'})->{'tail'}));
goto end_branch_1;;
};
  $__t1 = $v1_1;
  end_branch_1:;
  $__t2 = $__t1;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_takeFallback'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_takemajFallback';

// Data_String_CodePoints_take_closure
$GLOBALS['Data_String_CodePoints_take_closure'] = ($GLOBALS['Data_String_CodePoints__take'])($GLOBALS['Data_String_CodePoints_takeFallback']);

// Data_String_CodePoints_take
function majData_majString_majCodemajPoints_take(int $v_0, $v_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_take';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_take_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_take'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_take';

// Data_String_CodePoints_lastIndexOf'
function majData_majString_majCodemajPoints_lastmajIndexmajOf__prime__(string $p_0, $i_1 = null, $s_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_lastmajIndexmajOf__prime__';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($k_3) use ($s_2) {
  $__num = \func_num_args();
  $__res = \Data\String\CodePoints\majData_majString_majCodemajPoints_length(\Data\String\CodeUnits\majData_majString_majCodemajUnits_take($k_3, $s_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($GLOBALS['Data_String_CodeUnits_lastIndexOf__prime__'])($p_0))(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length(\Data\String\CodePoints\majData_majString_majCodemajPoints_take($i_1, $s_2))))($s_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_String_CodePoints_lastIndexOf__prime__'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_lastmajIndexmajOf__prime__';

// Data_String_CodePoints_splitAt
function majData_majString_majCodemajPoints_splitmajAt(int $i_0, $s_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_splitmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $before_2_0 = \Data\String\CodePoints\majData_majString_majCodemajPoints_take($i_0, $s_1);
  $__res = (object)["before" => $before_2_0, "after" => \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length($before_2_0), $s_1)];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_splitAt'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_splitmajAt';

// Data_String_CodePoints_eqCodePoint
$GLOBALS['Data_String_CodePoints_eqCodePoint'] = (object)["eq" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($x_0 === $y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_String_CodePoints_ordCodePoint
$GLOBALS['Data_String_CodePoints_ordCodePoint'] = (object)["compare" => (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_String_CodePoints_compare'])($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_eqCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_CodePoints_drop
function majData_majString_majCodemajPoints_drop(int $n_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_drop';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length(\Data\String\CodePoints\majData_majString_majCodemajPoints_take($n_0, $s_1)), $s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_drop'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_drop';

// Data_String_CodePoints_indexOf'
function majData_majString_majCodemajPoints_indexmajOf__prime__(string $p_0, $i_1 = null, $s_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_indexmajOf__prime__';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $s_prime_3_0 = \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length(\Data\String\CodePoints\majData_majString_majCodemajPoints_take($i_1, $s_2)), $s_2);
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($k_4) use ($i_1, $s_prime_3_0) {
  $__num = \func_num_args();
  $__res = ($i_1 + \Data\String\CodePoints\majData_majString_majCodemajPoints_length(\Data\String\CodeUnits\majData_majString_majCodemajUnits_take($k_4, $s_prime_3_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\String\CodeUnits\majData_majString_majCodemajUnits_indexmajOf($p_0, $s_prime_3_0));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_String_CodePoints_indexOf__prime__'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_indexmajOf__prime__';

// Data_String_CodePoints_countTail
function majData_majString_majCodemajPoints_countmajTail($p_0, $s_1 = null, $accum_2 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_countmajTail';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_String_CodePoints_countTail_p_0 = $p_0;
  $__tco_var_Data_String_CodePoints_countTail_s_1 = $s_1;
  $__tco_var_Data_String_CodePoints_countTail_accum_2 = $accum_2;
  tco_loop_Data_String_CodePoints_countTail:;
  $p_0 = $__tco_var_Data_String_CodePoints_countTail_p_0;
  $s_1 = $__tco_var_Data_String_CodePoints_countTail_s_1;
  $accum_2 = $__tco_var_Data_String_CodePoints_countTail_accum_2;
  $v_3_0 = \Data\String\CodePoints\majData_majString_majCodemajPoints_uncons($s_1);
  $__t1 = null;;
  if (($v_3_0 instanceof \Data\Maybe\Data_Maybe_Just && ($p_0)((($v_3_0)->{'value0'})->{'head'}))) {
$__tco_2 = $p_0;
$__tco_3 = (($v_3_0)->{'value0'})->{'tail'};
$__tco_4 = ($accum_2 + 1);
$__tco_var_Data_String_CodePoints_countTail_p_0 = $__tco_2;
$__tco_var_Data_String_CodePoints_countTail_s_1 = $__tco_3;
$__tco_var_Data_String_CodePoints_countTail_accum_2 = $__tco_4;
goto tco_loop_Data_String_CodePoints_countTail;;
$__t1 = null;
goto end_branch_1;;
};
  $__t1 = $accum_2;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_String_CodePoints_countTail'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_countmajTail';

// Data_String_CodePoints_countFallback
function majData_majString_majCodemajPoints_countmajFallback($p_0, $s_1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_countmajFallback';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\CodePoints\majData_majString_majCodemajPoints_countmajTail($p_0, $s_1, 0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_countFallback'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_countmajFallback';

// Data_String_CodePoints_countPrefix_closure
$GLOBALS['Data_String_CodePoints_countPrefix_closure'] = (($GLOBALS['Data_String_CodePoints__countPrefix'])($GLOBALS['Data_String_CodePoints_countFallback']))($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0']);

// Data_String_CodePoints_countPrefix
function majData_majString_majCodemajPoints_countmajPrefix($v_0, $v_1 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_countmajPrefix';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_countPrefix_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_countPrefix'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_countmajPrefix';

// Data_String_CodePoints_dropWhile
function majData_majString_majCodemajPoints_dropmajWhile($p_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_dropmajWhile';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\CodeUnits\majData_majString_majCodemajUnits_drop(\Data\String\CodeUnits\majData_majString_majCodemajUnits_length(\Data\String\CodePoints\majData_majString_majCodemajPoints_take(\Data\String\CodePoints\majData_majString_majCodemajPoints_countmajPrefix($p_0, $s_1), $s_1)), $s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_dropWhile'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_dropmajWhile';

// Data_String_CodePoints_takeWhile
function majData_majString_majCodemajPoints_takemajWhile($p_0, $s_1 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_takemajWhile';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\CodePoints\majData_majString_majCodemajPoints_take(\Data\String\CodePoints\majData_majString_majCodemajPoints_countmajPrefix($p_0, $s_1), $s_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_takeWhile'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_takemajWhile';

// Data_String_CodePoints_codePointFromChar_closure
$GLOBALS['Data_String_CodePoints_codePointFromChar_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_CodePoints_CodePoint']))(($GLOBALS['Data_Enum_boundedEnumChar'])->{'fromEnum'});

// Data_String_CodePoints_codePointFromChar
function majData_majString_majCodemajPoints_codemajPointmajFrommajChar($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_codemajPointmajFrommajChar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_CodePoints_codePointFromChar_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_CodePoints_codePointFromChar'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_codemajPointmajFrommajChar';

// Data_String_CodePoints_codePointAtFallback
function majData_majString_majCodemajPoints_codemajPointmajAtmajFallback(int $n_0, $s_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_codemajPointmajAtmajFallback';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_String_CodePoints_codePointAtFallback_n_0 = $n_0;
  $__tco_var_Data_String_CodePoints_codePointAtFallback_s_1 = $s_1;
  tco_loop_Data_String_CodePoints_codePointAtFallback:;
  $n_0 = $__tco_var_Data_String_CodePoints_codePointAtFallback_n_0;
  $s_1 = $__tco_var_Data_String_CodePoints_codePointAtFallback_s_1;
  $v_2_0 = \Data\String\CodePoints\majData_majString_majCodemajPoints_uncons($s_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = null;;
switch ($n_0) {
case 0:
$__t4 = new \Data\Maybe\Data_Maybe_Just((($v_2_0)->{'value0'})->{'head'});
goto end_branch_4;;
break;
default:
;
break;
};
$__tco_2 = ($n_0 - 1);
$__tco_3 = (($v_2_0)->{'value0'})->{'tail'};
$__tco_var_Data_String_CodePoints_codePointAtFallback_n_0 = $__tco_2;
$__tco_var_Data_String_CodePoints_codePointAtFallback_s_1 = $__tco_3;
goto tco_loop_Data_String_CodePoints_codePointAtFallback;;
$__t4 = null;
end_branch_4:;
$__t1 = $__t4;
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_codePointAtFallback'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_codemajPointmajAtmajFallback';

// Data_String_CodePoints_codePointAt
function majData_majString_majCodemajPoints_codemajPointmajAt(int $v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majCodemajPoints_codemajPointmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((($GLOBALS['Data_String_CodePoints_lessThan'])($v_0))(0)) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  switch ($v_0) {
case 0:
$__t0 = match ($v1_1) { "" => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just(\Data\String\CodePoints\majData_majString_majCodemajPoints_unsafemajCodemajPointmajAt0($v1_1)) };
goto end_branch_0;;
break;
default:
;
break;
};
  $__t0 = \Data\String\CodePoints\majData_majString_majCodemajPoints__codemajPointmajAt($GLOBALS['Data_String_CodePoints_codePointAtFallback'], $GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $GLOBALS['Data_String_CodePoints_unsafeCodePointAt0'], $v_0, $v1_1);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_CodePoints_codePointAt'] = __NAMESPACE__ . '\\majData_majString_majCodemajPoints_codemajPointmajAt';

// Data_String_CodePoints_boundedCodePoint
$GLOBALS['Data_String_CodePoints_boundedCodePoint'] = (object)["bottom" => 0, "top" => 1114111, "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_ordCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_CodePoints_boundedEnumCodePoint
$GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'] = (object)["cardinality" => 1114112, "fromEnum" => function($v_0) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "toEnum" => function($n_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_String_CodePoints_greaterThanOrEq'])($n_0))(0)))((($GLOBALS['Data_String_CodePoints_lessThanOrEq'])($n_0))(1114111))) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($n_0);
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_boundedCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_enumCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_CodePoints_enumCodePoint
$GLOBALS['Data_String_CodePoints_enumCodePoint'] = (object)["succ" => function($a_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'])->{'toEnum'})(((($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'])->{'fromEnum'})($a_0) + 1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pred" => function($a_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'])->{'toEnum'})(((($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'])->{'fromEnum'})($a_0) - 1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Ord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_ordCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

