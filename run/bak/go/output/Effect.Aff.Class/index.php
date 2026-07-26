<?php

namespace Effect\Aff\Class;

// ALL IMPORTS: Control.Category, Control.Monad.Cont.Trans, Control.Monad.Except.Trans, Control.Monad.List.Trans, Control.Monad.Maybe.Trans, Control.Monad.RWS.Trans, Control.Monad.Reader.Trans, Control.Monad.State.Trans, Control.Monad.Trans.Class, Control.Monad.Writer.Trans, Control.Semigroupoid, Effect.Aff, Effect.Aff.Class, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Category, Control.Monad.Cont.Trans, Control.Monad.Except.Trans, Control.Monad.List.Trans, Control.Monad.Maybe.Trans, Control.Monad.RWS.Trans, Control.Monad.Reader.Trans, Control.Monad.State.Trans, Control.Monad.Trans.Class, Control.Monad.Writer.Trans, Control.Semigroupoid, Effect.Aff, Effect.Aff.Class, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Except.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.List.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Maybe.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.RWS.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.State.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Trans/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Effect.Aff/index.php';
require_once __DIR__ . '/../Effect.Aff.Class/index.php';
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
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffAff'] = function() { $v = ["liftAff" => (($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity'], "MonadEffect0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_monadEffectAff'] ?? \PhpursThunks::eval('Effect_Aff_monadEffectAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_Class_liftAff'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['liftAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffContT'] = function() { $v = function($dictMonadAff_0 = null) {
  $__num = \func_num_args();
  $MonadEffect0_1_0 = (($dictMonadAff_0)['MonadEffect0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadEffectContT_2_1 = (($GLOBALS['Control_Monad_Cont_Trans_monadEffectContT'] ?? \PhpursThunks::eval('Control_Monad_Cont_Trans_monadEffectContT')))($MonadEffect0_1_0);
  $__res = ["liftAff" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((((($MonadEffect0_1_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['bind']))(($dictMonadAff_0)['liftAff']), "MonadEffect0" => function($dollar__unused_3 = null) use ($monadEffectContT_2_1) {
  $__num = \func_num_args();
  $__res = $monadEffectContT_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffExceptT'] = function() { $v = function($dictMonadAff_0 = null) {
  $__num = \func_num_args();
  $MonadEffect0_1_0 = (($dictMonadAff_0)['MonadEffect0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadEffectExceptT_2_1 = (($GLOBALS['Control_Monad_Except_Trans_monadEffectExceptT'] ?? \PhpursThunks::eval('Control_Monad_Except_Trans_monadEffectExceptT')))($MonadEffect0_1_0);
  $__res = ["liftAff" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Monad_Except_Trans_monadTransExceptT'] ?? \PhpursThunks::eval('Control_Monad_Except_Trans_monadTransExceptT')))['lift'])((($MonadEffect0_1_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))))(($dictMonadAff_0)['liftAff']), "MonadEffect0" => function($dollar__unused_3 = null) use ($monadEffectExceptT_2_1) {
  $__num = \func_num_args();
  $__res = $monadEffectExceptT_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffListT'] = function() { $v = function($dictMonadAff_0 = null) {
  $__num = \func_num_args();
  $MonadEffect0_1_0 = (($dictMonadAff_0)['MonadEffect0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadEffectListT_2_1 = (($GLOBALS['Control_Monad_List_Trans_monadEffectListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_monadEffectListT')))($MonadEffect0_1_0);
  $__res = ["liftAff" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Control_Monad_List_Trans_fromEffect'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_fromEffect')))((((($MonadEffect0_1_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))))(($dictMonadAff_0)['liftAff']), "MonadEffect0" => function($dollar__unused_3 = null) use ($monadEffectListT_2_1) {
  $__num = \func_num_args();
  $__res = $monadEffectListT_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffMaybe'] = function() { $v = function($dictMonadAff_0 = null) {
  $__num = \func_num_args();
  $MonadEffect0_1_0 = (($dictMonadAff_0)['MonadEffect0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadEffectMaybe_2_1 = (($GLOBALS['Control_Monad_Maybe_Trans_monadEffectMaybe'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_monadEffectMaybe')))($MonadEffect0_1_0);
  $__res = ["liftAff" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Monad_liftM1'] ?? \PhpursThunks::eval('Control_Monad_liftM1')))((($MonadEffect0_1_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))))(($dictMonadAff_0)['liftAff']), "MonadEffect0" => function($dollar__unused_3 = null) use ($monadEffectMaybe_2_1) {
  $__num = \func_num_args();
  $__res = $monadEffectMaybe_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffRWS'] = function() { $v = function($dictMonadAff_0 = null) {
  $__num = \func_num_args();
  $MonadEffect0_1_0 = (($dictMonadAff_0)['MonadEffect0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Monad0_2_1 = (($MonadEffect0_1_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $liftAff1_3_2 = ($dictMonadAff_0)['liftAff'];
  $__res = function($dictMonoid_4 = null) use ($Monad0_2_1, $MonadEffect0_1_0, $liftAff1_3_2) {
  $__num = \func_num_args();
  $monadEffectRWS_5_3 = ((($GLOBALS['Control_Monad_RWS_Trans_monadEffectRWS'] ?? \PhpursThunks::eval('Control_Monad_RWS_Trans_monadEffectRWS')))($dictMonoid_4))($MonadEffect0_1_0);
  $__res = ["liftAff" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Control_Monad_RWS_Trans_monadTransRWST'] ?? \PhpursThunks::eval('Control_Monad_RWS_Trans_monadTransRWST')))($dictMonoid_4))['lift'])($Monad0_2_1)))($liftAff1_3_2), "MonadEffect0" => function($dollar__unused_6 = null) use ($monadEffectRWS_5_3) {
  $__num = \func_num_args();
  $__res = $monadEffectRWS_5_3;
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
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffReader'] = function() { $v = function($dictMonadAff_0 = null) {
  $__num = \func_num_args();
  $monadEffectReader_1_0 = (($GLOBALS['Control_Monad_Reader_Trans_monadEffectReader'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_monadEffectReader')))((($dictMonadAff_0)['MonadEffect0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["liftAff" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Reader_Trans_ReaderT'] ?? \PhpursThunks::eval('Control_Monad_Reader_Trans_ReaderT'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')))))(($dictMonadAff_0)['liftAff']), "MonadEffect0" => function($dollar__unused_2 = null) use ($monadEffectReader_1_0) {
  $__num = \func_num_args();
  $__res = $monadEffectReader_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffState'] = function() { $v = function($dictMonadAff_0 = null) {
  $__num = \func_num_args();
  $MonadEffect0_1_0 = (($dictMonadAff_0)['MonadEffect0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadEffectState_2_1 = (($GLOBALS['Control_Monad_State_Trans_monadEffectState'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadEffectState')))($MonadEffect0_1_0);
  $__res = ["liftAff" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Monad_State_Trans_monadTransStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadTransStateT')))['lift'])((($MonadEffect0_1_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))))(($dictMonadAff_0)['liftAff']), "MonadEffect0" => function($dollar__unused_3 = null) use ($monadEffectState_2_1) {
  $__num = \func_num_args();
  $__res = $monadEffectState_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_Class_monadAffWriter'] = function() { $v = function($dictMonadAff_0 = null) {
  $__num = \func_num_args();
  $MonadEffect0_1_0 = (($dictMonadAff_0)['MonadEffect0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Monad0_2_1 = (($MonadEffect0_1_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $liftAff1_3_2 = ($dictMonadAff_0)['liftAff'];
  $__res = function($dictMonoid_4 = null) use ($Monad0_2_1, $MonadEffect0_1_0, $liftAff1_3_2) {
  $__num = \func_num_args();
  $monadEffectWriter_5_3 = ((($GLOBALS['Control_Monad_Writer_Trans_monadEffectWriter'] ?? \PhpursThunks::eval('Control_Monad_Writer_Trans_monadEffectWriter')))($dictMonoid_4))($MonadEffect0_1_0);
  $__res = ["liftAff" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Control_Monad_Writer_Trans_monadTransWriterT'] ?? \PhpursThunks::eval('Control_Monad_Writer_Trans_monadTransWriterT')))($dictMonoid_4))['lift'])($Monad0_2_1)))($liftAff1_3_2), "MonadEffect0" => function($dollar__unused_6 = null) use ($monadEffectWriter_5_3) {
  $__num = \func_num_args();
  $__res = $monadEffectWriter_5_3;
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
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };












