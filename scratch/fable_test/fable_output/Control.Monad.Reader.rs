pub mod PureScript_Control_Monad_Reader {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_8ea61114::PureScript_Control_Monad_Reader_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Reader_unwrap() -> &dyn Any {
        static Control_Monad_Reader_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_unwrap.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Monad_Reader_withReader() -> &dyn Any {
        static Control_Monad_Reader_withReader: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_withReader.get_or_init(||
                                                        &PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_withReaderT())
    }
    pub fn Control_Monad_Reader_runReader() -> &dyn Any {
        static Control_Monad_Reader_runReader: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_runReader.get_or_init(||
                                                       &Func1::new(move |v|
                                                                       {
                                                                           let m =
                                                                               Sharpurs_Prelude::unbox(v);
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                               &&&PureScript_Control_Monad_Reader::Control_Monad_Reader_unwrap()),
                                                                                                            &&&m)
                                                                       }))
    }
    pub fn Control_Monad_Reader_mapReader() -> &dyn Any {
        static Control_Monad_Reader_mapReader: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_mapReader.get_or_init(||
                                                       &Func1::new(move |f|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                           &&&PureScript_Control_Monad_Reader_Trans::Control_Monad_Reader_Trans_mapReaderT()),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                              &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                 f),
                                                                                                                                                                              &&&PureScript_Control_Monad_Reader::Control_Monad_Reader_unwrap())))))
    }
}
