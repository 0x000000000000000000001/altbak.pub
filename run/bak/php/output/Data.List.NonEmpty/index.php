<?php

namespace Data\List\NonEmpty;

// ALL IMPORTS: Control.Bind, Control.Category, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.List, Data.List.NonEmpty, Data.List.Types, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Traversable, Data.Tuple, Data.Unfoldable, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Bind, Control.Category, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.List, Data.List.NonEmpty, Data.List.Types, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Traversable, Data.Tuple, Data.Unfoldable, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.NonEmpty/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semigroup.Traversable/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
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


// Data_List_NonEmpty_zipWith
$GLOBALS['Data_List_NonEmpty_zipWith'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = (function() use ($f_0, &$go__3_0) {
  $__fn = function($v_4 = null, $v1_5 = null, $v2_6 = null) use ($f_0, &$go__3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__3_0_0_v_4 = $v_4;
  $__tco_var_go__3_0_0_v1_5 = $v1_5;
  $__tco_var_go__3_0_0_v2_6 = $v2_6;
  tco_loop_go__3_0_0:;
  $v_4 = $__tco_var_go__3_0_0_v_4;
  $v1_5 = $__tco_var_go__3_0_0_v1_5;
  $v2_6 = $__tco_var_go__3_0_0_v2_6;
  $__t0 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Nil"))) {
$__t0 = $v2_6;
goto end_branch_0;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$__t0 = $v2_6;
goto end_branch_0;;
};
  if (((is_object($v_4) && (($v_4)->{'tag'} === "Cons")) && (is_object($v1_5) && (($v1_5)->{'tag'} === "Cons")))) {
$__tco_1 = ($v_4)->{'value1'};
$__tco_2 = ($v1_5)->{'value1'};
$__tco_3 = new Phpurs_Data2("Cons", (($f_0)(($v_4)->{'value0'}))(($v1_5)->{'value0'}), $v2_6);
$__tco_var_go__3_0_0_v_4 = $__tco_1;
$__tco_var_go__3_0_0_v1_5 = $__tco_2;
$__tco_var_go__3_0_0_v2_6 = $__tco_3;
goto tco_loop_go__3_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $go__4_1 = null;
  $go__4_1 = (function() use (&$go__4_1) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$go__4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_1_1_v_5 = $v_5;
  $__tco_var_go__4_1_1_v1_6 = $v1_6;
  tco_loop_go__4_1_1:;
  $v_5 = $__tco_var_go__4_1_1_v_5;
  $v1_6 = $__tco_var_go__4_1_1_v1_6;
  $__t1 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go__4_1_1_v_5 = $__tco_2;
$__tco_var_go__4_1_1_v1_6 = $__tco_3;
goto tco_loop_go__4_1_1;;
$__t1 = null;
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
  $__res = new Phpurs_Data2("NonEmpty", (($f_0)(($v_1)->{'value0'}))(($v1_2)->{'value0'}), (($go__4_1)(new Phpurs_Data0("Nil")))(((($go__3_0)(($v_1)->{'value1'}))(($v1_2)->{'value1'}))(new Phpurs_Data0("Nil"))));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_zipWithA
$GLOBALS['Data_List_NonEmpty_zipWithA'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $sequence11_1_0 = (($GLOBALS['Data_List_Types_traversable1NonEmptyList'])['sequence1'])((($dictApplicative_0)['Apply0'])(null));
  $__res = (function() use ($sequence11_1_0) {
  $__fn = function($f_2 = null, $xs_3 = null, $ys_4 = null) use ($sequence11_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($sequence11_1_0)(((($GLOBALS['Data_List_NonEmpty_zipWith'])($f_2))($xs_3))($ys_4));
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

// Data_List_NonEmpty_zip
$GLOBALS['Data_List_NonEmpty_zip'] = ($GLOBALS['Data_List_NonEmpty_zipWith'])($GLOBALS['Data_Tuple_Tuple']);

// Data_List_NonEmpty_wrappedOperation2
$GLOBALS['Data_List_NonEmpty_wrappedOperation2'] = (function() {
  $__fn = function($name_0 = null, $f_1 = null, $v_2 = null, $v1_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $v2_4_0 = (($f_1)(new Phpurs_Data2("Cons", ($v_2)->{'value0'}, ($v_2)->{'value1'})))(new Phpurs_Data2("Cons", ($v1_3)->{'value0'}, ($v1_3)->{'value1'}));
  $__t1 = null;;
  if ((is_object($v2_4_0) && (($v2_4_0)->{'tag'} === "Cons"))) {
$__t1 = new Phpurs_Data2("NonEmpty", ($v2_4_0)->{'value0'}, ($v2_4_0)->{'value1'});
goto end_branch_1;;
};
  if ((is_object($v2_4_0) && (($v2_4_0)->{'tag'} === "Nil"))) {
$__t1 = ($GLOBALS['Partial__crashWith'])(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("Impossible: empty list in NonEmptyList "))($name_0));
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

// Data_List_NonEmpty_wrappedOperation
$GLOBALS['Data_List_NonEmpty_wrappedOperation'] = (function() {
  $__fn = function($name_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v1_3_0 = ($f_1)(new Phpurs_Data2("Cons", ($v_2)->{'value0'}, ($v_2)->{'value1'}));
  $__t1 = null;;
  if ((is_object($v1_3_0) && (($v1_3_0)->{'tag'} === "Cons"))) {
$__t1 = new Phpurs_Data2("NonEmpty", ($v1_3_0)->{'value0'}, ($v1_3_0)->{'value1'});
goto end_branch_1;;
};
  if ((is_object($v1_3_0) && (($v1_3_0)->{'tag'} === "Nil"))) {
$__t1 = ($GLOBALS['Partial__crashWith'])(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("Impossible: empty list in NonEmptyList "))($name_0));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_updateAt
$GLOBALS['Data_List_NonEmpty_updateAt'] = (function() {
  $__fn = function($i_0 = null, $a_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if (((($GLOBALS['Data_Eq_eqInt'])['eq'])($i_0))(0)) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("NonEmpty", $a_1, ($v_2)->{'value1'}));
goto end_branch_1;;
};
  $__local_var_3_0 = ($v_2)->{'value0'};
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($v1_4 = null) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $__local_var_3_0, $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(((($GLOBALS['Data_List_updateAt'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($i_0))(1)))($a_1))(($v_2)->{'value1'}));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_unzip
$GLOBALS['Data_List_NonEmpty_unzip'] = function($ts_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ((($GLOBALS['Data_List_Types_functorNonEmptyList'])['map'])($GLOBALS['Data_Tuple_fst']))($ts_0), ((($GLOBALS['Data_List_Types_functorNonEmptyList'])['map'])($GLOBALS['Data_Tuple_snd']))($ts_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_unsnoc
$GLOBALS['Data_List_NonEmpty_unsnoc'] = function($v_0 = null) {
  $__num = \func_num_args();
  $v1_1_0 = ($GLOBALS['Data_List_unsnoc'])(($v_0)->{'value1'});
  $__t1 = null;;
  if ((is_object($v1_1_0) && (($v1_1_0)->{'tag'} === "Nothing"))) {
$__t1 = ["init" => new Phpurs_Data0("Nil"), "last" => ($v_0)->{'value0'}];
goto end_branch_1;;
};
  if ((is_object($v1_1_0) && (($v1_1_0)->{'tag'} === "Just"))) {
$__t1 = ["init" => new Phpurs_Data2("Cons", ($v_0)->{'value0'}, (($v1_1_0)->{'value0'})['init']), "last" => (($v1_1_0)->{'value0'})['last']];
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

// Data_List_NonEmpty_unionBy
$GLOBALS['Data_List_NonEmpty_unionBy'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation2'])("unionBy")))($GLOBALS['Data_List_unionBy']);

// Data_List_NonEmpty_union
$GLOBALS['Data_List_NonEmpty_union'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation2'])("union"))(($GLOBALS['Data_List_union'])($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_uncons
$GLOBALS['Data_List_NonEmpty_uncons'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ["head" => ($v_0)->{'value0'}, "tail" => ($v_0)->{'value1'}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_toList
$GLOBALS['Data_List_NonEmpty_toList'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", ($v_0)->{'value0'}, ($v_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_toUnfoldable
$GLOBALS['Data_List_NonEmpty_toUnfoldable'] = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($dictUnfoldable_0)['unfoldr'])(function($xs_1 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($xs_1) && (($xs_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($xs_1) && (($xs_1)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", ["head" => ($xs_1)->{'value0'}, "tail" => ($xs_1)->{'value1'}]);
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($rec_2 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($rec_2)['head'], ($rec_2)['tail']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_List_NonEmpty_toList']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_tail
$GLOBALS['Data_List_NonEmpty_tail'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_sortBy
$GLOBALS['Data_List_NonEmpty_sortBy'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("sortBy")))($GLOBALS['Data_List_sortBy']);

// Data_List_NonEmpty_sort
$GLOBALS['Data_List_NonEmpty_sort'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = function($xs_2 = null) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_NonEmpty_sortBy'])($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_snoc
$GLOBALS['Data_List_NonEmpty_snoc'] = (function() {
  $__fn = function($v_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", ($v_0)->{'value0'}, (((($GLOBALS['Data_List_Types_foldableList'])['foldr'])($GLOBALS['Data_List_Types_Cons']))(new Phpurs_Data2("Cons", $y_1, new Phpurs_Data0("Nil"))))(($v_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_singleton
$GLOBALS['Data_List_NonEmpty_singleton'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($a_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $a_0, ($GLOBALS['Data_List_Types_plusList'])['empty']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_List_NonEmpty_snoc'
$GLOBALS['Data_List_NonEmpty_snoc__prime__'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data2("NonEmpty", ($v_0)->{'value0'}, (((($GLOBALS['Data_List_Types_foldableList'])['foldr'])($GLOBALS['Data_List_Types_Cons']))(new Phpurs_Data2("Cons", $v1_1, new Phpurs_Data0("Nil"))))(($v_0)->{'value1'}));
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nil"))) {
$__t0 = ($GLOBALS['Data_List_NonEmpty_singleton'])($v1_1);
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

// Data_List_NonEmpty_reverse
$GLOBALS['Data_List_NonEmpty_reverse'] = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("reverse"))($GLOBALS['Data_List_reverse']);

// Data_List_NonEmpty_nubEq
$GLOBALS['Data_List_NonEmpty_nubEq'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("nubEq"))(($GLOBALS['Data_List_nubEq'])($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_nubByEq
$GLOBALS['Data_List_NonEmpty_nubByEq'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("nubByEq")))($GLOBALS['Data_List_nubByEq']);

// Data_List_NonEmpty_nubBy
$GLOBALS['Data_List_NonEmpty_nubBy'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("nubBy")))($GLOBALS['Data_List_nubBy']);

// Data_List_NonEmpty_nub
$GLOBALS['Data_List_NonEmpty_nub'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("nub"))(($GLOBALS['Data_List_nubBy'])(($dictOrd_0)['compare']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_modifyAt
$GLOBALS['Data_List_NonEmpty_modifyAt'] = (function() {
  $__fn = function($i_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if (((($GLOBALS['Data_Eq_eqInt'])['eq'])($i_0))(0)) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("NonEmpty", ($f_1)(($v_2)->{'value0'}), ($v_2)->{'value1'}));
goto end_branch_1;;
};
  $__local_var_3_0 = ($v_2)->{'value0'};
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($v1_4 = null) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $__local_var_3_0, $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(((($GLOBALS['Data_List_alterAt'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($i_0))(1)))(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_Maybe_Just']))($f_1)))(($v_2)->{'value1'}));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_lift
$GLOBALS['Data_List_NonEmpty_lift'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)(new Phpurs_Data2("Cons", ($v_1)->{'value0'}, ($v_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_mapMaybe
$GLOBALS['Data_List_NonEmpty_mapMaybe'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_mapMaybe']);

// Data_List_NonEmpty_partition
$GLOBALS['Data_List_NonEmpty_partition'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_partition']);

// Data_List_NonEmpty_span
$GLOBALS['Data_List_NonEmpty_span'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_span']);

// Data_List_NonEmpty_take
$GLOBALS['Data_List_NonEmpty_take'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_take']);

// Data_List_NonEmpty_takeWhile
$GLOBALS['Data_List_NonEmpty_takeWhile'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_takeWhile']);

// Data_List_NonEmpty_length
$GLOBALS['Data_List_NonEmpty_length'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semiring_semiringInt'])['add'])(1))(($GLOBALS['Data_List_length'])(($v_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_last
$GLOBALS['Data_List_NonEmpty_last'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "Cons"))) {
$__t1 = null;;
if ((is_object((($v_0)->{'value1'})->{'value1'}) && (((($v_0)->{'value1'})->{'value1'})->{'tag'} === "Nil"))) {
$__t1 = (($GLOBALS['Control_Category_categoryFn'])['identity'])((($v_0)->{'value1'})->{'value0'});
goto end_branch_1;;
};
if ((is_object(($GLOBALS['Data_List_last'])((($v_0)->{'value1'})->{'value1'})) && ((($GLOBALS['Data_List_last'])((($v_0)->{'value1'})->{'value1'}))->{'tag'} === "Nothing"))) {
$__t1 = ($v_0)->{'value0'};
goto end_branch_1;;
};
if ((is_object(($GLOBALS['Data_List_last'])((($v_0)->{'value1'})->{'value1'})) && ((($GLOBALS['Data_List_last'])((($v_0)->{'value1'})->{'value1'}))->{'tag'} === "Just"))) {
$__t1 = (($GLOBALS['Control_Category_categoryFn'])['identity'])((($GLOBALS['Data_List_last'])((($v_0)->{'value1'})->{'value1'}))->{'value0'});
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  $__t0 = ($v_0)->{'value0'};
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_intersectBy
$GLOBALS['Data_List_NonEmpty_intersectBy'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation2'])("intersectBy")))($GLOBALS['Data_List_intersectBy']);

// Data_List_NonEmpty_intersect
$GLOBALS['Data_List_NonEmpty_intersect'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation2'])("intersect"))(($GLOBALS['Data_List_intersect'])($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_insertAt
$GLOBALS['Data_List_NonEmpty_insertAt'] = (function() {
  $__fn = function($i_0 = null, $a_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if (((($GLOBALS['Data_Eq_eqInt'])['eq'])($i_0))(0)) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("NonEmpty", $a_1, new Phpurs_Data2("Cons", ($v_2)->{'value0'}, ($v_2)->{'value1'})));
goto end_branch_1;;
};
  $__local_var_3_0 = ($v_2)->{'value0'};
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($v1_4 = null) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $__local_var_3_0, $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(((($GLOBALS['Data_List_insertAt'])(((($GLOBALS['Data_Ring_ringInt'])['sub'])($i_0))(1)))($a_1))(($v_2)->{'value1'}));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_init
$GLOBALS['Data_List_NonEmpty_init'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = ($v_1)['init'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_List_unsnoc'])(($v_0)->{'value1'}));
  $__t1 = null;;
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data2("Cons", ($v_0)->{'value0'}, ($__local_var_1_0)->{'value0'});
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

// Data_List_NonEmpty_index
$GLOBALS['Data_List_NonEmpty_index'] = (function() {
  $__fn = function($v_0 = null, $i_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (((($GLOBALS['Data_Eq_eqInt'])['eq'])($i_1))(0)) {
$__t0 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
goto end_branch_0;;
};
  $__t0 = (($GLOBALS['Data_List_index'])(($v_0)->{'value1'}))(((($GLOBALS['Data_Ring_ringInt'])['sub'])($i_1))(1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_head
$GLOBALS['Data_List_NonEmpty_head'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_groupBy
$GLOBALS['Data_List_NonEmpty_groupBy'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("groupBy")))($GLOBALS['Data_List_groupBy']);

// Data_List_NonEmpty_groupAllBy
$GLOBALS['Data_List_NonEmpty_groupAllBy'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("groupAllBy")))($GLOBALS['Data_List_groupAllBy']);

// Data_List_NonEmpty_groupAll
$GLOBALS['Data_List_NonEmpty_groupAll'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("groupAll"))(($GLOBALS['Data_List_groupAll'])($dictOrd_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_group
$GLOBALS['Data_List_NonEmpty_group'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("group"))(($GLOBALS['Data_List_group'])($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_fromList
$GLOBALS['Data_List_NonEmpty_fromList'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data2("NonEmpty", ($v_0)->{'value0'}, ($v_0)->{'value1'}));
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

// Data_List_NonEmpty_fromFoldable
$GLOBALS['Data_List_NonEmpty_fromFoldable'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_fromList']))(((($dictFoldable_0)['foldr'])($GLOBALS['Data_List_Types_Cons']))(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_foldM
$GLOBALS['Data_List_NonEmpty_foldM'] = (function() {
  $__fn = function($dictMonad_0 = null, $f_1 = null, $b_2 = null, $v_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__local_var_4_0 = ($v_3)->{'value1'};
  $__res = ((((($dictMonad_0)['Bind1'])(null))['bind'])((($f_1)($b_2))(($v_3)->{'value0'})))(function($b_prime_5 = null) use ($__local_var_4_0, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_foldM'])($dictMonad_0))($f_1))($b_prime_5))($__local_var_4_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_findLastIndex
$GLOBALS['Data_List_NonEmpty_findLastIndex'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = (($GLOBALS['Data_List_findLastIndex'])($f_0))(($v_1)->{'value1'});
  $__t1 = null;;
  if ((is_object($v1_2_0) && (($v1_2_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ((($GLOBALS['Data_Semiring_semiringInt'])['add'])(($v1_2_0)->{'value0'}))(1));
goto end_branch_1;;
};
  if ((is_object($v1_2_0) && (($v1_2_0)->{'tag'} === "Nothing"))) {
$__t2 = null;;
if (($f_0)(($v_1)->{'value0'})) {
$__t2 = new Phpurs_Data1("Just", 0);
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("Nothing");
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_findIndex
$GLOBALS['Data_List_NonEmpty_findIndex'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($f_0)(($v_1)->{'value0'})) {
$__t0 = new Phpurs_Data1("Just", 0);
goto end_branch_0;;
};
  $__t0 = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($v1_2 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semiring_semiringInt'])['add'])($v1_2))(1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Data_List_findIndex'])($f_0))(($v_1)->{'value1'}));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_filterM
$GLOBALS['Data_List_NonEmpty_filterM'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))(($GLOBALS['Data_List_filterM'])($dictMonad_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_filter
$GLOBALS['Data_List_NonEmpty_filter'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_filter']);

// Data_List_NonEmpty_elemLastIndex
$GLOBALS['Data_List_NonEmpty_elemLastIndex'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_findLastIndex'])(function($v_2 = null) use ($dictEq_0, $x_1) {
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

// Data_List_NonEmpty_elemIndex
$GLOBALS['Data_List_NonEmpty_elemIndex'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_findIndex'])(function($v_2 = null) use ($dictEq_0, $x_1) {
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

// Data_List_NonEmpty_dropWhile
$GLOBALS['Data_List_NonEmpty_dropWhile'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_dropWhile']);

// Data_List_NonEmpty_drop
$GLOBALS['Data_List_NonEmpty_drop'] = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_drop']);

// Data_List_NonEmpty_cons'
$GLOBALS['Data_List_NonEmpty_cons__prime__'] = (function() {
  $__fn = function($x_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", $x_0, $xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_cons
$GLOBALS['Data_List_NonEmpty_cons'] = (function() {
  $__fn = function($y_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", $y_0, new Phpurs_Data2("Cons", ($v_1)->{'value0'}, ($v_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_concatMap
$GLOBALS['Data_List_NonEmpty_concatMap'] = (function() {
  $__fn = function($b_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_Types_bindNonEmptyList'])['bind'])($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_NonEmpty_concat
$GLOBALS['Data_List_NonEmpty_concat'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_bindNonEmptyList'])['bind'])($v_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_catMaybes
$GLOBALS['Data_List_NonEmpty_catMaybes'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_catMaybes'])(new Phpurs_Data2("Cons", ($v_0)->{'value0'}, ($v_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_NonEmpty_appendFoldable
$GLOBALS['Data_List_NonEmpty_appendFoldable'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $fromFoldable1_1_0 = ((($dictFoldable_0)['foldr'])($GLOBALS['Data_List_Types_Cons']))(new Phpurs_Data0("Nil"));
  $__res = (function() use ($fromFoldable1_1_0) {
  $__fn = function($v_2 = null, $ys_3 = null) use ($fromFoldable1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("NonEmpty", ($v_2)->{'value0'}, ((($GLOBALS['Data_List_Types_semigroupList'])['append'])(($v_2)->{'value1'}))(($fromFoldable1_1_0)($ys_3)));
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

