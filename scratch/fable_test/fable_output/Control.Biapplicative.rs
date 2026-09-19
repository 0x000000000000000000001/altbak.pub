pub mod PureScript_Control_Biapplicative {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_c52ab4fd::PureScript_Control_Biapply;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Biapplicative_Biapplicativeusd_Dict() -> &dyn Any {
        static Control_Biapplicative_Biapplicativeusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapplicative_Biapplicativeusd_Dict.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |x|
                                                                                    x.clone()))
    }
    pub fn Control_Biapplicative_bipure() -> &dyn Any {
        static Control_Biapplicative_bipure: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapplicative_bipure.get_or_init(||
                                                     &Func1::new(move |dict|
                                                                     find(string("bipure"),
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Biapplicative_biapplicativeTuple() -> &dyn Any {
        static Control_Biapplicative_biapplicativeTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapplicative_biapplicativeTuple.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapplicative::Control_Biapplicative_Biapplicativeusd_Dict(),
                                                                                                  &&&add(string("bipure"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__arg1|
                                                                                                                          Func1::new({
                                                                                                                                         let usd__arg1
                                                                                                                                             =
                                                                                                                                             usd__arg1.clone();
                                                                                                                                         move
                                                                                                                                             |usd__arg2|
                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                     usd__arg2.clone()))
                                                                                                                                     })),
                                                                                                         add(string("Biapply0"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused|
                                                                                                                              &PureScript_Control_Biapply::Control_Biapply_biapplyTuple()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))
    }
}
