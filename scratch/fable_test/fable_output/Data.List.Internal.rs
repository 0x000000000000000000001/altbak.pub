pub mod PureScript_Data_List_Internal {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_d662adf2::PureScript_Data_List_Types::Data_List_Types_List;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    #[derive(Clone, Debug,)]
    pub enum Data_List_Internal_Set {
        Data_List_Internal_Leafusd_Ctor,
        Data_List_Internal_Twousd_Ctor(&dyn Any, &dyn Any, &dyn Any),
        Data_List_Internal_Threeusd_Ctor(&dyn Any, &dyn Any, &dyn Any,
                                         &dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_List_Internal::Data_List_Internal_Set {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug,)]
    pub enum Data_List_Internal_TreeContext {
        Data_List_Internal_TwoLeftusd_Ctor(&dyn Any, &dyn Any),
        Data_List_Internal_TwoRightusd_Ctor(&dyn Any, &dyn Any),
        Data_List_Internal_ThreeLeftusd_Ctor(&dyn Any, &dyn Any, &dyn Any,
                                             &dyn Any),
        Data_List_Internal_ThreeMiddleusd_Ctor(&dyn Any, &dyn Any, &dyn Any,
                                               &dyn Any),
        Data_List_Internal_ThreeRightusd_Ctor(&dyn Any, &dyn Any, &dyn Any,
                                              &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_List_Internal::Data_List_Internal_TreeContext {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug,)]
    pub enum Data_List_Internal_KickUp {
        Data_List_Internal_KickUpusd_Ctor(&dyn Any, &dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_List_Internal::Data_List_Internal_KickUp {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_List_Internal_Leaf() -> &dyn Any {
        static Data_List_Internal_Leaf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_Leaf.get_or_init(||
                                                &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Leafusd_Ctor))
    }
    pub fn Data_List_Internal_Two() -> &dyn Any {
        static Data_List_Internal_Two: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_Two.get_or_init(||
                                               &Func1::new(move |usd__arg1|
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
                                                                                                     &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(usd__arg1,
                                                                                                                                                                                                        usd__arg2,
                                                                                                                                                                                                        usd__arg3.clone()))
                                                                                             })
                                                                          })))
    }
    pub fn Data_List_Internal_Three() -> &dyn Any {
        static Data_List_Internal_Three: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_Three.get_or_init(||
                                                 &Func1::new(move |usd__arg1|
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
                                                                                                       Func1::new({
                                                                                                                      let usd__arg3
                                                                                                                          =
                                                                                                                          usd__arg3.clone();
                                                                                                                      move
                                                                                                                          |usd__arg4|
                                                                                                                          Func1::new({
                                                                                                                                         let usd__arg4
                                                                                                                                             =
                                                                                                                                             usd__arg4.clone();
                                                                                                                                         move
                                                                                                                                             |usd__arg5|
                                                                                                                                             &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Threeusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                  usd__arg2,
                                                                                                                                                                                                                                                  usd__arg3,
                                                                                                                                                                                                                                                  usd__arg4,
                                                                                                                                                                                                                                                  usd__arg5.clone()))
                                                                                                                                     })
                                                                                                                  })
                                                                                               })
                                                                            })))
    }
    pub fn Data_List_Internal_TwoLeft() -> &dyn Any {
        static Data_List_Internal_TwoLeft: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_TwoLeft.get_or_init(||
                                                   &Func1::new(move
                                                                   |usd__arg1|
                                                                   Func1::new({
                                                                                  let usd__arg1
                                                                                      =
                                                                                      usd__arg1.clone();
                                                                                  move
                                                                                      |usd__arg2|
                                                                                      &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_TwoLeftusd_Ctor(usd__arg1,
                                                                                                                                                                                                     usd__arg2.clone()))
                                                                              })))
    }
    pub fn Data_List_Internal_TwoRight() -> &dyn Any {
        static Data_List_Internal_TwoRight: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_TwoRight.get_or_init(||
                                                    &Func1::new(move
                                                                    |usd__arg1|
                                                                    Func1::new({
                                                                                   let usd__arg1
                                                                                       =
                                                                                       usd__arg1.clone();
                                                                                   move
                                                                                       |usd__arg2|
                                                                                       &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_TwoRightusd_Ctor(usd__arg1,
                                                                                                                                                                                                       usd__arg2.clone()))
                                                                               })))
    }
    pub fn Data_List_Internal_ThreeLeft() -> &dyn Any {
        static Data_List_Internal_ThreeLeft: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_ThreeLeft.get_or_init(||
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
                                                                                                           Func1::new({
                                                                                                                          let usd__arg3
                                                                                                                              =
                                                                                                                              usd__arg3.clone();
                                                                                                                          move
                                                                                                                              |usd__arg4|
                                                                                                                              &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeLeftusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                               usd__arg2,
                                                                                                                                                                                                                                               usd__arg3,
                                                                                                                                                                                                                                               usd__arg4.clone()))
                                                                                                                      })
                                                                                                   })
                                                                                })))
    }
    pub fn Data_List_Internal_ThreeMiddle() -> &dyn Any {
        static Data_List_Internal_ThreeMiddle: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_ThreeMiddle.get_or_init(||
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
                                                                                                             Func1::new({
                                                                                                                            let usd__arg3
                                                                                                                                =
                                                                                                                                usd__arg3.clone();
                                                                                                                            move
                                                                                                                                |usd__arg4|
                                                                                                                                &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeMiddleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                   usd__arg2,
                                                                                                                                                                                                                                                   usd__arg3,
                                                                                                                                                                                                                                                   usd__arg4.clone()))
                                                                                                                        })
                                                                                                     })
                                                                                  })))
    }
    pub fn Data_List_Internal_ThreeRight() -> &dyn Any {
        static Data_List_Internal_ThreeRight: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_ThreeRight.get_or_init(||
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
                                                                                                            Func1::new({
                                                                                                                           let usd__arg3
                                                                                                                               =
                                                                                                                               usd__arg3.clone();
                                                                                                                           move
                                                                                                                               |usd__arg4|
                                                                                                                               &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeRightusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                 usd__arg2,
                                                                                                                                                                                                                                                 usd__arg3,
                                                                                                                                                                                                                                                 usd__arg4.clone()))
                                                                                                                       })
                                                                                                    })
                                                                                 })))
    }
    pub fn Data_List_Internal_KickUp() -> &dyn Any {
        static Data_List_Internal_KickUp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_KickUp.get_or_init(||
                                                  &Func1::new(move |usd__arg1|
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
                                                                                                        &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(usd__arg1,
                                                                                                                                                                                                                 usd__arg2,
                                                                                                                                                                                                                 usd__arg3.clone()))
                                                                                                })
                                                                             })))
    }
    pub fn Data_List_Internal_fromZipper_004041() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Data_List_Internal::Data_List_Internal_fromZipper_tco(&v,
                                                                                                            v1)
                                   }))
    }
    pub fn Data_List_Internal_fromZipper_004041_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Internal_fromZipper_004041_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Internal_fromZipper_004041_002d1.get_or_init(||
                                                                   Lazy(Data_List_Internal_fromZipper_004041.clone()))
    }
    pub fn Data_List_Internal_fromZipper_tco(v: &dyn Any, v1: &dyn Any)
     -> &dyn Any {
        let matchValue: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        match matchValue.as_ref() {
            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                               matchValue_1_1)
            => {
                let tree_1 = matchValue_1;
                let ctx = matchValue_1_1.clone();
                let matchValue_3:
                        LrcPtr<PureScript_Data_List_Internal::Data_List_Internal_TreeContext> =
                    Sharpurs_Prelude::unbox(&matchValue_1_0);
                match matchValue_3.as_ref() {
                    PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_TwoRightusd_Ctor(matchValue_3_1_0,
                                                                                                                       matchValue_3_1_1)
                    =>
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Internal_fromZipper_004041_002d1.Value,
                                                                                        &&&ctx),
                                                     &&&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(matchValue_3_1_0,
                                                                                                                                                          matchValue_3_1_1,
                                                                                                                                                          &tree_1))),
                    PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeLeftusd_Ctor(matchValue_3_2_0,
                                                                                                                        matchValue_3_2_1,
                                                                                                                        matchValue_3_2_2,
                                                                                                                        matchValue_3_2_3)
                    =>
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Internal_fromZipper_004041_002d1.Value,
                                                                                        &&&ctx),
                                                     &&&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Threeusd_Ctor(&tree_1,
                                                                                                                                                            matchValue_3_2_0,
                                                                                                                                                            matchValue_3_2_1,
                                                                                                                                                            matchValue_3_2_2,
                                                                                                                                                            matchValue_3_2_3))),
                    PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeMiddleusd_Ctor(matchValue_3_3_0,
                                                                                                                          matchValue_3_3_1,
                                                                                                                          matchValue_3_3_2,
                                                                                                                          matchValue_3_3_3)
                    =>
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Internal_fromZipper_004041_002d1.Value,
                                                                                        &&&ctx),
                                                     &&&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Threeusd_Ctor(matchValue_3_3_0,
                                                                                                                                                            matchValue_3_3_1,
                                                                                                                                                            &tree_1,
                                                                                                                                                            matchValue_3_3_2,
                                                                                                                                                            matchValue_3_3_3))),
                    PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeRightusd_Ctor(matchValue_3_4_0,
                                                                                                                         matchValue_3_4_1,
                                                                                                                         matchValue_3_4_2,
                                                                                                                         matchValue_3_4_3)
                    =>
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Internal_fromZipper_004041_002d1.Value,
                                                                                        &&&ctx),
                                                     &&&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Threeusd_Ctor(matchValue_3_4_0,
                                                                                                                                                            matchValue_3_4_1,
                                                                                                                                                            matchValue_3_4_2,
                                                                                                                                                            matchValue_3_4_3,
                                                                                                                                                            &tree_1))),
                    PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_TwoLeftusd_Ctor(matchValue_3_0_0,
                                                                                                                      matchValue_3_0_1)
                    =>
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Internal_fromZipper_004041_002d1.Value,
                                                                                        &&&ctx),
                                                     &&&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(&tree_1,
                                                                                                                                                          matchValue_3_0_0,
                                                                                                                                                          matchValue_3_0_1))),
                }
            }
            _ => &matchValue_1,
        }
    }
    pub fn Data_List_Internal_fromZipper() -> &dyn Any {
        static Data_List_Internal_fromZipper: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_fromZipper.get_or_init(||
                                                      Data_List_Internal_fromZipper_004041_002d1.Value)
    }
    pub fn Data_List_Internal_insertAndLookupBy() -> &dyn Any {
        static Data_List_Internal_insertAndLookupBy: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Internal_insertAndLookupBy.get_or_init(||
                                                             &Func1::new(move
                                                                             |comp|
                                                                             &Func1::new({
                                                                                             let comp
                                                                                                 =
                                                                                                 comp.clone();
                                                                                             move
                                                                                                 |k|
                                                                                                 &Func1::new({
                                                                                                                 let k
                                                                                                                     =
                                                                                                                     k.clone();
                                                                                                                 move
                                                                                                                     |orig|
                                                                                                                     {
                                                                                                                         let up_2 =
                                                                                                                             Func0::new({
                                                                                                                                            let up_tco
                                                                                                                                                =
                                                                                                                                                up_tco.clone();
                                                                                                                                            move
                                                                                                                                                ||
                                                                                                                                                &Func1::new({
                                                                                                                                                                let up_tco
                                                                                                                                                                    =
                                                                                                                                                                    up_tco.clone();
                                                                                                                                                                move
                                                                                                                                                                    |v|
                                                                                                                                                                    Func1::new({
                                                                                                                                                                                   let up_tco
                                                                                                                                                                                       =
                                                                                                                                                                                       up_tco.clone();
                                                                                                                                                                                   let v
                                                                                                                                                                                       =
                                                                                                                                                                                       v.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |v1|
                                                                                                                                                                                       up_tco(v)(v1.clone())
                                                                                                                                                                               })
                                                                                                                                                            })
                                                                                                                                        });
                                                                                                                         let up_1 =
                                                                                                                             Lazy(up_2);
                                                                                                                         fn up_tco(v_1:
                                                                                                                                       _)
                                                                                                                          ->
                                                                                                                              Func1<&dyn Any,
                                                                                                                                    &dyn Any> {
                                                                                                                             Func1::new({
                                                                                                                                            let up_tco
                                                                                                                                                =
                                                                                                                                                up_tco.clone();
                                                                                                                                            let v_1
                                                                                                                                                =
                                                                                                                                                v_1.clone();
                                                                                                                                            move
                                                                                                                                                |v1_1|
                                                                                                                                                {
                                                                                                                                                    let matchValue:
                                                                                                                                                            LrcPtr<Data_List_Types_List> =
                                                                                                                                                        Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                    let matchValue_1:
                                                                                                                                                            LrcPtr<PureScript_Data_List_Internal::Data_List_Internal_KickUp> =
                                                                                                                                                        Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                    match matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                           matchValue_1_1)
                                                                                                                                                        =>
                                                                                                                                                        {
                                                                                                                                                            let ctx =
                                                                                                                                                                matchValue_1_1.clone();
                                                                                                                                                            let matchValue_3:
                                                                                                                                                                    LrcPtr<PureScript_Data_List_Internal::Data_List_Internal_TreeContext> =
                                                                                                                                                                Sharpurs_Prelude::unbox(&matchValue_1_0);
                                                                                                                                                            let matchValue_4:
                                                                                                                                                                    LrcPtr<PureScript_Data_List_Internal::Data_List_Internal_KickUp> =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&matchValue_1);
                                                                                                                                                            match matchValue_3.as_ref()
                                                                                                                                                                {
                                                                                                                                                                PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_TwoRightusd_Ctor(matchValue_3_1_0,
                                                                                                                                                                                                                                                                   matchValue_3_1_1)
                                                                                                                                                                =>
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Internal::Data_List_Internal_fromZipper(),
                                                                                                                                                                                                                                    &&&ctx),
                                                                                                                                                                                                 &&&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Threeusd_Ctor(matchValue_3_1_0,
                                                                                                                                                                                                                                                                                                        matchValue_3_1_1,
                                                                                                                                                                                                                                                                                                        &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                         _,
                                                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                        &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                         x,
                                                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                        &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                         _,
                                                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                         }))),
                                                                                                                                                                PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeLeftusd_Ctor(matchValue_3_2_0,
                                                                                                                                                                                                                                                                    matchValue_3_2_1,
                                                                                                                                                                                                                                                                    matchValue_3_2_2,
                                                                                                                                                                                                                                                                    matchValue_3_2_3)
                                                                                                                                                                =>
                                                                                                                                                                up_tco(&ctx)(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                                                                                         &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                                                                                         &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                          })),
                                                                                                                                                                                                                                                                                      matchValue_3_2_0,
                                                                                                                                                                                                                                                                                      &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(matchValue_3_2_1,
                                                                                                                                                                                                                                                                                                                                                                                         matchValue_3_2_2,
                                                                                                                                                                                                                                                                                                                                                                                         matchValue_3_2_3))))),
                                                                                                                                                                PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeMiddleusd_Ctor(matchValue_3_3_0,
                                                                                                                                                                                                                                                                      matchValue_3_3_1,
                                                                                                                                                                                                                                                                      matchValue_3_3_2,
                                                                                                                                                                                                                                                                      matchValue_3_3_3)
                                                                                                                                                                =>
                                                                                                                                                                up_tco(&ctx)(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(matchValue_3_3_0,
                                                                                                                                                                                                                                                                                                                                                                                         matchValue_3_3_1,
                                                                                                                                                                                                                                                                                                                                                                                         &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                          })),
                                                                                                                                                                                                                                                                                      &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                           PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                       x,
                                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                       },
                                                                                                                                                                                                                                                                                      &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                                                                                         matchValue_3_3_2,
                                                                                                                                                                                                                                                                                                                                                                                         matchValue_3_3_3))))),
                                                                                                                                                                PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeRightusd_Ctor(matchValue_3_4_0,
                                                                                                                                                                                                                                                                     matchValue_3_4_1,
                                                                                                                                                                                                                                                                     matchValue_3_4_2,
                                                                                                                                                                                                                                                                     matchValue_3_4_3)
                                                                                                                                                                =>
                                                                                                                                                                up_tco(&ctx)(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(matchValue_3_4_0,
                                                                                                                                                                                                                                                                                                                                                                                         matchValue_3_4_1,
                                                                                                                                                                                                                                                                                                                                                                                         matchValue_3_4_2)),
                                                                                                                                                                                                                                                                                      matchValue_3_4_3,
                                                                                                                                                                                                                                                                                      &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                                                                                         &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                                                                                         &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                          }))))),
                                                                                                                                                                PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_TwoLeftusd_Ctor(matchValue_3_0_0,
                                                                                                                                                                                                                                                                  matchValue_3_0_1)
                                                                                                                                                                =>
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Internal::Data_List_Internal_fromZipper(),
                                                                                                                                                                                                                                    &&&ctx),
                                                                                                                                                                                                 &&&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Threeusd_Ctor(&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                         _,
                                                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                        &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                         x,
                                                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                        &match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                         _,
                                                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                        matchValue_3_0_0,
                                                                                                                                                                                                                                                                                                        matchValue_3_0_1))),
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        _
                                                                                                                                                        =>
                                                                                                                                                        &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                            },
                                                                                                                                                                                                                                                           &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                            x,
                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                            },
                                                                                                                                                                                                                                                           &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                            })),
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                        })
                                                                                                                         }
                                                                                                                         let up =
                                                                                                                             up_1.Value;
                                                                                                                         {
                                                                                                                             let down_2 =
                                                                                                                                 Func0::new({
                                                                                                                                                let down_tco
                                                                                                                                                    =
                                                                                                                                                    down_tco.clone();
                                                                                                                                                move
                                                                                                                                                    ||
                                                                                                                                                    &Func1::new({
                                                                                                                                                                    let down_tco
                                                                                                                                                                        =
                                                                                                                                                                        down_tco.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |v_2|
                                                                                                                                                                        Func1::new({
                                                                                                                                                                                       let down_tco
                                                                                                                                                                                           =
                                                                                                                                                                                           down_tco.clone();
                                                                                                                                                                                       let v_2
                                                                                                                                                                                           =
                                                                                                                                                                                           v_2.clone();
                                                                                                                                                                                       move
                                                                                                                                                                                           |v1_2|
                                                                                                                                                                                           down_tco(v_2)(v1_2.clone())
                                                                                                                                                                                   })
                                                                                                                                                                })
                                                                                                                                            });
                                                                                                                             let down_1 =
                                                                                                                                 Lazy(down_2);
                                                                                                                             let down_tco =
                                                                                                                                 Func1::new({
                                                                                                                                                let orig
                                                                                                                                                    =
                                                                                                                                                    orig.clone();
                                                                                                                                                let up_tco
                                                                                                                                                    =
                                                                                                                                                    up_tco.clone();
                                                                                                                                                move
                                                                                                                                                    |v_3|
                                                                                                                                                    fix1(&(move
                                                                                                                                                               |down_tco,
                                                                                                                                                                v_3|
                                                                                                                                                               Func1::new({
                                                                                                                                                                              let down_tco
                                                                                                                                                                                  =
                                                                                                                                                                                  down_tco.clone();
                                                                                                                                                                              let up_tco
                                                                                                                                                                                  =
                                                                                                                                                                                  up_tco.clone();
                                                                                                                                                                              let v_3
                                                                                                                                                                                  =
                                                                                                                                                                                  v_3.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |v1_3|
                                                                                                                                                                                  {
                                                                                                                                                                                      let matchValue_6 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v_3);
                                                                                                                                                                                      let matchValue_7:
                                                                                                                                                                                              LrcPtr<PureScript_Data_List_Internal::Data_List_Internal_Set> =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(v1_3);
                                                                                                                                                                                      match matchValue_7.as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Twousd_Ctor(matchValue_7_1_0,
                                                                                                                                                                                                                                                                                matchValue_7_1_1,
                                                                                                                                                                                                                                                                                matchValue_7_1_2)
                                                                                                                                                                                          =>
                                                                                                                                                                                          {
                                                                                                                                                                                              let right_3 =
                                                                                                                                                                                                  matchValue_7_1_2.clone();
                                                                                                                                                                                              let left_3 =
                                                                                                                                                                                                  matchValue_7_1_0.clone();
                                                                                                                                                                                              let k1_5 =
                                                                                                                                                                                                  matchValue_7_1_1.clone();
                                                                                                                                                                                              let ctx_2 =
                                                                                                                                                                                                  matchValue_6;
                                                                                                                                                                                              let matchValue_9:
                                                                                                                                                                                                      LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&comp,
                                                                                                                                                                                                                                                                                                &&&k),
                                                                                                                                                                                                                                                             &&&k1_5));
                                                                                                                                                                                              match matchValue_9.as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  &add(string("found"),
                                                                                                                                                                                                       &&true,
                                                                                                                                                                                                       add(string("result"),
                                                                                                                                                                                                           &&orig,
                                                                                                                                                                                                           empty::<string,
                                                                                                                                                                                                                   &dyn Any>())),
                                                                                                                                                                                                  Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  down_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_TwoLeftusd_Ctor(&k1_5,
                                                                                                                                                                                                                                                                                                                                                                                          &right_3)),
                                                                                                                                                                                                                                                                           &ctx_2)))(&left_3),
                                                                                                                                                                                                  _
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  down_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_TwoRightusd_Ctor(&left_3,
                                                                                                                                                                                                                                                                                                                                                                                           &k1_5)),
                                                                                                                                                                                                                                                                           &ctx_2)))(&right_3),
                                                                                                                                                                                              }
                                                                                                                                                                                          }
                                                                                                                                                                                          PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Threeusd_Ctor(matchValue_7_2_0,
                                                                                                                                                                                                                                                                                  matchValue_7_2_1,
                                                                                                                                                                                                                                                                                  matchValue_7_2_2,
                                                                                                                                                                                                                                                                                  matchValue_7_2_3,
                                                                                                                                                                                                                                                                                  matchValue_7_2_4)
                                                                                                                                                                                          =>
                                                                                                                                                                                          {
                                                                                                                                                                                              let right_4 =
                                                                                                                                                                                                  matchValue_7_2_4.clone();
                                                                                                                                                                                              let mid_2 =
                                                                                                                                                                                                  matchValue_7_2_2.clone();
                                                                                                                                                                                              let left_4 =
                                                                                                                                                                                                  matchValue_7_2_0.clone();
                                                                                                                                                                                              let k2_3 =
                                                                                                                                                                                                  matchValue_7_2_3.clone();
                                                                                                                                                                                              let k1_6 =
                                                                                                                                                                                                  matchValue_7_2_1.clone();
                                                                                                                                                                                              let ctx_3 =
                                                                                                                                                                                                  matchValue_6;
                                                                                                                                                                                              let matchValue_10:
                                                                                                                                                                                                      LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&comp,
                                                                                                                                                                                                                                                                                                &&&k),
                                                                                                                                                                                                                                                             &&&k1_6));
                                                                                                                                                                                              if let Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_10.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &add(string("found"),
                                                                                                                                                                                                       &&true,
                                                                                                                                                                                                       add(string("result"),
                                                                                                                                                                                                           &&orig,
                                                                                                                                                                                                           empty::<string,
                                                                                                                                                                                                                   &dyn Any>()))
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  let v3 =
                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&comp,
                                                                                                                                                                                                                                                                          &&&k),
                                                                                                                                                                                                                                       &&&k2_3);
                                                                                                                                                                                                  let matchValue_11:
                                                                                                                                                                                                          LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&&matchValue_10);
                                                                                                                                                                                                  let matchValue_12:
                                                                                                                                                                                                          LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v3);
                                                                                                                                                                                                  if let Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_12.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                      &add(string("found"),
                                                                                                                                                                                                           &&true,
                                                                                                                                                                                                           add(string("result"),
                                                                                                                                                                                                               &&orig,
                                                                                                                                                                                                               empty::<string,
                                                                                                                                                                                                                       &dyn Any>()))
                                                                                                                                                                                                  } else {
                                                                                                                                                                                                      if let Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                                             =
                                                                                                                                                                                                             matchValue_12.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                          if let Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_11.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                              down_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeLeftusd_Ctor(&k1_6,
                                                                                                                                                                                                                                                                                                                                                                                                        &mid_2,
                                                                                                                                                                                                                                                                                                                                                                                                        &k2_3,
                                                                                                                                                                                                                                                                                                                                                                                                        &right_4)),
                                                                                                                                                                                                                                                                                       &ctx_3)))(&left_4)
                                                                                                                                                                                                          } else {
                                                                                                                                                                                                              if let Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     matchValue_11.as_ref()
                                                                                                                                                                                                                 {
                                                                                                                                                                                                                  down_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeMiddleusd_Ctor(&left_4,
                                                                                                                                                                                                                                                                                                                                                                                                              &k1_6,
                                                                                                                                                                                                                                                                                                                                                                                                              &k2_3,
                                                                                                                                                                                                                                                                                                                                                                                                              &right_4)),
                                                                                                                                                                                                                                                                                           &ctx_3)))(&mid_2)
                                                                                                                                                                                                              } else {
                                                                                                                                                                                                                  down_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeRightusd_Ctor(&left_4,
                                                                                                                                                                                                                                                                                                                                                                                                             &k1_6,
                                                                                                                                                                                                                                                                                                                                                                                                             &mid_2,
                                                                                                                                                                                                                                                                                                                                                                                                             &k2_3)),
                                                                                                                                                                                                                                                                                           &ctx_3)))(&right_4)
                                                                                                                                                                                                              }
                                                                                                                                                                                                          }
                                                                                                                                                                                                      } else {
                                                                                                                                                                                                          if let Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_11.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                              down_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeLeftusd_Ctor(&k1_6,
                                                                                                                                                                                                                                                                                                                                                                                                        &mid_2,
                                                                                                                                                                                                                                                                                                                                                                                                        &k2_3,
                                                                                                                                                                                                                                                                                                                                                                                                        &right_4)),
                                                                                                                                                                                                                                                                                       &ctx_3)))(&left_4)
                                                                                                                                                                                                          } else {
                                                                                                                                                                                                              down_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_TreeContext::Data_List_Internal_ThreeRightusd_Ctor(&left_4,
                                                                                                                                                                                                                                                                                                                                                                                                         &k1_6,
                                                                                                                                                                                                                                                                                                                                                                                                         &mid_2,
                                                                                                                                                                                                                                                                                                                                                                                                         &k2_3)),
                                                                                                                                                                                                                                                                                       &ctx_3)))(&right_4)
                                                                                                                                                                                                          }
                                                                                                                                                                                                      }
                                                                                                                                                                                                  }
                                                                                                                                                                                              }
                                                                                                                                                                                          }
                                                                                                                                                                                          _
                                                                                                                                                                                          =>
                                                                                                                                                                                          &add(string("found"),
                                                                                                                                                                                               &&false,
                                                                                                                                                                                               add(string("result"),
                                                                                                                                                                                                   &up_tco(&matchValue_6)(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_KickUp::Data_List_Internal_KickUpusd_Ctor(&LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Leafusd_Ctor),
                                                                                                                                                                                                                                                                                                                                   &k,
                                                                                                                                                                                                                                                                                                                                   &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Leafusd_Ctor)))),
                                                                                                                                                                                                   empty::<string,
                                                                                                                                                                                                           &dyn Any>())),
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                          })),
                                                                                                                                                         v_3.clone())
                                                                                                                                            });
                                                                                                                             let down =
                                                                                                                                 down_1.Value;
                                                                                                                             down_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))(orig.clone())
                                                                                                                         }
                                                                                                                     }
                                                                                                             })
                                                                                         })))
    }
    pub fn Data_List_Internal_emptySet() -> &dyn Any {
        static Data_List_Internal_emptySet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Internal_emptySet.get_or_init(||
                                                    &LrcPtr::new(PureScript_Data_List_Internal::Data_List_Internal_Set::Data_List_Internal_Leafusd_Ctor))
    }
}
