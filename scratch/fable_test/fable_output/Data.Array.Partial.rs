pub mod PureScript_Data_Array_Partial {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_2d8e16c::PureScript_Data_Array;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Array_Partial_tail() -> &dyn Any {
        static Data_Array_Partial_tail: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_Partial_tail.get_or_init(||
                                                &Func1::new(move |usd__unused|
                                                                &Func1::new(move
                                                                                |xs|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                       &&&1_i32),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                       xs)),
                                                                                                                 xs))))
    }
    pub fn Data_Array_Partial_last() -> &dyn Any {
        static Data_Array_Partial_last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_Partial_last.get_or_init(||
                                                &Func1::new(move |usd__unused|
                                                                &Func1::new(move
                                                                                |xs|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unsafeIndex(),
                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                    xs),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                          &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                                          xs)),
                                                                                                                                                    &&&1_i32)))))
    }
    pub fn Data_Array_Partial_init() -> &dyn Any {
        static Data_Array_Partial_init: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_Partial_init.get_or_init(||
                                                &Func1::new(move |usd__unused|
                                                                &Func1::new(move
                                                                                |xs|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                       &&&0_i32),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                             &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                                                                             xs)),
                                                                                                                                                                                       &&&1_i32)),
                                                                                                                 xs))))
    }
    pub fn Data_Array_Partial_head() -> &dyn Any {
        static Data_Array_Partial_head: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_Partial_head.get_or_init(||
                                                &Func1::new(move |usd__unused|
                                                                &Func1::new(move
                                                                                |xs|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unsafeIndex(),
                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                    xs),
                                                                                                                 &&&0_i32))))
    }
}
