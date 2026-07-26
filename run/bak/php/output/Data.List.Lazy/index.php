<?php

namespace Data\List\Lazy;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Lazy, Data.List.Internal, Data.List.Lazy, Data.List.Lazy.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Lazy, Data.List.Internal, Data.List.Lazy, Data.List.Lazy.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Internal/index.php';
require_once __DIR__ . '/../Data.List.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Lazy.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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


// Data_List_Lazy_tailRecM2
$GLOBALS['Data_List_Lazy_tailRecM2'] = (function() {
  $__fn = function($f_0 = null, $a_1 = null, $b_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_Rec_Class_monadRecMaybe'])['tailRecM'])(function($o_3 = null) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($f_0)(($o_3)['a']))(($o_3)['b']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(["a" => $a_1, "b" => $b_2]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_any
$GLOBALS['Data_List_Lazy_any'] = (function() use (&$__fn) {
$semigroupDisj1_0_0 = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0 || $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return (($GLOBALS['Data_List_Lazy_Types_foldableList'])['foldMap'])(["mempty" => false, "Semigroup0" => function($dollar__unused_1 = null) use ($semigroupDisj1_0_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
})();

// Data_List_Lazy_Pattern
$GLOBALS['Data_List_Lazy_Pattern'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_zipWith
$GLOBALS['Data_List_Lazy_zipWith'] = (function() {
  $__fn = function($f_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_zipWith_f_0 = $f_0;
  $__tco_var_Data_List_Lazy_zipWith_xs_1 = $xs_1;
  $__tco_var_Data_List_Lazy_zipWith_ys_2 = $ys_2;
  tco_loop_Data_List_Lazy_zipWith:;
  $f_0 = $__tco_var_Data_List_Lazy_zipWith_f_0;
  $xs_1 = $__tco_var_Data_List_Lazy_zipWith_xs_1;
  $ys_2 = $__tco_var_Data_List_Lazy_zipWith_ys_2;
  $__local_var_3_0 = ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($f_0, $xs_1) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Lazy_force'])($xs_1);
  $__res = function($v1_5 = null) use ($__local_var_4_0, $f_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if (((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Cons")) && (is_object($v1_5) && (($v1_5)->{'tag'} === "Cons")))) {
$__t1 = new Phpurs_Data2("Cons", (($f_0)(($__local_var_4_0)->{'value0'}))(($v1_5)->{'value0'}), ((($GLOBALS['Data_List_Lazy_zipWith'])($f_0))(($__local_var_4_0)->{'value1'}))(($v1_5)->{'value1'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_4 = null) use ($__local_var_3_0, $ys_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_force'])($__local_var_3_0))(($GLOBALS['Data_Lazy_force'])($ys_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_zipWithA
$GLOBALS['Data_List_Lazy_zipWithA'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $sequence1_1_0 = ((($GLOBALS['Data_List_Lazy_Types_traversableList'])['traverse'])($dictApplicative_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2 = null, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($sequence1_1_0)(((($GLOBALS['Data_List_Lazy_zipWith'])($f_2))($xs_3))($ys_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_zip
$GLOBALS['Data_List_Lazy_zip'] = ($GLOBALS['Data_List_Lazy_zipWith'])($GLOBALS['Data_Tuple_Tuple']);

// Data_List_Lazy_updateAt
$GLOBALS['Data_List_Lazy_updateAt'] = (function() {
  $__fn = function($n_0 = null, $x_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_updateAt_n_0 = $n_0;
  $__tco_var_Data_List_Lazy_updateAt_x_1 = $x_1;
  $__tco_var_Data_List_Lazy_updateAt_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_updateAt:;
  $n_0 = $__tco_var_Data_List_Lazy_updateAt_n_0;
  $x_1 = $__tco_var_Data_List_Lazy_updateAt_x_1;
  $xs_2 = $__tco_var_Data_List_Lazy_updateAt_xs_2;
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($n_0, $x_1, $xs_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Lazy_force'])($xs_2);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Cons"))) {
$__t1 = match ($n_0) { 0 => new Phpurs_Data2("Cons", $x_1, ($__local_var_4_0)->{'value1'}), default => new Phpurs_Data2("Cons", ($__local_var_4_0)->{'value0'}, ((($GLOBALS['Data_List_Lazy_updateAt'])(($n_0 - 1)))($x_1))(($__local_var_4_0)->{'value1'})) };
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_unzip
$GLOBALS['Data_List_Lazy_unzip'] = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])['foldr'])(function($v_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($v_0)->{'value0'};
  $__local_var_2_1 = ($v_0)->{'value1'};
  $__res = function($v1_3 = null) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_4_2 = ($v1_3)->{'value0'};
  $__local_var_5_3 = ($v1_3)->{'value1'};
  $__res = new Phpurs_Data2("Tuple", ($GLOBALS['Data_Lazy_defer'])(function($v_6 = null) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_1_0, $__local_var_4_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ($GLOBALS['Data_Lazy_defer'])(function($v_6 = null) use ($__local_var_2_1, $__local_var_5_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_2_1, $__local_var_5_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", $GLOBALS['Data_List_Lazy_Types_nil'], $GLOBALS['Data_List_Lazy_Types_nil']));

// Data_List_Lazy_uncons
$GLOBALS['Data_List_Lazy_uncons'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $v_1_0 = ($GLOBALS['Data_List_Lazy_Types_step'])($xs_0);
  $__t1 = null;;
  if ((is_object($v_1_0) && (($v_1_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  if ((is_object($v_1_0) && (($v_1_0)->{'tag'} === "Cons"))) {
$__t1 = new Phpurs_Data1("Just", ["head" => ($v_1_0)->{'value0'}, "tail" => ($v_1_0)->{'value1'}]);
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_toUnfoldable
$GLOBALS['Data_List_Lazy_toUnfoldable'] = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = (($dictUnfoldable_0)['unfoldr'])(function($xs_1 = null) {
  $__num = \func_num_args();
  $__local_var_2_0 = ($GLOBALS['Data_List_Lazy_uncons'])($xs_1);
  $__t1 = null;;
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", (($__local_var_2_0)->{'value0'})['head'], (($__local_var_2_0)->{'value0'})['tail']));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_takeWhile
$GLOBALS['Data_List_Lazy_takeWhile'] = function($p_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_takeWhile_p_0 = $p_0;
  tco_loop_Data_List_Lazy_takeWhile:;
  $p_0 = $__tco_var_Data_List_Lazy_takeWhile_p_0;
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])['map'])(function($v_1 = null) use ($p_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (((is_object($v_1) && (($v_1)->{'tag'} === "Cons")) && ($p_0)(($v_1)->{'value0'}))) {
$__t0 = new Phpurs_Data2("Cons", ($v_1)->{'value0'}, (($GLOBALS['Data_List_Lazy_takeWhile'])($p_0))(($v_1)->{'value1'}));
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nil");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_take
$GLOBALS['Data_List_Lazy_take'] = function($n_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_take_n_0 = $n_0;
  tco_loop_Data_List_Lazy_take:;
  $n_0 = $__tco_var_Data_List_Lazy_take_n_0;
  $__t1 = null;;
  if (($n_0 <= 0)) {
$__t1 = function($v_1 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_nil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_1;;
};
  $__t1 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])['map'])(function($v1_1 = null) use ($n_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data2("Cons", ($v1_1)->{'value0'}, (($GLOBALS['Data_List_Lazy_take'])(($n_0 - 1)))(($v1_1)->{'value1'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_tail
$GLOBALS['Data_List_Lazy_tail'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($GLOBALS['Data_List_Lazy_uncons'])($xs_0);
  $__t1 = null;;
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($__local_var_1_0)->{'value0'})['tail']);
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_stripPrefix
$GLOBALS['Data_List_Lazy_stripPrefix'] = (function() {
  $__fn = function($dictEq_0 = null, $v_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_tailRecM2'])((function() use ($dictEq_0) {
  $__fn = function($prefix_3 = null, $input_4 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_5_0 = ($GLOBALS['Data_List_Lazy_Types_step'])($prefix_3);
  $__t1 = null;;
  if ((is_object($v1_5_0) && (($v1_5_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data1("Done", $input_4));
goto end_branch_1;;
};
  if ((is_object($v1_5_0) && (($v1_5_0)->{'tag'} === "Cons"))) {
$v2_6_2 = ($GLOBALS['Data_List_Lazy_Types_step'])($input_4);
$__t3 = null;;
if (((is_object($v2_6_2) && (($v2_6_2)->{'tag'} === "Cons")) && ((($dictEq_0)['eq'])(($v1_5_0)->{'value0'}))(($v2_6_2)->{'value0'}))) {
$__t3 = new Phpurs_Data1("Just", new Phpurs_Data1("Loop", ["a" => ($v1_5_0)->{'value1'}, "b" => ($v2_6_2)->{'value1'}]));
goto end_branch_3;;
};
$__t3 = new Phpurs_Data0("Nothing");
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($v_1))($s_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_span
$GLOBALS['Data_List_Lazy_span'] = (function() {
  $__fn = function($p_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_Lazy_span_p_0 = $p_0;
  $__tco_var_Data_List_Lazy_span_xs_1 = $xs_1;
  tco_loop_Data_List_Lazy_span:;
  $p_0 = $__tco_var_Data_List_Lazy_span_p_0;
  $xs_1 = $__tco_var_Data_List_Lazy_span_xs_1;
  $v_2_0 = ($GLOBALS['Data_List_Lazy_uncons'])($xs_1);
  $__t1 = null;;
  if (((is_object($v_2_0) && (($v_2_0)->{'tag'} === "Just")) && ($p_0)((($v_2_0)->{'value0'})['head']))) {
$__local_var_3_2 = (($v_2_0)->{'value0'})['head'];
$v1_4_3 = (($GLOBALS['Data_List_Lazy_span'])($p_0))((($v_2_0)->{'value0'})['tail']);
$__local_var_5_4 = ($v1_4_3)['init'];
$__t1 = ["init" => ($GLOBALS['Data_Lazy_defer'])(function($v_6 = null) use ($__local_var_3_2, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_3_2, $__local_var_5_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "rest" => ($v1_4_3)['rest']];
goto end_branch_1;;
};
  $__t1 = ["init" => $GLOBALS['Data_List_Lazy_Types_nil'], "rest" => $xs_1];
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_snoc
$GLOBALS['Data_List_Lazy_snoc'] = (function() {
  $__fn = function($xs_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])['foldr'])($GLOBALS['Data_List_Lazy_Types_cons']))(($GLOBALS['Data_Lazy_defer'])(function($v_2 = null) use ($x_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $x_1, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($xs_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_singleton
$GLOBALS['Data_List_Lazy_singleton'] = function($a_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_1 = null) use ($a_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_0, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_showPattern
$GLOBALS['Data_List_Lazy_showPattern'] = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = ["show" => function($v_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = (("(Pattern " . ((($GLOBALS['Data_List_Lazy_Types_showList'])($dictShow_0))['show'])($v_1)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_scanlLazy
$GLOBALS['Data_List_Lazy_scanlLazy'] = (function() {
  $__fn = function($f_0 = null, $acc_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_scanlLazy_f_0 = $f_0;
  $__tco_var_Data_List_Lazy_scanlLazy_acc_1 = $acc_1;
  $__tco_var_Data_List_Lazy_scanlLazy_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_scanlLazy:;
  $f_0 = $__tco_var_Data_List_Lazy_scanlLazy_f_0;
  $acc_1 = $__tco_var_Data_List_Lazy_scanlLazy_acc_1;
  $xs_2 = $__tco_var_Data_List_Lazy_scanlLazy_xs_2;
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($acc_1, $f_0, $xs_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Lazy_force'])($xs_2);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Cons"))) {
$acc__prime___5_2 = (($f_0)($acc_1))(($__local_var_4_0)->{'value0'});
$__t1 = new Phpurs_Data2("Cons", $acc__prime___5_2, ((($GLOBALS['Data_List_Lazy_scanlLazy'])($f_0))($acc__prime___5_2))(($__local_var_4_0)->{'value1'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_reverse
$GLOBALS['Data_List_Lazy_reverse'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_defer'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))(function($v_1 = null) use ($xs_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])['foldl'])((function() {
  $__fn = function($b_2 = null, $a_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_4 = null) use ($a_3, $b_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_3, $b_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_List_Lazy_Types_nil']))($xs_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_replicateM
$GLOBALS['Data_List_Lazy_replicateM'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_replicateM_dictMonad_0 = $dictMonad_0;
  tco_loop_Data_List_Lazy_replicateM:;
  $dictMonad_0 = $__tco_var_Data_List_Lazy_replicateM_dictMonad_0;
  $__local_var_1_0 = (($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']);
  $__local_var_2_1 = (($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']);
  $__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($n_3 = null, $m_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if (($n_3 < 1)) {
$__t2 = (($__local_var_1_0)['pure'])($GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_2;;
};
  $__t2 = ((($__local_var_2_1)['bind'])($m_4))(function($a_5 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, $m_4, $n_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)['bind'])(((($GLOBALS['Data_List_Lazy_replicateM'])($dictMonad_0))(($n_3 - 1)))($m_4)))(function($as_6 = null) use ($__local_var_1_0, $a_5) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)['pure'])(($GLOBALS['Data_Lazy_defer'])(function($v_7 = null) use ($a_5, $as_6) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_5, $as_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_repeat
$GLOBALS['Data_List_Lazy_repeat'] = function($x_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = ($GLOBALS['Data_Lazy_defer'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))(function($v_2 = null) use (&$go_1_0, $x_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use (&$go_1_0, $x_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $x_0, $go_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_replicate
$GLOBALS['Data_List_Lazy_replicate'] = (function() {
  $__fn = function($i_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_List_Lazy_take'])($i_0))(($GLOBALS['Data_List_Lazy_repeat'])($xs_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_range
$GLOBALS['Data_List_Lazy_range'] = (function() {
  $__fn = function($start_0 = null, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if (($start_0 > $end_1)) {
$__t1 = ((($GLOBALS['Data_List_Lazy_Types_unfoldableList'])['unfoldr'])(function($x_2 = null) use ($end_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($x_2 >= $end_1)) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", $x_2, ($x_2 - 1)));
goto end_branch_2;;
};
  $__t2 = new Phpurs_Data0("Nothing");
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($start_0);
goto end_branch_1;;
};
  $__t1 = ((($GLOBALS['Data_List_Lazy_Types_unfoldableList'])['unfoldr'])(function($x_2 = null) use ($end_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($x_2 <= $end_1)) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", $x_2, ($x_2 + 1)));
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($start_0);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_partition
$GLOBALS['Data_List_Lazy_partition'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])['foldr'])((function() use ($f_0) {
  $__fn = function($x_1 = null, $v_2 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($f_0)($x_1)) {
$__t0 = ["yes" => ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($v_2, $x_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $x_1, ($v_2)['yes']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "no" => ($v_2)['no']];
goto end_branch_0;;
};
  $__t0 = ["yes" => ($v_2)['yes'], "no" => ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($v_2, $x_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $x_1, ($v_2)['no']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(["yes" => $GLOBALS['Data_List_Lazy_Types_nil'], "no" => $GLOBALS['Data_List_Lazy_Types_nil']]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_null
$GLOBALS['Data_List_Lazy_null'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_isNothing']))($GLOBALS['Data_List_Lazy_uncons']);

// Data_List_Lazy_nubBy
$GLOBALS['Data_List_Lazy_nubBy'] = function($p_0 = null) {
  $__num = \func_num_args();
  $goStep_1_0 = null;
  $go_1_0 = null;
  $goStep_1_0 = (function() use (&$go_1_0, $p_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go_1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons"))) {
$v2_4_2 = ((($GLOBALS['Data_List_Internal_insertAndLookupBy'])($p_0))(($v1_3)->{'value0'}))($v_2);
$__t3 = null;;
if (($v2_4_2)['found']) {
$__t3 = ($GLOBALS['Data_List_Lazy_Types_step'])((($go_1_0)(($v2_4_2)['result']))(($v1_3)->{'value1'}));
goto end_branch_3;;
};
$__t3 = new Phpurs_Data2("Cons", ($v1_3)->{'value0'}, (($go_1_0)(($v2_4_2)['result']))(($v1_3)->{'value1'}));
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go_1_0 = (function() use (&$goStep_1_0) {
  $__fn = function($s_2 = null, $v_3 = null) use (&$goStep_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_4_4 = ($goStep_1_0)($s_2);
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_5 = null) use ($__local_var_4_4, $v_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_4_4)(($GLOBALS['Data_Lazy_force'])($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go_1_0)(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_nub
$GLOBALS['Data_List_Lazy_nub'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_nubBy'])(($dictOrd_0)['compare']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_newtypePattern
$GLOBALS['Data_List_Lazy_newtypePattern'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_mapMaybe
$GLOBALS['Data_List_Lazy_mapMaybe'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_mapMaybe_f_0 = $f_0;
  tco_loop_Data_List_Lazy_mapMaybe:;
  $f_0 = $__tco_var_Data_List_Lazy_mapMaybe_f_0;
  $go_1_0 = null;
  $go_1_0 = function($v_2 = null) use (&$__tco_var_Data_List_Lazy_mapMaybe_f_0, $f_0, &$go_1_0) {
  $__num = \func_num_args();
  $__tco_var_go_1_0_0_v_2 = $v_2;
  tco_loop_go_1_0_0:;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$v1_3_1 = ($f_0)(($v_2)->{'value0'});
$__t2 = null;;
if ((is_object($v1_3_1) && (($v1_3_1)->{'tag'} === "Nothing"))) {
$__tco_3 = ($GLOBALS['Data_List_Lazy_Types_step'])(($v_2)->{'value1'});
$__tco_var_go_1_0_0_v_2 = $__tco_3;
goto tco_loop_go_1_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object($v1_3_1) && (($v1_3_1)->{'tag'} === "Just"))) {
$__t2 = new Phpurs_Data2("Cons", ($v1_3_1)->{'value0'}, (($GLOBALS['Data_List_Lazy_mapMaybe'])($f_0))(($v_2)->{'value1'}));
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])['map'])($go_1_0)))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_some
$GLOBALS['Data_List_Lazy_some'] = (function() {
  $__fn = function($dictAlternative_0 = null, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)['Applicative0'])($GLOBALS['Prim_undefined']))['Apply0'])($GLOBALS['Prim_undefined']))['apply'])(((((((((($dictAlternative_0)['Plus1'])($GLOBALS['Prim_undefined']))['Alt0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']))['map'])($GLOBALS['Data_List_Lazy_Types_cons']))($v_2)))((($dictLazy_1)['defer'])(function($v1_3 = null) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_many'])($dictAlternative_0))($dictLazy_1))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_many
$GLOBALS['Data_List_Lazy_many'] = (function() {
  $__fn = function($dictAlternative_0 = null, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)['Plus1'])($GLOBALS['Prim_undefined']))['Alt0'])($GLOBALS['Prim_undefined']))['alt'])(((($GLOBALS['Data_List_Lazy_some'])($dictAlternative_0))($dictLazy_1))($v_2)))((((($dictAlternative_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])($GLOBALS['Data_List_Lazy_Types_nil']));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_length
$GLOBALS['Data_List_Lazy_length'] = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])['foldl'])((function() {
  $__fn = function($l_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($l_0 + 1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(0);

// Data_List_Lazy_last
$GLOBALS['Data_List_Lazy_last'] = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = function($v_1 = null) use (&$go_0_0) {
  $__num = \func_num_args();
  $__tco_var_go_0_0_0_v_1 = $v_1;
  tco_loop_go_0_0_0:;
  $v_1 = $__tco_var_go_0_0_0_v_1;
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Cons"))) {
$__t2 = null;;
if (($GLOBALS['Data_List_Lazy_null'])(($v_1)->{'value1'})) {
$__t2 = new Phpurs_Data1("Just", ($v_1)->{'value0'});
goto end_branch_2;;
};
$__tco_1 = ($GLOBALS['Data_List_Lazy_Types_step'])(($v_1)->{'value1'});
$__tco_var_go_0_0_0_v_1 = $__tco_1;
goto tco_loop_go_0_0_0;;
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
return (($GLOBALS['Control_Semigroupoid_composeImpl'])($go_0_0))($GLOBALS['Data_List_Lazy_Types_step']);
})();

// Data_List_Lazy_iterate
$GLOBALS['Data_List_Lazy_iterate'] = (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_0 = null;
  $go_2_0 = ($GLOBALS['Data_Lazy_defer'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))(function($v_3 = null) use ($f_0, &$go_2_0, $x_1) {
  $__num = \func_num_args();
  $__local_var_4_1 = ((($GLOBALS['Data_List_Lazy_Types_functorList'])['map'])($f_0))($go_2_0);
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_5 = null) use ($__local_var_4_1, $x_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $x_1, $__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_insertAt
$GLOBALS['Data_List_Lazy_insertAt'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_insertAt_v_0 = $v_0;
  $__tco_var_Data_List_Lazy_insertAt_v1_1 = $v1_1;
  $__tco_var_Data_List_Lazy_insertAt_v2_2 = $v2_2;
  tco_loop_Data_List_Lazy_insertAt:;
  $v_0 = $__tco_var_Data_List_Lazy_insertAt_v_0;
  $v1_1 = $__tco_var_Data_List_Lazy_insertAt_v1_1;
  $v2_2 = $__tco_var_Data_List_Lazy_insertAt_v2_2;
  $__res = match ($v_0) { 0 => ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($v1_1, $v2_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $v1_1, $v2_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), default => ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($v1_1, $v2_2, $v_0) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Lazy_force'])($v2_2);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data2("Cons", $v1_1, $GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Cons"))) {
$__t1 = new Phpurs_Data2("Cons", ($__local_var_4_0)->{'value0'}, ((($GLOBALS['Data_List_Lazy_insertAt'])(($v_0 - 1)))($v1_1))(($__local_var_4_0)->{'value1'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}) };
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_init
$GLOBALS['Data_List_Lazy_init'] = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = function($v_1 = null) use (&$go_0_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Cons"))) {
$__t6 = null;;
if (($GLOBALS['Data_List_Lazy_null'])(($v_1)->{'value1'})) {
$__t6 = new Phpurs_Data1("Just", $GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_6;;
};
$__local_var_2_2 = ($go_0_0)(($GLOBALS['Data_List_Lazy_Types_step'])(($v_1)->{'value1'}));
$__t3 = null;;
if ((is_object($__local_var_2_2) && (($__local_var_2_2)->{'tag'} === "Just"))) {
$__local_var_3_4 = ($v_1)->{'value0'};
$__local_var_4_5 = ($__local_var_2_2)->{'value0'};
$__t3 = new Phpurs_Data1("Just", ($GLOBALS['Data_Lazy_defer'])(function($v_5 = null) use ($__local_var_3_4, $__local_var_4_5) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_3_4, $__local_var_4_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_3;;
};
$__t3 = new Phpurs_Data0("Nothing");
end_branch_3:;
$__t6 = $__t3;
end_branch_6:;
$__t1 = $__t6;
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
return (($GLOBALS['Control_Semigroupoid_composeImpl'])($go_0_0))($GLOBALS['Data_List_Lazy_Types_step']);
})();

// Data_List_Lazy_index
$GLOBALS['Data_List_Lazy_index'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_0_0_v_2 = $v_2;
  $__tco_var_go_1_0_0_v1_3 = $v1_3;
  tco_loop_go_1_0_0:;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $v1_3 = $__tco_var_go_1_0_0_v1_3;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$__t3 = null;;
switch ($v1_3) {
case 0:
$__t3 = new Phpurs_Data1("Just", ($v_2)->{'value0'});
goto end_branch_3;;
break;
default:
;
break;
};
$__tco_1 = ($GLOBALS['Data_List_Lazy_Types_step'])(($v_2)->{'value1'});
$__tco_2 = ($v1_3 - 1);
$__tco_var_go_1_0_0_v_2 = $__tco_1;
$__tco_var_go_1_0_0_v1_3 = $__tco_2;
goto tco_loop_go_1_0_0;;
$__t3 = null;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go_1_0)(($GLOBALS['Data_List_Lazy_Types_step'])($xs_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_head
$GLOBALS['Data_List_Lazy_head'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($GLOBALS['Data_List_Lazy_uncons'])($xs_0);
  $__t1 = null;;
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($__local_var_1_0)->{'value0'})['head']);
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_transpose
$GLOBALS['Data_List_Lazy_transpose'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_transpose_xs_0 = $xs_0;
  tco_loop_Data_List_Lazy_transpose:;
  $xs_0 = $__tco_var_Data_List_Lazy_transpose_xs_0;
  $v_1_0 = ($GLOBALS['Data_List_Lazy_uncons'])($xs_0);
  $__t1 = null;;
  if ((is_object($v_1_0) && (($v_1_0)->{'tag'} === "Nothing"))) {
$__t1 = $xs_0;
goto end_branch_1;;
};
  if ((is_object($v_1_0) && (($v_1_0)->{'tag'} === "Just"))) {
$v1_2_2 = ($GLOBALS['Data_List_Lazy_uncons'])((($v_1_0)->{'value0'})['head']);
$__t3 = null;;
if ((is_object($v1_2_2) && (($v1_2_2)->{'tag'} === "Nothing"))) {
$__tco_4 = (($v_1_0)->{'value0'})['tail'];
$__tco_var_Data_List_Lazy_transpose_xs_0 = $__tco_4;
goto tco_loop_Data_List_Lazy_transpose;;
$__t3 = null;
goto end_branch_3;;
};
if ((is_object($v1_2_2) && (($v1_2_2)->{'tag'} === "Just"))) {
$__local_var_3_5 = (($v1_2_2)->{'value0'})['head'];
$__local_var_4_6 = (($v1_2_2)->{'value0'})['tail'];
$__local_var_5_7 = (($GLOBALS['Data_List_Lazy_mapMaybe'])($GLOBALS['Data_List_Lazy_head']))((($v_1_0)->{'value0'})['tail']);
$__local_var_6_8 = ($GLOBALS['Data_Lazy_defer'])(function($v_6 = null) use ($__local_var_3_5, $__local_var_5_7) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_3_5, $__local_var_5_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_7_9 = (($GLOBALS['Data_List_Lazy_mapMaybe'])($GLOBALS['Data_List_Lazy_tail']))((($v_1_0)->{'value0'})['tail']);
$__local_var_7_9 = ($GLOBALS['Data_List_Lazy_transpose'])(($GLOBALS['Data_Lazy_defer'])(function($v_8 = null) use ($__local_var_4_6, $__local_var_7_9) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_4_6, $__local_var_7_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__t3 = ($GLOBALS['Data_Lazy_defer'])(function($v_8 = null) use ($__local_var_6_8, $__local_var_7_9) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_6_8, $__local_var_7_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_groupBy
$GLOBALS['Data_List_Lazy_groupBy'] = function($eq_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_groupBy_eq_0 = $eq_0;
  tco_loop_Data_List_Lazy_groupBy:;
  $eq_0 = $__tco_var_Data_List_Lazy_groupBy_eq_0;
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])['map'])(function($v_1 = null) use ($eq_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Cons"))) {
$__local_var_2_1 = ($v_1)->{'value0'};
$v1_3_2 = (($GLOBALS['Data_List_Lazy_span'])(($eq_0)($__local_var_2_1)))(($v_1)->{'value1'});
$__local_var_4_3 = ($v1_3_2)['init'];
$__t0 = new Phpurs_Data2("Cons", ($GLOBALS['Data_Lazy_defer'])(function($v2_5 = null) use ($__local_var_2_1, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $__local_var_2_1, $__local_var_4_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($GLOBALS['Data_List_Lazy_groupBy'])($eq_0))(($v1_3_2)['rest']));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_group
$GLOBALS['Data_List_Lazy_group'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_groupBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_fromStep
$GLOBALS['Data_List_Lazy_fromStep'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))(($GLOBALS['Data_Lazy_applicativeLazy'])['pure']);

// Data_List_Lazy_insertBy
$GLOBALS['Data_List_Lazy_insertBy'] = (function() {
  $__fn = function($cmp_0 = null, $x_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_insertBy_cmp_0 = $cmp_0;
  $__tco_var_Data_List_Lazy_insertBy_x_1 = $x_1;
  $__tco_var_Data_List_Lazy_insertBy_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_insertBy:;
  $cmp_0 = $__tco_var_Data_List_Lazy_insertBy_cmp_0;
  $x_1 = $__tco_var_Data_List_Lazy_insertBy_x_1;
  $xs_2 = $__tco_var_Data_List_Lazy_insertBy_xs_2;
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($cmp_0, $x_1, $xs_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Lazy_force'])($xs_2);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data2("Cons", $x_1, $GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Cons"))) {
$__t2 = null;;
if ((is_object((($cmp_0)($x_1))(($__local_var_4_0)->{'value0'})) && (((($cmp_0)($x_1))(($__local_var_4_0)->{'value0'}))->{'tag'} === "GT"))) {
$__t2 = new Phpurs_Data2("Cons", ($__local_var_4_0)->{'value0'}, ((($GLOBALS['Data_List_Lazy_insertBy'])($cmp_0))($x_1))(($__local_var_4_0)->{'value1'}));
goto end_branch_2;;
};
$__t2 = new Phpurs_Data2("Cons", $x_1, ($GLOBALS['Data_List_Lazy_fromStep'])($__local_var_4_0));
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
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_insert
$GLOBALS['Data_List_Lazy_insert'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_insertBy'])(($dictOrd_0)['compare']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_fromFoldable
$GLOBALS['Data_List_Lazy_fromFoldable'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_0)['foldr'])($GLOBALS['Data_List_Lazy_Types_cons']))($GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_foldrLazy
$GLOBALS['Data_List_Lazy_foldrLazy'] = (function() {
  $__fn = function($dictLazy_0 = null, $op_1 = null, $z_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go_3_0 = null;
  $go_3_0 = function($xs_4 = null) use ($dictLazy_0, &$go_3_0, $op_1, $z_2) {
  $__num = \func_num_args();
  $v_5_1 = ($GLOBALS['Data_List_Lazy_Types_step'])($xs_4);
  $__t2 = null;;
  if ((is_object($v_5_1) && (($v_5_1)->{'tag'} === "Cons"))) {
$__local_var_6_3 = ($v_5_1)->{'value0'};
$__local_var_7_4 = ($v_5_1)->{'value1'};
$__t2 = (($dictLazy_0)['defer'])(function($v1_8 = null) use ($__local_var_6_3, $__local_var_7_4, &$go_3_0, $op_1) {
  $__num = \func_num_args();
  $__res = (($op_1)($__local_var_6_3))(($go_3_0)($__local_var_7_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_2;;
};
  if ((is_object($v_5_1) && (($v_5_1)->{'tag'} === "Nil"))) {
$__t2 = $z_2;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_3_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_foldM
$GLOBALS['Data_List_Lazy_foldM'] = (function() {
  $__fn = function($dictMonad_0 = null, $f_1 = null, $b_2 = null, $xs_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_List_Lazy_foldM_dictMonad_0 = $dictMonad_0;
  $__tco_var_Data_List_Lazy_foldM_f_1 = $f_1;
  $__tco_var_Data_List_Lazy_foldM_b_2 = $b_2;
  $__tco_var_Data_List_Lazy_foldM_xs_3 = $xs_3;
  tco_loop_Data_List_Lazy_foldM:;
  $dictMonad_0 = $__tco_var_Data_List_Lazy_foldM_dictMonad_0;
  $f_1 = $__tco_var_Data_List_Lazy_foldM_f_1;
  $b_2 = $__tco_var_Data_List_Lazy_foldM_b_2;
  $xs_3 = $__tco_var_Data_List_Lazy_foldM_xs_3;
  $v_4_0 = ($GLOBALS['Data_List_Lazy_uncons'])($xs_3);
  $__t1 = null;;
  if ((is_object($v_4_0) && (($v_4_0)->{'tag'} === "Nothing"))) {
$__t1 = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])($b_2);
goto end_branch_1;;
};
  if ((is_object($v_4_0) && (($v_4_0)->{'tag'} === "Just"))) {
$__local_var_5_2 = (($v_4_0)->{'value0'})['tail'];
$__t1 = ((((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])((($f_1)($b_2))((($v_4_0)->{'value0'})['head'])))(function($b__prime___6 = null) use ($__local_var_5_2, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_foldM'])($dictMonad_0))($f_1))($b__prime___6))($__local_var_5_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_findIndex
$GLOBALS['Data_List_Lazy_findIndex'] = function($fn_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use ($fn_0, &$go_1_0) {
  $__fn = function($n_2 = null, $list_3 = null) use ($fn_0, &$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_0_0_n_2 = $n_2;
  $__tco_var_go_1_0_0_list_3 = $list_3;
  tco_loop_go_1_0_0:;
  $n_2 = $__tco_var_go_1_0_0_n_2;
  $list_3 = $__tco_var_go_1_0_0_list_3;
  $__local_var_4_0 = ($GLOBALS['Data_List_Lazy_uncons'])($list_3);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Just"))) {
$__t4 = null;;
if (($fn_0)((($__local_var_4_0)->{'value0'})['head'])) {
$__t4 = new Phpurs_Data1("Just", $n_2);
goto end_branch_4;;
};
$__tco_2 = ($n_2 + 1);
$__tco_3 = (($__local_var_4_0)->{'value0'})['tail'];
$__tco_var_go_1_0_0_n_2 = $__tco_2;
$__tco_var_go_1_0_0_list_3 = $__tco_3;
goto tco_loop_go_1_0_0;;
$__t4 = null;
end_branch_4:;
$__t1 = $__t4;
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go_1_0)(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_findLastIndex
$GLOBALS['Data_List_Lazy_findLastIndex'] = (function() {
  $__fn = function($fn_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($GLOBALS['Data_List_Lazy_findIndex'])($fn_0))(($GLOBALS['Data_List_Lazy_reverse'])($xs_1));
  $__t1 = null;;
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ((($GLOBALS['Data_List_Lazy_length'])($xs_1) - 1) - ($__local_var_2_0)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_filterM
$GLOBALS['Data_List_Lazy_filterM'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_filterM_dictMonad_0 = $dictMonad_0;
  tco_loop_Data_List_Lazy_filterM:;
  $dictMonad_0 = $__tco_var_Data_List_Lazy_filterM_dictMonad_0;
  $__local_var_1_0 = (($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']);
  $__local_var_2_1 = (($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']);
  $__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($p_3 = null, $list_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_5_2 = ($GLOBALS['Data_List_Lazy_uncons'])($list_4);
  $__t3 = null;;
  if ((is_object($v_5_2) && (($v_5_2)->{'tag'} === "Nothing"))) {
$__t3 = (($__local_var_1_0)['pure'])($GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_3;;
};
  if ((is_object($v_5_2) && (($v_5_2)->{'tag'} === "Just"))) {
$__local_var_6_4 = (($v_5_2)->{'value0'})['head'];
$__local_var_7_5 = (($v_5_2)->{'value0'})['tail'];
$__t3 = ((($__local_var_2_1)['bind'])(($p_3)($__local_var_6_4)))(function($b_8 = null) use ($__local_var_1_0, $__local_var_2_1, $__local_var_6_4, $__local_var_7_5, $dictMonad_0, $p_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)['bind'])(((($GLOBALS['Data_List_Lazy_filterM'])($dictMonad_0))($p_3))($__local_var_7_5)))(function($xs__prime___9 = null) use ($__local_var_1_0, $__local_var_6_4, $b_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($b_8) {
$__t6 = ($GLOBALS['Data_Lazy_defer'])(function($v_10 = null) use ($__local_var_6_4, $xs__prime___9) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_6_4, $xs__prime___9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_6;;
};
  $__t6 = $xs__prime___9;
  end_branch_6:;
  $__res = (($__local_var_1_0)['pure'])($__t6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_filter
$GLOBALS['Data_List_Lazy_filter'] = function($p_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_filter_p_0 = $p_0;
  tco_loop_Data_List_Lazy_filter:;
  $p_0 = $__tco_var_Data_List_Lazy_filter_p_0;
  $go_1_0 = null;
  $go_1_0 = function($v_2 = null) use (&$__tco_var_Data_List_Lazy_filter_p_0, &$go_1_0, $p_0) {
  $__num = \func_num_args();
  $__tco_var_go_1_0_0_v_2 = $v_2;
  tco_loop_go_1_0_0:;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$__t2 = null;;
if (($p_0)(($v_2)->{'value0'})) {
$__t2 = new Phpurs_Data2("Cons", ($v_2)->{'value0'}, (($GLOBALS['Data_List_Lazy_filter'])($p_0))(($v_2)->{'value1'}));
goto end_branch_2;;
};
$__tco_1 = ($GLOBALS['Data_List_Lazy_Types_step'])(($v_2)->{'value1'});
$__tco_var_go_1_0_0_v_2 = $__tco_1;
goto tco_loop_go_1_0_0;;
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])['map'])($go_1_0)))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_intersectBy
$GLOBALS['Data_List_Lazy_intersectBy'] = (function() {
  $__fn = function($eq_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Data_List_Lazy_filter'])(function($x_3 = null) use ($eq_0, $ys_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_Lazy_any'])(($eq_0)($x_3)))($ys_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_intersect
$GLOBALS['Data_List_Lazy_intersect'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_intersectBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_nubByEq
$GLOBALS['Data_List_Lazy_nubByEq'] = function($eq_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_Lazy_nubByEq_eq_0 = $eq_0;
  tco_loop_Data_List_Lazy_nubByEq:;
  $eq_0 = $__tco_var_Data_List_Lazy_nubByEq_eq_0;
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])['map'])(function($v_1 = null) use ($eq_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Cons"))) {
$__local_var_2_1 = ($v_1)->{'value0'};
$__t0 = new Phpurs_Data2("Cons", $__local_var_2_1, (($GLOBALS['Data_List_Lazy_nubByEq'])($eq_0))((($GLOBALS['Data_List_Lazy_filter'])(function($y_3 = null) use ($__local_var_2_1, $eq_0) {
  $__num = \func_num_args();
  $__res = ( ! (($eq_0)($__local_var_2_1))($y_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_1)->{'value1'})));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_nubEq
$GLOBALS['Data_List_Lazy_nubEq'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_nubByEq'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_eqPattern
$GLOBALS['Data_List_Lazy_eqPattern'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_eq1List'])['eq1'])($dictEq_0))($x_1))($y_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_ordPattern
$GLOBALS['Data_List_Lazy_ordPattern'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)['Eq0'])($GLOBALS['Prim_undefined']);
  $eqPattern1_2_1 = ["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_eq1List'])['eq1'])($__local_var_1_0))($x_2))($y_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["compare" => (function() use ($dictOrd_0) {
  $__fn = function($x_3 = null, $y_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_ordList'])($dictOrd_0))['compare'])($x_3))($y_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_3 = null) use ($eqPattern1_2_1) {
  $__num = \func_num_args();
  $__res = $eqPattern1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_elemLastIndex
$GLOBALS['Data_List_Lazy_elemLastIndex'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_Lazy_findLastIndex'])(function($v_2 = null) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)['eq'])($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_elemIndex
$GLOBALS['Data_List_Lazy_elemIndex'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_Lazy_findIndex'])(function($v_2 = null) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)['eq'])($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_dropWhile
$GLOBALS['Data_List_Lazy_dropWhile'] = function($p_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = function($v_2 = null) use (&$go_1_0, $p_0) {
  $__num = \func_num_args();
  $__tco_var_go_1_0_0_v_2 = $v_2;
  tco_loop_go_1_0_0:;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $__t0 = null;;
  if (((is_object($v_2) && (($v_2)->{'tag'} === "Cons")) && ($p_0)(($v_2)->{'value0'}))) {
$__tco_1 = ($GLOBALS['Data_List_Lazy_Types_step'])(($v_2)->{'value1'});
$__tco_var_go_1_0_0_v_2 = $__tco_1;
goto tco_loop_go_1_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  $__t0 = ($GLOBALS['Data_List_Lazy_fromStep'])($v_2);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($go_1_0))($GLOBALS['Data_List_Lazy_Types_step']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_drop
$GLOBALS['Data_List_Lazy_drop'] = function($n_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_0_0_v_2 = $v_2;
  $__tco_var_go_1_0_0_v1_3 = $v1_3;
  tco_loop_go_1_0_0:;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $v1_3 = $__tco_var_go_1_0_0_v1_3;
  $__t0 = null;;
  switch ($v_2) {
case 0:
$__t0 = $v1_3;
goto end_branch_0;;
break;
default:
;
break;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons"))) {
$__tco_1 = ($v_2 - 1);
$__tco_2 = ($GLOBALS['Data_List_Lazy_Types_step'])(($v1_3)->{'value1'});
$__tco_var_go_1_0_0_v_2 = $__tco_1;
$__tco_var_go_1_0_0_v1_3 = $__tco_2;
goto tco_loop_go_1_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])['map'])(($go_1_0)($n_0))))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_slice
$GLOBALS['Data_List_Lazy_slice'] = (function() {
  $__fn = function($start_0 = null, $end_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Data_List_Lazy_take'])(($end_1 - $start_0)))((($GLOBALS['Data_List_Lazy_drop'])($start_0))($xs_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_deleteBy
$GLOBALS['Data_List_Lazy_deleteBy'] = (function() {
  $__fn = function($eq_0 = null, $x_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_deleteBy_eq_0 = $eq_0;
  $__tco_var_Data_List_Lazy_deleteBy_x_1 = $x_1;
  $__tco_var_Data_List_Lazy_deleteBy_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_deleteBy:;
  $eq_0 = $__tco_var_Data_List_Lazy_deleteBy_eq_0;
  $x_1 = $__tco_var_Data_List_Lazy_deleteBy_x_1;
  $xs_2 = $__tco_var_Data_List_Lazy_deleteBy_xs_2;
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($eq_0, $x_1, $xs_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Lazy_force'])($xs_2);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Cons"))) {
$__t2 = null;;
if ((($eq_0)($x_1))(($__local_var_4_0)->{'value0'})) {
$__t2 = ($GLOBALS['Data_List_Lazy_Types_step'])(($__local_var_4_0)->{'value1'});
goto end_branch_2;;
};
$__t2 = new Phpurs_Data2("Cons", ($__local_var_4_0)->{'value0'}, ((($GLOBALS['Data_List_Lazy_deleteBy'])($eq_0))($x_1))(($__local_var_4_0)->{'value1'}));
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
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_unionBy
$GLOBALS['Data_List_Lazy_unionBy'] = (function() {
  $__fn = function($eq_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])['append'])($xs_1))((((($GLOBALS['Data_List_Lazy_Types_foldableList'])['foldl'])((function() use ($eq_0) {
  $__fn = function($b_3 = null, $a_4 = null) use ($eq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_deleteBy'])($eq_0))($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($GLOBALS['Data_List_Lazy_nubByEq'])($eq_0))($ys_2)))($xs_1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_union
$GLOBALS['Data_List_Lazy_union'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_unionBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_deleteAt
$GLOBALS['Data_List_Lazy_deleteAt'] = (function() {
  $__fn = function($n_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_Lazy_deleteAt_n_0 = $n_0;
  $__tco_var_Data_List_Lazy_deleteAt_xs_1 = $xs_1;
  tco_loop_Data_List_Lazy_deleteAt:;
  $n_0 = $__tco_var_Data_List_Lazy_deleteAt_n_0;
  $xs_1 = $__tco_var_Data_List_Lazy_deleteAt_xs_1;
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_2 = null) use ($n_0, $xs_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = ($GLOBALS['Data_Lazy_force'])($xs_1);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Cons"))) {
$__t1 = match ($n_0) { 0 => ($GLOBALS['Data_List_Lazy_Types_step'])(($__local_var_3_0)->{'value1'}), default => new Phpurs_Data2("Cons", ($__local_var_3_0)->{'value0'}, (($GLOBALS['Data_List_Lazy_deleteAt'])(($n_0 - 1)))(($__local_var_3_0)->{'value1'})) };
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_delete
$GLOBALS['Data_List_Lazy_delete'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_deleteBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_difference
$GLOBALS['Data_List_Lazy_difference'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_Lazy_Types_foldableList'])['foldl'])((function() use ($dictEq_0) {
  $__fn = function($b_1 = null, $a_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_deleteBy'])(($dictEq_0)['eq']))($a_2))($b_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_cycle
$GLOBALS['Data_List_Lazy_cycle'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = ($GLOBALS['Data_Lazy_defer'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))(function($v_2 = null) use (&$go_1_0, $xs_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])['append'])($xs_0))($go_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_concatMap
$GLOBALS['Data_List_Lazy_concatMap'] = (function() {
  $__fn = function($b_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_Types_bindList'])['bind'])($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_concat
$GLOBALS['Data_List_Lazy_concat'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_bindList'])['bind'])($v_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_catMaybes
$GLOBALS['Data_List_Lazy_catMaybes'] = ($GLOBALS['Data_List_Lazy_mapMaybe'])(($GLOBALS['Control_Category_categoryFn'])['identity']);

// Data_List_Lazy_alterAt
$GLOBALS['Data_List_Lazy_alterAt'] = (function() {
  $__fn = function($n_0 = null, $f_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_alterAt_n_0 = $n_0;
  $__tco_var_Data_List_Lazy_alterAt_f_1 = $f_1;
  $__tco_var_Data_List_Lazy_alterAt_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_alterAt:;
  $n_0 = $__tco_var_Data_List_Lazy_alterAt_n_0;
  $f_1 = $__tco_var_Data_List_Lazy_alterAt_f_1;
  $xs_2 = $__tco_var_Data_List_Lazy_alterAt_xs_2;
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_3 = null) use ($f_1, $n_0, $xs_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Lazy_force'])($xs_2);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Cons"))) {
$__t2 = null;;
switch ($n_0) {
case 0:
$v2_5_3 = ($f_1)(($__local_var_4_0)->{'value0'});
$__t4 = null;;
if ((is_object($v2_5_3) && (($v2_5_3)->{'tag'} === "Nothing"))) {
$__t4 = ($GLOBALS['Data_List_Lazy_Types_step'])(($__local_var_4_0)->{'value1'});
goto end_branch_4;;
};
if ((is_object($v2_5_3) && (($v2_5_3)->{'tag'} === "Just"))) {
$__t4 = new Phpurs_Data2("Cons", ($v2_5_3)->{'value0'}, ($__local_var_4_0)->{'value1'});
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t2 = $__t4;
goto end_branch_2;;
break;
default:
;
break;
};
$__t2 = new Phpurs_Data2("Cons", ($__local_var_4_0)->{'value0'}, ((($GLOBALS['Data_List_Lazy_alterAt'])(($n_0 - 1)))($f_1))(($__local_var_4_0)->{'value1'}));
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
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_modifyAt
$GLOBALS['Data_List_Lazy_modifyAt'] = (function() {
  $__fn = function($n_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_List_Lazy_alterAt'])($n_0))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($f_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

