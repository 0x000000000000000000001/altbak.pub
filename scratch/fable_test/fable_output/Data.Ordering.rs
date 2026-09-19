pub mod PureScript_Data_Ordering {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
    pub enum Data_Ordering_Ordering {
        Data_Ordering_LTusd_Ctor,
        Data_Ordering_GTusd_Ctor,
        Data_Ordering_EQusd_Ctor,
    }
    impl core::fmt::Display for
     PureScript_Data_Ordering::Data_Ordering_Ordering {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Ordering_LT() -> &dyn Any {
        static Data_Ordering_LT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ordering_LT.get_or_init(||
                                         &LrcPtr::new(PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor))
    }
    pub fn Data_Ordering_GT() -> &dyn Any {
        static Data_Ordering_GT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ordering_GT.get_or_init(||
                                         &LrcPtr::new(PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor))
    }
    pub fn Data_Ordering_EQ() -> &dyn Any {
        static Data_Ordering_EQ: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ordering_EQ.get_or_init(||
                                         &LrcPtr::new(PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))
    }
    pub fn Data_Ordering_showOrdering() -> &dyn Any {
        static Data_Ordering_showOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ordering_showOrdering.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                    &&&add(string("show"),
                                                                                           &&Func1::new(move
                                                                                                            |v|
                                                                                                            {
                                                                                                                let matchValue:
                                                                                                                        LrcPtr<PureScript_Data_Ordering::Data_Ordering_Ordering> =
                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                match matchValue.as_ref()
                                                                                                                    {
                                                                                                                    PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                    =>
                                                                                                                    &string("GT"),
                                                                                                                    PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                    =>
                                                                                                                    &string("EQ"),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    &string("LT"),
                                                                                                                }
                                                                                                            }),
                                                                                           empty::<string,
                                                                                                   &dyn Any>())))
    }
    pub fn Data_Ordering_semigroupOrdering() -> &dyn Any {
        static Data_Ordering_semigroupOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ordering_semigroupOrdering.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                         &&&add(string("append"),
                                                                                                &&Func1::new(move
                                                                                                                 |v|
                                                                                                                 &Func1::new({
                                                                                                                                 let v
                                                                                                                                     =
                                                                                                                                     v.clone();
                                                                                                                                 move
                                                                                                                                     |v1|
                                                                                                                                     {
                                                                                                                                         let matchValue:
                                                                                                                                                 LrcPtr<PureScript_Data_Ordering::Data_Ordering_Ordering> =
                                                                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                                                                         let matchValue_1 =
                                                                                                                                             Sharpurs_Prelude::unbox(v1);
                                                                                                                                         match matchValue.as_ref()
                                                                                                                                             {
                                                                                                                                             PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                             =>
                                                                                                                                             &LrcPtr::new(PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                                                                                             PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                             =>
                                                                                                                                             &matchValue_1,
                                                                                                                                             _
                                                                                                                                             =>
                                                                                                                                             &LrcPtr::new(PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                                                                                         }
                                                                                                                                     }
                                                                                                                             })),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Ordering_invert() -> &dyn Any {
        static Data_Ordering_invert: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ordering_invert.get_or_init(||
                                             &Func1::new(move |v|
                                                             {
                                                                 let matchValue:
                                                                         LrcPtr<PureScript_Data_Ordering::Data_Ordering_Ordering> =
                                                                     Sharpurs_Prelude::unbox(v);
                                                                 match matchValue.as_ref()
                                                                     {
                                                                     PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                     =>
                                                                     &LrcPtr::new(PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor),
                                                                     PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                     =>
                                                                     &LrcPtr::new(PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                     _ =>
                                                                     &LrcPtr::new(PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                 }
                                                             }))
    }
    pub fn Data_Ordering_eqOrdering() -> &dyn Any {
        static Data_Ordering_eqOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ordering_eqOrdering.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                  &&&add(string("eq"),
                                                                                         &&Func1::new(move
                                                                                                          |v|
                                                                                                          &Func1::new({
                                                                                                                          let v
                                                                                                                              =
                                                                                                                              v.clone();
                                                                                                                          move
                                                                                                                              |v1|
                                                                                                                              {
                                                                                                                                  let matchValue:
                                                                                                                                          LrcPtr<PureScript_Data_Ordering::Data_Ordering_Ordering> =
                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                  let matchValue_1:
                                                                                                                                          LrcPtr<PureScript_Data_Ordering::Data_Ordering_Ordering> =
                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                  if let PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                         =
                                                                                                                                         matchValue.as_ref()
                                                                                                                                     {
                                                                                                                                      if let PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                             =
                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                         {
                                                                                                                                          &true
                                                                                                                                      } else {
                                                                                                                                          &false
                                                                                                                                      }
                                                                                                                                  } else {
                                                                                                                                      if let PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                             =
                                                                                                                                             matchValue.as_ref()
                                                                                                                                         {
                                                                                                                                          if let PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                                 =
                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                             {
                                                                                                                                              &true
                                                                                                                                          } else {
                                                                                                                                              &false
                                                                                                                                          }
                                                                                                                                      } else {
                                                                                                                                          if let PureScript_Data_Ordering::Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                 =
                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                             {
                                                                                                                                              &true
                                                                                                                                          } else {
                                                                                                                                              &false
                                                                                                                                          }
                                                                                                                                      }
                                                                                                                                  }
                                                                                                                              }
                                                                                                                      })),
                                                                                         empty::<string,
                                                                                                 &dyn Any>())))
    }
}
