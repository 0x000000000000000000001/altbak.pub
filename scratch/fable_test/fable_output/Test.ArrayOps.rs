pub mod PureScript_Test_ArrayOps {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_2d8e16c::PureScript_Data_Array;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Test_ArrayOps_add() -> &dyn Any {
        static Test_ArrayOps_add: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOps_add.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                           &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()))
    }
    pub fn Test_ArrayOps_range() -> &dyn Any {
        static Test_ArrayOps_range: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOps_range.get_or_init(||
                                            &Func1::new(move |start|
                                                            &Func1::new({
                                                                            let start
                                                                                =
                                                                                start.clone();
                                                                            move
                                                                                |end_var|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_range(),
                                                                                                                                                    &&&start),
                                                                                                                 end_var)
                                                                        })))
    }
    pub fn Test_ArrayOps_filterEvens() -> &dyn Any {
        static Test_ArrayOps_filterEvens: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOps_filterEvens.get_or_init(||
                                                  &Func1::new(move |arr|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_filter(),
                                                                                                                                      &&&Func1::new(move
                                                                                                                                                        |x|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                                                                                                                                  x),
                                                                                                                                                                                                                                                               &&&2_i32)),
                                                                                                                                                                                         &&&0_i32))),
                                                                                                   arr)))
    }
    pub fn Test_ArrayOps_sumEvens() -> &dyn Any {
        static Test_ArrayOps_sumEvens: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOps_sumEvens.get_or_init(||
                                               &Func1::new(move |n|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_foldl(),
                                                                                                                                                                      &&&PureScript_Test_ArrayOps::Test_ArrayOps_add()),
                                                                                                                                   &&&0_i32),
                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ArrayOps::Test_ArrayOps_filterEvens(),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ArrayOps::Test_ArrayOps_range(),
                                                                                                                                                                                                         &&&1_i32),
                                                                                                                                                                      n)))))
    }
    pub fn Test_ArrayOps_describe() -> &dyn Any {
        static Test_ArrayOps_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOps_describe.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                &&&string("Array Processing (900 elements):")))
    }
    pub fn Test_ArrayOps_act() -> &dyn Any {
        static Test_ArrayOps_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOps_act.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                 &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                 &&&900_i32)),
                                                                           &&&Func1::new(move
                                                                                             |dummy|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                 &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                    &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ArrayOps::Test_ArrayOps_sumEvens(),
                                                                                                                                                                                                    dummy))))))
    }
}
