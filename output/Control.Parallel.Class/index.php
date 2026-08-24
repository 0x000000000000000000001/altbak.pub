<?php

namespace Control\Parallel\Class;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad.Cont.Trans, Control.Monad.Except.Trans, Control.Monad.Maybe.Trans, Control.Monad.Reader.Trans, Control.Monad.Writer.Trans, Control.Parallel.Class, Control.Plus, Control.Semigroupoid, Data.Either, Data.Function, Data.Functor, Data.Functor.Compose, Data.Functor.Costar, Data.Maybe, Data.Newtype, Data.Profunctor.Star, Data.Unit, Effect.Class, Effect.Ref, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad.Cont.Trans, Control.Monad.Except.Trans, Control.Monad.Maybe.Trans, Control.Monad.Reader.Trans, Control.Monad.Writer.Trans, Control.Parallel.Class, Control.Plus, Control.Semigroupoid, Data.Either, Data.Function, Data.Functor, Data.Functor.Compose, Data.Functor.Costar, Data.Maybe, Data.Newtype, Data.Profunctor.Star, Data.Unit, Effect.Class, Effect.Ref, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Except.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Maybe.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Trans/index.php';
require_once __DIR__ . '/../Control.Parallel.Class/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Functor.Costar/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Profunctor.Star/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect.Class/index.php';
require_once __DIR__ . '/../Effect.Ref/index.php';
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




