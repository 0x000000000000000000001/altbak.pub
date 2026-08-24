<?php
$exports['runRecordsFFI'] = function($limit) {
    $n = (int)$limit;
    $r = (object)["a" => 0, "b" => (object)["c" => 0, "d" => (object)["e" => 0, "f" => 0]]];
    $go = function($i, $rec) use (&$go) {
        if ($i === 0) return $rec;
        $newR = (object)[
            "a" => $rec->a + 1,
            "b" => (object)[
                "c" => $rec->b->c + 2,
                "d" => (object)[
                    "e" => $rec->b->d->e + 3,
                    "f" => $rec->b->d->f + ($i % 5)
                ]
            ]
        ];
        return $go($i - 1, $newR);
    };
    $finalRec = $go($n, $r);
    return $finalRec->b->d->f;
};
return $exports;
