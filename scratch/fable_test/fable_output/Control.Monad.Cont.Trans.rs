pub mod PureScript_Control_Monad_Cont_Trans {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_c6e4dcf3::PureScript_Control_Monad_Cont_Class;
    use crate::module_6641b520::PureScript_Control_Monad_Reader_Class;
    use crate::module_8aa70462::PureScript_Control_Monad_ST_Class;
    use crate::module_e3c9db92::PureScript_Control_Monad_State_Class;
    use crate::module_f5fe307f::PureScript_Control_Monad_Trans_Class;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Cont_Trans_ContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_ContT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_ContT.get_or_init(||
                                                       &Func1::new(move |x|
                                                                       x.clone()))
    }
    pub fn Control_Monad_Cont_Trans_withContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_withContT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_withContT.get_or_init(||
                                                           &Func1::new(move
                                                                           |f|
                                                                           &Func1::new({
                                                                                           let f
                                                                                               =
                                                                                               f.clone();
                                                                                           move
                                                                                               |v|
                                                                                               {
                                                                                                   let matchValue =
                                                                                                       Sharpurs_Prelude::unbox(&&f);
                                                                                                   let matchValue_1 =
                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                    &&&Func1::new({
                                                                                                                                                      let matchValue_1
                                                                                                                                                          =
                                                                                                                                                          matchValue_1.clone();
                                                                                                                                                      move
                                                                                                                                                          |k|
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                              k))
                                                                                                                                                  }))
                                                                                               }
                                                                                       })))
    }
    pub fn Control_Monad_Cont_Trans_runContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_runContT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_runContT.get_or_init(||
                                                          &Func1::new(move |v|
                                                                          &Func1::new({
                                                                                          let v
                                                                                              =
                                                                                              v.clone();
                                                                                          move
                                                                                              |k|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&v),
                                                                                                                               &&&Sharpurs_Prelude::unbox(k))
                                                                                      })))
    }
    pub fn Control_Monad_Cont_Trans_newtypeContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_newtypeContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_newtypeContT.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                               &&&add(string("Coercible0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &Sharpurs_Prelude::Prim_undefined()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))
    }
    pub fn Control_Monad_Cont_Trans_monadTransContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monadTransContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monadTransContT.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_MonadTransusd_Dict(),
                                                                                                  &&&add(string("lift"),
                                                                                                         &&Func1::new(move
                                                                                                                          |dictMonad|
                                                                                                                          {
                                                                                                                              let Bind1 =
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                              &Func1::new({
                                                                                                                                              let Bind1
                                                                                                                                                  =
                                                                                                                                                  Bind1.clone();
                                                                                                                                              move
                                                                                                                                                  |m|
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                     let m
                                                                                                                                                                                                         =
                                                                                                                                                                                                         m.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |k|
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                &&&Bind1),
                                                                                                                                                                                                                                                                             &&&m),
                                                                                                                                                                                                                                          k)
                                                                                                                                                                                                 }))
                                                                                                                                          })
                                                                                                                          }),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Control_Monad_Cont_Trans_lift() -> &dyn Any {
        static Control_Monad_Cont_Trans_lift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_lift.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                       &&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadTransContT()))
    }
    pub fn Control_Monad_Cont_Trans_mapContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_mapContT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_mapContT.get_or_init(||
                                                          &Func1::new(move |f|
                                                                          &Func1::new({
                                                                                          let f
                                                                                              =
                                                                                              f.clone();
                                                                                          move
                                                                                              |v|
                                                                                              {
                                                                                                  let matchValue =
                                                                                                      Sharpurs_Prelude::unbox(&&f);
                                                                                                  let matchValue_1 =
                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                   &&&Func1::new({
                                                                                                                                                     let matchValue_1
                                                                                                                                                         =
                                                                                                                                                         matchValue_1.clone();
                                                                                                                                                     move
                                                                                                                                                         |k|
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                             k))
                                                                                                                                                 }))
                                                                                              }
                                                                                      })))
    }
    pub fn Control_Monad_Cont_Trans_functorContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_functorContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_functorContT.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictFunctor|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                               &&&add(string("map"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |f|
                                                                                                                                       &Func1::new({
                                                                                                                                                       let f
                                                                                                                                                           =
                                                                                                                                                           f.clone();
                                                                                                                                                       move
                                                                                                                                                           |v|
                                                                                                                                                           {
                                                                                                                                                               let matchValue =
                                                                                                                                                                   Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                               let matchValue_1 =
                                                                                                                                                                   Sharpurs_Prelude::unbox(v);
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                  let matchValue_1
                                                                                                                                                                                                                      =
                                                                                                                                                                                                                      matchValue_1.clone();
                                                                                                                                                                                                                  move
                                                                                                                                                                                                                      |k|
                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                                         let k
                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                             k.clone();
                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                             |a|
                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                 &&&k),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                 a))
                                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                                              }))
                                                                                                                                                           }
                                                                                                                                                   })),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
    }
    pub fn Control_Monad_Cont_Trans_applyContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_applyContT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Cont_Trans_applyContT.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictApply|
                                                                            {
                                                                                let functorContT1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_functorContT(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                               Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                                 &&&add(string("apply"),
                                                                                                                        &&Func1::new(move
                                                                                                                                         |v|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let v
                                                                                                                                                             =
                                                                                                                                                             v.clone();
                                                                                                                                                         move
                                                                                                                                                             |v1|
                                                                                                                                                             {
                                                                                                                                                                 let matchValue =
                                                                                                                                                                     Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                 let matchValue_1 =
                                                                                                                                                                     Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let matchValue_1
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        matchValue_1.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |k|
                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                                                                                                           let k
                                                                                                                                                                                                                                                                               =
                                                                                                                                                                                                                                                                               k.clone();
                                                                                                                                                                                                                                                                           move
                                                                                                                                                                                                                                                                               |g|
                                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                  let g
                                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                                      g.clone();
                                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                                      |a|
                                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&k,
                                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                                                                                                                                                                                          a))
                                                                                                                                                                                                                                                                                                                              }))
                                                                                                                                                                                                                                                                       }))
                                                                                                                                                                                                                }))
                                                                                                                                                             }
                                                                                                                                                     })),
                                                                                                                        add(string("Functor0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let functorContT1
                                                                                                                                                 =
                                                                                                                                                 functorContT1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &functorContT1
                                                                                                                                         }),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>())))
                                                                            }))
    }
    pub fn Control_Monad_Cont_Trans_bindContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_bindContT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_bindContT.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictBind|
                                                                           {
                                                                               let applyContT1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_applyContT(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictBind)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                                &&&add(string("bind"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |v|
                                                                                                                                        &Func1::new({
                                                                                                                                                        let v
                                                                                                                                                            =
                                                                                                                                                            v.clone();
                                                                                                                                                        move
                                                                                                                                                            |k|
                                                                                                                                                            {
                                                                                                                                                                let matchValue =
                                                                                                                                                                    Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                let matchValue_1 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(k);
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                                   let matchValue_1
                                                                                                                                                                                                                       =
                                                                                                                                                                                                                       matchValue_1.clone();
                                                                                                                                                                                                                   move
                                                                                                                                                                                                                       |k_prime|
                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                                                          let k_prime
                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                              k_prime.clone();
                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                              |a|
                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                                             a)),
                                                                                                                                                                                                                                                                                                               &&&k_prime)
                                                                                                                                                                                                                                                                      }))
                                                                                                                                                                                                               }))
                                                                                                                                                            }
                                                                                                                                                    })),
                                                                                                                       add(string("Apply0"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let applyContT1
                                                                                                                                                =
                                                                                                                                                applyContT1.clone();
                                                                                                                                            move
                                                                                                                                                |usd__unused|
                                                                                                                                                &applyContT1
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>())))
                                                                           }))
    }
    pub fn Control_Monad_Cont_Trans_semigroupContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_semigroupContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_semigroupContT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictApply|
                                                                                {
                                                                                    let applyContT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_applyContT(),
                                                                                                                         dictApply);
                                                                                    &Func1::new({
                                                                                                    let applyContT1
                                                                                                        =
                                                                                                        applyContT1.clone();
                                                                                                    move
                                                                                                        |dictSemigroup|
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                                         &&&add(string("append"),
                                                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                     &&&applyContT1),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                     dictSemigroup)),
                                                                                                                                                empty::<string,
                                                                                                                                                        &dyn Any>()))
                                                                                                })
                                                                                }))
    }
    pub fn Control_Monad_Cont_Trans_applicativeContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_applicativeContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_applicativeContT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictApplicative|
                                                                                  {
                                                                                      let applyContT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_applyContT(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                                       &&&add(string("pure"),
                                                                                                                              &&Func1::new(move
                                                                                                                                               |a|
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                  let a
                                                                                                                                                                                                      =
                                                                                                                                                                                                      a.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |k|
                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(k,
                                                                                                                                                                                                                                       &&&a)
                                                                                                                                                                                              }))),
                                                                                                                              add(string("Apply0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let applyContT1
                                                                                                                                                       =
                                                                                                                                                       applyContT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &applyContT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_Cont_Trans_monadContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monadContT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monadContT.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictMonad|
                                                                            {
                                                                                let applicativeContT1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_applicativeContT(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                               Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                let bindContT1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_bindContT(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                               Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                                 &&&add(string("Applicative0"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let applicativeContT1
                                                                                                                                             =
                                                                                                                                             applicativeContT1.clone();
                                                                                                                                         move
                                                                                                                                             |usd__unused|
                                                                                                                                             &applicativeContT1
                                                                                                                                     }),
                                                                                                                        add(string("Bind1"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let bindContT1
                                                                                                                                                 =
                                                                                                                                                 bindContT1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused_1|
                                                                                                                                                 &bindContT1
                                                                                                                                         }),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>())))
                                                                            }))
    }
    pub fn Control_Monad_Cont_Trans_monadAskContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monadAskContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monadAskContT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictMonadAsk|
                                                                               {
                                                                                   let monadContT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadContT(),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadAskusd_Dict(),
                                                                                                                    &&&add(string("ask"),
                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                   &&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadTransContT()),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_ask(),
                                                                                                                                                                                                dictMonadAsk)),
                                                                                                                           add(string("Monad0"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let monadContT1
                                                                                                                                                    =
                                                                                                                                                    monadContT1.clone();
                                                                                                                                                move
                                                                                                                                                    |usd__unused|
                                                                                                                                                    &monadContT1
                                                                                                                                            }),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>())))
                                                                               }))
    }
    pub fn Control_Monad_Cont_Trans_monadReaderContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monadReaderContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monadReaderContT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonadReader|
                                                                                  {
                                                                                      let MonadAsk0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadAsk0"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadReader)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let Bind1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&MonadAsk0)),
                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let ask =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_ask(),
                                                                                                                           &&&MonadAsk0);
                                                                                      let monadAskContT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadAskContT(),
                                                                                                                           &&&MonadAsk0);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadReaderusd_Dict(),
                                                                                                                       &&&add(string("local"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let Bind1
                                                                                                                                                   =
                                                                                                                                                   Bind1.clone();
                                                                                                                                               let ask
                                                                                                                                                   =
                                                                                                                                                   ask.clone();
                                                                                                                                               let dictMonadReader
                                                                                                                                                   =
                                                                                                                                                   dictMonadReader.clone();
                                                                                                                                               move
                                                                                                                                                   |f|
                                                                                                                                                   &Func1::new({
                                                                                                                                                                   let f
                                                                                                                                                                       =
                                                                                                                                                                       f.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |v|
                                                                                                                                                                       {
                                                                                                                                                                           let matchValue =
                                                                                                                                                                               Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                           let matchValue_1 =
                                                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                                              let matchValue_1
                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                  matchValue_1.clone();
                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                  |k|
                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                         &&&Bind1),
                                                                                                                                                                                                                                                                                                      &&&ask),
                                                                                                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                                                                                                     let k
                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                         k.clone();
                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                         |r|
                                                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_local(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&dictMonadReader),
                                                                                                                                                                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_local(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&dictMonadReader),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         r))),
                                                                                                                                                                                                                                                                                                                                                                                                &&&k)))
                                                                                                                                                                                                                                                                                 }))
                                                                                                                                                                                                                          }))
                                                                                                                                                                       }
                                                                                                                                                               })
                                                                                                                                           }),
                                                                                                                              add(string("MonadAsk0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let monadAskContT1
                                                                                                                                                       =
                                                                                                                                                       monadAskContT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &monadAskContT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_Cont_Trans_monadContContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monadContContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monadContContT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonad|
                                                                                {
                                                                                    let monadContT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadContT(),
                                                                                                                         dictMonad);
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Class::Control_Monad_Cont_Class_MonadContusd_Dict(),
                                                                                                                     &&&add(string("callCC"),
                                                                                                                            &&Func1::new(move
                                                                                                                                             |f|
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                let f
                                                                                                                                                                                                    =
                                                                                                                                                                                                    f.clone();
                                                                                                                                                                                                move
                                                                                                                                                                                                    |k|
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                                                                                                                                     let k
                                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                                         k.clone();
                                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                                         |a|
                                                                                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                                                                                                                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                            let a
                                                                                                                                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                                                                                                                                a.clone();
                                                                                                                                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                                                                                                                                |v1|
                                                                                                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&k,
                                                                                                                                                                                                                                                                                                                                                                                                                 &&&a)
                                                                                                                                                                                                                                                                                                                                                                        }))
                                                                                                                                                                                                                                                                                                                 }))),
                                                                                                                                                                                                                                     k)
                                                                                                                                                                                            }))),
                                                                                                                            add(string("Monad0"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let monadContT1
                                                                                                                                                     =
                                                                                                                                                     monadContT1.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused|
                                                                                                                                                     &monadContT1
                                                                                                                                             }),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>())))
                                                                                }))
    }
    pub fn Control_Monad_Cont_Trans_monadEffectContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monadEffectContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monadEffectContT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonadEffect|
                                                                                  {
                                                                                      let Monad0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadEffect)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let monadContT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadContT(),
                                                                                                                           &&&Monad0);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                                                       &&&add(string("liftEffect"),
                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_lift(),
                                                                                                                                                                                                                                      &&&Monad0)),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                   dictMonadEffect)),
                                                                                                                              add(string("Monad0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let monadContT1
                                                                                                                                                       =
                                                                                                                                                       monadContT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &monadContT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_Cont_Trans_monadStateContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monadStateContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monadStateContT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictMonadState|
                                                                                 {
                                                                                     let Monad0 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                 Sharpurs_Prelude::unbox(dictMonadState)),
                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                                     let monadContT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadContT(),
                                                                                                                          &&&Monad0);
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_MonadStateusd_Dict(),
                                                                                                                      &&&add(string("state"),
                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_lift(),
                                                                                                                                                                                                                                     &&&Monad0)),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                                                                  dictMonadState)),
                                                                                                                             add(string("Monad0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let monadContT1
                                                                                                                                                      =
                                                                                                                                                      monadContT1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &monadContT1
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Monad_Cont_Trans_monadSTContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monadSTContT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monadSTContT.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictMonadST|
                                                                              {
                                                                                  let Monad0 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                              Sharpurs_Prelude::unbox(dictMonadST)),
                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                  let monadContT1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadContT(),
                                                                                                                       &&&Monad0);
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                                                   &&&add(string("liftST"),
                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_lift(),
                                                                                                                                                                                                                                  &&&Monad0)),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_liftST(),
                                                                                                                                                                                               dictMonadST)),
                                                                                                                          add(string("Monad0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let monadContT1
                                                                                                                                                   =
                                                                                                                                                   monadContT1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &monadContT1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
    pub fn Control_Monad_Cont_Trans_monoidContT() -> &dyn Any {
        static Control_Monad_Cont_Trans_monoidContT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Cont_Trans_monoidContT.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictApplicative|
                                                                             {
                                                                                 let applicativeContT1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_applicativeContT(),
                                                                                                                      dictApplicative);
                                                                                 let semigroupContT1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_semigroupContT(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 &Func1::new({
                                                                                                 let applicativeContT1
                                                                                                     =
                                                                                                     applicativeContT1.clone();
                                                                                                 let semigroupContT1
                                                                                                     =
                                                                                                     semigroupContT1.clone();
                                                                                                 move
                                                                                                     |dictMonoid|
                                                                                                     {
                                                                                                         let semigroupContT2 =
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&semigroupContT1,
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                                          &&&add(string("mempty"),
                                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                      &&&applicativeContT1),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                      dictMonoid)),
                                                                                                                                                 add(string("Semigroup0"),
                                                                                                                                                     &&Func1::new({
                                                                                                                                                                      let semigroupContT2
                                                                                                                                                                          =
                                                                                                                                                                          semigroupContT2.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |usd__unused|
                                                                                                                                                                          &semigroupContT2
                                                                                                                                                                  }),
                                                                                                                                                     empty::<string,
                                                                                                                                                             &dyn Any>())))
                                                                                                     }
                                                                                             })
                                                                             }))
    }
}
