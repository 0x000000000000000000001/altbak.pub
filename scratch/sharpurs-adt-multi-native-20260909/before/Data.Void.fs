[<AutoOpen>]
module PureScript_Data_Void

open System
open System.Collections.Generic

let Data_Void_Void  = (box (fun (x: obj) -> (box x)))

let Data_Void_absurd  = (box (fun (a: obj) -> (
                                                                                                                                                                                                        
                                                                                                                                                                                                        let rec spin_tco (v: obj) : obj = ((match ((unbox ((box v)))) with | b -> ((spin_tco ((box b)))))) 
                                                                                                                                                                                                        and spin = box ((fun (v: obj) -> spin_tco v)) 
                                                                                                                                                                                                        in
                                                                                                                                                                                                        (spin_tco ((box a)))
                                                                                                                                                                                                        )))
