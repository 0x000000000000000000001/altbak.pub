<?php

namespace Control\Monad\Gen\Common;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Gen.Common, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Either, Data.Functor, Data.Identity, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Tuple, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Monad.Gen, Control.Monad.Gen.Class, Control.Monad.Gen.Common, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Either, Data.Functor, Data.Identity, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Tuple, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.Gen/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Gen.Common/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
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


// Control_Monad_Gen_Common_max
$GLOBALS['Control_Monad_Gen_Common_max'] = (function() {
  $__fn = function($x_0 = null, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = ((($GLOBALS['Data_Ord_ordInt'])['compare'])($x_0))($y_1);
  $__t1 = null;;
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "LT"))) {
$__t1 = $y_1;
goto end_branch_1;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "EQ"))) {
$__t1 = $x_0;
goto end_branch_1;;
};
  if ((is_object($v_2_0) && (($v_2_0)->{'tag'} === "GT"))) {
$__t1 = $x_0;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Gen_Common_genTuple
$GLOBALS['Control_Monad_Gen_Common_genTuple'] = function($dictApply_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Apply_lift2'])($dictApply_0))($GLOBALS['Data_Tuple_Tuple']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Gen_Common_genNonEmpty
$GLOBALS['Control_Monad_Gen_Common_genNonEmpty'] = (function() {
  $__fn = function($dictMonadRec_0 = null, $dictMonadGen_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Apply0_2_0 = (((((($dictMonadGen_1)['Monad0'])($GLOBALS['Prim_undefined']))['Bind1'])($GLOBALS['Prim_undefined']))['Apply0'])($GLOBALS['Prim_undefined']);
  $unfoldable1_3_1 = (($GLOBALS['Control_Monad_Gen_unfoldable'])($dictMonadRec_0))($dictMonadGen_1);
  $__res = function($dictUnfoldable_4 = null) use ($Apply0_2_0, $dictMonadGen_1, $unfoldable1_3_1) {
  $__num = \func_num_args();
  $unfoldable2_5_2 = ($unfoldable1_3_1)($dictUnfoldable_4);
  $__res = function($gen_6 = null) use ($Apply0_2_0, $dictMonadGen_1, $unfoldable2_5_2) {
  $__num = \func_num_args();
  $__res = ((($Apply0_2_0)['apply'])(((((($Apply0_2_0)['Functor0'])($GLOBALS['Prim_undefined']))['map'])($GLOBALS['Data_NonEmpty_NonEmpty']))($gen_6)))(((($dictMonadGen_1)['resize'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Control_Monad_Gen_Common_max'])(0)))(function($v_7 = null) {
  $__num = \func_num_args();
  $__res = ($v_7 - 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($unfoldable2_5_2)($gen_6)));
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Gen_Common_genMaybe'
$GLOBALS['Control_Monad_Gen_Common_genMaybe__prime__'] = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadGen_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $Bind1_2_1 = (($Monad0_1_0)['Bind1'])($GLOBALS['Prim_undefined']);
  $__res = (function() use ($Bind1_2_1, $Monad0_1_0, $dictMonadGen_0) {
  $__fn = function($bias_3 = null, $gen_4 = null) use ($Bind1_2_1, $Monad0_1_0, $dictMonadGen_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Bind1_2_1)['bind'])(((($dictMonadGen_0)['chooseFloat'])(0.0))(1.0)))(function($n_5 = null) use ($Bind1_2_1, $Monad0_1_0, $bias_3, $gen_4) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($n_5 < $bias_3)) {
$__t2 = ((((((($Bind1_2_1)['Apply0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']))['map'])($GLOBALS['Data_Maybe_Just']))($gen_4);
goto end_branch_2;;
};
  $__t2 = (((($Monad0_1_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data0("Nothing"));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Gen_Common_genMaybe
$GLOBALS['Control_Monad_Gen_Common_genMaybe'] = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_Gen_Common_genMaybe__prime__'])($dictMonadGen_0))(0.75);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Gen_Common_genIdentity
$GLOBALS['Control_Monad_Gen_Common_genIdentity'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = (($dictFunctor_0)['map'])($GLOBALS['Data_Identity_Identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Gen_Common_genEither'
$GLOBALS['Control_Monad_Gen_Common_genEither__prime__'] = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $Bind1_1_0 = (((($dictMonadGen_0)['Monad0'])($GLOBALS['Prim_undefined']))['Bind1'])($GLOBALS['Prim_undefined']);
  $__local_var_2_1 = (((($Bind1_1_0)['Apply0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']);
  $__res = (function() use ($Bind1_1_0, $__local_var_2_1, $dictMonadGen_0) {
  $__fn = function($bias_3 = null, $genA_4 = null, $genB_5 = null) use ($Bind1_1_0, $__local_var_2_1, $dictMonadGen_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Bind1_1_0)['bind'])(((($dictMonadGen_0)['chooseFloat'])(0.0))(1.0)))(function($n_6 = null) use ($__local_var_2_1, $bias_3, $genA_4, $genB_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($n_6 < $bias_3)) {
$__t2 = ((($__local_var_2_1)['map'])($GLOBALS['Data_Either_Left']))($genA_4);
goto end_branch_2;;
};
  $__t2 = ((($__local_var_2_1)['map'])($GLOBALS['Data_Either_Right']))($genB_5);
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Gen_Common_genEither
$GLOBALS['Control_Monad_Gen_Common_genEither'] = function($dictMonadGen_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_Gen_Common_genEither__prime__'])($dictMonadGen_0))(0.5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

