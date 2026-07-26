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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


// Control_Parallel_Class_ParCont
$GLOBALS['Control_Parallel_Class_ParCont'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_sequential
$GLOBALS['Control_Parallel_Class_sequential'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['sequential'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_parallel
$GLOBALS['Control_Parallel_Class_parallel'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['parallel'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_newtypeParCont
$GLOBALS['Control_Parallel_Class_newtypeParCont'] = ["Coercible0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Parallel_Class_monadParWriterT
$GLOBALS['Control_Parallel_Class_monadParWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonoid_0)['Semigroup0'])(null);
  $applyWriterT_2_1 = function($dictApply_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $Functor0_3_1 = (($dictApply_2)['Functor0'])(null);
  $functorWriterT1_4_2 = ["map" => function($f_4 = null) use ($Functor0_3_1) {
  $__num = \func_num_args();
  $__res = (($Functor0_3_1)['map'])(function($v_5 = null) use ($f_4) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_4)(($v_5)->{'value0'}), ($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["apply" => (function() use ($Functor0_3_1, $__local_var_1_0, $dictApply_2) {
  $__fn = function($v_5 = null, $v1_6 = null) use ($Functor0_3_1, $__local_var_1_0, $dictApply_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictApply_2)['apply'])(((($Functor0_3_1)['map'])((function() use ($__local_var_1_0) {
  $__fn = function($v3_7 = null, $v4_8 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", (($v3_7)->{'value0'})(($v4_8)->{'value0'}), ((($__local_var_1_0)['append'])(($v3_7)->{'value1'}))(($v4_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($v_5)))($v1_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_5 = null) use ($functorWriterT1_4_2) {
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
  $__res = function($dictParallel_3 = null) use ($applyWriterT_2_1) {
  $__num = \func_num_args();
  $applyWriterT1_4_4 = ($applyWriterT_2_1)((($dictParallel_3)['Apply0'])(null));
  $applyWriterT2_5_5 = ($applyWriterT_2_1)((($dictParallel_3)['Apply1'])(null));
  $__res = ["parallel" => function($v_6 = null) use ($dictParallel_3) {
  $__num = \func_num_args();
  $__res = (($dictParallel_3)['parallel'])($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_6 = null) use ($dictParallel_3) {
  $__num = \func_num_args();
  $__res = (($dictParallel_3)['sequential'])($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_6 = null) use ($applyWriterT1_4_4) {
  $__num = \func_num_args();
  $__res = $applyWriterT1_4_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar__unused_6 = null) use ($applyWriterT2_5_5) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_5_5;
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

// Control_Parallel_Class_monadParStar
$GLOBALS['Control_Parallel_Class_monadParStar'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $parallel1_1_0 = ($dictParallel_0)['parallel'];
  $sequential1_2_1 = ($dictParallel_0)['sequential'];
  $applyStar_3_2 = ($GLOBALS['Data_Profunctor_Star_applyStar'])((($dictParallel_0)['Apply0'])(null));
  $applyStar1_4_3 = ($GLOBALS['Data_Profunctor_Star_applyStar'])((($dictParallel_0)['Apply1'])(null));
  $__res = ["parallel" => function($v_5 = null) use ($parallel1_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($parallel1_1_0))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_5 = null) use ($sequential1_2_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($sequential1_2_1))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_5 = null) use ($applyStar_3_2) {
  $__num = \func_num_args();
  $__res = $applyStar_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar__unused_5 = null) use ($applyStar1_4_3) {
  $__num = \func_num_args();
  $__res = $applyStar1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_monadParReaderT
$GLOBALS['Control_Parallel_Class_monadParReaderT'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $applyReaderT_1_0 = ($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'])((($dictParallel_0)['Apply0'])(null));
  $applyReaderT1_2_1 = ($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'])((($dictParallel_0)['Apply1'])(null));
  $__res = ["parallel" => ($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'])(($dictParallel_0)['parallel']), "sequential" => ($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'])(($dictParallel_0)['sequential']), "Apply0" => function($_dollar__unused_3 = null) use ($applyReaderT_1_0) {
  $__num = \func_num_args();
  $__res = $applyReaderT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar__unused_3 = null) use ($applyReaderT1_2_1) {
  $__num = \func_num_args();
  $__res = $applyReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_monadParMaybeT
$GLOBALS['Control_Parallel_Class_monadParMaybeT'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictParallel_0)['Apply1'])(null);
  $Functor0_2_1 = (($__local_var_1_0)['Functor0'])(null);
  $__local_var_3_2 = (($GLOBALS['Data_Maybe_applyMaybe'])['Functor0'])(null);
  $functorCompose2_4_3 = ["map" => (function() use ($Functor0_2_1, $__local_var_3_2) {
  $__fn = function($f_4 = null, $v_5 = null) use ($Functor0_2_1, $__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Functor0_2_1)['map'])((($__local_var_3_2)['map'])($f_4)))($v_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $applyCompose_4_3 = ["apply" => (function() use ($Functor0_2_1, $__local_var_1_0) {
  $__fn = function($v_5 = null, $v1_6 = null) use ($Functor0_2_1, $__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['apply'])(((($Functor0_2_1)['map'])(($GLOBALS['Data_Maybe_applyMaybe'])['apply']))($v_5)))($v1_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_5 = null) use ($functorCompose2_4_3) {
  $__num = \func_num_args();
  $__res = $functorCompose2_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonad_5 = null) use ($applyCompose_4_3, $dictParallel_0) {
  $__num = \func_num_args();
  $applyMaybeT_6_5 = ($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'])($dictMonad_5);
  $__res = ["parallel" => function($v_7 = null) use ($dictParallel_0) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)['parallel'])($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_7 = null) use ($dictParallel_0) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)['sequential'])($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_7 = null) use ($applyMaybeT_6_5) {
  $__num = \func_num_args();
  $__res = $applyMaybeT_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar__unused_7 = null) use ($applyCompose_4_3) {
  $__num = \func_num_args();
  $__res = $applyCompose_4_3;
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

// Control_Parallel_Class_monadParExceptT
$GLOBALS['Control_Parallel_Class_monadParExceptT'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictParallel_0)['Apply1'])(null);
  $Functor0_2_1 = (($__local_var_1_0)['Functor0'])(null);
  $__local_var_3_2 = (($GLOBALS['Data_Either_applyEither'])['Functor0'])(null);
  $functorCompose2_4_3 = ["map" => (function() use ($Functor0_2_1, $__local_var_3_2) {
  $__fn = function($f_4 = null, $v_5 = null) use ($Functor0_2_1, $__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Functor0_2_1)['map'])((($__local_var_3_2)['map'])($f_4)))($v_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $applyCompose_4_3 = ["apply" => (function() use ($Functor0_2_1, $__local_var_1_0) {
  $__fn = function($v_5 = null, $v1_6 = null) use ($Functor0_2_1, $__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['apply'])(((($Functor0_2_1)['map'])(($GLOBALS['Data_Either_applyEither'])['apply']))($v_5)))($v1_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_5 = null) use ($functorCompose2_4_3) {
  $__num = \func_num_args();
  $__res = $functorCompose2_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonad_5 = null) use ($applyCompose_4_3, $dictParallel_0) {
  $__num = \func_num_args();
  $applyExceptT_6_5 = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_5);
  $__res = ["parallel" => function($v_7 = null) use ($dictParallel_0) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)['parallel'])($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_7 = null) use ($dictParallel_0) {
  $__num = \func_num_args();
  $__res = (($dictParallel_0)['sequential'])($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_7 = null) use ($applyExceptT_6_5) {
  $__num = \func_num_args();
  $__res = $applyExceptT_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar__unused_7 = null) use ($applyCompose_4_3) {
  $__num = \func_num_args();
  $__res = $applyCompose_4_3;
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

// Control_Parallel_Class_monadParCostar
$GLOBALS['Control_Parallel_Class_monadParCostar'] = function($dictParallel_0 = null) {
  $__num = \func_num_args();
  $sequential1_1_0 = ($dictParallel_0)['sequential'];
  $parallel1_2_1 = ($dictParallel_0)['parallel'];
  $__res = ["parallel" => function($v_3 = null) use ($sequential1_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($v_3))($sequential1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_3 = null) use ($parallel1_2_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($v_3))($parallel1_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_3 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_Costar_applyCostar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar__unused_3 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_Costar_applyCostar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_monadParParCont
$GLOBALS['Control_Parallel_Class_monadParParCont'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $functorContT1_1_0 = ["map" => (function() {
  $__fn = function($f_1 = null, $v_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($a_4 = null) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $applyContT_2_1 = ["apply" => (function() {
  $__fn = function($v_2 = null, $v1_3 = null, $k_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($v_2)(function($g_5 = null) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  $__res = ($v1_3)(function($a_6 = null) use ($g_5, $k_4) {
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_2 = null) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["parallel" => $GLOBALS['Control_Parallel_Class_ParCont'], "sequential" => function($v_3 = null) {
  $__num = \func_num_args();
  $__res = $v_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_3 = null) use ($applyContT_2_1) {
  $__num = \func_num_args();
  $__res = $applyContT_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar__unused_3 = null) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Parallel_Class_applyParCont'])($dictMonadEffect_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_functorParCont
$GLOBALS['Control_Parallel_Class_functorParCont'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => function($f_1 = null) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($GLOBALS['Control_Parallel_Class_monadParParCont'])($dictMonadEffect_0))['parallel']))(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((function() use ($f_1) {
  $__fn = function($v_2 = null, $k_3 = null) use ($f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_2)(function($a_4 = null) use ($f_1, $k_3) {
  $__num = \func_num_args();
  $__res = ($k_3)(($f_1)($a_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($GLOBALS['Control_Parallel_Class_monadParParCont'])($dictMonadEffect_0))['sequential']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_applyParCont
$GLOBALS['Control_Parallel_Class_applyParCont'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $Bind1_1_0 = (((($dictMonadEffect_0)['Monad0'])(null))['Bind1'])(null);
  $discard1_2_1 = (($GLOBALS['Control_Bind_discardUnit'])['discard'])($Bind1_1_0);
  $__res = ["apply" => (function() use ($Bind1_1_0, $dictMonadEffect_0, $discard1_2_1) {
  $__fn = function($v_3 = null, $v1_4 = null, $k_5 = null) use ($Bind1_1_0, $dictMonadEffect_0, $discard1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Bind1_1_0)['bind'])((($dictMonadEffect_0)['liftEffect'])(($GLOBALS['Effect_Ref__new'])(new Phpurs_Data0("Nothing")))))(function($ra_6 = null) use ($Bind1_1_0, $dictMonadEffect_0, $discard1_2_1, $k_5, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)['bind'])((($dictMonadEffect_0)['liftEffect'])(($GLOBALS['Effect_Ref__new'])(new Phpurs_Data0("Nothing")))))(function($rb_7 = null) use ($Bind1_1_0, $dictMonadEffect_0, $discard1_2_1, $k_5, $ra_6, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = (($discard1_2_1)(($v_3)(function($a_8 = null) use ($Bind1_1_0, $dictMonadEffect_0, $k_5, $ra_6, $rb_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)['bind'])((($dictMonadEffect_0)['liftEffect'])(($GLOBALS['Effect_Ref_read'])($rb_7))))(function($mb_9 = null) use ($a_8, $dictMonadEffect_0, $k_5, $ra_6) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($mb_9) && (($mb_9)->{'tag'} === "Nothing"))) {
$__t2 = (($dictMonadEffect_0)['liftEffect'])((($GLOBALS['Effect_Ref_write'])(new Phpurs_Data1("Just", $a_8)))($ra_6));
goto end_branch_2;;
};
  if ((is_object($mb_9) && (($mb_9)->{'tag'} === "Just"))) {
$__t2 = ($k_5)(($a_8)(($mb_9)->{'value0'}));
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
})))(function($_dollar__unused_8 = null) use ($Bind1_1_0, $dictMonadEffect_0, $k_5, $ra_6, $rb_7, $v1_4) {
  $__num = \func_num_args();
  $__res = ($v1_4)(function($b_9 = null) use ($Bind1_1_0, $dictMonadEffect_0, $k_5, $ra_6, $rb_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_1_0)['bind'])((($dictMonadEffect_0)['liftEffect'])(($GLOBALS['Effect_Ref_read'])($ra_6))))(function($ma_10 = null) use ($b_9, $dictMonadEffect_0, $k_5, $rb_7) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((is_object($ma_10) && (($ma_10)->{'tag'} === "Nothing"))) {
$__t3 = (($dictMonadEffect_0)['liftEffect'])((($GLOBALS['Effect_Ref_write'])(new Phpurs_Data1("Just", $b_9)))($rb_7));
goto end_branch_3;;
};
  if ((is_object($ma_10) && (($ma_10)->{'tag'} === "Just"))) {
$__t3 = ($k_5)((($ma_10)->{'value0'})($b_9));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_3 = null) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Parallel_Class_functorParCont'])($dictMonadEffect_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_applicativeParCont
$GLOBALS['Control_Parallel_Class_applicativeParCont'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $applyParCont1_1_0 = ($GLOBALS['Control_Parallel_Class_applyParCont'])($dictMonadEffect_0);
  $__res = ["pure" => ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])((($GLOBALS['Control_Parallel_Class_monadParParCont'])($dictMonadEffect_0))['parallel']))((function() {
  $__fn = function($a_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()), "Apply0" => function($_dollar__unused_2 = null) use ($applyParCont1_1_0) {
  $__num = \func_num_args();
  $__res = $applyParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_altParCont
$GLOBALS['Control_Parallel_Class_altParCont'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadEffect_0)['Monad0'])(null);
  $Bind1_2_1 = (($Monad0_1_0)['Bind1'])(null);
  $discard1_3_2 = (($GLOBALS['Control_Bind_discardUnit'])['discard'])($Bind1_2_1);
  $__local_var_4_3 = (($Monad0_1_0)['Applicative0'])(null);
  $functorParCont1_5_4 = ($GLOBALS['Control_Parallel_Class_functorParCont'])($dictMonadEffect_0);
  $__res = ["alt" => (function() use ($Bind1_2_1, $__local_var_4_3, $dictMonadEffect_0, $discard1_3_2) {
  $__fn = function($v_6 = null, $v1_7 = null, $k_8 = null) use ($Bind1_2_1, $__local_var_4_3, $dictMonadEffect_0, $discard1_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Bind1_2_1)['bind'])((($dictMonadEffect_0)['liftEffect'])(($GLOBALS['Effect_Ref__new'])(false))))(function($done_9 = null) use ($Bind1_2_1, $__local_var_4_3, $dictMonadEffect_0, $discard1_3_2, $k_8, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = (($discard1_3_2)(($v_6)(function($a_10 = null) use ($Bind1_2_1, $__local_var_4_3, $dictMonadEffect_0, $discard1_3_2, $done_9, $k_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)['bind'])((($dictMonadEffect_0)['liftEffect'])(($GLOBALS['Effect_Ref_read'])($done_9))))(function($b_11 = null) use ($__local_var_4_3, $a_10, $dictMonadEffect_0, $discard1_3_2, $done_9, $k_8) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($b_11) {
$__t5 = (($__local_var_4_3)['pure'])($GLOBALS['Data_Unit_unit']);
goto end_branch_5;;
};
  $__t5 = (($discard1_3_2)((($dictMonadEffect_0)['liftEffect'])((($GLOBALS['Effect_Ref_write'])(true))($done_9))))(function($_dollar__unused_12 = null) use ($a_10, $k_8) {
  $__num = \func_num_args();
  $__res = ($k_8)($a_10);
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
})))(function($_dollar__unused_10 = null) use ($Bind1_2_1, $__local_var_4_3, $dictMonadEffect_0, $discard1_3_2, $done_9, $k_8, $v1_7) {
  $__num = \func_num_args();
  $__res = ($v1_7)(function($a_11 = null) use ($Bind1_2_1, $__local_var_4_3, $dictMonadEffect_0, $discard1_3_2, $done_9, $k_8) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)['bind'])((($dictMonadEffect_0)['liftEffect'])(($GLOBALS['Effect_Ref_read'])($done_9))))(function($b_12 = null) use ($__local_var_4_3, $a_11, $dictMonadEffect_0, $discard1_3_2, $done_9, $k_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($b_12) {
$__t6 = (($__local_var_4_3)['pure'])($GLOBALS['Data_Unit_unit']);
goto end_branch_6;;
};
  $__t6 = (($discard1_3_2)((($dictMonadEffect_0)['liftEffect'])((($GLOBALS['Effect_Ref_write'])(true))($done_9))))(function($_dollar__unused_13 = null) use ($a_11, $k_8) {
  $__num = \func_num_args();
  $__res = ($k_8)($a_11);
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_6 = null) use ($functorParCont1_5_4) {
  $__num = \func_num_args();
  $__res = $functorParCont1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_plusParCont
$GLOBALS['Control_Parallel_Class_plusParCont'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $altParCont1_1_0 = ($GLOBALS['Control_Parallel_Class_altParCont'])($dictMonadEffect_0);
  $__res = ["empty" => function($v_2 = null) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__res = (((((($dictMonadEffect_0)['Monad0'])(null))['Applicative0'])(null))['pure'])($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar__unused_2 = null) use ($altParCont1_1_0) {
  $__num = \func_num_args();
  $__res = $altParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Parallel_Class_alternativeParCont
$GLOBALS['Control_Parallel_Class_alternativeParCont'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $applicativeParCont1_1_0 = ($GLOBALS['Control_Parallel_Class_applicativeParCont'])($dictMonadEffect_0);
  $plusParCont1_2_1 = ($GLOBALS['Control_Parallel_Class_plusParCont'])($dictMonadEffect_0);
  $__res = ["Applicative0" => function($_dollar__unused_3 = null) use ($applicativeParCont1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_3 = null) use ($plusParCont1_2_1) {
  $__num = \func_num_args();
  $__res = $plusParCont1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

