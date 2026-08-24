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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_Char_Gen_genUnicodeChar
function majData_majChar_majGen_genmajUnicodemajChar($dictMonadGen_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majChar_majGen_genmajUnicodemajChar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_1) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_2_0)->{'value0'};
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = null;;
if (($x_1 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t2 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_2;;
};
$__t2 = $GLOBALS['Data_Bounded_topChar'];
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_0)->{'chooseInt'})(0))(65536));
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
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_1) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_2_0)->{'value0'};
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = null;;
if (($x_1 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t2 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_2;;
};
$__t2 = $GLOBALS['Data_Bounded_topChar'];
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_0)->{'chooseInt'})(48))(57));
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
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_1) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_2_0)->{'value0'};
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = null;;
if (($x_1 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t2 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_2;;
};
$__t2 = $GLOBALS['Data_Bounded_topChar'];
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_0)->{'chooseInt'})(0))(127));
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
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_1) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_2_0)->{'value0'};
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = null;;
if (($x_1 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t2 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_2;;
};
$__t2 = $GLOBALS['Data_Bounded_topChar'];
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_0)->{'chooseInt'})(32))(127));
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
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_1) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_2_0)->{'value0'};
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = null;;
if (($x_1 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t2 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_2;;
};
$__t2 = $GLOBALS['Data_Bounded_topChar'];
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_0)->{'chooseInt'})(65))(90));
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
  $__res = ((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_1) {
  $__num = \func_num_args();
  $v_2_0 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($v_2_0)->{'value0'};
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = null;;
if (($x_1 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t2 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_2;;
};
$__t2 = $GLOBALS['Data_Bounded_topChar'];
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_0)->{'chooseInt'})(97))(122));
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
  $foldableNonEmpty1_1_0 = (object)["foldMap" => function($dictMonoid_1) {
  $__num = \func_num_args();
  $Semigroup0_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $__res = function($f_3) use ($Semigroup0_2_0, $dictMonoid_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($Semigroup0_2_0, $dictMonoid_1, $f_3) {
  $__num = \func_num_args();
  $Semigroup0_5_1 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $__res = ((($Semigroup0_2_0)->{'append'})(($f_3)(($v_4)->{'value0'})))(\Data\Foldable\majData_majFoldable_foldrmajArray(function($x_6) use ($Semigroup0_5_1, $f_3) {
  $__num = \func_num_args();
  $__res = function($acc_7) use ($Semigroup0_5_1, $f_3, $x_6) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_5_1)->{'append'})(($f_3)($x_6)))($acc_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($dictMonoid_1)->{'mempty'}, ($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $f_1) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_1, (($f_1)($b_2))(($v_3)->{'value0'}), ($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $f_1) {
  $__num = \func_num_args();
  $__res = (($f_1)(($v_3)->{'value0'}))(\Data\Foldable\majData_majFoldable_foldrmajArray($f_1, $b_2, ($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_1_0 = (object)["foldMap1" => function($dictSemigroup_2) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictSemigroup_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictSemigroup_2, $f_3) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray(function($s_5) use ($dictSemigroup_2, $f_3) {
  $__num = \func_num_args();
  $__res = function($a1_6) use ($dictSemigroup_2, $f_3, $s_5) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_2)->{'append'})($s_5))(($f_3)($a1_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_3)(($v_4)->{'value0'}), ($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr1" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__local_var_4_3 = ($f_2)(($v_3)->{'value0'});
  $__local_var_5_4 = \Data\Foldable\majData_majFoldable_foldrmajArray(function($a1_5) use ($f_2) {
  $__num = \func_num_args();
  $__local_var_6_4 = ($f_2)($a1_5);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_7) use ($__local_var_6_4, $a1_5) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v2_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = $a1_5;
goto end_branch_5;;
};
  if ($v2_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ($__local_var_6_4)(($v2_7)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), ($v_3)->{'value1'});
  $__t7 = null;;
  if ($__local_var_5_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t7 = ($v_3)->{'value0'};
goto end_branch_7;;
};
  if ($__local_var_5_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t7 = ($__local_var_4_3)(($__local_var_5_4)->{'value0'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = \Data\Foldable\majData_majFoldable_foldlmajArray($f_2, ($v_3)->{'value0'}, ($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_2) use ($foldableNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_9 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_2) {
  $__num = \func_num_args();
  $v_3_9 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_2);
  $__t10 = null;;
  if ($v_3_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t10 = ($v_3_9)->{'value0'};
goto end_branch_10;;
};
  if ($v_3_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = null;;
if (($x_2 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t11 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_11;;
};
$__t11 = $GLOBALS['Data_Bounded_topChar'];
end_branch_11:;
$__t10 = $__t11;
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_0)->{'chooseInt'})(97))(122)), [((((((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($x_2) {
  $__num = \func_num_args();
  $v_3_12 = \Data\Enum\majData_majEnum_charmajTomajEnum($x_2);
  $__t13 = null;;
  if ($v_3_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($v_3_12)->{'value0'};
goto end_branch_13;;
};
  if ($v_3_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = null;;
if (($x_2 < \Data\Enum\majData_majEnum_tomajCharmajCode($GLOBALS['Data_Bounded_bottomChar']))) {
$__t14 = $GLOBALS['Data_Bounded_bottomChar'];
goto end_branch_14;;
};
$__t14 = $GLOBALS['Data_Bounded_topChar'];
end_branch_14:;
$__t13 = $__t14;
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictMonadGen_0)->{'chooseInt'})(65))(90))]);
  $__res = ((((((($dictMonadGen_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'bind'})(((($dictMonadGen_0)->{'chooseInt'})(0))(((((((($__local_var_1_0)->{'Foldable0'})(null))->{'foldl'})(function($c_3) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($c_3) {
  $__num = \func_num_args();
  $__res = (1 + $c_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0))($__local_var_2_9) - 1))))(function($n_3) use ($__local_var_1_0, $__local_var_2_9) {
  $__num = \func_num_args();
  $go__go_4_16 = null;
  $go__go_4_16 = (function() use ($__local_var_1_0, $__local_var_2_9, &$go__go_4_16) {
  $__fn = function(int $v_5, $v1_6 = null) use ($__local_var_1_0, $__local_var_2_9, &$go__go_4_16, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_16_16_v_5 = $v_5;
  $__tco_var_go__go_4_16_16_v1_6 = $v1_6;
  tco_loop_go__go_4_16_16:;
  $v_5 = $__tco_var_go__go_4_16_16_v_5;
  $v1_6 = $__tco_var_go__go_4_16_16_v1_6;
  $__t16 = null;;
  if ($v1_6 instanceof \Control\Monad\Gen\Control_Monad_Gen_Cons) {
$__t19 = null;;
if (($v1_6)->{'value1'} instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t19 = ($v1_6)->{'value0'};
goto end_branch_19;;
};
if (($v_5 <= 0)) {
$__t19 = ($v1_6)->{'value0'};
goto end_branch_19;;
};
$__tco_17 = ($v_5 - 1);
$__tco_18 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_16_16_v_5 = $__tco_17;
$__tco_var_go__go_4_16_16_v1_6 = $__tco_18;
goto tco_loop_go__go_4_16_16;;
$__t19 = null;
end_branch_19:;
$__t16 = $__t19;
goto end_branch_16;;
};
  if ($v1_6 instanceof \Control\Monad\Gen\Control_Monad_Gen_Nil) {
$__t16 = (((($__local_var_1_0)->{'foldMap1'})($GLOBALS['Data_Semigroup_Last_semigroupLast']))(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__local_var_2_9);
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_4_16)($n_3))((((((($__local_var_1_0)->{'Foldable0'})(null))->{'foldr'})($GLOBALS['Control_Monad_Gen_Cons']))(new \Control\Monad\Gen\Control_Monad_Gen_Nil()))($__local_var_2_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Char_Gen_genAlpha'] = __NAMESPACE__ . '\\majData_majChar_majGen_genmajAlpha';

