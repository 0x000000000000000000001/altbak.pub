let rec f_tco (x: obj) (y: obj) : obj = box x
and f = box (fun (x: obj) -> (fun (y: obj) -> f_tco x y))
