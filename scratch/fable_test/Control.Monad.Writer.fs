[<AutoOpen>]
module PureScript_Control_Monad_Writer

open System
open System.Collections.Generic

let Control_Monad_Writer_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Monad_Writer_writer  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Control_Monad_Writer_Trans_WriterT)))))) (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Data_Identity_applicativeIdentity)))))))

let Control_Monad_Writer_runWriter  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))))))) (box ((box Control_Monad_Writer_Trans_runWriterT))))

let Control_Monad_Writer_mapWriter  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_Writer_Trans_mapWriterT))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((box Control_Monad_Writer_unwrap))))))))))))

let Control_Monad_Writer_execWriter  = (box (fun (m: obj) -> (sharpurs_apply (box ((box Data_Tuple_snd))) (box ((sharpurs_apply (box ((box Control_Monad_Writer_runWriter))) (box ((box m)))))))))
