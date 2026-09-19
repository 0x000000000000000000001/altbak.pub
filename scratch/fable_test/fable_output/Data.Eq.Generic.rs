pub mod PureScript_Data_Eq_Generic {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Product;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Eq_Generic_GenericEqusd_Dict() -> &dyn Any {
        static Data_Eq_Generic_GenericEqusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Generic_GenericEqusd_Dict.get_or_init(||
                                                          &Func1::new(move |x|
                                                                          x.clone()))
    }
    pub fn Data_Eq_Generic_genericEqNoConstructors() -> &dyn Any {
        static Data_Eq_Generic_genericEqNoConstructors:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Generic_genericEqNoConstructors.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_GenericEqusd_Dict(),
                                                                                                 &&&add(string("genericEq\'"),
                                                                                                        &&Func1::new(move
                                                                                                                         |v|
                                                                                                                         &Func1::new(move
                                                                                                                                         |v1|
                                                                                                                                         &true)),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))
    }
    pub fn Data_Eq_Generic_genericEqNoArguments() -> &dyn Any {
        static Data_Eq_Generic_genericEqNoArguments: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Eq_Generic_genericEqNoArguments.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_GenericEqusd_Dict(),
                                                                                              &&&add(string("genericEq\'"),
                                                                                                     &&Func1::new(move
                                                                                                                      |v|
                                                                                                                      &Func1::new(move
                                                                                                                                      |v1|
                                                                                                                                      &true)),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_Eq_Generic_genericEqArgument() -> &dyn Any {
        static Data_Eq_Generic_genericEqArgument: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Generic_genericEqArgument.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictEq|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_GenericEqusd_Dict(),
                                                                                                           &&&add(string("genericEq\'"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let dictEq
                                                                                                                                       =
                                                                                                                                       dictEq.clone();
                                                                                                                                   move
                                                                                                                                       |v|
                                                                                                                                       &Func1::new({
                                                                                                                                                       let v
                                                                                                                                                           =
                                                                                                                                                           v.clone();
                                                                                                                                                       move
                                                                                                                                                           |v1|
                                                                                                                                                           {
                                                                                                                                                               let matchValue =
                                                                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                               let matchValue_1 =
                                                                                                                                                                   Sharpurs_Prelude::unbox(v1);
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                      &&&dictEq),
                                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                                &&&matchValue_1)
                                                                                                                                                           }
                                                                                                                                                   })
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Eq_Generic_genericEq_prime() -> &dyn Any {
        static Data_Eq_Generic_genericEq_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Generic_genericEq_prime.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("genericEq\'"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Eq_Generic_genericEqConstructor() -> &dyn Any {
        static Data_Eq_Generic_genericEqConstructor: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Eq_Generic_genericEqConstructor.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictGenericEq|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_GenericEqusd_Dict(),
                                                                                                              &&&add(string("genericEq\'"),
                                                                                                                     &&Func1::new({
                                                                                                                                      let dictGenericEq
                                                                                                                                          =
                                                                                                                                          dictGenericEq.clone();
                                                                                                                                      move
                                                                                                                                          |v|
                                                                                                                                          &Func1::new({
                                                                                                                                                          let v
                                                                                                                                                              =
                                                                                                                                                              v.clone();
                                                                                                                                                          move
                                                                                                                                                              |v1|
                                                                                                                                                              {
                                                                                                                                                                  let matchValue =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_genericEq_prime(),
                                                                                                                                                                                                                                                                         &&&dictGenericEq),
                                                                                                                                                                                                                                      &&&matchValue),
                                                                                                                                                                                                   &&&matchValue_1)
                                                                                                                                                              }
                                                                                                                                                      })
                                                                                                                                  }),
                                                                                                                     empty::<string,
                                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Eq_Generic_genericEqProduct() -> &dyn Any {
        static Data_Eq_Generic_genericEqProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Generic_genericEqProduct.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictGenericEq|
                                                                         &Func1::new({
                                                                                         let dictGenericEq
                                                                                             =
                                                                                             dictGenericEq.clone();
                                                                                         move
                                                                                             |dictGenericEq1|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_GenericEqusd_Dict(),
                                                                                                                              &&&add(string("genericEq\'"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let dictGenericEq1
                                                                                                                                                          =
                                                                                                                                                          dictGenericEq1.clone();
                                                                                                                                                      move
                                                                                                                                                          |v|
                                                                                                                                                          &Func1::new({
                                                                                                                                                                          let v
                                                                                                                                                                              =
                                                                                                                                                                              v.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v1|
                                                                                                                                                                              {
                                                                                                                                                                                  let matchValue:
                                                                                                                                                                                          LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                  let matchValue_1:
                                                                                                                                                                                          LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                         &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_genericEq_prime(),
                                                                                                                                                                                                                                                                                                                                                               &&&dictGenericEq),
                                                                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                   Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                            })),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_genericEq_prime(),
                                                                                                                                                                                                                                                                                                                            &&&dictGenericEq1),
                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                      &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                             Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                         }))
                                                                                                                                                                              }
                                                                                                                                                                      })
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>()))
                                                                                     })))
    }
    pub fn Data_Eq_Generic_genericEqSum() -> &dyn Any {
        static Data_Eq_Generic_genericEqSum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Generic_genericEqSum.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictGenericEq|
                                                                     &Func1::new({
                                                                                     let dictGenericEq
                                                                                         =
                                                                                         dictGenericEq.clone();
                                                                                     move
                                                                                         |dictGenericEq1|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_GenericEqusd_Dict(),
                                                                                                                          &&&add(string("genericEq\'"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let dictGenericEq1
                                                                                                                                                      =
                                                                                                                                                      dictGenericEq1.clone();
                                                                                                                                                  move
                                                                                                                                                      |v|
                                                                                                                                                      &Func1::new({
                                                                                                                                                                      let v
                                                                                                                                                                          =
                                                                                                                                                                          v.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |v1|
                                                                                                                                                                          {
                                                                                                                                                                              let matchValue:
                                                                                                                                                                                      LrcPtr<Data_Generic_Rep_Sum> =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                              let matchValue_1:
                                                                                                                                                                                      LrcPtr<Data_Generic_Rep_Sum> =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                              if let Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(matchValue_1_0)
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  if let Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_genericEq_prime(),
                                                                                                                                                                                                                                                                                             &&&dictGenericEq1),
                                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(x)
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                                                                                             }),
                                                                                                                                                                                                                       &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                              {
                                                                                                                                                                                                                              Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(x)
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                              _
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                                          })
                                                                                                                                                                                  } else {
                                                                                                                                                                                      &false
                                                                                                                                                                                  }
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(matchValue_1_0_0)
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_genericEq_prime(),
                                                                                                                                                                                                                                                                                             &&&dictGenericEq),
                                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(x)
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                                                                                             }),
                                                                                                                                                                                                                       &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                              {
                                                                                                                                                                                                                              Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(x)
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                              _
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                                          })
                                                                                                                                                                                  } else {
                                                                                                                                                                                      &false
                                                                                                                                                                                  }
                                                                                                                                                                              }
                                                                                                                                                                          }
                                                                                                                                                                  })
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>()))
                                                                                 })))
    }
    pub fn Data_Eq_Generic_genericEq() -> &dyn Any {
        static Data_Eq_Generic_genericEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Generic_genericEq.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictGeneric|
                                                                  &Func1::new({
                                                                                  let dictGeneric
                                                                                      =
                                                                                      dictGeneric.clone();
                                                                                  move
                                                                                      |dictGenericEq|
                                                                                      &Func1::new({
                                                                                                      let dictGenericEq
                                                                                                          =
                                                                                                          dictGenericEq.clone();
                                                                                                      move
                                                                                                          |x|
                                                                                                          &Func1::new({
                                                                                                                          let x
                                                                                                                              =
                                                                                                                              x.clone();
                                                                                                                          move
                                                                                                                              |y|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq_Generic::Data_Eq_Generic_genericEq_prime(),
                                                                                                                                                                                                                                     &&&dictGenericEq),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                        &&&dictGeneric),
                                                                                                                                                                                                                                     &&&x)),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                     &&&dictGeneric),
                                                                                                                                                                                                  y))
                                                                                                                      })
                                                                                                  })
                                                                              })))
    }
}
