pub mod PureScript_Data_Profunctor_Cochoice {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Profunctor_Cochoice_Cochoiceusd_Dict() -> &dyn Any {
        static Data_Profunctor_Cochoice_Cochoiceusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Cochoice_Cochoiceusd_Dict.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |x|
                                                                                  x.clone()))
    }
    pub fn Data_Profunctor_Cochoice_unright() -> &dyn Any {
        static Data_Profunctor_Cochoice_unright: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Cochoice_unright.get_or_init(||
                                                         &Func1::new(move
                                                                         |dict|
                                                                         find(string("unright"),
                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Profunctor_Cochoice_unleft() -> &dyn Any {
        static Data_Profunctor_Cochoice_unleft: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Cochoice_unleft.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("unleft"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}
