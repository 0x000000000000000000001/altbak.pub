pub mod PureScript_Effect_Aff_Class {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_dbc0dd87::PureScript_Control_Monad_Cont_Trans;
    use crate::module_8ea4b64e::PureScript_Control_Monad_Except_Trans;
    use crate::module_34515613::PureScript_Control_Monad_List_Trans;
    use crate::module_9b44f963::PureScript_Control_Monad_Maybe_Trans;
    use crate::module_9e2ffde7::PureScript_Control_Monad_RWS_Trans;
    use crate::module_8ea61114::PureScript_Control_Monad_Reader_Trans;
    use crate::module_3bc71566::PureScript_Control_Monad_State_Trans;
    use crate::module_f5fe307f::PureScript_Control_Monad_Trans_Class;
    use crate::module_2b2441be::PureScript_Control_Monad_Writer_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_361110b3::PureScript_Effect_Aff;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Effect_Aff_Class_lift() -> &dyn Any {
        static Effect_Aff_Class_lift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_lift.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                               &&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadTransContT()))
    }
    pub fn Effect_Aff_Class_lift1() -> &dyn Any {
        static Effect_Aff_Class_lift1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_lift1.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadTransExceptT()))
    }
    pub fn Effect_Aff_Class_lift2() -> &dyn Any {
        static Effect_Aff_Class_lift2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_lift2.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_monadTransListT()))
    }
    pub fn Effect_Aff_Class_lift3() -> &dyn Any {
        static Effect_Aff_Class_lift3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_lift3.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                &&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadTransMaybeT()))
    }
    pub fn Effect_Aff_Class_lift4() -> &dyn Any {
        static Effect_Aff_Class_lift4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_lift4.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                &&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadTransReaderT()))
    }
    pub fn Effect_Aff_Class_lift5() -> &dyn Any {
        static Effect_Aff_Class_lift5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_lift5.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                &&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadTransStateT()))
    }
    pub fn Effect_Aff_Class_MonadAffusd_Dict() -> &dyn Any {
        static Effect_Aff_Class_MonadAffusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_MonadAffusd_Dict.get_or_init(||
                                                          &Func1::new(move |x|
                                                                          x.clone()))
    }
    pub fn Effect_Aff_Class_monadAffAff() -> &dyn Any {
        static Effect_Aff_Class_monadAffAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffAff.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                      &&&add(string("liftAff"),
                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                               &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                             add(string("MonadEffect0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Effect_Aff::Effect_Aff_monadEffectAff()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Effect_Aff_Class_liftAff() -> &dyn Any {
        static Effect_Aff_Class_liftAff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_liftAff.get_or_init(||
                                                 &Func1::new(move |dict|
                                                                 find(string("liftAff"),
                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Effect_Aff_Class_monadAffContT() -> &dyn Any {
        static Effect_Aff_Class_monadAffContT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffContT.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadAff|
                                                                       {
                                                                           let MonadEffect0 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadEffect0"),
                                                                                                                       Sharpurs_Prelude::unbox(dictMonadAff)),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                           let monadEffectContT =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_monadEffectContT(),
                                                                                                                &&&MonadEffect0);
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                                            &&&add(string("liftAff"),
                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_lift(),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&MonadEffect0)),
                                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_liftAff(),
                                                                                                                                                                                        dictMonadAff)),
                                                                                                                   add(string("MonadEffect0"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let monadEffectContT
                                                                                                                                            =
                                                                                                                                            monadEffectContT.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused|
                                                                                                                                            &monadEffectContT
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>())))
                                                                       }))
    }
    pub fn Effect_Aff_Class_monadAffExceptT() -> &dyn Any {
        static Effect_Aff_Class_monadAffExceptT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffExceptT.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonadAff|
                                                                         {
                                                                             let MonadEffect0 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadEffect0"),
                                                                                                                         Sharpurs_Prelude::unbox(dictMonadAff)),
                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                                             let monadEffectExceptT =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_monadEffectExceptT(),
                                                                                                                  &&&MonadEffect0);
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                                              &&&add(string("liftAff"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_lift1(),
                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&MonadEffect0)),
                                                                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_liftAff(),
                                                                                                                                                                                          dictMonadAff)),
                                                                                                                     add(string("MonadEffect0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let monadEffectExceptT
                                                                                                                                              =
                                                                                                                                              monadEffectExceptT.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &monadEffectExceptT
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Aff_Class_monadAffListT() -> &dyn Any {
        static Effect_Aff_Class_monadAffListT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffListT.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadAff|
                                                                       {
                                                                           let MonadEffect0 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadEffect0"),
                                                                                                                       Sharpurs_Prelude::unbox(dictMonadAff)),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                           let monadEffectListT =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_monadEffectListT(),
                                                                                                                &&&MonadEffect0);
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                                            &&&add(string("liftAff"),
                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_lift2(),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&MonadEffect0)),
                                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_liftAff(),
                                                                                                                                                                                        dictMonadAff)),
                                                                                                                   add(string("MonadEffect0"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let monadEffectListT
                                                                                                                                            =
                                                                                                                                            monadEffectListT.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused|
                                                                                                                                            &monadEffectListT
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>())))
                                                                       }))
    }
    pub fn Effect_Aff_Class_monadAffMaybe() -> &dyn Any {
        static Effect_Aff_Class_monadAffMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffMaybe.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadAff|
                                                                       {
                                                                           let MonadEffect0 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadEffect0"),
                                                                                                                       Sharpurs_Prelude::unbox(dictMonadAff)),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                           let monadEffectMaybe =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Maybe_Trans::Control_Monad_Maybe_Trans_monadEffectMaybe(),
                                                                                                                &&&MonadEffect0);
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                                            &&&add(string("liftAff"),
                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_lift3(),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&MonadEffect0)),
                                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_liftAff(),
                                                                                                                                                                                        dictMonadAff)),
                                                                                                                   add(string("MonadEffect0"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let monadEffectMaybe
                                                                                                                                            =
                                                                                                                                            monadEffectMaybe.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused|
                                                                                                                                            &monadEffectMaybe
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>())))
                                                                       }))
    }
    pub fn Effect_Aff_Class_monadAffRWS() -> &dyn Any {
        static Effect_Aff_Class_monadAffRWS: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffRWS.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictMonadAff|
                                                                     {
                                                                         let MonadEffect0 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadEffect0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictMonadAff)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         let Monad0 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                     Sharpurs_Prelude::unbox(&&MonadEffect0)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         let liftAff1 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_liftAff(),
                                                                                                              dictMonadAff);
                                                                         &Func1::new({
                                                                                         let Monad0
                                                                                             =
                                                                                             Monad0.clone();
                                                                                         let MonadEffect0
                                                                                             =
                                                                                             MonadEffect0.clone();
                                                                                         let liftAff1
                                                                                             =
                                                                                             liftAff1.clone();
                                                                                         move
                                                                                             |dictMonoid|
                                                                                             {
                                                                                                 let monadEffectRWS =
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_RWS_Trans::Control_Monad_RWS_Trans_monadEffectRWS(),
                                                                                                                                                                         dictMonoid),
                                                                                                                                      &&&MonadEffect0);
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                                                                  &&&add(string("liftAff"),
                                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_RWS_Trans::Control_Monad_RWS_Trans_monadTransRWST(),
                                                                                                                                                                                                                                                                                                                       dictMonoid)),
                                                                                                                                                                                                                                                 &&&Monad0)),
                                                                                                                                                                           &&&liftAff1),
                                                                                                                                         add(string("MonadEffect0"),
                                                                                                                                             &&Func1::new({
                                                                                                                                                              let monadEffectRWS
                                                                                                                                                                  =
                                                                                                                                                                  monadEffectRWS.clone();
                                                                                                                                                              move
                                                                                                                                                                  |usd__unused|
                                                                                                                                                                  &monadEffectRWS
                                                                                                                                                          }),
                                                                                                                                             empty::<string,
                                                                                                                                                     &dyn Any>())))
                                                                                             }
                                                                                     })
                                                                     }))
    }
    pub fn Effect_Aff_Class_monadAffReader() -> &dyn Any {
        static Effect_Aff_Class_monadAffReader: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffReader.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictMonadAff|
                                                                        {
                                                                            let MonadEffect0 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadEffect0"),
                                                                                                                        Sharpurs_Prelude::unbox(dictMonadAff)),
                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                            let monadEffectReader =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_monadEffectReader(),
                                                                                                                 &&&MonadEffect0);
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                                             &&&add(string("liftAff"),
                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_lift4(),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&MonadEffect0)),
                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_liftAff(),
                                                                                                                                                                                         dictMonadAff)),
                                                                                                                    add(string("MonadEffect0"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let monadEffectReader
                                                                                                                                             =
                                                                                                                                             monadEffectReader.clone();
                                                                                                                                         move
                                                                                                                                             |usd__unused|
                                                                                                                                             &monadEffectReader
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>())))
                                                                        }))
    }
    pub fn Effect_Aff_Class_monadAffState() -> &dyn Any {
        static Effect_Aff_Class_monadAffState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffState.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadAff|
                                                                       {
                                                                           let MonadEffect0 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadEffect0"),
                                                                                                                       Sharpurs_Prelude::unbox(dictMonadAff)),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                           let monadEffectState =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_monadEffectState(),
                                                                                                                &&&MonadEffect0);
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                                            &&&add(string("liftAff"),
                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_lift5(),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&MonadEffect0)),
                                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_liftAff(),
                                                                                                                                                                                        dictMonadAff)),
                                                                                                                   add(string("MonadEffect0"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let monadEffectState
                                                                                                                                            =
                                                                                                                                            monadEffectState.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused|
                                                                                                                                            &monadEffectState
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>())))
                                                                       }))
    }
    pub fn Effect_Aff_Class_monadAffWriter() -> &dyn Any {
        static Effect_Aff_Class_monadAffWriter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Aff_Class_monadAffWriter.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictMonadAff|
                                                                        {
                                                                            let MonadEffect0 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadEffect0"),
                                                                                                                        Sharpurs_Prelude::unbox(dictMonadAff)),
                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                            let Monad0 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                        Sharpurs_Prelude::unbox(&&MonadEffect0)),
                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                            let liftAff1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_liftAff(),
                                                                                                                 dictMonadAff);
                                                                            &Func1::new({
                                                                                            let Monad0
                                                                                                =
                                                                                                Monad0.clone();
                                                                                            let MonadEffect0
                                                                                                =
                                                                                                MonadEffect0.clone();
                                                                                            let liftAff1
                                                                                                =
                                                                                                liftAff1.clone();
                                                                                            move
                                                                                                |dictMonoid|
                                                                                                {
                                                                                                    let monadEffectWriter =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadEffectWriter(),
                                                                                                                                                                            dictMonoid),
                                                                                                                                         &&&MonadEffect0);
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Aff_Class::Effect_Aff_Class_MonadAffusd_Dict(),
                                                                                                                                     &&&add(string("liftAff"),
                                                                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_monadTransWriterT(),
                                                                                                                                                                                                                                                                                                                          dictMonoid)),
                                                                                                                                                                                                                                                    &&&Monad0)),
                                                                                                                                                                              &&&liftAff1),
                                                                                                                                            add(string("MonadEffect0"),
                                                                                                                                                &&Func1::new({
                                                                                                                                                                 let monadEffectWriter
                                                                                                                                                                     =
                                                                                                                                                                     monadEffectWriter.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |usd__unused|
                                                                                                                                                                     &monadEffectWriter
                                                                                                                                                             }),
                                                                                                                                                empty::<string,
                                                                                                                                                        &dyn Any>())))
                                                                                                }
                                                                                        })
                                                                        }))
    }
}
