pub mod PureScript_Test_StateMonad {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Test_StateMonad_State() -> &dyn Any {
        static Test_StateMonad_State: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_State.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Test_StateMonad_runState() -> &dyn Any {
        static Test_StateMonad_runState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_runState.get_or_init(||
                                                 &Func1::new(move |v|
                                                                 &Func1::new({
                                                                                 let v
                                                                                     =
                                                                                     v.clone();
                                                                                 move
                                                                                     |s|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&v),
                                                                                                                      &&&Sharpurs_Prelude::unbox(s))
                                                                             })))
    }
    pub fn Test_StateMonad_put() -> &dyn Any {
        static Test_StateMonad_put: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_put.get_or_init(||
                                            &Func1::new(move |s|
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_State(),
                                                                                             &&&Func1::new({
                                                                                                               let s
                                                                                                                   =
                                                                                                                   s.clone();
                                                                                                               move
                                                                                                                   |v|
                                                                                                                   &add(string("val"),
                                                                                                                        &&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                                                        add(string("state"),
                                                                                                                            &&s,
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>()))
                                                                                                           }))))
    }
    pub fn Test_StateMonad_pureState() -> &dyn Any {
        static Test_StateMonad_pureState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_pureState.get_or_init(||
                                                  &Func1::new(move |a|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_State(),
                                                                                                   &&&Func1::new({
                                                                                                                     let a
                                                                                                                         =
                                                                                                                         a.clone();
                                                                                                                     move
                                                                                                                         |s|
                                                                                                                         &add(string("val"),
                                                                                                                              &&a,
                                                                                                                              add(string("state"),
                                                                                                                                  s.clone(),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>()))
                                                                                                                 }))))
    }
    pub fn Test_StateMonad_get() -> &dyn Any {
        static Test_StateMonad_get: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_get.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_State(),
                                                                             &&&Func1::new(move
                                                                                               |s|
                                                                                               &add(string("val"),
                                                                                                    s.clone(),
                                                                                                    add(string("state"),
                                                                                                        s.clone(),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))))
    }
    pub fn Test_StateMonad_describe() -> &dyn Any {
        static Test_StateMonad_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_describe.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                  &&&string("State Monad (1.2k Binds, 60 Stack Depth):")))
    }
    pub fn Test_StateMonad_bindState() -> &dyn Any {
        static Test_StateMonad_bindState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_bindState.get_or_init(||
                                                  &Func1::new(move |v|
                                                                  &Func1::new({
                                                                                  let v
                                                                                      =
                                                                                      v.clone();
                                                                                  move
                                                                                      |g|
                                                                                      {
                                                                                          let matchValue =
                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                          let matchValue_1 =
                                                                                              Sharpurs_Prelude::unbox(g);
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_State(),
                                                                                                                           &&&Func1::new({
                                                                                                                                             let matchValue_1
                                                                                                                                                 =
                                                                                                                                                 matchValue_1.clone();
                                                                                                                                             move
                                                                                                                                                 |s|
                                                                                                                                                 {
                                                                                                                                                     let r1 =
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                          s);
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                    &&find(string("val"),
                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&r1)))),
                                                                                                                                                                                      &&find(string("state"),
                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&r1)))
                                                                                                                                                 }
                                                                                                                                         }))
                                                                                      }
                                                                              })))
    }
    pub fn Test_StateMonad_modify() -> &dyn Any {
        static Test_StateMonad_modify: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_modify.get_or_init(||
                                               &Func1::new(move |f|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_bindState(),
                                                                                                                                   &&&PureScript_Test_StateMonad::Test_StateMonad_get()),
                                                                                                &&&Func1::new({
                                                                                                                  let f
                                                                                                                      =
                                                                                                                      f.clone();
                                                                                                                  move
                                                                                                                      |s|
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_put(),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                          s))
                                                                                                              }))))
    }
    pub fn Test_StateMonad_chainModifications_004024() -> &dyn Any {
        &Func1::new(move |v|
                        PureScript_Test_StateMonad::Test_StateMonad_chainModifications_tco(v))
    }
    pub fn Test_StateMonad_chainModifications_004024_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_StateMonad_chainModifications_004024_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_StateMonad_chainModifications_004024_002d1.get_or_init(||
                                                                        Lazy(Test_StateMonad_chainModifications_004024.clone()))
    }
    pub fn Test_StateMonad_chainModifications_tco(v: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32, &matchValue) {
            0_i32 =>
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_pureState(),
                                             &&&PureScript_Data_Unit::Data_Unit_unit()),
            _ =>
            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_bindState(),
                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_modify(),
                                                                                                                   &&&Func1::new(move
                                                                                                                                     |x|
                                                                                                                                     &(Sharpurs_Prelude::unbox(x)
                                                                                                                                           +
                                                                                                                                           Sharpurs_Prelude::unbox(&&&1_i32))))),
                                             &&&Func1::new(move |v1|
                                                               PureScript_Test_StateMonad::Test_StateMonad_chainModifications_tco(&&(Sharpurs_Prelude::unbox(&&&matchValue)
                                                                                                                                         -
                                                                                                                                         Sharpurs_Prelude::unbox(&&&1_i32))))),
        }
    }
    pub fn Test_StateMonad_chainModifications() -> &dyn Any {
        static Test_StateMonad_chainModifications: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_chainModifications.get_or_init(||
                                                           Test_StateMonad_chainModifications_004024_002d1.Value)
    }
    pub fn Test_StateMonad_runManyTimes_004028() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Test_StateMonad::Test_StateMonad_runManyTimes_tco(&v,
                                                                                                        v1)
                                   }))
    }
    pub fn Test_StateMonad_runManyTimes_004028_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_StateMonad_runManyTimes_004028_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_StateMonad_runManyTimes_004028_002d1.get_or_init(||
                                                                  Lazy(Test_StateMonad_runManyTimes_004028.clone()))
    }
    pub fn Test_StateMonad_runManyTimes_tco(v: &dyn Any, v1: &dyn Any)
     -> &dyn Any {
        let v = v.clone();
        let v1 = v1.clone();
        '_Test_StateMonad_runManyTimes_tco:
            loop  {
                break '_Test_StateMonad_runManyTimes_tco
                    ({
                         let matchValue = Sharpurs_Prelude::unbox(&&v);
                         let matchValue_1 = Sharpurs_Prelude::unbox(&&v1);
                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                         &matchValue)
                             {
                             0_i32 => acc.clone(),
                             _ => {
                                 let v_temp =
                                     &(Sharpurs_Prelude::unbox(&&&matchValue)
                                           -
                                           Sharpurs_Prelude::unbox(&&&1_i32));
                                 let v1_temp =
                                     &(Sharpurs_Prelude::unbox(&&&matchValue_1)
                                           +
                                           Sharpurs_Prelude::unbox(&&find(string("state"),
                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_runState(),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_chainModifications(),
                                                                                                                                                                                                          &&&60_i32)),
                                                                                                                                    &&&0_i32)))));
                                 v.set(v_temp);
                                 v1.set(v1_temp);
                                 continue '_Test_StateMonad_runManyTimes_tco
                             }
                         }
                     }) ;
            }
    }
    pub fn Test_StateMonad_runManyTimes() -> &dyn Any {
        static Test_StateMonad_runManyTimes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_runManyTimes.get_or_init(||
                                                     Test_StateMonad_runManyTimes_004028_002d1.Value)
    }
    pub fn Test_StateMonad_act() -> &dyn Any {
        static Test_StateMonad_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonad_act.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                   &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                   &&&20_i32)),
                                                                             &&&Func1::new(move
                                                                                               |dummy|
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                   &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                      &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonad::Test_StateMonad_runManyTimes(),
                                                                                                                                                                                                                                         dummy),
                                                                                                                                                                                                      &&&0_i32))))))
    }
}
