<?php

namespace Data\Profunctor\Star;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Distributive, Data.Either, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Newtype, Data.Profunctor, Data.Profunctor.Choice, Data.Profunctor.Closed, Data.Profunctor.Star, Data.Profunctor.Strong, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Distributive, Data.Either, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Newtype, Data.Profunctor, Data.Profunctor.Choice, Data.Profunctor.Closed, Data.Profunctor.Star, Data.Profunctor.Strong, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Distributive/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Invariant/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Profunctor/index.php';
require_once __DIR__ . '/../Data.Profunctor.Choice/index.php';
require_once __DIR__ . '/../Data.Profunctor.Closed/index.php';
require_once __DIR__ . '/../Data.Profunctor.Star/index.php';
require_once __DIR__ . '/../Data.Profunctor.Strong/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_Profunctor_Star_Star
function majData_majProfunctor_majStar_majStar($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_majStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_Star'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_majStar';

// Data_Profunctor_Star_semigroupoidStar
function majData_majProfunctor_majStar_semigroupoidmajStar($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_semigroupoidmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["compose" => function($v_1) use ($dictBind_0) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($dictBind_0, $v_1) {
  $__num = \func_num_args();
  $__res = function($x_3) use ($dictBind_0, $v1_2, $v_1) {
  $__num = \func_num_args();
  $__res = ((($dictBind_0)->{'bind'})(($v1_2)($x_3)))($v_1);
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
$GLOBALS['Data_Profunctor_Star_semigroupoidStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_semigroupoidmajStar';

// Data_Profunctor_Star_profunctorStar
function majData_majProfunctor_majStar_profunctormajStar($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_profunctormajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["dimap" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($g_2) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFunctor_0, $f_1, $g_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictFunctor_0)->{'map'})($g_2)))($v_3)))($f_1);
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
$GLOBALS['Data_Profunctor_Star_profunctorStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_profunctormajStar';

// Data_Profunctor_Star_strongStar
function majData_majProfunctor_majStar_strongmajStar($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_strongmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $profunctorStar1_1_0 = (object)["dimap" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($g_2) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFunctor_0, $f_1, $g_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictFunctor_0)->{'map'})($g_2)))($v_3)))($f_1);
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
  $__res = (object)["first" => function($v_2) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictFunctor_0, $v_2) {
  $__num = \func_num_args();
  $__local_var_4_1 = ($v1_3)->{'value1'};
  $__res = ((($dictFunctor_0)->{'map'})(function($v2_5) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple($v2_5, $__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_2)(($v1_3)->{'value0'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "second" => function($v_2) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictFunctor_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictFunctor_0)->{'map'})(($GLOBALS['Data_Tuple_Tuple'])(($v1_3)->{'value0'})))(($v_2)(($v1_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Profunctor0" => function($_dollar___unused_2) use ($profunctorStar1_1_0) {
  $__num = \func_num_args();
  $__res = $profunctorStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_strongStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_strongmajStar';

// Data_Profunctor_Star_newtypeStar
$GLOBALS['Data_Profunctor_Star_newtypeStar'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Profunctor_Star_invariantStar
function majData_majProfunctor_majStar_invariantmajStar($dictInvariant_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_invariantmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["imap" => function($f_1) use ($dictInvariant_0) {
  $__num = \func_num_args();
  $__res = function($g_2) use ($dictInvariant_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictInvariant_0, $f_1, $g_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictInvariant_0)->{'imap'})($f_1))($g_2)))($v_3);
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
$GLOBALS['Data_Profunctor_Star_invariantStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_invariantmajStar';

// Data_Profunctor_Star_hoistStar
function majData_majProfunctor_majStar_hoistmajStar($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_hoistmajStar';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_hoistStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_hoistmajStar';

// Data_Profunctor_Star_functorStar
function majData_majProfunctor_majStar_functormajStar($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_functormajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => function($f_1) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictFunctor_0)->{'map'})($f_1)))($v_2);
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
$GLOBALS['Data_Profunctor_Star_functorStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_functormajStar';

// Data_Profunctor_Star_distributiveStar
function majData_majProfunctor_majStar_distributivemajStar($dictDistributive_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_distributivemajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_Profunctor_Star_distributiveStar_dictDistributive_0 = $dictDistributive_0;
  tco_loop_Data_Profunctor_Star_distributiveStar:;
  $dictDistributive_0 = $__tco_var_Data_Profunctor_Star_distributiveStar_dictDistributive_0;
  $__local_var_1_0 = (($dictDistributive_0)->{'Functor0'})(null);
  $functorStar1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_1_0)->{'map'})($f_2)))($v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["distribute" => function($dictFunctor_2) use ($dictDistributive_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictDistributive_0, $dictFunctor_2) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($dictDistributive_0, $dictFunctor_2, $f_3) {
  $__num = \func_num_args();
  $__res = (((($dictDistributive_0)->{'collect'})($dictFunctor_2))(function($v_5) use ($a_4) {
  $__num = \func_num_args();
  $__res = ($v_5)($a_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_3);
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
}, "collect" => function($dictFunctor_2) use ($dictDistributive_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictDistributive_0, $dictFunctor_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_Profunctor_Star_distributiveStar'])($dictDistributive_0))->{'distribute'})($dictFunctor_2)))((($dictFunctor_2)->{'map'})($f_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_2) use ($functorStar1_1_0) {
  $__num = \func_num_args();
  $__res = $functorStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_distributiveStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_distributivemajStar';

// Data_Profunctor_Star_closedStar
function majData_majProfunctor_majStar_closedmajStar($dictDistributive_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_closedmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictDistributive_0)->{'Functor0'})(null);
  $profunctorStar1_1_0 = (object)["dimap" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($g_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $f_2, $g_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_1_0)->{'map'})($g_3)))($v_4)))($f_2);
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
  $__res = (object)["closed" => function($v_2) use ($dictDistributive_0) {
  $__num = \func_num_args();
  $__res = function($g_3) use ($dictDistributive_0, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictDistributive_0)->{'distribute'})($GLOBALS['Data_Functor_functorFn']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($v_2))($g_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Profunctor0" => function($_dollar___unused_2) use ($profunctorStar1_1_0) {
  $__num = \func_num_args();
  $__res = $profunctorStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_closedStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_closedmajStar';

// Data_Profunctor_Star_choiceStar
function majData_majProfunctor_majStar_choicemajStar($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_choicemajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_2_1 = (($Apply0_1_0)->{'Functor0'})(null);
  $pure_3_2 = ($dictApplicative_0)->{'pure'};
  $pure1_4_3 = ($dictApplicative_0)->{'pure'};
  $__local_var_5_4 = (($Apply0_1_0)->{'Functor0'})(null);
  $profunctorStar1_5_4 = (object)["dimap" => function($f_6) use ($__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($g_7) use ($__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_5_4, $f_6, $g_7) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_5_4)->{'map'})($g_7)))($v_8)))($f_6);
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
  $__res = (object)["left" => function($v_6) use ($Functor0_2_1, $pure_3_2) {
  $__num = \func_num_args();
  $__local_var_7_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_2_1)->{'map'})($GLOBALS['Data_Either_Left'])))($v_6);
  $__local_var_8_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_3_2))($GLOBALS['Data_Either_Right']);
  $__res = function($v2_9) use ($__local_var_7_6, $__local_var_8_7) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v2_9 instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($__local_var_7_6)(($v2_9)->{'value0'});
