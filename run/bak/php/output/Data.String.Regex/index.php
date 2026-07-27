<?php

namespace Data\String\Regex;

// ALL IMPORTS: Control.Semigroupoid, Data.Array.NonEmpty, Data.Either, Data.Function, Data.Maybe, Data.Semigroup, Data.Show, Data.String, Data.String.CodeUnits, Data.String.Pattern, Data.String.Regex, Data.String.Regex.Flags, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Array.NonEmpty, Data.Either, Data.Function, Data.Maybe, Data.Semigroup, Data.Show, Data.String, Data.String.CodeUnits, Data.String.Pattern, Data.String.Regex, Data.String.Regex.Flags, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.String/index.php';
require_once __DIR__ . '/../Data.String.CodeUnits/index.php';
require_once __DIR__ . '/../Data.String.Pattern/index.php';
require_once __DIR__ . '/../Data.String.Regex/index.php';
require_once __DIR__ . '/../Data.String.Regex.Flags/index.php';
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
$ffi_Data_String_Regex = \call_user_func(function() {
  $exports = [];
$showRegexImpl = function($r) use (&$showRegexImpl) {
    return $r->pattern;
};

$regexImpl = function($left, $right = null, $s1 = null, $s2 = null) use (&$regexImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$regexImpl) {
            return $regexImpl(...\array_merge($__args, $more));
        };
    }
    $pattern = '/' . $s1 . '/' . $s2;
    // Strip modifiers that PHP doesn't support
    $php_flags = str_replace(['g', 'y'], '', $s2);
    $pcre = "\x01" . $s1 . "\x01" . $php_flags;
    
    if (@preg_match($pcre, '') === false) {
        return $left(error_get_last()['message'] ?? "Invalid regex");
    }
    return $right((object)["pattern" => $pattern, "pcre" => $pcre, "source" => $s1, "flags" => $s2]);
};

$source = function($r) use (&$source) {
    return $r->source;
};

$flagsImpl = function($r) use (&$flagsImpl) {
    return (object)[
        "multiline" => strpos($r->flags, 'm') !== false,
        "ignoreCase" => strpos($r->flags, 'i') !== false,
        "global" => strpos($r->flags, 'g') !== false,
        "dotAll" => strpos($r->flags, 's') !== false,
        "sticky" => strpos($r->flags, 'y') !== false,
        "unicode" => strpos($r->flags, 'u') !== false
    ];
};

$test = function($r, $s = null) use (&$test) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$test) {
            return $test(...\array_merge($__args, $more));
        };
    }
    return preg_match($r->pcre, $s) === 1;
};

$_match = function($just, $nothing = null, $r = null, $s = null) use (&$_match) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_match) {
            return $_match(...\array_merge($__args, $more));
        };
    }
    if (strpos($r->flags, 'g') !== false) {
        $matched = preg_match_all($r->pcre, $s, $matches);
        if ($matched) {
            $res = [];
            foreach ($matches[0] as $m) {
                $res[] = $m === "" ? $nothing : $just($m);
            }
            return $just($res);
        }
    } else {
        $matched = preg_match($r->pcre, $s, $matches);
        if ($matched) {
            $res = [];
            foreach ($matches as $m) {
                $res[] = $m === "" ? $nothing : $just($m);
            }
            return $just($res);
        }
    }
    return $nothing;
};

$replace = function($r, $s1 = null, $s2 = null) use (&$replace) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$replace) {
            return $replace(...\array_merge($__args, $more));
        };
    }
    $limit = strpos($r->flags, 'g') !== false ? -1 : 1;
    // $s1 in PCRE uses $1 for groups whereas JS uses $1 or \1. We assume s1 is compatible.
    // However, JS replace uses $1, PCRE preg_replace also uses $1.
    return preg_replace($r->pcre, $s1, $s2, $limit);
};

