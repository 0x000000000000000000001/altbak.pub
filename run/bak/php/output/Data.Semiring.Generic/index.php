<?php

namespace Data\Semiring\Generic;

// ALL IMPORTS: Data.Function, Data.Generic.Rep, Data.Semiring, Data.Semiring.Generic, Prelude, Prim
// TO REQUIRE: Data.Function, Data.Generic.Rep, Data.Semiring, Data.Semiring.Generic, Prelude
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Semiring.Generic/index.php';
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




// Data_Semiring_Generic_genericZero'
function majData_majSemiring_majGeneric_genericmajZero__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajZero__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericZero__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericZero__prime__'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajZero__prime__';

// Data_Semiring_Generic_genericZero
function majData_majSemiring_majGeneric_genericmajZero($dictGeneric_0, $dictGenericSemiring_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajZero';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictGeneric_0)->{'to'})(($dictGenericSemiring_1)->{'genericZero__prime__'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericZero'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajZero';

// Data_Semiring_Generic_genericSemiringNoArguments
$GLOBALS['Data_Semiring_Generic_genericSemiringNoArguments'] = (object)["genericAdd__prime__" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Generic\Rep\Data_Generic_Rep_NoArguments();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericZero__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_NoArguments(), "genericMul__prime__" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Generic\Rep\Data_Generic_Rep_NoArguments();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericOne__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_NoArguments()];

// Data_Semiring_Generic_genericSemiringArgument
function majData_majSemiring_majGeneric_genericmajSemiringmajArgument($dictSemiring_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajSemiringmajArgument';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericAdd__prime__" => (function() use ($dictSemiring_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictSemiring_0)->{'add'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericZero__prime__" => ($dictSemiring_0)->{'zero'}, "genericMul__prime__" => (function() use ($dictSemiring_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictSemiring_0)->{'mul'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericOne__prime__" => ($dictSemiring_0)->{'one'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericSemiringArgument'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajSemiringmajArgument';

// Data_Semiring_Generic_genericOne'
function majData_majSemiring_majGeneric_genericmajOne__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajOne__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericOne__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericOne__prime__'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajOne__prime__';

// Data_Semiring_Generic_genericOne
function majData_majSemiring_majGeneric_genericmajOne($dictGeneric_0, $dictGenericSemiring_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajOne';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictGeneric_0)->{'to'})(($dictGenericSemiring_1)->{'genericOne__prime__'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericOne'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajOne';

// Data_Semiring_Generic_genericMul'
function majData_majSemiring_majGeneric_genericmajMul__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajMul__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericMul__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericMul__prime__'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajMul__prime__';

// Data_Semiring_Generic_genericMul
function majData_majSemiring_majGeneric_genericmajMul($dictGeneric_0, $dictGenericSemiring_1 = null, $x_2 = null, $y_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajMul';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($dictGeneric_0)->{'to'})(((($dictGenericSemiring_1)->{'genericMul__prime__'})((($dictGeneric_0)->{'from'})($x_2)))((($dictGeneric_0)->{'from'})($y_3)));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericMul'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajMul';

// Data_Semiring_Generic_genericAdd'
function majData_majSemiring_majGeneric_genericmajAdd__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajAdd__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericAdd__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericAdd__prime__'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajAdd__prime__';

// Data_Semiring_Generic_genericSemiringConstructor
function majData_majSemiring_majGeneric_genericmajSemiringmajConstructor($dictGenericSemiring_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajSemiringmajConstructor';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericAdd__prime__" => (function() use ($dictGenericSemiring_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictGenericSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictGenericSemiring_0)->{'genericAdd__prime__'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericZero__prime__" => ($dictGenericSemiring_0)->{'genericZero__prime__'}, "genericMul__prime__" => (function() use ($dictGenericSemiring_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictGenericSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictGenericSemiring_0)->{'genericMul__prime__'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericOne__prime__" => ($dictGenericSemiring_0)->{'genericOne__prime__'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericSemiringConstructor'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajSemiringmajConstructor';

// Data_Semiring_Generic_genericSemiringProduct
function majData_majSemiring_majGeneric_genericmajSemiringmajProduct($dictGenericSemiring_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajSemiringmajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $genericZero_prime1_1_0 = ($dictGenericSemiring_0)->{'genericZero__prime__'};
  $genericOne_prime1_2_1 = ($dictGenericSemiring_0)->{'genericOne__prime__'};
  $__res = function($dictGenericSemiring1_3) use ($dictGenericSemiring_0, $genericOne_prime1_2_1, $genericZero_prime1_1_0) {
  $__num = \func_num_args();
  $__res = (object)["genericAdd__prime__" => (function() use ($dictGenericSemiring1_3, $dictGenericSemiring_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictGenericSemiring1_3, $dictGenericSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Generic\Rep\Data_Generic_Rep_Product(((($dictGenericSemiring_0)->{'genericAdd__prime__'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictGenericSemiring1_3)->{'genericAdd__prime__'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericZero__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_Product($genericZero_prime1_1_0, ($dictGenericSemiring1_3)->{'genericZero__prime__'}), "genericMul__prime__" => (function() use ($dictGenericSemiring1_3, $dictGenericSemiring_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictGenericSemiring1_3, $dictGenericSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Generic\Rep\Data_Generic_Rep_Product(((($dictGenericSemiring_0)->{'genericMul__prime__'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictGenericSemiring1_3)->{'genericMul__prime__'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericOne__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_Product($genericOne_prime1_2_1, ($dictGenericSemiring1_3)->{'genericOne__prime__'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericSemiringProduct'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajSemiringmajProduct';

// Data_Semiring_Generic_genericAdd
function majData_majSemiring_majGeneric_genericmajAdd($dictGeneric_0, $dictGenericSemiring_1 = null, $x_2 = null, $y_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemiring_majGeneric_genericmajAdd';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($dictGeneric_0)->{'to'})(((($dictGenericSemiring_1)->{'genericAdd__prime__'})((($dictGeneric_0)->{'from'})($x_2)))((($dictGeneric_0)->{'from'})($y_3)));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Semiring_Generic_genericAdd'] = __NAMESPACE__ . '\\majData_majSemiring_majGeneric_genericmajAdd';

