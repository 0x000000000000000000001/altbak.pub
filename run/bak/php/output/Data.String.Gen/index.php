<?php

namespace Data\String\Gen;

// ALL IMPORTS: Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Rec.Class, Data.Char.Gen, Data.Function, Data.Functor, Data.Ord, Data.String.CodeUnits, Data.String.Gen, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Rec.Class, Data.Char.Gen, Data.Function, Data.Functor, Data.Ord, Data.String.CodeUnits, Data.String.Gen, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Data.Char.Gen/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.String.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.Gen/index.php';
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




// Data_String_Gen_max
$GLOBALS['Data_String_Gen_max'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($x_1, $y_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_3_1 = (($__local_var_0_0)($x_1))($y_2);
  $__t2 = null;;
  if ($v_3_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t2 = $y_2;
goto end_branch_2;;
};
  if ($v_3_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = $x_1;
goto end_branch_2;;
};
  if ($v_3_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t2 = $x_1;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_String_Gen_genString
function majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Bind1_2_0 = (((($dictMonadGen_1)->{'Monad0'})(null))->{'Bind1'})(null);
  $unfoldable1_3_1 = ((($GLOBALS['Control_Monad_Gen_unfoldable'])($dictMonadRec_0))($dictMonadGen_1))($GLOBALS['Data_Unfoldable_unfoldableArray']);
  $__res = function($genChar_4) use ($Bind1_2_0, $dictMonadGen_1, $unfoldable1_3_1) {
  $__num = \func_num_args();
  $__res = (($dictMonadGen_1)->{'sized'})(function($size_5) use ($Bind1_2_0, $dictMonadGen_1, $genChar_4, $unfoldable1_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})(((($dictMonadGen_1)->{'chooseInt'})(1))((($GLOBALS['Data_String_Gen_max'])(1))($size_5))))(function($newSize_6) use ($Bind1_2_0, $dictMonadGen_1, $genChar_4, $unfoldable1_3_1) {
  $__num = \func_num_args();
  $__res = ((($dictMonadGen_1)->{'resize'})(function($v_7) use ($newSize_6) {
  $__num = \func_num_args();
  $__res = $newSize_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((((((($Bind1_2_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_String_CodeUnits_fromCharArray']))(($unfoldable1_3_1)($genChar_4)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajString';

// Data_String_Gen_genUnicodeString
function majData_majString_majGen_genmajUnicodemajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajUnicodemajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Gen\majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1, \Data\Char\Gen\majData_majChar_majGen_genmajUnicodemajChar($dictMonadGen_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genUnicodeString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajUnicodemajString';

// Data_String_Gen_genDigitString
function majData_majString_majGen_genmajDigitmajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajDigitmajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Gen\majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1, \Data\Char\Gen\majData_majChar_majGen_genmajDigitmajChar($dictMonadGen_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genDigitString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajDigitmajString';

// Data_String_Gen_genAsciiString'
function majData_majString_majGen_genmajAsciimajString__prime__($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAsciimajString__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Gen\majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1, ($GLOBALS['Data_Char_Gen_genAsciiChar__prime__'])($dictMonadGen_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAsciiString__prime__'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAsciimajString__prime__';

// Data_String_Gen_genAsciiString
function majData_majString_majGen_genmajAsciimajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAsciimajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Gen\majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1, \Data\Char\Gen\majData_majChar_majGen_genmajAsciimajChar($dictMonadGen_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAsciiString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAsciimajString';

// Data_String_Gen_genAlphaUppercaseString
function majData_majString_majGen_genmajAlphamajUppercasemajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAlphamajUppercasemajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Gen\majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1, \Data\Char\Gen\majData_majChar_majGen_genmajAlphamajUppercase($dictMonadGen_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAlphaUppercaseString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAlphamajUppercasemajString';

// Data_String_Gen_genAlphaString
function majData_majString_majGen_genmajAlphamajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAlphamajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Gen\majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1, \Data\Char\Gen\majData_majChar_majGen_genmajAlpha($dictMonadGen_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAlphaString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAlphamajString';

// Data_String_Gen_genAlphaLowercaseString
function majData_majString_majGen_genmajAlphamajLowercasemajString($dictMonadRec_0, $dictMonadGen_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majGen_genmajAlphamajLowercasemajString';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Gen\majData_majString_majGen_genmajString($dictMonadRec_0, $dictMonadGen_1, \Data\Char\Gen\majData_majChar_majGen_genmajAlphamajLowercase($dictMonadGen_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Gen_genAlphaLowercaseString'] = __NAMESPACE__ . '\\majData_majString_majGen_genmajAlphamajLowercasemajString';

