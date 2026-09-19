pub mod PureScript_Control_MonadPlus {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_MonadPlus_MonadPlususd_Dict() -> &dyn Any {
        static Control_MonadPlus_MonadPlususd_Dict: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_MonadPlus_MonadPlususd_Dict.get_or_init(||
                                                            &Func1::new(move
                                                                            |x|
                                                                            x.clone()))
    }
    pub fn Control_MonadPlus_monadPlusArray() -> &dyn Any {
        static Control_MonadPlus_monadPlusArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_MonadPlus_monadPlusArray.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                          &&&add(string("Monad0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Control_Monad::Control_Monad_monadArray()),
                                                                                                 add(string("Alternative1"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused_1|
                                                                                                                      &PureScript_Control_Alternative::Control_Alternative_alternativeArray()),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
}
