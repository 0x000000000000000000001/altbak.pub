<?php
$exports['runRecordsFFI'] = function($limit) {
    $n = (int)$limit;
    $rec = (object)["a" => (object)["b" => (object)["c" => (object)["d" => (object)["e" => 0]]]]];
    $go = function($i, $r) use (&$go, $n) {
        if ($i >= $n) return $r;
        $newR = (object)[
            "a" => (object)[
                "b" => (object)[
                    "c" => (object)[
                        "d" => (object)[
                            "e" => $r->a->b->c->d->e + 1
                        ]
                    ]
                ]
            ]
        ];
        return $go($i + 1, $newR);
    };
    $finalRec = $go(0, $rec);
    return $finalRec->a->b->c->d->e;
};
return $exports;
