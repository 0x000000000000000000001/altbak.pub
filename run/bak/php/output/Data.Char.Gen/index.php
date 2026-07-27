<?php

namespace Data\Char\Gen;

// ALL IMPORTS: Control.Monad.Gen, Control.Monad.Gen.Class, Data.Bounded, Data.Char.Gen, Data.Enum, Data.Foldable, Data.Functor, Data.NonEmpty, Prelude, Prim
// TO REQUIRE: Control.Monad.Gen, Control.Monad.Gen.Class, Data.Bounded, Data.Char.Gen, Data.Enum, Data.Foldable, Data.Functor, Data.NonEmpty, Prelude
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Class/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Char.Gen/index.php';
require_once __DIR__ . '/../Data.Enum/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
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




// Data_Char_Gen_toEnumWithDefaults
$GLOBALS['Data_Char_Gen_toEnumWithDefaults'] = ($GLOBALS['Data_Enum_toEnumWithDefaults'])($GLOBALS['Data_Enum_boundedEnumChar']);

// Data_Char_Gen_foldable1NonEmpty
$GLOBALS['Data_Char_Gen_foldable1NonEmpty'] = \Data\NonEmpty\majData_majNonmajEmpty_foldable1majNonmajEmpty($GLOBALS['Data_Foldable_foldableArray']);

// Data_Char_Gen_genUnicodeChar
function majData_majChar_majGen_genmajUnicodemajChar($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majChar_majGen_genmajUnicodemajChar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'])(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'}))(($GLOBALS['Data_Bounded_boundedChar'])->{'top'})))(((($dictMonadGen_0)->{'chooseInt'})(0))(65536));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Char_Gen_genUnicodeChar'] = __NAMESPACE__ . '\\majData_majChar_majGen_genmajUnicodemajChar';

// Data_Char_Gen_genDigitChar
function majData_majChar_majGen_genmajDigitmajChar($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majChar_majGen_genmajDigitmajChar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'])(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'}))(($GLOBALS['Data_Bounded_boundedChar'])->{'top'})))(((($dictMonadGen_0)->{'chooseInt'})(48))(57));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Char_Gen_genDigitChar'] = __NAMESPACE__ . '\\majData_majChar_majGen_genmajDigitmajChar';

// Data_Char_Gen_genAsciiChar'
function majData_majChar_majGen_genmajAsciimajChar__prime__($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majChar_majGen_genmajAsciimajChar__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'])(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'}))(($GLOBALS['Data_Bounded_boundedChar'])->{'top'})))(((($dictMonadGen_0)->{'chooseInt'})(0))(127));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Char_Gen_genAsciiChar__prime__'] = __NAMESPACE__ . '\\majData_majChar_majGen_genmajAsciimajChar__prime__';

// Data_Char_Gen_genAsciiChar
function majData_majChar_majGen_genmajAsciimajChar($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majChar_majGen_genmajAsciimajChar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'])(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'}))(($GLOBALS['Data_Bounded_boundedChar'])->{'top'})))(((($dictMonadGen_0)->{'chooseInt'})(32))(127));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Char_Gen_genAsciiChar'] = __NAMESPACE__ . '\\majData_majChar_majGen_genmajAsciimajChar';

// Data_Char_Gen_genAlphaUppercase
function majData_majChar_majGen_genmajAlphamajUppercase($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majChar_majGen_genmajAlphamajUppercase';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'])(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'}))(($GLOBALS['Data_Bounded_boundedChar'])->{'top'})))(((($dictMonadGen_0)->{'chooseInt'})(65))(90));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Char_Gen_genAlphaUppercase'] = __NAMESPACE__ . '\\majData_majChar_majGen_genmajAlphamajUppercase';

// Data_Char_Gen_genAlphaLowercase
function majData_majChar_majGen_genmajAlphamajLowercase($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majChar_majGen_genmajAlphamajLowercase';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'])(($GLOBALS['Data_Bounded_boundedChar'])->{'bottom'}))(($GLOBALS['Data_Bounded_boundedChar'])->{'top'})))(((($dictMonadGen_0)->{'chooseInt'})(97))(122));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Char_Gen_genAlphaLowercase'] = __NAMESPACE__ . '\\majData_majChar_majGen_genmajAlphamajLowercase';

// Data_Char_Gen_genAlpha
function majData_majChar_majGen_genmajAlpha($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majChar_majGen_genmajAlpha';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Control\Monad\Gen\majControl_majMonad_majGen_onemajOf($dictMonadGen_0, $GLOBALS['Data_Char_Gen_foldable1NonEmpty'], new \Data\NonEmpty\Data_NonEmpty_NonEmpty(\Data\Char\Gen\majData_majChar_majGen_genmajAlphamajLowercase($dictMonadGen_0), [\Data\Char\Gen\majData_majChar_majGen_genmajAlphamajUppercase($dictMonadGen_0)]));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Char_Gen_genAlpha'] = __NAMESPACE__ . '\\majData_majChar_majGen_genmajAlpha';

