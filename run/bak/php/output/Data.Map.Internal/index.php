<?php

namespace Data\Map\Internal;

// ALL IMPORTS: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Plus, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.List, Data.List.Types, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude, Prim, Prim.TypeError
// TO REQUIRE: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Plus, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.List, Data.List.Types, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_Map_Internal_Leaf { public $tag = 'Leaf'; public function __construct() {} }
final class Data_Map_Internal_Node { public $tag = 'Node'; public function __construct(public int $value0, public int $value1, public  $value2, public  $value3, public  $value4, public  $value5) {} }
final class Data_Map_Internal_IterLeaf { public $tag = 'IterLeaf'; public function __construct() {} }
final class Data_Map_Internal_IterEmit { public $tag = 'IterEmit'; public function __construct(public  $value0, public  $value1, public  $value2) {} }
final class Data_Map_Internal_IterNode { public $tag = 'IterNode'; public function __construct(public  $value0, public  $value1) {} }
final class Data_Map_Internal_IterDone { public $tag = 'IterDone'; public function __construct() {} }
final class Data_Map_Internal_IterNext { public $tag = 'IterNext'; public function __construct(public  $value0, public  $value1, public  $value2) {} }
final class Data_Map_Internal_Split { public $tag = 'Split'; public function __construct(public  $value0, public  $value1, public  $value2) {} }
final class Data_Map_Internal_SplitLast { public $tag = 'SplitLast'; public function __construct(public  $value0, public  $value1, public  $value2) {} }

// Data_Map_Internal_greaterThan
$GLOBALS['Data_Map_Internal_greaterThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return function($a1_1) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__res = function($a2_2) use ($__local_var_0_0, $a1_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = true;
goto end_branch_1;;
};
  $__t1 = false;
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
})();

// Data_Map_Internal_lessThanOrEq
$GLOBALS['Data_Map_Internal_lessThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return function($a1_1) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__res = function($a2_2) use ($__local_var_0_0, $a1_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = false;
goto end_branch_1;;
};
  $__t1 = true;
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
})();

// Data_Map_Internal_identity
function majData_majMap_majInternal_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_identity'] = __NAMESPACE__ . '\\majData_majMap_majInternal_identity';

// Data_Map_Internal_lessThan
$GLOBALS['Data_Map_Internal_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return function($a1_1) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__res = function($a2_2) use ($__local_var_0_0, $a1_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = true;
goto end_branch_1;;
};
  $__t1 = false;
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
})();

// Data_Map_Internal_abs
$GLOBALS['Data_Map_Internal_abs'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return function($x_1) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((function() use ($__local_var_0_0, $x_1, &$__fn) {
$__t2 = null;;
if ((($__local_var_0_0)($x_1))(0) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t2 = false;
goto end_branch_2;;
};
$__t2 = true;
end_branch_2:;
return $__t2;
})()) {
$__t1 = $x_1;
goto end_branch_1;;
};
  $__t1 = ( - $x_1);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
})();

