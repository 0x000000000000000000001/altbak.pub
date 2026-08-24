<?php

namespace Control\Monad\Identity\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Identity.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Functor, Data.Newtype, Data.Ord, Data.Traversable, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Identity.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Functor, Data.Newtype, Data.Ord, Data.Traversable, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Identity.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
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




// Control_Monad_Identity_Trans_IdentityT
function majControl_majMonad_majIdentity_majTrans_majIdentitymajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_majIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_IdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_majIdentitymajT';

// Control_Monad_Identity_Trans_monadSTIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajSmajTmajIdentitymajT($dictMonadST_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajSmajTmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadST_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadSTIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajSmajTmajIdentitymajT';

// Control_Monad_Identity_Trans_traversableIdentityT
function majControl_majMonad_majIdentity_majTrans_traversablemajIdentitymajT($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_traversablemajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictTraversable_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_traversableIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_traversablemajIdentitymajT';

// Control_Monad_Identity_Trans_runIdentityT
function majControl_majMonad_majIdentity_majTrans_runmajIdentitymajT($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_runmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_runIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_runmajIdentitymajT';

// Control_Monad_Identity_Trans_plusIdentityT
function majControl_majMonad_majIdentity_majTrans_plusmajIdentitymajT($dictPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_plusmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictPlus_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_plusIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_plusmajIdentitymajT';

// Control_Monad_Identity_Trans_newtypeIdentityT
$GLOBALS['Control_Monad_Identity_Trans_newtypeIdentityT'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Identity_Trans_monadWriterIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajWritermajIdentitymajT($dictMonadWriter_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajWritermajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadWriter_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadWriterIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajWritermajIdentitymajT';

// Control_Monad_Identity_Trans_monadTransIdentityT
$GLOBALS['Control_Monad_Identity_Trans_monadTransIdentityT'] = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  $__res = function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Identity_Trans_monadThrowIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajThrowmajIdentitymajT($dictMonadThrow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajThrowmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadThrow_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadThrowIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajThrowmajIdentitymajT';

// Control_Monad_Identity_Trans_monadTellIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajTellmajIdentitymajT($dictMonadTell_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajTellmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadTell_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadTellIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajTellmajIdentitymajT';

// Control_Monad_Identity_Trans_monadStateIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajStatemajIdentitymajT($dictMonadState_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajStatemajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadState_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadStateIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajStatemajIdentitymajT';

// Control_Monad_Identity_Trans_monadRecIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajRecmajIdentitymajT($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajRecmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadRec_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadRecIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajRecmajIdentitymajT';

// Control_Monad_Identity_Trans_monadReaderIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajReadermajIdentitymajT($dictMonadReader_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajReadermajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadReader_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadReaderIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajReadermajIdentitymajT';

// Control_Monad_Identity_Trans_monadPlusIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajPlusmajIdentitymajT($dictMonadPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajPlusmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadPlus_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadPlusIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajPlusmajIdentitymajT';

// Control_Monad_Identity_Trans_monadIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajIdentitymajT($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonad_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajIdentitymajT';

// Control_Monad_Identity_Trans_monadErrorIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajErrormajIdentitymajT($dictMonadError_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajErrormajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadError_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadErrorIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajErrormajIdentitymajT';

// Control_Monad_Identity_Trans_monadEffectIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajEffectmajIdentitymajT($dictMonadEffect_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajEffectmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadEffect_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadEffectIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajEffectmajIdentitymajT';

// Control_Monad_Identity_Trans_monadContIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajContmajIdentitymajT($dictMonadCont_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajContmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadCont_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadContIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajContmajIdentitymajT';

// Control_Monad_Identity_Trans_monadAskIdentityT
function majControl_majMonad_majIdentity_majTrans_monadmajAskmajIdentitymajT($dictMonadAsk_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_monadmajAskmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictMonadAsk_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_monadAskIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_monadmajAskmajIdentitymajT';

// Control_Monad_Identity_Trans_mapIdentityT
function majControl_majMonad_majIdentity_majTrans_mapmajIdentitymajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_mapmajIdentitymajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_mapIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_mapmajIdentitymajT';

// Control_Monad_Identity_Trans_functorIdentityT
function majControl_majMonad_majIdentity_majTrans_functormajIdentitymajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_functormajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictFunctor_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_functorIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_functormajIdentitymajT';

// Control_Monad_Identity_Trans_foldableIdentityT
function majControl_majMonad_majIdentity_majTrans_foldablemajIdentitymajT($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_foldablemajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictFoldable_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_foldableIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_foldablemajIdentitymajT';

// Control_Monad_Identity_Trans_extendIdentityI
function majControl_majMonad_majIdentity_majTrans_extendmajIdentitymajI($dictExtend_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_extendmajIdentitymajI';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorIdentityT1_1_0 = (($dictExtend_0)->{'Functor0'})(null);
  $__res = (object)["extend" => function($f_2) use ($dictExtend_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictExtend_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictExtend_0)->{'extend'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorIdentityT1_1_0) {
  $__num = \func_num_args();
  $__res = $functorIdentityT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_extendIdentityI'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_extendmajIdentitymajI';

// Control_Monad_Identity_Trans_eqIdentityT
function majControl_majMonad_majIdentity_majTrans_eqmajIdentitymajT($dictEq1_0, $dictEq_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_eqmajIdentitymajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq" => function($x_2) use ($dictEq1_0, $dictEq_1) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictEq1_0, $dictEq_1, $x_2) {
  $__num = \func_num_args();
  $__res = (((($dictEq1_0)->{'eq1'})($dictEq_1))($x_2))($y_3);
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
$GLOBALS['Control_Monad_Identity_Trans_eqIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_eqmajIdentitymajT';

// Control_Monad_Identity_Trans_ordIdentityT
function majControl_majMonad_majIdentity_majTrans_ordmajIdentitymajT($dictOrd1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_ordmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd1_0)->{'Eq10'})(null);
  $__res = function($dictOrd_2) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd_2)->{'Eq0'})(null);
  $eqIdentityT2_3_1 = (object)["eq" => function($x_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($y_5) use ($__local_var_1_0, $__local_var_3_1, $x_4) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'eq1'})($__local_var_3_1))($x_4))($y_5);
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
  $__res = (((($dictOrd1_0)->{'compare1'})($dictOrd_2))($x_4))($y_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_4) use ($eqIdentityT2_3_1) {
  $__num = \func_num_args();
  $__res = $eqIdentityT2_3_1;
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
$GLOBALS['Control_Monad_Identity_Trans_ordIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_ordmajIdentitymajT';

// Control_Monad_Identity_Trans_eq1IdentityT
function majControl_majMonad_majIdentity_majTrans_eq1majIdentitymajT($dictEq1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_eq1majIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq1" => function($dictEq_1) use ($dictEq1_0) {
  $__num = \func_num_args();
  $__res = function($x_2) use ($dictEq1_0, $dictEq_1) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictEq1_0, $dictEq_1, $x_2) {
  $__num = \func_num_args();
  $__res = (((($dictEq1_0)->{'eq1'})($dictEq_1))($x_2))($y_3);
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
$GLOBALS['Control_Monad_Identity_Trans_eq1IdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_eq1majIdentitymajT';

// Control_Monad_Identity_Trans_ord1IdentityT
function majControl_majMonad_majIdentity_majTrans_ord1majIdentitymajT($dictOrd1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_ord1majIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd1_0)->{'Eq10'})(null);
  $eq1IdentityT1_1_0 = (object)["eq1" => function($dictEq_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($x_3) use ($__local_var_1_0, $dictEq_2) {
  $__num = \func_num_args();
  $__res = function($y_4) use ($__local_var_1_0, $dictEq_2, $x_3) {
  $__num = \func_num_args();
  $__res = (((($__local_var_1_0)->{'eq1'})($dictEq_2))($x_3))($y_4);
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
  $__res = (((($dictOrd1_0)->{'compare1'})($dictOrd_2))($x_3))($y_4);
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
}, "Eq10" => function($_dollar___unused_2) use ($eq1IdentityT1_1_0) {
  $__num = \func_num_args();
  $__res = $eq1IdentityT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_ord1IdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_ord1majIdentitymajT';

// Control_Monad_Identity_Trans_comonadIdentityT
function majControl_majMonad_majIdentity_majTrans_comonadmajIdentitymajT($dictComonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_comonadmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonad_0)->{'Extend0'})(null);
  $functorIdentityT1_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $extendIdentityI1_1_0 = (object)["extend" => function($f_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'extend'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_3))(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorIdentityT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorIdentityT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["extract" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictComonad_0)->{'extract'}))($GLOBALS['Control_Monad_Identity_Trans_runIdentityT']), "Extend0" => function($_dollar___unused_2) use ($extendIdentityI1_1_0) {
  $__num = \func_num_args();
  $__res = $extendIdentityI1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_comonadIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_comonadmajIdentitymajT';

// Control_Monad_Identity_Trans_bindIdentityT
function majControl_majMonad_majIdentity_majTrans_bindmajIdentitymajT($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_bindmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictBind_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_bindIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_bindmajIdentitymajT';

// Control_Monad_Identity_Trans_applyIdentityT
function majControl_majMonad_majIdentity_majTrans_applymajIdentitymajT($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_applymajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictApply_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_applyIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_applymajIdentitymajT';

// Control_Monad_Identity_Trans_applicativeIdentityT
function majControl_majMonad_majIdentity_majTrans_applicativemajIdentitymajT($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_applicativemajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictApplicative_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_applicativeIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_applicativemajIdentitymajT';

// Control_Monad_Identity_Trans_alternativeIdentityT
function majControl_majMonad_majIdentity_majTrans_alternativemajIdentitymajT($dictAlternative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_alternativemajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictAlternative_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_alternativeIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_alternativemajIdentitymajT';

// Control_Monad_Identity_Trans_altIdentityT
function majControl_majMonad_majIdentity_majTrans_altmajIdentitymajT($dictAlt_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majIdentity_majTrans_altmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictAlt_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_Identity_Trans_altIdentityT'] = __NAMESPACE__ . '\\majControl_majMonad_majIdentity_majTrans_altmajIdentitymajT';