goto end_branch_8;;
};
  if ($v2_9 instanceof \Data\Either\Data_Either_Right) {
$__t8 = ($__local_var_8_7)(($v2_9)->{'value0'});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "right" => function($v_6) use ($Functor0_2_1, $pure1_4_3) {
  $__num = \func_num_args();
  $__local_var_7_9 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure1_4_3))($GLOBALS['Data_Either_Left']);
  $__local_var_8_10 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_2_1)->{'map'})($GLOBALS['Data_Either_Right'])))($v_6);
  $__res = function($v2_9) use ($__local_var_7_9, $__local_var_8_10) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v2_9 instanceof \Data\Either\Data_Either_Left) {
$__t11 = ($__local_var_7_9)(($v2_9)->{'value0'});
goto end_branch_11;;
};
  if ($v2_9 instanceof \Data\Either\Data_Either_Right) {
$__t11 = ($__local_var_8_10)(($v2_9)->{'value0'});
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Profunctor0" => function($_dollar___unused_6) use ($profunctorStar1_5_4) {
  $__num = \func_num_args();
  $__res = $profunctorStar1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_choiceStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_choicemajStar';

// Data_Profunctor_Star_categoryStar
function majData_majProfunctor_majStar_categorymajStar($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_categorymajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonad_0)->{'Bind1'})(null);
  $semigroupoidStar1_1_0 = (object)["compose" => function($v_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($__local_var_1_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($x_4) use ($__local_var_1_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'bind'})(($v1_3)($x_4)))($v_2);
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
  $__res = (object)["identity" => ((($dictMonad_0)->{'Applicative0'})(null))->{'pure'}, "Semigroupoid0" => function($_dollar___unused_2) use ($semigroupoidStar1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupoidStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_categoryStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_categorymajStar';

// Data_Profunctor_Star_applyStar
function majData_majProfunctor_majStar_applymajStar($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_applymajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $functorStar1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_1_0)->{'map'})($f_2)))($v_3);
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
  $__res = function($a_4) use ($dictApply_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(($v_2)($a_4)))(($v1_3)($a_4));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorStar1_1_0) {
  $__num = \func_num_args();
  $__res = $functorStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_applyStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_applymajStar';

// Data_Profunctor_Star_bindStar
function majData_majProfunctor_majStar_bindmajStar($dictBind_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_bindmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBind_0)->{'Apply0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorStar1_2_1 = (object)["map" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_2_1)->{'map'})($f_3)))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar1_1_0 = (object)["apply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($__local_var_1_0, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(($v_3)($a_5)))(($v1_4)($a_5));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorStar1_2_1) {
  $__num = \func_num_args();
  $__res = $functorStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => function($v_2) use ($dictBind_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictBind_0, $v_2) {
  $__num = \func_num_args();
  $__res = function($x_4) use ($dictBind_0, $f_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictBind_0)->{'bind'})(($v_2)($x_4)))(function($a_5) use ($f_3, $x_4) {
  $__num = \func_num_args();
  $__res = (($f_3)($a_5))($x_4);
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
}, "Apply0" => function($_dollar___unused_2) use ($applyStar1_1_0) {
  $__num = \func_num_args();
  $__res = $applyStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_bindStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_bindmajStar';

// Data_Profunctor_Star_applicativeStar
function majData_majProfunctor_majStar_applicativemajStar($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_applicativemajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorStar1_2_1 = (object)["map" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_2_1)->{'map'})($f_3)))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar1_1_0 = (object)["apply" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($__local_var_1_0, $v_3) {
  $__num = \func_num_args();
  $__res = function($a_5) use ($__local_var_1_0, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(($v_3)($a_5)))(($v1_4)($a_5));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorStar1_2_1) {
  $__num = \func_num_args();
  $__res = $functorStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pure" => function($a_2) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($a_2, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)->{'pure'})($a_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_2) use ($applyStar1_1_0) {
  $__num = \func_num_args();
  $__res = $applyStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_applicativeStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_applicativemajStar';

// Data_Profunctor_Star_monadStar
function majData_majProfunctor_majStar_monadmajStar($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_monadmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Apply0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $functorStar1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_3_2)->{'map'})($f_4)))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar1_2_1 = (object)["apply" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($__local_var_2_1, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'apply'})(($v_4)($a_6)))(($v1_5)($a_6));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorStar1_3_2) {
  $__num = \func_num_args();
  $__res = $functorStar1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeStar1_1_0 = (object)["pure" => function($a_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $a_3) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)->{'pure'})($a_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_3) use ($applyStar1_2_1) {
  $__num = \func_num_args();
  $__res = $applyStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_6 = (($dictMonad_0)->{'Bind1'})(null);
  $__local_var_3_7 = (($__local_var_2_6)->{'Apply0'})(null);
  $__local_var_4_8 = (($__local_var_3_7)->{'Functor0'})(null);
  $functorStar1_4_8 = (object)["map" => function($f_5) use ($__local_var_4_8) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_8, $f_5) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_4_8)->{'map'})($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar1_3_7 = (object)["apply" => function($v_5) use ($__local_var_3_7) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_3_7, $v_5) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($__local_var_3_7, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_7)->{'apply'})(($v_5)($a_7)))(($v1_6)($a_7));
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
}, "Functor0" => function($_dollar___unused_5) use ($functorStar1_4_8) {
  $__num = \func_num_args();
  $__res = $functorStar1_4_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindStar1_2_6 = (object)["bind" => function($v_4) use ($__local_var_2_6) {
  $__num = \func_num_args();
  $__res = function($f_5) use ($__local_var_2_6, $v_4) {
  $__num = \func_num_args();
  $__res = function($x_6) use ($__local_var_2_6, $f_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_6)->{'bind'})(($v_4)($x_6)))(function($a_7) use ($f_5, $x_6) {
  $__num = \func_num_args();
  $__res = (($f_5)($a_7))($x_6);
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
}, "Apply0" => function($_dollar___unused_4) use ($applyStar1_3_7) {
  $__num = \func_num_args();
  $__res = $applyStar1_3_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeStar1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_3) use ($bindStar1_2_6) {
  $__num = \func_num_args();
  $__res = $bindStar1_2_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_monadStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_monadmajStar';

// Data_Profunctor_Star_altStar
function majData_majProfunctor_majStar_altmajStar($dictAlt_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_altmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictAlt_0)->{'Functor0'})(null);
  $functorStar1_1_0 = (object)["map" => function($f_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_1_0)->{'map'})($f_2)))($v_3);
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
  $__res = function($a_4) use ($dictAlt_0, $v1_3, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictAlt_0)->{'alt'})(($v_2)($a_4)))(($v1_3)($a_4));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorStar1_1_0) {
  $__num = \func_num_args();
  $__res = $functorStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_altStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_altmajStar';

// Data_Profunctor_Star_plusStar
function majData_majProfunctor_majStar_plusmajStar($dictPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_plusmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $empty_1_0 = ($dictPlus_0)->{'empty'};
  $__local_var_2_1 = (($dictPlus_0)->{'Alt0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $functorStar1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_3_2)->{'map'})($f_4)))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altStar1_2_1 = (object)["alt" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($__local_var_2_1, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'alt'})(($v_4)($a_6)))(($v1_5)($a_6));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorStar1_3_2) {
  $__num = \func_num_args();
  $__res = $functorStar1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => function($v_3) use ($empty_1_0) {
  $__num = \func_num_args();
  $__res = $empty_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_3) use ($altStar1_2_1) {
  $__num = \func_num_args();
  $__res = $altStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_plusStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_plusmajStar';

// Data_Profunctor_Star_alternativeStar
function majData_majProfunctor_majStar_alternativemajStar($dictAlternative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_alternativemajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictAlternative_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Apply0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $functorStar1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_3_2)->{'map'})($f_4)))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar1_2_1 = (object)["apply" => function($v_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($__local_var_2_1, $v_4) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($__local_var_2_1, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'apply'})(($v_4)($a_6)))(($v1_5)($a_6));
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
}, "Functor0" => function($_dollar___unused_4) use ($functorStar1_3_2) {
  $__num = \func_num_args();
  $__res = $functorStar1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeStar1_1_0 = (object)["pure" => function($a_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $a_3) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)->{'pure'})($a_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_3) use ($applyStar1_2_1) {
  $__num = \func_num_args();
  $__res = $applyStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_6 = (($dictAlternative_0)->{'Plus1'})(null);
  $empty_3_7 = ($__local_var_2_6)->{'empty'};
  $__local_var_4_8 = (($__local_var_2_6)->{'Alt0'})(null);
  $__local_var_5_9 = (($__local_var_4_8)->{'Functor0'})(null);
  $functorStar1_5_9 = (object)["map" => function($f_6) use ($__local_var_5_9) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_9, $f_6) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_5_9)->{'map'})($f_6)))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altStar1_4_8 = (object)["alt" => function($v_6) use ($__local_var_4_8) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_8, $v_6) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($__local_var_4_8, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_8)->{'alt'})(($v_6)($a_8)))(($v1_7)($a_8));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStar1_5_9) {
  $__num = \func_num_args();
  $__res = $functorStar1_5_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusStar1_2_6 = (object)["empty" => function($v_5) use ($empty_3_7) {
  $__num = \func_num_args();
  $__res = $empty_3_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_5) use ($altStar1_4_8) {
  $__num = \func_num_args();
  $__res = $altStar1_4_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Applicative0" => function($_dollar___unused_3) use ($applicativeStar1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_3) use ($plusStar1_2_6) {
  $__num = \func_num_args();
  $__res = $plusStar1_2_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_alternativeStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_alternativemajStar';

// Data_Profunctor_Star_monadPlusStar
function majData_majProfunctor_majStar_monadmajPlusmajStar($dictMonadPlus_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majProfunctor_majStar_monadmajPlusmajStar';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonadPlus_0)->{'Monad0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Applicative0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Apply0'})(null);
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
  $applyStar1_3_2 = (object)["apply" => function($v_5) use ($__local_var_3_2) {
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
  $applicativeStar1_2_1 = (object)["pure" => function($a_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_2_1, $a_4) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)->{'pure'})($a_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_4) use ($applyStar1_3_2) {
  $__num = \func_num_args();
  $__res = $applyStar1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_3_7 = (($__local_var_1_0)->{'Bind1'})(null);
  $__local_var_4_8 = (($__local_var_3_7)->{'Apply0'})(null);
  $__local_var_5_9 = (($__local_var_4_8)->{'Functor0'})(null);
  $functorStar1_5_9 = (object)["map" => function($f_6) use ($__local_var_5_9) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_9, $f_6) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_5_9)->{'map'})($f_6)))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar1_4_8 = (object)["apply" => function($v_6) use ($__local_var_4_8) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_8, $v_6) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($__local_var_4_8, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_8)->{'apply'})(($v_6)($a_8)))(($v1_7)($a_8));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStar1_5_9) {
  $__num = \func_num_args();
  $__res = $functorStar1_5_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $bindStar1_3_7 = (object)["bind" => function($v_5) use ($__local_var_3_7) {
  $__num = \func_num_args();
  $__res = function($f_6) use ($__local_var_3_7, $v_5) {
  $__num = \func_num_args();
  $__res = function($x_7) use ($__local_var_3_7, $f_6, $v_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_7)->{'bind'})(($v_5)($x_7)))(function($a_8) use ($f_6, $x_7) {
  $__num = \func_num_args();
  $__res = (($f_6)($a_8))($x_7);
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
}, "Apply0" => function($_dollar___unused_5) use ($applyStar1_4_8) {
  $__num = \func_num_args();
  $__res = $applyStar1_4_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $monadStar1_1_0 = (object)["Applicative0" => function($_dollar___unused_4) use ($applicativeStar1_2_1) {
  $__num = \func_num_args();
  $__res = $applicativeStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_4) use ($bindStar1_3_7) {
  $__num = \func_num_args();
  $__res = $bindStar1_3_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_14 = (($dictMonadPlus_0)->{'Alternative1'})(null);
  $__local_var_3_15 = (($__local_var_2_14)->{'Applicative0'})(null);
  $__local_var_4_16 = (($__local_var_3_15)->{'Apply0'})(null);
  $__local_var_5_17 = (($__local_var_4_16)->{'Functor0'})(null);
  $functorStar1_5_17 = (object)["map" => function($f_6) use ($__local_var_5_17) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_17, $f_6) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_5_17)->{'map'})($f_6)))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyStar1_4_16 = (object)["apply" => function($v_6) use ($__local_var_4_16) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($__local_var_4_16, $v_6) {
  $__num = \func_num_args();
  $__res = function($a_8) use ($__local_var_4_16, $v1_7, $v_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_16)->{'apply'})(($v_6)($a_8)))(($v1_7)($a_8));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorStar1_5_17) {
  $__num = \func_num_args();
  $__res = $functorStar1_5_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applicativeStar1_3_15 = (object)["pure" => function($a_5) use ($__local_var_3_15) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_3_15, $a_5) {
  $__num = \func_num_args();
  $__res = (($__local_var_3_15)->{'pure'})($a_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_5) use ($applyStar1_4_16) {
  $__num = \func_num_args();
  $__res = $applyStar1_4_16;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_21 = (($__local_var_2_14)->{'Plus1'})(null);
  $empty_5_22 = ($__local_var_4_21)->{'empty'};
  $__local_var_6_23 = (($__local_var_4_21)->{'Alt0'})(null);
  $__local_var_7_24 = (($__local_var_6_23)->{'Functor0'})(null);
  $functorStar1_7_24 = (object)["map" => function($f_8) use ($__local_var_7_24) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($__local_var_7_24, $f_8) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_7_24)->{'map'})($f_8)))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altStar1_6_23 = (object)["alt" => function($v_8) use ($__local_var_6_23) {
  $__num = \func_num_args();
  $__res = function($v1_9) use ($__local_var_6_23, $v_8) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($__local_var_6_23, $v1_9, $v_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_23)->{'alt'})(($v_8)($a_10)))(($v1_9)($a_10));
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
}, "Functor0" => function($_dollar___unused_8) use ($functorStar1_7_24) {
  $__num = \func_num_args();
  $__res = $functorStar1_7_24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusStar1_4_21 = (object)["empty" => function($v_7) use ($empty_5_22) {
  $__num = \func_num_args();
  $__res = $empty_5_22;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($_dollar___unused_7) use ($altStar1_6_23) {
  $__num = \func_num_args();
  $__res = $altStar1_6_23;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeStar1_2_14 = (object)["Applicative0" => function($_dollar___unused_5) use ($applicativeStar1_3_15) {
  $__num = \func_num_args();
  $__res = $applicativeStar1_3_15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_5) use ($plusStar1_4_21) {
  $__num = \func_num_args();
  $__res = $plusStar1_4_21;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["Monad0" => function($_dollar___unused_3) use ($monadStar1_1_0) {
  $__num = \func_num_args();
  $__res = $monadStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar___unused_3) use ($alternativeStar1_2_14) {
  $__num = \func_num_args();
  $__res = $alternativeStar1_2_14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Profunctor_Star_monadPlusStar'] = __NAMESPACE__ . '\\majData_majProfunctor_majStar_monadmajPlusmajStar';

