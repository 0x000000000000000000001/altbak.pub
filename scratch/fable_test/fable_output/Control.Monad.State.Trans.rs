pub mod PureScript_Control_Monad_State_Trans {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty as empty_1;
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
    use crate::module_209e0d2c::PureScript_Control_Lazy;
    use crate::module_c6e4dcf3::PureScript_Control_Monad_Cont_Class;
    use crate::module_7a9a81dd::PureScript_Control_Monad_Error_Class;
    use crate::module_6641b520::PureScript_Control_Monad_Reader_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_8aa70462::PureScript_Control_Monad_ST_Class;
    use crate::module_e3c9db92::PureScript_Control_Monad_State_Class;
    use crate::module_f5fe307f::PureScript_Control_Monad_Trans_Class;
    use crate::module_c8a91fca::PureScript_Control_Monad_Writer_Class;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_53a2e11::PureScript_Control_MonadPlus;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Control_Monad_State_Trans_StateT() -> &dyn Any {
        static Control_Monad_State_Trans_StateT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_StateT.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Control_Monad_State_Trans_withStateT() -> &dyn Any {
        static Control_Monad_State_Trans_withStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_withStateT.get_or_init(||
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
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                            &&&matchValue_1),
                                                                                                                                                                         &&&matchValue))
                                                                                                 }
                                                                                         })))
    }
    pub fn Control_Monad_State_Trans_runStateT() -> &dyn Any {
        static Control_Monad_State_Trans_runStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_runStateT.get_or_init(||
                                                            &Func1::new(move
                                                                            |v|
                                                                            &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Monad_State_Trans_newtypeStateT() -> &dyn Any {
        static Control_Monad_State_Trans_newtypeStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_newtypeStateT.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                 &&&add(string("Coercible0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &Sharpurs_Prelude::Prim_undefined()),
                                                                                                        empty_1::<string,
                                                                                                                  &dyn Any>())))
    }
    pub fn Control_Monad_State_Trans_monadTransStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadTransStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadTransStateT.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_MonadTransusd_Dict(),
                                                                                                    &&&add(string("lift"),
                                                                                                           &&Func1::new(move
                                                                                                                            |dictMonad|
                                                                                                                            {
                                                                                                                                let Bind1 =
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                let pure_var =
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                &Func1::new({
                                                                                                                                                let Bind1
                                                                                                                                                    =
                                                                                                                                                    Bind1.clone();
                                                                                                                                                let pure_var
                                                                                                                                                    =
                                                                                                                                                    pure_var.clone();
                                                                                                                                                move
                                                                                                                                                    |m|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                                                       let m
                                                                                                                                                                                                           =
                                                                                                                                                                                                           m.clone();
                                                                                                                                                                                                       move
                                                                                                                                                                                                           |s|
                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                  &&&Bind1),
                                                                                                                                                                                                                                                                               &&&m),
                                                                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                                                                              let s
                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                  s.clone();
                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                  |x|
                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                      &&&pure_var),
                                                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x.clone(),
                                                                                                                                                                                                                                                                                                                                                             &s)))
                                                                                                                                                                                                                                                          }))
                                                                                                                                                                                                   }))
                                                                                                                                            })
                                                                                                                            }),
                                                                                                           empty_1::<string,
                                                                                                                     &dyn Any>())))
    }
    pub fn Control_Monad_State_Trans_lift() -> &dyn Any {
        static Control_Monad_State_Trans_lift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_lift.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                        &&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadTransStateT()))
    }
    pub fn Control_Monad_State_Trans_mapStateT() -> &dyn Any {
        static Control_Monad_State_Trans_mapStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_mapStateT.get_or_init(||
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
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                        &&&matchValue_1))
                                                                                                }
                                                                                        })))
    }
    pub fn Control_Monad_State_Trans_lazyStateT() -> &dyn Any {
        static Control_Monad_State_Trans_lazyStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_lazyStateT.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_Lazyusd_Dict(),
                                                                                              &&&add(string("defer"),
                                                                                                     &&Func1::new(move
                                                                                                                      |f|
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let f
                                                                                                                                                                             =
                                                                                                                                                                             f.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |s|
                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                            &&&PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                                                                                                              s)
                                                                                                                                                                     }))),
                                                                                                     empty_1::<string,
                                                                                                               &dyn Any>())))
    }
    pub fn Control_Monad_State_Trans_functorStateT() -> &dyn Any {
        static Control_Monad_State_Trans_functorStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_functorStateT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictFunctor|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                                 &&&add(string("map"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let dictFunctor
                                                                                                                                             =
                                                                                                                                             dictFunctor.clone();
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
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                        let matchValue_1
                                                                                                                                                                                                                            =
                                                                                                                                                                                                                            matchValue_1.clone();
                                                                                                                                                                                                                        move
                                                                                                                                                                                                                            |s|
                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                   &&&dictFunctor),
                                                                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                  |v1|
                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                      let matchValue_3:
                                                                                                                                                                                                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                               &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                                                                                                                                                              &match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                                                               }))
                                                                                                                                                                                                                                                                                                                  })),
                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                s))
                                                                                                                                                                                                                    }))
                                                                                                                                                                 }
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty_1::<string,
                                                                                                                                  &dyn Any>()))))
    }
    pub fn Control_Monad_State_Trans_execStateT() -> &dyn Any {
        static Control_Monad_State_Trans_execStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_execStateT.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictFunctor|
                                                                             &Func1::new({
                                                                                             let dictFunctor
                                                                                                 =
                                                                                                 dictFunctor.clone();
                                                                                             move
                                                                                                 |v|
                                                                                                 &Func1::new({
                                                                                                                 let v
                                                                                                                     =
                                                                                                                     v.clone();
                                                                                                                 move
                                                                                                                     |s|
                                                                                                                     {
                                                                                                                         let matchValue =
                                                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                                                         let matchValue_1 =
                                                                                                                             Sharpurs_Prelude::unbox(s);
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                &&&dictFunctor),
                                                                                                                                                                                             &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                             &&&matchValue_1))
                                                                                                                     }
                                                                                                             })
                                                                                         })))
    }
    pub fn Control_Monad_State_Trans_evalStateT() -> &dyn Any {
        static Control_Monad_State_Trans_evalStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_evalStateT.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictFunctor|
                                                                             &Func1::new({
                                                                                             let dictFunctor
                                                                                                 =
                                                                                                 dictFunctor.clone();
                                                                                             move
                                                                                                 |v|
                                                                                                 &Func1::new({
                                                                                                                 let v
                                                                                                                     =
                                                                                                                     v.clone();
                                                                                                                 move
                                                                                                                     |s|
                                                                                                                     {
                                                                                                                         let matchValue =
                                                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                                                         let matchValue_1 =
                                                                                                                             Sharpurs_Prelude::unbox(s);
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                &&&dictFunctor),
                                                                                                                                                                                             &&&PureScript_Data_Tuple::Data_Tuple_fst()),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                             &&&matchValue_1))
                                                                                                                     }
                                                                                                             })
                                                                                         })))
    }
    pub fn Control_Monad_State_Trans_monadStateT_004030() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT_tco(dictMonad))
    }
    pub fn Control_Monad_State_Trans_monadStateT_004030_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_State_Trans_monadStateT_004030_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadStateT_004030_002d1.get_or_init(||
                                                                           Lazy(Control_Monad_State_Trans_monadStateT_004030.clone()))
    }
    pub fn Control_Monad_State_Trans_bindStateT_004032() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_bindStateT_tco(dictMonad))
    }
    pub fn Control_Monad_State_Trans_bindStateT_004032_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_State_Trans_bindStateT_004032_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_State_Trans_bindStateT_004032_002d1.get_or_init(||
                                                                          Lazy(Control_Monad_State_Trans_bindStateT_004032.clone()))
    }
    pub fn Control_Monad_State_Trans_applyStateT_004034() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_applyStateT_tco(dictMonad))
    }
    pub fn Control_Monad_State_Trans_applyStateT_004034_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_State_Trans_applyStateT_004034_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_State_Trans_applyStateT_004034_002d1.get_or_init(||
                                                                           Lazy(Control_Monad_State_Trans_applyStateT_004034.clone()))
    }
    pub fn Control_Monad_State_Trans_applicativeStateT_004036() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_applicativeStateT_tco(dictMonad))
    }
    pub fn Control_Monad_State_Trans_applicativeStateT_004036_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_State_Trans_applicativeStateT_004036_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_State_Trans_applicativeStateT_004036_002d1.get_or_init(||
                                                                                 Lazy(Control_Monad_State_Trans_applicativeStateT_004036.clone()))
    }
    pub fn Control_Monad_State_Trans_monadStateT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                         &&&add(string("Applicative0"),
                                                &&Func1::new({
                                                                 let dictMonad
                                                                     =
                                                                     dictMonad.clone();
                                                                 move
                                                                     |usd__unused|
                                                                     PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_applicativeStateT_tco(&&dictMonad)
                                                             }),
                                                add(string("Bind1"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused_1|
                                                                         PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_bindStateT_tco(&&dictMonad)
                                                                 }),
                                                    empty_1::<string,
                                                              &dyn Any>())))
    }
    pub fn Control_Monad_State_Trans_monadStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadStateT.get_or_init(||
                                                              Control_Monad_State_Trans_monadStateT_004030_002d1.Value)
    }
    pub fn Control_Monad_State_Trans_bindStateT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                         &&&add(string("bind"),
                                                &&Func1::new({
                                                                 let Bind1 =
                                                                     Bind1.clone();
                                                                 move |v|
                                                                     &Func1::new({
                                                                                     let v
                                                                                         =
                                                                                         v.clone();
                                                                                     move
                                                                                         |f|
                                                                                         {
                                                                                             let matchValue =
                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                             let matchValue_1 =
                                                                                                 Sharpurs_Prelude::unbox(f);
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                              &&&Func1::new({
                                                                                                                                                let matchValue_1
                                                                                                                                                    =
                                                                                                                                                    matchValue_1.clone();
                                                                                                                                                move
                                                                                                                                                    |s|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                           &&&Bind1),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                           s)),
                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                       |v1|
                                                                                                                                                                                                       {
                                                                                                                                                                                                           let matchValue_3:
                                                                                                                                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                          &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                             })),
                                                                                                                                                                                                                                            &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                               })
                                                                                                                                                                                                       }))
                                                                                                                                            }))
                                                                                         }
                                                                                 })
                                                             }),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_applyStateT_tco(&&dictMonad)
                                                                 }),
                                                    empty_1::<string,
                                                              &dyn Any>())))
    }
    pub fn Control_Monad_State_Trans_bindStateT() -> &dyn Any {
        static Control_Monad_State_Trans_bindStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_bindStateT.get_or_init(||
                                                             Control_Monad_State_Trans_bindStateT_004032_002d1.Value)
    }
    pub fn Control_Monad_State_Trans_applyStateT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let functorStateT1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_functorStateT(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                         &&&add(string("apply"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_ap(),
                                                                                  &&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT_tco(dictMonad)),
                                                add(string("Functor0"),
                                                    &&Func1::new({
                                                                     let functorStateT1
                                                                         =
                                                                         functorStateT1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &functorStateT1
                                                                 }),
                                                    empty_1::<string,
                                                              &dyn Any>())))
    }
    pub fn Control_Monad_State_Trans_applyStateT() -> &dyn Any {
        static Control_Monad_State_Trans_applyStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_applyStateT.get_or_init(||
                                                              Control_Monad_State_Trans_applyStateT_004034_002d1.Value)
    }
    pub fn Control_Monad_State_Trans_applicativeStateT_tco(dictMonad:
                                                               &dyn Any)
     -> &dyn Any {
        let pure_var =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                       Sharpurs_Prelude::unbox(dictMonad)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                         &&&add(string("pure"),
                                                &&Func1::new({
                                                                 let pure_var
                                                                     =
                                                                     pure_var.clone();
                                                                 move |a|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                      &&&Func1::new({
                                                                                                                        let a
                                                                                                                            =
                                                                                                                            a.clone();
                                                                                                                        move
                                                                                                                            |s|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                &&&pure_var),
                                                                                                                                                             &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                       s.clone())))
                                                                                                                    }))
                                                             }),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_applyStateT_tco(&&dictMonad)
                                                                 }),
                                                    empty_1::<string,
                                                              &dyn Any>())))
    }
    pub fn Control_Monad_State_Trans_applicativeStateT() -> &dyn Any {
        static Control_Monad_State_Trans_applicativeStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_applicativeStateT.get_or_init(||
                                                                    Control_Monad_State_Trans_applicativeStateT_004036_002d1.Value)
    }
    pub fn Control_Monad_State_Trans_semigroupStateT() -> &dyn Any {
        static Control_Monad_State_Trans_semigroupStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_semigroupStateT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonad|
                                                                                  {
                                                                                      let applyStateT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_applyStateT(),
                                                                                                                           dictMonad);
                                                                                      &Func1::new({
                                                                                                      let applyStateT1
                                                                                                          =
                                                                                                          applyStateT1.clone();
                                                                                                      move
                                                                                                          |dictSemigroup|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                                           &&&add(string("append"),
                                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                       &&&applyStateT1),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                       dictSemigroup)),
                                                                                                                                                  empty_1::<string,
                                                                                                                                                            &dyn Any>()))
                                                                                                  })
                                                                                  }))
    }
    pub fn Control_Monad_State_Trans_monadAskStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadAskStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadAskStateT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictMonadAsk|
                                                                                 {
                                                                                     let monadStateT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                    Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadAskusd_Dict(),
                                                                                                                      &&&add(string("ask"),
                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadTransStateT()),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_ask(),
                                                                                                                                                                                                  dictMonadAsk)),
                                                                                                                             add(string("Monad0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let monadStateT1
                                                                                                                                                      =
                                                                                                                                                      monadStateT1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &monadStateT1
                                                                                                                                              }),
                                                                                                                                 empty_1::<string,
                                                                                                                                           &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Monad_State_Trans_monadReaderStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadReaderStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadReaderStateT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonadReader|
                                                                                    {
                                                                                        let monadAskStateT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadAskStateT(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadAsk0"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonadReader)),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadReaderusd_Dict(),
                                                                                                                         &&&add(string("local"),
                                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                     &&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_mapStateT()),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_local(),
                                                                                                                                                                                                     dictMonadReader)),
                                                                                                                                add(string("MonadAsk0"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let monadAskStateT1
                                                                                                                                                         =
                                                                                                                                                         monadAskStateT1.clone();
                                                                                                                                                     move
                                                                                                                                                         |usd__unused|
                                                                                                                                                         &monadAskStateT1
                                                                                                                                                 }),
                                                                                                                                    empty_1::<string,
                                                                                                                                              &dyn Any>())))
                                                                                    }))
    }
    pub fn Control_Monad_State_Trans_monadContStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadContStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadContStateT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonadCont|
                                                                                  {
                                                                                      let monadStateT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
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
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                      let f
                                                                                                                                                                                                          =
                                                                                                                                                                                                          f.clone();
                                                                                                                                                                                                      move
                                                                                                                                                                                                          |s|
                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Class::Control_Monad_Cont_Class_callCC(),
                                                                                                                                                                                                                                                                              &&&dictMonadCont),
                                                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                                                             let s
                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                 s.clone();
                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                 |c|
                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                  let c
                                                                                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                                                                                      c.clone();
                                                                                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                                                                                      |a|
                                                                                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                         let a
                                                                                                                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                                                                                                                             a.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                                                                                                                             |s_prime|
                                                                                                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&c,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        s_prime.clone())))
                                                                                                                                                                                                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                                                                                                                                  &&&s)
                                                                                                                                                                                                                                                         }))
                                                                                                                                                                                                  }))
                                                                                                                                           }),
                                                                                                                              add(string("Monad0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let monadStateT1
                                                                                                                                                       =
                                                                                                                                                       monadStateT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &monadStateT1
                                                                                                                                               }),
                                                                                                                                  empty_1::<string,
                                                                                                                                            &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_State_Trans_monadEffectState() -> &dyn Any {
        static Control_Monad_State_Trans_monadEffectState:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadEffectState.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadEffect|
                                                                                   {
                                                                                       let Monad0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictMonadEffect)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let monadStateT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
                                                                                                                            &&&Monad0);
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                                                        &&&add(string("liftEffect"),
                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_lift(),
                                                                                                                                                                                                                                       &&&Monad0)),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                    dictMonadEffect)),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadStateT1
                                                                                                                                                        =
                                                                                                                                                        monadStateT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadStateT1
                                                                                                                                                }),
                                                                                                                                   empty_1::<string,
                                                                                                                                             &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_State_Trans_monadRecStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadRecStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadRecStateT.get_or_init(||
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
                                                                                     let Applicative0 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                 Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                                     let monadStateT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
                                                                                                                          &&&Monad0);
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                                                      &&&add(string("tailRecM"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let Applicative0
                                                                                                                                                  =
                                                                                                                                                  Applicative0.clone();
                                                                                                                                              let Bind1
                                                                                                                                                  =
                                                                                                                                                  Bind1.clone();
                                                                                                                                              let dictMonadRec
                                                                                                                                                  =
                                                                                                                                                  dictMonadRec.clone();
                                                                                                                                              move
                                                                                                                                                  |f|
                                                                                                                                                  &Func1::new({
                                                                                                                                                                  let f
                                                                                                                                                                      =
                                                                                                                                                                      f.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |a|
                                                                                                                                                                      {
                                                                                                                                                                          let f_prime =
                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                              |v|
                                                                                                                                                                                              {
                                                                                                                                                                                                  let matchValue:
                                                                                                                                                                                                          LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                  let st =
                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                    }));
                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                         &&&Bind1),
                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&st,
                                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                            })),
                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                     |v2|
                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                         let matchValue_1:
                                                                                                                                                                                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(v2);
                                                                                                                                                                                                                                                         let s1 =
                                                                                                                                                                                                                                                             match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                             };
                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                             &&&Applicative0),
                                                                                                                                                                                                                                                                                          &&{
                                                                                                                                                                                                                                                                                                let matchValue_2:
                                                                                                                                                                                                                                                                                                        LrcPtr<Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                                              });
                                                                                                                                                                                                                                                                                                match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                    Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_2_1_0)
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(matchValue_2_1_0,
                                                                                                                                                                                                                                                                                                                                                                                                                                            &s1)))),
                                                                                                                                                                                                                                                                                                    Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_2_0_0)
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(matchValue_2_0_0,
                                                                                                                                                                                                                                                                                                                                                                                                                                            &s1)))),
                                                                                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                              });
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                             let a
                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                 a.clone();
                                                                                                                                                                                                                             let f_prime
                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                 f_prime.clone();
                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                 |s_1|
                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                                                                                                                                        &&&dictMonadRec),
                                                                                                                                                                                                                                                                                                     &&&f_prime),
                                                                                                                                                                                                                                                                  &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                                                            s_1.clone())))
                                                                                                                                                                                                                         }))
                                                                                                                                                                      }
                                                                                                                                                              })
                                                                                                                                          }),
                                                                                                                             add(string("Monad0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let monadStateT1
                                                                                                                                                      =
                                                                                                                                                      monadStateT1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &monadStateT1
                                                                                                                                              }),
                                                                                                                                 empty_1::<string,
                                                                                                                                           &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Monad_State_Trans_monadStateStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadStateStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadStateStateT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonad|
                                                                                   {
                                                                                       let pure_var =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       let monadStateT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
                                                                                                                            dictMonad);
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_MonadStateusd_Dict(),
                                                                                                                        &&&add(string("state"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let pure_var
                                                                                                                                                    =
                                                                                                                                                    pure_var.clone();
                                                                                                                                                move
                                                                                                                                                    |f|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                        &&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT()),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                           &&&pure_var),
                                                                                                                                                                                                                        f))
                                                                                                                                            }),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadStateT1
                                                                                                                                                        =
                                                                                                                                                        monadStateT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadStateT1
                                                                                                                                                }),
                                                                                                                                   empty_1::<string,
                                                                                                                                             &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_State_Trans_monadTellStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadTellStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadTellStateT.get_or_init(||
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
                                                                                      let monadStateT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
                                                                                                                           &&&Monad1);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadTellusd_Dict(),
                                                                                                                       &&&add(string("tell"),
                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_lift(),
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
                                                                                                                                                       let monadStateT1
                                                                                                                                                           =
                                                                                                                                                           monadStateT1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused_1|
                                                                                                                                                           &monadStateT1
                                                                                                                                                   }),
                                                                                                                                      empty_1::<string,
                                                                                                                                                &dyn Any>()))))
                                                                                  }))
    }
    pub fn Control_Monad_State_Trans_monadWriterStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadWriterStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadWriterStateT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonadWriter|
                                                                                    {
                                                                                        let MonadTell1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadTell1"),
                                                                                                                                    Sharpurs_Prelude::unbox(dictMonadWriter)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let Monad1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad1"),
                                                                                                                                    Sharpurs_Prelude::unbox(&&MonadTell1)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let Bind1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                    Sharpurs_Prelude::unbox(&&Monad1)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let Applicative0 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                    Sharpurs_Prelude::unbox(&&Monad1)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let pure_var =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                             &&&Applicative0);
                                                                                        let pure1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                             &&&Applicative0);
                                                                                        let Monoid0 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Monoid0"),
                                                                                                                                    Sharpurs_Prelude::unbox(dictMonadWriter)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let monadTellStateT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadTellStateT(),
                                                                                                                             &&&MonadTell1);
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadWriterusd_Dict(),
                                                                                                                         &&&add(string("listen"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let Bind1
                                                                                                                                                     =
                                                                                                                                                     Bind1.clone();
                                                                                                                                                 let dictMonadWriter
                                                                                                                                                     =
                                                                                                                                                     dictMonadWriter.clone();
                                                                                                                                                 let pure_var
                                                                                                                                                     =
                                                                                                                                                     pure_var.clone();
                                                                                                                                                 move
                                                                                                                                                     |m|
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                        let m
                                                                                                                                                                                                            =
                                                                                                                                                                                                            m.clone();
                                                                                                                                                                                                        move
                                                                                                                                                                                                            |s|
                                                                                                                                                                                                            {
                                                                                                                                                                                                                let m_prime =
                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&m);
                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                       &&&Bind1),
                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_listen(),
                                                                                                                                                                                                                                                                                                                                                          &&&dictMonadWriter),
                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&m_prime,
                                                                                                                                                                                                                                                                                                                                                          s))),
                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                   |v|
                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                       let matchValue:
                                                                                                                                                                                                                                                                               LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                                                                       let activePatternResult:
                                                                                                                                                                                                                                                                               LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                           Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                              });
                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                           &&&pure_var),
                                                                                                                                                                                                                                                                                                        &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                           },
                                                                                                                                                                                                                                                                                                                                                                                                                          &match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                           })),
                                                                                                                                                                                                                                                                                                                                                                  &match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                                                                   })))
                                                                                                                                                                                                                                                                   }))
                                                                                                                                                                                                            }
                                                                                                                                                                                                    }))
                                                                                                                                             }),
                                                                                                                                add(string("pass"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let Bind1
                                                                                                                                                         =
                                                                                                                                                         Bind1.clone();
                                                                                                                                                     let dictMonadWriter
                                                                                                                                                         =
                                                                                                                                                         dictMonadWriter.clone();
                                                                                                                                                     let pure1
                                                                                                                                                         =
                                                                                                                                                         pure1.clone();
                                                                                                                                                     move
                                                                                                                                                         |m_1|
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                                            let m_1
                                                                                                                                                                                                                =
                                                                                                                                                                                                                m_1.clone();
                                                                                                                                                                                                            move
                                                                                                                                                                                                                |s_1|
                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_pass(),
                                                                                                                                                                                                                                                                                    &&&dictMonadWriter),
                                                                                                                                                                                                                                                 &&{
                                                                                                                                                                                                                                                       let m_prime_1 =
                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&m_1);
                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                              &&&Bind1),
                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&m_prime_1,
                                                                                                                                                                                                                                                                                                                                                              s_1)),
                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                          |v_1|
                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                              let matchValue_1:
                                                                                                                                                                                                                                                                                                                      LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                                                                                                                              let activePatternResult_1:
                                                                                                                                                                                                                                                                                                                      LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                     });
                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                  &&&pure1),
                                                                                                                                                                                                                                                                                                                                               &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                  },
                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                  })),
                                                                                                                                                                                                                                                                                                                                                                                                         &match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                          })))
                                                                                                                                                                                                                                                                                                          }))
                                                                                                                                                                                                                                                   })
                                                                                                                                                                                                        }))
                                                                                                                                                 }),
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
                                                                                                                                                             let monadTellStateT1
                                                                                                                                                                 =
                                                                                                                                                                 monadTellStateT1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |usd__unused_1|
                                                                                                                                                                 &monadTellStateT1
                                                                                                                                                         }),
                                                                                                                                            empty_1::<string,
                                                                                                                                                      &dyn Any>())))))
                                                                                    }))
    }
    pub fn Control_Monad_State_Trans_monadThrowStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadThrowStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadThrowStateT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadThrow|
                                                                                   {
                                                                                       let Monad0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictMonadThrow)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let monadStateT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonadThrow)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadThrowusd_Dict(),
                                                                                                                        &&&add(string("throwError"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let Monad0
                                                                                                                                                    =
                                                                                                                                                    Monad0.clone();
                                                                                                                                                let dictMonadThrow
                                                                                                                                                    =
                                                                                                                                                    dictMonadThrow.clone();
                                                                                                                                                move
                                                                                                                                                    |e|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                                           &&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadTransStateT()),
                                                                                                                                                                                                                        &&&Monad0),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                                                                                                                           &&&dictMonadThrow),
                                                                                                                                                                                                                        e))
                                                                                                                                            }),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadStateT1
                                                                                                                                                        =
                                                                                                                                                        monadStateT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadStateT1
                                                                                                                                                }),
                                                                                                                                   empty_1::<string,
                                                                                                                                             &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_State_Trans_monadErrorStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadErrorStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadErrorStateT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadError|
                                                                                   {
                                                                                       let monadThrowStateT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadThrowStateT(),
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
                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                                                                               let matchValue_1
                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                   matchValue_1.clone();
                                                                                                                                                                                                                               move
                                                                                                                                                                                                                                   |s|
                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_catchError(),
                                                                                                                                                                                                                                                                                                                                          &&&dictMonadError),
                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                          s)),
                                                                                                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                                                                                                      let s
                                                                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                                                                          s.clone();
                                                                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                                                                          |e|
                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                                                         e)),
                                                                                                                                                                                                                                                                                                                           &&&s)
                                                                                                                                                                                                                                                                                  }))
                                                                                                                                                                                                                           }))
                                                                                                                                                                        }
                                                                                                                                                                })
                                                                                                                                            }),
                                                                                                                               add(string("MonadThrow0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadThrowStateT1
                                                                                                                                                        =
                                                                                                                                                        monadThrowStateT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadThrowStateT1
                                                                                                                                                }),
                                                                                                                                   empty_1::<string,
                                                                                                                                             &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_State_Trans_monadSTStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadSTStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadSTStateT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonadST|
                                                                                {
                                                                                    let Monad0 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                Sharpurs_Prelude::unbox(dictMonadST)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    let monadStateT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
                                                                                                                         &&&Monad0);
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                                                     &&&add(string("liftST"),
                                                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_lift(),
                                                                                                                                                                                                                                    &&&Monad0)),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_liftST(),
                                                                                                                                                                                                 dictMonadST)),
                                                                                                                            add(string("Monad0"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let monadStateT1
                                                                                                                                                     =
                                                                                                                                                     monadStateT1.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused|
                                                                                                                                                     &monadStateT1
                                                                                                                                             }),
                                                                                                                                empty_1::<string,
                                                                                                                                          &dyn Any>())))
                                                                                }))
    }
    pub fn Control_Monad_State_Trans_monoidStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monoidStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monoidStateT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictMonad|
                                                                               {
                                                                                   let applicativeStateT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_applicativeStateT(),
                                                                                                                        dictMonad);
                                                                                   let semigroupStateT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_semigroupStateT(),
                                                                                                                        dictMonad);
                                                                                   &Func1::new({
                                                                                                   let applicativeStateT1
                                                                                                       =
                                                                                                       applicativeStateT1.clone();
                                                                                                   let semigroupStateT1
                                                                                                       =
                                                                                                       semigroupStateT1.clone();
                                                                                                   move
                                                                                                       |dictMonoid|
                                                                                                       {
                                                                                                           let semigroupStateT2 =
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&semigroupStateT1,
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                                            &&&add(string("mempty"),
                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                        &&&applicativeStateT1),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                        dictMonoid)),
                                                                                                                                                   add(string("Semigroup0"),
                                                                                                                                                       &&Func1::new({
                                                                                                                                                                        let semigroupStateT2
                                                                                                                                                                            =
                                                                                                                                                                            semigroupStateT2.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |usd__unused|
                                                                                                                                                                            &semigroupStateT2
                                                                                                                                                                    }),
                                                                                                                                                       empty_1::<string,
                                                                                                                                                                 &dyn Any>())))
                                                                                                       }
                                                                                               })
                                                                               }))
    }
    pub fn Control_Monad_State_Trans_altStateT() -> &dyn Any {
        static Control_Monad_State_Trans_altStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_altStateT.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictMonad|
                                                                            &Func1::new(move
                                                                                            |dictAlt|
                                                                                            {
                                                                                                let functorStateT1 =
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_functorStateT(),
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
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                        let matchValue_1
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            matchValue_1.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |s|
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictAlt),
                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                   s)),
                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                s))
                                                                                                                                                                                                                                    }))
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     }),
                                                                                                                                        add(string("Functor0"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let functorStateT1
                                                                                                                                                                 =
                                                                                                                                                                 functorStateT1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |usd__unused|
                                                                                                                                                                 &functorStateT1
                                                                                                                                                         }),
                                                                                                                                            empty_1::<string,
                                                                                                                                                      &dyn Any>())))
                                                                                            })))
    }
    pub fn Control_Monad_State_Trans_plusStateT() -> &dyn Any {
        static Control_Monad_State_Trans_plusStateT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_State_Trans_plusStateT.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictMonad|
                                                                             {
                                                                                 let altStateT1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_altStateT(),
                                                                                                                      dictMonad);
                                                                                 &Func1::new({
                                                                                                 let altStateT1
                                                                                                     =
                                                                                                     altStateT1.clone();
                                                                                                 move
                                                                                                     |dictPlus|
                                                                                                     {
                                                                                                         let empty =
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                                              dictPlus);
                                                                                                         let altStateT2 =
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&altStateT1,
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictPlus)),
                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                                                                          &&&add(string("empty"),
                                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_StateT(),
                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                     let empty
                                                                                                                                                                                                         =
                                                                                                                                                                                                         empty.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |v|
                                                                                                                                                                                                         &empty
                                                                                                                                                                                                 })),
                                                                                                                                                 add(string("Alt0"),
                                                                                                                                                     &&Func1::new({
                                                                                                                                                                      let altStateT2
                                                                                                                                                                          =
                                                                                                                                                                          altStateT2.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |usd__unused|
                                                                                                                                                                          &altStateT2
                                                                                                                                                                  }),
                                                                                                                                                     empty_1::<string,
                                                                                                                                                               &dyn Any>())))
                                                                                                     }
                                                                                             })
                                                                             }))
    }
    pub fn Control_Monad_State_Trans_alternativeStateT() -> &dyn Any {
        static Control_Monad_State_Trans_alternativeStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_alternativeStateT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonad|
                                                                                    {
                                                                                        let applicativeStateT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_applicativeStateT(),
                                                                                                                             dictMonad);
                                                                                        let plusStateT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_plusStateT(),
                                                                                                                             dictMonad);
                                                                                        &Func1::new({
                                                                                                        let applicativeStateT1
                                                                                                            =
                                                                                                            applicativeStateT1.clone();
                                                                                                        let plusStateT1
                                                                                                            =
                                                                                                            plusStateT1.clone();
                                                                                                        move
                                                                                                            |dictAlternative|
                                                                                                            {
                                                                                                                let plusStateT2 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&plusStateT1,
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                                                                                 &&&add(string("Applicative0"),
                                                                                                                                                        &&Func1::new(move
                                                                                                                                                                         |usd__unused|
                                                                                                                                                                         &applicativeStateT1),
                                                                                                                                                        add(string("Plus1"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let plusStateT2
                                                                                                                                                                                 =
                                                                                                                                                                                 plusStateT2.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |usd__unused_1|
                                                                                                                                                                                 &plusStateT2
                                                                                                                                                                         }),
                                                                                                                                                            empty_1::<string,
                                                                                                                                                                      &dyn Any>())))
                                                                                                            }
                                                                                                    })
                                                                                    }))
    }
    pub fn Control_Monad_State_Trans_monadPlusStateT() -> &dyn Any {
        static Control_Monad_State_Trans_monadPlusStateT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Trans_monadPlusStateT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonadPlus|
                                                                                  {
                                                                                      let Monad0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadPlus)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let monadStateT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadStateT(),
                                                                                                                           &&&Monad0);
                                                                                      let alternativeStateT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_alternativeStateT(),
                                                                                                                                                              &&&Monad0),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alternative1"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonadPlus)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                                                       &&&add(string("Monad0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let monadStateT1
                                                                                                                                                   =
                                                                                                                                                   monadStateT1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &monadStateT1
                                                                                                                                           }),
                                                                                                                              add(string("Alternative1"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let alternativeStateT1
                                                                                                                                                       =
                                                                                                                                                       alternativeStateT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused_1|
                                                                                                                                                       &alternativeStateT1
                                                                                                                                               }),
                                                                                                                                  empty_1::<string,
                                                                                                                                            &dyn Any>())))
                                                                                  }))
    }
}
