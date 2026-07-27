<?php

namespace Effect;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Data.Functor, Data.Monoid, Data.Semigroup, Effect, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Data.Functor, Data.Monoid, Data.Semigroup, Effect, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Effect/index.php';
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
$ffi_Effect = \call_user_func(function() {
  $exports = [];
$pureE = function($x) { return function() use($x) { return $x; }; };
$bindE = function($a, $f = null) use (&$bindE) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$bindE) {
            return $bindE(...\array_merge($__args, $more));
        };
    }
    return function() use($a, $f) {
        $a_res = $a();
        $res = $f($a_res);
        return $res();
    };
};

$untilE = function($f) {
    return function() use ($f) {
        while (!$f()) {}
    };
};

$whileE = function($f, $a = null) use (&$whileE) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$whileE) {
            return $whileE(...\array_merge($__args, $more));
        };
    }
    return function() use ($f, $a) {
        while ($f()) {
            $a();
        }
    };
};

$forE = function($lo, $hi = null, $f = null) use (&$forE) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$forE) {
            return $forE(...\array_merge($__args, $more));
        };
    }
    return function() use ($lo, $hi, $f) {
        for ($i = $lo; $i < $hi; $i++) {
            $f($i)();
        }
    };
};

$foreachE = function($as, $f = null) use (&$foreachE) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foreachE) {
            return $foreachE(...\array_merge($__args, $more));
        };
    }
    return function() use ($as, $f) {
        foreach ($as as $a) {
            $f($a)();
        }
    };
};

$exports['pureE'] = $pureE;
$exports['bindE'] = $bindE;
$exports['untilE'] = $untilE;
$exports['whileE'] = $whileE;
$exports['forE'] = $forE;
$exports['foreachE'] = $foreachE;

return $exports;
  return $exports;
});
function majEffect_bindmajE($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_bindmajE';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect;
  $f = ($ffi_Effect['bindE'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_bindE'] = __NAMESPACE__ . '\\majEffect_bindmajE';

function majEffect_formajE(int $v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_formajE';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Effect;
  $f = ($ffi_Effect['forE'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Effect_forE'] = __NAMESPACE__ . '\\majEffect_formajE';

function majEffect_foreachmajE($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_foreachmajE';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect;
  $f = ($ffi_Effect['foreachE'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_foreachE'] = __NAMESPACE__ . '\\majEffect_foreachmajE';

function majEffect_puremajE($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_puremajE';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect;
  $f = ($ffi_Effect['pureE'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_pureE'] = __NAMESPACE__ . '\\majEffect_puremajE';

function majEffect_untilmajE($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_untilmajE';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect;
  $f = ($ffi_Effect['untilE'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_untilE'] = __NAMESPACE__ . '\\majEffect_untilmajE';

function majEffect_whilemajE($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_whilemajE';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect;
  $f = ($ffi_Effect['whileE'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_whileE'] = __NAMESPACE__ . '\\majEffect_whilemajE';





// Effect_monadEffect
$GLOBALS['Effect_monadEffect'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_applicativeEffect'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_bindEffect'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_bindEffect
$GLOBALS['Effect_bindEffect'] = (object)["bind" => $GLOBALS['Effect_bindE'], "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_applyEffect'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_applyEffect
$GLOBALS['Effect_applyEffect'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Effect_monadEffect'])->{'Bind1'})(null);
return (object)["apply" => (function() use ($__local_var_0_0) {
  $__fn = function($f_1, $a_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_0_0)->{'bind'})($f_1))(function($f_prime_3) use ($__local_var_0_0, $a_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_0_0)->{'bind'})($a_2))(function($a_prime_4) use ($f_prime_3) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Effect_monadEffect'])->{'Applicative0'})(null))->{'pure'})(($f_prime_3)($a_prime_4));
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_functorEffect'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Effect_applicativeEffect
$GLOBALS['Effect_applicativeEffect'] = (object)["pure" => $GLOBALS['Effect_pureE'], "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_applyEffect'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_functorEffect
$GLOBALS['Effect_functorEffect'] = (object)["map" => (function() {
  $__fn = function($f_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Effect_applicativeEffect'])->{'Apply0'})(null))->{'apply'})((($GLOBALS['Effect_applicativeEffect'])->{'pure'})($f_0)))($a_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Effect_lift2
function majEffect_lift2($f_0, $a_1 = null, $b_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_lift2';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = (($GLOBALS['Effect_monadEffect'])->{'Bind1'})(null);
  $__res = ((($__local_var_3_0)->{'bind'})(((((($GLOBALS['Effect_applyEffect'])->{'Functor0'})(null))->{'map'})($f_0))($a_1)))(function($f_prime_4) use ($__local_var_3_0, $b_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_0)->{'bind'})($b_2))(function($a_prime_5) use ($f_prime_4) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Effect_monadEffect'])->{'Applicative0'})(null))->{'pure'})(($f_prime_4)($a_prime_5));
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
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Effect_lift2'] = __NAMESPACE__ . '\\majEffect_lift2';

// Effect_semigroupEffect
function majEffect_semigroupmajEffect($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_semigroupmajEffect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["append" => ($GLOBALS['Effect_lift2'])(($dictSemigroup_0)->{'append'})];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_semigroupEffect'] = __NAMESPACE__ . '\\majEffect_semigroupmajEffect';

// Effect_monoidEffect
function majEffect_monoidmajEffect($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_monoidmajEffect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["mempty" => \Effect\majEffect_puremajE(($dictMonoid_0)->{'mempty'}), "Semigroup0" => function($_dollar__unused_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (object)["append" => ($GLOBALS['Effect_lift2'])(((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_monoidEffect'] = __NAMESPACE__ . '\\majEffect_monoidmajEffect';

