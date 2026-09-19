pub mod PureScript_Data_List_Partial {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_d662adf2::PureScript_Data_List_Types::Data_List_Types_List;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_List_Partial_tail() -> &dyn Any {
        static Data_List_Partial_tail: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Partial_tail.get_or_init(||
                                               &Func1::new(move |usd__unused|
                                                               &Func1::new(move
                                                                               |v|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&Func1::new({
                                                                                                                                  let v
                                                                                                                                      =
                                                                                                                                      v.clone();
                                                                                                                                  move
                                                                                                                                      |usd__unused_1|
                                                                                                                                      {
                                                                                                                                          let matchValue:
                                                                                                                                                  LrcPtr<Data_List_Types_List> =
                                                                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                                                                          if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                    matchValue_1_1)
                                                                                                                                                 =
                                                                                                                                                 matchValue.as_ref()
                                                                                                                                             {
                                                                                                                                              &match matchValue.as_ref()
                                                                                                                                                   {
                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                      x)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                                   _
                                                                                                                                                   =>
                                                                                                                                                   unreachable!(),
                                                                                                                                               }
                                                                                                                                          } else {
                                                                                                                                              panic!("{}",
                                                                                                                                                     string("Match failure"),)
                                                                                                                                          }
                                                                                                                                      }
                                                                                                                              }),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))))
    }
    pub fn Data_List_Partial_last_004010() -> &dyn Any {
        &Func1::new(move |usd__unused|
                        Func1::new({
                                       let usd__unused = usd__unused.clone();
                                       move |v|
                                           PureScript_Data_List_Partial::Data_List_Partial_last_tco(&usd__unused,
                                                                                                    v)
                                   }))
    }
    pub fn Data_List_Partial_last_004010_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Partial_last_004010_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Partial_last_004010_002d1.get_or_init(||
                                                            Lazy(Data_List_Partial_last_004010.clone()))
    }
    pub fn Data_List_Partial_last_tco(usd__unused: &dyn Any, v: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&Func1::new({
                                                           let Data_List_Partial_last_004010_002d1
                                                               =
                                                               Data_List_Partial_last_004010_002d1.clone();
                                                           let v = v.clone();
                                                           move
                                                               |usd__unused_1|
                                                               {
                                                                   let matchValue:
                                                                           LrcPtr<Data_List_Types_List> =
                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                   if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                             matchValue_1_1)
                                                                          =
                                                                          matchValue.as_ref()
                                                                      {
                                                                       if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                              =
                                                                              Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                     {
                                                                                                                     Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                        x)
                                                                                                                     =>
                                                                                                                     x.clone(),
                                                                                                                     _
                                                                                                                     =>
                                                                                                                     unreachable!(),
                                                                                                                 }).as_ref()
                                                                          {
                                                                           &match matchValue.as_ref()
                                                                                {
                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                   _)
                                                                                =>
                                                                                x.clone(),
                                                                                _
                                                                                =>
                                                                                unreachable!(),
                                                                            }
                                                                       } else {
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Partial_last_004010_002d1.Value,
                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                            &&&match matchValue.as_ref()
                                                                                                                   {
                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                      x)
                                                                                                                   =>
                                                                                                                   x.clone(),
                                                                                                                   _
                                                                                                                   =>
                                                                                                                   unreachable!(),
                                                                                                               })
                                                                       }
                                                                   } else {
                                                                       panic!("{}",
                                                                              string("Match failure: PureScript_Data_List_Types.Data_List_Types_List"),)
                                                                   }
                                                               }
                                                       }),
                                         &&&Sharpurs_Prelude::Prim_undefined())
    }
    pub fn Data_List_Partial_last() -> &dyn Any {
        static Data_List_Partial_last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Partial_last.get_or_init(||
                                               Data_List_Partial_last_004010_002d1.Value)
    }
    pub fn Data_List_Partial_init_004014() -> &dyn Any {
        &Func1::new(move |usd__unused|
                        Func1::new({
                                       let usd__unused = usd__unused.clone();
                                       move |v|
                                           PureScript_Data_List_Partial::Data_List_Partial_init_tco(&usd__unused,
                                                                                                    v)
                                   }))
    }
    pub fn Data_List_Partial_init_004014_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Partial_init_004014_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Partial_init_004014_002d1.get_or_init(||
                                                            Lazy(Data_List_Partial_init_004014.clone()))
    }
    pub fn Data_List_Partial_init_tco(usd__unused: &dyn Any, v: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&Func1::new({
                                                           let Data_List_Partial_init_004014_002d1
                                                               =
                                                               Data_List_Partial_init_004014_002d1.clone();
                                                           let v = v.clone();
                                                           move
                                                               |usd__unused_1|
                                                               {
                                                                   let matchValue:
                                                                           LrcPtr<Data_List_Types_List> =
                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                   if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                             matchValue_1_1)
                                                                          =
                                                                          matchValue.as_ref()
                                                                      {
                                                                       if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                              =
                                                                              Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                     {
                                                                                                                     Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                        x)
                                                                                                                     =>
                                                                                                                     x.clone(),
                                                                                                                     _
                                                                                                                     =>
                                                                                                                     unreachable!(),
                                                                                                                 }).as_ref()
                                                                          {
                                                                           &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)
                                                                       } else {
                                                                           &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                {
                                                                                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                   _)
                                                                                                                                                =>
                                                                                                                                                x.clone(),
                                                                                                                                                _
                                                                                                                                                =>
                                                                                                                                                unreachable!(),
                                                                                                                                            },
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Partial_init_004014_002d1.Value,
                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                   {
                                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                      x)
                                                                                                                                                                                   =>
                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                   _
                                                                                                                                                                                   =>
                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                               })))
                                                                       }
                                                                   } else {
                                                                       panic!("{}",
                                                                              string("Match failure: PureScript_Data_List_Types.Data_List_Types_List"),)
                                                                   }
                                                               }
                                                       }),
                                         &&&Sharpurs_Prelude::Prim_undefined())
    }
    pub fn Data_List_Partial_init() -> &dyn Any {
        static Data_List_Partial_init: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Partial_init.get_or_init(||
                                               Data_List_Partial_init_004014_002d1.Value)
    }
    pub fn Data_List_Partial_head() -> &dyn Any {
        static Data_List_Partial_head: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Partial_head.get_or_init(||
                                               &Func1::new(move |usd__unused|
                                                               &Func1::new(move
                                                                               |v|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&Func1::new({
                                                                                                                                  let v
                                                                                                                                      =
                                                                                                                                      v.clone();
                                                                                                                                  move
                                                                                                                                      |usd__unused_1|
                                                                                                                                      {
                                                                                                                                          let matchValue:
                                                                                                                                                  LrcPtr<Data_List_Types_List> =
                                                                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                                                                          if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                    matchValue_1_1)
                                                                                                                                                 =
                                                                                                                                                 matchValue.as_ref()
                                                                                                                                             {
                                                                                                                                              &match matchValue.as_ref()
                                                                                                                                                   {
                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                      _)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                                   _
                                                                                                                                                   =>
                                                                                                                                                   unreachable!(),
                                                                                                                                               }
                                                                                                                                          } else {
                                                                                                                                              panic!("{}",
                                                                                                                                                     string("Match failure"),)
                                                                                                                                          }
                                                                                                                                      }
                                                                                                                              }),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))))
    }
}
