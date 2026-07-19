<?php

namespace Control\Monad\Reader\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Reader.Trans, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Distributive, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Reader.Trans, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Distributive, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Distributive/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Effect.Class/index.php';
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
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_ReaderT'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_ReaderT"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_withReaderT'] = function() { $v = (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_withReaderT"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($v_1))($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_runReaderT'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_runReaderT"), recVars=[];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_newtypeReaderT'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadTransReaderT'] = function() { $v = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_mapReaderT'] = function() { $v = (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_mapReaderT"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_0))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_functorReaderT'] = function() { $v = function($dictFunctor_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_functorReaderT"), recVars=[];
  $__res = (object)["map" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_mapReaderT'))))(($dictFunctor_0)->map)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_distributiveReaderT'] = function() { $v = function($dictDistributive_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_distributiveReaderT"), recVars=["Control_Monad_Reader_Trans_distributiveReaderT"];
  while (true) {
$functorReaderT1_1_0 = (object)["map" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_mapReaderT'))))(((($dictDistributive_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)];
$__res = (object)["distribute" => function($dictFunctor_2) use ($dictDistributive_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Reader_Trans_distributiveReaderT"];
  $collect1_3_1 = (($dictDistributive_0)->collect)($dictFunctor_2);
  $__res = (function() use ($collect1_3_1) {
  $__fn = function($a_4, $e_5 = null) use ($collect1_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Reader_Trans_distributiveReaderT"];
  $__res = (($collect1_3_1)(function($r_6) use ($e_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Reader_Trans_distributiveReaderT"];
  $__res = ($r_6)($e_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "collect" => (function() use ($dictDistributive_0) {
  $__fn = function($dictFunctor_2, $f_3 = null) use ($dictDistributive_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Reader_Trans_distributiveReaderT"];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Control_Monad_Reader_Trans_distributiveReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_distributiveReaderT')))($dictDistributive_0))->distribute)($dictFunctor_2)))((($dictFunctor_2)->map)($f_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_2) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Reader_Trans_distributiveReaderT"];
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_applyReaderT'] = function() { $v = function($dictApply_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_applyReaderT"), recVars=[];
  $functorReaderT1_1_0 = (object)["map" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_mapReaderT'))))(((($dictApply_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)];
  $__res = (object)["apply" => (function() use ($dictApply_0) {
  $__fn = function($v_2, $v1_3 = null, $r_4 = null) use ($dictApply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictApply_0)->apply)(($v_2)($r_4)))(($v1_3)($r_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_2) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_bindReaderT'] = function() { $v = function($dictBind_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_bindReaderT"), recVars=[];
  $applyReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_applyReaderT')))((($dictBind_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["bind" => (function() use ($dictBind_0) {
  $__fn = function($v_2, $k_3 = null, $r_4 = null) use ($dictBind_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictBind_0)->bind)(($v_2)($r_4)))(function($a_5) use ($k_3, $r_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($k_3)($a_5))($r_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_2) use ($applyReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_semigroupReaderT'] = function() { $v = function($dictApply_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_semigroupReaderT"), recVars=[];
  $lift2_1_0 = (($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))((($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_applyReaderT')))($dictApply_0));
  $__res = function($dictSemigroup_2) use ($lift2_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["append" => ($lift2_1_0)(($dictSemigroup_2)->append)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_applicativeReaderT'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_applicativeReaderT"), recVars=[];
  $applyReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_applyReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_applyReaderT')))((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["pure" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))(($dictApplicative_0)->pure)), "Apply0" => function($dollar__unused_2) use ($applyReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadReaderT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadReaderT"), recVars=[];
  $applicativeReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_applicativeReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_applicativeReaderT')))((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $bindReaderT1_2_1 = (($GLOBALS['Control_Monad_Reader_Trans_bindReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_bindReaderT')))((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["Applicative0" => function($dollar__unused_3) use ($applicativeReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applicativeReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3) use ($bindReaderT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $bindReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadAskReaderT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadAskReaderT"), recVars=[];
  $monadReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))($dictMonad_0);
  $__res = (object)["ask" => ((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure, "Monad0" => function($dollar__unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadReaderReaderT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadReaderReaderT"), recVars=[];
  $monadReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))($dictMonad_0);
  $monadAskReaderT1_2_1 = (object)["ask" => ((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure, "Monad0" => function($dollar__unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["local" => ($GLOBALS['Control_Monad_Reader_Trans_withReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_withReaderT')), "MonadAsk0" => function($dollar__unused_3) use ($monadAskReaderT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadAskReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadContReaderT'] = function() { $v = function($dictMonadCont_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadContReaderT"), recVars=[];
  $monadReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))((($dictMonadCont_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["callCC" => (function() use ($dictMonadCont_0) {
  $__fn = function($f_2, $r_3 = null) use ($dictMonadCont_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictMonadCont_0)->callCC)(function($c_4) use ($f_2, $r_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_2)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))($c_4))))($r_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadEffectReader'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadEffectReader"), recVars=[];
  $monadReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))((($dictMonadEffect_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["liftEffect" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))(($dictMonadEffect_0)->liftEffect), "Monad0" => function($dollar__unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadRecReaderT'] = function() { $v = function($dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadRecReaderT"), recVars=[];
  $Monad0_1_0 = (($dictMonadRec_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (($Monad0_1_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $pure_3_2 = ((($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure;
  $monadReaderT1_4_3 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))($Monad0_1_0);
  $__res = (object)["tailRecM" => (function() use ($__local_var_2_1, $dictMonadRec_0, $pure_3_2) {
  $__fn = function($k_5, $a_6 = null, $r_7 = null) use ($__local_var_2_1, $dictMonadRec_0, $pure_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictMonadRec_0)->tailRecM)(function($a__prime___8) use ($__local_var_2_1, $k_5, $pure_3_2, $r_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($__local_var_2_1)->bind)((($k_5)($a__prime___8))($r_7)))($pure_3_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_6);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_5) use ($monadReaderT1_4_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadStateReaderT'] = function() { $v = function($dictMonadState_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadStateReaderT"), recVars=[];
  $monadReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))((($dictMonadState_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["state" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))(($dictMonadState_0)->state), "Monad0" => function($dollar__unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadTellReaderT'] = function() { $v = function($dictMonadTell_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadTellReaderT"), recVars=[];
  $Semigroup0_1_0 = (($dictMonadTell_0)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadReaderT1_2_1 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))((($dictMonadTell_0)->Monad1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["tell" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))(($dictMonadTell_0)->tell), "Semigroup0" => function($dollar__unused_3) use ($Semigroup0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $Semigroup0_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($dollar__unused_3) use ($monadReaderT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadWriterReaderT'] = function() { $v = function($dictMonadWriter_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadWriterReaderT"), recVars=[];
  $Monoid0_1_0 = (($dictMonadWriter_0)->Monoid0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadTellReaderT1_2_1 = (($GLOBALS['Control_Monad_Reader_Trans_monadTellReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadTellReaderT')))((($dictMonadWriter_0)->MonadTell1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["listen" => (($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_mapReaderT')))(($dictMonadWriter_0)->listen), "pass" => (($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_mapReaderT')))(($dictMonadWriter_0)->pass), "Monoid0" => function($dollar__unused_3) use ($Monoid0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $Monoid0_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($dollar__unused_3) use ($monadTellReaderT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadTellReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadThrowReaderT'] = function() { $v = function($dictMonadThrow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadThrowReaderT"), recVars=[];
  $monadReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))((($dictMonadThrow_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["throwError" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))(($dictMonadThrow_0)->throwError), "Monad0" => function($dollar__unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadErrorReaderT'] = function() { $v = function($dictMonadError_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadErrorReaderT"), recVars=[];
  $monadThrowReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadThrowReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadThrowReaderT')))((($dictMonadError_0)->MonadThrow0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["catchError" => (function() use ($dictMonadError_0) {
  $__fn = function($v_2, $h_3 = null, $r_4 = null) use ($dictMonadError_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictMonadError_0)->catchError)(($v_2)($r_4)))(function($e_5) use ($h_3, $r_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($h_3)($e_5))($r_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($dollar__unused_2) use ($monadThrowReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadThrowReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadSTReaderT'] = function() { $v = function($dictMonadST_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadSTReaderT"), recVars=[];
  $monadReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))((($dictMonadST_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["liftST" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))(($dictMonadST_0)->liftST), "Monad0" => function($dollar__unused_2) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monoidReaderT'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monoidReaderT"), recVars=[];
  $semigroupReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_semigroupReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_semigroupReaderT')))((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictMonoid_2) use ($dictApplicative_0, $semigroupReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $semigroupReaderT2_3_1 = ($semigroupReaderT1_1_0)((($dictMonoid_2)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["mempty" => (((($GLOBALS['Control_Monad_Reader_Trans_applicativeReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_applicativeReaderT')))($dictApplicative_0))->pure)(($dictMonoid_2)->mempty), "Semigroup0" => function($dollar__unused_4) use ($semigroupReaderT2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $semigroupReaderT2_3_1;
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
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_altReaderT'] = function() { $v = function($dictAlt_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_altReaderT"), recVars=[];
  $functorReaderT1_1_0 = (object)["map" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_mapReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_mapReaderT'))))(((($dictAlt_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)];
  $__res = (object)["alt" => (function() use ($dictAlt_0) {
  $__fn = function($v_2, $v1_3 = null, $r_4 = null) use ($dictAlt_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictAlt_0)->alt)(($v_2)($r_4)))(($v1_3)($r_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_2) use ($functorReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $functorReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_plusReaderT'] = function() { $v = function($dictPlus_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_plusReaderT"), recVars=[];
  $altReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_altReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_altReaderT')))((($dictPlus_0)->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__local_var_2_1 = ($dictPlus_0)->empty;
  $__res = (object)["empty" => function($v_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $__local_var_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($dollar__unused_2) use ($altReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $altReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_alternativeReaderT'] = function() { $v = function($dictAlternative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_alternativeReaderT"), recVars=[];
  $applicativeReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_applicativeReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_applicativeReaderT')))((($dictAlternative_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $plusReaderT1_2_1 = (($GLOBALS['Control_Monad_Reader_Trans_plusReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_plusReaderT')))((($dictAlternative_0)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["Applicative0" => function($dollar__unused_3) use ($applicativeReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applicativeReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_3) use ($plusReaderT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $plusReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Reader_Trans_monadPlusReaderT'] = function() { $v = function($dictMonadPlus_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Reader_Trans_monadPlusReaderT"), recVars=[];
  $monadReaderT1_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadReaderT')))((($dictMonadPlus_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $alternativeReaderT1_2_1 = (($GLOBALS['Control_Monad_Reader_Trans_alternativeReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_alternativeReaderT')))((($dictMonadPlus_0)->Alternative1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["Monad0" => function($dollar__unused_3) use ($monadReaderT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadReaderT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($dollar__unused_3) use ($alternativeReaderT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $alternativeReaderT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };































