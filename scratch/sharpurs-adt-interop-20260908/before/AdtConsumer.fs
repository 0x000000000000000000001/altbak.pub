let AdtConsumer_shared  = (box (fun (child: obj) -> (box (AdtPilot_Tusd_Ctor((box AdtPilot_Busd_Ctor), (box child), (box 11), (box child))))))

let AdtConsumer_saturated  = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (box (AdtPilot_Tusd_Ctor((box color), (box left), (box value), (box right))))))))))))

let AdtConsumer_roundTrip  = (box (fun (child: obj) -> (sharpurs_apply (box ((box AdtPilot_depth))) (box ((box (AdtPilot_Tusd_Ctor((box AdtPilot_Busd_Ctor), (box child), (box 11), (box child)))))))))

let AdtConsumer_rootValue  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box 0)) | AdtPilot_Tusd_Ctor(_, _, value, _) -> ((box value)))))

let AdtConsumer_rootColor  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box AdtPilot_Busd_Ctor)) | AdtPilot_Tusd_Ctor(color, _, _, _) -> ((box color)))))

let AdtConsumer_partial  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (fun (usd__arg3: obj) -> (box (AdtPilot_Tusd_Ctor((box AdtPilot_Busd_Ctor), usd__arg1, usd__arg2, usd__arg3))))))))

let AdtConsumer_namedChild  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Tusd_Ctor(_, Unbox((AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Rusd_Ctor), _, _, _) as child)), _, _) -> ((box child)) | _ -> ((box AdtPilot_Eusd_Ctor)))))

let AdtConsumer_leftChild  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box AdtPilot_Eusd_Ctor)) | AdtPilot_Tusd_Ctor(_, child, _, _) -> ((box child)))))

let AdtConsumer_deepPattern  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Rusd_Ctor), Unbox(AdtPilot_Eusd_Ctor), Unbox(LitInt 7 ()), Unbox(AdtPilot_Eusd_Ctor))), Unbox(LitInt 11 ()), Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), Unbox(AdtPilot_Eusd_Ctor), value, Unbox(AdtPilot_Eusd_Ctor)))) -> ((box value)) | _ -> ((box 0)))))

let AdtConsumer_applyValue  = (box (fun (f: obj) -> (box (fun (value: obj) -> (sharpurs_apply (box ((box f))) (box ((box value))))))))

let AdtConsumer_throughGeneric  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_applyValue))) (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (fun (usd__arg3: obj) -> (fun (usd__arg4: obj) -> (box (AdtPilot_Tusd_Ctor(usd__arg1, usd__arg2, usd__arg3, usd__arg4)))))))))))))) (box ((box AdtPilot_Rusd_Ctor)))))) (box ((box AdtPilot_Eusd_Ctor)))))) (box ((box value)))))) (box ((box AdtPilot_Eusd_Ctor))))))

let AdtConsumer_throughPartial  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_applyValue))) (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (AdtPilot_Tusd_Ctor((box AdtPilot_Busd_Ctor), (box AdtPilot_Eusd_Ctor), usd__arg1, usd__arg2)))))))))))) (box ((box value)))))) (box ((box AdtPilot_Eusd_Ctor))))))