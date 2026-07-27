<?php

namespace Data\Bitraversable;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Category, Data.Bifoldable, Data.Bifunctor, Data.Bitraversable, Data.Const, Data.Either, Data.Functor, Data.Functor.Clown, Data.Functor.Flip, Data.Functor.Joker, Data.Functor.Product2, Data.Traversable, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Category, Data.Bifoldable, Data.Bifunctor, Data.Bitraversable, Data.Const, Data.Either, Data.Functor, Data.Functor.Clown, Data.Functor.Flip, Data.Functor.Joker, Data.Functor.Product2, Data.Traversable, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Data.Bifoldable/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Bitraversable/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Clown/index.php';
require_once __DIR__ . '/../Data.Functor.Flip/index.php';
require_once __DIR__ . '/../Data.Functor.Joker/index.php';
require_once __DIR__ . '/../Data.Functor.Product2/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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




// Data_Bitraversable_identity
function majData_majBitraversable_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bitraversable_identity'] = __NAMESPACE__ . '\\majData_majBitraversable_identity';

// Data_Bitraversable_bitraverse
function majData_majBitraversable_bitraverse($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bitraverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'bitraverse'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bitraversable_bitraverse'] = __NAMESPACE__ . '\\majData_majBitraversable_bitraverse';

// Data_Bitraversable_lfor
function majData_majBitraversable_lfor($dictBitraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_lfor';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $bitraverse2_2_0 = (($dictBitraversable_0)->{'bitraverse'})($dictApplicative_1);
  $pure_3_1 = ($dictApplicative_1)->{'pure'};
  $__res = (function() use ($bitraverse2_2_0, $pure_3_1) {
  $__fn = function($t_4, $f_5 = null) use ($bitraverse2_2_0, $pure_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($bitraverse2_2_0)($f_5))($pure_3_1))($t_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bitraversable_lfor'] = __NAMESPACE__ . '\\majData_majBitraversable_lfor';

// Data_Bitraversable_ltraverse
function majData_majBitraversable_ltraverse($dictBitraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_ltraverse';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $bitraverse2_2_0 = (($dictBitraversable_0)->{'bitraverse'})($dictApplicative_1);
  $pure_3_1 = ($dictApplicative_1)->{'pure'};
  $__res = function($f_4) use ($bitraverse2_2_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = (($bitraverse2_2_0)($f_4))($pure_3_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bitraversable_ltraverse'] = __NAMESPACE__ . '\\majData_majBitraversable_ltraverse';

// Data_Bitraversable_rfor
function majData_majBitraversable_rfor($dictBitraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_rfor';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $bitraverse2_2_0 = (($dictBitraversable_0)->{'bitraverse'})($dictApplicative_1);
  $pure_3_1 = ($dictApplicative_1)->{'pure'};
  $__res = (function() use ($bitraverse2_2_0, $pure_3_1) {
  $__fn = function($t_4, $f_5 = null) use ($bitraverse2_2_0, $pure_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($bitraverse2_2_0)($pure_3_1))($f_5))($t_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bitraversable_rfor'] = __NAMESPACE__ . '\\majData_majBitraversable_rfor';

// Data_Bitraversable_rtraverse
function majData_majBitraversable_rtraverse($dictBitraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_rtraverse';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictBitraversable_0)->{'bitraverse'})($dictApplicative_1))(($dictApplicative_1)->{'pure'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bitraversable_rtraverse'] = __NAMESPACE__ . '\\majData_majBitraversable_rtraverse';

// Data_Bitraversable_bitraversableTuple
$GLOBALS['Data_Bitraversable_bitraversableTuple'] = (object)["bitraverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = (function() use ($Apply0_1_0) {
  $__fn = function($f_2, $g_3 = null, $v_4 = null) use ($Apply0_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Apply0_1_0)->{'apply'})(((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Tuple_Tuple']))(($f_2)(($v_4)->{'value0'}))))(($g_3)(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_1 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($v_2) use ($Apply0_1_1) {
  $__num = \func_num_args();
  $__res = ((($Apply0_1_1)->{'apply'})(((((($Apply0_1_1)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Tuple_Tuple']))(($v_2)->{'value0'})))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bifunctor_bifunctorTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bifoldable_bifoldableTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bitraversable_bitraversableJoker
function majData_majBitraversable_bitraversablemajJoker($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bitraversablemajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $bifunctorJoker_2_1 = (object)["bimap" => (function() use ($__local_var_1_0) {
  $__fn = function($v_2, $g_3 = null, $v1_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)->{'map'})($g_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__local_var_3_2 = (($dictTraversable_0)->{'Foldable1'})(null);
  $bifoldableJoker_4_3 = (object)["bifoldr" => (function() use ($__local_var_3_2) {
  $__fn = function($v_4, $r_5 = null, $u_6 = null, $v1_7 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($__local_var_3_2)->{'foldr'})($r_5))($u_6))($v1_7);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() use ($__local_var_3_2) {
  $__fn = function($v_4, $r_5 = null, $u_6 = null, $v1_7 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($__local_var_3_2)->{'foldl'})($r_5))($u_6))($v1_7);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $foldMap1_5_3 = (($__local_var_3_2)->{'foldMap'})($dictMonoid_4);
  $__res = (function() use ($foldMap1_5_3) {
  $__fn = function($v_6, $r_7 = null, $v1_8 = null) use ($foldMap1_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($foldMap1_5_3)($r_7))($v1_8);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bitraverse" => function($dictApplicative_5) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $traverse1_6_5 = (($dictTraversable_0)->{'traverse'})($dictApplicative_5);
  $__res = (function() use ($dictApplicative_5, $traverse1_6_5) {
  $__fn = function($v_7, $r_8 = null, $v1_9 = null) use ($dictApplicative_5, $traverse1_6_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Joker_Joker']))((($traverse1_6_5)($r_8))($v1_9));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_5) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $sequence1_6_6 = (($dictTraversable_0)->{'sequence'})($dictApplicative_5);
  $__res = function($v_7) use ($dictApplicative_5, $sequence1_6_6) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Joker_Joker']))(($sequence1_6_6)($v_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar__unused_5) use ($bifunctorJoker_2_1) {
  $__num = \func_num_args();
  $__res = $bifunctorJoker_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_5) use ($bifoldableJoker_4_3) {
  $__num = \func_num_args();
  $__res = $bifoldableJoker_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bitraversable_bitraversableJoker'] = __NAMESPACE__ . '\\majData_majBitraversable_bitraversablemajJoker';

// Data_Bitraversable_bitraversableEither
$GLOBALS['Data_Bitraversable_bitraversableEither'] = (object)["bitraverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = (function() use ($__local_var_1_0) {
  $__fn = function($v_2, $v1_3 = null, $v2_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if ($v2_4 instanceof \Data\Either\Data_Either_Left) {
$__t1 = ((($__local_var_1_0)->{'map'})($GLOBALS['Data_Either_Left']))(($v_2)(($v2_4)->{'value0'}));
goto end_branch_1;;
};
  if ($v2_4 instanceof \Data\Either\Data_Either_Right) {
$__t1 = ((($__local_var_1_0)->{'map'})($GLOBALS['Data_Either_Right']))(($v1_3)(($v2_4)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__local_var_1_2 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($__local_var_1_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v_2 instanceof \Data\Either\Data_Either_Left) {
$__t3 = ((($__local_var_1_2)->{'map'})($GLOBALS['Data_Either_Left']))(($v_2)->{'value0'});
goto end_branch_3;;
};
  if ($v_2 instanceof \Data\Either\Data_Either_Right) {
$__t3 = ((($__local_var_1_2)->{'map'})($GLOBALS['Data_Either_Right']))(($v_2)->{'value0'});
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
}, "Bifunctor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bifunctor_bifunctorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bifoldable_bifoldableEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bitraversable_bitraversableConst
$GLOBALS['Data_Bitraversable_bitraversableConst'] = (object)["bitraverse" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null, $v1_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Const_Const']))(($f_1)($v1_3));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bisequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Const_Const']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Bifunctor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bifunctor_bifunctorConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Bifoldable_bifoldableConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Bitraversable_bitraversableClown
function majData_majBitraversable_bitraversablemajClown($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bitraversablemajClown';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $bifunctorClown_2_1 = (object)["bimap" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2, $v_3 = null, $v1_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)->{'map'})($f_2))($v1_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__local_var_3_2 = (($dictTraversable_0)->{'Foldable1'})(null);
  $bifoldableClown_4_3 = (object)["bifoldr" => (function() use ($__local_var_3_2) {
  $__fn = function($l_4, $v_5 = null, $u_6 = null, $v1_7 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($__local_var_3_2)->{'foldr'})($l_4))($u_6))($v1_7);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() use ($__local_var_3_2) {
  $__fn = function($l_4, $v_5 = null, $u_6 = null, $v1_7 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($__local_var_3_2)->{'foldl'})($l_4))($u_6))($v1_7);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $foldMap1_5_3 = (($__local_var_3_2)->{'foldMap'})($dictMonoid_4);
  $__res = (function() use ($foldMap1_5_3) {
  $__fn = function($l_6, $v_7 = null, $v1_8 = null) use ($foldMap1_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($foldMap1_5_3)($l_6))($v1_8);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bitraverse" => function($dictApplicative_5) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $traverse1_6_5 = (($dictTraversable_0)->{'traverse'})($dictApplicative_5);
  $__res = (function() use ($dictApplicative_5, $traverse1_6_5) {
  $__fn = function($l_7, $v_8 = null, $v1_9 = null) use ($dictApplicative_5, $traverse1_6_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Clown_Clown']))((($traverse1_6_5)($l_7))($v1_9));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_5) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $sequence1_6_6 = (($dictTraversable_0)->{'sequence'})($dictApplicative_5);
  $__res = function($v_7) use ($dictApplicative_5, $sequence1_6_6) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Clown_Clown']))(($sequence1_6_6)($v_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar__unused_5) use ($bifunctorClown_2_1) {
  $__num = \func_num_args();
  $__res = $bifunctorClown_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_5) use ($bifoldableClown_4_3) {
  $__num = \func_num_args();
  $__res = $bifoldableClown_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bitraversable_bitraversableClown'] = __NAMESPACE__ . '\\majData_majBitraversable_bitraversablemajClown';

// Data_Bitraversable_bisequenceDefault
function majData_majBitraversable_bisequencemajDefault($dictBitraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bisequencemajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictBitraversable_0)->{'bitraverse'})($dictApplicative_1))($GLOBALS['Data_Bitraversable_identity']))($GLOBALS['Data_Bitraversable_identity']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bitraversable_bisequenceDefault'] = __NAMESPACE__ . '\\majData_majBitraversable_bisequencemajDefault';

// Data_Bitraversable_bisequence
function majData_majBitraversable_bisequence($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bisequence';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'bisequence'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bitraversable_bisequence'] = __NAMESPACE__ . '\\majData_majBitraversable_bisequence';

// Data_Bitraversable_bitraversableFlip
function majData_majBitraversable_bitraversablemajFlip($dictBitraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bitraversablemajFlip';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBitraversable_0)->{'Bifunctor0'})(null);
  $bifunctorFlip_2_1 = (object)["bimap" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2, $g_3 = null, $v_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_1_0)->{'bimap'})($g_3))($f_2))($v_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__local_var_3_2 = (($dictBitraversable_0)->{'Bifoldable1'})(null);
  $bifoldableFlip_4_3 = (object)["bifoldr" => (function() use ($__local_var_3_2) {
  $__fn = function($r_4, $l_5 = null, $u_6 = null, $v_7 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($__local_var_3_2)->{'bifoldr'})($l_5))($r_4))($u_6))($v_7);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() use ($__local_var_3_2) {
  $__fn = function($r_4, $l_5 = null, $u_6 = null, $v_7 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($__local_var_3_2)->{'bifoldl'})($l_5))($r_4))($u_6))($v_7);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $bifoldMap2_5_3 = (($__local_var_3_2)->{'bifoldMap'})($dictMonoid_4);
  $__res = (function() use ($bifoldMap2_5_3) {
  $__fn = function($r_6, $l_7 = null, $v_8 = null) use ($bifoldMap2_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($bifoldMap2_5_3)($l_7))($r_6))($v_8);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bitraverse" => function($dictApplicative_5) use ($dictBitraversable_0) {
  $__num = \func_num_args();
  $bitraverse2_6_5 = (($dictBitraversable_0)->{'bitraverse'})($dictApplicative_5);
  $__res = (function() use ($bitraverse2_6_5, $dictApplicative_5) {
  $__fn = function($r_7, $l_8 = null, $v_9 = null) use ($bitraverse2_6_5, $dictApplicative_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Flip_Flip']))(((($bitraverse2_6_5)($l_8))($r_7))($v_9));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_5) use ($dictBitraversable_0) {
  $__num = \func_num_args();
  $bisequence2_6_6 = (($dictBitraversable_0)->{'bisequence'})($dictApplicative_5);
  $__res = function($v_7) use ($bisequence2_6_6, $dictApplicative_5) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Flip_Flip']))(($bisequence2_6_6)($v_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar__unused_5) use ($bifunctorFlip_2_1) {
  $__num = \func_num_args();
  $__res = $bifunctorFlip_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_5) use ($bifoldableFlip_4_3) {
  $__num = \func_num_args();
  $__res = $bifoldableFlip_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bitraversable_bitraversableFlip'] = __NAMESPACE__ . '\\majData_majBitraversable_bitraversablemajFlip';

// Data_Bitraversable_bitraversableProduct2
function majData_majBitraversable_bitraversablemajProduct2($dictBitraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bitraversablemajProduct2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBitraversable_0)->{'Bifunctor0'})(null);
  $bifoldableProduct2_2_1 = ($GLOBALS['Data_Bifoldable_bifoldableProduct2'])((($dictBitraversable_0)->{'Bifoldable1'})(null));
  $__res = function($dictBitraversable1_3) use ($__local_var_1_0, $bifoldableProduct2_2_1, $dictBitraversable_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictBitraversable1_3)->{'Bifunctor0'})(null);
  $bifunctorProduct21_5_3 = (object)["bimap" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5, $g_6 = null, $v_7 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Functor\Product2\Data_Functor_Product2_Product2((((($__local_var_1_0)->{'bimap'})($f_5))($g_6))(($v_7)->{'value0'}), (((($__local_var_4_2)->{'bimap'})($f_5))($g_6))(($v_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $bifoldableProduct21_6_4 = ($bifoldableProduct2_2_1)((($dictBitraversable1_3)->{'Bifoldable1'})(null));
  $__res = (object)["bitraverse" => function($dictApplicative_7) use ($dictBitraversable1_3, $dictBitraversable_0) {
  $__num = \func_num_args();
  $Apply0_8_5 = (($dictApplicative_7)->{'Apply0'})(null);
  $bitraverse3_9_6 = (($dictBitraversable_0)->{'bitraverse'})($dictApplicative_7);
  $bitraverse4_10_7 = (($dictBitraversable1_3)->{'bitraverse'})($dictApplicative_7);
  $__res = (function() use ($Apply0_8_5, $bitraverse3_9_6, $bitraverse4_10_7) {
  $__fn = function($l_11, $r_12 = null, $v_13 = null) use ($Apply0_8_5, $bitraverse3_9_6, $bitraverse4_10_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Apply0_8_5)->{'apply'})(((((($Apply0_8_5)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product2_Product2']))(((($bitraverse3_9_6)($l_11))($r_12))(($v_13)->{'value0'}))))(((($bitraverse4_10_7)($l_11))($r_12))(($v_13)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_7) use ($dictBitraversable1_3, $dictBitraversable_0) {
  $__num = \func_num_args();
  $Apply0_8_8 = (($dictApplicative_7)->{'Apply0'})(null);
  $bisequence3_9_9 = (($dictBitraversable_0)->{'bisequence'})($dictApplicative_7);
  $bisequence4_10_10 = (($dictBitraversable1_3)->{'bisequence'})($dictApplicative_7);
  $__res = function($v_11) use ($Apply0_8_8, $bisequence3_9_9, $bisequence4_10_10) {
  $__num = \func_num_args();
  $__res = ((($Apply0_8_8)->{'apply'})(((((($Apply0_8_8)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product2_Product2']))(($bisequence3_9_9)(($v_11)->{'value0'}))))(($bisequence4_10_10)(($v_11)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar__unused_7) use ($bifunctorProduct21_5_3) {
  $__num = \func_num_args();
  $__res = $bifunctorProduct21_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_7) use ($bifoldableProduct21_6_4) {
  $__num = \func_num_args();
  $__res = $bifoldableProduct21_6_4;
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
$GLOBALS['Data_Bitraversable_bitraversableProduct2'] = __NAMESPACE__ . '\\majData_majBitraversable_bitraversablemajProduct2';

// Data_Bitraversable_bitraverseDefault
function majData_majBitraversable_bitraversemajDefault($dictBitraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bitraversemajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $bisequence2_2_0 = (($dictBitraversable_0)->{'bisequence'})($dictApplicative_1);
  $__res = (function() use ($bisequence2_2_0, $dictBitraversable_0) {
  $__fn = function($f_3, $g_4 = null, $t_5 = null) use ($bisequence2_2_0, $dictBitraversable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($bisequence2_2_0)((((((($dictBitraversable_0)->{'Bifunctor0'})(null))->{'bimap'})($f_3))($g_4))($t_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bitraversable_bitraverseDefault'] = __NAMESPACE__ . '\\majData_majBitraversable_bitraversemajDefault';

// Data_Bitraversable_bifor
function majData_majBitraversable_bifor($dictBitraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBitraversable_bifor';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $bitraverse2_2_0 = (($dictBitraversable_0)->{'bitraverse'})($dictApplicative_1);
  $__res = (function() use ($bitraverse2_2_0) {
  $__fn = function($t_3, $f_4 = null, $g_5 = null) use ($bitraverse2_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($bitraverse2_2_0)($f_4))($g_5))($t_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bitraversable_bifor'] = __NAMESPACE__ . '\\majData_majBitraversable_bifor';

