<?php

namespace Data\List\ZipList;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.List.Lazy, Data.List.Lazy.Types, Data.List.ZipList, Data.Monoid, Data.Newtype, Data.Ord, Data.Semigroup, Data.Show, Data.Traversable, Partial.Unsafe, Prelude, Prim, Prim.TypeError
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.List.Lazy, Data.List.Lazy.Types, Data.List.ZipList, Data.Monoid, Data.Newtype, Data.Ord, Data.Semigroup, Data.Show, Data.Traversable, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.List.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Lazy.Types/index.php';
require_once __DIR__ . '/../Data.List.ZipList/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_List_ZipList_ZipList
function majData_majList_majZipmajList_majZipmajList($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majZipmajList_majZipmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_ZipList_ZipList'] = __NAMESPACE__ . '\\majData_majList_majZipmajList_majZipmajList';

// Data_List_ZipList_traversableZipList
$GLOBALS['Data_List_ZipList_traversableZipList'] = $GLOBALS['Data_List_Lazy_Types_traversableList'];

// Data_List_ZipList_showZipList
function majData_majList_majZipmajList_showmajZipmajList($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majZipmajList_showmajZipmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $showList_1_0 = (object)["show" => function($xs_1) use ($dictShow_0) {
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
  $__res = (object)["show" => function($v_2) use ($showList_1_0) {
  $__num = \func_num_args();
  $__res = (("(ZipList " . (($showList_1_0)->{'show'})($v_2)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_ZipList_showZipList'] = __NAMESPACE__ . '\\majData_majList_majZipmajList_showmajZipmajList';

// Data_List_ZipList_semigroupZipList
$GLOBALS['Data_List_ZipList_semigroupZipList'] = $GLOBALS['Data_List_Lazy_Types_semigroupList'];

// Data_List_ZipList_ordZipList
function majData_majList_majZipmajList_ordmajZipmajList($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majZipmajList_ordmajZipmajList';
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
$GLOBALS['Data_List_ZipList_ordZipList'] = __NAMESPACE__ . '\\majData_majList_majZipmajList_ordmajZipmajList';

// Data_List_ZipList_newtypeZipList
$GLOBALS['Data_List_ZipList_newtypeZipList'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_ZipList_monoidZipList
$GLOBALS['Data_List_ZipList_monoidZipList'] = $GLOBALS['Data_List_Lazy_Types_monoidList'];

// Data_List_ZipList_functorZipList
$GLOBALS['Data_List_ZipList_functorZipList'] = $GLOBALS['Data_List_Lazy_Types_functorList'];

// Data_List_ZipList_foldableZipList
$GLOBALS['Data_List_ZipList_foldableZipList'] = $GLOBALS['Data_List_Lazy_Types_foldableList'];

// Data_List_ZipList_eqZipList
function majData_majList_majZipmajList_eqmajZipmajList($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majZipmajList_eqmajZipmajList';
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
$GLOBALS['Data_List_ZipList_eqZipList'] = __NAMESPACE__ . '\\majData_majList_majZipmajList_eqmajZipmajList';

// Data_List_ZipList_applyZipList
$GLOBALS['Data_List_ZipList_applyZipList'] = (object)["apply" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = \Data\List\Lazy\majData_majList_majLazy_zipmajWith($GLOBALS['Data_Function_apply'], $v_0, $v1_1);
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

// Data_List_ZipList_zipListIsNotBind
function majData_majList_majZipmajList_zipmajListmajIsmajNotmajBind($_dollar___unused_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majZipmajList_zipmajListmajIsmajNotmajBind';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["bind" => \Partial\majPartial__crashmajWith("bind: unreachable"), "Apply0" => function($_dollar___unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_ZipList_applyZipList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_ZipList_zipListIsNotBind'] = __NAMESPACE__ . '\\majData_majList_majZipmajList_zipmajListmajIsmajNotmajBind';

// Data_List_ZipList_applicativeZipList
$GLOBALS['Data_List_ZipList_applicativeZipList'] = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_0) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_List_Lazy_repeat']), "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_ZipList_applyZipList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_ZipList_altZipList
$GLOBALS['Data_List_ZipList_altZipList'] = (object)["alt" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = \Data\List\Lazy\majData_majList_majLazy_drop(\Data\List\Lazy\majData_majList_majLazy_length($v_0), $v1_1);
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_3) use ($__local_var_2_0, $v_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = \Data\Lazy\majData_majLazy_force($v_0);
  $__t2 = null;;
  if ($__local_var_4_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($__local_var_2_0);
goto end_branch_2;;
};
  if ($__local_var_4_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t2 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($__local_var_4_1)->{'value0'}, ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})(($__local_var_4_1)->{'value1'}))($__local_var_2_0));
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

// Data_List_ZipList_plusZipList
$GLOBALS['Data_List_ZipList_plusZipList'] = (object)["empty" => $GLOBALS['Data_List_Lazy_Types_nil'], "Alt0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_ZipList_altZipList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_ZipList_alternativeZipList
$GLOBALS['Data_List_ZipList_alternativeZipList'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_ZipList_applicativeZipList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_ZipList_plusZipList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

