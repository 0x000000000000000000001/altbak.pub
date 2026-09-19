pub mod PureScript_Effect {
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
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub mod Effect_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::count;
        pub fn pureE() -> &dyn Any {
            static pureE: MutCell<Option<&dyn Any>> = MutCell::new(None);
            pureE.get_or_init(||
                                  &Func1::new(move |a|
                                                  &Func1::new({
                                                                  let a =
                                                                      a.clone();
                                                                  move |_arg|
                                                                      a
                                                              })))
        }
        pub fn bindE() -> &dyn Any {
            static bindE: MutCell<Option<&dyn Any>> = MutCell::new(None);
            bindE.get_or_init(||
                                  &Func1::new(move |a|
                                                  &Func1::new({
                                                                  let a =
                                                                      a.clone();
                                                                  move |f|
                                                                      &Func1::new({
                                                                                      let f
                                                                                          =
                                                                                          f.clone();
                                                                                      move
                                                                                          |_arg|
                                                                                          f(a(defaultOf()))(defaultOf())
                                                                                  })
                                                              })))
        }
        pub fn untilE() -> &dyn Any {
            static untilE: MutCell<Option<&dyn Any>> = MutCell::new(None);
            untilE.get_or_init(||
                                   &Func1::new(move |f|
                                                   &Func1::new({
                                                                   let f =
                                                                       f.clone();
                                                                   move |_arg|
                                                                       {
                                                                           let condition:
                                                                                   MutCell<bool> =
                                                                               MutCell::new(false);
                                                                           while !condition.get()
                                                                                 {
                                                                               condition.set(Sharpurs_Prelude::unbox(&f(defaultOf())));
                                                                           }
                                                                           defaultOf()
                                                                       }
                                                               })))
        }
        pub fn whileE() -> &dyn Any {
            static whileE: MutCell<Option<&dyn Any>> = MutCell::new(None);
            whileE.get_or_init(||
                                   &Func1::new(move |f|
                                                   &Func1::new({
                                                                   let f =
                                                                       f.clone();
                                                                   move |a|
                                                                       &Func1::new({
                                                                                       let a
                                                                                           =
                                                                                           a.clone();
                                                                                       move
                                                                                           |_arg|
                                                                                           {
                                                                                               let f_ =
                                                                                                   f;
                                                                                               let condition:
                                                                                                       MutCell<bool> =
                                                                                                   MutCell::new(Sharpurs_Prelude::unbox(&f_(defaultOf())));
                                                                                               while condition.get()
                                                                                                     {
                                                                                                   {
                                                                                                       let value =
                                                                                                           a(defaultOf());
                                                                                                       ()
                                                                                                   }
                                                                                                   condition.set(Sharpurs_Prelude::unbox(&f_(defaultOf())))
                                                                                               }
                                                                                               defaultOf()
                                                                                           }
                                                                                   })
                                                               })))
        }
        pub fn forE() -> &dyn Any {
            static forE: MutCell<Option<&dyn Any>> = MutCell::new(None);
            forE.get_or_init(||
                                 &Func1::new(move |lo|
                                                 &Func1::new({
                                                                 let lo =
                                                                     lo.clone();
                                                                 move |hi|
                                                                     &Func1::new({
                                                                                     let hi
                                                                                         =
                                                                                         hi.clone();
                                                                                     move
                                                                                         |f|
                                                                                         &Func1::new({
                                                                                                         let f
                                                                                                             =
                                                                                                             f.clone();
                                                                                                         move
                                                                                                             |_arg|
                                                                                                             {
                                                                                                                 let l:
                                                                                                                         i32 =
                                                                                                                     Sharpurs_Prelude::unbox(&lo);
                                                                                                                 let h:
                                                                                                                         i32 =
                                                                                                                     Sharpurs_Prelude::unbox(&hi);
                                                                                                                 for i
                                                                                                                     in
                                                                                                                     l..=h
                                                                                                                             -
                                                                                                                             1_i32
                                                                                                                     {
                                                                                                                     let value =
                                                                                                                         f(&i)(defaultOf());
                                                                                                                     ()
                                                                                                                 }
                                                                                                                 defaultOf()
                                                                                                             }
                                                                                                     })
                                                                                 })
                                                             })))
        }
        pub fn foreachE() -> &dyn Any {
            static foreachE: MutCell<Option<&dyn Any>> = MutCell::new(None);
            foreachE.get_or_init(||
                                     &Func1::new(move |arr|
                                                     &Func1::new({
                                                                     let arr =
                                                                         arr.clone();
                                                                     move |f|
                                                                         &Func1::new({
                                                                                         let f
                                                                                             =
                                                                                             f.clone();
                                                                                         move
                                                                                             |_arg|
                                                                                             {
                                                                                                 let arr_ =
                                                                                                     Sharpurs_Prelude::unbox(&arr);
                                                                                                 for idx
                                                                                                     in
                                                                                                     0_i32..=count(arr_.clone())
                                                                                                                 -
                                                                                                                 1_i32
                                                                                                     {
                                                                                                     let value =
                                                                                                         f(arr_[idx].clone())(defaultOf());
                                                                                                     ()
                                                                                                 }
                                                                                                 defaultOf()
                                                                                             }
                                                                                     })
                                                                 })))
        }
    }
    pub fn Effect_bindE() -> &dyn Any {
        static Effect_bindE: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Effect_bindE.get_or_init(|| &PureScript_Effect::Effect_FFI::bindE())
    }
    pub fn Effect_forE() -> &dyn Any {
        static Effect_forE: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Effect_forE.get_or_init(|| &PureScript_Effect::Effect_FFI::forE())
    }
    pub fn Effect_foreachE() -> &dyn Any {
        static Effect_foreachE: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_foreachE.get_or_init(||
                                        &PureScript_Effect::Effect_FFI::foreachE())
    }
    pub fn Effect_pureE() -> &dyn Any {
        static Effect_pureE: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Effect_pureE.get_or_init(|| &PureScript_Effect::Effect_FFI::pureE())
    }
    pub fn Effect_untilE() -> &dyn Any {
        static Effect_untilE: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Effect_untilE.get_or_init(|| &PureScript_Effect::Effect_FFI::untilE())
    }
    pub fn Effect_whileE() -> &dyn Any {
        static Effect_whileE: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Effect_whileE.get_or_init(|| &PureScript_Effect::Effect_FFI::whileE())
    }
    pub fn Effect_monadEffect_004063() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                         &&&add(string("Applicative0"),
                                                &&Func1::new({
                                                                 let Effect_applicativeEffect_004066_002d1
                                                                     =
                                                                     Effect_applicativeEffect_004066_002d1.clone();
                                                                 move
                                                                     |usd__unused|
                                                                     &Effect_applicativeEffect_004066_002d1.Value
                                                             }),
                                                add(string("Bind1"),
                                                    &&Func1::new({
                                                                     let Effect_bindEffect_004064_002d1
                                                                         =
                                                                         Effect_bindEffect_004064_002d1.clone();
                                                                     move
                                                                         |usd__unused_1|
                                                                         &Effect_bindEffect_004064_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Effect_monadEffect_004063_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_monadEffect_004063_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_monadEffect_004063_002d1.get_or_init(||
                                                        Lazy(Effect_monadEffect_004063.clone()))
    }
    pub fn Effect_bindEffect_004064() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                         &&&add(string("bind"),
                                                &&PureScript_Effect::Effect_bindE(),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let Effect_applyEffect_004065_002d1
                                                                         =
                                                                         Effect_applyEffect_004065_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Effect_applyEffect_004065_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Effect_bindEffect_004064_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_bindEffect_004064_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_bindEffect_004064_002d1.get_or_init(||
                                                       Lazy(Effect_bindEffect_004064.clone()))
    }
    pub fn Effect_applyEffect_004065() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                         &&&add(string("apply"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_ap(),
                                                                                  &&&Effect_monadEffect_004063_002d1.Value),
                                                add(string("Functor0"),
                                                    &&Func1::new({
                                                                     let Effect_functorEffect_004067_002d1
                                                                         =
                                                                         Effect_functorEffect_004067_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Effect_functorEffect_004067_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Effect_applyEffect_004065_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_applyEffect_004065_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_applyEffect_004065_002d1.get_or_init(||
                                                        Lazy(Effect_applyEffect_004065.clone()))
    }
    pub fn Effect_applicativeEffect_004066() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                         &&&add(string("pure"),
                                                &&PureScript_Effect::Effect_pureE(),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let Effect_applyEffect_004065_002d1
                                                                         =
                                                                         Effect_applyEffect_004065_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Effect_applyEffect_004065_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Effect_applicativeEffect_004066_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_applicativeEffect_004066_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_applicativeEffect_004066_002d1.get_or_init(||
                                                              Lazy(Effect_applicativeEffect_004066.clone()))
    }
    pub fn Effect_functorEffect_004067() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                         &&&add(string("map"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_liftA1(),
                                                                                  &&&Effect_applicativeEffect_004066_002d1.Value),
                                                empty::<string, &dyn Any>()))
    }
    pub fn Effect_functorEffect_004067_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Effect_functorEffect_004067_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Effect_functorEffect_004067_002d1.get_or_init(||
                                                          Lazy(Effect_functorEffect_004067.clone()))
    }
    pub fn Effect_monadEffect() -> &dyn Any {
        static Effect_monadEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_monadEffect.get_or_init(||
                                           Effect_monadEffect_004063_002d1.Value)
    }
    pub fn Effect_bindEffect() -> &dyn Any {
        static Effect_bindEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_bindEffect.get_or_init(|| Effect_bindEffect_004064_002d1.Value)
    }
    pub fn Effect_applyEffect() -> &dyn Any {
        static Effect_applyEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_applyEffect.get_or_init(||
                                           Effect_applyEffect_004065_002d1.Value)
    }
    pub fn Effect_applicativeEffect() -> &dyn Any {
        static Effect_applicativeEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_applicativeEffect.get_or_init(||
                                                 Effect_applicativeEffect_004066_002d1.Value)
    }
    pub fn Effect_functorEffect() -> &dyn Any {
        static Effect_functorEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_functorEffect.get_or_init(||
                                             Effect_functorEffect_004067_002d1.Value)
    }
    pub fn Effect_semigroupEffect() -> &dyn Any {
        static Effect_semigroupEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_semigroupEffect.get_or_init(||
                                               &Func1::new(move
                                                               |dictSemigroup|
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                &&&add(string("append"),
                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                            &&&PureScript_Effect::Effect_applyEffect()),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                            dictSemigroup)),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Effect_monoidEffect() -> &dyn Any {
        static Effect_monoidEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_monoidEffect.get_or_init(||
                                            &Func1::new(move |dictMonoid|
                                                            {
                                                                let semigroupEffect1 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                               Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                 &&&add(string("mempty"),
                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_pureE(),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                             dictMonoid)),
                                                                                                        add(string("Semigroup0"),
                                                                                                            &&Func1::new({
                                                                                                                             let semigroupEffect1
                                                                                                                                 =
                                                                                                                                 semigroupEffect1.clone();
                                                                                                                             move
                                                                                                                                 |usd__unused|
                                                                                                                                 &semigroupEffect1
                                                                                                                         }),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>())))
                                                            }))
    }
}
