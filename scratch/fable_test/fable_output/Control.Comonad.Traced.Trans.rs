pub mod PureScript_Control_Comonad_Traced_Trans {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_35294a53::PureScript_Control_Comonad_Trans_Class;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Traced_Trans_TracedT() -> &dyn Any {
        static Control_Comonad_Traced_Trans_TracedT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Comonad_Traced_Trans_TracedT.get_or_init(||
                                                             &Func1::new(move
                                                                             |x|
                                                                             x.clone()))
    }
    pub fn Control_Comonad_Traced_Trans_runTracedT() -> &dyn Any {
        static Control_Comonad_Traced_Trans_runTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Trans_runTracedT.get_or_init(||
                                                                &Func1::new(move
                                                                                |v|
                                                                                &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Comonad_Traced_Trans_newtypeTracedT() -> &dyn Any {
        static Control_Comonad_Traced_Trans_newtypeTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Trans_newtypeTracedT.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                     &&&add(string("Coercible0"),
                                                                                                            &&Func1::new(move
                                                                                                                             |usd__unused|
                                                                                                                             &Sharpurs_Prelude::Prim_undefined()),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>())))
    }
    pub fn Control_Comonad_Traced_Trans_functorTracedT() -> &dyn Any {
        static Control_Comonad_Traced_Trans_functorTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Trans_functorTracedT.get_or_init(||
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
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_TracedT(),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                   &&&dictFunctor),
                                                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                                                  |g|
                                                                                                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                                                                                                  let g
                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                      g.clone();
                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                      |t|
                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                                                                                                                                                                          t))
                                                                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                                                                             &&&matchValue_1))
                                                                                                                                                                     }
                                                                                                                                                             })
                                                                                                                                         }),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>()))))
    }
    pub fn Control_Comonad_Traced_Trans_extendTracedT() -> &dyn Any {
        static Control_Comonad_Traced_Trans_extendTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Trans_extendTracedT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictExtend|
                                                                                   {
                                                                                       let Functor0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictExtend)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let functorTracedT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_functorTracedT(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(dictExtend)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       &Func1::new({
                                                                                                       let Functor0
                                                                                                           =
                                                                                                           Functor0.clone();
                                                                                                       let dictExtend
                                                                                                           =
                                                                                                           dictExtend.clone();
                                                                                                       let functorTracedT1
                                                                                                           =
                                                                                                           functorTracedT1.clone();
                                                                                                       move
                                                                                                           |dictSemigroup|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                                                                            &&&add(string("extend"),
                                                                                                                                                   &&Func1::new({
                                                                                                                                                                    let dictSemigroup
                                                                                                                                                                        =
                                                                                                                                                                        dictSemigroup.clone();
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
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_TracedT(),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                                                                                          &&&dictExtend),
                                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                         |w_prime|
                                                                                                                                                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                                                                                                                                                         let w_prime
                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                             w_prime.clone();
                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                             |t|
                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                 &&&matchValue),
                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_TracedT(),
                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         let t
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             t.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |h|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             let h
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 h.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |t_prime|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&h),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&t),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     t_prime))
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         })
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     })),
                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&w_prime)))
                                                                                                                                                                                                                                                                                                                                     }))),
                                                                                                                                                                                                                                                                    &&&matchValue_1))
                                                                                                                                                                                            }
                                                                                                                                                                                    })
                                                                                                                                                                }),
                                                                                                                                                   add(string("Functor0"),
                                                                                                                                                       &&Func1::new(move
                                                                                                                                                                        |usd__unused|
                                                                                                                                                                        &functorTracedT1),
                                                                                                                                                       empty::<string,
                                                                                                                                                               &dyn Any>())))
                                                                                                   })
                                                                                   }))
    }
    pub fn Control_Comonad_Traced_Trans_comonadTransTracedT() -> &dyn Any {
        static Control_Comonad_Traced_Trans_comonadTransTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Trans_comonadTransTracedT.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |dictMonoid|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_ComonadTransusd_Dict(),
                                                                                                                          &&&add(string("lower"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let dictMonoid
                                                                                                                                                      =
                                                                                                                                                      dictMonoid.clone();
                                                                                                                                                  move
                                                                                                                                                      |dictComonad|
                                                                                                                                                      {
                                                                                                                                                          let Functor0 =
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictComonad)),
                                                                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                          &Func1::new({
                                                                                                                                                                          let Functor0
                                                                                                                                                                              =
                                                                                                                                                                              Functor0.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v|
                                                                                                                                                                              {
                                                                                                                                                                                  let w =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                         &&&Functor0),
                                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                                        |f|
                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                                                                                                            &&&dictMonoid)))),
                                                                                                                                                                                                                   &&&w)
                                                                                                                                                                              }
                                                                                                                                                                      })
                                                                                                                                                      }
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>()))))
    }
    pub fn Control_Comonad_Traced_Trans_comonadTracedT() -> &dyn Any {
        static Control_Comonad_Traced_Trans_comonadTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Trans_comonadTracedT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictComonad|
                                                                                    {
                                                                                        let extendTracedT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_extendTracedT(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(dictComonad)),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        &Func1::new({
                                                                                                        let dictComonad
                                                                                                            =
                                                                                                            dictComonad.clone();
                                                                                                        let extendTracedT1
                                                                                                            =
                                                                                                            extendTracedT1.clone();
                                                                                                        move
                                                                                                            |dictMonoid|
                                                                                                            {
                                                                                                                let extendTracedT2 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&extendTracedT1,
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_Comonadusd_Dict(),
                                                                                                                                                 &&&add(string("extract"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let dictMonoid
                                                                                                                                                                             =
                                                                                                                                                                             dictMonoid.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             {
                                                                                                                                                                                 let w =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                                                                                        &&&dictComonad),
                                                                                                                                                                                                                                                     &&&w),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                     &&&dictMonoid))
                                                                                                                                                                             }
                                                                                                                                                                     }),
                                                                                                                                                        add(string("Extend0"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let extendTracedT2
                                                                                                                                                                                 =
                                                                                                                                                                                 extendTracedT2.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |usd__unused|
                                                                                                                                                                                 &extendTracedT2
                                                                                                                                                                         }),
                                                                                                                                                            empty::<string,
                                                                                                                                                                    &dyn Any>())))
                                                                                                            }
                                                                                                    })
                                                                                    }))
    }
}
