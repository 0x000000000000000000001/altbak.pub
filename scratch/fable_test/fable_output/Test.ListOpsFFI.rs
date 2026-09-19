pub mod PureScript_Test_ListOpsFFI {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Test_ListOpsFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::Func2;
        use fable_library_rust::Native_::LrcPtr;
        #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
        pub enum List {
            Nil,
            Cons(i32,
                 LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>),
        }
        impl core::fmt::Display for
         PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn range(start: i32, finish: i32)
         -> LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List> {
            let go =
                Func2::new({
                               let start = start.clone();
                               move
                                   |current: i32,
                                    acc:
                                        LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>|
                                   {
                                       let current: MutCell<i32> =
                                           MutCell::new(current);
                                       let acc:
                                               MutCell<LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>> =
                                           MutCell::new(acc.clone());
                                       '_go:
                                           loop  {
                                               break '_go
                                                   (if current.get() < start {
                                                        acc.get()
                                                    } else {
                                                        let current_temp:
                                                                i32 =
                                                            current.get() -
                                                                1_i32;
                                                        let acc_temp:
                                                                LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List> =
                                                            LrcPtr::new(PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List::Cons(current.get(),
                                                                                                                                    acc.get()));
                                                        current.set(current_temp);
                                                        acc.set(acc_temp);
                                                        continue '_go
                                                    }) ;
                                           }
                                   }
                           });
            go(finish,
               LrcPtr::new(PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List::Nil))
        }
        pub fn filterEvens(list:
                               LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>)
         -> LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List> {
            fn go(rest:
                      LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>,
                  acc:
                      LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>)
             ->
                 LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List> {
                let rest:
                        MutCell<LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>> =
                    MutCell::new(rest.clone());
                let acc:
                        MutCell<LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>> =
                    MutCell::new(acc.clone());
                '_go:
                    loop  {
                        break '_go
                            (match rest.get().as_ref() {
                                 PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List::Cons(rest_1_0,
                                                                                             rest_1_1)
                                 => {
                                     let value: i32 = rest_1_0.clone();
                                     {
                                         let rest_temp:
                                                 LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List> =
                                             rest_1_1.clone();
                                         let acc_temp:
                                                 LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List> =
                                             if value % 2_i32 == 0_i32 {
                                                 LrcPtr::new(PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List::Cons(value,
                                                                                                                         acc.get()))
                                             } else { acc.get() };
                                         rest.set(rest_temp);
                                         acc.set(acc_temp);
                                         continue '_go
                                     }
                                 }
                                 _ => acc.get(),
                             }) ;
                    }
            }
            go(list,
               LrcPtr::new(PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List::Nil))
        }
        pub fn foldl<a: Clone +
                     'static>(f: Func2<a, i32, a>, acc: a,
                              _arg:
                                  LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>)
         -> a {
            let f = MutCell::new(f.clone());
            let acc: MutCell<a> = MutCell::new(acc.clone());
            let _arg:
                    MutCell<LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List>> =
                MutCell::new(_arg.clone());
            '_foldl:
                loop  {
                    break '_foldl
                        (match _arg.get().as_ref() {
                             PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List::Cons(_arg_1_0,
                                                                                         _arg_1_1)
                             => {
                                 let f_temp = f.get();
                                 let acc_temp: a =
                                     f(acc.get(), _arg_1_0.clone());
                                 let _arg_temp:
                                         LrcPtr<PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::List> =
                                     _arg_1_1.clone();
                                 f.set(f_temp);
                                 acc.set(acc_temp);
                                 _arg.set(_arg_temp);
                                 continue '_foldl
                             }
                             _ => acc.get(),
                         }) ;
                }
        }
        pub fn runListOpsFFI(n: &dyn Any) -> &dyn Any {
            &PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::foldl(Func2::new(move
                                                                                   |x:
                                                                                        i32,
                                                                                    y:
                                                                                        i32|
                                                                                   x
                                                                                       +
                                                                                       y),
                                                                    0_i32,
                                                                    PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::filterEvens(PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::range(1_i32,
                                                                                                                                                                                        Sharpurs_Prelude::unbox(n))))
        }
    }
    pub fn Test_ListOpsFFI_runListOpsFFI() -> &dyn Any {
        static Test_ListOpsFFI_runListOpsFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOpsFFI_runListOpsFFI.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Test_ListOpsFFI::Test_ListOpsFFI_FFI::runListOpsFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_ListOpsFFI_describe() -> &dyn Any {
        static Test_ListOpsFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOpsFFI_describe.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                  &&&string("List Processing FFI (900 elements):")))
    }
    pub fn Test_ListOpsFFI_act() -> &dyn Any {
        static Test_ListOpsFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ListOpsFFI_act.get_or_init(||
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
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ListOpsFFI::Test_ListOpsFFI_runListOpsFFI(),
                                                                                                                                                                                                      dummy))))))
    }
}
