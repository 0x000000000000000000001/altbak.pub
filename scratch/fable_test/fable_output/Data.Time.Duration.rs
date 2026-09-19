pub mod PureScript_Data_Time_Duration {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Time_Duration_negate() -> &dyn Any {
        static Data_Time_Duration_negate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_negate.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringNumber()))
    }
    pub fn Data_Time_Duration_identity() -> &dyn Any {
        static Data_Time_Duration_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_identity.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                     &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Time_Duration_Seconds() -> &dyn Any {
        static Data_Time_Duration_Seconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_Seconds.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Time_Duration_Minutes() -> &dyn Any {
        static Data_Time_Duration_Minutes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_Minutes.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Time_Duration_Milliseconds() -> &dyn Any {
        static Data_Time_Duration_Milliseconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_Milliseconds.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Data_Time_Duration_Hours() -> &dyn Any {
        static Data_Time_Duration_Hours: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_Hours.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Time_Duration_Durationusd_Dict() -> &dyn Any {
        static Data_Time_Duration_Durationusd_Dict: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Time_Duration_Durationusd_Dict.get_or_init(||
                                                            &Func1::new(move
                                                                            |x|
                                                                            x.clone()))
    }
    pub fn Data_Time_Duration_Days() -> &dyn Any {
        static Data_Time_Duration_Days: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_Days.get_or_init(||
                                                &Func1::new(move |x|
                                                                x.clone()))
    }
    pub fn Data_Time_Duration_toDuration() -> &dyn Any {
        static Data_Time_Duration_toDuration: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_toDuration.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("toDuration"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Time_Duration_showSeconds() -> &dyn Any {
        static Data_Time_Duration_showSeconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_showSeconds.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                        &&&add(string("show"),
                                                                                               &&Func1::new(move
                                                                                                                |v|
                                                                                                                {
                                                                                                                    let n =
                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                        &&&string("(Seconds ")),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Show::Data_Show_showNumber()),
                                                                                                                                                                                                                                                              &&&n)),
                                                                                                                                                                                        &&&string(")")))
                                                                                                                }),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))
    }
    pub fn Data_Time_Duration_showMinutes() -> &dyn Any {
        static Data_Time_Duration_showMinutes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_showMinutes.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                        &&&add(string("show"),
                                                                                               &&Func1::new(move
                                                                                                                |v|
                                                                                                                {
                                                                                                                    let n =
                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                        &&&string("(Minutes ")),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Show::Data_Show_showNumber()),
                                                                                                                                                                                                                                                              &&&n)),
                                                                                                                                                                                        &&&string(")")))
                                                                                                                }),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))
    }
    pub fn Data_Time_Duration_showMilliseconds() -> &dyn Any {
        static Data_Time_Duration_showMilliseconds: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Time_Duration_showMilliseconds.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                             &&&add(string("show"),
                                                                                                    &&Func1::new(move
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let n =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                             &&&string("(Milliseconds ")),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Show::Data_Show_showNumber()),
                                                                                                                                                                                                                                                                   &&&n)),
                                                                                                                                                                                             &&&string(")")))
                                                                                                                     }),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Time_Duration_showHours() -> &dyn Any {
        static Data_Time_Duration_showHours: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_showHours.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                      &&&add(string("show"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              {
                                                                                                                  let n =
                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                      &&&string("(Hours ")),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Show::Data_Show_showNumber()),
                                                                                                                                                                                                                                                            &&&n)),
                                                                                                                                                                                      &&&string(")")))
                                                                                                              }),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Time_Duration_showDays() -> &dyn Any {
        static Data_Time_Duration_showDays: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_showDays.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                     &&&add(string("show"),
                                                                                            &&Func1::new(move
                                                                                                             |v|
                                                                                                             {
                                                                                                                 let n =
                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                        &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                     &&&string("(Days ")),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Show::Data_Show_showNumber()),
                                                                                                                                                                                                                                                           &&&n)),
                                                                                                                                                                                     &&&string(")")))
                                                                                                             }),
                                                                                            empty::<string,
                                                                                                    &dyn Any>())))
    }
    pub fn Data_Time_Duration_semigroupSeconds() -> &dyn Any {
        static Data_Time_Duration_semigroupSeconds: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Time_Duration_semigroupSeconds.get_or_init(||
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
                                                                                                                                             let matchValue =
                                                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                                                             let matchValue_1 =
                                                                                                                                                 Sharpurs_Prelude::unbox(v1);
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Seconds(),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                 &&&matchValue_1))
                                                                                                                                         }
                                                                                                                                 })),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Time_Duration_semigroupMinutes() -> &dyn Any {
        static Data_Time_Duration_semigroupMinutes: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Time_Duration_semigroupMinutes.get_or_init(||
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
                                                                                                                                             let matchValue =
                                                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                                                             let matchValue_1 =
                                                                                                                                                 Sharpurs_Prelude::unbox(v1);
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Minutes(),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                 &&&matchValue_1))
                                                                                                                                         }
                                                                                                                                 })),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Time_Duration_semigroupMilliseconds() -> &dyn Any {
        static Data_Time_Duration_semigroupMilliseconds:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_semigroupMilliseconds.get_or_init(||
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
                                                                                                                                                  let matchValue =
                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                  let matchValue_1 =
                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds(),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                                      &&&matchValue_1))
                                                                                                                                              }
                                                                                                                                      })),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_Time_Duration_semigroupHours() -> &dyn Any {
        static Data_Time_Duration_semigroupHours: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_semigroupHours.get_or_init(||
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
                                                                                                                                           let matchValue =
                                                                                                                                               Sharpurs_Prelude::unbox(&&v);
                                                                                                                                           let matchValue_1 =
                                                                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Hours(),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                               &&&matchValue_1))
                                                                                                                                       }
                                                                                                                               })),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Time_Duration_semigroupDays() -> &dyn Any {
        static Data_Time_Duration_semigroupDays: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_semigroupDays.get_or_init(||
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
                                                                                                                                          let matchValue =
                                                                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                                                                          let matchValue_1 =
                                                                                                                                              Sharpurs_Prelude::unbox(v1);
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Days(),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                                                 &&&matchValue),
                                                                                                                                                                                                              &&&matchValue_1))
                                                                                                                                      }
                                                                                                                              })),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>())))
    }
    pub fn Data_Time_Duration_ordSeconds() -> &dyn Any {
        static Data_Time_Duration_ordSeconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_ordSeconds.get_or_init(||
                                                      &PureScript_Data_Ord::Data_Ord_ordNumber())
    }
    pub fn Data_Time_Duration_ordMinutes() -> &dyn Any {
        static Data_Time_Duration_ordMinutes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_ordMinutes.get_or_init(||
                                                      &PureScript_Data_Ord::Data_Ord_ordNumber())
    }
    pub fn Data_Time_Duration_ordMilliseconds() -> &dyn Any {
        static Data_Time_Duration_ordMilliseconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_ordMilliseconds.get_or_init(||
                                                           &PureScript_Data_Ord::Data_Ord_ordNumber())
    }
    pub fn Data_Time_Duration_ordHours() -> &dyn Any {
        static Data_Time_Duration_ordHours: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_ordHours.get_or_init(||
                                                    &PureScript_Data_Ord::Data_Ord_ordNumber())
    }
    pub fn Data_Time_Duration_ordDays() -> &dyn Any {
        static Data_Time_Duration_ordDays: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_ordDays.get_or_init(||
                                                   &PureScript_Data_Ord::Data_Ord_ordNumber())
    }
    pub fn Data_Time_Duration_newtypeSeconds() -> &dyn Any {
        static Data_Time_Duration_newtypeSeconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_newtypeSeconds.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                           &&&add(string("Coercible0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &Sharpurs_Prelude::Prim_undefined()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Time_Duration_newtypeMinutes() -> &dyn Any {
        static Data_Time_Duration_newtypeMinutes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_newtypeMinutes.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                           &&&add(string("Coercible0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &Sharpurs_Prelude::Prim_undefined()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Time_Duration_newtypeMilliseconds() -> &dyn Any {
        static Data_Time_Duration_newtypeMilliseconds:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_newtypeMilliseconds.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                &&&add(string("Coercible0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &Sharpurs_Prelude::Prim_undefined()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>())))
    }
    pub fn Data_Time_Duration_newtypeHours() -> &dyn Any {
        static Data_Time_Duration_newtypeHours: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_newtypeHours.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                         &&&add(string("Coercible0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &Sharpurs_Prelude::Prim_undefined()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Time_Duration_newtypeDays() -> &dyn Any {
        static Data_Time_Duration_newtypeDays: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_newtypeDays.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                        &&&add(string("Coercible0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &Sharpurs_Prelude::Prim_undefined()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))
    }
    pub fn Data_Time_Duration_monoidSeconds() -> &dyn Any {
        static Data_Time_Duration_monoidSeconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_monoidSeconds.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                          &&&add(string("mempty"),
                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Seconds(),
                                                                                                                                   &&&0.0_f64),
                                                                                                 add(string("Semigroup0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Data_Time_Duration::Data_Time_Duration_semigroupSeconds()),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Time_Duration_monoidMinutes() -> &dyn Any {
        static Data_Time_Duration_monoidMinutes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_monoidMinutes.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                          &&&add(string("mempty"),
                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Minutes(),
                                                                                                                                   &&&0.0_f64),
                                                                                                 add(string("Semigroup0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Data_Time_Duration::Data_Time_Duration_semigroupMinutes()),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Time_Duration_monoidMilliseconds() -> &dyn Any {
        static Data_Time_Duration_monoidMilliseconds:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_monoidMilliseconds.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                               &&&add(string("mempty"),
                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds(),
                                                                                                                                        &&&0.0_f64),
                                                                                                      add(string("Semigroup0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Time_Duration::Data_Time_Duration_semigroupMilliseconds()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Time_Duration_monoidHours() -> &dyn Any {
        static Data_Time_Duration_monoidHours: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_monoidHours.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                        &&&add(string("mempty"),
                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Hours(),
                                                                                                                                 &&&0.0_f64),
                                                                                               add(string("Semigroup0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_Time_Duration::Data_Time_Duration_semigroupHours()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Time_Duration_monoidDays() -> &dyn Any {
        static Data_Time_Duration_monoidDays: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_monoidDays.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                       &&&add(string("mempty"),
                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Days(),
                                                                                                                                &&&0.0_f64),
                                                                                              add(string("Semigroup0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_Time_Duration::Data_Time_Duration_semigroupDays()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Time_Duration_fromDuration() -> &dyn Any {
        static Data_Time_Duration_fromDuration: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_fromDuration.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("fromDuration"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Time_Duration_negateDuration() -> &dyn Any {
        static Data_Time_Duration_negateDuration: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_negateDuration.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictDuration|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_toDuration(),
                                                                                                                                                                                 dictDuration)),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                       &&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds()),
                                                                                                                                                                                                                    &&&PureScript_Data_Time_Duration::Data_Time_Duration_negate())),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_fromDuration(),
                                                                                                                                                                                 dictDuration)))))
    }
    pub fn Data_Time_Duration_eqSeconds() -> &dyn Any {
        static Data_Time_Duration_eqSeconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_eqSeconds.get_or_init(||
                                                     &PureScript_Data_Eq::Data_Eq_eqNumber())
    }
    pub fn Data_Time_Duration_eqMinutes() -> &dyn Any {
        static Data_Time_Duration_eqMinutes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_eqMinutes.get_or_init(||
                                                     &PureScript_Data_Eq::Data_Eq_eqNumber())
    }
    pub fn Data_Time_Duration_eqMilliseconds() -> &dyn Any {
        static Data_Time_Duration_eqMilliseconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_eqMilliseconds.get_or_init(||
                                                          &PureScript_Data_Eq::Data_Eq_eqNumber())
    }
    pub fn Data_Time_Duration_eqHours() -> &dyn Any {
        static Data_Time_Duration_eqHours: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_eqHours.get_or_init(||
                                                   &PureScript_Data_Eq::Data_Eq_eqNumber())
    }
    pub fn Data_Time_Duration_eqDays() -> &dyn Any {
        static Data_Time_Duration_eqDays: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_eqDays.get_or_init(||
                                                  &PureScript_Data_Eq::Data_Eq_eqNumber())
    }
    pub fn Data_Time_Duration_durationSeconds() -> &dyn Any {
        static Data_Time_Duration_durationSeconds: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_durationSeconds.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Durationusd_Dict(),
                                                                                            &&&add(string("fromDuration"),
                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                        &&&PureScript_Data_Time_Duration::Data_Time_Duration_Seconds()),
                                                                                                                                     &&&Func1::new(move
                                                                                                                                                       |v|
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                           v),
                                                                                                                                                                                        &&&1000.0_f64))),
                                                                                                   add(string("toDuration"),
                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                            &&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds()),
                                                                                                                                         &&&Func1::new(move
                                                                                                                                                           |v_1|
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingNumber()),
                                                                                                                                                                                                                               v_1),
                                                                                                                                                                                            &&&1000.0_f64))),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Time_Duration_durationMinutes() -> &dyn Any {
        static Data_Time_Duration_durationMinutes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_durationMinutes.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Durationusd_Dict(),
                                                                                            &&&add(string("fromDuration"),
                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                        &&&PureScript_Data_Time_Duration::Data_Time_Duration_Minutes()),
                                                                                                                                     &&&Func1::new(move
                                                                                                                                                       |v|
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                           v),
                                                                                                                                                                                        &&&60000.0_f64))),
                                                                                                   add(string("toDuration"),
                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                            &&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds()),
                                                                                                                                         &&&Func1::new(move
                                                                                                                                                           |v_1|
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingNumber()),
                                                                                                                                                                                                                               v_1),
                                                                                                                                                                                            &&&60000.0_f64))),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Time_Duration_durationMilliseconds() -> &dyn Any {
        static Data_Time_Duration_durationMilliseconds:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_durationMilliseconds.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Durationusd_Dict(),
                                                                                                 &&&add(string("fromDuration"),
                                                                                                        &&PureScript_Data_Time_Duration::Data_Time_Duration_identity(),
                                                                                                        add(string("toDuration"),
                                                                                                            &&PureScript_Data_Time_Duration::Data_Time_Duration_identity(),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Time_Duration_durationHours() -> &dyn Any {
        static Data_Time_Duration_durationHours: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_durationHours.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Durationusd_Dict(),
                                                                                          &&&add(string("fromDuration"),
                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                      &&&PureScript_Data_Time_Duration::Data_Time_Duration_Hours()),
                                                                                                                                   &&&Func1::new(move
                                                                                                                                                     |v|
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                            &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                         v),
                                                                                                                                                                                      &&&3600000.0_f64))),
                                                                                                 add(string("toDuration"),
                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                          &&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds()),
                                                                                                                                       &&&Func1::new(move
                                                                                                                                                         |v_1|
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                                                &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingNumber()),
                                                                                                                                                                                                                             v_1),
                                                                                                                                                                                          &&&3600000.0_f64))),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Time_Duration_durationDays() -> &dyn Any {
        static Data_Time_Duration_durationDays: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_durationDays.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Durationusd_Dict(),
                                                                                         &&&add(string("fromDuration"),
                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                     &&&PureScript_Data_Time_Duration::Data_Time_Duration_Days()),
                                                                                                                                  &&&Func1::new(move
                                                                                                                                                    |v|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                           &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                        v),
                                                                                                                                                                                     &&&86400000.0_f64))),
                                                                                                add(string("toDuration"),
                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_over(),
                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                         &&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds()),
                                                                                                                                      &&&Func1::new(move
                                                                                                                                                        |v_1|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingNumber()),
                                                                                                                                                                                                                            v_1),
                                                                                                                                                                                         &&&86400000.0_f64))),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Time_Duration_convertDuration() -> &dyn Any {
        static Data_Time_Duration_convertDuration: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Duration_convertDuration.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictDuration|
                                                                           {
                                                                               let fromDuration1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_fromDuration(),
                                                                                                                    dictDuration);
                                                                               &Func1::new({
                                                                                               let fromDuration1
                                                                                                   =
                                                                                                   fromDuration1.clone();
                                                                                               move
                                                                                                   |dictDuration1|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_toDuration(),
                                                                                                                                                                                                          dictDuration1)),
                                                                                                                                    &&&fromDuration1)
                                                                                           })
                                                                           }))
    }
}
