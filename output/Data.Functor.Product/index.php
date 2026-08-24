<?php

namespace Data\Functor\Product;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Semigroupoid, Data.Bifunctor, Data.Eq, Data.Functor, Data.Functor.Product, Data.HeytingAlgebra, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Semigroupoid, Data.Bifunctor, Data.Eq, Data.Functor, Data.Functor.Product, Data.HeytingAlgebra, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Product/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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




// Data_Functor_Product_Product
function majData_majFunctor_majProduct_majProduct($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_majProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Product_Product'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_majProduct';

// Data_Functor_Product_showProduct
function majData_majFunctor_majProduct_showmajProduct($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_showmajProduct';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = (((("(product " . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow1_1)->{'show'})(($v_2)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Functor_Product_showProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_showmajProduct';

// Data_Functor_Product_product
function majData_majFunctor_majProduct_product($fa_0, $ga_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_product';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($fa_0, $ga_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Functor_Product_product'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_product';

// Data_Functor_Product_newtypeProduct
$GLOBALS['Data_Functor_Product_newtypeProduct'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Functor_Product_functorProduct
function majData_majFunctor_majProduct_functormajProduct($dictFunctor_0, $dictFunctor1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_functormajProduct';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["map" => function($f_2) use ($dictFunctor1_1, $dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFunctor1_1, $dictFunctor_0, $f_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictFunctor_0)->{'map'})($f_2))(($v_3)->{'value0'}), ((($dictFunctor1_1)->{'map'})($f_2))(($v_3)->{'value1'}));
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
$GLOBALS['Data_Functor_Product_functorProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_functormajProduct';

// Data_Functor_Product_eq1Product
function majData_majFunctor_majProduct_eq1majProduct($dictEq1_0, $dictEq11_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_eq1majProduct';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq1" => function($dictEq_2) use ($dictEq11_1, $dictEq1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictEq11_1, $dictEq1_0, $dictEq_2) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($dictEq11_1, $dictEq1_0, $dictEq_2, $v_3) {
  $__num = \func_num_args();
  $__res = ((((($dictEq1_0)->{'eq1'})($dictEq_2))(($v_3)->{'value0'}))(($v1_4)->{'value0'}) && (((($dictEq11_1)->{'eq1'})($dictEq_2))(($v_3)->{'value1'}))(($v1_4)->{'value1'}));
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
$GLOBALS['Data_Functor_Product_eq1Product'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_eq1majProduct';

// Data_Functor_Product_eqProduct
function majData_majFunctor_majProduct_eqmajProduct($dictEq1_0, $dictEq11_1 = null, $dictEq_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_eqmajProduct';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (object)["eq" => function($v_3) use ($dictEq11_1, $dictEq1_0, $dictEq_2) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($dictEq11_1, $dictEq1_0, $dictEq_2, $v_3) {
  $__num = \func_num_args();
  $__res = ((((($dictEq1_0)->{'eq1'})($dictEq_2))(($v_3)->{'value0'}))(($v1_4)->{'value0'}) && (((($dictEq11_1)->{'eq1'})($dictEq_2))(($v_3)->{'value1'}))(($v1_4)->{'value1'}));
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
$GLOBALS['Data_Functor_Product_eqProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_eqmajProduct';

// Data_Functor_Product_ord1Product
function majData_majFunctor_majProduct_ord1majProduct($dictOrd1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_ord1majProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd1_0)->{'Eq10'})(null);
  $__res = function($dictOrd11_2) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd11_2)->{'Eq10'})(null);
  $eq1Product2_3_1 = (object)["eq1" => function($dictEq_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_1_0, $__local_var_3_1, $dictEq_4) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_1_0, $__local_var_3_1, $dictEq_4, $v_5) {
  $__num = \func_num_args();
  $__res = ((((($__local_var_1_0)->{'eq1'})($dictEq_4))(($v_5)->{'value0'}))(($v1_6)->{'value0'}) && (((($__local_var_3_1)->{'eq1'})($dictEq_4))(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
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
  $__res = (object)["compare1" => function($dictOrd_4) use ($dictOrd11_2, $dictOrd1_0) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($dictOrd11_2, $dictOrd1_0, $dictOrd_4) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($dictOrd11_2, $dictOrd1_0, $dictOrd_4, $v_5) {
  $__num = \func_num_args();
  $v2_7_3 = (((($dictOrd1_0)->{'compare1'})($dictOrd_4))(($v_5)->{'value0'}))(($v1_6)->{'value0'});
  $__t4 = null;;
  if ($v2_7_3 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t4 = (((($dictOrd11_2)->{'compare1'})($dictOrd_4))(($v_5)->{'value1'}))(($v1_6)->{'value1'});
goto end_branch_4;;
};
  $__t4 = $v2_7_3;
  end_branch_4:;
  $__res = $__t4;
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
}, "Eq10" => function($_dollar___unused_4) use ($eq1Product2_3_1) {
  $__num = \func_num_args();
  $__res = $eq1Product2_3_1;
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
$GLOBALS['Data_Functor_Product_ord1Product'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_ord1majProduct';

// Data_Functor_Product_ordProduct
function majData_majFunctor_majProduct_ordmajProduct($dictOrd1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_ordmajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd1_0)->{'Eq10'})(null);
  $__res = function($dictOrd11_2) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd11_2)->{'Eq10'})(null);
  $__res = function($dictOrd_4) use ($__local_var_1_0, $__local_var_3_1, $dictOrd11_2, $dictOrd1_0) {
  $__num = \func_num_args();
  $__local_var_5_2 = (($dictOrd_4)->{'Eq0'})(null);
  $eqProduct3_5_2 = (object)["eq" => function($v_6) use ($__local_var_1_0, $__local_var_3_1, $__local_var_5_2) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_1_0, $__local_var_3_1, $__local_var_5_2, $v_6) {
  $__num = \func_num_args();
  $__res = ((((($__local_var_1_0)->{'eq1'})($__local_var_5_2))(($v_6)->{'value0'}))(($v1_7)->{'value0'}) && (((($__local_var_3_1)->{'eq1'})($__local_var_5_2))(($v_6)->{'value1'}))(($v1_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($v_6) use ($dictOrd11_2, $dictOrd1_0, $dictOrd_4) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($dictOrd11_2, $dictOrd1_0, $dictOrd_4, $v_6) {
  $__num = \func_num_args();
  $v2_8_4 = (((($dictOrd1_0)->{'compare1'})($dictOrd_4))(($v_6)->{'value0'}))(($v1_7)->{'value0'});
  $__t5 = null;;
  if ($v2_8_4 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t5 = (((($dictOrd11_2)->{'compare1'})($dictOrd_4))(($v_6)->{'value1'}))(($v1_7)->{'value1'});
goto end_branch_5;;
};
  $__t5 = $v2_8_4;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_6) use ($eqProduct3_5_2) {
  $__num = \func_num_args();
  $__res = $eqProduct3_5_2;
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
$GLOBALS['Data_Functor_Product_ordProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_ordmajProduct';

// Data_Functor_Product_bihoistProduct
function majData_majFunctor_majProduct_bihoistmajProduct($natF_0, $natG_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_bihoistmajProduct';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($natF_0)(($v_2)->{'value0'}), ($natG_1)(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Functor_Product_bihoistProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_bihoistmajProduct';

// Data_Functor_Product_applyProduct
function majData_majFunctor_majProduct_applymajProduct($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_applymajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($dictApply1_2) use ($__local_var_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictApply1_2)->{'Functor0'})(null);
  $functorProduct2_3_1 = (object)["map" => function($f_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_1_0, $__local_var_3_1, $f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_1_0)->{'map'})($f_4))(($v_5)->{'value0'}), ((($__local_var_3_1)->{'map'})($f_4))(($v_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($v_4) use ($dictApply1_2, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictApply1_2, $dictApply_0, $v_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictApply_0)->{'apply'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictApply1_2)->{'apply'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorProduct2_3_1) {
  $__num = \func_num_args();
  $__res = $functorProduct2_3_1;
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
$GLOBALS['Data_Functor_Product_applyProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_applymajProduct';

// Data_Functor_Product_bindProduct
function majData_majFunctor_majProduct_bindmajProduct($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_bindmajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBind_0)->{'Apply0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__res = function($dictBind1_3) use ($__local_var_1_0, $__local_var_2_1, $dictBind_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictBind1_3)->{'Apply0'})(null);
  $__local_var_5_3 = (($__local_var_4_2)->{'Functor0'})(null);
  $functorProduct2_5_3 = (object)["map" => function($f_6) use ($__local_var_2_1, $__local_var_5_3) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_2_1, $__local_var_5_3, $f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_2_1)->{'map'})($f_6))(($v_7)->{'value0'}), ((($__local_var_5_3)->{'map'})($f_6))(($v_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyProduct2_4_2 = (object)["apply" => function($v_6) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_1_0, $__local_var_4_2, $v_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_1_0)->{'apply'})(($v_6)->{'value0'}))(($v1_7)->{'value0'}), ((($__local_var_4_2)->{'apply'})(($v_6)->{'value1'}))(($v1_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorProduct2_5_3) {
  $__num = \func_num_args();
  $__res = $functorProduct2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_5) use ($dictBind1_3, $dictBind_0) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($dictBind1_3, $dictBind_0, $v_5) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictBind_0)->{'bind'})(($v_5)->{'value0'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_fst']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($f_6))), ((($dictBind1_3)->{'bind'})(($v_5)->{'value1'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($f_6))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($applyProduct2_4_2) {
  $__num = \func_num_args();
  $__res = $applyProduct2_4_2;
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
$GLOBALS['Data_Functor_Product_bindProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_bindmajProduct';

// Data_Functor_Product_applicativeProduct
function majData_majFunctor_majProduct_applicativemajProduct($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_applicativemajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__res = function($dictApplicative1_3) use ($__local_var_1_0, $__local_var_2_1, $dictApplicative_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictApplicative1_3)->{'Apply0'})(null);
  $__local_var_5_3 = (($__local_var_4_2)->{'Functor0'})(null);
  $functorProduct2_5_3 = (object)["map" => function($f_6) use ($__local_var_2_1, $__local_var_5_3) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_2_1, $__local_var_5_3, $f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_2_1)->{'map'})($f_6))(($v_7)->{'value0'}), ((($__local_var_5_3)->{'map'})($f_6))(($v_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyProduct2_4_2 = (object)["apply" => function($v_6) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_1_0, $__local_var_4_2, $v_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_1_0)->{'apply'})(($v_6)->{'value0'}))(($v1_7)->{'value0'}), ((($__local_var_4_2)->{'apply'})(($v_6)->{'value1'}))(($v1_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorProduct2_5_3) {
  $__num = \func_num_args();
  $__res = $functorProduct2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => function($a_5) use ($dictApplicative1_3, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($dictApplicative_0)->{'pure'})($a_5), (($dictApplicative1_3)->{'pure'})($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($applyProduct2_4_2) {
  $__num = \func_num_args();
  $__res = $applyProduct2_4_2;
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
$GLOBALS['Data_Functor_Product_applicativeProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_applicativemajProduct';

// Data_Functor_Product_monadProduct
function majData_majFunctor_majProduct_monadmajProduct($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct_monadmajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Apply0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $applicativeProduct1_3_2 = function($dictApplicative1_4) use ($__local_var_1_0, $__local_var_2_1, $__local_var_3_2) {
  $__num = \func_num_args();
  $__local_var_5_3 = (($dictApplicative1_4)->{'Apply0'})(null);
  $__local_var_6_4 = (($__local_var_5_3)->{'Functor0'})(null);
  $functorProduct2_6_4 = (object)["map" => function($f_7) use ($__local_var_3_2, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_3_2, $__local_var_6_4, $f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_3_2)->{'map'})($f_7))(($v_8)->{'value0'}), ((($__local_var_6_4)->{'map'})($f_7))(($v_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyProduct2_5_3 = (object)["apply" => function($v_7) use ($__local_var_2_1, $__local_var_5_3) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_2_1, $__local_var_5_3, $v_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_2_1)->{'apply'})(($v_7)->{'value0'}))(($v1_8)->{'value0'}), ((($__local_var_5_3)->{'apply'})(($v_7)->{'value1'}))(($v1_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorProduct2_6_4) {
  $__num = \func_num_args();
  $__res = $functorProduct2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => function($a_6) use ($__local_var_1_0, $dictApplicative1_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($__local_var_1_0)->{'pure'})($a_6), (($dictApplicative1_4)->{'pure'})($a_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($applyProduct2_5_3) {
  $__num = \func_num_args();
  $__res = $applyProduct2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_4_8 = (($dictMonad_0)->{'Bind1'})(null);
  $__local_var_5_9 = (($__local_var_4_8)->{'Apply0'})(null);
  $__local_var_6_10 = (($__local_var_5_9)->{'Functor0'})(null);
  $bindProduct1_6_10 = function($dictBind1_7) use ($__local_var_4_8, $__local_var_5_9, $__local_var_6_10) {
  $__num = \func_num_args();
  $__local_var_8_11 = (($dictBind1_7)->{'Apply0'})(null);
  $__local_var_9_12 = (($__local_var_8_11)->{'Functor0'})(null);
  $functorProduct2_9_12 = (object)["map" => function($f_10) use ($__local_var_6_10, $__local_var_9_12) {
  $__num = \func_num_args();
  $__res = function($v_11) use ($__local_var_6_10, $__local_var_9_12, $f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_6_10)->{'map'})($f_10))(($v_11)->{'value0'}), ((($__local_var_9_12)->{'map'})($f_10))(($v_11)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyProduct2_8_11 = (object)["apply" => function($v_10) use ($__local_var_5_9, $__local_var_8_11) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($__local_var_5_9, $__local_var_8_11, $v_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_5_9)->{'apply'})(($v_10)->{'value0'}))(($v1_11)->{'value0'}), ((($__local_var_8_11)->{'apply'})(($v_10)->{'value1'}))(($v1_11)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_10) use ($functorProduct2_9_12) {
  $__num = \func_num_args();
  $__res = $functorProduct2_9_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_9) use ($__local_var_4_8, $dictBind1_7) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($__local_var_4_8, $dictBind1_7, $v_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_4_8)->{'bind'})(($v_9)->{'value0'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_fst']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($f_10))), ((($dictBind1_7)->{'bind'})(($v_9)->{'value1'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($f_10))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($applyProduct2_8_11) {
  $__num = \func_num_args();
  $__res = $applyProduct2_8_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonad1_7) use ($applicativeProduct1_3_2, $bindProduct1_6_10) {
  $__num = \func_num_args();
  $applicativeProduct2_8_16 = ($applicativeProduct1_3_2)((($dictMonad1_7)->{'Applicative0'})(null));
  $bindProduct2_9_17 = ($bindProduct1_6_10)((($dictMonad1_7)->{'Bind1'})(null));
  $__res = (object)["Applicative0" => function($_dollar___unused_10) use ($applicativeProduct2_8_16) {
  $__num = \func_num_args();
  $__res = $applicativeProduct2_8_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_10) use ($bindProduct2_9_17) {
  $__num = \func_num_args();
  $__res = $bindProduct2_9_17;
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
$GLOBALS['Data_Functor_Product_monadProduct'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct_monadmajProduct';