$_replaceBy = function($just, $nothing = null, $r = null, $f = null, $s = null) use (&$_replaceBy) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_replaceBy) {
            return $_replaceBy(...\array_merge($__args, $more));
        };
    }
    $limit = strpos($r->flags, 'g') !== false ? -1 : 1;
    return preg_replace_callback($r->pcre, function($matches) use ($f, $just, $nothing) {
        $match = $matches[0];
        $groups = [];
        for ($i = 1; $i < \count($matches); $i++) {
            $groups[] = (!isset($matches[$i]) || $matches[$i] === "") ? $nothing : $just($matches[$i]);
        }
        $fn = $f($match);
        return $fn($groups);
    }, $s, $limit);
};

$_search = function($just, $nothing = null, $r = null, $s = null) use (&$_search) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_search) {
            return $_search(...\array_merge($__args, $more));
        };
    }
    if (preg_match($r->pcre, $s, $matches, PREG_OFFSET_CAPTURE)) {
        return $just($matches[0][1]);
    }
    return $nothing;
};

$split = function($r, $s = null) use (&$split) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$split) {
            return $split(...\array_merge($__args, $more));
        };
    }
    $limit = strpos($r->flags, 'g') !== false ? -1 : 2;
    return preg_split($r->pcre, $s, $limit);
};

$exports['showRegexImpl'] = $showRegexImpl;
$exports['regexImpl'] = $regexImpl;
$exports['source'] = $source;
$exports['flagsImpl'] = $flagsImpl;
$exports['test'] = $test;
$exports['_match'] = $_match;
$exports['replace'] = $replace;
$exports['_replaceBy'] = $_replaceBy;
$exports['_search'] = $_search;
$exports['split'] = $split;
return $exports;
  return $exports;
});
function majData_majString_majRegex__match($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex__match';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['_match'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_String_Regex__match'] = __NAMESPACE__ . '\\majData_majString_majRegex__match';

function majData_majString_majRegex__replacemajBy($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex__replacemajBy';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['_replaceBy'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_String_Regex__replaceBy'] = __NAMESPACE__ . '\\majData_majString_majRegex__replacemajBy';

function majData_majString_majRegex__search($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex__search';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['_search'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_String_Regex__search'] = __NAMESPACE__ . '\\majData_majString_majRegex__search';

function majData_majString_majRegex_flagsmajImpl($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex_flagsmajImpl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['flagsImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_Regex_flagsImpl'] = __NAMESPACE__ . '\\majData_majString_majRegex_flagsmajImpl';

function majData_majString_majRegex_regexmajImpl($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex_regexmajImpl';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['regexImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Data_String_Regex_regexImpl'] = __NAMESPACE__ . '\\majData_majString_majRegex_regexmajImpl';

function majData_majString_majRegex_replace($v0, $v1 = null, $v2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex_replace';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['replace'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_String_Regex_replace'] = __NAMESPACE__ . '\\majData_majString_majRegex_replace';

function majData_majString_majRegex_showmajRegexmajImpl($v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex_showmajRegexmajImpl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['showRegexImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_Regex_showRegexImpl'] = __NAMESPACE__ . '\\majData_majString_majRegex_showmajRegexmajImpl';

function majData_majString_majRegex_source($v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex_source';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['source'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_String_Regex_source'] = __NAMESPACE__ . '\\majData_majString_majRegex_source';

function majData_majString_majRegex_split($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex_split';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['split'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_Regex_split'] = __NAMESPACE__ . '\\majData_majString_majRegex_split';

function majData_majString_majRegex_test($v0, $v1 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majString_majRegex_test';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_String_Regex;
  $f = ($ffi_Data_String_Regex['test'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_String_Regex_test'] = __NAMESPACE__ . '\\majData_majString_majRegex_test';





// Data_String_Regex_showRegex
$GLOBALS['Data_String_Regex_showRegex'] = (object)["show" => $GLOBALS['Data_String_Regex_showRegexImpl']];

// Data_String_Regex_search_closure
$GLOBALS['Data_String_Regex_search_closure'] = (($GLOBALS['Data_String_Regex__search'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_Regex_search
function majData_majString_majRegex_search($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majRegex_search';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_Regex_search_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Regex_search'] = __NAMESPACE__ . '\\majData_majString_majRegex_search';

// Data_String_Regex_replace'_closure
$GLOBALS['Data_String_Regex_replace__prime___closure'] = (($GLOBALS['Data_String_Regex__replaceBy'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_Regex_replace'
function majData_majString_majRegex_replace__prime__($v_0, $v_1 = null, $v_2 = null): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majRegex_replace__prime__';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_String_Regex_replace__prime___closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_String_Regex_replace__prime__'] = __NAMESPACE__ . '\\majData_majString_majRegex_replace__prime__';

// Data_String_Regex_renderFlags
function majData_majString_majRegex_rendermajFlags($v_0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majRegex_rendermajFlags';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if (($v_0)->{'global'}) {
$__t0 = "g";
goto end_branch_0;;
};
  $__t0 = "";
  end_branch_0:;
  $__t1 = null;;
  if (($v_0)->{'ignoreCase'}) {
$__t1 = "i";
goto end_branch_1;;
};
  $__t1 = "";
  end_branch_1:;
  $__t2 = null;;
  if (($v_0)->{'multiline'}) {
$__t2 = "m";
goto end_branch_2;;
};
  $__t2 = "";
  end_branch_2:;
  $__t3 = null;;
  if (($v_0)->{'dotAll'}) {
$__t3 = "s";
goto end_branch_3;;
};
  $__t3 = "";
  end_branch_3:;
  $__t4 = null;;
  if (($v_0)->{'sticky'}) {
$__t4 = "y";
goto end_branch_4;;
};
  $__t4 = "";
  end_branch_4:;
  $__t5 = null;;
  if (($v_0)->{'unicode'}) {
$__t5 = "u";
goto end_branch_5;;
};
  $__t5 = "";
  end_branch_5:;
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($__t0))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($__t1))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($__t2))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($__t3))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})($__t4))($__t5)))));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_Regex_renderFlags'] = __NAMESPACE__ . '\\majData_majString_majRegex_rendermajFlags';

