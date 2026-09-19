pub mod PureScript_Data_DivisionRing {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_DivisionRing_DivisionRingusd_Dict() -> &dyn Any {
        static Data_DivisionRing_DivisionRingusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DivisionRing_DivisionRingusd_Dict.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               x.clone()))
    }
    pub fn Data_DivisionRing_recip() -> &dyn Any {
        static Data_DivisionRing_recip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DivisionRing_recip.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("recip"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_DivisionRing_rightDiv() -> &dyn Any {
        static Data_DivisionRing_rightDiv: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DivisionRing_rightDiv.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictDivisionRing|
                                                                   {
                                                                       let Semiring0 =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                   Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictDivisionRing)),
                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                       &Func1::new({
                                                                                       let Semiring0
                                                                                           =
                                                                                           Semiring0.clone();
                                                                                       let dictDivisionRing
                                                                                           =
                                                                                           dictDivisionRing.clone();
                                                                                       move
                                                                                           |a|
                                                                                           &Func1::new({
                                                                                                           let a
                                                                                                               =
                                                                                                               a.clone();
                                                                                                           move
                                                                                                               |b|
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                      &&&Semiring0),
                                                                                                                                                                                   &&&a),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DivisionRing::Data_DivisionRing_recip(),
                                                                                                                                                                                                                      &&&dictDivisionRing),
                                                                                                                                                                                   b))
                                                                                                       })
                                                                                   })
                                                                   }))
    }
    pub fn Data_DivisionRing_leftDiv() -> &dyn Any {
        static Data_DivisionRing_leftDiv: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DivisionRing_leftDiv.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictDivisionRing|
                                                                  {
                                                                      let Semiring0 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictDivisionRing)),
                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                      &Func1::new({
                                                                                      let Semiring0
                                                                                          =
                                                                                          Semiring0.clone();
                                                                                      let dictDivisionRing
                                                                                          =
                                                                                          dictDivisionRing.clone();
                                                                                      move
                                                                                          |a|
                                                                                          &Func1::new({
                                                                                                          let a
                                                                                                              =
                                                                                                              a.clone();
                                                                                                          move
                                                                                                              |b|
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                     &&&Semiring0),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DivisionRing::Data_DivisionRing_recip(),
                                                                                                                                                                                                                                                        &&&dictDivisionRing),
                                                                                                                                                                                                                     b)),
                                                                                                                                               &&&a)
                                                                                                      })
                                                                                  })
                                                                  }))
    }
    pub fn Data_DivisionRing_divisionringNumber() -> &dyn Any {
        static Data_DivisionRing_divisionringNumber: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_DivisionRing_divisionringNumber.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DivisionRing::Data_DivisionRing_DivisionRingusd_Dict(),
                                                                                              &&&add(string("recip"),
                                                                                                     &&Func1::new(move
                                                                                                                      |x|
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                             &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingNumber()),
                                                                                                                                                                                          &&&1.0_f64),
                                                                                                                                                       x)),
                                                                                                     add(string("Ring0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Ring::Data_Ring_ringNumber()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
}
