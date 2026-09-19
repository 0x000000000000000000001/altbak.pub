pub mod PureScript_Data_Time_Component_Gen {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_25dacb64::PureScript_Data_Enum_Gen;
    use crate::module_f62dae61::PureScript_Data_Time_Component;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Time_Component_Gen_genSecond() -> &dyn Any {
        static Data_Time_Component_Gen_genSecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_Gen_genSecond.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadGen|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum_Gen::Data_Enum_Gen_genBoundedEnum(),
                                                                                                                                              dictMonadGen),
                                                                                                           &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumSecond())))
    }
    pub fn Data_Time_Component_Gen_genMinute() -> &dyn Any {
        static Data_Time_Component_Gen_genMinute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_Gen_genMinute.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadGen|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum_Gen::Data_Enum_Gen_genBoundedEnum(),
                                                                                                                                              dictMonadGen),
                                                                                                           &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumMinute())))
    }
    pub fn Data_Time_Component_Gen_genMillisecond() -> &dyn Any {
        static Data_Time_Component_Gen_genMillisecond:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_Gen_genMillisecond.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictMonadGen|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum_Gen::Data_Enum_Gen_genBoundedEnum(),
                                                                                                                                                   dictMonadGen),
                                                                                                                &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumMillisecond())))
    }
    pub fn Data_Time_Component_Gen_genHour() -> &dyn Any {
        static Data_Time_Component_Gen_genHour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_Gen_genHour.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictMonadGen|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum_Gen::Data_Enum_Gen_genBoundedEnum(),
                                                                                                                                            dictMonadGen),
                                                                                                         &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumHour())))
    }
}
