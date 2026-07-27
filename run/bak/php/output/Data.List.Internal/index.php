<?php

namespace Data\List\Internal;

// ALL IMPORTS: Data.List.Internal, Data.List.Types, Data.Ordering, Prelude, Prim
// TO REQUIRE: Data.List.Internal, Data.List.Types, Data.Ordering, Prelude
require_once __DIR__ . '/../Data.List.Internal/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
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


final class Data_List_Internal_Leaf { public $tag = 'Leaf'; public function __construct() {} }
final class Data_List_Internal_Two { public $tag = 'Two'; public function __construct(public  $value0, public  $value1, public  $value2) {} }
final class Data_List_Internal_Three { public $tag = 'Three'; public function __construct(public  $value0, public  $value1, public  $value2, public  $value3, public  $value4) {} }
final class Data_List_Internal_TwoLeft { public $tag = 'TwoLeft'; public function __construct(public  $value0, public  $value1) {} }
final class Data_List_Internal_TwoRight { public $tag = 'TwoRight'; public function __construct(public  $value0, public  $value1) {} }
final class Data_List_Internal_ThreeLeft { public $tag = 'ThreeLeft'; public function __construct(public  $value0, public  $value1, public  $value2, public  $value3) {} }
final class Data_List_Internal_ThreeMiddle { public $tag = 'ThreeMiddle'; public function __construct(public  $value0, public  $value1, public  $value2, public  $value3) {} }
final class Data_List_Internal_ThreeRight { public $tag = 'ThreeRight'; public function __construct(public  $value0, public  $value1, public  $value2, public  $value3) {} }
final class Data_List_Internal_KickUp { public $tag = 'KickUp'; public function __construct(public  $value0, public  $value1, public  $value2) {} }

// Data_List_Internal_Leaf
$GLOBALS['Data_List_Internal_Leaf'] = ($GLOBALS['__phpurs_data0_Leaf'] ??= new \Data\List\Internal\Data_List_Internal_Leaf());

// Data_List_Internal_Two
$GLOBALS['Data_List_Internal_Two'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\List\Internal\Data_List_Internal_Two($value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Internal_Three
$GLOBALS['Data_List_Internal_Three'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null, $value4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__res = new \Data\List\Internal\Data_List_Internal_Three($value0, $value1, $value2, $value3, $value4);
  goto __end;;
  __end:
  return $__num > 5 ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
  };
  return $__fn;
})();

