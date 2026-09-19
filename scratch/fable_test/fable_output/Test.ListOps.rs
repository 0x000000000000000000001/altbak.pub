pub mod PureScript_Test_ListOps {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    #[derive(Clone, Debug,)]
    pub enum Test_ListOps_List {
        Test_ListOps_Nilusd_Ctor,
        Test_ListOps_Consusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for PureScript_Test_ListOps::Test_ListOps_List {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Test_ListOps_add() -> &dyn Any {
        static Test_ListOps_add: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_add.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                          &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()))
    }
    pub fn Test_ListOps_Nil() -> &dyn Any {
        static Test_ListOps_Nil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_Nil.get_or_init(||
                                         &LrcPtr::new(PureScript_Test_ListOps::Test_ListOps_List::Test_ListOps_Nilusd_Ctor))
    }
    pub fn Test_ListOps_Cons() -> &dyn Any {
        static Test_ListOps_Cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_Cons.get_or_init(||
                                          &Func1::new(move |usd__arg1|
                                                          Func1::new({
                                                                         let usd__arg1
                                                                             =
                                                                             usd__arg1.clone();
                                                                         move
                                                                             |usd__arg2|
                                                                             &LrcPtr::new(PureScript_Test_ListOps::Test_ListOps_List::Test_ListOps_Consusd_Ctor(usd__arg1,
                                                                                                                                                                usd__arg2.clone()))
                                                                     })))
    }
    pub fn Test_ListOps_range() -> &dyn Any {
        static Test_ListOps_range: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_range.get_or_init(||
                                           &Func1::new(move |start|
                                                           &Func1::new({
                                                                           let start
                                                                               =
                                                                               start.clone();
                                                                           move
                                                                               |end_var|
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
                                                                                                                              |curr|
                                                                                                                              Func1::new({
                                                                                                                                             let curr
                                                                                                                                                 =
                                                                                                                                                 curr.clone();
                                                                                                                                             let go_tco
                                                                                                                                                 =
                                                                                                                                                 go_tco.clone();
                                                                                                                                             move
                                                                                                                                                 |acc|
                                                                                                                                                 go_tco(curr)(acc.clone())
                                                                                                                                         })
                                                                                                                      })
                                                                                                  });
                                                                                   let go_1 =
                                                                                       Lazy(go_2);
                                                                                   fn go_tco(curr_1:
                                                                                                 _)
                                                                                    ->
                                                                                        Func1<&dyn Any,
                                                                                              &dyn Any> {
                                                                                       Func1::new({
                                                                                                      let curr_1
                                                                                                          =
                                                                                                          curr_1.clone();
                                                                                                      let go_tco
                                                                                                          =
                                                                                                          go_tco.clone();
                                                                                                      move
                                                                                                          |acc_1|
                                                                                                          {
                                                                                                              let matchValue =
                                                                                                                  Sharpurs_Prelude::unbox(&&(Sharpurs_Prelude::unbox(&&&curr_1)
                                                                                                                                                 <
                                                                                                                                                 Sharpurs_Prelude::unbox(&&&start)));
                                                                                                              match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                               &matchValue)
                                                                                                                  {
                                                                                                                  0_i32
                                                                                                                  =>
                                                                                                                  acc_1.clone(),
                                                                                                                  _
                                                                                                                  =>
                                                                                                                  go_tco(&(Sharpurs_Prelude::unbox(&&&curr_1)
                                                                                                                               -
                                                                                                                               Sharpurs_Prelude::unbox(&&&1_i32)))(&LrcPtr::new(PureScript_Test_ListOps::Test_ListOps_List::Test_ListOps_Consusd_Ctor(&curr_1,
                                                                                                                                                                                                                                                      acc_1.clone()))),
                                                                                                              }
                                                                                                          }
                                                                                                  })
                                                                                   }
                                                                                   let go =
                                                                                       go_1.Value;
                                                                                   go_tco(end_var.clone())(&LrcPtr::new(PureScript_Test_ListOps::Test_ListOps_List::Test_ListOps_Nilusd_Ctor))
                                                                               }
                                                                       })))
    }
    pub fn Test_ListOps_foldl_004026() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           Func1::new({
                                                          let v1 = v1.clone();
                                                          move |v2|
                                                              PureScript_Test_ListOps::Test_ListOps_foldl_tco(&v,
                                                                                                              &v1,
                                                                                                              v2)
                                                      })
                                   }))
    }
    pub fn Test_ListOps_foldl_004026_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_ListOps_foldl_004026_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_ListOps_foldl_004026_002d1.get_or_init(||
                                                        Lazy(Test_ListOps_foldl_004026.clone()))
    }
    pub fn Test_ListOps_foldl_tco(v: &dyn Any, v1: &dyn Any, v2: &dyn Any)
     -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        let matchValue_2: LrcPtr<PureScript_Test_ListOps::Test_ListOps_List> =
            Sharpurs_Prelude::unbox(v2);
        match matchValue_2.as_ref() {
            PureScript_Test_ListOps::Test_ListOps_List::Test_ListOps_Consusd_Ctor(matchValue_2_1_0,
                                                                                  matchValue_2_1_1)
            => {
                let f = matchValue;
                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Test_ListOps_foldl_004026_002d1.Value,
                                                                                                                       &&&f),
                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                          &&&matchValue_1),
                                                                                                                       &&matchValue_2_1_0)),
                                                 &&matchValue_2_1_1)
            }
            _ => &matchValue_1,
        }
    }
    pub fn Test_ListOps_foldl() -> &dyn Any {
        static Test_ListOps_foldl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_foldl.get_or_init(||
                                           Test_ListOps_foldl_004026_002d1.Value)
    }
    pub fn Test_ListOps_filterEvens() -> &dyn Any {
        static Test_ListOps_filterEvens: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_filterEvens.get_or_init(||
                                                 &Func1::new(move |lst|
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
                                                                                                let matchValue:
                                                                                                        LrcPtr<PureScript_Test_ListOps::Test_ListOps_List> =
                                                                                                    Sharpurs_Prelude::unbox(&&v_1);
                                                                                                let matchValue_1 =
                                                                                                    Sharpurs_Prelude::unbox(v1_1);
                                                                                                match matchValue.as_ref()
                                                                                                    {
                                                                                                    PureScript_Test_ListOps::Test_ListOps_List::Test_ListOps_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                          matchValue_1_1)
                                                                                                    =>
                                                                                                    {
                                                                                                        let xs =
                                                                                                            matchValue_1_1.clone();
                                                                                                        let x =
                                                                                                            matchValue_1_0.clone();
                                                                                                        let acc_1 =
                                                                                                            matchValue_1;
                                                                                                        let matchValue_3 =
                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                            &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                                                                                                               &&&x),
                                                                                                                                                                                                                                            &&&2_i32)),
                                                                                                                                                                      &&&0_i32));
                                                                                                        match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                         &matchValue_3)
                                                                                                            {
                                                                                                            0_i32
                                                                                                            =>
                                                                                                            go_tco(&xs)(&LrcPtr::new(PureScript_Test_ListOps::Test_ListOps_List::Test_ListOps_Consusd_Ctor(&x,
                                                                                                                                                                                                           &acc_1))),
                                                                                                            _
                                                                                                            =>
                                                                                                            go_tco(&xs)(&acc_1),
                                                                                                        }
                                                                                                    }
                                                                                                    _
                                                                                                    =>
                                                                                                    &matchValue_1,
                                                                                                }
                                                                                            }
                                                                                    })
                                                                     }
                                                                     let go =
                                                                         go_1.Value;
                                                                     go_tco(lst.clone())(&LrcPtr::new(PureScript_Test_ListOps::Test_ListOps_List::Test_ListOps_Nilusd_Ctor))
                                                                 }))
    }
    pub fn Test_ListOps_sumEvens() -> &dyn Any {
        static Test_ListOps_sumEvens: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_sumEvens.get_or_init(||
                                              &Func1::new(move |n|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ListOps::Test_ListOps_foldl(),
                                                                                                                                                                     &&&PureScript_Test_ListOps::Test_ListOps_add()),
                                                                                                                                  &&&0_i32),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ListOps::Test_ListOps_filterEvens(),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ListOps::Test_ListOps_range(),
                                                                                                                                                                                                        &&&1_i32),
                                                                                                                                                                     n)))))
    }
    pub fn Test_ListOps_describe() -> &dyn Any {
        static Test_ListOps_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_describe.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                               &&&string("List Processing (900 elements):")))
    }
    pub fn Test_ListOps_act() -> &dyn Any {
        static Test_ListOps_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOps_act.get_or_init(||
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
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ListOps::Test_ListOps_sumEvens(),
                                                                                                                                                                                                   dummy))))))
    }
}
