pub mod PureScript_Data_Profunctor_Costrong {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Profunctor_Costrong_Costrongusd_Dict() -> &dyn Any {
        static Data_Profunctor_Costrong_Costrongusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Costrong_Costrongusd_Dict.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |x|
                                                                                  x.clone()))
    }
    pub fn Data_Profunctor_Costrong_unsecond() -> &dyn Any {
        static Data_Profunctor_Costrong_unsecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Costrong_unsecond.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("unsecond"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Profunctor_Costrong_unfirst() -> &dyn Any {
        static Data_Profunctor_Costrong_unfirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Costrong_unfirst.get_or_init(||
                                                         &Func1::new(move
                                                                         |dict|
                                                                         find(string("unfirst"),
                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}
