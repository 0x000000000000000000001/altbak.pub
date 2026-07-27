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
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(product "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow1_1)->{'show'})(($v_2)->{'value1'})))(")"))));
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
$GLOBALS['Data_Functor_Product_newtypeProduct'] = (object)["Coercible0" => function($_dollar__unused_0) {
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
  $__res = (object)["map" => (function() use ($dictFunctor1_1, $dictFunctor_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictFunctor1_1, $dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Bifunctor_bifunctorTuple'])->{'bimap'})((($dictFunctor_0)->{'map'})($f_2)))((($dictFunctor1_1)->{'map'})($f_2)))($v_3);
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
  $eq12_3_0 = (($dictEq1_0)->{'eq1'})($dictEq_2);
  $eq13_4_1 = (($dictEq11_1)->{'eq1'})($dictEq_2);
  $__res = (function() use ($eq12_3_0, $eq13_4_1) {
  $__fn = function($v_5, $v1_6 = null) use ($eq12_3_0, $eq13_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($eq12_3_0)(($v_5)->{'value0'}))(($v1_6)->{'value0'})))((($eq13_4_1)(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
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
  $__res = (object)["eq" => ((\Data\Functor\Product\majData_majFunctor_majProduct_eq1majProduct($dictEq1_0, $dictEq11_1))->{'eq1'})($dictEq_2)];
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
  $eq1Product1_1_0 = ($GLOBALS['Data_Functor_Product_eq1Product'])((($dictOrd1_0)->{'Eq10'})(null));
  $__res = function($dictOrd11_2) use ($dictOrd1_0, $eq1Product1_1_0) {
  $__num = \func_num_args();
  $eq1Product2_3_1 = ($eq1Product1_1_0)((($dictOrd11_2)->{'Eq10'})(null));
  $__res = (object)["compare1" => function($dictOrd_4) use ($dictOrd11_2, $dictOrd1_0) {
  $__num = \func_num_args();
  $compare12_5_2 = (($dictOrd1_0)->{'compare1'})($dictOrd_4);
  $compare13_6_3 = (($dictOrd11_2)->{'compare1'})($dictOrd_4);
  $__res = (function() use ($compare12_5_2, $compare13_6_3) {
  $__fn = function($v_7, $v1_8 = null) use ($compare12_5_2, $compare13_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v2_9_4 = (($compare12_5_2)(($v_7)->{'value0'}))(($v1_8)->{'value0'});
  $__t5 = null;;
  if ($v2_9_4 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t5 = (($compare13_6_3)(($v_7)->{'value1'}))(($v1_8)->{'value1'});
goto end_branch_5;;
};
  $__t5 = $v2_9_4;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_4) use ($eq1Product2_3_1) {
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
  $ord1Product1_1_0 = ($GLOBALS['Data_Functor_Product_ord1Product'])($dictOrd1_0);
  $eqProduct1_2_1 = ($GLOBALS['Data_Functor_Product_eqProduct'])((($dictOrd1_0)->{'Eq10'})(null));
  $__res = function($dictOrd11_3) use ($eqProduct1_2_1, $ord1Product1_1_0) {
  $__num = \func_num_args();
  $eqProduct2_4_2 = ($eqProduct1_2_1)((($dictOrd11_3)->{'Eq10'})(null));
  $__res = function($dictOrd_5) use ($dictOrd11_3, $eqProduct2_4_2, $ord1Product1_1_0) {
  $__num = \func_num_args();
  $eqProduct3_6_3 = ($eqProduct2_4_2)((($dictOrd_5)->{'Eq0'})(null));
  $__res = (object)["compare" => ((($ord1Product1_1_0)($dictOrd11_3))->{'compare1'})($dictOrd_5), "Eq0" => function($_dollar__unused_7) use ($eqProduct3_6_3) {
  $__num = \func_num_args();
  $__res = $eqProduct3_6_3;
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
  $__res = (((($GLOBALS['Data_Bifunctor_bifunctorTuple'])->{'bimap'})($natF_0))($natG_1))($v_2);
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
  $functorProduct1_1_0 = ($GLOBALS['Data_Functor_Product_functorProduct'])((($dictApply_0)->{'Functor0'})(null));
  $__res = function($dictApply1_2) use ($dictApply_0, $functorProduct1_1_0) {
  $__num = \func_num_args();
  $functorProduct2_3_1 = ($functorProduct1_1_0)((($dictApply1_2)->{'Functor0'})(null));
  $__res = (object)["apply" => (function() use ($dictApply1_2, $dictApply_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictApply1_2, $dictApply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictApply_0)->{'apply'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictApply1_2)->{'apply'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4) use ($functorProduct2_3_1) {
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
  $applyProduct1_1_0 = ($GLOBALS['Data_Functor_Product_applyProduct'])((($dictBind_0)->{'Apply0'})(null));
  $__res = function($dictBind1_2) use ($applyProduct1_1_0, $dictBind_0) {
  $__num = \func_num_args();
  $applyProduct2_3_1 = ($applyProduct1_1_0)((($dictBind1_2)->{'Apply0'})(null));
  $__res = (object)["bind" => (function() use ($dictBind1_2, $dictBind_0) {
  $__fn = function($v_4, $f_5 = null) use ($dictBind1_2, $dictBind_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictBind_0)->{'bind'})(($v_4)->{'value0'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_fst']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($f_5))), ((($dictBind1_2)->{'bind'})(($v_4)->{'value1'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($f_5))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_4) use ($applyProduct2_3_1) {
  $__num = \func_num_args();
  $__res = $applyProduct2_3_1;
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
  $applyProduct1_1_0 = ($GLOBALS['Data_Functor_Product_applyProduct'])((($dictApplicative_0)->{'Apply0'})(null));
  $__res = function($dictApplicative1_2) use ($applyProduct1_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $applyProduct2_3_1 = ($applyProduct1_1_0)((($dictApplicative1_2)->{'Apply0'})(null));
  $__res = (object)["pure" => function($a_4) use ($dictApplicative1_2, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($dictApplicative_0)->{'pure'})($a_4), (($dictApplicative1_2)->{'pure'})($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_4) use ($applyProduct2_3_1) {
  $__num = \func_num_args();
  $__res = $applyProduct2_3_1;
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
  $applicativeProduct1_1_0 = ($GLOBALS['Data_Functor_Product_applicativeProduct'])((($dictMonad_0)->{'Applicative0'})(null));
  $bindProduct1_2_1 = ($GLOBALS['Data_Functor_Product_bindProduct'])((($dictMonad_0)->{'Bind1'})(null));
  $__res = function($dictMonad1_3) use ($applicativeProduct1_1_0, $bindProduct1_2_1) {
  $__num = \func_num_args();
  $applicativeProduct2_4_2 = ($applicativeProduct1_1_0)((($dictMonad1_3)->{'Applicative0'})(null));
  $bindProduct2_5_3 = ($bindProduct1_2_1)((($dictMonad1_3)->{'Bind1'})(null));
  $__res = (object)["Applicative0" => function($_dollar__unused_6) use ($applicativeProduct2_4_2) {
  $__num = \func_num_args();
  $__res = $applicativeProduct2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_6) use ($bindProduct2_5_3) {
  $__num = \func_num_args();
  $__res = $bindProduct2_5_3;
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

