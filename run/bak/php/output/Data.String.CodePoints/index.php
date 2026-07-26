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
$GLOBALS['Data_String_CodePoints__codePointAt'] = $ffi_Data_String_CodePoints['_codePointAt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodePoints__countPrefix'] = $ffi_Data_String_CodePoints['_countPrefix'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodePoints__fromCodePointArray'] = $ffi_Data_String_CodePoints['_fromCodePointArray'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodePoints__singleton'] = $ffi_Data_String_CodePoints['_singleton'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodePoints__take'] = $ffi_Data_String_CodePoints['_take'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodePoints__toCodePointArray'] = $ffi_Data_String_CodePoints['_toCodePointArray'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_String_CodePoints__unsafeCodePointAt0'] = $ffi_Data_String_CodePoints['_unsafeCodePointAt0'] ?? new class { public function __invoke(...$args) { return $this; } };


// Data_String_CodePoints_CodePoint
$GLOBALS['Data_String_CodePoints_CodePoint'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodePoints_unsurrogate
$GLOBALS['Data_String_CodePoints_unsurrogate'] = (function() {
  $__fn = function($lead_0 = null, $trail_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semiring_semiringInt'])['add'])(((($GLOBALS['Data_Semiring_semiringInt'])['add'])(((($GLOBALS['Data_Semiring_semiringInt'])['mul'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($lead_0))(55296)))(1024)))(((($GLOBALS['Data_Ring_ringInt'])['sub'])($trail_1))(56320))))(65536);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_showCodePoint
$GLOBALS['Data_String_CodePoints_showCodePoint'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(CodePoint 0x"))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(($GLOBALS['Data_String_Common_toUpper'])((($GLOBALS['Data_Int_toStringAs'])(16))($v_0))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_CodePoints_uncons
$GLOBALS['Data_String_CodePoints_uncons'] = function($s_0 = null) {
  $__num = \func_num_args();
  $v_1_0 = ($GLOBALS['Data_String_CodeUnits_length'])($s_0);
  $__t4 = null;;
  switch ($v_1_0) {
case 0:
$__t4 = new Phpurs_Data0("Nothing");
goto end_branch_4;;
break;
default:
;
break;
};
  switch ($v_1_0) {
case 1:
$__t4 = new Phpurs_Data1("Just", ["head" => (($GLOBALS['Data_Enum_boundedEnumChar'])['fromEnum'])((($GLOBALS['Data_String_Unsafe_charAt'])(0))($s_0)), "tail" => ""]);
goto end_branch_4;;
break;
default:
;
break;
};
  $cu1_2_1 = (($GLOBALS['Data_Enum_boundedEnumChar'])['fromEnum'])((($GLOBALS['Data_String_Unsafe_charAt'])(1))($s_0));
  $cu0_3_2 = (($GLOBALS['Data_Enum_boundedEnumChar'])['fromEnum'])((($GLOBALS['Data_String_Unsafe_charAt'])(0))($s_0));
  $__t3 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((55296 <= $cu0_3_2)))(($cu0_3_2 <= 56319))))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((56320 <= $cu1_2_1)))(($cu1_2_1 <= 57343)))) {
$__t3 = new Phpurs_Data1("Just", ["head" => (($GLOBALS['Data_String_CodePoints_unsurrogate'])($cu0_3_2))($cu1_2_1), "tail" => (($GLOBALS['Data_String_CodeUnits_drop'])(2))($s_0)]);
goto end_branch_3;;
};
  $__t3 = new Phpurs_Data1("Just", ["head" => $cu0_3_2, "tail" => (($GLOBALS['Data_String_CodeUnits_drop'])(1))($s_0)]);
  end_branch_3:;
  $__t4 = $__t3;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodePoints_unconsButWithTuple
$GLOBALS['Data_String_CodePoints_unconsButWithTuple'] = function($s_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($v_1)['head'], ($v_1)['tail']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_String_CodePoints_uncons'])($s_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodePoints_toCodePointArrayFallback
$GLOBALS['Data_String_CodePoints_toCodePointArrayFallback'] = function($s_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Unfoldable_unfoldableArray'])['unfoldr'])($GLOBALS['Data_String_CodePoints_unconsButWithTuple']))($s_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodePoints_unsafeCodePointAt0Fallback
$GLOBALS['Data_String_CodePoints_unsafeCodePointAt0Fallback'] = function($s_0 = null) {
  $__num = \func_num_args();
  $cu0_1_0 = (($GLOBALS['Data_Enum_boundedEnumChar'])['fromEnum'])((($GLOBALS['Data_String_Unsafe_charAt'])(0))($s_0));
  $__t1 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((55296 <= $cu0_1_0)))(($cu0_1_0 <= 56319))))((($GLOBALS['Data_String_CodeUnits_length'])($s_0) > 1))) {
$cu1_2_2 = (($GLOBALS['Data_Enum_boundedEnumChar'])['fromEnum'])((($GLOBALS['Data_String_Unsafe_charAt'])(1))($s_0));
$__t3 = null;;
if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((56320 <= $cu1_2_2)))(($cu1_2_2 <= 57343))) {
$__t3 = (($GLOBALS['Data_String_CodePoints_unsurrogate'])($cu0_1_0))($cu1_2_2);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodePoints_unsafeCodePointAt0
$GLOBALS['Data_String_CodePoints_unsafeCodePointAt0'] = ($GLOBALS['Data_String_CodePoints__unsafeCodePointAt0'])($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0Fallback']);

// Data_String_CodePoints_toCodePointArray
$GLOBALS['Data_String_CodePoints_toCodePointArray'] = (($GLOBALS['Data_String_CodePoints__toCodePointArray'])($GLOBALS['Data_String_CodePoints_toCodePointArrayFallback']))($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0']);

// Data_String_CodePoints_length
$GLOBALS['Data_String_CodePoints_length'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_Array_length']))($GLOBALS['Data_String_CodePoints_toCodePointArray']);

// Data_String_CodePoints_lastIndexOf
$GLOBALS['Data_String_CodePoints_lastIndexOf'] = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($i_2 = null) use ($s_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_String_CodePoints_length'])((($GLOBALS['Data_String_CodeUnits_take'])($i_2))($s_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_CodeUnits_lastIndexOf'])($p_0))($s_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_indexOf
$GLOBALS['Data_String_CodePoints_indexOf'] = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($i_2 = null) use ($s_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_String_CodePoints_length'])((($GLOBALS['Data_String_CodeUnits_take'])($i_2))($s_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_CodeUnits_indexOf'])($p_0))($s_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_fromCharCode
$GLOBALS['Data_String_CodePoints_fromCharCode'] = (function() use (&$__fn) {
$bottom2_0_0 = ((($GLOBALS['Data_Enum_boundedEnumChar'])['Bounded0'])(null))['bottom'];
return ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_CodeUnits_singleton']))(function($x_1 = null) use ($bottom2_0_0) {
  $__num = \func_num_args();
  $v_2_1 = (($GLOBALS['Data_Enum_boundedEnumChar'])['toEnum'])($x_1);
  $__t2 = null;;
  if ((is_object($v_2_1) && (($v_2_1)->{'tag'} === "Just"))) {
$__t2 = ($v_2_1)->{'value0'};
goto end_branch_2;;
};
  if ((is_object($v_2_1) && (($v_2_1)->{'tag'} === "Nothing"))) {
$__t3 = null;;
if (($x_1 < (($GLOBALS['Data_Enum_boundedEnumChar'])['fromEnum'])($bottom2_0_0))) {
$__t3 = ($GLOBALS['Data_Bounded_boundedChar'])['bottom'];
goto end_branch_3;;
};
$__t3 = ($GLOBALS['Data_Bounded_boundedChar'])['top'];
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
})();

// Data_String_CodePoints_singletonFallback
$GLOBALS['Data_String_CodePoints_singletonFallback'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($v_0 <= 65535)) {
$__t0 = ($GLOBALS['Data_String_CodePoints_fromCharCode'])($v_0);
goto end_branch_0;;
};
  $__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(($GLOBALS['Data_String_CodePoints_fromCharCode'])(((($GLOBALS['Data_Semiring_semiringInt'])['add'])(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])['div'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($v_0))(65536)))(1024)))(55296))))(($GLOBALS['Data_String_CodePoints_fromCharCode'])(((($GLOBALS['Data_Semiring_semiringInt'])['add'])(((($GLOBALS['Data_EuclideanRing_euclideanRingInt'])['mod'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($v_0))(65536)))(1024)))(56320)));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_String_CodePoints_fromCodePointArray