// Data_String_Regex_regex
function majData_majString_majRegex_regex(string $s_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majRegex_regex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\String\Regex\majData_majString_majRegex_regexmajImpl($GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], $s_0, \Data\String\Regex\majData_majString_majRegex_rendermajFlags($f_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Regex_regex'] = __NAMESPACE__ . '\\majData_majString_majRegex_regex';

// Data_String_Regex_parseFlags
function majData_majString_majRegex_parsemajFlags(string $s_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majRegex_parsemajFlags';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["global" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Maybe_isJust'], ($GLOBALS['Data_String_CodeUnits_indexOf'])("g"), $s_0), "ignoreCase" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Maybe_isJust'], ($GLOBALS['Data_String_CodeUnits_indexOf'])("i"), $s_0), "multiline" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Maybe_isJust'], ($GLOBALS['Data_String_CodeUnits_indexOf'])("m"), $s_0), "dotAll" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Maybe_isJust'], ($GLOBALS['Data_String_CodeUnits_indexOf'])("s"), $s_0), "sticky" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Maybe_isJust'], ($GLOBALS['Data_String_CodeUnits_indexOf'])("y"), $s_0), "unicode" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Maybe_isJust'], ($GLOBALS['Data_String_CodeUnits_indexOf'])("u"), $s_0)];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_Regex_parseFlags'] = __NAMESPACE__ . '\\majData_majString_majRegex_parsemajFlags';

// Data_String_Regex_match_closure
$GLOBALS['Data_String_Regex_match_closure'] = (($GLOBALS['Data_String_Regex__match'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Data_String_Regex_match
function majData_majString_majRegex_match($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majRegex_match';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_String_Regex_match_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_String_Regex_match'] = __NAMESPACE__ . '\\majData_majString_majRegex_match';

// Data_String_Regex_flags_closure
$GLOBALS['Data_String_Regex_flags_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_String_Regex_Flags_RegexFlags']))($GLOBALS['Data_String_Regex_flagsImpl']);

// Data_String_Regex_flags
function majData_majString_majRegex_flags($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majString_majRegex_flags';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_String_Regex_flags_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_String_Regex_flags'] = __NAMESPACE__ . '\\majData_majString_majRegex_flags';

