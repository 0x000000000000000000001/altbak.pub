<?php
$exports['unfoldr1ArrayImpl'] = function($isNothing) {
    return function($fromJust) use ($isNothing) {
        return function($fst) use ($isNothing, $fromJust) {
            return function($snd) use ($isNothing, $fromJust, $fst) {
                return function($f) use ($isNothing, $fromJust, $fst, $snd) {
                    return function($b) use ($isNothing, $fromJust, $fst, $snd, $f) {
                        $result = [];
                        $value = $b;
                        while (true) {
                            $tuple = $f($value);
                            $result[] = $fst($tuple);
                            $maybe = $snd($tuple);
                            if ($isNothing($maybe)) {
                                return $result;
                            }
                            $value = $fromJust($maybe);
                        }
                    };
                };
            };
        };
    };
};
