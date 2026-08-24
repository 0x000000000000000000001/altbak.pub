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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_List_Lazy_NonEmpty_uncons
function majData_majList_majLazy_majNonmajEmpty_uncons($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $v1_1_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__res = (object)["head" => ($v1_1_0)->{'value0'}, "tail" => ($v1_1_0)->{'value1'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_uncons'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_uncons';

// Data_List_Lazy_NonEmpty_toList
function majData_majList_majLazy_majNonmajEmpty_tomajList($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_tomajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $v1_1_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_2_1 = ($v1_1_0)->{'value0'};
  $__local_var_3_2 = ($v1_1_0)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_4) use ($__local_var_2_1, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_2_1, $__local_var_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_toList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_tomajList';

// Data_List_Lazy_NonEmpty_toUnfoldable
function majData_majList_majLazy_majNonmajEmpty_tomajUnfoldable($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_tomajUnfoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable_0)->{'unfoldr'})(function($xs_1) {
  $__num = \func_num_args();
  $__local_var_2_0 = \Data\List\Lazy\majData_majList_majLazy_uncons($xs_1);
  $__t1 = null;;
  if ($__local_var_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple((($__local_var_2_0)->{'value0'})->{'head'}, (($__local_var_2_0)->{'value0'})->{'tail'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_List_Lazy_NonEmpty_toList']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_toUnfoldable'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_tomajUnfoldable';

// Data_List_Lazy_NonEmpty_tail
function majData_majList_majLazy_majNonmajEmpty_tail($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_tail';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (\Data\Lazy\majData_majLazy_force($v_0))->{'value1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_tail'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_tail';

// Data_List_Lazy_NonEmpty_singleton
function majData_majList_majLazy_majNonmajEmpty_singleton($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_1) use ($a_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_0, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_singleton'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_singleton';

// Data_List_Lazy_NonEmpty_repeat
function majData_majList_majLazy_majNonmajEmpty_repeat($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_repeat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_1) use ($x_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($x_0, \Data\List\Lazy\majData_majList_majLazy_repeat($x_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_repeat'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_repeat';

// Data_List_Lazy_NonEmpty_length
function majData_majList_majLazy_majNonmajEmpty_length($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (1 + \Data\List\Lazy\majData_majList_majLazy_length((\Data\Lazy\majData_majLazy_force($v_0))->{'value1'}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_length'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_length';

// Data_List_Lazy_NonEmpty_last
function majData_majList_majLazy_majNonmajEmpty_last($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_last';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $v1_1_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_2_1 = \Data\List\Lazy\majData_majList_majLazy_last(($v1_1_0)->{'value1'});
  $__t2 = null;;
  if ($__local_var_2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = ($v1_1_0)->{'value0'};
goto end_branch_2;;
};
  if ($__local_var_2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($__local_var_2_1)->{'value0'};
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_last'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_last';

// Data_List_Lazy_NonEmpty_iterate
function majData_majList_majLazy_majNonmajEmpty_iterate($f_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_iterate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_2) use ($f_0, $x_1) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($x_1, \Data\List\Lazy\majData_majList_majLazy_iterate($f_0, ($f_0)($x_1)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_iterate'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_iterate';

// Data_List_Lazy_NonEmpty_init
function majData_majList_majLazy_majNonmajEmpty_init($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_init';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $v1_1_0 = \Data\Lazy\majData_majLazy_force($v_0);
  $__local_var_2_1 = \Data\List\Lazy\majData_majList_majLazy_init(($v1_1_0)->{'value1'});
  $__t2 = null;;
  if ($__local_var_2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = $GLOBALS['Data_List_Lazy_Types_nil'];
goto end_branch_2;;
};
  if ($__local_var_2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_3_3 = ($__local_var_2_1)->{'value0'};
$__t2 = \Data\Lazy\majData_majLazy_defer(function($v_4) use ($__local_var_3_3, $v1_1_0) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v1_1_0)->{'value0'}, $__local_var_3_3);
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_init'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_init';

// Data_List_Lazy_NonEmpty_head
function majData_majList_majLazy_majNonmajEmpty_head($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_head';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (\Data\Lazy\majData_majLazy_force($v_0))->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_head'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_head';

// Data_List_Lazy_NonEmpty_fromList
function majData_majList_majLazy_majNonmajEmpty_frommajList($l_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_frommajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $v_1_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($l_0);
  $__t1 = null;;
  if ($v_1_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_2_2 = ($v_1_0)->{'value0'};
$__local_var_3_3 = ($v_1_0)->{'value1'};
$__t1 = new \Data\Maybe\Data_Maybe_Just(\Data\Lazy\majData_majLazy_defer(function($v1_4) use ($__local_var_2_2, $__local_var_3_3) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_2_2, $__local_var_3_3);
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_fromList'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_frommajList';

// Data_List_Lazy_NonEmpty_fromFoldable
function majData_majList_majLazy_majNonmajEmpty_frommajFoldable($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_frommajFoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_NonEmpty_fromList']))(((($dictFoldable_0)->{'foldr'})($GLOBALS['Data_List_Lazy_Types_cons']))($GLOBALS['Data_List_Lazy_Types_nil']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_fromFoldable'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_frommajFoldable';

// Data_List_Lazy_NonEmpty_cons
function majData_majList_majLazy_majNonmajEmpty_cons($y_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_cons';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v1_2) use ($v_1, $y_0) {
  $__num = \func_num_args();
  $v2_3_0 = \Data\Lazy\majData_majLazy_force($v_1);
  $__local_var_4_1 = ($v2_3_0)->{'value0'};
  $__local_var_5_2 = ($v2_3_0)->{'value1'};
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($y_0, \Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_4_1, $__local_var_5_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_4_1, $__local_var_5_2);
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_cons'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_cons';

// Data_List_Lazy_NonEmpty_concatMap
function majData_majList_majLazy_majNonmajEmpty_concatmajMap($b_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_concatmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = \Data\Lazy\majData_majLazy_force($a_1);
  $__local_var_3_1 = ($v1_2_0)->{'value1'};
  $v2_4_2 = \Data\Lazy\majData_majLazy_force(($b_0)(($v1_2_0)->{'value0'}));
  $__local_var_5_3 = ($v2_4_2)->{'value0'};
  $__local_var_6_4 = ($v2_4_2)->{'value1'};
  $__res = \Data\Lazy\majData_majLazy_defer(function($v3_7) use ($__local_var_3_1, $__local_var_5_3, $__local_var_6_4, $b_0) {
  $__num = \func_num_args();
  $__local_var_8_5 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_toList']))($b_0);
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
$__local_var_12_9 = ($__local_var_10_6)->{'value1'};
$__local_var_12_9 = \Data\Lazy\majData_majLazy_defer(function($v_13) use ($__local_var_12_9, $__local_var_8_5) {
  $__num = \func_num_args();
  $__local_var_14_10 = \Data\Lazy\majData_majLazy_force($__local_var_12_9);
  $__t11 = null;;
  if ($__local_var_14_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t11 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_11;;
};
  if ($__local_var_14_10 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_15_12 = ($__local_var_8_5)(($__local_var_14_10)->{'value0'});
$__local_var_16_13 = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})(($__local_var_14_10)->{'value1'}))($__local_var_8_5);
$__t11 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_17) use ($__local_var_15_12, $__local_var_16_13) {
  $__num = \func_num_args();
  $__local_var_18_14 = \Data\Lazy\majData_majLazy_force($__local_var_15_12);
  $__t15 = null;;
  if ($__local_var_18_14 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t15 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_16_13);
goto end_branch_15;;
};
  if ($__local_var_18_14 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_19_16 = ($__local_var_18_14)->{'value1'};
$__t15 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_18_14)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_20) use ($__local_var_16_13, $__local_var_19_16) {
  $__num = \func_num_args();
  $__local_var_21_17 = \Data\Lazy\majData_majLazy_force($__local_var_19_16);
  $__t18 = null;;
  if ($__local_var_21_17 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t18 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_16_13);
goto end_branch_18;;
};
  if ($__local_var_21_17 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t18 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_21_17)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_21_17)->{'value1'}))($__local_var_16_13));
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
$__t7 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(\Data\Lazy\majData_majLazy_defer(function($v_13) use ($__local_var_11_8, $__local_var_12_9) {
  $__num = \func_num_args();
  $__local_var_14_20 = \Data\Lazy\majData_majLazy_force($__local_var_11_8);
  $__t21 = null;;
  if ($__local_var_14_20 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t21 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_12_9);
goto end_branch_21;;
};
  if ($__local_var_14_20 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_15_22 = ($__local_var_14_20)->{'value1'};
$__t21 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_14_20)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_16) use ($__local_var_12_9, $__local_var_15_22) {
  $__num = \func_num_args();
  $__local_var_17_23 = \Data\Lazy\majData_majLazy_force($__local_var_15_22);
  $__t24 = null;;
  if ($__local_var_17_23 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t24 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_12_9);
goto end_branch_24;;
};
  if ($__local_var_17_23 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_18_25 = ($__local_var_17_23)->{'value1'};
$__t24 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_17_23)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_19) use ($__local_var_12_9, $__local_var_18_25) {
  $__num = \func_num_args();
  $__local_var_20_26 = \Data\Lazy\majData_majLazy_force($__local_var_18_25);
  $__t27 = null;;
  if ($__local_var_20_26 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t27 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_12_9);
goto end_branch_27;;
};
  if ($__local_var_20_26 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t27 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_20_26)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_20_26)->{'value1'}))($__local_var_12_9));
goto end_branch_27;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t27 = null;
  end_branch_27:;
  $__res = $__t27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
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
  $__local_var_10_29 = \Data\Lazy\majData_majLazy_force($__local_var_6_4);
  $__t30 = null;;
  if ($__local_var_10_29 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t30 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_8_5);
goto end_branch_30;;
};
  if ($__local_var_10_29 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_11_31 = ($__local_var_10_29)->{'value1'};
$__t30 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_10_29)->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_12) use ($__local_var_11_31, $__local_var_8_5) {
  $__num = \func_num_args();
  $__local_var_13_32 = \Data\Lazy\majData_majLazy_force($__local_var_11_31);
  $__t33 = null;;
  if ($__local_var_13_32 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t33 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_8_5);
goto end_branch_33;;
};
  if ($__local_var_13_32 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t33 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_13_32)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_13_32)->{'value1'}))($__local_var_8_5));
goto end_branch_33;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t33 = null;
  end_branch_33:;
  $__res = $__t33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto end_branch_30;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t30 = null;
  end_branch_30:;
  $__res = $__t30;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_concatMap'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_concatmajMap';

// Data_List_Lazy_NonEmpty_appendFoldable
function majData_majList_majLazy_majNonmajEmpty_appendmajFoldable($dictFoldable_0, $nel_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majNonmajEmpty_appendmajFoldable';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_3) use ($dictFoldable_0, $nel_1, $ys_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = (\Data\Lazy\majData_majLazy_force($nel_1))->{'value1'};
  $__local_var_5_1 = (((($dictFoldable_0)->{'foldr'})($GLOBALS['Data_List_Lazy_Types_cons']))($GLOBALS['Data_List_Lazy_Types_nil']))($ys_2);
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty((\Data\Lazy\majData_majLazy_force($nel_1))->{'value0'}, \Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_4_0, $__local_var_5_1) {
  $__num = \func_num_args();
  $__local_var_7_2 = \Data\Lazy\majData_majLazy_force($__local_var_4_0);
  $__t3 = null;;
  if ($__local_var_7_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_5_1);
goto end_branch_3;;
};
  if ($__local_var_7_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t3 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_7_2)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_7_2)->{'value1'}))($__local_var_5_1));
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
});
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_NonEmpty_appendFoldable'] = __NAMESPACE__ . '\\majData_majList_majLazy_majNonmajEmpty_appendmajFoldable';

