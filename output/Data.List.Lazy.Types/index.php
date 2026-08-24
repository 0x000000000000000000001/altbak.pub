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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_List_Lazy_Types_Nil { public $tag = 'Nil'; public function __construct() {} }
final class Data_List_Lazy_Types_Cons { public $tag = 'Cons'; public function __construct(public  $value0, public  $value1) {} }

// Data_List_Lazy_Types_List
function majData_majList_majLazy_majTypes_majList($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_majList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_List'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_majList';

// Data_List_Lazy_Types_Nil
$GLOBALS['Data_List_Lazy_Types_Nil'] = ($GLOBALS['__phpurs_data0_Nil'] ??= new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil());

// Data_List_Lazy_Types_Cons
$GLOBALS['Data_List_Lazy_Types_Cons'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_Types_NonEmptyList
function majData_majList_majLazy_majTypes_majNonmajEmptymajList($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_majNonmajEmptymajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_NonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_majNonmajEmptymajList';

// Data_List_Lazy_Types_nil
$GLOBALS['Data_List_Lazy_Types_nil'] = \Data\Lazy\majData_majLazy_defer(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_List_Lazy_Types_newtypeNonEmptyList
$GLOBALS['Data_List_Lazy_Types_newtypeNonEmptyList'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_newtypeList
$GLOBALS['Data_List_Lazy_Types_newtypeList'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_step_closure
$GLOBALS['Data_List_Lazy_Types_step_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Lazy_force']))($GLOBALS['Unsafe_Coerce_unsafeCoerce']);

// Data_List_Lazy_Types_step
function majData_majList_majLazy_majTypes_step($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_step';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_Types_step_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_step'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_step';

// Data_List_Lazy_Types_semigroupList
$GLOBALS['Data_List_Lazy_Types_semigroupList'] = (object)["append" => function($xs_0) {
  $__num = \func_num_args();
  $__res = function($ys_1) use ($xs_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($xs_0, $ys_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = \Data\Lazy\majData_majLazy_force($xs_0);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_1);
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_3_0)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_3_0)->{'value1'}))($ys_1));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_monoidList
$GLOBALS['Data_List_Lazy_Types_monoidList'] = (object)["mempty" => $GLOBALS['Data_List_Lazy_Types_nil'], "Semigroup0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_semigroupList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_lazyList
$GLOBALS['Data_List_Lazy_Types_lazyList'] = (object)["defer" => function($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))($f_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_functorList
$GLOBALS['Data_List_Lazy_Types_functorList'] = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($xs_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($f_0, $xs_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = \Data\Lazy\majData_majLazy_force($xs_1);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($f_0)(($__local_var_3_0)->{'value0'}), ((($GLOBALS['Data_List_Lazy_Types_functorList'])->{'map'})($f_0))(($__local_var_3_0)->{'value1'}));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_functorNonEmpty
$GLOBALS['Data_List_Lazy_Types_functorNonEmpty'] = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($m_1) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = ($m_1)->{'value1'};
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)(($m_1)->{'value0'}), \Data\Lazy\majData_majLazy_defer(function($v_3) use ($__local_var_2_0, $f_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = \Data\Lazy\majData_majLazy_force($__local_var_2_0);
  $__t2 = null;;
  if ($__local_var_4_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t2 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_2;;
};
  if ($__local_var_4_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t2 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($f_0)(($__local_var_4_1)->{'value0'}), ((($GLOBALS['Data_List_Lazy_Types_functorList'])->{'map'})($f_0))(($__local_var_4_1)->{'value1'}));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_functorNonEmptyList
$GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'] = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($f_0, $v_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_functorNonEmpty'])->{'map'})($f_0))(\Data\Lazy\majData_majLazy_force($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_eq1List
$GLOBALS['Data_List_Lazy_Types_eq1List'] = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($xs_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($ys_2) use ($dictEq_0, $xs_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil;
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && (((($dictEq_0)->{'eq'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}) && (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_4)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_5)->{'value1'})))));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_1)))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_2));
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

// Data_List_Lazy_Types_eq1NonEmptyList
$GLOBALS['Data_List_Lazy_Types_eq1NonEmptyList'] = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $eqNonEmpty1_1_0 = (object)["eq" => function($x_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($y_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil;
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && (((($dictEq_0)->{'eq'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}) && (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_4)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_5)->{'value1'})))));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'}) && (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($x_1)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($y_2)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($v_2) use ($eqNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($eqNonEmpty1_1_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($eqNonEmpty1_1_0)->{'eq'})(\Data\Lazy\majData_majLazy_force($v_2)))(\Data\Lazy\majData_majLazy_force($v1_3));
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

// Data_List_Lazy_Types_eqList
function majData_majList_majLazy_majTypes_eqmajList($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_eqmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq" => function($xs_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($ys_2) use ($dictEq_0, $xs_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil;
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && (((($dictEq_0)->{'eq'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}) && (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_4)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_5)->{'value1'})))));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_1)))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_2));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_eqList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_eqmajList';

// Data_List_Lazy_Types_eqNonEmptyList
function majData_majList_majLazy_majTypes_eqmajNonmajEmptymajList($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_eqmajNonmajEmptymajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (object)["eq" => function($x_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($y_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil;
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && (((($dictEq_0)->{'eq'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}) && (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_4)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_5)->{'value1'})))));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'}) && (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($x_1)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($y_2)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["eq" => function($x_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($__local_var_1_0, $x_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'eq'})(\Data\Lazy\majData_majLazy_force($x_2)))(\Data\Lazy\majData_majLazy_force($y_3));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_eqNonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_eqmajNonmajEmptymajList';

