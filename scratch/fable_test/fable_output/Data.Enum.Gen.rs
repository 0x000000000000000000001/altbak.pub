pub mod PureScript_Data_Enum_Gen {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_e72de349::PureScript_Control_Monad_Gen;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_6a1c5ce6::PureScript_Data_Enum;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty::Data_NonEmpty_NonEmpty;
    use crate::module_b1754f14::PureScript_Data_Unfoldable1;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Enum_Gen_foldable1NonEmpty() -> &dyn Any {
        static Data_Enum_Gen_foldable1NonEmpty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_Gen_foldable1NonEmpty.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_foldable1NonEmpty(),
                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()))
    }
    pub fn Data_Enum_Gen_genBoundedEnum() -> &dyn Any {
        static Data_Enum_Gen_genBoundedEnum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_Gen_genBoundedEnum.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictMonadGen|
                                                                     {
                                                                         let Applicative0 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         &Func1::new({
                                                                                         let Applicative0
                                                                                             =
                                                                                             Applicative0.clone();
                                                                                         let dictMonadGen
                                                                                             =
                                                                                             dictMonadGen.clone();
                                                                                         move
                                                                                             |dictBoundedEnum|
                                                                                             {
                                                                                                 let Enum1 =
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                                                             Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                 let Bounded0 =
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Bounded0"),
                                                                                                                                             Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                 let matchValue:
                                                                                                         LrcPtr<Data_Maybe_Maybe> =
                                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                   &&&Enum1),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                   &&&Bounded0)));
                                                                                                 match matchValue.as_ref()
                                                                                                     {
                                                                                                     Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                     =>
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                         &&&Applicative0),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                         &&&Bounded0)),
                                                                                                     Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                     =>
                                                                                                     {
                                                                                                         let possibilities =
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_enumFromTo(),
                                                                                                                                                                                                                                                       &&&Enum1),
                                                                                                                                                                                                                    &&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldable1Array()),
                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                    }),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                 &&&Bounded0));
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_elements(),
                                                                                                                                                                                                                &&&dictMonadGen),
                                                                                                                                                                             &&&PureScript_Data_Enum_Gen::Data_Enum_Gen_foldable1NonEmpty()),
                                                                                                                                          &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                                                 &&&Bounded0),
                                                                                                                                                                                                                &possibilities)))
                                                                                                     }
                                                                                                 }
                                                                                             }
                                                                                     })
                                                                     }))
    }
}
