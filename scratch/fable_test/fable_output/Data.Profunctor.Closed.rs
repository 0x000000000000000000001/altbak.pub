pub mod PureScript_Data_Profunctor_Closed {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Profunctor_Closed_Closedusd_Dict() -> &dyn Any {
        static Data_Profunctor_Closed_Closedusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Closed_Closedusd_Dict.get_or_init(||
                                                              &Func1::new(move
                                                                              |x|
                                                                              x.clone()))
    }
    pub fn Data_Profunctor_Closed_closedFunction() -> &dyn Any {
        static Data_Profunctor_Closed_closedFunction:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Closed_closedFunction.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Closed::Data_Profunctor_Closed_Closedusd_Dict(),
                                                                                               &&&add(string("closed"),
                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                      add(string("Profunctor0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Profunctor::Data_Profunctor_profunctorFn()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Profunctor_Closed_closed() -> &dyn Any {
        static Data_Profunctor_Closed_closed: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Closed_closed.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("closed"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}
