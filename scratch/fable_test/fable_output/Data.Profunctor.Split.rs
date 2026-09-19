pub mod PureScript_Data_Profunctor_Split {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c353b055::PureScript_Data_Exists;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug,)]
    pub enum Data_Profunctor_Split_SplitF {
        Data_Profunctor_Split_SplitFusd_Ctor(&dyn Any, &dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Profunctor_Split::Data_Profunctor_Split_SplitF {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Profunctor_Split_identity() -> &dyn Any {
        static Data_Profunctor_Split_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_identity.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                        &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Profunctor_Split_SplitF() -> &dyn Any {
        static Data_Profunctor_Split_SplitF: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_SplitF.get_or_init(||
                                                     &Func1::new(move
                                                                     |usd__arg1|
                                                                     Func1::new({
                                                                                    let usd__arg1
                                                                                        =
                                                                                        usd__arg1.clone();
                                                                                    move
                                                                                        |usd__arg2|
                                                                                        Func1::new({
                                                                                                       let usd__arg2
                                                                                                           =
                                                                                                           usd__arg2.clone();
                                                                                                       move
                                                                                                           |usd__arg3|
                                                                                                           &LrcPtr::new(PureScript_Data_Profunctor_Split::Data_Profunctor_Split_SplitF::Data_Profunctor_Split_SplitFusd_Ctor(usd__arg1,
                                                                                                                                                                                                                             usd__arg2,
                                                                                                                                                                                                                             usd__arg3.clone()))
                                                                                                   })
                                                                                })))
    }
    pub fn Data_Profunctor_Split_Split() -> &dyn Any {
        static Data_Profunctor_Split_Split: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_Split.get_or_init(||
                                                    &Func1::new(move |x|
                                                                    x.clone()))
    }
    pub fn Data_Profunctor_Split_unSplit() -> &dyn Any {
        static Data_Profunctor_Split_unSplit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_unSplit.get_or_init(||
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
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Exists::Data_Exists_runExists(),
                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                    |v1|
                                                                                                                                                                                    {
                                                                                                                                                                                        let matchValue_3:
                                                                                                                                                                                                LrcPtr<PureScript_Data_Profunctor_Split::Data_Profunctor_Split_SplitF> =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                               &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                      PureScript_Data_Profunctor_Split::Data_Profunctor_Split_SplitF::Data_Profunctor_Split_SplitFusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                           _,
                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                                            &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                   PureScript_Data_Profunctor_Split::Data_Profunctor_Split_SplitF::Data_Profunctor_Split_SplitFusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                        x,
                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                         &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                PureScript_Data_Profunctor_Split::Data_Profunctor_Split_SplitF::Data_Profunctor_Split_SplitFusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                     _,
                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                            })
                                                                                                                                                                                    })),
                                                                                                                               &&&matchValue_1)
                                                                                          }
                                                                                  })))
    }
    pub fn Data_Profunctor_Split_split() -> &dyn Any {
        static Data_Profunctor_Split_split: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_split.get_or_init(||
                                                    &Func1::new(move |f|
                                                                    &Func1::new({
                                                                                    let f
                                                                                        =
                                                                                        f.clone();
                                                                                    move
                                                                                        |g|
                                                                                        &Func1::new({
                                                                                                        let g
                                                                                                            =
                                                                                                            g.clone();
                                                                                                        move
                                                                                                            |fx|
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_Split(),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Exists::Data_Exists_mkExists(),
                                                                                                                                                                                &&&LrcPtr::new(PureScript_Data_Profunctor_Split::Data_Profunctor_Split_SplitF::Data_Profunctor_Split_SplitFusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                    &g,
                                                                                                                                                                                                                                                                                                    fx.clone()))))
                                                                                                    })
                                                                                })))
    }
    pub fn Data_Profunctor_Split_profunctorSplit() -> &dyn Any {
        static Data_Profunctor_Split_profunctorSplit:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_profunctorSplit.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_Profunctorusd_Dict(),
                                                                                               &&&add(string("dimap"),
                                                                                                      &&Func1::new(move
                                                                                                                       |f|
                                                                                                                       &Func1::new({
                                                                                                                                       let f
                                                                                                                                           =
                                                                                                                                           f.clone();
                                                                                                                                       move
                                                                                                                                           |g|
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_unSplit(),
                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                              let g
                                                                                                                                                                                                  =
                                                                                                                                                                                                  g.clone();
                                                                                                                                                                                              move
                                                                                                                                                                                                  |h|
                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                  let h
                                                                                                                                                                                                                      =
                                                                                                                                                                                                                      h.clone();
                                                                                                                                                                                                                  move
                                                                                                                                                                                                                      |i|
                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_split(),
                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                &&&h),
                                                                                                                                                                                                                                                                                                                             &&&f)),
                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                             &&&g),
                                                                                                                                                                                                                                                                                          i))
                                                                                                                                                                                                              })
                                                                                                                                                                                          }))
                                                                                                                                   })),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))
    }
    pub fn Data_Profunctor_Split_lowerSplit() -> &dyn Any {
        static Data_Profunctor_Split_lowerSplit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_lowerSplit.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictInvariant|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_unSplit(),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imap(),
                                                                                                                                                                                dictInvariant)))))
    }
    pub fn Data_Profunctor_Split_liftSplit() -> &dyn Any {
        static Data_Profunctor_Split_liftSplit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_liftSplit.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_split(),
                                                                                                                            &&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_identity()),
                                                                                         &&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_identity()))
    }
    pub fn Data_Profunctor_Split_hoistSplit() -> &dyn Any {
        static Data_Profunctor_Split_hoistSplit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_hoistSplit.get_or_init(||
                                                         &Func1::new(move
                                                                         |nat|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_unSplit(),
                                                                                                          &&&Func1::new({
                                                                                                                            let nat
                                                                                                                                =
                                                                                                                                nat.clone();
                                                                                                                            move
                                                                                                                                |f|
                                                                                                                                &Func1::new({
                                                                                                                                                let f
                                                                                                                                                    =
                                                                                                                                                    f.clone();
                                                                                                                                                move
                                                                                                                                                    |g|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_split(),
                                                                                                                                                                                                                                                                                              &&&f),
                                                                                                                                                                                                                                                           g)),
                                                                                                                                                                                     &&&nat)
                                                                                                                                            })
                                                                                                                        }))))
    }
    pub fn Data_Profunctor_Split_functorSplit() -> &dyn Any {
        static Data_Profunctor_Split_functorSplit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Split_functorSplit.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                            &&&add(string("map"),
                                                                                                   &&Func1::new(move
                                                                                                                    |f|
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_unSplit(),
                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                       let f
                                                                                                                                                                           =
                                                                                                                                                                           f.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |g|
                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                           let g
                                                                                                                                                                                               =
                                                                                                                                                                                               g.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |h|
                                                                                                                                                                                               &Func1::new({
                                                                                                                                                                                                               let h
                                                                                                                                                                                                                   =
                                                                                                                                                                                                                   h.clone();
                                                                                                                                                                                                               move
                                                                                                                                                                                                                   |fx|
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Split::Data_Profunctor_Split_split(),
                                                                                                                                                                                                                                                                                                                          &&&g),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                             &&&f),
                                                                                                                                                                                                                                                                                                                          &&&h)),
                                                                                                                                                                                                                                                    fx)
                                                                                                                                                                                                           })
                                                                                                                                                                                       })
                                                                                                                                                                   }))),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>())))
    }
}
