pub mod PureScript_Data_CommutativeRing {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_CommutativeRing_ringRecord() -> &dyn Any {
        static Data_CommutativeRing_ringRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_ringRecord.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_ringRecord(),
                                                                                         &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_CommutativeRing_CommutativeRingRecordusd_Dict() -> &dyn Any {
        static Data_CommutativeRing_CommutativeRingRecordusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_CommutativeRingRecordusd_Dict.get_or_init(||
                                                                           &Func1::new(move
                                                                                           |x|
                                                                                           x.clone()))
    }
    pub fn Data_CommutativeRing_CommutativeRingusd_Dict() -> &dyn Any {
        static Data_CommutativeRing_CommutativeRingusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_CommutativeRingusd_Dict.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Data_CommutativeRing_commutativeRingUnit() -> &dyn Any {
        static Data_CommutativeRing_commutativeRingUnit:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_commutativeRingUnit.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                                  &&&add(string("Ring0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Ring::Data_Ring_ringUnit()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_CommutativeRing_commutativeRingRecordNil() -> &dyn Any {
        static Data_CommutativeRing_commutativeRingRecordNil:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_commutativeRingRecordNil.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingRecordusd_Dict(),
                                                                                                       &&&add(string("RingRecord0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Data_Ring::Data_Ring_ringRecordNil()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))
    }
    pub fn Data_CommutativeRing_commutativeRingRecordCons() -> &dyn Any {
        static Data_CommutativeRing_commutativeRingRecordCons:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_commutativeRingRecordCons.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictIsSymbol|
                                                                                       {
                                                                                           let ringRecordCons =
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_ringRecordCons(),
                                                                                                                                                                   dictIsSymbol),
                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                                           &Func1::new({
                                                                                                           let ringRecordCons
                                                                                                               =
                                                                                                               ringRecordCons.clone();
                                                                                                           move
                                                                                                               |usd__unused|
                                                                                                               &Func1::new(move
                                                                                                                               |dictCommutativeRingRecord|
                                                                                                                               {
                                                                                                                                   let ringRecordCons1 =
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&ringRecordCons,
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("RingRecord0"),
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictCommutativeRingRecord)),
                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                   &Func1::new({
                                                                                                                                                   let ringRecordCons1
                                                                                                                                                       =
                                                                                                                                                       ringRecordCons1.clone();
                                                                                                                                                   move
                                                                                                                                                       |dictCommutativeRing|
                                                                                                                                                       {
                                                                                                                                                           let ringRecordCons2 =
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&ringRecordCons1,
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(dictCommutativeRing)),
                                                                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingRecordusd_Dict(),
                                                                                                                                                                                            &&&add(string("RingRecord0"),
                                                                                                                                                                                                   &&Func1::new({
                                                                                                                                                                                                                    let ringRecordCons2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        ringRecordCons2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |usd__unused_1|
                                                                                                                                                                                                                        &ringRecordCons2
                                                                                                                                                                                                                }),
                                                                                                                                                                                                   empty::<string,
                                                                                                                                                                                                           &dyn Any>()))
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                               })
                                                                                                       })
                                                                                       }))
    }
    pub fn Data_CommutativeRing_commutativeRingRecord() -> &dyn Any {
        static Data_CommutativeRing_commutativeRingRecord:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_commutativeRingRecord.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |usd__unused|
                                                                                   &Func1::new(move
                                                                                                   |dictCommutativeRingRecord|
                                                                                                   {
                                                                                                       let ringRecord1 =
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_ringRecord(),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("RingRecord0"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictCommutativeRingRecord)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                                                                        &&&add(string("Ring0"),
                                                                                                                                               &&Func1::new({
                                                                                                                                                                let ringRecord1
                                                                                                                                                                    =
                                                                                                                                                                    ringRecord1.clone();
                                                                                                                                                                move
                                                                                                                                                                    |usd__unused_1|
                                                                                                                                                                    &ringRecord1
                                                                                                                                                            }),
                                                                                                                                               empty::<string,
                                                                                                                                                       &dyn Any>()))
                                                                                                   })))
    }
    pub fn Data_CommutativeRing_commutativeRingProxy() -> &dyn Any {
        static Data_CommutativeRing_commutativeRingProxy:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_commutativeRingProxy.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                                   &&&add(string("Ring0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Ring::Data_Ring_ringProxy()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
    }
    pub fn Data_CommutativeRing_commutativeRingNumber() -> &dyn Any {
        static Data_CommutativeRing_commutativeRingNumber:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_commutativeRingNumber.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                                    &&&add(string("Ring0"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused|
                                                                                                                            &PureScript_Data_Ring::Data_Ring_ringNumber()),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
    }
    pub fn Data_CommutativeRing_commutativeRingInt() -> &dyn Any {
        static Data_CommutativeRing_commutativeRingInt:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_commutativeRingInt.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                                 &&&add(string("Ring0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))
    }
    pub fn Data_CommutativeRing_commutativeRingFn() -> &dyn Any {
        static Data_CommutativeRing_commutativeRingFn:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_CommutativeRing_commutativeRingFn.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictCommutativeRing|
                                                                               {
                                                                                   let ringFn =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_ringFn(),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(dictCommutativeRing)),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                                                    &&&add(string("Ring0"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let ringFn
                                                                                                                                                =
                                                                                                                                                ringFn.clone();
                                                                                                                                            move
                                                                                                                                                |usd__unused|
                                                                                                                                                &ringFn
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))
                                                                               }))
    }
}
