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
\PhpursThunks::$thunks['Data_Char_Gen_toEnumWithDefaults'] = function() { $v = (($GLOBALS['Data_Enum_toEnumWithDefaults'] ?? \PhpursThunks::eval('Data_Enum_toEnumWithDefaults')))(($GLOBALS['Data_Enum_boundedEnumChar'] ?? \PhpursThunks::eval('Data_Enum_boundedEnumChar'))); return $v; };
\PhpursThunks::$thunks['Data_Char_Gen_foldable1NonEmpty'] = function() { $v = (($GLOBALS['Data_NonEmpty_foldable1NonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldable1NonEmpty')))(($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'))); return $v; };
\PhpursThunks::$thunks['Data_Char_Gen_genUnicodeChar'] = function() { $v = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'] ?? \PhpursThunks::eval('Data_Char_Gen_toEnumWithDefaults')))(($GLOBALS['Data_Bounded_bottomChar'] ?? \PhpursThunks::eval('Data_Bounded_bottomChar'))))(($GLOBALS['Data_Bounded_topChar'] ?? \PhpursThunks::eval('Data_Bounded_topChar')))))(((($dictMonadGen_0)->{'chooseInt'})(0))(65536));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Char_Gen_genDigitChar'] = function() { $v = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'] ?? \PhpursThunks::eval('Data_Char_Gen_toEnumWithDefaults')))(($GLOBALS['Data_Bounded_bottomChar'] ?? \PhpursThunks::eval('Data_Bounded_bottomChar'))))(($GLOBALS['Data_Bounded_topChar'] ?? \PhpursThunks::eval('Data_Bounded_topChar')))))(((($dictMonadGen_0)->{'chooseInt'})(48))(57));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Char_Gen_genAsciiChar__prime__'] = function() { $v = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'] ?? \PhpursThunks::eval('Data_Char_Gen_toEnumWithDefaults')))(($GLOBALS['Data_Bounded_bottomChar'] ?? \PhpursThunks::eval('Data_Bounded_bottomChar'))))(($GLOBALS['Data_Bounded_topChar'] ?? \PhpursThunks::eval('Data_Bounded_topChar')))))(((($dictMonadGen_0)->{'chooseInt'})(0))(127));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Char_Gen_genAsciiChar'] = function() { $v = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'] ?? \PhpursThunks::eval('Data_Char_Gen_toEnumWithDefaults')))(($GLOBALS['Data_Bounded_bottomChar'] ?? \PhpursThunks::eval('Data_Bounded_bottomChar'))))(($GLOBALS['Data_Bounded_topChar'] ?? \PhpursThunks::eval('Data_Bounded_topChar')))))(((($dictMonadGen_0)->{'chooseInt'})(32))(127));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Char_Gen_genAlphaUppercase'] = function() { $v = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'] ?? \PhpursThunks::eval('Data_Char_Gen_toEnumWithDefaults')))(($GLOBALS['Data_Bounded_bottomChar'] ?? \PhpursThunks::eval('Data_Bounded_bottomChar'))))(($GLOBALS['Data_Bounded_topChar'] ?? \PhpursThunks::eval('Data_Bounded_topChar')))))(((($dictMonadGen_0)->{'chooseInt'})(65))(90));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Char_Gen_genAlphaLowercase'] = function() { $v = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'})(((($GLOBALS['Data_Char_Gen_toEnumWithDefaults'] ?? \PhpursThunks::eval('Data_Char_Gen_toEnumWithDefaults')))(($GLOBALS['Data_Bounded_bottomChar'] ?? \PhpursThunks::eval('Data_Bounded_bottomChar'))))(($GLOBALS['Data_Bounded_topChar'] ?? \PhpursThunks::eval('Data_Bounded_topChar')))))(((($dictMonadGen_0)->{'chooseInt'})(97))(122));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Char_Gen_genAlpha'] = function() { $v = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Control_Monad_Gen_oneOf'] ?? \PhpursThunks::eval('Control_Monad_Gen_oneOf')))($dictMonadGen_0))(($GLOBALS['Data_Char_Gen_foldable1NonEmpty'] ?? \PhpursThunks::eval('Data_Char_Gen_foldable1NonEmpty'))))(new Phpurs_Data2("NonEmpty", (($GLOBALS['Data_Char_Gen_genAlphaLowercase'] ?? \PhpursThunks::eval('Data_Char_Gen_genAlphaLowercase')))($dictMonadGen_0), [(($GLOBALS['Data_Char_Gen_genAlphaUppercase'] ?? \PhpursThunks::eval('Data_Char_Gen_genAlphaUppercase')))($dictMonadGen_0)]));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };











