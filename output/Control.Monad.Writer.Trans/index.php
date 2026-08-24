<?php

namespace Control\Monad\Writer\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.Monad.Writer.Trans, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.Monad.Writer.Trans, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Trans/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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




// Control_Monad_Writer_Trans_WriterT
function majControl_majMonad_majWriter_majTrans_majWritermajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_majWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_WriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_majWritermajT';

// Control_Monad_Writer_Trans_runWriterT
function majControl_majMonad_majWriter_majTrans_runmajWritermajT($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_runmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_runWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_runmajWritermajT';

// Control_Monad_Writer_Trans_newtypeWriterT
$GLOBALS['Control_Monad_Writer_Trans_newtypeWriterT'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Writer_Trans_monadTransWriterT
function majControl_majMonad_majWriter_majTrans_monadmajTransmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajTransmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["lift" => function($dictMonad_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $pure_3_1 = ((($dictMonad_1)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_4) use ($Bind1_2_0, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($m_4))(function($a_5) use ($dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_5, ($dictMonoid_0)->{'mempty'}));
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
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_monadTransWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajTransmajWritermajT';

// Control_Monad_Writer_Trans_mapWriterT
function majControl_majMonad_majWriter_majTrans_mapmajWritermajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_mapmajWritermajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_mapWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_mapmajWritermajT';

// Control_Monad_Writer_Trans_functorWriterT
function majControl_majMonad_majWriter_majTrans_functormajWritermajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_functormajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (($dictFunctor_0)->{'map'})(function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_1)(($v_2)->{'value0'}), ($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = ($__local_var_2_0)($v_3);
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
$GLOBALS['Control_Monad_Writer_Trans_functorWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_functormajWritermajT';

// Control_Monad_Writer_Trans_execWriterT
function majControl_majMonad_majWriter_majTrans_execmajWritermajT($dictFunctor_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_execmajWritermajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)->{'map'})($GLOBALS['Data_Tuple_snd']))($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_execWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_execmajWritermajT';

// Control_Monad_Writer_Trans_applyWriterT
function majControl_majMonad_majWriter_majTrans_applymajWritermajT($dictSemigroup_0, $dictApply_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_applymajWritermajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Functor0_2_0 = (($dictApply_1)->{'Functor0'})(null);
  $__local_var_3_1 = (($dictApply_1)->{'Functor0'})(null);
  $functorWriterT1_3_1 = (object)["map" => function($f_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__local_var_5_2 = (($__local_var_3_1)->{'map'})(function($v_5) use ($f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_4)(($v_5)->{'value0'}), ($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_5_2)($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($v_4) use ($Functor0_2_0, $dictApply_1, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($Functor0_2_0, $dictApply_1, $dictSemigroup_0, $v_4) {
  $__num = \func_num_args();
  $__res = ((($dictApply_1)->{'apply'})(((($Functor0_2_0)->{'map'})(function($v3_6) use ($dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($v4_7) use ($dictSemigroup_0, $v3_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_6)->{'value0'})(($v4_7)->{'value0'}), ((($dictSemigroup_0)->{'append'})(($v3_6)->{'value1'}))(($v4_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_4)))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorWriterT1_3_1) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_applyWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_applymajWritermajT';

// Control_Monad_Writer_Trans_bindWriterT
function majControl_majMonad_majWriter_majTrans_bindmajWritermajT($dictSemigroup_0, $dictBind_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_bindmajWritermajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Apply0_2_0 = (($dictBind_1)->{'Apply0'})(null);
  $Functor0_3_1 = (($Apply0_2_0)->{'Functor0'})(null);
  $Functor0_4_2 = (($Apply0_2_0)->{'Functor0'})(null);
  $__local_var_5_3 = (($Apply0_2_0)->{'Functor0'})(null);
  $functorWriterT1_5_3 = (object)["map" => function($f_6) use ($__local_var_5_3) {
  $__num = \func_num_args();
  $__local_var_7_4 = (($__local_var_5_3)->{'map'})(function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v_7)->{'value0'}), ($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_4) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_4)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_4_2 = (object)["apply" => function($v_6) use ($Apply0_2_0, $Functor0_4_2, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($Apply0_2_0, $Functor0_4_2, $dictSemigroup_0, $v_6) {
  $__num = \func_num_args();
  $__res = ((($Apply0_2_0)->{'apply'})(((($Functor0_4_2)->{'map'})(function($v3_8) use ($dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($v4_9) use ($dictSemigroup_0, $v3_8) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_8)->{'value0'})(($v4_9)->{'value0'}), ((($dictSemigroup_0)->{'append'})(($v3_8)->{'value1'}))(($v4_9)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6)))($v1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorWriterT1_5_3) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_5) use ($Functor0_3_1, $dictBind_1, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($k_6) use ($Functor0_3_1, $dictBind_1, $dictSemigroup_0, $v_5) {
  $__num = \func_num_args();
  $__res = ((($dictBind_1)->{'bind'})($v_5))(function($v1_7) use ($Functor0_3_1, $dictSemigroup_0, $k_6) {
  $__num = \func_num_args();
  $__local_var_8_7 = ($v1_7)->{'value1'};
  $__res = ((($Functor0_3_1)->{'map'})(function($v3_9) use ($__local_var_8_7, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_9)->{'value0'}, ((($dictSemigroup_0)->{'append'})($__local_var_8_7))(($v3_9)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_6)(($v1_7)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_5) use ($applyWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_bindWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_bindmajWritermajT';

// Control_Monad_Writer_Trans_semigroupWriterT
function majControl_majMonad_majWriter_majTrans_semigroupmajWritermajT($dictApply_0, $dictSemigroup_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_semigroupmajWritermajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Functor0_2_0 = (($dictApply_0)->{'Functor0'})(null);
  $__local_var_3_1 = (($dictApply_0)->{'Functor0'})(null);
  $functorWriterT1_3_1 = (object)["map" => function($f_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__local_var_5_2 = (($__local_var_3_1)->{'map'})(function($v_5) use ($f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_4)(($v_5)->{'value0'}), ($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_6) use ($__local_var_5_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_5_2)($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT1_2_0 = (object)["apply" => function($v_4) use ($Functor0_2_0, $dictApply_0, $dictSemigroup_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($Functor0_2_0, $dictApply_0, $dictSemigroup_1, $v_4) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($Functor0_2_0)->{'map'})(function($v3_6) use ($dictSemigroup_1) {
  $__num = \func_num_args();
  $__res = function($v4_7) use ($dictSemigroup_1, $v3_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_6)->{'value0'})(($v4_7)->{'value0'}), ((($dictSemigroup_1)->{'append'})(($v3_6)->{'value1'}))(($v4_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_4)))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorWriterT1_3_1) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictSemigroup1_3) use ($applyWriterT1_2_0) {
  $__num = \func_num_args();
  $Functor0_4_5 = (($applyWriterT1_2_0)->{'Functor0'})(null);
  $__local_var_5_6 = ($dictSemigroup1_3)->{'append'};
  $__res = (object)["append" => function($a_6) use ($Functor0_4_5, $__local_var_5_6, $applyWriterT1_2_0) {
  $__num = \func_num_args();
  $__res = function($b_7) use ($Functor0_4_5, $__local_var_5_6, $a_6, $applyWriterT1_2_0) {
  $__num = \func_num_args();
  $__res = ((($applyWriterT1_2_0)->{'apply'})(((($Functor0_4_5)->{'map'})($__local_var_5_6))($a_6)))($b_7);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_semigroupWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_semigroupmajWritermajT';

// Control_Monad_Writer_Trans_applicativeWriterT
function majControl_majMonad_majWriter_majTrans_applicativemajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_applicativemajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictApplicative_2) use ($__local_var_1_0, $dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictApplicative_2)->{'Apply0'})(null);
  $Functor0_4_2 = (($__local_var_3_1)->{'Functor0'})(null);
  $__local_var_5_3 = (($__local_var_3_1)->{'Functor0'})(null);
  $functorWriterT1_5_3 = (object)["map" => function($f_6) use ($__local_var_5_3) {
  $__num = \func_num_args();
  $__local_var_7_4 = (($__local_var_5_3)->{'map'})(function($v_7) use ($f_6) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_6)(($v_7)->{'value0'}), ($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_8) use ($__local_var_7_4) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_4)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_3_1 = (object)["apply" => function($v_6) use ($Functor0_4_2, $__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($Functor0_4_2, $__local_var_1_0, $__local_var_3_1, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_1)->{'apply'})(((($Functor0_4_2)->{'map'})(function($v3_8) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v4_9) use ($__local_var_1_0, $v3_8) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_8)->{'value0'})(($v4_9)->{'value0'}), ((($__local_var_1_0)->{'append'})(($v3_8)->{'value1'}))(($v4_9)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6)))($v1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorWriterT1_5_3) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => function($a_4) use ($dictApplicative_2, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_2)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_4, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($applyWriterT2_3_1) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_3_1;
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
$GLOBALS['Control_Monad_Writer_Trans_applicativeWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_applicativemajWritermajT';

// Control_Monad_Writer_Trans_monadWriterT
function majControl_majMonad_majWriter_majTrans_monadmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonad_3) use ($__local_var_1_0, $__local_var_2_1, $dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictMonad_3)->{'Applicative0'})(null);
  $__local_var_5_3 = (($__local_var_4_2)->{'Apply0'})(null);
  $Functor0_6_4 = (($__local_var_5_3)->{'Functor0'})(null);
  $__local_var_7_5 = (($__local_var_5_3)->{'Functor0'})(null);
  $functorWriterT1_7_5 = (object)["map" => function($f_8) use ($__local_var_7_5) {
  $__num = \func_num_args();
  $__local_var_9_6 = (($__local_var_7_5)->{'map'})(function($v_9) use ($f_8) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_8)(($v_9)->{'value0'}), ($v_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_10) use ($__local_var_9_6) {
  $__num = \func_num_args();
  $__res = ($__local_var_9_6)($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_5_3 = (object)["apply" => function($v_8) use ($Functor0_6_4, $__local_var_1_0, $__local_var_5_3) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($Functor0_6_4, $__local_var_1_0, $__local_var_5_3, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_3)->{'apply'})(((($Functor0_6_4)->{'map'})(function($v3_10) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v4_11) use ($__local_var_1_0, $v3_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_10)->{'value0'})(($v4_11)->{'value0'}), ((($__local_var_1_0)->{'append'})(($v3_10)->{'value1'}))(($v4_11)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8)))($v1_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorWriterT1_7_5) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_4_2 = (object)["pure" => function($a_6) use ($__local_var_4_2, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_4_2)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_6, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_6) use ($applyWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_5_10 = (($dictMonad_3)->{'Bind1'})(null);
  $Apply0_6_11 = (($__local_var_5_10)->{'Apply0'})(null);
  $Functor0_7_12 = (($Apply0_6_11)->{'Functor0'})(null);
  $Functor0_8_13 = (($Apply0_6_11)->{'Functor0'})(null);
  $__local_var_9_14 = (($Apply0_6_11)->{'Functor0'})(null);
  $functorWriterT1_9_14 = (object)["map" => function($f_10) use ($__local_var_9_14) {
  $__num = \func_num_args();
  $__local_var_11_15 = (($__local_var_9_14)->{'map'})(function($v_11) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v_11)->{'value0'}), ($v_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_15) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_15)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_8_13 = (object)["apply" => function($v_10) use ($Apply0_6_11, $Functor0_8_13, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($Apply0_6_11, $Functor0_8_13, $__local_var_2_1, $v_10) {
  $__num = \func_num_args();
  $__res = ((($Apply0_6_11)->{'apply'})(((($Functor0_8_13)->{'map'})(function($v3_12) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v4_13) use ($__local_var_2_1, $v3_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_12)->{'value0'})(($v4_13)->{'value0'}), ((($__local_var_2_1)->{'append'})(($v3_12)->{'value1'}))(($v4_13)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10)))($v1_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_10) use ($functorWriterT1_9_14) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_9_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_5_10 = (object)["bind" => function($v_9) use ($Functor0_7_12, $__local_var_2_1, $__local_var_5_10) {
  $__num = \func_num_args();
  $__res = function($k_10) use ($Functor0_7_12, $__local_var_2_1, $__local_var_5_10, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_10)->{'bind'})($v_9))(function($v1_11) use ($Functor0_7_12, $__local_var_2_1, $k_10) {
  $__num = \func_num_args();
  $__local_var_12_18 = ($v1_11)->{'value1'};
  $__res = ((($Functor0_7_12)->{'map'})(function($v3_13) use ($__local_var_12_18, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_13)->{'value0'}, ((($__local_var_2_1)->{'append'})($__local_var_12_18))(($v3_13)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_10)(($v1_11)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_9) use ($applyWriterT2_8_13) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_8_13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_6) use ($applicativeWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_6) use ($bindWriterT2_5_10) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_5_10;
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
$GLOBALS['Control_Monad_Writer_Trans_monadWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajWritermajT';

// Control_Monad_Writer_Trans_monadAskWriterT
function majControl_majMonad_majWriter_majTrans_monadmajAskmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajAskmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadTransWriterT1_1_0 = (object)["lift" => function($dictMonad_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $pure_3_1 = ((($dictMonad_1)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_4) use ($Bind1_2_0, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($m_4))(function($a_5) use ($dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_5, ($dictMonoid_0)->{'mempty'}));
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
}];
  $__local_var_2_3 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_4 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonadAsk_4) use ($__local_var_2_3, $__local_var_3_4, $dictMonoid_0, $monadTransWriterT1_1_0) {
  $__num = \func_num_args();
  $__local_var_5_5 = (($dictMonadAsk_4)->{'Monad0'})(null);
  $__local_var_6_6 = (($__local_var_5_5)->{'Applicative0'})(null);
  $__local_var_7_7 = (($__local_var_6_6)->{'Apply0'})(null);
  $Functor0_8_8 = (($__local_var_7_7)->{'Functor0'})(null);
  $__local_var_9_9 = (($__local_var_7_7)->{'Functor0'})(null);
  $functorWriterT1_9_9 = (object)["map" => function($f_10) use ($__local_var_9_9) {
  $__num = \func_num_args();
  $__local_var_11_10 = (($__local_var_9_9)->{'map'})(function($v_11) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v_11)->{'value0'}), ($v_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_10) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_10)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_7_7 = (object)["apply" => function($v_10) use ($Functor0_8_8, $__local_var_2_3, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($Functor0_8_8, $__local_var_2_3, $__local_var_7_7, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'apply'})(((($Functor0_8_8)->{'map'})(function($v3_12) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($v4_13) use ($__local_var_2_3, $v3_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_12)->{'value0'})(($v4_13)->{'value0'}), ((($__local_var_2_3)->{'append'})(($v3_12)->{'value1'}))(($v4_13)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10)))($v1_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_10) use ($functorWriterT1_9_9) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_9_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_6_6 = (object)["pure" => function($a_8) use ($__local_var_6_6, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_6_6)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_8, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($applyWriterT2_7_7) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_7_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_7_14 = (($__local_var_5_5)->{'Bind1'})(null);
  $Apply0_8_15 = (($__local_var_7_14)->{'Apply0'})(null);
  $Functor0_9_16 = (($Apply0_8_15)->{'Functor0'})(null);
  $Functor0_10_17 = (($Apply0_8_15)->{'Functor0'})(null);
  $__local_var_11_18 = (($Apply0_8_15)->{'Functor0'})(null);
  $functorWriterT1_11_18 = (object)["map" => function($f_12) use ($__local_var_11_18) {
  $__num = \func_num_args();
  $__local_var_13_19 = (($__local_var_11_18)->{'map'})(function($v_13) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v_13)->{'value0'}), ($v_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_19) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_19)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_10_17 = (object)["apply" => function($v_12) use ($Apply0_8_15, $Functor0_10_17, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v1_13) use ($Apply0_8_15, $Functor0_10_17, $__local_var_3_4, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Apply0_8_15)->{'apply'})(((($Functor0_10_17)->{'map'})(function($v3_14) use ($__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v4_15) use ($__local_var_3_4, $v3_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_14)->{'value0'})(($v4_15)->{'value0'}), ((($__local_var_3_4)->{'append'})(($v3_14)->{'value1'}))(($v4_15)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12)))($v1_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_12) use ($functorWriterT1_11_18) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_11_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_7_14 = (object)["bind" => function($v_11) use ($Functor0_9_16, $__local_var_3_4, $__local_var_7_14) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Functor0_9_16, $__local_var_3_4, $__local_var_7_14, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_14)->{'bind'})($v_11))(function($v1_13) use ($Functor0_9_16, $__local_var_3_4, $k_12) {
  $__num = \func_num_args();
  $__local_var_14_22 = ($v1_13)->{'value1'};
  $__res = ((($Functor0_9_16)->{'map'})(function($v3_15) use ($__local_var_14_22, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_15)->{'value0'}, ((($__local_var_3_4)->{'append'})($__local_var_14_22))(($v3_15)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_12)(($v1_13)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_11) use ($applyWriterT2_10_17) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_10_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_5_5 = (object)["Applicative0" => function($_dollar___unused_8) use ($applicativeWriterT2_6_6) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_6_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_8) use ($bindWriterT2_7_14) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_7_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["ask" => ((($monadTransWriterT1_1_0)->{'lift'})((($dictMonadAsk_4)->{'Monad0'})(null)))(($dictMonadAsk_4)->{'ask'}), "Monad0" => function($_dollar___unused_6) use ($monadWriterT2_5_5) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_5;
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
$GLOBALS['Control_Monad_Writer_Trans_monadAskWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajAskmajWritermajT';

// Control_Monad_Writer_Trans_monadReaderWriterT
function majControl_majMonad_majWriter_majTrans_monadmajReadermajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajReadermajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadTransWriterT1_1_0 = (object)["lift" => function($dictMonad_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $pure_3_1 = ((($dictMonad_1)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_4) use ($Bind1_2_0, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($m_4))(function($a_5) use ($dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_5, ($dictMonoid_0)->{'mempty'}));
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
}];
  $__local_var_2_3 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_4 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $monadAskWriterT1_3_4 = function($dictMonadAsk_4) use ($__local_var_2_3, $__local_var_3_4, $dictMonoid_0, $monadTransWriterT1_1_0) {
  $__num = \func_num_args();
  $__local_var_5_5 = (($dictMonadAsk_4)->{'Monad0'})(null);
  $__local_var_6_6 = (($__local_var_5_5)->{'Applicative0'})(null);
  $__local_var_7_7 = (($__local_var_6_6)->{'Apply0'})(null);
  $Functor0_8_8 = (($__local_var_7_7)->{'Functor0'})(null);
  $__local_var_9_9 = (($__local_var_7_7)->{'Functor0'})(null);
  $functorWriterT1_9_9 = (object)["map" => function($f_10) use ($__local_var_9_9) {
  $__num = \func_num_args();
  $__local_var_11_10 = (($__local_var_9_9)->{'map'})(function($v_11) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v_11)->{'value0'}), ($v_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_10) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_10)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_7_7 = (object)["apply" => function($v_10) use ($Functor0_8_8, $__local_var_2_3, $__local_var_7_7) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($Functor0_8_8, $__local_var_2_3, $__local_var_7_7, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'apply'})(((($Functor0_8_8)->{'map'})(function($v3_12) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($v4_13) use ($__local_var_2_3, $v3_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_12)->{'value0'})(($v4_13)->{'value0'}), ((($__local_var_2_3)->{'append'})(($v3_12)->{'value1'}))(($v4_13)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10)))($v1_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_10) use ($functorWriterT1_9_9) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_9_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_6_6 = (object)["pure" => function($a_8) use ($__local_var_6_6, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_6_6)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_8, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($applyWriterT2_7_7) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_7_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_7_14 = (($__local_var_5_5)->{'Bind1'})(null);
  $Apply0_8_15 = (($__local_var_7_14)->{'Apply0'})(null);
  $Functor0_9_16 = (($Apply0_8_15)->{'Functor0'})(null);
  $Functor0_10_17 = (($Apply0_8_15)->{'Functor0'})(null);
  $__local_var_11_18 = (($Apply0_8_15)->{'Functor0'})(null);
  $functorWriterT1_11_18 = (object)["map" => function($f_12) use ($__local_var_11_18) {
  $__num = \func_num_args();
  $__local_var_13_19 = (($__local_var_11_18)->{'map'})(function($v_13) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v_13)->{'value0'}), ($v_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_19) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_19)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_10_17 = (object)["apply" => function($v_12) use ($Apply0_8_15, $Functor0_10_17, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v1_13) use ($Apply0_8_15, $Functor0_10_17, $__local_var_3_4, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Apply0_8_15)->{'apply'})(((($Functor0_10_17)->{'map'})(function($v3_14) use ($__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v4_15) use ($__local_var_3_4, $v3_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_14)->{'value0'})(($v4_15)->{'value0'}), ((($__local_var_3_4)->{'append'})(($v3_14)->{'value1'}))(($v4_15)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12)))($v1_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_12) use ($functorWriterT1_11_18) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_11_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_7_14 = (object)["bind" => function($v_11) use ($Functor0_9_16, $__local_var_3_4, $__local_var_7_14) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Functor0_9_16, $__local_var_3_4, $__local_var_7_14, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_14)->{'bind'})($v_11))(function($v1_13) use ($Functor0_9_16, $__local_var_3_4, $k_12) {
  $__num = \func_num_args();
  $__local_var_14_22 = ($v1_13)->{'value1'};
  $__res = ((($Functor0_9_16)->{'map'})(function($v3_15) use ($__local_var_14_22, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_15)->{'value0'}, ((($__local_var_3_4)->{'append'})($__local_var_14_22))(($v3_15)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_12)(($v1_13)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_11) use ($applyWriterT2_10_17) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_10_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_5_5 = (object)["Applicative0" => function($_dollar___unused_8) use ($applicativeWriterT2_6_6) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_6_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_8) use ($bindWriterT2_7_14) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_7_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["ask" => ((($monadTransWriterT1_1_0)->{'lift'})((($dictMonadAsk_4)->{'Monad0'})(null)))(($dictMonadAsk_4)->{'ask'}), "Monad0" => function($_dollar___unused_6) use ($monadWriterT2_5_5) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonadReader_4) use ($monadAskWriterT1_3_4) {
  $__num = \func_num_args();
  $monadAskWriterT2_5_26 = ($monadAskWriterT1_3_4)((($dictMonadReader_4)->{'MonadAsk0'})(null));
  $__res = (object)["local" => function($f_6) use ($dictMonadReader_4) {
  $__num = \func_num_args();
  $__local_var_7_27 = (($dictMonadReader_4)->{'local'})($f_6);
  $__res = function($v_8) use ($__local_var_7_27) {
  $__num = \func_num_args();
  $__res = ($__local_var_7_27)($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadAsk0" => function($_dollar___unused_6) use ($monadAskWriterT2_5_26) {
  $__num = \func_num_args();
  $__res = $monadAskWriterT2_5_26;
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
$GLOBALS['Control_Monad_Writer_Trans_monadReaderWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajReadermajWritermajT';

// Control_Monad_Writer_Trans_monadContWriterT
function majControl_majMonad_majWriter_majTrans_monadmajContmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajContmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonadCont_3) use ($__local_var_1_0, $__local_var_2_1, $dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictMonadCont_3)->{'Monad0'})(null);
  $__local_var_5_3 = (($__local_var_4_2)->{'Applicative0'})(null);
  $__local_var_6_4 = (($__local_var_5_3)->{'Apply0'})(null);
  $Functor0_7_5 = (($__local_var_6_4)->{'Functor0'})(null);
  $__local_var_8_6 = (($__local_var_6_4)->{'Functor0'})(null);
  $functorWriterT1_8_6 = (object)["map" => function($f_9) use ($__local_var_8_6) {
  $__num = \func_num_args();
  $__local_var_10_7 = (($__local_var_8_6)->{'map'})(function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_7) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_7)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_6_4 = (object)["apply" => function($v_9) use ($Functor0_7_5, $__local_var_1_0, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($v1_10) use ($Functor0_7_5, $__local_var_1_0, $__local_var_6_4, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_4)->{'apply'})(((($Functor0_7_5)->{'map'})(function($v3_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v4_12) use ($__local_var_1_0, $v3_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_11)->{'value0'})(($v4_12)->{'value0'}), ((($__local_var_1_0)->{'append'})(($v3_11)->{'value1'}))(($v4_12)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9)))($v1_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_9) use ($functorWriterT1_8_6) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_5_3 = (object)["pure" => function($a_7) use ($__local_var_5_3, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_5_3)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_7, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($applyWriterT2_6_4) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_11 = (($__local_var_4_2)->{'Bind1'})(null);
  $Apply0_7_12 = (($__local_var_6_11)->{'Apply0'})(null);
  $Functor0_8_13 = (($Apply0_7_12)->{'Functor0'})(null);
  $Functor0_9_14 = (($Apply0_7_12)->{'Functor0'})(null);
  $__local_var_10_15 = (($Apply0_7_12)->{'Functor0'})(null);
  $functorWriterT1_10_15 = (object)["map" => function($f_11) use ($__local_var_10_15) {
  $__num = \func_num_args();
  $__local_var_12_16 = (($__local_var_10_15)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_16) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_16)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_9_14 = (object)["apply" => function($v_11) use ($Apply0_7_12, $Functor0_9_14, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Apply0_7_12, $Functor0_9_14, $__local_var_2_1, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_12)->{'apply'})(((($Functor0_9_14)->{'map'})(function($v3_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_2_1, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_2_1)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_15) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_6_11 = (object)["bind" => function($v_10) use ($Functor0_8_13, $__local_var_2_1, $__local_var_6_11) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Functor0_8_13, $__local_var_2_1, $__local_var_6_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_11)->{'bind'})($v_10))(function($v1_12) use ($Functor0_8_13, $__local_var_2_1, $k_11) {
  $__num = \func_num_args();
  $__local_var_13_19 = ($v1_12)->{'value1'};
  $__res = ((($Functor0_8_13)->{'map'})(function($v3_14) use ($__local_var_13_19, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_14)->{'value0'}, ((($__local_var_2_1)->{'append'})($__local_var_13_19))(($v3_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_11)(($v1_12)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_10) use ($applyWriterT2_9_14) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_9_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_4_2 = (object)["Applicative0" => function($_dollar___unused_7) use ($applicativeWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_7) use ($bindWriterT2_6_11) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_6_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["callCC" => function($f_5) use ($dictMonadCont_3, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadCont_3)->{'callCC'})(function($c_6) use ($dictMonoid_0, $f_5) {
  $__num = \func_num_args();
  $__res = ($f_5)(function($a_7) use ($c_6, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = ($c_6)(new \Data\Tuple\Data_Tuple_Tuple($a_7, ($dictMonoid_0)->{'mempty'}));
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
}, "Monad0" => function($_dollar___unused_5) use ($monadWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_4_2;
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
$GLOBALS['Control_Monad_Writer_Trans_monadContWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajContmajWritermajT';

// Control_Monad_Writer_Trans_monadEffectWriter
function majControl_majMonad_majWriter_majTrans_monadmajEffectmajWriter($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajEffectmajWriter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonadEffect_3) use ($__local_var_1_0, $__local_var_2_1, $dictMonoid_0) {
  $__num = \func_num_args();
  $Monad0_4_2 = (($dictMonadEffect_3)->{'Monad0'})(null);
  $__local_var_5_3 = (($Monad0_4_2)->{'Applicative0'})(null);
  $__local_var_6_4 = (($__local_var_5_3)->{'Apply0'})(null);
  $Functor0_7_5 = (($__local_var_6_4)->{'Functor0'})(null);
  $__local_var_8_6 = (($__local_var_6_4)->{'Functor0'})(null);
  $functorWriterT1_8_6 = (object)["map" => function($f_9) use ($__local_var_8_6) {
  $__num = \func_num_args();
  $__local_var_10_7 = (($__local_var_8_6)->{'map'})(function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_7) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_7)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_6_4 = (object)["apply" => function($v_9) use ($Functor0_7_5, $__local_var_1_0, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($v1_10) use ($Functor0_7_5, $__local_var_1_0, $__local_var_6_4, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_4)->{'apply'})(((($Functor0_7_5)->{'map'})(function($v3_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v4_12) use ($__local_var_1_0, $v3_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_11)->{'value0'})(($v4_12)->{'value0'}), ((($__local_var_1_0)->{'append'})(($v3_11)->{'value1'}))(($v4_12)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9)))($v1_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_9) use ($functorWriterT1_8_6) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_5_3 = (object)["pure" => function($a_7) use ($__local_var_5_3, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_5_3)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_7, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($applyWriterT2_6_4) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_11 = (($Monad0_4_2)->{'Bind1'})(null);
  $Apply0_7_12 = (($__local_var_6_11)->{'Apply0'})(null);
  $Functor0_8_13 = (($Apply0_7_12)->{'Functor0'})(null);
  $Functor0_9_14 = (($Apply0_7_12)->{'Functor0'})(null);
  $__local_var_10_15 = (($Apply0_7_12)->{'Functor0'})(null);
  $functorWriterT1_10_15 = (object)["map" => function($f_11) use ($__local_var_10_15) {
  $__num = \func_num_args();
  $__local_var_12_16 = (($__local_var_10_15)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_16) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_16)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_9_14 = (object)["apply" => function($v_11) use ($Apply0_7_12, $Functor0_9_14, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Apply0_7_12, $Functor0_9_14, $__local_var_2_1, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_12)->{'apply'})(((($Functor0_9_14)->{'map'})(function($v3_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_2_1, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_2_1)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_15) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_6_11 = (object)["bind" => function($v_10) use ($Functor0_8_13, $__local_var_2_1, $__local_var_6_11) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Functor0_8_13, $__local_var_2_1, $__local_var_6_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_11)->{'bind'})($v_10))(function($v1_12) use ($Functor0_8_13, $__local_var_2_1, $k_11) {
  $__num = \func_num_args();
  $__local_var_13_19 = ($v1_12)->{'value1'};
  $__res = ((($Functor0_8_13)->{'map'})(function($v3_14) use ($__local_var_13_19, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_14)->{'value0'}, ((($__local_var_2_1)->{'append'})($__local_var_13_19))(($v3_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_11)(($v1_12)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_10) use ($applyWriterT2_9_14) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_9_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_5_3 = (object)["Applicative0" => function($_dollar___unused_7) use ($applicativeWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_7) use ($bindWriterT2_6_11) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_6_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_22 = (($Monad0_4_2)->{'Bind1'})(null);
  $pure_7_23 = ((($Monad0_4_2)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_8) use ($Bind1_6_22, $dictMonoid_0, $pure_7_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_22)->{'bind'})($m_8))(function($a_9) use ($dictMonoid_0, $pure_7_23) {
  $__num = \func_num_args();
  $__res = ($pure_7_23)(new \Data\Tuple\Data_Tuple_Tuple($a_9, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadEffect_3)->{'liftEffect'}), "Monad0" => function($_dollar___unused_6) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
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
$GLOBALS['Control_Monad_Writer_Trans_monadEffectWriter'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajEffectmajWriter';

// Control_Monad_Writer_Trans_monadRecWriterT
function majControl_majMonad_majWriter_majTrans_monadmajRecmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajRecmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonadRec_4) use ($Semigroup0_1_0, $__local_var_2_1, $__local_var_3_2, $dictMonoid_0) {
  $__num = \func_num_args();
  $Monad0_5_3 = (($dictMonadRec_4)->{'Monad0'})(null);
  $Bind1_6_4 = (($Monad0_5_3)->{'Bind1'})(null);
  $Applicative0_7_5 = (($Monad0_5_3)->{'Applicative0'})(null);
  $__local_var_8_6 = (($Monad0_5_3)->{'Applicative0'})(null);
  $__local_var_9_7 = (($__local_var_8_6)->{'Apply0'})(null);
  $Functor0_10_8 = (($__local_var_9_7)->{'Functor0'})(null);
  $__local_var_11_9 = (($__local_var_9_7)->{'Functor0'})(null);
  $functorWriterT1_11_9 = (object)["map" => function($f_12) use ($__local_var_11_9) {
  $__num = \func_num_args();
  $__local_var_13_10 = (($__local_var_11_9)->{'map'})(function($v_13) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v_13)->{'value0'}), ($v_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_10) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_10)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_9_7 = (object)["apply" => function($v_12) use ($Functor0_10_8, $__local_var_2_1, $__local_var_9_7) {
  $__num = \func_num_args();
  $__res = function($v1_13) use ($Functor0_10_8, $__local_var_2_1, $__local_var_9_7, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_7)->{'apply'})(((($Functor0_10_8)->{'map'})(function($v3_14) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v4_15) use ($__local_var_2_1, $v3_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_14)->{'value0'})(($v4_15)->{'value0'}), ((($__local_var_2_1)->{'append'})(($v3_14)->{'value1'}))(($v4_15)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12)))($v1_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_12) use ($functorWriterT1_11_9) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_11_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_8_6 = (object)["pure" => function($a_10) use ($__local_var_8_6, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_8_6)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_10, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_10) use ($applyWriterT2_9_7) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_9_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_9_14 = (($Monad0_5_3)->{'Bind1'})(null);
  $Apply0_10_15 = (($__local_var_9_14)->{'Apply0'})(null);
  $Functor0_11_16 = (($Apply0_10_15)->{'Functor0'})(null);
  $Functor0_12_17 = (($Apply0_10_15)->{'Functor0'})(null);
  $__local_var_13_18 = (($Apply0_10_15)->{'Functor0'})(null);
  $functorWriterT1_13_18 = (object)["map" => function($f_14) use ($__local_var_13_18) {
  $__num = \func_num_args();
  $__local_var_15_19 = (($__local_var_13_18)->{'map'})(function($v_15) use ($f_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_14)(($v_15)->{'value0'}), ($v_15)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_16) use ($__local_var_15_19) {
  $__num = \func_num_args();
  $__res = ($__local_var_15_19)($v_16);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_12_17 = (object)["apply" => function($v_14) use ($Apply0_10_15, $Functor0_12_17, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_15) use ($Apply0_10_15, $Functor0_12_17, $__local_var_3_2, $v_14) {
  $__num = \func_num_args();
  $__res = ((($Apply0_10_15)->{'apply'})(((($Functor0_12_17)->{'map'})(function($v3_16) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v4_17) use ($__local_var_3_2, $v3_16) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_16)->{'value0'})(($v4_17)->{'value0'}), ((($__local_var_3_2)->{'append'})(($v3_16)->{'value1'}))(($v4_17)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_14)))($v1_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_14) use ($functorWriterT1_13_18) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_13_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_9_14 = (object)["bind" => function($v_13) use ($Functor0_11_16, $__local_var_3_2, $__local_var_9_14) {
  $__num = \func_num_args();
  $__res = function($k_14) use ($Functor0_11_16, $__local_var_3_2, $__local_var_9_14, $v_13) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_14)->{'bind'})($v_13))(function($v1_15) use ($Functor0_11_16, $__local_var_3_2, $k_14) {
  $__num = \func_num_args();
  $__local_var_16_22 = ($v1_15)->{'value1'};
  $__res = ((($Functor0_11_16)->{'map'})(function($v3_17) use ($__local_var_16_22, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_17)->{'value0'}, ((($__local_var_3_2)->{'append'})($__local_var_16_22))(($v3_17)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_14)(($v1_15)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_13) use ($applyWriterT2_12_17) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_12_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_8_6 = (object)["Applicative0" => function($_dollar___unused_10) use ($applicativeWriterT2_8_6) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_10) use ($bindWriterT2_9_14) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_9_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["tailRecM" => function($f_9) use ($Applicative0_7_5, $Bind1_6_4, $Semigroup0_1_0, $dictMonadRec_4, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($Applicative0_7_5, $Bind1_6_4, $Semigroup0_1_0, $dictMonadRec_4, $dictMonoid_0, $f_9) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_4)->{'tailRecM'})(function($v_11) use ($Applicative0_7_5, $Bind1_6_4, $Semigroup0_1_0, $f_9) {
  $__num = \func_num_args();
  $__local_var_12_25 = ($v_11)->{'value1'};
  $__res = ((($Bind1_6_4)->{'bind'})(($f_9)(($v_11)->{'value0'})))(function($v2_13) use ($Applicative0_7_5, $Semigroup0_1_0, $__local_var_12_25) {
  $__num = \func_num_args();
  $__t26 = null;;
  if (($v2_13)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$__t26 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop(new \Data\Tuple\Data_Tuple_Tuple((($v2_13)->{'value0'})->{'value0'}, ((($Semigroup0_1_0)->{'append'})($__local_var_12_25))(($v2_13)->{'value1'})));
goto end_branch_26;;
};
  if (($v2_13)->{'value0'} instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t26 = new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(new \Data\Tuple\Data_Tuple_Tuple((($v2_13)->{'value0'})->{'value0'}, ((($Semigroup0_1_0)->{'append'})($__local_var_12_25))(($v2_13)->{'value1'})));
goto end_branch_26;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t26 = null;
  end_branch_26:;
  $__res = (($Applicative0_7_5)->{'pure'})($__t26);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple($a_10, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_9) use ($monadWriterT2_8_6) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_8_6;
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
$GLOBALS['Control_Monad_Writer_Trans_monadRecWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajRecmajWritermajT';

// Control_Monad_Writer_Trans_monadStateWriterT
function majControl_majMonad_majWriter_majTrans_monadmajStatemajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajStatemajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadTransWriterT1_1_0 = (object)["lift" => function($dictMonad_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $pure_3_1 = ((($dictMonad_1)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_4) use ($Bind1_2_0, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($m_4))(function($a_5) use ($dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_5, ($dictMonoid_0)->{'mempty'}));
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
}];
  $__local_var_2_3 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_4 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonadState_4) use ($__local_var_2_3, $__local_var_3_4, $dictMonoid_0, $monadTransWriterT1_1_0) {
  $__num = \func_num_args();
  $Monad0_5_5 = (($dictMonadState_4)->{'Monad0'})(null);
  $__local_var_6_6 = (($dictMonadState_4)->{'Monad0'})(null);
  $__local_var_7_7 = (($__local_var_6_6)->{'Applicative0'})(null);
  $__local_var_8_8 = (($__local_var_7_7)->{'Apply0'})(null);
  $Functor0_9_9 = (($__local_var_8_8)->{'Functor0'})(null);
  $__local_var_10_10 = (($__local_var_8_8)->{'Functor0'})(null);
  $functorWriterT1_10_10 = (object)["map" => function($f_11) use ($__local_var_10_10) {
  $__num = \func_num_args();
  $__local_var_12_11 = (($__local_var_10_10)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_11)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_8_8 = (object)["apply" => function($v_11) use ($Functor0_9_9, $__local_var_2_3, $__local_var_8_8) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Functor0_9_9, $__local_var_2_3, $__local_var_8_8, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_8)->{'apply'})(((($Functor0_9_9)->{'map'})(function($v3_13) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_2_3, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_2_3)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_10) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_7_7 = (object)["pure" => function($a_9) use ($__local_var_7_7, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_7_7)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_9, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($applyWriterT2_8_8) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_8_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_8_15 = (($__local_var_6_6)->{'Bind1'})(null);
  $Apply0_9_16 = (($__local_var_8_15)->{'Apply0'})(null);
  $Functor0_10_17 = (($Apply0_9_16)->{'Functor0'})(null);
  $Functor0_11_18 = (($Apply0_9_16)->{'Functor0'})(null);
  $__local_var_12_19 = (($Apply0_9_16)->{'Functor0'})(null);
  $functorWriterT1_12_19 = (object)["map" => function($f_13) use ($__local_var_12_19) {
  $__num = \func_num_args();
  $__local_var_14_20 = (($__local_var_12_19)->{'map'})(function($v_14) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_13)(($v_14)->{'value0'}), ($v_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_20) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_20)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_11_18 = (object)["apply" => function($v_13) use ($Apply0_9_16, $Functor0_11_18, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v1_14) use ($Apply0_9_16, $Functor0_11_18, $__local_var_3_4, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Apply0_9_16)->{'apply'})(((($Functor0_11_18)->{'map'})(function($v3_15) use ($__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v4_16) use ($__local_var_3_4, $v3_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_15)->{'value0'})(($v4_16)->{'value0'}), ((($__local_var_3_4)->{'append'})(($v3_15)->{'value1'}))(($v4_16)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_13)))($v1_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_13) use ($functorWriterT1_12_19) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_12_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_8_15 = (object)["bind" => function($v_12) use ($Functor0_10_17, $__local_var_3_4, $__local_var_8_15) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Functor0_10_17, $__local_var_3_4, $__local_var_8_15, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_15)->{'bind'})($v_12))(function($v1_14) use ($Functor0_10_17, $__local_var_3_4, $k_13) {
  $__num = \func_num_args();
  $__local_var_15_23 = ($v1_14)->{'value1'};
  $__res = ((($Functor0_10_17)->{'map'})(function($v3_16) use ($__local_var_15_23, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_16)->{'value0'}, ((($__local_var_3_4)->{'append'})($__local_var_15_23))(($v3_16)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_13)(($v1_14)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_12) use ($applyWriterT2_11_18) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_11_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_6_6 = (object)["Applicative0" => function($_dollar___unused_9) use ($applicativeWriterT2_7_7) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_7_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_9) use ($bindWriterT2_8_15) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_8_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["state" => function($f_7) use ($Monad0_5_5, $dictMonadState_4, $monadTransWriterT1_1_0) {
  $__num = \func_num_args();
  $__res = ((($monadTransWriterT1_1_0)->{'lift'})($Monad0_5_5))((($dictMonadState_4)->{'state'})($f_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_7) use ($monadWriterT2_6_6) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_6_6;
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
$GLOBALS['Control_Monad_Writer_Trans_monadStateWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajStatemajWritermajT';

// Control_Monad_Writer_Trans_monadTellWriterT
function majControl_majMonad_majWriter_majTrans_monadmajTellmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajTellmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonad_4) use ($Semigroup0_1_0, $__local_var_2_1, $__local_var_3_2, $dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_5_3 = (($dictMonad_4)->{'Applicative0'})(null);
  $__local_var_6_4 = (($__local_var_5_3)->{'Apply0'})(null);
  $Functor0_7_5 = (($__local_var_6_4)->{'Functor0'})(null);
  $__local_var_8_6 = (($__local_var_6_4)->{'Functor0'})(null);
  $functorWriterT1_8_6 = (object)["map" => function($f_9) use ($__local_var_8_6) {
  $__num = \func_num_args();
  $__local_var_10_7 = (($__local_var_8_6)->{'map'})(function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_7) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_7)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_6_4 = (object)["apply" => function($v_9) use ($Functor0_7_5, $__local_var_2_1, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($v1_10) use ($Functor0_7_5, $__local_var_2_1, $__local_var_6_4, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_4)->{'apply'})(((($Functor0_7_5)->{'map'})(function($v3_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v4_12) use ($__local_var_2_1, $v3_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_11)->{'value0'})(($v4_12)->{'value0'}), ((($__local_var_2_1)->{'append'})(($v3_11)->{'value1'}))(($v4_12)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9)))($v1_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_9) use ($functorWriterT1_8_6) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_5_3 = (object)["pure" => function($a_7) use ($__local_var_5_3, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_5_3)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_7, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($applyWriterT2_6_4) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_11 = (($dictMonad_4)->{'Bind1'})(null);
  $Apply0_7_12 = (($__local_var_6_11)->{'Apply0'})(null);
  $Functor0_8_13 = (($Apply0_7_12)->{'Functor0'})(null);
  $Functor0_9_14 = (($Apply0_7_12)->{'Functor0'})(null);
  $__local_var_10_15 = (($Apply0_7_12)->{'Functor0'})(null);
  $functorWriterT1_10_15 = (object)["map" => function($f_11) use ($__local_var_10_15) {
  $__num = \func_num_args();
  $__local_var_12_16 = (($__local_var_10_15)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_16) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_16)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_9_14 = (object)["apply" => function($v_11) use ($Apply0_7_12, $Functor0_9_14, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Apply0_7_12, $Functor0_9_14, $__local_var_3_2, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_12)->{'apply'})(((($Functor0_9_14)->{'map'})(function($v3_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_3_2, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_3_2)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_15) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_6_11 = (object)["bind" => function($v_10) use ($Functor0_8_13, $__local_var_3_2, $__local_var_6_11) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Functor0_8_13, $__local_var_3_2, $__local_var_6_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_11)->{'bind'})($v_10))(function($v1_12) use ($Functor0_8_13, $__local_var_3_2, $k_11) {
  $__num = \func_num_args();
  $__local_var_13_19 = ($v1_12)->{'value1'};
  $__res = ((($Functor0_8_13)->{'map'})(function($v3_14) use ($__local_var_13_19, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_14)->{'value0'}, ((($__local_var_3_2)->{'append'})($__local_var_13_19))(($v3_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_11)(($v1_12)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_10) use ($applyWriterT2_9_14) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_9_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_5_3 = (object)["Applicative0" => function($_dollar___unused_7) use ($applicativeWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_7) use ($bindWriterT2_6_11) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_6_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_4)->{'Applicative0'})(null))->{'pure'}))(($GLOBALS['Data_Tuple_Tuple'])($GLOBALS['Data_Unit_unit']))), "Semigroup0" => function($_dollar___unused_6) use ($Semigroup0_1_0) {
  $__num = \func_num_args();
  $__res = $Semigroup0_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_6) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
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
$GLOBALS['Control_Monad_Writer_Trans_monadTellWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajTellmajWritermajT';

// Control_Monad_Writer_Trans_monadWriterWriterT
function majControl_majMonad_majWriter_majTrans_monadmajWritermajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajWritermajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $monadTellWriterT1_3_2 = function($dictMonad_4) use ($Semigroup0_1_0, $__local_var_2_1, $__local_var_3_2, $dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_5_3 = (($dictMonad_4)->{'Applicative0'})(null);
  $__local_var_6_4 = (($__local_var_5_3)->{'Apply0'})(null);
  $Functor0_7_5 = (($__local_var_6_4)->{'Functor0'})(null);
  $__local_var_8_6 = (($__local_var_6_4)->{'Functor0'})(null);
  $functorWriterT1_8_6 = (object)["map" => function($f_9) use ($__local_var_8_6) {
  $__num = \func_num_args();
  $__local_var_10_7 = (($__local_var_8_6)->{'map'})(function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_7) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_7)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_6_4 = (object)["apply" => function($v_9) use ($Functor0_7_5, $__local_var_2_1, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($v1_10) use ($Functor0_7_5, $__local_var_2_1, $__local_var_6_4, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_4)->{'apply'})(((($Functor0_7_5)->{'map'})(function($v3_11) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v4_12) use ($__local_var_2_1, $v3_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_11)->{'value0'})(($v4_12)->{'value0'}), ((($__local_var_2_1)->{'append'})(($v3_11)->{'value1'}))(($v4_12)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9)))($v1_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_9) use ($functorWriterT1_8_6) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_5_3 = (object)["pure" => function($a_7) use ($__local_var_5_3, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_5_3)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_7, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($applyWriterT2_6_4) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_11 = (($dictMonad_4)->{'Bind1'})(null);
  $Apply0_7_12 = (($__local_var_6_11)->{'Apply0'})(null);
  $Functor0_8_13 = (($Apply0_7_12)->{'Functor0'})(null);
  $Functor0_9_14 = (($Apply0_7_12)->{'Functor0'})(null);
  $__local_var_10_15 = (($Apply0_7_12)->{'Functor0'})(null);
  $functorWriterT1_10_15 = (object)["map" => function($f_11) use ($__local_var_10_15) {
  $__num = \func_num_args();
  $__local_var_12_16 = (($__local_var_10_15)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_16) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_16)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_9_14 = (object)["apply" => function($v_11) use ($Apply0_7_12, $Functor0_9_14, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Apply0_7_12, $Functor0_9_14, $__local_var_3_2, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_12)->{'apply'})(((($Functor0_9_14)->{'map'})(function($v3_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_3_2, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_3_2)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_15) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_6_11 = (object)["bind" => function($v_10) use ($Functor0_8_13, $__local_var_3_2, $__local_var_6_11) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Functor0_8_13, $__local_var_3_2, $__local_var_6_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_11)->{'bind'})($v_10))(function($v1_12) use ($Functor0_8_13, $__local_var_3_2, $k_11) {
  $__num = \func_num_args();
  $__local_var_13_19 = ($v1_12)->{'value1'};
  $__res = ((($Functor0_8_13)->{'map'})(function($v3_14) use ($__local_var_13_19, $__local_var_3_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_14)->{'value0'}, ((($__local_var_3_2)->{'append'})($__local_var_13_19))(($v3_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_11)(($v1_12)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_10) use ($applyWriterT2_9_14) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_9_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_5_3 = (object)["Applicative0" => function($_dollar___unused_7) use ($applicativeWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_7) use ($bindWriterT2_6_11) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_6_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_4)->{'Applicative0'})(null))->{'pure'}))(($GLOBALS['Data_Tuple_Tuple'])($GLOBALS['Data_Unit_unit']))), "Semigroup0" => function($_dollar___unused_6) use ($Semigroup0_1_0) {
  $__num = \func_num_args();
  $__res = $Semigroup0_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar___unused_6) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonad_4) use ($dictMonoid_0, $monadTellWriterT1_3_2) {
  $__num = \func_num_args();
  $Bind1_5_23 = (($dictMonad_4)->{'Bind1'})(null);
  $Applicative0_6_24 = (($dictMonad_4)->{'Applicative0'})(null);
  $monadTellWriterT2_7_25 = ($monadTellWriterT1_3_2)($dictMonad_4);
  $__res = (object)["listen" => function($v_8) use ($Applicative0_6_24, $Bind1_5_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_23)->{'bind'})($v_8))(function($v1_9) use ($Applicative0_6_24) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_24)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple(new \Data\Tuple\Data_Tuple_Tuple(($v1_9)->{'value0'}, ($v1_9)->{'value1'}), ($v1_9)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pass" => function($v_8) use ($Applicative0_6_24, $Bind1_5_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_5_23)->{'bind'})($v_8))(function($v1_9) use ($Applicative0_6_24) {
  $__num = \func_num_args();
  $__res = (($Applicative0_6_24)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple((($v1_9)->{'value0'})->{'value0'}, ((($v1_9)->{'value0'})->{'value1'})(($v1_9)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monoid0" => function($_dollar___unused_8) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = $dictMonoid_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($_dollar___unused_8) use ($monadTellWriterT2_7_25) {
  $__num = \func_num_args();
  $__res = $monadTellWriterT2_7_25;
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
$GLOBALS['Control_Monad_Writer_Trans_monadWriterWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajWritermajWritermajT';

// Control_Monad_Writer_Trans_monadThrowWriterT
function majControl_majMonad_majWriter_majTrans_monadmajThrowmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajThrowmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadTransWriterT1_1_0 = (object)["lift" => function($dictMonad_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $pure_3_1 = ((($dictMonad_1)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_4) use ($Bind1_2_0, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($m_4))(function($a_5) use ($dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_5, ($dictMonoid_0)->{'mempty'}));
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
}];
  $__local_var_2_3 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_4 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonadThrow_4) use ($__local_var_2_3, $__local_var_3_4, $dictMonoid_0, $monadTransWriterT1_1_0) {
  $__num = \func_num_args();
  $Monad0_5_5 = (($dictMonadThrow_4)->{'Monad0'})(null);
  $__local_var_6_6 = (($dictMonadThrow_4)->{'Monad0'})(null);
  $__local_var_7_7 = (($__local_var_6_6)->{'Applicative0'})(null);
  $__local_var_8_8 = (($__local_var_7_7)->{'Apply0'})(null);
  $Functor0_9_9 = (($__local_var_8_8)->{'Functor0'})(null);
  $__local_var_10_10 = (($__local_var_8_8)->{'Functor0'})(null);
  $functorWriterT1_10_10 = (object)["map" => function($f_11) use ($__local_var_10_10) {
  $__num = \func_num_args();
  $__local_var_12_11 = (($__local_var_10_10)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_11)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_8_8 = (object)["apply" => function($v_11) use ($Functor0_9_9, $__local_var_2_3, $__local_var_8_8) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Functor0_9_9, $__local_var_2_3, $__local_var_8_8, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_8)->{'apply'})(((($Functor0_9_9)->{'map'})(function($v3_13) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_2_3, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_2_3)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_10) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_7_7 = (object)["pure" => function($a_9) use ($__local_var_7_7, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_7_7)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_9, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($applyWriterT2_8_8) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_8_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_8_15 = (($__local_var_6_6)->{'Bind1'})(null);
  $Apply0_9_16 = (($__local_var_8_15)->{'Apply0'})(null);
  $Functor0_10_17 = (($Apply0_9_16)->{'Functor0'})(null);
  $Functor0_11_18 = (($Apply0_9_16)->{'Functor0'})(null);
  $__local_var_12_19 = (($Apply0_9_16)->{'Functor0'})(null);
  $functorWriterT1_12_19 = (object)["map" => function($f_13) use ($__local_var_12_19) {
  $__num = \func_num_args();
  $__local_var_14_20 = (($__local_var_12_19)->{'map'})(function($v_14) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_13)(($v_14)->{'value0'}), ($v_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_20) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_20)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_11_18 = (object)["apply" => function($v_13) use ($Apply0_9_16, $Functor0_11_18, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v1_14) use ($Apply0_9_16, $Functor0_11_18, $__local_var_3_4, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Apply0_9_16)->{'apply'})(((($Functor0_11_18)->{'map'})(function($v3_15) use ($__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v4_16) use ($__local_var_3_4, $v3_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_15)->{'value0'})(($v4_16)->{'value0'}), ((($__local_var_3_4)->{'append'})(($v3_15)->{'value1'}))(($v4_16)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_13)))($v1_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_13) use ($functorWriterT1_12_19) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_12_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_8_15 = (object)["bind" => function($v_12) use ($Functor0_10_17, $__local_var_3_4, $__local_var_8_15) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Functor0_10_17, $__local_var_3_4, $__local_var_8_15, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_15)->{'bind'})($v_12))(function($v1_14) use ($Functor0_10_17, $__local_var_3_4, $k_13) {
  $__num = \func_num_args();
  $__local_var_15_23 = ($v1_14)->{'value1'};
  $__res = ((($Functor0_10_17)->{'map'})(function($v3_16) use ($__local_var_15_23, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_16)->{'value0'}, ((($__local_var_3_4)->{'append'})($__local_var_15_23))(($v3_16)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_13)(($v1_14)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_12) use ($applyWriterT2_11_18) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_11_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_6_6 = (object)["Applicative0" => function($_dollar___unused_9) use ($applicativeWriterT2_7_7) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_7_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_9) use ($bindWriterT2_8_15) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_8_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["throwError" => function($e_7) use ($Monad0_5_5, $dictMonadThrow_4, $monadTransWriterT1_1_0) {
  $__num = \func_num_args();
  $__res = ((($monadTransWriterT1_1_0)->{'lift'})($Monad0_5_5))((($dictMonadThrow_4)->{'throwError'})($e_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_7) use ($monadWriterT2_6_6) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_6_6;
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
$GLOBALS['Control_Monad_Writer_Trans_monadThrowWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajThrowmajWritermajT';

// Control_Monad_Writer_Trans_monadErrorWriterT
function majControl_majMonad_majWriter_majTrans_monadmajErrormajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajErrormajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $monadTransWriterT1_1_0 = (object)["lift" => function($dictMonad_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $pure_3_1 = ((($dictMonad_1)->{'Applicative0'})(null))->{'pure'};
  $__res = function($m_4) use ($Bind1_2_0, $dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($m_4))(function($a_5) use ($dictMonoid_0, $pure_3_1) {
  $__num = \func_num_args();
  $__res = ($pure_3_1)(new \Data\Tuple\Data_Tuple_Tuple($a_5, ($dictMonoid_0)->{'mempty'}));
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
}];
  $__local_var_2_3 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_4 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $monadThrowWriterT1_3_4 = function($dictMonadThrow_4) use ($__local_var_2_3, $__local_var_3_4, $dictMonoid_0, $monadTransWriterT1_1_0) {
  $__num = \func_num_args();
  $Monad0_5_5 = (($dictMonadThrow_4)->{'Monad0'})(null);
  $__local_var_6_6 = (($dictMonadThrow_4)->{'Monad0'})(null);
  $__local_var_7_7 = (($__local_var_6_6)->{'Applicative0'})(null);
  $__local_var_8_8 = (($__local_var_7_7)->{'Apply0'})(null);
  $Functor0_9_9 = (($__local_var_8_8)->{'Functor0'})(null);
  $__local_var_10_10 = (($__local_var_8_8)->{'Functor0'})(null);
  $functorWriterT1_10_10 = (object)["map" => function($f_11) use ($__local_var_10_10) {
  $__num = \func_num_args();
  $__local_var_12_11 = (($__local_var_10_10)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_11) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_11)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_8_8 = (object)["apply" => function($v_11) use ($Functor0_9_9, $__local_var_2_3, $__local_var_8_8) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Functor0_9_9, $__local_var_2_3, $__local_var_8_8, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_8)->{'apply'})(((($Functor0_9_9)->{'map'})(function($v3_13) use ($__local_var_2_3) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_2_3, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_2_3)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_10) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_7_7 = (object)["pure" => function($a_9) use ($__local_var_7_7, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_7_7)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_9, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($applyWriterT2_8_8) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_8_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_8_15 = (($__local_var_6_6)->{'Bind1'})(null);
  $Apply0_9_16 = (($__local_var_8_15)->{'Apply0'})(null);
  $Functor0_10_17 = (($Apply0_9_16)->{'Functor0'})(null);
  $Functor0_11_18 = (($Apply0_9_16)->{'Functor0'})(null);
  $__local_var_12_19 = (($Apply0_9_16)->{'Functor0'})(null);
  $functorWriterT1_12_19 = (object)["map" => function($f_13) use ($__local_var_12_19) {
  $__num = \func_num_args();
  $__local_var_14_20 = (($__local_var_12_19)->{'map'})(function($v_14) use ($f_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_13)(($v_14)->{'value0'}), ($v_14)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_15) use ($__local_var_14_20) {
  $__num = \func_num_args();
  $__res = ($__local_var_14_20)($v_15);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_11_18 = (object)["apply" => function($v_13) use ($Apply0_9_16, $Functor0_11_18, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v1_14) use ($Apply0_9_16, $Functor0_11_18, $__local_var_3_4, $v_13) {
  $__num = \func_num_args();
  $__res = ((($Apply0_9_16)->{'apply'})(((($Functor0_11_18)->{'map'})(function($v3_15) use ($__local_var_3_4) {
  $__num = \func_num_args();
  $__res = function($v4_16) use ($__local_var_3_4, $v3_15) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_15)->{'value0'})(($v4_16)->{'value0'}), ((($__local_var_3_4)->{'append'})(($v3_15)->{'value1'}))(($v4_16)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_13)))($v1_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_13) use ($functorWriterT1_12_19) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_12_19;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_8_15 = (object)["bind" => function($v_12) use ($Functor0_10_17, $__local_var_3_4, $__local_var_8_15) {
  $__num = \func_num_args();
  $__res = function($k_13) use ($Functor0_10_17, $__local_var_3_4, $__local_var_8_15, $v_12) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_15)->{'bind'})($v_12))(function($v1_14) use ($Functor0_10_17, $__local_var_3_4, $k_13) {
  $__num = \func_num_args();
  $__local_var_15_23 = ($v1_14)->{'value1'};
  $__res = ((($Functor0_10_17)->{'map'})(function($v3_16) use ($__local_var_15_23, $__local_var_3_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_16)->{'value0'}, ((($__local_var_3_4)->{'append'})($__local_var_15_23))(($v3_16)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_13)(($v1_14)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_12) use ($applyWriterT2_11_18) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_11_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_6_6 = (object)["Applicative0" => function($_dollar___unused_9) use ($applicativeWriterT2_7_7) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_7_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_9) use ($bindWriterT2_8_15) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_8_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["throwError" => function($e_7) use ($Monad0_5_5, $dictMonadThrow_4, $monadTransWriterT1_1_0) {
  $__num = \func_num_args();
  $__res = ((($monadTransWriterT1_1_0)->{'lift'})($Monad0_5_5))((($dictMonadThrow_4)->{'throwError'})($e_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_7) use ($monadWriterT2_6_6) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_6_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonadError_4) use ($monadThrowWriterT1_3_4) {
  $__num = \func_num_args();
  $monadThrowWriterT2_5_27 = ($monadThrowWriterT1_3_4)((($dictMonadError_4)->{'MonadThrow0'})(null));
  $__res = (object)["catchError" => function($v_6) use ($dictMonadError_4) {
  $__num = \func_num_args();
  $__res = function($h_7) use ($dictMonadError_4, $v_6) {
  $__num = \func_num_args();
  $__res = ((($dictMonadError_4)->{'catchError'})($v_6))(function($e_8) use ($h_7) {
  $__num = \func_num_args();
  $__res = ($h_7)($e_8);
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
}, "MonadThrow0" => function($_dollar___unused_6) use ($monadThrowWriterT2_5_27) {
  $__num = \func_num_args();
  $__res = $monadThrowWriterT2_5_27;
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
$GLOBALS['Control_Monad_Writer_Trans_monadErrorWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajErrormajWritermajT';

// Control_Monad_Writer_Trans_monadSTWriterT
function majControl_majMonad_majWriter_majTrans_monadmajSmajTmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajSmajTmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonadST_3) use ($__local_var_1_0, $__local_var_2_1, $dictMonoid_0) {
  $__num = \func_num_args();
  $Monad0_4_2 = (($dictMonadST_3)->{'Monad0'})(null);
  $__local_var_5_3 = (($Monad0_4_2)->{'Applicative0'})(null);
  $__local_var_6_4 = (($__local_var_5_3)->{'Apply0'})(null);
  $Functor0_7_5 = (($__local_var_6_4)->{'Functor0'})(null);
  $__local_var_8_6 = (($__local_var_6_4)->{'Functor0'})(null);
  $functorWriterT1_8_6 = (object)["map" => function($f_9) use ($__local_var_8_6) {
  $__num = \func_num_args();
  $__local_var_10_7 = (($__local_var_8_6)->{'map'})(function($v_10) use ($f_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_9)(($v_10)->{'value0'}), ($v_10)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_11) use ($__local_var_10_7) {
  $__num = \func_num_args();
  $__res = ($__local_var_10_7)($v_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_6_4 = (object)["apply" => function($v_9) use ($Functor0_7_5, $__local_var_1_0, $__local_var_6_4) {
  $__num = \func_num_args();
  $__res = function($v1_10) use ($Functor0_7_5, $__local_var_1_0, $__local_var_6_4, $v_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_4)->{'apply'})(((($Functor0_7_5)->{'map'})(function($v3_11) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v4_12) use ($__local_var_1_0, $v3_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_11)->{'value0'})(($v4_12)->{'value0'}), ((($__local_var_1_0)->{'append'})(($v3_11)->{'value1'}))(($v4_12)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9)))($v1_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_9) use ($functorWriterT1_8_6) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_5_3 = (object)["pure" => function($a_7) use ($__local_var_5_3, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_5_3)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_7, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_7) use ($applyWriterT2_6_4) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_11 = (($Monad0_4_2)->{'Bind1'})(null);
  $Apply0_7_12 = (($__local_var_6_11)->{'Apply0'})(null);
  $Functor0_8_13 = (($Apply0_7_12)->{'Functor0'})(null);
  $Functor0_9_14 = (($Apply0_7_12)->{'Functor0'})(null);
  $__local_var_10_15 = (($Apply0_7_12)->{'Functor0'})(null);
  $functorWriterT1_10_15 = (object)["map" => function($f_11) use ($__local_var_10_15) {
  $__num = \func_num_args();
  $__local_var_12_16 = (($__local_var_10_15)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_16) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_16)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_9_14 = (object)["apply" => function($v_11) use ($Apply0_7_12, $Functor0_9_14, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Apply0_7_12, $Functor0_9_14, $__local_var_2_1, $v_11) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_12)->{'apply'})(((($Functor0_9_14)->{'map'})(function($v3_13) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_2_1, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_2_1)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_15) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_6_11 = (object)["bind" => function($v_10) use ($Functor0_8_13, $__local_var_2_1, $__local_var_6_11) {
  $__num = \func_num_args();
  $__res = function($k_11) use ($Functor0_8_13, $__local_var_2_1, $__local_var_6_11, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_11)->{'bind'})($v_10))(function($v1_12) use ($Functor0_8_13, $__local_var_2_1, $k_11) {
  $__num = \func_num_args();
  $__local_var_13_19 = ($v1_12)->{'value1'};
  $__res = ((($Functor0_8_13)->{'map'})(function($v3_14) use ($__local_var_13_19, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_14)->{'value0'}, ((($__local_var_2_1)->{'append'})($__local_var_13_19))(($v3_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_11)(($v1_12)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_10) use ($applyWriterT2_9_14) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_9_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_5_3 = (object)["Applicative0" => function($_dollar___unused_7) use ($applicativeWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_7) use ($bindWriterT2_6_11) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_6_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Bind1_6_22 = (($Monad0_4_2)->{'Bind1'})(null);
  $pure_7_23 = ((($Monad0_4_2)->{'Applicative0'})(null))->{'pure'};
  $__res = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_8) use ($Bind1_6_22, $dictMonoid_0, $pure_7_23) {
  $__num = \func_num_args();
  $__res = ((($Bind1_6_22)->{'bind'})($m_8))(function($a_9) use ($dictMonoid_0, $pure_7_23) {
  $__num = \func_num_args();
  $__res = ($pure_7_23)(new \Data\Tuple\Data_Tuple_Tuple($a_9, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadST_3)->{'liftST'}), "Monad0" => function($_dollar___unused_6) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
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
$GLOBALS['Control_Monad_Writer_Trans_monadSTWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajSmajTmajWritermajT';

// Control_Monad_Writer_Trans_monoidWriterT
function majControl_majMonad_majWriter_majTrans_monoidmajWritermajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monoidmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($dictMonoid_2) use ($__local_var_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $__local_var_4_2 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_5_3 = (($__local_var_4_2)->{'Functor0'})(null);
  $__local_var_6_4 = (($__local_var_4_2)->{'Functor0'})(null);
  $functorWriterT1_6_4 = (object)["map" => function($f_7) use ($__local_var_6_4) {
  $__num = \func_num_args();
  $__local_var_8_5 = (($__local_var_6_4)->{'map'})(function($v_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_5)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_4_2 = (object)["apply" => function($v_7) use ($Functor0_5_3, $__local_var_3_1, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($Functor0_5_3, $__local_var_3_1, $__local_var_4_2, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_2)->{'apply'})(((($Functor0_5_3)->{'map'})(function($v3_9) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($v4_10) use ($__local_var_3_1, $v3_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_9)->{'value0'})(($v4_10)->{'value0'}), ((($__local_var_3_1)->{'append'})(($v3_9)->{'value1'}))(($v4_10)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7)))($v1_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorWriterT1_6_4) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT1_3_1 = (object)["pure" => function($a_5) use ($dictApplicative_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_5, ($dictMonoid_2)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($applyWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_9 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $Functor0_5_10 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_6_11 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorWriterT1_6_11 = (object)["map" => function($f_7) use ($__local_var_6_11) {
  $__num = \func_num_args();
  $__local_var_8_12 = (($__local_var_6_11)->{'map'})(function($v_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_12) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_12)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT1_5_10 = (object)["apply" => function($v_7) use ($Functor0_5_10, $__local_var_1_0, $__local_var_4_9) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($Functor0_5_10, $__local_var_1_0, $__local_var_4_9, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(((($Functor0_5_10)->{'map'})(function($v3_9) use ($__local_var_4_9) {
  $__num = \func_num_args();
  $__res = function($v4_10) use ($__local_var_4_9, $v3_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_9)->{'value0'})(($v4_10)->{'value0'}), ((($__local_var_4_9)->{'append'})(($v3_9)->{'value1'}))(($v4_10)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7)))($v1_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorWriterT1_6_11) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_6_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid1_6) use ($applicativeWriterT1_3_1, $applyWriterT1_5_10) {
  $__num = \func_num_args();
  $Functor0_7_15 = (($applyWriterT1_5_10)->{'Functor0'})(null);
  $__local_var_8_16 = ((($dictMonoid1_6)->{'Semigroup0'})(null))->{'append'};
  $semigroupWriterT3_7_15 = (object)["append" => function($a_9) use ($Functor0_7_15, $__local_var_8_16, $applyWriterT1_5_10) {
  $__num = \func_num_args();
  $__res = function($b_10) use ($Functor0_7_15, $__local_var_8_16, $a_9, $applyWriterT1_5_10) {
  $__num = \func_num_args();
  $__res = ((($applyWriterT1_5_10)->{'apply'})(((($Functor0_7_15)->{'map'})($__local_var_8_16))($a_9)))($b_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => (($applicativeWriterT1_3_1)->{'pure'})(($dictMonoid1_6)->{'mempty'}), "Semigroup0" => function($_dollar___unused_8) use ($semigroupWriterT3_7_15) {
  $__num = \func_num_args();
  $__res = $semigroupWriterT3_7_15;
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
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_monoidWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monoidmajWritermajT';

// Control_Monad_Writer_Trans_altWriterT
function majControl_majMonad_majWriter_majTrans_altmajWritermajT($dictAlt_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_altmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictAlt_0)->{'Functor0'})(null);
  $functorWriterT1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($__local_var_1_0)->{'map'})(function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_2)(($v_3)->{'value0'}), ($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = ($__local_var_3_1)($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["alt" => function($v_2) use ($dictAlt_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictAlt_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictAlt_0)->{'alt'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorWriterT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_altWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_altmajWritermajT';

// Control_Monad_Writer_Trans_plusWriterT
function majControl_majMonad_majWriter_majTrans_plusmajWritermajT($dictPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_plusmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictPlus_0)->{'Alt0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorWriterT1_2_1 = (object)["map" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($__local_var_2_1)->{'map'})(function($v_4) use ($f_3) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_3)(($v_4)->{'value0'}), ($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_5) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = ($__local_var_4_2)($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altWriterT1_1_0 = (object)["alt" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'alt'})($v_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorWriterT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => ($dictPlus_0)->{'empty'}, "Alt0" => function($_dollar___unused_2) use ($altWriterT1_1_0) {
  $__num = \func_num_args();
  $__res = $altWriterT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Writer_Trans_plusWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_plusmajWritermajT';

// Control_Monad_Writer_Trans_alternativeWriterT
function majControl_majMonad_majWriter_majTrans_alternativemajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_alternativemajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictAlternative_2) use ($__local_var_1_0, $dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictAlternative_2)->{'Applicative0'})(null);
  $__local_var_4_2 = (($__local_var_3_1)->{'Apply0'})(null);
  $Functor0_5_3 = (($__local_var_4_2)->{'Functor0'})(null);
  $__local_var_6_4 = (($__local_var_4_2)->{'Functor0'})(null);
  $functorWriterT1_6_4 = (object)["map" => function($f_7) use ($__local_var_6_4) {
  $__num = \func_num_args();
  $__local_var_8_5 = (($__local_var_6_4)->{'map'})(function($v_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_5) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_5)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_4_2 = (object)["apply" => function($v_7) use ($Functor0_5_3, $__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($Functor0_5_3, $__local_var_1_0, $__local_var_4_2, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_2)->{'apply'})(((($Functor0_5_3)->{'map'})(function($v3_9) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v4_10) use ($__local_var_1_0, $v3_9) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_9)->{'value0'})(($v4_10)->{'value0'}), ((($__local_var_1_0)->{'append'})(($v3_9)->{'value1'}))(($v4_10)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7)))($v1_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorWriterT1_6_4) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_3_1 = (object)["pure" => function($a_5) use ($__local_var_3_1, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_3_1)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_5, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($applyWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_9 = (($dictAlternative_2)->{'Plus1'})(null);
  $__local_var_5_10 = (($__local_var_4_9)->{'Alt0'})(null);
  $__local_var_6_11 = (($__local_var_5_10)->{'Functor0'})(null);
  $functorWriterT1_6_11 = (object)["map" => function($f_7) use ($__local_var_6_11) {
  $__num = \func_num_args();
  $__local_var_8_12 = (($__local_var_6_11)->{'map'})(function($v_8) use ($f_7) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_9) use ($__local_var_8_12) {
  $__num = \func_num_args();
  $__res = ($__local_var_8_12)($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altWriterT1_5_10 = (object)["alt" => function($v_7) use ($__local_var_5_10) {
  $__num = \func_num_args();
  $__res = function($v1_8) use ($__local_var_5_10, $v_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_10)->{'alt'})($v_7))($v1_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_7) use ($functorWriterT1_6_11) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_6_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusWriterT1_4_9 = (object)["empty" => ($__local_var_4_9)->{'empty'}, "Alt0" => function($_dollar___unused_6) use ($altWriterT1_5_10) {
  $__num = \func_num_args();
  $__res = $altWriterT1_5_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_5) use ($applicativeWriterT2_3_1) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_5) use ($plusWriterT1_4_9) {
  $__num = \func_num_args();
  $__res = $plusWriterT1_4_9;
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
$GLOBALS['Control_Monad_Writer_Trans_alternativeWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_alternativemajWritermajT';

// Control_Monad_Writer_Trans_monadPlusWriterT
function majControl_majMonad_majWriter_majTrans_monadmajPlusmajWritermajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majWriter_majTrans_monadmajPlusmajWritermajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_3_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonadPlus_4) use ($__local_var_1_0, $__local_var_2_1, $__local_var_3_2, $dictMonoid_0) {
  $__num = \func_num_args();
  $__local_var_5_3 = (($dictMonadPlus_4)->{'Monad0'})(null);
  $__local_var_6_4 = (($__local_var_5_3)->{'Applicative0'})(null);
  $__local_var_7_5 = (($__local_var_6_4)->{'Apply0'})(null);
  $Functor0_8_6 = (($__local_var_7_5)->{'Functor0'})(null);
  $__local_var_9_7 = (($__local_var_7_5)->{'Functor0'})(null);
  $functorWriterT1_9_7 = (object)["map" => function($f_10) use ($__local_var_9_7) {
  $__num = \func_num_args();
  $__local_var_11_8 = (($__local_var_9_7)->{'map'})(function($v_11) use ($f_10) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_10)(($v_11)->{'value0'}), ($v_11)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_12) use ($__local_var_11_8) {
  $__num = \func_num_args();
  $__res = ($__local_var_11_8)($v_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_7_5 = (object)["apply" => function($v_10) use ($Functor0_8_6, $__local_var_1_0, $__local_var_7_5) {
  $__num = \func_num_args();
  $__res = function($v1_11) use ($Functor0_8_6, $__local_var_1_0, $__local_var_7_5, $v_10) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_5)->{'apply'})(((($Functor0_8_6)->{'map'})(function($v3_12) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v4_13) use ($__local_var_1_0, $v3_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_12)->{'value0'})(($v4_13)->{'value0'}), ((($__local_var_1_0)->{'append'})(($v3_12)->{'value1'}))(($v4_13)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10)))($v1_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_10) use ($functorWriterT1_9_7) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_9_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_6_4 = (object)["pure" => function($a_8) use ($__local_var_6_4, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_6_4)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_8, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_8) use ($applyWriterT2_7_5) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_7_12 = (($__local_var_5_3)->{'Bind1'})(null);
  $Apply0_8_13 = (($__local_var_7_12)->{'Apply0'})(null);
  $Functor0_9_14 = (($Apply0_8_13)->{'Functor0'})(null);
  $Functor0_10_15 = (($Apply0_8_13)->{'Functor0'})(null);
  $__local_var_11_16 = (($Apply0_8_13)->{'Functor0'})(null);
  $functorWriterT1_11_16 = (object)["map" => function($f_12) use ($__local_var_11_16) {
  $__num = \func_num_args();
  $__local_var_13_17 = (($__local_var_11_16)->{'map'})(function($v_13) use ($f_12) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_12)(($v_13)->{'value0'}), ($v_13)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_14) use ($__local_var_13_17) {
  $__num = \func_num_args();
  $__res = ($__local_var_13_17)($v_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_10_15 = (object)["apply" => function($v_12) use ($Apply0_8_13, $Functor0_10_15, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_13) use ($Apply0_8_13, $Functor0_10_15, $__local_var_2_1, $v_12) {
  $__num = \func_num_args();
  $__res = ((($Apply0_8_13)->{'apply'})(((($Functor0_10_15)->{'map'})(function($v3_14) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v4_15) use ($__local_var_2_1, $v3_14) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_14)->{'value0'})(($v4_15)->{'value0'}), ((($__local_var_2_1)->{'append'})(($v3_14)->{'value1'}))(($v4_15)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_12)))($v1_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_12) use ($functorWriterT1_11_16) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_11_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindWriterT2_7_12 = (object)["bind" => function($v_11) use ($Functor0_9_14, $__local_var_2_1, $__local_var_7_12) {
  $__num = \func_num_args();
  $__res = function($k_12) use ($Functor0_9_14, $__local_var_2_1, $__local_var_7_12, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_12)->{'bind'})($v_11))(function($v1_13) use ($Functor0_9_14, $__local_var_2_1, $k_12) {
  $__num = \func_num_args();
  $__local_var_14_20 = ($v1_13)->{'value1'};
  $__res = ((($Functor0_9_14)->{'map'})(function($v3_15) use ($__local_var_14_20, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v3_15)->{'value0'}, ((($__local_var_2_1)->{'append'})($__local_var_14_20))(($v3_15)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_12)(($v1_13)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_11) use ($applyWriterT2_10_15) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_10_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadWriterT2_5_3 = (object)["Applicative0" => function($_dollar___unused_8) use ($applicativeWriterT2_6_4) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_8) use ($bindWriterT2_7_12) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_7_12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_6_23 = (($dictMonadPlus_4)->{'Alternative1'})(null);
  $__local_var_7_24 = (($__local_var_6_23)->{'Applicative0'})(null);
  $__local_var_8_25 = (($__local_var_7_24)->{'Apply0'})(null);
  $Functor0_9_26 = (($__local_var_8_25)->{'Functor0'})(null);
  $__local_var_10_27 = (($__local_var_8_25)->{'Functor0'})(null);
  $functorWriterT1_10_27 = (object)["map" => function($f_11) use ($__local_var_10_27) {
  $__num = \func_num_args();
  $__local_var_12_28 = (($__local_var_10_27)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_28) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_28)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_8_25 = (object)["apply" => function($v_11) use ($Functor0_9_26, $__local_var_3_2, $__local_var_8_25) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($Functor0_9_26, $__local_var_3_2, $__local_var_8_25, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_8_25)->{'apply'})(((($Functor0_9_26)->{'map'})(function($v3_13) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v4_14) use ($__local_var_3_2, $v3_13) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($v3_13)->{'value0'})(($v4_14)->{'value0'}), ((($__local_var_3_2)->{'append'})(($v3_13)->{'value1'}))(($v4_14)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_11)))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_27) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_27;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeWriterT2_7_24 = (object)["pure" => function($a_9) use ($__local_var_7_24, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_7_24)->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($a_9, ($dictMonoid_0)->{'mempty'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_9) use ($applyWriterT2_8_25) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_8_25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_8_32 = (($__local_var_6_23)->{'Plus1'})(null);
  $__local_var_9_33 = (($__local_var_8_32)->{'Alt0'})(null);
  $__local_var_10_34 = (($__local_var_9_33)->{'Functor0'})(null);
  $functorWriterT1_10_34 = (object)["map" => function($f_11) use ($__local_var_10_34) {
  $__num = \func_num_args();
  $__local_var_12_35 = (($__local_var_10_34)->{'map'})(function($v_12) use ($f_11) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_11)(($v_12)->{'value0'}), ($v_12)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($v_13) use ($__local_var_12_35) {
  $__num = \func_num_args();
  $__res = ($__local_var_12_35)($v_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altWriterT1_9_33 = (object)["alt" => function($v_11) use ($__local_var_9_33) {
  $__num = \func_num_args();
  $__res = function($v1_12) use ($__local_var_9_33, $v_11) {
  $__num = \func_num_args();
  $__res = ((($__local_var_9_33)->{'alt'})($v_11))($v1_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_11) use ($functorWriterT1_10_34) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_10_34;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusWriterT1_8_32 = (object)["empty" => ($__local_var_8_32)->{'empty'}, "Alt0" => function($_dollar___unused_10) use ($altWriterT1_9_33) {
  $__num = \func_num_args();
  $__res = $altWriterT1_9_33;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeWriterT2_6_23 = (object)["Applicative0" => function($_dollar___unused_9) use ($applicativeWriterT2_7_24) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_7_24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_9) use ($plusWriterT1_8_32) {
  $__num = \func_num_args();
  $__res = $plusWriterT1_8_32;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Monad0" => function($_dollar___unused_7) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar___unused_7) use ($alternativeWriterT2_6_23) {
  $__num = \func_num_args();
  $__res = $alternativeWriterT2_6_23;
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
$GLOBALS['Control_Monad_Writer_Trans_monadPlusWriterT'] = __NAMESPACE__ . '\\majControl_majMonad_majWriter_majTrans_monadmajPlusmajWritermajT';

