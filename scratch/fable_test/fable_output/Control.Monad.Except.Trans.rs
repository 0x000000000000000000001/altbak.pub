pub mod PureScript_Control_Monad_Except_Trans {
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
    use crate::module_86d6df2::PureScript_Control_Category;
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
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Control_Monad_Except_Trans_ExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_ExceptT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_ExceptT.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Control_Monad_Except_Trans_withExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_withExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_withExceptT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictFunctor|
                                                                               &Func1::new({
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
                                                                                                                           let mapLeft =
                                                                                                                               &Func1::new(move
                                                                                                                                               |v1|
                                                                                                                                               &Func1::new({
                                                                                                                                                               let v1
                                                                                                                                                                   =
                                                                                                                                                                   v1.clone();
                                                                                                                                                               move
                                                                                                                                                                   |v2|
                                                                                                                                                                   {
                                                                                                                                                                       let matchValue_3 =
                                                                                                                                                                           Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                                                       let matchValue_4:
                                                                                                                                                                               LrcPtr<Data_Either_Either> =
                                                                                                                                                                           Sharpurs_Prelude::unbox(v2);
                                                                                                                                                                       match matchValue_4.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                           Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_4_0_0)
                                                                                                                                                                           =>
                                                                                                                                                                           &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue_3,
                                                                                                                                                                                                                                                                      &&matchValue_4_0_0))),
                                                                                                                                                                           Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0)
                                                                                                                                                                           =>
                                                                                                                                                                           &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0)),
                                                                                                                                                                       }
                                                                                                                                                                   }
                                                                                                                                                           }));
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                               &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT()),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                     &&&dictFunctor),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&mapLeft,
                                                                                                                                                                                                                                                                     &&&matchValue)),
                                                                                                                                                                                               &&&matchValue_1))
                                                                                                                       }
                                                                                                               })
                                                                                           })))
    }
    pub fn Control_Monad_Except_Trans_runExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_runExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_runExceptT.get_or_init(||
                                                              &Func1::new(move
                                                                              |v|
                                                                              &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Monad_Except_Trans_newtypeExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_newtypeExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_newtypeExceptT.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                   &&&add(string("Coercible0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &Sharpurs_Prelude::Prim_undefined()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
    }
    pub fn Control_Monad_Except_Trans_monadTransExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadTransExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadTransExceptT.get_or_init(||
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
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT(),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                &&&Bind1),
                                                                                                                                                                                                                                                             m),
                                                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                                                            |a|
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                &&&pure_var),
                                                                                                                                                                                                                                                                             &&&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(a.clone()))))))
                                                                                                                                              })
                                                                                                                              }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
    }
    pub fn Control_Monad_Except_Trans_lift() -> &dyn Any {
        static Control_Monad_Except_Trans_lift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_lift.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                         &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadTransExceptT()))
    }
    pub fn Control_Monad_Except_Trans_mapExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_mapExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_mapExceptT.get_or_init(||
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
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                          &&&matchValue_1))
                                                                                                  }
                                                                                          })))
    }
    pub fn Control_Monad_Except_Trans_functorExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_functorExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_functorExceptT.get_or_init(||
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
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_mapExceptT(),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                      &&&dictFunctor),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Either::Data_Either_functorEither()),
                                                                                                                                                                                                                                                      f)))
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))
    }
    pub fn Control_Monad_Except_Trans_except() -> &dyn Any {
        static Control_Monad_Except_Trans_except: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_except.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictApplicative|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                              &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT()),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                              dictApplicative))))
    }
    pub fn Control_Monad_Except_Trans_monadExceptT_004026() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT_tco(dictMonad))
    }
    pub fn Control_Monad_Except_Trans_monadExceptT_004026_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Except_Trans_monadExceptT_004026_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadExceptT_004026_002d1.get_or_init(||
                                                                             Lazy(Control_Monad_Except_Trans_monadExceptT_004026.clone()))
    }
    pub fn Control_Monad_Except_Trans_bindExceptT_004028() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_bindExceptT_tco(dictMonad))
    }
    pub fn Control_Monad_Except_Trans_bindExceptT_004028_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Except_Trans_bindExceptT_004028_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_bindExceptT_004028_002d1.get_or_init(||
                                                                            Lazy(Control_Monad_Except_Trans_bindExceptT_004028.clone()))
    }
    pub fn Control_Monad_Except_Trans_applyExceptT_004030() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_applyExceptT_tco(dictMonad))
    }
    pub fn Control_Monad_Except_Trans_applyExceptT_004030_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Except_Trans_applyExceptT_004030_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_applyExceptT_004030_002d1.get_or_init(||
                                                                             Lazy(Control_Monad_Except_Trans_applyExceptT_004030.clone()))
    }
    pub fn Control_Monad_Except_Trans_applicativeExceptT_004032()
     -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_applicativeExceptT_tco(dictMonad))
    }
    pub fn Control_Monad_Except_Trans_applicativeExceptT_004032_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Except_Trans_applicativeExceptT_004032_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_applicativeExceptT_004032_002d1.get_or_init(||
                                                                                   Lazy(Control_Monad_Except_Trans_applicativeExceptT_004032.clone()))
    }
    pub fn Control_Monad_Except_Trans_monadExceptT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                         &&&add(string("Applicative0"),
                                                &&Func1::new({
                                                                 let dictMonad
                                                                     =
                                                                     dictMonad.clone();
                                                                 move
                                                                     |usd__unused|
                                                                     PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_applicativeExceptT_tco(&&dictMonad)
                                                             }),
                                                add(string("Bind1"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused_1|
                                                                         PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_bindExceptT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Except_Trans_monadExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadExceptT.get_or_init(||
                                                                Control_Monad_Except_Trans_monadExceptT_004026_002d1.Value)
    }
    pub fn Control_Monad_Except_Trans_bindExceptT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let pure_var =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                       Sharpurs_Prelude::unbox(dictMonad)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                         &&&add(string("bind"),
                                                &&Func1::new({
                                                                 let Bind1 =
                                                                     Bind1.clone();
                                                                 let pure_var
                                                                     =
                                                                     pure_var.clone();
                                                                 move |v|
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
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                       &&&Bind1),
                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                             &&&pure_var),
                                                                                                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                                                                                                            |usd__arg1|
                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))))),
                                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                                      let matchValue_1
                                                                                                                                                                                                                          =
                                                                                                                                                                                                                          matchValue_1.clone();
                                                                                                                                                                                                                      move
                                                                                                                                                                                                                          |a|
                                                                                                                                                                                                                          &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                      a))
                                                                                                                                                                                                                  }))))
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
                                                                         PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_applyExceptT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Except_Trans_bindExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_bindExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_bindExceptT.get_or_init(||
                                                               Control_Monad_Except_Trans_bindExceptT_004028_002d1.Value)
    }
    pub fn Control_Monad_Except_Trans_applyExceptT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let functorExceptT1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_functorExceptT(),
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
                                                                                  &&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT_tco(dictMonad)),
                                                add(string("Functor0"),
                                                    &&Func1::new({
                                                                     let functorExceptT1
                                                                         =
                                                                         functorExceptT1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &functorExceptT1
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Except_Trans_applyExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_applyExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_applyExceptT.get_or_init(||
                                                                Control_Monad_Except_Trans_applyExceptT_004030_002d1.Value)
    }
    pub fn Control_Monad_Except_Trans_applicativeExceptT_tco(dictMonad:
                                                                 &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                         &&&add(string("pure"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT()),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                     &&&Func1::new(move
                                                                                                                                       |usd__arg1|
                                                                                                                                       &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone()))))),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_applyExceptT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Except_Trans_applicativeExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_applicativeExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_applicativeExceptT.get_or_init(||
                                                                      Control_Monad_Except_Trans_applicativeExceptT_004032_002d1.Value)
    }
    pub fn Control_Monad_Except_Trans_semigroupExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_semigroupExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_semigroupExceptT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonad|
                                                                                    {
                                                                                        let applyExceptT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_applyExceptT(),
                                                                                                                             dictMonad);
                                                                                        &Func1::new({
                                                                                                        let applyExceptT1
                                                                                                            =
                                                                                                            applyExceptT1.clone();
                                                                                                        move
                                                                                                            |dictSemigroup|
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                                             &&&add(string("append"),
                                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                         &&&applyExceptT1),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                         dictSemigroup)),
                                                                                                                                                    empty::<string,
                                                                                                                                                            &dyn Any>()))
                                                                                                    })
                                                                                    }))
    }
    pub fn Control_Monad_Except_Trans_monadAskExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadAskExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadAskExceptT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadAsk|
                                                                                   {
                                                                                       let monadExceptT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadAskusd_Dict(),
                                                                                                                        &&&add(string("ask"),
                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                       &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadTransExceptT()),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_ask(),
                                                                                                                                                                                                    dictMonadAsk)),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadExceptT1
                                                                                                                                                        =
                                                                                                                                                        monadExceptT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadExceptT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Except_Trans_monadReaderExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadReaderExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadReaderExceptT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonadReader|
                                                                                      {
                                                                                          let monadAskExceptT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadAskExceptT(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadAsk0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictMonadReader)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadReaderusd_Dict(),
                                                                                                                           &&&add(string("local"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let dictMonadReader
                                                                                                                                                       =
                                                                                                                                                       dictMonadReader.clone();
                                                                                                                                                   move
                                                                                                                                                       |f|
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_mapExceptT(),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_local(),
                                                                                                                                                                                                                                                              &&&dictMonadReader),
                                                                                                                                                                                                                           f))
                                                                                                                                               }),
                                                                                                                                  add(string("MonadAsk0"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let monadAskExceptT1
                                                                                                                                                           =
                                                                                                                                                           monadAskExceptT1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &monadAskExceptT1
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))
                                                                                      }))
    }
    pub fn Control_Monad_Except_Trans_monadContExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadContExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadContExceptT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonadCont|
                                                                                    {
                                                                                        let monadExceptT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
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
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                         &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT()),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Class::Control_Monad_Cont_Class_callCC(),
                                                                                                                                                                                                                                                            &&&dictMonadCont),
                                                                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                                                                           let f
                                                                                                                                                                                                                                               =
                                                                                                                                                                                                                                               f.clone();
                                                                                                                                                                                                                                           move
                                                                                                                                                                                                                                               |c|
                                                                                                                                                                                                                                               &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                                                                                                                             let c
                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                 c.clone();
                                                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                                                 |a|
                                                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT()),
                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&c,
                                                                                                                                                                                                                                                                                                                                                                                                     &&&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(a.clone()))))
                                                                                                                                                                                                                                                                                                                         })))
                                                                                                                                                                                                                                       })))
                                                                                                                                             }),
                                                                                                                                add(string("Monad0"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let monadExceptT1
                                                                                                                                                         =
                                                                                                                                                         monadExceptT1.clone();
                                                                                                                                                     move
                                                                                                                                                         |usd__unused|
                                                                                                                                                         &monadExceptT1
                                                                                                                                                 }),
                                                                                                                                    empty::<string,
                                                                                                                                            &dyn Any>())))
                                                                                    }))
    }
    pub fn Control_Monad_Except_Trans_monadEffectExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadEffectExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadEffectExceptT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonadEffect|
                                                                                      {
                                                                                          let Monad0 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                      Sharpurs_Prelude::unbox(dictMonadEffect)),
                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                                          let monadExceptT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
                                                                                                                               &&&Monad0);
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                                                           &&&add(string("liftEffect"),
                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_lift(),
                                                                                                                                                                                                                                          &&&Monad0)),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                       dictMonadEffect)),
                                                                                                                                  add(string("Monad0"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let monadExceptT1
                                                                                                                                                           =
                                                                                                                                                           monadExceptT1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &monadExceptT1
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))
                                                                                      }))
    }
    pub fn Control_Monad_Except_Trans_monadRecExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadRecExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadRecExceptT.get_or_init(||
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
                                                                                       let monadExceptT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
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
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                        &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT()),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                                                           &&&dictMonadRec),
                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                          let f
                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                              f.clone();
                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                              |a|
                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                  let m =
                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                 a));
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                         &&&Bind1),
                                                                                                                                                                                                                                                                                                                      &&&m),
                                                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                                                     |m_prime|
                                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                         &&&Applicative0),
                                                                                                                                                                                                                                                                                                                                      &&{
                                                                                                                                                                                                                                                                                                                                            let matchValue:
                                                                                                                                                                                                                                                                                                                                                    LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(m_prime);
                                                                                                                                                                                                                                                                                                                                            if let Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                                                                                                                                   matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                let activePatternResult:
                                                                                                                                                                                                                                                                                                                                                        LrcPtr<Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                                                                           Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                           _
                                                                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                                                                                                                                                                                       });
                                                                                                                                                                                                                                                                                                                                                if let Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(activePatternResult_0_0)
                                                                                                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                                                                                                       activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                                                                                                         Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                                                                                                                                                    let activePatternResult_1:
                                                                                                                                                                                                                                                                                                                                                            LrcPtr<Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                               Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                               _
                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                                                                                                                                                                                           });
                                                                                                                                                                                                                                                                                                                                                    if let Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(activePatternResult_1_1_0)
                                                                                                                                                                                                                                                                                                                                                           =
                                                                                                                                                                                                                                                                                                                                                           activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    }))))
                                                                                                                                                                                                                                                                                                                                                    } else {
                                                                                                                                                                                                                                                                                                                                                        panic!("{}",
                                                                                                                                                                                                                                                                                                                                                               string("Match failure: PureScript_Data_Either.Data_Either_Either"),)
                                                                                                                                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                                                                                                                            } else {
                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               _
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           }))))
                                                                                                                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                                                                                                                        })))
                                                                                                                                                                                                                                              }
                                                                                                                                                                                                                                      })))
                                                                                                                                            }),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadExceptT1
                                                                                                                                                        =
                                                                                                                                                        monadExceptT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadExceptT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Except_Trans_monadStateExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadStateExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadStateExceptT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonadState|
                                                                                     {
                                                                                         let Monad0 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                     Sharpurs_Prelude::unbox(dictMonadState)),
                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                                         let monadExceptT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                        Sharpurs_Prelude::unbox(dictMonadState)),
                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_MonadStateusd_Dict(),
                                                                                                                          &&&add(string("state"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let Monad0
                                                                                                                                                      =
                                                                                                                                                      Monad0.clone();
                                                                                                                                                  let dictMonadState
                                                                                                                                                      =
                                                                                                                                                      dictMonadState.clone();
                                                                                                                                                  move
                                                                                                                                                      |f|
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                                             &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadTransExceptT()),
                                                                                                                                                                                                                          &&&Monad0),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                                                                                                                             &&&dictMonadState),
                                                                                                                                                                                                                          f))
                                                                                                                                              }),
                                                                                                                                 add(string("Monad0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let monadExceptT1
                                                                                                                                                          =
                                                                                                                                                          monadExceptT1.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &monadExceptT1
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }))
    }
    pub fn Control_Monad_Except_Trans_monadTellExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadTellExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadTellExceptT.get_or_init(||
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
                                                                                        let monadExceptT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
                                                                                                                             &&&Monad1);
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadTellusd_Dict(),
                                                                                                                         &&&add(string("tell"),
                                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_lift(),
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
                                                                                                                                                         let monadExceptT1
                                                                                                                                                             =
                                                                                                                                                             monadExceptT1.clone();
                                                                                                                                                         move
                                                                                                                                                             |usd__unused_1|
                                                                                                                                                             &monadExceptT1
                                                                                                                                                     }),
                                                                                                                                        empty::<string,
                                                                                                                                                &dyn Any>()))))
                                                                                    }))
    }
    pub fn Control_Monad_Except_Trans_monadWriterExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadWriterExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadWriterExceptT.get_or_init(||
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
                                                                                          let pure_var =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(&&Monad1)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          let Applicative0 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                      Sharpurs_Prelude::unbox(&&Monad1)),
                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                                          let Monoid0 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Monoid0"),
                                                                                                                                      Sharpurs_Prelude::unbox(dictMonadWriter)),
                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                                          let monadTellExceptT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadTellExceptT(),
                                                                                                                               &&&MonadTell1);
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadWriterusd_Dict(),
                                                                                                                           &&&add(string("listen"),
                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_mapExceptT(),
                                                                                                                                                                    &&&Func1::new({
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
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                 &&&Bind1),
                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_listen(),
                                                                                                                                                                                                                                                                                                                                    &&&dictMonadWriter),
                                                                                                                                                                                                                                                                                                 m)),
                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                             |v|
                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                 let matchValue:
                                                                                                                                                                                                                                                         LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                     &&&pure_var),
                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Either::Data_Either_functorEither()),
                                                                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                          |r|
                                                                                                                                                                                                                                                                                                                                                                          &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(r.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                  &match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                   })))),
                                                                                                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                        }))
                                                                                                                                                                                                                                             }))
                                                                                                                                                                                  })),
                                                                                                                                  add(string("pass"),
                                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_mapExceptT(),
                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                          let Applicative0
                                                                                                                                                                                              =
                                                                                                                                                                                              Applicative0.clone();
                                                                                                                                                                                          let Bind1
                                                                                                                                                                                              =
                                                                                                                                                                                              Bind1.clone();
                                                                                                                                                                                          let dictMonadWriter
                                                                                                                                                                                              =
                                                                                                                                                                                              dictMonadWriter.clone();
                                                                                                                                                                                          move
                                                                                                                                                                                              |m_1|
                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_pass(),
                                                                                                                                                                                                                                                                  &&&dictMonadWriter),
                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                        &&&Bind1),
                                                                                                                                                                                                                                                                                                     m_1),
                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                    |a_1|
                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                        &&&Applicative0),
                                                                                                                                                                                                                                                                                                                     &&{
                                                                                                                                                                                                                                                                                                                           let matchValue_1:
                                                                                                                                                                                                                                                                                                                                   LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(a_1);
                                                                                                                                                                                                                                                                                                                           match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                               Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                   let activePatternResult:
                                                                                                                                                                                                                                                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::_007cUnbox_007c(matchValue_1_1_0);
                                                                                                                                                                                                                                                                                                                                   &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       })),
                                                                                                                                                                                                                                                                                                                                                                                           &match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                            }))
                                                                                                                                                                                                                                                                                                                               }
                                                                                                                                                                                                                                                                                                                               Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                               &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)),
                                                                                                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Category::Control_Category_categoryFn()))),
                                                                                                                                                                                                                                                                                                                           }
                                                                                                                                                                                                                                                                                                                       }))))
                                                                                                                                                                                      })),
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
                                                                                                                                                               let monadTellExceptT1
                                                                                                                                                                   =
                                                                                                                                                                   monadTellExceptT1.clone();
                                                                                                                                                               move
                                                                                                                                                                   |usd__unused_1|
                                                                                                                                                                   &monadTellExceptT1
                                                                                                                                                           }),
                                                                                                                                              empty::<string,
                                                                                                                                                      &dyn Any>())))))
                                                                                      }))
    }
    pub fn Control_Monad_Except_Trans_monadThrowExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadThrowExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadThrowExceptT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictMonad|
                                                                                     {
                                                                                         let monadExceptT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
                                                                                                                              dictMonad);
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadThrowusd_Dict(),
                                                                                                                          &&&add(string("throwError"),
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                      &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT()),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                        |usd__arg1|
                                                                                                                                                                                                                        &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))))),
                                                                                                                                 add(string("Monad0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let monadExceptT1
                                                                                                                                                          =
                                                                                                                                                          monadExceptT1.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &monadExceptT1
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }))
    }
    pub fn Control_Monad_Except_Trans_monadErrorExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadErrorExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadErrorExceptT.get_or_init(||
                                                                     &Func1::new(move
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
                                                                                         let monadThrowExceptT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadThrowExceptT(),
                                                                                                                              dictMonad);
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadErrorusd_Dict(),
                                                                                                                          &&&add(string("catchError"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let Bind1
                                                                                                                                                      =
                                                                                                                                                      Bind1.clone();
                                                                                                                                                  let pure_var
                                                                                                                                                      =
                                                                                                                                                      pure_var.clone();
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
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT(),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                        &&&Bind1),
                                                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                          let matchValue_1
                                                                                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                                                                                              matchValue_1.clone();
                                                                                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                                                                                              |a|
                                                                                                                                                                                                                                                                                                                                              &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                                                                          a))
                                                                                                                                                                                                                                                                                                                                      })),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                           &&&pure_var),
                                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                          |usd__arg1|
                                                                                                                                                                                                                                                                                                                                          &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone())))))))
                                                                                                                                                                          }
                                                                                                                                                                  })
                                                                                                                                              }),
                                                                                                                                 add(string("MonadThrow0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let monadThrowExceptT1
                                                                                                                                                          =
                                                                                                                                                          monadThrowExceptT1.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &monadThrowExceptT1
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }))
    }
    pub fn Control_Monad_Except_Trans_monadSTExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadSTExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadSTExceptT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonadST|
                                                                                  {
                                                                                      let Monad0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadST)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let monadExceptT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
                                                                                                                           &&&Monad0);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                                                       &&&add(string("liftST"),
                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_lift(),
                                                                                                                                                                                                                                      &&&Monad0)),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_liftST(),
                                                                                                                                                                                                   dictMonadST)),
                                                                                                                              add(string("Monad0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let monadExceptT1
                                                                                                                                                       =
                                                                                                                                                       monadExceptT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &monadExceptT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_Except_Trans_monoidExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monoidExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monoidExceptT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictMonad|
                                                                                 {
                                                                                     let applicativeExceptT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_applicativeExceptT(),
                                                                                                                          dictMonad);
                                                                                     let semigroupExceptT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_semigroupExceptT(),
                                                                                                                          dictMonad);
                                                                                     &Func1::new({
                                                                                                     let applicativeExceptT1
                                                                                                         =
                                                                                                         applicativeExceptT1.clone();
                                                                                                     let semigroupExceptT1
                                                                                                         =
                                                                                                         semigroupExceptT1.clone();
                                                                                                     move
                                                                                                         |dictMonoid|
                                                                                                         {
                                                                                                             let semigroupExceptT2 =
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&semigroupExceptT1,
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                                              &&&add(string("mempty"),
                                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                          &&&applicativeExceptT1),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                          dictMonoid)),
                                                                                                                                                     add(string("Semigroup0"),
                                                                                                                                                         &&Func1::new({
                                                                                                                                                                          let semigroupExceptT2
                                                                                                                                                                              =
                                                                                                                                                                              semigroupExceptT2.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |usd__unused|
                                                                                                                                                                              &semigroupExceptT2
                                                                                                                                                                      }),
                                                                                                                                                         empty::<string,
                                                                                                                                                                 &dyn Any>())))
                                                                                                         }
                                                                                                 })
                                                                                 }))
    }
    pub fn Control_Monad_Except_Trans_altExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_altExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_altExceptT.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictSemigroup|
                                                                              &Func1::new({
                                                                                              let dictSemigroup
                                                                                                  =
                                                                                                  dictSemigroup.clone();
                                                                                              move
                                                                                                  |dictMonad|
                                                                                                  {
                                                                                                      let Bind1 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                  Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                      let Applicative0 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                  Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                      let functorExceptT1 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_functorExceptT(),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                                                                       &&&add(string("alt"),
                                                                                                                                              &&Func1::new({
                                                                                                                                                               let Applicative0
                                                                                                                                                                   =
                                                                                                                                                                   Applicative0.clone();
                                                                                                                                                               let Bind1
                                                                                                                                                                   =
                                                                                                                                                                   Bind1.clone();
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
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_ExceptT(),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                     &&&Bind1),
                                                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                                                                 let matchValue_1
                                                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                                                     matchValue_1.clone();
                                                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                                                     |rm|
                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                         let matchValue_3:
                                                                                                                                                                                                                                                                                                 LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(rm);
                                                                                                                                                                                                                                                                                         match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                             Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_3_0_0)
                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&Bind1),
                                                                                                                                                                                                                                                                                                                                                                 &&&matchValue_1),
                                                                                                                                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                let matchValue_3
                                                                                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                                                                                    matchValue_3.clone();
                                                                                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                                                                                    |rn|
                                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                                        let matchValue_4:
                                                                                                                                                                                                                                                                                                                                                                LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(rn);
                                                                                                                                                                                                                                                                                                                                                        match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                                            Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_4_0_0)
                                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                                                                                &&&Applicative0),
                                                                                                                                                                                                                                                                                                                                                                                             &&&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&matchValue_4_0_0)))),
                                                                                                                                                                                                                                                                                                                                                            Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0)
                                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                                                                                &&&Applicative0),
                                                                                                                                                                                                                                                                                                                                                                                             &&&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0))),
                                                                                                                                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                                                                                                                            })),
                                                                                                                                                                                                                                                                                             Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_3_1_0)
                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                 &&&Applicative0),
                                                                                                                                                                                                                                                                                                                              &&&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_3_1_0))),
                                                                                                                                                                                                                                                                                         }
                                                                                                                                                                                                                                                                                     }
                                                                                                                                                                                                                                                                             })))
                                                                                                                                                                                       }
                                                                                                                                                                               })
                                                                                                                                                           }),
                                                                                                                                              add(string("Functor0"),
                                                                                                                                                  &&Func1::new({
                                                                                                                                                                   let functorExceptT1
                                                                                                                                                                       =
                                                                                                                                                                       functorExceptT1.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |usd__unused|
                                                                                                                                                                       &functorExceptT1
                                                                                                                                                               }),
                                                                                                                                                  empty::<string,
                                                                                                                                                          &dyn Any>())))
                                                                                                  }
                                                                                          })))
    }
    pub fn Control_Monad_Except_Trans_plusExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_plusExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_plusExceptT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictMonoid|
                                                                               {
                                                                                   let mempty =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                        dictMonoid);
                                                                                   let altExceptT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_altExceptT(),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                   &Func1::new({
                                                                                                   let altExceptT1
                                                                                                       =
                                                                                                       altExceptT1.clone();
                                                                                                   let mempty
                                                                                                       =
                                                                                                       mempty.clone();
                                                                                                   move
                                                                                                       |dictMonad|
                                                                                                       {
                                                                                                           let altExceptT2 =
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&altExceptT1,
                                                                                                                                                dictMonad);
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                                                                            &&&add(string("empty"),
                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadThrowExceptT(),
                                                                                                                                                                                                                                                           dictMonad)),
                                                                                                                                                                                     &&&mempty),
                                                                                                                                                   add(string("Alt0"),
                                                                                                                                                       &&Func1::new({
                                                                                                                                                                        let altExceptT2
                                                                                                                                                                            =
                                                                                                                                                                            altExceptT2.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |usd__unused|
                                                                                                                                                                            &altExceptT2
                                                                                                                                                                    }),
                                                                                                                                                       empty::<string,
                                                                                                                                                               &dyn Any>())))
                                                                                                       }
                                                                                               })
                                                                               }))
    }
    pub fn Control_Monad_Except_Trans_alternativeExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_alternativeExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_alternativeExceptT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonoid|
                                                                                      {
                                                                                          let plusExceptT1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_plusExceptT(),
                                                                                                                               dictMonoid);
                                                                                          &Func1::new({
                                                                                                          let plusExceptT1
                                                                                                              =
                                                                                                              plusExceptT1.clone();
                                                                                                          move
                                                                                                              |dictMonad|
                                                                                                              {
                                                                                                                  let applicativeExceptT1 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_applicativeExceptT(),
                                                                                                                                                       dictMonad);
                                                                                                                  let plusExceptT2 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&plusExceptT1,
                                                                                                                                                       dictMonad);
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                                                                                   &&&add(string("Applicative0"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let applicativeExceptT1
                                                                                                                                                                               =
                                                                                                                                                                               applicativeExceptT1.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |usd__unused|
                                                                                                                                                                               &applicativeExceptT1
                                                                                                                                                                       }),
                                                                                                                                                          add(string("Plus1"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let plusExceptT2
                                                                                                                                                                                   =
                                                                                                                                                                                   plusExceptT2.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |usd__unused_1|
                                                                                                                                                                                   &plusExceptT2
                                                                                                                                                                           }),
                                                                                                                                                              empty::<string,
                                                                                                                                                                      &dyn Any>())))
                                                                                                              }
                                                                                                      })
                                                                                      }))
    }
    pub fn Control_Monad_Except_Trans_monadPlusExceptT() -> &dyn Any {
        static Control_Monad_Except_Trans_monadPlusExceptT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_Trans_monadPlusExceptT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonoid|
                                                                                    {
                                                                                        let alternativeExceptT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_alternativeExceptT(),
                                                                                                                             dictMonoid);
                                                                                        &Func1::new({
                                                                                                        let alternativeExceptT1
                                                                                                            =
                                                                                                            alternativeExceptT1.clone();
                                                                                                        move
                                                                                                            |dictMonad|
                                                                                                            {
                                                                                                                let monadExceptT1 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadExceptT(),
                                                                                                                                                     dictMonad);
                                                                                                                let alternativeExceptT2 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&alternativeExceptT1,
                                                                                                                                                     dictMonad);
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                                                                                 &&&add(string("Monad0"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let monadExceptT1
                                                                                                                                                                             =
                                                                                                                                                                             monadExceptT1.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |usd__unused|
                                                                                                                                                                             &monadExceptT1
                                                                                                                                                                     }),
                                                                                                                                                        add(string("Alternative1"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let alternativeExceptT2
                                                                                                                                                                                 =
                                                                                                                                                                                 alternativeExceptT2.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |usd__unused_1|
                                                                                                                                                                                 &alternativeExceptT2
                                                                                                                                                                         }),
                                                                                                                                                            empty::<string,
                                                                                                                                                                    &dyn Any>())))
                                                                                                            }
                                                                                                    })
                                                                                    }))
    }
}
