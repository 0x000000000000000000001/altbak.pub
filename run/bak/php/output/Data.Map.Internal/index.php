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


// Data_Map_Internal_greaterThan
$GLOBALS['Data_Map_Internal_greaterThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT"));
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1 = null, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (is_object((($__local_var_0_0)($a1_1))($a2_2)) && (((($__local_var_0_0)($a1_1))($a2_2))->{'tag'} === "GT"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Map_Internal_lessThanOrEq
$GLOBALS['Data_Map_Internal_lessThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT"));
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1 = null, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (is_object((($__local_var_0_0)($a1_1))($a2_2)) && (((($__local_var_0_0)($a1_1))($a2_2))->{'tag'} === "GT")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Map_Internal_identity
$GLOBALS['Data_Map_Internal_identity'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_lessThan
$GLOBALS['Data_Map_Internal_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT"));
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1 = null, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (is_object((($__local_var_0_0)($a1_1))($a2_2)) && (((($__local_var_0_0)($a1_1))($a2_2))->{'tag'} === "LT"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Map_Internal_abs
$GLOBALS['Data_Map_Internal_abs'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT"));
return function($x_1 = null) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (( ! (is_object((($__local_var_0_0)($x_1))(0)) && (((($__local_var_0_0)($x_1))(0))->{'tag'} === "LT")))) {
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

// Data_Map_Internal_Leaf
$GLOBALS['Data_Map_Internal_Leaf'] = ($GLOBALS['__phpurs_data0_Leaf'] ??= new Phpurs_Data0("Leaf"));

// Data_Map_Internal_Node
$GLOBALS['Data_Map_Internal_Node'] = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null, $value3 = null, $value4 = null, $value5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__res = new Phpurs_Data6("Node", $value0, $value1, $value2, $value3, $value4, $value5);
  goto __end;;
  __end:
  return $__num > 6 ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_IterLeaf
$GLOBALS['Data_Map_Internal_IterLeaf'] = ($GLOBALS['__phpurs_data0_IterLeaf'] ??= new Phpurs_Data0("IterLeaf"));

// Data_Map_Internal_IterEmit
$GLOBALS['Data_Map_Internal_IterEmit'] = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("IterEmit", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_IterNode
$GLOBALS['Data_Map_Internal_IterNode'] = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("IterNode", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_IterDone
$GLOBALS['Data_Map_Internal_IterDone'] = ($GLOBALS['__phpurs_data0_IterDone'] ??= new Phpurs_Data0("IterDone"));

// Data_Map_Internal_IterNext
$GLOBALS['Data_Map_Internal_IterNext'] = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("IterNext", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_Split
$GLOBALS['Data_Map_Internal_Split'] = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("Split", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_SplitLast
$GLOBALS['Data_Map_Internal_SplitLast'] = (function() {
  $__fn = function($value0 = null, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("SplitLast", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_unsafeNode
$GLOBALS['Data_Map_Internal_unsafeNode'] = (function() {
  $__fn = function($k_0 = null, $v_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ((is_object($l_2) && (($l_2)->{'tag'} === "Leaf"))) {
$__t1 = null;;
if ((is_object($r_3) && (($r_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data6("Node", 1, 1, $k_0, $v_1, $l_2, $r_3);
goto end_branch_1;;
};
if ((is_object($r_3) && (($r_3)->{'tag'} === "Node"))) {
$__t1 = new Phpurs_Data6("Node", (1 + ($r_3)->{'value0'}), (1 + ($r_3)->{'value1'}), $k_0, $v_1, $l_2, $r_3);
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ((is_object($l_2) && (($l_2)->{'tag'} === "Node"))) {
$__t2 = null;;
if ((is_object($r_3) && (($r_3)->{'tag'} === "Leaf"))) {
$__t2 = new Phpurs_Data6("Node", (1 + ($l_2)->{'value0'}), (1 + ($l_2)->{'value1'}), $k_0, $v_1, $l_2, $r_3);
goto end_branch_2;;
};
if ((is_object($r_3) && (($r_3)->{'tag'} === "Node"))) {
$__t3 = null;;
if ((($GLOBALS['Data_Map_Internal_greaterThan'])(($l_2)->{'value0'}))(($r_3)->{'value0'})) {
$__t3 = (1 + ($l_2)->{'value0'});
goto end_branch_3;;
};
$__t3 = (1 + ($r_3)->{'value0'});
end_branch_3:;
$__t2 = new Phpurs_Data6("Node", $__t3, ((1 + ($l_2)->{'value1'}) + ($r_3)->{'value1'}), $k_0, $v_1, $l_2, $r_3);
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
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_toMapIter
$GLOBALS['Data_Map_Internal_toMapIter'] = function($a_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("IterNode", $a_0, new Phpurs_Data0("IterLeaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_stepWith
$GLOBALS['Data_Map_Internal_stepWith'] = (function() {
  $__fn = function($f_0 = null, $next_1 = null, $done_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = function($v_4 = null) use ($done_2, $f_0, &$go__3_0, $next_1) {
  $__num = \func_num_args();
  $__tco_var_go__3_0_0_v_4 = $v_4;
  tco_loop_go__3_0_0:;
  $v_4 = $__tco_var_go__3_0_0_v_4;
  $__t0 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "IterLeaf"))) {
$__t0 = ($done_2)($GLOBALS['Data_Unit_unit']);
goto end_branch_0;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "IterEmit"))) {
$__t0 = ($next_1)(($v_4)->{'value0'}, ($v_4)->{'value1'}, ($v_4)->{'value2'});
goto end_branch_0;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "IterNode"))) {
$__tco_1 = (($f_0)(($v_4)->{'value1'}))(($v_4)->{'value0'});
$__tco_var_go__3_0_0_v_4 = $__tco_1;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__3_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_size
$GLOBALS['Data_Map_Internal_size'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Leaf"))) {
$__t0 = 0;
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Node"))) {
$__t0 = ($v_0)->{'value1'};
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

// Data_Map_Internal_singleton
$GLOBALS['Data_Map_Internal_singleton'] = (function() {
  $__fn = function($k_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data6("Node", 1, 1, $k_0, $v_1, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_unsafeBalancedNode
$GLOBALS['Data_Map_Internal_unsafeBalancedNode'] = (function() {
  $__fn = function($k_0 = null, $v_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ((is_object($l_2) && (($l_2)->{'tag'} === "Leaf"))) {
$__t1 = null;;
if ((is_object($r_3) && (($r_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data6("Node", 1, 1, $k_0, $v_1, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
goto end_branch_1;;
};
if (((is_object($r_3) && (($r_3)->{'tag'} === "Node")) && (($GLOBALS['Data_Map_Internal_greaterThan'])(($r_3)->{'value0'}))(1))) {
$__t2 = null;;
if ((function() use ($r_3, &$__fn) {
$__t3 = null;;
if ((is_object(($r_3)->{'value5'}) && ((($r_3)->{'value5'})->{'tag'} === "Leaf"))) {
$__t3 = 0;
goto end_branch_3;;
};
if ((is_object(($r_3)->{'value5'}) && ((($r_3)->{'value5'})->{'tag'} === "Node"))) {
$__t3 = (($r_3)->{'value5'})->{'value0'};
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
return ((is_object(($r_3)->{'value4'}) && ((($r_3)->{'value4'})->{'tag'} === "Node")) && (($GLOBALS['Data_Map_Internal_greaterThan'])((($r_3)->{'value4'})->{'value0'}))($__t3));
})()) {
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeNode'])((($r_3)->{'value4'})->{'value2'}, (($r_3)->{'value4'})->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, $l_2, (($r_3)->{'value4'})->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeNode'])(($r_3)->{'value2'}, ($r_3)->{'value3'}, (($r_3)->{'value4'})->{'value5'}, ($r_3)->{'value5'}));
goto end_branch_2;;
};
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeNode'])(($r_3)->{'value2'}, ($r_3)->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, $l_2, ($r_3)->{'value4'}), ($r_3)->{'value5'});
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
$__t1 = ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, $l_2, $r_3);
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ((is_object($l_2) && (($l_2)->{'tag'} === "Node"))) {
$__t4 = null;;
if ((is_object($r_3) && (($r_3)->{'tag'} === "Node"))) {
$__t5 = null;;
if ((($GLOBALS['Data_Map_Internal_greaterThan'])(($r_3)->{'value0'}))((($l_2)->{'value0'} + 1))) {
$__t6 = null;;
if ((function() use ($r_3, &$__fn) {
$__t7 = null;;
if ((is_object(($r_3)->{'value5'}) && ((($r_3)->{'value5'})->{'tag'} === "Leaf"))) {
$__t7 = 0;
goto end_branch_7;;
};
if ((is_object(($r_3)->{'value5'}) && ((($r_3)->{'value5'})->{'tag'} === "Node"))) {
$__t7 = (($r_3)->{'value5'})->{'value0'};
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
return ((is_object(($r_3)->{'value4'}) && ((($r_3)->{'value4'})->{'tag'} === "Node")) && (($GLOBALS['Data_Map_Internal_greaterThan'])((($r_3)->{'value4'})->{'value0'}))($__t7));
})()) {
$__t6 = ($GLOBALS['Data_Map_Internal_unsafeNode'])((($r_3)->{'value4'})->{'value2'}, (($r_3)->{'value4'})->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, $l_2, (($r_3)->{'value4'})->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeNode'])(($r_3)->{'value2'}, ($r_3)->{'value3'}, (($r_3)->{'value4'})->{'value5'}, ($r_3)->{'value5'}));
goto end_branch_6;;
};
$__t6 = ($GLOBALS['Data_Map_Internal_unsafeNode'])(($r_3)->{'value2'}, ($r_3)->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, $l_2, ($r_3)->{'value4'}), ($r_3)->{'value5'});
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
if ((($GLOBALS['Data_Map_Internal_greaterThan'])(($l_2)->{'value0'}))((($r_3)->{'value0'} + 1))) {
$__t8 = null;;
if ((function() use ($l_2, &$__fn) {
$__t9 = null;;
if ((is_object(($l_2)->{'value4'}) && ((($l_2)->{'value4'})->{'tag'} === "Leaf"))) {
$__t9 = 0;
goto end_branch_9;;
};
if ((is_object(($l_2)->{'value4'}) && ((($l_2)->{'value4'})->{'tag'} === "Node"))) {
$__t9 = (($l_2)->{'value4'})->{'value0'};
goto end_branch_9;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
end_branch_9:;
return ((is_object(($l_2)->{'value5'}) && ((($l_2)->{'value5'})->{'tag'} === "Node")) && (($GLOBALS['Data_Map_Internal_lessThanOrEq'])($__t9))((($l_2)->{'value5'})->{'value0'}));
})()) {
$__t8 = ($GLOBALS['Data_Map_Internal_unsafeNode'])((($l_2)->{'value5'})->{'value2'}, (($l_2)->{'value5'})->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])(($l_2)->{'value2'}, ($l_2)->{'value3'}, ($l_2)->{'value4'}, (($l_2)->{'value5'})->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, (($l_2)->{'value5'})->{'value5'}, $r_3));
goto end_branch_8;;
};
$__t8 = ($GLOBALS['Data_Map_Internal_unsafeNode'])(($l_2)->{'value2'}, ($l_2)->{'value3'}, ($l_2)->{'value4'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, ($l_2)->{'value5'}, $r_3));
end_branch_8:;
$__t5 = $__t8;
goto end_branch_5;;
};
$__t5 = ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, $l_2, $r_3);
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
if (((is_object($r_3) && (($r_3)->{'tag'} === "Leaf")) && (($GLOBALS['Data_Map_Internal_greaterThan'])(($l_2)->{'value0'}))(1))) {
$__t10 = null;;
if ((function() use ($l_2, &$__fn) {
$__t11 = null;;
if ((is_object(($l_2)->{'value4'}) && ((($l_2)->{'value4'})->{'tag'} === "Leaf"))) {
$__t11 = 0;
goto end_branch_11;;
};
if ((is_object(($l_2)->{'value4'}) && ((($l_2)->{'value4'})->{'tag'} === "Node"))) {
$__t11 = (($l_2)->{'value4'})->{'value0'};
goto end_branch_11;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
end_branch_11:;
return ((is_object(($l_2)->{'value5'}) && ((($l_2)->{'value5'})->{'tag'} === "Node")) && (($GLOBALS['Data_Map_Internal_lessThanOrEq'])($__t11))((($l_2)->{'value5'})->{'value0'}));
})()) {
$__t10 = ($GLOBALS['Data_Map_Internal_unsafeNode'])((($l_2)->{'value5'})->{'value2'}, (($l_2)->{'value5'})->{'value3'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])(($l_2)->{'value2'}, ($l_2)->{'value3'}, ($l_2)->{'value4'}, (($l_2)->{'value5'})->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, (($l_2)->{'value5'})->{'value5'}, $r_3));
goto end_branch_10;;
};
$__t10 = ($GLOBALS['Data_Map_Internal_unsafeNode'])(($l_2)->{'value2'}, ($l_2)->{'value3'}, ($l_2)->{'value4'}, ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, ($l_2)->{'value5'}, $r_3));
end_branch_10:;
$__t4 = $__t10;
goto end_branch_4;;
};
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeNode'])($k_0, $v_1, $l_2, $r_3);
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
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_unsafeSplit
$GLOBALS['Data_Map_Internal_unsafeSplit'] = (function() {
  $__fn = function($comp_0 = null, $k_1 = null, $m_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_Map_Internal_unsafeSplit_comp_0 = $comp_0;
  $__tco_var_Data_Map_Internal_unsafeSplit_k_1 = $k_1;
  $__tco_var_Data_Map_Internal_unsafeSplit_m_2 = $m_2;
  tco_loop_Data_Map_Internal_unsafeSplit:;
  $comp_0 = $__tco_var_Data_Map_Internal_unsafeSplit_comp_0;
  $k_1 = $__tco_var_Data_Map_Internal_unsafeSplit_k_1;
  $m_2 = $__tco_var_Data_Map_Internal_unsafeSplit_m_2;
  $__t0 = null;;
  if ((is_object($m_2) && (($m_2)->{'tag'} === "Leaf"))) {
$__t0 = new Phpurs_Data3("Split", new Phpurs_Data0("Nothing"), new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
goto end_branch_0;;
};
  if ((is_object($m_2) && (($m_2)->{'tag'} === "Node"))) {
$v_3_1 = (($comp_0)($k_1))(($m_2)->{'value2'});
$__t2 = null;;
if ((is_object($v_3_1) && (($v_3_1)->{'tag'} === "LT"))) {
$v1_4_3 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, $k_1, ($m_2)->{'value4'});
$__t2 = new Phpurs_Data3("Split", ($v1_4_3)->{'value0'}, ($v1_4_3)->{'value1'}, ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($m_2)->{'value2'}, ($m_2)->{'value3'}, ($v1_4_3)->{'value2'}, ($m_2)->{'value5'}));
goto end_branch_2;;
};
if ((is_object($v_3_1) && (($v_3_1)->{'tag'} === "GT"))) {
$v1_4_4 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, $k_1, ($m_2)->{'value5'});
$__t2 = new Phpurs_Data3("Split", ($v1_4_4)->{'value0'}, ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($m_2)->{'value2'}, ($m_2)->{'value3'}, ($m_2)->{'value4'}, ($v1_4_4)->{'value1'}), ($v1_4_4)->{'value2'});
goto end_branch_2;;
};
if ((is_object($v_3_1) && (($v_3_1)->{'tag'} === "EQ"))) {
$__t2 = new Phpurs_Data3("Split", new Phpurs_Data1("Just", ($m_2)->{'value3'}), ($m_2)->{'value4'}, ($m_2)->{'value5'});
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_unsafeSplitLast
$GLOBALS['Data_Map_Internal_unsafeSplitLast'] = (function() {
  $__fn = function($k_0 = null, $v_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_Map_Internal_unsafeSplitLast_k_0 = $k_0;
  $__tco_var_Data_Map_Internal_unsafeSplitLast_v_1 = $v_1;
  $__tco_var_Data_Map_Internal_unsafeSplitLast_l_2 = $l_2;
  $__tco_var_Data_Map_Internal_unsafeSplitLast_r_3 = $r_3;
  tco_loop_Data_Map_Internal_unsafeSplitLast:;
  $k_0 = $__tco_var_Data_Map_Internal_unsafeSplitLast_k_0;
  $v_1 = $__tco_var_Data_Map_Internal_unsafeSplitLast_v_1;
  $l_2 = $__tco_var_Data_Map_Internal_unsafeSplitLast_l_2;
  $r_3 = $__tco_var_Data_Map_Internal_unsafeSplitLast_r_3;
  $__t0 = null;;
  if ((is_object($r_3) && (($r_3)->{'tag'} === "Leaf"))) {
$__t0 = new Phpurs_Data3("SplitLast", $k_0, $v_1, $l_2);
goto end_branch_0;;
};
  if ((is_object($r_3) && (($r_3)->{'tag'} === "Node"))) {
$v1_4_1 = ($GLOBALS['Data_Map_Internal_unsafeSplitLast'])(($r_3)->{'value2'}, ($r_3)->{'value3'}, ($r_3)->{'value4'}, ($r_3)->{'value5'});
$__t0 = new Phpurs_Data3("SplitLast", ($v1_4_1)->{'value0'}, ($v1_4_1)->{'value1'}, ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])($k_0, $v_1, $l_2, ($v1_4_1)->{'value2'}));
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
})();

// Data_Map_Internal_unsafeJoinNodes
$GLOBALS['Data_Map_Internal_unsafeJoinNodes'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Leaf"))) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Node"))) {
$v2_2_1 = ($GLOBALS['Data_Map_Internal_unsafeSplitLast'])(($v_0)->{'value2'}, ($v_0)->{'value3'}, ($v_0)->{'value4'}, ($v_0)->{'value5'});
$__t0 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v2_2_1)->{'value0'}, ($v2_2_1)->{'value1'}, ($v2_2_1)->{'value2'}, $v1_1);
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

// Data_Map_Internal_unsafeDifference
$GLOBALS['Data_Map_Internal_unsafeDifference'] = (function() {
  $__fn = function($comp_0 = null, $l_1 = null, $r_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_Map_Internal_unsafeDifference_comp_0 = $comp_0;
  $__tco_var_Data_Map_Internal_unsafeDifference_l_1 = $l_1;
  $__tco_var_Data_Map_Internal_unsafeDifference_r_2 = $r_2;
  tco_loop_Data_Map_Internal_unsafeDifference:;
  $comp_0 = $__tco_var_Data_Map_Internal_unsafeDifference_comp_0;
  $l_1 = $__tco_var_Data_Map_Internal_unsafeDifference_l_1;
  $r_2 = $__tco_var_Data_Map_Internal_unsafeDifference_r_2;
  $__t0 = null;;
  if ((is_object($l_1) && (($l_1)->{'tag'} === "Leaf"))) {
$__t0 = new Phpurs_Data0("Leaf");
goto end_branch_0;;
};
  if ((is_object($r_2) && (($r_2)->{'tag'} === "Leaf"))) {
$__t0 = $l_1;
goto end_branch_0;;
};
  if ((is_object($r_2) && (($r_2)->{'tag'} === "Node"))) {
$v_3_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, ($r_2)->{'value2'}, $l_1);
$__t0 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($GLOBALS['Data_Map_Internal_unsafeDifference'])($comp_0, ($v_3_1)->{'value1'}, ($r_2)->{'value4'}), ($GLOBALS['Data_Map_Internal_unsafeDifference'])($comp_0, ($v_3_1)->{'value2'}, ($r_2)->{'value5'}));
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

// Data_Map_Internal_unsafeIntersectionWith
$GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] = (function() {
  $__fn = function($comp_0 = null, $app_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_Map_Internal_unsafeIntersectionWith_comp_0 = $comp_0;
  $__tco_var_Data_Map_Internal_unsafeIntersectionWith_app_1 = $app_1;
  $__tco_var_Data_Map_Internal_unsafeIntersectionWith_l_2 = $l_2;
  $__tco_var_Data_Map_Internal_unsafeIntersectionWith_r_3 = $r_3;
  tco_loop_Data_Map_Internal_unsafeIntersectionWith:;
  $comp_0 = $__tco_var_Data_Map_Internal_unsafeIntersectionWith_comp_0;
  $app_1 = $__tco_var_Data_Map_Internal_unsafeIntersectionWith_app_1;
  $l_2 = $__tco_var_Data_Map_Internal_unsafeIntersectionWith_l_2;
  $r_3 = $__tco_var_Data_Map_Internal_unsafeIntersectionWith_r_3;
  $__t0 = null;;
  if ((is_object($l_2) && (($l_2)->{'tag'} === "Leaf"))) {
$__t0 = new Phpurs_Data0("Leaf");
goto end_branch_0;;
};
  if ((is_object($r_3) && (($r_3)->{'tag'} === "Leaf"))) {
$__t0 = new Phpurs_Data0("Leaf");
goto end_branch_0;;
};
  if ((is_object($r_3) && (($r_3)->{'tag'} === "Node"))) {
$v_4_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, ($r_3)->{'value2'}, $l_2);
$l_prime_5_2 = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($comp_0, $app_1, ($v_4_1)->{'value1'}, ($r_3)->{'value4'});
$r_prime_6_3 = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($comp_0, $app_1, ($v_4_1)->{'value2'}, ($r_3)->{'value5'});
$__t4 = null;;
if ((is_object(($v_4_1)->{'value0'}) && ((($v_4_1)->{'value0'})->{'tag'} === "Just"))) {
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($r_3)->{'value2'}, (($app_1)((($v_4_1)->{'value0'})->{'value0'}))(($r_3)->{'value3'}), $l_prime_5_2, $r_prime_6_3);
goto end_branch_4;;
};
if ((is_object(($v_4_1)->{'value0'}) && ((($v_4_1)->{'value0'})->{'tag'} === "Nothing"))) {
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
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_unsafeUnionWith
$GLOBALS['Data_Map_Internal_unsafeUnionWith'] = (function() {
  $__fn = function($comp_0 = null, $app_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_Map_Internal_unsafeUnionWith_comp_0 = $comp_0;
  $__tco_var_Data_Map_Internal_unsafeUnionWith_app_1 = $app_1;
  $__tco_var_Data_Map_Internal_unsafeUnionWith_l_2 = $l_2;
  $__tco_var_Data_Map_Internal_unsafeUnionWith_r_3 = $r_3;
  tco_loop_Data_Map_Internal_unsafeUnionWith:;
  $comp_0 = $__tco_var_Data_Map_Internal_unsafeUnionWith_comp_0;
  $app_1 = $__tco_var_Data_Map_Internal_unsafeUnionWith_app_1;
  $l_2 = $__tco_var_Data_Map_Internal_unsafeUnionWith_l_2;
  $r_3 = $__tco_var_Data_Map_Internal_unsafeUnionWith_r_3;
  $__t0 = null;;
  if ((is_object($l_2) && (($l_2)->{'tag'} === "Leaf"))) {
$__t0 = $r_3;
goto end_branch_0;;
};
  if ((is_object($r_3) && (($r_3)->{'tag'} === "Leaf"))) {
$__t0 = $l_2;
goto end_branch_0;;
};
  if ((is_object($r_3) && (($r_3)->{'tag'} === "Node"))) {
$v_4_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($comp_0, ($r_3)->{'value2'}, $l_2);
$l_prime_5_2 = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($comp_0, $app_1, ($v_4_1)->{'value1'}, ($r_3)->{'value4'});
$r_prime_6_3 = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($comp_0, $app_1, ($v_4_1)->{'value2'}, ($r_3)->{'value5'});
$__t4 = null;;
if ((is_object(($v_4_1)->{'value0'}) && ((($v_4_1)->{'value0'})->{'tag'} === "Just"))) {
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($r_3)->{'value2'}, (($app_1)((($v_4_1)->{'value0'})->{'value0'}))(($r_3)->{'value3'}), $l_prime_5_2, $r_prime_6_3);
goto end_branch_4;;
};
if ((is_object(($v_4_1)->{'value0'}) && ((($v_4_1)->{'value0'})->{'tag'} === "Nothing"))) {
$__t4 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($r_3)->{'value2'}, ($r_3)->{'value3'}, $l_prime_5_2, $r_prime_6_3);
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
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_unionWith
$GLOBALS['Data_Map_Internal_unionWith'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($app_2 = null, $m1_3 = null, $m2_4 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $app_2, $m1_3, $m2_4);
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

// Data_Map_Internal_union
$GLOBALS['Data_Map_Internal_union'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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

// Data_Map_Internal_update
$GLOBALS['Data_Map_Internal_update'] = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null, $k_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = function($v_4 = null) use ($dictOrd_0, $f_1, &$go__3_0, $k_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Node"))) {
$v1_5_2 = ((($dictOrd_0)['compare'])($k_2))(($v_4)->{'value2'});
$__t3 = null;;
if ((is_object($v1_5_2) && (($v1_5_2)->{'tag'} === "LT"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_4)->{'value2'}, ($v_4)->{'value3'}, ($go__3_0)(($v_4)->{'value4'}), ($v_4)->{'value5'});
goto end_branch_3;;
};
if ((is_object($v1_5_2) && (($v1_5_2)->{'tag'} === "GT"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_4)->{'value2'}, ($v_4)->{'value3'}, ($v_4)->{'value4'}, ($go__3_0)(($v_4)->{'value5'}));
goto end_branch_3;;
};
if ((is_object($v1_5_2) && (($v1_5_2)->{'tag'} === "EQ"))) {
$v2_6_4 = ($f_1)(($v_4)->{'value3'});
$__t5 = null;;
if ((is_object($v2_6_4) && (($v2_6_4)->{'tag'} === "Nothing"))) {
$__t5 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($v_4)->{'value4'}, ($v_4)->{'value5'});
goto end_branch_5;;
};
if ((is_object($v2_6_4) && (($v2_6_4)->{'tag'} === "Just"))) {
$__t5 = new Phpurs_Data6("Node", ($v_4)->{'value0'}, ($v_4)->{'value1'}, ($v_4)->{'value2'}, ($v2_6_4)->{'value0'}, ($v_4)->{'value4'}, ($v_4)->{'value5'});
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
  $__res = $go__3_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_showTree
$GLOBALS['Data_Map_Internal_showTree'] = (function() {
  $__fn = function($dictShow_0 = null, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($dictShow1_1, $dictShow_0, &$go__2_0) {
  $__fn = function($ind_3 = null, $v_4 = null) use ($dictShow1_1, $dictShow_0, &$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Leaf"))) {
$__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])($ind_3))("Leaf");
goto end_branch_1;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Node"))) {
$__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])($ind_3))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("["))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($GLOBALS['Data_Show_showInt'])['show'])(($v_4)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("] "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow_0)['show'])(($v_4)->{'value2'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(" => "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow1_1)['show'])(($v_4)->{'value3'})))("
")))))))))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($go__2_0)(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])($ind_3))("    ")))(($v_4)->{'value4'})))("
")))((($go__2_0)(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])($ind_3))("    ")))(($v_4)->{'value5'})));
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
  $__res = ($go__2_0)("");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_semigroupMap
$GLOBALS['Data_Map_Internal_semigroupMap'] = (function() {
  $__fn = function($_dollar__unused_0 = null, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $compare_2_0 = ($dictOrd_1)['compare'];
  $__res = function($dictSemigroup_3 = null) use ($compare_2_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = ($dictSemigroup_3)['append'];
  $__res = ["append" => (function() use ($__local_var_4_1, $compare_2_0) {
  $__fn = function($m1_5 = null, $m2_6 = null) use ($__local_var_4_1, $compare_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_2_0, $__local_var_4_1, $m1_5, $m2_6);
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
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_pop
$GLOBALS['Data_Map_Internal_pop'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($k_2 = null, $m_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_4_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($compare_1_0, $k_2, $m_3);
  $__local_var_5_2 = ($v_4_1)->{'value1'};
  $__local_var_6_3 = ($v_4_1)->{'value2'};
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($a_7 = null) use ($__local_var_5_2, $__local_var_6_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", $a_7, ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])($__local_var_5_2, $__local_var_6_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4_1)->{'value0'});
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

// Data_Map_Internal_member
$GLOBALS['Data_Map_Internal_member'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($dictOrd_0, &$go__2_0, $k_1) {
  $__num = \func_num_args();
  $__tco_var_go__2_0_0_v_3 = $v_3;
  tco_loop_go__2_0_0:;
  $v_3 = $__tco_var_go__2_0_0_v_3;
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t0 = false;
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_1 = ((($dictOrd_0)['compare'])($k_1))(($v_3)->{'value2'});
$__t2 = null;;
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "LT"))) {
$__tco_3 = ($v_3)->{'value4'};
$__tco_var_go__2_0_0_v_3 = $__tco_3;
goto tco_loop_go__2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "GT"))) {
$__tco_4 = ($v_3)->{'value5'};
$__tco_var_go__2_0_0_v_3 = $__tco_4;
goto tco_loop_go__2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "EQ"))) {
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_mapMaybeWithKey
$GLOBALS['Data_Map_Internal_mapMaybeWithKey'] = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($f_1, &$go__2_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
goto end_branch_1;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v2_4_2 = (($f_1)(($v_3)->{'value2'}))(($v_3)->{'value3'});
$__t3 = null;;
if ((is_object($v2_4_2) && (($v2_4_2)->{'tag'} === "Just"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v2_4_2)->{'value0'}, ($go__2_0)(($v_3)->{'value4'}), ($go__2_0)(($v_3)->{'value5'}));
goto end_branch_3;;
};
if ((is_object($v2_4_2) && (($v2_4_2)->{'tag'} === "Nothing"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($go__2_0)(($v_3)->{'value4'}), ($go__2_0)(($v_3)->{'value5'}));
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_mapMaybe
$GLOBALS['Data_Map_Internal_mapMaybe'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Map_Internal_mapMaybeWithKey'])($dictOrd_0)))($GLOBALS['Data_Function_const']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_lookupLE
$GLOBALS['Data_Map_Internal_lookupLE'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($dictOrd_0, &$go__2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_2 = ((($dictOrd_0)['compare'])($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "LT"))) {
$__t3 = ($go__2_0)(($v_3)->{'value4'});
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "GT"))) {
$v2_5_4 = ($go__2_0)(($v_3)->{'value5'});
$__t5 = null;;
if ((is_object($v2_5_4) && (($v2_5_4)->{'tag'} === "Nothing"))) {
$__t5 = new Phpurs_Data1("Just", ["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
goto end_branch_5;;
};
$__t5 = $v2_5_4;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "EQ"))) {
$__t3 = new Phpurs_Data1("Just", ["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_lookupGE
$GLOBALS['Data_Map_Internal_lookupGE'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($dictOrd_0, &$go__2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_2 = ((($dictOrd_0)['compare'])($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "LT"))) {
$v2_5_4 = ($go__2_0)(($v_3)->{'value4'});
$__t5 = null;;
if ((is_object($v2_5_4) && (($v2_5_4)->{'tag'} === "Nothing"))) {
$__t5 = new Phpurs_Data1("Just", ["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
goto end_branch_5;;
};
$__t5 = $v2_5_4;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "GT"))) {
$__t3 = ($go__2_0)(($v_3)->{'value5'});
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "EQ"))) {
$__t3 = new Phpurs_Data1("Just", ["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_lookup
$GLOBALS['Data_Map_Internal_lookup'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($dictOrd_0, &$go__2_0, $k_1) {
  $__num = \func_num_args();
  $__tco_var_go__2_0_0_v_3 = $v_3;
  tco_loop_go__2_0_0:;
  $v_3 = $__tco_var_go__2_0_0_v_3;
  $__t0 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_1 = ((($dictOrd_0)['compare'])($k_1))(($v_3)->{'value2'});
$__t2 = null;;
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "LT"))) {
$__tco_3 = ($v_3)->{'value4'};
$__tco_var_go__2_0_0_v_3 = $__tco_3;
goto tco_loop_go__2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "GT"))) {
$__tco_4 = ($v_3)->{'value5'};
$__tco_var_go__2_0_0_v_3 = $__tco_4;
goto tco_loop_go__2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ((is_object($v1_4_1) && (($v1_4_1)->{'tag'} === "EQ"))) {
$__t2 = new Phpurs_Data1("Just", ($v_3)->{'value3'});
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_iterMapU
$GLOBALS['Data_Map_Internal_iterMapU'] = (function() {
  $__fn = function($iter_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Leaf"))) {
$__t0 = $iter_0;
goto end_branch_0;;
};
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Node"))) {
$__t1 = null;;
if ((is_object(($v_1)->{'value4'}) && ((($v_1)->{'value4'})->{'tag'} === "Leaf"))) {
$__t2 = null;;
if ((is_object(($v_1)->{'value5'}) && ((($v_1)->{'value5'})->{'tag'} === "Leaf"))) {
$__t2 = new Phpurs_Data3("IterEmit", ($v_1)->{'value2'}, ($v_1)->{'value3'}, $iter_0);
goto end_branch_2;;
};
$__t2 = new Phpurs_Data3("IterEmit", ($v_1)->{'value2'}, ($v_1)->{'value3'}, new Phpurs_Data2("IterNode", ($v_1)->{'value5'}, $iter_0));
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
if ((is_object(($v_1)->{'value5'}) && ((($v_1)->{'value5'})->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data3("IterEmit", ($v_1)->{'value2'}, ($v_1)->{'value3'}, new Phpurs_Data2("IterNode", ($v_1)->{'value4'}, $iter_0));
goto end_branch_1;;
};
$__t1 = new Phpurs_Data3("IterEmit", ($v_1)->{'value2'}, ($v_1)->{'value3'}, new Phpurs_Data2("IterNode", ($v_1)->{'value4'}, new Phpurs_Data2("IterNode", ($v_1)->{'value5'}, $iter_0)));
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_stepUnorderedCps
$GLOBALS['Data_Map_Internal_stepUnorderedCps'] = ($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapU']);

// Data_Map_Internal_stepUnfoldrUnordered
$GLOBALS['Data_Map_Internal_stepUnfoldrUnordered'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapU']))((function() {
  $__fn = function($k_0 = null, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", $k_0, $v_1), $next_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_toUnfoldableUnordered
$GLOBALS['Data_Map_Internal_toUnfoldableUnordered'] = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable_0)['unfoldr'])($GLOBALS['Data_Map_Internal_stepUnfoldrUnordered'])))($GLOBALS['Data_Map_Internal_toMapIter']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_stepUnordered
$GLOBALS['Data_Map_Internal_stepUnordered'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapU']))((function() {
  $__fn = function($k_0 = null, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("IterNext", $k_0, $v_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("IterDone");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_iterMapR
$GLOBALS['Data_Map_Internal_iterMapR'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = (function() use (&$go__0_0) {
  $__fn = function($iter_1 = null, $v_2 = null) use (&$go__0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__0_0_0_iter_1 = $iter_1;
  $__tco_var_go__0_0_0_v_2 = $v_2;
  tco_loop_go__0_0_0:;
  $iter_1 = $__tco_var_go__0_0_0_iter_1;
  $v_2 = $__tco_var_go__0_0_0_v_2;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Leaf"))) {
$__t0 = $iter_1;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Node"))) {
$__t3 = null;;
if ((is_object(($v_2)->{'value5'}) && ((($v_2)->{'value5'})->{'tag'} === "Leaf"))) {
$__tco_4 = new Phpurs_Data3("IterEmit", ($v_2)->{'value2'}, ($v_2)->{'value3'}, $iter_1);
$__tco_5 = ($v_2)->{'value4'};
$__tco_var_go__0_0_0_iter_1 = $__tco_4;
$__tco_var_go__0_0_0_v_2 = $__tco_5;
goto tco_loop_go__0_0_0;;
$__t3 = null;
goto end_branch_3;;
};
$__tco_1 = new Phpurs_Data3("IterEmit", ($v_2)->{'value2'}, ($v_2)->{'value3'}, new Phpurs_Data2("IterNode", ($v_2)->{'value4'}, $iter_1));
$__tco_2 = ($v_2)->{'value5'};
$__tco_var_go__0_0_0_iter_1 = $__tco_1;
$__tco_var_go__0_0_0_v_2 = $__tco_2;
goto tco_loop_go__0_0_0;;
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
return $go__0_0;
})();

// Data_Map_Internal_stepDescCps
$GLOBALS['Data_Map_Internal_stepDescCps'] = ($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapR']);

// Data_Map_Internal_stepDesc
$GLOBALS['Data_Map_Internal_stepDesc'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapR']))((function() {
  $__fn = function($k_0 = null, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("IterNext", $k_0, $v_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("IterDone");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_iterMapL
$GLOBALS['Data_Map_Internal_iterMapL'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = (function() use (&$go__0_0) {
  $__fn = function($iter_1 = null, $v_2 = null) use (&$go__0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__0_0_0_iter_1 = $iter_1;
  $__tco_var_go__0_0_0_v_2 = $v_2;
  tco_loop_go__0_0_0:;
  $iter_1 = $__tco_var_go__0_0_0_iter_1;
  $v_2 = $__tco_var_go__0_0_0_v_2;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Leaf"))) {
$__t0 = $iter_1;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Node"))) {
$__t3 = null;;
if ((is_object(($v_2)->{'value5'}) && ((($v_2)->{'value5'})->{'tag'} === "Leaf"))) {
$__tco_4 = new Phpurs_Data3("IterEmit", ($v_2)->{'value2'}, ($v_2)->{'value3'}, $iter_1);
$__tco_5 = ($v_2)->{'value4'};
$__tco_var_go__0_0_0_iter_1 = $__tco_4;
$__tco_var_go__0_0_0_v_2 = $__tco_5;
goto tco_loop_go__0_0_0;;
$__t3 = null;
goto end_branch_3;;
};
$__tco_1 = new Phpurs_Data3("IterEmit", ($v_2)->{'value2'}, ($v_2)->{'value3'}, new Phpurs_Data2("IterNode", ($v_2)->{'value5'}, $iter_1));
$__tco_2 = ($v_2)->{'value4'};
$__tco_var_go__0_0_0_iter_1 = $__tco_1;
$__tco_var_go__0_0_0_v_2 = $__tco_2;
goto tco_loop_go__0_0_0;;
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
return $go__0_0;
})();

// Data_Map_Internal_stepAscCps
$GLOBALS['Data_Map_Internal_stepAscCps'] = ($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapL']);

// Data_Map_Internal_stepAsc
$GLOBALS['Data_Map_Internal_stepAsc'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapL']))((function() {
  $__fn = function($k_0 = null, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("IterNext", $k_0, $v_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("IterDone");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_eqMapIter
$GLOBALS['Data_Map_Internal_eqMapIter'] = (function() {
  $__fn = function($dictEq_0 = null, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($dictEq1_1, $dictEq_0, &$go__2_0) {
  $__fn = function($a_3 = null, $b_4 = null) use ($dictEq1_1, $dictEq_0, &$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_5_1 = ($GLOBALS['Data_Map_Internal_stepAsc'])($a_3);
  $__t2 = null;;
  if ((is_object($v_5_1) && (($v_5_1)->{'tag'} === "IterNext"))) {
$v2_6_3 = ($GLOBALS['Data_Map_Internal_stepAsc'])($b_4);
$__t2 = ((is_object($v2_6_3) && (($v2_6_3)->{'tag'} === "IterNext")) && (((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($dictEq_0)['eq'])(($v_5_1)->{'value0'}))(($v2_6_3)->{'value0'})))(((($dictEq1_1)['eq'])(($v_5_1)->{'value1'}))(($v2_6_3)->{'value1'})) && (($go__2_0)(($v_5_1)->{'value2'}))(($v2_6_3)->{'value2'})));
goto end_branch_2;;
};
  if ((is_object($v_5_1) && (($v_5_1)->{'tag'} === "IterDone"))) {
$__t2 = true;
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
  $__res = ["eq" => $go__2_0];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_ordMapIter
$GLOBALS['Data_Map_Internal_ordMapIter'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $eqMapIter1_1_0 = ($GLOBALS['Data_Map_Internal_eqMapIter'])((($dictOrd_0)['Eq0'])(null));
  $__res = function($dictOrd1_2 = null) use ($dictOrd_0, $eqMapIter1_1_0) {
  $__num = \func_num_args();
  $eqMapIter2_3_1 = ($eqMapIter1_1_0)((($dictOrd1_2)['Eq0'])(null));
  $go__4_2 = null;
  $go__4_2 = (function() use ($dictOrd1_2, $dictOrd_0, &$go__4_2) {
  $__fn = function($a_5 = null, $b_6 = null) use ($dictOrd1_2, $dictOrd_0, &$go__4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_2_2_a_5 = $a_5;
  $__tco_var_go__4_2_2_b_6 = $b_6;
  tco_loop_go__4_2_2:;
  $a_5 = $__tco_var_go__4_2_2_a_5;
  $b_6 = $__tco_var_go__4_2_2_b_6;
  $v_7_2 = ($GLOBALS['Data_Map_Internal_stepAsc'])($b_6);
  $v1_8_3 = ($GLOBALS['Data_Map_Internal_stepAsc'])($a_5);
  $__t4 = null;;
  if ((is_object($v1_8_3) && (($v1_8_3)->{'tag'} === "IterNext"))) {
$__t5 = null;;
if ((is_object($v_7_2) && (($v_7_2)->{'tag'} === "IterNext"))) {
$v3_9_6 = ((($dictOrd_0)['compare'])(($v1_8_3)->{'value0'}))(($v_7_2)->{'value0'});
$__t7 = null;;
if ((is_object($v3_9_6) && (($v3_9_6)->{'tag'} === "EQ"))) {
$v4_10_8 = ((($dictOrd1_2)['compare'])(($v1_8_3)->{'value1'}))(($v_7_2)->{'value1'});
$__t9 = null;;
if ((is_object($v4_10_8) && (($v4_10_8)->{'tag'} === "EQ"))) {
$__tco_10 = ($v1_8_3)->{'value2'};
$__tco_11 = ($v_7_2)->{'value2'};
$__tco_var_go__4_2_2_a_5 = $__tco_10;
$__tco_var_go__4_2_2_b_6 = $__tco_11;
goto tco_loop_go__4_2_2;;
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
if ((is_object($v_7_2) && (($v_7_2)->{'tag'} === "IterDone"))) {
$__t5 = new Phpurs_Data0("GT");
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ((is_object($v1_8_3) && (($v1_8_3)->{'tag'} === "IterDone"))) {
$__t12 = null;;
if ((is_object($v_7_2) && (($v_7_2)->{'tag'} === "IterDone"))) {
$__t12 = new Phpurs_Data0("EQ");
goto end_branch_12;;
};
$__t12 = new Phpurs_Data0("LT");
end_branch_12:;
$__t4 = $__t12;
goto end_branch_4;;
};
  if ((is_object($v_7_2) && (($v_7_2)->{'tag'} === "IterDone"))) {
$__t4 = new Phpurs_Data0("GT");
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
  $__res = ["compare" => $go__4_2, "Eq0" => function($_dollar__unused_4 = null) use ($eqMapIter2_3_1) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_stepUnfoldr
$GLOBALS['Data_Map_Internal_stepUnfoldr'] = ((($GLOBALS['Data_Map_Internal_stepWith'])($GLOBALS['Data_Map_Internal_iterMapL']))((function() {
  $__fn = function($k_0 = null, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", $k_0, $v_1), $next_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Map_Internal_toUnfoldable
$GLOBALS['Data_Map_Internal_toUnfoldable'] = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable_0)['unfoldr'])($GLOBALS['Data_Map_Internal_stepUnfoldr'])))($GLOBALS['Data_Map_Internal_toMapIter']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_toUnfoldable1
$GLOBALS['Data_Map_Internal_toUnfoldable1'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Unfoldable_unfoldableArray'])['unfoldr'])($GLOBALS['Data_Map_Internal_stepUnfoldr'])))($GLOBALS['Data_Map_Internal_toMapIter']);

// Data_Map_Internal_showMap
$GLOBALS['Data_Map_Internal_showMap'] = (function() {
  $__fn = function($dictShow_0 = null, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $show1_2_0 = ($GLOBALS['Data_Show_showArrayImpl'])(((($GLOBALS['Data_Tuple_showTuple'])($dictShow_0))($dictShow1_1))['show']);
  $__res = ["show" => function($as_3 = null) use ($show1_2_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(fromFoldable "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(($show1_2_0)(($GLOBALS['Data_Map_Internal_toUnfoldable1'])($as_3))))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_isSubmap
$GLOBALS['Data_Map_Internal_isSubmap'] = (function() {
  $__fn = function($dictOrd_0 = null, $dictEq_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($dictEq_1, $dictOrd_0, &$go__2_0) {
  $__fn = function($m1_3 = null, $m2_4 = null) use ($dictEq_1, $dictOrd_0, &$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($m1_3) && (($m1_3)->{'tag'} === "Leaf"))) {
$__t1 = true;
goto end_branch_1;;
};
  if ((is_object($m1_3) && (($m1_3)->{'tag'} === "Node"))) {
$__local_var_5_2 = ($m1_3)->{'value2'};
$go__6_3 = null;
$go__6_3 = function($v_7 = null) use ($__local_var_5_2, $dictOrd_0, &$go__6_3) {
  $__num = \func_num_args();
  $__tco_var_go__6_3_3_v_7 = $v_7;
  tco_loop_go__6_3_3:;
  $v_7 = $__tco_var_go__6_3_3_v_7;
  $__t3 = null;;
  if ((is_object($v_7) && (($v_7)->{'tag'} === "Leaf"))) {
$__t3 = new Phpurs_Data0("Nothing");
goto end_branch_3;;
};
  if ((is_object($v_7) && (($v_7)->{'tag'} === "Node"))) {
$v1_8_4 = ((($dictOrd_0)['compare'])($__local_var_5_2))(($v_7)->{'value2'});
$__t5 = null;;
if ((is_object($v1_8_4) && (($v1_8_4)->{'tag'} === "LT"))) {
$__tco_6 = ($v_7)->{'value4'};
$__tco_var_go__6_3_3_v_7 = $__tco_6;
goto tco_loop_go__6_3_3;;
$__t5 = null;
goto end_branch_5;;
};
if ((is_object($v1_8_4) && (($v1_8_4)->{'tag'} === "GT"))) {
$__tco_7 = ($v_7)->{'value5'};
$__tco_var_go__6_3_3_v_7 = $__tco_7;
goto tco_loop_go__6_3_3;;
$__t5 = null;
goto end_branch_5;;
};
if ((is_object($v1_8_4) && (($v1_8_4)->{'tag'} === "EQ"))) {
$__t5 = new Phpurs_Data1("Just", ($v_7)->{'value3'});
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
$v1_7_4 = ($go__6_3)($m2_4);
$__t5 = null;;
if ((is_object($v1_7_4) && (($v1_7_4)->{'tag'} === "Nothing"))) {
$__t5 = false;
goto end_branch_5;;
};
if ((is_object($v1_7_4) && (($v1_7_4)->{'tag'} === "Just"))) {
$__t5 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($dictEq_1)['eq'])(($m1_3)->{'value3'}))(($v1_7_4)->{'value0'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($go__2_0)(($m1_3)->{'value4'}))($m2_4)))((($go__2_0)(($m1_3)->{'value5'}))($m2_4)));
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_isEmpty
$GLOBALS['Data_Map_Internal_isEmpty'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (is_object($v_0) && (($v_0)->{'tag'} === "Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_intersectionWith
$GLOBALS['Data_Map_Internal_intersectionWith'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($app_2 = null, $m1_3 = null, $m2_4 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $app_2, $m1_3, $m2_4);
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

// Data_Map_Internal_intersection
$GLOBALS['Data_Map_Internal_intersection'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
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

// Data_Map_Internal_insertWith
$GLOBALS['Data_Map_Internal_insertWith'] = (function() {
  $__fn = function($dictOrd_0 = null, $app_1 = null, $k_2 = null, $v_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $go__4_0 = null;
  $go__4_0 = function($v1_5 = null) use ($app_1, $dictOrd_0, &$go__4_0, $k_2, $v_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data6("Node", 1, 1, $k_2, $v_3, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
goto end_branch_1;;
};
  if ((is_object($v1_5) && (($v1_5)->{'tag'} === "Node"))) {
$v2_6_2 = ((($dictOrd_0)['compare'])($k_2))(($v1_5)->{'value2'});
$__t3 = null;;
if ((is_object($v2_6_2) && (($v2_6_2)->{'tag'} === "LT"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_5)->{'value2'}, ($v1_5)->{'value3'}, ($go__4_0)(($v1_5)->{'value4'}), ($v1_5)->{'value5'});
goto end_branch_3;;
};
if ((is_object($v2_6_2) && (($v2_6_2)->{'tag'} === "GT"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_5)->{'value2'}, ($v1_5)->{'value3'}, ($v1_5)->{'value4'}, ($go__4_0)(($v1_5)->{'value5'}));
goto end_branch_3;;
};
if ((is_object($v2_6_2) && (($v2_6_2)->{'tag'} === "EQ"))) {
$__t3 = new Phpurs_Data6("Node", ($v1_5)->{'value0'}, ($v1_5)->{'value1'}, $k_2, (($app_1)(($v1_5)->{'value3'}))($v_3), ($v1_5)->{'value4'}, ($v1_5)->{'value5'});
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
  $__res = $go__4_0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_insert
$GLOBALS['Data_Map_Internal_insert'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = function($v1_4 = null) use ($dictOrd_0, &$go__3_0, $k_1, $v_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data6("Node", 1, 1, $k_1, $v_2, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
goto end_branch_1;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Node"))) {
$v2_5_2 = ((($dictOrd_0)['compare'])($k_1))(($v1_4)->{'value2'});
$__t3 = null;;
if ((is_object($v2_5_2) && (($v2_5_2)->{'tag'} === "LT"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_4)->{'value2'}, ($v1_4)->{'value3'}, ($go__3_0)(($v1_4)->{'value4'}), ($v1_4)->{'value5'});
goto end_branch_3;;
};
if ((is_object($v2_5_2) && (($v2_5_2)->{'tag'} === "GT"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v1_4)->{'value2'}, ($v1_4)->{'value3'}, ($v1_4)->{'value4'}, ($go__3_0)(($v1_4)->{'value5'}));
goto end_branch_3;;
};
if ((is_object($v2_5_2) && (($v2_5_2)->{'tag'} === "EQ"))) {
$__t3 = new Phpurs_Data6("Node", ($v1_4)->{'value0'}, ($v1_4)->{'value1'}, $k_1, $v_2, ($v1_4)->{'value4'}, ($v1_4)->{'value5'});
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
  $__res = $go__3_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_functorMap
$GLOBALS['Data_Map_Internal_functorMap'] = ["map" => function($f_0 = null) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = function($v_2 = null) use ($f_0, &$go__1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
goto end_branch_1;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Node"))) {
$__t1 = new Phpurs_Data6("Node", ($v_2)->{'value0'}, ($v_2)->{'value1'}, ($v_2)->{'value2'}, ($f_0)(($v_2)->{'value3'}), ($go__1_0)(($v_2)->{'value4'}), ($go__1_0)(($v_2)->{'value5'}));
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
  $__res = $go__1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_functorWithIndexMap
$GLOBALS['Data_Map_Internal_functorWithIndexMap'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = function($v_2 = null) use ($f_0, &$go__1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
goto end_branch_1;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Node"))) {
$__t1 = new Phpurs_Data6("Node", ($v_2)->{'value0'}, ($v_2)->{'value1'}, ($v_2)->{'value2'}, (($f_0)(($v_2)->{'value2'}))(($v_2)->{'value3'}), ($go__1_0)(($v_2)->{'value4'}), ($go__1_0)(($v_2)->{'value5'}));
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
  $__res = $go__1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_foldableMap
$GLOBALS['Data_Map_Internal_foldableMap'] = ["foldr" => (function() {
  $__fn = function($f_0 = null, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($f_0, &$go__2_0) {
  $__fn = function($m_prime_3 = null, $z_prime_4 = null) use ($f_0, &$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($m_prime_3) && (($m_prime_3)->{'tag'} === "Leaf"))) {
$__t1 = $z_prime_4;
goto end_branch_1;;
};
  if ((is_object($m_prime_3) && (($m_prime_3)->{'tag'} === "Node"))) {
$__t1 = ($go__2_0)(($m_prime_3)->{'value4'}, (($f_0)(($m_prime_3)->{'value3'}))(($go__2_0)(($m_prime_3)->{'value5'}, $z_prime_4)));
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
  $__res = function($m_3 = null) use (&$go__2_0, $z_1) {
  $__num = \func_num_args();
  $__res = ($go__2_0)($m_3, $z_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0 = null, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_2 = null;
  $go__2_2 = (function() use ($f_0, &$go__2_2) {
  $__fn = function($z_prime_3 = null, $m_prime_4 = null) use ($f_0, &$go__2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if ((is_object($m_prime_4) && (($m_prime_4)->{'tag'} === "Leaf"))) {
$__t3 = $z_prime_3;
goto end_branch_3;;
};
  if ((is_object($m_prime_4) && (($m_prime_4)->{'tag'} === "Node"))) {
$__t3 = ($go__2_2)((($f_0)(($go__2_2)($z_prime_3, ($m_prime_4)->{'value4'})))(($m_prime_4)->{'value3'}), ($m_prime_4)->{'value5'});
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
  $__res = function($m_3 = null) use (&$go__2_2, $z_1) {
  $__num = \func_num_args();
  $__res = ($go__2_2)($z_1, $m_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_4 = ($dictMonoid_0)['mempty'];
  $__local_var_2_5 = (($dictMonoid_0)['Semigroup0'])(null);
  $__res = function($f_3 = null) use ($__local_var_2_5, $mempty_1_4) {
  $__num = \func_num_args();
  $go__4_6 = null;
  $go__4_6 = function($v_5 = null) use ($__local_var_2_5, $f_3, &$go__4_6, $mempty_1_4) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Leaf"))) {
$__t7 = $mempty_1_4;
goto end_branch_7;;
};
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Node"))) {
$__t7 = ((($__local_var_2_5)['append'])(($go__4_6)(($v_5)->{'value4'})))(((($__local_var_2_5)['append'])(($f_3)(($v_5)->{'value3'})))(($go__4_6)(($v_5)->{'value5'})));
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
  $__res = $go__4_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_foldableWithIndexMap
$GLOBALS['Data_Map_Internal_foldableWithIndexMap'] = ["foldrWithIndex" => (function() {
  $__fn = function($f_0 = null, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($f_0, &$go__2_0) {
  $__fn = function($m_prime_3 = null, $z_prime_4 = null) use ($f_0, &$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($m_prime_3) && (($m_prime_3)->{'tag'} === "Leaf"))) {
$__t1 = $z_prime_4;
goto end_branch_1;;
};
  if ((is_object($m_prime_3) && (($m_prime_3)->{'tag'} === "Node"))) {
$__t1 = ($go__2_0)(($m_prime_3)->{'value4'}, ((($f_0)(($m_prime_3)->{'value2'}))(($m_prime_3)->{'value3'}))(($go__2_0)(($m_prime_3)->{'value5'}, $z_prime_4)));
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
  $__res = function($m_3 = null) use (&$go__2_0, $z_1) {
  $__num = \func_num_args();
  $__res = ($go__2_0)($m_3, $z_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0 = null, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_2 = null;
  $go__2_2 = (function() use ($f_0, &$go__2_2) {
  $__fn = function($z_prime_3 = null, $m_prime_4 = null) use ($f_0, &$go__2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if ((is_object($m_prime_4) && (($m_prime_4)->{'tag'} === "Leaf"))) {
$__t3 = $z_prime_3;
goto end_branch_3;;
};
  if ((is_object($m_prime_4) && (($m_prime_4)->{'tag'} === "Node"))) {
$__t3 = ($go__2_2)(((($f_0)(($m_prime_4)->{'value2'}))(($go__2_2)($z_prime_3, ($m_prime_4)->{'value4'})))(($m_prime_4)->{'value3'}), ($m_prime_4)->{'value5'});
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
  $__res = function($m_3 = null) use (&$go__2_2, $z_1) {
  $__num = \func_num_args();
  $__res = ($go__2_2)($z_1, $m_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_4 = ($dictMonoid_0)['mempty'];
  $__local_var_2_5 = (($dictMonoid_0)['Semigroup0'])(null);
  $__res = function($f_3 = null) use ($__local_var_2_5, $mempty_1_4) {
  $__num = \func_num_args();
  $go__4_6 = null;
  $go__4_6 = function($v_5 = null) use ($__local_var_2_5, $f_3, &$go__4_6, $mempty_1_4) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Leaf"))) {
$__t7 = $mempty_1_4;
goto end_branch_7;;
};
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Node"))) {
$__t7 = ((($__local_var_2_5)['append'])(($go__4_6)(($v_5)->{'value4'})))(((($__local_var_2_5)['append'])((($f_3)(($v_5)->{'value2'}))(($v_5)->{'value3'})))(($go__4_6)(($v_5)->{'value5'})));
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
  $__res = $go__4_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_foldableMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_keys
$GLOBALS['Data_Map_Internal_keys'] = ((($GLOBALS['Data_Map_Internal_foldableWithIndexMap'])['foldrWithIndex'])((function() {
  $__fn = function($k_0 = null, $v_1 = null, $acc_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data2("Cons", $k_0, $acc_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nil"));

// Data_Map_Internal_traversableMap
$GLOBALS['Data_Map_Internal_traversableMap'] = ["traverse" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)['Apply0'])(null);
  $__res = function($f_2 = null) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $go__3_1 = null;
  $go__3_1 = function($v_4 = null) use ($Apply0_1_0, $dictApplicative_0, $f_2, &$go__3_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Leaf"))) {
$__t2 = (($dictApplicative_0)['pure'])(new Phpurs_Data0("Leaf"));
goto end_branch_2;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Node"))) {
$__local_var_5_3 = ($v_4)->{'value0'};
$__local_var_6_4 = ($v_4)->{'value2'};
$__local_var_7_5 = ($v_4)->{'value1'};
$__t2 = ((($Apply0_1_0)['apply'])(((($Apply0_1_0)['apply'])(((((($Apply0_1_0)['Functor0'])(null))['map'])((function() use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__fn = function($l_prime_8 = null, $v_prime_9 = null, $r_prime_10 = null) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data6("Node", $__local_var_5_3, $__local_var_7_5, $__local_var_6_4, $v_prime_9, $l_prime_8, $r_prime_10);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(($go__3_1)(($v_4)->{'value4'}))))(($f_2)(($v_4)->{'value3'}))))(($go__3_1)(($v_4)->{'value5'}));
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
  $__res = $go__3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Map_Internal_traversableMap'])['traverse'])($dictApplicative_0))($GLOBALS['Data_Map_Internal_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_foldableMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_traversableWithIndexMap
$GLOBALS['Data_Map_Internal_traversableWithIndexMap'] = ["traverseWithIndex" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)['Apply0'])(null);
  $__res = function($f_2 = null) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $go__3_1 = null;
  $go__3_1 = function($v_4 = null) use ($Apply0_1_0, $dictApplicative_0, $f_2, &$go__3_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Leaf"))) {
$__t2 = (($dictApplicative_0)['pure'])(new Phpurs_Data0("Leaf"));
goto end_branch_2;;
};
  if ((is_object($v_4) && (($v_4)->{'tag'} === "Node"))) {
$__local_var_5_3 = ($v_4)->{'value0'};
$__local_var_6_4 = ($v_4)->{'value2'};
$__local_var_7_5 = ($v_4)->{'value1'};
$__t2 = ((($Apply0_1_0)['apply'])(((($Apply0_1_0)['apply'])(((((($Apply0_1_0)['Functor0'])(null))['map'])((function() use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__fn = function($l_prime_8 = null, $v_prime_9 = null, $r_prime_10 = null) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data6("Node", $__local_var_5_3, $__local_var_7_5, $__local_var_6_4, $v_prime_9, $l_prime_8, $r_prime_10);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(($go__3_1)(($v_4)->{'value4'}))))((($f_2)($__local_var_6_4))(($v_4)->{'value3'}))))(($go__3_1)(($v_4)->{'value5'}));
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
  $__res = $go__3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorWithIndexMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_foldableWithIndexMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_traversableMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Map_Internal_values
$GLOBALS['Data_Map_Internal_values'] = ((($GLOBALS['Data_Map_Internal_foldableMap'])['foldr'])($GLOBALS['Data_List_Types_Cons']))(new Phpurs_Data0("Nil"));

// Data_Map_Internal_foldSubmapBy
$GLOBALS['Data_Map_Internal_foldSubmapBy'] = (function() {
  $__fn = function($dictOrd_0 = null, $appendFn_1 = null, $memptyValue_2 = null, $kmin_3 = null, $kmax_4 = null, $f_5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__t0 = null;;
  if ((is_object($kmin_3) && (($kmin_3)->{'tag'} === "Just"))) {
$__local_var_6_1 = ($kmin_3)->{'value0'};
$__t0 = function($k_7 = null) use ($__local_var_6_1, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = (is_object(((($dictOrd_0)['compare'])($k_7))($__local_var_6_1)) && ((((($dictOrd_0)['compare'])($k_7))($__local_var_6_1))->{'tag'} === "LT"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_0;;
};
  if ((is_object($kmin_3) && (($kmin_3)->{'tag'} === "Nothing"))) {
$__t0 = function($v_6 = null) {
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
  $__t3 = null;;
  if ((is_object($kmax_4) && (($kmax_4)->{'tag'} === "Just"))) {
$__local_var_7_4 = ($kmax_4)->{'value0'};
$__t3 = function($k_8 = null) use ($__local_var_7_4, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = (is_object(((($dictOrd_0)['compare'])($k_8))($__local_var_7_4)) && ((((($dictOrd_0)['compare'])($k_8))($__local_var_7_4))->{'tag'} === "GT"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_3;;
};
  if ((is_object($kmax_4) && (($kmax_4)->{'tag'} === "Nothing"))) {
$__t3 = function($v_7 = null) {
  $__num = \func_num_args();
  $__res = false;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $tooLarge_7_3 = $__t3;
  $__t6 = null;;
  if ((is_object($kmin_3) && (($kmin_3)->{'tag'} === "Just"))) {
$__t7 = null;;
if ((is_object($kmax_4) && (($kmax_4)->{'tag'} === "Just"))) {
$__local_var_8_8 = ($kmax_4)->{'value0'};
$__local_var_9_9 = ($kmin_3)->{'value0'};
$__t7 = function($k_10 = null) use ($__local_var_8_8, $__local_var_9_9, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(( ! (is_object(((($dictOrd_0)['compare'])($__local_var_9_9))($k_10)) && ((((($dictOrd_0)['compare'])($__local_var_9_9))($k_10))->{'tag'} === "GT")))))(( ! (is_object(((($dictOrd_0)['compare'])($k_10))($__local_var_8_8)) && ((((($dictOrd_0)['compare'])($k_10))($__local_var_8_8))->{'tag'} === "GT"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_7;;
};
if ((is_object($kmax_4) && (($kmax_4)->{'tag'} === "Nothing"))) {
$__local_var_8_10 = ($kmin_3)->{'value0'};
$__t7 = function($k_9 = null) use ($__local_var_8_10, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = ( ! (is_object(((($dictOrd_0)['compare'])($__local_var_8_10))($k_9)) && ((((($dictOrd_0)['compare'])($__local_var_8_10))($k_9))->{'tag'} === "GT")));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_7;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
end_branch_7:;
$__t6 = $__t7;
goto end_branch_6;;
};
  if ((is_object($kmin_3) && (($kmin_3)->{'tag'} === "Nothing"))) {
$__t11 = null;;
if ((is_object($kmax_4) && (($kmax_4)->{'tag'} === "Just"))) {
$__local_var_8_12 = ($kmax_4)->{'value0'};
$__t11 = function($k_9 = null) use ($__local_var_8_12, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = ( ! (is_object(((($dictOrd_0)['compare'])($k_9))($__local_var_8_12)) && ((((($dictOrd_0)['compare'])($k_9))($__local_var_8_12))->{'tag'} === "GT")));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_11;;
};
if ((is_object($kmax_4) && (($kmax_4)->{'tag'} === "Nothing"))) {
$__t11 = function($v_8 = null) {
  $__num = \func_num_args();
  $__res = true;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_11;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
end_branch_11:;
$__t6 = $__t11;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $inBounds_8_6 = $__t6;
  $go__9_14 = null;
  $go__9_14 = function($v_10 = null) use ($appendFn_1, $f_5, &$go__9_14, $inBounds_8_6, $memptyValue_2, $tooLarge_7_3, $tooSmall_6_0) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ((is_object($v_10) && (($v_10)->{'tag'} === "Leaf"))) {
$__t15 = $memptyValue_2;
goto end_branch_15;;
};
  if ((is_object($v_10) && (($v_10)->{'tag'} === "Node"))) {
$__t16 = null;;
if (($tooSmall_6_0)(($v_10)->{'value2'})) {
$__t16 = $memptyValue_2;
goto end_branch_16;;
};
$__t16 = ($go__9_14)(($v_10)->{'value4'});
end_branch_16:;
$__t17 = null;;
if (($inBounds_8_6)(($v_10)->{'value2'})) {
$__t17 = (($f_5)(($v_10)->{'value2'}))(($v_10)->{'value3'});
goto end_branch_17;;
};
$__t17 = $memptyValue_2;
end_branch_17:;
$__t18 = null;;
if (($tooLarge_7_3)(($v_10)->{'value2'})) {
$__t18 = $memptyValue_2;
goto end_branch_18;;
};
$__t18 = ($go__9_14)(($v_10)->{'value5'});
end_branch_18:;
$__t15 = (($appendFn_1)((($appendFn_1)($__t16))($__t17)))($__t18);
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__9_14;
  goto __end;;
  __end:
  return $__num > 6 ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_foldSubmap
$GLOBALS['Data_Map_Internal_foldSubmap'] = (function() {
  $__fn = function($dictOrd_0 = null, $dictMonoid_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Map_Internal_foldSubmapBy'])($dictOrd_0))(((($dictMonoid_1)['Semigroup0'])(null))['append']))(($dictMonoid_1)['mempty']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_findMin
$GLOBALS['Data_Map_Internal_findMin'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_Map_Internal_findMin_v_0 = $v_0;
  tco_loop_Data_Map_Internal_findMin:;
  $v_0 = $__tco_var_Data_Map_Internal_findMin_v_0;
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Leaf"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Node"))) {
$__t2 = null;;
if ((is_object(($v_0)->{'value4'}) && ((($v_0)->{'value4'})->{'tag'} === "Leaf"))) {
$__t2 = new Phpurs_Data1("Just", ["key" => ($v_0)->{'value2'}, "value" => ($v_0)->{'value3'}]);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_lookupGT
$GLOBALS['Data_Map_Internal_lookupGT'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($dictOrd_0, &$go__2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_2 = ((($dictOrd_0)['compare'])($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "LT"))) {
$v2_5_4 = ($go__2_0)(($v_3)->{'value4'});
$__t5 = null;;
if ((is_object($v2_5_4) && (($v2_5_4)->{'tag'} === "Nothing"))) {
$__t5 = new Phpurs_Data1("Just", ["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
goto end_branch_5;;
};
$__t5 = $v2_5_4;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "GT"))) {
$__t3 = ($go__2_0)(($v_3)->{'value5'});
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "EQ"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_findMin'])(($v_3)->{'value5'});
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_findMax
$GLOBALS['Data_Map_Internal_findMax'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_Map_Internal_findMax_v_0 = $v_0;
  tco_loop_Data_Map_Internal_findMax:;
  $v_0 = $__tco_var_Data_Map_Internal_findMax_v_0;
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Leaf"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Node"))) {
$__t2 = null;;
if ((is_object(($v_0)->{'value5'}) && ((($v_0)->{'value5'})->{'tag'} === "Leaf"))) {
$__t2 = new Phpurs_Data1("Just", ["key" => ($v_0)->{'value2'}, "value" => ($v_0)->{'value3'}]);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_lookupLT
$GLOBALS['Data_Map_Internal_lookupLT'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($dictOrd_0, &$go__2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_2 = ((($dictOrd_0)['compare'])($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "LT"))) {
$__t3 = ($go__2_0)(($v_3)->{'value4'});
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "GT"))) {
$v2_5_4 = ($go__2_0)(($v_3)->{'value5'});
$__t5 = null;;
if ((is_object($v2_5_4) && (($v2_5_4)->{'tag'} === "Nothing"))) {
$__t5 = new Phpurs_Data1("Just", ["key" => ($v_3)->{'value2'}, "value" => ($v_3)->{'value3'}]);
goto end_branch_5;;
};
$__t5 = $v2_5_4;
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "EQ"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_findMax'])(($v_3)->{'value4'});
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_filterWithKey
$GLOBALS['Data_Map_Internal_filterWithKey'] = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($f_1, &$go__2_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
goto end_branch_1;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$__t2 = null;;
if ((($f_1)(($v_3)->{'value2'}))(($v_3)->{'value3'})) {
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($go__2_0)(($v_3)->{'value4'}), ($go__2_0)(($v_3)->{'value5'}));
goto end_branch_2;;
};
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($go__2_0)(($v_3)->{'value4'}), ($go__2_0)(($v_3)->{'value5'}));
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_filterKeys
$GLOBALS['Data_Map_Internal_filterKeys'] = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($f_1, &$go__2_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
goto end_branch_1;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$__t2 = null;;
if (($f_1)(($v_3)->{'value2'})) {
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($go__2_0)(($v_3)->{'value4'}), ($go__2_0)(($v_3)->{'value5'}));
goto end_branch_2;;
};
$__t2 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($go__2_0)(($v_3)->{'value4'}), ($go__2_0)(($v_3)->{'value5'}));
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_filter
$GLOBALS['Data_Map_Internal_filter'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Map_Internal_filterWithKey'])($dictOrd_0)))($GLOBALS['Data_Function_const']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_eqMap
$GLOBALS['Data_Map_Internal_eqMap'] = (function() {
  $__fn = function($dictEq_0 = null, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($xs_2 = null, $ys_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($xs_2) && (($xs_2)->{'tag'} === "Leaf"))) {
$__t0 = (is_object($ys_3) && (($ys_3)->{'tag'} === "Leaf"));
goto end_branch_0;;
};
  if ((is_object($xs_2) && (($xs_2)->{'tag'} === "Node"))) {
$__t0 = ((is_object($ys_3) && (($ys_3)->{'tag'} === "Node")) && ((($xs_2)->{'value1'} === ($ys_3)->{'value1'}) && ((((($GLOBALS['Data_Map_Internal_eqMapIter'])($dictEq_0))($dictEq1_1))['eq'])(new Phpurs_Data2("IterNode", $xs_2, new Phpurs_Data0("IterLeaf"))))(new Phpurs_Data2("IterNode", $ys_3, new Phpurs_Data0("IterLeaf")))));
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
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_ordMap
$GLOBALS['Data_Map_Internal_ordMap'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $ordMapIter1_1_0 = ($GLOBALS['Data_Map_Internal_ordMapIter'])($dictOrd_0);
  $eqMap1_2_1 = ($GLOBALS['Data_Map_Internal_eqMap'])((($dictOrd_0)['Eq0'])(null));
  $__res = function($dictOrd1_3 = null) use ($eqMap1_2_1, $ordMapIter1_1_0) {
  $__num = \func_num_args();
  $eqMap2_4_2 = ($eqMap1_2_1)((($dictOrd1_3)['Eq0'])(null));
  $__res = ["compare" => (function() use ($dictOrd1_3, $ordMapIter1_1_0) {
  $__fn = function($xs_5 = null, $ys_6 = null) use ($dictOrd1_3, $ordMapIter1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t3 = null;;
  if ((is_object($xs_5) && (($xs_5)->{'tag'} === "Leaf"))) {
$__t4 = null;;
if ((is_object($ys_6) && (($ys_6)->{'tag'} === "Leaf"))) {
$__t4 = new Phpurs_Data0("EQ");
goto end_branch_4;;
};
$__t4 = new Phpurs_Data0("LT");
end_branch_4:;
$__t3 = $__t4;
goto end_branch_3;;
};
  if ((is_object($ys_6) && (($ys_6)->{'tag'} === "Leaf"))) {
$__t3 = new Phpurs_Data0("GT");
goto end_branch_3;;
};
  $__t3 = (((($ordMapIter1_1_0)($dictOrd1_3))['compare'])(new Phpurs_Data2("IterNode", $xs_5, new Phpurs_Data0("IterLeaf"))))(new Phpurs_Data2("IterNode", $ys_6, new Phpurs_Data0("IterLeaf")));
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_5 = null) use ($eqMap2_4_2) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_eq1Map
$GLOBALS['Data_Map_Internal_eq1Map'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq1" => function($dictEq1_1 = null) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Map_Internal_eqMap'])($dictEq_0))($dictEq1_1))['eq'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_ord1Map
$GLOBALS['Data_Map_Internal_ord1Map'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $ordMap1_1_0 = ($GLOBALS['Data_Map_Internal_ordMap'])($dictOrd_0);
  $__local_var_2_1 = (($dictOrd_0)['Eq0'])(null);
  $eq1Map1_3_2 = ["eq1" => function($dictEq1_3 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Map_Internal_eqMap'])($__local_var_2_1))($dictEq1_3))['eq'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["compare1" => function($dictOrd1_4 = null) use ($ordMap1_1_0) {
  $__num = \func_num_args();
  $__res = (($ordMap1_1_0)($dictOrd1_4))['compare'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_4 = null) use ($eq1Map1_3_2) {
  $__num = \func_num_args();
  $__res = $eq1Map1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_empty
$GLOBALS['Data_Map_Internal_empty'] = new Phpurs_Data0("Leaf");

// Data_Map_Internal_fromFoldable
$GLOBALS['Data_Map_Internal_fromFoldable'] = (function() {
  $__fn = function($dictOrd_0 = null, $dictFoldable_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_1)['foldl'])((function() use ($dictOrd_0) {
  $__fn = function($m_2 = null, $v_3 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0))(($v_3)->{'value0'}))(($v_3)->{'value1'}))($m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_fromFoldableWith
$GLOBALS['Data_Map_Internal_fromFoldableWith'] = (function() {
  $__fn = function($dictOrd_0 = null, $dictFoldable_1 = null, $f_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $f_prime_3_0 = (($GLOBALS['Data_Map_Internal_insertWith'])($dictOrd_0))((function() use ($f_2) {
  $__fn = function($b_3 = null, $a_4 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_2)($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  $__res = ((($dictFoldable_1)['foldl'])((function() use ($f_prime_3_0) {
  $__fn = function($m_4 = null, $v_5 = null) use ($f_prime_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($f_prime_3_0)(($v_5)->{'value0'}))(($v_5)->{'value1'}))($m_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_fromFoldableWithIndex
$GLOBALS['Data_Map_Internal_fromFoldableWithIndex'] = (function() {
  $__fn = function($dictOrd_0 = null, $dictFoldableWithIndex_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldableWithIndex_1)['foldlWithIndex'])((function() use ($dictOrd_0) {
  $__fn = function($k_2 = null, $m_3 = null, $v_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_Map_Internal_insert'])($dictOrd_0))($k_2))($v_4))($m_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_monoidSemigroupMap
$GLOBALS['Data_Map_Internal_monoidSemigroupMap'] = (function() {
  $__fn = function($_dollar__unused_0 = null, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupMap2_2_0 = (($GLOBALS['Data_Map_Internal_semigroupMap'])(null))($dictOrd_1);
  $__res = function($dictSemigroup_3 = null) use ($semigroupMap2_2_0) {
  $__num = \func_num_args();
  $semigroupMap3_4_1 = ($semigroupMap2_2_0)($dictSemigroup_3);
  $__res = ["mempty" => new Phpurs_Data0("Leaf"), "Semigroup0" => function($_dollar__unused_5 = null) use ($semigroupMap3_4_1) {
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_submap
$GLOBALS['Data_Map_Internal_submap'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0, $dictOrd_0) {
  $__fn = function($kmin_2 = null, $kmax_3 = null) use ($compare_1_0, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((($GLOBALS['Data_Map_Internal_foldSubmapBy'])($dictOrd_0))((function() use ($compare_1_0) {
  $__fn = function($m1_4 = null, $m2_5 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_4, $m2_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf")))($kmin_2))($kmax_3))($GLOBALS['Data_Map_Internal_singleton']);
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

// Data_Map_Internal_unions
$GLOBALS['Data_Map_Internal_unions'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = function($dictFoldable_2 = null) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_2)['foldl'])((function() use ($compare_1_0) {
  $__fn = function($m1_3 = null, $m2_4 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_3, $m2_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_difference
$GLOBALS['Data_Map_Internal_difference'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeDifference'])($compare_1_0, $m1_2, $m2_3);
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

// Data_Map_Internal_delete
$GLOBALS['Data_Map_Internal_delete'] = (function() {
  $__fn = function($dictOrd_0 = null, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = function($v_3 = null) use ($dictOrd_0, &$go__2_0, $k_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
goto end_branch_1;;
};
  if ((is_object($v_3) && (($v_3)->{'tag'} === "Node"))) {
$v1_4_2 = ((($dictOrd_0)['compare'])($k_1))(($v_3)->{'value2'});
$__t3 = null;;
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "LT"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($go__2_0)(($v_3)->{'value4'}), ($v_3)->{'value5'});
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "GT"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])(($v_3)->{'value2'}, ($v_3)->{'value3'}, ($v_3)->{'value4'}, ($go__2_0)(($v_3)->{'value5'}));
goto end_branch_3;;
};
if ((is_object($v1_4_2) && (($v1_4_2)->{'tag'} === "EQ"))) {
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
  $__res = $go__2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Map_Internal_checkValid
$GLOBALS['Data_Map_Internal_checkValid'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = function($v_2 = null) use ($dictOrd_0, &$go__1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Leaf"))) {
$__t1 = true;
goto end_branch_1;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Node"))) {
$__t2 = null;;
if ((is_object(($v_2)->{'value4'}) && ((($v_2)->{'value4'})->{'tag'} === "Leaf"))) {
$__t3 = null;;
if ((is_object(($v_2)->{'value5'}) && ((($v_2)->{'value5'})->{'tag'} === "Leaf"))) {
$__t3 = true;
goto end_branch_3;;
};
if ((is_object(($v_2)->{'value5'}) && ((($v_2)->{'value5'})->{'tag'} === "Node"))) {
$__t3 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($v_2)->{'value0'} === 2)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($v_2)->{'value5'})->{'value0'} === 1)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($GLOBALS['Data_Map_Internal_greaterThan'])(($v_2)->{'value1'}))((($v_2)->{'value5'})->{'value1'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((is_object(((($dictOrd_0)['compare'])((($v_2)->{'value5'})->{'value2'}))(($v_2)->{'value2'})) && ((((($dictOrd_0)['compare'])((($v_2)->{'value5'})->{'value2'}))(($v_2)->{'value2'}))->{'tag'} === "GT"))))(($go__1_0)(($v_2)->{'value5'})))));
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
if ((is_object(($v_2)->{'value4'}) && ((($v_2)->{'value4'})->{'tag'} === "Node"))) {
$__t4 = null;;
if ((is_object(($v_2)->{'value5'}) && ((($v_2)->{'value5'})->{'tag'} === "Leaf"))) {
$__t4 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($v_2)->{'value0'} === 2)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($v_2)->{'value4'})->{'value0'} === 1)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($GLOBALS['Data_Map_Internal_greaterThan'])(($v_2)->{'value1'}))((($v_2)->{'value4'})->{'value1'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((is_object(((($dictOrd_0)['compare'])((($v_2)->{'value4'})->{'value2'}))(($v_2)->{'value2'})) && ((((($dictOrd_0)['compare'])((($v_2)->{'value4'})->{'value2'}))(($v_2)->{'value2'}))->{'tag'} === "LT"))))(($go__1_0)(($v_2)->{'value4'})))));
goto end_branch_4;;
};
if ((is_object(($v_2)->{'value5'}) && ((($v_2)->{'value5'})->{'tag'} === "Node"))) {
$__t4 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($GLOBALS['Data_Map_Internal_greaterThan'])(($v_2)->{'value0'}))((($v_2)->{'value5'})->{'value0'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((is_object(((($dictOrd_0)['compare'])((($v_2)->{'value5'})->{'value2'}))(($v_2)->{'value2'})) && ((((($dictOrd_0)['compare'])((($v_2)->{'value5'})->{'value2'}))(($v_2)->{'value2'}))->{'tag'} === "GT"))))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($GLOBALS['Data_Map_Internal_greaterThan'])(($v_2)->{'value0'}))((($v_2)->{'value4'})->{'value0'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((is_object(((($dictOrd_0)['compare'])((($v_2)->{'value4'})->{'value2'}))(($v_2)->{'value2'})) && ((((($dictOrd_0)['compare'])((($v_2)->{'value4'})->{'value2'}))(($v_2)->{'value2'}))->{'tag'} === "LT"))))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])((($GLOBALS['Data_Map_Internal_lessThan'])(($GLOBALS['Data_Map_Internal_abs'])(((($v_2)->{'value5'})->{'value0'} - (($v_2)->{'value4'})->{'value0'}))))(2)))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((((($v_2)->{'value5'})->{'value1'} + (($v_2)->{'value4'})->{'value1'}) + 1) === ($v_2)->{'value1'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(($go__1_0)(($v_2)->{'value4'})))(($go__1_0)(($v_2)->{'value5'}))))))));
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
  $__res = $go__1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_catMaybes
$GLOBALS['Data_Map_Internal_catMaybes'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Map_Internal_mapMaybeWithKey'])($dictOrd_0)))($GLOBALS['Data_Function_const']))($GLOBALS['Data_Map_Internal_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_applyMap
$GLOBALS['Data_Map_Internal_applyMap'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = ["apply" => (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Map_Internal_identity'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_1 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_bindMap
$GLOBALS['Data_Map_Internal_bindMap'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $applyMap1_1_0 = ["apply" => (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'])($compare_1_0, $GLOBALS['Data_Map_Internal_identity'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_1 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["bind" => (function() use ($dictOrd_0) {
  $__fn = function($m_2 = null, $f_3 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Map_Internal_mapMaybeWithKey'])($dictOrd_0))(function($k_4 = null) use ($dictOrd_0, $f_3) {
  $__num = \func_num_args();
  $go__5_2 = null;
  $go__5_2 = function($v_6 = null) use ($dictOrd_0, &$go__5_2, $k_4) {
  $__num = \func_num_args();
  $__tco_var_go__5_2_2_v_6 = $v_6;
  tco_loop_go__5_2_2:;
  $v_6 = $__tco_var_go__5_2_2_v_6;
  $__t2 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Leaf"))) {
$__t2 = new Phpurs_Data0("Nothing");
goto end_branch_2;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Node"))) {
$v1_7_3 = ((($dictOrd_0)['compare'])($k_4))(($v_6)->{'value2'});
$__t4 = null;;
if ((is_object($v1_7_3) && (($v1_7_3)->{'tag'} === "LT"))) {
$__tco_5 = ($v_6)->{'value4'};
$__tco_var_go__5_2_2_v_6 = $__tco_5;
goto tco_loop_go__5_2_2;;
$__t4 = null;
goto end_branch_4;;
};
if ((is_object($v1_7_3) && (($v1_7_3)->{'tag'} === "GT"))) {
$__tco_6 = ($v_6)->{'value5'};
$__tco_var_go__5_2_2_v_6 = $__tco_6;
goto tco_loop_go__5_2_2;;
$__t4 = null;
goto end_branch_4;;
};
if ((is_object($v1_7_3) && (($v1_7_3)->{'tag'} === "EQ"))) {
$__t4 = new Phpurs_Data1("Just", ($v_6)->{'value3'});
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($go__5_2))($f_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_2 = null) use ($applyMap1_1_0) {
  $__num = \func_num_args();
  $__res = $applyMap1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_anyWithKey
$GLOBALS['Data_Map_Internal_anyWithKey'] = function($predicate_0 = null) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = function($v_2 = null) use (&$go__1_0, $predicate_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Leaf"))) {
$__t1 = false;
goto end_branch_1;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Node"))) {
$__t1 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['disj'])((($predicate_0)(($v_2)->{'value2'}))(($v_2)->{'value3'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['disj'])(($go__1_0)(($v_2)->{'value4'})))(($go__1_0)(($v_2)->{'value5'})));
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
  $__res = $go__1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_any
$GLOBALS['Data_Map_Internal_any'] = function($predicate_0 = null) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = function($v_2 = null) use (&$go__1_0, $predicate_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Leaf"))) {
$__t1 = false;
goto end_branch_1;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Node"))) {
$__t1 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['disj'])(($predicate_0)(($v_2)->{'value3'})))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['disj'])(($go__1_0)(($v_2)->{'value4'})))(($go__1_0)(($v_2)->{'value5'})));
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
  $__res = $go__1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_alter
$GLOBALS['Data_Map_Internal_alter'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = (function() use ($compare_1_0) {
  $__fn = function($f_2 = null, $k_3 = null, $m_4 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_5_1 = ($GLOBALS['Data_Map_Internal_unsafeSplit'])($compare_1_0, $k_3, $m_4);
  $v2_6_2 = ($f_2)(($v_5_1)->{'value0'});
  $__t3 = null;;
  if ((is_object($v2_6_2) && (($v2_6_2)->{'tag'} === "Nothing"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeJoinNodes'])(($v_5_1)->{'value1'}, ($v_5_1)->{'value2'});
goto end_branch_3;;
};
  if ((is_object($v2_6_2) && (($v2_6_2)->{'tag'} === "Just"))) {
$__t3 = ($GLOBALS['Data_Map_Internal_unsafeBalancedNode'])($k_3, ($v2_6_2)->{'value0'}, ($v_5_1)->{'value1'}, ($v_5_1)->{'value2'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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

// Data_Map_Internal_altMap
$GLOBALS['Data_Map_Internal_altMap'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = ["alt" => (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_1 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Map_Internal_plusMap
$GLOBALS['Data_Map_Internal_plusMap'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $altMap1_1_0 = ["alt" => (function() use ($compare_1_0) {
  $__fn = function($m1_2 = null, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Map_Internal_unsafeUnionWith'])($compare_1_0, $GLOBALS['Data_Function_const'], $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_1 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Map_Internal_functorMap'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["empty" => new Phpurs_Data0("Leaf"), "Alt0" => function($_dollar__unused_2 = null) use ($altMap1_1_0) {
  $__num = \func_num_args();
  $__res = $altMap1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

