type AdtConsumer_ConsumerBox =
  | AdtConsumer_ConsumerBoxusd_Ctor of obj * obj

let AdtConsumer_ConsumerBox  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (AdtConsumer_ConsumerBoxusd_Ctor(usd__arg1, usd__arg2)))))))

let AdtConsumer_wrap  = (box (fun (child: obj) -> (box (AdtConsumer_ConsumerBoxusd_Ctor((box child), (box 42))))))

let AdtConsumer_unwrap  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtConsumer_ConsumerBoxusd_Ctor(child, _) -> ((box child)))))

let AdtConsumer_shared  = (box (fun (child: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((box AdtPilot_B)))))) (box ((box child)))))) (box ((box 11)))))) (box ((box child))))))

let AdtConsumer_saturated  = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((box color)))))) (box ((box left)))))) (box ((box value)))))) (box ((box right))))))))))))

let AdtConsumer_roundTrip  = (box (fun (child: obj) -> (sharpurs_apply (box ((box AdtPilot_depth))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((box AdtPilot_B)))))) (box ((box child)))))) (box ((box 11)))))) (box ((box child)))))))))

let AdtConsumer_rootValue  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box 0)) | AdtPilot_Tusd_Ctor(_, _, value, _) -> ((box value)))))

let AdtConsumer_rootColor  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box AdtPilot_B)) | AdtPilot_Tusd_Ctor(color, _, _, _) -> ((box color)))))

let AdtConsumer_partial  = (sharpurs_apply (box ((box AdtPilot_T))) (box ((box AdtPilot_B))))

let AdtConsumer_orderedPartial  = (box (fun (child: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackColor))) (box ((box 1)))))) (box ((box AdtPilot_B))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackTree))) (box ((box 2)))))) (box ((box child)))))))))

let AdtConsumer_orderedConstruction  = (box (fun (child: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackColor))) (box ((box 1)))))) (box ((box AdtPilot_B))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackTree))) (box ((box 2)))))) (box ((box child))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackInt))) (box ((box 3)))))) (box ((box 11))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackTree))) (box ((box 4)))))) (box ((box child)))))))))

let AdtConsumer_namedChild  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Tusd_Ctor(_, Unbox((AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Rusd_Ctor), _, _, _) as child)), _, _) -> ((box child)) | _ -> ((box AdtPilot_E)))))

let AdtConsumer_leftChild  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box AdtPilot_E)) | AdtPilot_Tusd_Ctor(_, child, _, _) -> ((box child)))))

let AdtConsumer_deepPattern  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Rusd_Ctor), Unbox(AdtPilot_Eusd_Ctor), Unbox(LitInt 7 ()), Unbox(AdtPilot_Eusd_Ctor))), Unbox(LitInt 11 ()), Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), Unbox(AdtPilot_Eusd_Ctor), value, Unbox(AdtPilot_Eusd_Ctor)))) -> ((box value)) | _ -> ((box 0)))))

let AdtConsumer_boxedValue  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtConsumer_ConsumerBoxusd_Ctor(_, value) -> ((box value)))))

let AdtConsumer_boxedPattern  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtConsumer_ConsumerBoxusd_Ctor(Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), _, Unbox(LitInt 11 ()), _)), Unbox(LitInt 42 ())) -> ((box 7)) | _ -> ((box 0)))))

let AdtConsumer_applyValue  = (box (fun (f: obj) -> (box (fun (value: obj) -> (sharpurs_apply (box ((box f))) (box ((box value))))))))

let AdtConsumer_throughGeneric  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_applyValue))) (box ((box AdtPilot_T)))))) (box ((box AdtPilot_R)))))) (box ((box AdtPilot_E)))))) (box ((box value)))))) (box ((box AdtPilot_E))))))

let AdtConsumer_throughPartial  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_applyValue))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((box AdtPilot_B)))))) (box ((box AdtPilot_E))))))))) (box ((box value)))))) (box ((box AdtPilot_E))))))