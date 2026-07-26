<?php
$exports['fromNumberImpl'] = function($just) {
  return function($nothing) use ($just) {
    return function($n) use ($just, $nothing) {
      if (\is_int($n) || (\is_float($n) && \floor($n) == $n && !\is_infinite($n) && !\is_nan($n))) {
        return $just((int)$n);
      }
      return $nothing;
    };
  };
};

$exports['toNumber'] = function($n) {
  return (float)$n;
};

$exports['fromStringAsImpl'] = function($just) {
  return function($nothing) use ($just) {
    return function($radix) use ($just, $nothing) {
      return function($s) use ($just, $nothing, $radix) {
        $i = \intval($s, $radix);
        // intval returns 0 on failure for some invalid strings in older PHP,
        // but we should just try to convert back to check or just return $just.
        // Actually, PHP doesn't have a direct equivalent to JS pattern matching here easily,
        // so we'll just check if it's numeric in that base
        if (\preg_match('/^[\+\-]?[0-9a-zA-Z]+$/', $s)) {
            $parsed = \intval($s, $radix);
            // intval bounds checking
            if (\strval($parsed) === $s || \base_convert(\strval($parsed), 10, $radix) === \strtolower(\ltrim($s, '+'))) {
                return $just($parsed);
            }
        }
        return $nothing;
      };
    };
  };
};

$exports['toStringAs'] = function($radix) {
  return function($i) use ($radix) {
    return \base_convert((string)$i, 10, $radix);
  };
};

$exports['quot'] = function($x) {
  return function($y) use ($x) {
    return \intdiv($x, $y);
  };
};

$exports['rem'] = function($x) {
  return function($y) use ($x) {
    return $x % $y;
  };
};

$exports['pow'] = function($x) {
  return function($y) use ($x) {
    return (int)\pow($x, $y);
  };
};
