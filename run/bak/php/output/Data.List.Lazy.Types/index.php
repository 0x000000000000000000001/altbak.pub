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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_List_Lazy_Types_Nil { public $tag = 'Nil'; public function __construct() {} }
final class Data_List_Lazy_Types_Cons { public $tag = 'Cons'; public function __construct(public  $value0, public  $value1) {} }

// Data_List_Lazy_Types_identity
function majData_majList_majLazy_majTypes_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_identity'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_identity';

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
$GLOBALS['Data_List_Lazy_Types_newtypeNonEmptyList'] = (object)["Coercible0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_newtypeList
$GLOBALS['Data_List_Lazy_Types_newtypeList'] = (object)["Coercible0" => function($_dollar__unused_0) {
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
$GLOBALS['Data_List_Lazy_Types_semigroupList'] = (object)["append" => (function() {
  $__fn = function($xs_0, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_2) use ($ys_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_1);
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v_2)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($v_2)->{'value1'}))($ys_1));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Lazy_Types_monoidList
$GLOBALS['Data_List_Lazy_Types_monoidList'] = (object)["mempty" => $GLOBALS['Data_List_Lazy_Types_nil'], "Semigroup0" => function($_dollar__unused_0) {
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
$GLOBALS['Data_List_Lazy_Types_functorList'] = (object)["map" => (function() {
  $__fn = function($f_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_2) use ($f_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($f_0)(($v_2)->{'value0'}), ((($GLOBALS['Data_List_Lazy_Types_functorList'])->{'map'})($f_0))(($v_2)->{'value1'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Lazy_Types_map1
function majData_majList_majLazy_majTypes_map1($f_0, $m_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majTypes_map1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)(($m_1)->{'value0'}), ((($GLOBALS['Data_List_Lazy_Types_functorList'])->{'map'})($f_0))(($m_1)->{'value1'}));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_map1'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_map1';

// Data_List_Lazy_Types_functorNonEmptyList
$GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'] = (object)["map" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(($GLOBALS['Data_List_Lazy_Types_map1'])($f_0)))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Lazy_Types_eq1List
$GLOBALS['Data_List_Lazy_Types_eq1List'] = (object)["eq1" => (function() {
  $__fn = function($dictEq_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = (function() use ($dictEq_0, &$go__3_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictEq_0, &$go__3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil;
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($v1_5 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && (((($dictEq_0)->{'eq'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}) && (($go__3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_4)->{'value1'})))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_5)->{'value1'})))));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_1)))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Lazy_Types_eq1NonEmptyList
$GLOBALS['Data_List_Lazy_Types_eq1NonEmptyList'] = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_eq1Lazy'])->{'eq1'})(\Data\NonEmpty\majData_majNonmajEmpty_eqmajNonmajEmpty($GLOBALS['Data_List_Lazy_Types_eq1List'], $dictEq_0));
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
  $__res = (object)["eq" => (($GLOBALS['Data_List_Lazy_Types_eq1List'])->{'eq1'})($dictEq_0)];
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
  $__local_var_1_0 = \Data\NonEmpty\majData_majNonmajEmpty_eqmajNonmajEmpty($GLOBALS['Data_List_Lazy_Types_eq1List'], $dictEq_0);
  $__res = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'eq'})(\Data\Lazy\majData_majLazy_force($x_2)))(\Data\Lazy\majData_majLazy_force($y_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_eqNonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_eqmajNonmajEmptymajList';

// Data_List_Lazy_Types_ord1List
$GLOBALS['Data_List_Lazy_Types_ord1List'] = (object)["compare1" => (function() {
  $__fn = function($dictOrd_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = (function() use ($dictOrd_0, &$go__3_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictOrd_0, &$go__3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__3_0_0_v_4 = $v_4;
  $__tco_var_go__3_0_0_v1_5 = $v1_5;
  tco_loop_go__3_0_0:;
  $v_4 = $__tco_var_go__3_0_0_v_4;
  $v1_5 = $__tco_var_go__3_0_0_v1_5;
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
$__tco_var_go__3_0_0_v_4 = $__tco_4;
$__tco_var_go__3_0_0_v1_5 = $__tco_5;
goto tco_loop_go__3_0_0;;
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
  $__res = (($go__3_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_1)))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($ys_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Eq10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_eq1List'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_ordNonEmpty
$GLOBALS['Data_List_Lazy_Types_ordNonEmpty'] = ($GLOBALS['Data_NonEmpty_ordNonEmpty'])($GLOBALS['Data_List_Lazy_Types_ord1List']);

// Data_List_Lazy_Types_ord1NonEmptyList
$GLOBALS['Data_List_Lazy_Types_ord1NonEmptyList'] = (object)["compare1" => function($dictOrd_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_ord1Lazy'])->{'compare1'})(($GLOBALS['Data_List_Lazy_Types_ordNonEmpty'])($dictOrd_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_0) {
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
  $eqList1_1_0 = (object)["eq" => (($GLOBALS['Data_List_Lazy_Types_eq1List'])->{'eq1'})((($dictOrd_0)->{'Eq0'})(null))];
  $__res = (object)["compare" => (($GLOBALS['Data_List_Lazy_Types_ord1List'])->{'compare1'})($dictOrd_0), "Eq0" => function($_dollar__unused_2) use ($eqList1_1_0) {
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
  $__res = \Data\Lazy\majData_majLazy_ordmajLazy(($GLOBALS['Data_List_Lazy_Types_ordNonEmpty'])($dictOrd_0));
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
$GLOBALS['Data_List_Lazy_Types_foldableList'] = (object)["foldr" => (function() {
  $__fn = function($op_0, $z_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})((function() use ($op_0) {
  $__fn = function($b_3, $a_4 = null) use ($op_0, &$__fn) {
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
})()))($z_1))((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($b_3, $a_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_5) use ($a_4, $b_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_List_Lazy_Types_nil']))($xs_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => function($op_0) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0, $op_0) {
  $__fn = function($b_2, $xs_3 = null) use (&$go__1_0, $op_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_b_2 = $b_2;
  $__tco_var_go__1_0_0_xs_3 = $xs_3;
  tco_loop_go__1_0_0:;
  $b_2 = $__tco_var_go__1_0_0_b_2;
  $xs_3 = $__tco_var_go__1_0_0_xs_3;
  $v_4_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_3);
  $__t1 = null;;
  if ($v_4_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $b_2;
goto end_branch_1;;
};
  if ($v_4_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_2 = (($op_0)($b_2))(($v_4_0)->{'value0'});
$__tco_3 = ($v_4_0)->{'value1'};
$__tco_var_go__1_0_0_b_2 = $__tco_2;
$__tco_var_go__1_0_0_xs_3 = $__tco_3;
goto tco_loop_go__1_0_0;;
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
  $__res = $go__1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_1 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($b_3, $a_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})($b_3))(($f_2)($a_4));
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
}];

// Data_List_Lazy_Types_foldableNonEmpty
$GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'] = (object)["foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap1_1_0 = (($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldMap'})($dictMonoid_0);
  $__res = (function() use ($dictMonoid_0, $foldMap1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictMonoid_0, $foldMap1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})(($f_2)(($v_3)->{'value0'})))((($foldMap1_1_0)($f_2))(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})($f_0))((($f_0)($b_1))(($v_2)->{'value0'})))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)(($v_2)->{'value0'}))((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldr'})($f_0))($b_1))(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Lazy_Types_extendList
$GLOBALS['Data_List_Lazy_Types_extendList'] = (object)["extend" => (function() {
  $__fn = function($f_0, $l_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($l_1);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = $GLOBALS['Data_List_Lazy_Types_nil'];
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_3_2 = ($f_0)($l_1);
$__local_var_4_3 = ((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldr'})((function() use ($f_0) {
  $__fn = function($a_4, $v_5 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_3 = ($v_5)->{'acc'};
  $__local_var_7_4 = ($v_5)->{'val'};
  $acc_prime_8_5 = \Data\Lazy\majData_majLazy_defer(function($v_8) use ($__local_var_6_3, $a_4) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_4, $__local_var_6_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_9_6 = ($f_0)($acc_prime_8_5);
  $__res = (object)["val" => \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_7_4, $__local_var_9_6) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_9_6, $__local_var_7_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "acc" => $acc_prime_8_5];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["val" => $GLOBALS['Data_List_Lazy_Types_nil'], "acc" => $GLOBALS['Data_List_Lazy_Types_nil']]))(($v_2_0)->{'value1'}))->{'val'};
$__t1 = \Data\Lazy\majData_majLazy_defer(function($v_5) use ($__local_var_3_2, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_3_2, $__local_var_4_3);
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
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_extendNonEmptyList
$GLOBALS['Data_List_Lazy_Types_extendNonEmptyList'] = (object)["extend" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (\Data\Lazy\majData_majLazy_force($v_1))->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v2_3) use ($__local_var_2_0, $f_0, $v_1) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)($v_1), ((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldr'})((function() use ($f_0) {
  $__fn = function($a_4, $v1_5 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_1 = ($v1_5)->{'acc'};
  $__local_var_7_2 = ($v1_5)->{'val'};
  $__local_var_8_3 = ($f_0)(\Data\Lazy\majData_majLazy_defer(function($v2_8) use ($__local_var_6_1, $a_4) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_4, $__local_var_6_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__res = (object)["val" => \Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_7_2, $__local_var_8_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_8_3, $__local_var_7_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "acc" => \Data\Lazy\majData_majLazy_defer(function($v_8) use ($__local_var_6_1, $a_4) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_4, $__local_var_6_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["val" => $GLOBALS['Data_List_Lazy_Types_nil'], "acc" => $GLOBALS['Data_List_Lazy_Types_nil']]))($__local_var_2_0))->{'val'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_foldableNonEmptyList
$GLOBALS['Data_List_Lazy_Types_foldableNonEmptyList'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = \Data\Lazy\majData_majLazy_force($v_2);
  $__res = (($f_0)(($__local_var_3_0)->{'value0'}))((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldr'})($f_0))($b_1))(($__local_var_3_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_1 = \Data\Lazy\majData_majLazy_force($v_2);
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})($f_0))((($f_0)($b_1))(($__local_var_3_1)->{'value0'})))(($__local_var_3_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap1_1_2 = (($GLOBALS['Data_List_Lazy_Types_foldableNonEmpty'])->{'foldMap'})($dictMonoid_0);
  $__res = (function() use ($foldMap1_1_2) {
  $__fn = function($f_2, $v_3 = null) use ($foldMap1_1_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap1_1_2)($f_2))(\Data\Lazy\majData_majLazy_force($v_3));
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
$__t1 = "";
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_2_0)->{'value0'})))((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})((function() use ($dictShow_0) {
  $__fn = function($shown_3, $x_prime_4 = null) use ($dictShow_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($shown_3))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(","))((($dictShow_0)->{'show'})($x_prime_4)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(""))(($v_2_0)->{'value1'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(fromFoldable ["))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($__t1))("])"));
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
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(NonEmptyList "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(((\Data\Lazy\majData_majLazy_showmajLazy(\Data\NonEmpty\majData_majNonmajEmpty_showmajNonmajEmpty($dictShow_0, \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_showmajList($dictShow_0))))->{'show'})($v_1)))(")"));
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
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = "Nil";
goto end_branch_0;;
};
  if ($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("("))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_1)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" : "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(((\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_showmajList($dictShow_0))->{'show'})(($v_1)->{'value1'})))(")"))));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_showStep'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_showmajStep';

// Data_List_Lazy_Types_foldableWithIndexList
$GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'] = (object)["foldrWithIndex" => (function() {
  $__fn = function($f_0, $b_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_3_0 = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})(function($v1_3) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($v1_3)->{'value1'};
  $__local_var_5_1 = ($v1_3)->{'value0'};
  $__res = function($a_6) use ($__local_var_4_0, $__local_var_5_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_5_1 + 1), \Data\Lazy\majData_majLazy_defer(function($v_7) use ($__local_var_4_0, $a_6) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_6, $__local_var_4_0);
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
}))(new \Data\Tuple\Data_Tuple_Tuple(0, $GLOBALS['Data_List_Lazy_Types_nil'])))($xs_2);
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})(function($v1_4) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_5_3 = ($v1_4)->{'value1'};
  $__local_var_6_4 = ($v1_4)->{'value0'};
  $__res = function($a_7) use ($__local_var_5_3, $__local_var_6_4, $f_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_6_4 - 1), ((($f_0)(($__local_var_6_4 - 1)))($a_7))($__local_var_5_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple(($v_3_0)->{'value0'}, $b_1)))(($v_3_0)->{'value1'}))->{'value1'};
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0, $acc_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))(((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})(function($v_2) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_3_5 = ($v_2)->{'value1'};
  $__local_var_4_6 = ($v_2)->{'value0'};
  $__res = function($a_5) use ($__local_var_3_5, $__local_var_4_6, $f_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_4_6 + 1), ((($f_0)($__local_var_4_6))($__local_var_3_5))($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple(0, $acc_1)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_7 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_7) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'])->{'foldlWithIndex'})((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($i_3, $acc_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})($acc_4)))(($f_2)($i_3));
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
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_foldableWithIndexNonEmpty
$GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmpty'] = \Data\NonEmpty\majData_majNonmajEmpty_foldablemajWithmajIndexmajNonmajEmpty($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList']);

// Data_List_Lazy_Types_foldableWithIndexNonEmptyList
$GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmptyList'] = (object)["foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMapWithIndex1_1_0 = (($GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmpty'])->{'foldMapWithIndex'})($dictMonoid_0);
  $__res = (function() use ($foldMapWithIndex1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($foldMapWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMapWithIndex1_1_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))(function($v2_4) {
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
})))(\Data\Lazy\majData_majLazy_force($v_3));
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
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmpty'])->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = 0;
goto end_branch_2;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
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
})))($b_1))(\Data\Lazy\majData_majLazy_force($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldrWithIndex" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmpty'])->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = 0;
goto end_branch_3;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
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
})))($b_1))(\Data\Lazy\majData_majLazy_force($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_functorWithIndexList
$GLOBALS['Data_List_Lazy_Types_functorWithIndexList'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'])->{'foldrWithIndex'})((function() use ($f_0) {
  $__fn = function($i_1, $x_2 = null, $acc_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_4_0 = (($f_0)($i_1))($x_2);
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_5) use ($__local_var_4_0, $acc_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_4_0, $acc_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_mapWithIndex
$GLOBALS['Data_List_Lazy_Types_mapWithIndex'] = (\Data\NonEmpty\majData_majNonmajEmpty_functormajWithmajIndex($GLOBALS['Data_List_Lazy_Types_functorWithIndexList']))->{'mapWithIndex'};

// Data_List_Lazy_Types_functorWithIndexNonEmptyList
$GLOBALS['Data_List_Lazy_Types_functorWithIndexNonEmptyList'] = (object)["mapWithIndex" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v1_2) use ($f_0, $v_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_Lazy_Types_mapWithIndex'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
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
})))(\Data\Lazy\majData_majLazy_force($v_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
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
  $__res = (($GLOBALS['Data_List_Lazy_Types_lazyList'])->{'defer'})(function($v1_1) use ($v_0) {
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
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Types_toList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majTypes_tomajList';

// Data_List_Lazy_Types_semigroupNonEmptyList
$GLOBALS['Data_List_Lazy_Types_semigroupNonEmptyList'] = (object)["append" => (function() {
  $__fn = function($v_0, $as_prime_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_3_1 = ($v1_2_0)->{'value0'};
  $__local_var_4_2 = ($v1_2_0)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v2_5) use ($__local_var_3_1, $__local_var_4_2, $as_prime_1) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_3_1, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})($__local_var_4_2))(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_tomajList($as_prime_1)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Lazy_Types_traversableList
$GLOBALS['Data_List_Lazy_Types_traversableList'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldr'})((function() use ($Apply0_1_0, $f_2) {
  $__fn = function($a_3, $b_4 = null) use ($Apply0_1_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Apply0_1_0)->{'apply'})(((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_List_Lazy_Types_cons']))(($f_2)($a_3))))($b_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_traversableList'])->{'traverse'})($dictApplicative_0))($GLOBALS['Data_List_Lazy_Types_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_traversableNonEmpty
$GLOBALS['Data_List_Lazy_Types_traversableNonEmpty'] = \Data\NonEmpty\majData_majNonmajEmpty_traversablemajNonmajEmpty($GLOBALS['Data_List_Lazy_Types_traversableList']);

// Data_List_Lazy_Types_traversableNonEmptyList
$GLOBALS['Data_List_Lazy_Types_traversableNonEmptyList'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse1_1_0 = (($GLOBALS['Data_List_Lazy_Types_traversableNonEmpty'])->{'traverse'})($dictApplicative_0);
  $__res = (function() use ($dictApplicative_0, $traverse1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictApplicative_0, $traverse1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($xxs_4) {
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
}))((($traverse1_1_0)($f_2))(\Data\Lazy\majData_majLazy_force($v_3)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $sequence1_1_1 = (($GLOBALS['Data_List_Lazy_Types_traversableNonEmpty'])->{'sequence'})($dictApplicative_0);
  $__res = function($v_2) use ($dictApplicative_0, $sequence1_1_1) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($xxs_3) {
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
}))(($sequence1_1_1)(\Data\Lazy\majData_majLazy_force($v_2)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
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
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'])->{'foldrWithIndex'})((function() use ($Apply0_1_0, $f_2) {
  $__fn = function($i_3, $a_4 = null, $b_5 = null) use ($Apply0_1_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Apply0_1_0)->{'apply'})(((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_List_Lazy_Types_cons']))((($f_2)($i_3))($a_4))))($b_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorWithIndexList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableWithIndexList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_traversableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_traverseWithIndex
$GLOBALS['Data_List_Lazy_Types_traverseWithIndex'] = (\Data\NonEmpty\majData_majNonmajEmpty_traversablemajWithmajIndexmajNonmajEmpty($GLOBALS['Data_List_Lazy_Types_traversableWithIndexList']))->{'traverseWithIndex'};

// Data_List_Lazy_Types_traversableWithIndexNonEmptyList
$GLOBALS['Data_List_Lazy_Types_traversableWithIndexNonEmptyList'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverseWithIndex1_1_0 = ($GLOBALS['Data_List_Lazy_Types_traverseWithIndex'])($dictApplicative_0);
  $__res = (function() use ($dictApplicative_0, $traverseWithIndex1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictApplicative_0, $traverseWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(function($xxs_4) {
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
}))((($traverseWithIndex1_1_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))(function($v2_4) {
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
})))(\Data\Lazy\majData_majLazy_force($v_3)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorWithIndexNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_foldableWithIndexNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_traversableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_unfoldable1List
$GLOBALS['Data_List_Lazy_Types_unfoldable1List'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = (function() use (&$go__0_0) {
  $__fn = function($f_1, $b_2 = null) use (&$go__0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_List_Lazy_Types_lazyList'])->{'defer'})(function($v_3) use ($b_2, $f_1, &$go__0_0) {
  $__num = \func_num_args();
  $v1_4_1 = ($f_1)($b_2);
  $__t2 = null;;
  if (($v1_4_1)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_5_3 = ($v1_4_1)->{'value0'};
$__local_var_6_4 = (($go__0_0)($f_1))((($v1_4_1)->{'value1'})->{'value0'});
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
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
return (object)["unfoldr1" => $go__0_0];
})();

// Data_List_Lazy_Types_unfoldableList
$GLOBALS['Data_List_Lazy_Types_unfoldableList'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = (function() use (&$go__0_0) {
  $__fn = function($f_1, $b_2 = null) use (&$go__0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_List_Lazy_Types_lazyList'])->{'defer'})(function($v_3) use ($b_2, $f_1, &$go__0_0) {
  $__num = \func_num_args();
  $v1_4_1 = ($f_1)($b_2);
  $__t2 = null;;
  if ($v1_4_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = $GLOBALS['Data_List_Lazy_Types_nil'];
goto end_branch_2;;
};
  if ($v1_4_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_5_3 = (($v1_4_1)->{'value0'})->{'value0'};
$__local_var_6_4 = (($go__0_0)($f_1))((($v1_4_1)->{'value0'})->{'value1'});
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
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
return (object)["unfoldr" => $go__0_0, "Unfoldable10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_unfoldable1List'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_List_Lazy_Types_unfoldr1
$GLOBALS['Data_List_Lazy_Types_unfoldr1'] = (\Data\NonEmpty\majData_majNonmajEmpty_unfoldable1majNonmajEmpty($GLOBALS['Data_List_Lazy_Types_unfoldableList']))->{'unfoldr1'};

// Data_List_Lazy_Types_unfoldable1NonEmptyList
$GLOBALS['Data_List_Lazy_Types_unfoldable1NonEmptyList'] = (object)["unfoldr1" => (function() {
  $__fn = function($f_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_Lazy_Types_unfoldr1'])($f_0))($b_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Lazy_Types_comonadNonEmptyList
$GLOBALS['Data_List_Lazy_Types_comonadNonEmptyList'] = (object)["extract" => function($v_0) {
  $__num = \func_num_args();
  $__res = (\Data\Lazy\majData_majLazy_force($v_0))->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_extendNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_monadList
$GLOBALS['Data_List_Lazy_Types_monadList'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applicativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_bindList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_bindList
$GLOBALS['Data_List_Lazy_Types_bindList'] = (object)["bind" => (function() {
  $__fn = function($xs_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($f_1)(($v_2)->{'value0'})))(((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($v_2)->{'value1'}))($f_1)));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_applyList
$GLOBALS['Data_List_Lazy_Types_applyList'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Data_List_Lazy_Types_monadList'])->{'Bind1'})(null);
return (object)["apply" => (function() use ($__local_var_0_0) {
  $__fn = function($f_1, $a_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_0_0)->{'bind'})($f_1))(function($f_prime_3) use ($__local_var_0_0, $a_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_0_0)->{'bind'})($a_2))(function($a_prime_4) use ($f_prime_3) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_monadList'])->{'Applicative0'})(null))->{'pure'})(($f_prime_3)($a_prime_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

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
}, "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_applyNonEmptyList
$GLOBALS['Data_List_Lazy_Types_applyNonEmptyList'] = (object)["apply" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v2_2_0 = \Data\Lazy\majData_majLazy_force($v1_1);
  $v3_3_1 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_4_2 = ($v2_2_0)->{'value0'};
  $__local_var_5_3 = ($v2_2_0)->{'value1'};
  $__local_var_6_4 = ($v3_3_1)->{'value0'};
  $__local_var_7_5 = ($v3_3_1)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v4_8) use ($__local_var_4_2, $__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($__local_var_6_4)($__local_var_4_2), ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(((($GLOBALS['Data_List_Lazy_Types_applyList'])->{'apply'})($__local_var_7_5))(\Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_4_2, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(((($GLOBALS['Data_List_Lazy_Types_applyList'])->{'apply'})(\Data\Lazy\majData_majLazy_defer(function($v_9) use ($__local_var_6_4, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_6_4, $__local_var_7_5);
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
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_bindNonEmptyList
$GLOBALS['Data_List_Lazy_Types_bindNonEmptyList'] = (object)["bind" => (function() {
  $__fn = function($v_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_3_1 = ($v1_2_0)->{'value1'};
  $v2_4_2 = \Data\Lazy\majData_majLazy_force(($f_1)(($v1_2_0)->{'value0'}));
  $__local_var_5_3 = ($v2_4_2)->{'value0'};
  $__local_var_6_4 = ($v2_4_2)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v3_7) use ($__local_var_3_1, $__local_var_5_3, $__local_var_6_4, $f_1) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_5_3, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})($__local_var_6_4))(((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})($__local_var_3_1))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_toList']))($f_1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applyNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_altNonEmptyList
$GLOBALS['Data_List_Lazy_Types_altNonEmptyList'] = (object)["alt" => ($GLOBALS['Data_List_Lazy_Types_semigroupNonEmptyList'])->{'append'}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_altList
$GLOBALS['Data_List_Lazy_Types_altList'] = (object)["alt" => ($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_plusList
$GLOBALS['Data_List_Lazy_Types_plusList'] = (object)["empty" => $GLOBALS['Data_List_Lazy_Types_nil'], "Alt0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_altList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_alternativeList
$GLOBALS['Data_List_Lazy_Types_alternativeList'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applicativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_plusList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_monadPlusList
$GLOBALS['Data_List_Lazy_Types_monadPlusList'] = (object)["Monad0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_monadList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar__unused_0) {
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
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_0, ($GLOBALS['Data_List_Lazy_Types_plusList'])->{'empty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applyNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_Types_monadNonEmptyList
$GLOBALS['Data_List_Lazy_Types_monadNonEmptyList'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_applicativeNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_bindNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

