<?php

namespace Data\Functor\Joker;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Biapplicative, Control.Biapply, Control.Bind, Control.Monad, Control.Semigroupoid, Data.Bifunctor, Data.Either, Data.Eq, Data.Function, Data.Functor, Data.Functor.Joker, Data.Newtype, Data.Ord, Data.Profunctor, Data.Profunctor.Choice, Data.Semigroup, Data.Show, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Biapplicative, Control.Biapply, Control.Bind, Control.Monad, Control.Semigroupoid, Data.Bifunctor, Data.Either, Data.Eq, Data.Function, Data.Functor, Data.Functor.Joker, Data.Newtype, Data.Ord, Data.Profunctor, Data.Profunctor.Choice, Data.Semigroup, Data.Show, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Biapplicative/index.php';
require_once __DIR__ . '/../Control.Biapply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Joker/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Profunctor/index.php';
require_once __DIR__ . '/../Data.Profunctor.Choice/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
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




// Data_Functor_Joker_Joker
function majData_majFunctor_majJoker_majJoker($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_majJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_Joker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_majJoker';

// Data_Functor_Joker_showJoker
function majData_majFunctor_majJoker_showmajJoker($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_showmajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = (("(Joker " . (($dictShow_0)->{'show'})($v_1)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_showJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_showmajJoker';

// Data_Functor_Joker_profunctorJoker
function majData_majFunctor_majJoker_profunctormajJoker($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_profunctormajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["dimap" => function($v_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($g_2) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictFunctor_0, $g_2) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})($g_2))($v1_3);
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
$GLOBALS['Data_Functor_Joker_profunctorJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_profunctormajJoker';

// Data_Functor_Joker_ordJoker
function majData_majFunctor_majJoker_ordmajJoker($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_ordmajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictOrd_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_ordJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_ordmajJoker';

// Data_Functor_Joker_newtypeJoker
$GLOBALS['Data_Functor_Joker_newtypeJoker'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Functor_Joker_hoistJoker
function majData_majFunctor_majJoker_hoistmajJoker($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_hoistmajJoker';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Functor_Joker_hoistJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_hoistmajJoker';

// Data_Functor_Joker_functorJoker
function majData_majFunctor_majJoker_functormajJoker($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_functormajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})($f_1))($v_2);
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
$GLOBALS['Data_Functor_Joker_functorJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_functormajJoker';

// Data_Functor_Joker_eqJoker
function majData_majFunctor_majJoker_eqmajJoker($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_eqmajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $dictEq_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_eqJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_eqmajJoker';

// Data_Functor_Joker_choiceJoker
function majData_majFunctor_majJoker_choicemajJoker($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_choicemajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $profunctorJoker1_1_0 = (object)["dimap" => function($v_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($g_2) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictFunctor_0, $g_2) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})($g_2))($v1_3);
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
  $__res = (object)["left" => function($v_2) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})($GLOBALS['Data_Either_Left']))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "right" => function($v_2) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})($GLOBALS['Data_Either_Right']))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Profunctor0" => function($_dollar___unused_2) use ($profunctorJoker1_1_0) {
  $__num = \func_num_args();
  $__res = $profunctorJoker1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_choiceJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_choicemajJoker';

// Data_Functor_Joker_bifunctorJoker
function majData_majFunctor_majJoker_bifunctormajJoker($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_bifunctormajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["bimap" => function($v_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($g_2) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictFunctor_0, $g_2) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})($g_2))($v1_3);
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
$GLOBALS['Data_Functor_Joker_bifunctorJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_bifunctormajJoker';

// Data_Functor_Joker_biapplyJoker
function majData_majFunctor_majJoker_biapplymajJoker($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_biapplymajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $bifunctorJoker1_1_0 = (object)["bimap" => function($v_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($g_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $g_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})($g_3))($v1_4);
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
  $__res = (object)["biapply" => function($v_2) use ($dictApply_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictApply_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar___unused_2) use ($bifunctorJoker1_1_0) {
  $__num = \func_num_args();
  $__res = $bifunctorJoker1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_biapplyJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_biapplymajJoker';

// Data_Functor_Joker_biapplicativeJoker
function majData_majFunctor_majJoker_biapplicativemajJoker($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_biapplicativemajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $bifunctorJoker1_2_1 = (object)["bimap" => function($v_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($g_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $g_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'map'})($g_4))($v1_5);
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
  $biapplyJoker1_1_0 = (object)["biapply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})($v_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar___unused_3) use ($bifunctorJoker1_2_1) {
  $__num = \func_num_args();
  $__res = $bifunctorJoker1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bipure" => function($v_2) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)->{'pure'})($b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Biapply0" => function($_dollar___unused_2) use ($biapplyJoker1_1_0) {
  $__num = \func_num_args();
  $__res = $biapplyJoker1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_biapplicativeJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_biapplicativemajJoker';

// Data_Functor_Joker_applyJoker
function majData_majFunctor_majJoker_applymajJoker($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_applymajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $functorJoker1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})($f_2))($v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["apply" => function($v_2) use ($dictApply_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictApply_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorJoker1_1_0) {
  $__num = \func_num_args();
  $__res = $functorJoker1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_applyJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_applymajJoker';

// Data_Functor_Joker_bindJoker
function majData_majFunctor_majJoker_bindmajJoker($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_bindmajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBind_0)->{'Apply0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorJoker1_2_1 = (object)["map" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'map'})($f_3))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyJoker1_1_0 = (object)["apply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})($v_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorJoker1_2_1) {
  $__num = \func_num_args();
  $__res = $functorJoker1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_2) use ($dictBind_0) {
  $__num = \func_num_args();
  $__res = function($amb_3) use ($dictBind_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictBind_0)->{'bind'})($v_2))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($amb_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyJoker1_1_0) {
  $__num = \func_num_args();
  $__res = $applyJoker1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_bindJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_bindmajJoker';

// Data_Functor_Joker_applicativeJoker
function majData_majFunctor_majJoker_applicativemajJoker($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_applicativemajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorJoker1_2_1 = (object)["map" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'map'})($f_3))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyJoker1_1_0 = (object)["apply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})($v_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorJoker1_2_1) {
  $__num = \func_num_args();
  $__res = $functorJoker1_2_1;
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
}))(($dictApplicative_0)->{'pure'}), "Apply0" => function($_dollar___unused_2) use ($applyJoker1_1_0) {
  $__num = \func_num_args();
  $__res = $applyJoker1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_applicativeJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_applicativemajJoker';

// Data_Functor_Joker_monadJoker
function majData_majFunctor_majJoker_monadmajJoker($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFunctor_majJoker_monadmajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Apply0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $functorJoker1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})($f_4))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyJoker1_2_1 = (object)["apply" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'apply'})($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorJoker1_3_2) {
  $__num = \func_num_args();
  $__res = $functorJoker1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeJoker1_1_0 = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($__local_var_1_0)->{'pure'}), "Apply0" => function($_dollar___unused_3) use ($applyJoker1_2_1) {
  $__num = \func_num_args();
  $__res = $applyJoker1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_6 = (($dictMonad_0)->{'Bind1'})(null);
  $__local_var_3_7 = (($__local_var_2_6)->{'Apply0'})(null);
  $__local_var_4_8 = (($__local_var_3_7)->{'Functor0'})(null);
  $functorJoker1_4_8 = (object)["map" => function($f_5) use ($__local_var_4_8) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_8, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_8)->{'map'})($f_5))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyJoker1_3_7 = (object)["apply" => function($v_5) use ($__local_var_3_7) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_7, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_7)->{'apply'})($v_5))($v1_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorJoker1_4_8) {
  $__num = \func_num_args();
  $__res = $functorJoker1_4_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindJoker1_2_6 = (object)["bind" => function($v_4) use ($__local_var_2_6) {
  $__num = \func_num_args();
  $__res = function($amb_5) use ($__local_var_2_6, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_6)->{'bind'})($v_4))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($amb_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($applyJoker1_3_7) {
  $__num = \func_num_args();
  $__res = $applyJoker1_3_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeJoker1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeJoker1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindJoker1_2_6) {
  $__num = \func_num_args();
  $__res = $bindJoker1_2_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Functor_Joker_monadJoker'] = __NAMESPACE__ . '\\majData_majFunctor_majJoker_monadmajJoker';

