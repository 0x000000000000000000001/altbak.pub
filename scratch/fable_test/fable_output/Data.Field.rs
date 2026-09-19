pub mod PureScript_Data_Field {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Field_Fieldusd_Dict() -> &dyn Any {
        static Data_Field_Fieldusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Field_Fieldusd_Dict.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Field_field() -> &dyn Any {
        static Data_Field_field: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Field_field.get_or_init(||
                                         &Func1::new(move |dictEuclideanRing|
                                                         &Func1::new({
                                                                         let dictEuclideanRing
                                                                             =
                                                                             dictEuclideanRing.clone();
                                                                         move
                                                                             |dictDivisionRing|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Field::Data_Field_Fieldusd_Dict(),
                                                                                                              &&&add(string("EuclideanRing0"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |usd__unused|
                                                                                                                                      &dictEuclideanRing),
                                                                                                                     add(string("DivisionRing1"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let dictDivisionRing
                                                                                                                                              =
                                                                                                                                              dictDivisionRing.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused_1|
                                                                                                                                              &dictDivisionRing
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                     })))
    }
}
