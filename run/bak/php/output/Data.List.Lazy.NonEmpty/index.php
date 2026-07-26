<?php

namespace Data\List\Lazy\NonEmpty;

// ALL IMPORTS: Control.Applicative, Control.Bind, Control.Semigroupoid, Data.Foldable, Data.Function, Data.Functor, Data.Lazy, Data.List.Lazy, Data.List.Lazy.NonEmpty, Data.List.Lazy.Types, Data.Maybe, Data.NonEmpty, Data.Semigroup, Data.Semiring, Data.Tuple, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Bind, Control.Semigroupoid, Data.Foldable, Data.Function, Data.Functor, Data.Lazy, Data.List.Lazy, Data.List.Lazy.NonEmpty, Data.List.Lazy.Types, Data.Maybe, Data.NonEmpty, Data.Semigroup, Data.Semiring, Data.Tuple, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Lazy.NonEmpty/index.php';
require_once __DIR__ . '/../Data.List.Lazy.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
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


// Data_List_Lazy_NonEmpty_uncons
$GLOBALS['Data_List_Lazy_NonEmpty_uncons'] = function($v_0 = null) {
  $__num = \func_num_args();
  $v1_1_0 = ($GLOBALS['Data_Lazy_force'])($v_0);
  $__res = ["head" => ($v1_1_0)->{'value0'}, "tail" => ($v1_1_0)->{'value1'}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_NonEmpty_toList
$GLOBALS['Data_List_Lazy_NonEmpty_toList'] = function($v_0 = null) {
  $__num = \func_num_args();
  $v1_1_0 = ($GLOBALS['Data_Lazy_force'])($v_0);
  $__local_var_2_1 = ($v1_1_0)->{'value0'};
  $__local_var_3_2 = ($v1_1_0)->{'value1'};
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_4 = null) use ($__local_var_2_1, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_2_1, $__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_NonEmpty_toUnfoldable
$GLOBALS['Data_List_Lazy_NonEmpty_toUnfoldable'] = function($dictUnfoldable_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable_0)['unfoldr'])(function($xs_1 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])['map'])(function($rec_2 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($rec_2)['head'], ($rec_2)['tail']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_List_Lazy_uncons'])($xs_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_List_Lazy_NonEmpty_toList']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_NonEmpty_tail
$GLOBALS['Data_List_Lazy_NonEmpty_tail'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_force'])($v_0))->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_NonEmpty_singleton
$GLOBALS['Data_List_Lazy_NonEmpty_singleton'] = ($GLOBALS['Data_List_Lazy_Types_applicativeNonEmptyList'])['pure'];

// Data_List_Lazy_NonEmpty_repeat
$GLOBALS['Data_List_Lazy_NonEmpty_repeat'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_1 = null) use ($x_0) {
  $__num = \func_num_args();
  $go__2_0 = null;
  $go__2_0 = (($GLOBALS['Data_List_Lazy_Types_lazyList'])['defer'])(function($v_3 = null) use (&$go__2_0, $x_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_4 = null) use (&$go__2_0, $x_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $x_0, $go__2_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = new Phpurs_Data2("NonEmpty", $x_0, $go__2_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_NonEmpty_length
$GLOBALS['Data_List_Lazy_NonEmpty_length'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (1 + ($GLOBALS['Data_List_Lazy_length'])((($GLOBALS['Data_Lazy_force'])($v_0))->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_NonEmpty_last
$GLOBALS['Data_List_Lazy_NonEmpty_last'] = function($v_0 = null) {
  $__num = \func_num_args();
  $v1_1_0 = ($GLOBALS['Data_Lazy_force'])($v_0);
  $__local_var_2_1 = ($GLOBALS['Data_List_Lazy_last'])(($v1_1_0)->{'value1'});
  $__t2 = null;;
  if ((is_object($__local_var_2_1) && (($__local_var_2_1)->{'tag'} === "Nothing"))) {
$__t2 = ($v1_1_0)->{'value0'};
goto end_branch_2;;
};
  if ((is_object($__local_var_2_1) && (($__local_var_2_1)->{'tag'} === "Just"))) {
$__t2 = ($__local_var_2_1)->{'value0'};
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

// Data_List_Lazy_NonEmpty_iterate
$GLOBALS['Data_List_Lazy_NonEmpty_iterate'] = (function() {
  $__fn = function($f_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_2 = null) use ($f_0, $x_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $x_1, (($GLOBALS['Data_List_Lazy_iterate'])($f_0))(($f_0)($x_1)));
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

// Data_List_Lazy_NonEmpty_init
$GLOBALS['Data_List_Lazy_NonEmpty_init'] = function($v_0 = null) {
  $__num = \func_num_args();
  $v1_1_0 = ($GLOBALS['Data_Lazy_force'])($v_0);
  $__local_var_2_1 = ($GLOBALS['Data_List_Lazy_init'])(($v1_1_0)->{'value1'});
  $__t2 = null;;
  if ((is_object($__local_var_2_1) && (($__local_var_2_1)->{'tag'} === "Nothing"))) {
$__t2 = $GLOBALS['Data_List_Lazy_Types_nil'];
goto end_branch_2;;
};
  if ((is_object($__local_var_2_1) && (($__local_var_2_1)->{'tag'} === "Just"))) {
$__local_var_3_3 = ($__local_var_2_1)->{'value0'};
$__t2 = ($GLOBALS['Data_Lazy_defer'])(function($v_4 = null) use ($__local_var_3_3, $v1_1_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", ($v1_1_0)->{'value0'}, $__local_var_3_3);
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
};

// Data_List_Lazy_NonEmpty_head
$GLOBALS['Data_List_Lazy_NonEmpty_head'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Lazy_force'])($v_0))->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_NonEmpty_fromList
$GLOBALS['Data_List_Lazy_NonEmpty_fromList'] = function($l_0 = null) {
  $__num = \func_num_args();
  $v_1_0 = ($GLOBALS['Data_List_Lazy_Types_step'])($l_0);
  $__t1 = null;;
  if ((is_object($v_1_0) && (($v_1_0)->{'tag'} === "Nil"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  if ((is_object($v_1_0) && (($v_1_0)->{'tag'} === "Cons"))) {
$__local_var_2_2 = ($v_1_0)->{'value0'};
$__local_var_3_3 = ($v_1_0)->{'value1'};
$__t1 = new Phpurs_Data1("Just", ($GLOBALS['Data_Lazy_defer'])(function($v1_4 = null) use ($__local_var_2_2, $__local_var_3_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", $__local_var_2_2, $__local_var_3_3);
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
};

// Data_List_Lazy_NonEmpty_fromFoldable
$GLOBALS['Data_List_Lazy_NonEmpty_fromFoldable'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_NonEmpty_fromList']))(((($dictFoldable_0)['foldr'])($GLOBALS['Data_List_Lazy_Types_cons']))($GLOBALS['Data_List_Lazy_Types_nil']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_List_Lazy_NonEmpty_cons
$GLOBALS['Data_List_Lazy_NonEmpty_cons'] = (function() {
  $__fn = function($y_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v1_2 = null) use ($v_1, $y_0) {
  $__num = \func_num_args();
  $v2_3_0 = ($GLOBALS['Data_Lazy_force'])($v_1);
  $__local_var_4_1 = ($v2_3_0)->{'value0'};
  $__local_var_5_2 = ($v2_3_0)->{'value1'};
  $__res = new Phpurs_Data2("NonEmpty", $y_0, ($GLOBALS['Data_Lazy_defer'])(function($v_6 = null) use ($__local_var_4_1, $__local_var_5_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Cons", $__local_var_4_1, $__local_var_5_2);
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_NonEmpty_concatMap
$GLOBALS['Data_List_Lazy_NonEmpty_concatMap'] = (function() {
  $__fn = function($b_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_Types_bindNonEmptyList'])['bind'])($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Lazy_NonEmpty_appendFoldable
$GLOBALS['Data_List_Lazy_NonEmpty_appendFoldable'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $fromFoldable1_1_0 = ((($dictFoldable_0)['foldr'])($GLOBALS['Data_List_Lazy_Types_cons']))($GLOBALS['Data_List_Lazy_Types_nil']);
  $__res = (function() use ($fromFoldable1_1_0) {
  $__fn = function($nel_2 = null, $ys_3 = null) use ($fromFoldable1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Lazy_defer'])(function($v_4 = null) use ($fromFoldable1_1_0, $nel_2, $ys_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("NonEmpty", (($GLOBALS['Data_Lazy_force'])($nel_2))->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])['append'])((($GLOBALS['Data_Lazy_force'])($nel_2))->{'value1'}))(($fromFoldable1_1_0)($ys_3)));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

