[<AutoOpen>]
module PureScript_Effect_Exception_Unsafe

open System
open System.Collections.Generic

let Effect_Exception_Unsafe_unsafeThrowException  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Unsafe_unsafePerformEffect)))))) (box ((box Effect_Exception_throwException))))

let Effect_Exception_Unsafe_unsafeThrow  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Exception_Unsafe_unsafeThrowException)))))) (box ((box Effect_Exception_error))))
