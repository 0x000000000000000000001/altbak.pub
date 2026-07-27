<?php

namespace Data\Functor\Compose;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Newtype, Data.Ord, Data.Semigroup, Data.Show, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Newtype, Data.Ord, Data.Semigroup, Data.Show, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.App/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
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




// Data_Functor_Compose_Compose
function majData_majFunctor_majCompose_majCompose($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_majCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_Compose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_majCompose';

// Data_Functor_Compose_showCompose
function majData_majFunctor_majCompose_showmajCompose($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_showmajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Compose "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})($v_1)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_showCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_showmajCompose';

// Data_Functor_Compose_newtypeCompose
$GLOBALS['Data_Functor_Compose_newtypeCompose'] = (object)["Coercible0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Functor_Compose_functorCompose
function majData_majFunctor_majCompose_functormajCompose($dictFunctor_0, $dictFunctor1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_functormajCompose';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["map" => (function() use ($dictFunctor1_1, $dictFunctor_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictFunctor1_1, $dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)->{'map'})((($dictFunctor1_1)->{'map'})($f_2)))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Functor_Compose_functorCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_functormajCompose';

// Data_Functor_Compose_eqCompose
function majData_majFunctor_majCompose_eqmajCompose($dictEq1_0, $dictEq11_1 = null, $dictEq_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_eqmajCompose';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $eq11_3_0 = (($dictEq11_1)->{'eq1'})($dictEq_2);
  $eq11_3_0 = (($dictEq1_0)->{'eq1'})((object)["eq" => (function() use ($eq11_3_0) {
  $__fn = function($x_4, $y_5 = null) use ($eq11_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_3_0)($x_4))($y_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]);
  $__res = (object)["eq" => (function() use ($eq11_3_0) {
  $__fn = function($v_4, $v1_5 = null) use ($eq11_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_3_0)($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Compose_eqCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_eqmajCompose';

// Data_Functor_Compose_ordCompose
function majData_majFunctor_majCompose_ordmajCompose($dictOrd1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_ordmajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd1_0)->{'Eq10'})(null);
  $__res = function($dictOrd11_2) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd11_2)->{'Eq10'})(null);
  $__local_var_4_2 = (($dictOrd11_2)->{'Eq10'})(null);
  $__res = function($dictOrd_5) use ($__local_var_1_0, $__local_var_3_1, $__local_var_4_2, $dictOrd11_2, $dictOrd1_0) {
  $__num = \func_num_args();
  $compare11_6_3 = (($dictOrd11_2)->{'compare1'})($dictOrd_5);
  $eq11_7_4 = (($__local_var_3_1)->{'eq1'})((($dictOrd_5)->{'Eq0'})(null));
  $eqApp2_8_5 = (object)["eq" => (function() use ($eq11_7_4) {
  $__fn = function($x_8, $y_9 = null) use ($eq11_7_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_7_4)($x_8))($y_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $compare11_6_3 = (($dictOrd1_0)->{'compare1'})((object)["compare" => (function() use ($compare11_6_3) {
  $__fn = function($x_9, $y_10 = null) use ($compare11_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($compare11_6_3)($x_9))($y_10);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_9) use ($eqApp2_8_5) {
  $__num = \func_num_args();
  $__res = $eqApp2_8_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  $eq11_7_7 = (($__local_var_4_2)->{'eq1'})((($dictOrd_5)->{'Eq0'})(null));
  $eq11_7_7 = (($__local_var_1_0)->{'eq1'})((object)["eq" => (function() use ($eq11_7_7) {
  $__fn = function($x_8, $y_9 = null) use ($eq11_7_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_7_7)($x_8))($y_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]);
  $eqCompose3_8_9 = (object)["eq" => (function() use ($eq11_7_7) {
  $__fn = function($v_8, $v1_9 = null) use ($eq11_7_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_7_7)($v_8))($v1_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => (function() use ($compare11_6_3) {
  $__fn = function($v_9, $v1_10 = null) use ($compare11_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($compare11_6_3)($v_9))($v1_10);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_9) use ($eqCompose3_8_9) {
  $__num = \func_num_args();
  $__res = $eqCompose3_8_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_ordCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_ordmajCompose';

// Data_Functor_Compose_eq1Compose
function majData_majFunctor_majCompose_eq1majCompose($dictEq1_0, $dictEq11_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_eq1majCompose';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq1" => function($dictEq_2) use ($dictEq11_1, $dictEq1_0) {
  $__num = \func_num_args();
  $eq11_3_0 = (($dictEq11_1)->{'eq1'})($dictEq_2);
  $__res = (($dictEq1_0)->{'eq1'})((object)["eq" => (function() use ($eq11_3_0) {
  $__fn = function($x_4, $y_5 = null) use ($eq11_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_3_0)($x_4))($y_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Functor_Compose_eq1Compose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_eq1majCompose';

// Data_Functor_Compose_ord1Compose
function majData_majFunctor_majCompose_ord1majCompose($dictOrd1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_ord1majCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ordCompose1_1_0 = ($GLOBALS['Data_Functor_Compose_ordCompose'])($dictOrd1_0);
  $__local_var_2_1 = (($dictOrd1_0)->{'Eq10'})(null);
  $__res = function($dictOrd11_3) use ($__local_var_2_1, $ordCompose1_1_0) {
  $__num = \func_num_args();
  $ordCompose2_4_2 = ($ordCompose1_1_0)($dictOrd11_3);
  $__local_var_5_3 = (($dictOrd11_3)->{'Eq10'})(null);
  $eq1Compose2_6_4 = (object)["eq1" => function($dictEq_6) use ($__local_var_2_1, $__local_var_5_3) {
  $__num = \func_num_args();
  $eq11_7_4 = (($__local_var_5_3)->{'eq1'})($dictEq_6);
  $__res = (($__local_var_2_1)->{'eq1'})((object)["eq" => (function() use ($eq11_7_4) {
  $__fn = function($x_8, $y_9 = null) use ($eq11_7_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($eq11_7_4)($x_8))($y_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare1" => function($dictOrd_7) use ($ordCompose2_4_2) {
  $__num = \func_num_args();
  $__res = (($ordCompose2_4_2)($dictOrd_7))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_7) use ($eq1Compose2_6_4) {
  $__num = \func_num_args();
  $__res = $eq1Compose2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_ord1Compose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_ord1majCompose';

// Data_Functor_Compose_bihoistCompose
function majData_majFunctor_majCompose_bihoistmajCompose($dictFunctor_0, $natF_1 = null, $natG_2 = null, $v_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_bihoistmajCompose';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ($natF_1)(((($dictFunctor_0)->{'map'})($natG_2))($v_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Functor_Compose_bihoistCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_bihoistmajCompose';

// Data_Functor_Compose_applyCompose
function majData_majFunctor_majCompose_applymajCompose($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_applymajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($dictApply1_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $apply1_3_1 = ($dictApply1_2)->{'apply'};
  $__local_var_4_2 = (($dictApply1_2)->{'Functor0'})(null);
  $functorCompose2_5_3 = (object)["map" => (function() use ($Functor0_1_0, $__local_var_4_2) {
  $__fn = function($f_5, $v_6 = null) use ($Functor0_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Functor0_1_0)->{'map'})((($__local_var_4_2)->{'map'})($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["apply" => (function() use ($Functor0_1_0, $apply1_3_1, $dictApply_0) {
  $__fn = function($v_6, $v1_7 = null) use ($Functor0_1_0, $apply1_3_1, $dictApply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})($apply1_3_1))($v_6)))($v1_7);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_6) use ($functorCompose2_5_3) {
  $__num = \func_num_args();
  $__res = $functorCompose2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_applyCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_applymajCompose';

// Data_Functor_Compose_applicativeCompose
function majData_majFunctor_majCompose_applicativemajCompose($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_applicativemajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pure_1_0 = ($dictApplicative_0)->{'pure'};
  $__local_var_2_1 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $__res = function($dictApplicative1_4) use ($Functor0_3_2, $__local_var_2_1, $pure_1_0) {
  $__num = \func_num_args();
  $__local_var_5_3 = (($dictApplicative1_4)->{'Apply0'})(null);
  $apply1_6_4 = ($__local_var_5_3)->{'apply'};
  $__local_var_7_5 = (($__local_var_5_3)->{'Functor0'})(null);
  $functorCompose2_8_6 = (object)["map" => (function() use ($Functor0_3_2, $__local_var_7_5) {
  $__fn = function($f_8, $v_9 = null) use ($Functor0_3_2, $__local_var_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Functor0_3_2)->{'map'})((($__local_var_7_5)->{'map'})($f_8)))($v_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $applyCompose2_7_5 = (object)["apply" => (function() use ($Functor0_3_2, $__local_var_2_1, $apply1_6_4) {
  $__fn = function($v_9, $v1_10 = null) use ($Functor0_3_2, $__local_var_2_1, $apply1_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_1)->{'apply'})(((($Functor0_3_2)->{'map'})($apply1_6_4))($v_9)))($v1_10);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_9) use ($functorCompose2_8_6) {
  $__num = \func_num_args();
  $__res = $functorCompose2_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Functor_Compose_Compose']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_1_0))(($dictApplicative1_4)->{'pure'})), "Apply0" => function($_dollar__unused_8) use ($applyCompose2_7_5) {
  $__num = \func_num_args();
  $__res = $applyCompose2_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_applicativeCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_applicativemajCompose';

// Data_Functor_Compose_altCompose
function majData_majFunctor_majCompose_altmajCompose($dictAlt_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_altmajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictAlt_0)->{'Functor0'})(null);
  $__res = function($dictFunctor_2) use ($__local_var_1_0, $dictAlt_0) {
  $__num = \func_num_args();
  $functorCompose2_3_1 = (object)["map" => (function() use ($__local_var_1_0, $dictFunctor_2) {
  $__fn = function($f_3, $v_4 = null) use ($__local_var_1_0, $dictFunctor_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'map'})((($dictFunctor_2)->{'map'})($f_3)))($v_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["alt" => (function() use ($dictAlt_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictAlt_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictAlt_0)->{'alt'})($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4) use ($functorCompose2_3_1) {
  $__num = \func_num_args();
  $__res = $functorCompose2_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_altCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_altmajCompose';

// Data_Functor_Compose_plusCompose
function majData_majFunctor_majCompose_plusmajCompose($dictPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_plusmajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $empty_1_0 = ($dictPlus_0)->{'empty'};
  $__local_var_2_1 = (($dictPlus_0)->{'Alt0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $__res = function($dictFunctor_4) use ($__local_var_2_1, $__local_var_3_2, $empty_1_0) {
  $__num = \func_num_args();
  $functorCompose2_5_3 = (object)["map" => (function() use ($__local_var_3_2, $dictFunctor_4) {
  $__fn = function($f_5, $v_6 = null) use ($__local_var_3_2, $dictFunctor_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_3_2)->{'map'})((($dictFunctor_4)->{'map'})($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $altCompose2_6_4 = (object)["alt" => (function() use ($__local_var_2_1) {
  $__fn = function($v_6, $v1_7 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_1)->{'alt'})($v_6))($v1_7);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_6) use ($functorCompose2_5_3) {
  $__num = \func_num_args();
  $__res = $functorCompose2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => $empty_1_0, "Alt0" => function($_dollar__unused_7) use ($altCompose2_6_4) {
  $__num = \func_num_args();
  $__res = $altCompose2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_plusCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_plusmajCompose';

// Data_Functor_Compose_alternativeCompose
function majData_majFunctor_majCompose_alternativemajCompose($dictAlternative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majCompose_alternativemajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $applicativeCompose1_1_0 = ($GLOBALS['Data_Functor_Compose_applicativeCompose'])((($dictAlternative_0)->{'Applicative0'})(null));
  $__local_var_2_1 = (($dictAlternative_0)->{'Plus1'})(null);
  $empty_3_2 = ($__local_var_2_1)->{'empty'};
  $__local_var_4_3 = (($__local_var_2_1)->{'Alt0'})(null);
  $__local_var_5_4 = (($__local_var_4_3)->{'Functor0'})(null);
  $plusCompose1_4_3 = function($dictFunctor_6) use ($__local_var_4_3, $__local_var_5_4, $empty_3_2) {
  $__num = \func_num_args();
  $functorCompose2_7_5 = (object)["map" => (function() use ($__local_var_5_4, $dictFunctor_6) {
  $__fn = function($f_7, $v_8 = null) use ($__local_var_5_4, $dictFunctor_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_5_4)->{'map'})((($dictFunctor_6)->{'map'})($f_7)))($v_8);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $altCompose2_8_6 = (object)["alt" => (function() use ($__local_var_4_3) {
  $__fn = function($v_8, $v1_9 = null) use ($__local_var_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_4_3)->{'alt'})($v_8))($v1_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_8) use ($functorCompose2_7_5) {
  $__num = \func_num_args();
  $__res = $functorCompose2_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => $empty_3_2, "Alt0" => function($_dollar__unused_9) use ($altCompose2_8_6) {
  $__num = \func_num_args();
  $__res = $altCompose2_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictApplicative_5) use ($applicativeCompose1_1_0, $plusCompose1_4_3) {
  $__num = \func_num_args();
  $applicativeCompose2_6_8 = ($applicativeCompose1_1_0)($dictApplicative_5);
  $plusCompose2_7_9 = ($plusCompose1_4_3)((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null));
  $__res = (object)["Applicative0" => function($_dollar__unused_8) use ($applicativeCompose2_6_8) {
  $__num = \func_num_args();
  $__res = $applicativeCompose2_6_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_8) use ($plusCompose2_7_9) {
  $__num = \func_num_args();
  $__res = $plusCompose2_7_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Compose_alternativeCompose'] = __NAMESPACE__ . '\\majData_majFunctor_majCompose_alternativemajCompose';

