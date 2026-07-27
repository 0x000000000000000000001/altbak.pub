<?php

namespace Data\Bounded\Generic;

// ALL IMPORTS: Data.Bounded, Data.Bounded.Generic, Data.Generic.Rep, Prim
// TO REQUIRE: Data.Bounded, Data.Bounded.Generic, Data.Generic.Rep
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Bounded.Generic/index.php';
require_once __DIR__ . '/../Data.Generic.Rep/index.php';

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




// Data_Bounded_Generic_genericTopNoArguments
$GLOBALS['Data_Bounded_Generic_genericTopNoArguments'] = (object)["genericTop__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_NoArguments()];

// Data_Bounded_Generic_genericTopArgument
function majData_majBounded_majGeneric_genericmajTopmajArgument($dictBounded_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajTopmajArgument';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericTop__prime__" => ($dictBounded_0)->{'top'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericTopArgument'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajTopmajArgument';

// Data_Bounded_Generic_genericTop'
function majData_majBounded_majGeneric_genericmajTop__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajTop__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericTop__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericTop__prime__'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajTop__prime__';

// Data_Bounded_Generic_genericTopConstructor
function majData_majBounded_majGeneric_genericmajTopmajConstructor($dictGenericTop_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajTopmajConstructor';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericTop__prime__" => ($dictGenericTop_0)->{'genericTop__prime__'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericTopConstructor'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajTopmajConstructor';

// Data_Bounded_Generic_genericTopProduct
function majData_majBounded_majGeneric_genericmajTopmajProduct($dictGenericTop_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajTopmajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $genericTop_prime1_1_0 = ($dictGenericTop_0)->{'genericTop__prime__'};
  $__res = function($dictGenericTop1_2) use ($genericTop_prime1_1_0) {
  $__num = \func_num_args();
  $__res = (object)["genericTop__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_Product($genericTop_prime1_1_0, ($dictGenericTop1_2)->{'genericTop__prime__'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericTopProduct'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajTopmajProduct';

// Data_Bounded_Generic_genericTopSum
function majData_majBounded_majGeneric_genericmajTopmajSum($dictGenericTop_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajTopmajSum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericTop__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_Inr(($dictGenericTop_0)->{'genericTop__prime__'})];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericTopSum'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajTopmajSum';

// Data_Bounded_Generic_genericTop
function majData_majBounded_majGeneric_genericmajTop($dictGeneric_0, $dictGenericTop_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajTop';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictGeneric_0)->{'to'})(($dictGenericTop_1)->{'genericTop__prime__'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericTop'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajTop';

// Data_Bounded_Generic_genericBottomNoArguments
$GLOBALS['Data_Bounded_Generic_genericBottomNoArguments'] = (object)["genericBottom__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_NoArguments()];

// Data_Bounded_Generic_genericBottomArgument
function majData_majBounded_majGeneric_genericmajBottommajArgument($dictBounded_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajBottommajArgument';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericBottom__prime__" => ($dictBounded_0)->{'bottom'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericBottomArgument'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajBottommajArgument';

// Data_Bounded_Generic_genericBottom'
function majData_majBounded_majGeneric_genericmajBottom__prime__($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajBottom__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'genericBottom__prime__'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericBottom__prime__'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajBottom__prime__';

// Data_Bounded_Generic_genericBottomConstructor
function majData_majBounded_majGeneric_genericmajBottommajConstructor($dictGenericBottom_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajBottommajConstructor';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericBottom__prime__" => ($dictGenericBottom_0)->{'genericBottom__prime__'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericBottomConstructor'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajBottommajConstructor';

// Data_Bounded_Generic_genericBottomProduct
function majData_majBounded_majGeneric_genericmajBottommajProduct($dictGenericBottom_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajBottommajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $genericBottom_prime1_1_0 = ($dictGenericBottom_0)->{'genericBottom__prime__'};
  $__res = function($dictGenericBottom1_2) use ($genericBottom_prime1_1_0) {
  $__num = \func_num_args();
  $__res = (object)["genericBottom__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_Product($genericBottom_prime1_1_0, ($dictGenericBottom1_2)->{'genericBottom__prime__'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericBottomProduct'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajBottommajProduct';

// Data_Bounded_Generic_genericBottomSum
function majData_majBounded_majGeneric_genericmajBottommajSum($dictGenericBottom_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajBottommajSum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["genericBottom__prime__" => new \Data\Generic\Rep\Data_Generic_Rep_Inl(($dictGenericBottom_0)->{'genericBottom__prime__'})];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericBottomSum'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajBottommajSum';

// Data_Bounded_Generic_genericBottom
function majData_majBounded_majGeneric_genericmajBottom($dictGeneric_0, $dictGenericBottom_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBounded_majGeneric_genericmajBottom';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictGeneric_0)->{'to'})(($dictGenericBottom_1)->{'genericBottom__prime__'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bounded_Generic_genericBottom'] = __NAMESPACE__ . '\\majData_majBounded_majGeneric_genericmajBottom';

