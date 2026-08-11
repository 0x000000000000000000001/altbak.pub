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
  $__local_var_3_0 = ((($GLOBALS['Data_Tuple_functorTuple'])->{'map'})((($dictUnfoldable_0)->{'unfoldr'})((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})($f_1))))(($f_1)($b_2));
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
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(NonEmpty "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow1_1)->{'show'})(($v_2)->{'value1'})))(")"))));
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
function majData_majNonmajEmpty_onemajOf($dictAlternative_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_onemajOf';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'alt'})((((($dictAlternative_0)->{'Applicative0'})(null))->{'pure'})(($v_1)->{'value0'})))(($v_1)->{'value1'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
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
}, "Functor0" => function($_dollar__unused_2) use ($functorNonEmpty1_1_0) {
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
  $foldMap1_2_0 = (($dictFoldable_0)->{'foldMap'})($dictMonoid_1);
  $__res = function($f_3) use ($dictMonoid_1, $foldMap1_2_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictMonoid_1, $f_3, $foldMap1_2_0) {
  $__num = \func_num_args();
  $__res = ((((($dictMonoid_1)->{'Semigroup0'})(null))->{'append'})(($f_3)(($v_4)->{'value0'})))((($foldMap1_2_0)($f_3))(($v_4)->{'value1'}));
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
  $foldMap1_3_1 = (($__local_var_1_0)->{'foldMap'})($dictMonoid_2);
  $__res = function($f_4) use ($dictMonoid_2, $foldMap1_3_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($dictMonoid_2, $f_4, $foldMap1_3_1) {
  $__num = \func_num_args();
  $__res = ((((($dictMonoid_2)->{'Semigroup0'})(null))->{'append'})(($f_4)(($v_5)->{'value0'})))((($foldMap1_3_1)($f_4))(($v_5)->{'value1'}));
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
  $foldMapWithIndex1_3_3 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_2);
  $__res = function($f_4) use ($dictMonoid_2, $foldMapWithIndex1_3_3) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($dictMonoid_2, $f_4, $foldMapWithIndex1_3_3) {
  $__num = \func_num_args();
  $__res = ((((($dictMonoid_2)->{'Semigroup0'})(null))->{'append'})((($f_4)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_5)->{'value0'})))((($foldMapWithIndex1_3_3)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($GLOBALS['Data_Maybe_Just'])))(($v_5)->{'value1'}));
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
}, "Foldable0" => function($_dollar__unused_2) use ($foldableNonEmpty1_1_0) {
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
  $foldMap1_4_3 = (($__local_var_2_2)->{'foldMap'})($dictMonoid_3);
  $__res = function($f_5) use ($dictMonoid_3, $foldMap1_4_3) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($dictMonoid_3, $f_5, $foldMap1_4_3) {
  $__num = \func_num_args();
  $__res = ((((($dictMonoid_3)->{'Semigroup0'})(null))->{'append'})(($f_5)(($v_6)->{'value0'})))((($foldMap1_4_3)($f_5))(($v_6)->{'value1'}));
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
  $sequence1_5_6 = (($dictTraversable_0)->{'sequence'})($dictApplicative_3);
  $__res = function($v_6) use ($Apply0_4_5, $sequence1_5_6) {
  $__num = \func_num_args();
  $__res = ((($Apply0_4_5)->{'apply'})(((((($Apply0_4_5)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($v_6)->{'value0'})))(($sequence1_5_6)(($v_6)->{'value1'}));
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
  $traverse1_5_8 = (($dictTraversable_0)->{'traverse'})($dictApplicative_3);
  $__res = function($f_6) use ($Apply0_4_7, $traverse1_5_8) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Apply0_4_7, $f_6, $traverse1_5_8) {
  $__num = \func_num_args();
  $__res = ((($Apply0_4_7)->{'apply'})(((((($Apply0_4_7)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($f_6)(($v_7)->{'value0'}))))((($traverse1_5_8)($f_6))(($v_7)->{'value1'}));
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
}, "Functor0" => function($_dollar__unused_3) use ($functorNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_3) use ($foldableNonEmpty1_2_2) {
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
  $functorWithIndex1_1_0 = ($GLOBALS['Data_NonEmpty_functorWithIndex'])((($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null));
  $foldableWithIndexNonEmpty1_2_1 = ($GLOBALS['Data_NonEmpty_foldableWithIndexNonEmpty'])((($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null));
  $traversableNonEmpty1_3_2 = ($GLOBALS['Data_NonEmpty_traversableNonEmpty'])((($dictTraversableWithIndex_0)->{'Traversable2'})(null));
  $__res = (object)["traverseWithIndex" => function($dictApplicative_4) use ($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $Apply0_5_3 = (($dictApplicative_4)->{'Apply0'})(null);
  $traverseWithIndex1_6_4 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_4);
  $__res = function($f_7) use ($Apply0_5_3, $traverseWithIndex1_6_4) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($Apply0_5_3, $f_7, $traverseWithIndex1_6_4) {
  $__num = \func_num_args();
  $__res = ((($Apply0_5_3)->{'apply'})(((((($Apply0_5_3)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))((($f_7)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_8)->{'value0'}))))((($traverseWithIndex1_6_4)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_7))($GLOBALS['Data_Maybe_Just'])))(($v_8)->{'value1'}));
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
}, "FunctorWithIndex0" => function($_dollar__unused_4) use ($functorWithIndex1_1_0) {
  $__num = \func_num_args();
  $__res = $functorWithIndex1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_4) use ($foldableWithIndexNonEmpty1_2_1) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexNonEmpty1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_4) use ($traversableNonEmpty1_3_2) {
  $__num = \func_num_args();
  $__res = $traversableNonEmpty1_3_2;
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
  $foldMap1_2_0 = (($dictFoldable_0)->{'foldMap'})($dictMonoid_1);
  $__res = function($f_3) use ($dictMonoid_1, $foldMap1_2_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictMonoid_1, $f_3, $foldMap1_2_0) {
  $__num = \func_num_args();
  $__res = ((((($dictMonoid_1)->{'Semigroup0'})(null))->{'append'})(($f_3)(($v_4)->{'value0'})))((($foldMap1_2_0)($f_3))(($v_4)->{'value1'}));
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
}, "Foldable0" => function($_dollar__unused_2) use ($foldableNonEmpty1_1_0) {
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
function majData_majNonmajEmpty_foldl1($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_foldl1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_NonEmpty_foldable1NonEmpty'])($dictFoldable_0))->{'foldl1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_foldl1'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_foldl1';

// Data_NonEmpty_eqNonEmpty
function majData_majNonmajEmpty_eqmajNonmajEmpty($dictEq1_0, $dictEq_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNonmajEmpty_eqmajNonmajEmpty';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $eq11_2_0 = (($dictEq1_0)->{'eq1'})($dictEq_1);
  $__res = (object)["eq" => function($x_3) use ($dictEq_1, $eq11_2_0) {
  $__num = \func_num_args();
  $__res = function($y_4) use ($dictEq_1, $eq11_2_0, $x_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($dictEq_1)->{'eq'})(($x_3)->{'value0'}))(($y_4)->{'value0'})))((($eq11_2_0)(($x_3)->{'value1'}))(($y_4)->{'value1'}));
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
  $eqNonEmpty1_1_0 = ($GLOBALS['Data_NonEmpty_eqNonEmpty'])((($dictOrd1_0)->{'Eq10'})(null));
  $__res = function($dictOrd_2) use ($dictOrd1_0, $eqNonEmpty1_1_0) {
  $__num = \func_num_args();
  $compare11_3_1 = (($dictOrd1_0)->{'compare1'})($dictOrd_2);
  $eqNonEmpty2_4_2 = ($eqNonEmpty1_1_0)((($dictOrd_2)->{'Eq0'})(null));
  $__res = (object)["compare" => function($x_5) use ($compare11_3_1, $dictOrd_2) {
  $__num = \func_num_args();
  $__res = function($y_6) use ($compare11_3_1, $dictOrd_2, $x_5) {
  $__num = \func_num_args();
  $v_7_3 = ((($dictOrd_2)->{'compare'})(($x_5)->{'value0'}))(($y_6)->{'value0'});
  $__t4 = null;;
  if ($v_7_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_4;;
};
  if ($v_7_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  $__t4 = (($compare11_3_1)(($x_5)->{'value1'}))(($y_6)->{'value1'});
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar__unused_5) use ($eqNonEmpty2_4_2) {
  $__num = \func_num_args();
  $__res = $eqNonEmpty2_4_2;
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
  $__res = ((($GLOBALS['Data_NonEmpty_eqNonEmpty'])($dictEq1_0))($dictEq_1))->{'eq'};
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
  $ordNonEmpty1_1_0 = ($GLOBALS['Data_NonEmpty_ordNonEmpty'])($dictOrd1_0);
  $__local_var_2_1 = (($dictOrd1_0)->{'Eq10'})(null);
  $eq1NonEmpty1_2_1 = (object)["eq1" => function($dictEq_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_NonEmpty_eqNonEmpty'])($__local_var_2_1))($dictEq_3))->{'eq'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare1" => function($dictOrd_3) use ($ordNonEmpty1_1_0) {
  $__num = \func_num_args();
  $__res = (($ordNonEmpty1_1_0)($dictOrd_3))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_3) use ($eq1NonEmpty1_2_1) {
  $__num = \func_num_args();
  $__res = $eq1NonEmpty1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_NonEmpty_ord1NonEmpty'] = __NAMESPACE__ . '\\majData_majNonmajEmpty_ord1majNonmajEmpty';

