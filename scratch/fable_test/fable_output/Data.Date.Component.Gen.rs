pub mod PureScript_Data_Date_Component_Gen {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_851afbc9::PureScript_Control_Monad_Gen_Class;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_b6aac8e0::PureScript_Data_Date_Component;
    use crate::module_25dacb64::PureScript_Data_Enum_Gen;
    use crate::module_6a1c5ce6::PureScript_Data_Enum;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Date_Component_Gen_toEnum() -> &dyn Any {
        static Data_Date_Component_Gen_toEnum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Gen_toEnum.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                        &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumYear()))
    }
    pub fn Data_Date_Component_Gen_genYear() -> &dyn Any {
        static Data_Date_Component_Gen_genYear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Gen_genYear.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictMonadGen|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                       |usd__unused|
                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined())))),
                                                                                                                                                                               &&&PureScript_Data_Date_Component_Gen::Data_Date_Component_Gen_toEnum())),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseInt(),
                                                                                                                                                                                                                  dictMonadGen),
                                                                                                                                                                               &&&1900_i32),
                                                                                                                                            &&&2100_i32))))
    }
    pub fn Data_Date_Component_Gen_genWeekday() -> &dyn Any {
        static Data_Date_Component_Gen_genWeekday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Gen_genWeekday.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictMonadGen|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum_Gen::Data_Enum_Gen_genBoundedEnum(),
                                                                                                                                               dictMonadGen),
                                                                                                            &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumWeekday())))
    }
    pub fn Data_Date_Component_Gen_genMonth() -> &dyn Any {
        static Data_Date_Component_Gen_genMonth: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Gen_genMonth.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonadGen|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum_Gen::Data_Enum_Gen_genBoundedEnum(),
                                                                                                                                             dictMonadGen),
                                                                                                          &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth())))
    }
    pub fn Data_Date_Component_Gen_genDay() -> &dyn Any {
        static Data_Date_Component_Gen_genDay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Gen_genDay.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadGen|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum_Gen::Data_Enum_Gen_genBoundedEnum(),
                                                                                                                                           dictMonadGen),
                                                                                                        &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay())))
    }
}
