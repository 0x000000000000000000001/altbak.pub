[<AutoOpen>]
module PureScript_Test_Main

open System
open System.Collections.Generic

let Test_Main_main  = (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((box Data_Unit_unit))))
