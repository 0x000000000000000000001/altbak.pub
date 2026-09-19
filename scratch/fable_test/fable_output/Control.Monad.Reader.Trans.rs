pub mod PureScript_Control_Monad_Reader_Trans {
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
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_c6e4dcf3::PureScript_Control_Monad_Cont_Class;
    use crate::module_7a9a81dd::PureScript_Control_Monad_Error_Class;
    use crate::module_6641b520::PureScript_Control_Monad_Reader_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_8aa70462::PureScript_Control_Monad_ST_Class;
    use crate::module_e3c9db92::PureScript_Control_Monad_State_Class;
    use crate::module_f5fe307f::PureScript_Control_Monad_Trans_Class;
    use crate::module_c8a91fca::PureScript_Control_Monad_Writer_Class;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_53a2e11::PureScript_Control_MonadPlus;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_83e4823d::PureScript_Data_Distributive;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Control_Monad_Reader_Trans_ReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_ReaderT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_ReaderT.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Control_Monad_Reader_Trans_withReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_withReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_withReaderT.get_or_init(||
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
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                              &&&matchValue_1),
                                                                                                                                                                           &&&matchValue))
                                                                                                   }
                                                                                           })))
    }
    pub fn Control_Monad_Reader_Trans_runReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_runReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_runReaderT.get_or_init(||
                                                              &Func1::new(move
                                                                              |v|
                                                                              &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Monad_Reader_Trans_newtypeReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_newtypeReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_newtypeReaderT.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                   &&&add(string("Coercible0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &Sharpurs_Prelude::Prim_undefined()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
    }
    pub fn Control_Monad_Reader_Trans_monadTransReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadTransReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadTransReaderT.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_MonadTransusd_Dict(),
                                                                                                      &&&add(string("lift"),
                                                                                                             &&Func1::new(move
                                                                                                                              |dictMonad|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                  &&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT()),
                                                                                                                                                               &&&PureScript_Data_Function::Data_Function_const())),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
    }
    pub fn Control_Monad_Reader_Trans_lift() -> &dyn Any {
        static Control_Monad_Reader_Trans_lift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_lift.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                         &&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadTransReaderT()))
    }
    pub fn Control_Monad_Reader_Trans_mapReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_mapReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_mapReaderT.get_or_init(||
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
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                          &&&matchValue_1))
                                                                                                  }
                                                                                          })))
    }
    pub fn Control_Monad_Reader_Trans_functorReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_functorReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_functorReaderT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictFunctor|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                                   &&&add(string("map"),
                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                               &&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_mapReaderT()),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                               dictFunctor)),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))
    }
    pub fn Control_Monad_Reader_Trans_distributiveReaderT_004024()
     -> &dyn Any {
        &Func1::new(move |dictDistributive|
                        PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_distributiveReaderT_tco(dictDistributive))
    }
    pub fn Control_Monad_Reader_Trans_distributiveReaderT_004024_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Reader_Trans_distributiveReaderT_004024_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_distributiveReaderT_004024_002d1.get_or_init(||
                                                                                    Lazy(Control_Monad_Reader_Trans_distributiveReaderT_004024.clone()))
    }
    pub fn Control_Monad_Reader_Trans_distributiveReaderT_tco(dictDistributive:
                                                                  &dyn Any)
     -> &dyn Any {
        let functorReaderT1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_functorReaderT(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                       Sharpurs_Prelude::unbox(dictDistributive)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_Distributiveusd_Dict(),
                                         &&&add(string("distribute"),
                                                &&Func1::new({
                                                                 let dictDistributive
                                                                     =
                                                                     dictDistributive.clone();
                                                                 move
                                                                     |dictFunctor|
                                                                     &Func1::new({
                                                                                     let dictFunctor
                                                                                         =
                                                                                         dictFunctor.clone();
                                                                                     move
                                                                                         |a|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                          &&&Func1::new({
                                                                                                                                            let a
                                                                                                                                                =
                                                                                                                                                a.clone();
                                                                                                                                            move
                                                                                                                                                |e|
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_collect(),
                                                                                                                                                                                                                                                                                          &&&dictDistributive),
                                                                                                                                                                                                                                                       &&&dictFunctor),
                                                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                                                      let e
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          e.clone();
                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                          |r|
                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(r),
                                                                                                                                                                                                                                                                           &&&e)
                                                                                                                                                                                                                                  })),
                                                                                                                                                                                 &&&a)
                                                                                                                                        }))
                                                                                 })
                                                             }),
                                                add(string("collect"),
                                                    &&Func1::new({
                                                                     let dictDistributive
                                                                         =
                                                                         dictDistributive.clone();
                                                                     move
                                                                         |dictFunctor_1|
                                                                         &Func1::new({
                                                                                         let dictFunctor_1
                                                                                             =
                                                                                             dictFunctor_1.clone();
                                                                                         move
                                                                                             |f|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_distribute(),
                                                                                                                                                                                                                                       &&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_distributiveReaderT_tco(&&dictDistributive)),
                                                                                                                                                                                                    &&&dictFunctor_1)),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                    &&&dictFunctor_1),
                                                                                                                                                                 f))
                                                                                     })
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new({
                                                                         let functorReaderT1
                                                                             =
                                                                             functorReaderT1.clone();
                                                                         move
                                                                             |usd__unused|
                                                                             &functorReaderT1
                                                                     }),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Control_Monad_Reader_Trans_distributiveReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_distributiveReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_distributiveReaderT.get_or_init(||
                                                                       Control_Monad_Reader_Trans_distributiveReaderT_004024_002d1.Value)
    }
    pub fn Control_Monad_Reader_Trans_applyReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_applyReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_applyReaderT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictApply|
                                                                                {
                                                                                    let functorReaderT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_functorReaderT(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                                     &&&add(string("apply"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let dictApply
                                                                                                                                                 =
                                                                                                                                                 dictApply.clone();
                                                                                                                                             move
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
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                                                            let matchValue_1
                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                matchValue_1.clone();
                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                |r|
                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                       &&&dictApply),
                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                       r)),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                    r))
                                                                                                                                                                                                                        }))
                                                                                                                                                                     }
                                                                                                                                                             })
                                                                                                                                         }),
                                                                                                                            add(string("Functor0"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let functorReaderT1
                                                                                                                                                     =
                                                                                                                                                     functorReaderT1.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused|
                                                                                                                                                     &functorReaderT1
                                                                                                                                             }),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>())))
                                                                                }))
    }
    pub fn Control_Monad_Reader_Trans_bindReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_bindReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_bindReaderT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictBind|
                                                                               {
                                                                                   let applyReaderT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_applyReaderT(),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(dictBind)),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                                    &&&add(string("bind"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let dictBind
                                                                                                                                                =
                                                                                                                                                dictBind.clone();
                                                                                                                                            move
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
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                                                           let matchValue_1
                                                                                                                                                                                                                               =
                                                                                                                                                                                                                               matchValue_1.clone();
                                                                                                                                                                                                                           move
                                                                                                                                                                                                                               |r|
                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                      &&&dictBind),
                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                      r)),
                                                                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                                                                  let r
                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                      r.clone();
                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                      |a|
                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                                                     a)),
                                                                                                                                                                                                                                                                                                                       &&&r)
                                                                                                                                                                                                                                                                              }))
                                                                                                                                                                                                                       }))
                                                                                                                                                                    }
                                                                                                                                                            })
                                                                                                                                        }),
                                                                                                                           add(string("Apply0"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let applyReaderT1
                                                                                                                                                    =
                                                                                                                                                    applyReaderT1.clone();
                                                                                                                                                move
                                                                                                                                                    |usd__unused|
                                                                                                                                                    &applyReaderT1
                                                                                                                                            }),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>())))
                                                                               }))
    }
    pub fn Control_Monad_Reader_Trans_semigroupReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_semigroupReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_semigroupReaderT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictApply|
                                                                                    {
                                                                                        let applyReaderT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_applyReaderT(),
                                                                                                                             dictApply);
                                                                                        &Func1::new({
                                                                                                        let applyReaderT1
                                                                                                            =
                                                                                                            applyReaderT1.clone();
                                                                                                        move
                                                                                                            |dictSemigroup|
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                                             &&&add(string("append"),
                                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                         &&&applyReaderT1),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                         dictSemigroup)),
                                                                                                                                                    empty::<string,
                                                                                                                                                            &dyn Any>()))
                                                                                                    })
                                                                                    }))
    }
    pub fn Control_Monad_Reader_Trans_applicativeReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_applicativeReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_applicativeReaderT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictApplicative|
                                                                                      {
                                                                                          let applyReaderT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_applyReaderT(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                                           &&&add(string("pure"),
                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                       &&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT()),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                          &&&PureScript_Data_Function::Data_Function_const()),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                          dictApplicative))),
                                                                                                                                  add(string("Apply0"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let applyReaderT1
                                                                                                                                                           =
                                                                                                                                                           applyReaderT1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &applyReaderT1
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))
                                                                                      }))
    }
    pub fn Control_Monad_Reader_Trans_monadReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadReaderT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonad|
                                                                                {
                                                                                    let applicativeReaderT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_applicativeReaderT(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    let bindReaderT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_bindReaderT(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                                     &&&add(string("Applicative0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let applicativeReaderT1
                                                                                                                                                 =
                                                                                                                                                 applicativeReaderT1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &applicativeReaderT1
                                                                                                                                         }),
                                                                                                                            add(string("Bind1"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let bindReaderT1
                                                                                                                                                     =
                                                                                                                                                     bindReaderT1.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused_1|
                                                                                                                                                     &bindReaderT1
                                                                                                                                             }),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>())))
                                                                                }))
    }
    pub fn Control_Monad_Reader_Trans_monadAskReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadAskReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadAskReaderT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonad|
                                                                                   {
                                                                                       let monadReaderT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                            dictMonad);
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadAskusd_Dict(),
                                                                                                                        &&&add(string("ask"),
                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadReaderT1
                                                                                                                                                        =
                                                                                                                                                        monadReaderT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadReaderT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Reader_Trans_monadReaderReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadReaderReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadReaderReaderT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonad|
                                                                                      {
                                                                                          let monadAskReaderT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadAskReaderT(),
                                                                                                                               dictMonad);
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadReaderusd_Dict(),
                                                                                                                           &&&add(string("local"),
                                                                                                                                  &&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_withReaderT(),
                                                                                                                                  add(string("MonadAsk0"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let monadAskReaderT1
                                                                                                                                                           =
                                                                                                                                                           monadAskReaderT1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &monadAskReaderT1
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))
                                                                                      }))
    }
    pub fn Control_Monad_Reader_Trans_monadContReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadContReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadContReaderT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonadCont|
                                                                                    {
                                                                                        let monadReaderT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonadCont)),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Class::Control_Monad_Cont_Class_MonadContusd_Dict(),
                                                                                                                         &&&add(string("callCC"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let dictMonadCont
                                                                                                                                                     =
                                                                                                                                                     dictMonadCont.clone();
                                                                                                                                                 move
                                                                                                                                                     |f|
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                        let f
                                                                                                                                                                                                            =
                                                                                                                                                                                                            f.clone();
                                                                                                                                                                                                        move
                                                                                                                                                                                                            |r|
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Class::Control_Monad_Cont_Class_callCC(),
                                                                                                                                                                                                                                                                                &&&dictMonadCont),
                                                                                                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                                                                                                               let r
                                                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                                                   r.clone();
                                                                                                                                                                                                                                                               move
                                                                                                                                                                                                                                                                   |c|
                                                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT()),
                                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Function::Data_Function_const()),
                                                                                                                                                                                                                                                                                                                                                                                                                                        c)))),
                                                                                                                                                                                                                                                                                                    &&&r)
                                                                                                                                                                                                                                                           }))
                                                                                                                                                                                                    }))
                                                                                                                                             }),
                                                                                                                                add(string("Monad0"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let monadReaderT1
                                                                                                                                                         =
                                                                                                                                                         monadReaderT1.clone();
                                                                                                                                                     move
                                                                                                                                                         |usd__unused|
                                                                                                                                                         &monadReaderT1
                                                                                                                                                 }),
                                                                                                                                    empty::<string,
                                                                                                                                            &dyn Any>())))
                                                                                    }))
    }
    pub fn Control_Monad_Reader_Trans_monadEffectReader() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadEffectReader:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadEffectReader.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonadEffect|
                                                                                     {
                                                                                         let Monad0 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                     Sharpurs_Prelude::unbox(dictMonadEffect)),
                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                                         let monadReaderT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                              &&&Monad0);
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                                                          &&&add(string("liftEffect"),
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_lift(),
                                                                                                                                                                                                                                         &&&Monad0)),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                      dictMonadEffect)),
                                                                                                                                 add(string("Monad0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let monadReaderT1
                                                                                                                                                          =
                                                                                                                                                          monadReaderT1.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &monadReaderT1
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }))
    }
    pub fn Control_Monad_Reader_Trans_monadRecReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadRecReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadRecReaderT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadRec|
                                                                                   {
                                                                                       let Monad0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let Bind1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                   Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let pure_var =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       let monadReaderT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                            &&&Monad0);
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                                                        &&&add(string("tailRecM"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let Bind1
                                                                                                                                                    =
                                                                                                                                                    Bind1.clone();
                                                                                                                                                let dictMonadRec
                                                                                                                                                    =
                                                                                                                                                    dictMonadRec.clone();
                                                                                                                                                let pure_var
                                                                                                                                                    =
                                                                                                                                                    pure_var.clone();
                                                                                                                                                move
                                                                                                                                                    |k|
                                                                                                                                                    &Func1::new({
                                                                                                                                                                    let k
                                                                                                                                                                        =
                                                                                                                                                                        k.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |a|
                                                                                                                                                                        {
                                                                                                                                                                            let k_prime =
                                                                                                                                                                                &Func1::new(move
                                                                                                                                                                                                |r|
                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                let r
                                                                                                                                                                                                                    =
                                                                                                                                                                                                                    r.clone();
                                                                                                                                                                                                                move
                                                                                                                                                                                                                    |a_prime|
                                                                                                                                                                                                                    {
                                                                                                                                                                                                                        let f =
                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&k,
                                                                                                                                                                                                                                                                                       a_prime));
                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bindFlipped(),
                                                                                                                                                                                                                                                                                                                               &&&Bind1),
                                                                                                                                                                                                                                                                                            &&&pure_var),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                            &&&r))
                                                                                                                                                                                                                    }
                                                                                                                                                                                                            }));
                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                                                                               let a
                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                   a.clone();
                                                                                                                                                                                                                               let k_prime
                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                   k_prime.clone();
                                                                                                                                                                                                                               move
                                                                                                                                                                                                                                   |r_1|
                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                                                                                                                                          &&&dictMonadRec),
                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&k_prime,
                                                                                                                                                                                                                                                                                                                                          r_1)),
                                                                                                                                                                                                                                                                    &&&a)
                                                                                                                                                                                                                           }))
                                                                                                                                                                        }
                                                                                                                                                                })
                                                                                                                                            }),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadReaderT1
                                                                                                                                                        =
                                                                                                                                                        monadReaderT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadReaderT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Reader_Trans_monadStateReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadStateReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadStateReaderT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonadState|
                                                                                     {
                                                                                         let Monad0 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                     Sharpurs_Prelude::unbox(dictMonadState)),
                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                                         let monadReaderT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                              &&&Monad0);
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_MonadStateusd_Dict(),
                                                                                                                          &&&add(string("state"),
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_lift(),
                                                                                                                                                                                                                                         &&&Monad0)),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                                                                      dictMonadState)),
                                                                                                                                 add(string("Monad0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let monadReaderT1
                                                                                                                                                          =
                                                                                                                                                          monadReaderT1.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &monadReaderT1
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }))
    }
    pub fn Control_Monad_Reader_Trans_monadTellReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadTellReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadTellReaderT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonadTell|
                                                                                    {
                                                                                        let Monad1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad1"),
                                                                                                                                    Sharpurs_Prelude::unbox(dictMonadTell)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let Semigroup0 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                    Sharpurs_Prelude::unbox(dictMonadTell)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let monadReaderT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                             &&&Monad1);
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadTellusd_Dict(),
                                                                                                                         &&&add(string("tell"),
                                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_lift(),
                                                                                                                                                                                                                                        &&&Monad1)),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_tell(),
                                                                                                                                                                                                     dictMonadTell)),
                                                                                                                                add(string("Semigroup0"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let Semigroup0
                                                                                                                                                         =
                                                                                                                                                         Semigroup0.clone();
                                                                                                                                                     move
                                                                                                                                                         |usd__unused|
                                                                                                                                                         &Semigroup0
                                                                                                                                                 }),
                                                                                                                                    add(string("Monad1"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let monadReaderT1
                                                                                                                                                             =
                                                                                                                                                             monadReaderT1.clone();
                                                                                                                                                         move
                                                                                                                                                             |usd__unused_1|
                                                                                                                                                             &monadReaderT1
                                                                                                                                                     }),
                                                                                                                                        empty::<string,
                                                                                                                                                &dyn Any>()))))
                                                                                    }))
    }
    pub fn Control_Monad_Reader_Trans_monadWriterReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadWriterReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadWriterReaderT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonadWriter|
                                                                                      {
                                                                                          let Monoid0 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Monoid0"),
                                                                                                                                      Sharpurs_Prelude::unbox(dictMonadWriter)),
                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                                          let monadTellReaderT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadTellReaderT(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadTell1"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictMonadWriter)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadWriterusd_Dict(),
                                                                                                                           &&&add(string("listen"),
                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_mapReaderT(),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_listen(),
                                                                                                                                                                                                       dictMonadWriter)),
                                                                                                                                  add(string("pass"),
                                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_mapReaderT(),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_pass(),
                                                                                                                                                                                                           dictMonadWriter)),
                                                                                                                                      add(string("Monoid0"),
                                                                                                                                          &&Func1::new({
                                                                                                                                                           let Monoid0
                                                                                                                                                               =
                                                                                                                                                               Monoid0.clone();
                                                                                                                                                           move
                                                                                                                                                               |usd__unused|
                                                                                                                                                               &Monoid0
                                                                                                                                                       }),
                                                                                                                                          add(string("MonadTell1"),
                                                                                                                                              &&Func1::new({
                                                                                                                                                               let monadTellReaderT1
                                                                                                                                                                   =
                                                                                                                                                                   monadTellReaderT1.clone();
                                                                                                                                                               move
                                                                                                                                                                   |usd__unused_1|
                                                                                                                                                                   &monadTellReaderT1
                                                                                                                                                           }),
                                                                                                                                              empty::<string,
                                                                                                                                                      &dyn Any>())))))
                                                                                      }))
    }
    pub fn Control_Monad_Reader_Trans_monadThrowReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadThrowReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadThrowReaderT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonadThrow|
                                                                                     {
                                                                                         let Monad0 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                     Sharpurs_Prelude::unbox(dictMonadThrow)),
                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                                         let monadReaderT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                              &&&Monad0);
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadThrowusd_Dict(),
                                                                                                                          &&&add(string("throwError"),
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_lift(),
                                                                                                                                                                                                                                         &&&Monad0)),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                                                                      dictMonadThrow)),
                                                                                                                                 add(string("Monad0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let monadReaderT1
                                                                                                                                                          =
                                                                                                                                                          monadReaderT1.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &monadReaderT1
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }))
    }
    pub fn Control_Monad_Reader_Trans_monadErrorReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadErrorReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadErrorReaderT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonadError|
                                                                                     {
                                                                                         let monadThrowReaderT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadThrowReaderT(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadThrow0"),
                                                                                                                                                                        Sharpurs_Prelude::unbox(dictMonadError)),
                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadErrorusd_Dict(),
                                                                                                                          &&&add(string("catchError"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let dictMonadError
                                                                                                                                                      =
                                                                                                                                                      dictMonadError.clone();
                                                                                                                                                  move
                                                                                                                                                      |v|
                                                                                                                                                      &Func1::new({
                                                                                                                                                                      let v
                                                                                                                                                                          =
                                                                                                                                                                          v.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |h|
                                                                                                                                                                          {
                                                                                                                                                                              let matchValue =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                              let matchValue_1 =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(h);
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                 let matchValue_1
                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                     matchValue_1.clone();
                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                     |r|
                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_catchError(),
                                                                                                                                                                                                                                                                                                                                            &&&dictMonadError),
                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                            r)),
                                                                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                                                                        let r
                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                            r.clone();
                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                            |e|
                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                                                           e)),
                                                                                                                                                                                                                                                                                                                             &&&r)
                                                                                                                                                                                                                                                                                    }))
                                                                                                                                                                                                                             }))
                                                                                                                                                                          }
                                                                                                                                                                  })
                                                                                                                                              }),
                                                                                                                                 add(string("MonadThrow0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let monadThrowReaderT1
                                                                                                                                                          =
                                                                                                                                                          monadThrowReaderT1.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &monadThrowReaderT1
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }))
    }
    pub fn Control_Monad_Reader_Trans_monadSTReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadSTReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadSTReaderT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonadST|
                                                                                  {
                                                                                      let Monad0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadST)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let monadReaderT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                           &&&Monad0);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                                                       &&&add(string("liftST"),
                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_lift(),
                                                                                                                                                                                                                                      &&&Monad0)),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_liftST(),
                                                                                                                                                                                                   dictMonadST)),
                                                                                                                              add(string("Monad0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let monadReaderT1
                                                                                                                                                       =
                                                                                                                                                       monadReaderT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &monadReaderT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_Reader_Trans_monoidReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monoidReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monoidReaderT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictApplicative|
                                                                                 {
                                                                                     let applicativeReaderT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_applicativeReaderT(),
                                                                                                                          dictApplicative);
                                                                                     let semigroupReaderT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_semigroupReaderT(),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                    Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                     &Func1::new({
                                                                                                     let applicativeReaderT1
                                                                                                         =
                                                                                                         applicativeReaderT1.clone();
                                                                                                     let semigroupReaderT1
                                                                                                         =
                                                                                                         semigroupReaderT1.clone();
                                                                                                     move
                                                                                                         |dictMonoid|
                                                                                                         {
                                                                                                             let semigroupReaderT2 =
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&semigroupReaderT1,
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                                              &&&add(string("mempty"),
                                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                          &&&applicativeReaderT1),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                          dictMonoid)),
                                                                                                                                                     add(string("Semigroup0"),
                                                                                                                                                         &&Func1::new({
                                                                                                                                                                          let semigroupReaderT2
                                                                                                                                                                              =
                                                                                                                                                                              semigroupReaderT2.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |usd__unused|
                                                                                                                                                                              &semigroupReaderT2
                                                                                                                                                                      }),
                                                                                                                                                         empty::<string,
                                                                                                                                                                 &dyn Any>())))
                                                                                                         }
                                                                                                 })
                                                                                 }))
    }
    pub fn Control_Monad_Reader_Trans_altReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_altReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_altReaderT.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictAlt|
                                                                              {
                                                                                  let functorReaderT1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_functorReaderT(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictAlt)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                                                   &&&add(string("alt"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictAlt
                                                                                                                                               =
                                                                                                                                               dictAlt.clone();
                                                                                                                                           move
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
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                          let matchValue_1
                                                                                                                                                                                                                              =
                                                                                                                                                                                                                              matchValue_1.clone();
                                                                                                                                                                                                                          move
                                                                                                                                                                                                                              |r|
                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                                                                                                                                                     &&&dictAlt),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                     r)),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                  r))
                                                                                                                                                                                                                      }))
                                                                                                                                                                   }
                                                                                                                                                           })
                                                                                                                                       }),
                                                                                                                          add(string("Functor0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let functorReaderT1
                                                                                                                                                   =
                                                                                                                                                   functorReaderT1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &functorReaderT1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
    pub fn Control_Monad_Reader_Trans_plusReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_plusReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_plusReaderT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictPlus|
                                                                               {
                                                                                   let altReaderT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_altReaderT(),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(dictPlus)),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                                                    &&&add(string("empty"),
                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_ReaderT(),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                                                                                                                                   dictPlus))),
                                                                                                                           add(string("Alt0"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let altReaderT1
                                                                                                                                                    =
                                                                                                                                                    altReaderT1.clone();
                                                                                                                                                move
                                                                                                                                                    |usd__unused|
                                                                                                                                                    &altReaderT1
                                                                                                                                            }),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>())))
                                                                               }))
    }
    pub fn Control_Monad_Reader_Trans_alternativeReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_alternativeReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_alternativeReaderT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictAlternative|
                                                                                      {
                                                                                          let applicativeReaderT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_applicativeReaderT(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          let plusReaderT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_plusReaderT(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                                                           &&&add(string("Applicative0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let applicativeReaderT1
                                                                                                                                                       =
                                                                                                                                                       applicativeReaderT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &applicativeReaderT1
                                                                                                                                               }),
                                                                                                                                  add(string("Plus1"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let plusReaderT1
                                                                                                                                                           =
                                                                                                                                                           plusReaderT1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused_1|
                                                                                                                                                           &plusReaderT1
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))
                                                                                      }))
    }
    pub fn Control_Monad_Reader_Trans_monadPlusReaderT() -> &dyn Any {
        static Control_Monad_Reader_Trans_monadPlusReaderT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Trans_monadPlusReaderT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonadPlus|
                                                                                    {
                                                                                        let monadReaderT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadReaderT(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonadPlus)),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        let alternativeReaderT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_alternativeReaderT(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alternative1"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonadPlus)),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                                                         &&&add(string("Monad0"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let monadReaderT1
                                                                                                                                                     =
                                                                                                                                                     monadReaderT1.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused|
                                                                                                                                                     &monadReaderT1
                                                                                                                                             }),
                                                                                                                                add(string("Alternative1"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let alternativeReaderT1
                                                                                                                                                         =
                                                                                                                                                         alternativeReaderT1.clone();
                                                                                                                                                     move
                                                                                                                                                         |usd__unused_1|
                                                                                                                                                         &alternativeReaderT1
                                                                                                                                                 }),
                                                                                                                                    empty::<string,
                                                                                                                                            &dyn Any>())))
                                                                                    }))
    }
}
