<?php

namespace Data\List;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.List, Data.List.Internal, Data.List.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.List, Data.List.Internal, Data.List.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.Internal/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
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
require_once __DIR__ . '/../Data.Unit/index.php';
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
\PhpursThunks::$thunks['Data_List_tailRecM2'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $a_1 = null, $b_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Control_Monad_Rec_Class_monadRecMaybe'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_monadRecMaybe')))['tailRecM'])(function($o_3 = null) use ($f_0) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_any'] = function() { $v = (function() use (&$__fn) {
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
return ((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))['foldMap'])(["mempty" => false, "Semigroup0" => function($dollar__unused_1 = null) use ($semigroupDisj1_0_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Pattern'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_updateAt'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_updateAt_v_0 = $v_0;
  $__tco_var_Data_List_updateAt_v1_1 = $v1_1;
  $__tco_var_Data_List_updateAt_v2_2 = $v2_2;
  tco_loop_Data_List_updateAt:;
  $v_0 = $__tco_var_Data_List_updateAt_v_0;
  $v1_1 = $__tco_var_Data_List_updateAt_v1_1;
  $v2_2 = $__tco_var_Data_List_updateAt_v2_2;
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Cons"))) {
$__t3 = null;;
switch ($v_0) {
case 0:
$__t3 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", $v1_1, ($v2_2)->{'value1'}));
goto end_branch_3;;
break;
default:
;
break;
};
$__local_var_3_1 = (((($GLOBALS['Data_List_updateAt'] ?? \PhpursThunks::eval('Data_List_updateAt')))(($v_0 - 1)))($v1_1))(($v2_2)->{'value1'});
$__t2 = null;;
if ((is_object($__local_var_3_1) && (($__local_var_3_1)->{'tag'} === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", ($v2_2)->{'value0'}, ($__local_var_3_1)->{'value0'}));
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("Nothing");
end_branch_2:;
$__t3 = $__t2;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_unzip'] = function() { $v = (((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))['foldr'])(function($v_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($v_0)->{'value0'};
  $__local_var_2_1 = ($v_0)->{'value1'};
  $__res = function($v1_3 = null) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", new Phpurs_Data2("Cons", $__local_var_1_0, ($v1_3)->{'value0'}), new Phpurs_Data2("Cons", $__local_var_2_1, ($v1_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", new Phpurs_Data0("Nil"), new Phpurs_Data0("Nil"))); return $v; };
\PhpursThunks::$thunks['Data_List_uncons'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", ["head" => ($v_0)->{'value0'}, "tail" => ($v_0)->{'value1'}]);
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_toUnfoldable'] = function() { $v = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = (($dictUnfoldable_0)['unfoldr'])(function($xs_1 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($xs_1) && (($xs_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($xs_1) && (($xs_1)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($xs_1)->{'value0'}, ($xs_1)->{'value1'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_tail'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", ($v_0)->{'value1'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_stripPrefix'] = function() { $v = (function() {
  $__fn = function($dictEq_0 = null, $v_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_List_tailRecM2'] ?? \PhpursThunks::eval('Data_List_tailRecM2')))((function() use ($dictEq_0) {
  $__fn = function($prefix_3 = null, $input_4 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($input_4) && (($input_4)->{'tag'} === "Cons"))) {
$__t1 = null;;
if ((is_object($prefix_3) && (($prefix_3)->{'tag'} === "Cons"))) {
$__t2 = null;;
if (((($dictEq_0)['eq'])(($prefix_3)->{'value0'}))(($input_4)->{'value0'})) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data1("Loop", ["a" => ($prefix_3)->{'value1'}, "b" => ($input_4)->{'value1'}]));
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("Nothing");
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
if ((is_object($prefix_3) && (($prefix_3)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data1("Done", $input_4));
goto end_branch_1;;
};
$__t1 = new Phpurs_Data0("Nothing");
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ((is_object($prefix_3) && (($prefix_3)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data1("Done", $input_4));
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_span'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_span_v_0 = $v_0;
  $__tco_var_Data_List_span_v1_1 = $v1_1;
  tco_loop_Data_List_span:;
  $v_0 = $__tco_var_Data_List_span_v_0;
  $v1_1 = $__tco_var_Data_List_span_v1_1;
  $__t0 = null;;
  if (((is_object($v1_1) && (($v1_1)->{'tag'} === "Cons")) && ($v_0)(($v1_1)->{'value0'}))) {
$v2_2_1 = ((($GLOBALS['Data_List_span'] ?? \PhpursThunks::eval('Data_List_span')))($v_0))(($v1_1)->{'value1'});
$__t0 = ["init" => new Phpurs_Data2("Cons", ($v1_1)->{'value0'}, ($v2_2_1)['init']), "rest" => ($v2_2_1)['rest']];
goto end_branch_0;;
};
  $__t0 = ["init" => new Phpurs_Data0("Nil"), "rest" => $v1_1];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_snoc'] = function() { $v = (function() {
  $__fn = function($xs_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))['foldr'])(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))(new Phpurs_Data2("Cons", $x_1, new Phpurs_Data0("Nil"))))($xs_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_singleton'] = function() { $v = function($a_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_0, new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_sortBy'] = function() { $v = function($cmp_0 = null) {
  $__num = \func_num_args();
  $merge_1_0 = null;
  $merge_1_0 = (function() use ($cmp_0, &$merge_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($cmp_0, &$merge_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$__t2 = null;;
if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons"))) {
$__t3 = null;;
if ((is_object((($cmp_0)(($v_2)->{'value0'}))(($v1_3)->{'value0'})) && (((($cmp_0)(($v_2)->{'value0'}))(($v1_3)->{'value0'}))->{'tag'} === "GT"))) {
$__t3 = new Phpurs_Data2("Cons", ($v1_3)->{'value0'}, (($merge_1_0)($v_2))(($v1_3)->{'value1'}));
goto end_branch_3;;
};
$__t3 = new Phpurs_Data2("Cons", ($v_2)->{'value0'}, (($merge_1_0)(($v_2)->{'value1'}))($v1_3));
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$__t2 = $v_2;
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t1 = $v1_3;
goto end_branch_1;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$__t1 = $v_2;
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
  $mergePairs_2_4 = null;
  $mergePairs_2_4 = function($v_3 = null) use (&$mergePairs_2_4, &$merge_1_0) {
  $__num = \func_num_args();
  $__t5 = null;;
  if (((is_object($v_3) && (($v_3)->{'tag'} === "Cons")) && (is_object(($v_3)->{'value1'}) && ((($v_3)->{'value1'})->{'tag'} === "Cons")))) {
$__t5 = new Phpurs_Data2("Cons", (($merge_1_0)(($v_3)->{'value0'}))((($v_3)->{'value1'})->{'value0'}), ($mergePairs_2_4)((($v_3)->{'value1'})->{'value1'}));
goto end_branch_5;;
};
  $__t5 = $v_3;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $mergeAll_3_6 = null;
  $mergeAll_3_6 = function($v_4 = null) use (&$mergeAll_3_6, &$mergePairs_2_4) {
  $__num = \func_num_args();
  $__tco_var_mergeAll_3_6_6_v_4 = $v_4;
  tco_loop_mergeAll_3_6_6:;
  $v_4 = $__tco_var_mergeAll_3_6_6_v_4;
  $__t7 = null;;
  if (((is_object($v_4) && (($v_4)->{'tag'} === "Cons")) && (is_object(($v_4)->{'value1'}) && ((($v_4)->{'value1'})->{'tag'} === "Nil")))) {
$__t7 = ($v_4)->{'value0'};
goto end_branch_7;;
};
  $__tco_6 = ($mergePairs_2_4)($v_4);
  $__tco_var_mergeAll_3_6_6_v_4 = $__tco_6;
  goto tco_loop_mergeAll_3_6_6;;
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $sequences_4_7 = null;
  $descending_4_7 = null;
  $ascending_4_7 = null;
  $sequences_4_7 = function($v_5 = null) use (&$ascending_4_7, $cmp_0, &$descending_4_7) {
  $__num = \func_num_args();
  $__t8 = null;;
  if (((is_object($v_5) && (($v_5)->{'tag'} === "Cons")) && (is_object(($v_5)->{'value1'}) && ((($v_5)->{'value1'})->{'tag'} === "Cons")))) {
$__t10 = null;;
if ((is_object((($cmp_0)(($v_5)->{'value0'}))((($v_5)->{'value1'})->{'value0'})) && (((($cmp_0)(($v_5)->{'value0'}))((($v_5)->{'value1'})->{'value0'}))->{'tag'} === "GT"))) {
$__t10 = ((($descending_4_7)((($v_5)->{'value1'})->{'value0'}))(new Phpurs_Data2("Cons", ($v_5)->{'value0'}, new Phpurs_Data0("Nil"))))((($v_5)->{'value1'})->{'value1'});
goto end_branch_10;;
};
$__local_var_6_9 = ($v_5)->{'value0'};
$__t10 = ((($ascending_4_7)((($v_5)->{'value1'})->{'value0'}))(function($v1_7 = null) use ($__local_var_6_9) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_6_9, $v1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_5)->{'value1'})->{'value1'});
end_branch_10:;
$__t8 = $__t10;
goto end_branch_8;;
};
  $__t8 = new Phpurs_Data2("Cons", $v_5, new Phpurs_Data0("Nil"));
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $descending_4_7 = (function() use ($cmp_0, &$descending_4_7, &$sequences_4_7) {
  $__fn = function($v_5 = null, $v1_6 = null, $v2_7 = null) use ($cmp_0, &$descending_4_7, &$sequences_4_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t11 = null;;
  if (((is_object($v2_7) && (($v2_7)->{'tag'} === "Cons")) && (is_object((($cmp_0)($v_5))(($v2_7)->{'value0'})) && (((($cmp_0)($v_5))(($v2_7)->{'value0'}))->{'tag'} === "GT")))) {
$__t11 = ((($descending_4_7)(($v2_7)->{'value0'}))(new Phpurs_Data2("Cons", $v_5, $v1_6)))(($v2_7)->{'value1'});
goto end_branch_11;;
};
  $__t11 = new Phpurs_Data2("Cons", new Phpurs_Data2("Cons", $v_5, $v1_6), ($sequences_4_7)($v2_7));
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $ascending_4_7 = (function() use (&$ascending_4_7, $cmp_0, &$sequences_4_7) {
  $__fn = function($v_5 = null, $v1_6 = null, $v2_7 = null) use (&$ascending_4_7, $cmp_0, &$sequences_4_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t12 = null;;
  if ((function() use ($cmp_0, $v2_7, $v_5, &$__fn) {
$__local_var_8_13 = (($cmp_0)($v_5))(($v2_7)->{'value0'});
return ((is_object($v2_7) && (($v2_7)->{'tag'} === "Cons")) && ((is_object($__local_var_8_13) && (($__local_var_8_13)->{'tag'} === "LT")) || ( ! (is_object($__local_var_8_13) && (($__local_var_8_13)->{'tag'} === "GT")))));
})()) {
$__t12 = ((($ascending_4_7)(($v2_7)->{'value0'}))(function($ys_8 = null) use ($v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ($v1_6)(new Phpurs_Data2("Cons", $v_5, $ys_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v2_7)->{'value1'});
goto end_branch_12;;
};
  $__t12 = new Phpurs_Data2("Cons", ($v1_6)(new Phpurs_Data2("Cons", $v_5, new Phpurs_Data0("Nil"))), ($sequences_4_7)($v2_7));
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($mergeAll_3_6))($sequences_4_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_sort'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = function($xs_2 = null) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_sortBy'] ?? \PhpursThunks::eval('Data_List_sortBy')))($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_showPattern'] = function() { $v = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = ["show" => function($v_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = (("(Pattern " . (((($GLOBALS['Data_List_Types_showList'] ?? \PhpursThunks::eval('Data_List_Types_showList')))($dictShow_0))['show'])($v_1)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_reverse'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_0_0_0_v_1 = $v_1;
  $__tco_var_go_0_0_0_v1_2 = $v1_2;
  tco_loop_go_0_0_0:;
  $v_1 = $__tco_var_go_0_0_0_v_1;
  $v1_2 = $__tco_var_go_0_0_0_v1_2;
  $__t0 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Nil"))) {
$__t0 = $v_1;
goto end_branch_0;;
};
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Cons"))) {
$__tco_1 = new Phpurs_Data2("Cons", ($v1_2)->{'value0'}, $v_1);
$__tco_2 = ($v1_2)->{'value1'};
$__tco_var_go_0_0_0_v_1 = $__tco_1;
$__tco_var_go_0_0_0_v1_2 = $__tco_2;
goto tco_loop_go_0_0_0;;
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
return ($go_0_0)(new Phpurs_Data0("Nil"));
})(); return $v; };
\PhpursThunks::$thunks['Data_List_take'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($v_1 = null, $v1_2 = null, $v2_3 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go_0_0_0_v_1 = $v_1;
  $__tco_var_go_0_0_0_v1_2 = $v1_2;
  $__tco_var_go_0_0_0_v2_3 = $v2_3;
  tco_loop_go_0_0_0:;
  $v_1 = $__tco_var_go_0_0_0_v_1;
  $v1_2 = $__tco_var_go_0_0_0_v1_2;
  $v2_3 = $__tco_var_go_0_0_0_v2_3;
  $__t0 = null;;
  if (($v1_2 < 1)) {
$go_4_1 = null;
$go_4_1 = (function() use (&$__tco_var_go_0_0_0_v_1, &$__tco_var_go_0_0_0_v1_2, &$__tco_var_go_0_0_0_v2_3, &$go_4_1) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$__tco_var_go_0_0_0_v_1, &$__tco_var_go_0_0_0_v1_2, &$__tco_var_go_0_0_0_v2_3, &$go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_1_1_v_5 = $v_5;
  $__tco_var_go_4_1_1_v1_6 = $v1_6;
  tco_loop_go_4_1_1:;
  $v_5 = $__tco_var_go_4_1_1_v_5;
  $v1_6 = $__tco_var_go_4_1_1_v1_6;
  $__t1 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go_4_1_1_v_5 = $__tco_2;
$__tco_var_go_4_1_1_v1_6 = $__tco_3;
goto tco_loop_go_4_1_1;;
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
$__t0 = (($go_4_1)(new Phpurs_Data0("Nil")))($v_1);
goto end_branch_0;;
};
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Nil"))) {
$go_4_2 = null;
$go_4_2 = (function() use (&$__tco_var_go_0_0_0_v_1, &$__tco_var_go_0_0_0_v1_2, &$__tco_var_go_0_0_0_v2_3, &$go_4_2) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$__tco_var_go_0_0_0_v_1, &$__tco_var_go_0_0_0_v1_2, &$__tco_var_go_0_0_0_v2_3, &$go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_2_2_v_5 = $v_5;
  $__tco_var_go_4_2_2_v1_6 = $v1_6;
  tco_loop_go_4_2_2:;
  $v_5 = $__tco_var_go_4_2_2_v_5;
  $v1_6 = $__tco_var_go_4_2_2_v1_6;
  $__t2 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t2 = $v_5;
goto end_branch_2;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_3 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_4 = ($v1_6)->{'value1'};
$__tco_var_go_4_2_2_v_5 = $__tco_3;
$__tco_var_go_4_2_2_v1_6 = $__tco_4;
goto tco_loop_go_4_2_2;;
$__t2 = null;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t0 = (($go_4_2)(new Phpurs_Data0("Nil")))($v_1);
goto end_branch_0;;
};
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Cons"))) {
$__tco_3 = new Phpurs_Data2("Cons", ($v2_3)->{'value0'}, $v_1);
$__tco_4 = ($v1_2 - 1);
$__tco_5 = ($v2_3)->{'value1'};
$__tco_var_go_0_0_0_v_1 = $__tco_3;
$__tco_var_go_0_0_0_v1_2 = $__tco_4;
$__tco_var_go_0_0_0_v2_3 = $__tco_5;
goto tco_loop_go_0_0_0;;
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
return ($go_0_0)(new Phpurs_Data0("Nil"));
})(); return $v; };
\PhpursThunks::$thunks['Data_List_takeWhile'] = function() { $v = function($p_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $p_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go_1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_0_0_v_2 = $v_2;
  $__tco_var_go_1_0_0_v1_3 = $v1_3;
  tco_loop_go_1_0_0:;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $v1_3 = $__tco_var_go_1_0_0_v1_3;
  $__t1 = null;;
  if (((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons")) && ($p_0)(($v1_3)->{'value0'}))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_3)->{'value0'}, $v_2);
$__tco_3 = ($v1_3)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_2;
$__tco_var_go_1_0_0_v1_3 = $__tco_3;
goto tco_loop_go_1_0_0;;
$__t1 = null;
goto end_branch_1;;
};
  $go_4_0 = null;
  $go_4_0 = (function() use (&$__tco_var_go_1_0_0_v_2, &$__tco_var_go_1_0_0_v1_3, &$go_4_0) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$__tco_var_go_1_0_0_v_2, &$__tco_var_go_1_0_0_v1_3, &$go_4_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_0_0_v_5 = $v_5;
  $__tco_var_go_4_0_0_v1_6 = $v1_6;
  tco_loop_go_4_0_0:;
  $v_5 = $__tco_var_go_4_0_0_v_5;
  $v1_6 = $__tco_var_go_4_0_0_v1_6;
  $__t0 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t0 = $v_5;
goto end_branch_0;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_1 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_2 = ($v1_6)->{'value1'};
$__tco_var_go_4_0_0_v_5 = $__tco_1;
$__tco_var_go_4_0_0_v1_6 = $__tco_2;
goto tco_loop_go_4_0_0;;
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
  $__t1 = (($go_4_0)(new Phpurs_Data0("Nil")))($v_2);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go_1_0)(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_unsnoc'] = function() { $v = function($lst_0 = null) {
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
if ((is_object(($v_2)->{'value1'}) && ((($v_2)->{'value1'})->{'tag'} === "Nil"))) {
$__t3 = new Phpurs_Data1("Just", ["revInit" => $v1_3, "last" => ($v_2)->{'value0'}]);
goto end_branch_3;;
};
$__tco_1 = ($v_2)->{'value1'};
$__tco_2 = new Phpurs_Data2("Cons", ($v_2)->{'value0'}, $v1_3);
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
  $__local_var_2_1 = (($go_1_0)($lst_0))(new Phpurs_Data0("Nil"));
  $__t2 = null;;
  if ((is_object($__local_var_2_1) && (($__local_var_2_1)->{'tag'} === "Just"))) {
$go_3_3 = null;
$go_3_3 = (function() use (&$go_3_3) {
  $__fn = function($v_4 = null, $v1_5 = null) use (&$go_3_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_3_3_3_v_4 = $v_4;
  $__tco_var_go_3_3_3_v1_5 = $v1_5;
  tco_loop_go_3_3_3:;
  $v_4 = $__tco_var_go_3_3_3_v_4;
  $v1_5 = $__tco_var_go_3_3_3_v1_5;
  $__t3 = null;;
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$__t3 = $v_4;
goto end_branch_3;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v1_5)->{'value0'}, $v_4);
$__tco_5 = ($v1_5)->{'value1'};
$__tco_var_go_3_3_3_v_4 = $__tco_4;
$__tco_var_go_3_3_3_v1_5 = $__tco_5;
goto tco_loop_go_3_3_3;;
$__t3 = null;
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
$__t2 = new Phpurs_Data1("Just", ["init" => (($go_3_3)(new Phpurs_Data0("Nil")))((($__local_var_2_1)->{'value0'})['revInit']), "last" => (($__local_var_2_1)->{'value0'})['last']]);
goto end_branch_2;;
};
  $__t2 = new Phpurs_Data0("Nothing");
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_zipWith'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go_3_0 = null;
  $go_3_0 = (function() use ($f_0, &$go_3_0) {
  $__fn = function($v_4 = null, $v1_5 = null, $v2_6 = null) use ($f_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go_3_0_0_v_4 = $v_4;
  $__tco_var_go_3_0_0_v1_5 = $v1_5;
  $__tco_var_go_3_0_0_v2_6 = $v2_6;
  tco_loop_go_3_0_0:;
  $v_4 = $__tco_var_go_3_0_0_v_4;
  $v1_5 = $__tco_var_go_3_0_0_v1_5;
  $v2_6 = $__tco_var_go_3_0_0_v2_6;
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
$__tco_var_go_3_0_0_v_4 = $__tco_1;
$__tco_var_go_3_0_0_v1_5 = $__tco_2;
$__tco_var_go_3_0_0_v2_6 = $__tco_3;
goto tco_loop_go_3_0_0;;
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
  $go_4_1 = null;
  $go_4_1 = (function() use (&$go_4_1) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_1_1_v_5 = $v_5;
  $__tco_var_go_4_1_1_v1_6 = $v1_6;
  tco_loop_go_4_1_1:;
  $v_5 = $__tco_var_go_4_1_1_v_5;
  $v1_6 = $__tco_var_go_4_1_1_v1_6;
  $__t1 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go_4_1_1_v_5 = $__tco_2;
$__tco_var_go_4_1_1_v1_6 = $__tco_3;
goto tco_loop_go_4_1_1;;
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
  $__res = (($go_4_1)(new Phpurs_Data0("Nil")))(((($go_3_0)($xs_1))($ys_2))(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_zip'] = function() { $v = (function() {
  $__fn = function($xs_0 = null, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null) use (&$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go_2_0_0_v_3 = $v_3;
  $__tco_var_go_2_0_0_v1_4 = $v1_4;
  $__tco_var_go_2_0_0_v2_5 = $v2_5;
  tco_loop_go_2_0_0:;
  $v_3 = $__tco_var_go_2_0_0_v_3;
  $v1_4 = $__tco_var_go_2_0_0_v1_4;
  $v2_5 = $__tco_var_go_2_0_0_v2_5;
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Nil"))) {
$__t0 = $v2_5;
goto end_branch_0;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Nil"))) {
$__t0 = $v2_5;
goto end_branch_0;;
};
  if (((is_object($v_3) && (($v_3)->{'tag'} === "Cons")) && (is_object($v1_4) && (($v1_4)->{'tag'} === "Cons")))) {
$__tco_1 = ($v_3)->{'value1'};
$__tco_2 = ($v1_4)->{'value1'};
$__tco_3 = new Phpurs_Data2("Cons", new Phpurs_Data2("Tuple", ($v_3)->{'value0'}, ($v1_4)->{'value0'}), $v2_5);
$__tco_var_go_2_0_0_v_3 = $__tco_1;
$__tco_var_go_2_0_0_v1_4 = $__tco_2;
$__tco_var_go_2_0_0_v2_5 = $__tco_3;
goto tco_loop_go_2_0_0;;
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
  $go_3_1 = null;
  $go_3_1 = (function() use (&$go_3_1) {
  $__fn = function($v_4 = null, $v1_5 = null) use (&$go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_3_1_1_v_4 = $v_4;
  $__tco_var_go_3_1_1_v1_5 = $v1_5;
  tco_loop_go_3_1_1:;
  $v_4 = $__tco_var_go_3_1_1_v_4;
  $v1_5 = $__tco_var_go_3_1_1_v1_5;
  $__t1 = null;;
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$__t1 = $v_4;
goto end_branch_1;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_5)->{'value0'}, $v_4);
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go_3_1_1_v_4 = $__tco_2;
$__tco_var_go_3_1_1_v1_5 = $__tco_3;
goto tco_loop_go_3_1_1;;
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
  $__res = (($go_3_1)(new Phpurs_Data0("Nil")))(((($go_2_0)($xs_0))($ys_1))(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_zipWithA'] = function() { $v = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $sequence1_1_0 = (((($GLOBALS['Data_List_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Types_traversableList')))['traverse'])($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity']);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2 = null, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go_5_1 = null;
  $go_5_1 = (function() use ($f_2, &$go_5_1) {
  $__fn = function($v_6 = null, $v1_7 = null, $v2_8 = null) use ($f_2, &$go_5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go_5_1_1_v_6 = $v_6;
  $__tco_var_go_5_1_1_v1_7 = $v1_7;
  $__tco_var_go_5_1_1_v2_8 = $v2_8;
  tco_loop_go_5_1_1:;
  $v_6 = $__tco_var_go_5_1_1_v_6;
  $v1_7 = $__tco_var_go_5_1_1_v1_7;
  $v2_8 = $__tco_var_go_5_1_1_v2_8;
  $__t1 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Nil"))) {
$__t1 = $v2_8;
goto end_branch_1;;
};
  if ((is_object($v1_7) && (($v1_7)->{'tag'} === "Nil"))) {
$__t1 = $v2_8;
goto end_branch_1;;
};
  if (((is_object($v_6) && (($v_6)->{'tag'} === "Cons")) && (is_object($v1_7) && (($v1_7)->{'tag'} === "Cons")))) {
$__tco_2 = ($v_6)->{'value1'};
$__tco_3 = ($v1_7)->{'value1'};
$__tco_4 = new Phpurs_Data2("Cons", (($f_2)(($v_6)->{'value0'}))(($v1_7)->{'value0'}), $v2_8);
$__tco_var_go_5_1_1_v_6 = $__tco_2;
$__tco_var_go_5_1_1_v1_7 = $__tco_3;
$__tco_var_go_5_1_1_v2_8 = $__tco_4;
goto tco_loop_go_5_1_1;;
$__t1 = null;
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
  $go_6_2 = null;
  $go_6_2 = (function() use (&$go_6_2) {
  $__fn = function($v_7 = null, $v1_8 = null) use (&$go_6_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_6_2_2_v_7 = $v_7;
  $__tco_var_go_6_2_2_v1_8 = $v1_8;
  tco_loop_go_6_2_2:;
  $v_7 = $__tco_var_go_6_2_2_v_7;
  $v1_8 = $__tco_var_go_6_2_2_v1_8;
  $__t2 = null;;
  if ((is_object($v1_8) && (($v1_8)->{'tag'} === "Nil"))) {
$__t2 = $v_7;
goto end_branch_2;;
};
  if ((is_object($v1_8) && (($v1_8)->{'tag'} === "Cons"))) {
$__tco_3 = new Phpurs_Data2("Cons", ($v1_8)->{'value0'}, $v_7);
$__tco_4 = ($v1_8)->{'value1'};
$__tco_var_go_6_2_2_v_7 = $__tco_3;
$__tco_var_go_6_2_2_v1_8 = $__tco_4;
goto tco_loop_go_6_2_2;;
$__t2 = null;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($sequence1_1_0)((($go_6_2)(new Phpurs_Data0("Nil")))(((($go_5_1)($xs_3))($ys_4))(new Phpurs_Data0("Nil"))));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_range'] = function() { $v = (function() {
  $__fn = function($start_0 = null, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if (($start_0 === $end_1)) {
$__t2 = new Phpurs_Data2("Cons", $start_0, new Phpurs_Data0("Nil"));
goto end_branch_2;;
};
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0) {
  $__fn = function($s_3 = null, $e_4 = null, $step_5 = null, $rest_6 = null) use (&$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_go_2_0_0_s_3 = $s_3;
  $__tco_var_go_2_0_0_e_4 = $e_4;
  $__tco_var_go_2_0_0_step_5 = $step_5;
  $__tco_var_go_2_0_0_rest_6 = $rest_6;
  tco_loop_go_2_0_0:;
  $s_3 = $__tco_var_go_2_0_0_s_3;
  $e_4 = $__tco_var_go_2_0_0_e_4;
  $step_5 = $__tco_var_go_2_0_0_step_5;
  $rest_6 = $__tco_var_go_2_0_0_rest_6;
  $__t4 = null;;
  if (($s_3 === $e_4)) {
$__t4 = new Phpurs_Data2("Cons", $s_3, $rest_6);
goto end_branch_4;;
};
  $__tco_0 = ($s_3 + $step_5);
  $__tco_1 = $e_4;
  $__tco_2 = $step_5;
  $__tco_3 = new Phpurs_Data2("Cons", $s_3, $rest_6);
  $__tco_var_go_2_0_0_s_3 = $__tco_0;
  $__tco_var_go_2_0_0_e_4 = $__tco_1;
  $__tco_var_go_2_0_0_step_5 = $__tco_2;
  $__tco_var_go_2_0_0_rest_6 = $__tco_3;
  goto tco_loop_go_2_0_0;;
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();
  $__t1 = null;;
  if (($start_0 > $end_1)) {
$__t1 = 1;
goto end_branch_1;;
};
  $__t1 = -1;
  end_branch_1:;
  $__t2 = (((($go_2_0)($end_1))($start_0))($__t1))(new Phpurs_Data0("Nil"));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_partition'] = function() { $v = (function() {
  $__fn = function($p_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))['foldr'])((function() use ($p_0) {
  $__fn = function($x_2 = null, $v_3 = null) use ($p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($p_0)($x_2)) {
$__t0 = ["no" => ($v_3)['no'], "yes" => new Phpurs_Data2("Cons", $x_2, ($v_3)['yes'])];
goto end_branch_0;;
};
  $__t0 = ["no" => new Phpurs_Data2("Cons", $x_2, ($v_3)['no']), "yes" => ($v_3)['yes']];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(["no" => new Phpurs_Data0("Nil"), "yes" => new Phpurs_Data0("Nil")]))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_null'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (is_object($v_0) && (($v_0)->{'tag'} === "Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_nubBy'] = function() { $v = function($p_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $p_0) {
  $__fn = function($v_2 = null, $v1_3 = null, $v2_4 = null) use (&$go_1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go_1_0_0_v_2 = $v_2;
  $__tco_var_go_1_0_0_v1_3 = $v1_3;
  $__tco_var_go_1_0_0_v2_4 = $v2_4;
  tco_loop_go_1_0_0:;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $v1_3 = $__tco_var_go_1_0_0_v1_3;
  $v2_4 = $__tco_var_go_1_0_0_v2_4;
  $__t0 = null;;
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Nil"))) {
$__t0 = $v1_3;
goto end_branch_0;;
};
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Cons"))) {
$v3_5_1 = (((($GLOBALS['Data_List_Internal_insertAndLookupBy'] ?? \PhpursThunks::eval('Data_List_Internal_insertAndLookupBy')))($p_0))(($v2_4)->{'value0'}))($v_2);
$__t5 = null;;
if (($v3_5_1)['found']) {
$__tco_6 = ($v3_5_1)['result'];
$__tco_7 = $v1_3;
$__tco_8 = ($v2_4)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_6;
$__tco_var_go_1_0_0_v1_3 = $__tco_7;
$__tco_var_go_1_0_0_v2_4 = $__tco_8;
goto tco_loop_go_1_0_0;;
$__t5 = null;
goto end_branch_5;;
};
$__tco_2 = ($v3_5_1)['result'];
$__tco_3 = new Phpurs_Data2("Cons", ($v2_4)->{'value0'}, $v1_3);
$__tco_4 = ($v2_4)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_2;
$__tco_var_go_1_0_0_v1_3 = $__tco_3;
$__tco_var_go_1_0_0_v2_4 = $__tco_4;
goto tco_loop_go_1_0_0;;
$__t5 = null;
end_branch_5:;
$__t0 = $__t5;
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
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_reverse'] ?? \PhpursThunks::eval('Data_List_reverse'))))((($go_1_0)(new Phpurs_Data0("Leaf")))(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_nub'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_nubBy'] ?? \PhpursThunks::eval('Data_List_nubBy')))(($dictOrd_0)['compare']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_newtypePattern'] = function() { $v = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_mapMaybe'] = function() { $v = function($f_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use ($f_0, &$go_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($f_0, &$go_1_0, &$__fn) {
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
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$go_4_1 = null;
$go_4_1 = (function() use (&$__tco_var_go_1_0_0_v_2, &$__tco_var_go_1_0_0_v1_3, &$go_4_1) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$__tco_var_go_1_0_0_v_2, &$__tco_var_go_1_0_0_v1_3, &$go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_1_1_v_5 = $v_5;
  $__tco_var_go_4_1_1_v1_6 = $v1_6;
  tco_loop_go_4_1_1:;
  $v_5 = $__tco_var_go_4_1_1_v_5;
  $v1_6 = $__tco_var_go_4_1_1_v1_6;
  $__t1 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go_4_1_1_v_5 = $__tco_2;
$__tco_var_go_4_1_1_v1_6 = $__tco_3;
goto tco_loop_go_4_1_1;;
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
$__t0 = (($go_4_1)(new Phpurs_Data0("Nil")))($v_2);
goto end_branch_0;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons"))) {
$v2_4_2 = ($f_0)(($v1_3)->{'value0'});
$__t3 = null;;
if ((is_object($v2_4_2) && (($v2_4_2)->{'tag'} === "Nothing"))) {
$__tco_4 = $v_2;
$__tco_5 = ($v1_3)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_4;
$__tco_var_go_1_0_0_v1_3 = $__tco_5;
goto tco_loop_go_1_0_0;;
$__t3 = null;
goto end_branch_3;;
};
if ((is_object($v2_4_2) && (($v2_4_2)->{'tag'} === "Just"))) {
$__tco_6 = new Phpurs_Data2("Cons", ($v2_4_2)->{'value0'}, $v_2);
$__tco_7 = ($v1_3)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_6;
$__tco_var_go_1_0_0_v1_3 = $__tco_7;
goto tco_loop_go_1_0_0;;
$__t3 = null;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
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
  $__res = ($go_1_0)(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_manyRec'] = function() { $v = (function() {
  $__fn = function($dictMonadRec_0 = null, $dictAlternative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Alt0_2_0 = (((($dictAlternative_1)['Plus1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Alt0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_3_1 = (($dictAlternative_1)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($p_4 = null) use ($Alt0_2_0, $__local_var_3_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)['tailRecM'])(function($acc_5 = null) use ($Alt0_2_0, $__local_var_3_1, $dictMonadRec_0, $p_4) {
  $__num = \func_num_args();
  $__res = ((((((($dictMonadRec_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['bind'])(((($Alt0_2_0)['alt'])(((((($Alt0_2_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Control_Monad_Rec_Class_Loop'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_Loop'))))($p_4)))((($__local_var_3_1)['pure'])(new Phpurs_Data1("Done", ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')))))))(function($aa_6 = null) use ($__local_var_3_1, $acc_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($aa_6) && (($aa_6)->{'tag'} === "Loop"))) {
$__t2 = new Phpurs_Data1("Loop", new Phpurs_Data2("Cons", ($aa_6)->{'value0'}, $acc_5));
goto end_branch_2;;
};
  if ((is_object($aa_6) && (($aa_6)->{'tag'} === "Done"))) {
$go_7_3 = null;
$go_7_3 = (function() use (&$go_7_3) {
  $__fn = function($v_8 = null, $v1_9 = null) use (&$go_7_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_7_3_3_v_8 = $v_8;
  $__tco_var_go_7_3_3_v1_9 = $v1_9;
  tco_loop_go_7_3_3:;
  $v_8 = $__tco_var_go_7_3_3_v_8;
  $v1_9 = $__tco_var_go_7_3_3_v1_9;
  $__t3 = null;;
  if ((is_object($v1_9) && (($v1_9)->{'tag'} === "Nil"))) {
$__t3 = $v_8;
goto end_branch_3;;
};
  if ((is_object($v1_9) && (($v1_9)->{'tag'} === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v1_9)->{'value0'}, $v_8);
$__tco_5 = ($v1_9)->{'value1'};
$__tco_var_go_7_3_3_v_8 = $__tco_4;
$__tco_var_go_7_3_3_v1_9 = $__tco_5;
goto tco_loop_go_7_3_3;;
$__t3 = null;
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
$__t2 = new Phpurs_Data1("Done", (($go_7_3)(new Phpurs_Data0("Nil")))($acc_5));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = (($__local_var_3_1)['pure'])($__t2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_someRec'] = function() { $v = (function() {
  $__fn = function($dictMonadRec_0 = null, $dictAlternative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $manyRec2_2_0 = ((($GLOBALS['Data_List_manyRec'] ?? \PhpursThunks::eval('Data_List_manyRec')))($dictMonadRec_0))($dictAlternative_1);
  $__res = function($v_3 = null) use ($dictAlternative_1, $manyRec2_2_0) {
  $__num = \func_num_args();
  $__res = ((((((($dictAlternative_1)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['apply'])(((((((((($dictAlternative_1)['Plus1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Alt0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))($v_3)))(($manyRec2_2_0)($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_some'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0 = null, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['apply'])(((((((((($dictAlternative_0)['Plus1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Alt0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))($v_2)))((($dictLazy_1)['defer'])(function($v1_3 = null) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_many'] ?? \PhpursThunks::eval('Data_List_many')))($dictAlternative_0))($dictLazy_1))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_many'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0 = null, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)['Plus1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Alt0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['alt'])((((($GLOBALS['Data_List_some'] ?? \PhpursThunks::eval('Data_List_some')))($dictAlternative_0))($dictLazy_1))($v_2)))((((($dictAlternative_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['pure'])(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_length'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($b_1 = null, $v_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_0_0_0_b_1 = $b_1;
  $__tco_var_go_0_0_0_v_2 = $v_2;
  tco_loop_go_0_0_0:;
  $b_1 = $__tco_var_go_0_0_0_b_1;
  $v_2 = $__tco_var_go_0_0_0_v_2;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nil"))) {
$__t0 = $b_1;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Cons"))) {
$__tco_1 = ($b_1 + 1);
$__tco_2 = ($v_2)->{'value1'};
$__tco_var_go_0_0_0_b_1 = $__tco_1;
$__tco_var_go_0_0_0_v_2 = $__tco_2;
goto tco_loop_go_0_0_0;;
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
return ($go_0_0)(0);
})(); return $v; };
\PhpursThunks::$thunks['Data_List_last'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_last_v_0 = $v_0;
  tco_loop_Data_List_last:;
  $v_0 = $__tco_var_Data_List_last_v_0;
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__t2 = null;;
if ((is_object(($v_0)->{'value1'}) && ((($v_0)->{'value1'})->{'tag'} === "Nil"))) {
$__t2 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
goto end_branch_2;;
};
$__tco_1 = ($v_0)->{'value1'};
$__tco_var_Data_List_last_v_0 = $__tco_1;
goto tco_loop_Data_List_last;;
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
}; return $v; };
\PhpursThunks::$thunks['Data_List_insertBy'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_insertBy_v_0 = $v_0;
  $__tco_var_Data_List_insertBy_v1_1 = $v1_1;
  $__tco_var_Data_List_insertBy_v2_2 = $v2_2;
  tco_loop_Data_List_insertBy:;
  $v_0 = $__tco_var_Data_List_insertBy_v_0;
  $v1_1 = $__tco_var_Data_List_insertBy_v1_1;
  $v2_2 = $__tco_var_Data_List_insertBy_v2_2;
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data2("Cons", $v1_1, new Phpurs_Data0("Nil"));
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Cons"))) {
$__t1 = null;;
if ((is_object((($v_0)($v1_1))(($v2_2)->{'value0'})) && (((($v_0)($v1_1))(($v2_2)->{'value0'}))->{'tag'} === "GT"))) {
$__t1 = new Phpurs_Data2("Cons", ($v2_2)->{'value0'}, (((($GLOBALS['Data_List_insertBy'] ?? \PhpursThunks::eval('Data_List_insertBy')))($v_0))($v1_1))(($v2_2)->{'value1'}));
goto end_branch_1;;
};
$__t1 = new Phpurs_Data2("Cons", $v1_1, $v2_2);
end_branch_1:;
$__t0 = $__t1;
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_insertAt'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_insertAt_v_0 = $v_0;
  $__tco_var_Data_List_insertAt_v1_1 = $v1_1;
  $__tco_var_Data_List_insertAt_v2_2 = $v2_2;
  tco_loop_Data_List_insertAt:;
  $v_0 = $__tco_var_Data_List_insertAt_v_0;
  $v1_1 = $__tco_var_Data_List_insertAt_v1_1;
  $v2_2 = $__tco_var_Data_List_insertAt_v2_2;
  $__t0 = null;;
  switch ($v_0) {
case 0:
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", $v1_1, $v2_2));
goto end_branch_0;;
break;
default:
;
break;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Cons"))) {
$__local_var_3_1 = (((($GLOBALS['Data_List_insertAt'] ?? \PhpursThunks::eval('Data_List_insertAt')))(($v_0 - 1)))($v1_1))(($v2_2)->{'value1'});
$__t2 = null;;
if ((is_object($__local_var_3_1) && (($__local_var_3_1)->{'tag'} === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", ($v2_2)->{'value0'}, ($__local_var_3_1)->{'value0'}));
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("Nothing");
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_insert'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_insertBy'] ?? \PhpursThunks::eval('Data_List_insertBy')))(($dictOrd_0)['compare']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_init'] = function() { $v = function($lst_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Data_List_unsnoc'] ?? \PhpursThunks::eval('Data_List_unsnoc')))($lst_0);
  $__t1 = null;;
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($__local_var_1_0)->{'value0'})['init']);
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_index'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_index_v_0 = $v_0;
  $__tco_var_Data_List_index_v1_1 = $v1_1;
  tco_loop_Data_List_index:;
  $v_0 = $__tco_var_Data_List_index_v_0;
  $v1_1 = $__tco_var_Data_List_index_v1_1;
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__t3 = null;;
switch ($v1_1) {
case 0:
$__t3 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
goto end_branch_3;;
break;
default:
;
break;
};
$__tco_1 = ($v_0)->{'value1'};
$__tco_2 = ($v1_1 - 1);
$__tco_var_Data_List_index_v_0 = $__tco_1;
$__tco_var_Data_List_index_v1_1 = $__tco_2;
goto tco_loop_Data_List_index;;
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_head'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_transpose'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_transpose_v_0 = $v_0;
  tco_loop_Data_List_transpose:;
  $v_0 = $__tco_var_Data_List_transpose_v_0;
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Cons"))) {
$__t1 = null;;
if ((is_object(($v_0)->{'value0'}) && ((($v_0)->{'value0'})->{'tag'} === "Nil"))) {
$__tco_2 = ($v_0)->{'value1'};
$__tco_var_Data_List_transpose_v_0 = $__tco_2;
goto tco_loop_Data_List_transpose;;
$__t1 = null;
goto end_branch_1;;
};
if ((is_object(($v_0)->{'value0'}) && ((($v_0)->{'value0'})->{'tag'} === "Cons"))) {
$go_1_3 = null;
$go_1_3 = (function() use (&$go_1_3) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go_1_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_3_3_v_2 = $v_2;
  $__tco_var_go_1_3_3_v1_3 = $v1_3;
  tco_loop_go_1_3_3:;
  $v_2 = $__tco_var_go_1_3_3_v_2;
  $v1_3 = $__tco_var_go_1_3_3_v1_3;
  $__t3 = null;;
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$go_4_4 = null;
$go_4_4 = (function() use (&$__tco_var_go_1_3_3_v_2, &$__tco_var_go_1_3_3_v1_3, &$go_4_4) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$__tco_var_go_1_3_3_v_2, &$__tco_var_go_1_3_3_v1_3, &$go_4_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_4_4_v_5 = $v_5;
  $__tco_var_go_4_4_4_v1_6 = $v1_6;
  tco_loop_go_4_4_4:;
  $v_5 = $__tco_var_go_4_4_4_v_5;
  $v1_6 = $__tco_var_go_4_4_4_v1_6;
  $__t4 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t4 = $v_5;
goto end_branch_4;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_5 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_6 = ($v1_6)->{'value1'};
$__tco_var_go_4_4_4_v_5 = $__tco_5;
$__tco_var_go_4_4_4_v1_6 = $__tco_6;
goto tco_loop_go_4_4_4;;
$__t4 = null;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t3 = (($go_4_4)(new Phpurs_Data0("Nil")))($v_2);
goto end_branch_3;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons"))) {
$__t5 = null;;
if ((is_object(($v1_3)->{'value0'}) && ((($v1_3)->{'value0'})->{'tag'} === "Nil"))) {
$__tco_6 = $v_2;
$__tco_7 = ($v1_3)->{'value1'};
$__tco_var_go_1_3_3_v_2 = $__tco_6;
$__tco_var_go_1_3_3_v1_3 = $__tco_7;
goto tco_loop_go_1_3_3;;
$__t5 = null;
goto end_branch_5;;
};
if ((is_object(($v1_3)->{'value0'}) && ((($v1_3)->{'value0'})->{'tag'} === "Cons"))) {
$__tco_8 = new Phpurs_Data2("Cons", (($v1_3)->{'value0'})->{'value0'}, $v_2);
$__tco_9 = ($v1_3)->{'value1'};
$__tco_var_go_1_3_3_v_2 = $__tco_8;
$__tco_var_go_1_3_3_v1_3 = $__tco_9;
goto tco_loop_go_1_3_3;;
$__t5 = null;
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t3 = $__t5;
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
$go_1_4 = null;
$go_1_4 = (function() use (&$go_1_4) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go_1_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_4_4_v_2 = $v_2;
  $__tco_var_go_1_4_4_v1_3 = $v1_3;
  tco_loop_go_1_4_4:;
  $v_2 = $__tco_var_go_1_4_4_v_2;
  $v1_3 = $__tco_var_go_1_4_4_v1_3;
  $__t4 = null;;
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$go_4_5 = null;
$go_4_5 = (function() use (&$__tco_var_go_1_4_4_v_2, &$__tco_var_go_1_4_4_v1_3, &$go_4_5) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$__tco_var_go_1_4_4_v_2, &$__tco_var_go_1_4_4_v1_3, &$go_4_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_5_5_v_5 = $v_5;
  $__tco_var_go_4_5_5_v1_6 = $v1_6;
  tco_loop_go_4_5_5:;
  $v_5 = $__tco_var_go_4_5_5_v_5;
  $v1_6 = $__tco_var_go_4_5_5_v1_6;
  $__t5 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t5 = $v_5;
goto end_branch_5;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_6 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_7 = ($v1_6)->{'value1'};
$__tco_var_go_4_5_5_v_5 = $__tco_6;
$__tco_var_go_4_5_5_v1_6 = $__tco_7;
goto tco_loop_go_4_5_5;;
$__t5 = null;
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t4 = (($go_4_5)(new Phpurs_Data0("Nil")))($v_2);
goto end_branch_4;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons"))) {
$__t6 = null;;
if ((is_object(($v1_3)->{'value0'}) && ((($v1_3)->{'value0'})->{'tag'} === "Nil"))) {
$__tco_7 = $v_2;
$__tco_8 = ($v1_3)->{'value1'};
$__tco_var_go_1_4_4_v_2 = $__tco_7;
$__tco_var_go_1_4_4_v1_3 = $__tco_8;
goto tco_loop_go_1_4_4;;
$__t6 = null;
goto end_branch_6;;
};
if ((is_object(($v1_3)->{'value0'}) && ((($v1_3)->{'value0'})->{'tag'} === "Cons"))) {
$__tco_9 = new Phpurs_Data2("Cons", (($v1_3)->{'value0'})->{'value1'}, $v_2);
$__tco_10 = ($v1_3)->{'value1'};
$__tco_var_go_1_4_4_v_2 = $__tco_9;
$__tco_var_go_1_4_4_v1_3 = $__tco_10;
goto tco_loop_go_1_4_4;;
$__t6 = null;
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t4 = $__t6;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = new Phpurs_Data2("Cons", new Phpurs_Data2("Cons", (($v_0)->{'value0'})->{'value0'}, (($go_1_3)(new Phpurs_Data0("Nil")))(($v_0)->{'value1'})), (($GLOBALS['Data_List_transpose'] ?? \PhpursThunks::eval('Data_List_transpose')))(new Phpurs_Data2("Cons", (($v_0)->{'value0'})->{'value1'}, (($go_1_4)(new Phpurs_Data0("Nil")))(($v_0)->{'value1'}))));
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_groupBy'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_groupBy_v_0 = $v_0;
  $__tco_var_Data_List_groupBy_v1_1 = $v1_1;
  tco_loop_Data_List_groupBy:;
  $v_0 = $__tco_var_Data_List_groupBy_v_0;
  $v1_1 = $__tco_var_Data_List_groupBy_v1_1;
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Cons"))) {
$v2_2_1 = ((($GLOBALS['Data_List_span'] ?? \PhpursThunks::eval('Data_List_span')))(($v_0)(($v1_1)->{'value0'})))(($v1_1)->{'value1'});
$__t0 = new Phpurs_Data2("Cons", new Phpurs_Data2("NonEmpty", ($v1_1)->{'value0'}, ($v2_2_1)['init']), ((($GLOBALS['Data_List_groupBy'] ?? \PhpursThunks::eval('Data_List_groupBy')))($v_0))(($v2_2_1)['rest']));
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_groupAllBy'] = function() { $v = function($p_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_List_groupBy'] ?? \PhpursThunks::eval('Data_List_groupBy')))((function() use ($p_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (is_object((($p_0)($x_1))($y_2)) && (((($p_0)($x_1))($y_2))->{'tag'} === "EQ"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))((($GLOBALS['Data_List_sortBy'] ?? \PhpursThunks::eval('Data_List_sortBy')))($p_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_group'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_groupBy'] ?? \PhpursThunks::eval('Data_List_groupBy')))(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_groupAll'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_List_group'] ?? \PhpursThunks::eval('Data_List_group')))((($dictOrd_0)['Eq0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))))(function($xs_2 = null) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_sortBy'] ?? \PhpursThunks::eval('Data_List_sortBy')))($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_fromFoldable'] = function() { $v = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_0)['foldr'])(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_foldM'] = function() { $v = (function() {
  $__fn = function($dictMonad_0 = null, $v_1 = null, $v1_2 = null, $v2_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_List_foldM_dictMonad_0 = $dictMonad_0;
  $__tco_var_Data_List_foldM_v_1 = $v_1;
  $__tco_var_Data_List_foldM_v1_2 = $v1_2;
  $__tco_var_Data_List_foldM_v2_3 = $v2_3;
  tco_loop_Data_List_foldM:;
  $dictMonad_0 = $__tco_var_Data_List_foldM_dictMonad_0;
  $v_1 = $__tco_var_Data_List_foldM_v_1;
  $v1_2 = $__tco_var_Data_List_foldM_v1_2;
  $v2_3 = $__tco_var_Data_List_foldM_v2_3;
  $__t0 = null;;
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Nil"))) {
$__t0 = (((($dictMonad_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['pure'])($v1_2);
goto end_branch_0;;
};
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Cons"))) {
$__local_var_4_1 = ($v2_3)->{'value1'};
$__t0 = ((((($dictMonad_0)['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['bind'])((($v_1)($v1_2))(($v2_3)->{'value0'})))(function($b__prime___5 = null) use ($__local_var_4_1, $dictMonad_0, $v_1) {
  $__num = \func_num_args();
  $__res = ((((($GLOBALS['Data_List_foldM'] ?? \PhpursThunks::eval('Data_List_foldM')))($dictMonad_0))($v_1))($b__prime___5))($__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_findIndex'] = function() { $v = function($fn_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use ($fn_0, &$go_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($fn_0, &$go_1_0, &$__fn) {
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
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons"))) {
$__t3 = null;;
if (($fn_0)(($v1_3)->{'value0'})) {
$__t3 = new Phpurs_Data1("Just", $v_2);
goto end_branch_3;;
};
$__tco_1 = ($v_2 + 1);
$__tco_2 = ($v1_3)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_1;
$__tco_var_go_1_0_0_v1_3 = $__tco_2;
goto tco_loop_go_1_0_0;;
$__t3 = null;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
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
  $__res = ($go_1_0)(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_findLastIndex'] = function() { $v = (function() {
  $__fn = function($fn_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_0 = null;
  $go_2_0 = (function() use ($fn_0, &$go_2_0) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($fn_0, &$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_2_0_0_v_3 = $v_3;
  $__tco_var_go_2_0_0_v1_4 = $v1_4;
  tco_loop_go_2_0_0:;
  $v_3 = $__tco_var_go_2_0_0_v_3;
  $v1_4 = $__tco_var_go_2_0_0_v1_4;
  $__t0 = null;;
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Cons"))) {
$__t3 = null;;
if (($fn_0)(($v1_4)->{'value0'})) {
$__t3 = new Phpurs_Data1("Just", $v_3);
goto end_branch_3;;
};
$__tco_1 = ($v_3 + 1);
$__tco_2 = ($v1_4)->{'value1'};
$__tco_var_go_2_0_0_v_3 = $__tco_1;
$__tco_var_go_2_0_0_v1_4 = $__tco_2;
goto tco_loop_go_2_0_0;;
$__t3 = null;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
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
  $go_3_1 = null;
  $go_3_1 = (function() use (&$go_3_1) {
  $__fn = function($v_4 = null, $v1_5 = null) use (&$go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_3_1_1_v_4 = $v_4;
  $__tco_var_go_3_1_1_v1_5 = $v1_5;
  tco_loop_go_3_1_1:;
  $v_4 = $__tco_var_go_3_1_1_v_4;
  $v1_5 = $__tco_var_go_3_1_1_v1_5;
  $__t1 = null;;
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$__t1 = $v_4;
goto end_branch_1;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_5)->{'value0'}, $v_4);
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go_3_1_1_v_4 = $__tco_2;
$__tco_var_go_3_1_1_v1_5 = $__tco_3;
goto tco_loop_go_3_1_1;;
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
  $__local_var_3_1 = (($go_2_0)(0))((($go_3_1)(new Phpurs_Data0("Nil")))($xs_1));
  $__t3 = null;;
  if ((is_object($__local_var_3_1) && (($__local_var_3_1)->{'tag'} === "Just"))) {
$go_4_4 = null;
$go_4_4 = (function() use (&$go_4_4) {
  $__fn = function($b_5 = null, $v_6 = null) use (&$go_4_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_4_4_b_5 = $b_5;
  $__tco_var_go_4_4_4_v_6 = $v_6;
  tco_loop_go_4_4_4:;
  $b_5 = $__tco_var_go_4_4_4_b_5;
  $v_6 = $__tco_var_go_4_4_4_v_6;
  $__t4 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Nil"))) {
$__t4 = $b_5;
goto end_branch_4;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Cons"))) {
$__tco_5 = ($b_5 + 1);
$__tco_6 = ($v_6)->{'value1'};
$__tco_var_go_4_4_4_b_5 = $__tco_5;
$__tco_var_go_4_4_4_v_6 = $__tco_6;
goto tco_loop_go_4_4_4;;
$__t4 = null;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t3 = new Phpurs_Data1("Just", (((($go_4_4)(0))($xs_1) - 1) - ($__local_var_3_1)->{'value0'}));
goto end_branch_3;;
};
  $__t3 = new Phpurs_Data0("Nothing");
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_filterM'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_List_filterM_dictMonad_0 = $dictMonad_0;
  tco_loop_Data_List_filterM:;
  $dictMonad_0 = $__tco_var_Data_List_filterM_dictMonad_0;
  $__local_var_1_0 = (($dictMonad_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (($dictMonad_0)['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Nil"))) {
$__t2 = (($__local_var_1_0)['pure'])(new Phpurs_Data0("Nil"));
goto end_branch_2;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Cons"))) {
$__local_var_5_3 = ($v1_4)->{'value0'};
$__local_var_6_4 = ($v1_4)->{'value1'};
$__t2 = ((($__local_var_2_1)['bind'])(($v_3)($__local_var_5_3)))(function($b_7 = null) use ($__local_var_1_0, $__local_var_2_1, $__local_var_5_3, $__local_var_6_4, $dictMonad_0, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)['bind'])((((($GLOBALS['Data_List_filterM'] ?? \PhpursThunks::eval('Data_List_filterM')))($dictMonad_0))($v_3))($__local_var_6_4)))(function($xs__prime___8 = null) use ($__local_var_1_0, $__local_var_5_3, $b_7) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($b_7) {
$__t5 = new Phpurs_Data2("Cons", $__local_var_5_3, $xs__prime___8);
goto end_branch_5;;
};
  $__t5 = $xs__prime___8;
  end_branch_5:;
  $__res = (($__local_var_1_0)['pure'])($__t5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
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
}; return $v; };
\PhpursThunks::$thunks['Data_List_filter'] = function() { $v = function($p_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $p_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use (&$go_1_0, $p_0, &$__fn) {
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
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Nil"))) {
$go_4_1 = null;
$go_4_1 = (function() use (&$__tco_var_go_1_0_0_v_2, &$__tco_var_go_1_0_0_v1_3, &$go_4_1) {
  $__fn = function($v_5 = null, $v1_6 = null) use (&$__tco_var_go_1_0_0_v_2, &$__tco_var_go_1_0_0_v1_3, &$go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_4_1_1_v_5 = $v_5;
  $__tco_var_go_4_1_1_v1_6 = $v1_6;
  tco_loop_go_4_1_1:;
  $v_5 = $__tco_var_go_4_1_1_v_5;
  $v1_6 = $__tco_var_go_4_1_1_v1_6;
  $__t1 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil"))) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go_4_1_1_v_5 = $__tco_2;
$__tco_var_go_4_1_1_v1_6 = $__tco_3;
goto tco_loop_go_4_1_1;;
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
$__t0 = (($go_4_1)(new Phpurs_Data0("Nil")))($v_2);
goto end_branch_0;;
};
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Cons"))) {
$__t4 = null;;
if (($p_0)(($v1_3)->{'value0'})) {
$__tco_5 = new Phpurs_Data2("Cons", ($v1_3)->{'value0'}, $v_2);
$__tco_6 = ($v1_3)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_5;
$__tco_var_go_1_0_0_v1_3 = $__tco_6;
goto tco_loop_go_1_0_0;;
$__t4 = null;
goto end_branch_4;;
};
$__tco_2 = $v_2;
$__tco_3 = ($v1_3)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_2;
$__tco_var_go_1_0_0_v1_3 = $__tco_3;
goto tco_loop_go_1_0_0;;
$__t4 = null;
end_branch_4:;
$__t0 = $__t4;
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
  $__res = ($go_1_0)(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_intersectBy'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  $go_3_0 = null;
  $go_3_0 = (function() use (&$go_3_0, $v2_2, $v_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use (&$go_3_0, $v2_2, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_3_0_0_v_4 = $v_4;
  $__tco_var_go_3_0_0_v1_5 = $v1_5;
  tco_loop_go_3_0_0:;
  $v_4 = $__tco_var_go_3_0_0_v_4;
  $v1_5 = $__tco_var_go_3_0_0_v1_5;
  $__t0 = null;;
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$go_6_1 = null;
$go_6_1 = (function() use (&$__tco_var_go_3_0_0_v_4, &$__tco_var_go_3_0_0_v1_5, &$go_6_1) {
  $__fn = function($v_7 = null, $v1_8 = null) use (&$__tco_var_go_3_0_0_v_4, &$__tco_var_go_3_0_0_v1_5, &$go_6_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_6_1_1_v_7 = $v_7;
  $__tco_var_go_6_1_1_v1_8 = $v1_8;
  tco_loop_go_6_1_1:;
  $v_7 = $__tco_var_go_6_1_1_v_7;
  $v1_8 = $__tco_var_go_6_1_1_v1_8;
  $__t1 = null;;
  if ((is_object($v1_8) && (($v1_8)->{'tag'} === "Nil"))) {
$__t1 = $v_7;
goto end_branch_1;;
};
  if ((is_object($v1_8) && (($v1_8)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_8)->{'value0'}, $v_7);
$__tco_3 = ($v1_8)->{'value1'};
$__tco_var_go_6_1_1_v_7 = $__tco_2;
$__tco_var_go_6_1_1_v1_8 = $__tco_3;
goto tco_loop_go_6_1_1;;
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
$__t0 = (($go_6_1)(new Phpurs_Data0("Nil")))($v_4);
goto end_branch_0;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Cons"))) {
$__t4 = null;;
if (((($GLOBALS['Data_List_any'] ?? \PhpursThunks::eval('Data_List_any')))(($v_0)(($v1_5)->{'value0'})))($v2_2)) {
$__tco_5 = new Phpurs_Data2("Cons", ($v1_5)->{'value0'}, $v_4);
$__tco_6 = ($v1_5)->{'value1'};
$__tco_var_go_3_0_0_v_4 = $__tco_5;
$__tco_var_go_3_0_0_v1_5 = $__tco_6;
goto tco_loop_go_3_0_0;;
$__t4 = null;
goto end_branch_4;;
};
$__tco_2 = $v_4;
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go_3_0_0_v_4 = $__tco_2;
$__tco_var_go_3_0_0_v1_5 = $__tco_3;
goto tco_loop_go_3_0_0;;
$__t4 = null;
end_branch_4:;
$__t0 = $__t4;
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
  $__t1 = (($go_3_0)(new Phpurs_Data0("Nil")))($v1_1);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_intersect'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_intersectBy'] ?? \PhpursThunks::eval('Data_List_intersectBy')))(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_nubByEq'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_nubByEq_v_0 = $v_0;
  $__tco_var_Data_List_nubByEq_v1_1 = $v1_1;
  tco_loop_Data_List_nubByEq:;
  $v_0 = $__tco_var_Data_List_nubByEq_v_0;
  $v1_1 = $__tco_var_Data_List_nubByEq_v1_1;
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Cons"))) {
$__local_var_2_1 = ($v1_1)->{'value0'};
$go_3_2 = null;
$go_3_2 = (function() use ($__local_var_2_1, &$go_3_2, $v_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($__local_var_2_1, &$go_3_2, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_3_2_2_v_4 = $v_4;
  $__tco_var_go_3_2_2_v1_5 = $v1_5;
  tco_loop_go_3_2_2:;
  $v_4 = $__tco_var_go_3_2_2_v_4;
  $v1_5 = $__tco_var_go_3_2_2_v1_5;
  $__t2 = null;;
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$go_6_3 = null;
$go_6_3 = (function() use (&$__tco_var_go_3_2_2_v_4, &$__tco_var_go_3_2_2_v1_5, &$go_6_3) {
  $__fn = function($v_7 = null, $v1_8 = null) use (&$__tco_var_go_3_2_2_v_4, &$__tco_var_go_3_2_2_v1_5, &$go_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_6_3_3_v_7 = $v_7;
  $__tco_var_go_6_3_3_v1_8 = $v1_8;
  tco_loop_go_6_3_3:;
  $v_7 = $__tco_var_go_6_3_3_v_7;
  $v1_8 = $__tco_var_go_6_3_3_v1_8;
  $__t3 = null;;
  if ((is_object($v1_8) && (($v1_8)->{'tag'} === "Nil"))) {
$__t3 = $v_7;
goto end_branch_3;;
};
  if ((is_object($v1_8) && (($v1_8)->{'tag'} === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v1_8)->{'value0'}, $v_7);
$__tco_5 = ($v1_8)->{'value1'};
$__tco_var_go_6_3_3_v_7 = $__tco_4;
$__tco_var_go_6_3_3_v1_8 = $__tco_5;
goto tco_loop_go_6_3_3;;
$__t3 = null;
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
$__t2 = (($go_6_3)(new Phpurs_Data0("Nil")))($v_4);
goto end_branch_2;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Cons"))) {
$__t6 = null;;
if (( ! (($v_0)($__local_var_2_1))(($v1_5)->{'value0'}))) {
$__tco_7 = new Phpurs_Data2("Cons", ($v1_5)->{'value0'}, $v_4);
$__tco_8 = ($v1_5)->{'value1'};
$__tco_var_go_3_2_2_v_4 = $__tco_7;
$__tco_var_go_3_2_2_v1_5 = $__tco_8;
goto tco_loop_go_3_2_2;;
$__t6 = null;
goto end_branch_6;;
};
$__tco_4 = $v_4;
$__tco_5 = ($v1_5)->{'value1'};
$__tco_var_go_3_2_2_v_4 = $__tco_4;
$__tco_var_go_3_2_2_v1_5 = $__tco_5;
goto tco_loop_go_3_2_2;;
$__t6 = null;
end_branch_6:;
$__t2 = $__t6;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t0 = new Phpurs_Data2("Cons", $__local_var_2_1, ((($GLOBALS['Data_List_nubByEq'] ?? \PhpursThunks::eval('Data_List_nubByEq')))($v_0))((($go_3_2)(new Phpurs_Data0("Nil")))(($v1_1)->{'value1'})));
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_nubEq'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_nubByEq'] ?? \PhpursThunks::eval('Data_List_nubByEq')))(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_eqPattern'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_3_0 = null;
  $go_3_0 = (function() use ($dictEq_0, &$go_3_0) {
  $__fn = function($v_4 = null, $v1_5 = null, $v2_6 = null) use ($dictEq_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if (( ! $v2_6)) {
$__t1 = false;
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Nil"))) {
$__t1 = ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil")) && $v2_6);
goto end_branch_1;;
};
  $__t1 = ((is_object($v_4) && (($v_4)->{'tag'} === "Cons")) && ((is_object($v1_5) && (($v1_5)->{'tag'} === "Cons")) && ((($go_3_0)(($v_4)->{'value1'}))(($v1_5)->{'value1'}))(($v2_6 && ((($dictEq_0)['eq'])(($v1_5)->{'value0'}))(($v_4)->{'value0'})))));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($go_3_0)($x_1))($y_2))(true);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_ordPattern'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)['Eq0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eqPattern1_2_1 = ["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_4_1 = null;
  $go_4_1 = (function() use ($__local_var_1_0, &$go_4_1) {
  $__fn = function($v_5 = null, $v1_6 = null, $v2_7 = null) use ($__local_var_1_0, &$go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t2 = null;;
  if (( ! $v2_7)) {
$__t2 = false;
goto end_branch_2;;
};
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Nil"))) {
$__t2 = ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nil")) && $v2_7);
goto end_branch_2;;
};
  $__t2 = ((is_object($v_5) && (($v_5)->{'tag'} === "Cons")) && ((is_object($v1_6) && (($v1_6)->{'tag'} === "Cons")) && ((($go_4_1)(($v_5)->{'value1'}))(($v1_6)->{'value1'}))(($v2_7 && ((($__local_var_1_0)['eq'])(($v1_6)->{'value0'}))(($v_5)->{'value0'})))));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($go_4_1)($x_2))($y_3))(true);
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
  $__res = ((((($GLOBALS['Data_List_Types_ordList'] ?? \PhpursThunks::eval('Data_List_Types_ordList')))($dictOrd_0))['compare'])($x_3))($y_4);
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
}; return $v; };
\PhpursThunks::$thunks['Data_List_elemLastIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_List_findLastIndex'] ?? \PhpursThunks::eval('Data_List_findLastIndex')))(function($v_2 = null) use ($dictEq_0, $x_1) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_elemIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_0 = null;
  $go_2_0 = (function() use ($dictEq_0, &$go_2_0, $x_1) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($dictEq_0, &$go_2_0, $x_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_2_0_0_v_3 = $v_3;
  $__tco_var_go_2_0_0_v1_4 = $v1_4;
  tco_loop_go_2_0_0:;
  $v_3 = $__tco_var_go_2_0_0_v_3;
  $v1_4 = $__tco_var_go_2_0_0_v1_4;
  $__t0 = null;;
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Cons"))) {
$__t3 = null;;
if (((($dictEq_0)['eq'])(($v1_4)->{'value0'}))($x_1)) {
$__t3 = new Phpurs_Data1("Just", $v_3);
goto end_branch_3;;
};
$__tco_1 = ($v_3 + 1);
$__tco_2 = ($v1_4)->{'value1'};
$__tco_var_go_2_0_0_v_3 = $__tco_1;
$__tco_var_go_2_0_0_v1_4 = $__tco_2;
goto tco_loop_go_2_0_0;;
$__t3 = null;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
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
  $__res = ($go_2_0)(0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_dropWhile'] = function() { $v = function($p_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = function($v_2 = null) use (&$go_1_0, $p_0) {
  $__num = \func_num_args();
  $__tco_var_go_1_0_0_v_2 = $v_2;
  tco_loop_go_1_0_0:;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $__t0 = null;;
  if (((is_object($v_2) && (($v_2)->{'tag'} === "Cons")) && ($p_0)(($v_2)->{'value0'}))) {
$__tco_1 = ($v_2)->{'value1'};
$__tco_var_go_1_0_0_v_2 = $__tco_1;
goto tco_loop_go_1_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  $__t0 = $v_2;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_dropEnd'] = function() { $v = (function() {
  $__fn = function($n_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0) {
  $__fn = function($b_3 = null, $v_4 = null) use (&$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_2_0_0_b_3 = $b_3;
  $__tco_var_go_2_0_0_v_4 = $v_4;
  tco_loop_go_2_0_0:;
  $b_3 = $__tco_var_go_2_0_0_b_3;
  $v_4 = $__tco_var_go_2_0_0_v_4;
  $__t0 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Nil"))) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Cons"))) {
$__tco_1 = ($b_3 + 1);
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go_2_0_0_b_3 = $__tco_1;
$__tco_var_go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go_2_0_0;;
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
  $__res = ((($GLOBALS['Data_List_take'] ?? \PhpursThunks::eval('Data_List_take')))(((($go_2_0)(0))($xs_1) - $n_0)))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_drop'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_drop_v_0 = $v_0;
  $__tco_var_Data_List_drop_v1_1 = $v1_1;
  tco_loop_Data_List_drop:;
  $v_0 = $__tco_var_Data_List_drop_v_0;
  $v1_1 = $__tco_var_Data_List_drop_v1_1;
  $__t0 = null;;
  if (($v_0 < 1)) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Cons"))) {
$__tco_1 = ($v_0 - 1);
$__tco_2 = ($v1_1)->{'value1'};
$__tco_var_Data_List_drop_v_0 = $__tco_1;
$__tco_var_Data_List_drop_v1_1 = $__tco_2;
goto tco_loop_Data_List_drop;;
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_slice'] = function() { $v = (function() {
  $__fn = function($start_0 = null, $end_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_List_take'] ?? \PhpursThunks::eval('Data_List_take')))(($end_1 - $start_0)))(((($GLOBALS['Data_List_drop'] ?? \PhpursThunks::eval('Data_List_drop')))($start_0))($xs_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_takeEnd'] = function() { $v = (function() {
  $__fn = function($n_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0) {
  $__fn = function($b_3 = null, $v_4 = null) use (&$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_2_0_0_b_3 = $b_3;
  $__tco_var_go_2_0_0_v_4 = $v_4;
  tco_loop_go_2_0_0:;
  $b_3 = $__tco_var_go_2_0_0_b_3;
  $v_4 = $__tco_var_go_2_0_0_v_4;
  $__t0 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Nil"))) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Cons"))) {
$__tco_1 = ($b_3 + 1);
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go_2_0_0_b_3 = $__tco_1;
$__tco_var_go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go_2_0_0;;
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
  $__res = ((($GLOBALS['Data_List_drop'] ?? \PhpursThunks::eval('Data_List_drop')))(((($go_2_0)(0))($xs_1) - $n_0)))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_deleteBy'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_deleteBy_v_0 = $v_0;
  $__tco_var_Data_List_deleteBy_v1_1 = $v1_1;
  $__tco_var_Data_List_deleteBy_v2_2 = $v2_2;
  tco_loop_Data_List_deleteBy:;
  $v_0 = $__tco_var_Data_List_deleteBy_v_0;
  $v1_1 = $__tco_var_Data_List_deleteBy_v1_1;
  $v2_2 = $__tco_var_Data_List_deleteBy_v2_2;
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Cons"))) {
$__t1 = null;;
if ((($v_0)($v1_1))(($v2_2)->{'value0'})) {
$__t1 = ($v2_2)->{'value1'};
goto end_branch_1;;
};
$__t1 = new Phpurs_Data2("Cons", ($v2_2)->{'value0'}, (((($GLOBALS['Data_List_deleteBy'] ?? \PhpursThunks::eval('Data_List_deleteBy')))($v_0))($v1_1))(($v2_2)->{'value1'}));
end_branch_1:;
$__t0 = $__t1;
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
})(); return $v; };
\PhpursThunks::$thunks['Data_List_unionBy'] = function() { $v = (function() {
  $__fn = function($eq2_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go_3_0 = null;
  $go_3_0 = (function() use ($eq2_0, &$go_3_0) {
  $__fn = function($b_4 = null, $v_5 = null) use ($eq2_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_3_0_0_b_4 = $b_4;
  $__tco_var_go_3_0_0_v_5 = $v_5;
  tco_loop_go_3_0_0:;
  $b_4 = $__tco_var_go_3_0_0_b_4;
  $v_5 = $__tco_var_go_3_0_0_v_5;
  $__t0 = null;;
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Nil"))) {
$__t0 = $b_4;
goto end_branch_0;;
};
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Cons"))) {
$__tco_1 = (((($GLOBALS['Data_List_deleteBy'] ?? \PhpursThunks::eval('Data_List_deleteBy')))($eq2_0))(($v_5)->{'value0'}))($b_4);
$__tco_2 = ($v_5)->{'value1'};
$__tco_var_go_3_0_0_b_4 = $__tco_1;
$__tco_var_go_3_0_0_v_5 = $__tco_2;
goto tco_loop_go_3_0_0;;
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
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))['foldr'])(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))((($go_3_0)(((($GLOBALS['Data_List_nubByEq'] ?? \PhpursThunks::eval('Data_List_nubByEq')))($eq2_0))($ys_2)))($xs_1)))($xs_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_union'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_unionBy'] ?? \PhpursThunks::eval('Data_List_unionBy')))(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_deleteAt'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_deleteAt_v_0 = $v_0;
  $__tco_var_Data_List_deleteAt_v1_1 = $v1_1;
  tco_loop_Data_List_deleteAt:;
  $v_0 = $__tco_var_Data_List_deleteAt_v_0;
  $v1_1 = $__tco_var_Data_List_deleteAt_v1_1;
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Cons"))) {
$__t3 = null;;
switch ($v_0) {
case 0:
$__t3 = new Phpurs_Data1("Just", ($v1_1)->{'value1'});
goto end_branch_3;;
break;
default:
;
break;
};
$__local_var_2_1 = ((($GLOBALS['Data_List_deleteAt'] ?? \PhpursThunks::eval('Data_List_deleteAt')))(($v_0 - 1)))(($v1_1)->{'value1'});
$__t2 = null;;
if ((is_object($__local_var_2_1) && (($__local_var_2_1)->{'tag'} === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", ($v1_1)->{'value0'}, ($__local_var_2_1)->{'value0'}));
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("Nothing");
end_branch_2:;
$__t3 = $__t2;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_delete'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_deleteBy'] ?? \PhpursThunks::eval('Data_List_deleteBy')))(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_difference'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use ($dictEq_0, &$go_1_0) {
  $__fn = function($b_2 = null, $v_3 = null) use ($dictEq_0, &$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_0_0_b_2 = $b_2;
  $__tco_var_go_1_0_0_v_3 = $v_3;
  tco_loop_go_1_0_0:;
  $b_2 = $__tco_var_go_1_0_0_b_2;
  $v_3 = $__tco_var_go_1_0_0_v_3;
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Nil"))) {
$__t0 = $b_2;
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Cons"))) {
$__tco_1 = (((($GLOBALS['Data_List_deleteBy'] ?? \PhpursThunks::eval('Data_List_deleteBy')))(($dictEq_0)['eq']))(($v_3)->{'value0'}))($b_2);
$__tco_2 = ($v_3)->{'value1'};
$__tco_var_go_1_0_0_b_2 = $__tco_1;
$__tco_var_go_1_0_0_v_3 = $__tco_2;
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
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_concatMap'] = function() { $v = (function() {
  $__fn = function($b_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Types_bindList')))['bind'])($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_concat'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Types_bindList')))['bind'])($v_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_catMaybes'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_0_0_0_v_1 = $v_1;
  $__tco_var_go_0_0_0_v1_2 = $v1_2;
  tco_loop_go_0_0_0:;
  $v_1 = $__tco_var_go_0_0_0_v_1;
  $v1_2 = $__tco_var_go_0_0_0_v1_2;
  $__t0 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Nil"))) {
$go_3_1 = null;
$go_3_1 = (function() use (&$__tco_var_go_0_0_0_v_1, &$__tco_var_go_0_0_0_v1_2, &$go_3_1) {
  $__fn = function($v_4 = null, $v1_5 = null) use (&$__tco_var_go_0_0_0_v_1, &$__tco_var_go_0_0_0_v1_2, &$go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_3_1_1_v_4 = $v_4;
  $__tco_var_go_3_1_1_v1_5 = $v1_5;
  tco_loop_go_3_1_1:;
  $v_4 = $__tco_var_go_3_1_1_v_4;
  $v1_5 = $__tco_var_go_3_1_1_v1_5;
  $__t1 = null;;
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$__t1 = $v_4;
goto end_branch_1;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_5)->{'value0'}, $v_4);
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go_3_1_1_v_4 = $__tco_2;
$__tco_var_go_3_1_1_v1_5 = $__tco_3;
goto tco_loop_go_3_1_1;;
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
$__t0 = (($go_3_1)(new Phpurs_Data0("Nil")))($v_1);
goto end_branch_0;;
};
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Cons"))) {
$__t2 = null;;
if ((is_object(($v1_2)->{'value0'}) && ((($v1_2)->{'value0'})->{'tag'} === "Nothing"))) {
$__tco_3 = $v_1;
$__tco_4 = ($v1_2)->{'value1'};
$__tco_var_go_0_0_0_v_1 = $__tco_3;
$__tco_var_go_0_0_0_v1_2 = $__tco_4;
goto tco_loop_go_0_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object(($v1_2)->{'value0'}) && ((($v1_2)->{'value0'})->{'tag'} === "Just"))) {
$__tco_5 = new Phpurs_Data2("Cons", (($v1_2)->{'value0'})->{'value0'}, $v_1);
$__tco_6 = ($v1_2)->{'value1'};
$__tco_var_go_0_0_0_v_1 = $__tco_5;
$__tco_var_go_0_0_0_v1_2 = $__tco_6;
goto tco_loop_go_0_0_0;;
$__t2 = null;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
return ($go_0_0)(new Phpurs_Data0("Nil"));
})(); return $v; };
\PhpursThunks::$thunks['Data_List_alterAt'] = function() { $v = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_alterAt_v_0 = $v_0;
  $__tco_var_Data_List_alterAt_v1_1 = $v1_1;
  $__tco_var_Data_List_alterAt_v2_2 = $v2_2;
  tco_loop_Data_List_alterAt:;
  $v_0 = $__tco_var_Data_List_alterAt_v_0;
  $v1_1 = $__tco_var_Data_List_alterAt_v1_1;
  $v2_2 = $__tco_var_Data_List_alterAt_v2_2;
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Cons"))) {
$__t3 = null;;
switch ($v_0) {
case 0:
$v3_3_4 = ($v1_1)(($v2_2)->{'value0'});
$__t5 = null;;
if ((is_object($v3_3_4) && (($v3_3_4)->{'tag'} === "Nothing"))) {
$__t5 = ($v2_2)->{'value1'};
goto end_branch_5;;
};
if ((is_object($v3_3_4) && (($v3_3_4)->{'tag'} === "Just"))) {
$__t5 = new Phpurs_Data2("Cons", ($v3_3_4)->{'value0'}, ($v2_2)->{'value1'});
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t3 = new Phpurs_Data1("Just", $__t5);
goto end_branch_3;;
break;
default:
;
break;
};
$__local_var_3_1 = (((($GLOBALS['Data_List_alterAt'] ?? \PhpursThunks::eval('Data_List_alterAt')))(($v_0 - 1)))($v1_1))(($v2_2)->{'value1'});
$__t2 = null;;
if ((is_object($__local_var_3_1) && (($__local_var_3_1)->{'tag'} === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", ($v2_2)->{'value0'}, ($__local_var_3_1)->{'value0'}));
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("Nothing");
end_branch_2:;
$__t3 = $__t2;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data0("Nothing");
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_modifyAt'] = function() { $v = (function() {
  $__fn = function($n_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_alterAt'] ?? \PhpursThunks::eval('Data_List_alterAt')))($n_0))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))($f_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };














































