// Control_Parallel_Class_ParCont
function majControl_majParallel_majClass_majParmajCont($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_majParmajCont';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_ParCont'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_majParmajCont';

// Control_Parallel_Class_sequential
function majControl_majParallel_majClass_sequential($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_sequential';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'sequential'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_sequential'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_sequential';

// Control_Parallel_Class_parallel
function majControl_majParallel_majClass_parallel($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_parallel';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'parallel'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_parallel'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_parallel';

// Control_Parallel_Class_newtypeParCont
$GLOBALS['Control_Parallel_Class_newtypeParCont'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Parallel_Class_monadParWriterT
function majControl_majParallel_majClass_monadmajParmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_monadmajParmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $applyWriterT_2_1 = function($dictApply_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $Functor0_3_1 = (($dictApply_2)->{'Functor0'})(null);
  $__local_var_4_2 = (($dictApply_2)->{'Functor0'})(null);
  $functorWriterT1_4_2 = (object)["map" => function($f_5) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__local_var_6_3 = (($__local_var_4_2)->{'map'})(function($v_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_5)(($v_6)->{'value0'}), ($v_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_7) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__res = ($__local_var_6_3)($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($v_5) use ($Functor0_3_1, $__local_var_1_0, $dictApply_2) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($Functor0_3_1, $__local_var_1_0, $dictApply_2, $v_5) {
  $__num = \func_num_args();
  $__res = ((($dictApply_2)->{'apply'})(((($Functor0_3_1)->{'map'})(function($v3_7) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v4_8) use ($__local_var_1_0, $v3_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_7)->{'value0'})(($v4_8)->{'value0'}), ((($__local_var_1_0)->{'append'})(($v3_7)->{'value1'}))(($v4_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5)))($v1_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorWriterT1_4_2) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictParallel_3) use ($applyWriterT_2_1) {
  $__num = \func_num_args();
  $applyWriterT1_4_6 = ($applyWriterT_2_1)((($dictParallel_3)->{'Apply0'})(null));
  $applyWriterT2_5_7 = ($applyWriterT_2_1)((($dictParallel_3)->{'Apply1'})(null));
  $__res = (object)["parallel" => function($v_6) use ($dictParallel_3) {
  $__num = \func_num_args();
  $__res = (($dictParallel_3)->{'parallel'})($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_6) use ($dictParallel_3) {
  $__num = \func_num_args();
  $__res = (($dictParallel_3)->{'sequential'})($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($applyWriterT1_4_6) {
  $__num = \func_num_args();
  $__res = $applyWriterT1_4_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar___unused_6) use ($applyWriterT2_5_7) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_5_7;
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
$GLOBALS['Control_Parallel_Class_monadParWriterT'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_monadmajParmajWritermajT';

// Control_Parallel_Class_monadParStar
function majControl_majParallel_majClass_monadmajParmajStar($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_monadmajParmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $parallel1_1_0 = ($dictParallel_0)->{'parallel'};
  $sequential1_2_1 = ($dictParallel_0)->{'sequential'};
  $__local_var_3_2 = (($dictParallel_0)->{'Apply0'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Functor0'})(null);
  $functorStar1_4_3 = (object)["map" => function($f_5) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_3, $f_5) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_4_3)->{'map'})($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar_3_2 = (object)["apply" => function($v_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_2, $v_5) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($__local_var_3_2, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'apply'})(($v_5)($a_7)))(($v1_6)($a_7));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorStar1_4_3) {
  $__num = \func_num_args();
  $__res = $functorStar1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_6 = (($dictParallel_0)->{'Apply1'})(null);
  $__local_var_5_7 = (($__local_var_4_6)->{'Functor0'})(null);
  $functorStar1_5_7 = (object)["map" => function($f_6) use ($__local_var_5_7) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_7, $f_6) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_5_7)->{'map'})($f_6)))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar1_4_6 = (object)["apply" => function($v_6) use ($__local_var_4_6) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_6, $v_6) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($__local_var_4_6, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_6)->{'apply'})(($v_6)($a_8)))(($v1_7)($a_8));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStar1_5_7) {
  $__num = \func_num_args();
  $__res = $functorStar1_5_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["parallel" => function($v_5) use ($parallel1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($parallel1_1_0))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_5) use ($sequential1_2_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($sequential1_2_1))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($applyStar_3_2) {
  $__num = \func_num_args();
  $__res = $applyStar_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar___unused_5) use ($applyStar1_4_6) {
  $__num = \func_num_args();
  $__res = $applyStar1_4_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_monadParStar'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_monadmajParmajStar';

// Control_Parallel_Class_monadParReaderT
function majControl_majParallel_majClass_monadmajParmajReadermajT($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_monadmajParmajReadermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictParallel_0)->{'Apply0'})(null);
  $functorReaderT1_2_1 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_1_0)->{'Functor0'})(null))->{'map'})];
  $applyReaderT_1_0 = (object)["apply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($r_5) use ($__local_var_1_0, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(($v_3)($r_5)))(($v1_4)($r_5));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_3 = (($dictParallel_0)->{'Apply1'})(null);
  $functorReaderT1_3_4 = (object)["map" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Reader_Trans_mapReaderT']))(((($__local_var_2_3)->{'Functor0'})(null))->{'map'})];
  $applyReaderT1_2_3 = (object)["apply" => function($v_4) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_3, $v_4) {
  $__num = \func_num_args();
  $__res = function($r_6) use ($__local_var_2_3, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_3)->{'apply'})(($v_4)($r_6)))(($v1_5)($r_6));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorReaderT1_3_4) {
  $__num = \func_num_args();
  $__res = $functorReaderT1_3_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["parallel" => ($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'])(($dictParallel_0)->{'parallel'}), "sequential" => ($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'])(($dictParallel_0)->{'sequential'}), "Apply0" => function($_dollar___unused_3) use ($applyReaderT_1_0) {
  $__num = \func_num_args();
  $__res = $applyReaderT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar___unused_3) use ($applyReaderT1_2_3) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_2_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_monadParReaderT'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_monadmajParmajReadermajT';

// Control_Parallel_Class_monadParMaybeT
function majControl_majParallel_majClass_monadmajParmajMaybemajT($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_monadmajParmajMaybemajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictParallel_0)->{'Apply1'})(null);
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorCompose2_4_3 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})(function($v1_6) use ($f_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(($f_4)(($v1_6)->{'value0'}));
goto end_branch_3;;
};
  $__t3 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyCompose_1_0 = (object)["apply" => function($v_5) use ($Functor0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($Functor0_2_1, $__local_var_1_0, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})(function($v_7) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($v_7) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v_7 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = null;;
if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = new \Data\Maybe\Data_Maybe_Just((($v_7)->{'value0'})(($v1_8)->{'value0'}));
goto end_branch_6;;
};
$__t6 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5)))($v1_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorCompose2_4_3) {
  $__num = \func_num_args();
  $__res = $functorCompose2_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonad_2) use ($applyCompose_1_0, $dictParallel_0) {
  $__num = \func_num_args();
  $__local_var_3_8 = (((((($dictMonad_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_3_8 = (object)["map" => function($f_4) use ($__local_var_3_8) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_8, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_8)->{'map'})(function($v1_6) use ($f_4) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v1_6 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = new \Data\Maybe\Data_Maybe_Just(($f_4)(($v1_6)->{'value0'}));
goto end_branch_9;;
};
  $__t9 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_4_11 = (($dictMonad_2)->{'Bind1'})(null);
  $Applicative0_5_12 = (($dictMonad_2)->{'Applicative0'})(null);
  $Bind1_4_11 = (object)["bind" => function($v_6) use ($Applicative0_5_12, $Bind1_4_11) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($Applicative0_5_12, $Bind1_4_11, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_11)->{'bind'})($v_6))(function($v1_8) use ($Applicative0_5_12, $f_7) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t13 = (($Applicative0_5_12)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_13;;
};
  if ($v1_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t13 = ($f_7)(($v1_8)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
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
}, "Apply0" => function($_dollar___unused_6) use ($dictMonad_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_5_15 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Maybe_Just'])), "Apply0" => function($_dollar___unused_5) use ($dictMonad_2) {
  $__num = \func_num_args();
  $__local_var_6_15 = (((((($dictMonad_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorMaybeT1_6_15 = (object)["map" => function($f_7) use ($__local_var_6_15) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_15, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_15)->{'map'})(function($v1_9) use ($f_7) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v1_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t16 = new \Data\Maybe\Data_Maybe_Just(($f_7)(($v1_9)->{'value0'}));
goto end_branch_16;;
};
  $__t16 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_18 = (($dictMonad_2)->{'Bind1'})(null);
  $Applicative0_8_19 = (($dictMonad_2)->{'Applicative0'})(null);
  $Bind1_7_18 = (object)["bind" => function($v_9) use ($Applicative0_8_19, $Bind1_7_18) {
  $__num = \func_num_args();
  $__res = function($f_10) use ($Applicative0_8_19, $Bind1_7_18, $v_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_18)->{'bind'})($v_9))(function($v1_11) use ($Applicative0_8_19, $f_10) {
  $__num = \func_num_args();
  $__t20 = null;;
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t20 = (($Applicative0_8_19)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_20;;
};
  if ($v1_11 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t20 = ($f_10)(($v1_11)->{'value0'});
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = $__t20;
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
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_22 = ($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'])($dictMonad_2);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_22, $Bind1_7_18) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_22, $Bind1_7_18, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_18)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_22, $Bind1_7_18, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_18)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_22, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_22)->{'pure'})(($f_prime__11)($a_prime__12));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorMaybeT1_6_15) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_6_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyMaybeT_3_8 = (object)["apply" => function($f_6) use ($Applicative0_5_15, $Bind1_4_11) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Applicative0_5_15, $Bind1_4_11, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_11)->{'bind'})($f_6))(function($f_prime__8) use ($Applicative0_5_15, $Bind1_4_11, $a_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_11)->{'bind'})($a_7))(function($a_prime__9) use ($Applicative0_5_15, $f_prime__8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_15)->{'pure'})(($f_prime__8)($a_prime__9));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorMaybeT1_3_8) {
  $__num = \func_num_args();
  $__res = $functorMaybeT1_3_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["parallel" => function($v_4) use ($dictParallel_0) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)->{'parallel'})($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_4) use ($dictParallel_0) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)->{'sequential'})($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($applyMaybeT_3_8) {
  $__num = \func_num_args();
  $__res = $applyMaybeT_3_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar___unused_4) use ($applyCompose_1_0) {
  $__num = \func_num_args();
  $__res = $applyCompose_1_0;
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
$GLOBALS['Control_Parallel_Class_monadParMaybeT'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_monadmajParmajMaybemajT';

// Control_Parallel_Class_monadParExceptT
function majControl_majParallel_majClass_monadmajParmajExceptmajT($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_monadmajParmajExceptmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictParallel_0)->{'Apply1'})(null);
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorCompose2_4_3 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})(function($m_6) use ($f_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($m_6 instanceof \Data\Either\Data_Either_Left) {
$__t3 = new \Data\Either\Data_Either_Left(($m_6)->{'value0'});
goto end_branch_3;;
};
  if ($m_6 instanceof \Data\Either\Data_Either_Right) {
$__t3 = new \Data\Either\Data_Either_Right(($f_4)(($m_6)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyCompose_1_0 = (object)["apply" => function($v_5) use ($Functor0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($Functor0_2_1, $__local_var_1_0, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})(function($v_7) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($v_7) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v_7 instanceof \Data\Either\Data_Either_Left) {
$__t5 = new \Data\Either\Data_Either_Left(($v_7)->{'value0'});
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\Either\Data_Either_Right) {
$__t6 = null;;
if ($v1_8 instanceof \Data\Either\Data_Either_Left) {
$__t6 = new \Data\Either\Data_Either_Left(($v1_8)->{'value0'});
goto end_branch_6;;
};
if ($v1_8 instanceof \Data\Either\Data_Either_Right) {
$__t6 = new \Data\Either\Data_Either_Right((($v_7)->{'value0'})(($v1_8)->{'value0'}));
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5)))($v1_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorCompose2_4_3) {
  $__num = \func_num_args();
  $__res = $functorCompose2_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonad_2) use ($applyCompose_1_0, $dictParallel_0) {
  $__num = \func_num_args();
  $__local_var_3_8 = (((((($dictMonad_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_3_8 = (object)["map" => function($f_4) use ($__local_var_3_8) {
  $__num = \func_num_args();
  $__local_var_5_9 = (($__local_var_3_8)->{'map'})(function($m_5) use ($f_4) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($m_5 instanceof \Data\Either\Data_Either_Left) {
$__t9 = new \Data\Either\Data_Either_Left(($m_5)->{'value0'});
goto end_branch_9;;
};
  if ($m_5 instanceof \Data\Either\Data_Either_Right) {
$__t9 = new \Data\Either\Data_Either_Right(($f_4)(($m_5)->{'value0'}));
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_6) use ($__local_var_5_9) {
  $__num = \func_num_args();
  $__res = ($__local_var_5_9)($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_4_12 = (($dictMonad_2)->{'Bind1'})(null);
  $pure_5_13 = ((($dictMonad_2)->{'Applicative0'})(null))->{'pure'};
  $Bind1_4_12 = (object)["bind" => function($v_6) use ($Bind1_4_12, $pure_5_13) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($Bind1_4_12, $pure_5_13, $v_6) {
  $__num = \func_num_args();
  $__local_var_8_14 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_5_13))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_4_12)->{'bind'})($v_6))(function($v2_9) use ($__local_var_8_14, $k_7) {
  $__num = \func_num_args();
  $__t15 = null;;
  if ($v2_9 instanceof \Data\Either\Data_Either_Left) {
$__t15 = ($__local_var_8_14)(($v2_9)->{'value0'});
goto end_branch_15;;
};
  if ($v2_9 instanceof \Data\Either\Data_Either_Right) {
$__t15 = ($k_7)(($v2_9)->{'value0'});
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
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
}, "Apply0" => function($_dollar___unused_6) use ($dictMonad_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_5_17 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_2)->{'Applicative0'})(null))->{'pure'}))($GLOBALS['Data_Either_Right'])), "Apply0" => function($_dollar___unused_5) use ($dictMonad_2) {
  $__num = \func_num_args();
  $__local_var_6_17 = (((((($dictMonad_2)->{'Bind1'})(null))->{'Apply0'})(null))->{'Functor0'})(null);
  $functorExceptT1_6_17 = (object)["map" => function($f_7) use ($__local_var_6_17) {
  $__num = \func_num_args();
  $__local_var_8_18 = (($__local_var_6_17)->{'map'})(function($m_8) use ($f_7) {
  $__num = \func_num_args();
  $__t18 = null;;
  if ($m_8 instanceof \Data\Either\Data_Either_Left) {
$__t18 = new \Data\Either\Data_Either_Left(($m_8)->{'value0'});
goto end_branch_18;;
};
  if ($m_8 instanceof \Data\Either\Data_Either_Right) {
$__t18 = new \Data\Either\Data_Either_Right(($f_7)(($m_8)->{'value0'}));
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_18) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_18)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_7_21 = (($dictMonad_2)->{'Bind1'})(null);
  $pure_8_22 = ((($dictMonad_2)->{'Applicative0'})(null))->{'pure'};
  $Bind1_7_21 = (object)["bind" => function($v_9) use ($Bind1_7_21, $pure_8_22) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Bind1_7_21, $pure_8_22, $v_9) {
  $__num = \func_num_args();
  $__local_var_11_23 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_8_22))($GLOBALS['Data_Either_Left']);
  $__res = ((($Bind1_7_21)->{'bind'})($v_9))(function($v2_12) use ($__local_var_11_23, $k_10) {
  $__num = \func_num_args();
  $__t24 = null;;
  if ($v2_12 instanceof \Data\Either\Data_Either_Left) {
$__t24 = ($__local_var_11_23)(($v2_12)->{'value0'});
goto end_branch_24;;
};
  if ($v2_12 instanceof \Data\Either\Data_Either_Right) {
$__t24 = ($k_10)(($v2_12)->{'value0'});
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
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
}, "Apply0" => function($_dollar___unused_9) use ($dictMonad_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_8_26 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_2);
  $__res = (object)["apply" => function($f_9) use ($Applicative0_8_26, $Bind1_7_21) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_8_26, $Bind1_7_21, $f_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_21)->{'bind'})($f_9))(function($f_prime__11) use ($Applicative0_8_26, $Bind1_7_21, $a_10) {
  $__num = \func_num_args();
  $__res = ((($Bind1_7_21)->{'bind'})($a_10))(function($a_prime__12) use ($Applicative0_8_26, $f_prime__11) {
  $__num = \func_num_args();
  $__res = (($Applicative0_8_26)->{'pure'})(($f_prime__11)($a_prime__12));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorExceptT1_6_17) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_6_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyExceptT_3_8 = (object)["apply" => function($f_6) use ($Applicative0_5_17, $Bind1_4_12) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Applicative0_5_17, $Bind1_4_12, $f_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_12)->{'bind'})($f_6))(function($f_prime__8) use ($Applicative0_5_17, $Bind1_4_12, $a_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_12)->{'bind'})($a_7))(function($a_prime__9) use ($Applicative0_5_17, $f_prime__8) {
  $__num = \func_num_args();
  $__res = (($Applicative0_5_17)->{'pure'})(($f_prime__8)($a_prime__9));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorExceptT1_3_8) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_3_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["parallel" => function($v_4) use ($dictParallel_0) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)->{'parallel'})($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_4) use ($dictParallel_0) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)->{'sequential'})($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($applyExceptT_3_8) {
  $__num = \func_num_args();
  $__res = $applyExceptT_3_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar___unused_4) use ($applyCompose_1_0) {
  $__num = \func_num_args();
  $__res = $applyCompose_1_0;
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
$GLOBALS['Control_Parallel_Class_monadParExceptT'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_monadmajParmajExceptmajT';

// Control_Parallel_Class_monadParCostar
function majControl_majParallel_majClass_monadmajParmajCostar($dictParallel_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_monadmajParmajCostar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequential1_1_0 = ($dictParallel_0)->{'sequential'};
  $parallel1_2_1 = ($dictParallel_0)->{'parallel'};
  $__res = (object)["parallel" => function($v_3) use ($sequential1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_3))($sequential1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_3) use ($parallel1_2_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_3))($parallel1_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_3) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_Costar_applyCostar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar___unused_3) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_Costar_applyCostar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_monadParCostar'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_monadmajParmajCostar';

// Control_Parallel_Class_monadParParCont
function majControl_majParallel_majClass_monadmajParmajParmajCont($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_monadmajParmajParmajCont';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT1_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyContT_1_0 = (object)["apply" => function($v_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
  $__num = \func_num_args();
  $__res = ($k_4)(($g_5)($a_6));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["parallel" => function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_2) {
  $__num = \func_num_args();
  $__res = $v_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyContT_1_0) {
  $__num = \func_num_args();
  $__res = $applyContT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar___unused_2) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Parallel_Class_applyParCont'])($dictMonadEffect_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_monadParParCont'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_monadmajParmajParmajCont';

// Control_Parallel_Class_functorParCont
function majControl_majParallel_majClass_functormajParmajCont($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_functormajParmajCont';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorContT_1_0 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($f_1, $v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["map" => function($f_2) use ($functorContT_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($functorContT_1_0)->{'map'})($f_2)))(function($v_3) {
  $__num = \func_num_args();
  $__res = $v_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_functorParCont'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_functormajParmajCont';

// Control_Parallel_Class_applyParCont
function majControl_majParallel_majClass_applymajParmajCont($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_applymajParmajCont';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (((($dictMonadEffect_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $__res = (object)["apply" => function($v_2) use ($Bind1_1_0, $dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($Bind1_1_0, $dictMonadEffect_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($Bind1_1_0, $dictMonadEffect_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(new \Data\Maybe\Data_Maybe_Nothing()))))(function($ra_5) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(new \Data\Maybe\Data_Maybe_Nothing()))))(function($rb_6) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})(($v_2)(function($a_7) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($rb_6))))(function($mb_8) use ($a_7, $dictMonadEffect_0, $k_4, $ra_5) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($mb_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(new \Data\Maybe\Data_Maybe_Just($a_7), $ra_5));
goto end_branch_1;;
};
  if ($mb_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($k_4)(($a_7)(($mb_8)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar___unused_7) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($b_8) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($ra_5))))(function($ma_9) use ($b_8, $dictMonadEffect_0, $k_4, $rb_6) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($ma_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(new \Data\Maybe\Data_Maybe_Just($b_8), $rb_6));
goto end_branch_2;;
};
  if ($ma_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($k_4)((($ma_9)->{'value0'})($b_8));
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
});
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Functor0" => function($_dollar___unused_2) {
  $__num = \func_num_args();
  $functorContT_3_3 = (object)["map" => function($f_3) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($f_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($f_3, $v_4) {
  $__num = \func_num_args();
  $__res = ($v_4)(function($a_6) use ($f_3, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($f_3)($a_6));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["map" => function($f_4) use ($functorContT_3_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($functorContT_3_3)->{'map'})($f_4)))(function($v_5) {
  $__num = \func_num_args();
  $__res = $v_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_applyParCont'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_applymajParmajCont';

// Control_Parallel_Class_applicativeParCont
function majControl_majParallel_majClass_applicativemajParmajCont($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_applicativemajParmajCont';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (((($dictMonadEffect_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $applyParCont1_1_0 = (object)["apply" => function($v_2) use ($Bind1_1_0, $dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($Bind1_1_0, $dictMonadEffect_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($Bind1_1_0, $dictMonadEffect_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(new \Data\Maybe\Data_Maybe_Nothing()))))(function($ra_5) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(new \Data\Maybe\Data_Maybe_Nothing()))))(function($rb_6) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})(($v_2)(function($a_7) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($rb_6))))(function($mb_8) use ($a_7, $dictMonadEffect_0, $k_4, $ra_5) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($mb_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(new \Data\Maybe\Data_Maybe_Just($a_7), $ra_5));
goto end_branch_1;;
};
  if ($mb_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($k_4)(($a_7)(($mb_8)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar___unused_7) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($b_8) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($ra_5))))(function($ma_9) use ($b_8, $dictMonadEffect_0, $k_4, $rb_6) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($ma_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(new \Data\Maybe\Data_Maybe_Just($b_8), $rb_6));
goto end_branch_2;;
};
  if ($ma_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($k_4)((($ma_9)->{'value0'})($b_8));
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
});
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Functor0" => function($_dollar___unused_2) {
  $__num = \func_num_args();
  $functorContT_3_3 = (object)["map" => function($f_3) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($f_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($f_3, $v_4) {
  $__num = \func_num_args();
  $__res = ($v_4)(function($a_6) use ($f_3, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($f_3)($a_6));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["map" => function($f_4) use ($functorContT_3_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($functorContT_3_3)->{'map'})($f_4)))(function($v_5) {
  $__num = \func_num_args();
  $__res = $v_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Apply0" => function($_dollar___unused_2) use ($applyParCont1_1_0) {
  $__num = \func_num_args();
  $__res = $applyParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_applicativeParCont'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_applicativemajParmajCont';

// Control_Parallel_Class_altParCont
function majControl_majParallel_majClass_altmajParmajCont($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_altmajParmajCont';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadEffect_0)->{'Monad0'})(null);
  $Bind1_2_1 = (($Monad0_1_0)->{'Bind1'})(null);
  $Applicative0_3_2 = (($Monad0_1_0)->{'Applicative0'})(null);
  $functorContT_4_3 = (object)["map" => function($f_4) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($f_4) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($f_4, $v_5) {
  $__num = \func_num_args();
  $__res = ($v_5)(function($a_7) use ($f_4, $k_6) {
  $__num = \func_num_args();
  $__res = ($k_6)(($f_4)($a_7));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorParCont1_4_3 = (object)["map" => function($f_5) use ($functorContT_4_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($functorContT_4_3)->{'map'})($f_5)))(function($v_6) {
  $__num = \func_num_args();
  $__res = $v_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["alt" => function($v_5) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadEffect_0, $v_5) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadEffect_0, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(false))))(function($done_8) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadEffect_0, $k_7, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})(($v_5)(function($a_9) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadEffect_0, $done_8, $k_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($done_8))))(function($b_10) use ($Applicative0_3_2, $Bind1_2_1, $a_9, $dictMonadEffect_0, $done_8, $k_7) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($b_10) {
$__t5 = (($Applicative0_3_2)->{'pure'})($GLOBALS['Data_Unit_unit']);
goto end_branch_5;;
};
  $__t5 = ((($Bind1_2_1)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(true, $done_8))))(function($_dollar___unused_11) use ($a_9, $k_7) {
  $__num = \func_num_args();
  $__res = ($k_7)($a_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar___unused_9) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadEffect_0, $done_8, $k_7, $v1_6) {
  $__num = \func_num_args();
  $__res = ($v1_6)(function($a_10) use ($Applicative0_3_2, $Bind1_2_1, $dictMonadEffect_0, $done_8, $k_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($done_8))))(function($b_11) use ($Applicative0_3_2, $Bind1_2_1, $a_10, $dictMonadEffect_0, $done_8, $k_7) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($b_11) {
$__t6 = (($Applicative0_3_2)->{'pure'})($GLOBALS['Data_Unit_unit']);
goto end_branch_6;;
};
  $__t6 = ((($Bind1_2_1)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(true, $done_8))))(function($_dollar___unused_12) use ($a_10, $k_7) {
  $__num = \func_num_args();
  $__res = ($k_7)($a_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_6:;
  $__res = $__t6;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorParCont1_4_3) {
  $__num = \func_num_args();
  $__res = $functorParCont1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_altParCont'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_altmajParmajCont';

// Control_Parallel_Class_plusParCont
function majControl_majParallel_majClass_plusmajParmajCont($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_plusmajParmajCont';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Applicative0_1_0 = (((($dictMonadEffect_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $Monad0_2_1 = (($dictMonadEffect_0)->{'Monad0'})(null);
  $Bind1_3_2 = (($Monad0_2_1)->{'Bind1'})(null);
  $Applicative0_4_3 = (($Monad0_2_1)->{'Applicative0'})(null);
  $functorContT_5_4 = (object)["map" => function($f_5) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($k_7) use ($f_5, $v_6) {
  $__num = \func_num_args();
  $__res = ($v_6)(function($a_8) use ($f_5, $k_7) {
  $__num = \func_num_args();
  $__res = ($k_7)(($f_5)($a_8));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorParCont1_5_4 = (object)["map" => function($f_6) use ($functorContT_5_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($functorContT_5_4)->{'map'})($f_6)))(function($v_7) {
  $__num = \func_num_args();
  $__res = $v_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altParCont1_2_1 = (object)["alt" => function($v_6) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadEffect_0, $v_6) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadEffect_0, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(false))))(function($done_9) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadEffect_0, $k_8, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)->{'bind'})(($v_6)(function($a_10) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadEffect_0, $done_9, $k_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($done_9))))(function($b_11) use ($Applicative0_4_3, $Bind1_3_2, $a_10, $dictMonadEffect_0, $done_9, $k_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($b_11) {
$__t6 = (($Applicative0_4_3)->{'pure'})($GLOBALS['Data_Unit_unit']);
goto end_branch_6;;
};
  $__t6 = ((($Bind1_3_2)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(true, $done_9))))(function($_dollar___unused_12) use ($a_10, $k_8) {
  $__num = \func_num_args();
  $__res = ($k_8)($a_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar___unused_10) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadEffect_0, $done_9, $k_8, $v1_7) {
  $__num = \func_num_args();
  $__res = ($v1_7)(function($a_11) use ($Applicative0_4_3, $Bind1_3_2, $dictMonadEffect_0, $done_9, $k_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($done_9))))(function($b_12) use ($Applicative0_4_3, $Bind1_3_2, $a_11, $dictMonadEffect_0, $done_9, $k_8) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ($b_12) {
$__t7 = (($Applicative0_4_3)->{'pure'})($GLOBALS['Data_Unit_unit']);
goto end_branch_7;;
};
  $__t7 = ((($Bind1_3_2)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(true, $done_9))))(function($_dollar___unused_13) use ($a_11, $k_8) {
  $__num = \func_num_args();
  $__res = ($k_8)($a_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_7:;
  $__res = $__t7;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorParCont1_5_4) {
  $__num = \func_num_args();
  $__res = $functorParCont1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => function($v_3) use ($Applicative0_1_0) {
  $__num = \func_num_args();
  $__res = (($Applicative0_1_0)->{'pure'})($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_3) use ($altParCont1_2_1) {
  $__num = \func_num_args();
  $__res = $altParCont1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_plusParCont'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_plusmajParmajCont';

// Control_Parallel_Class_alternativeParCont
function majControl_majParallel_majClass_alternativemajParmajCont($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majParallel_majClass_alternativemajParmajCont';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Bind1_1_0 = (((($dictMonadEffect_0)->{'Monad0'})(null))->{'Bind1'})(null);
  $applyParCont1_1_0 = (object)["apply" => function($v_2) use ($Bind1_1_0, $dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($Bind1_1_0, $dictMonadEffect_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($k_4) use ($Bind1_1_0, $dictMonadEffect_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(new \Data\Maybe\Data_Maybe_Nothing()))))(function($ra_5) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(new \Data\Maybe\Data_Maybe_Nothing()))))(function($rb_6) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})(($v_2)(function($a_7) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($rb_6))))(function($mb_8) use ($a_7, $dictMonadEffect_0, $k_4, $ra_5) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($mb_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(new \Data\Maybe\Data_Maybe_Just($a_7), $ra_5));
goto end_branch_1;;
};
  if ($mb_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($k_4)(($a_7)(($mb_8)->{'value0'}));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar___unused_7) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($b_8) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($ra_5))))(function($ma_9) use ($b_8, $dictMonadEffect_0, $k_4, $rb_6) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($ma_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(new \Data\Maybe\Data_Maybe_Just($b_8), $rb_6));
goto end_branch_2;;
};
  if ($ma_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($k_4)((($ma_9)->{'value0'})($b_8));
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
});
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}, "Functor0" => function($_dollar___unused_2) {
  $__num = \func_num_args();
  $functorContT_3_3 = (object)["map" => function($f_3) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($f_3) {
  $__num = \func_num_args();
  $__res = function($k_5) use ($f_3, $v_4) {
  $__num = \func_num_args();
  $__res = ($v_4)(function($a_6) use ($f_3, $k_5) {
  $__num = \func_num_args();
  $__res = ($k_5)(($f_3)($a_6));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["map" => function($f_4) use ($functorContT_3_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($functorContT_3_3)->{'map'})($f_4)))(function($v_5) {
  $__num = \func_num_args();
  $__res = $v_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeParCont1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_2) {
  $__num = \func_num_args();
  $__res = function($k_3) use ($a_2) {
  $__num = \func_num_args();
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Apply0" => function($_dollar___unused_2) use ($applyParCont1_1_0) {
  $__num = \func_num_args();
  $__res = $applyParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Applicative0_2_6 = (((($dictMonadEffect_0)->{'Monad0'})(null))->{'Applicative0'})(null);
  $Monad0_3_7 = (($dictMonadEffect_0)->{'Monad0'})(null);
  $Bind1_4_8 = (($Monad0_3_7)->{'Bind1'})(null);
  $Applicative0_5_9 = (($Monad0_3_7)->{'Applicative0'})(null);
  $functorContT_6_10 = (object)["map" => function($f_6) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = function($k_8) use ($f_6, $v_7) {
  $__num = \func_num_args();
  $__res = ($v_7)(function($a_9) use ($f_6, $k_8) {
  $__num = \func_num_args();
  $__res = ($k_8)(($f_6)($a_9));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $functorParCont1_6_10 = (object)["map" => function($f_7) use ($functorContT_6_10) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($functorContT_6_10)->{'map'})($f_7)))(function($v_8) {
  $__num = \func_num_args();
  $__res = $v_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altParCont1_3_7 = (object)["alt" => function($v_7) use ($Applicative0_5_9, $Bind1_4_8, $dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($Applicative0_5_9, $Bind1_4_8, $dictMonadEffect_0, $v_7) {
  $__num = \func_num_args();
  $__res = function($k_9) use ($Applicative0_5_9, $Bind1_4_8, $dictMonadEffect_0, $v1_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_8)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef__new(false))))(function($done_10) use ($Applicative0_5_9, $Bind1_4_8, $dictMonadEffect_0, $k_9, $v1_8, $v_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_8)->{'bind'})(($v_7)(function($a_11) use ($Applicative0_5_9, $Bind1_4_8, $dictMonadEffect_0, $done_10, $k_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_8)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($done_10))))(function($b_12) use ($Applicative0_5_9, $Bind1_4_8, $a_11, $dictMonadEffect_0, $done_10, $k_9) {
  $__num = \func_num_args();
  $__t12 = null;;
  if ($b_12) {
$__t12 = (($Applicative0_5_9)->{'pure'})($GLOBALS['Data_Unit_unit']);
goto end_branch_12;;
};
  $__t12 = ((($Bind1_4_8)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(true, $done_10))))(function($_dollar___unused_13) use ($a_11, $k_9) {
  $__num = \func_num_args();
  $__res = ($k_9)($a_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar___unused_11) use ($Applicative0_5_9, $Bind1_4_8, $dictMonadEffect_0, $done_10, $k_9, $v1_8) {
  $__num = \func_num_args();
  $__res = ($v1_8)(function($a_12) use ($Applicative0_5_9, $Bind1_4_8, $dictMonadEffect_0, $done_10, $k_9) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_8)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_read($done_10))))(function($b_13) use ($Applicative0_5_9, $Bind1_4_8, $a_12, $dictMonadEffect_0, $done_10, $k_9) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($b_13) {
$__t13 = (($Applicative0_5_9)->{'pure'})($GLOBALS['Data_Unit_unit']);
goto end_branch_13;;
};
  $__t13 = ((($Bind1_4_8)->{'bind'})((($dictMonadEffect_0)->{'liftEffect'})(\Effect\Ref\majEffect_majRef_write(true, $done_10))))(function($_dollar___unused_14) use ($a_12, $k_9) {
  $__num = \func_num_args();
  $__res = ($k_9)($a_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_13:;
  $__res = $__t13;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorParCont1_6_10) {
  $__num = \func_num_args();
  $__res = $functorParCont1_6_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusParCont1_2_6 = (object)["empty" => function($v_4) use ($Applicative0_2_6) {
  $__num = \func_num_args();
  $__res = (($Applicative0_2_6)->{'pure'})($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_4) use ($altParCont1_3_7) {
  $__num = \func_num_args();
  $__res = $altParCont1_3_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeParCont1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_3) use ($plusParCont1_2_6) {
  $__num = \func_num_args();
  $__res = $plusParCont1_2_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Parallel_Class_alternativeParCont'] = __NAMESPACE__ . '\\majControl_majParallel_majClass_alternativemajParmajCont';

