pub mod PureScript_Test_BenchCheck {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Test_BenchCheck_act() -> &dyn Any {
        static Test_BenchCheck_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_BenchCheck_act.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                   &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                &&&PureScript_Bench::Bench_benchNow()),
                                                                             &&&Func1::new(move
                                                                                               |t1|
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                      &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                   &&&PureScript_Bench::Bench_benchNow()),
                                                                                                                                &&&Func1::new({
                                                                                                                                                  let t1
                                                                                                                                                      =
                                                                                                                                                      t1.clone();
                                                                                                                                                  move
                                                                                                                                                      |t2|
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                          &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                             &&&string("Delta: ")),
                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Show::Data_Show_showNumber()),
                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Ring::Data_Ring_ringNumber()),
                                                                                                                                                                                                                                                                                                                                   t2),
                                                                                                                                                                                                                                                                                                &&&t1))))
                                                                                                                                              })))))
    }
}