$GLOBALS['Data_String_CodePoints_fromCodePointArray'] = ($GLOBALS['Data_String_CodePoints__fromCodePointArray'])($GLOBALS['Data_String_CodePoints_singletonFallback']);

// Data_String_CodePoints_singleton
$GLOBALS['Data_String_CodePoints_singleton'] = ($GLOBALS['Data_String_CodePoints__singleton'])($GLOBALS['Data_String_CodePoints_singletonFallback']);

// Data_String_CodePoints_takeFallback
$GLOBALS['Data_String_CodePoints_takeFallback'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if (($v_0 < 1)) {
$__t2 = "";
goto end_branch_2;;
};
  $v2_2_0 = ($GLOBALS['Data_String_CodePoints_uncons'])($v1_1);
  $__t1 = null;;
  if ((is_object($v2_2_0) && (($v2_2_0)->{'tag'} === "Just"))) {
$__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(($GLOBALS['Data_String_CodePoints_singleton'])((($v2_2_0)->{'value0'})['head'])))((($GLOBALS['Data_String_CodePoints_takeFallback'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($v_0))(1)))((($v2_2_0)->{'value0'})['tail']));
goto end_branch_1;;
};
  $__t1 = $v1_1;
  end_branch_1:;
  $__t2 = $__t1;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_take
