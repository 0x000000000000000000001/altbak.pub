pub mod PureScript_Control_Monad_Writer {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_2b2441be::PureScript_Control_Monad_Writer_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Writer_unwrap() -> &dyn Any {
        static Control_Monad_Writer_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_unwrap.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Monad_Writer_writer() -> &dyn Any {
        static Control_Monad_Writer_writer: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_writer.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                        &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_WriterT()),
                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                        &&&PureScript_Data_Identity::Data_Identity_applicativeIdentity())))
    }
    pub fn Control_Monad_Writer_runWriter() -> &dyn Any {
        static Control_Monad_Writer_runWriter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_runWriter.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                        &&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_runWriterT()))
    }
    pub fn Control_Monad_Writer_mapWriter() -> &dyn Any {
        static Control_Monad_Writer_mapWriter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_mapWriter.get_or_init(||
                                                       &Func1::new(move |f|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Trans::Control_Monad_Writer_Trans_mapWriterT(),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                              &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                 f),
                                                                                                                                                                              &&&PureScript_Control_Monad_Writer::Control_Monad_Writer_unwrap())))))
    }
    pub fn Control_Monad_Writer_execWriter() -> &dyn Any {
        static Control_Monad_Writer_execWriter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_execWriter.get_or_init(||
                                                        &Func1::new(move |m|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_snd(),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer::Control_Monad_Writer_runWriter(),
                                                                                                                                            m))))
    }
}