// Data_Map_Internal_identity1
function majData_majMap_majInternal_identity1($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_identity1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_identity1'] = __NAMESPACE__ . '\\majData_majMap_majInternal_identity1';

// Data_Map_Internal_identity2
function majData_majMap_majInternal_identity2($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_identity2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_identity2'] = __NAMESPACE__ . '\\majData_majMap_majInternal_identity2';

// Data_Map_Internal_Leaf
$GLOBALS['Data_Map_Internal_Leaf'] = ($GLOBALS['__phpurs_data0_Leaf'] ??= new \Data\Map\Internal\Data_Map_Internal_Leaf());

// Data_Map_Internal_Node
$GLOBALS['Data_Map_Internal_Node'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null, $value4 = null, $value5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node($value0, $value1, $value2, $value3, $value4, $value5);
  goto __end;;
  __end:
  return $__num > 6 ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_IterLeaf
$GLOBALS['Data_Map_Internal_IterLeaf'] = ($GLOBALS['__phpurs_data0_IterLeaf'] ??= new \Data\Map\Internal\Data_Map_Internal_IterLeaf());

// Data_Map_Internal_IterEmit
$GLOBALS['Data_Map_Internal_IterEmit'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterEmit($value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_IterNode
$GLOBALS['Data_Map_Internal_IterNode'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterNode($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_IterDone
$GLOBALS['Data_Map_Internal_IterDone'] = ($GLOBALS['__phpurs_data0_IterDone'] ??= new \Data\Map\Internal\Data_Map_Internal_IterDone());

// Data_Map_Internal_IterNext
$GLOBALS['Data_Map_Internal_IterNext'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterNext($value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_Split
$GLOBALS['Data_Map_Internal_Split'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Split($value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_SplitLast
$GLOBALS['Data_Map_Internal_SplitLast'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_SplitLast($value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_unsafeNode
function majData_majMap_majInternal_unsafemajNode($k_0, $__local_var_1 = null, $l_2 = null, $__local_var_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unsafemajNode';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ($l_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = null;;
if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $k_0, $__local_var_1, $l_2, $__local_var_3);
goto end_branch_1;;
};
if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node((1 + ($__local_var_3)->{'value0'}), (1 + ($__local_var_3)->{'value1'}), $k_0, $__local_var_1, $l_2, $__local_var_3);
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($l_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = new \Data\Map\Internal\Data_Map_Internal_Node((1 + ($l_2)->{'value0'}), (1 + ($l_2)->{'value1'}), $k_0, $__local_var_1, $l_2, $__local_var_3);
goto end_branch_2;;
};
if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = null;;
if ((($GLOBALS['Data_Map_Internal_greaterThan'])(($l_2)->{'value0'}))(($__local_var_3)->{'value0'})) {
$__t3 = (1 + ($l_2)->{'value0'});
goto end_branch_3;;
};
$__t3 = (1 + ($__local_var_3)->{'value0'});
end_branch_3:;
$__t2 = new \Data\Map\Internal\Data_Map_Internal_Node($__t3, ((1 + ($l_2)->{'value1'}) + ($__local_var_3)->{'value1'}), $k_0, $__local_var_1, $l_2, $__local_var_3);
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
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Map_Internal_unsafeNode'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unsafemajNode';

// Data_Map_Internal_toMapIter
function majData_majMap_majInternal_tomajMapmajIter($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_tomajMapmajIter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterNode($a_0, new \Data\Map\Internal\Data_Map_Internal_IterLeaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_toMapIter'] = __NAMESPACE__ . '\\majData_majMap_majInternal_tomajMapmajIter';

// Data_Map_Internal_stepWith
function majData_majMap_majInternal_stepmajWith($f_0, $next_1 = null, $done_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajWith';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($done_2, $f_0, &$go__go_3_0, $next_1) {
  $__num = \func_num_args();
  $__tco_var_go__go_3_0_0_v_4 = $v_4;
  tco_loop_go__go_3_0_0:;
  $v_4 = $__tco_var_go__go_3_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_IterLeaf) {
$__t0 = ($done_2)($GLOBALS['Data_Unit_unit']);
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_IterEmit) {
$__t0 = ($next_1)(($v_4)->{'value0'}, ($v_4)->{'value1'}, ($v_4)->{'value2'});
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_IterNode) {
$__tco_1 = (($f_0)(($v_4)->{'value1'}))(($v_4)->{'value0'});
$__tco_var_go__go_3_0_0_v_4 = $__tco_1;
goto tco_loop_go__go_3_0_0;;
$__t0 = null;
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
  $__res = $go__go_3_0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepWith'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajWith';

// Data_Map_Internal_size
function majData_majMap_majInternal_size($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_size';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t0 = ($v_0)->{'value1'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_size'] = __NAMESPACE__ . '\\majData_majMap_majInternal_size';

// Data_Map_Internal_singleton
function majData_majMap_majInternal_singleton($k_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_singleton';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $k_0, $v_1, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_singleton'] = __NAMESPACE__ . '\\majData_majMap_majInternal_singleton';

// Data_Map_Internal_unsafeBalancedNode
function majData_majMap_majInternal_unsafemajBalancedmajNode($k_0, $__local_var_1 = null, $l_2 = null, $__local_var_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unsafemajBalancedmajNode';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ($l_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = null;;
if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $k_0, $__local_var_1, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_1;;
};
if (($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node && (($GLOBALS['Data_Map_Internal_greaterThan'])(($__local_var_3)->{'value0'}))(1))) {
$__t2 = null;;
if ((function() use ($__local_var_3, &$__fn) {
$__t3 = null;;
if (($__local_var_3)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = 0;
goto end_branch_3;;
};
if (($__local_var_3)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = (($__local_var_3)->{'value5'})->{'value0'};
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
return (($__local_var_3)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Node && (($GLOBALS['Data_Map_Internal_greaterThan'])((($__local_var_3)->{'value4'})->{'value0'}))($__t3));
})()) {
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeNode'])((($__local_var_3)->{'value4'})->{'value2'}, (($__local_var_3)->{'value4'})->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, $l_2, (($__local_var_3)->{'value4'})->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeNode'])(($__local_var_3)->{'value2'}, ($__local_var_3)->{'value3'}, (($__local_var_3)->{'value4'})->{'value5'}, ($__local_var_3)->{'value5'}));
goto end_branch_2;;
};
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeNode'])(($__local_var_3)->{'value2'}, ($__local_var_3)->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, $l_2, ($__local_var_3)->{'value4'}), ($__local_var_3)->{'value5'});
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
$__t1 = ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, $l_2, $__local_var_3);
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($l_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t4 = null;;
if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t5 = null;;
if ((($GLOBALS['Data_Map_Internal_greaterThan'])(($__local_var_3)->{'value0'}))((($l_2)->{'value0'} + 1))) {
$__t6 = null;;
if ((function() use ($__local_var_3, &$__fn) {
$__t7 = null;;
if (($__local_var_3)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t7 = 0;
goto end_branch_7;;
};
if (($__local_var_3)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t7 = (($__local_var_3)->{'value5'})->{'value0'};
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
return (($__local_var_3)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Node && (($GLOBALS['Data_Map_Internal_greaterThan'])((($__local_var_3)->{'value4'})->{'value0'}))($__t7));
})()) {
$__t6 = ($GLOBALS['Data_Map_Internal_unsafeNode'])((($__local_var_3)->{'value4'})->{'value2'}, (($__local_var_3)->{'value4'})->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, $l_2, (($__local_var_3)->{'value4'})->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeNode'])(($__local_var_3)->{'value2'}, ($__local_var_3)->{'value3'}, (($__local_var_3)->{'value4'})->{'value5'}, ($__local_var_3)->{'value5'}));
goto end_branch_6;;
};
$__t6 = ($GLOBALS['Data_Map_Internal_unsafeNode'])(($__local_var_3)->{'value2'}, ($__local_var_3)->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, $l_2, ($__local_var_3)->{'value4'}), ($__local_var_3)->{'value5'});
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
if ((($GLOBALS['Data_Map_Internal_greaterThan'])(($l_2)->{'value0'}))((($__local_var_3)->{'value0'} + 1))) {
$__t8 = null;;
if ((function() use ($l_2, &$__fn) {
$__t9 = null;;
if (($l_2)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t9 = 0;
goto end_branch_9;;
};
if (($l_2)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t9 = (($l_2)->{'value4'})->{'value0'};
goto end_branch_9;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
end_branch_9:;
return (($l_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Node && (($GLOBALS['Data_Map_Internal_lessThanOrEq'])($__t9))((($l_2)->{'value5'})->{'value0'}));
})()) {
$__t8 = ($GLOBALS['Data_Map_Internal_unsafeNode'])((($l_2)->{'value5'})->{'value2'}, (($l_2)->{'value5'})->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])(($l_2)->{'value2'}, ($l_2)->{'value3'}, ($l_2)->{'value4'}, (($l_2)->{'value5'})->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, (($l_2)->{'value5'})->{'value5'}, $__local_var_3));
goto end_branch_8;;
};
$__t8 = ($GLOBALS['Data_Map_Internal_unsafeNode'])(($l_2)->{'value2'}, ($l_2)->{'value3'}, ($l_2)->{'value4'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, ($l_2)->{'value5'}, $__local_var_3));
end_branch_8:;
$__t5 = $__t8;
goto end_branch_5;;
};
$__t5 = ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, $l_2, $__local_var_3);
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
if (($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf && (($GLOBALS['Data_Map_Internal_greaterThan'])(($l_2)->{'value0'}))(1))) {
$__t10 = null;;
if ((function() use ($l_2, &$__fn) {
$__t11 = null;;
if (($l_2)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t11 = 0;
goto end_branch_11;;
};
if (($l_2)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t11 = (($l_2)->{'value4'})->{'value0'};
goto end_branch_11;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
end_branch_11:;
return (($l_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Node && (($GLOBALS['Data_Map_Internal_lessThanOrEq'])($__t11))((($l_2)->{'value5'})->{'value0'}));
})()) {
$__t10 = ($GLOBALS['Data_Map_Internal_unsafeNode'])((($l_2)->{'value5'})->{'value2'}, (($l_2)->{'value5'})->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])(($l_2)->{'value2'}, ($l_2)->{'value3'}, ($l_2)->{'value4'}, (($l_2)->{'value5'})->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, (($l_2)->{'value5'})->{'value5'}, $__local_var_3));
goto end_branch_10;;
};
$__t10 = ($GLOBALS['Data_Map_Internal_unsafeNode'])(($l_2)->{'value2'}, ($l_2)->{'value3'}, ($l_2)->{'value4'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, ($l_2)->{'value5'}, $__local_var_3));
end_branch_10:;
$__t4 = $__t10;
goto end_branch_4;;
};
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $__local_var_1, $l_2, $__local_var_3);
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
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Map_Internal_unsafeBalancedNode'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unsafemajBalancedmajNode';

// Data_Map_Internal_unsafeSplit
function majData_majMap_majInternal_unsafemajSplit($comp_0, $__local_var_1 = null, $m_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unsafemajSplit';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_Map_Internal_unsafeSplit_comp_0 = $comp_0;
  $__tco_var_Data_Map_Internal_unsafeSplit___local_var_1 = $__local_var_1;
  $__tco_var_Data_Map_Internal_unsafeSplit_m_2 = $m_2;
  tco_loop_Data_Map_Internal_unsafeSplit:;
  $comp_0 = $__tco_var_Data_Map_Internal_unsafeSplit_comp_0;
  $__local_var_1 = $__tco_var_Data_Map_Internal_unsafeSplit___local_var_1;
  $m_2 = $__tco_var_Data_Map_Internal_unsafeSplit_m_2;
  $__t0 = null;;
  if ($m_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Map\Internal\Data_Map_Internal_Split(new \Data\Maybe\Data_Maybe_Nothing(), new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_0;;
};
  if ($m_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v_3_1 = (($comp_0)($__local_var_1))(($m_2)->{'value2'});
$__t2 = null;;
if ($v_3_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$v1_4_3 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, $__local_var_1, ($m_2)->{'value4'});
$__t2 = new \Data\Map\Internal\Data_Map_Internal_Split(($v1_4_3)->{'value0'}, ($v1_4_3)->{'value1'}, ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($m_2)->{'value2'}, ($m_2)->{'value3'}, ($v1_4_3)->{'value2'}, ($m_2)->{'value5'}));
goto end_branch_2;;
};
if ($v_3_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$v1_4_4 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, $__local_var_1, ($m_2)->{'value5'});
$__t2 = new \Data\Map\Internal\Data_Map_Internal_Split(($v1_4_4)->{'value0'}, ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($m_2)->{'value2'}, ($m_2)->{'value3'}, ($m_2)->{'value4'}, ($v1_4_4)->{'value1'}), ($v1_4_4)->{'value2'});
goto end_branch_2;;
};
if ($v_3_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = new \Data\Map\Internal\Data_Map_Internal_Split(new \Data\Maybe\Data_Maybe_Just(($m_2)->{'value3'}), ($m_2)->{'value4'}, ($m_2)->{'value5'});
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
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_unsafeSplit'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unsafemajSplit';

// Data_Map_Internal_unsafeSplitLast
function majData_majMap_majInternal_unsafemajSplitmajLast($k_0, $__local_var_1 = null, $l_2 = null, $__local_var_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unsafemajSplitmajLast';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_Map_Internal_unsafeSplitLast_k_0 = $k_0;
  $__tco_var_Data_Map_Internal_unsafeSplitLast___local_var_1 = $__local_var_1;
  $__tco_var_Data_Map_Internal_unsafeSplitLast_l_2 = $l_2;
  $__tco_var_Data_Map_Internal_unsafeSplitLast___local_var_3 = $__local_var_3;
  tco_loop_Data_Map_Internal_unsafeSplitLast:;
  $k_0 = $__tco_var_Data_Map_Internal_unsafeSplitLast_k_0;
  $__local_var_1 = $__tco_var_Data_Map_Internal_unsafeSplitLast___local_var_1;
  $l_2 = $__tco_var_Data_Map_Internal_unsafeSplitLast_l_2;
  $__local_var_3 = $__tco_var_Data_Map_Internal_unsafeSplitLast___local_var_3;
  $__t0 = null;;
  if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Map\Internal\Data_Map_Internal_SplitLast($k_0, $__local_var_1, $l_2);
goto end_branch_0;;
};
  if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_1 = ($GLOBALS['Data_Map_Internal_unsafeSplitLast'])(($__local_var_3)->{'value2'}, ($__local_var_3)->{'value3'}, ($__local_var_3)->{'value4'}, ($__local_var_3)->{'value5'});
$__t0 = new \Data\Map\Internal\Data_Map_Internal_SplitLast(($v1_4_1)->{'value0'}, ($v1_4_1)->{'value1'}, ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])($k_0, $__local_var_1, $l_2, ($v1_4_1)->{'value2'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Map_Internal_unsafeSplitLast'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unsafemajSplitmajLast';

// Data_Map_Internal_unsafeJoinNodes
function majData_majMap_majInternal_unsafemajJoinmajNodes($v_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unsafemajJoinmajNodes';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = $__local_var_1;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_2_1 = ($GLOBALS['Data_Map_Internal_unsafeSplitLast'])(($v_0)->{'value2'}, ($v_0)->{'value3'}, ($v_0)->{'value4'}, ($v_0)->{'value5'});
$__t0 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v2_2_1)->{'value0'}, ($v2_2_1)->{'value1'}, ($v2_2_1)->{'value2'}, $__local_var_1);
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_unsafeJoinNodes'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unsafemajJoinmajNodes';

// Data_Map_Internal_unsafeDifference
function majData_majMap_majInternal_unsafemajDifference($comp_0, $__local_var_1 = null, $r_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unsafemajDifference';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_Map_Internal_unsafeDifference_comp_0 = $comp_0;
  $__tco_var_Data_Map_Internal_unsafeDifference___local_var_1 = $__local_var_1;
  $__tco_var_Data_Map_Internal_unsafeDifference_r_2 = $r_2;
  tco_loop_Data_Map_Internal_unsafeDifference:;
  $comp_0 = $__tco_var_Data_Map_Internal_unsafeDifference_comp_0;
  $__local_var_1 = $__tco_var_Data_Map_Internal_unsafeDifference___local_var_1;
  $r_2 = $__tco_var_Data_Map_Internal_unsafeDifference_r_2;
  $__t0 = null;;
  if ($__local_var_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_0;;
};
  if ($r_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = $__local_var_1;
goto end_branch_0;;
};
  if ($r_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v_3_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, ($r_2)->{'value2'}, $__local_var_1);
$__t0 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($GLOBALS['Data_Map_Internal_unsafeDifference'])($comp_0, ($v_3_1)->{'value1'}, ($r_2)->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeDifference'])($comp_0, ($v_3_1)->{'value2'}, ($r_2)->{'value5'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_unsafeDifference'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unsafemajDifference';

// Data_Map_Internal_unsafeIntersectionWith
function majData_majMap_majInternal_unsafemajIntersectionmajWith($comp_0, $__local_var_1 = null, $l_2 = null, $__local_var_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unsafemajIntersectionmajWith';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_Map_Internal_unsafeIntersectionWith_comp_0 = $comp_0;
  $__tco_var_Data_Map_Internal_unsafeIntersectionWith___local_var_1 = $__local_var_1;
  $__tco_var_Data_Map_Internal_unsafeIntersectionWith_l_2 = $l_2;
  $__tco_var_Data_Map_Internal_unsafeIntersectionWith___local_var_3 = $__local_var_3;
  tco_loop_Data_Map_Internal_unsafeIntersectionWith:;
  $comp_0 = $__tco_var_Data_Map_Internal_unsafeIntersectionWith_comp_0;
  $__local_var_1 = $__tco_var_Data_Map_Internal_unsafeIntersectionWith___local_var_1;
  $l_2 = $__tco_var_Data_Map_Internal_unsafeIntersectionWith_l_2;
  $__local_var_3 = $__tco_var_Data_Map_Internal_unsafeIntersectionWith___local_var_3;
  $__t0 = null;;
  if ($l_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_0;;
};
  if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_0;;
};
  if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v_4_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, ($__local_var_3)->{'value2'}, $l_2);
$l_prime_5_2 = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($comp_0, $__local_var_1, ($v_4_1)->{'value1'}, ($__local_var_3)->{'value4'});
$r_prime_6_3 = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($comp_0, $__local_var_1, ($v_4_1)->{'value2'}, ($__local_var_3)->{'value5'});
$__t4 = null;;
if (($v_4_1)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($__local_var_3)->{'value2'}, (($__local_var_1)((($v_4_1)->{'value0'})->{'value0'}))(($__local_var_3)->{'value3'}), $l_prime_5_2, $r_prime_6_3);
goto end_branch_4;;
};
if (($v_4_1)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])($l_prime_5_2, $r_prime_6_3);
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
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
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unsafemajIntersectionmajWith';

// Data_Map_Internal_unsafeUnionWith
function majData_majMap_majInternal_unsafemajUnionmajWith($comp_0, $__local_var_1 = null, $l_2 = null, $__local_var_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unsafemajUnionmajWith';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_Map_Internal_unsafeUnionWith_comp_0 = $comp_0;
  $__tco_var_Data_Map_Internal_unsafeUnionWith___local_var_1 = $__local_var_1;
  $__tco_var_Data_Map_Internal_unsafeUnionWith_l_2 = $l_2;
  $__tco_var_Data_Map_Internal_unsafeUnionWith___local_var_3 = $__local_var_3;
  tco_loop_Data_Map_Internal_unsafeUnionWith:;
  $comp_0 = $__tco_var_Data_Map_Internal_unsafeUnionWith_comp_0;
  $__local_var_1 = $__tco_var_Data_Map_Internal_unsafeUnionWith___local_var_1;
  $l_2 = $__tco_var_Data_Map_Internal_unsafeUnionWith_l_2;
  $__local_var_3 = $__tco_var_Data_Map_Internal_unsafeUnionWith___local_var_3;
  $__t0 = null;;
  if ($l_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = $__local_var_3;
goto end_branch_0;;
};
  if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = $l_2;
goto end_branch_0;;
};
  if ($__local_var_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v_4_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, ($__local_var_3)->{'value2'}, $l_2);
$l_prime_5_2 = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($comp_0, $__local_var_1, ($v_4_1)->{'value1'}, ($__local_var_3)->{'value4'});
$r_prime_6_3 = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($comp_0, $__local_var_1, ($v_4_1)->{'value2'}, ($__local_var_3)->{'value5'});
$__t4 = null;;
if (($v_4_1)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($__local_var_3)->{'value2'}, (($__local_var_1)((($v_4_1)->{'value0'})->{'value0'}))(($__local_var_3)->{'value3'}), $l_prime_5_2, $r_prime_6_3);
goto end_branch_4;;
};
if (($v_4_1)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($__local_var_3)->{'value2'}, ($__local_var_3)->{'value3'}, $l_prime_5_2, $r_prime_6_3);
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
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
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Map_Internal_unsafeUnionWith'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unsafemajUnionmajWith';

// Data_Map_Internal_unionWith
function majData_majMap_majInternal_unionmajWith($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unionmajWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($app_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m1_3) use ($app_2, $compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_4) use ($app_2, $compare_1_0, $m1_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $app_2, $m1_3, $m2_4);
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
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_unionWith'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unionmajWith';

// Data_Map_Internal_union
function majData_majMap_majInternal_union($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_union';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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
$GLOBALS['Data_Map_Internal_union'] = __NAMESPACE__ . '\\majData_majMap_majInternal_union';

// Data_Map_Internal_update
function majData_majMap_majInternal_update($dictOrd_0, $f_1 = null, $k_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_update';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictOrd_0, $f_1, &$go__go_3_0, $k_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_5_2 = ((($dictOrd_0)->{'compare'})($k_2))(($v_4)->{'value2'});
$__t3 = null;;
if ($v1_5_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_4)->{'value2'}, ($v_4)->{'value3'}, ($go__go_3_0)(($v_4)->{'value4'}), ($v_4)->{'value5'});
goto end_branch_3;;
};
if ($v1_5_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_4)->{'value2'}, ($v_4)->{'value3'}, ($v_4)->{'value4'}, ($go__go_3_0)(($v_4)->{'value5'}));
goto end_branch_3;;
};
if ($v1_5_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$v2_6_4 = ($f_1)(($v_4)->{'value3'});
$__t5 = null;;
if ($v2_6_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($v_4)->{'value4'}, ($v_4)->{'value5'});
goto end_branch_5;;
};
if ($v2_6_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = new \Data\Map\Internal\Data_Map_Internal_Node(($v_4)->{'value0'}, ($v_4)->{'value1'}, ($v_4)->{'value2'}, ($v2_6_4)->{'value0'}, ($v_4)->{'value4'}, ($v_4)->{'value5'});
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
  $__res = $go__go_3_0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_update'] = __NAMESPACE__ . '\\majData_majMap_majInternal_update';

// Data_Map_Internal_showTree
function majData_majMap_majInternal_showmajTree($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_showmajTree';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($ind_3) use ($dictShow1_1, $dictShow_0, &$go__go_2_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictShow1_1, $dictShow_0, &$go__go_2_0, $ind_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($ind_3))("Leaf");
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($ind_3))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("["))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_Show_showInt'])->{'show'})(($v_4)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("] "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_4)->{'value2'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" => "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow1_1)->{'show'})(($v_4)->{'value3'})))("
")))))))))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($go__go_2_0)(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($ind_3))("    ")))(($v_4)->{'value4'})))("
")))((($go__go_2_0)(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($ind_3))("    ")))(($v_4)->{'value5'})));
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
};
  $__res = ($go__go_2_0)("");
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_showTree'] = __NAMESPACE__ . '\\majData_majMap_majInternal_showmajTree';

// Data_Map_Internal_semigroupMap
function majData_majMap_majInternal_semigroupmajMap($_dollar__unused_0, $dictOrd_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_semigroupmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $compare_2_0 = ($dictOrd_1)->{'compare'};
  $__res = function($dictSemigroup_3) use ($compare_2_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = ($dictSemigroup_3)->{'append'};
  $__res = (object)["append" => function($m1_5) use ($__local_var_4_1, $compare_2_0) {
  $__num = \func_num_args();
  $__res = function($m2_6) use ($__local_var_4_1, $compare_2_0, $m1_5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_2_0, $__local_var_4_1, $m1_5, $m2_6);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_semigroupMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_semigroupmajMap';

// Data_Map_Internal_pop
function majData_majMap_majInternal_pop($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_pop';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($k_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m_3) use ($compare_1_0, $k_2) {
  $__num = \func_num_args();
  $v_4_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($compare_1_0, $k_2, $m_3);
  $__local_var_5_2 = ($v_4_1)->{'value1'};
  $__local_var_6_3 = ($v_4_1)->{'value2'};
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($a_7) use ($__local_var_5_2, $__local_var_6_3) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_7, ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])($__local_var_5_2, $__local_var_6_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4_1)->{'value0'});
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
$GLOBALS['Data_Map_Internal_pop'] = __NAMESPACE__ . '\\majData_majMap_majInternal_pop';

// Data_Map_Internal_member
function majData_majMap_majInternal_member($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_member';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__tco_var_go__go_2_0_0_v_3 = $v_3;
  tco_loop_go__go_2_0_0:;
  $v_3 = $__tco_var_go__go_2_0_0_v_3;
  $__t0 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = false;
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_1 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t2 = null;;
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_3 = ($v_3)->{'value4'};
$__tco_var_go__go_2_0_0_v_3 = $__tco_3;
goto tco_loop_go__go_2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__tco_4 = ($v_3)->{'value5'};
$__tco_var_go__go_2_0_0_v_3 = $__tco_4;
goto tco_loop_go__go_2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = true;
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_member'] = __NAMESPACE__ . '\\majData_majMap_majInternal_member';

// Data_Map_Internal_mapMaybeWithKey
function majData_majMap_majInternal_mapmajMaybemajWithmajKey($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_mapmajMaybemajWithmajKey';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($f_1, &$go__go_2_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_4_2 = (($f_1)(($v_3)->{'value2'}))(($v_3)->{'value3'});
$__t3 = null;;
if ($v2_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v2_4_2)->{'value0'}, ($go__go_2_0)(($v_3)->{'value4'}), ($go__go_2_0)(($v_3)->{'value5'}));
goto end_branch_3;;
};
if ($v2_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($go__go_2_0)(($v_3)->{'value4'}), ($go__go_2_0)(($v_3)->{'value5'}));
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_mapMaybeWithKey'] = __NAMESPACE__ . '\\majData_majMap_majInternal_mapmajMaybemajWithmajKey';

// Data_Map_Internal_mapMaybe
function majData_majMap_majInternal_mapmajMaybe($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_mapmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Map_Internal_mapMaybeWithKey'])($dictOrd_0)))($GLOBALS['Data_Function_const']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_mapMaybe'] = __NAMESPACE__ . '\\majData_majMap_majInternal_mapmajMaybe';

// Data_Map_Internal_lookupLE
function majData_majMap_majInternal_lookupmajLmajE($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_lookupmajLmajE';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_2 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($go__go_2_0)(($v_3)->{'value4'});
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$v2_5_4 = ($go__go_2_0)(($v_3)->{'value5'});
$__t5 = null;;
if ($v2_5_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = new \Data\Maybe\Data_Maybe_Just((object)["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
goto end_branch_5;;
};
$__t5 = $v2_5_4;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = new \Data\Maybe\Data_Maybe_Just((object)["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_lookupLE'] = __NAMESPACE__ . '\\majData_majMap_majInternal_lookupmajLmajE';

// Data_Map_Internal_lookupGE
function majData_majMap_majInternal_lookupmajGmajE($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_lookupmajGmajE';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_2 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$v2_5_4 = ($go__go_2_0)(($v_3)->{'value4'});
$__t5 = null;;
if ($v2_5_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = new \Data\Maybe\Data_Maybe_Just((object)["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
goto end_branch_5;;
};
$__t5 = $v2_5_4;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($go__go_2_0)(($v_3)->{'value5'});
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = new \Data\Maybe\Data_Maybe_Just((object)["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_lookupGE'] = __NAMESPACE__ . '\\majData_majMap_majInternal_lookupmajGmajE';

// Data_Map_Internal_lookup
function majData_majMap_majInternal_lookup($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_lookup';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__tco_var_go__go_2_0_0_v_3 = $v_3;
  tco_loop_go__go_2_0_0:;
  $v_3 = $__tco_var_go__go_2_0_0_v_3;
  $__t0 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_1 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t2 = null;;
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_3 = ($v_3)->{'value4'};
$__tco_var_go__go_2_0_0_v_3 = $__tco_3;
goto tco_loop_go__go_2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__tco_4 = ($v_3)->{'value5'};
$__tco_var_go__go_2_0_0_v_3 = $__tco_4;
goto tco_loop_go__go_2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_4_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($v_3)->{'value3'});
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_lookup'] = __NAMESPACE__ . '\\majData_majMap_majInternal_lookup';

// Data_Map_Internal_iterMapU
function majData_majMap_majInternal_itermajMapmajU($iter_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_itermajMapmajU';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_1 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = $iter_0;
goto end_branch_0;;
};
  if ($v_1 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if (($v_1)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = null;;
if (($v_1)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = new \Data\Map\Internal\Data_Map_Internal_IterEmit(($v_1)->{'value2'}, ($v_1)->{'value3'}, $iter_0);
goto end_branch_3;;
};
$__t3 = new \Data\Map\Internal\Data_Map_Internal_IterEmit(($v_1)->{'value2'}, ($v_1)->{'value3'}, new \Data\Map\Internal\Data_Map_Internal_IterNode(($v_1)->{'value5'}, $iter_0));
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
$__t1 = null;;
if (($v_1)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_IterEmit(($v_1)->{'value2'}, ($v_1)->{'value3'}, new \Data\Map\Internal\Data_Map_Internal_IterNode(($v_1)->{'value4'}, $iter_0));
goto end_branch_1;;
};
$__t1 = new \Data\Map\Internal\Data_Map_Internal_IterEmit(($v_1)->{'value2'}, ($v_1)->{'value3'}, new \Data\Map\Internal\Data_Map_Internal_IterNode(($v_1)->{'value4'}, new \Data\Map\Internal\Data_Map_Internal_IterNode(($v_1)->{'value5'}, $iter_0)));
end_branch_1:;
$__t2 = $__t1;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_iterMapU'] = __NAMESPACE__ . '\\majData_majMap_majInternal_itermajMapmajU';

// Data_Map_Internal_stepUnorderedCps_closure
$GLOBALS['Data_Map_Internal_stepUnorderedCps_closure'] = ($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapU']);

// Data_Map_Internal_stepUnorderedCps
function majData_majMap_majInternal_stepmajUnorderedmajCps($v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajUnorderedmajCps';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Map_Internal_stepUnorderedCps_closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepUnorderedCps'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajUnorderedmajCps';

// Data_Map_Internal_stepUnfoldrUnordered_closure
$GLOBALS['Data_Map_Internal_stepUnfoldrUnordered_closure'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapU']))((function() {
  $__fn = function($k_0, $__local_var_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(new \Data\Tuple\Data_Tuple_Tuple($k_0, $__local_var_1), $next_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_stepUnfoldrUnordered
function majData_majMap_majInternal_stepmajUnfoldrmajUnordered($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajUnfoldrmajUnordered';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_stepUnfoldrUnordered_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepUnfoldrUnordered'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajUnfoldrmajUnordered';

// Data_Map_Internal_toUnfoldableUnordered
function majData_majMap_majInternal_tomajUnfoldablemajUnordered($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_tomajUnfoldablemajUnordered';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable_0)->{'unfoldr'})($GLOBALS['Data_Map_Internal_stepUnfoldrUnordered'])))($GLOBALS['Data_Map_Internal_toMapIter']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_toUnfoldableUnordered'] = __NAMESPACE__ . '\\majData_majMap_majInternal_tomajUnfoldablemajUnordered';

// Data_Map_Internal_stepUnordered_closure
$GLOBALS['Data_Map_Internal_stepUnordered_closure'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapU']))((function() {
  $__fn = function($k_0, $__local_var_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterNext($k_0, $__local_var_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterDone();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_stepUnordered
function majData_majMap_majInternal_stepmajUnordered($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajUnordered';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_stepUnordered_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepUnordered'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajUnordered';

// Data_Map_Internal_iterMapR_closure
$GLOBALS['Data_Map_Internal_iterMapR_closure'] = (function() use (&$__fn) {
$go__go_0_0 = null;
$go__go_0_0 = (function() use (&$go__go_0_0) {
  $__fn = function($iter_1, $v_2 = null) use (&$go__go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_0_0_0_iter_1 = $iter_1;
  $__tco_var_go__go_0_0_0_v_2 = $v_2;
  tco_loop_go__go_0_0_0:;
  $iter_1 = $__tco_var_go__go_0_0_0_iter_1;
  $v_2 = $__tco_var_go__go_0_0_0_v_2;
  $__t0 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = $iter_1;
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = null;;
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__tco_4 = new \Data\Map\Internal\Data_Map_Internal_IterEmit(($v_2)->{'value2'}, ($v_2)->{'value3'}, $iter_1);
$__tco_5 = ($v_2)->{'value4'};
$__tco_var_go__go_0_0_0_iter_1 = $__tco_4;
$__tco_var_go__go_0_0_0_v_2 = $__tco_5;
goto tco_loop_go__go_0_0_0;;
$__t3 = null;
goto end_branch_3;;
};
$__tco_1 = new \Data\Map\Internal\Data_Map_Internal_IterEmit(($v_2)->{'value2'}, ($v_2)->{'value3'}, new \Data\Map\Internal\Data_Map_Internal_IterNode(($v_2)->{'value4'}, $iter_1));
$__tco_2 = ($v_2)->{'value5'};
$__tco_var_go__go_0_0_0_iter_1 = $__tco_1;
$__tco_var_go__go_0_0_0_v_2 = $__tco_2;
goto tco_loop_go__go_0_0_0;;
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
return $go__go_0_0;
})();

// Data_Map_Internal_iterMapR
function majData_majMap_majInternal_itermajMapmajR($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_itermajMapmajR';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_iterMapR_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_iterMapR'] = __NAMESPACE__ . '\\majData_majMap_majInternal_itermajMapmajR';

// Data_Map_Internal_stepDescCps_closure
$GLOBALS['Data_Map_Internal_stepDescCps_closure'] = ($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapR']);

// Data_Map_Internal_stepDescCps
function majData_majMap_majInternal_stepmajDescmajCps($v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajDescmajCps';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Map_Internal_stepDescCps_closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepDescCps'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajDescmajCps';

// Data_Map_Internal_stepDesc_closure
$GLOBALS['Data_Map_Internal_stepDesc_closure'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapR']))((function() {
  $__fn = function($k_0, $__local_var_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterNext($k_0, $__local_var_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterDone();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_stepDesc
function majData_majMap_majInternal_stepmajDesc($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajDesc';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_stepDesc_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepDesc'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajDesc';

// Data_Map_Internal_iterMapL_closure
$GLOBALS['Data_Map_Internal_iterMapL_closure'] = (function() use (&$__fn) {
$go__go_0_0 = null;
$go__go_0_0 = (function() use (&$go__go_0_0) {
  $__fn = function($iter_1, $v_2 = null) use (&$go__go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_0_0_0_iter_1 = $iter_1;
  $__tco_var_go__go_0_0_0_v_2 = $v_2;
  tco_loop_go__go_0_0_0:;
  $iter_1 = $__tco_var_go__go_0_0_0_iter_1;
  $v_2 = $__tco_var_go__go_0_0_0_v_2;
  $__t0 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = $iter_1;
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = null;;
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__tco_4 = new \Data\Map\Internal\Data_Map_Internal_IterEmit(($v_2)->{'value2'}, ($v_2)->{'value3'}, $iter_1);
$__tco_5 = ($v_2)->{'value4'};
$__tco_var_go__go_0_0_0_iter_1 = $__tco_4;
$__tco_var_go__go_0_0_0_v_2 = $__tco_5;
goto tco_loop_go__go_0_0_0;;
$__t3 = null;
goto end_branch_3;;
};
$__tco_1 = new \Data\Map\Internal\Data_Map_Internal_IterEmit(($v_2)->{'value2'}, ($v_2)->{'value3'}, new \Data\Map\Internal\Data_Map_Internal_IterNode(($v_2)->{'value5'}, $iter_1));
$__tco_2 = ($v_2)->{'value4'};
$__tco_var_go__go_0_0_0_iter_1 = $__tco_1;
$__tco_var_go__go_0_0_0_v_2 = $__tco_2;
goto tco_loop_go__go_0_0_0;;
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
return $go__go_0_0;
})();

// Data_Map_Internal_iterMapL
function majData_majMap_majInternal_itermajMapmajL($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_itermajMapmajL';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_iterMapL_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_iterMapL'] = __NAMESPACE__ . '\\majData_majMap_majInternal_itermajMapmajL';

// Data_Map_Internal_stepAscCps_closure
$GLOBALS['Data_Map_Internal_stepAscCps_closure'] = ($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapL']);

// Data_Map_Internal_stepAscCps
function majData_majMap_majInternal_stepmajAscmajCps($v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajAscmajCps';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Map_Internal_stepAscCps_closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepAscCps'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajAscmajCps';

// Data_Map_Internal_stepAsc_closure
$GLOBALS['Data_Map_Internal_stepAsc_closure'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapL']))((function() {
  $__fn = function($k_0, $__local_var_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterNext($k_0, $__local_var_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Map\Internal\Data_Map_Internal_IterDone();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_stepAsc
function majData_majMap_majInternal_stepmajAsc($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajAsc';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_stepAsc_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepAsc'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajAsc';

// Data_Map_Internal_eqMapIter
function majData_majMap_majInternal_eqmajMapmajIter($dictEq_0, $dictEq1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_eqmajMapmajIter';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use ($dictEq1_1, $dictEq_0, &$go__go_2_0) {
  $__fn = function($a_3, $b_4 = null) use ($dictEq1_1, $dictEq_0, &$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_a_3 = $a_3;
  $__tco_var_go__go_2_0_0_b_4 = $b_4;
  tco_loop_go__go_2_0_0:;
  $a_3 = $__tco_var_go__go_2_0_0_a_3;
  $b_4 = $__tco_var_go__go_2_0_0_b_4;
  $v_5_0 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_3);
  $__t1 = null;;
  if ($v_5_0 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v2_6_2 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_4);
$__t3 = null;;
if (($v2_6_2 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext && ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($dictEq_0)->{'eq'})(($v_5_0)->{'value0'}))(($v2_6_2)->{'value0'})))(((($dictEq1_1)->{'eq'})(($v_5_0)->{'value1'}))(($v2_6_2)->{'value1'})))) {
$__tco_4 = ($v_5_0)->{'value2'};
$__tco_5 = ($v2_6_2)->{'value2'};
$__tco_var_go__go_2_0_0_a_3 = $__tco_4;
$__tco_var_go__go_2_0_0_b_4 = $__tco_5;
goto tco_loop_go__go_2_0_0;;
$__t3 = null;
goto end_branch_3;;
};
$__t3 = false;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  if ($v_5_0 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t1 = true;
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
  $__res = (object)["eq" => $go__go_2_0];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_eqMapIter'] = __NAMESPACE__ . '\\majData_majMap_majInternal_eqmajMapmajIter';

// Data_Map_Internal_ordMapIter
function majData_majMap_majInternal_ordmajMapmajIter($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_ordmajMapmajIter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eqMapIter1_1_0 = ($GLOBALS['Data_Map_Internal_eqMapIter'])((($dictOrd_0)->{'Eq0'})(null));
  $__res = function($dictOrd1_2) use ($dictOrd_0, $eqMapIter1_1_0) {
  $__num = \func_num_args();
  $eqMapIter2_3_1 = ($eqMapIter1_1_0)((($dictOrd1_2)->{'Eq0'})(null));
  $go__go_4_2 = null;
  $go__go_4_2 = (function() use ($dictOrd1_2, $dictOrd_0, &$go__go_4_2) {
  $__fn = function($a_5, $b_6 = null) use ($dictOrd1_2, $dictOrd_0, &$go__go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_2_2_a_5 = $a_5;
  $__tco_var_go__go_4_2_2_b_6 = $b_6;
  tco_loop_go__go_4_2_2:;
  $a_5 = $__tco_var_go__go_4_2_2_a_5;
  $b_6 = $__tco_var_go__go_4_2_2_b_6;
  $v_7_2 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($b_6);
  $v1_8_3 = \Data\Map\Internal\majData_majMap_majInternal_stepmajAsc($a_5);
  $__t4 = null;;
  if ($v1_8_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$__t5 = null;;
if ($v_7_2 instanceof \Data\Map\Internal\Data_Map_Internal_IterNext) {
$v3_9_6 = ((($dictOrd_0)->{'compare'})(($v1_8_3)->{'value0'}))(($v_7_2)->{'value0'});
$__t7 = null;;
if ($v3_9_6 instanceof \Data\Ordering\Data_Ordering_EQ) {
$v4_10_8 = ((($dictOrd1_2)->{'compare'})(($v1_8_3)->{'value1'}))(($v_7_2)->{'value1'});
$__t9 = null;;
if ($v4_10_8 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_10 = ($v1_8_3)->{'value2'};
$__tco_11 = ($v_7_2)->{'value2'};
$__tco_var_go__go_4_2_2_a_5 = $__tco_10;
$__tco_var_go__go_4_2_2_b_6 = $__tco_11;
goto tco_loop_go__go_4_2_2;;
$__t9 = null;
goto end_branch_9;;
};
$__t9 = $v4_10_8;
end_branch_9:;
$__t7 = $__t9;
goto end_branch_7;;
};
$__t7 = $v3_9_6;
end_branch_7:;
$__t5 = $__t7;
goto end_branch_5;;
};
if ($v_7_2 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($v1_8_3 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t12 = null;;
if ($v_7_2 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t12 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_12;;
};
$__t12 = new \Data\Ordering\Data_Ordering_LT();
end_branch_12:;
$__t4 = $__t12;
goto end_branch_4;;
};
  if ($v_7_2 instanceof \Data\Map\Internal\Data_Map_Internal_IterDone) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
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
  $__res = (object)["compare" => $go__go_4_2, "Eq0" => function($_dollar__unused_4) use ($eqMapIter2_3_1) {
  $__num = \func_num_args();
  $__res = $eqMapIter2_3_1;
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
$GLOBALS['Data_Map_Internal_ordMapIter'] = __NAMESPACE__ . '\\majData_majMap_majInternal_ordmajMapmajIter';

// Data_Map_Internal_stepUnfoldr_closure
$GLOBALS['Data_Map_Internal_stepUnfoldr_closure'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapL']))((function() {
  $__fn = function($k_0, $__local_var_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(new \Data\Tuple\Data_Tuple_Tuple($k_0, $__local_var_1), $next_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_stepUnfoldr
function majData_majMap_majInternal_stepmajUnfoldr($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_stepmajUnfoldr';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_stepUnfoldr_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_stepUnfoldr'] = __NAMESPACE__ . '\\majData_majMap_majInternal_stepmajUnfoldr';

// Data_Map_Internal_toUnfoldable
function majData_majMap_majInternal_tomajUnfoldable($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_tomajUnfoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable_0)->{'unfoldr'})($GLOBALS['Data_Map_Internal_stepUnfoldr'])))($GLOBALS['Data_Map_Internal_toMapIter']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_toUnfoldable'] = __NAMESPACE__ . '\\majData_majMap_majInternal_tomajUnfoldable';

// Data_Map_Internal_toUnfoldable1
$GLOBALS['Data_Map_Internal_toUnfoldable1'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Unfoldable_unfoldableArray'])->{'unfoldr'})($GLOBALS['Data_Map_Internal_stepUnfoldr'])))($GLOBALS['Data_Map_Internal_toMapIter']);

// Data_Map_Internal_showMap
function majData_majMap_majInternal_showmajMap($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_showmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $show1_2_0 = ($GLOBALS['Data_Show_showArrayImpl'])(((($GLOBALS['Data_Tuple_showTuple'])($dictShow_0))($dictShow1_1))->{'show'});
  $__res = (object)["show" => function($as_3) use ($show1_2_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(fromFoldable "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(($show1_2_0)(($GLOBALS['Data_Map_Internal_toUnfoldable1'])($as_3))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_showMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_showmajMap';

// Data_Map_Internal_isSubmap
function majData_majMap_majInternal_ismajSubmap($dictOrd_0, $dictEq_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_ismajSubmap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($m1_3) use ($dictEq_1, $dictOrd_0, &$go__go_2_0) {
  $__num = \func_num_args();
  $__res = function($m2_4) use ($dictEq_1, $dictOrd_0, &$go__go_2_0, $m1_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($m1_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = true;
goto end_branch_1;;
};
  if ($m1_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__local_var_5_2 = ($m1_3)->{'value2'};
$go__go_6_3 = null;
$go__go_6_3 = function($v_7) use ($__local_var_5_2, $dictOrd_0, &$go__go_6_3) {
  $__num = \func_num_args();
  $__tco_var_go__go_6_3_3_v_7 = $v_7;
  tco_loop_go__go_6_3_3:;
  $v_7 = $__tco_var_go__go_6_3_3_v_7;
  $__t3 = null;;
  if ($v_7 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_3;;
};
  if ($v_7 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_8_4 = ((($dictOrd_0)->{'compare'})($__local_var_5_2))(($v_7)->{'value2'});
$__t5 = null;;
if ($v1_8_4 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_6 = ($v_7)->{'value4'};
$__tco_var_go__go_6_3_3_v_7 = $__tco_6;
goto tco_loop_go__go_6_3_3;;
$__t5 = null;
goto end_branch_5;;
};
if ($v1_8_4 instanceof \Data\Ordering\Data_Ordering_GT) {
$__tco_7 = ($v_7)->{'value5'};
$__tco_var_go__go_6_3_3_v_7 = $__tco_7;
goto tco_loop_go__go_6_3_3;;
$__t5 = null;
goto end_branch_5;;
};
if ($v1_8_4 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t5 = new \Data\Maybe\Data_Maybe_Just(($v_7)->{'value3'});
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$v1_7_4 = ($go__go_6_3)($m2_4);
$__t5 = null;;
if ($v1_7_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = false;
goto end_branch_5;;
};
if ($v1_7_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($dictEq_1)->{'eq'})(($m1_3)->{'value3'}))(($v1_7_4)->{'value0'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($go__go_2_0)(($m1_3)->{'value4'}))($m2_4)))((($go__go_2_0)(($m1_3)->{'value5'}))($m2_4)));
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t1 = $__t5;
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
};
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_isSubmap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_ismajSubmap';

// Data_Map_Internal_isEmpty
function majData_majMap_majInternal_ismajEmpty($v_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_ismajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = true;
goto end_branch_0;;
};
  $__t0 = false;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_isEmpty'] = __NAMESPACE__ . '\\majData_majMap_majInternal_ismajEmpty';

// Data_Map_Internal_intersectionWith
function majData_majMap_majInternal_intersectionmajWith($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_intersectionmajWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($app_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m1_3) use ($app_2, $compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_4) use ($app_2, $compare_1_0, $m1_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $app_2, $m1_3, $m2_4);
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
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_intersectionWith'] = __NAMESPACE__ . '\\majData_majMap_majInternal_intersectionmajWith';

// Data_Map_Internal_intersection
function majData_majMap_majInternal_intersection($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_intersection';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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
$GLOBALS['Data_Map_Internal_intersection'] = __NAMESPACE__ . '\\majData_majMap_majInternal_intersection';

// Data_Map_Internal_insertWith
function majData_majMap_majInternal_insertmajWith($dictOrd_0, $app_1 = null, $k_2 = null, $v_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_insertmajWith';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $go__go_4_0 = null;
  $go__go_4_0 = function($v1_5) use ($app_1, $dictOrd_0, &$go__go_4_0, $k_2, $v_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $k_2, $v_3, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_1;;
};
  if ($v1_5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_6_2 = ((($dictOrd_0)->{'compare'})($k_2))(($v1_5)->{'value2'});
$__t3 = null;;
if ($v2_6_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_5)->{'value2'}, ($v1_5)->{'value3'}, ($go__go_4_0)(($v1_5)->{'value4'}), ($v1_5)->{'value5'});
goto end_branch_3;;
};
if ($v2_6_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_5)->{'value2'}, ($v1_5)->{'value3'}, ($v1_5)->{'value4'}, ($go__go_4_0)(($v1_5)->{'value5'}));
goto end_branch_3;;
};
if ($v2_6_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = new \Data\Map\Internal\Data_Map_Internal_Node(($v1_5)->{'value0'}, ($v1_5)->{'value1'}, $k_2, (($app_1)(($v1_5)->{'value3'}))($v_3), ($v1_5)->{'value4'}, ($v1_5)->{'value5'});
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
  $__res = $go__go_4_0;
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Map_Internal_insertWith'] = __NAMESPACE__ . '\\majData_majMap_majInternal_insertmajWith';

// Data_Map_Internal_insert
function majData_majMap_majInternal_insert($dictOrd_0, $k_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_insert';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__go_3_0 = null;
  $go__go_3_0 = function($v1_4) use ($dictOrd_0, &$go__go_3_0, $k_1, $v_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(1, 1, $k_1, $v_2, new \Data\Map\Internal\Data_Map_Internal_Leaf(), new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_1;;
};
  if ($v1_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v2_5_2 = ((($dictOrd_0)->{'compare'})($k_1))(($v1_4)->{'value2'});
$__t3 = null;;
if ($v2_5_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_4)->{'value2'}, ($v1_4)->{'value3'}, ($go__go_3_0)(($v1_4)->{'value4'}), ($v1_4)->{'value5'});
goto end_branch_3;;
};
if ($v2_5_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_4)->{'value2'}, ($v1_4)->{'value3'}, ($v1_4)->{'value4'}, ($go__go_3_0)(($v1_4)->{'value5'}));
goto end_branch_3;;
};
if ($v2_5_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = new \Data\Map\Internal\Data_Map_Internal_Node(($v1_4)->{'value0'}, ($v1_4)->{'value1'}, $k_1, $v_2, ($v1_4)->{'value4'}, ($v1_4)->{'value5'});
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
  $__res = $go__go_3_0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_insert'] = __NAMESPACE__ . '\\majData_majMap_majInternal_insert';

// Data_Map_Internal_functorMap
$GLOBALS['Data_Map_Internal_functorMap'] = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $go__go_1_0 = null;
  $go__go_1_0 = function($v_2) use ($f_0, &$go__go_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(($v_2)->{'value0'}, ($v_2)->{'value1'}, ($v_2)->{'value2'}, ($f_0)(($v_2)->{'value3'}), ($go__go_1_0)(($v_2)->{'value4'}), ($go__go_1_0)(($v_2)->{'value5'}));
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
  $__res = $go__go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_functorWithIndexMap
$GLOBALS['Data_Map_Internal_functorWithIndexMap'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $go__go_1_0 = null;
  $go__go_1_0 = function($v_2) use ($f_0, &$go__go_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Node(($v_2)->{'value0'}, ($v_2)->{'value1'}, ($v_2)->{'value2'}, (($f_0)(($v_2)->{'value2'}))(($v_2)->{'value3'}), ($go__go_1_0)(($v_2)->{'value4'}), ($go__go_1_0)(($v_2)->{'value5'}));
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
  $__res = $go__go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_foldableMap
$GLOBALS['Data_Map_Internal_foldableMap'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use ($f_0, &$go__go_2_0) {
  $__fn = function($m_prime_3, $__local_var_4 = null) use ($f_0, &$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($m_prime_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = $__local_var_4;
goto end_branch_1;;
};
  if ($m_prime_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = ($go__go_2_0)(($m_prime_3)->{'value4'}, (($f_0)(($m_prime_3)->{'value3'}))(($go__go_2_0)(($m_prime_3)->{'value5'}, $__local_var_4)));
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
  $__res = function($m_3) use (&$go__go_2_0, $z_1) {
  $__num = \func_num_args();
  $__res = ($go__go_2_0)($m_3, $z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_2 = null;
  $go__go_2_2 = (function() use ($f_0, &$go__go_2_2) {
  $__fn = function($z_prime_3, $__local_var_4 = null) use ($f_0, &$go__go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if ($__local_var_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = $z_prime_3;
goto end_branch_3;;
};
  if ($__local_var_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = ($go__go_2_2)((($f_0)(($go__go_2_2)($z_prime_3, ($__local_var_4)->{'value4'})))(($__local_var_4)->{'value3'}), ($__local_var_4)->{'value5'});
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
  $__res = function($m_3) use (&$go__go_2_2, $z_1) {
  $__num = \func_num_args();
  $__res = ($go__go_2_2)($z_1, $m_3);
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
  $mempty_1_4 = ($dictMonoid_0)->{'mempty'};
  $__local_var_2_5 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($f_3) use ($__local_var_2_5, $mempty_1_4) {
  $__num = \func_num_args();
  $go__go_4_6 = null;
  $go__go_4_6 = function($v_5) use ($__local_var_2_5, $f_3, &$go__go_4_6, $mempty_1_4) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t7 = $mempty_1_4;
goto end_branch_7;;
};
  if ($v_5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t7 = ((($__local_var_2_5)->{'append'})(($go__go_4_6)(($v_5)->{'value4'})))(((($__local_var_2_5)->{'append'})(($f_3)(($v_5)->{'value3'})))(($go__go_4_6)(($v_5)->{'value5'})));
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__go_4_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_foldableWithIndexMap
$GLOBALS['Data_Map_Internal_foldableWithIndexMap'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use ($f_0, &$go__go_2_0) {
  $__fn = function($m_prime_3, $__local_var_4 = null) use ($f_0, &$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($m_prime_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = $__local_var_4;
goto end_branch_1;;
};
  if ($m_prime_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = ($go__go_2_0)(($m_prime_3)->{'value4'}, ((($f_0)(($m_prime_3)->{'value2'}))(($m_prime_3)->{'value3'}))(($go__go_2_0)(($m_prime_3)->{'value5'}, $__local_var_4)));
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
  $__res = function($m_3) use (&$go__go_2_0, $z_1) {
  $__num = \func_num_args();
  $__res = ($go__go_2_0)($m_3, $z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_2 = null;
  $go__go_2_2 = (function() use ($f_0, &$go__go_2_2) {
  $__fn = function($z_prime_3, $__local_var_4 = null) use ($f_0, &$go__go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if ($__local_var_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = $z_prime_3;
goto end_branch_3;;
};
  if ($__local_var_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t3 = ($go__go_2_2)(((($f_0)(($__local_var_4)->{'value2'}))(($go__go_2_2)($z_prime_3, ($__local_var_4)->{'value4'})))(($__local_var_4)->{'value3'}), ($__local_var_4)->{'value5'});
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
  $__res = function($m_3) use (&$go__go_2_2, $z_1) {
  $__num = \func_num_args();
  $__res = ($go__go_2_2)($z_1, $m_3);
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
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_4 = ($dictMonoid_0)->{'mempty'};
  $__local_var_2_5 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($f_3) use ($__local_var_2_5, $mempty_1_4) {
  $__num = \func_num_args();
  $go__go_4_6 = null;
  $go__go_4_6 = function($v_5) use ($__local_var_2_5, $f_3, &$go__go_4_6, $mempty_1_4) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($v_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t7 = $mempty_1_4;
goto end_branch_7;;
};
  if ($v_5 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t7 = ((($__local_var_2_5)->{'append'})(($go__go_4_6)(($v_5)->{'value4'})))(((($__local_var_2_5)->{'append'})((($f_3)(($v_5)->{'value2'}))(($v_5)->{'value3'})))(($go__go_4_6)(($v_5)->{'value5'})));
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__go_4_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_foldableMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_keys_closure
$GLOBALS['Data_Map_Internal_keys_closure'] = ((($GLOBALS['Data_Map_Internal_foldableWithIndexMap'])->{'foldrWithIndex'})(function($k_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($k_0) {
  $__num = \func_num_args();
  $__res = function($acc_2) use ($k_0) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($k_0, $acc_2);
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
}))(new \Data\List\Types\Data_List_Types_Nil());

// Data_Map_Internal_keys
function majData_majMap_majInternal_keys($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_keys';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_keys_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_keys'] = __NAMESPACE__ . '\\majData_majMap_majInternal_keys';

// Data_Map_Internal_traversableMap
$GLOBALS['Data_Map_Internal_traversableMap'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $go__go_3_1 = null;
  $go__go_3_1 = function($v_4) use ($Apply0_1_0, $dictApplicative_0, $f_2, &$go__go_3_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = (($dictApplicative_0)->{'pure'})(new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_2;;
};
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__local_var_5_3 = ($v_4)->{'value0'};
$__local_var_6_4 = ($v_4)->{'value2'};
$__local_var_7_5 = ($v_4)->{'value1'};
$__t2 = ((($Apply0_1_0)->{'apply'})(((($Apply0_1_0)->{'apply'})(((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})(function($l_prime_8) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = function($v_prime_9) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5, $l_prime_8) {
  $__num = \func_num_args();
  $__res = function($r_prime_10) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5, $l_prime_8, $v_prime_9) {
  $__num = \func_num_args();
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node($__local_var_5_3, $__local_var_7_5, $__local_var_6_4, $v_prime_9, $l_prime_8, $r_prime_10);
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
}))(($go__go_3_1)(($v_4)->{'value4'}))))(($f_2)(($v_4)->{'value3'}))))(($go__go_3_1)(($v_4)->{'value5'}));
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
  $__res = $go__go_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Map_Internal_traversableMap'])->{'traverse'})($dictApplicative_0))($GLOBALS['Data_Map_Internal_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_foldableMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_traversableWithIndexMap
$GLOBALS['Data_Map_Internal_traversableWithIndexMap'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $go__go_3_1 = null;
  $go__go_3_1 = function($v_4) use ($Apply0_1_0, $dictApplicative_0, $f_2, &$go__go_3_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = (($dictApplicative_0)->{'pure'})(new \Data\Map\Internal\Data_Map_Internal_Leaf());
goto end_branch_2;;
};
  if ($v_4 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__local_var_5_3 = ($v_4)->{'value0'};
$__local_var_6_4 = ($v_4)->{'value2'};
$__local_var_7_5 = ($v_4)->{'value1'};
$__t2 = ((($Apply0_1_0)->{'apply'})(((($Apply0_1_0)->{'apply'})(((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})(function($l_prime_8) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = function($v_prime_9) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5, $l_prime_8) {
  $__num = \func_num_args();
  $__res = function($r_prime_10) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5, $l_prime_8, $v_prime_9) {
  $__num = \func_num_args();
  $__res = new \Data\Map\Internal\Data_Map_Internal_Node($__local_var_5_3, $__local_var_7_5, $__local_var_6_4, $v_prime_9, $l_prime_8, $r_prime_10);
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
}))(($go__go_3_1)(($v_4)->{'value4'}))))((($f_2)($__local_var_6_4))(($v_4)->{'value3'}))))(($go__go_3_1)(($v_4)->{'value5'}));
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
  $__res = $go__go_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorWithIndexMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_foldableWithIndexMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_traversableMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_values_closure
$GLOBALS['Data_Map_Internal_values_closure'] = ((($GLOBALS['Data_Map_Internal_foldableMap'])->{'foldr'})($GLOBALS['Data_List_Types_Cons']))(new \Data\List\Types\Data_List_Types_Nil());

// Data_Map_Internal_values
function majData_majMap_majInternal_values($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_values';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Map_Internal_values_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_values'] = __NAMESPACE__ . '\\majData_majMap_majInternal_values';

// Data_Map_Internal_foldSubmapBy
function majData_majMap_majInternal_foldmajSubmapmajBy($dictOrd_0, $appendFn_1 = null, $memptyValue_2 = null, $kmin_3 = null, $kmax_4 = null, $f_5 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_foldmajSubmapmajBy';
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__t0 = null;;
  if ($kmin_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_6_1 = ($kmin_3)->{'value0'};
$__t0 = function($k_7) use ($__local_var_6_1, $dictOrd_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (((($dictOrd_0)->{'compare'})($k_7))($__local_var_6_1) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t2 = true;
goto end_branch_2;;
};
  $__t2 = false;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_0;;
};
  if ($kmin_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = function($v_6) {
  $__num = \func_num_args();
  $__res = false;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $tooSmall_6_0 = $__t0;
  $__t4 = null;;
  if ($kmax_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_7_5 = ($kmax_4)->{'value0'};
$__t4 = function($k_8) use ($__local_var_7_5, $dictOrd_0) {
  $__num = \func_num_args();
  $__t6 = null;;
  if (((($dictOrd_0)->{'compare'})($k_8))($__local_var_7_5) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t6 = true;
goto end_branch_6;;
};
  $__t6 = false;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_4;;
};
  if ($kmax_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = function($v_7) {
  $__num = \func_num_args();
  $__res = false;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $tooLarge_7_4 = $__t4;
  $__t8 = null;;
  if ($kmin_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = null;;
if ($kmax_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_8_10 = ($kmax_4)->{'value0'};
$__local_var_9_11 = ($kmin_3)->{'value0'};
$__t9 = function($k_10) use ($__local_var_8_10, $__local_var_9_11, $dictOrd_0) {
  $__num = \func_num_args();
  $__t12 = null;;
  if (((($dictOrd_0)->{'compare'})($__local_var_9_11))($k_10) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t12 = false;
goto end_branch_12;;
};
  $__t12 = true;
  end_branch_12:;
  $__t13 = null;;
  if (((($dictOrd_0)->{'compare'})($k_10))($__local_var_8_10) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t13 = false;
goto end_branch_13;;
};
  $__t13 = true;
  end_branch_13:;
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})($__t12))($__t13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_9;;
};
if ($kmax_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__local_var_8_14 = ($kmin_3)->{'value0'};
$__t9 = function($k_9) use ($__local_var_8_14, $dictOrd_0) {
  $__num = \func_num_args();
  $__t15 = null;;
  if (((($dictOrd_0)->{'compare'})($__local_var_8_14))($k_9) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t15 = false;
goto end_branch_15;;
};
  $__t15 = true;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_9;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
end_branch_9:;
$__t8 = $__t9;
goto end_branch_8;;
};
  if ($kmin_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t16 = null;;
if ($kmax_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_8_17 = ($kmax_4)->{'value0'};
$__t16 = function($k_9) use ($__local_var_8_17, $dictOrd_0) {
  $__num = \func_num_args();
  $__t18 = null;;
  if (((($dictOrd_0)->{'compare'})($k_9))($__local_var_8_17) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t18 = false;
goto end_branch_18;;
};
  $__t18 = true;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_16;;
};
if ($kmax_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t16 = function($v_8) {
  $__num = \func_num_args();
  $__res = true;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_16;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t16 = null;
end_branch_16:;
$__t8 = $__t16;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $inBounds_8_8 = $__t8;
  $go__go_9_20 = null;
  $go__go_9_20 = function($v_10) use ($appendFn_1, $f_5, &$go__go_9_20, $inBounds_8_8, $memptyValue_2, $tooLarge_7_4, $tooSmall_6_0) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v_10 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t21 = $memptyValue_2;
goto end_branch_21;;
};
  if ($v_10 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t22 = null;;
if (($tooSmall_6_0)(($v_10)->{'value2'})) {
$__t22 = $memptyValue_2;
goto end_branch_22;;
};
$__t22 = ($go__go_9_20)(($v_10)->{'value4'});
end_branch_22:;
$__t23 = null;;
if (($inBounds_8_8)(($v_10)->{'value2'})) {
$__t23 = (($f_5)(($v_10)->{'value2'}))(($v_10)->{'value3'});
goto end_branch_23;;
};
$__t23 = $memptyValue_2;
end_branch_23:;
$__t24 = null;;
if (($tooLarge_7_4)(($v_10)->{'value2'})) {
$__t24 = $memptyValue_2;
goto end_branch_24;;
};
$__t24 = ($go__go_9_20)(($v_10)->{'value5'});
end_branch_24:;
$__t21 = (($appendFn_1)((($appendFn_1)($__t22))($__t23)))($__t24);
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__go_9_20;
  goto __end;;
  __end:
  return 6 < $__num ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
}
$GLOBALS['Data_Map_Internal_foldSubmapBy'] = __NAMESPACE__ . '\\majData_majMap_majInternal_foldmajSubmapmajBy';

// Data_Map_Internal_foldSubmap
function majData_majMap_majInternal_foldmajSubmap($dictOrd_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_foldmajSubmap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Map_Internal_foldSubmapBy'])($dictOrd_0))(((($dictMonoid_1)->{'Semigroup0'})(null))->{'append'}))(($dictMonoid_1)->{'mempty'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_foldSubmap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_foldmajSubmap';

// Data_Map_Internal_findMin
function majData_majMap_majInternal_findmajMin($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_findmajMin';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_Map_Internal_findMin_v_0 = $v_0;
  tco_loop_Data_Map_Internal_findMin:;
  $v_0 = $__tco_var_Data_Map_Internal_findMin_v_0;
  $__t0 = null;;
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if (($v_0)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = new \Data\Maybe\Data_Maybe_Just((object)["key" => ($v_0)->{'value2'}, "value" => ($v_0)->{'value3'}]);
goto end_branch_2;;
};
$__tco_1 = ($v_0)->{'value4'};
$__tco_var_Data_Map_Internal_findMin_v_0 = $__tco_1;
goto tco_loop_Data_Map_Internal_findMin;;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_findMin'] = __NAMESPACE__ . '\\majData_majMap_majInternal_findmajMin';

// Data_Map_Internal_lookupGT
function majData_majMap_majInternal_lookupmajGmajT($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_lookupmajGmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_2 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$v2_5_4 = ($go__go_2_0)(($v_3)->{'value4'});
$__t5 = null;;
if ($v2_5_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = new \Data\Maybe\Data_Maybe_Just((object)["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
goto end_branch_5;;
};
$__t5 = $v2_5_4;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($go__go_2_0)(($v_3)->{'value5'});
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = \Data\Map\Internal\majData_majMap_majInternal_findmajMin(($v_3)->{'value5'});
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_lookupGT'] = __NAMESPACE__ . '\\majData_majMap_majInternal_lookupmajGmajT';

// Data_Map_Internal_findMax
function majData_majMap_majInternal_findmajMax($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_findmajMax';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_Map_Internal_findMax_v_0 = $v_0;
  tco_loop_Data_Map_Internal_findMax:;
  $v_0 = $__tco_var_Data_Map_Internal_findMax_v_0;
  $__t0 = null;;
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if (($v_0)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = new \Data\Maybe\Data_Maybe_Just((object)["key" => ($v_0)->{'value2'}, "value" => ($v_0)->{'value3'}]);
goto end_branch_2;;
};
$__tco_1 = ($v_0)->{'value5'};
$__tco_var_Data_Map_Internal_findMax_v_0 = $__tco_1;
goto tco_loop_Data_Map_Internal_findMax;;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_findMax'] = __NAMESPACE__ . '\\majData_majMap_majInternal_findmajMax';

// Data_Map_Internal_lookupLT
function majData_majMap_majInternal_lookupmajLmajT($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_lookupmajLmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_2 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($go__go_2_0)(($v_3)->{'value4'});
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$v2_5_4 = ($go__go_2_0)(($v_3)->{'value5'});
$__t5 = null;;
if ($v2_5_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = new \Data\Maybe\Data_Maybe_Just((object)["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
goto end_branch_5;;
};
$__t5 = $v2_5_4;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = \Data\Map\Internal\majData_majMap_majInternal_findmajMax(($v_3)->{'value4'});
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_lookupLT'] = __NAMESPACE__ . '\\majData_majMap_majInternal_lookupmajLmajT';

// Data_Map_Internal_filterWithKey
function majData_majMap_majInternal_filtermajWithmajKey($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_filtermajWithmajKey';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($f_1, &$go__go_2_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if ((($f_1)(($v_3)->{'value2'}))(($v_3)->{'value3'})) {
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($go__go_2_0)(($v_3)->{'value4'}), ($go__go_2_0)(($v_3)->{'value5'}));
goto end_branch_2;;
};
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($go__go_2_0)(($v_3)->{'value4'}), ($go__go_2_0)(($v_3)->{'value5'}));
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
};
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_filterWithKey'] = __NAMESPACE__ . '\\majData_majMap_majInternal_filtermajWithmajKey';

// Data_Map_Internal_filterKeys
function majData_majMap_majInternal_filtermajKeys($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_filtermajKeys';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($f_1, &$go__go_2_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if (($f_1)(($v_3)->{'value2'})) {
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($go__go_2_0)(($v_3)->{'value4'}), ($go__go_2_0)(($v_3)->{'value5'}));
goto end_branch_2;;
};
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($go__go_2_0)(($v_3)->{'value4'}), ($go__go_2_0)(($v_3)->{'value5'}));
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
};
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_filterKeys'] = __NAMESPACE__ . '\\majData_majMap_majInternal_filtermajKeys';

// Data_Map_Internal_filter
function majData_majMap_majInternal_filter($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_filter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Map_Internal_filterWithKey'])($dictOrd_0)))($GLOBALS['Data_Function_const']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_filter'] = __NAMESPACE__ . '\\majData_majMap_majInternal_filter';

// Data_Map_Internal_eqMap
function majData_majMap_majInternal_eqmajMap($dictEq_0, $dictEq1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_eqmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq" => function($xs_2) use ($dictEq1_1, $dictEq_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($dictEq1_1, $dictEq_0, $xs_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($xs_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = null;;
if ($ys_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = true;
goto end_branch_1;;
};
$__t1 = false;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($xs_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if (($ys_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node && (($xs_2)->{'value1'} === ($ys_3)->{'value1'}))) {
$__t2 = ((((($GLOBALS['Data_Map_Internal_eqMapIter'])($dictEq_0))($dictEq1_1))->{'eq'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_2, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_3, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()));
goto end_branch_2;;
};
$__t2 = false;
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_eqMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_eqmajMap';

// Data_Map_Internal_ordMap
function majData_majMap_majInternal_ordmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_ordmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ordMapIter1_1_0 = ($GLOBALS['Data_Map_Internal_ordMapIter'])($dictOrd_0);
  $eqMap1_2_1 = ($GLOBALS['Data_Map_Internal_eqMap'])((($dictOrd_0)->{'Eq0'})(null));
  $__res = function($dictOrd1_3) use ($eqMap1_2_1, $ordMapIter1_1_0) {
  $__num = \func_num_args();
  $eqMap2_4_2 = ($eqMap1_2_1)((($dictOrd1_3)->{'Eq0'})(null));
  $__res = (object)["compare" => function($xs_5) use ($dictOrd1_3, $ordMapIter1_1_0) {
  $__num = \func_num_args();
  $__res = function($ys_6) use ($dictOrd1_3, $ordMapIter1_1_0, $xs_5) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($xs_5 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t5 = null;;
if ($ys_6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  $__t3 = null;;
  if ($ys_6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_3;;
};
  $__t3 = (((($ordMapIter1_1_0)($dictOrd1_3))->{'compare'})(new \Data\Map\Internal\Data_Map_Internal_IterNode($xs_5, new \Data\Map\Internal\Data_Map_Internal_IterLeaf())))(new \Data\Map\Internal\Data_Map_Internal_IterNode($ys_6, new \Data\Map\Internal\Data_Map_Internal_IterLeaf()));
  end_branch_3:;
  $__t4 = $__t3;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar__unused_5) use ($eqMap2_4_2) {
  $__num = \func_num_args();
  $__res = $eqMap2_4_2;
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
$GLOBALS['Data_Map_Internal_ordMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_ordmajMap';

// Data_Map_Internal_eq1Map
function majData_majMap_majInternal_eq1majMap($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_eq1majMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq1" => function($dictEq1_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Map_Internal_eqMap'])($dictEq_0))($dictEq1_1))->{'eq'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_eq1Map'] = __NAMESPACE__ . '\\majData_majMap_majInternal_eq1majMap';

// Data_Map_Internal_ord1Map
function majData_majMap_majInternal_ord1majMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_ord1majMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ordMap1_1_0 = ($GLOBALS['Data_Map_Internal_ordMap'])($dictOrd_0);
  $__local_var_2_1 = (($dictOrd_0)->{'Eq0'})(null);
  $eq1Map1_2_1 = (object)["eq1" => function($dictEq1_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Map_Internal_eqMap'])($__local_var_2_1))($dictEq1_3))->{'eq'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare1" => function($dictOrd1_3) use ($ordMap1_1_0) {
  $__num = \func_num_args();
  $__res = (($ordMap1_1_0)($dictOrd1_3))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_3) use ($eq1Map1_2_1) {
  $__num = \func_num_args();
  $__res = $eq1Map1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_ord1Map'] = __NAMESPACE__ . '\\majData_majMap_majInternal_ord1majMap';

// Data_Map_Internal_empty
$GLOBALS['Data_Map_Internal_empty'] = new \Data\Map\Internal\Data_Map_Internal_Leaf();

// Data_Map_Internal_fromFoldable
function majData_majMap_majInternal_frommajFoldable($dictOrd_0, $dictFoldable_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_frommajFoldable';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_1)->{'foldl'})(function($m_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictOrd_0, $m_2) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0))(($v_3)->{'value0'}))(($v_3)->{'value1'}))($m_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_fromFoldable'] = __NAMESPACE__ . '\\majData_majMap_majInternal_frommajFoldable';

// Data_Map_Internal_fromFoldableWith
function majData_majMap_majInternal_frommajFoldablemajWith($dictOrd_0, $dictFoldable_1 = null, $f_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_frommajFoldablemajWith';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $f_prime_3_0 = (($GLOBALS['Data_Map_Internal_insertWith'])($dictOrd_0))(function($b_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($b_3, $f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = ((($dictFoldable_1)->{'foldl'})(function($m_4) use ($f_prime_3_0) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($f_prime_3_0, $m_4) {
  $__num = \func_num_args();
  $__res = ((($f_prime_3_0)(($v_5)->{'value0'}))(($v_5)->{'value1'}))($m_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Map_Internal_fromFoldableWith'] = __NAMESPACE__ . '\\majData_majMap_majInternal_frommajFoldablemajWith';

// Data_Map_Internal_fromFoldableWithIndex
function majData_majMap_majInternal_frommajFoldablemajWithmajIndex($dictOrd_0, $dictFoldableWithIndex_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_frommajFoldablemajWithmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldableWithIndex_1)->{'foldlWithIndex'})(function($k_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($m_3) use ($dictOrd_0, $k_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictOrd_0, $k_2, $m_3) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0))($k_2))($v_4))($m_3);
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
}))(new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_fromFoldableWithIndex'] = __NAMESPACE__ . '\\majData_majMap_majInternal_frommajFoldablemajWithmajIndex';

// Data_Map_Internal_monoidSemigroupMap
function majData_majMap_majInternal_monoidmajSemigroupmajMap($_dollar__unused_0, $dictOrd_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_monoidmajSemigroupmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupMap2_2_0 = (($GLOBALS['Data_Map_Internal_semigroupMap'])(null))($dictOrd_1);
  $__res = function($dictSemigroup_3) use ($semigroupMap2_2_0) {
  $__num = \func_num_args();
  $semigroupMap3_4_1 = ($semigroupMap2_2_0)($dictSemigroup_3);
  $__res = (object)["mempty" => new \Data\Map\Internal\Data_Map_Internal_Leaf(), "Semigroup0" => function($_dollar__unused_5) use ($semigroupMap3_4_1) {
  $__num = \func_num_args();
  $__res = $semigroupMap3_4_1;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_monoidSemigroupMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_monoidmajSemigroupmajMap';

// Data_Map_Internal_submap
function majData_majMap_majInternal_submap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_submap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $union1_1_0 = function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($kmin_2) use ($dictOrd_0, $union1_1_0) {
  $__num = \func_num_args();
  $__res = function($kmax_3) use ($dictOrd_0, $kmin_2, $union1_1_0) {
  $__num = \func_num_args();
  $__res = (((((($GLOBALS['Data_Map_Internal_foldSubmapBy'])($dictOrd_0))($union1_1_0))(new \Data\Map\Internal\Data_Map_Internal_Leaf()))($kmin_2))($kmax_3))($GLOBALS['Data_Map_Internal_singleton']);
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
$GLOBALS['Data_Map_Internal_submap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_submap';

// Data_Map_Internal_unions
function majData_majMap_majInternal_unions($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_unions';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $union1_1_0 = function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictFoldable_2) use ($union1_1_0) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_2)->{'foldl'})($union1_1_0))(new \Data\Map\Internal\Data_Map_Internal_Leaf());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_unions'] = __NAMESPACE__ . '\\majData_majMap_majInternal_unions';

// Data_Map_Internal_difference
function majData_majMap_majInternal_difference($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_difference';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $m1_2, $m2_3);
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
$GLOBALS['Data_Map_Internal_difference'] = __NAMESPACE__ . '\\majData_majMap_majInternal_difference';

// Data_Map_Internal_delete
function majData_majMap_majInternal_delete($dictOrd_0, $k_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_delete';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function($v_3) use ($dictOrd_0, &$go__go_2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = new \Data\Map\Internal\Data_Map_Internal_Leaf();
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_4_2 = ((($dictOrd_0)->{'compare'})($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($go__go_2_0)(($v_3)->{'value4'}), ($v_3)->{'value5'});
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($v_3)->{'value4'}, ($go__go_2_0)(($v_3)->{'value5'}));
goto end_branch_3;;
};
if ($v1_4_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($v_3)->{'value4'}, ($v_3)->{'value5'});
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
  $__res = $go__go_2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Map_Internal_delete'] = __NAMESPACE__ . '\\majData_majMap_majInternal_delete';

// Data_Map_Internal_checkValid
function majData_majMap_majInternal_checkmajValid($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_checkmajValid';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = function($v_2) use ($dictOrd_0, &$go__go_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = true;
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t2 = null;;
if (($v_2)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = null;;
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t3 = true;
goto end_branch_3;;
};
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t4 = null;;
if (((($dictOrd_0)->{'compare'})((($v_2)->{'value5'})->{'value2'}))(($v_2)->{'value2'}) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = true;
goto end_branch_4;;
};
$__t4 = false;
end_branch_4:;
$__t3 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($v_2)->{'value0'} === 2)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($v_2)->{'value5'})->{'value0'} === 1)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Map_Internal_greaterThan'])(($v_2)->{'value1'}))((($v_2)->{'value5'})->{'value1'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})($__t4))(($go__go_1_0)(($v_2)->{'value5'})))));
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
if (($v_2)->{'value4'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t5 = null;;
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t6 = null;;
if (((($dictOrd_0)->{'compare'})((($v_2)->{'value4'})->{'value2'}))(($v_2)->{'value2'}) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t6 = true;
goto end_branch_6;;
};
$__t6 = false;
end_branch_6:;
$__t5 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($v_2)->{'value0'} === 2)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($v_2)->{'value4'})->{'value0'} === 1)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Map_Internal_greaterThan'])(($v_2)->{'value1'}))((($v_2)->{'value4'})->{'value1'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})($__t6))(($go__go_1_0)(($v_2)->{'value4'})))));
goto end_branch_5;;
};
if (($v_2)->{'value5'} instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t7 = null;;
if (((($dictOrd_0)->{'compare'})((($v_2)->{'value5'})->{'value2'}))(($v_2)->{'value2'}) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t7 = true;
goto end_branch_7;;
};
$__t7 = false;
end_branch_7:;
$__t8 = null;;
if (((($dictOrd_0)->{'compare'})((($v_2)->{'value4'})->{'value2'}))(($v_2)->{'value2'}) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = true;
goto end_branch_8;;
};
$__t8 = false;
end_branch_8:;
$__t5 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Map_Internal_greaterThan'])(($v_2)->{'value0'}))((($v_2)->{'value5'})->{'value0'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})($__t7))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Map_Internal_greaterThan'])(($v_2)->{'value0'}))((($v_2)->{'value4'})->{'value0'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})($__t8))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})((($GLOBALS['Data_Map_Internal_lessThan'])(($GLOBALS['Data_Map_Internal_abs'])(((($v_2)->{'value5'})->{'value0'} - (($v_2)->{'value4'})->{'value0'}))))(2)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((((($v_2)->{'value5'})->{'value1'} + (($v_2)->{'value4'})->{'value1'}) + 1) === ($v_2)->{'value1'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(($go__go_1_0)(($v_2)->{'value4'})))(($go__go_1_0)(($v_2)->{'value5'}))))))));
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t2 = $__t5;
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
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
};
  $__res = $go__go_1_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_checkValid'] = __NAMESPACE__ . '\\majData_majMap_majInternal_checkmajValid';

// Data_Map_Internal_catMaybes
function majData_majMap_majInternal_catmajMaybes($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_catmajMaybes';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($GLOBALS['Data_Map_Internal_mapMaybeWithKey'])($dictOrd_0), $GLOBALS['Data_Function_const'], $GLOBALS['Data_Map_Internal_identity1']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_catMaybes'] = __NAMESPACE__ . '\\majData_majMap_majInternal_catmajMaybes';

// Data_Map_Internal_applyMap
function majData_majMap_majInternal_applymajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_applymajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (object)["apply" => function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Map_Internal_identity2'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_applyMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_applymajMap';

// Data_Map_Internal_bindMap
function majData_majMap_majInternal_bindmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_bindmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $applyMap1_1_0 = (object)["apply" => function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Map_Internal_identity2'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($m_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictOrd_0, $m_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Map_Internal_mapMaybeWithKey'])($dictOrd_0))(function($k_4) use ($dictOrd_0, $f_3) {
  $__num = \func_num_args();
  $go__go_5_2 = null;
  $go__go_5_2 = function($v_6) use ($dictOrd_0, &$go__go_5_2, $k_4) {
  $__num = \func_num_args();
  $__tco_var_go__go_5_2_2_v_6 = $v_6;
  tco_loop_go__go_5_2_2:;
  $v_6 = $__tco_var_go__go_5_2_2_v_6;
  $__t2 = null;;
  if ($v_6 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_2;;
};
  if ($v_6 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$v1_7_3 = ((($dictOrd_0)->{'compare'})($k_4))(($v_6)->{'value2'});
$__t4 = null;;
if ($v1_7_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_5 = ($v_6)->{'value4'};
$__tco_var_go__go_5_2_2_v_6 = $__tco_5;
goto tco_loop_go__go_5_2_2;;
$__t4 = null;
goto end_branch_4;;
};
if ($v1_7_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__tco_6 = ($v_6)->{'value5'};
$__tco_var_go__go_5_2_2_v_6 = $__tco_6;
goto tco_loop_go__go_5_2_2;;
$__t4 = null;
goto end_branch_4;;
};
if ($v1_7_3 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t4 = new \Data\Maybe\Data_Maybe_Just(($v_6)->{'value3'});
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t2 = $__t4;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($go__go_5_2))($f_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($m_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_2) use ($applyMap1_1_0) {
  $__num = \func_num_args();
  $__res = $applyMap1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_bindMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_bindmajMap';

// Data_Map_Internal_anyWithKey
function majData_majMap_majInternal_anymajWithmajKey($predicate_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_anymajWithmajKey';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = function($v_2) use (&$go__go_1_0, $predicate_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = false;
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})((($predicate_0)(($v_2)->{'value2'}))(($v_2)->{'value3'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($go__go_1_0)(($v_2)->{'value4'})))(($go__go_1_0)(($v_2)->{'value5'})));
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
  $__res = $go__go_1_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_anyWithKey'] = __NAMESPACE__ . '\\majData_majMap_majInternal_anymajWithmajKey';

// Data_Map_Internal_any
function majData_majMap_majInternal_any($predicate_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_any';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = function($v_2) use (&$go__go_1_0, $predicate_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Leaf) {
$__t1 = false;
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Map\Internal\Data_Map_Internal_Node) {
$__t1 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($predicate_0)(($v_2)->{'value3'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})(($go__go_1_0)(($v_2)->{'value4'})))(($go__go_1_0)(($v_2)->{'value5'})));
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
  $__res = $go__go_1_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_any'] = __NAMESPACE__ . '\\majData_majMap_majInternal_any';

// Data_Map_Internal_alter
function majData_majMap_majInternal_alter($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_alter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($f_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($compare_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($m_4) use ($compare_1_0, $f_2, $k_3) {
  $__num = \func_num_args();
  $v_5_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($compare_1_0, $k_3, $m_4);
  $v2_6_2 = ($f_2)(($v_5_1)->{'value0'});
  $__t3 = null;;
  if ($v2_6_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($v_5_1)->{'value1'}, ($v_5_1)->{'value2'});
goto end_branch_3;;
};
  if ($v2_6_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])($k_3, ($v2_6_2)->{'value0'}, ($v_5_1)->{'value1'}, ($v_5_1)->{'value2'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_alter'] = __NAMESPACE__ . '\\majData_majMap_majInternal_alter';

// Data_Map_Internal_altMap
function majData_majMap_majInternal_altmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_altmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (object)["alt" => function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_altMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_altmajMap';

// Data_Map_Internal_plusMap
function majData_majMap_majInternal_plusmajMap($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majMap_majInternal_plusmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $altMap1_1_0 = (object)["alt" => function($m1_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = function($m2_3) use ($compare_1_0, $m1_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => new \Data\Map\Internal\Data_Map_Internal_Leaf(), "Alt0" => function($_dollar__unused_2) use ($altMap1_1_0) {
  $__num = \func_num_args();
  $__res = $altMap1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Map_Internal_plusMap'] = __NAMESPACE__ . '\\majData_majMap_majInternal_plusmajMap';