// Data_List_Lazy_Types_ord1List
$GLOBALS['Data_List_Lazy_Types_ord1List'] = (object)["compare1" => function($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($xs_1) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($ys_2) use ($dictOrd_0, $xs_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = (function() use ($dictOrd_0, &$go__go_3_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictOrd_0, &$go__go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_0_0_v_4 = $v_4;
  $__tco_var_go__go_3_0_0_v1_5 = $v1_5;
  tco_loop_go__go_3_0_0:;
  $v_4 = $__tco_var_go__go_3_0_0_v_4;
  $v1_5 = $__tco_var_go__go_3_0_0_v1_5;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = null;;
if ($v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if (($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && $v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons)) {
$v2_6_2 = ((($dictOrd_0)->{'compare'})(($v_4)->{'value0'}))(($v1_5)->{'value0'});
$__t3 = null;;
if ($v2_6_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_4 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_4)->{'value1'});
$__tco_5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_5)->{'value1'});
$__tco_var_go__go_3_0_0_v_4 = $__tco_4;
$__tco_var_go__go_3_0_0_v1_5 = $__tco_5;
goto tco_loop_go__go_3_0_0;;
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
  $__res = (($go__go_3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_1)))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_2));
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
}, "Eq10" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_eq1List'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_ord1NonEmptyList
$GLOBALS['Data_List_Lazy_Types_ord1NonEmptyList'] = (object)["compare1" => function($dictOrd_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqNonEmpty2_1_0 = (object)["eq" => function($x_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($__local_var_1_0, $x_2) {
  $__num = \func_num_args();
  $go__go_4_1 = null;
  $go__go_4_1 = function($v_5) use ($__local_var_1_0, &$go__go_4_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_1_0, &$go__go_4_1, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t2 = $v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil;
goto end_branch_2;;
};
  $__t2 = ($v_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && (((($__local_var_1_0)->{'eq'})(($v_5)->{'value0'}))(($v1_6)->{'value0'}) && (($go__go_4_1)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_5)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_6)->{'value1'})))));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (((($__local_var_1_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && (($go__go_4_1)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($x_2)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($y_3)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordNonEmpty1_1_0 = (object)["compare" => function($x_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictOrd_0, $x_2) {
  $__num = \func_num_args();
  $v_4_4 = ((($dictOrd_0)->{'compare'})(($x_2)->{'value0'}))(($y_3)->{'value0'});
  $__t6 = null;;
  if ($v_4_4 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t6 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_6;;
};
  if ($v_4_4 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t6 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_6;;
};
  $go__go_5_5 = null;
  $go__go_5_5 = (function() use ($dictOrd_0, &$go__go_5_5) {
  $__fn = function($v_6, $v1_7 = null) use ($dictOrd_0, &$go__go_5_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_5_5_v_6 = $v_6;
  $__tco_var_go__go_5_5_5_v1_7 = $v1_7;
  tco_loop_go__go_5_5_5:;
  $v_6 = $__tco_var_go__go_5_5_5_v_6;
  $v1_7 = $__tco_var_go__go_5_5_5_v1_7;
  $__t5 = null;;
  if ($v_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t6 = null;;
if ($v1_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t6 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
  if ($v1_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  if (($v_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && $v1_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons)) {
$v2_8_7 = ((($dictOrd_0)->{'compare'})(($v_6)->{'value0'}))(($v1_7)->{'value0'});
$__t8 = null;;
if ($v2_8_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_9 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_6)->{'value1'});
$__tco_10 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_7)->{'value1'});
$__tco_var_go__go_5_5_5_v_6 = $__tco_9;
$__tco_var_go__go_5_5_5_v1_7 = $__tco_10;
goto tco_loop_go__go_5_5_5;;
$__t8 = null;
goto end_branch_8;;
};
$__t8 = $v2_8_7;
end_branch_8:;
$__t5 = $__t8;
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
  $__t6 = (($go__go_5_5)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($x_2)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($y_3)->{'value1'}));
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_2) use ($eqNonEmpty2_1_0) {
  $__num = \func_num_args();
  $__res = $eqNonEmpty2_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($v_2) use ($ordNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($ordNonEmpty1_1_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($ordNonEmpty1_1_0)->{'compare'})(\Data\Lazy\majData_majLazy_force($v_2)))(\Data\Lazy\majData_majLazy_force($v1_3));
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
}, "Eq10" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_eq1NonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_ordList
function majData_majList_majLazy_majTypes_ordmajList($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_ordmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqList1_1_0 = (object)["eq" => function($xs_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($__local_var_1_0, $xs_2) {
  $__num = \func_num_args();
  $go__go_4_1 = null;
  $go__go_4_1 = function($v_5) use ($__local_var_1_0, &$go__go_4_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_1_0, &$go__go_4_1, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t2 = $v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil;
goto end_branch_2;;
};
  $__t2 = ($v_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && (((($__local_var_1_0)->{'eq'})(($v_5)->{'value0'}))(($v1_6)->{'value0'}) && (($go__go_4_1)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_5)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_6)->{'value1'})))));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (($go__go_4_1)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_2)))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($xs_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($dictOrd_0, $xs_2) {
  $__num = \func_num_args();
  $go__go_4_4 = null;
  $go__go_4_4 = (function() use ($dictOrd_0, &$go__go_4_4) {
  $__fn = function($v_5, $v1_6 = null) use ($dictOrd_0, &$go__go_4_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_4_4_v_5 = $v_5;
  $__tco_var_go__go_4_4_4_v1_6 = $v1_6;
  tco_loop_go__go_4_4_4:;
  $v_5 = $__tco_var_go__go_4_4_4_v_5;
  $v1_6 = $__tco_var_go__go_4_4_4_v1_6;
  $__t4 = null;;
  if ($v_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = null;;
if ($v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if (($v_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && $v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons)) {
$v2_7_6 = ((($dictOrd_0)->{'compare'})(($v_5)->{'value0'}))(($v1_6)->{'value0'});
$__t7 = null;;
if ($v2_7_6 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_8 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_5)->{'value1'});
$__tco_9 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_6)->{'value1'});
$__tco_var_go__go_4_4_4_v_5 = $__tco_8;
$__tco_var_go__go_4_4_4_v1_6 = $__tco_9;
goto tco_loop_go__go_4_4_4;;
$__t7 = null;
goto end_branch_7;;
};
$__t7 = $v2_7_6;
end_branch_7:;
$__t4 = $__t7;
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
  $__res = (($go__go_4_4)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_2)))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_2) use ($eqList1_1_0) {
  $__num = \func_num_args();
  $__res = $eqList1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_ordList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_ordmajList';

// Data_List_Lazy_Types_ordNonEmptyList
function majData_majList_majLazy_majTypes_ordmajNonmajEmptymajList($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_ordmajNonmajEmptymajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqNonEmpty2_1_0 = (object)["eq" => function($x_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($__local_var_1_0, $x_2) {
  $__num = \func_num_args();
  $go__go_4_1 = null;
  $go__go_4_1 = function($v_5) use ($__local_var_1_0, &$go__go_4_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_1_0, &$go__go_4_1, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t2 = $v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil;
goto end_branch_2;;
};
  $__t2 = ($v_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($v1_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && (((($__local_var_1_0)->{'eq'})(($v_5)->{'value0'}))(($v1_6)->{'value0'}) && (($go__go_4_1)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_5)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_6)->{'value1'})))));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (((($__local_var_1_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && (($go__go_4_1)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($x_2)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($y_3)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_4 = (object)["compare" => function($x_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictOrd_0, $x_2) {
  $__num = \func_num_args();
  $v_4_4 = ((($dictOrd_0)->{'compare'})(($x_2)->{'value0'}))(($y_3)->{'value0'});
  $__t6 = null;;
  if ($v_4_4 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t6 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_6;;
};
  if ($v_4_4 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t6 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_6;;
};
  $go__go_5_5 = null;
  $go__go_5_5 = (function() use ($dictOrd_0, &$go__go_5_5) {
  $__fn = function($v_6, $v1_7 = null) use ($dictOrd_0, &$go__go_5_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_5_5_v_6 = $v_6;
  $__tco_var_go__go_5_5_5_v1_7 = $v1_7;
  tco_loop_go__go_5_5_5:;
  $v_6 = $__tco_var_go__go_5_5_5_v_6;
  $v1_7 = $__tco_var_go__go_5_5_5_v1_7;
  $__t5 = null;;
  if ($v_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t6 = null;;
if ($v1_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t6 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
  if ($v1_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  if (($v_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && $v1_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons)) {
$v2_8_7 = ((($dictOrd_0)->{'compare'})(($v_6)->{'value0'}))(($v1_7)->{'value0'});
$__t8 = null;;
if ($v2_8_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_9 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_6)->{'value1'});
$__tco_10 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_7)->{'value1'});
$__tco_var_go__go_5_5_5_v_6 = $__tco_9;
$__tco_var_go__go_5_5_5_v1_7 = $__tco_10;
goto tco_loop_go__go_5_5_5;;
$__t8 = null;
goto end_branch_8;;
};
$__t8 = $v2_8_7;
end_branch_8:;
$__t5 = $__t8;
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
  $__t6 = (($go__go_5_5)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($x_2)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($y_3)->{'value1'}));
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_2) use ($eqNonEmpty2_1_0) {
  $__num = \func_num_args();
  $__res = $eqNonEmpty2_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_8 = (($__local_var_2_4)->{'Eq0'})(null);
  $eqLazy1_3_8 = (object)["eq" => function($x_4) use ($__local_var_3_8) {
  $__num = \func_num_args();
  $__res = function($y_5) use ($__local_var_3_8, $x_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_8)->{'eq'})(\Data\Lazy\majData_majLazy_force($x_4)))(\Data\Lazy\majData_majLazy_force($y_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($x_4) use ($__local_var_2_4) {
  $__num = \func_num_args();
  $__res = function($y_5) use ($__local_var_2_4, $x_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_4)->{'compare'})(\Data\Lazy\majData_majLazy_force($x_4)))(\Data\Lazy\majData_majLazy_force($y_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_4) use ($eqLazy1_3_8) {
  $__num = \func_num_args();
  $__res = $eqLazy1_3_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_ordNonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_ordmajNonmajEmptymajList';

// Data_List_Lazy_Types_cons
function majData_majList_majLazy_majTypes_cons($x_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_cons';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($x_0, $xs_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_0, $xs_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_cons'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_cons';

// Data_List_Lazy_Types_foldableList
$GLOBALS['Data_List_Lazy_Types_foldableList'] = (object)["foldr" => function($op_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($op_0) {
  $__num = \func_num_args();
  $__res = function($xs_2) use ($op_0, $z_1) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})(function($b_3) use ($op_0) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($b_3, $op_0) {
  $__num = \func_num_args();
  $__res = (($op_0)($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($z_1))((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})(function($b_3) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($b_3) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_5) use ($a_4, $b_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_List_Lazy_Types_nil']))($xs_2));
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
}, "foldl" => function($op_0) {
  $__num = \func_num_args();
  $go__go_1_0 = null;
  $go__go_1_0 = (function() use (&$go__go_1_0, $op_0) {
  $__fn = function($b_2, $xs_3 = null) use (&$go__go_1_0, $op_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_0_0_b_2 = $b_2;
  $__tco_var_go__go_1_0_0_xs_3 = $xs_3;
  tco_loop_go__go_1_0_0:;
  $b_2 = $__tco_var_go__go_1_0_0_b_2;
  $xs_3 = $__tco_var_go__go_1_0_0_xs_3;
  $v_4_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_3);
  $__t1 = null;;
  if ($v_4_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $b_2;
goto end_branch_1;;
};
  if ($v_4_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_2 = (($op_0)($b_2))(($v_4_0)->{'value0'});
$__tco_3 = ($v_4_0)->{'value1'};
$__tco_var_go__go_1_0_0_b_2 = $__tco_2;
$__tco_var_go__go_1_0_0_xs_3 = $__tco_3;
goto tco_loop_go__go_1_0_0;;
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
  $__res = $go__go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_2 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_1, $mempty_2_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})(function($b_4) use ($Semigroup0_1_1, $f_3) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($Semigroup0_1_1, $b_4, $f_3) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_1)->{'append'})($b_4))(($f_3)($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_foldableNonEmpty
$GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'] = (object)["foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($f_2) use ($Semigroup0_1_0, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Semigroup0_1_0, $dictMonoid_0, $f_2) {
  $__num = \func_num_args();
  $Semigroup0_4_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $go__go_5_2 = null;
  $go__go_5_2 = (function() use ($Semigroup0_4_1, $f_2, &$go__go_5_2) {
  $__fn = function($b_6, $xs_7 = null) use ($Semigroup0_4_1, $f_2, &$go__go_5_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_2_2_b_6 = $b_6;
  $__tco_var_go__go_5_2_2_xs_7 = $xs_7;
  tco_loop_go__go_5_2_2:;
  $b_6 = $__tco_var_go__go_5_2_2_b_6;
  $xs_7 = $__tco_var_go__go_5_2_2_xs_7;
  $v_8_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_7);
  $__t3 = null;;
  if ($v_8_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t3 = $b_6;
goto end_branch_3;;
};
  if ($v_8_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_4 = ((($Semigroup0_4_1)->{'append'})($b_6))(($f_2)(($v_8_2)->{'value0'}));
$__tco_5 = ($v_8_2)->{'value1'};
$__tco_var_go__go_5_2_2_b_6 = $__tco_4;
$__tco_var_go__go_5_2_2_xs_7 = $__tco_5;
goto tco_loop_go__go_5_2_2;;
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
  $__res = ((($Semigroup0_1_0)->{'append'})(($f_2)(($v_3)->{'value0'})))((($go__go_5_2)(($dictMonoid_0)->{'mempty'}))(($v_3)->{'value1'}));
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
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $go__go_3_3 = null;
  $go__go_3_3 = (function() use ($f_0, &$go__go_3_3) {
  $__fn = function($b_4, $xs_5 = null) use ($f_0, &$go__go_3_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_3_3_b_4 = $b_4;
  $__tco_var_go__go_3_3_3_xs_5 = $xs_5;
  tco_loop_go__go_3_3_3:;
  $b_4 = $__tco_var_go__go_3_3_3_b_4;
  $xs_5 = $__tco_var_go__go_3_3_3_xs_5;
  $v_6_3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_5);
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t4 = $b_4;
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_5 = (($f_0)($b_4))(($v_6_3)->{'value0'});
$__tco_6 = ($v_6_3)->{'value1'};
$__tco_var_go__go_3_3_3_b_4 = $__tco_5;
$__tco_var_go__go_3_3_3_xs_5 = $__tco_6;
goto tco_loop_go__go_3_3_3;;
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
  $__res = (($go__go_3_3)((($f_0)($b_1))(($v_2)->{'value0'})))(($v_2)->{'value1'});
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
}, "foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $go__go_3_4 = null;
  $go__go_3_4 = (function() use ($f_0, &$go__go_3_4) {
  $__fn = function($b_4, $xs_5 = null) use ($f_0, &$go__go_3_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_4_4_b_4 = $b_4;
  $__tco_var_go__go_3_4_4_xs_5 = $xs_5;
  tco_loop_go__go_3_4_4:;
  $b_4 = $__tco_var_go__go_3_4_4_b_4;
  $xs_5 = $__tco_var_go__go_3_4_4_xs_5;
  $v_6_4 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_5);
  $__t5 = null;;
  if ($v_6_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = $b_4;
goto end_branch_5;;
};
  if ($v_6_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_6 = (($f_0)(($v_6_4)->{'value0'}))($b_4);
$__tco_7 = ($v_6_4)->{'value1'};
$__tco_var_go__go_3_4_4_b_4 = $__tco_6;
$__tco_var_go__go_3_4_4_xs_5 = $__tco_7;
goto tco_loop_go__go_3_4_4;;
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
  $go__go_4_5 = null;
  $go__go_4_5 = (function() use (&$go__go_4_5) {
  $__fn = function($b_5, $xs_6 = null) use (&$go__go_4_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_5_5_b_5 = $b_5;
  $__tco_var_go__go_4_5_5_xs_6 = $xs_6;
  tco_loop_go__go_4_5_5:;
  $b_5 = $__tco_var_go__go_4_5_5_b_5;
  $xs_6 = $__tco_var_go__go_4_5_5_xs_6;
  $v_7_5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_6);
  $__t6 = null;;
  if ($v_7_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t6 = $b_5;
goto end_branch_6;;
};
  if ($v_7_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_8_7 = ($v_7_5)->{'value0'};
$__tco_8 = \Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_8_7, $b_5) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_8_7, $b_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__tco_9 = ($v_7_5)->{'value1'};
$__tco_var_go__go_4_5_5_b_5 = $__tco_8;
$__tco_var_go__go_4_5_5_xs_6 = $__tco_9;
goto tco_loop_go__go_4_5_5;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($f_0)(($v_2)->{'value0'}))((($go__go_3_4)($b_1))((($go__go_4_5)($GLOBALS['Data_List_Lazy_Types_nil']))(($v_2)->{'value1'})));
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

// Data_List_Lazy_Types_extendList
$GLOBALS['Data_List_Lazy_Types_extendList'] = (object)["extend" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($l_1) use ($f_0) {
  $__num = \func_num_args();
  $v_2_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($l_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $GLOBALS['Data_List_Lazy_Types_nil'];
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_3_2 = ($f_0)($l_1);
$go__go_4_3 = null;
$go__go_4_3 = (function() use ($f_0, &$go__go_4_3) {
  $__fn = function($b_5, $xs_6 = null) use ($f_0, &$go__go_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_3_3_b_5 = $b_5;
  $__tco_var_go__go_4_3_3_xs_6 = $xs_6;
  tco_loop_go__go_4_3_3:;
  $b_5 = $__tco_var_go__go_4_3_3_b_5;
  $xs_6 = $__tco_var_go__go_4_3_3_xs_6;
  $v_7_3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_6);
  $__t4 = null;;
  if ($v_7_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t4 = $b_5;
goto end_branch_4;;
};
  if ($v_7_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_8_5 = ($b_5)->{'acc'};
$__local_var_9_6 = ($b_5)->{'val'};
$acc_prime__10_7 = \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_8_5, $v_7_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v_7_3)->{'value0'}, $__local_var_8_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_11_8 = ($f_0)($acc_prime__10_7);
$__tco_9 = (object)["val" => \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_11_8, $__local_var_9_6) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_11_8, $__local_var_9_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "acc" => $acc_prime__10_7];
$__tco_10 = ($v_7_3)->{'value1'};
$__tco_var_go__go_4_3_3_b_5 = $__tco_9;
$__tco_var_go__go_4_3_3_xs_6 = $__tco_10;
goto tco_loop_go__go_4_3_3;;
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
$go__go_5_4 = null;
$go__go_5_4 = (function() use (&$go__go_5_4) {
  $__fn = function($b_6, $xs_7 = null) use (&$go__go_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_4_4_b_6 = $b_6;
  $__tco_var_go__go_5_4_4_xs_7 = $xs_7;
  tco_loop_go__go_5_4_4:;
  $b_6 = $__tco_var_go__go_5_4_4_b_6;
  $xs_7 = $__tco_var_go__go_5_4_4_xs_7;
  $v_8_4 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_7);
  $__t5 = null;;
  if ($v_8_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = $b_6;
goto end_branch_5;;
};
  if ($v_8_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_9_6 = ($v_8_4)->{'value0'};
$__tco_7 = \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_9_6, $b_6) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_9_6, $b_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__tco_8 = ($v_8_4)->{'value1'};
$__tco_var_go__go_5_4_4_b_6 = $__tco_7;
$__tco_var_go__go_5_4_4_xs_7 = $__tco_8;
goto tco_loop_go__go_5_4_4;;
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
$__local_var_5_4 = ((($go__go_4_3)((object)["val" => $GLOBALS['Data_List_Lazy_Types_nil'], "acc" => $GLOBALS['Data_List_Lazy_Types_nil']]))((($go__go_5_4)($GLOBALS['Data_List_Lazy_Types_nil']))(($v_2_0)->{'value1'})))->{'val'};
$__t1 = \Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_3_2, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_3_2, $__local_var_5_4);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_extendNonEmptyList
$GLOBALS['Data_List_Lazy_Types_extendNonEmptyList'] = (object)["extend" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (\Data\Lazy\majData_majLazy_force($v_1))->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v2_3) use ($__local_var_2_0, $f_0, $v_1) {
  $__num = \func_num_args();
  $go__go_4_1 = null;
  $go__go_4_1 = (function() use ($f_0, &$go__go_4_1) {
  $__fn = function($b_5, $xs_6 = null) use ($f_0, &$go__go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_1_1_b_5 = $b_5;
  $__tco_var_go__go_4_1_1_xs_6 = $xs_6;
  tco_loop_go__go_4_1_1:;
  $b_5 = $__tco_var_go__go_4_1_1_b_5;
  $xs_6 = $__tco_var_go__go_4_1_1_xs_6;
  $v_7_1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_6);
  $__t2 = null;;
  if ($v_7_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t2 = $b_5;
goto end_branch_2;;
};
  if ($v_7_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_8_3 = ($b_5)->{'acc'};
$__local_var_9_4 = ($b_5)->{'val'};
$__local_var_10_5 = ($f_0)(\Data\Lazy\majData_majLazy_defer(function($v2_10) use ($__local_var_8_3, $v_7_1) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_7_1)->{'value0'}, $__local_var_8_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_6 = (object)["val" => \Data\Lazy\majData_majLazy_defer(function($v_11) use ($__local_var_10_5, $__local_var_9_4) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_10_5, $__local_var_9_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "acc" => \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_8_3, $v_7_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v_7_1)->{'value0'}, $__local_var_8_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
$__tco_7 = ($v_7_1)->{'value1'};
$__tco_var_go__go_4_1_1_b_5 = $__tco_6;
$__tco_var_go__go_4_1_1_xs_6 = $__tco_7;
goto tco_loop_go__go_4_1_1;;
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
  $go__go_5_2 = null;
  $go__go_5_2 = (function() use (&$go__go_5_2) {
  $__fn = function($b_6, $xs_7 = null) use (&$go__go_5_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_2_2_b_6 = $b_6;
  $__tco_var_go__go_5_2_2_xs_7 = $xs_7;
  tco_loop_go__go_5_2_2:;
  $b_6 = $__tco_var_go__go_5_2_2_b_6;
  $xs_7 = $__tco_var_go__go_5_2_2_xs_7;
  $v_8_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_7);
  $__t3 = null;;
  if ($v_8_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t3 = $b_6;
goto end_branch_3;;
};
  if ($v_8_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_9_4 = ($v_8_2)->{'value0'};
$__tco_5 = \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_9_4, $b_6) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_9_4, $b_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__tco_6 = ($v_8_2)->{'value1'};
$__tco_var_go__go_5_2_2_b_6 = $__tco_5;
$__tco_var_go__go_5_2_2_xs_7 = $__tco_6;
goto tco_loop_go__go_5_2_2;;
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
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)($v_1), ((($go__go_4_1)((object)["val" => $GLOBALS['Data_List_Lazy_Types_nil'], "acc" => $GLOBALS['Data_List_Lazy_Types_nil']]))((($go__go_5_2)($GLOBALS['Data_List_Lazy_Types_nil']))($__local_var_2_0)))->{'val'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_foldableNonEmptyList
$GLOBALS['Data_List_Lazy_Types_foldableNonEmptyList'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'])->{'foldr'})($f_0))($b_1))(\Data\Lazy\majData_majLazy_force($v_2));
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
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'])->{'foldl'})($f_0))($b_1))(\Data\Lazy\majData_majLazy_force($v_2));
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictMonoid_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'])->{'foldMap'})($dictMonoid_0))($f_1))(\Data\Lazy\majData_majLazy_force($v_2));
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

// Data_List_Lazy_Types_showList
function majData_majList_majLazy_majTypes_showmajList($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_showmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["show" => function($xs_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $v_2_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = "(fromFoldable [])";
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$go__go_3_2 = null;
$go__go_3_2 = (function() use ($dictShow_0, &$go__go_3_2) {
  $__fn = function($b_4, $xs_5 = null) use ($dictShow_0, &$go__go_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_2_2_b_4 = $b_4;
  $__tco_var_go__go_3_2_2_xs_5 = $xs_5;
  tco_loop_go__go_3_2_2:;
  $b_4 = $__tco_var_go__go_3_2_2_b_4;
  $xs_5 = $__tco_var_go__go_3_2_2_xs_5;
  $v_6_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_5);
  $__t3 = null;;
  if ($v_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t3 = $b_4;
goto end_branch_3;;
};
  if ($v_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_4 = (($b_4 . ",") . (($dictShow_0)->{'show'})(($v_6_2)->{'value0'}));
$__tco_5 = ($v_6_2)->{'value1'};
$__tco_var_go__go_3_2_2_b_4 = $__tco_4;
$__tco_var_go__go_3_2_2_xs_5 = $__tco_5;
goto tco_loop_go__go_3_2_2;;
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
$__t1 = ((("(fromFoldable [" . (($dictShow_0)->{'show'})(($v_2_0)->{'value0'})) . (($go__go_3_2)(""))(($v_2_0)->{'value1'})) . "])");
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_showList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_showmajList';

// Data_List_Lazy_Types_showNonEmptyList
function majData_majList_majLazy_majTypes_showmajNonmajEmptymajList($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_showmajNonmajEmptymajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (object)["show" => function($xs_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $v_2_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = "(fromFoldable [])";
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$go__go_3_2 = null;
$go__go_3_2 = (function() use ($dictShow_0, &$go__go_3_2) {
  $__fn = function($b_4, $xs_5 = null) use ($dictShow_0, &$go__go_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_2_2_b_4 = $b_4;
  $__tco_var_go__go_3_2_2_xs_5 = $xs_5;
  tco_loop_go__go_3_2_2:;
  $b_4 = $__tco_var_go__go_3_2_2_b_4;
  $xs_5 = $__tco_var_go__go_3_2_2_xs_5;
  $v_6_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_5);
  $__t3 = null;;
  if ($v_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t3 = $b_4;
goto end_branch_3;;
};
  if ($v_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_4 = (($b_4 . ",") . (($dictShow_0)->{'show'})(($v_6_2)->{'value0'}));
$__tco_5 = ($v_6_2)->{'value1'};
$__tco_var_go__go_3_2_2_b_4 = $__tco_4;
$__tco_var_go__go_3_2_2_xs_5 = $__tco_5;
goto tco_loop_go__go_3_2_2;;
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
$__t1 = ((("(fromFoldable [" . (($dictShow_0)->{'show'})(($v_2_0)->{'value0'})) . (($go__go_3_2)(""))(($v_2_0)->{'value1'})) . "])");
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
  $__local_var_2_4 = (object)["show" => function($v_2) use ($__local_var_1_0, $dictShow_0) {
  $__num = \func_num_args();
  $__res = (((("(NonEmpty " . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . " ") . (($__local_var_1_0)->{'show'})(($v_2)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $showLazy_1_0 = (object)["show" => function($x_3) use ($__local_var_2_4) {
  $__num = \func_num_args();
  $__res = (("(defer \\_ -> " . (($__local_var_2_4)->{'show'})(\Data\Lazy\majData_majLazy_force($x_3))) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["show" => function($v_2) use ($showLazy_1_0) {
  $__num = \func_num_args();
  $__res = (("(NonEmptyList " . (($showLazy_1_0)->{'show'})($v_2)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_showNonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_showmajNonmajEmptymajList';

// Data_List_Lazy_Types_showStep
function majData_majList_majLazy_majTypes_showmajStep($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_showmajStep';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $showList1_1_0 = (object)["show" => function($xs_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $v_2_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = "(fromFoldable [])";
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$go__go_3_2 = null;
$go__go_3_2 = (function() use ($dictShow_0, &$go__go_3_2) {
  $__fn = function($b_4, $xs_5 = null) use ($dictShow_0, &$go__go_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_2_2_b_4 = $b_4;
  $__tco_var_go__go_3_2_2_xs_5 = $xs_5;
  tco_loop_go__go_3_2_2:;
  $b_4 = $__tco_var_go__go_3_2_2_b_4;
  $xs_5 = $__tco_var_go__go_3_2_2_xs_5;
  $v_6_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_5);
  $__t3 = null;;
  if ($v_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t3 = $b_4;
goto end_branch_3;;
};
  if ($v_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_4 = (($b_4 . ",") . (($dictShow_0)->{'show'})(($v_6_2)->{'value0'}));
$__tco_5 = ($v_6_2)->{'value1'};
$__tco_var_go__go_3_2_2_b_4 = $__tco_4;
$__tco_var_go__go_3_2_2_xs_5 = $__tco_5;
goto tco_loop_go__go_3_2_2;;
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
$__t1 = ((("(fromFoldable [" . (($dictShow_0)->{'show'})(($v_2_0)->{'value0'})) . (($go__go_3_2)(""))(($v_2_0)->{'value1'})) . "])");
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
  $__res = (object)["show" => function($v_2) use ($dictShow_0, $showList1_1_0) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t4 = "Nil";
goto end_branch_4;;
};
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t4 = (((("(" . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . " : ") . (($showList1_1_0)->{'show'})(($v_2)->{'value1'})) . ")");
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_showStep'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_showmajStep';

// Data_List_Lazy_Types_foldableWithIndexList
$GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($xs_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = (function() use (&$go__go_3_0) {
  $__fn = function($b_4, $xs_5 = null) use (&$go__go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_0_0_b_4 = $b_4;
  $__tco_var_go__go_3_0_0_xs_5 = $xs_5;
  tco_loop_go__go_3_0_0:;
  $b_4 = $__tco_var_go__go_3_0_0_b_4;
  $xs_5 = $__tco_var_go__go_3_0_0_xs_5;
  $v_6_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_5);
  $__t1 = null;;
  if ($v_6_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $b_4;
goto end_branch_1;;
};
  if ($v_6_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_7_2 = ($b_4)->{'value1'};
$__local_var_8_3 = ($v_6_0)->{'value0'};
$__tco_4 = new \Data\Tuple\Data_Tuple_Tuple((($b_4)->{'value0'} + 1), \Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_7_2, $__local_var_8_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_8_3, $__local_var_7_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_5 = ($v_6_0)->{'value1'};
$__tco_var_go__go_3_0_0_b_4 = $__tco_4;
$__tco_var_go__go_3_0_0_xs_5 = $__tco_5;
goto tco_loop_go__go_3_0_0;;
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
  $v_3_0 = (($go__go_3_0)(new \Data\Tuple\Data_Tuple_Tuple(0, $GLOBALS['Data_List_Lazy_Types_nil'])))($xs_2);
  $go__go_4_2 = null;
  $go__go_4_2 = (function() use ($f_0, &$go__go_4_2) {
  $__fn = function($b_5, $xs_6 = null) use ($f_0, &$go__go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_2_2_b_5 = $b_5;
  $__tco_var_go__go_4_2_2_xs_6 = $xs_6;
  tco_loop_go__go_4_2_2:;
  $b_5 = $__tco_var_go__go_4_2_2_b_5;
  $xs_6 = $__tco_var_go__go_4_2_2_xs_6;
  $v_7_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_6);
  $__t3 = null;;
  if ($v_7_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t3 = $b_5;
goto end_branch_3;;
};
  if ($v_7_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_4 = new \Data\Tuple\Data_Tuple_Tuple((($b_5)->{'value0'} - 1), ((($f_0)((($b_5)->{'value0'} - 1)))(($v_7_2)->{'value0'}))(($b_5)->{'value1'}));
$__tco_5 = ($v_7_2)->{'value1'};
$__tco_var_go__go_4_2_2_b_5 = $__tco_4;
$__tco_var_go__go_4_2_2_xs_6 = $__tco_5;
goto tco_loop_go__go_4_2_2;;
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
  $__res = ((($go__go_4_2)(new \Data\Tuple\Data_Tuple_Tuple(($v_3_0)->{'value0'}, $b_1)))(($v_3_0)->{'value1'}))->{'value1'};
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($acc_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_3 = null;
  $go__go_2_3 = (function() use ($f_0, &$go__go_2_3) {
  $__fn = function($b_3, $xs_4 = null) use ($f_0, &$go__go_2_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_3_3_b_3 = $b_3;
  $__tco_var_go__go_2_3_3_xs_4 = $xs_4;
  tco_loop_go__go_2_3_3:;
  $b_3 = $__tco_var_go__go_2_3_3_b_3;
  $xs_4 = $__tco_var_go__go_2_3_3_xs_4;
  $v_5_3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_4);
  $__t4 = null;;
  if ($v_5_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t4 = $b_3;
goto end_branch_4;;
};
  if ($v_5_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_5 = new \Data\Tuple\Data_Tuple_Tuple((($b_3)->{'value0'} + 1), ((($f_0)(($b_3)->{'value0'}))(($b_3)->{'value1'}))(($v_5_3)->{'value0'}));
$__tco_6 = ($v_5_3)->{'value1'};
$__tco_var_go__go_2_3_3_b_3 = $__tco_5;
$__tco_var_go__go_2_3_3_xs_4 = $__tco_6;
goto tco_loop_go__go_2_3_3;;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))(($go__go_2_3)(new \Data\Tuple\Data_Tuple_Tuple(0, $acc_1)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_4 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_5 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_4, $mempty_2_5) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'])->{'foldlWithIndex'})(function($i_4) use ($Semigroup0_1_4, $f_3) {
  $__num = \func_num_args();
  $__res = function($acc_5) use ($Semigroup0_1_4, $f_3, $i_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Semigroup0_1_4)->{'append'})($acc_5)))(($f_3)($i_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_foldableWithIndexNonEmptyList
$GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmptyList'] = (object)["foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictMonoid_0, $f_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_1))(function($v2_3) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
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
});
  $__local_var_4_2 = \Data\Lazy\majData_majLazy_force($v_2);
  $Semigroup0_5_3 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_6_4 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_3_0))($GLOBALS['Data_Maybe_Just']);
  $go__go_7_5 = null;
  $go__go_7_5 = (function() use ($Semigroup0_5_3, $__local_var_6_4, &$go__go_7_5) {
  $__fn = function($b_8, $xs_9 = null) use ($Semigroup0_5_3, $__local_var_6_4, &$go__go_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_5_5_b_8 = $b_8;
  $__tco_var_go__go_7_5_5_xs_9 = $xs_9;
  tco_loop_go__go_7_5_5:;
  $b_8 = $__tco_var_go__go_7_5_5_b_8;
  $xs_9 = $__tco_var_go__go_7_5_5_xs_9;
  $v_10_5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_9);
  $__t6 = null;;
  if ($v_10_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t6 = $b_8;
goto end_branch_6;;
};
  if ($v_10_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_7 = new \Data\Tuple\Data_Tuple_Tuple((($b_8)->{'value0'} + 1), \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_5_3)->{'append'})(($b_8)->{'value1'}), ($__local_var_6_4)(($b_8)->{'value0'}), ($v_10_5)->{'value0'}));
$__tco_8 = ($v_10_5)->{'value1'};
$__tco_var_go__go_7_5_5_b_8 = $__tco_7;
$__tco_var_go__go_7_5_5_xs_9 = $__tco_8;
goto tco_loop_go__go_7_5_5;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})((($__local_var_3_0)(new \Data\Maybe\Data_Maybe_Nothing()))(($__local_var_4_2)->{'value0'})))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Tuple_snd'], ($go__go_7_5)(new \Data\Tuple\Data_Tuple_Tuple(0, ($dictMonoid_0)->{'mempty'})), ($__local_var_4_2)->{'value1'}));
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__local_var_3_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = 0;
goto end_branch_6;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = (1 + ($v2_3)->{'value0'});
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_4_8 = \Data\Lazy\majData_majLazy_force($v_2);
  $__local_var_5_9 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_3_6))($GLOBALS['Data_Maybe_Just']);
  $go__go_6_10 = null;
  $go__go_6_10 = (function() use ($__local_var_5_9, &$go__go_6_10) {
  $__fn = function($b_7, $xs_8 = null) use ($__local_var_5_9, &$go__go_6_10, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_10_10_b_7 = $b_7;
  $__tco_var_go__go_6_10_10_xs_8 = $xs_8;
  tco_loop_go__go_6_10_10:;
  $b_7 = $__tco_var_go__go_6_10_10_b_7;
  $xs_8 = $__tco_var_go__go_6_10_10_xs_8;
  $v_9_10 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_8);
  $__t11 = null;;
  if ($v_9_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t11 = $b_7;
goto end_branch_11;;
};
  if ($v_9_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_12 = new \Data\Tuple\Data_Tuple_Tuple((($b_7)->{'value0'} + 1), ((($__local_var_5_9)(($b_7)->{'value0'}))(($b_7)->{'value1'}))(($v_9_10)->{'value0'}));
$__tco_13 = ($v_9_10)->{'value1'};
$__tco_var_go__go_6_10_10_b_7 = $__tco_12;
$__tco_var_go__go_6_10_10_xs_8 = $__tco_13;
goto tco_loop_go__go_6_10_10;;
$__t11 = null;
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Tuple_snd'], ($go__go_6_10)(new \Data\Tuple\Data_Tuple_Tuple(0, ((($__local_var_3_6)(new \Data\Maybe\Data_Maybe_Nothing()))($b_1))(($__local_var_4_8)->{'value0'}))), ($__local_var_4_8)->{'value1'});
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
}, "foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__local_var_3_11 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = 0;
goto end_branch_11;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = (1 + ($v2_3)->{'value0'});
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_4_13 = \Data\Lazy\majData_majLazy_force($v_2);
  $__local_var_5_14 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_3_11))($GLOBALS['Data_Maybe_Just']);
  $go__go_6_15 = null;
  $go__go_6_15 = (function() use (&$go__go_6_15) {
  $__fn = function($b_7, $xs_8 = null) use (&$go__go_6_15, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_15_15_b_7 = $b_7;
  $__tco_var_go__go_6_15_15_xs_8 = $xs_8;
  tco_loop_go__go_6_15_15:;
  $b_7 = $__tco_var_go__go_6_15_15_b_7;
  $xs_8 = $__tco_var_go__go_6_15_15_xs_8;
  $v_9_15 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_8);
  $__t16 = null;;
  if ($v_9_15 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t16 = $b_7;
goto end_branch_16;;
};
  if ($v_9_15 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_10_17 = ($b_7)->{'value1'};
$__local_var_11_18 = ($v_9_15)->{'value0'};
$__tco_19 = new \Data\Tuple\Data_Tuple_Tuple((($b_7)->{'value0'} + 1), \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_10_17, $__local_var_11_18) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_11_18, $__local_var_10_17);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_20 = ($v_9_15)->{'value1'};
$__tco_var_go__go_6_15_15_b_7 = $__tco_19;
$__tco_var_go__go_6_15_15_xs_8 = $__tco_20;
goto tco_loop_go__go_6_15_15;;
$__t16 = null;
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $v_6_15 = (($go__go_6_15)(new \Data\Tuple\Data_Tuple_Tuple(0, $GLOBALS['Data_List_Lazy_Types_nil'])))(($__local_var_4_13)->{'value1'});
  $go__go_7_17 = null;
  $go__go_7_17 = (function() use ($__local_var_5_14, &$go__go_7_17) {
  $__fn = function($b_8, $xs_9 = null) use ($__local_var_5_14, &$go__go_7_17, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_17_17_b_8 = $b_8;
  $__tco_var_go__go_7_17_17_xs_9 = $xs_9;
  tco_loop_go__go_7_17_17:;
  $b_8 = $__tco_var_go__go_7_17_17_b_8;
  $xs_9 = $__tco_var_go__go_7_17_17_xs_9;
  $v_10_17 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_9);
  $__t18 = null;;
  if ($v_10_17 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t18 = $b_8;
goto end_branch_18;;
};
  if ($v_10_17 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_19 = new \Data\Tuple\Data_Tuple_Tuple((($b_8)->{'value0'} - 1), ((($__local_var_5_14)((($b_8)->{'value0'} - 1)))(($v_10_17)->{'value0'}))(($b_8)->{'value1'}));
$__tco_20 = ($v_10_17)->{'value1'};
$__tco_var_go__go_7_17_17_b_8 = $__tco_19;
$__tco_var_go__go_7_17_17_xs_9 = $__tco_20;
goto tco_loop_go__go_7_17_17;;
$__t18 = null;
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($__local_var_3_11)(new \Data\Maybe\Data_Maybe_Nothing()))(($__local_var_4_13)->{'value0'}))(((($go__go_7_17)(new \Data\Tuple\Data_Tuple_Tuple(($v_6_15)->{'value0'}, $b_1)))(($v_6_15)->{'value1'}))->{'value1'});
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_functorWithIndexList
$GLOBALS['Data_List_Lazy_Types_functorWithIndexList'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($xs_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0) {
  $__fn = function($b_3, $xs_4 = null) use (&$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_xs_4 = $xs_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $xs_4 = $__tco_var_go__go_2_0_0_xs_4;
  $v_5_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_4);
  $__t1 = null;;
  if ($v_5_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $b_3;
goto end_branch_1;;
};
  if ($v_5_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_6_2 = ($b_3)->{'value1'};
$__local_var_7_3 = ($v_5_0)->{'value0'};
$__tco_4 = new \Data\Tuple\Data_Tuple_Tuple((($b_3)->{'value0'} + 1), \Data\Lazy\majData_majLazy_defer(function($v_8) use ($__local_var_6_2, $__local_var_7_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_7_3, $__local_var_6_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_5 = ($v_5_0)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_4;
$__tco_var_go__go_2_0_0_xs_4 = $__tco_5;
goto tco_loop_go__go_2_0_0;;
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
  $v_2_0 = (($go__go_2_0)(new \Data\Tuple\Data_Tuple_Tuple(0, $GLOBALS['Data_List_Lazy_Types_nil'])))($xs_1);
  $go__go_3_2 = null;
  $go__go_3_2 = (function() use ($f_0, &$go__go_3_2) {
  $__fn = function($b_4, $xs_5 = null) use ($f_0, &$go__go_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_2_2_b_4 = $b_4;
  $__tco_var_go__go_3_2_2_xs_5 = $xs_5;
  tco_loop_go__go_3_2_2:;
  $b_4 = $__tco_var_go__go_3_2_2_b_4;
  $xs_5 = $__tco_var_go__go_3_2_2_xs_5;
  $v_6_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_5);
  $__t3 = null;;
  if ($v_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t3 = $b_4;
goto end_branch_3;;
};
  if ($v_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_7_4 = ($b_4)->{'value1'};
$__local_var_8_5 = (($f_0)((($b_4)->{'value0'} - 1)))(($v_6_2)->{'value0'});
$__tco_6 = new \Data\Tuple\Data_Tuple_Tuple((($b_4)->{'value0'} - 1), \Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_7_4, $__local_var_8_5) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_8_5, $__local_var_7_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_7 = ($v_6_2)->{'value1'};
$__tco_var_go__go_3_2_2_b_4 = $__tco_6;
$__tco_var_go__go_3_2_2_xs_5 = $__tco_7;
goto tco_loop_go__go_3_2_2;;
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
  $__res = ((($go__go_3_2)(new \Data\Tuple\Data_Tuple_Tuple(($v_2_0)->{'value0'}, $GLOBALS['Data_List_Lazy_Types_nil'])))(($v_2_0)->{'value1'}))->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_functorWithIndexNonEmptyList
$GLOBALS['Data_List_Lazy_Types_functorWithIndexNonEmptyList'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v1_2) use ($f_0, $v_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
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
});
  $__local_var_4_2 = \Data\Lazy\majData_majLazy_force($v_1);
  $__local_var_5_3 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_3_0))($GLOBALS['Data_Maybe_Just']);
  $go__go_6_4 = null;
  $go__go_6_4 = (function() use (&$go__go_6_4) {
  $__fn = function($b_7, $xs_8 = null) use (&$go__go_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_4_4_b_7 = $b_7;
  $__tco_var_go__go_6_4_4_xs_8 = $xs_8;
  tco_loop_go__go_6_4_4:;
  $b_7 = $__tco_var_go__go_6_4_4_b_7;
  $xs_8 = $__tco_var_go__go_6_4_4_xs_8;
  $v_9_4 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_8);
  $__t5 = null;;
  if ($v_9_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = $b_7;
goto end_branch_5;;
};
  if ($v_9_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_10_6 = ($b_7)->{'value1'};
$__local_var_11_7 = ($v_9_4)->{'value0'};
$__tco_8 = new \Data\Tuple\Data_Tuple_Tuple((($b_7)->{'value0'} + 1), \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_10_6, $__local_var_11_7) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_11_7, $__local_var_10_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_9 = ($v_9_4)->{'value1'};
$__tco_var_go__go_6_4_4_b_7 = $__tco_8;
$__tco_var_go__go_6_4_4_xs_8 = $__tco_9;
goto tco_loop_go__go_6_4_4;;
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
  $v_6_4 = (($go__go_6_4)(new \Data\Tuple\Data_Tuple_Tuple(0, $GLOBALS['Data_List_Lazy_Types_nil'])))(($__local_var_4_2)->{'value1'});
  $go__go_7_6 = null;
  $go__go_7_6 = (function() use ($__local_var_5_3, &$go__go_7_6) {
  $__fn = function($b_8, $xs_9 = null) use ($__local_var_5_3, &$go__go_7_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_6_6_b_8 = $b_8;
  $__tco_var_go__go_7_6_6_xs_9 = $xs_9;
  tco_loop_go__go_7_6_6:;
  $b_8 = $__tco_var_go__go_7_6_6_b_8;
  $xs_9 = $__tco_var_go__go_7_6_6_xs_9;
  $v_10_6 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_9);
  $__t7 = null;;
  if ($v_10_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t7 = $b_8;
goto end_branch_7;;
};
  if ($v_10_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_11_8 = ($b_8)->{'value1'};
$__local_var_12_9 = (($__local_var_5_3)((($b_8)->{'value0'} - 1)))(($v_10_6)->{'value0'});
$__tco_10 = new \Data\Tuple\Data_Tuple_Tuple((($b_8)->{'value0'} - 1), \Data\Lazy\majData_majLazy_defer(function($v_13) use ($__local_var_11_8, $__local_var_12_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_12_9, $__local_var_11_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_11 = ($v_10_6)->{'value1'};
$__tco_var_go__go_7_6_6_b_8 = $__tco_10;
$__tco_var_go__go_7_6_6_xs_9 = $__tco_11;
goto tco_loop_go__go_7_6_6;;
$__t7 = null;
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty((($__local_var_3_0)(new \Data\Maybe\Data_Maybe_Nothing()))(($__local_var_4_2)->{'value0'}), ((($go__go_7_6)(new \Data\Tuple\Data_Tuple_Tuple(($v_6_4)->{'value0'}, $GLOBALS['Data_List_Lazy_Types_nil'])))(($v_6_4)->{'value1'}))->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_toList
function majData_majList_majLazy_majTypes_tomajList($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_tomajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Data\Lazy\majData_majLazy_defer((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))(function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $v2_2_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_3_1 = ($v2_2_0)->{'value0'};
  $__local_var_4_2 = ($v2_2_0)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_5) use ($__local_var_3_1, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_3_1, $__local_var_4_2);
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_toList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_tomajList';

// Data_List_Lazy_Types_semigroupNonEmptyList
$GLOBALS['Data_List_Lazy_Types_semigroupNonEmptyList'] = (object)["append" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($as_prime__1) use ($v_0) {
  $__num = \func_num_args();
  $v1_2_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_3_1 = ($v1_2_0)->{'value0'};
  $__local_var_4_2 = ($v1_2_0)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v2_5) use ($__local_var_3_1, $__local_var_4_2, $as_prime__1) {
  $__num = \func_num_args();
  $__local_var_6_3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_tomajList($as_prime__1);
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_3_1, \Data\Lazy\majData_majLazy_defer(function($v_7) use ($__local_var_4_2, $__local_var_6_3) {
  $__num = \func_num_args();
  $__local_var_8_4 = \Data\Lazy\majData_majLazy_force($__local_var_4_2);
  $__t5 = null;;
  if ($__local_var_8_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_6_3);
goto end_branch_5;;
};
  if ($__local_var_8_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t5 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_8_4)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_8_4)->{'value1'}))($__local_var_6_3));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_traversableList
$GLOBALS['Data_List_Lazy_Types_traversableList'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_2_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_3) use ($Apply0_1_0, $Functor0_2_1, $dictApplicative_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictApplicative_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil']);
  $__res = function($xs_5) use ($Apply0_1_0, $Functor0_2_1, $__local_var_4_2, $f_3) {
  $__num = \func_num_args();
  $go__go_6_3 = null;
  $go__go_6_3 = (function() use ($Apply0_1_0, $Functor0_2_1, $f_3, &$go__go_6_3) {
  $__fn = function($b_7, $xs_8 = null) use ($Apply0_1_0, $Functor0_2_1, $f_3, &$go__go_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_3_3_b_7 = $b_7;
  $__tco_var_go__go_6_3_3_xs_8 = $xs_8;
  tco_loop_go__go_6_3_3:;
  $b_7 = $__tco_var_go__go_6_3_3_b_7;
  $xs_8 = $__tco_var_go__go_6_3_3_xs_8;
  $v_9_3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_8);
  $__t4 = null;;
  if ($v_9_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t4 = $b_7;
goto end_branch_4;;
};
  if ($v_9_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_5 = ((($Apply0_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})($GLOBALS['Data_List_Lazy_Types_cons']))(($f_3)(($v_9_3)->{'value0'}))))($b_7);
$__tco_6 = ($v_9_3)->{'value1'};
$__tco_var_go__go_6_3_3_b_7 = $__tco_5;
$__tco_var_go__go_6_3_3_xs_8 = $__tco_6;
goto tco_loop_go__go_6_3_3;;
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
  $go__go_7_4 = null;
  $go__go_7_4 = (function() use (&$go__go_7_4) {
  $__fn = function($b_8, $xs_9 = null) use (&$go__go_7_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_4_4_b_8 = $b_8;
  $__tco_var_go__go_7_4_4_xs_9 = $xs_9;
  tco_loop_go__go_7_4_4:;
  $b_8 = $__tco_var_go__go_7_4_4_b_8;
  $xs_9 = $__tco_var_go__go_7_4_4_xs_9;
  $v_10_4 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_9);
  $__t5 = null;;
  if ($v_10_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = $b_8;
goto end_branch_5;;
};
  if ($v_10_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_11_6 = ($v_10_4)->{'value0'};
$__tco_7 = \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_11_6, $b_8) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_11_6, $b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__tco_8 = ($v_10_4)->{'value1'};
$__tco_var_go__go_7_4_4_b_8 = $__tco_7;
$__tco_var_go__go_7_4_4_xs_9 = $__tco_8;
goto tco_loop_go__go_7_4_4;;
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
  $__res = (($go__go_6_3)($__local_var_4_2))((($go__go_7_4)($GLOBALS['Data_List_Lazy_Types_nil']))($xs_5));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_traversableList'])->{'traverse'})($dictApplicative_0))(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_traversableNonEmptyList
$GLOBALS['Data_List_Lazy_Types_traversableNonEmptyList'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $__local_var_4_1 = \Data\Lazy\majData_majLazy_force($v_3);
  $Apply0_5_2 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_6_3 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $go__go_7_4 = null;
  $go__go_7_4 = (function() use ($Apply0_5_2, $Functor0_6_3, $f_2, &$go__go_7_4) {
  $__fn = function($b_8, $xs_9 = null) use ($Apply0_5_2, $Functor0_6_3, $f_2, &$go__go_7_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_4_4_b_8 = $b_8;
  $__tco_var_go__go_7_4_4_xs_9 = $xs_9;
  tco_loop_go__go_7_4_4:;
  $b_8 = $__tco_var_go__go_7_4_4_b_8;
  $xs_9 = $__tco_var_go__go_7_4_4_xs_9;
  $v_10_4 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_9);
  $__t5 = null;;
  if ($v_10_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = $b_8;
goto end_branch_5;;
};
  if ($v_10_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_6 = ((($Apply0_5_2)->{'apply'})(((($Functor0_6_3)->{'map'})($GLOBALS['Data_List_Lazy_Types_cons']))(($f_2)(($v_10_4)->{'value0'}))))($b_8);
$__tco_7 = ($v_10_4)->{'value1'};
$__tco_var_go__go_7_4_4_b_8 = $__tco_6;
$__tco_var_go__go_7_4_4_xs_9 = $__tco_7;
goto tco_loop_go__go_7_4_4;;
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
  $go__go_8_5 = null;
  $go__go_8_5 = (function() use (&$go__go_8_5) {
  $__fn = function($b_9, $xs_10 = null) use (&$go__go_8_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_5_5_b_9 = $b_9;
  $__tco_var_go__go_8_5_5_xs_10 = $xs_10;
  tco_loop_go__go_8_5_5:;
  $b_9 = $__tco_var_go__go_8_5_5_b_9;
  $xs_10 = $__tco_var_go__go_8_5_5_xs_10;
  $v_11_5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_10);
  $__t6 = null;;
  if ($v_11_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t6 = $b_9;
goto end_branch_6;;
};
  if ($v_11_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_12_7 = ($v_11_5)->{'value0'};
$__tco_8 = \Data\Lazy\majData_majLazy_defer(function($v_13) use ($__local_var_12_7, $b_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_12_7, $b_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__tco_9 = ($v_11_5)->{'value1'};
$__tco_var_go__go_8_5_5_b_9 = $__tco_8;
$__tco_var_go__go_8_5_5_xs_10 = $__tco_9;
goto tco_loop_go__go_8_5_5;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Functor0_1_0)->{'map'})(function($xxs_4) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v1_5) use ($xxs_4) {
  $__num = \func_num_args();
  $__res = $xxs_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((((($dictApplicative_0)->{'Apply0'})(null))->{'apply'})(((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($f_2)(($__local_var_4_1)->{'value0'}))))((($go__go_7_4)((($dictApplicative_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil'])))((($go__go_8_5)($GLOBALS['Data_List_Lazy_Types_nil']))(($__local_var_4_1)->{'value1'}))));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_6 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_6, $dictApplicative_0) {
  $__num = \func_num_args();
  $__local_var_3_7 = \Data\Lazy\majData_majLazy_force($v_2);
  $Apply0_4_8 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_5_9 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $go__go_6_10 = null;
  $go__go_6_10 = (function() use ($Apply0_4_8, $Functor0_5_9, &$go__go_6_10) {
  $__fn = function($b_7, $xs_8 = null) use ($Apply0_4_8, $Functor0_5_9, &$go__go_6_10, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_10_10_b_7 = $b_7;
  $__tco_var_go__go_6_10_10_xs_8 = $xs_8;
  tco_loop_go__go_6_10_10:;
  $b_7 = $__tco_var_go__go_6_10_10_b_7;
  $xs_8 = $__tco_var_go__go_6_10_10_xs_8;
  $v_9_10 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_8);
  $__t11 = null;;
  if ($v_9_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t11 = $b_7;
goto end_branch_11;;
};
  if ($v_9_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_12 = ((($Apply0_4_8)->{'apply'})(((($Functor0_5_9)->{'map'})($GLOBALS['Data_List_Lazy_Types_cons']))(($v_9_10)->{'value0'})))($b_7);
$__tco_13 = ($v_9_10)->{'value1'};
$__tco_var_go__go_6_10_10_b_7 = $__tco_12;
$__tco_var_go__go_6_10_10_xs_8 = $__tco_13;
goto tco_loop_go__go_6_10_10;;
$__t11 = null;
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_7_11 = null;
  $go__go_7_11 = (function() use (&$go__go_7_11) {
  $__fn = function($b_8, $xs_9 = null) use (&$go__go_7_11, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_11_11_b_8 = $b_8;
  $__tco_var_go__go_7_11_11_xs_9 = $xs_9;
  tco_loop_go__go_7_11_11:;
  $b_8 = $__tco_var_go__go_7_11_11_b_8;
  $xs_9 = $__tco_var_go__go_7_11_11_xs_9;
  $v_10_11 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_9);
  $__t12 = null;;
  if ($v_10_11 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t12 = $b_8;
goto end_branch_12;;
};
  if ($v_10_11 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_11_13 = ($v_10_11)->{'value0'};
$__tco_14 = \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_11_13, $b_8) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_11_13, $b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__tco_15 = ($v_10_11)->{'value1'};
$__tco_var_go__go_7_11_11_b_8 = $__tco_14;
$__tco_var_go__go_7_11_11_xs_9 = $__tco_15;
goto tco_loop_go__go_7_11_11;;
$__t12 = null;
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Functor0_1_6)->{'map'})(function($xxs_3) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v1_4) use ($xxs_3) {
  $__num = \func_num_args();
  $__res = $xxs_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((((($dictApplicative_0)->{'Apply0'})(null))->{'apply'})(((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($__local_var_3_7)->{'value0'})))((($go__go_6_10)((($dictApplicative_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil'])))((($go__go_7_11)($GLOBALS['Data_List_Lazy_Types_nil']))(($__local_var_3_7)->{'value1'}))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_traversableWithIndexList
$GLOBALS['Data_List_Lazy_Types_traversableWithIndexList'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_2_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_3) use ($Apply0_1_0, $Functor0_2_1, $dictApplicative_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictApplicative_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil']);
  $__res = function($xs_5) use ($Apply0_1_0, $Functor0_2_1, $__local_var_4_2, $f_3) {
  $__num = \func_num_args();
  $go__go_6_3 = null;
  $go__go_6_3 = (function() use (&$go__go_6_3) {
  $__fn = function($b_7, $xs_8 = null) use (&$go__go_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_3_3_b_7 = $b_7;
  $__tco_var_go__go_6_3_3_xs_8 = $xs_8;
  tco_loop_go__go_6_3_3:;
  $b_7 = $__tco_var_go__go_6_3_3_b_7;
  $xs_8 = $__tco_var_go__go_6_3_3_xs_8;
  $v_9_3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_8);
  $__t4 = null;;
  if ($v_9_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t4 = $b_7;
goto end_branch_4;;
};
  if ($v_9_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_10_5 = ($b_7)->{'value1'};
$__local_var_11_6 = ($v_9_3)->{'value0'};
$__tco_7 = new \Data\Tuple\Data_Tuple_Tuple((($b_7)->{'value0'} + 1), \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_10_5, $__local_var_11_6) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_11_6, $__local_var_10_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_8 = ($v_9_3)->{'value1'};
$__tco_var_go__go_6_3_3_b_7 = $__tco_7;
$__tco_var_go__go_6_3_3_xs_8 = $__tco_8;
goto tco_loop_go__go_6_3_3;;
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
  $v_6_3 = (($go__go_6_3)(new \Data\Tuple\Data_Tuple_Tuple(0, $GLOBALS['Data_List_Lazy_Types_nil'])))($xs_5);
  $go__go_7_5 = null;
  $go__go_7_5 = (function() use ($Apply0_1_0, $Functor0_2_1, $f_3, &$go__go_7_5) {
  $__fn = function($b_8, $xs_9 = null) use ($Apply0_1_0, $Functor0_2_1, $f_3, &$go__go_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_5_5_b_8 = $b_8;
  $__tco_var_go__go_7_5_5_xs_9 = $xs_9;
  tco_loop_go__go_7_5_5:;
  $b_8 = $__tco_var_go__go_7_5_5_b_8;
  $xs_9 = $__tco_var_go__go_7_5_5_xs_9;
  $v_10_5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_9);
  $__t6 = null;;
  if ($v_10_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t6 = $b_8;
goto end_branch_6;;
};
  if ($v_10_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_7 = new \Data\Tuple\Data_Tuple_Tuple((($b_8)->{'value0'} - 1), ((($Apply0_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})($GLOBALS['Data_List_Lazy_Types_cons']))((($f_3)((($b_8)->{'value0'} - 1)))(($v_10_5)->{'value0'}))))(($b_8)->{'value1'}));
$__tco_8 = ($v_10_5)->{'value1'};
$__tco_var_go__go_7_5_5_b_8 = $__tco_7;
$__tco_var_go__go_7_5_5_xs_9 = $__tco_8;
goto tco_loop_go__go_7_5_5;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($go__go_7_5)(new \Data\Tuple\Data_Tuple_Tuple(($v_6_3)->{'value0'}, $__local_var_4_2)))(($v_6_3)->{'value1'}))->{'value1'};
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorWithIndexList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_traversableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_traversableWithIndexNonEmptyList
$GLOBALS['Data_List_Lazy_Types_traversableWithIndexNonEmptyList'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $__local_var_4_1 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))(function($v2_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v2_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($v2_4 instanceof \Data\Maybe\Data_Maybe_Just) {
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
});
  $__local_var_5_3 = \Data\Lazy\majData_majLazy_force($v_3);
  $Apply0_6_4 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_7_5 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_8_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_4_1))($GLOBALS['Data_Maybe_Just']);
  $go__go_9_7 = null;
  $go__go_9_7 = (function() use (&$go__go_9_7) {
  $__fn = function($b_10, $xs_11 = null) use (&$go__go_9_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_9_7_7_b_10 = $b_10;
  $__tco_var_go__go_9_7_7_xs_11 = $xs_11;
  tco_loop_go__go_9_7_7:;
  $b_10 = $__tco_var_go__go_9_7_7_b_10;
  $xs_11 = $__tco_var_go__go_9_7_7_xs_11;
  $v_12_7 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_11);
  $__t8 = null;;
  if ($v_12_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t8 = $b_10;
goto end_branch_8;;
};
  if ($v_12_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_13_9 = ($b_10)->{'value1'};
$__local_var_14_10 = ($v_12_7)->{'value0'};
$__tco_11 = new \Data\Tuple\Data_Tuple_Tuple((($b_10)->{'value0'} + 1), \Data\Lazy\majData_majLazy_defer(function($v_15) use ($__local_var_13_9, $__local_var_14_10) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_14_10, $__local_var_13_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__tco_12 = ($v_12_7)->{'value1'};
$__tco_var_go__go_9_7_7_b_10 = $__tco_11;
$__tco_var_go__go_9_7_7_xs_11 = $__tco_12;
goto tco_loop_go__go_9_7_7;;
$__t8 = null;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $v_9_7 = (($go__go_9_7)(new \Data\Tuple\Data_Tuple_Tuple(0, $GLOBALS['Data_List_Lazy_Types_nil'])))(($__local_var_5_3)->{'value1'});
  $go__go_10_9 = null;
  $go__go_10_9 = (function() use ($Apply0_6_4, $Functor0_7_5, $__local_var_8_6, &$go__go_10_9) {
  $__fn = function($b_11, $xs_12 = null) use ($Apply0_6_4, $Functor0_7_5, $__local_var_8_6, &$go__go_10_9, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_10_9_9_b_11 = $b_11;
  $__tco_var_go__go_10_9_9_xs_12 = $xs_12;
  tco_loop_go__go_10_9_9:;
  $b_11 = $__tco_var_go__go_10_9_9_b_11;
  $xs_12 = $__tco_var_go__go_10_9_9_xs_12;
  $v_13_9 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_12);
  $__t10 = null;;
  if ($v_13_9 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t10 = $b_11;
goto end_branch_10;;
};
  if ($v_13_9 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_11 = new \Data\Tuple\Data_Tuple_Tuple((($b_11)->{'value0'} - 1), ((($Apply0_6_4)->{'apply'})(((($Functor0_7_5)->{'map'})($GLOBALS['Data_List_Lazy_Types_cons']))((($__local_var_8_6)((($b_11)->{'value0'} - 1)))(($v_13_9)->{'value0'}))))(($b_11)->{'value1'}));
$__tco_12 = ($v_13_9)->{'value1'};
$__tco_var_go__go_10_9_9_b_11 = $__tco_11;
$__tco_var_go__go_10_9_9_xs_12 = $__tco_12;
goto tco_loop_go__go_10_9_9;;
$__t10 = null;
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Functor0_1_0)->{'map'})(function($xxs_4) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v1_5) use ($xxs_4) {
  $__num = \func_num_args();
  $__res = $xxs_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((((($dictApplicative_0)->{'Apply0'})(null))->{'apply'})(((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))((($__local_var_4_1)(new \Data\Maybe\Data_Maybe_Nothing()))(($__local_var_5_3)->{'value0'}))))(((($go__go_10_9)(new \Data\Tuple\Data_Tuple_Tuple(($v_9_7)->{'value0'}, (($dictApplicative_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil']))))(($v_9_7)->{'value1'}))->{'value1'}));
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorWithIndexNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_traversableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_unfoldable1List
$GLOBALS['Data_List_Lazy_Types_unfoldable1List'] = (function() use (&$__fn) {
$go__go_0_0 = null;
$go__go_0_0 = function($f_1) use (&$go__go_0_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($f_1, &$go__go_0_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))(function($v_3) use ($b_2, $f_1, &$go__go_0_0) {
  $__num = \func_num_args();
  $v1_4_1 = ($f_1)($b_2);
  $__t2 = null;;
  if (($v1_4_1)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_5_3 = ($v1_4_1)->{'value0'};
$__local_var_6_4 = (($go__go_0_0)($f_1))((($v1_4_1)->{'value1'})->{'value0'});
$__t2 = \Data\Lazy\majData_majLazy_defer(function($v_7) use ($__local_var_5_3, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_5_3, $__local_var_6_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_2;;
};
  if (($v1_4_1)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__local_var_5_5 = ($v1_4_1)->{'value0'};
$__t2 = \Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_5_5) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_5_5, $GLOBALS['Data_List_Lazy_Types_nil']);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
return (object)["unfoldr1" => $go__go_0_0];
})();

// Data_List_Lazy_Types_unfoldableList
$GLOBALS['Data_List_Lazy_Types_unfoldableList'] = (function() use (&$__fn) {
$go__go_0_0 = null;
$go__go_0_0 = function($f_1) use (&$go__go_0_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($f_1, &$go__go_0_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))(function($v_3) use ($b_2, $f_1, &$go__go_0_0) {
  $__num = \func_num_args();
  $v1_4_1 = ($f_1)($b_2);
  $__t2 = null;;
  if ($v1_4_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = $GLOBALS['Data_List_Lazy_Types_nil'];
goto end_branch_2;;
};
  if ($v1_4_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_5_3 = (($v1_4_1)->{'value0'})->{'value0'};
$__local_var_6_4 = (($go__go_0_0)($f_1))((($v1_4_1)->{'value0'})->{'value1'});
$__t2 = \Data\Lazy\majData_majLazy_defer(function($v_7) use ($__local_var_5_3, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_5_3, $__local_var_6_4);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
return (object)["unfoldr" => $go__go_0_0, "Unfoldable10" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_unfoldable1List'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_List_Lazy_Types_unfoldable1NonEmpty
$GLOBALS['Data_List_Lazy_Types_unfoldable1NonEmpty'] = (object)["unfoldr1" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = ($f_0)($b_1);
  $go__go_3_1 = null;
  $go__go_3_1 = function($f_4) use (&$go__go_3_1) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($f_4, &$go__go_3_1) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_step']))(function($v_6) use ($b_5, $f_4, &$go__go_3_1) {
  $__num = \func_num_args();
  $v1_7_2 = ($f_4)($b_5);
  $__t3 = null;;
  if ($v1_7_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = $GLOBALS['Data_List_Lazy_Types_nil'];
goto end_branch_3;;
};
  if ($v1_7_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_8_4 = (($v1_7_2)->{'value0'})->{'value0'};
$__local_var_9_5 = (($go__go_3_1)($f_4))((($v1_7_2)->{'value0'})->{'value1'});
$__t3 = \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_8_4, $__local_var_9_5) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_8_4, $__local_var_9_5);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_2_0 = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_2_0)->{'value0'}, (($go__go_3_1)(function($v1_4) use ($f_0) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($v1_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = new \Data\Maybe\Data_Maybe_Just(($f_0)(($v1_4)->{'value0'}));
goto end_branch_6;;
};
  $__t6 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($__local_var_2_0)->{'value1'}));
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($__local_var_2_0)->{'value0'}, ($__local_var_2_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_unfoldable1NonEmptyList
$GLOBALS['Data_List_Lazy_Types_unfoldable1NonEmptyList'] = (object)["unfoldr1" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_unfoldable1NonEmpty'])->{'unfoldr1'})($f_0))($b_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_comonadNonEmptyList
$GLOBALS['Data_List_Lazy_Types_comonadNonEmptyList'] = (object)["extract" => function($v_0) {
  $__num = \func_num_args();
  $__res = (\Data\Lazy\majData_majLazy_force($v_0))->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_extendNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_monadList
$GLOBALS['Data_List_Lazy_Types_monadList'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applicativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_bindList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_bindList
$GLOBALS['Data_List_Lazy_Types_bindList'] = (object)["bind" => function($xs_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($xs_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($f_1, $xs_0) {
  $__num = \func_num_args();
  $__local_var_3_0 = \Data\Lazy\majData_majLazy_force($xs_0);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_4_2 = ($f_1)(($__local_var_3_0)->{'value0'});
$__local_var_5_3 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_3_0)->{'value1'}))($f_1);
$__t1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_4_2, $__local_var_5_3) {
  $__num = \func_num_args();
  $__local_var_7_4 = \Data\Lazy\majData_majLazy_force($__local_var_4_2);
  $__t5 = null;;
  if ($__local_var_7_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_5_3);
goto end_branch_5;;
};
  if ($__local_var_7_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t5 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_7_4)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_7_4)->{'value1'}))($__local_var_5_3));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_applyList
$GLOBALS['Data_List_Lazy_Types_applyList'] = (object)["apply" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($a_1) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = function($f_prime__2) use ($a_1) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_3) use ($a_1, $f_prime__2) {
  $__num = \func_num_args();
  $__local_var_4_0 = \Data\Lazy\majData_majLazy_force($a_1);
  $__t1 = null;;
  if ($__local_var_4_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_1;;
};
  if ($__local_var_4_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_5_2 = (($GLOBALS['Data_List_Lazy_Types_applicativeList'])->{'pure'})(($f_prime__2)(($__local_var_4_0)->{'value0'}));
$__local_var_6_3 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_4_0)->{'value1'}))(function($a_prime__6) use ($f_prime__2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_Lazy_Types_applicativeList'])->{'pure'})(($f_prime__2)($a_prime__6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_7) use ($__local_var_5_2, $__local_var_6_3) {
  $__num = \func_num_args();
  $__local_var_8_4 = \Data\Lazy\majData_majLazy_force($__local_var_5_2);
  $__t5 = null;;
  if ($__local_var_8_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_6_3);
goto end_branch_5;;
};
  if ($__local_var_8_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_9_6 = ($__local_var_8_4)->{'value1'};
$__t5 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_8_4)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_6_3, $__local_var_9_6) {
  $__num = \func_num_args();
  $__local_var_11_7 = \Data\Lazy\majData_majLazy_force($__local_var_9_6);
  $__t8 = null;;
  if ($__local_var_11_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t8 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_6_3);
goto end_branch_8;;
};
  if ($__local_var_11_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t8 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_11_7)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_11_7)->{'value1'}))($__local_var_6_3));
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_3) use ($__local_var_2_0, $f_0) {
  $__num = \func_num_args();
  $__local_var_4_10 = \Data\Lazy\majData_majLazy_force($f_0);
  $__t11 = null;;
  if ($__local_var_4_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t11 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_11;;
};
  if ($__local_var_4_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_5_12 = ($__local_var_2_0)(($__local_var_4_10)->{'value0'});
$__local_var_6_13 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_4_10)->{'value1'}))($__local_var_2_0);
$__t11 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_7) use ($__local_var_5_12, $__local_var_6_13) {
  $__num = \func_num_args();
  $__local_var_8_14 = \Data\Lazy\majData_majLazy_force($__local_var_5_12);
  $__t15 = null;;
  if ($__local_var_8_14 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t15 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_6_13);
goto end_branch_15;;
};
  if ($__local_var_8_14 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_9_16 = ($__local_var_8_14)->{'value1'};
$__t15 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_8_14)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_6_13, $__local_var_9_16) {
  $__num = \func_num_args();
  $__local_var_11_17 = \Data\Lazy\majData_majLazy_force($__local_var_9_16);
  $__t18 = null;;
  if ($__local_var_11_17 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t18 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_6_13);
goto end_branch_18;;
};
  if ($__local_var_11_17 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t18 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_11_17)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_11_17)->{'value1'}))($__local_var_6_13));
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_applicativeList
$GLOBALS['Data_List_Lazy_Types_applicativeList'] = (object)["pure" => function($a_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_1) use ($a_0) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_0, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_applyNonEmptyList
$GLOBALS['Data_List_Lazy_Types_applyNonEmptyList'] = (object)["apply" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $v2_2_0 = \Data\Lazy\majData_majLazy_force($v1_1);
  $v3_3_1 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_4_2 = ($v2_2_0)->{'value0'};
  $__local_var_5_3 = ($v2_2_0)->{'value1'};
  $__local_var_6_4 = ($v3_3_1)->{'value0'};
  $__local_var_7_5 = ($v3_3_1)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v4_8) use ($__local_var_4_2, $__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__num = \func_num_args();
  $__local_var_9_6 = \Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_4_2, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_10_7 = function($f_prime__10) use ($__local_var_9_6) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_11) use ($__local_var_9_6, $f_prime__10) {
  $__num = \func_num_args();
  $__local_var_12_7 = \Data\Lazy\majData_majLazy_force($__local_var_9_6);
  $__t8 = null;;
  if ($__local_var_12_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t8 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_8;;
};
  if ($__local_var_12_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_13_9 = ($f_prime__10)(($__local_var_12_7)->{'value0'});
$__local_var_13_9 = \Data\Lazy\majData_majLazy_defer(function($v_14) use ($__local_var_13_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_13_9, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_14_11 = ($__local_var_12_7)->{'value1'};
$__local_var_14_11 = \Data\Lazy\majData_majLazy_defer(function($v_15) use ($__local_var_14_11, $f_prime__10) {
  $__num = \func_num_args();
  $__local_var_16_12 = \Data\Lazy\majData_majLazy_force($__local_var_14_11);
  $__t13 = null;;
  if ($__local_var_16_12 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t13 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_13;;
};
  if ($__local_var_16_12 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_17_14 = ($f_prime__10)(($__local_var_16_12)->{'value0'});
$__local_var_17_14 = \Data\Lazy\majData_majLazy_defer(function($v_18) use ($__local_var_17_14) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_17_14, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_18_16 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_16_12)->{'value1'}))(function($a_prime__18) use ($f_prime__10) {
  $__num = \func_num_args();
  $__local_var_19_16 = ($f_prime__10)($a_prime__18);
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_20) use ($__local_var_19_16) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_19_16, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t13 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_19) use ($__local_var_17_14, $__local_var_18_16) {
  $__num = \func_num_args();
  $__local_var_20_18 = \Data\Lazy\majData_majLazy_force($__local_var_17_14);
  $__t19 = null;;
  if ($__local_var_20_18 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t19 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_18_16);
goto end_branch_19;;
};
  if ($__local_var_20_18 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_21_20 = ($__local_var_20_18)->{'value1'};
$__t19 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_20_18)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_22) use ($__local_var_18_16, $__local_var_21_20) {
  $__num = \func_num_args();
  $__local_var_23_21 = \Data\Lazy\majData_majLazy_force($__local_var_21_20);
  $__t22 = null;;
  if ($__local_var_23_21 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t22 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_18_16);
goto end_branch_22;;
};
  if ($__local_var_23_21 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t22 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_23_21)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_23_21)->{'value1'}))($__local_var_18_16));
goto end_branch_22;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t22 = null;
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_19;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t19 = null;
  end_branch_19:;
  $__res = $__t19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t8 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_15) use ($__local_var_13_9, $__local_var_14_11) {
  $__num = \func_num_args();
  $__local_var_16_24 = \Data\Lazy\majData_majLazy_force($__local_var_13_9);
  $__t25 = null;;
  if ($__local_var_16_24 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t25 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_14_11);
goto end_branch_25;;
};
  if ($__local_var_16_24 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_17_26 = ($__local_var_16_24)->{'value1'};
$__t25 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_16_24)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_18) use ($__local_var_14_11, $__local_var_17_26) {
  $__num = \func_num_args();
  $__local_var_19_27 = \Data\Lazy\majData_majLazy_force($__local_var_17_26);
  $__t28 = null;;
  if ($__local_var_19_27 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t28 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_14_11);
goto end_branch_28;;
};
  if ($__local_var_19_27 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_20_29 = ($__local_var_19_27)->{'value1'};
$__t28 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_19_27)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_21) use ($__local_var_14_11, $__local_var_20_29) {
  $__num = \func_num_args();
  $__local_var_22_30 = \Data\Lazy\majData_majLazy_force($__local_var_20_29);
  $__t31 = null;;
  if ($__local_var_22_30 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t31 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_14_11);
goto end_branch_31;;
};
  if ($__local_var_22_30 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t31 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_22_30)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_22_30)->{'value1'}))($__local_var_14_11));
goto end_branch_31;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t31 = null;
  end_branch_31:;
  $__res = $__t31;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_28;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t28 = null;
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_9_6 = \Data\Lazy\majData_majLazy_defer(function($v_11) use ($__local_var_10_7, $__local_var_7_5) {
  $__num = \func_num_args();
  $__local_var_12_33 = \Data\Lazy\majData_majLazy_force($__local_var_7_5);
  $__t34 = null;;
  if ($__local_var_12_33 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t34 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_34;;
};
  if ($__local_var_12_33 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_13_35 = ($__local_var_10_7)(($__local_var_12_33)->{'value0'});
$__local_var_14_36 = ($__local_var_12_33)->{'value1'};
$__local_var_14_36 = \Data\Lazy\majData_majLazy_defer(function($v_15) use ($__local_var_10_7, $__local_var_14_36) {
  $__num = \func_num_args();
  $__local_var_16_37 = \Data\Lazy\majData_majLazy_force($__local_var_14_36);
  $__t38 = null;;
  if ($__local_var_16_37 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t38 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_38;;
};
  if ($__local_var_16_37 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_17_39 = ($__local_var_10_7)(($__local_var_16_37)->{'value0'});
$__local_var_18_40 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_16_37)->{'value1'}))($__local_var_10_7);
$__t38 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_19) use ($__local_var_17_39, $__local_var_18_40) {
  $__num = \func_num_args();
  $__local_var_20_41 = \Data\Lazy\majData_majLazy_force($__local_var_17_39);
  $__t42 = null;;
  if ($__local_var_20_41 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t42 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_18_40);
goto end_branch_42;;
};
  if ($__local_var_20_41 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_21_43 = ($__local_var_20_41)->{'value1'};
$__t42 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_20_41)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_22) use ($__local_var_18_40, $__local_var_21_43) {
  $__num = \func_num_args();
  $__local_var_23_44 = \Data\Lazy\majData_majLazy_force($__local_var_21_43);
  $__t45 = null;;
  if ($__local_var_23_44 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t45 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_18_40);
goto end_branch_45;;
};
  if ($__local_var_23_44 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t45 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_23_44)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_23_44)->{'value1'}))($__local_var_18_40));
goto end_branch_45;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t45 = null;
  end_branch_45:;
  $__res = $__t45;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_42;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t42 = null;
  end_branch_42:;
  $__res = $__t42;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_38;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t38 = null;
  end_branch_38:;
  $__res = $__t38;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t34 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_15) use ($__local_var_13_35, $__local_var_14_36) {
  $__num = \func_num_args();
  $__local_var_16_47 = \Data\Lazy\majData_majLazy_force($__local_var_13_35);
  $__t48 = null;;
  if ($__local_var_16_47 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t48 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_14_36);
goto end_branch_48;;
};
  if ($__local_var_16_47 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_17_49 = ($__local_var_16_47)->{'value1'};
$__t48 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_16_47)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_18) use ($__local_var_14_36, $__local_var_17_49) {
  $__num = \func_num_args();
  $__local_var_19_50 = \Data\Lazy\majData_majLazy_force($__local_var_17_49);
  $__t51 = null;;
  if ($__local_var_19_50 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t51 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_14_36);
goto end_branch_51;;
};
  if ($__local_var_19_50 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_20_52 = ($__local_var_19_50)->{'value1'};
$__t51 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_19_50)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_21) use ($__local_var_14_36, $__local_var_20_52) {
  $__num = \func_num_args();
  $__local_var_22_53 = \Data\Lazy\majData_majLazy_force($__local_var_20_52);
  $__t54 = null;;
  if ($__local_var_22_53 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t54 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_14_36);
goto end_branch_54;;
};
  if ($__local_var_22_53 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t54 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_22_53)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_22_53)->{'value1'}))($__local_var_14_36));
goto end_branch_54;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t54 = null;
  end_branch_54:;
  $__res = $__t54;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_51;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t51 = null;
  end_branch_51:;
  $__res = $__t51;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_48;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t48 = null;
  end_branch_48:;
  $__res = $__t48;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_34;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t34 = null;
  end_branch_34:;
  $__res = $__t34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_10_56 = \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_6_4, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_6_4, $__local_var_7_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_11_57 = function($f_prime__11) use ($__local_var_5_3) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_5_3, $f_prime__11) {
  $__num = \func_num_args();
  $__local_var_13_57 = \Data\Lazy\majData_majLazy_force($__local_var_5_3);
  $__t58 = null;;
  if ($__local_var_13_57 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t58 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_58;;
};
  if ($__local_var_13_57 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_14_59 = ($f_prime__11)(($__local_var_13_57)->{'value0'});
$__local_var_14_59 = \Data\Lazy\majData_majLazy_defer(function($v_15) use ($__local_var_14_59) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_14_59, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_15_61 = ($__local_var_13_57)->{'value1'};
$__local_var_15_61 = \Data\Lazy\majData_majLazy_defer(function($v_16) use ($__local_var_15_61, $f_prime__11) {
  $__num = \func_num_args();
  $__local_var_17_62 = \Data\Lazy\majData_majLazy_force($__local_var_15_61);
  $__t63 = null;;
  if ($__local_var_17_62 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t63 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_63;;
};
  if ($__local_var_17_62 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_18_64 = ($f_prime__11)(($__local_var_17_62)->{'value0'});
$__local_var_18_64 = \Data\Lazy\majData_majLazy_defer(function($v_19) use ($__local_var_18_64) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_18_64, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_19_66 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_17_62)->{'value1'}))(function($a_prime__19) use ($f_prime__11) {
  $__num = \func_num_args();
  $__local_var_20_66 = ($f_prime__11)($a_prime__19);
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_21) use ($__local_var_20_66) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_20_66, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t63 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_20) use ($__local_var_18_64, $__local_var_19_66) {
  $__num = \func_num_args();
  $__local_var_21_68 = \Data\Lazy\majData_majLazy_force($__local_var_18_64);
  $__t69 = null;;
  if ($__local_var_21_68 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t69 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_19_66);
goto end_branch_69;;
};
  if ($__local_var_21_68 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_22_70 = ($__local_var_21_68)->{'value1'};
$__t69 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_21_68)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_23) use ($__local_var_19_66, $__local_var_22_70) {
  $__num = \func_num_args();
  $__local_var_24_71 = \Data\Lazy\majData_majLazy_force($__local_var_22_70);
  $__t72 = null;;
  if ($__local_var_24_71 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t72 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_19_66);
goto end_branch_72;;
};
  if ($__local_var_24_71 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t72 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_24_71)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_24_71)->{'value1'}))($__local_var_19_66));
goto end_branch_72;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t72 = null;
  end_branch_72:;
  $__res = $__t72;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_69;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t69 = null;
  end_branch_69:;
  $__res = $__t69;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_63;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t63 = null;
  end_branch_63:;
  $__res = $__t63;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t58 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_16) use ($__local_var_14_59, $__local_var_15_61) {
  $__num = \func_num_args();
  $__local_var_17_74 = \Data\Lazy\majData_majLazy_force($__local_var_14_59);
  $__t75 = null;;
  if ($__local_var_17_74 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t75 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_15_61);
goto end_branch_75;;
};
  if ($__local_var_17_74 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_18_76 = ($__local_var_17_74)->{'value1'};
$__t75 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_17_74)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_19) use ($__local_var_15_61, $__local_var_18_76) {
  $__num = \func_num_args();
  $__local_var_20_77 = \Data\Lazy\majData_majLazy_force($__local_var_18_76);
  $__t78 = null;;
  if ($__local_var_20_77 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t78 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_15_61);
goto end_branch_78;;
};
  if ($__local_var_20_77 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_21_79 = ($__local_var_20_77)->{'value1'};
$__t78 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_20_77)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_22) use ($__local_var_15_61, $__local_var_21_79) {
  $__num = \func_num_args();
  $__local_var_23_80 = \Data\Lazy\majData_majLazy_force($__local_var_21_79);
  $__t81 = null;;
  if ($__local_var_23_80 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t81 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_15_61);
goto end_branch_81;;
};
  if ($__local_var_23_80 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t81 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_23_80)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_23_80)->{'value1'}))($__local_var_15_61));
goto end_branch_81;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t81 = null;
  end_branch_81:;
  $__res = $__t81;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_78;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t78 = null;
  end_branch_78:;
  $__res = $__t78;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_75;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t75 = null;
  end_branch_75:;
  $__res = $__t75;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_58;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t58 = null;
  end_branch_58:;
  $__res = $__t58;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_10_56 = \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_10_56, $__local_var_11_57) {
  $__num = \func_num_args();
  $__local_var_13_83 = \Data\Lazy\majData_majLazy_force($__local_var_10_56);
  $__t84 = null;;
  if ($__local_var_13_83 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t84 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_84;;
};
  if ($__local_var_13_83 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_14_85 = ($__local_var_11_57)(($__local_var_13_83)->{'value0'});
$__local_var_15_86 = ($__local_var_13_83)->{'value1'};
$__local_var_15_86 = \Data\Lazy\majData_majLazy_defer(function($v_16) use ($__local_var_11_57, $__local_var_15_86) {
  $__num = \func_num_args();
  $__local_var_17_87 = \Data\Lazy\majData_majLazy_force($__local_var_15_86);
  $__t88 = null;;
  if ($__local_var_17_87 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t88 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_88;;
};
  if ($__local_var_17_87 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_18_89 = ($__local_var_11_57)(($__local_var_17_87)->{'value0'});
$__local_var_19_90 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_17_87)->{'value1'}))($__local_var_11_57);
$__t88 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_20) use ($__local_var_18_89, $__local_var_19_90) {
  $__num = \func_num_args();
  $__local_var_21_91 = \Data\Lazy\majData_majLazy_force($__local_var_18_89);
  $__t92 = null;;
  if ($__local_var_21_91 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t92 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_19_90);
goto end_branch_92;;
};
  if ($__local_var_21_91 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_22_93 = ($__local_var_21_91)->{'value1'};
$__t92 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_21_91)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_23) use ($__local_var_19_90, $__local_var_22_93) {
  $__num = \func_num_args();
  $__local_var_24_94 = \Data\Lazy\majData_majLazy_force($__local_var_22_93);
  $__t95 = null;;
  if ($__local_var_24_94 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t95 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_19_90);
goto end_branch_95;;
};
  if ($__local_var_24_94 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t95 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_24_94)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_24_94)->{'value1'}))($__local_var_19_90));
goto end_branch_95;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t95 = null;
  end_branch_95:;
  $__res = $__t95;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_92;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t92 = null;
  end_branch_92:;
  $__res = $__t92;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_88;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t88 = null;
  end_branch_88:;
  $__res = $__t88;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t84 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_16) use ($__local_var_14_85, $__local_var_15_86) {
  $__num = \func_num_args();
  $__local_var_17_97 = \Data\Lazy\majData_majLazy_force($__local_var_14_85);
  $__t98 = null;;
  if ($__local_var_17_97 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t98 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_15_86);
goto end_branch_98;;
};
  if ($__local_var_17_97 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_18_99 = ($__local_var_17_97)->{'value1'};
$__t98 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_17_97)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_19) use ($__local_var_15_86, $__local_var_18_99) {
  $__num = \func_num_args();
  $__local_var_20_100 = \Data\Lazy\majData_majLazy_force($__local_var_18_99);
  $__t101 = null;;
  if ($__local_var_20_100 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t101 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_15_86);
goto end_branch_101;;
};
  if ($__local_var_20_100 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_21_102 = ($__local_var_20_100)->{'value1'};
$__t101 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_20_100)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_22) use ($__local_var_15_86, $__local_var_21_102) {
  $__num = \func_num_args();
  $__local_var_23_103 = \Data\Lazy\majData_majLazy_force($__local_var_21_102);
  $__t104 = null;;
  if ($__local_var_23_103 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t104 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_15_86);
goto end_branch_104;;
};
  if ($__local_var_23_103 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t104 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_23_103)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_23_103)->{'value1'}))($__local_var_15_86));
goto end_branch_104;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t104 = null;
  end_branch_104:;
  $__res = $__t104;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_101;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t101 = null;
  end_branch_101:;
  $__res = $__t101;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_98;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t98 = null;
  end_branch_98:;
  $__res = $__t98;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_84;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t84 = null;
  end_branch_84:;
  $__res = $__t84;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($__local_var_6_4)($__local_var_4_2), \Data\Lazy\majData_majLazy_defer(function($v_11) use ($__local_var_10_56, $__local_var_9_6) {
  $__num = \func_num_args();
  $__local_var_12_106 = \Data\Lazy\majData_majLazy_force($__local_var_9_6);
  $__t107 = null;;
  if ($__local_var_12_106 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t107 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_10_56);
goto end_branch_107;;
};
  if ($__local_var_12_106 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t107 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_12_106)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_12_106)->{'value1'}))($__local_var_10_56));
goto end_branch_107;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t107 = null;
  end_branch_107:;
  $__res = $__t107;
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_bindNonEmptyList
$GLOBALS['Data_List_Lazy_Types_bindNonEmptyList'] = (object)["bind" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($v_0) {
  $__num = \func_num_args();
  $v1_2_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_3_1 = ($v1_2_0)->{'value1'};
  $v2_4_2 = \Data\Lazy\majData_majLazy_force(($f_1)(($v1_2_0)->{'value0'}));
  $__local_var_5_3 = ($v2_4_2)->{'value0'};
  $__local_var_6_4 = ($v2_4_2)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v3_7) use ($__local_var_3_1, $__local_var_5_3, $__local_var_6_4, $f_1) {
  $__num = \func_num_args();
  $__local_var_8_5 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_toList']))($f_1);
  $__local_var_8_5 = \Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_3_1, $__local_var_8_5) {
  $__num = \func_num_args();
  $__local_var_10_6 = \Data\Lazy\majData_majLazy_force($__local_var_3_1);
  $__t7 = null;;
  if ($__local_var_10_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t7 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_7;;
};
  if ($__local_var_10_6 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_11_8 = ($__local_var_8_5)(($__local_var_10_6)->{'value0'});
$__local_var_12_9 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_10_6)->{'value1'}))($__local_var_8_5);
$__t7 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_13) use ($__local_var_11_8, $__local_var_12_9) {
  $__num = \func_num_args();
  $__local_var_14_10 = \Data\Lazy\majData_majLazy_force($__local_var_11_8);
  $__t11 = null;;
  if ($__local_var_14_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t11 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_12_9);
goto end_branch_11;;
};
  if ($__local_var_14_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_15_12 = ($__local_var_14_10)->{'value1'};
$__t11 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_14_10)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_16) use ($__local_var_12_9, $__local_var_15_12) {
  $__num = \func_num_args();
  $__local_var_17_13 = \Data\Lazy\majData_majLazy_force($__local_var_15_12);
  $__t14 = null;;
  if ($__local_var_17_13 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t14 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_12_9);
goto end_branch_14;;
};
  if ($__local_var_17_13 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t14 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_17_13)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_17_13)->{'value1'}))($__local_var_12_9));
goto end_branch_14;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t14 = null;
  end_branch_14:;
  $__res = $__t14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_5_3, \Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_6_4, $__local_var_8_5) {
  $__num = \func_num_args();
  $__local_var_10_16 = \Data\Lazy\majData_majLazy_force($__local_var_6_4);
  $__t17 = null;;
  if ($__local_var_10_16 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t17 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_8_5);
goto end_branch_17;;
};
  if ($__local_var_10_16 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t17 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_10_16)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_10_16)->{'value1'}))($__local_var_8_5));
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applyNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_altNonEmptyList
$GLOBALS['Data_List_Lazy_Types_altNonEmptyList'] = (object)["alt" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($as_prime__1) use ($v_0) {
  $__num = \func_num_args();
  $v1_2_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_3_1 = ($v1_2_0)->{'value0'};
  $__local_var_4_2 = ($v1_2_0)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v2_5) use ($__local_var_3_1, $__local_var_4_2, $as_prime__1) {
  $__num = \func_num_args();
  $__local_var_6_3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_tomajList($as_prime__1);
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_3_1, \Data\Lazy\majData_majLazy_defer(function($v_7) use ($__local_var_4_2, $__local_var_6_3) {
  $__num = \func_num_args();
  $__local_var_8_4 = \Data\Lazy\majData_majLazy_force($__local_var_4_2);
  $__t5 = null;;
  if ($__local_var_8_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t5 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_6_3);
goto end_branch_5;;
};
  if ($__local_var_8_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_9_6 = ($__local_var_8_4)->{'value1'};
$__t5 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_8_4)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_6_3, $__local_var_9_6) {
  $__num = \func_num_args();
  $__local_var_11_7 = \Data\Lazy\majData_majLazy_force($__local_var_9_6);
  $__t8 = null;;
  if ($__local_var_11_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t8 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_6_3);
goto end_branch_8;;
};
  if ($__local_var_11_7 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t8 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_11_7)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_11_7)->{'value1'}))($__local_var_6_3));
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_altList
$GLOBALS['Data_List_Lazy_Types_altList'] = (object)["alt" => function($xs_0) {
  $__num = \func_num_args();
  $__res = function($ys_1) use ($xs_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($xs_0, $ys_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = \Data\Lazy\majData_majLazy_force($xs_0);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_1);
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_3_0)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_3_0)->{'value1'}))($ys_1));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_plusList
$GLOBALS['Data_List_Lazy_Types_plusList'] = (object)["empty" => $GLOBALS['Data_List_Lazy_Types_nil'], "Alt0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_altList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_alternativeList
$GLOBALS['Data_List_Lazy_Types_alternativeList'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applicativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_plusList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_monadPlusList
$GLOBALS['Data_List_Lazy_Types_monadPlusList'] = (object)["Monad0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_monadList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_alternativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_applicativeNonEmptyList
$GLOBALS['Data_List_Lazy_Types_applicativeNonEmptyList'] = (object)["pure" => function($a_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_1) use ($a_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_0, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applyNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_monadNonEmptyList
$GLOBALS['Data_List_Lazy_Types_monadNonEmptyList'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applicativeNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_bindNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

