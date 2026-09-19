pub mod PureScript_Data_Ring_Generic {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_NoArguments;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Product;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Ring_Generic_GenericRingusd_Dict() -> &dyn Any {
        static Data_Ring_Generic_GenericRingusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_Generic_GenericRingusd_Dict.get_or_init(||
                                                              &Func1::new(move
                                                                              |x|
                                                                              x.clone()))
    }
    pub fn Data_Ring_Generic_genericSub_prime() -> &dyn Any {
        static Data_Ring_Generic_genericSub_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_Generic_genericSub_prime.get_or_init(||
                                                           &Func1::new(move
                                                                           |dict|
                                                                           find(string("genericSub\'"),
                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Ring_Generic_genericSub() -> &dyn Any {
        static Data_Ring_Generic_genericSub: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_Generic_genericSub.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictGeneric|
                                                                     {
                                                                         let to_var =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                              dictGeneric);
                                                                         &Func1::new({
                                                                                         let dictGeneric
                                                                                             =
                                                                                             dictGeneric.clone();
                                                                                         let to_var
                                                                                             =
                                                                                             to_var.clone();
                                                                                         move
                                                                                             |dictGenericRing|
                                                                                             &Func1::new({
                                                                                                             let dictGenericRing
                                                                                                                 =
                                                                                                                 dictGenericRing.clone();
                                                                                                             move
                                                                                                                 |x|
                                                                                                                 &Func1::new({
                                                                                                                                 let x
                                                                                                                                     =
                                                                                                                                     x.clone();
                                                                                                                                 move
                                                                                                                                     |y|
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                         &&&to_var),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring_Generic::Data_Ring_Generic_genericSub_prime(),
                                                                                                                                                                                                                                                                               &&&dictGenericRing),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                                                                  &&&dictGeneric),
                                                                                                                                                                                                                                                                               &&&x)),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                               &&&dictGeneric),
                                                                                                                                                                                                                                            y)))
                                                                                                                             })
                                                                                                         })
                                                                                     })
                                                                     }))
    }
    pub fn Data_Ring_Generic_genericRingProduct() -> &dyn Any {
        static Data_Ring_Generic_genericRingProduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Ring_Generic_genericRingProduct.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictGenericRing|
                                                                             &Func1::new({
                                                                                             let dictGenericRing
                                                                                                 =
                                                                                                 dictGenericRing.clone();
                                                                                             move
                                                                                                 |dictGenericRing1|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring_Generic::Data_Ring_Generic_GenericRingusd_Dict(),
                                                                                                                                  &&&add(string("genericSub\'"),
                                                                                                                                         &&Func1::new({
                                                                                                                                                          let dictGenericRing1
                                                                                                                                                              =
                                                                                                                                                              dictGenericRing1.clone();
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
                                                                                                                                                                                      &LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring_Generic::Data_Ring_Generic_genericSub_prime(),
                                                                                                                                                                                                                                                                                                                                                                     &&&dictGenericRing),
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
                                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring_Generic::Data_Ring_Generic_genericSub_prime(),
                                                                                                                                                                                                                                                                                                                                                                     &&&dictGenericRing1),
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
                                                                                                                                                                                                                                                                                                  })))
                                                                                                                                                                                  }
                                                                                                                                                                          })
                                                                                                                                                      }),
                                                                                                                                         empty::<string,
                                                                                                                                                 &dyn Any>()))
                                                                                         })))
    }
    pub fn Data_Ring_Generic_genericRingNoArguments() -> &dyn Any {
        static Data_Ring_Generic_genericRingNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_Generic_genericRingNoArguments.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring_Generic::Data_Ring_Generic_GenericRingusd_Dict(),
                                                                                                  &&&add(string("genericSub\'"),
                                                                                                         &&Func1::new(move
                                                                                                                          |v|
                                                                                                                          &Func1::new(move
                                                                                                                                          |v1|
                                                                                                                                          &LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor))),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_Ring_Generic_genericRingConstructor() -> &dyn Any {
        static Data_Ring_Generic_genericRingConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_Generic_genericRingConstructor.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictGenericRing|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring_Generic::Data_Ring_Generic_GenericRingusd_Dict(),
                                                                                                                  &&&add(string("genericSub\'"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let dictGenericRing
                                                                                                                                              =
                                                                                                                                              dictGenericRing.clone();
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
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring_Generic::Data_Ring_Generic_genericSub_prime(),
                                                                                                                                                                                                                                                                                                                &&&dictGenericRing),
                                                                                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                                                                                          &&&matchValue_1))
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Ring_Generic_genericRingArgument() -> &dyn Any {
        static Data_Ring_Generic_genericRingArgument:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_Generic_genericRingArgument.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictRing|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring_Generic::Data_Ring_Generic_GenericRingusd_Dict(),
                                                                                                               &&&add(string("genericSub\'"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let dictRing
                                                                                                                                           =
                                                                                                                                           dictRing.clone();
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
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                             &&&dictRing),
                                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                                       &&&matchValue_1))
                                                                                                                                                               }
                                                                                                                                                       })
                                                                                                                                   }),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
    }
}
