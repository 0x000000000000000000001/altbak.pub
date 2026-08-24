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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
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
  $__res = (("(Compose " . (($dictShow_0)->{'show'})($v_1)) . ")");
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
$GLOBALS['Data_Functor_Compose_newtypeCompose'] = (object)["Coercible0" => function($_dollar___unused_0) {
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
  $__res = (object)["map" => function($f_2) use ($dictFunctor1_1, $dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFunctor1_1, $dictFunctor_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})((($dictFunctor1_1)->{'map'})($f_2)))($v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
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
  $eqApp1_3_0 = (object)["eq" => function($x_3) use ($dictEq11_1, $dictEq_2) {
  $__num = \func_num_args();
  $__res = function($y_4) use ($dictEq11_1, $dictEq_2, $x_3) {
  $__num = \func_num_args();
  $__res = (((($dictEq11_1)->{'eq1'})($dictEq_2))($x_3))($y_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["eq" => function($v_4) use ($dictEq1_0, $eqApp1_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq1_0, $eqApp1_3_0, $v_4) {
  $__num = \func_num_args();
  $__res = (((($dictEq1_0)->{'eq1'})($eqApp1_3_0))($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
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
  $__local_var_6_3 = (($dictOrd_5)->{'Eq0'})(null);
  $eqApp2_6_3 = (object)["eq" => function($x_7) use ($__local_var_3_1, $__local_var_6_3) {
  $__num = \func_num_args();
  $__res = function($y_8) use ($__local_var_3_1, $__local_var_6_3, $x_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_1)->{'eq1'})($__local_var_6_3))($x_7))($y_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordApp1_6_3 = (object)["compare" => function($x_7) use ($dictOrd11_2, $dictOrd_5) {
  $__num = \func_num_args();
  $__res = function($y_8) use ($dictOrd11_2, $dictOrd_5, $x_7) {
  $__num = \func_num_args();
  $__res = (((($dictOrd11_2)->{'compare1'})($dictOrd_5))($x_7))($y_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_7) use ($eqApp2_6_3) {
  $__num = \func_num_args();
  $__res = $eqApp2_6_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_7_6 = (($dictOrd_5)->{'Eq0'})(null);
  $eqApp1_8_7 = (object)["eq" => function($x_8) use ($__local_var_4_2, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($y_9) use ($__local_var_4_2, $__local_var_7_6, $x_8) {
  $__num = \func_num_args();
  $__res = (((($__local_var_4_2)->{'eq1'})($__local_var_7_6))($x_8))($y_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $eqCompose3_7_6 = (object)["eq" => function($v_9) use ($__local_var_1_0, $eqApp1_8_7) {
  $__num = \func_num_args();
  $__res = function($v1_10) use ($__local_var_1_0, $eqApp1_8_7, $v_9) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'eq1'})($eqApp1_8_7))($v_9))($v1_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($v_8) use ($dictOrd1_0, $ordApp1_6_3) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($dictOrd1_0, $ordApp1_6_3, $v_8) {
  $__num = \func_num_args();
  $__res = (((($dictOrd1_0)->{'compare1'})($ordApp1_6_3))($v_8))($v1_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_8) use ($eqCompose3_7_6) {
  $__num = \func_num_args();
  $__res = $eqCompose3_7_6;
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
  $eqApp1_3_0 = (object)["eq" => function($x_3) use ($dictEq11_1, $dictEq_2) {
  $__num = \func_num_args();
  $__res = function($y_4) use ($dictEq11_1, $dictEq_2, $x_3) {
  $__num = \func_num_args();
  $__res = (((($dictEq11_1)->{'eq1'})($dictEq_2))($x_3))($y_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($v_4) use ($dictEq1_0, $eqApp1_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq1_0, $eqApp1_3_0, $v_4) {
  $__num = \func_num_args();
  $__res = (((($dictEq1_0)->{'eq1'})($eqApp1_3_0))($v_4))($v1_5);
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
  $__local_var_1_0 = (($dictOrd1_0)->{'Eq10'})(null);
  $__res = function($dictOrd11_2) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd11_2)->{'Eq10'})(null);
  $__local_var_4_2 = (($dictOrd11_2)->{'Eq10'})(null);
  $eq1Compose2_4_2 = (object)["eq1" => function($dictEq_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $eqApp1_6_3 = (object)["eq" => function($x_6) use ($__local_var_4_2, $dictEq_5) {
  $__num = \func_num_args();
  $__res = function($y_7) use ($__local_var_4_2, $dictEq_5, $x_6) {
  $__num = \func_num_args();
  $__res = (((($__local_var_4_2)->{'eq1'})($dictEq_5))($x_6))($y_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($v_7) use ($__local_var_1_0, $eqApp1_6_3) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_1_0, $eqApp1_6_3, $v_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'eq1'})($eqApp1_6_3))($v_7))($v1_8);
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
  $__res = (object)["compare1" => function($dictOrd_5) use ($__local_var_3_1, $dictOrd11_2, $dictOrd1_0) {
  $__num = \func_num_args();
  $__local_var_6_5 = (($dictOrd_5)->{'Eq0'})(null);
  $eqApp2_6_5 = (object)["eq" => function($x_7) use ($__local_var_3_1, $__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($y_8) use ($__local_var_3_1, $__local_var_6_5, $x_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_1)->{'eq1'})($__local_var_6_5))($x_7))($y_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordApp1_6_5 = (object)["compare" => function($x_7) use ($dictOrd11_2, $dictOrd_5) {
  $__num = \func_num_args();
  $__res = function($y_8) use ($dictOrd11_2, $dictOrd_5, $x_7) {
  $__num = \func_num_args();
  $__res = (((($dictOrd11_2)->{'compare1'})($dictOrd_5))($x_7))($y_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_7) use ($eqApp2_6_5) {
  $__num = \func_num_args();
  $__res = $eqApp2_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($v_7) use ($dictOrd1_0, $ordApp1_6_5) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($dictOrd1_0, $ordApp1_6_5, $v_7) {
  $__num = \func_num_args();
  $__res = (((($dictOrd1_0)->{'compare1'})($ordApp1_6_5))($v_7))($v1_8);
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
}, "Eq10" => function($_dollar___unused_5) use ($eq1Compose2_4_2) {
  $__num = \func_num_args();
  $__res = $eq1Compose2_4_2;
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
  $__local_var_2_1 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($dictApply1_3) use ($Functor0_1_0, $__local_var_2_1, $dictApply_0) {
  $__num = \func_num_args();
  $apply_4_2 = ($dictApply1_3)->{'apply'};
  $__local_var_5_3 = (($dictApply1_3)->{'Functor0'})(null);
  $functorCompose2_5_3 = (object)["map" => function($f_6) use ($__local_var_2_1, $__local_var_5_3) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_2_1, $__local_var_5_3, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'map'})((($__local_var_5_3)->{'map'})($f_6)))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($v_6) use ($Functor0_1_0, $apply_4_2, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($Functor0_1_0, $apply_4_2, $dictApply_0, $v_6) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})($apply_4_2))($v_6)))($v1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorCompose2_5_3) {
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
  $__local_var_4_3 = (($__local_var_2_1)->{'Functor0'})(null);
  $applyCompose1_4_3 = function($dictApply1_5) use ($Functor0_3_2, $__local_var_2_1, $__local_var_4_3) {
  $__num = \func_num_args();
  $apply_6_4 = ($dictApply1_5)->{'apply'};
  $__local_var_7_5 = (($dictApply1_5)->{'Functor0'})(null);
  $functorCompose2_7_5 = (object)["map" => function($f_8) use ($__local_var_4_3, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_4_3, $__local_var_7_5, $f_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_3)->{'map'})((($__local_var_7_5)->{'map'})($f_8)))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($v_8) use ($Functor0_3_2, $__local_var_2_1, $apply_6_4) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($Functor0_3_2, $__local_var_2_1, $apply_6_4, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'apply'})(((($Functor0_3_2)->{'map'})($apply_6_4))($v_8)))($v1_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorCompose2_7_5) {
  $__num = \func_num_args();
  $__res = $functorCompose2_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictApplicative1_5) use ($applyCompose1_4_3, $pure_1_0) {
  $__num = \func_num_args();
  $applyCompose2_6_8 = ($applyCompose1_4_3)((($dictApplicative1_5)->{'Apply0'})(null));
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_1_0))(($dictApplicative1_5)->{'pure'})), "Apply0" => function($_dollar___unused_7) use ($applyCompose2_6_8) {
  $__num = \func_num_args();
  $__res = $applyCompose2_6_8;
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
  $functorCompose2_3_1 = (object)["map" => function($f_3) use ($__local_var_1_0, $dictFunctor_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $dictFunctor_2, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})((($dictFunctor_2)->{'map'})($f_3)))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["alt" => function($v_4) use ($dictAlt_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictAlt_0, $v_4) {
  $__num = \func_num_args();
  $__res = ((($dictAlt_0)->{'alt'})($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorCompose2_3_1) {
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
  $functorCompose2_5_3 = (object)["map" => function($f_5) use ($__local_var_3_2, $dictFunctor_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_3_2, $dictFunctor_4, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})((($dictFunctor_4)->{'map'})($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altCompose2_5_3 = (object)["alt" => function($v_6) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_2_1, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'alt'})($v_6))($v1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorCompose2_5_3) {
  $__num = \func_num_args();
  $__res = $functorCompose2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => $empty_1_0, "Alt0" => function($_dollar___unused_6) use ($altCompose2_5_3) {
  $__num = \func_num_args();
  $__res = $altCompose2_5_3;
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
  $__local_var_1_0 = (($dictAlternative_0)->{'Applicative0'})(null);
  $pure_2_1 = ($__local_var_1_0)->{'pure'};
  $__local_var_3_2 = (($__local_var_1_0)->{'Apply0'})(null);
  $Functor0_4_3 = (($__local_var_3_2)->{'Functor0'})(null);
  $__local_var_5_4 = (($__local_var_3_2)->{'Functor0'})(null);
  $applicativeCompose1_4_3 = function($dictApplicative1_6) use ($Functor0_4_3, $__local_var_3_2, $__local_var_5_4, $pure_2_1) {
  $__num = \func_num_args();
  $__local_var_7_5 = (($dictApplicative1_6)->{'Apply0'})(null);
  $apply_8_6 = ($__local_var_7_5)->{'apply'};
  $__local_var_9_7 = (($__local_var_7_5)->{'Functor0'})(null);
  $functorCompose2_9_7 = (object)["map" => function($f_10) use ($__local_var_5_4, $__local_var_9_7) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_5_4, $__local_var_9_7, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'map'})((($__local_var_9_7)->{'map'})($f_10)))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyCompose2_7_5 = (object)["apply" => function($v_10) use ($Functor0_4_3, $__local_var_3_2, $apply_8_6) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($Functor0_4_3, $__local_var_3_2, $apply_8_6, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'apply'})(((($Functor0_4_3)->{'map'})($apply_8_6))($v_10)))($v1_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_10) use ($functorCompose2_9_7) {
  $__num = \func_num_args();
  $__res = $functorCompose2_9_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_2_1))(($dictApplicative1_6)->{'pure'})), "Apply0" => function($_dollar___unused_8) use ($applyCompose2_7_5) {
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
  $__local_var_5_11 = (($dictAlternative_0)->{'Plus1'})(null);
  $empty_6_12 = ($__local_var_5_11)->{'empty'};
  $__local_var_7_13 = (($__local_var_5_11)->{'Alt0'})(null);
  $__local_var_8_14 = (($__local_var_7_13)->{'Functor0'})(null);
  $plusCompose1_8_14 = function($dictFunctor_9) use ($__local_var_7_13, $__local_var_8_14, $empty_6_12) {
  $__num = \func_num_args();
  $functorCompose2_10_15 = (object)["map" => function($f_10) use ($__local_var_8_14, $dictFunctor_9) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_8_14, $dictFunctor_9, $f_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_14)->{'map'})((($dictFunctor_9)->{'map'})($f_10)))($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altCompose2_10_15 = (object)["alt" => function($v_11) use ($__local_var_7_13) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($__local_var_7_13, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_13)->{'alt'})($v_11))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorCompose2_10_15) {
  $__num = \func_num_args();
  $__res = $functorCompose2_10_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => $empty_6_12, "Alt0" => function($_dollar___unused_11) use ($altCompose2_10_15) {
  $__num = \func_num_args();
  $__res = $altCompose2_10_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictApplicative_9) use ($applicativeCompose1_4_3, $plusCompose1_8_14) {
  $__num = \func_num_args();
  $applicativeCompose2_10_18 = ($applicativeCompose1_4_3)($dictApplicative_9);
  $plusCompose2_11_19 = ($plusCompose1_8_14)((((($dictApplicative_9)->{'Apply0'})(null))->{'Functor0'})(null));
  $__res = (object)["Applicative0" => function($_dollar___unused_12) use ($applicativeCompose2_10_18) {
  $__num = \func_num_args();
  $__res = $applicativeCompose2_10_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_12) use ($plusCompose2_11_19) {
  $__num = \func_num_args();
  $__res = $plusCompose2_11_19;
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

