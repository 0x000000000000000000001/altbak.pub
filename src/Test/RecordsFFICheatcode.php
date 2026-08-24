<?php
$exports['runRecordsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $rec = (object)["a" => (object)["b" => (object)["c" => (object)["d" => (object)["e" => 0]]]]];
    for ($i = 0; $i < $n; $i++) {
        $rec = (object)[
            "a" => (object)[
                "b" => (object)[
                    "c" => (object)[
                        "d" => (object)[
                            "e" => $rec->a->b->c->d->e + 1
                        ]
                    ]
                ]
            ]
        ];
    }
    return $rec->a->b->c->d->e;
};
return $exports;
