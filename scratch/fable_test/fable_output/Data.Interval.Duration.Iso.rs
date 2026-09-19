pub mod PureScript_Data_Interval_Duration_Iso {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::NativeArray_::new_array;
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_8669b9ba::PureScript_Data_Interval_Duration;
    use crate::module_8669b9ba::PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent;
    use crate::module_eb82fa3::PureScript_Data_List_NonEmpty;
    use crate::module_d662adf2::PureScript_Data_List_Types;
    use crate::module_d662adf2::PureScript_Data_List_Types::Data_List_Types_List;
    use crate::module_843b47b7::PureScript_Data_List;
    use crate::module_ed2bf3e0::PureScript_Data_Map_Internal;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_3b72fe33::PureScript_Data_Monoid_Additive;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_91782ab6::PureScript_Data_Number;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    #[derive(Clone, Debug,)]
    pub enum Data_Interval_Duration_Iso_Error {
        Data_Interval_Duration_Iso_IsEmptyusd_Ctor,
        Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor,
        Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(&dyn Any),
        Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(&dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Interval_Duration_Iso_empty() -> &dyn Any {
        static Data_Interval_Duration_Iso_empty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_empty.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                          &&&PureScript_Data_List_Types::Data_List_Types_plusList()))
    }
    pub fn Data_Interval_Duration_Iso_foldMap() -> &dyn Any {
        static Data_Interval_Duration_Iso_foldMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_foldMap.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                               &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                            &&&PureScript_Data_List_Types::Data_List_Types_monoidList()))
    }
    pub fn Data_Interval_Duration_Iso_monoidAdditive() -> &dyn Any {
        static Data_Interval_Duration_Iso_monoidAdditive:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_monoidAdditive.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_monoidAdditive(),
                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()))
    }
    pub fn Data_Interval_Duration_Iso_heytingAlgebraFunction() -> &dyn Any {
        static Data_Interval_Duration_Iso_heytingAlgebraFunction:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_heytingAlgebraFunction.get_or_init(||
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraFunction(),
                                                                                                           &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()))
    }
    pub fn Data_Interval_Duration_Iso_monoidFn() -> &dyn Any {
        static Data_Interval_Duration_Iso_monoidFn: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Interval_Duration_Iso_monoidFn.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_monoidFn(),
                                                                                             &&&PureScript_Data_List_Types::Data_List_Types_monoidList()))
    }
    pub fn Data_Interval_Duration_Iso_IsoDuration() -> &dyn Any {
        static Data_Interval_Duration_Iso_IsoDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_IsoDuration.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               x.clone()))
    }
    pub fn Data_Interval_Duration_Iso_IsEmpty() -> &dyn Any {
        static Data_Interval_Duration_Iso_IsEmpty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_IsEmpty.get_or_init(||
                                                           &LrcPtr::new(PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_IsEmptyusd_Ctor))
    }
    pub fn Data_Interval_Duration_Iso_InvalidWeekComponentUsage()
     -> &dyn Any {
        static Data_Interval_Duration_Iso_InvalidWeekComponentUsage:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_InvalidWeekComponentUsage.get_or_init(||
                                                                             &LrcPtr::new(PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor))
    }
    pub fn Data_Interval_Duration_Iso_ContainsNegativeValue() -> &dyn Any {
        static Data_Interval_Duration_Iso_ContainsNegativeValue:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_ContainsNegativeValue.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |usd__arg1|
                                                                                         &LrcPtr::new(PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Interval_Duration_Iso_InvalidFractionalUse() -> &dyn Any {
        static Data_Interval_Duration_Iso_InvalidFractionalUse:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_InvalidFractionalUse.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |usd__arg1|
                                                                                        &LrcPtr::new(PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Interval_Duration_Iso_unIsoDuration() -> &dyn Any {
        static Data_Interval_Duration_Iso_unIsoDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_unIsoDuration.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |v|
                                                                                 &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_Interval_Duration_Iso_showIsoDuration() -> &dyn Any {
        static Data_Interval_Duration_Iso_showIsoDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_showIsoDuration.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                    &&&add(string("show"),
                                                                                                           &&Func1::new(move
                                                                                                                            |v|
                                                                                                                            {
                                                                                                                                let d =
                                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                       &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                    &&&string("(IsoDuration ")),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_showDuration()),
                                                                                                                                                                                                                                                                          &&&d)),
                                                                                                                                                                                                    &&&string(")")))
                                                                                                                            }),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
    }
    pub fn Data_Interval_Duration_Iso_showError() -> &dyn Any {
        static Data_Interval_Duration_Iso_showError: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Interval_Duration_Iso_showError.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                              &&&add(string("show"),
                                                                                                     &&Func1::new(move
                                                                                                                      |v|
                                                                                                                      {
                                                                                                                          let matchValue:
                                                                                                                                  LrcPtr<PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error> =
                                                                                                                              Sharpurs_Prelude::unbox(v);
                                                                                                                          match matchValue.as_ref()
                                                                                                                              {
                                                                                                                              PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                                                              =>
                                                                                                                              &string("(InvalidWeekComponentUsage)"),
                                                                                                                              PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_2_0)
                                                                                                                              =>
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                  &&&string("(ContainsNegativeValue ")),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_showDurationComponent()),
                                                                                                                                                                                                                                                                        &&matchValue_2_0)),
                                                                                                                                                                                                  &&&string(")"))),
                                                                                                                              PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(matchValue_3_0)
                                                                                                                              =>
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                  &&&string("(InvalidFractionalUse ")),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_showDurationComponent()),
                                                                                                                                                                                                                                                                        &&matchValue_3_0)),
                                                                                                                                                                                                  &&&string(")"))),
                                                                                                                              _
                                                                                                                              =>
                                                                                                                              &string("(IsEmpty)"),
                                                                                                                          }
                                                                                                                      }),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_Interval_Duration_Iso_prettyError() -> &dyn Any {
        static Data_Interval_Duration_Iso_prettyError:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_prettyError.get_or_init(||
                                                               &Func1::new(move
                                                                               |v|
                                                                               {
                                                                                   let matchValue:
                                                                                           LrcPtr<PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error> =
                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                   match matchValue.as_ref()
                                                                                       {
                                                                                       PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                       =>
                                                                                       &string("Week component of Duration is used with other components"),
                                                                                       PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_2_0)
                                                                                       =>
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                           &&&string("Component `")),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                    &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_showDurationComponent()),
                                                                                                                                                                                                                                 &&matchValue_2_0)),
                                                                                                                                                           &&&string("` contains negative value"))),
                                                                                       PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(matchValue_3_0)
                                                                                       =>
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                           &&&string("Invalid usage of Fractional value at component `")),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                    &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_showDurationComponent()),
                                                                                                                                                                                                                                 &&matchValue_3_0)),
                                                                                                                                                           &&&string("`"))),
                                                                                       _
                                                                                       =>
                                                                                       &string("Duration is empty (has no components)"),
                                                                                   }
                                                                               }))
    }
    pub fn Data_Interval_Duration_Iso_eqIsoDuration() -> &dyn Any {
        static Data_Interval_Duration_Iso_eqIsoDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_eqIsoDuration.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                  &&&add(string("eq"),
                                                                                                         &&Func1::new(move
                                                                                                                          |x|
                                                                                                                          &Func1::new({
                                                                                                                                          let x
                                                                                                                                              =
                                                                                                                                              x.clone();
                                                                                                                                          move
                                                                                                                                              |y|
                                                                                                                                              {
                                                                                                                                                  let matchValue =
                                                                                                                                                      Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                  let matchValue_1 =
                                                                                                                                                      Sharpurs_Prelude::unbox(y);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                         &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_eqDuration()),
                                                                                                                                                                                                                      &&&matchValue),
                                                                                                                                                                                   &&&matchValue_1)
                                                                                                                                              }
                                                                                                                                      })),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_Interval_Duration_Iso_ordIsoDuration() -> &dyn Any {
        static Data_Interval_Duration_Iso_ordIsoDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_ordIsoDuration.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                   &&&add(string("compare"),
                                                                                                          &&Func1::new(move
                                                                                                                           |x|
                                                                                                                           &Func1::new({
                                                                                                                                           let x
                                                                                                                                               =
                                                                                                                                               x.clone();
                                                                                                                                           move
                                                                                                                                               |y|
                                                                                                                                               {
                                                                                                                                                   let matchValue =
                                                                                                                                                       Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                   let matchValue_1 =
                                                                                                                                                       Sharpurs_Prelude::unbox(y);
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                          &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_ordDuration()),
                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                    &&&matchValue_1)
                                                                                                                                               }
                                                                                                                                       })),
                                                                                                          add(string("Eq0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_eqIsoDuration()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Interval_Duration_Iso_eqError() -> &dyn Any {
        static Data_Interval_Duration_Iso_eqError: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_eqError.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                            &&&add(string("eq"),
                                                                                                   &&Func1::new(move
                                                                                                                    |x|
                                                                                                                    &Func1::new({
                                                                                                                                    let x
                                                                                                                                        =
                                                                                                                                        x.clone();
                                                                                                                                    move
                                                                                                                                        |y|
                                                                                                                                        {
                                                                                                                                            let matchValue:
                                                                                                                                                    LrcPtr<PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error> =
                                                                                                                                                Sharpurs_Prelude::unbox(&&x);
                                                                                                                                            let matchValue_1:
                                                                                                                                                    LrcPtr<PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error> =
                                                                                                                                                Sharpurs_Prelude::unbox(y);
                                                                                                                                            if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                                                                                   =
                                                                                                                                                   matchValue.as_ref()
                                                                                                                                               {
                                                                                                                                                if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                                                                                       =
                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                   {
                                                                                                                                                    &true
                                                                                                                                                } else {
                                                                                                                                                    &false
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_2_0)
                                                                                                                                                       =
                                                                                                                                                       matchValue.as_ref()
                                                                                                                                                   {
                                                                                                                                                    if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_1_2_0)
                                                                                                                                                           =
                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                       {
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_eqDurationComponent()),
                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                   PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(x)
                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                               }),
                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(x)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                            })
                                                                                                                                                    } else {
                                                                                                                                                        &false
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(matchValue_3_0)
                                                                                                                                                           =
                                                                                                                                                           matchValue.as_ref()
                                                                                                                                                       {
                                                                                                                                                        if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(matchValue_1_3_0)
                                                                                                                                                               =
                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                           {
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_eqDurationComponent()),
                                                                                                                                                                                                                                &&&match matchValue.as_ref()
                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                       PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(x)
                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                       _
                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                       unreachable!(),
                                                                                                                                                                                                                                   }),
                                                                                                                                                                                             &&&match matchValue_1.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(x)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                })
                                                                                                                                                        } else {
                                                                                                                                                            &false
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_IsEmptyusd_Ctor
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
                                                                                                                                        }
                                                                                                                                })),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>())))
    }
    pub fn Data_Interval_Duration_Iso_ordError() -> &dyn Any {
        static Data_Interval_Duration_Iso_ordError: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Interval_Duration_Iso_ordError.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                             &&&add(string("compare"),
                                                                                                    &&Func1::new(move
                                                                                                                     |x|
                                                                                                                     &Func1::new({
                                                                                                                                     let x
                                                                                                                                         =
                                                                                                                                         x.clone();
                                                                                                                                     move
                                                                                                                                         |y|
                                                                                                                                         {
                                                                                                                                             let matchValue:
                                                                                                                                                     LrcPtr<PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error> =
                                                                                                                                                 Sharpurs_Prelude::unbox(&&x);
                                                                                                                                             let matchValue_1:
                                                                                                                                                     LrcPtr<PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error> =
                                                                                                                                                 Sharpurs_Prelude::unbox(y);
                                                                                                                                             if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                                                                                    =
                                                                                                                                                    matchValue.as_ref()
                                                                                                                                                {
                                                                                                                                                 if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_IsEmptyusd_Ctor
                                                                                                                                                        =
                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                    {
                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                 } else {
                                                                                                                                                     if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                                                                                            =
                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                        {
                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                     } else {
                                                                                                                                                         if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_1_2_0)
                                                                                                                                                                =
                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                            {
                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                         } else {
                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                         }
                                                                                                                                                     }
                                                                                                                                                 }
                                                                                                                                             } else {
                                                                                                                                                 if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_2_0)
                                                                                                                                                        =
                                                                                                                                                        matchValue.as_ref()
                                                                                                                                                    {
                                                                                                                                                     if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_IsEmptyusd_Ctor
                                                                                                                                                            =
                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                        {
                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                     } else {
                                                                                                                                                         if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                            {
                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                         } else {
                                                                                                                                                             if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_1_2_0)
                                                                                                                                                                    =
                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                {
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_ordDurationComponent()),
                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(x)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                            _
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(x)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                         _
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                     })
                                                                                                                                                             } else {
                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                     }
                                                                                                                                                 } else {
                                                                                                                                                     if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(matchValue_3_0)
                                                                                                                                                            =
                                                                                                                                                            matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                         if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                            {
                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                         } else {
                                                                                                                                                             if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_1_2_0)
                                                                                                                                                                    =
                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                {
                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                             } else {
                                                                                                                                                                 if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(matchValue_1_3_0)
                                                                                                                                                                        =
                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                            &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_ordDurationComponent()),
                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(x)
                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                      &&&match matchValue_1.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(x)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                             _
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                         })
                                                                                                                                                                 } else {
                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                 }
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                     } else {
                                                                                                                                                         if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_IsEmptyusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                            {
                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                         } else {
                                                                                                                                                             if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor
                                                                                                                                                                    =
                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                {
                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                             } else {
                                                                                                                                                                 if let PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(matchValue_1_2_0)
                                                                                                                                                                        =
                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                 } else {
                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                 }
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                     }
                                                                                                                                                 }
                                                                                                                                             }
                                                                                                                                         }
                                                                                                                                 })),
                                                                                                    add(string("Eq0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_eqError()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))))
    }
    pub fn Data_Interval_Duration_Iso_checkWeekUsage() -> &dyn Any {
        static Data_Interval_Duration_Iso_checkWeekUsage:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_checkWeekUsage.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |v|
                                                                                  {
                                                                                      let matchValue =
                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                      {
                                                                                          let activePatternResult =
                                                                                              Sharpurs_Prelude::_007cHasProp_007c__007c(string("asMap"),
                                                                                                                                        &matchValue);
                                                                                          if activePatternResult.is_some()
                                                                                             {
                                                                                              let asMap =
                                                                                                  getValue(activePatternResult);
                                                                                              let matchValue_1 =
                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                  &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isJust(),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_lookup(),
                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_ordDurationComponent()),
                                                                                                                                                                                                                                                                                                        &&&LrcPtr::new(Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor)),
                                                                                                                                                                                                                                                                     &&&asMap))),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_size(),
                                                                                                                                                                                                                                                                     &&&asMap)),
                                                                                                                                                                                               &&&1_i32)));
                                                                                              match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                               &matchValue_1)
                                                                                                  {
                                                                                                  0_i32
                                                                                                  =>
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                      &&&PureScript_Data_List_Types::Data_List_Types_applicativeList()),
                                                                                                                                   &&&LrcPtr::new(PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidWeekComponentUsageusd_Ctor)),
                                                                                                  _
                                                                                                  =>
                                                                                                  &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_empty(),
                                                                                              }
                                                                                          } else {
                                                                                              panic!("{}",
                                                                                                     LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Interval.Duration.Iso.fs"),
                                  Data1: 49_i32,
                                  Data2: 78_i32,}).get_Message(),)
                                                                                          }
                                                                                      }
                                                                                  }))
    }
    pub fn Data_Interval_Duration_Iso_checkNegativeValues() -> &dyn Any {
        static Data_Interval_Duration_Iso_checkNegativeValues:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_checkNegativeValues.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |v|
                                                                                       {
                                                                                           let matchValue =
                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                           {
                                                                                               let activePatternResult =
                                                                                                   Sharpurs_Prelude::_007cHasProp_007c__007c(string("asList"),
                                                                                                                                             &matchValue);
                                                                                               if activePatternResult.is_some()
                                                                                                  {
                                                                                                   let asList =
                                                                                                       getValue(activePatternResult);
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                          &&&PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_foldMap()),
                                                                                                                                                                       &&&asList),
                                                                                                                                    &&&Func1::new(move
                                                                                                                                                      |v1|
                                                                                                                                                      {
                                                                                                                                                          let matchValue_1:
                                                                                                                                                                  LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                              Sharpurs_Prelude::unbox(v1);
                                                                                                                                                          let matchValue_2 =
                                                                                                                                                              Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                             x)
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                        &&&0.0_f64));
                                                                                                                                                          match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                           &matchValue_2)
                                                                                                                                                              {
                                                                                                                                                              0_i32
                                                                                                                                                              =>
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                                                                                               &&&PureScript_Data_List_Types::Data_List_Types_plusList()),
                                                                                                                                                              _
                                                                                                                                                              =>
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                  &&&PureScript_Data_List_Types::Data_List_Types_applicativeList()),
                                                                                                                                                                                               &&&LrcPtr::new(PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_ContainsNegativeValueusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                 }))),
                                                                                                                                                          }
                                                                                                                                                      }))
                                                                                               } else {
                                                                                                   panic!("{}",
                                                                                                          LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Interval.Duration.Iso.fs"),
                                  Data1: 51_i32,
                                  Data2: 83_i32,}).get_Message(),)
                                                                                               }
                                                                                           }
                                                                                       }))
    }
    pub fn Data_Interval_Duration_Iso_checkFractionalUse() -> &dyn Any {
        static Data_Interval_Duration_Iso_checkFractionalUse:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_checkFractionalUse.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |v|
                                                                                      {
                                                                                          let matchValue =
                                                                                              Sharpurs_Prelude::unbox(v);
                                                                                          {
                                                                                              let activePatternResult =
                                                                                                  Sharpurs_Prelude::_007cHasProp_007c__007c(string("asList"),
                                                                                                                                            &matchValue);
                                                                                              if activePatternResult.is_some()
                                                                                                 {
                                                                                                  let asList =
                                                                                                      getValue(activePatternResult);
                                                                                                  let isFractional =
                                                                                                      &Func1::new(move
                                                                                                                      |a|
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_notEq(),
                                                                                                                                                                                                                             &&&PureScript_Data_Eq::Data_Eq_eqNumber()),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number::Data_Number_floor(),
                                                                                                                                                                                                                             a)),
                                                                                                                                                       a));
                                                                                                  let checkRest =
                                                                                                      &Func1::new(move
                                                                                                                      |rest|
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                             &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_monoidAdditive()),
                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Number::Data_Number_abs()),
                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive()))),
                                                                                                                                                                                                                                                                rest))),
                                                                                                                                                       &&&0.0_f64));
                                                                                                  let matchValue_1:
                                                                                                          LrcPtr<Data_List_Types_List> =
                                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                   |v2|
                                                                                                                                                                                   find(string("rest"),
                                                                                                                                                                                        Sharpurs_Prelude::unbox(v2))),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_span(),
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_heytingAlgebraFunction()),
                                                                                                                                                                                                                                                                                                             &&&isFractional))),
                                                                                                                                                                                                    &&&asList)));
                                                                                                  if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                            matchValue_1_1_1)
                                                                                                         =
                                                                                                         matchValue_1.as_ref()
                                                                                                     {
                                                                                                      let activePatternResult_1:
                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                          Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                                 {
                                                                                                                                                 Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                    _)
                                                                                                                                                 =>
                                                                                                                                                 x.clone(),
                                                                                                                                                 _
                                                                                                                                                 =>
                                                                                                                                                 unreachable!(),
                                                                                                                                             });
                                                                                                      if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&checkRest,
                                                                                                                                                                   &&&match matchValue_1.as_ref()
                                                                                                                                                                          {
                                                                                                                                                                          Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                             x)
                                                                                                                                                                          =>
                                                                                                                                                                          x.clone(),
                                                                                                                                                                          _
                                                                                                                                                                          =>
                                                                                                                                                                          unreachable!(),
                                                                                                                                                                      }))
                                                                                                         {
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                              &&&PureScript_Data_List_Types::Data_List_Types_applicativeList()),
                                                                                                                                           &&&LrcPtr::new(PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_InvalidFractionalUseusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                            })))
                                                                                                      } else {
                                                                                                          &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_empty()
                                                                                                      }
                                                                                                  } else {
                                                                                                      &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_empty()
                                                                                                  }
                                                                                              } else {
                                                                                                  panic!("{}",
                                                                                                         LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Interval.Duration.Iso.fs"),
                                  Data1: 53_i32,
                                  Data2: 82_i32,}).get_Message(),)
                                                                                              }
                                                                                          }
                                                                                      }))
    }
    pub fn Data_Interval_Duration_Iso_checkEmptiness() -> &dyn Any {
        static Data_Interval_Duration_Iso_checkEmptiness:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_checkEmptiness.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |v|
                                                                                  {
                                                                                      let matchValue =
                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                      {
                                                                                          let activePatternResult =
                                                                                              Sharpurs_Prelude::_007cHasProp_007c__007c(string("asList"),
                                                                                                                                        &matchValue);
                                                                                          if activePatternResult.is_some()
                                                                                             {
                                                                                              let asList =
                                                                                                  getValue(activePatternResult);
                                                                                              let matchValue_1 =
                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_null(),
                                                                                                                                                            &&&asList));
                                                                                              match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                               &matchValue_1)
                                                                                                  {
                                                                                                  0_i32
                                                                                                  =>
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                      &&&PureScript_Data_List_Types::Data_List_Types_applicativeList()),
                                                                                                                                   &&&LrcPtr::new(PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_Error::Data_Interval_Duration_Iso_IsEmptyusd_Ctor)),
                                                                                                  _
                                                                                                  =>
                                                                                                  &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_empty(),
                                                                                              }
                                                                                          } else {
                                                                                              panic!("{}",
                                                                                                     LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Interval.Duration.Iso.fs"),
                                  Data1: 55_i32,
                                  Data2: 78_i32,}).get_Message(),)
                                                                                          }
                                                                                      }
                                                                                  }))
    }
    pub fn Data_Interval_Duration_Iso_checkValidIsoDuration() -> &dyn Any {
        static Data_Interval_Duration_Iso_checkValidIsoDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_checkValidIsoDuration.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |v|
                                                                                         {
                                                                                             let asMap =
                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_fold(),
                                                                                                                                                                                                                                        &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()),
                                                                                                                                                                                                     &&&PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_monoidFn()),
                                                                                                                                                                  &&&new_array(&[&PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_checkWeekUsage(),
                                                                                                                                                                                 &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_checkEmptiness(),
                                                                                                                                                                                 &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_checkFractionalUse(),
                                                                                                                                                                                 &PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_checkNegativeValues()])),
                                                                                                                              &&&add(string("asList"),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_toUnfoldable(),
                                                                                                                                                                                                                                              &&&PureScript_Data_List_Types::Data_List_Types_unfoldableList()),
                                                                                                                                                                                                           &&&asMap)),
                                                                                                                                     add(string("asMap"),
                                                                                                                                         &&asMap,
                                                                                                                                         empty::<string,
                                                                                                                                                 &dyn Any>())))
                                                                                         }))
    }
    pub fn Data_Interval_Duration_Iso_mkIsoDuration() -> &dyn Any {
        static Data_Interval_Duration_Iso_mkIsoDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Iso_mkIsoDuration.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |d|
                                                                                 {
                                                                                     let matchValue:
                                                                                             LrcPtr<Data_Maybe_Maybe> =
                                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_fromList(),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_checkValidIsoDuration(),
                                                                                                                                                                                       d)));
                                                                                     match matchValue.as_ref()
                                                                                         {
                                                                                         Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                         =>
                                                                                         &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration_Iso::Data_Interval_Duration_Iso_IsoDuration(),
                                                                                                                                                                                     d))),
                                                                                         Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                         =>
                                                                                         &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                        =>
                                                                                                                                                        x.clone(),
                                                                                                                                                        _
                                                                                                                                                        =>
                                                                                                                                                        unreachable!(),
                                                                                                                                                    })),
                                                                                     }
                                                                                 }))
    }
}
