pub mod PureScript_Data_BooleanAlgebra {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_BooleanAlgebra_heytingAlgebraRecord() -> &dyn Any {
        static Data_BooleanAlgebra_heytingAlgebraRecord:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_heytingAlgebraRecord.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraRecord(),
                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_BooleanAlgebra_BooleanAlgebraRecordusd_Dict() -> &dyn Any {
        static Data_BooleanAlgebra_BooleanAlgebraRecordusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_BooleanAlgebraRecordusd_Dict.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |x|
                                                                                         x.clone()))
    }
    pub fn Data_BooleanAlgebra_BooleanAlgebrausd_Dict() -> &dyn Any {
        static Data_BooleanAlgebra_BooleanAlgebrausd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_BooleanAlgebrausd_Dict.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |x|
                                                                                   x.clone()))
    }
    pub fn Data_BooleanAlgebra_booleanAlgebraUnit() -> &dyn Any {
        static Data_BooleanAlgebra_booleanAlgebraUnit:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_booleanAlgebraUnit.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebrausd_Dict(),
                                                                                                &&&add(string("HeytingAlgebra0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraUnit()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>())))
    }
    pub fn Data_BooleanAlgebra_booleanAlgebraRecordNil() -> &dyn Any {
        static Data_BooleanAlgebra_booleanAlgebraRecordNil:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_booleanAlgebraRecordNil.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebraRecordusd_Dict(),
                                                                                                     &&&add(string("HeytingAlgebraRecord0"),
                                                                                                            &&Func1::new(move
                                                                                                                             |usd__unused|
                                                                                                                             &PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraRecordNil()),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>())))
    }
    pub fn Data_BooleanAlgebra_booleanAlgebraRecordCons() -> &dyn Any {
        static Data_BooleanAlgebra_booleanAlgebraRecordCons:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_booleanAlgebraRecordCons.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictIsSymbol|
                                                                                     {
                                                                                         let heytingAlgebraRecordCons =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraRecordCons(),
                                                                                                                                                                 dictIsSymbol),
                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                                         &Func1::new({
                                                                                                         let heytingAlgebraRecordCons
                                                                                                             =
                                                                                                             heytingAlgebraRecordCons.clone();
                                                                                                         move
                                                                                                             |usd__unused|
                                                                                                             &Func1::new(move
                                                                                                                             |dictBooleanAlgebraRecord|
                                                                                                                             {
                                                                                                                                 let heytingAlgebraRecordCons1 =
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&heytingAlgebraRecordCons,
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebraRecord0"),
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictBooleanAlgebraRecord)),
                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                 &Func1::new({
                                                                                                                                                 let heytingAlgebraRecordCons1
                                                                                                                                                     =
                                                                                                                                                     heytingAlgebraRecordCons1.clone();
                                                                                                                                                 move
                                                                                                                                                     |dictBooleanAlgebra|
                                                                                                                                                     {
                                                                                                                                                         let heytingAlgebraRecordCons2 =
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&heytingAlgebraRecordCons1,
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebra0"),
                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictBooleanAlgebra)),
                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebraRecordusd_Dict(),
                                                                                                                                                                                          &&&add(string("HeytingAlgebraRecord0"),
                                                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                                                  let heytingAlgebraRecordCons2
                                                                                                                                                                                                                      =
                                                                                                                                                                                                                      heytingAlgebraRecordCons2.clone();
                                                                                                                                                                                                                  move
                                                                                                                                                                                                                      |usd__unused_1|
                                                                                                                                                                                                                      &heytingAlgebraRecordCons2
                                                                                                                                                                                                              }),
                                                                                                                                                                                                 empty::<string,
                                                                                                                                                                                                         &dyn Any>()))
                                                                                                                                                     }
                                                                                                                                             })
                                                                                                                             })
                                                                                                     })
                                                                                     }))
    }
    pub fn Data_BooleanAlgebra_booleanAlgebraRecord() -> &dyn Any {
        static Data_BooleanAlgebra_booleanAlgebraRecord:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_booleanAlgebraRecord.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |usd__unused|
                                                                                 &Func1::new(move
                                                                                                 |dictBooleanAlgebraRecord|
                                                                                                 {
                                                                                                     let heytingAlgebraRecord1 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_heytingAlgebraRecord(),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebraRecord0"),
                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictBooleanAlgebraRecord)),
                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebrausd_Dict(),
                                                                                                                                      &&&add(string("HeytingAlgebra0"),
                                                                                                                                             &&Func1::new({
                                                                                                                                                              let heytingAlgebraRecord1
                                                                                                                                                                  =
                                                                                                                                                                  heytingAlgebraRecord1.clone();
                                                                                                                                                              move
                                                                                                                                                                  |usd__unused_1|
                                                                                                                                                                  &heytingAlgebraRecord1
                                                                                                                                                          }),
                                                                                                                                             empty::<string,
                                                                                                                                                     &dyn Any>()))
                                                                                                 })))
    }
    pub fn Data_BooleanAlgebra_booleanAlgebraProxy() -> &dyn Any {
        static Data_BooleanAlgebra_booleanAlgebraProxy:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_booleanAlgebraProxy.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebrausd_Dict(),
                                                                                                 &&&add(string("HeytingAlgebra0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraProxy()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))
    }
    pub fn Data_BooleanAlgebra_booleanAlgebraFn() -> &dyn Any {
        static Data_BooleanAlgebra_booleanAlgebraFn: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_BooleanAlgebra_booleanAlgebraFn.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictBooleanAlgebra|
                                                                             {
                                                                                 let heytingAlgebraFunction =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraFunction(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebra0"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictBooleanAlgebra)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebrausd_Dict(),
                                                                                                                  &&&add(string("HeytingAlgebra0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let heytingAlgebraFunction
                                                                                                                                              =
                                                                                                                                              heytingAlgebraFunction.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &heytingAlgebraFunction
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))
                                                                             }))
    }
    pub fn Data_BooleanAlgebra_booleanAlgebraBoolean() -> &dyn Any {
        static Data_BooleanAlgebra_booleanAlgebraBoolean:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_BooleanAlgebra_booleanAlgebraBoolean.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebrausd_Dict(),
                                                                                                   &&&add(string("HeytingAlgebra0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
    }
}
