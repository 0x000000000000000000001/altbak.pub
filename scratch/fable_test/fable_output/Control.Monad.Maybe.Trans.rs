pub mod PureScript_Control_Monad_Maybe_Trans {
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
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Control_Monad_Maybe_Trans_MaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_MaybeT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_MaybeT.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Control_Monad_Maybe_Trans_runMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_runMaybeT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_runMaybeT.get_or_init(||
                                                            &Func1::new(move
                                                                            |v|
                                                                            &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Monad_Maybe_Trans_newtypeMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_newtypeMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_newtypeMaybeT.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                 &&&add(string("Coercible0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &Sharpurs_Prelude::Prim_undefined()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))
    }
    pub fn Control_Monad_Maybe_Trans_monadTransMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadTransMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadTransMaybeT.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_MonadTransusd_Dict(),
                                                                                                    &&&add(string("lift"),
                                                                                                           &&Func1::new(move
                                                                                                                            |dictMonad|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT()),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_liftM1(),
                                                                                                                                                                                                                                   dictMonad),
                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                  |usd__arg1|
                                                                                                                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))))),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
    }
    pub fn Control_Monad_Maybe_Trans_lift() -> &dyn Any {
        static Control_Monad_Maybe_Trans_lift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_lift.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                        &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadTransMaybeT()))
    }
    pub fn Control_Monad_Maybe_Trans_mapMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_mapMaybeT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_mapMaybeT.get_or_init(||
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
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT(),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                        &&&matchValue_1))
                                                                                                }
                                                                                        })))
    }
    pub fn Control_Monad_Maybe_Trans_functorMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_functorMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_functorMaybeT.get_or_init(||
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
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT(),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                               &&&dictFunctor),
                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                               &&&matchValue)),
                                                                                                                                                                                                                                         &&&matchValue_1))
                                                                                                                                                                 }
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))))
    }
    pub fn Control_Monad_Maybe_Trans_monadMaybeT_004022() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT_tco(dictMonad))
    }
    pub fn Control_Monad_Maybe_Trans_monadMaybeT_004022_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Maybe_Trans_monadMaybeT_004022_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadMaybeT_004022_002d1.get_or_init(||
                                                                           Lazy(Control_Monad_Maybe_Trans_monadMaybeT_004022.clone()))
    }
    pub fn Control_Monad_Maybe_Trans_bindMaybeT_004024() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_bindMaybeT_tco(dictMonad))
    }
    pub fn Control_Monad_Maybe_Trans_bindMaybeT_004024_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Maybe_Trans_bindMaybeT_004024_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_bindMaybeT_004024_002d1.get_or_init(||
                                                                          Lazy(Control_Monad_Maybe_Trans_bindMaybeT_004024.clone()))
    }
    pub fn Control_Monad_Maybe_Trans_applyMaybeT_004026() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_applyMaybeT_tco(dictMonad))
    }
    pub fn Control_Monad_Maybe_Trans_applyMaybeT_004026_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Maybe_Trans_applyMaybeT_004026_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_applyMaybeT_004026_002d1.get_or_init(||
                                                                           Lazy(Control_Monad_Maybe_Trans_applyMaybeT_004026.clone()))
    }
    pub fn Control_Monad_Maybe_Trans_applicativeMaybeT_004028() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_applicativeMaybeT_tco(dictMonad))
    }
    pub fn Control_Monad_Maybe_Trans_applicativeMaybeT_004028_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_Maybe_Trans_applicativeMaybeT_004028_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_applicativeMaybeT_004028_002d1.get_or_init(||
                                                                                 Lazy(Control_Monad_Maybe_Trans_applicativeMaybeT_004028.clone()))
    }
    pub fn Control_Monad_Maybe_Trans_monadMaybeT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                         &&&add(string("Applicative0"),
                                                &&Func1::new({
                                                                 let dictMonad
                                                                     =
                                                                     dictMonad.clone();
                                                                 move
                                                                     |usd__unused|
                                                                     PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_applicativeMaybeT_tco(&&dictMonad)
                                                             }),
                                                add(string("Bind1"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused_1|
                                                                         PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_bindMaybeT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Maybe_Trans_monadMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadMaybeT.get_or_init(||
                                                              Control_Monad_Maybe_Trans_monadMaybeT_004022_002d1.Value)
    }
    pub fn Control_Monad_Maybe_Trans_bindMaybeT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                         &&&add(string("bind"),
                                                &&Func1::new({
                                                                 let Applicative0
                                                                     =
                                                                     Applicative0.clone();
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
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                       &&&Bind1),
                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                   let matchValue_1
                                                                                                                                                                                       =
                                                                                                                                                                                       matchValue_1.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |v1|
                                                                                                                                                                                       {
                                                                                                                                                                                           let matchValue_3:
                                                                                                                                                                                                   LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                           match matchValue_3.as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_3_1_0)
                                                                                                                                                                                               =>
                                                                                                                                                                                               &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                           &&matchValue_3_1_0)),
                                                                                                                                                                                               _
                                                                                                                                                                                               =>
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                   &&&Applicative0),
                                                                                                                                                                                                                                &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                                                           }
                                                                                                                                                                                       }
                                                                                                                                                                               })))
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
                                                                         PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_applyMaybeT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Maybe_Trans_bindMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_bindMaybeT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_bindMaybeT.get_or_init(||
                                                             Control_Monad_Maybe_Trans_bindMaybeT_004024_002d1.Value)
    }
    pub fn Control_Monad_Maybe_Trans_applyMaybeT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let functorMaybeT1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_functorMaybeT(),
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
                                                                                  &&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT_tco(dictMonad)),
                                                add(string("Functor0"),
                                                    &&Func1::new({
                                                                     let functorMaybeT1
                                                                         =
                                                                         functorMaybeT1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &functorMaybeT1
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Maybe_Trans_applyMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_applyMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_applyMaybeT.get_or_init(||
                                                              Control_Monad_Maybe_Trans_applyMaybeT_004026_002d1.Value)
    }
    pub fn Control_Monad_Maybe_Trans_applicativeMaybeT_tco(dictMonad:
                                                               &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                         &&&add(string("pure"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT()),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                     &&&Func1::new(move
                                                                                                                                       |usd__arg1|
                                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone()))))),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_applyMaybeT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Maybe_Trans_applicativeMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_applicativeMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_applicativeMaybeT.get_or_init(||
                                                                    Control_Monad_Maybe_Trans_applicativeMaybeT_004028_002d1.Value)
    }
    pub fn Control_Monad_Maybe_Trans_semigroupMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_semigroupMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_semigroupMaybeT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonad|
                                                                                  {
                                                                                      let applyMaybeT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_applyMaybeT(),
                                                                                                                           dictMonad);
                                                                                      &Func1::new({
                                                                                                      let applyMaybeT1
                                                                                                          =
                                                                                                          applyMaybeT1.clone();
                                                                                                      move
                                                                                                          |dictSemigroup|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                                           &&&add(string("append"),
                                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                       &&&applyMaybeT1),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                       dictSemigroup)),
                                                                                                                                                  empty::<string,
                                                                                                                                                          &dyn Any>()))
                                                                                                  })
                                                                                  }))
    }
    pub fn Control_Monad_Maybe_Trans_monadAskMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadAskMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadAskMaybeT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictMonadAsk|
                                                                                 {
                                                                                     let monadMaybeT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                    Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadAskusd_Dict(),
                                                                                                                      &&&add(string("ask"),
                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadTransMaybeT()),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_ask(),
                                                                                                                                                                                                  dictMonadAsk)),
                                                                                                                             add(string("Monad0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let monadMaybeT1
                                                                                                                                                      =
                                                                                                                                                      monadMaybeT1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &monadMaybeT1
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Monad_Maybe_Trans_monadReaderMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadReaderMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadReaderMaybeT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonadReader|
                                                                                    {
                                                                                        let monadAskMaybeT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadAskMaybeT(),
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
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_mapMaybeT(),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_local(),
                                                                                                                                                                                                                                                            &&&dictMonadReader),
                                                                                                                                                                                                                         f))
                                                                                                                                             }),
                                                                                                                                add(string("MonadAsk0"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let monadAskMaybeT1
                                                                                                                                                         =
                                                                                                                                                         monadAskMaybeT1.clone();
                                                                                                                                                     move
                                                                                                                                                         |usd__unused|
                                                                                                                                                         &monadAskMaybeT1
                                                                                                                                                 }),
                                                                                                                                    empty::<string,
                                                                                                                                            &dyn Any>())))
                                                                                    }))
    }
    pub fn Control_Monad_Maybe_Trans_monadContMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadContMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadContMaybeT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonadCont|
                                                                                  {
                                                                                      let monadMaybeT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
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
                                                                                                                                                                                                                       &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT()),
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
                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT()),
                                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(a.clone()))))
                                                                                                                                                                                                                                                                                                                       })))
                                                                                                                                                                                                                                     })))
                                                                                                                                           }),
                                                                                                                              add(string("Monad0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let monadMaybeT1
                                                                                                                                                       =
                                                                                                                                                       monadMaybeT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &monadMaybeT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_Maybe_Trans_monadEffectMaybe() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadEffectMaybe:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadEffectMaybe.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadEffect|
                                                                                   {
                                                                                       let Monad0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictMonadEffect)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let monadMaybeT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
                                                                                                                            &&&Monad0);
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                                                        &&&add(string("liftEffect"),
                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_lift(),
                                                                                                                                                                                                                                       &&&Monad0)),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                    dictMonadEffect)),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadMaybeT1
                                                                                                                                                        =
                                                                                                                                                        monadMaybeT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadMaybeT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Maybe_Trans_monadRecMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadRecMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadRecMaybeT.get_or_init(||
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
                                                                                     let monadMaybeT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
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
                                                                                                                                                                                                                      &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT()),
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
                                                                                                                                                                                                                                                                                                                                                  LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(m_prime);
                                                                                                                                                                                                                                                                                                                                          if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                                 matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                              let activePatternResult:
                                                                                                                                                                                                                                                                                                                                                      LrcPtr<Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                                                         Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
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
                                                                                                                                                                                                                                                                                                                                                                                             Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
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
                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match activePatternResult_1.as_ref()
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
                                                                                                                                                                                                                                                                                                                                                             string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                                                                                                                                                                                                                                                                                                                                                  }
                                                                                                                                                                                                                                                                                                                                              }
                                                                                                                                                                                                                                                                                                                                          } else {
                                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
                                                                                                                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                                                                                                                      })))
                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                    })))
                                                                                                                                          }),
                                                                                                                             add(string("Monad0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let monadMaybeT1
                                                                                                                                                      =
                                                                                                                                                      monadMaybeT1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &monadMaybeT1
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Monad_Maybe_Trans_monadStateMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadStateMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadStateMaybeT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadState|
                                                                                   {
                                                                                       let Monad0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictMonadState)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let monadMaybeT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
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
                                                                                                                                                                                                                                                           &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadTransMaybeT()),
                                                                                                                                                                                                                        &&&Monad0),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                                                                                                                           &&&dictMonadState),
                                                                                                                                                                                                                        f))
                                                                                                                                            }),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadMaybeT1
                                                                                                                                                        =
                                                                                                                                                        monadMaybeT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadMaybeT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Maybe_Trans_monadTellMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadTellMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadTellMaybeT.get_or_init(||
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
                                                                                      let monadMaybeT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
                                                                                                                           &&&Monad1);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadTellusd_Dict(),
                                                                                                                       &&&add(string("tell"),
                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_lift(),
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
                                                                                                                                                       let monadMaybeT1
                                                                                                                                                           =
                                                                                                                                                           monadMaybeT1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused_1|
                                                                                                                                                           &monadMaybeT1
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>()))))
                                                                                  }))
    }
    pub fn Control_Monad_Maybe_Trans_monadWriterMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadWriterMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadWriterMaybeT.get_or_init(||
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
                                                                                        let monadTellMaybeT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadTellMaybeT(),
                                                                                                                             &&&MonadTell1);
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_MonadWriterusd_Dict(),
                                                                                                                         &&&add(string("listen"),
                                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_mapMaybeT(),
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
                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
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
                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_mapMaybeT(),
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
                                                                                                                                                                                                                                                                                                                                 LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(a_1);
                                                                                                                                                                                                                                                                                                                         match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                             Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                 let activePatternResult:
                                                                                                                                                                                                                                                                                                                                         LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::_007cUnbox_007c(matchValue_1_1_0);
                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match activePatternResult.as_ref()
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
                                                                                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
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
                                                                                                                                                             let monadTellMaybeT1
                                                                                                                                                                 =
                                                                                                                                                                 monadTellMaybeT1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |usd__unused_1|
                                                                                                                                                                 &monadTellMaybeT1
                                                                                                                                                         }),
                                                                                                                                            empty::<string,
                                                                                                                                                    &dyn Any>())))))
                                                                                    }))
    }
    pub fn Control_Monad_Maybe_Trans_monadThrowMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadThrowMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadThrowMaybeT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadThrow|
                                                                                   {
                                                                                       let Monad0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictMonadThrow)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let monadMaybeT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
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
                                                                                                                                                                                                                                                           &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadTransMaybeT()),
                                                                                                                                                                                                                        &&&Monad0),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                                                                                                                           &&&dictMonadThrow),
                                                                                                                                                                                                                        e))
                                                                                                                                            }),
                                                                                                                               add(string("Monad0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadMaybeT1
                                                                                                                                                        =
                                                                                                                                                        monadMaybeT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadMaybeT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Maybe_Trans_monadErrorMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadErrorMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadErrorMaybeT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictMonadError|
                                                                                   {
                                                                                       let monadThrowMaybeT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadThrowMaybeT(),
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
                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT()),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_catchError(),
                                                                                                                                                                                                                                                                                                                      &&&dictMonadError),
                                                                                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                                                  let matchValue_1
                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                      matchValue_1.clone();
                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                      |a|
                                                                                                                                                                                                                                                                      &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                  a))
                                                                                                                                                                                                                                                              })))
                                                                                                                                                                        }
                                                                                                                                                                })
                                                                                                                                            }),
                                                                                                                               add(string("MonadThrow0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let monadThrowMaybeT1
                                                                                                                                                        =
                                                                                                                                                        monadThrowMaybeT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &monadThrowMaybeT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Maybe_Trans_monadSTMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadSTMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadSTMaybeT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonadST|
                                                                                {
                                                                                    let Monad0 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                Sharpurs_Prelude::unbox(dictMonadST)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    let monadMaybeT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
                                                                                                                         &&&Monad0);
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                                                     &&&add(string("liftST"),
                                                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_lift(),
                                                                                                                                                                                                                                    &&&Monad0)),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_liftST(),
                                                                                                                                                                                                 dictMonadST)),
                                                                                                                            add(string("Monad0"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let monadMaybeT1
                                                                                                                                                     =
                                                                                                                                                     monadMaybeT1.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused|
                                                                                                                                                     &monadMaybeT1
                                                                                                                                             }),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>())))
                                                                                }))
    }
    pub fn Control_Monad_Maybe_Trans_monoidMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monoidMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monoidMaybeT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictMonad|
                                                                               {
                                                                                   let applicativeMaybeT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_applicativeMaybeT(),
                                                                                                                        dictMonad);
                                                                                   let semigroupMaybeT1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_semigroupMaybeT(),
                                                                                                                        dictMonad);
                                                                                   &Func1::new({
                                                                                                   let applicativeMaybeT1
                                                                                                       =
                                                                                                       applicativeMaybeT1.clone();
                                                                                                   let semigroupMaybeT1
                                                                                                       =
                                                                                                       semigroupMaybeT1.clone();
                                                                                                   move
                                                                                                       |dictMonoid|
                                                                                                       {
                                                                                                           let semigroupMaybeT2 =
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&semigroupMaybeT1,
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                                            &&&add(string("mempty"),
                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                        &&&applicativeMaybeT1),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                        dictMonoid)),
                                                                                                                                                   add(string("Semigroup0"),
                                                                                                                                                       &&Func1::new({
                                                                                                                                                                        let semigroupMaybeT2
                                                                                                                                                                            =
                                                                                                                                                                            semigroupMaybeT2.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |usd__unused|
                                                                                                                                                                            &semigroupMaybeT2
                                                                                                                                                                    }),
                                                                                                                                                       empty::<string,
                                                                                                                                                               &dyn Any>())))
                                                                                                       }
                                                                                               })
                                                                               }))
    }
    pub fn Control_Monad_Maybe_Trans_altMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_altMaybeT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_altMaybeT.get_or_init(||
                                                            &Func1::new(move
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
                                                                                let functorMaybeT1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_functorMaybeT(),
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
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT(),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                               &&&Bind1),
                                                                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                                                                                           let matchValue_1
                                                                                                                                                                                                                                                               =
                                                                                                                                                                                                                                                               matchValue_1.clone();
                                                                                                                                                                                                                                                           move
                                                                                                                                                                                                                                                               |m|
                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                   let matchValue_3:
                                                                                                                                                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(m);
                                                                                                                                                                                                                                                                   if let Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                                                          matchValue_3.as_ref()
                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                       &matchValue_1
                                                                                                                                                                                                                                                                   } else {
                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                           &&&Applicative0),
                                                                                                                                                                                                                                                                                                        &&&matchValue_3)
                                                                                                                                                                                                                                                                   }
                                                                                                                                                                                                                                                               }
                                                                                                                                                                                                                                                       })))
                                                                                                                                                                 }
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        add(string("Functor0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let functorMaybeT1
                                                                                                                                                 =
                                                                                                                                                 functorMaybeT1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &functorMaybeT1
                                                                                                                                         }),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>())))
                                                                            }))
    }
    pub fn Control_Monad_Maybe_Trans_plusMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_plusMaybeT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_plusMaybeT.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictMonad|
                                                                             {
                                                                                 let altMaybeT1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_altMaybeT(),
                                                                                                                      dictMonad);
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                                                  &&&add(string("empty"),
                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_MaybeT(),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                                                              &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))),
                                                                                                                         add(string("Alt0"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let altMaybeT1
                                                                                                                                                  =
                                                                                                                                                  altMaybeT1.clone();
                                                                                                                                              move
                                                                                                                                                  |usd__unused|
                                                                                                                                                  &altMaybeT1
                                                                                                                                          }),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>())))
                                                                             }))
    }
    pub fn Control_Monad_Maybe_Trans_alternativeMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_alternativeMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_alternativeMaybeT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonad|
                                                                                    {
                                                                                        let applicativeMaybeT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_applicativeMaybeT(),
                                                                                                                             dictMonad);
                                                                                        let plusMaybeT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_plusMaybeT(),
                                                                                                                             dictMonad);
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                                                         &&&add(string("Applicative0"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let applicativeMaybeT1
                                                                                                                                                     =
                                                                                                                                                     applicativeMaybeT1.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused|
                                                                                                                                                     &applicativeMaybeT1
                                                                                                                                             }),
                                                                                                                                add(string("Plus1"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let plusMaybeT1
                                                                                                                                                         =
                                                                                                                                                         plusMaybeT1.clone();
                                                                                                                                                     move
                                                                                                                                                         |usd__unused_1|
                                                                                                                                                         &plusMaybeT1
                                                                                                                                                 }),
                                                                                                                                    empty::<string,
                                                                                                                                            &dyn Any>())))
                                                                                    }))
    }
    pub fn Control_Monad_Maybe_Trans_monadPlusMaybeT() -> &dyn Any {
        static Control_Monad_Maybe_Trans_monadPlusMaybeT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Maybe_Trans_monadPlusMaybeT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonad|
                                                                                  {
                                                                                      let monadMaybeT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadMaybeT(),
                                                                                                                           dictMonad);
                                                                                      let alternativeMaybeT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_alternativeMaybeT(),
                                                                                                                           dictMonad);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                                                       &&&add(string("Monad0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let monadMaybeT1
                                                                                                                                                   =
                                                                                                                                                   monadMaybeT1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &monadMaybeT1
                                                                                                                                           }),
                                                                                                                              add(string("Alternative1"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let alternativeMaybeT1
                                                                                                                                                       =
                                                                                                                                                       alternativeMaybeT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused_1|
                                                                                                                                                       &alternativeMaybeT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
}
