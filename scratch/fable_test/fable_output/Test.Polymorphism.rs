pub mod PureScript_Test_Polymorphism {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Test_Polymorphism_Monoidishusd_Dict() -> &dyn Any {
        static Test_Polymorphism_Monoidishusd_Dict: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Test_Polymorphism_Monoidishusd_Dict.get_or_init(||
                                                            &Func1::new(move
                                                                            |x|
                                                                            x.clone()))
    }
    pub fn Test_Polymorphism_mempty_() -> &dyn Any {
        static Test_Polymorphism_mempty_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Polymorphism_mempty_.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("mempty_"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Test_Polymorphism_mappend_() -> &dyn Any {
        static Test_Polymorphism_mappend_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Polymorphism_mappend_.get_or_init(||
                                                   &Func1::new(move |dict|
                                                                   find(string("mappend_"),
                                                                        Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Test_Polymorphism_polyLoop() -> &dyn Any {
        static Test_Polymorphism_polyLoop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Polymorphism_polyLoop.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonoidish|
                                                                   &Func1::new({
                                                                                   let dictMonoidish
                                                                                       =
                                                                                       dictMonoidish.clone();
                                                                                   move
                                                                                       |n_init|
                                                                                       &Func1::new({
                                                                                                       let n_init
                                                                                                           =
                                                                                                           n_init.clone();
                                                                                                       move
                                                                                                           |acc_init|
                                                                                                           {
                                                                                                               let go_2 =
                                                                                                                   Func0::new({
                                                                                                                                  let go_tco
                                                                                                                                      =
                                                                                                                                      go_tco.clone();
                                                                                                                                  move
                                                                                                                                      ||
                                                                                                                                      &Func1::new({
                                                                                                                                                      let go_tco
                                                                                                                                                          =
                                                                                                                                                          go_tco.clone();
                                                                                                                                                      move
                                                                                                                                                          |v|
                                                                                                                                                          Func1::new({
                                                                                                                                                                         let go_tco
                                                                                                                                                                             =
                                                                                                                                                                             go_tco.clone();
                                                                                                                                                                         let v
                                                                                                                                                                             =
                                                                                                                                                                             v.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v1|
                                                                                                                                                                             go_tco(v)(v1.clone())
                                                                                                                                                                     })
                                                                                                                                                  })
                                                                                                                              });
                                                                                                               let go_1 =
                                                                                                                   Lazy(go_2);
                                                                                                               fn go_tco(v_1:
                                                                                                                             _)
                                                                                                                ->
                                                                                                                    Func1<&dyn Any,
                                                                                                                          &dyn Any> {
                                                                                                                   Func1::new({
                                                                                                                                  let go_tco
                                                                                                                                      =
                                                                                                                                      go_tco.clone();
                                                                                                                                  let v_1
                                                                                                                                      =
                                                                                                                                      v_1.clone();
                                                                                                                                  move
                                                                                                                                      |v1_1|
                                                                                                                                      {
                                                                                                                                          let matchValue =
                                                                                                                                              Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                          let matchValue_1 =
                                                                                                                                              Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                          match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                                                                                                          &matchValue)
                                                                                                                                              {
                                                                                                                                              0_i32
                                                                                                                                              =>
                                                                                                                                              acc.clone(),
                                                                                                                                              _
                                                                                                                                              =>
                                                                                                                                              go_tco(&(Sharpurs_Prelude::unbox(&&&matchValue)
                                                                                                                                                           -
                                                                                                                                                           Sharpurs_Prelude::unbox(&&&1_i32)))(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Polymorphism::Test_Polymorphism_mappend_(),
                                                                                                                                                                                                                                                                                                      &&&dictMonoidish),
                                                                                                                                                                                                                                                                   &&&matchValue_1),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Polymorphism::Test_Polymorphism_mempty_(),
                                                                                                                                                                                                                                                                   &&&dictMonoidish))),
                                                                                                                                          }
                                                                                                                                      }
                                                                                                                              })
                                                                                                               }
                                                                                                               let go =
                                                                                                                   go_1.Value;
                                                                                                               go_tco(&n_init)(acc_init.clone())
                                                                                                           }
                                                                                                   })
                                                                               })))
    }
    pub fn Test_Polymorphism_intMonoidish() -> &dyn Any {
        static Test_Polymorphism_intMonoidish: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Polymorphism_intMonoidish.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Polymorphism::Test_Polymorphism_Monoidishusd_Dict(),
                                                                                        &&&add(string("mempty_"),
                                                                                               &&1_i32,
                                                                                               add(string("mappend_"),
                                                                                                   &&Func1::new(move
                                                                                                                    |x|
                                                                                                                    &Func1::new({
                                                                                                                                    let x
                                                                                                                                        =
                                                                                                                                        x.clone();
                                                                                                                                    move
                                                                                                                                        |y|
                                                                                                                                        &(Sharpurs_Prelude::unbox(&&&x)
                                                                                                                                              +
                                                                                                                                              Sharpurs_Prelude::unbox(y))
                                                                                                                                })),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Test_Polymorphism_describe() -> &dyn Any {
        static Test_Polymorphism_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Polymorphism_describe.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                    &&&string("Polymorphism (10M Type Class Dict Lookups):")))
    }
    pub fn Test_Polymorphism_act() -> &dyn Any {
        static Test_Polymorphism_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Polymorphism_act.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                     &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                     &&&10000000_i32)),
                                                                               &&&Func1::new(move
                                                                                                 |dummy|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                     &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                        &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Polymorphism::Test_Polymorphism_polyLoop(),
                                                                                                                                                                                                                                                                              &&&PureScript_Test_Polymorphism::Test_Polymorphism_intMonoidish()),
                                                                                                                                                                                                                                           dummy),
                                                                                                                                                                                                        &&&0_i32))))))
    }
}
