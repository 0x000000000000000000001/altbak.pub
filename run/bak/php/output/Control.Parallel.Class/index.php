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
\PhpursThunks::$thunks['Control_Parallel_Class_ParCont'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_ParCont"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_sequential'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_sequential"), recVars=[];
  $__res = ($dict_0)->sequential;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_parallel'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_parallel"), recVars=[];
  $__res = ($dict_0)->parallel;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_newtypeParCont'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_monadParWriterT'] = function() { $v = function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_monadParWriterT"), recVars=[];
  $applyWriterT_1_0 = (($GLOBALS['Control_Monad_Writer_Trans_applyWriterT'] ?? \PhpursThunks::eval('Control_Monad_Writer_Trans_applyWriterT')))((($dictMonoid_0)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictParallel_2) use ($applyWriterT_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $applyWriterT1_3_1 = ($applyWriterT_1_0)((($dictParallel_2)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $applyWriterT2_4_2 = ($applyWriterT_1_0)((($dictParallel_2)->Apply1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["parallel" => function($v_5) use ($dictParallel_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictParallel_2)->parallel)($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_5) use ($dictParallel_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictParallel_2)->sequential)($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_5) use ($applyWriterT1_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyWriterT1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($dollar__unused_5) use ($applyWriterT2_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyWriterT2_4_2;
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
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_monadParStar'] = function() { $v = function($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_monadParStar"), recVars=[];
  $parallel1_1_0 = ($dictParallel_0)->parallel;
  $sequential1_2_1 = ($dictParallel_0)->sequential;
  $applyStar_3_2 = (($GLOBALS['Data_Profunctor_Star_applyStar'] ?? \PhpursThunks::eval('Data_Profunctor_Star_applyStar')))((($dictParallel_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $applyStar1_4_3 = (($GLOBALS['Data_Profunctor_Star_applyStar'] ?? \PhpursThunks::eval('Data_Profunctor_Star_applyStar')))((($dictParallel_0)->Apply1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["parallel" => function($v_5) use ($parallel1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($parallel1_1_0))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_5) use ($sequential1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($sequential1_2_1))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_5) use ($applyStar_3_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyStar_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($dollar__unused_5) use ($applyStar1_4_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyStar1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_monadParReaderT'] = function() { $v = function($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_monadParReaderT"), recVars=[];
  $applyReaderT_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_applyReaderT')))((($dictParallel_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $applyReaderT1_2_1 = (($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_applyReaderT')))((($dictParallel_0)->Apply1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["parallel" => (($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_mapReaderT')))(($dictParallel_0)->parallel), "sequential" => (($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_mapReaderT')))(($dictParallel_0)->sequential), "Apply0" => function($dollar__unused_3) use ($applyReaderT_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyReaderT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($dollar__unused_3) use ($applyReaderT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_monadParMaybeT'] = function() { $v = function($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_monadParMaybeT"), recVars=[];
  $applyCompose_1_0 = ((($GLOBALS['Data_Functor_Compose_applyCompose'] ?? \PhpursThunks::eval('Data_Functor_Compose_applyCompose')))((($dictParallel_0)->Apply1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))))(($GLOBALS['Data_Maybe_applyMaybe'] ?? \PhpursThunks::eval('Data_Maybe_applyMaybe')));
  $__res = function($dictMonad_2) use ($applyCompose_1_0, $dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $applyMaybeT_3_1 = (($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applyMaybeT')))($dictMonad_2);
  $__res = (object)["parallel" => function($v_4) use ($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictParallel_0)->parallel)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_4) use ($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictParallel_0)->sequential)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_4) use ($applyMaybeT_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyMaybeT_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($dollar__unused_4) use ($applyCompose_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_monadParExceptT'] = function() { $v = function($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_monadParExceptT"), recVars=[];
  $applyCompose_1_0 = ((($GLOBALS['Data_Functor_Compose_applyCompose'] ?? \PhpursThunks::eval('Data_Functor_Compose_applyCompose')))((($dictParallel_0)->Apply1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))))(($GLOBALS['Data_Either_applyEither'] ?? \PhpursThunks::eval('Data_Either_applyEither')));
  $__res = function($dictMonad_2) use ($applyCompose_1_0, $dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $applyExceptT_3_1 = (($GLOBALS['Control_Monad_Except_Trans_applyExceptT'] ?? \PhpursThunks::eval('Control_Monad_Except_Trans_applyExceptT')))($dictMonad_2);
  $__res = (object)["parallel" => function($v_4) use ($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictParallel_0)->parallel)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_4) use ($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictParallel_0)->sequential)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_4) use ($applyExceptT_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyExceptT_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($dollar__unused_4) use ($applyCompose_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_monadParCostar'] = function() { $v = function($dictParallel_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_monadParCostar"), recVars=[];
  $sequential1_1_0 = ($dictParallel_0)->sequential;
  $parallel1_2_1 = ($dictParallel_0)->parallel;
  $__res = (object)["parallel" => function($v_3) use ($sequential1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($v_3))($sequential1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequential" => function($v_3) use ($parallel1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($v_3))($parallel1_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Functor_Costar_applyCostar'] ?? \PhpursThunks::eval('Data_Functor_Costar_applyCostar'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($dollar__unused_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Functor_Costar_applyCostar'] ?? \PhpursThunks::eval('Data_Functor_Costar_applyCostar'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_monadParParCont'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_monadParParCont"), recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  while (true) {
$functorContT1_1_0 = (object)["map" => (function() {
  $__fn = function($f_1, $v_2 = null, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
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
$applyContT_2_1 = (object)["apply" => (function() {
  $__fn = function($v_2, $v1_3 = null, $k_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ($v_2)(function($g_5) use ($k_4, $v1_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ($v1_3)(function($a_6) use ($g_5, $k_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
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
})(), "Functor0" => function($dollar__unused_2) use ($functorContT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = $functorContT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$__res = (object)["parallel" => ($GLOBALS['Control_Parallel_Class_ParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_ParCont')), "sequential" => function($v_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = $v_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_3) use ($applyContT_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = $applyContT_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($dollar__unused_3) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = (($GLOBALS['Control_Parallel_Class_applyParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_applyParCont')))($dictMonadEffect_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_functorParCont'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_functorParCont"), recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  while (true) {
$__res = (object)["map" => function($f_1) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Parallel_Class_monadParParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_monadParParCont')))($dictMonadEffect_0))->parallel))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((function() use ($f_1) {
  $__fn = function($v_2, $k_3 = null) use ($f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ($v_2)(function($a_4) use ($f_1, $k_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
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
})()))(((($GLOBALS['Control_Parallel_Class_monadParParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_monadParParCont')))($dictMonadEffect_0))->sequential));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_applyParCont'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_applyParCont"), recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  while (true) {
$Bind1_1_0 = (((($dictMonadEffect_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (object)["apply" => (function() use ($Bind1_1_0, $dictMonadEffect_0) {
  $__fn = function($v_2, $v1_3 = null, $k_4 = null) use ($Bind1_1_0, $dictMonadEffect_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ((($Bind1_1_0)->bind)((($dictMonadEffect_0)->liftEffect)((($GLOBALS['Effect_Ref__new'] ?? \PhpursThunks::eval('Effect_Ref__new')))(new Phpurs_Data0("Nothing")))))(function($ra_5) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $v1_3, $v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ((($Bind1_1_0)->bind)((($dictMonadEffect_0)->liftEffect)((($GLOBALS['Effect_Ref__new'] ?? \PhpursThunks::eval('Effect_Ref__new')))(new Phpurs_Data0("Nothing")))))(function($rb_6) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $v1_3, $v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ((($Bind1_1_0)->bind)(($v_2)(function($a_7) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ((($Bind1_1_0)->bind)((($dictMonadEffect_0)->liftEffect)((($GLOBALS['Effect_Ref_read'] ?? \PhpursThunks::eval('Effect_Ref_read')))($rb_6))))(function($mb_8) use ($a_7, $dictMonadEffect_0, $k_4, $ra_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  if ((is_object($mb_8) && (($mb_8)->tag === "Nothing"))) {
$__t1 = (($dictMonadEffect_0)->liftEffect)(((($GLOBALS['Effect_Ref_write'] ?? \PhpursThunks::eval('Effect_Ref_write')))(new Phpurs_Data1("Just", $a_7)))($ra_5));
} else {
if ((is_object($mb_8) && (($mb_8)->tag === "Just"))) {
$__t1 = ($k_4)(($a_7)(($mb_8)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_7) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6, $v1_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ($v1_3)(function($b_8) use ($Bind1_1_0, $dictMonadEffect_0, $k_4, $ra_5, $rb_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = ((($Bind1_1_0)->bind)((($dictMonadEffect_0)->liftEffect)((($GLOBALS['Effect_Ref_read'] ?? \PhpursThunks::eval('Effect_Ref_read')))($ra_5))))(function($ma_9) use ($b_8, $dictMonadEffect_0, $k_4, $rb_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  if ((is_object($ma_9) && (($ma_9)->tag === "Nothing"))) {
$__t2 = (($dictMonadEffect_0)->liftEffect)(((($GLOBALS['Effect_Ref_write'] ?? \PhpursThunks::eval('Effect_Ref_write')))(new Phpurs_Data1("Just", $b_8)))($rb_6));
} else {
if ((is_object($ma_9) && (($ma_9)->tag === "Just"))) {
$__t2 = ($k_4)((($ma_9)->value0)($b_8));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_2) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Parallel_Class_monadParParCont","Control_Parallel_Class_functorParCont","Control_Parallel_Class_applyParCont"];
  $__res = (($GLOBALS['Control_Parallel_Class_functorParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_functorParCont')))($dictMonadEffect_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_applicativeParCont'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_applicativeParCont"), recVars=[];
  $applyParCont1_1_0 = (($GLOBALS['Control_Parallel_Class_applyParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_applyParCont')))($dictMonadEffect_0);
  $__res = (object)["pure" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Parallel_Class_monadParParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_monadParParCont')))($dictMonadEffect_0))->parallel))((function() {
  $__fn = function($a_2, $k_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($k_3)($a_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()), "Apply0" => function($dollar__unused_2) use ($applyParCont1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_altParCont'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_altParCont"), recVars=[];
  $Monad0_1_0 = (($dictMonadEffect_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Bind1_2_1 = (($Monad0_1_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_3_2 = (($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorParCont1_4_3 = (($GLOBALS['Control_Parallel_Class_functorParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_functorParCont')))($dictMonadEffect_0);
  $__res = (object)["alt" => (function() use ($Bind1_2_1, $__local_var_3_2, $dictMonadEffect_0) {
  $__fn = function($v_5, $v1_6 = null, $k_7 = null) use ($Bind1_2_1, $__local_var_3_2, $dictMonadEffect_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($Bind1_2_1)->bind)((($dictMonadEffect_0)->liftEffect)((($GLOBALS['Effect_Ref__new'] ?? \PhpursThunks::eval('Effect_Ref__new')))(false))))(function($done_8) use ($Bind1_2_1, $__local_var_3_2, $dictMonadEffect_0, $k_7, $v1_6, $v_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($Bind1_2_1)->bind)(($v_5)(function($a_9) use ($Bind1_2_1, $__local_var_3_2, $dictMonadEffect_0, $done_8, $k_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($Bind1_2_1)->bind)((($dictMonadEffect_0)->liftEffect)((($GLOBALS['Effect_Ref_read'] ?? \PhpursThunks::eval('Effect_Ref_read')))($done_8))))(function($b_10) use ($Bind1_2_1, $__local_var_3_2, $a_9, $dictMonadEffect_0, $done_8, $k_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ($b_10) {
$__t4 = (($__local_var_3_2)->pure)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
} else {
$__t4 = ((($Bind1_2_1)->bind)((($dictMonadEffect_0)->liftEffect)(((($GLOBALS['Effect_Ref_write'] ?? \PhpursThunks::eval('Effect_Ref_write')))(true))($done_8))))(function($dollar__unused_11) use ($a_9, $k_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($k_7)($a_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_9) use ($Bind1_2_1, $__local_var_3_2, $dictMonadEffect_0, $done_8, $k_7, $v1_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($v1_6)(function($a_10) use ($Bind1_2_1, $__local_var_3_2, $dictMonadEffect_0, $done_8, $k_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($Bind1_2_1)->bind)((($dictMonadEffect_0)->liftEffect)((($GLOBALS['Effect_Ref_read'] ?? \PhpursThunks::eval('Effect_Ref_read')))($done_8))))(function($b_11) use ($Bind1_2_1, $__local_var_3_2, $a_10, $dictMonadEffect_0, $done_8, $k_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ($b_11) {
$__t5 = (($__local_var_3_2)->pure)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
} else {
$__t5 = ((($Bind1_2_1)->bind)((($dictMonadEffect_0)->liftEffect)(((($GLOBALS['Effect_Ref_write'] ?? \PhpursThunks::eval('Effect_Ref_write')))(true))($done_8))))(function($dollar__unused_12) use ($a_10, $k_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($k_7)($a_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
};
  $__res = $__t5;
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
})(), "Functor0" => function($dollar__unused_5) use ($functorParCont1_4_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $functorParCont1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_plusParCont'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_plusParCont"), recVars=[];
  $altParCont1_1_0 = (($GLOBALS['Control_Parallel_Class_altParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_altParCont')))($dictMonadEffect_0);
  $__res = (object)["empty" => function($v_2) use ($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((((($dictMonadEffect_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($dollar__unused_2) use ($altParCont1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $altParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Parallel_Class_alternativeParCont'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Parallel_Class_alternativeParCont"), recVars=[];
  $applicativeParCont1_1_0 = (($GLOBALS['Control_Parallel_Class_applicativeParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_applicativeParCont')))($dictMonadEffect_0);
  $plusParCont1_2_1 = (($GLOBALS['Control_Parallel_Class_plusParCont'] ?? \PhpursThunks::eval('Control_Parallel_Class_plusParCont')))($dictMonadEffect_0);
  $__res = (object)["Applicative0" => function($dollar__unused_3) use ($applicativeParCont1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applicativeParCont1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_3) use ($plusParCont1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $plusParCont1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };



















