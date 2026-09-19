pub mod PureScript_Control_Comonad_Env {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use crate::module_ca268cc0::PureScript_Control_Comonad_Env_Trans;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Env_unwrap() -> &dyn Any {
        static Control_Comonad_Env_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_unwrap.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                    &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Comonad_Env_withEnv() -> &dyn Any {
        static Control_Comonad_Env_withEnv: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_withEnv.get_or_init(||
                                                    &PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_withEnvT())
    }
    pub fn Control_Comonad_Env_runEnv() -> &dyn Any {
        static Control_Comonad_Env_runEnv: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_runEnv.get_or_init(||
                                                   &Func1::new(move |v|
                                                                   {
                                                                       let x =
                                                                           Sharpurs_Prelude::unbox(v);
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                              &&&PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                                                                           &&&PureScript_Control_Comonad_Env::Control_Comonad_Env_unwrap()),
                                                                                                        &&&x)
                                                                   }))
    }
    pub fn Control_Comonad_Env_mapEnv() -> &dyn Any {
        static Control_Comonad_Env_mapEnv: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_mapEnv.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_functorEnvT(),
                                                                                                                       &&&PureScript_Data_Identity::Data_Identity_functorIdentity())))
    }
    pub fn Control_Comonad_Env_env() -> &dyn Any {
        static Control_Comonad_Env_env: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_env.get_or_init(||
                                                &Func1::new(move |e|
                                                                &Func1::new({
                                                                                let e
                                                                                    =
                                                                                    e.clone();
                                                                                move
                                                                                    |a|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                        &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                |usd__arg1|
                                                                                                                                                                                                                                                Func1::new({
                                                                                                                                                                                                                                                               let usd__arg1
                                                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                                                   usd__arg1.clone();
                                                                                                                                                                                                                                                               move
                                                                                                                                                                                                                                                                   |usd__arg2|
                                                                                                                                                                                                                                                                   &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                           usd__arg2.clone()))
                                                                                                                                                                                                                                                           })),
                                                                                                                                                                                                                              &&&e)),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_Identity(),
                                                                                                                                                                                           a)))
                                                                            })))
    }
}
