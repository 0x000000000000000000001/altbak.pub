<?php

namespace Data\NonEmpty;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semigroup.Foldable, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semigroup.Foldable, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
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


final class Data_NonEmpty_NonEmpty { public $tag = 'NonEmpty'; public function __construct(public  $value0, public  $value1) {} }

// Data_NonEmpty_NonEmpty
$GLOBALS['Data_NonEmpty_NonEmpty'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_NonEmpty_unfoldable1NonEmpty
function majData_majNonmajEmpty_unfoldable1majNonmajEmpty($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_unfoldable1majNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["unfoldr1" => function($f_1) use ($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($dictUnfoldable_0, $f_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = ($f_1)($b_2);
  $__local_var_3_0 = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_3_0)->{'value0'}, ((($dictUnfoldable_0)->{'unfoldr'})(function($v1_4) use ($f_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($f_1)(($v1_4)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($__local_var_3_0)->{'value1'}));
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($__local_var_3_0)->{'value0'}, ($__local_var_3_0)->{'value1'});
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
$GLOBALS['Data_NonEmpty_unfoldable1NonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_unfoldable1majNonmajEmpty';

// Data_NonEmpty_tail
function majData_majNonmajEmpty_tail($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_tail';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_tail'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_tail';

// Data_NonEmpty_singleton
function majData_majNonmajEmpty_singleton($dictPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $empty_1_0 = ($dictPlus_0)->{'empty'};
  $__res = function($a_2) use ($empty_1_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_2, $empty_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_singleton'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_singleton';

// Data_NonEmpty_showNonEmpty
function majData_majNonmajEmpty_showmajNonmajEmpty($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_showmajNonmajEmpty';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = (((("(NonEmpty " . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow1_1)->{'show'})(($v_2)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_NonEmpty_showNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_showmajNonmajEmpty';

// Data_NonEmpty_semigroupNonEmpty
function majData_majNonmajEmpty_semigroupmajNonmajEmpty($dictApplicative_0, $dictSemigroup_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_semigroupmajNonmajEmpty';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["append" => function($v_2) use ($dictApplicative_0, $dictSemigroup_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictApplicative_0, $dictSemigroup_1, $v_2) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_2)->{'value0'}, ((($dictSemigroup_1)->{'append'})(($v_2)->{'value1'}))(((($dictSemigroup_1)->{'append'})((($dictApplicative_0)->{'pure'})(($v1_3)->{'value0'})))(($v1_3)->{'value1'})));
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
$GLOBALS['Data_NonEmpty_semigroupNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_semigroupmajNonmajEmpty';

// Data_NonEmpty_oneOf
function majData_majNonmajEmpty_onemajOf($dictAlternative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_onemajOf';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Alt0_1_0 = (((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null);
  $Applicative0_2_1 = (($dictAlternative_0)->{'Applicative0'})(null);
  $__res = function($v_3) use ($Alt0_1_0, $Applicative0_2_1) {
  $__num = \func_num_args();
  $__res = ((($Alt0_1_0)->{'alt'})((($Applicative0_2_1)->{'pure'})(($v_3)->{'value0'})))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_oneOf'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_onemajOf';

// Data_NonEmpty_head
function majData_majNonmajEmpty_head($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_head';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_head'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_head';

// Data_NonEmpty_functorNonEmpty
function majData_majNonmajEmpty_functormajNonmajEmpty($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_functormajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($m_2) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_1)(($m_2)->{'value0'}), ((($dictFunctor_0)->{'map'})($f_1))(($m_2)->{'value1'}));
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
$GLOBALS['Data_NonEmpty_functorNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_functormajNonmajEmpty';

// Data_NonEmpty_functorWithIndex
function majData_majNonmajEmpty_functormajWithmajIndex($dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_functormajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFunctorWithIndex_0)->{'Functor0'})(null);
  $functorNonEmpty1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($m_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_2)(($m_3)->{'value0'}), ((($__local_var_1_0)->{'map'})($f_2))(($m_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mapWithIndex" => function($f_2) use ($dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFunctorWithIndex_0, $f_2) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty((($f_2)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_3)->{'value0'}), ((($dictFunctorWithIndex_0)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))($GLOBALS['Data_Maybe_Just'])))(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_functorWithIndex'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_functormajWithmajIndex';

// Data_NonEmpty_fromNonEmpty
function majData_majNonmajEmpty_frommajNonmajEmpty($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_frommajNonmajEmpty';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)(($v_1)->{'value0'}))(($v_1)->{'value1'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_NonEmpty_fromNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_frommajNonmajEmpty';

// Data_NonEmpty_foldableNonEmpty
function majData_majNonmajEmpty_foldablemajNonmajEmpty($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_foldablemajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["foldMap" => function($dictMonoid_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $Semigroup0_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $__res = function($f_3) use ($Semigroup0_2_0, $dictFoldable_0, $dictMonoid_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($Semigroup0_2_0, $dictFoldable_0, $dictMonoid_1, $f_3) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_2_0)->{'append'})(($f_3)(($v_4)->{'value0'})))((((($dictFoldable_0)->{'foldMap'})($dictMonoid_1))($f_3))(($v_4)->{'value1'}));
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
}, "foldl" => function($f_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldl'})($f_1))((($f_1)($b_2))(($v_3)->{'value0'})))(($v_3)->{'value1'});
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
}, "foldr" => function($f_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = (($f_1)(($v_3)->{'value0'}))((((($dictFoldable_0)->{'foldr'})($f_1))($b_2))(($v_3)->{'value1'}));
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
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_foldableNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_foldablemajNonmajEmpty';

// Data_NonEmpty_foldableWithIndexNonEmpty
function majData_majNonmajEmpty_foldablemajWithmajIndexmajNonmajEmpty($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_foldablemajWithmajIndexmajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFoldableWithIndex_0)->{'Foldable0'})(null);
  $foldableNonEmpty1_1_0 = (object)["foldMap" => function($dictMonoid_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $Semigroup0_3_1 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $__res = function($f_4) use ($Semigroup0_3_1, $__local_var_1_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Semigroup0_3_1, $__local_var_1_0, $dictMonoid_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_3_1)->{'append'})(($f_4)(($v_5)->{'value0'})))((((($__local_var_1_0)->{'foldMap'})($dictMonoid_2))($f_4))(($v_5)->{'value1'}));
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
}, "foldl" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $b_3, $f_2) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'foldl'})($f_2))((($f_2)($b_3))(($v_4)->{'value0'})))(($v_4)->{'value1'});
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
}, "foldr" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $b_3, $f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)(($v_4)->{'value0'}))((((($__local_var_1_0)->{'foldr'})($f_2))($b_3))(($v_4)->{'value1'}));
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
  $__res = (object)["foldMapWithIndex" => function($dictMonoid_2) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $Semigroup0_3_3 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $__res = function($f_4) use ($Semigroup0_3_3, $dictFoldableWithIndex_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Semigroup0_3_3, $dictFoldableWithIndex_0, $dictMonoid_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_3_3)->{'append'})((($f_4)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_5)->{'value0'})))((((($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_2))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Maybe_Just'])))(($v_5)->{'value1'}));
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
}, "foldlWithIndex" => function($f_2) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($dictFoldableWithIndex_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($b_3, $dictFoldableWithIndex_0, $f_2) {
  $__num = \func_num_args();
  $__res = (((($dictFoldableWithIndex_0)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))($GLOBALS['Data_Maybe_Just'])))(((($f_2)(new \Data\Maybe\Data_Maybe_Nothing()))($b_3))(($v_4)->{'value0'})))(($v_4)->{'value1'});
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
}, "foldrWithIndex" => function($f_2) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($dictFoldableWithIndex_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($b_3, $dictFoldableWithIndex_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($f_2)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_4)->{'value0'}))((((($dictFoldableWithIndex_0)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))($GLOBALS['Data_Maybe_Just'])))($b_3))(($v_4)->{'value1'}));
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
}, "Foldable0" => function($_dollar___unused_2) use ($foldableNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_foldableWithIndexNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_foldablemajWithmajIndexmajNonmajEmpty';

// Data_NonEmpty_traversableNonEmpty
function majData_majNonmajEmpty_traversablemajNonmajEmpty($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_traversablemajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $functorNonEmpty1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($m_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_2)(($m_3)->{'value0'}), ((($__local_var_1_0)->{'map'})($f_2))(($m_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_2 = (($dictTraversable_0)->{'Foldable1'})(null);
  $foldableNonEmpty1_2_2 = (object)["foldMap" => function($dictMonoid_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $Semigroup0_4_3 = (($dictMonoid_3)->{'Semigroup0'})(null);
  $__res = function($f_5) use ($Semigroup0_4_3, $__local_var_2_2, $dictMonoid_3) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($Semigroup0_4_3, $__local_var_2_2, $dictMonoid_3, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_4_3)->{'append'})(($f_5)(($v_6)->{'value0'})))((((($__local_var_2_2)->{'foldMap'})($dictMonoid_3))($f_5))(($v_6)->{'value1'}));
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
}, "foldl" => function($f_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($__local_var_2_2, $f_3) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_2_2, $b_4, $f_3) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_2)->{'foldl'})($f_3))((($f_3)($b_4))(($v_5)->{'value0'})))(($v_5)->{'value1'});
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
}, "foldr" => function($f_3) use ($__local_var_2_2) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($__local_var_2_2, $f_3) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_2_2, $b_4, $f_3) {
  $__num = \func_num_args();
  $__res = (($f_3)(($v_5)->{'value0'}))((((($__local_var_2_2)->{'foldr'})($f_3))($b_4))(($v_5)->{'value1'}));
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
  $__res = (object)["sequence" => function($dictApplicative_3) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $Apply0_4_5 = (($dictApplicative_3)->{'Apply0'})(null);
  $Functor0_5_6 = (((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_6) use ($Apply0_4_5, $Functor0_5_6, $dictApplicative_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = ((($Apply0_4_5)->{'apply'})(((($Functor0_5_6)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($v_6)->{'value0'})))(((($dictTraversable_0)->{'sequence'})($dictApplicative_3))(($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "traverse" => function($dictApplicative_3) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $Apply0_4_7 = (($dictApplicative_3)->{'Apply0'})(null);
  $Functor0_5_8 = (((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_6) use ($Apply0_4_7, $Functor0_5_8, $dictApplicative_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Apply0_4_7, $Functor0_5_8, $dictApplicative_3, $dictTraversable_0, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Apply0_4_7)->{'apply'})(((($Functor0_5_8)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($f_6)(($v_7)->{'value0'}))))((((($dictTraversable_0)->{'traverse'})($dictApplicative_3))($f_6))(($v_7)->{'value1'}));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_3) use ($foldableNonEmpty1_2_2) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_2_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_traversableNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_traversablemajNonmajEmpty';

// Data_NonEmpty_traversableWithIndexNonEmpty
function majData_majNonmajEmpty_traversablemajWithmajIndexmajNonmajEmpty($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_traversablemajWithmajIndexmajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorNonEmpty1_2_1 = (object)["map" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($m_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_3)(($m_4)->{'value0'}), ((($__local_var_2_1)->{'map'})($f_3))(($m_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorWithIndex1_1_0 = (object)["mapWithIndex" => function($f_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty((($f_3)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_4)->{'value0'}), ((($__local_var_1_0)->{'mapWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_3))($GLOBALS['Data_Maybe_Just'])))(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorNonEmpty1_2_1) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_4 = (($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null);
  $__local_var_3_5 = (($__local_var_2_4)->{'Foldable0'})(null);
  $foldableNonEmpty1_3_5 = (object)["foldMap" => function($dictMonoid_4) use ($__local_var_3_5) {
  $__num = \func_num_args();
  $Semigroup0_5_6 = (($dictMonoid_4)->{'Semigroup0'})(null);
  $__res = function($f_6) use ($Semigroup0_5_6, $__local_var_3_5, $dictMonoid_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Semigroup0_5_6, $__local_var_3_5, $dictMonoid_4, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_5_6)->{'append'})(($f_6)(($v_7)->{'value0'})))((((($__local_var_3_5)->{'foldMap'})($dictMonoid_4))($f_6))(($v_7)->{'value1'}));
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
}, "foldl" => function($f_4) use ($__local_var_3_5) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($__local_var_3_5, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_3_5, $b_5, $f_4) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_5)->{'foldl'})($f_4))((($f_4)($b_5))(($v_6)->{'value0'})))(($v_6)->{'value1'});
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
}, "foldr" => function($f_4) use ($__local_var_3_5) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($__local_var_3_5, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_3_5, $b_5, $f_4) {
  $__num = \func_num_args();
  $__res = (($f_4)(($v_6)->{'value0'}))((((($__local_var_3_5)->{'foldr'})($f_4))($b_5))(($v_6)->{'value1'}));
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
  $foldableWithIndexNonEmpty1_2_4 = (object)["foldMapWithIndex" => function($dictMonoid_4) use ($__local_var_2_4) {
  $__num = \func_num_args();
  $Semigroup0_5_8 = (($dictMonoid_4)->{'Semigroup0'})(null);
  $__res = function($f_6) use ($Semigroup0_5_8, $__local_var_2_4, $dictMonoid_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Semigroup0_5_8, $__local_var_2_4, $dictMonoid_4, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_5_8)->{'append'})((($f_6)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_7)->{'value0'})))((((($__local_var_2_4)->{'foldMapWithIndex'})($dictMonoid_4))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_6))($GLOBALS['Data_Maybe_Just'])))(($v_7)->{'value1'}));
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
}, "foldlWithIndex" => function($f_4) use ($__local_var_2_4) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($__local_var_2_4, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_2_4, $b_5, $f_4) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_4)->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Maybe_Just'])))(((($f_4)(new \Data\Maybe\Data_Maybe_Nothing()))($b_5))(($v_6)->{'value0'})))(($v_6)->{'value1'});
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
}, "foldrWithIndex" => function($f_4) use ($__local_var_2_4) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($__local_var_2_4, $f_4) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_2_4, $b_5, $f_4) {
  $__num = \func_num_args();
  $__res = ((($f_4)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_6)->{'value0'}))((((($__local_var_2_4)->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Maybe_Just'])))($b_5))(($v_6)->{'value1'}));
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
}, "Foldable0" => function($_dollar___unused_4) use ($foldableNonEmpty1_3_5) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_3_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_10 = (($dictTraversableWithIndex_0)->{'Traversable2'})(null);
  $__local_var_4_11 = (($__local_var_3_10)->{'Functor0'})(null);
  $functorNonEmpty1_4_11 = (object)["map" => function($f_5) use ($__local_var_4_11) {
  $__num = \func_num_args();
  $__res = function($m_6) use ($__local_var_4_11, $f_5) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_5)(($m_6)->{'value0'}), ((($__local_var_4_11)->{'map'})($f_5))(($m_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_5_13 = (($__local_var_3_10)->{'Foldable1'})(null);
  $foldableNonEmpty1_5_13 = (object)["foldMap" => function($dictMonoid_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $Semigroup0_7_14 = (($dictMonoid_6)->{'Semigroup0'})(null);
  $__res = function($f_8) use ($Semigroup0_7_14, $__local_var_5_13, $dictMonoid_6) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($Semigroup0_7_14, $__local_var_5_13, $dictMonoid_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_7_14)->{'append'})(($f_8)(($v_9)->{'value0'})))((((($__local_var_5_13)->{'foldMap'})($dictMonoid_6))($f_8))(($v_9)->{'value1'}));
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
}, "foldl" => function($f_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($__local_var_5_13, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_5_13, $b_7, $f_6) {
  $__num = \func_num_args();
  $__res = (((($__local_var_5_13)->{'foldl'})($f_6))((($f_6)($b_7))(($v_8)->{'value0'})))(($v_8)->{'value1'});
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
}, "foldr" => function($f_6) use ($__local_var_5_13) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($__local_var_5_13, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_5_13, $b_7, $f_6) {
  $__num = \func_num_args();
  $__res = (($f_6)(($v_8)->{'value0'}))((((($__local_var_5_13)->{'foldr'})($f_6))($b_7))(($v_8)->{'value1'}));
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
  $traversableNonEmpty1_3_10 = (object)["sequence" => function($dictApplicative_6) use ($__local_var_3_10) {
  $__num = \func_num_args();
  $Apply0_7_16 = (($dictApplicative_6)->{'Apply0'})(null);
  $Functor0_8_17 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_9) use ($Apply0_7_16, $Functor0_8_17, $__local_var_3_10, $dictApplicative_6) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_16)->{'apply'})(((($Functor0_8_17)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($v_9)->{'value0'})))(((($__local_var_3_10)->{'sequence'})($dictApplicative_6))(($v_9)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "traverse" => function($dictApplicative_6) use ($__local_var_3_10) {
  $__num = \func_num_args();
  $Apply0_7_18 = (($dictApplicative_6)->{'Apply0'})(null);
  $Functor0_8_19 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_9) use ($Apply0_7_18, $Functor0_8_19, $__local_var_3_10, $dictApplicative_6) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($Apply0_7_18, $Functor0_8_19, $__local_var_3_10, $dictApplicative_6, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_18)->{'apply'})(((($Functor0_8_19)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($f_9)(($v_10)->{'value0'}))))((((($__local_var_3_10)->{'traverse'})($dictApplicative_6))($f_9))(($v_10)->{'value1'}));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorNonEmpty1_4_11) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_4_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_6) use ($foldableNonEmpty1_5_13) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_5_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["traverseWithIndex" => function($dictApplicative_4) use ($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $Apply0_5_21 = (($dictApplicative_4)->{'Apply0'})(null);
  $Functor0_6_22 = (((($dictApplicative_4)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_7) use ($Apply0_5_21, $Functor0_6_22, $dictApplicative_4, $dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Apply0_5_21, $Functor0_6_22, $dictApplicative_4, $dictTraversableWithIndex_0, $f_7) {
  $__num = \func_num_args();
  $__res = ((($Apply0_5_21)->{'apply'})(((($Functor0_6_22)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))((($f_7)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_8)->{'value0'}))))((((($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_4))((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_7))($GLOBALS['Data_Maybe_Just'])))(($v_8)->{'value1'}));
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
}, "FunctorWithIndex0" => function($_dollar___unused_4) use ($functorWithIndex1_1_0) {
  $__num = \func_num_args();
  $__res = $functorWithIndex1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_4) use ($foldableWithIndexNonEmpty1_2_4) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexNonEmpty1_2_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_4) use ($traversableNonEmpty1_3_10) {
  $__num = \func_num_args();
  $__res = $traversableNonEmpty1_3_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_traversableWithIndexNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_traversablemajWithmajIndexmajNonmajEmpty';

// Data_NonEmpty_foldable1NonEmpty
function majData_majNonmajEmpty_foldable1majNonmajEmpty($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_foldable1majNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldableNonEmpty1_1_0 = (object)["foldMap" => function($dictMonoid_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $Semigroup0_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $__res = function($f_3) use ($Semigroup0_2_0, $dictFoldable_0, $dictMonoid_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($Semigroup0_2_0, $dictFoldable_0, $dictMonoid_1, $f_3) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_2_0)->{'append'})(($f_3)(($v_4)->{'value0'})))((((($dictFoldable_0)->{'foldMap'})($dictMonoid_1))($f_3))(($v_4)->{'value1'}));
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
}, "foldl" => function($f_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldl'})($f_1))((($f_1)($b_2))(($v_3)->{'value0'})))(($v_3)->{'value1'});
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
}, "foldr" => function($f_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = (($f_1)(($v_3)->{'value0'}))((((($dictFoldable_0)->{'foldr'})($f_1))($b_2))(($v_3)->{'value1'}));
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
  $__res = (object)["foldMap1" => function($dictSemigroup_2) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictFoldable_0, $dictSemigroup_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldable_0, $dictSemigroup_2, $f_3) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldl'})(function($s_5) use ($dictSemigroup_2, $f_3) {
  $__num = \func_num_args();
  $__res = function($a1_6) use ($dictSemigroup_2, $f_3, $s_5) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_2)->{'append'})($s_5))(($f_3)($a1_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_3)(($v_4)->{'value0'})))(($v_4)->{'value1'});
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
}, "foldr1" => function($f_2) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFoldable_0, $f_2) {
  $__num = \func_num_args();
  $__local_var_4_2 = ($f_2)(($v_3)->{'value0'});
  $__local_var_5_3 = (((($dictFoldable_0)->{'foldr'})(function($a1_5) use ($f_2) {
  $__num = \func_num_args();
  $__local_var_6_3 = ($f_2)($a1_5);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))(function($v2_7) use ($__local_var_6_3, $a1_5) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($v2_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = $a1_5;
goto end_branch_4;;
};
  if ($v2_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = ($__local_var_6_3)(($v2_7)->{'value0'});
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Maybe\Data_Maybe_Nothing()))(($v_3)->{'value1'});
  $__t6 = null;;
  if ($__local_var_5_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t6 = ($v_3)->{'value0'};
goto end_branch_6;;
};
  if ($__local_var_5_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = ($__local_var_4_2)(($__local_var_5_3)->{'value0'});
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_2) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFoldable_0, $f_2) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldl'})($f_2))(($v_3)->{'value0'}))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_2) use ($foldableNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_foldable1NonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_foldable1majNonmajEmpty';

// Data_NonEmpty_foldl1
function majData_majNonmajEmpty_foldl1($dictFoldable_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_foldl1';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)->{'foldl'})($f_1))(($v_2)->{'value0'}))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_NonEmpty_foldl1'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_foldl1';

// Data_NonEmpty_eqNonEmpty
function majData_majNonmajEmpty_eqmajNonmajEmpty($dictEq1_0, $dictEq_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_eqmajNonmajEmpty';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq" => function($x_2) use ($dictEq1_0, $dictEq_1) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictEq1_0, $dictEq_1, $x_2) {
  $__num = \func_num_args();
  $__res = (((($dictEq_1)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && (((($dictEq1_0)->{'eq1'})($dictEq_1))(($x_2)->{'value1'}))(($y_3)->{'value1'}));
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
$GLOBALS['Data_NonEmpty_eqNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_eqmajNonmajEmpty';

// Data_NonEmpty_ordNonEmpty
function majData_majNonmajEmpty_ordmajNonmajEmpty($dictOrd1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_ordmajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd1_0)->{'Eq10'})(null);
  $__res = function($dictOrd_2) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd_2)->{'Eq0'})(null);
  $eqNonEmpty2_3_1 = (object)["eq" => function($x_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($y_5) use ($__local_var_1_0, $__local_var_3_1, $x_4) {
  $__num = \func_num_args();
  $__res = (((($__local_var_3_1)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'}) && (((($__local_var_1_0)->{'eq1'})($__local_var_3_1))(($x_4)->{'value1'}))(($y_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($x_4) use ($dictOrd1_0, $dictOrd_2) {
  $__num = \func_num_args();
  $__res = function($y_5) use ($dictOrd1_0, $dictOrd_2, $x_4) {
  $__num = \func_num_args();
  $v_6_3 = ((($dictOrd_2)->{'compare'})(($x_4)->{'value0'}))(($y_5)->{'value0'});
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  $__t4 = (((($dictOrd1_0)->{'compare1'})($dictOrd_2))(($x_4)->{'value1'}))(($y_5)->{'value1'});
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_4) use ($eqNonEmpty2_3_1) {
  $__num = \func_num_args();
  $__res = $eqNonEmpty2_3_1;
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
$GLOBALS['Data_NonEmpty_ordNonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_ordmajNonmajEmpty';

// Data_NonEmpty_eq1NonEmpty
function majData_majNonmajEmpty_eq1majNonmajEmpty($dictEq1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_eq1majNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq1" => function($dictEq_1) use ($dictEq1_0) {
  $__num = \func_num_args();
  $__res = function($x_2) use ($dictEq1_0, $dictEq_1) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictEq1_0, $dictEq_1, $x_2) {
  $__num = \func_num_args();
  $__res = (((($dictEq_1)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && (((($dictEq1_0)->{'eq1'})($dictEq_1))(($x_2)->{'value1'}))(($y_3)->{'value1'}));
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
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_eq1NonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_eq1majNonmajEmpty';

// Data_NonEmpty_ord1NonEmpty
function majData_majNonmajEmpty_ord1majNonmajEmpty($dictOrd1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_ord1majNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd1_0)->{'Eq10'})(null);
  $eq1NonEmpty1_1_0 = (object)["eq1" => function($dictEq_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($x_3) use ($__local_var_1_0, $dictEq_2) {
  $__num = \func_num_args();
  $__res = function($y_4) use ($__local_var_1_0, $dictEq_2, $x_3) {
  $__num = \func_num_args();
  $__res = (((($dictEq_2)->{'eq'})(($x_3)->{'value0'}))(($y_4)->{'value0'}) && (((($__local_var_1_0)->{'eq1'})($dictEq_2))(($x_3)->{'value1'}))(($y_4)->{'value1'}));
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
  $__res = (object)["compare1" => function($dictOrd_2) use ($dictOrd1_0) {
  $__num = \func_num_args();
  $__res = function($x_3) use ($dictOrd1_0, $dictOrd_2) {
  $__num = \func_num_args();
  $__res = function($y_4) use ($dictOrd1_0, $dictOrd_2, $x_3) {
  $__num = \func_num_args();
  $v_5_2 = ((($dictOrd_2)->{'compare'})(($x_3)->{'value0'}))(($y_4)->{'value0'});
  $__t3 = null;;
  if ($v_5_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_3;;
};
  if ($v_5_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_3;;
};
  $__t3 = (((($dictOrd1_0)->{'compare1'})($dictOrd_2))(($x_3)->{'value1'}))(($y_4)->{'value1'});
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
}, "Eq10" => function($_dollar___unused_2) use ($eq1NonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $eq1NonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_ord1NonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_ord1majNonmajEmpty';