// Data_List_Internal_TwoLeft
$GLOBALS['Data_List_Internal_TwoLeft'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Internal\Data_List_Internal_TwoLeft($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Internal_TwoRight
$GLOBALS['Data_List_Internal_TwoRight'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Internal\Data_List_Internal_TwoRight($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Internal_ThreeLeft
$GLOBALS['Data_List_Internal_ThreeLeft'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new \Data\List\Internal\Data_List_Internal_ThreeLeft($value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_List_Internal_ThreeMiddle
$GLOBALS['Data_List_Internal_ThreeMiddle'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new \Data\List\Internal\Data_List_Internal_ThreeMiddle($value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_List_Internal_ThreeRight
$GLOBALS['Data_List_Internal_ThreeRight'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new \Data\List\Internal\Data_List_Internal_ThreeRight($value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_List_Internal_KickUp
$GLOBALS['Data_List_Internal_KickUp'] = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\List\Internal\Data_List_Internal_KickUp($value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_List_Internal_fromZipper
function majData_majList_majInternal_frommajZipper($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majInternal_frommajZipper';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_Internal_fromZipper_v_0 = $v_0;
  $__tco_var_Data_List_Internal_fromZipper_v1_1 = $v1_1;
  tco_loop_Data_List_Internal_fromZipper:;
  $v_0 = $__tco_var_Data_List_Internal_fromZipper_v_0;
  $v1_1 = $__tco_var_Data_List_Internal_fromZipper_v1_1;
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = null;;
if (($v_0)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_TwoLeft) {
$__tco_2 = ($v_0)->{'value1'};
$__tco_3 = new \Data\List\Internal\Data_List_Internal_Two($v1_1, (($v_0)->{'value0'})->{'value0'}, (($v_0)->{'value0'})->{'value1'});
$__tco_var_Data_List_Internal_fromZipper_v_0 = $__tco_2;
$__tco_var_Data_List_Internal_fromZipper_v1_1 = $__tco_3;
goto tco_loop_Data_List_Internal_fromZipper;;
$__t1 = null;
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_TwoRight) {
$__tco_4 = ($v_0)->{'value1'};
$__tco_5 = new \Data\List\Internal\Data_List_Internal_Two((($v_0)->{'value0'})->{'value0'}, (($v_0)->{'value0'})->{'value1'}, $v1_1);
$__tco_var_Data_List_Internal_fromZipper_v_0 = $__tco_4;
$__tco_var_Data_List_Internal_fromZipper_v1_1 = $__tco_5;
goto tco_loop_Data_List_Internal_fromZipper;;
$__t1 = null;
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_ThreeLeft) {
$__tco_6 = ($v_0)->{'value1'};
$__tco_7 = new \Data\List\Internal\Data_List_Internal_Three($v1_1, (($v_0)->{'value0'})->{'value0'}, (($v_0)->{'value0'})->{'value1'}, (($v_0)->{'value0'})->{'value2'}, (($v_0)->{'value0'})->{'value3'});
$__tco_var_Data_List_Internal_fromZipper_v_0 = $__tco_6;
$__tco_var_Data_List_Internal_fromZipper_v1_1 = $__tco_7;
goto tco_loop_Data_List_Internal_fromZipper;;
$__t1 = null;
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_ThreeMiddle) {
$__tco_8 = ($v_0)->{'value1'};
$__tco_9 = new \Data\List\Internal\Data_List_Internal_Three((($v_0)->{'value0'})->{'value0'}, (($v_0)->{'value0'})->{'value1'}, $v1_1, (($v_0)->{'value0'})->{'value2'}, (($v_0)->{'value0'})->{'value3'});
$__tco_var_Data_List_Internal_fromZipper_v_0 = $__tco_8;
$__tco_var_Data_List_Internal_fromZipper_v1_1 = $__tco_9;
goto tco_loop_Data_List_Internal_fromZipper;;
$__t1 = null;
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_ThreeRight) {
$__tco_10 = ($v_0)->{'value1'};
$__tco_11 = new \Data\List\Internal\Data_List_Internal_Three((($v_0)->{'value0'})->{'value0'}, (($v_0)->{'value0'})->{'value1'}, (($v_0)->{'value0'})->{'value2'}, (($v_0)->{'value0'})->{'value3'}, $v1_1);
$__tco_var_Data_List_Internal_fromZipper_v_0 = $__tco_10;
$__tco_var_Data_List_Internal_fromZipper_v1_1 = $__tco_11;
goto tco_loop_Data_List_Internal_fromZipper;;
$__t1 = null;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Internal_fromZipper'] = __NAMESPACE__ . '\\majData_majList_majInternal_frommajZipper';

// Data_List_Internal_insertAndLookupBy
function majData_majList_majInternal_insertmajAndmajLookupmajBy($comp_0, $k_1 = null, $orig_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majInternal_insertmajAndmajLookupmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $up_3_0 = null;
  $up_3_0 = (function() use (&$up_3_0) {
  $__fn = function($v_4, $v1_5 = null) use (&$up_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_up_3_0_0_v_4 = $v_4;
  $__tco_var_up_3_0_0_v1_5 = $v1_5;
  tco_loop_up_3_0_0:;
  $v_4 = $__tco_var_up_3_0_0_v_4;
  $v1_5 = $__tco_var_up_3_0_0_v1_5;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Internal\Data_List_Internal_Two(($v1_5)->{'value0'}, ($v1_5)->{'value1'}, ($v1_5)->{'value2'});
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = null;;
if (($v_4)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_TwoLeft) {
$__t1 = \Data\List\Internal\majData_majList_majInternal_frommajZipper(($v_4)->{'value1'}, new \Data\List\Internal\Data_List_Internal_Three(($v1_5)->{'value0'}, ($v1_5)->{'value1'}, ($v1_5)->{'value2'}, (($v_4)->{'value0'})->{'value0'}, (($v_4)->{'value0'})->{'value1'}));
goto end_branch_1;;
};
if (($v_4)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_TwoRight) {
$__t1 = \Data\List\Internal\majData_majList_majInternal_frommajZipper(($v_4)->{'value1'}, new \Data\List\Internal\Data_List_Internal_Three((($v_4)->{'value0'})->{'value0'}, (($v_4)->{'value0'})->{'value1'}, ($v1_5)->{'value0'}, ($v1_5)->{'value1'}, ($v1_5)->{'value2'}));
goto end_branch_1;;
};
if (($v_4)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_ThreeLeft) {
$__tco_2 = ($v_4)->{'value1'};
$__tco_3 = new \Data\List\Internal\Data_List_Internal_KickUp(new \Data\List\Internal\Data_List_Internal_Two(($v1_5)->{'value0'}, ($v1_5)->{'value1'}, ($v1_5)->{'value2'}), (($v_4)->{'value0'})->{'value0'}, new \Data\List\Internal\Data_List_Internal_Two((($v_4)->{'value0'})->{'value1'}, (($v_4)->{'value0'})->{'value2'}, (($v_4)->{'value0'})->{'value3'}));
$__tco_var_up_3_0_0_v_4 = $__tco_2;
$__tco_var_up_3_0_0_v1_5 = $__tco_3;
goto tco_loop_up_3_0_0;;
$__t1 = null;
goto end_branch_1;;
};
if (($v_4)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_ThreeMiddle) {
$__tco_4 = ($v_4)->{'value1'};
$__tco_5 = new \Data\List\Internal\Data_List_Internal_KickUp(new \Data\List\Internal\Data_List_Internal_Two((($v_4)->{'value0'})->{'value0'}, (($v_4)->{'value0'})->{'value1'}, ($v1_5)->{'value0'}), ($v1_5)->{'value1'}, new \Data\List\Internal\Data_List_Internal_Two(($v1_5)->{'value2'}, (($v_4)->{'value0'})->{'value2'}, (($v_4)->{'value0'})->{'value3'}));
$__tco_var_up_3_0_0_v_4 = $__tco_4;
$__tco_var_up_3_0_0_v1_5 = $__tco_5;
goto tco_loop_up_3_0_0;;
$__t1 = null;
goto end_branch_1;;
};
if (($v_4)->{'value0'} instanceof \Data\List\Internal\Data_List_Internal_ThreeRight) {
$__tco_6 = ($v_4)->{'value1'};
$__tco_7 = new \Data\List\Internal\Data_List_Internal_KickUp(new \Data\List\Internal\Data_List_Internal_Two((($v_4)->{'value0'})->{'value0'}, (($v_4)->{'value0'})->{'value1'}, (($v_4)->{'value0'})->{'value2'}), (($v_4)->{'value0'})->{'value3'}, new \Data\List\Internal\Data_List_Internal_Two(($v1_5)->{'value0'}, ($v1_5)->{'value1'}, ($v1_5)->{'value2'}));
$__tco_var_up_3_0_0_v_4 = $__tco_6;
$__tco_var_up_3_0_0_v1_5 = $__tco_7;
goto tco_loop_up_3_0_0;;
$__t1 = null;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $down_4_1 = null;
  $down_4_1 = (function() use ($comp_0, &$down_4_1, $k_1, $orig_2, &$up_3_0) {
  $__fn = function($v_5, $v1_6 = null) use ($comp_0, &$down_4_1, $k_1, $orig_2, &$up_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_down_4_1_1_v_5 = $v_5;
  $__tco_var_down_4_1_1_v1_6 = $v1_6;
  tco_loop_down_4_1_1:;
  $v_5 = $__tco_var_down_4_1_1_v_5;
  $v1_6 = $__tco_var_down_4_1_1_v1_6;
  $__t1 = null;;
  if ($v1_6 instanceof \Data\List\Internal\Data_List_Internal_Leaf) {
$__t1 = (object)["found" => false, "result" => (($up_3_0)($v_5))(new \Data\List\Internal\Data_List_Internal_KickUp(new \Data\List\Internal\Data_List_Internal_Leaf(), $k_1, new \Data\List\Internal\Data_List_Internal_Leaf()))];
goto end_branch_1;;
};
  if ($v1_6 instanceof \Data\List\Internal\Data_List_Internal_Two) {
$v2_7_2 = (($comp_0)($k_1))(($v1_6)->{'value1'});
$__t5 = null;;
if ($v2_7_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t5 = (object)["found" => true, "result" => $orig_2];
goto end_branch_5;;
};
if ($v2_7_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_6 = new \Data\List\Types\Data_List_Types_Cons(new \Data\List\Internal\Data_List_Internal_TwoLeft(($v1_6)->{'value1'}, ($v1_6)->{'value2'}), $v_5);
$__tco_7 = ($v1_6)->{'value0'};
$__tco_var_down_4_1_1_v_5 = $__tco_6;
$__tco_var_down_4_1_1_v1_6 = $__tco_7;
goto tco_loop_down_4_1_1;;
$__t5 = null;
goto end_branch_5;;
};
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(new \Data\List\Internal\Data_List_Internal_TwoRight(($v1_6)->{'value0'}, ($v1_6)->{'value1'}), $v_5);
$__tco_4 = ($v1_6)->{'value2'};
$__tco_var_down_4_1_1_v_5 = $__tco_3;
$__tco_var_down_4_1_1_v1_6 = $__tco_4;
goto tco_loop_down_4_1_1;;
$__t5 = null;
end_branch_5:;
$__t1 = $__t5;
goto end_branch_1;;
};
  if ($v1_6 instanceof \Data\List\Internal\Data_List_Internal_Three) {
$v2_7_8 = (($comp_0)($k_1))(($v1_6)->{'value1'});
$__t17 = null;;
if ($v2_7_8 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t17 = (object)["found" => true, "result" => $orig_2];
goto end_branch_17;;
};
$v3_8_9 = (($comp_0)($k_1))(($v1_6)->{'value3'});
$__t12 = null;;
if ($v3_8_9 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t12 = (object)["found" => true, "result" => $orig_2];
goto end_branch_12;;
};
if ($v2_7_8 instanceof \Data\Ordering\Data_Ordering_LT) {
$__tco_13 = new \Data\List\Types\Data_List_Types_Cons(new \Data\List\Internal\Data_List_Internal_ThreeLeft(($v1_6)->{'value1'}, ($v1_6)->{'value2'}, ($v1_6)->{'value3'}, ($v1_6)->{'value4'}), $v_5);
$__tco_14 = ($v1_6)->{'value0'};
$__tco_var_down_4_1_1_v_5 = $__tco_13;
$__tco_var_down_4_1_1_v1_6 = $__tco_14;
goto tco_loop_down_4_1_1;;
$__t12 = null;
goto end_branch_12;;
};
if (($v2_7_8 instanceof \Data\Ordering\Data_Ordering_GT && $v3_8_9 instanceof \Data\Ordering\Data_Ordering_LT)) {
$__tco_15 = new \Data\List\Types\Data_List_Types_Cons(new \Data\List\Internal\Data_List_Internal_ThreeMiddle(($v1_6)->{'value0'}, ($v1_6)->{'value1'}, ($v1_6)->{'value3'}, ($v1_6)->{'value4'}), $v_5);
$__tco_16 = ($v1_6)->{'value2'};
$__tco_var_down_4_1_1_v_5 = $__tco_15;
$__tco_var_down_4_1_1_v1_6 = $__tco_16;
goto tco_loop_down_4_1_1;;
$__t12 = null;
goto end_branch_12;;
};
$__tco_10 = new \Data\List\Types\Data_List_Types_Cons(new \Data\List\Internal\Data_List_Internal_ThreeRight(($v1_6)->{'value0'}, ($v1_6)->{'value1'}, ($v1_6)->{'value2'}, ($v1_6)->{'value3'}), $v_5);
$__tco_11 = ($v1_6)->{'value4'};
$__tco_var_down_4_1_1_v_5 = $__tco_10;
$__tco_var_down_4_1_1_v1_6 = $__tco_11;
goto tco_loop_down_4_1_1;;
$__t12 = null;
end_branch_12:;
$__t17 = $__t12;
end_branch_17:;
$__t1 = $__t17;
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
  $__res = (($down_4_1)(new \Data\List\Types\Data_List_Types_Nil()))($orig_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Internal_insertAndLookupBy'] = __NAMESPACE__ . '\\majData_majList_majInternal_insertmajAndmajLookupmajBy';

// Data_List_Internal_emptySet
$GLOBALS['Data_List_Internal_emptySet'] = new \Data\List\Internal\Data_List_Internal_Leaf();

