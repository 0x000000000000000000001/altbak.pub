<?php

namespace Data\List\Lazy\Types;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Comonad, Control.Extend, Control.Lazy, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.Lazy, Data.List.Lazy.Types, Data.Maybe, Data.Monoid, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Comonad, Control.Extend, Control.Lazy, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.Lazy, Data.List.Lazy.Types, Data.Maybe, Data.Monoid, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Lazy.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
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
\PhpursThunks::$thunks['Data_List_Lazy_Types_List'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_Nil'] = function() { $v = ($GLOBALS['__phpurs_data0_Nil'] ??= new Phpurs_Data0("Nil")); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_Cons'] = function() { $v = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Cons", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_NonEmptyList'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_nil'] = function() { $v = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("Nil");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_newtypeNonEmptyList'] = function() { $v = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_newtypeList'] = function() { $v = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_step'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force'))))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_semigroupList'] = function() { $v = ["append" => (function() {
  $__fn = function($xs_0 = null, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($xs_0, $ys_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($xs_0);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nil"))) {
$__t1 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($ys_1);
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Cons"))) {
$__t1 = new Phpurs_Data2("Cons", ($__local_var_3_0)->{'value0'}, (((($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList')))['append'])(($__local_var_3_0)->{'value1'}))($ys_1));
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_monoidList'] = function() { $v = ["mempty" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')), "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_lazyList'] = function() { $v = ["defer" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step'))))($f_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_functorList'] = function() { $v = ["map" => (function() {
  $__fn = function($f_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($f_0, $xs_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($xs_1);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Cons"))) {
$__t1 = new Phpurs_Data2("Cons", ($f_0)(($__local_var_3_0)->{'value0'}), (((($GLOBALS['Data_List_Lazy_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorList')))['map'])($f_0))(($__local_var_3_0)->{'value1'}));
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
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_functorNonEmptyList'] = function() { $v = ["map" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($f_0, $v_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_1);
  $__res = new Phpurs_Data2("NonEmpty", ($f_0)(($__local_var_3_0)->{'value0'}), (((($GLOBALS['Data_List_Lazy_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorList')))['map'])($f_0))(($__local_var_3_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_eq1List'] = function() { $v = ["eq1" => (function() {
  $__fn = function($dictEq_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go_3_0 = null;
  $go_3_0 = (function() use ($dictEq_0, &$go_3_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictEq_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Nil"))) {
$__t1 = (is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"));
goto end_branch_1;;
};
  $__t1 = ((is_object($v_4) && (($v_4)->{'tag'} === "Cons")) && ((is_object($v1_5) && (($v1_5)->{'tag'} === "Cons")) && (((($dictEq_0)['eq'])(($v_4)->{'value0'}))(($v1_5)->{'value0'}) && (($go_3_0)((($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v_4)->{'value1'})))((($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v1_5)->{'value1'})))));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_3_0)((($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($xs_1)))((($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($ys_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_eqNonEmpty'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictEq_0)['eq'])(($x_1)->{'value0'}))(($y_2)->{'value0'}) && ((((($GLOBALS['Data_List_Lazy_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eq1List')))['eq1'])($dictEq_0))(($x_1)->{'value1'}))(($y_2)->{'value1'}));
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
\PhpursThunks::$thunks['Data_List_Lazy_Types_eq1NonEmptyList'] = function() { $v = ["eq1" => function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_eq1Lazy'] ?? \PhpursThunks::eval('Data_Lazy_eq1Lazy')))['eq1'])((($GLOBALS['Data_List_Lazy_Types_eqNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eqNonEmpty')))($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_eqList'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq" => ((($GLOBALS['Data_List_Lazy_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eq1List')))['eq1'])($dictEq_0)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_eqNonEmptyList'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Data_List_Lazy_Types_eqNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eqNonEmpty')))($dictEq_0);
  $__res = ["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['eq'])((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($x_2)))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($y_3));
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
\PhpursThunks::$thunks['Data_List_Lazy_Types_ord1List'] = function() { $v = ["compare1" => (function() {
  $__fn = function($dictOrd_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go_3_0 = null;
  $go_3_0 = (function() use ($dictOrd_0, &$go_3_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictOrd_0, &$go_3_0, &$__fn) {
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
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Nil"))) {
$__t1 = null;;
if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("EQ");
goto end_branch_1;;
};
$__t1 = new Phpurs_Data0("LT");
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Nil"))) {
$__t0 = new Phpurs_Data0("GT");
goto end_branch_0;;
};
  if (((is_object($v_4) && (($v_4)->{'tag'} === "Cons")) && (is_object($v1_5) && (($v1_5)->{'tag'} === "Cons")))) {
$v2_6_2 = ((($dictOrd_0)['compare'])(($v_4)->{'value0'}))(($v1_5)->{'value0'});
$__t3 = null;;
if ((is_object($v2_6_2) && (($v2_6_2)->{'tag'} === "EQ"))) {
$__tco_4 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v_4)->{'value1'});
$__tco_5 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v1_5)->{'value1'});
$__tco_var_go_3_0_0_v_4 = $__tco_4;
$__tco_var_go_3_0_0_v1_5 = $__tco_5;
goto tco_loop_go_3_0_0;;
$__t3 = null;
goto end_branch_3;;
};
$__t3 = $v2_6_2;
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
  $__res = (($go_3_0)((($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($xs_1)))((($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($ys_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Eq10" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eq1List'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_ordNonEmpty'] = function() { $v = (($GLOBALS['Data_NonEmpty_ordNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_ordNonEmpty')))(($GLOBALS['Data_List_Lazy_Types_ord1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_ord1List'))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_ord1NonEmptyList'] = function() { $v = ["compare1" => function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_ordLazy'] ?? \PhpursThunks::eval('Data_Lazy_ordLazy')))((($GLOBALS['Data_List_Lazy_Types_ordNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_ordNonEmpty')))($dictOrd_0)))['compare'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_eq1NonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eq1NonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_ordList'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ["compare" => ((($GLOBALS['Data_List_Lazy_Types_ord1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_ord1List')))['compare1'])($dictOrd_0), "Eq0" => function($dollar__unused_1 = null) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = ["eq" => ((($GLOBALS['Data_List_Lazy_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eq1List')))['eq1'])((($dictOrd_0)['Eq0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_ordNonEmptyList'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_ordLazy'] ?? \PhpursThunks::eval('Data_Lazy_ordLazy')))((($GLOBALS['Data_List_Lazy_Types_ordNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_ordNonEmpty')))($dictOrd_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_cons'] = function() { $v = (function() {
  $__fn = function($x_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($x_0, $xs_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $x_0, $xs_1);
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
\PhpursThunks::$thunks['Data_List_Lazy_Types_foldableList'] = function() { $v = ["foldr" => (function() {
  $__fn = function($op_0 = null, $z_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldl'])((function() use ($op_0) {
  $__fn = function($b_3 = null, $a_4 = null) use ($op_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($op_0)($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($z_1))(((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldl'])((function() {
  $__fn = function($b_3 = null, $a_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($a_4, $b_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'))))($xs_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => function($op_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $op_0) {
  $__fn = function($b_2 = null, $xs_3 = null) use (&$go_1_0, $op_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_0_0_b_2 = $b_2;
  $__tco_var_go_1_0_0_xs_3 = $xs_3;
  tco_loop_go_1_0_0:;
  $b_2 = $__tco_var_go_1_0_0_b_2;
  $xs_3 = $__tco_var_go_1_0_0_xs_3;
  $v_4_0 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($xs_3);
  $__t1 = null;;
  if ((is_object($v_4_0) && (($v_4_0)->{'tag'} === "Nil"))) {
$__t1 = $b_2;
goto end_branch_1;;
};
  if ((is_object($v_4_0) && (($v_4_0)->{'tag'} === "Cons"))) {
$__tco_2 = (($op_0)($b_2))(($v_4_0)->{'value0'});
$__tco_3 = ($v_4_0)->{'value1'};
$__tco_var_go_1_0_0_b_2 = $__tco_2;
$__tco_var_go_1_0_0_xs_3 = $__tco_3;
goto tco_loop_go_1_0_0;;
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
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_1 = ($dictMonoid_0)['mempty'];
  $__res = function($f_2 = null) use ($dictMonoid_0, $mempty_1_1) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldl'])((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($b_3 = null, $a_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_0)['Semigroup0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['append'])($b_3))(($f_2)($a_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_1_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_foldableNonEmpty'] = function() { $v = (($GLOBALS['Data_NonEmpty_foldableNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldableNonEmpty')))(($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList'))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_extendList'] = function() { $v = ["extend" => (function() {
  $__fn = function($f_0 = null, $l_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($l_1);
  $__t1 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "Nil"))) {
$__t1 = ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'));
goto end_branch_1;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "Cons"))) {
$__local_var_3_2 = ($f_0)($l_1);
$__local_var_4_3 = (((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldr'])((function() use ($f_0) {
  $__fn = function($a_4 = null, $v_5 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_3 = ($v_5)['acc'];
  $__local_var_7_4 = ($v_5)['val'];
  $acc__prime___8_5 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_8 = null) use ($__local_var_6_3, $a_4) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_4, $__local_var_6_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_9_6 = ($f_0)($acc__prime___8_5);
  $__res = ["val" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_10 = null) use ($__local_var_7_4, $__local_var_9_6) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_9_6, $__local_var_7_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "acc" => $acc__prime___8_5];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(["val" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')), "acc" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'))]))(($v_2_0)->{'value1'}))['val'];
$__t1 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($__local_var_3_2, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_3_2, $__local_var_4_3);
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_extendNonEmptyList'] = function() { $v = ["extend" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_1))->{'value1'};
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v2_3 = null) use ($__local_var_2_0, $f_0, $v_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", ($f_0)($v_1), (((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldr'])((function() use ($f_0) {
  $__fn = function($a_4 = null, $v1_5 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_1 = ($v1_5)['acc'];
  $__local_var_7_2 = ($v1_5)['val'];
  $__local_var_8_3 = ($f_0)((($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v2_8 = null) use ($__local_var_6_1, $a_4) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $a_4, $__local_var_6_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__res = ["val" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_9 = null) use ($__local_var_7_2, $__local_var_8_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_8_3, $__local_var_7_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "acc" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_8 = null) use ($__local_var_6_1, $a_4) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_4, $__local_var_6_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(["val" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')), "acc" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'))]))($__local_var_2_0))['val']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_foldableNonEmptyList'] = function() { $v = ["foldr" => (function() {
  $__fn = function($f_0 = null, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableNonEmpty')))['foldr'])($f_0))($b_1))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0 = null, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableNonEmpty')))['foldl'])($f_0))($b_1))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $foldMap1_1_0 = ((($GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableNonEmpty')))['foldMap'])($dictMonoid_0);
  $__res = (function() use ($foldMap1_1_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($foldMap1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap1_1_0)($f_2))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_showList'] = function() { $v = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = ["show" => function($xs_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $v_2_0 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($xs_1);
  $__t1 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "Nil"))) {
$__t1 = "(fromFoldable [])";
goto end_branch_1;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "Cons"))) {
$__t1 = ((("(fromFoldable [" . (($dictShow_0)['show'])(($v_2_0)->{'value0'})) . ((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldl'])((function() use ($dictShow_0) {
  $__fn = function($shown_3 = null, $x__prime___4 = null) use ($dictShow_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($shown_3 . ",") . (($dictShow_0)['show'])($x__prime___4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(""))(($v_2_0)->{'value1'})) . "])");
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_showNonEmptyList'] = function() { $v = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Data_List_Lazy_Types_showList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_showList')))($dictShow_0);
  $__res = ["show" => function($v_2 = null) use ($__local_var_1_0, $dictShow_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_2);
  $__res = (((("(NonEmptyList (defer \\_ -> (NonEmpty " . (($dictShow_0)['show'])(($__local_var_3_1)->{'value0'})) . " ") . (($__local_var_1_0)['show'])(($__local_var_3_1)->{'value1'})) . ")))");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_showStep'] = function() { $v = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = ["show" => function($v_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Nil"))) {
$__t0 = "Nil";
goto end_branch_0;;
};
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Cons"))) {
$__t0 = (((("(" . (($dictShow_0)['show'])(($v_1)->{'value0'})) . " : ") . (((($GLOBALS['Data_List_Lazy_Types_showList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_showList')))($dictShow_0))['show'])(($v_1)->{'value1'})) . ")");
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_foldableWithIndexList'] = function() { $v = ["foldrWithIndex" => (function() {
  $__fn = function($f_0 = null, $b_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_3_0 = ((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldl'])(function($v1_3 = null) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($v1_3)->{'value1'};
  $__local_var_5_1 = ($v1_3)->{'value0'};
  $__res = function($a_6 = null) use ($__local_var_4_0, $__local_var_5_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($__local_var_5_1 + 1), (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7 = null) use ($__local_var_4_0, $a_6) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_6, $__local_var_4_0);
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
}))(new Phpurs_Data2("Tuple", 0, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')))))($xs_2);
  $__res = (((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldl'])(function($v1_4 = null) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_5_3 = ($v1_4)->{'value1'};
  $__local_var_6_4 = ($v1_4)->{'value0'};
  $__res = function($a_7 = null) use ($__local_var_5_3, $__local_var_6_4, $f_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($__local_var_6_4 - 1), ((($f_0)(($__local_var_6_4 - 1)))($a_7))($__local_var_5_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", ($v_3_0)->{'value0'}, $b_1)))(($v_3_0)->{'value1'}))->{'value1'};
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0 = null, $acc_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Tuple_snd'] ?? \PhpursThunks::eval('Data_Tuple_snd'))))((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldl'])(function($v_2 = null) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_3_5 = ($v_2)->{'value1'};
  $__local_var_4_6 = ($v_2)->{'value0'};
  $__res = function($a_5 = null) use ($__local_var_3_5, $__local_var_4_6, $f_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($__local_var_4_6 + 1), ((($f_0)($__local_var_4_6))($__local_var_3_5))($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", 0, $acc_1)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_7 = ($dictMonoid_0)['mempty'];
  $__res = function($f_2 = null) use ($dictMonoid_0, $mempty_1_7) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexList')))['foldlWithIndex'])((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($i_3 = null, $acc_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($dictMonoid_0)['Semigroup0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['append'])($acc_4)))(($f_2)($i_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_foldableWithIndexNonEmpty'] = function() { $v = (($GLOBALS['Data_NonEmpty_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_foldableWithIndexNonEmpty')))(($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexList'))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_foldableWithIndexNonEmptyList'] = function() { $v = ["foldMapWithIndex" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $foldMapWithIndex1_1_0 = ((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexNonEmpty')))['foldMapWithIndex'])($dictMonoid_0);
  $__res = (function() use ($foldMapWithIndex1_1_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($foldMapWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMapWithIndex1_1_0)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_2))(function($v2_4 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Nothing"))) {
$__t1 = 0;
goto end_branch_1;;
};
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Just"))) {
$__t1 = (1 + ($v2_4)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => (function() {
  $__fn = function($f_0 = null, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexNonEmpty')))['foldlWithIndex'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_0))(function($v2_3 = null) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Nothing"))) {
$__t2 = 0;
goto end_branch_2;;
};
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Just"))) {
$__t2 = (1 + ($v2_3)->{'value0'});
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($b_1))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldrWithIndex" => (function() {
  $__fn = function($f_0 = null, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexNonEmpty')))['foldrWithIndex'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_0))(function($v2_3 = null) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Nothing"))) {
$__t3 = 0;
goto end_branch_3;;
};
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Just"))) {
$__t3 = (1 + ($v2_3)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($b_1))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_foldableNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_functorWithIndexList'] = function() { $v = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexList')))['foldrWithIndex'])((function() use ($f_0) {
  $__fn = function($i_1 = null, $x_2 = null, $acc_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_4_0 = (($f_0)($i_1))($x_2);
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($__local_var_4_0, $acc_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_4_0, $acc_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_mapWithIndex'] = function() { $v = ((($GLOBALS['Data_NonEmpty_functorWithIndex'] ?? \PhpursThunks::eval('Data_NonEmpty_functorWithIndex')))(($GLOBALS['Data_List_Lazy_Types_functorWithIndexList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorWithIndexList'))))['mapWithIndex']; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_functorWithIndexNonEmptyList'] = function() { $v = ["mapWithIndex" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v1_2 = null) use ($f_0, $v_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_mapWithIndex'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_mapWithIndex')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_0))(function($v2_3 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Nothing"))) {
$__t0 = 0;
goto end_branch_0;;
};
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Just"))) {
$__t0 = (1 + ($v2_3)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_toList'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step'))))(function($v1_1 = null) use ($v_0) {
  $__num = \func_num_args();
  $v2_2_0 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_0);
  $__local_var_3_1 = ($v2_2_0)->{'value0'};
  $__local_var_4_2 = ($v2_2_0)->{'value1'};
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5 = null) use ($__local_var_3_1, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_3_1, $__local_var_4_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_semigroupNonEmptyList'] = function() { $v = ["append" => (function() {
  $__fn = function($v_0 = null, $as__prime___1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_0);
  $__local_var_3_1 = ($v1_2_0)->{'value0'};
  $__local_var_4_2 = ($v1_2_0)->{'value1'};
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v2_5 = null) use ($__local_var_3_1, $__local_var_4_2, $as__prime___1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $__local_var_3_1, (((($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList')))['append'])($__local_var_4_2))((($GLOBALS['Data_List_Lazy_Types_toList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_toList')))($as__prime___1)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_traversableList'] = function() { $v = ["traverse" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($f_2 = null) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))['foldr'])((function() use ($Apply0_1_0, $f_2) {
  $__fn = function($a_3 = null, $b_4 = null) use ($Apply0_1_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Apply0_1_0)['apply'])(((((($Apply0_1_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_List_Lazy_Types_cons'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_cons'))))(($f_2)($a_3))))($b_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($dictApplicative_0)['pure'])(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traversableList')))['traverse'])($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_traversableNonEmpty'] = function() { $v = (($GLOBALS['Data_NonEmpty_traversableNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_traversableNonEmpty')))(($GLOBALS['Data_List_Lazy_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traversableList'))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_traversableNonEmptyList'] = function() { $v = ["traverse" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $traverse1_1_0 = ((($GLOBALS['Data_List_Lazy_Types_traversableNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traversableNonEmpty')))['traverse'])($dictApplicative_0);
  $__res = (function() use ($dictApplicative_0, $traverse1_1_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($dictApplicative_0, $traverse1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(function($xxs_4 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v1_5 = null) use ($xxs_4) {
  $__num = \func_num_args();
  $__res = $xxs_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($traverse1_1_0)($f_2))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_3)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $sequence1_1_1 = ((($GLOBALS['Data_List_Lazy_Types_traversableNonEmpty'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traversableNonEmpty')))['sequence'])($dictApplicative_0);
  $__res = function($v_2 = null) use ($dictApplicative_0, $sequence1_1_1) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(function($xxs_3 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v1_4 = null) use ($xxs_3) {
  $__num = \func_num_args();
  $__res = $xxs_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($sequence1_1_1)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_2)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_foldableNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_traversableWithIndexList'] = function() { $v = ["traverseWithIndex" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($f_2 = null) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexList')))['foldrWithIndex'])((function() use ($Apply0_1_0, $f_2) {
  $__fn = function($i_3 = null, $a_4 = null, $b_5 = null) use ($Apply0_1_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Apply0_1_0)['apply'])(((((($Apply0_1_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_List_Lazy_Types_cons'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_cons'))))((($f_2)($i_3))($a_4))))($b_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))((($dictApplicative_0)['pure'])(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorWithIndexList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorWithIndexList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traversableList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_traverseWithIndex'] = function() { $v = ((($GLOBALS['Data_NonEmpty_traversableWithIndexNonEmpty'] ?? \PhpursThunks::eval('Data_NonEmpty_traversableWithIndexNonEmpty')))(($GLOBALS['Data_List_Lazy_Types_traversableWithIndexList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traversableWithIndexList'))))['traverseWithIndex']; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_traversableWithIndexNonEmptyList'] = function() { $v = ["traverseWithIndex" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $traverseWithIndex1_1_0 = (($GLOBALS['Data_List_Lazy_Types_traverseWithIndex'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traverseWithIndex')))($dictApplicative_0);
  $__res = (function() use ($dictApplicative_0, $traverseWithIndex1_1_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($dictApplicative_0, $traverseWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(function($xxs_4 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v1_5 = null) use ($xxs_4) {
  $__num = \func_num_args();
  $__res = $xxs_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($traverseWithIndex1_1_0)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_2))(function($v2_4 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Nothing"))) {
$__t1 = 0;
goto end_branch_1;;
};
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Just"))) {
$__t1 = (1 + ($v2_4)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_3)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorWithIndexNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorWithIndexNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableWithIndexNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_traversableNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traversableNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_unfoldable1List'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($f_1 = null, $b_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step'))))(function($v_3 = null) use ($b_2, $f_1, &$go_0_0) {
  $__num = \func_num_args();
  $v1_4_1 = ($f_1)($b_2);
  $__t2 = null;;
  if ((is_object(($v1_4_1)->{'value1'}) && ((($v1_4_1)->{'value1'})->{'tag'} === "Just"))) {
$__local_var_5_3 = ($v1_4_1)->{'value0'};
$__local_var_6_4 = (($go_0_0)($f_1))((($v1_4_1)->{'value1'})->{'value0'});
$__t2 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7 = null) use ($__local_var_5_3, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_5_3, $__local_var_6_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_2;;
};
  if ((is_object(($v1_4_1)->{'value1'}) && ((($v1_4_1)->{'value1'})->{'tag'} === "Nothing"))) {
$__local_var_5_5 = ($v1_4_1)->{'value0'};
$__t2 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6 = null) use ($__local_var_5_5) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_5_5, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
return ["unfoldr1" => $go_0_0];
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_unfoldableList'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($f_1 = null, $b_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step'))))(function($v_3 = null) use ($b_2, $f_1, &$go_0_0) {
  $__num = \func_num_args();
  $v1_4_1 = ($f_1)($b_2);
  $__t2 = null;;
  if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "Nothing"))) {
$__t2 = ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'));
goto end_branch_2;;
};
  if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "Just"))) {
$__local_var_5_3 = (($v1_4_1)->{'value0'})->{'value0'};
$__local_var_6_4 = (($go_0_0)($f_1))((($v1_4_1)->{'value0'})->{'value1'});
$__t2 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7 = null) use ($__local_var_5_3, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_5_3, $__local_var_6_4);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
return ["unfoldr" => $go_0_0, "Unfoldable10" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_unfoldable1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_unfoldable1List'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_unfoldr1'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ($f_0)($b_1);
  $__res = new Phpurs_Data2("NonEmpty", ($__local_var_2_0)->{'value0'}, (((($GLOBALS['Data_List_Lazy_Types_unfoldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_unfoldableList')))['unfoldr'])(function($v1_3 = null) use ($f_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_3) && (($v1_3)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($f_0)(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($__local_var_2_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_unfoldable1NonEmptyList'] = function() { $v = ["unfoldr1" => (function() {
  $__fn = function($f_0 = null, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_unfoldr1'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_unfoldr1')))($f_0))($b_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_comonadNonEmptyList'] = function() { $v = ["extract" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_0))->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_extendNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_extendNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_monadList'] = function() { $v = ["Applicative0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_applicativeList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applicativeList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_bindList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_bindList'] = function() { $v = ["bind" => (function() {
  $__fn = function($xs_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2 = null) use ($f_1, $xs_0) {
  $__num = \func_num_args();
  $__local_var_3_0 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($xs_0);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Cons"))) {
$__t1 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))((((($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList')))['append'])(($f_1)(($__local_var_3_0)->{'value0'})))((((($GLOBALS['Data_List_Lazy_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_bindList')))['bind'])(($__local_var_3_0)->{'value1'}))($f_1)));
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
})(), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_applyList'] = function() { $v = ["apply" => (($GLOBALS['Control_Monad_ap'] ?? \PhpursThunks::eval('Control_Monad_ap')))(($GLOBALS['Data_List_Lazy_Types_monadList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_monadList'))), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_applicativeList'] = function() { $v = ["pure" => function($a_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_1 = null) use ($a_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $a_0, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_applyNonEmptyList'] = function() { $v = ["apply" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v2_2_0 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v1_1);
  $v3_3_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_0);
  $__local_var_4_2 = ($v2_2_0)->{'value0'};
  $__local_var_5_3 = ($v2_2_0)->{'value1'};
  $__local_var_6_4 = ($v3_3_1)->{'value0'};
  $__local_var_7_5 = ($v3_3_1)->{'value1'};
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v4_8 = null) use ($__local_var_4_2, $__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", ($__local_var_6_4)($__local_var_4_2), (((($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList')))['append'])((((($GLOBALS['Data_List_Lazy_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applyList')))['apply'])($__local_var_7_5))((($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_9 = null) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_4_2, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))((((($GLOBALS['Data_List_Lazy_Types_applyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applyList')))['apply'])((($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_9 = null) use ($__local_var_6_4, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_6_4, $__local_var_7_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($__local_var_5_3)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_bindNonEmptyList'] = function() { $v = ["bind" => (function() {
  $__fn = function($v_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_0);
  $__local_var_3_1 = ($v1_2_0)->{'value1'};
  $v2_4_2 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($f_1)(($v1_2_0)->{'value0'}));
  $__local_var_5_3 = ($v2_4_2)->{'value0'};
  $__local_var_6_4 = ($v2_4_2)->{'value1'};
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v3_7 = null) use ($__local_var_3_1, $__local_var_5_3, $__local_var_6_4, $f_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $__local_var_5_3, (((($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList')))['append'])($__local_var_6_4))((((($GLOBALS['Data_List_Lazy_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_bindList')))['bind'])($__local_var_3_1))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_toList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_toList'))))($f_1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_applyNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applyNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_altNonEmptyList'] = function() { $v = ["alt" => (($GLOBALS['Data_List_Lazy_Types_semigroupNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupNonEmptyList')))['append'], "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_altList'] = function() { $v = ["alt" => (($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList')))['append'], "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_plusList'] = function() { $v = ["empty" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')), "Alt0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_altList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_altList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_alternativeList'] = function() { $v = ["Applicative0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_applicativeList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applicativeList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_plusList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_plusList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_monadPlusList'] = function() { $v = ["Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_monadList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_monadList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_alternativeList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_alternativeList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_applicativeNonEmptyList'] = function() { $v = ["pure" => function($a_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_1 = null) use ($a_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $a_0, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_applyNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applyNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Types_monadNonEmptyList'] = function() { $v = ["Applicative0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_applicativeNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_applicativeNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_List_Lazy_Types_bindNonEmptyList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_bindNonEmptyList'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


































































