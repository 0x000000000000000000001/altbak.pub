<?php

namespace Data\Functor\Product2;

// ALL IMPORTS: Control.Biapplicative, Control.Biapply, Data.Bifunctor, Data.Eq, Data.Functor, Data.Functor.Product2, Data.HeytingAlgebra, Data.Ord, Data.Ordering, Data.Profunctor, Data.Semigroup, Data.Show, Prelude, Prim
// TO REQUIRE: Control.Biapplicative, Control.Biapply, Data.Bifunctor, Data.Eq, Data.Functor, Data.Functor.Product2, Data.HeytingAlgebra, Data.Ord, Data.Ordering, Data.Profunctor, Data.Semigroup, Data.Show, Prelude
require_once __DIR__ . '/../Control.Biapplicative/index.php';
require_once __DIR__ . '/../Control.Biapply/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Product2/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Profunctor/index.php';
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


final class Data_Functor_Product2_Product2 { public $tag = 'Product2'; public function __construct(public  $value0, public  $value1) {} }

// Data_Functor_Product2_Product2
$GLOBALS['Data_Functor_Product2_Product2'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Functor_Product2_showProduct2
function majData_majFunctor_majProduct2_showmajProduct2($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct2_showmajProduct2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Product2 "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow1_1)->{'show'})(($v_2)->{'value1'})))(")"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Functor_Product2_showProduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct2_showmajProduct2';

// Data_Functor_Product2_profunctorProduct2
function majData_majFunctor_majProduct2_profunctormajProduct2($dictProfunctor_0, $dictProfunctor1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct2_profunctormajProduct2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["dimap" => (function() use ($dictProfunctor1_1, $dictProfunctor_0) {
  $__fn = function($f_2, $g_3 = null, $v_4 = null) use ($dictProfunctor1_1, $dictProfunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2((((($dictProfunctor_0)->{'dimap'})($f_2))($g_3))(($v_4)->{'value0'}), (((($dictProfunctor1_1)->{'dimap'})($f_2))($g_3))(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Functor_Product2_profunctorProduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct2_profunctormajProduct2';

// Data_Functor_Product2_functorProduct2
function majData_majFunctor_majProduct2_functormajProduct2($dictFunctor_0, $dictFunctor1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct2_functormajProduct2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["map" => (function() use ($dictFunctor1_1, $dictFunctor_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictFunctor1_1, $dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2(((($dictFunctor_0)->{'map'})($f_2))(($v_3)->{'value0'}), ((($dictFunctor1_1)->{'map'})($f_2))(($v_3)->{'value1'}));
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
$GLOBALS['Data_Functor_Product2_functorProduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct2_functormajProduct2';

// Data_Functor_Product2_eqProduct2
function majData_majFunctor_majProduct2_eqmajProduct2($dictEq_0, $dictEq1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct2_eqmajProduct2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($x_2, $y_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($dictEq_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'})))(((($dictEq1_1)->{'eq'})(($x_2)->{'value1'}))(($y_3)->{'value1'}));
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
$GLOBALS['Data_Functor_Product2_eqProduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct2_eqmajProduct2';

// Data_Functor_Product2_ordProduct2
function majData_majFunctor_majProduct2_ordmajProduct2($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct2_ordmajProduct2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eqProduct21_1_0 = ($GLOBALS['Data_Functor_Product2_eqProduct2'])((($dictOrd_0)->{'Eq0'})(null));
  $__res = function($dictOrd1_2) use ($dictOrd_0, $eqProduct21_1_0) {
  $__num = \func_num_args();
  $eqProduct22_3_1 = ($eqProduct21_1_0)((($dictOrd1_2)->{'Eq0'})(null));
  $__res = (object)["compare" => (function() use ($dictOrd1_2, $dictOrd_0) {
  $__fn = function($x_4, $y_5 = null) use ($dictOrd1_2, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_6_2 = ((($dictOrd_0)->{'compare'})(($x_4)->{'value0'}))(($y_5)->{'value0'});
  $__t3 = null;;
  if ($v_6_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_3;;
};
  if ($v_6_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_3;;
};
  $__t3 = ((($dictOrd1_2)->{'compare'})(($x_4)->{'value1'}))(($y_5)->{'value1'});
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_4) use ($eqProduct22_3_1) {
  $__num = \func_num_args();
  $__res = $eqProduct22_3_1;
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
$GLOBALS['Data_Functor_Product2_ordProduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct2_ordmajProduct2';

// Data_Functor_Product2_bifunctorProduct2
function majData_majFunctor_majProduct2_bifunctormajProduct2($dictBifunctor_0, $dictBifunctor1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct2_bifunctormajProduct2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["bimap" => (function() use ($dictBifunctor1_1, $dictBifunctor_0) {
  $__fn = function($f_2, $g_3 = null, $v_4 = null) use ($dictBifunctor1_1, $dictBifunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2((((($dictBifunctor_0)->{'bimap'})($f_2))($g_3))(($v_4)->{'value0'}), (((($dictBifunctor1_1)->{'bimap'})($f_2))($g_3))(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Functor_Product2_bifunctorProduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct2_bifunctormajProduct2';

// Data_Functor_Product2_biapplyProduct2
function majData_majFunctor_majProduct2_biapplymajProduct2($dictBiapply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct2_biapplymajProduct2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBiapply_0)->{'Bifunctor0'})(null);
  $__res = function($dictBiapply1_2) use ($__local_var_1_0, $dictBiapply_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictBiapply1_2)->{'Bifunctor0'})(null);
  $bifunctorProduct22_4_2 = (object)["bimap" => (function() use ($__local_var_1_0, $__local_var_3_1) {
  $__fn = function($f_4, $g_5 = null, $v_6 = null) use ($__local_var_1_0, $__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2((((($__local_var_1_0)->{'bimap'})($f_4))($g_5))(($v_6)->{'value0'}), (((($__local_var_3_1)->{'bimap'})($f_4))($g_5))(($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["biapply" => (function() use ($dictBiapply1_2, $dictBiapply_0) {
  $__fn = function($v_5, $v1_6 = null) use ($dictBiapply1_2, $dictBiapply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2(((($dictBiapply_0)->{'biapply'})(($v_5)->{'value0'}))(($v1_6)->{'value0'}), ((($dictBiapply1_2)->{'biapply'})(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Bifunctor0" => function($_dollar__unused_5) use ($bifunctorProduct22_4_2) {
  $__num = \func_num_args();
  $__res = $bifunctorProduct22_4_2;
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
$GLOBALS['Data_Functor_Product2_biapplyProduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct2_biapplymajProduct2';

// Data_Functor_Product2_biapplicativeProduct2
function majData_majFunctor_majProduct2_biapplicativemajProduct2($dictBiapplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majProduct2_biapplicativemajProduct2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBiapplicative_0)->{'Biapply0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Bifunctor0'})(null);
  $__res = function($dictBiapplicative1_3) use ($__local_var_1_0, $__local_var_2_1, $dictBiapplicative_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictBiapplicative1_3)->{'Biapply0'})(null);
  $__local_var_5_3 = (($__local_var_4_2)->{'Bifunctor0'})(null);
  $bifunctorProduct22_6_4 = (object)["bimap" => (function() use ($__local_var_2_1, $__local_var_5_3) {
  $__fn = function($f_6, $g_7 = null, $v_8 = null) use ($__local_var_2_1, $__local_var_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2((((($__local_var_2_1)->{'bimap'})($f_6))($g_7))(($v_8)->{'value0'}), (((($__local_var_5_3)->{'bimap'})($f_6))($g_7))(($v_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $biapplyProduct22_6_4 = (object)["biapply" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($v_7, $v1_8 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2(((($__local_var_1_0)->{'biapply'})(($v_7)->{'value0'}))(($v1_8)->{'value0'}), ((($__local_var_4_2)->{'biapply'})(($v_7)->{'value1'}))(($v1_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Bifunctor0" => function($_dollar__unused_7) use ($bifunctorProduct22_6_4) {
  $__num = \func_num_args();
  $__res = $bifunctorProduct22_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bipure" => (function() use ($dictBiapplicative1_3, $dictBiapplicative_0) {
  $__fn = function($a_7, $b_8 = null) use ($dictBiapplicative1_3, $dictBiapplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2(((($dictBiapplicative_0)->{'bipure'})($a_7))($b_8), ((($dictBiapplicative1_3)->{'bipure'})($a_7))($b_8));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Biapply0" => function($_dollar__unused_7) use ($biapplyProduct22_6_4) {
  $__num = \func_num_args();
  $__res = $biapplyProduct22_6_4;
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
$GLOBALS['Data_Functor_Product2_biapplicativeProduct2'] = __NAMESPACE__ . '\\majData_majFunctor_majProduct2_biapplicativemajProduct2';