$GLOBALS['Data_String_CodePoints_take'] = ($GLOBALS['Data_String_CodePoints__take'])($GLOBALS['Data_String_CodePoints_takeFallback']);

// Data_String_CodePoints_lastIndexOf'
$GLOBALS['Data_String_CodePoints_lastIndexOf__prime__'] = (function() {
  $__fn = function($p_0 = null, $i_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($k_3 = null) use ($s_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_String_CodePoints_length'])((($GLOBALS['Data_String_CodeUnits_take'])($k_3))($s_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($GLOBALS['Data_String_CodeUnits_lastIndexOf__prime__'])($p_0))(($GLOBALS['Data_String_CodeUnits_length'])((($GLOBALS['Data_String_CodePoints_take'])($i_1))($s_2))))($s_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_splitAt
$GLOBALS['Data_String_CodePoints_splitAt'] = (function() {
  $__fn = function($i_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $before_2_0 = (($GLOBALS['Data_String_CodePoints_take'])($i_0))($s_1);
  $__res = ["before" => $before_2_0, "after" => (($GLOBALS['Data_String_CodeUnits_drop'])(($GLOBALS['Data_String_CodeUnits_length'])($before_2_0)))($s_1)];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_eqCodePoint
$GLOBALS['Data_String_CodePoints_eqCodePoint'] = ["eq" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Eq_eqInt'])['eq'])($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_String_CodePoints_ordCodePoint
$GLOBALS['Data_String_CodePoints_ordCodePoint'] = ["compare" => (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Ord_ordInt'])['compare'])($x_0))($y_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_eqCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_CodePoints_drop
$GLOBALS['Data_String_CodePoints_drop'] = (function() {
  $__fn = function($n_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_String_CodeUnits_drop'])(($GLOBALS['Data_String_CodeUnits_length'])((($GLOBALS['Data_String_CodePoints_take'])($n_0))($s_1))))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_indexOf'
$GLOBALS['Data_String_CodePoints_indexOf__prime__'] = (function() {
  $__fn = function($p_0 = null, $i_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $s_prime_3_0 = (($GLOBALS['Data_String_CodeUnits_drop'])(($GLOBALS['Data_String_CodeUnits_length'])((($GLOBALS['Data_String_CodePoints_take'])($i_1))($s_2))))($s_2);
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($k_4 = null) use ($i_1, $s_prime_3_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semiring_semiringInt'])['add'])($i_1))(($GLOBALS['Data_String_CodePoints_length'])((($GLOBALS['Data_String_CodeUnits_take'])($k_4))($s_prime_3_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_String_CodeUnits_indexOf'])($p_0))($s_prime_3_0));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_countTail
$GLOBALS['Data_String_CodePoints_countTail'] = (function() {
  $__fn = function($p_0 = null, $s_1 = null, $accum_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_3_0 = ($GLOBALS['Data_String_CodePoints_uncons'])($s_1);
  $__t1 = null;;
  if (((is_object($v_3_0) && (($v_3_0)->{'tag'} === "Just")) && ($p_0)((($v_3_0)->{'value0'})['head']))) {
$__t1 = ((($GLOBALS['Data_String_CodePoints_countTail'])($p_0))((($v_3_0)->{'value0'})['tail']))(((($GLOBALS['Data_Semiring_semiringInt'])['add'])($accum_2))(1));
goto end_branch_1;;
};
  $__t1 = $accum_2;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_countFallback
$GLOBALS['Data_String_CodePoints_countFallback'] = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_String_CodePoints_countTail'])($p_0))($s_1))(0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_countPrefix
$GLOBALS['Data_String_CodePoints_countPrefix'] = (($GLOBALS['Data_String_CodePoints__countPrefix'])($GLOBALS['Data_String_CodePoints_countFallback']))($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0']);

// Data_String_CodePoints_dropWhile
$GLOBALS['Data_String_CodePoints_dropWhile'] = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_String_CodeUnits_drop'])(($GLOBALS['Data_String_CodeUnits_length'])((($GLOBALS['Data_String_CodePoints_take'])((($GLOBALS['Data_String_CodePoints_countPrefix'])($p_0))($s_1)))($s_1))))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_takeWhile
$GLOBALS['Data_String_CodePoints_takeWhile'] = (function() {
  $__fn = function($p_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_String_CodePoints_take'])((($GLOBALS['Data_String_CodePoints_countPrefix'])($p_0))($s_1)))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_codePointFromChar
$GLOBALS['Data_String_CodePoints_codePointFromChar'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_String_CodePoints_CodePoint']))(($GLOBALS['Data_Enum_boundedEnumChar'])['fromEnum']);

// Data_String_CodePoints_codePointAtFallback
$GLOBALS['Data_String_CodePoints_codePointAtFallback'] = (function() {
  $__fn = function($n_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = ($GLOBALS['Data_String_CodePoints_uncons'])($s_1);
  $__t1 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "Just"))) {
$__t2 = null;;
if (((($GLOBALS['Data_Eq_eqInt'])['eq'])($n_0))(0)) {
$__t2 = new Phpurs_Data1("Just", (($v_2_0)->{'value0'})['head']);
goto end_branch_2;;
};
$__t2 = (($GLOBALS['Data_String_CodePoints_codePointAtFallback'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($n_0))(1)))((($v_2_0)->{'value0'})['tail']);
end_branch_2:;
$__t1 = $__t2;
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

// Data_String_CodePoints_codePointAt
$GLOBALS['Data_String_CodePoints_codePointAt'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($v_0 < 0)) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  switch ($v_0) {
case 0:
$__t0 = match ($v1_1) { "" => new Phpurs_Data0("Nothing"), default => new Phpurs_Data1("Just", ($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0'])($v1_1)) };
goto end_branch_0;;
break;
default:
;
break;
};
  $__t0 = (((((($GLOBALS['Data_String_CodePoints__codePointAt'])($GLOBALS['Data_String_CodePoints_codePointAtFallback']))($GLOBALS['Data_Maybe_Just']))(new Phpurs_Data0("Nothing")))($GLOBALS['Data_String_CodePoints_unsafeCodePointAt0']))($v_0))($v1_1);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_String_CodePoints_boundedCodePoint
$GLOBALS['Data_String_CodePoints_boundedCodePoint'] = ["bottom" => 0, "top" => 1114111, "Ord0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_ordCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_CodePoints_boundedEnumCodePoint
$GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'] = ["cardinality" => ((($GLOBALS['Data_Semiring_semiringInt'])['add'])(1114111))(1), "fromEnum" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "toEnum" => function($n_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(($n_0 >= 0)))(($n_0 <= 1114111))) {
$__t0 = new Phpurs_Data1("Just", $n_0);
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bounded0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_boundedCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Enum1" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_enumCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_String_CodePoints_enumCodePoint
$GLOBALS['Data_String_CodePoints_enumCodePoint'] = ["succ" => (($GLOBALS['Data_Enum_defaultSucc'])(($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'])['toEnum']))(($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'])['fromEnum']), "pred" => (($GLOBALS['Data_Enum_defaultPred'])(($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'])['toEnum']))(($GLOBALS['Data_String_CodePoints_boundedEnumCodePoint'])['fromEnum']), "Ord0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_String_CodePoints_ordCodePoint'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

